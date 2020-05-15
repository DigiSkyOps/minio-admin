package util

import (
	"github.com/minio/minio-go/v6"
	"github.com/minio/minio/pkg/madmin"
	"minio-admin/global"
)

var (
	MinioClient *SMinio
)

type SMinio struct {
	Admin *madmin.AdminClient
	Server *minio.Client
}

func init(){
	initMinioClient()
}

func initMinioClient(){
	madmClnt, err := madmin.New(global.Setting.MinioEndpoint,global.Setting.AccessKey, global.Setting.SecretKey,false)
	if err != nil {
		Logger.Fatal(err)
	}

	mserverClnt,err := minio.New(global.Setting.MinioEndpoint,global.Setting.AccessKey, global.Setting.SecretKey,false)
	if err != nil {
		Logger.Fatal(err)
	}
	MinioClient = &SMinio{
		Admin: madmClnt,
		Server: mserverClnt,
	}
}
