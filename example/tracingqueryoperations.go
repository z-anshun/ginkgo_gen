package tracing

import (
	"fmt"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	tracing_client "insight.io/api/sdk/v1alpha1/client/tracing"
	"insight.io/test/tools"

	"k8s.io/klog/v2"
)

//go:generate ginkgo --focus=查询拓扑中的节点指标-/service-graph/node-metrics
var _ = Describe("查询拓扑中的节点指标-/service-graph/node-metrics", func() {
	client := tools.Client{Client: tools.NewClient(nil, tools.Cfg)}

	//TODO: var param
	cluster := ""
	clusterName := ""
	end := ""
	namespace := ""
	serviceName := ""
	spanKind := ""
	start := ""

	defaultParams := func() *tracing_client.TracingQueryOperationsParams {
		return tracing_client.NewTracingQueryOperationsParams().
			WithCluster(&cluster).
			WithClusterName(&clusterName).
			WithEnd(&end).
			WithNamespace(&namespace).
			WithServiceName(&serviceName).
			WithSpanKind([]string{spanKind}).
			WithStart(&start)
	}

	Context("1. 所有参数填写正确", func() {
		params := defaultParams()
		resp, err := client.Client.Tracing.TracingQueryOperations(params, tools.GenAuthClientOption())
		if err != nil {
			klog.Errorln("TracingQueryOperations error:", err)
		}
		It("返回正确数据", func() {
			Expect(len(resp.Payload.Nodes)).Should(BeNumerically(">", 0))
		})
	})

	Context("2. service不填", func() {
		params := defaultParams()
		resp, err := client.Client.Tracing.TracingQueryOperations(params, tools.GenAuthClientOption())
		if err != nil {
			klog.Errorln("TracingQueryOperations error:", err)
		}
		It("正确的错误提示", func() {
			Expect(len(resp.Payload.Nodes)).Should(BeNumerically(">", 0))
		})
	})

	Context("3. service不存在", func() {
		params := defaultParams()
		resp, err := client.Client.Tracing.TracingQueryOperations(params, tools.GenAuthClientOption())
		if err != nil {
			klog.Errorln("TracingQueryOperations error:", err)
		}
		It("无数据返回", func() {
			Expect(len(resp.Payload.Nodes)).Should(BeNumerically(">", 0))
		})
	})

	Context("4. cluster不填", func() {
		params := defaultParams()
		resp, err := client.Client.Tracing.TracingQueryOperations(params, tools.GenAuthClientOption())
		if err != nil {
			klog.Errorln("TracingQueryOperations error:", err)
		}
		It("正确的错误提示", func() {
			Expect(len(resp.Payload.Nodes)).Should(BeNumerically(">", 0))
		})
	})

	Context("5. cluster不存在", func() {
		params := defaultParams()
		resp, err := client.Client.Tracing.TracingQueryOperations(params, tools.GenAuthClientOption())
		if err != nil {
			klog.Errorln("TracingQueryOperations error:", err)
		}
		It("无数据返回", func() {
			Expect(len(resp.Payload.Nodes)).Should(BeNumerically(">", 0))
		})
	})

	Context("6. namespace不填", func() {
		params := defaultParams()
		resp, err := client.Client.Tracing.TracingQueryOperations(params, tools.GenAuthClientOption())
		if err != nil {
			klog.Errorln("TracingQueryOperations error:", err)
		}
		It("正确的错误提示", func() {
			Expect(len(resp.Payload.Nodes)).Should(BeNumerically(">", 0))
		})
	})

	Context("7. namespace不存在", func() {
		params := defaultParams()
		resp, err := client.Client.Tracing.TracingQueryOperations(params, tools.GenAuthClientOption())
		if err != nil {
			klog.Errorln("TracingQueryOperations error:", err)
		}
		It("无数据返回", func() {
			Expect(len(resp.Payload.Nodes)).Should(BeNumerically(">", 0))
		})
	})

	Context("8. endTime格式错误", func() {
		params := defaultParams()
		resp, err := client.Client.Tracing.TracingQueryOperations(params, tools.GenAuthClientOption())
		if err != nil {
			klog.Errorln("TracingQueryOperations error:", err)
		}
		It("正确的错误提示", func() {
			Expect(len(resp.Payload.Nodes)).Should(BeNumerically(">", 0))
		})
	})

	Context("9. lookback > endTime", func() {
		params := defaultParams()
		resp, err := client.Client.Tracing.TracingQueryOperations(params, tools.GenAuthClientOption())
		if err != nil {
			klog.Errorln("TracingQueryOperations error:", err)
		}
		It("正确的错误提示", func() {
			Expect(len(resp.Payload.Nodes)).Should(BeNumerically(">", 0))
		})
	})

	Context("10. step > lookback", func() {
		params := defaultParams()
		resp, err := client.Client.Tracing.TracingQueryOperations(params, tools.GenAuthClientOption())
		if err != nil {
			klog.Errorln("TracingQueryOperations error:", err)
		}
		It("正确的错误提示", func() {
			Expect(len(resp.Payload.Nodes)).Should(BeNumerically(">", 0))
		})
	})

})
