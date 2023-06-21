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
	v1alpha1.POST("/namespaces/:namespace/pod", podApi.PodAction)
	v1alpha1.DELETE("/namespaces/:namespace/pods/:podID", podApi.DeletePod)
	v1alpha1.GET("/namespaces/:namespace/pods/:podID", podApi.GetPod)
}
