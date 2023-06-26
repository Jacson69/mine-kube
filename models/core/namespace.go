package core

import (
	v1 "k8s.io/api/core/v1"
)

type NamespacePost struct {
	v1.Namespace `json:",inline"`
}
