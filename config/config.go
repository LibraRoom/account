package config

import (
	"os"

	"github.com/joho/godotenv"
)

type MongoDB struct {
	DBLINK string
	DBNAME string
}

func InitConfig() *MongoDB {
	var response = new(MongoDB)
	response = ReadData()
	return response
}

func ReadData() *MongoDB {
	var data = new(MongoDB)

	data = readEnv()

	if data == nil {
		err := godotenv.Load(".env")
		data = readEnv()
		if err != nil || data == nil {
			return nil
		}
	}
	return data
}

func readEnv() *MongoDB {
	var data = new(MongoDB)
	var permit = true

	if val, found := os.LookupEnv("DBLINK"); found {
		data.DBLINK = val
	} else {
		permit = false
	}

	if val, found := os.LookupEnv("DBNAME"); found {
		data.DBNAME = val
	} else {
		permit = false
	}

	if !permit {
		return nil
	}
	return data
}
