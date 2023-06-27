package core

import (
	coreService "mine-kube/apis/core"
	"mine-kube/pkg/util/logger"

	"github.com/gin-gonic/gin"
)

func RegisterServiceRouter(v1alpha1 *gin.RouterGroup) {
	serviceApi, err := coreService.NewService()
	if err != nil {
		logger.Error(err)
		return
	}
	v1alpha1.GET("/clusters/:clusterID/namespaces/:namespace/services", serviceApi.GetServiceList)
	v1alpha1.POST("/clusters/:clusterID/namespaces/:namespace/services", serviceApi.ServiceAction)
	v1alpha1.DELETE("/clusters/:clusterID/namespaces/:namespace/services/:serviceID", serviceApi.DeleteService)
	v1alpha1.GET("/clusters/:clusterID/namespaces/:namespace/services/:serviceID", serviceApi.GetService)
}
