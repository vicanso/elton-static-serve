package main

import (
	"github.com/vicanso/elton"

	staticServe "github.com/vicanso/elton-static-serve"
)

func main() {
	e := elton.New()

	sf := new(staticServe.FS)
	// static file route
	e.GET("/*file", staticServe.New(sf, staticServe.Config{
		Path: "/tmp",
		// 客户端缓存一年
		MaxAge: 365 * 24 * 3600,
		// 缓存服务器缓存一个小时
		SMaxAge:             60 * 60,
		DenyQueryString:     true,
		DisableLastModified: true,
		// packr不支持Stat，因此需要用强ETag
		EnableStrongETag: true,
	}))

	err := e.ListenAndServe(":3000")
	if err != nil {
		panic(err)
	}
}
