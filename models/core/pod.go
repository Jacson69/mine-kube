package core

import "k8s.io/api/core/v1"

type PodPost struct {
	v1.Pod `json:",inline"`
}
