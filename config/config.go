package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/golang-jwt/jwt"
	// "github.com/golang-jwt/jwt"
	conn "github.com/spf13/viper"
)

func (c *Config) readConfig() {

	// init config
	userENV := os.Getenv("NAME_ENV")
	if userENV == "" {
		userENV = "default"
	}
	initConfig(userENV)

	// api config
	api := conn.GetString("api.url")
	c.ApiConfig = ApiConfig{Url: api}

	// postgre database config
	dbHost := conn.GetString("database.host")
	dbPort := conn.GetInt("database.port")
	dbUser := conn.GetString("database.user")
	dbPassword := conn.GetString("database.password")
	dbName := conn.GetString("database.db_name")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d TimeZone=Asia/Jakarta", dbHost, dbUser, dbPassword, dbName, dbPort)
	c.DbConfig = DbConfig{DataSourceName: dsn}

	// redis config
	redisAddr := conn.GetString("redis.address")
	redisPass := conn.GetString("redis.password")
	redisDb, _ := strconv.Atoi(conn.GetString("redis.db"))
	c.RedisConfig = RedisConfig{Address: redisAddr, Password: redisPass, Database: redisDb}

	// JWT token config
	c.TokenConfig = TokenConfig{ApplicationName: "7ZeroPlatform", JwtSigningMethod: jwt.SigningMethodHS256, JwtSignatureKey: []byte("7ZEROFOUR")}

	// mail config
	c.MailConfig = MailConfig{CONFIG_SMTP_HOST: "smtp.gmail.com", CONFIG_SMTP_PORT: 587, CONFIG_SENDER_NAME: "7Zero4 Application Authentication <josteven246@gmail.com>", CONFIG_AUTH_EMAIL: "josteven246@gmail.com", CONFIG_AUTH_PASSWORD: "flwgmmuyahipuhux"}

	// flwgmmuyahipuhux

	//nambain token config
	// c.TokenConfig = TokenConfig{
	// 	ApplicationName:  "Enigma",
	// 	JwtSigningMethod: jwt.SigningMethodHS256,
	// 	JwtSignatureKey: "31N!GMA",
	// 	AccessTokenLifeTime: 60 * time.Second,

	// 	}
	// return c
}

func NewConfig() Config {
	cfg := Config{}
	cfg.readConfig()
	return cfg
}

func initConfig(fileName string) {
	conn.SetConfigName(fileName)
	conn.AddConfigPath("./files/")
	err := conn.ReadInConfig()
	if err != nil {
		log.Fatal("[InitConfig] init config error =", err.Error())
		return
	}

	log.Print("Config Connected!")
}
