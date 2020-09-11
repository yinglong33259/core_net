package main

import (
	"github.com/containernetworking/cni/pkg/skel"
	"github.com/containernetworking/cni/pkg/version"
	log "github.com/sirupsen/logrus"
)

func main() {
	skel.PluginMain(cmdAdd, cmdCheck, cmdDel, version.All, "my core net plugin for kube!!!")
}

func cmdAdd(args *skel.CmdArgs) error {
	log.Info("Got a net add request")
	return nil
}

func cmdCheck(args *skel.CmdArgs) error {
	log.Info("Got a net check request")
	return nil
}

func cmdDel(args *skel.CmdArgs) error {
	log.Info("Got a net delete request")
	return nil
}
