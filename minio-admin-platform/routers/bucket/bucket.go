package bucket

import (
	"github.com/gin-gonic/gin"
	"minio-admin/util"
)

func ListBuckets(c *gin.Context){
	buckets, err :=  util.MinioClient.Server.ListBuckets()
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
