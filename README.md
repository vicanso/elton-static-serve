# cod-static-serve

Static serve for cod, it use to serve static file, such as html, image and etc.

```go
package main

import (
	"os"

	packr "github.com/gobuffalo/packr/v2"
	"github.com/vicanso/cod"

	staticServe "github.com/vicanso/cod-static-serve"
)

var (
	box = packr.New("asset", "./web")
)

type (
	staticFile struct {
		box *packr.Box
	}
)

func (sf *staticFile) Exists(file string) bool {
	return sf.box.Has(file)
}
func (sf *staticFile) Get(file string) ([]byte, error) {
	return sf.box.Find(file)
}
func (sf *staticFile) Stat(file string) os.FileInfo {
	return nil
}

func main() {
	d := cod.New()

	sf := &staticFile{
		box: box,
	}

	// static file route
	d.GET("/static/*file", staticServe.New(sf, staticServe.Config{
		// 客户端缓存一年
		MaxAge: 365 * 24 * 3600,
		// 缓存服务器缓存一个小时
		SMaxAge:             60 * 60,
		DenyQueryString:     true,
		DisableLastModified: true,
	}))

	d.ListenAndServe(":7001")
}
```