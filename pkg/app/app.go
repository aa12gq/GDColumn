package app

import "GDColumn/pkg/config"

// URL 传参 path 拼接站点的 URL
func URL(path string) string {
	return config.Get("app.url") + path
}

// V1URL 拼接带 v1 标示 URL
func V1URL(path string) string {
	return URL("/v1/" + path)
}
