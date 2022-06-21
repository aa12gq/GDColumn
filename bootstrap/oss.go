package bootstrap

import (
	"GDColumn/pkg/config"
	"GDColumn/pkg/logger"
	"GDColumn/pkg/oss"
)

func SetupOss() {
	var oss oss.OssStruct
	oss.Endpoint = "oss-cn-hangzhou.aliyuncs.com"
	oss.AccessKeyId = config.Get("oss.accessKeyId")
	oss.AccessKeySecret = config.Get("oss.accessKeySecret")
	oss.Bucket = "bitpig-column"
	oss.Region = "华东1(杭州)"

	if err := oss.OssCennect(&oss); err != nil {
		logger.LogIf(err)
	}
}
