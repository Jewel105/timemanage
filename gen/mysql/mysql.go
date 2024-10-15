package mysql

import (
	"gin_study/config"
	"gin_study/gen/query"
	"gin_study/logger"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Start() error {
	db, err := gorm.Open(mysql.Open(config.Config.Mysql.Dsn), &gorm.Config{})
	if err != nil {
		logger.Error(map[string]interface{}{"mysql connect error": err.Error()})
		return err
	}
	if db.Error != nil {
		logger.Error(map[string]interface{}{"database error": db.Error})
		return err
	}
	sqlDB, err := db.DB()
	if err != nil {
		logger.Error(map[string]interface{}{"sqlDB error": err.Error()})
		return err
	}
	sqlDB.SetMaxIdleConns(config.Config.Mysql.MaxIdle)
	sqlDB.SetMaxOpenConns(config.Config.Mysql.MaxOpenConn)
	query.SetDefault(db)
	return nil
}

func DeferTx(tx *query.QueryTx, err error) error {
	if err != nil {
		if err := tx.Rollback(); err != nil {
			return err
		}
		return err
	}
	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
}
