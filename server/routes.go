// created by lilei at 2021/10/7
package server

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"iris-blog-server/handler"
)

func regisRouter(app *iris.Application) {
	mvc.Configure(app.Party("/api"), func(m *mvc.Application) {
		m.Party("/user").Handle(new(handler.UserHandler))
	})
}
