# Golang context初探

## 什么是context

从go1.7开始，golang.org/x/net/context包正式作为context包进入了标准库。那么，这个包到底是做什么的呢？根据官方的文档说明：

> Package context defines the Context type, which carries deadlines, cancelation signals, and other request-scoped values across API boundaries and between processes.

也就是说，通过context，我们可以方便地对同一个请求所产生地goroutine进行约束管理，可以设定超时、deadline，甚至是取消这个请求相关的所有goroutine。形象地说，假如一个请求过来，需要A去做事情，而A让B去做一些事情，B让C去做一些事情，A、B、C是三个有关联的goroutine，那么问题来了：假如在A、B、C还在处理事情的时候请求被取消了，那么该如何优雅地同时关闭goroutine A、B、C呢？这个时候就轮到context包上场了。

## 如何使用context

在开始说明如何使用context之前，先来看看context有什么相关的方法定义。
先来看看context的代码示例：

```go
package main

import (
    "context"
    "log"
    "net/http"
    _ "net/http/pprof"
    "time"
)

func main() {
    go http.ListenAndServe(":8989", nil)
    ctx, cancel := context.WithCancel(context.Background())
    go func() {
        time.Sleep(3 * time.Second)
        cancel()
    }()
    log.Println(A(ctx))
    select {}
}

func C(ctx context.Context) string {
    select {
    case <-ctx.Done():
        return "C Done..."
    }
    return ""
}

func B(ctx context.Context) string {
    ctx, _ = context.WithCancel(ctx)
    go log.Println(C(ctx))
    select {
    case <-ctx.Done():
        return "B Done..."
    }
    return ""
}

func A(ctx context.Context) string {
    go log.Println(B(ctx))
    select {
    case <-ctx.Done():
        return "A Done..."
    }
    return ""
}
```

运行结果：

> 2019/07/24 13:25:52 A Done...
> 2019/07/24 13:25:52 B Done...
> 2019/07/24 13:25:52 C Done...

pprof截图，刚开始的时候：



![img](https://upload-images.jianshu.io/upload_images/4084325-b4a1982967a8729a.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/801/format/webp)

刚开始.png

A、B、C结束后：



![img](https://upload-images.jianshu.io/upload_images/4084325-1fe5c438788d951e.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/808/format/webp)

结束后.png

在这里，我们用http的pprof来查看运行时有多少个goroutine正在执行，
`go http.ListenAndServe(":8989", nil)`
这一句是启动一个http服务器，用来查看程序的一些运行信息。
我们用`context.Background()`来实例化一个context，然后调用`WithCancel()`方法来返回一个context和一个取消的方法，并在3秒后调用这个cancel方法关闭goroutine A、B、C。从程序的运行结果看，我们调用了一次cancel方法，子goroutine的`ctx.Done()`都收到了关闭信号。从pprof的截图也可以看到，从一开始有5个goroutine，到关闭后剩下两个，也就是A、B、C三个goroutine都已经关闭了。
这里的例子是直接调用了`context.WithCancel()`，我们也可以使用`context.WithTimeout()`和`context.WithDeadline()`来设置goroutine的超时时间和最终的运行时间。具体的用法可以看一下官方文档，这里就不细说了。另外有一个方法在例子中没有用到，那就是`context.WithValue()`。这个方法是用来传递在这次的请求处理中相关goroutine的共享变量，这与全局变量是有所区别的，因为它只在这次的请求范围内有效。

## context的使用规范

在Golang的 1.8 版本中，很多相关的包都加入了context，比如database包。那么，在使用context的时候有哪些需要注意呢？

- 不要把context存储在结构体中，而是要显式地进行传递
- 把context作为第一个参数，并且一般都把变量命名为ctx
- 就算是程序允许，也不要传入一个nil的context，如果不知道是否要用context的话，用`context.TODO()`来替代
- `context.WithValue()`只用来传递请求范围的值，不要用它来传递可选参数
- 就算是被多个不同的goroutine使用，context也是安全的