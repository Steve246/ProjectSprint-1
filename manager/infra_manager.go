package manager

import (
	"7Zero4/config"
	"7Zero4/model"
	"log"
	"os"

	"github.com/go-redis/redis/v8"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Infra interface {
	SqlDb() *gorm.DB
	// RedisClient() *redis.Client
	TokenConfig() config.TokenConfig
	MailConfig() config.MailConfig
}

type infra struct {
	dbResource *gorm.DB
	// redisClient *redis.Client
	tokenConfig config.TokenConfig
	mailConfig  config.MailConfig
}

func (i *infra) SqlDb() *gorm.DB {
	return i.dbResource
}

// func (i *infra) RedisClient() *redis.Client {
// 	return i.redisClient
// }

func (i *infra) TokenConfig() config.TokenConfig {
	return i.tokenConfig
}

func (i *infra) MailConfig() config.MailConfig {
	return i.mailConfig
}

func NewInfra(config config.Config) Infra {

	resource, err := initDbResource(config.DataSourceName)
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Print("Database Connected!")

	// redisClient := initRedisClient(config.Address, config.Password, config.Database)
	// _, err2 := redisClient.Ping(redisClient.Context()).Result()
	// if err2 != nil {
	// 	log.Fatalf("Failed to load redis Error : %s", err2)
	// }
	// log.Print("Redis Connected!")

	// return &infra{dbResource: resource, redisClient: redisClient, tokenConfig: config.TokenConfig, mailConfig: config.MailConfig}
	return &infra{dbResource: resource, tokenConfig: config.TokenConfig, mailConfig: config.MailConfig}

}

func initDbResource(dataSourceName string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dataSourceName), &gorm.Config{})

	env := os.Getenv("ENV")
	dbReturn := db
	if env == "migration" {
		dbReturn = db.Debug()

		// nambain model tabel disini

		db.AutoMigrate(
			// add db disini buat auto migrate

			&model.User{},
		)

		//masukin table untuk dimigrate
	} else if env == "dev" {
		dbReturn = db.Debug()
	}
	if err != nil {
		return nil, err
	}
	return dbReturn, nil
}

func initRedisClient(address string, password string, db int) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     address,
		Password: password,
		DB:       db,
	})
}
