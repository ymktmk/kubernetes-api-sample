## kubernetes-api-sample

### controller.Run()

1. Reflectorを作成して実行する(Watchしてキューに詰める)
2. キューから取り出す

ProcessFuncで処理します。これらの両方`stopCh` が閉じられるまで続けます

```go:
stopCh := make(chan struct{})
defer close(stopCh)
go controller.Run(stopCh)
select {}
```
