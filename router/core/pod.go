package core

import (
	coreService "mine-kube/apis/core"
	"mine-kube/pkg/util/logger"

	"github.com/gin-gonic/gin"
)

func RegisterPodRouter(v1alpha1 *gin.RouterGroup) {
	podApi, err := coreService.NewPod()
	if err != nil {
		logger.Error(err)
		return
	}
	v1alpha1.GET("/namespaces/:namespace/podList", podApi.GetPodList)
	//v1alpha1.POST("/clusters/:clusterID/namespaces/:namespace/deployments", deploymentApi.DeploymentAction)
	//v1alpha1.DELETE("/clusters/:clusterID/namespaces/:namespace/deployments/:deploymentID", deploymentApi.DeleteDeployment)
}
