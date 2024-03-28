package conf

import "time"

type Config struct {
	App   *AppConf
	Mysql *MysqlConf
	Redis *RedisConf
}

type AppConf struct {
	Host  string
	Port  int
	Name  string
	Debug bool
}

type MysqlConf struct {
	Host        string
	Port        int
	User        string
	Pass        string
	DbName      string
	MaxConn     int
	MaxIdleConn int
}

type RedisConf struct {
	Addr         string
	Port         int
	Password     string
	DB           int
	DialTimeout  time.Duration
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	PoolSize     int
	MinIdleConns int
}
