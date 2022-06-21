package config

import "GDColumn/pkg/config"

func init() {

	config.Add("oss", func() map[string]interface{} {
		return map[string]interface{}{

			"accessKeyId":     config.Env("ACCESSKEYID", ""),
			"accessKeySecret": config.Env("ACCESSKEYSECRET", ""),
		}
	})
}
