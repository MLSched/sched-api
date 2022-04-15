// Package kube Package Copyright (2022, ) Institute of Software, Chinese Academy of Sciences
package kube

import (
	"fmt"
	"github.com/kubesys/client-go/pkg/kubesys"
	"sched-api/apis"
)

/**
 * author: wuheng@iscas.ac.cn
 * since:  0.1
 */

type KubeListener struct {
	apis.Listener
}

func (listener *KubeListener) DoListening() {
	client := listener.Target.Client.(*kubesys.KubernetesClient)
	res, _ := client.ListResources(listener.Type, "")
	fmt.Print(string(res))
}
