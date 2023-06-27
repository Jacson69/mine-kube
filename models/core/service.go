package core

import v1 "k8s.io/api/core/v1"

type ServicePost struct {
	v1.Service `json:",inline"`
}
