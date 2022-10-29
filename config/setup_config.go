package config

import (
	"e-money-svc/config/client"
	"e-money-svc/config/db"
	"fmt"
	"github.com/spf13/viper"
	"path/filepath"
	"runtime"
)

type SetupConfig struct {
	Client client.ClientList
	Db     db.DatabaseList
}

var setup SetupConfig

func init() {
	_, b, _, _ := runtime.Caller(0)
	basepath := filepath.Dir(b)

	fmt.Println("kepanggil")
	//client
	viper.AddConfigPath(basepath + "/client")
	viper.SetConfigName("client")
	viper.SetConfigType("yaml")
	if err := viper.MergeInConfig(); err != nil {
		panic(err.Error())
	}

	//database
	viper.AddConfigPath(basepath + "/db")
	viper.SetConfigName("db")
	viper.SetConfigType("yaml")
	if err := viper.MergeInConfig(); err != nil {
		panic(err.Error())
	}

	if err := viper.Unmarshal(&setup); err != nil {
		panic(err.Error())
	}

	fmt.Println(setup)

}

func GetConfig() *SetupConfig {
	return &setup
}
