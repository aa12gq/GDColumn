package file

import (
	"GDColumn/pkg/helpers"
	aliyun "GDColumn/pkg/oss"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"mime/multipart"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"
	"fmt"
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

func SaveUploadImage(id uint64, c *gin.Context, file *multipart.FileHeader) (avatarPath,imageName,extName string, err error) {

	var avatar string
	publicPath := "public"
	dirName := fmt.Sprintf("/uploads/images/")
	os.MkdirAll(publicPath+dirName, 0755)

	num := strconv.FormatUint(id,10)
	suffix := path.Ext(file.Filename)
	newSuffix := path.Ext(file.Filename)
	if newSuffix != ".jpg"{
		newSuffix = ".jpg"
	}
	fileName := fileNameFromUploadFile(num,file)
	avatarPath = publicPath + dirName + fileName
	if err := c.SaveUploadedFile(file, avatarPath); err != nil {
		return avatar,file.Filename,newSuffix, err
	}

	pwd,_ := os.Getwd()
	imgPwd := fmt.Sprintf("%v/%v%v%v%v",pwd,publicPath,dirName,num,suffix)

	err = aliyun.Bucket.PutObjectFromFile(fmt.Sprintf("exampledir/%v%v",num,newSuffix),imgPwd);
	os.Remove(imgPwd)
	avatarPath = fmt.Sprintf("https://bitpig-column.oss-cn-hangzhou.aliyuncs.com/exampledir/%v%v",num,newSuffix)
	return avatarPath,file.Filename,newSuffix, err
}

func randomNameFromUploadFile(file *multipart.FileHeader) string {
	return helpers.RandomString(16) + filepath.Ext(file.Filename)
}

func fileNameFromUploadFile(name string, file *multipart.FileHeader) string {
	return name + filepath.Ext(file.Filename)
}
