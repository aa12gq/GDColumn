package file

import (
	"GDColumn/pkg/helpers"
	aliyun "GDColumn/pkg/oss"
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"mime/multipart"
	"os"
	"path"
	"path/filepath"
	"strings"
)

// Put 将数据存入文件
func Put(data []byte, to string) error {
	err := ioutil.WriteFile(to, data, 0644)
	if err != nil {
		return err
	}
	return nil
}

// Exists 判断文件是否存在
func Exists(fileToCheck string) bool {
	if _, err := os.Stat(fileToCheck); os.IsNotExist(err) {
		return false
	}
	return true
}

func FileNameWithoutExtension(fileName string) string {
	return strings.TrimSuffix(fileName, filepath.Ext(fileName))
}

func SaveUploadImage(id string, c *gin.Context, file *multipart.FileHeader) (avatarPath string, err error) {

	var avatar string
	publicPath := "public"
	dirName := fmt.Sprintf("/uploads/images/")
	os.MkdirAll(publicPath+dirName, 0755)
	if err := c.SaveUploadedFile(file, publicPath + dirName + fileNameFromUploadFile(id,file)); err != nil {
		return avatar, err
	}

	pwd,_ := os.Getwd()
	imgPwd := fmt.Sprintf("%v/%v%v%v%v",pwd,publicPath,dirName,id,path.Ext(file.Filename))
	err = aliyun.Bucket.PutObjectFromFile(fmt.Sprintf("exampledir/%v%v",id,".jpg"),
		imgPwd,oss.ContentType("image/jpg"));
	os.Remove(imgPwd)
	return fmt.Sprintf("https://bitpig-column.oss-cn-hangzhou.aliyuncs.com/exampledir/%v%v",id,".jpg"), err
}

func randomNameFromUploadFile(file *multipart.FileHeader) string {
	return helpers.RandomString(16) + filepath.Ext(file.Filename)
}

func fileNameFromUploadFile(name string, file *multipart.FileHeader) string {
	return name + filepath.Ext(file.Filename)
}
