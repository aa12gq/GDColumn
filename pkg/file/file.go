package file

import (
	"GDColumn/pkg/helpers"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"mime/multipart"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"fmt"
	"time"
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

func SaveUploadAvatar(id uint64, c *gin.Context, file *multipart.FileHeader) (string, error) {

	var avatar string
	// 确保目录存在，不存在创建
	publicPath := "public"
	dirName := fmt.Sprintf("/uploads/avatars/")
	os.MkdirAll(publicPath+dirName, 0755)
	//userModel := auth.CurrentUser(c)
	num := strconv.FormatUint(id,10)
	//num := strconv.FormatUint(userModel.AvatarID,10)
	// 保存文件
	fileName := randomNameFromUploadFile2(num,file)
	// public/uploads/avatars/2021/12/22/1/nFDacgaWKpWWOmOt.png
	avatarPath := publicPath + dirName + fileName
	if err := c.SaveUploadedFile(file, avatarPath); err != nil {
		return avatar, err
	}

	return avatarPath, nil
}

func randomNameFromUploadFile(file *multipart.FileHeader) string {
	return helpers.RandomString(16) + filepath.Ext(file.Filename)
}

func randomNameFromUploadFile2(random string, file *multipart.FileHeader) string {

	return random + filepath.Ext(file.Filename)
}


func SaveUPTest(c *gin.Context, file *multipart.FileHeader) (string, error) {

	var avatar string
	// 确保目录存在，不存在创建
	publicPath := "https://bitpig-column.oss-cn-hangzhou.aliyuncs.com/exampledir"
	t := time.Now()
	dirName := fmt.Sprintf("%s%s%s%s%s%s%s", strconv.Itoa(t.Year()), strconv.Itoa(int(t.Month())), strconv.Itoa(t.Day()), strconv.Itoa(t.Hour()), strconv.Itoa(t.Minute()), strconv.Itoa(t.Second()), strconv.Itoa(int(t.Unix())))
	//os.MkdirAll(publicPath+dirName, 0755)

	// 保存文件
	fileName := randomNameFromUploadFile(file)
	// public/uploads/avatars/2021/12/22/1/nFDacgaWKpWWOmOt.png
	avatarPath := publicPath + dirName + fileName
	if err := c.SaveUploadedFile(file, avatarPath); err != nil {
		return avatar, err
	}

	return avatarPath, nil
}