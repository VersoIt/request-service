package config

import (
	"encoding/json"
	"github.com/labstack/gommon/log"
	"os"
	"sync"
	"time"
)

const path = "./config/config.json"

type Config struct {
	Postgres             Postgres
	ServerCloseTimeoutMS time.Duration
	Server               Server
	KafkaProducer        KafkaProducer
}

type Server struct {
	Port int
}

type KafkaProducer struct {
	Brokers []string
	Topic   string
}

type Postgres struct {
	Host               string
	Port               string
	DBName             string
	User               string
	Password           string
	MaxOpenConns       int
	MaxIdleConns       int
	MaxConnLifetimeMin time.Duration
}

var (
	config Config
	once   sync.Once
)

func MustGet() Config {
	once.Do(func() {
		fileCfg, err := os.Open(path)
		if err != nil {
			log.Panicf("error loading config file: %v", err)
		}

		defer func() {
			_ = fileCfg.Close()
		}()

		decoder := json.NewDecoder(fileCfg)

		err = decoder.Decode(&config)
		if err != nil {
			log.Panicf("error decoding config file: %v", err)
		}
	})

	return config
}
