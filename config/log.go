package config

import "GDColumn/pkg/config"

func init() {
	config.Add("log", func() map[string]interface{} {
		return map[string]interface{}{

			"level": config.Env("LOG_LEVEL", "debug"),

			"type": config.Env("LOG_TYPE", "single"),

			/* ------------------ 滚动日志配置 ------------------ */
			// 日志文件路径
			"filename": config.Env("LOG_NAME", "storage/logs/logs.log"),
			// 每个日志文件保存的最大尺寸 单位：M
			"max_size": config.Env("LOG_MAX_SIZE", 64),
			// 最多保存日志文件数，0 为不限，MaxAge 到了还是会删
			"max_backup": config.Env("LOG_MAX_BACKUP", 5),
			// 最多保存多少天，7 表示一周前的日志会被删除，0 表示不删
			"max_age": config.Env("LOG_MAX_AGE", 30),
			// 是否压缩，压缩日志不方便查看，我们设置为 false（压缩可节省空间）
			"compress": config.Env("LOG_COMPRESS", false),
		}
	})
}
