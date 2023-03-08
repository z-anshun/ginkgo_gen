package tracing

import (
	"fmt"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	tracing_client "insight.io/api/sdk/v1alpha1/client/tracing"
	"insight.io/test/tools"

	"k8s.io/klog/v2"
)

//go:generate ginkgo --focus=查询操作列表-/traces​/operations
var _ = Describe("查询操作列表-/traces​/operations", func() {
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

	Context("2. cluster不存在", func() {
		params := defaultParams()
		resp, err := client.Client.Tracing.TracingQueryOperations(params, tools.GenAuthClientOption())
		if err != nil {
			klog.Errorln("TracingQueryOperations error:", err)
		}
		It("无数据返回", func() {
			Expect(len(resp.Payload.Nodes)).Should(BeNumerically(">", 0))
		})
	})

	Context("3. namespace不存在", func() {
		params := defaultParams()
		resp, err := client.Client.Tracing.TracingQueryOperations(params, tools.GenAuthClientOption())
		if err != nil {
			klog.Errorln("TracingQueryOperations error:", err)
		}
		It("无数据返回", func() {
			Expect(len(resp.Payload.Nodes)).Should(BeNumerically(">", 0))
		})
	})

	Context("4. serviceName不存在", func() {
		params := defaultParams()
		resp, err := client.Client.Tracing.TracingQueryOperations(params, tools.GenAuthClientOption())
		if err != nil {
			klog.Errorln("TracingQueryOperations error:", err)
		}
		It("无数据返回", func() {
			Expect(len(resp.Payload.Nodes)).Should(BeNumerically(">", 0))
		})
	})

})
