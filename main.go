package main

import (
	"github.com/kubesys/client-go/pkg/kubesys"
	"sched-api/apis"
	"sched-api/kube"
)

func main() {
	url := "https://39.100.71.73:6443"
	tok := ""
	client := kubesys.NewKubernetesClient(url, tok)
	client.Init()

	target := apis.NewTarget(client)
	listener := new(kube.KubeListener)
	listener.Target = target
	listener.Type = "Pod"

	listener.DoListening()

}
