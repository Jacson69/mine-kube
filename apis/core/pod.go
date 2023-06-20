package core

import (
	"github.com/gin-gonic/gin"
	"mine-kube/apis"
	"mine-kube/pkg/consts"
	"mine-kube/pkg/service"
	podService "mine-kube/pkg/service/core"
)

type Pod struct {
	apis.Base
	ds podService.PodInterface
	//podActionFunc map[string]func(dc *Pod, c *gin.Context)
}

func NewPod() (*Pod, error) {
	tmp, err := podService.NewPod()
	if err != nil {
		return nil, err
	}
	return &Pod{
		ds: tmp,
		//podActionFunc: newPodActionFunc(),
	}, nil
}

//func newPodActionFunc() map[string]func(dc *Pod, c *gin.Context) {
//	return map[string]func(dc *Pod, c *gin.Context){}
//}

func (p *Pod) GetPodList(c *gin.Context) {
	pagination := p.GetPagination(c)
	namespace := c.Param("namespace")
	list, count, err := p.ds.GetPodList(namespace, service.WithPagination(pagination))
	if err != nil {
		p.Error(c, consts.EooroGetPodList, err, "")
		return
	}
	p.PageOK(c, list, count, pagination, "")
}
