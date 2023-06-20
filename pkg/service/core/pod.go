package core

import (
	"context"
	v1 "k8s.io/api/core/v1"
	"k8s.io/client-go/tools/clientcmd"
	"mine-kube/pkg/client/k8s"
	baseService "mine-kube/pkg/service"
	"mine-kube/pkg/util"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type podService struct {
	client kubernetes.Interface
}
type PodInterface interface {
	GetPodList(namespace string, opts ...baseService.OpOption) ([]v1.Pod, *int64, error)
	//ScaleDeployment(clusterID string, namespace string, deploymentID string, replicas int32) error
	//GetDeployments(clusterID string, namespace string, opts ...baseService.OpOption) ([]appsv1.Deployment, *int64, error)
	//GetDeployment(clusterID string, namespace string, deploymentID string, opts ...baseService.OpOption) (*appsv1.Deployment, error)
	//CreateDeployment(clusterID string, namespace string, deploymentPost coreModels.DeploymentPost, opts ...baseService.OpOption) (*appsv1.Deployment, error)
	//DryRunDeployment(clusterID string, namespace string, deploymentPost coreModels.DeploymentPost, opts ...baseService.OpOption) (*appsv1.Deployment, error)
}

func NewPod() (PodInterface, error) {
	return newPod()
}
func newPod() (*podService, error) {
	//client, err := k8s.NewKubernetesClient(&k8s.KubernetesOptions{KubeConfig: "~/.kube/config"})
	client, err := k8s.NewKubernetesClient(&k8s.KubernetesOptions{KubeConfig: clientcmd.RecommendedHomeFile})
	if err != nil {
		return nil, err
	}
	return &podService{
		client: client.Kubernetes(),
	}, nil
}

func (ps *podService) GetPodList(namespace string, opts ...baseService.OpOption) ([]v1.Pod, *int64, error) {
	op := baseService.OpGet(opts...)
	list, err := ps.client.CoreV1().Pods(namespace).List(context.Background(), metav1.ListOptions{})
	if err != nil {
		return nil, nil, err
	}
	count := util.ConvertToInt64Ptr(len(list.Items))
	offset, end := baseService.CommonPaginate(list.Items,
		(op.Pagination.Page-1)*op.Pagination.PageSize,
		op.Pagination.PageSize)
	listItem := list.Items[offset:end]
	return listItem, count, nil
}
