package core

import (
	"context"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	coreModels "mine-kube/models/core"
	baseService "mine-kube/pkg/service"
	"mine-kube/pkg/service/cluster"
	"mine-kube/pkg/util"
)

type configmapService struct {
	baseService.BaseInterface
	ctx context.Context
	cs  cluster.Interface
}

type ConfigmapInterface interface {
	GetConfigmaps(clusterID string, namespace string, opts ...baseService.OpOption) ([]v1.ConfigMap, *int64, error)
	GetConfigmap(clusterID string, namespace string, configmapID string) (*v1.ConfigMap, error)
	DeleteConfigmap(clusterID string, namespace string, configmapID string) error
	CreateConfigmap(clusterID string, namespace string, configmapPost coreModels.ConfigmapPost) (*v1.ConfigMap, error)
}

func NewConfigmap() (ConfigmapInterface, error) {
	return newConfigmap()
}

func newConfigmap() (*configmapService, error) {
	bs, err := baseService.NewBase()
	if err != nil {
		return nil, err
	}
	clusterService, err := cluster.NewClusterService()
	if err != nil {
		return nil, err
	}
	return &configmapService{
		ctx:           context.Background(),
		BaseInterface: bs,
		cs:            clusterService,
	}, nil
}

func (cs *configmapService) GetConfigmaps(clusterID string, namespace string, opts ...baseService.OpOption) ([]v1.ConfigMap, *int64, error) {
	op := baseService.OpGet(opts...)
	clientSet, err := cs.cs.GetKubernetesClientSet(clusterID)
	if err != nil {
		return nil, nil, err
	}
	list, err := clientSet.Kubernetes().CoreV1().ConfigMaps(namespace).List(cs.ctx, metav1.ListOptions{})
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

func (cs *configmapService) GetConfigmap(clusterID string, namespace string, configmapID string) (*v1.ConfigMap, error) {
	clientSet, err := cs.cs.GetKubernetesClientSet(clusterID)
	if err != nil {
		return nil, err
	}
	return clientSet.Kubernetes().CoreV1().ConfigMaps(namespace).Get(cs.ctx, configmapID, metav1.GetOptions{})
}

func (cs *configmapService) DeleteConfigmap(clusterID string, namespace string, configmapID string) error {
	clientSet, err := cs.cs.GetKubernetesClientSet(clusterID)
	if err != nil {
		return err
	}
	err = clientSet.Kubernetes().CoreV1().ConfigMaps(namespace).Delete(cs.ctx, configmapID, metav1.DeleteOptions{})
	if err != nil {
		return err
	}
	return nil
}

func (cs *configmapService) CreateConfigmap(clusterID string, namespace string, configmapPost coreModels.ConfigmapPost) (*v1.ConfigMap, error) {
	clientSet, err := cs.cs.GetKubernetesClientSet(clusterID)
	if err != nil {
		return nil, err
	}
	createConfigmap := &v1.ConfigMap{
		ObjectMeta: metav1.ObjectMeta{
			Name: configmapPost.Name,
		},
		Data: configmapPost.Data,
	}
	configmap, err := clientSet.Kubernetes().CoreV1().ConfigMaps(namespace).Create(cs.ctx, createConfigmap, metav1.CreateOptions{})
	if err != nil {
		return nil, err
	}
	return configmap, nil
}
