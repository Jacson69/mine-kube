package consts

// cluster api error code
const (
	ERRGETCLUSTERS    = 10001
	ERRGETCLUSTER     = 10002
	ERRCREATECLUSTER  = 10003
	ERRGETNODEMETRICS = 10004
)

// deployment api error code
const (
	ErrorGetDeployments          = 10100
	ErrorCreateDeployment        = 10101
	ErrorScaleReplicasDeployment = 10102
	ErrorDeleteDeployment        = 10103
	ErrorGetDeployment           = 10104
)

const (
	ErrorGetPodList = 10200
	ErrorCreatePod  = 10201
)
