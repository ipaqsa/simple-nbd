package main

import (
	"github.com/ipaqsa/simple-nbd/pkg"
	"github.com/ipaqsa/simple-nbd/pkg/configurator"
	"os"
)

func init() {
	//comment if you don`t want to use create with config
	err := configurator.InitConfiguration(&pkg.Config, "0.0.1")
	if err != nil {
		println(err.Error())
		os.Exit(-1)
	}
	configurator.InitInfo(pkg.Config.Port)
}

func main() {
	//meta := server.CreateMetaFromDir("disks")
	//meta := createMeta()
	//s := server.NewServer(":7002", false, 1, 4096, 0xffffffff, meta)

	for _, i := range pkg.Config.Instances {
		println(i.Name, i.Path)
	}
	//s := server.NewServerWithConfig()
	//if err := s.Serve(); err != nil {
	//	log.Fatal(err)
	//}
}

//func createMeta() []server.ExportMeta {
//	var data []server.ExportMeta
//	m := *server.NewMeta("disk", "For test", "disk.img")
//
//	data = append(data, m)
//	return data
//}
