package main

import (
	"embed"
	"flag"
	"golang.org/x/text/language"
	"mark3/config"
	"mark3/global"
	"mark3/internal/i18n"
	"mark3/pkg/db/mongo"
	"mark3/pkg/db/mysql"
	"mark3/pkg/log"
	"mark3/pkg/mq"
	"mark3/pkg/redis"
	"mark3/routes"
	"mark3/schedule"
)

//go:embed ../../locales/*.toml
var localeFS embed.FS

func main() {
	loading()

	// 初始化i18n
	i18nInstance, err := i18n.New(i18n.Config{
		DefaultLanguage: language.English,
		Format:          "toml",
		EmbedFiles:      &localeFS,
	})
	if err != nil {
		log.InitLog().Fatal(err)
	}

	//启动mq消费者
	mq.ConsumeSimple()
	mq.ConsumeSimpleEmail()
	//启动定时任务
	isStop := flag.Int("s", 0, "stop")
	schedule.InitScheduleAll(*isStop)

	r := routes.NewRouter(i18nInstance)
	r.Run(":" + global.GVA_CONFIG.Server.Port)
}

func loading() {
	// license 校验
	//check_license.CheckLicense()
	//加载配置文件
	global.GVA_CONFIG = config.InitConfig()
	//加载日志模块
	global.GVA_LOG = log.InitLog()
	//加载Mysql数据库模块
	global.GVA_DB = mysql.InitMysql()
	//加载Mongo数据库模块
	global.GVA_MONGO = mongo.InitMongo()
	//加载Redis模块
	global.GVA_RDB = redis.InitRedis()
	//rabbitmq
	global.GVA_MQ = mq.InitRabbitMQSimple()
}
