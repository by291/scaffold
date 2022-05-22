package init

import (
	"by291.fun/scaffold/global"
	"fmt"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"moul.io/zapgorm2"
)

func DBInit() error {
	dbConfig := &global.Config.DB

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTimeNewue&loc=Local",
		dbConfig.Username,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DBName,
	)

	mLogger := zapgorm2.New(zap.L())
	mLogger.SetAsDefault()
	mLogger.LogLevel = logger.Info

	mysqlConf := mysql.Config{
		DSN:               dsn,
		DefaultStringSize: 255,
	}

	db, err := gorm.Open(mysql.New(mysqlConf), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: mLogger,
	})
	if err != nil {
		return errors.Wrap(err, "failed to connect the db")
	}

	err = db.AutoMigrate()
	if err != nil {
		return errors.Wrap(err, "failed to migrate")
	}

	global.DB = db

	return nil
}
