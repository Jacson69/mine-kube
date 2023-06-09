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
	ErrorDeleteDeployment        = 10102
	ErrorGetDeployment           = 10103
	ErrorScaleReplicasDeployment = 10104
)

const (
	ErrorGetPodList = 10200
	ErrorCreatePod  = 10201
	ErrorDeletePod  = 10202
	ErrorGetPod     = 10203
)

const (
	ErrorGetNamespaceList = 10300
	ErrorCreateNamespace  = 10301
	ErrorDeleteNamespace  = 10302
	ErrorGetNamespace     = 10303
)

const (
	ErrorGetConfigmapList = 10400
	ErrorCreateConfigmap  = 10401
	ErrorDeleteConfigmap  = 10402
	ErrorGetConfigmap     = 10403
)

const (
	ErrorGetServiceList = 10500
	ErrorCreateService  = 10501
	ErrorDeleteService  = 10502
	ErrorGetService     = 10503
)
