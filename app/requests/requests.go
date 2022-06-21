package requests

import (
	"GDColumn/pkg/response"
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type ValidatorFunc func(interface{}, *gin.Context) map[string][]string

func Validate(c *gin.Context, obj interface{}, handler ValidatorFunc) bool {

	if err := c.ShouldBind(obj); err != nil {
		response.BadRequest(c, err, "请求解析错误，请确认请求格式是否正确。上传文件请使用 multipart 标头，参数请使用 JSON 格式。")
		return false
	}

	errs := handler(obj, c)

	if len(errs) > 0 {
		response.ValidationError(c, errs)
		return false
	}

	return true
}


func validate(data interface{}, rules govalidator.MapData, messages govalidator.MapData) map[string][]string {

	opts := govalidator.Options{
		Data:          data,
		Rules:         rules,
		TagIdentifier: "valid", // 模型中的 Struct 标签标识符
		Messages:      messages,
	}

	return govalidator.New(opts).ValidateStruct()
}

func validateFile(c *gin.Context, data interface{},
	rules govalidator.MapData, messages govalidator.MapData) map[string][]string {
	opts := govalidator.Options{
		Request:		c.Request,
		Rules:			rules,
		Messages:		messages,
		TagIdentifier:	"valid",
	}
	// 调用 govalidator 的 Validate 方法来验证文件
	return govalidator.New(opts).Validate()
}