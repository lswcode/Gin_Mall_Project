package conf

import (
	"fmt"
	"gin_mall/dao"
	"strings"

	"gopkg.in/ini.v1"
)

var (
	AppMode  string
	HttpPort string

	Db         string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassWord string
	DbName     string

	AccessKey   string
	SerectKey   string
	Bucket      string
	QiniuServer string

	ValidEmail string
	SmtpHost   string
	SmtpEmail  string
	SmtpPass   string

	EsHost  string
	EsPort  string
	EsIndex string
)

func Init() {
	//从本地读取环境变量
	file, err := ini.Load("./conf/config.ini") // 使用第三方库 ini 读取配置文件
	if err != nil {
		fmt.Println("配置文件读取错误，请检查文件路径:", err)
	}
	LoadServer(file)
	LoadMysqlData(file)

	// MySQL
	// strings.Join() 将字符串切片拼接为单个字符串，第一个参数是字符串切片。第二个参数是连接符
	pathRead := strings.Join([]string{DbUser, ":", DbPassWord, "@tcp(", DbHost, ":", DbPort, ")/", DbName, "?charset=utf8&parseTime=true"}, "")
	pathWrite := strings.Join([]string{DbUser, ":", DbPassWord, "@tcp(", DbHost, ":", DbPort, ")/", DbName, "?charset=utf8&parseTime=true"}, "")
	dao.Database(pathRead, pathWrite)
}

func LoadServer(file *ini.File) {
	AppMode = file.Section("server").Key("AppMode").String() // 读取指定分区的指定key对应的值，并转换为字符串
	HttpPort = file.Section("server").Key("HttpPort").String()
}

func LoadMysqlData(file *ini.File) {
	Db = file.Section("mysql").Key("Db").String()
	DbHost = file.Section("mysql").Key("DbHost").String()
	DbPort = file.Section("mysql").Key("DbPort").String()
	DbUser = file.Section("mysql").Key("DbUser").String()
	DbPassWord = file.Section("mysql").Key("DbPassWord").String()
	DbName = file.Section("mysql").Key("DbName").String()
}
