package validators

import (
	"errors"
	"fmt"
	"GDColumn/pkg/database"
	"strconv"
	"strings"
	"unicode/utf8"

	"github.com/thedevsaddam/govalidator"
)

func init() {

	govalidator.AddCustomRule("not_exists", func(field string, rule string, message string, value interface{}) error {

		rng := strings.Split(strings.TrimPrefix(rule, "not_exists:"), ",")
		tableName := rng[0]
		dbFiled := rng[1]

		var exceptID string
		if len(rng) > 2 {
			exceptID = rng[2]
		}

		requestValue := value.(string)

		query := database.DB.Table(tableName).Where(dbFiled+" = ?", requestValue)
		if len(exceptID) > 0 {
			query.Where("id != ?", exceptID)
		}

		var count int64
		query.Count(&count)
		if count != 0 {
			if message != "" {
				return errors.New(message)
			}
			return fmt.Errorf("%v 已被占用", requestValue)
		}
		return nil
	})

	// max_cn:8 中文长度设定不超过 8
	govalidator.AddCustomRule("max_cn", func(field string, rule string, message string, value interface{}) error {
		valLength := utf8.RuneCountInString(value.(string))
		l, _ := strconv.Atoi(strings.TrimPrefix(rule, "max_cn:"))
		if valLength > l {
			// 如果有自定义错误消息的话，使用自定义消息
			if message != "" {
				return errors.New(message)
			}
			return fmt.Errorf("长度不能超过 %d 个字", l)
		}
		return nil
	})

	// min_cn:2 中文长度设定不小于 2
	govalidator.AddCustomRule("min_cn", func(field string, rule string, message string, value interface{}) error {
		valLength := utf8.RuneCountInString(value.(string))
		l, _ := strconv.Atoi(strings.TrimPrefix(rule, "min_cn:"))
		if valLength < l {
			// 如果有自定义错误消息的话，使用自定义消息
			if message != "" {
				return errors.New(message)
			}
			return fmt.Errorf("长度需大于 %d 个字", l)
		}
		return nil
	})

}