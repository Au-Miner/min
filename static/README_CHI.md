# Min

## What is Min
Min并不是求最小值的工具（x

Min是一个轻量级、简化版本的Gin轮子，Min专注于性能和简单性，适合想要深入学习Go框架的初学者，
Min实现了Gin的大部分基础功能，包括路由、中间件、请求/响应处理、错误处理、路由组等，用于快速搭建HTTP服务。

## Status
目前Min适用于学习和小型项目，其中性能和稳定性已经在MinerDB中得到了验证

## Design overview
![Min.png](Min.png)

## Key features
### Routing
目前实现了GET、POST

### Middleware

中间件机制允许在请求处理过程中插入自定义逻辑。

内置中间件：目前实现了日志记录、错误恢复。

自定义中间件：可以创建和使用自定义中间件处理请求前后逻辑。

### context
对于 Web 服务来说，无非就是根据请求`*http.Request`，构造响应`http.ResponseWriter`。

但是这个对于 HTTP 服务来说太细粒度了。所以我们采用 context 中间层来简化。

同时，我们也可以将动态参数放在 context 中，以及一些中间件产生的信息。

### Group
我们需要将一组 api 进行统一的管理，可以方便的对一组 api 添加中间件进行管理。



## Gettings Started
```
package main

import (
	"min"
	"net/http"
)

func main() {
	engine := min.New()

	group := engine.Group("/api")

	group.Use(min.Recovery(), min.Logger())

	group.GET("/echo/*name", func(ctx *min.Context) {
		ctx.JSON(http.StatusOK, min.H{"echo": ctx.Params["name"]})
	})
	group.POST("/ping", func(ctx *min.Context) {
		ctx.JSON(http.StatusOK, "pone")
	})

	group.POST("/panic", func(ctx *min.Context) {
		panic("panic")
	})

	engine.Run(":9999")
}
```

## Thanks
尤其感谢@ruanjiancheng的支持。Orz