package utile

import (
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

type Conf struct {
	Database struct {
		Host   string `yaml:"host"`
		Port   string `yaml:"port"`
		User   string `yaml:"user"`
		Passwd string `yaml:"passwd"`
		Db     string `yaml:"db"`
	} `yaml:"database"`
}

const (
	DB_Host     = "db_host"
	DB_Port     = "db_port"
	DB_User     = "db_user"
	DB_Password = "db_password"
	DB_Db       = "db_db_name"
)

func GetConf() {
	_, b, _, _ := runtime.Caller(0)
	basepath := filepath.Dir(b)

	path := strings.Replace(basepath, `pkg\utile`, "config.yaml", 2)
	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	var c Conf
	err = yaml.Unmarshal(yamlFile, &c)
	os.Setenv(DB_Host, c.Database.Host)
	os.Setenv(DB_Port, c.Database.Port)
	os.Setenv(DB_User, c.Database.User)
	os.Setenv(DB_Password, c.Database.Passwd)
	os.Setenv(DB_Db, c.Database.Db)

	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

}
func ReturnResponseWithJson(data interface{}) []byte {
	b, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return b
}
