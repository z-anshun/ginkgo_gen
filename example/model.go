package tracing

// Describe: 查询操作列表-/traces/operations
// 所有参数填写正确，返回正确数据
// cluster不存在，无数据返回
// namespace 不存在，无数据返回
type TracingQueryOperations struct {
	Cluster     *string
	ClusterName *string
	End         *string
	Namespace   *string
	ServiceName *string
	SpanKind    []string
	Start       *string
}
