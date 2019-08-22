#关于context


Go的设计者早考虑多个Goroutine共享数据，以及多Goroutine管理机制。Context介绍请参考Go Concurrency Patterns: Context，golang.org/x/net/context包就是这种机制的实现。

几种处理并发

WaitGroup
chan + select
Context

```
    func main() {
        stop := make(chan bool)

        go func() {
            for {
                select {
                case <-stop:
                    fmt.Println("监控退出，停止了...")
                    return
                default:
                    fmt.Println("goroutine监控中...")
                    time.Sleep(2 * time.Second)
                }
            }
        }()

        time.Sleep(10 * time.Second)
        fmt.Println("可以了，通知监控停止")
        stop<- true
        //为了检测监控过是否停止，如果没有监控输出，就表示停止了
        time.Sleep(5 * time.Second)
    }

    func main() {
        ctx, cancel := context.WithCancel(context.Background())
        go func(ctx context.Context) {
            for {
                select {
                case <-ctx.Done():
                    fmt.Println("监控退出，停止了...")
                    return
                default:
                    fmt.Println("goroutine监控中...")
                    time.Sleep(2 * time.Second)
                }
            }
        }(ctx)

        time.Sleep(10 * time.Second)
        fmt.Println("可以了，通知监控停止")
        cancel()
        //为了检测监控过是否停止，如果没有监控输出，就表示停止了
        time.Sleep(5 * time.Second)

    }
```

示例中启动了3个监控goroutine进行不断的监控，每一个都使用了Context进行跟踪，当我们使用cancel函数通知取消时，这3个goroutine都会被结束。这就是Context的控制能力，它就像一个控制器一样，按下开关后，所有基于这个Context或者衍生的子Context都会收到通知，这时就可以进行清理操作了，最终释放goroutine，这就优雅的解决了goroutine启动后不可控的问题。

```
    func main() {
        ctx, cancel := context.WithCancel(context.Background())
        go watch(ctx,"【监控1】")
        go watch(ctx,"【监控2】")
        go watch(ctx,"【监控3】")

        time.Sleep(10 * time.Second)
        fmt.Println("可以了，通知监控停止")
        cancel()
        //为了检测监控过是否停止，如果没有监控输出，就表示停止了
        time.Sleep(5 * time.Second)
    }

    func watch(ctx context.Context, name string) {
        for {
            select {
            case <-ctx.Done():
                fmt.Println(name,"监控退出，停止了...")
                return
            default:
                fmt.Println(name,"goroutine监控中...")
                time.Sleep(2 * time.Second)
            }
        }
    }
```

#关于tudo和background

TODO返回一个非空，空的上下文
在目前还不清楚要使用的上下文或尚不可用时
```
    context.TODO()
```

Background返回一个非空，空的上下文。
这是没有取消，没有值，并且没有期限。
它通常用于由主功能，初始化和测试，并作为输入的顶层上下文
```
    context.Background()
```

##主要方法
```
    func WithCancel(parent Context) (ctx Context, cancel CancelFunc)
    func WithDeadline(parent Context, deadline time.Time) (Context, CancelFunc)
    func WithTimeout(parent Context, timeout time.Duration) (Context, CancelFunc)
    func WithValue(parent Context, key interface{}, val interface{}) Context
```

// 获取失败, 则在该函数结束时结束 ...
```
    ctx, cancel = context.WithCancel(context.Background())
```

WithCancel 对应的是 cancelCtx ,其中，返回一个 cancelCtx ，同时返回一个 CancelFunc，CancelFunc 是 context 包中定义的一个函数类型：type CancelFunc func()。调用这个 CancelFunc 时，关闭对应的c.done，也就是让他的后代goroutine退出

WithDeadline 和 WithTimeout 对应的是 timerCtx ，WithDeadline 和 WithTimeout 是相似的，WithDeadline 是设置具体的 deadline 时间，到达 deadline 的时候，后代 goroutine 退出，而 WithTimeout 简单粗暴，直接 return WithDeadline(parent, time.Now().Add(timeout))

WithValue 对应 valueCtx ，WithValue 是在 Context 中设置一个 map，拿到这个 Context 以及它的后代的 goroutine 都可以拿到 map 里的值

context的创建
所有的context的父对象，也叫根对象，是一个空的context，它不能被取消，它没有值，从不会被取消，也没有超时时间，它常常作为处理request的顶层context存在，然后通过WithCancel、WithTimeout函数来创建子对象来获得cancel、timeout的能力

当顶层的request请求函数结束后，我们就可以cancel掉某个context，从而通知别的routine结束

WithValue方法可以把键值对加入context中，让不同的routine获取

#关于rabbitmq
1. 超级管理员(administrator)
2. 监控者(monitoring)
3. 策略制定者(policymaker)
4. 普通管理者(management)
5. 其他

是否登陆和监控