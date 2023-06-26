package core

import (
	coreService "mine-kube/apis/core"
	"mine-kube/pkg/util/logger"

	"github.com/gin-gonic/gin"
)

func RegisterConfigMapRouter(v1alpha1 *gin.RouterGroup) {
	configmapApi, err := coreService.NewConfigMap()
	if err != nil {
		logger.Error(err)
		return
	}
	v1alpha1.GET("/clusters/:clusterID/namespaces/:namespace/configmaps", configmapApi.GetConfigmapList)
	v1alpha1.POST("/clusters/:clusterID/namespaces/:namespace/configmaps", configmapApi.ConfigmapAction)
	v1alpha1.DELETE("/clusters/:clusterID/namespaces/:namespace/configmaps/:configmapID", configmapApi.DeleteConfigmap)
	v1alpha1.GET("/clusters/:clusterID/namespaces/:namespace/configmaps/:configmapID", configmapApi.GetConfigmap)
}
