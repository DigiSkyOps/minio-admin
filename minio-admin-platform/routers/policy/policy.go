package policy

import (
	"context"
	"github.com/gin-gonic/gin"
	"minio-admin/util"
)

type AddPolicyInfo struct {
	PolicyName string		`json:"policyname",binding:"required"`
	PolicyContent string	`json:"policycontent",binding:"required"`
}

type SetPolicyInfo struct {
	PolicyName string		`json:"policyname",binding:"required"`
	AccessKey string		`json:"accesskey",binding:"required"`
	IsGroup		bool		`json:"isgroup",binding:"required"`
}

type DelPolicyInfo struct {
	Name string	`json:"name",binding:"required"`
}

func ListPolicy(c *gin.Context){
	policys, err := util.MinioClient.Admin.ListCannedPolicies(context.Background())

	if err != nil {
		util.Logger.Error(err)
		c.JSON(200, gin.H{
			"code": -200,
			"message": err.Error(),
		})
	}else {
		c.JSON(200, gin.H{
			"code": 200,
			"data": policys,
		})
	}
}

func GetPolicyInfo(c *gin.Context){
	ctx := context.Background()
	var policyName string
	policyName = c.Query("name")
	policyInfo, err := util.MinioClient.Admin.InfoCannedPolicy(ctx,policyName)

	if err != nil {
		util.Logger.Error(err)
		c.JSON(200, gin.H{
			"code": -200,
			"message": err.Error(),
		})
	}else {
		c.JSON(200, gin.H{
			"code": 200,
			"data": policyInfo,
		})
	}
}



func AddPolicy(c *gin.Context){
	var addPolicyInfo AddPolicyInfo
	if err := c.ShouldBindJSON(&addPolicyInfo); err != nil {
		util.Logger.Error(err.Error())
		c.JSON(200, gin.H{
			"code": -200,
			"message": "缺少参数:" + err.Error(),
		})
		return
	}
	ctx := context.Background()
	//policy := `{"Version":"2012-10-17","Statement":[{"Effect":"Allow","Action":["s3:*"],"Resource":["arn:aws:s3:::nxx-n/*"]}]}`
	err := util.MinioClient.Admin.AddCannedPolicy(ctx,addPolicyInfo.PolicyName ,addPolicyInfo.PolicyContent)
	if err != nil {
		util.Logger.Error(err)
		c.JSON(200, gin.H{
			"code": -200,
			"message": err.Error(),
		})
	}else{
		c.JSON(200, gin.H{
			"code": 200,
			"data": "add policy success",
		})
	}
	return
}


func SetPolicy(c *gin.Context){
	var setPolicyInfo SetPolicyInfo
	if err := c.ShouldBindJSON(&setPolicyInfo); err != nil {
		util.Logger.Error(err.Error())
		c.JSON(200, gin.H{
			"code": -200,
			"message": "缺少参数:" + err.Error(),
		})
		return
	}
	ctx := context.Background()
	err := util.MinioClient.Admin.SetPolicy(ctx,setPolicyInfo.PolicyName ,setPolicyInfo.AccessKey,setPolicyInfo.IsGroup)
	if err != nil {
		util.Logger.Error(err)
		c.JSON(200, gin.H{
			"code": -200,
			"message": err.Error(),
		})
	}else{
		c.JSON(200, gin.H{
			"code": 200,
			"data": "set policy success",
		})
	}
	return
}


func DelPolicy(c *gin.Context){
	var delPolicyInfo DelPolicyInfo
	if err := c.ShouldBindJSON(&delPolicyInfo); err != nil {
		util.Logger.Error(err.Error())
		c.JSON(200, gin.H{
			"code": -200,
			"message": "缺少参数:" + err.Error(),
		})
		return
	}
	ctx := context.Background()
	err := util.MinioClient.Admin.RemoveCannedPolicy(ctx,delPolicyInfo.Name)
	if err != nil {
		util.Logger.Error(err)
		c.JSON(200, gin.H{
			"code": -200,
			"message": err.Error(),
		})
	}else{
		c.JSON(200, gin.H{
			"code": 200,
			"data": "remove policy success",
		})
	}
	return
}


