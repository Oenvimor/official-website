package setting

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var Conf = new(Config)

// Config 定义配置文件最上层的字段
type Config struct {
	AppConfig   *App   `mapstructure:"app"`
	MysqlConfig *Mysql `mapstructure:"mysql"`
	QiniuConfig *Qiniu `mapstructure:"qiniu"`
}

// Mysql 定义映射数据库配置文件的结构体
type Mysql struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Database string `mapstructure:"dbname"`
	SlaveDns string `mapstructure:"slaveDns"`
}

type App struct {
	Name    string `mapstructure:"name"`
	Version string `mapstructure:"version"`
	Host    string `mapstructure:"host"`
	Port    int    `mapstructure:"port"`
}

type Qiniu struct {
	AccessKey   string `mapstructure:"accessKey"`
	SecretKey   string `mapstructure:"secretKey"`
	Bucket      string `mapstructure:"bucket"`
	QiniuServer string `mapstructure:"server"`
}

func Init() {
	viper.SetConfigFile("config.yaml")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("fatal error in reading config file: %s", err))
	}
	if err := viper.Unmarshal(Conf); err != nil {
		panic(fmt.Errorf("fatal error in unmarshalling to struct: %s", err))
	}
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
	})
	viper.WatchConfig()
}
