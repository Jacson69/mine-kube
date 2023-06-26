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

type namespaceService struct {
	baseService.BaseInterface
	ctx context.Context
	cs  cluster.Interface
}

type NamespaceInterface interface {
	GetNamespaces(clusterID string, opts ...baseService.OpOption) ([]v1.Namespace, *int64, error)
	GetNamespace(clusterID string, namespaceID string) (*v1.Namespace, error)
	DeleteNamespace(clusterID string, namespaceID string) error
	CreateNamespace(clusterID string, namespacePost coreModels.NamespacePost) (*v1.Namespace, error)
}

func NewNamespace() (NamespaceInterface, error) {
	return newNamespace()
}

func newNamespace() (*namespaceService, error) {
	bs, err := baseService.NewBase()
	if err != nil {
		return nil, err
	}
	clusterService, err := cluster.NewClusterService()
	if err != nil {
		return nil, err
	}
	return &namespaceService{
		ctx:           context.Background(),
		BaseInterface: bs,
		cs:            clusterService,
	}, nil
}

func (ns *namespaceService) GetNamespaces(clusterID string, opts ...baseService.OpOption) ([]v1.Namespace, *int64, error) {
	op := baseService.OpGet(opts...)
	clientSet, err := ns.cs.GetKubernetesClientSet(clusterID)
	if err != nil {
		return nil, nil, err
	}
	list, err := clientSet.Kubernetes().CoreV1().Namespaces().List(ns.ctx, metav1.ListOptions{})
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

func (ns *namespaceService) GetNamespace(clusterID string, namespaceID string) (*v1.Namespace, error) {
	clientSet, err := ns.cs.GetKubernetesClientSet(clusterID)
	if err != nil {
		return nil, err
	}
	return clientSet.Kubernetes().CoreV1().Namespaces().Get(ns.ctx, namespaceID, metav1.GetOptions{})
}

func (ns *namespaceService) DeleteNamespace(clusterID string, namespaceID string) error {
	clientSet, err := ns.cs.GetKubernetesClientSet(clusterID)
	if err != nil {
		return err
	}
	err = clientSet.Kubernetes().CoreV1().Namespaces().Delete(ns.ctx, namespaceID, metav1.DeleteOptions{})
	if err != nil {
		return err
	}
	return nil
}
func (ns *namespaceService) CreateNamespace(clusterID string, namespacePost coreModels.NamespacePost) (*v1.Namespace, error) {
	clientSet, err := ns.cs.GetKubernetesClientSet(clusterID)
	if err != nil {
		return nil, err
	}

	createNamespace := &v1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: namespacePost.Name,
		},
	}
	namespace, err := clientSet.Kubernetes().CoreV1().Namespaces().Create(ns.ctx, createNamespace, metav1.CreateOptions{})
	if err != nil {
		return nil, err
	}
	return namespace, nil
}
