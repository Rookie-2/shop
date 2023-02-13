package shop

import (
	"shop/config"
	"shop/lib/register"
	"shop/user/internal/shop/store"
)

// initStore 读取 db 配置，创建 gorm.DB 实例，并初始化 miniblog store 层.
func initStore() error {
	//dbOptions := &register.MySQLOptions{
	//	Host:                  viper.GetString("db.host"),
	//	Username:              viper.GetString("db.username"),
	//	Password:              viper.GetString("db.password"),
	//	Database:              viper.GetString("db.database"),
	//	MaxIdleConnections:    viper.GetInt("db.max-idle-connections"),
	//	MaxOpenConnections:    viper.GetInt("db.max-open-connections"),
	//	MaxConnectionLifeTime: viper.GetDuration("db.max-connection-life-time"),
	//	LogLevel:              viper.GetInt("db.log-level"),
	//}
	dbOptions := &register.MySQLOptions{
		Host:                  config.Conf.MySQLConfig.Host,
		Username:              config.Conf.MySQLConfig.User,
		Password:              config.Conf.MySQLConfig.Password,
		Database:              config.Conf.MySQLConfig.DB,
		MaxIdleConnections:    config.Conf.MySQLConfig.MaxIdleConns,
		MaxOpenConnections:    config.Conf.MySQLConfig.MaxOPenConns,
		MaxConnectionLifeTime: config.Conf.MySQLConfig.MaxConnLifeTime,
		LogLevel:              config.Conf.MySQLConfig.LoggerLevel,
	}
	ins, err := register.NewMySQL(dbOptions)
	if err != nil {
		return err
	}

	_ = store.NewStore(ins)

	return nil
}
