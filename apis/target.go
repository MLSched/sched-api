// Package apis Copyright (2022, ) Institute of Software, Chinese Academy of Sciences
package apis

/**
 * author: wuheng@iscas.ac.cn
 * since:  0.1
 */

type Target struct {
	Client any
}

func NewTarget(client any) *Target {
	kubeTarget := new(Target)
	kubeTarget.Client = client
	return kubeTarget
}
