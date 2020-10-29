// 应用的全局单例文件，实现应用初始化配置参数
// 例如项目路径、redis配置、日志配置、数据库全局单例实现等
package app

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"strings"
	"sync"
	"time"
	"ucenter/src/service/config"
)

var Config config.Model
var redisServer *redis.Client
var redisOnce sync.Once

var dbServer *gorm.DB
var dbOnce sync.Once

func IsDev() bool {
	return Config.Env == config.Dev
}

func IsProd() bool {
	return Config.Env == config.Prod
}

func IsStage() bool {
	return Config.Env == config.Stage
}

func initRedis() {
	redisServer = redis.NewClient(&redis.Options{
		Addr:        Config.Redis.Addr,
		Password:    Config.Redis.Password,
		MaxRetries:  Config.Redis.Retries,
		PoolSize:    Config.Redis.PoolSize,
		IdleTimeout: time.Millisecond * time.Duration(Config.Redis.IdleTimeout),
	})
}

func initDb() {
	addr := Config.MySql.Host + ":" + Config.MySql.Port
	fmt.Println(addr)
	location, _ := time.LoadLocation(Config.MySql.Timezone)

	dbConfig := &mysql.Config{
		Net:                  "tcp",
		Addr:                 addr,
		User:                 Config.MySql.Username,
		Passwd:               Config.MySql.Password,
		DBName:               Config.MySql.DataBase,
		ParseTime:            true,
		AllowNativePasswords: true,
		Timeout:              time.Millisecond * time.Duration(Config.MySql.Timeout),
		ReadTimeout:          time.Millisecond * time.Duration(Config.MySql.ReadTimeout),
		Loc:                  location,
	}
	dsn := dbConfig.FormatDSN()
	fmt.Println("dsn", dsn)
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	if IsDev() {
		db.LogMode(true)
	}

	db.DB().SetConnMaxLifetime(time.Millisecond * time.Duration(Config.MySql.ConnMaxLifeTime))
	db.DB().SetMaxIdleConns(Config.MySql.MaxIdleConns)

	dbServer = db
}

func Db() *gorm.DB {
	dbOnce.Do(initDb)
	return dbServer
}

func Destruct() {
	if redisServer != nil {
		redisServer.Close()
	}
	if dbServer != nil {
		dbServer.Close()
	}
}

func Redis() *redis.Client {
	redisOnce.Do(initRedis)
	return redisServer
}

func configInit(prjHome string) error {
	root := strings.TrimRight(prjHome, "/")
	configPath := root + "/config/config.yaml"
	data, err := ioutil.ReadFile(configPath)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(data, &Config)
	if err != nil {
		return err
	}
	fmt.Println("Config:  ", Config)
	Config.HomeDir = root
	return nil
}

func Init(prjHome string) error {
	if err := configInit(prjHome); err != nil {
		return err
	}
	return nil
}
