package core

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"mine-kube/apis"
	"mine-kube/models/core"
	"mine-kube/pkg/consts"
	"mine-kube/pkg/service"
	configmapService "mine-kube/pkg/service/core"
)

type ConfigMap struct {
	apis.Base
	cm configmapService.ConfigmapInterface
}

func NewConfigMap() (*ConfigMap, error) {
	tmp, err := configmapService.NewConfigmap()
	if err != nil {
		return nil, err
	}
	return &ConfigMap{
		cm: tmp,
	}, nil
}

func (dc *ConfigMap) GetConfigmapList(c *gin.Context) {
	pagination := dc.GetPagination(c)
	clusterID := c.Param("clusterID")
	namespace := c.Param("namespace")

	configmaps, count, err := dc.cm.GetConfigmaps(clusterID, namespace, service.WithPagination(pagination))
	if err != nil {
		dc.Error(c, consts.ErrorGetConfigmapList, err, "")
		return
	}
	dc.PageOK(c, configmaps, count, pagination, "")
}

func (dc *ConfigMap) GetConfigmap(c *gin.Context) {
	clusterID := c.Param("clusterID")
	namespace := c.Param("namespace")
	configmapID := c.Param("configmapID")
	configmap, err := dc.cm.GetConfigmap(clusterID, namespace, configmapID)
	if err != nil {
		dc.Error(c, consts.ErrorGetConfigmap, err, "")
		return
	}
	dc.OK(c, configmap, fmt.Sprintf("get configmap %s success", configmapID))
}

func (dc *ConfigMap) DeleteConfigmap(c *gin.Context) {
	clusterID := c.Param("clusterID")
	namespace := c.Param("namespace")
	configmapID := c.Param("configmapID")
	err := dc.cm.DeleteConfigmap(clusterID, namespace, configmapID)
	if err != nil {
		dc.Error(c, consts.ErrorDeleteConfigmap, err, "")
		return
	}
	dc.OK(c, nil, fmt.Sprintf("delete configmap %s success", configmapID))
}

func (dc *ConfigMap) ConfigmapAction(c *gin.Context) {
	clusterID := c.Param("clusterID")
	namespace := c.Param("namespace")
	configmapPost := core.ConfigmapPost{}
	if err := c.ShouldBindJSON(&configmapPost); err != nil {
		dc.Error(c, consts.ERRCREATECLUSTER, err, "")
		return
	}
	configmap, err := dc.cm.CreateConfigmap(clusterID, namespace, configmapPost)
	if err != nil {
		dc.Error(c, consts.ErrorCreateConfigmap, err, "")
		return
	}
	dc.OK(c, configmap, "create configmap success")
}
