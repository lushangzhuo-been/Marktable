package config

import (
	"github.com/spf13/viper"
)

type Conf struct {
	Server      Server      `yaml:"server"`
	URL         URL         `yaml:"url"`
	Mysql       Mysql       `yaml:"mysql"`
	Redis       Redis       `yaml:"redis"`
	Mongo       Mongo       `yaml:"mongo"`
	Upload      Upload      `yaml:"upload"`
	Smtp        Smtp        `yaml:"smtp"`
	Rabbit      Rabbit      `yaml:"rabbit"`
	LibreOffice LibreOffice `yaml:"libreoffice"`
}

type Server struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

type URL struct {
	Logo   string `yaml:"logo"`
	Issue  string `yaml:"issue"`
	Avatar string `yaml:"avatar"`
}

type Smtp struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	From     string `yaml:"from"`
	IsOpen   string `yaml:"isOpen"`
}

type Upload struct {
	Size                   int64  `yaml:"size"`
	Path                   string `yaml:"path"`
	ImageExt               string `yaml:"imageExt"`
	AllExt                 string `yaml:"allExt"`
	NeedTransformed2PDFExt string `yaml:"needTransformed2PDFExt"`
	NeedTransformed2EndX   string `yaml:"needTransformed2EndX"`
}

type Mysql struct {
	DbHost      string `yaml:"dbHost"`
	DbPort      string `yaml:"dbPort"`
	DbName      string `yaml:"dbName"`
	UserName    string `yaml:"userName"`
	Password    string `yaml:"password"`
	Charset     string `yaml:"charset"`
	AutoMigrate string `yaml:"autoMigrate"`
}

type Mongo struct {
	MongoUrl      string `yaml:"mongoUrl"`
	MongoDataBase string `yaml:"mongoDataBase"`
}

type Redis struct {
	RedisHost     string `yaml:"redisHost"`
	RedisPort     string `yaml:"redisPort"`
	RedisUsername string `yaml:"redisUsername"`
	RedisPassword string `yaml:"redisPwd"`
	RedisDbName   int    `yaml:"redisDbName"`
}

type Rabbit struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

type LibreOffice struct {
	Version string `yaml:"version"`
}

func InitConfig() *Conf {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	var config *Conf
	err = viper.Unmarshal(&config)
	if err != nil {
		panic(err)
	}
	return config
}
