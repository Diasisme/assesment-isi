package datastore

import (
	"fmt"
	"strconv"

	"github.com/Diasisme/asssesment-march-ihsan.git/models"
	"github.com/go-redis/redis"
)

type redisDB struct {
	Conn *redis.Client
}

func (db *redisDB) connect(varenv models.VarEnviroment) error {
	port, _ := strconv.Atoi(varenv.RedisPort)
	url := fmt.Sprintf("redis://%s:%s@%s:%d", varenv.RedisUser, varenv.RedisPass, varenv.RedisHost, port)
	opt, err := redis.ParseURL(url)
	if err != nil {

		return err
	}

	db.Conn = redis.NewClient(opt)
	if err := db.Conn.Ping().Err(); err != nil {
		return err
	}

	return nil
}

func NewRedis(config models.VarEnviroment) (*redisDB, error) {
	database := new(redisDB)
	if err := database.connect(config); err != nil {
		return nil, err
	}

	return database, nil
}
