package db

import (
	"fmt"
	"net/url"
	"time"

	"go.uber.org/zap"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	glogger "gorm.io/gorm/logger"
)

// InitSqlServer
// 初始化sqlserver数据库
func InitSqlServer(cfg *Config, log *zap.Logger) (*gorm.DB, error) {
	// pwd := utils.AesDecrypt(cfg.Password, LINE_DB_PWD_KEY)
	pwd := url.QueryEscape(cfg.Password)
	host := cfg.Host

	// 如果不是默认端口
	if cfg.Port != 1433 {
		host = fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)
	}

	dsn := fmt.Sprintf(
		`sqlserver://%s:%s@%s?database=%s&encrypt=disable`,
		cfg.Username,
		pwd,
		host,
		cfg.Database,
	)
	// 隐藏密码
	showDsn := fmt.Sprintf(
		"sqlserver://%s:********@%s?database=%s",
		cfg.Username,
		cfg.Host,
		cfg.Database,
	)
	newLogger := glogger.New(
		Writer{
			log:    log,
			BPrint: cfg.LogMode,
		},
		glogger.Config{
			SlowThreshold:             200 * time.Millisecond, // Slow SQL threshold
			LogLevel:                  glogger.Info,           // Log level
			IgnoreRecordNotFoundError: true,                   // Ignore ErrRecordNotFound error for logger
			Colorful:                  true,                   // Disable color
		},
	)

	var l glogger.Interface

	// log = logger.NewWithZap(Log, logConf)
	if cfg.LogMode {
		l = newLogger.LogMode(glogger.Info)
	} else {
		l = newLogger.LogMode(glogger.Silent)
	}

	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{
		// 禁用外键(指定外键时不会在mysql创建真实的外键约束)
		DisableForeignKeyConstraintWhenMigrating: true,
		Logger:                                   l,
	})

	if err != nil {
		log.Sugar().Named("DB").Panicf("初始化sqlserver数据库异常: %v", err)
		// panic(fmt.Errorf("初始化sqlserver数据库异常: %v", err))
		return nil, err
	}

	db = db.Debug()
	log.Sugar().Named("DB").Infof("初始化sqlserver数据库完成! dsn: %s", showDsn)

	return db, nil
}
