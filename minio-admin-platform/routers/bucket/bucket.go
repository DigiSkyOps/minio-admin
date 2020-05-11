package bucket

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go/v6"
	"minio-admin/util"
)

var (
	MinioClient *minio.Client
)

func init(){
	MinioClient = initMinioClient()
}

func initMinioClient() *minio.Client {
	minioClient, err :=  minio.New(util.Setting.MinioEndpoint,util.Setting.AccessKey, util.Setting.SecretKey,false)
	if err != nil {
		fmt.Println(err)
		util.Logger.Error(err)
	}
	return minioClient
}


func ListBuckets(c *gin.Context){
	buckets, err :=  MinioClient.ListBuckets()
	if err != nil {
		util.Logger.Error(err)
		c.JSON(200, gin.H{
			"code": -200,
			"message": err.Error(),
		})
	}else {
		c.JSON(200, gin.H{
			"code": 200,
			"data": buckets,
		})
	}
}
