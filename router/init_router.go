package router

import (
	"fmt"
	"mine-kube/router/cluster"
	"mine-kube/router/core"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	return baseRouterV1()
}

func baseRouterV1() *gin.Engine {
	r := gin.New()
	addRouter(r.Group(fmt.Sprintf("/api/%s/%s", VERSION, SERVERNAME)))
	return r
}

func addRouter(v1alpha1 *gin.RouterGroup) {
	cluster.RegisterClusterRouter(v1alpha1)
	core.RegisterDeploymentRouter(v1alpha1)
	core.RegisterPodRouter(v1alpha1)
	core.RegisterNameSpaceRouter(v1alpha1)
	core.RegisterConfigMapRouter(v1alpha1)
	core.RegisterServiceRouter(v1alpha1)
}
