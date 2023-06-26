package core

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"mine-kube/apis"
	"mine-kube/models/core"
	"mine-kube/pkg/consts"
	"mine-kube/pkg/service"
	namespaceService "mine-kube/pkg/service/core"
)

type NameSpace struct {
	apis.Base
	ns namespaceService.NamespaceInterface
}

func NewNamespace() (*NameSpace, error) {
	tmp, err := namespaceService.NewNamespace()
	if err != nil {
		return nil, err
	}
	return &NameSpace{
		ns: tmp,
	}, nil
}

func (dc *NameSpace) GetNamespaceList(c *gin.Context) {
	pagination := dc.GetPagination(c)
	clusterID := c.Param("clusterID")
	namespaces, count, err := dc.ns.GetNamespaces(clusterID, service.WithPagination(pagination))
	if err != nil {
		dc.Error(c, consts.ErrorGetNamespaceList, err, "")
	}
	dc.PageOK(c, namespaces, count, pagination, "")
}

func (dc *NameSpace) GetNamespace(c *gin.Context) {
	clusterID := c.Param("clusterID")
	namespaceID := c.Param("namespaceID")
	namespace, err := dc.ns.GetNamespace(clusterID, namespaceID)
	if err != nil {
		dc.Error(c, consts.ErrorGetNamespace, err, "")
		return
	}
	dc.OK(c, namespace, fmt.Sprintf("get namespace %s success", namespaceID))
}

func (dc *NameSpace) DeleteNamespace(c *gin.Context) {
	clusterID := c.Param("clusterID")
	namespaceID := c.Param("namespaceID")
	err := dc.ns.DeleteNamespace(clusterID, namespaceID)
	if err != nil {
		dc.Error(c, consts.ErrorDeleteNamespace, err, "")
		return
	}
	dc.OK(c, nil, fmt.Sprintf("delete namespace %s success", namespaceID))
}

func (dc *NameSpace) NamespaceAction(c *gin.Context) {
	clusterID := c.Param("clusterID")
	namespacePost := core.NamespacePost{}
	if err := c.ShouldBindJSON(&namespacePost); err != nil {
		dc.Error(c, consts.ErrorCreateNamespace, err, "")
		return
	}
	namespace, err := dc.ns.CreateNamespace(clusterID, namespacePost)
	if err != nil {
		dc.Error(c, consts.ErrorCreateNamespace, err, "")
		return
	}
	dc.OK(c, namespace, "create namespace success")
}
