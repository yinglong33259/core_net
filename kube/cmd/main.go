package main

import (
	"encoding/json"
	"fmt"
	"github.com/containernetworking/cni/pkg/skel"
	"github.com/containernetworking/cni/pkg/types"
	"github.com/containernetworking/cni/pkg/types/current"
	"github.com/containernetworking/cni/pkg/version"
	log "github.com/sirupsen/logrus"
)

type PluginConf struct {
	types.NetConf
}

func parseConfig(stdin []byte) (*PluginConf, error) {
	conf := PluginConf{}

	if err := json.Unmarshal(stdin, &conf); err != nil {
		return nil, fmt.Errorf("failed to parse network configuration: %v", err)
	}

	// Parse previous result. Remove this if your plugin is not chained.
	if conf.RawPrevResult != nil {
		resultBytes, err := json.Marshal(conf.RawPrevResult)
		if err != nil {
			return nil, fmt.Errorf("could not serialize prevResult: %v", err)
		}
		res, err := version.NewResult(conf.CNIVersion, resultBytes)
		if err != nil {
			return nil, fmt.Errorf("could not parse prevResult: %v", err)
		}
		conf.RawPrevResult = nil
		conf.PrevResult, err = current.NewResultFromResult(res)
		if err != nil {
			return nil, fmt.Errorf("could not convert result to current version: %v", err)
		}
	}
	// End previous result parsing

	// Do any validation here
	//if conf.AnotherAwesomeArg == "" {
	//	return nil, fmt.Errorf("anotherAwesomeArg must be specified")
	//}

	return &conf, nil
}

func main() {
	skel.PluginMain(cmdAdd, cmdCheck, cmdDel, version.All, "corenet")
}

func cmdAdd(args *skel.CmdArgs) error {
	log.WithField("reqParas", args).Info("Got a net add request")
	return nil
}

func cmdCheck(args *skel.CmdArgs) error {
	log.WithField("reqParas", args).Info("Got a net check request")
	return nil
}

func cmdDel(args *skel.CmdArgs) error {
	log.WithField("reqParas", args).Info("Got a net delete request")
	return nil
}
