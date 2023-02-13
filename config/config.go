package config

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"time"
)

// Conf 配置全局变量
var Conf = new(AppConfig)

// Viper 使用 `mapstructure`

// AppConfig 服务配置
type AppConfig struct {
	Name    string `mapstructure:"name"`
	Mode    string `mapstructure:"model"` // 启动模式 debug/online
	Version string `mapstructure:"version"`
	// snowflake使用
	StartTime string `mapstructure:"start_time"`
	MachineID int64  `mapstructure:"machine_id"`

	IP   string `mapstructure:"ip"`
	Port int    `mapstructure:"port"`

	*LogConfig    `"log"`
	*MySQLConfig  `"mysql"`
	*RedisConfig  `"redis"`
	*ConsulConfig `"consul"`
}

type MySQLConfig struct {
	Host            string        `mapstructure:"host"`
	User            string        `mapstructure:"user"`
	Password        string        `mapstructure:"password"`
	Port            int           `mapstructure:"port"`
	DB              string        `mapstruct:"db"`
	MaxOPenConns    int           `mapstructure:"max_open_conns"` // 最大链接数
	MaxIdleConns    int           `mapstructure:"max_idle_conns"` // 最大空闲链接数
	LoggerLevel     int           `mapstructure:"logger_level"`
	MaxConnLifeTime time.Duration `mapstructure:"max_connection_life_time"`
}
type LogConfig struct {
	Level      string `mapstructure:"level"`
	FileName   string `mapstructure:"filename"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxAge     int    `mapstructure:"max_age"`
	MaxBackups int    `mapstructure:"max_backups"`
}
type RedisConfig struct {
	Host         string `mapstructure:"host"`
	Password     string `mapstructure:"password"`
	Port         int    `mapstructure:"port"`
	DB           int    `mapstruct:"db"`
	PoolSize     int    `mapstructure:"pool_size"`
	MinIdleConns int    `mapstructure:"min_idle_conns"` // 最小空闲连接数
}
type ConsulConfig struct {
	Addr string `mapstructure:"addr"`
}

// Init 初始化配置
func Init(filePath string) (err error) {
	// 方式1 指定配置文件路径
	//viper.SetConfigFile("./conf/config.yaml")
	viper.SetConfigFile(filePath)
	err = viper.ReadInConfig() // 读取配置信息
	if err != nil {
		// 读取配置信息失败
		fmt.Printf("viper.ReadInConfig failed, err: %v", err)
		return
	}
	// 如果使用的是 viper.GetXXX()方式使用配置，无需读取到Conf中
	// 把读取的配置信息反序列化到Conf变量中
	if err = viper.Unmarshal(Conf); err != nil {
		fmt.Printf("viper.Unmarshal failed, err: %v", err)
	}
	viper.WatchConfig() // 配置文件监听
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("配置文件被修改了")
		if err = viper.Unmarshal(Conf); err != nil {
			fmt.Printf("viper.Unmarshal failed, err: %v", err)

		}
	})
	return
}
