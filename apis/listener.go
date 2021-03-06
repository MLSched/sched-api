// Package apis Copyright (2022, ) Institute of Software, Chinese Academy of Sciences
package apis

/**
 * author: wuheng@iscas.ac.cn
 * since:  0.1
 */

type Listener struct {
	Target *Target
	Type   string
}

type ListenerActions interface {
	DoListening()
}
