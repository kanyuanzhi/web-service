package global

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"io/ioutil"
)

var Object *Obj

type Obj struct {
	Host string `json:"host,omitempty"`
	Port int    `json:"port,omitempty"`

	MysqlUsername string `json:"mysql_username,omitempty"`
	MysqlPassword string `json:"mysql_password,omitempty"`
	MysqlHost     string `json:"mysql_host,omitempty"`
	MysqlPort     uint   `json:"mysql_port,omitempty"`
	MysqlDbname   string `json:"mysql_dbname,omitempty"`

	// tcp settings

	LogMode bool `json:"log_mode,omitempty"`
}

func (g *Obj) Reload() {
	data, err := ioutil.ReadFile("conf/webService.json")
	if err != nil {
		panic(err.Error())
	}

	err = json.Unmarshal(data, &Object)
	if err != nil {
		panic(err.Error())
	}
	Log = logrus.New()
	Log.SetReportCaller(Object.LogMode)
	if Object.LogMode == true {
		Log.SetLevel(logrus.TraceLevel)
	} else {
		Log.SetLevel(logrus.InfoLevel)
	}
	Log.SetFormatter(&CustomFormatter{})
}

func init() {
	Object = &Obj{
		Host: "127.0.0.1",
		Port: 8080,

		MysqlUsername: "admin",
		MysqlPassword: "c3i123456",
		MysqlHost:     "127.0.0.1",
		MysqlPort:     3306,
		MysqlDbname:   "web_service",

		LogMode: true, // true：详细，打印log在代码中输出位置；false：简要，不打印文件输出位置，不打印debug和trace（性能高，生产环境使用）
	}

	Object.Reload()

	DB = initDB()
}
