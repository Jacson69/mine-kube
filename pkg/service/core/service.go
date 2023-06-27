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

type serviceService struct {
	baseService.BaseInterface
	ctx context.Context
	cs  cluster.Interface
}

type ServiceInterface interface {
	GetServices(clusterID string, namespace string, opts ...baseService.OpOption) ([]v1.Service, *int64, error)
	GetService(clusterID string, namespace string, serviceID string) (*v1.Service, error)
	DeleteService(clusterID string, namespace string, serviceID string) error
	CreateService(clusterID string, namespace string, servicePost coreModels.ServicePost) (*v1.Service, error)
}

func NewService() (ServiceInterface, error) {
	return newService()
}

func newService() (*serviceService, error) {
	bs, err := baseService.NewBase()
	if err != nil {
		return nil, err
	}
	clusterService, err := cluster.NewClusterService()
	if err != nil {
		return nil, err
	}
	return &serviceService{
		ctx:           context.Background(),
		BaseInterface: bs,
		cs:            clusterService,
	}, nil
}

func (ss *serviceService) GetServices(clusterID string, namespace string, opts ...baseService.OpOption) ([]v1.Service, *int64, error) {
	op := baseService.OpGet(opts...)
	clientSet, err := ss.cs.GetKubernetesClientSet(clusterID)
	if err != nil {
		return nil, nil, err
	}
	list, err := clientSet.Kubernetes().CoreV1().Services(namespace).List(ss.ctx, metav1.ListOptions{})
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

func (ss *serviceService) GetService(clusterID string, namespace string, serviceID string) (*v1.Service, error) {
	clientSet, err := ss.cs.GetKubernetesClientSet(clusterID)
	if err != nil {
		return nil, err
	}
	return clientSet.Kubernetes().CoreV1().Services(namespace).Get(ss.ctx, serviceID, metav1.GetOptions{})
}

func (ss *serviceService) DeleteService(clusterID string, namespace string, serviceID string) error {
	clientSet, err := ss.cs.GetKubernetesClientSet(clusterID)
	if err != nil {
		return err
	}
	err = clientSet.Kubernetes().CoreV1().Services(namespace).Delete(ss.ctx, serviceID, metav1.DeleteOptions{})
	if err != nil {
		return err
	}
	return nil
}

func (ss *serviceService) CreateService(clusterID string, namespace string, servicePost coreModels.ServicePost) (*v1.Service, error) {
	clientSet, err := ss.cs.GetKubernetesClientSet(clusterID)
	if err != nil {
		return nil, err
	}
	createService := &v1.Service{
		ObjectMeta: servicePost.ObjectMeta,
		Spec:       servicePost.Spec,
	}
	service, err := clientSet.Kubernetes().CoreV1().Services(namespace).Create(ss.ctx, createService, metav1.CreateOptions{})
	if err != nil {
		return nil, err
	}
	return service, nil
}
