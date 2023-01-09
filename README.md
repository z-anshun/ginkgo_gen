# Ginkgo_gen

## quick start
```shell
$ go build .
$ ./ginkgo_gen -f filePath
```

## example
```go
// Describe: 查询操作列表-/traces/operations
// 所有参数填写正确，返回正确数据
// cluster不存在，无数据返回
// namespace 不存在，无数据返回
// serviceName 不存在，无数据返回
type TracingQueryOperations struct {
	Cluster     *string
	ClusterName *string
	End         *string
	Namespace   *string
	ServiceName *string
	SpanKind    []string
	Start       *string
}
```
注意：这里结构体名，需要跟对应 grpc 的参数名一样，因为调用方法为 `New{{$.Name}}Params()`