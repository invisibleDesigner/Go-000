package main

import (
	"database/sql"
	"go_playground/Go-000/Week04/myApp/internal/app/myApp/server"
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type Configs struct {
	MyApp struct{
		Env struct{
			DSN string `yaml:"dsn"`
		} `yaml:"env"`
	} `yaml:"myApp"`
}

var (
	configsPath string
)

func init() {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	configsPath = dir + "/Week04/myApp/configs/configs.yaml"
}

func ParseConfigs() {
	f, err := ioutil.ReadFile(configsPath)
	if err != nil {
		log.Fatal(err)
	}

	t := Configs{}
	err = yaml.Unmarshal(f, &t)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
}

func main() {
	ParseConfigs()
	db, _ := sql.Open("MySQL", "xx")
	initializeMyApp(db)
	server.NewServer()
}

3