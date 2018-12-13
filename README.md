# go-cloud-service
A simple blog server

## Utils
提供了一套`http server`用的`utils`，主要有：
* `db`: `boltdb`的简单封装
* `exception`, `handleException`: 一套简单的异常处理方式
* `logrequest`: 配合异常处理的一个log中间件
* `middleware`: 一个简单的中间件包装工具
* `utils`: 发送请求之类的小工具

### Middleware
类似`express`的中间件机制，使用方法如下:
```go
import (
  "net/http"
  "github.com/leiysky/cloud-service/utils"
)

func main() {
  http.Handle("/path/you/want/to/handle", utils.HandlerCompose(
    Middleware1,
    Middleware2,
    // ...
    MiddlewareN,
  ))
  http.ListenAndServe(":8888", nil)
}

func Middleware1(w http.ResponseWriter, req *http.Request, next utils.NextFun) error {
  // do something
  err := next()
  // do something after Middleware2 and Middleware... returned
}

```

### Exception
提供了一套异常处理的方式，主要有以下几个规则：
* 使用**500以下**的状态码的异常为**Soft error**
* 使用**500以上**的状态码的异常为**Hard error**
* 异常均由`handleException`中间件处理，不同在于**Soft error**不会引起程序崩溃，**Hard error**会直接导致进程中断
* 使用`panic`发送一个`Exception`对象表示异常

（为了方便我给`next`也加上了`error`类型的返回值，也可以被上层中间件处理）

用法如下:
```go
func Middleware1(w http.ResponseWriter, req *http.Request, next utils.NextFun) error {
  // try to get some resource
  if resource != nil {
    return util.SendData(w, resource, msg, status) // actually the resource here should be a JSON string
  } else {
    panic(utils.Exception{"Not Found", 404})
  }
}
```
### Router
还没写，有时间补上
