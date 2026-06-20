# K8s Metrics Viewer

一个用 Go 编写的命令行工具，调用 Kubernetes Metrics API，统计并展示每个 Pod 的 CPU 和内存使用量，效果类似 `kubectl top pods`。

## 功能

- 调用 Metrics API（依赖集群已安装 metrics-server）
- 统计每个 Pod 下所有容器的 CPU（毫核）和内存（MiB）总用量
- 以对齐的表格形式打印结果

## 技术栈

- Go
- [client-go](https://github.com/kubernetes/client-go)
- [k8s.io/metrics](https://github.com/kubernetes/metrics)：Kubernetes Metrics API 客户端

## 前置条件

集群中需要已部署 `metrics-server`。可通过以下命令确认：

\`\`\`bash
kubectl top pods -A
\`\`\`

如果该命令能正常返回数据，说明环境满足要求。

## 运行方式

确保本地已配置好可用的 kubeconfig（默认读取 `~/.kube/config`），然后：

\`\`\`bash
go mod tidy
go run main.go
\`\`\`

## 示例输出

\`\`\`
NAMESPACE    POD                                      CPU          MEMORY
kube-system  coredns-8db54c48d-bcd9h                  2m           13Mi
kube-system  metrics-server-786d997795-f4ffd          4m           22Mi
...
\`\`\`

## 开发与测试环境

本项目在 WSL2 + K3s 单节点集群上开发和验证。
