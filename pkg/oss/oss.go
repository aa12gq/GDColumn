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

func (this *OssStruct) OssCennect(o *OssStruct) (err error) {
	client, err := oss.New(o.Endpoint, o.AccessKeyId, o.AccessKeySecret)
	if err != nil {
		fmt.Println("Error:", err)
	}
	Bucket, err = client.Bucket(o.Bucket)
	if err != nil {
		fmt.Println("Error:", err)
	}
	return
}
