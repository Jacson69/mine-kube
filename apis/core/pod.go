package core

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"mine-kube/apis"
	"mine-kube/models/core"
	"mine-kube/pkg/consts"
	"mine-kube/pkg/service"
	podService "mine-kube/pkg/service/core"
)

type Pod struct {
	apis.Base
	ds            podService.PodInterface
	podActionFunc map[string]func(dc *Pod, c *gin.Context)
}

func NewPod() (*Pod, error) {
	tmp, err := podService.NewPod()
	if err != nil {
		return nil, err
	}
	return &Pod{
		ds:            tmp,
		podActionFunc: newPodActionFunc(),
	}, nil
}

func newPodActionFunc() map[string]func(dc *Pod, c *gin.Context) {
	return map[string]func(dc *Pod, c *gin.Context){
		apis.CreateAction: createPod,
		apis.DryRunAction: dryRunPod,
	}
}

func createPod(p *Pod, c *gin.Context) {
	namespace := c.Param("namespace")
	podPost := core.PodPost{}
	if err := c.ShouldBindJSON(&podPost); err != nil {
		p.Error(c, consts.ERRCREATECLUSTER, err, "")
		return
	}
	pod, err := p.ds.CreatePod(namespace, podPost)
	if err != nil {
		p.Error(c, consts.ErrorCreatePod, err, "")
		return
	}
	p.OK(c, pod, "create pod success")
}

func dryRunPod(p *Pod, c *gin.Context) {
	namespace := c.Param("namespace")
	podPost := core.PodPost{}
	if err := c.ShouldBindJSON(&podPost); err != nil {
		p.Error(c, consts.ERRCREATECLUSTER, err, "")
		return
	}
	pod, err := p.ds.DryRunPod(namespace, podPost)
	if err != nil {
		p.Error(c, consts.ErrorCreatePod, err, "")
		return
	}
	p.OK(c, pod, "dry-run pod success")
}

func (p *Pod) GetPodList(c *gin.Context) {
	pagination := p.GetPagination(c)
	namespace := c.Param("namespace")
	list, count, err := p.ds.GetPodList(namespace, service.WithPagination(pagination))
	if err != nil {
		p.Error(c, consts.ErrorGetPodList, err, "")
		return
	}
	p.PageOK(c, list, count, pagination, "")
}

func (p *Pod) PodAction(c *gin.Context) {
	action := c.DefaultQuery("action", "create")
	actionFunc := p.podActionFunc[action]
	actionFunc(p, c)
}

func (p *Pod) DeletePod(c *gin.Context) {
	namespace := c.Param("namespace")
	podID := c.Param("podID")
	err := p.ds.DeletePod(namespace, podID)
	if err != nil {
		p.Error(c, consts.ErrorDeletePod, err, "")
		return
	}

	p.OK(c, nil, fmt.Sprintf("delete pod %s success", podID))
}

func (p *Pod) GetPod(c *gin.Context) {
	namespace := c.Param("namespace")
	podID := c.Param("podID")
	pod, err := p.ds.GetPod(namespace, podID)
	if err != nil {
		p.Error(c, consts.ErrorGetPod, err, "")
		return
	}
	p.OK(c, pod, fmt.Sprintf("get pod %s success", podID))
}
