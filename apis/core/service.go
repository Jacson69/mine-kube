package core

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"mine-kube/apis"
	"mine-kube/models/core"
	"mine-kube/pkg/consts"
	"mine-kube/pkg/service"
	serviceService "mine-kube/pkg/service/core"
)

type Service struct {
	apis.Base
	svc serviceService.ServiceInterface
}

func NewService() (*Service, error) {
	tmp, err := serviceService.NewService()
	if err != nil {
		return nil, err
	}
	return &Service{
		svc: tmp,
	}, nil
}

func (dc *Service) GetServiceList(c *gin.Context) {
	pagination := dc.GetPagination(c)
	clusterID := c.Param("clusterID")
	namespace := c.Param("namespace")
	services, count, err := dc.svc.GetServices(clusterID, namespace, service.WithPagination(pagination))
	if err != nil {
		dc.Error(c, consts.ErrorGetServiceList, err, "")
		return
	}
	dc.PageOK(c, services, count, pagination, "")
}

func (dc *Service) GetService(c *gin.Context) {
	clusterID := c.Param("clusterID")
	namespace := c.Param("namespace")
	serviceID := c.Param("serviceID")
	service, err := dc.svc.GetService(clusterID, namespace, serviceID)
	if err != nil {
		dc.Error(c, consts.ErrorGetService, err, "")
		return
	}
	dc.OK(c, service, fmt.Sprintf("get service %s success", serviceID))
}

func (dc *Service) DeleteService(c *gin.Context) {
	clusterID := c.Param("clusterID")
	namespace := c.Param("namespace")
	serviceID := c.Param("serviceID")
	err := dc.svc.DeleteService(clusterID, namespace, serviceID)
	if err != nil {
		dc.Error(c, consts.ErrorDeleteService, err, "")
		return
	}
	dc.OK(c, nil, fmt.Sprintf("delete service %s success", serviceID))
}

func (dc *Service) ServiceAction(c *gin.Context) {
	clusterID := c.Param("clusterID")
	namespace := c.Param("namespace")
	servicePost := core.ServicePost{}
	if err := c.ShouldBindJSON(&servicePost); err != nil {
		dc.Error(c, consts.ERRCREATECLUSTER, err, "")
		return
	}
	service, err := dc.svc.CreateService(clusterID, namespace, servicePost)
	if err != nil {
		dc.Error(c, consts.ErrorCreateService, err, "")
		return
	}
	dc.OK(c, service, "create service success")
}
