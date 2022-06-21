package oss

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"fmt"
)

var Bucket *oss.Bucket

type OssStruct struct {
	Endpoint 		string
	AccessKeyId 	string
	AccessKeySecret string
	Region 			string
	Bucket 			string
	Secure 			bool
	Cname 		    bool
}

func (this *OssStruct) OssCennect(o OssStruct) (err error) {
	//此处需要进入阿里云oss控制台配置域名
	this.Endpoint = o.Endpoint
	this.AccessKeyId = o.AccessKeyId
	this.AccessKeySecret = o.AccessKeySecret
	this.Region = o.Region
	this.Bucket = o.Bucket
	this.Secure = true
	this.Cname = true
	client, err := oss.New(this.Endpoint, this.AccessKeyId, this.AccessKeySecret, oss.UseCname(true))
	if err != nil {
		fmt.Println(err)
	}
	Bucket, err = client.Bucket("bitpig-column")

	return
}

//func (this *OssStruct) LocalUrl(c *gin.Context, file *multipart.FileHeader) (url string, err error) {
//
//	userModel := auth.CurrentUser(c)
//	//拼接文件名称
//	AvatarID := fmt.Sprintf("%s",userModel.AvatarID)
//	str := "https://bitpig-column.oss-cn-hangzhou.aliyuncs.com/exampledir" + AvatarID + ".jpg"
//	fmt.Println("拼接好的",str)
//
//	err = Bucket.UploadFile(file)
//	if err != nil {
//		url = "上传错误"
//	} else {
//		url = fmt.Sprintf("%s%s", "https://static.xxxxxxxi.cn/", str)
//	}
//	return url, err
//
//}
