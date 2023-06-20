package core

import (
	"context"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	coreModels "mine-kube/models/core"
	"mine-kube/pkg/client/k8s"
	baseService "mine-kube/pkg/service"
	"mine-kube/pkg/util"
)

type podService struct {
	client kubernetes.Interface
	ctx    context.Context
}
type PodInterface interface {
	GetPodList(namespace string, opts ...baseService.OpOption) ([]v1.Pod, *int64, error)
	//ScaleDeployment(clusterID string, namespace string, deploymentID string, replicas int32) error
	//GetDeployments(clusterID string, namespace string, opts ...baseService.OpOption) ([]appsv1.Deployment, *int64, error)
	//GetDeployment(clusterID string, namespace string, deploymentID string, opts ...baseService.OpOption) (*appsv1.Deployment, error)
	CreatePod(namespace string, podPost coreModels.PodPost) (*v1.Pod, error)
	//DryRunDeployment(clusterID string, namespace string, deploymentPost coreModels.DeploymentPost, opts ...baseService.OpOption) (*appsv1.Deployment, error)
}

func NewPod() (PodInterface, error) {
	return newPod()
}
func newPod() (*podService, error) {
	//client, err := k8s.NewKubernetesClient(&k8s.KubernetesOptions{KubeConfig: "~/.kube/config"})
	client, err := k8s.NewKubernetesClient(&k8s.KubernetesOptions{KubeConfig: "/root/.kube/config"})
	if err != nil {
		return nil, err
	}
	return &podService{
		client: client.Kubernetes(),
	}, nil
}

func (ps *podService) GetPodList(namespace string, opts ...baseService.OpOption) ([]v1.Pod, *int64, error) {
	op := baseService.OpGet(opts...)
	list, err := ps.client.CoreV1().Pods(namespace).List(ps.ctx, metav1.ListOptions{})
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

func (ps *podService) CreatePod(namespace string, podPost coreModels.PodPost) (*v1.Pod, error) {
	createPod := &v1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name: podPost.Name,
		},
		Spec: podPost.Spec,
	}
	pod, err := ps.client.CoreV1().Pods(namespace).Create(context.Background(), createPod, metav1.CreateOptions{})
	if err != nil {
		return nil, err
	}
	return pod, nil
}
