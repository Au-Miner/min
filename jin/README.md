# jin

## context

对于 Web 服务来说，无非就是根据请求`*http.Request`，构造响应`http.ResponseWriter`。

但是这个对于 HTTP 服务来说太细粒度了。所以我们采用 context 中间层来简化。

同时，我们也可以将动态参数放在 context 中，以及一些中间件产生的信息。

## Group

我们需要将一组 api 进行统一的管理，可以方便的对一组 api 添加中间件进行管理。

