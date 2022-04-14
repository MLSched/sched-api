// Package kube Package Copyright (2022, ) Institute of Software, Chinese Academy of Sciences
package kube

import (
	"github.com/kubesys/client-go/pkg/kubesys"
	apis "sched-api/apis"
)

type KubeTarget struct {
	apis.Target
	client *kubesys.KubernetesClient
}
