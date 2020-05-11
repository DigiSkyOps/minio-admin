package user

import (
	"fmt"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/minio/minio/pkg/madmin"
	"minio-admin/util"
)

var (
	MinioClient *madmin.AdminClient
)
//const (
//	AccountEnabled  AccountStatus = "enabled"
//	AccountDisabled AccountStatus = "disabled"
//)

type AddUserInfo struct {
	AccessKey string	`json:"accessKey" binding:"required"`
	SecretKey string	`json:"secretKey" binding:"required"`
}


type DelUserInfo struct {
	AccessKey string	`json:"accessKey" binding:"required"`
}
type SetUserInfo struct {
	AccessKey string		`json:"accessKey" binding:"required"`
	SecretKey string		`json:"secretKey" binding:"required"`
	Status    madmin.AccountStatus	`json:"status" binding:"required"`
}



type SetUserStatusInfo struct {
	AccessKey string		`json:"accessKey" binding:"required"`
	Status    madmin.AccountStatus	`json:"status" binding:"required"`
}


func init(){
	MinioClient = initMinioClient()
}

func initMinioClient() *madmin.AdminClient{
	madmClnt, err := madmin.New(util.Setting.MinioEndpoint,util.Setting.AccessKey, util.Setting.SecretKey,false)
	if err != nil {
		fmt.Println(err)
		util.Logger.Error(err)
	}
	return madmClnt
}



func ListUser(c *gin.Context){
	ctx := context.Background()
	users, err := MinioClient.ListUsers(ctx)

	if err != nil {
		util.Logger.Error(err)
		c.JSON(200, gin.H{
			"code": -200,
			"message": err.Error(),
		})
	}else {
		c.JSON(200, gin.H{
			"code": 200,
			"data": users,
		})
	}
}

func GetUserInfo(c *gin.Context){
	ctx := context.Background()
	var accessKey string
	accessKey = c.Query("accessKey")
	userInfo, err := MinioClient.GetUserInfo(ctx,accessKey)

	if err != nil {
		util.Logger.Error(err)
		c.JSON(200, gin.H{
			"code": -200,
			"message": err.Error(),
		})
	}else {
		c.JSON(200, gin.H{
			"code": 200,
			"data": userInfo,
		})
	}
}


func AddUser(c *gin.Context){
	var addUserInfo AddUserInfo
	if err := c.ShouldBindJSON(&addUserInfo); err != nil {
		util.Logger.Error(err.Error())
		c.JSON(200, gin.H{
			"code": -200,
			"message": "缺少参数:" + err.Error(),
		})
		return
	}
	ctx := context.Background()
	err := MinioClient.AddUser(ctx,addUserInfo.AccessKey ,addUserInfo.SecretKey)
	if err != nil {
		util.Logger.Error(err)
		c.JSON(200, gin.H{
			"code": -200,
			"message": err.Error(),
		})
	}else{
		c.JSON(200, gin.H{
			"code": 200,
			"data": "add user success",
		})
	}
	return
}


func SetUser(c *gin.Context){
	var setUserInfo SetUserInfo
	if err := c.ShouldBindJSON(&setUserInfo); err != nil {
		fmt.Println(setUserInfo)
		util.Logger.Error(err.Error())
		c.JSON(200, gin.H{
			"code": -200,
			"message": "缺少参数:" + err.Error(),
		})
		return
	}
	ctx := context.Background()
	err := MinioClient.SetUser(ctx,setUserInfo.AccessKey,setUserInfo.SecretKey,setUserInfo.Status)
	if err != nil {
		util.Logger.Error(err.Error())
		c.JSON(200, gin.H{
			"code": -200,
			"message": err.Error(),
		})
	}else{
		c.JSON(200, gin.H{
			"code": 200,
			"data": "set user success",
		})
	}
	return
}
func SetUserStatus(c *gin.Context){
	var setUserStatusInfo SetUserStatusInfo
	if err := c.ShouldBindJSON(&setUserStatusInfo); err != nil {
		util.Logger.Error(err.Error())
		c.JSON(200, gin.H{
			"code": -200,
			"message": "缺少参数:" + err.Error(),
		})
		return
	}
	ctx := context.Background()
	err := MinioClient.SetUserStatus(ctx,setUserStatusInfo.AccessKey,setUserStatusInfo.Status)
	if err != nil {
		util.Logger.Error(err.Error())
		c.JSON(200, gin.H{
			"code": -200,
			"message": err.Error(),
		})
	}else{
		c.JSON(200, gin.H{
			"code": 200,
			"data": "set user status success",
		})
	}
	return
}


func DelUser(c *gin.Context){
	var delUserInfo DelUserInfo
	if err := c.ShouldBindJSON(&delUserInfo); err != nil {
		util.Logger.Error(err.Error())
		c.JSON(200, gin.H{
			"code": -200,
			"message": "缺少参数:" + err.Error(),
		})
		return
	}
	ctx := context.Background()
	err := MinioClient.RemoveUser(ctx,delUserInfo.AccessKey)
	if err != nil {
		util.Logger.Error(err)
		c.JSON(200, gin.H{
			"code": -200,
			"message": err.Error(),
		})
	}else{
		c.JSON(200, gin.H{
			"code": 200,
			"data": "remove user success",
		})
	}
	return
}