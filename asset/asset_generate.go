package main

import (
	"log"
	"time"
	"github.com/shurcooL/vfsgen"
	"github.com/prometheus/alertmanager/asset"
	"github.com/prometheus/alertmanager/pkg/modtimevfs"
)

func main() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fs := modtimevfs.New(asset.Assets, time.Unix(1, 0))
	err := vfsgen.Generate(fs, vfsgen.Options{PackageName: "asset", BuildTags: "!dev", VariableName: "Assets"})
	if err != nil {
		log.Fatalln(err)
	}
}
