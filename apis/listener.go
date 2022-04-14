// Package apis Copyright (2022, ) Institute of Software, Chinese Academy of Sciences
package apis

/**
 * author: wuheng@iscas.ac.cn
 * since:  0.1
 */

type Listener interface {
	listen(t *Target)
}

type Target struct {
	// client *Client, substructure should have it
}
