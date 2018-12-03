# quiz01

## how to run

```bash
$ go run main.go
```

## goal

これはchannelによるキューイング処理のテストコードです。
channelを使ってジョブキューイング(またはキューワーカー)を実装して下さい。

- 実装は `main.go` のみ
- 最大+50行くらいの改修内容

## note

[`main.go` の `L22` から `L30`](https://github.com/speecan/quiz01/blob/master/main.go#L22-L30) の
`job` の数だけ `worker` を作成して逐次実行している部分をどうにかします。

一つの `job` は `DummyData [1024 * 1024]byte` というフィールドを持っており、
このままではgolangがデフォルトで確保するキューの数だけメモリも確保されてしまいます。

`maxQueue` の数だけキューイングし、そのキューから `worker` がそれぞれ `job` を実行することで、
goroutineの並列実行性とメモリの消費量を制御することができます。
> デフォルトの `numberOfJobs` は `30000` ですが、
> `numberOfJobs` が増えれば増えるほど逐次処理しきれなかった分だけ使用するメモリの量が増加します。
> ジョブキューイングが実装できていると `maxWorkers = 1000` と `maxQueue = 1000` で最大約4GBほど使用されるはずです。
> 数値を色々変えてみるとメモリの使用量も変化することが分かるはずです。

## 参考

[go tour の concurrency 以降](https://go-tour-jp.appspot.com/concurrency)などを参考にしてみて下さい。
