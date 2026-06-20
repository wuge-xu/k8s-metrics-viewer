package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	metricsv "k8s.io/metrics/pkg/client/clientset/versioned"

	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	kubeconfig := filepath.Join(os.Getenv("HOME"), ".kube", "config")

	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		fmt.Println("构建配置失败:", err)
		os.Exit(1)
	}

	metricsClient, err := metricsv.NewForConfig(config)
	if err != nil {
		fmt.Println("创建 metrics 客户端失败:", err)
		os.Exit(1)
	}

	podMetricsList, err := metricsClient.MetricsV1beta1().PodMetricses("").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		fmt.Println("获取 metrics 数据失败:", err)
		os.Exit(1)
	}

	fmt.Printf("%-12s %-40s %-12s %-12s\n", "NAMESPACE", "POD", "CPU", "MEMORY")

	for _, podMetrics := range podMetricsList.Items {
		var totalCPU, totalMemory int64

		for _, container := range podMetrics.Containers {
			totalCPU += container.Usage.Cpu().MilliValue()
			totalMemory += container.Usage.Memory().Value() / (1024 * 1024)
		}

		fmt.Printf("%-12s %-40s %-12s %-12s\n",
			podMetrics.Namespace,
			podMetrics.Name,
			fmt.Sprintf("%dm", totalCPU),
			fmt.Sprintf("%dMi", totalMemory),
		)
	}
}
