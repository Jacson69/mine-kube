package cluster

import (
	v1 "k8s.io/api/core/v1"
	"mine-kube/pkg/api/cluster/v1alpha1"
)

type Cluster struct {
	v1alpha1.Cluster
	HealthStatus string       `json:"health_status"`
	Version      string       `json:"version"`
	NodeList     *v1.NodeList `json:"node_list,omitempty"`
}

type Post struct {
	DisplayName   string `json:"displayname"`
	KubeConfig    string `json:"kubeconfig"`
	PrometheusURL string `json:"prometheusurl"`
}
