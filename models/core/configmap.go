package core

import v1 "k8s.io/api/core/v1"

type ConfigmapPost struct {
	v1.ConfigMap `json:",inline"`
}
