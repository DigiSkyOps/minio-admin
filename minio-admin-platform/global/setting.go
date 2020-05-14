package util
import (
	"github.com/go-ini/ini"
	"log"
	"os"
	"time"
)

var (
	Cfg *ini.File
)

type SettingStruct struct {
	Cfg *ini.File
	RunMode string
	HTTPPort int
	ReadTimeout time.Duration
	WriteTimeout time.Duration
	MinioEndpoint string
	AccessKey string
	SecretKey string
}

var Setting SettingStruct

func init() {
	var err error
	Cfg, err = ini.Load("conf/app.ini")
	if err != nil {
		log.Fatalf("Fail to parse 'conf/app.ini': %v", err)
	}
	LoadBase()
	LoadServer()
	LoadMinio()
	InitMinioClient()
}
func LoadBase() {
	Setting.RunMode = Cfg.Section("").Key("RUN_MODE").MustString("debug")
}
func LoadServer() {
	sec, err := Cfg.GetSection("server")
	if err != nil {
		log.Fatalf("Fail to get section 'server': %v", err)
	}
	Setting.RunMode = Cfg.Section("").Key("RUN_MODE").MustString("debug")
	Setting.HTTPPort = sec.Key("HTTP_PORT").MustInt(8000)
	Setting.ReadTimeout = time.Duration(sec.Key("READ_TIMEOUT").MustInt(60)) * time.Second
	Setting.WriteTimeout =  time.Duration(sec.Key("WRITE_TIMEOUT").MustInt(60)) * time.Second
}
func LoadMinio() {
	Setting.AccessKey = os.Getenv("ACCESSKEY")
	Setting.SecretKey = os.Getenv("SECRETKEY")
	Setting.MinioEndpoint  = os.Getenv("MINIO_ENDPOINT")
}
