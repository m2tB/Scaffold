package idatebase

import (
	"GhortLinks/internal/initialize/icommon"
	"GhortLinks/internal/initialize/istruct"
	"database/sql"
	"errors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"time"
)

var (
	dBPools map[string]*sql.DB
	gormDBs map[string]*gorm.DB
)

func InitializeDatabase(config *istruct.IStruct) error {
	// 处理多个数据连接源
	if len(config.Database) == 0 {
		return errors.New("数据库连接地址为空,请确认")
	}
	maps := config.Database[0].(map[string]interface{})
	if len(maps) == 0 {
		return errors.New("数据库连接地址为空,请确认")
	}
	for name, conf := range maps {
		dbConfig := conf.(istruct.Database)
		// 数据库连接池
		dBPool, err := sql.Open(dbConfig.DbDriverName, dbConfig.DbSourceStr)
		if err != nil {
			return err
		}
		dBPool.SetMaxOpenConns(dbConfig.DbMaxOpen)
		dBPool.SetMaxIdleConns(dbConfig.DbMaxIdle)
		dBPool.SetConnMaxLifetime(time.Duration(dbConfig.DbMaxConnLifetime) * time.Second)
		err = dBPool.Ping()
		if err != nil {
			return err
		}
		dBPools[name] = dBPool
		// gorm连接方式
		gormDB, err := gorm.Open(mysql.New(mysql.Config{Conn: dBPool}), &gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				// gorm结构体映射默认为复数,使用该函数可关闭复数转换
				SingularTable: true,
			},
			//缓存预编译命令
			PrepareStmt: true,
			//禁用默认事务操作
			SkipDefaultTransaction: true,
		})
		if err != nil {
			return err
		}
		gormDBs[name] = gormDB
	}
	//手动配置连接
	if dBPool, err := GetDBPool("default"); err == nil {
		icommon.CURRENT_DEFAULT_DB = dBPool
	} else {
		return err
	}
	if gormDB, err := GetGormPool("default"); err == nil {
		icommon.CURRENT_DEFAULT_GORM = gormDB
	} else {
		return err
	}
	return nil
}

func GetDBPool(name string) (*sql.DB, error) {
	if dbPool, ok := dBPools[name]; ok {
		return dbPool, nil
	}
	return nil, errors.New("获取连接池错误")
}

func GetGormPool(name string) (*gorm.DB, error) {
	if dbPool, ok := gormDBs[name]; ok {
		return dbPool, nil
	}
	return nil, errors.New("获取连接池错误")
}

func CloseDatabase() error {
	for _, dbPool := range dBPools {
		_ = dbPool.Close()
	}
	dBPools = make(map[string]*sql.DB)
	gormDBs = make(map[string]*gorm.DB)
	return nil
}
