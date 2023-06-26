package core

import (
	coreService "mine-kube/apis/core"
	"mine-kube/pkg/util/logger"

	"github.com/gin-gonic/gin"
)

func RegisterNameSpaceRouter(v1alpha1 *gin.RouterGroup) {
	namespaceApi, err := coreService.NewNamespace()
	if err != nil {
		logger.Error(err)
		return
	}
	v1alpha1.GET("/clusters/:clusterID/namespaces", namespaceApi.GetNamespaceList)
	v1alpha1.POST("/clusters/:clusterID/namespace", namespaceApi.NamespaceAction)
	v1alpha1.DELETE("/clusters/:clusterID/namespace/:namespaceID", namespaceApi.DeleteNamespace)
	v1alpha1.GET("/clusters/:clusterID/namespaces/:namespaceID", namespaceApi.GetNamespace)
}
