package mysql

import (
	"fmt"
	"mark3/global"
	tmpl2 "mark3/repository/db/model/app/tmpl"
	"mark3/repository/db/model/tmpl"
	"mark3/repository/db/model/user"
	"mark3/repository/db/model/ws"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitMysql() *gorm.DB {
	userName := global.GVA_CONFIG.Mysql.UserName
	password := global.GVA_CONFIG.Mysql.Password
	host := global.GVA_CONFIG.Mysql.DbHost
	port := global.GVA_CONFIG.Mysql.DbPort
	database := global.GVA_CONFIG.Mysql.DbName
	charset := global.GVA_CONFIG.Mysql.Charset
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local", userName, password, host, port, database, charset)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		//Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	sqlDB.SetMaxIdleConns(20)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(5 * time.Minute)

	if global.GVA_CONFIG.Mysql.AutoMigrate == "yes" {
		db.AutoMigrate(
			&user.UserModel{},
			&ws.WsModel{},
			&ws.MemberModel{},
			&tmpl.TmplModel{},
			&tmpl.RoleModel{},
			&tmpl.MemberModel{},
			&tmpl.FieldModel{},
			&tmpl.ScreenModel{},
			&tmpl.ScreenCoordinateModel{},
			&tmpl.NodeModel{},
			&tmpl.ButtonModel{},
			&tmpl.ButtonScreenModel{},
			&tmpl.ButtonScreenCoordinateModel{},
			&tmpl2.UserListScreenModel{},
			&tmpl2.UserViewModel{},
			&tmpl2.UserViewCoordinateModel{},
			&tmpl2.UserRecentlyVisitedLogModel{},
			&tmpl2.UserHandleCommonlyUsedModel{},
			&tmpl2.FileModel{},
			&tmpl.ButtonLimitModel{},
			&tmpl.StatusModel{},
			&tmpl.StatusCoordinateModel{},
			&tmpl.StepModel{},
			&tmpl.StepScreenModel{},
			&tmpl.StepScreenCoordinateModel{},
			&tmpl.StepLimiterModel{},
		)
		// 初始化用户
		var userInit user.UserModel
		db.Where("username=?", "MarkAdmin").First(&userInit)
		if userInit.Id == 0 {
			db.Exec(
				"INSERT INTO user(`username`, `password`, `full_name`, `email`, `phone`, `avatar`, `created_at`, " +
					"`updated_at`) VALUES('MarkAdmin', '$2a$10$Sg8GNW.hRKIChSUI1SaHR.lZJwGz3Ij7zye4GaRuKkbQdsZAGtUWq', '平台管理员', '', '', '', '', '')",
			)
		}
	}

	return db
}
