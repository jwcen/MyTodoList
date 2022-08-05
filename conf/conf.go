package conf

import (
	"MyTodoList/model"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gopkg.in/ini.v1"

	"fmt"
	"strings"
)

var (
	AppMode    string
	HttpPort   string
	Db         string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassword string
	DbName     string
)

// Init 初始化加载配置文件
func Init() {
	file, err := ini.Load("./conf/config.ini")
	if err != nil {
		fmt.Println("配置文件读取错误，请检查文件路径")
	}

	LoadServer(file)
	LoadMysql(file)
	// 配置数据库连接
	path := strings.Join([]string{DbUser, ":", DbPassword, "@tcp(", DbHost, ":", DbPort, ")/",
		DbName, "?charset=utf8mb4&parseTime=true"}, "")
	model.Database(path)
}

// LoadServer 加载服务器配置
func LoadServer(file *ini.File) {
	AppMode = file.Section("service").Key("AppMode").String()
	HttpPort = file.Section("service").Key("HttpPort").String()
}

// LoadMysql 加载MySQL配置
func LoadMysql(file *ini.File) {
	Db = file.Section("mysql").Key("Db").String()
	DbHost = file.Section("mysql").Key("DbHost").String()
	DbPort = file.Section("mysql").Key("DbPort").String()
	DbUser = file.Section("mysql").Key("DbUser").String()
	DbPassword = file.Section("mysql").Key("DbPassword").String()
	DbName = file.Section("mysql").Key("DbName").String()
}
