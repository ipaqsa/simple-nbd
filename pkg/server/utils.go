package server

import (
	"github.com/ipaqsa/simple-nbd/pkg"
	"github.com/pojntfx/go-nbd/pkg/backend"
	"github.com/pojntfx/go-nbd/pkg/server"
	"log"
	"os"
	"path"
	"strings"
)

func (s *Server) createExport(in ExportMeta) (*server.Export, error) {
	var export server.Export
	var f *os.File
	var err error

	if s.ReadOnly {
		f, err = os.OpenFile(in.Path, os.O_RDONLY, 0644)
		if err != nil {
			return nil, err
		}
	} else {
		f, err = os.OpenFile(in.Path, os.O_RDWR, 0644)
		if err != nil {
			return nil, err
		}
	}

	b := backend.NewFileBackend(f)

	export.Name = in.Name
	export.Description = in.Description
	export.Backend = b

	return &export, nil
}

func (s *Server) createExports() ([]server.Export, error) {
	var exports []server.Export

	for _, e := range s.MetasExport {
		exp, err := s.createExport(e)
		if err != nil {
			return nil, err
		}
		exports = append(exports, *exp)
	}
	return exports, nil
}

func NewMeta(name, description, path string) *ExportMeta {
	return &ExportMeta{Name: name, Description: description, Path: path, In: false}
}

func GetAddr(address string) string {
	splits := strings.Split(address, ":")
	return splits[0]
}

func CreateMetaFromDir(dirpath string) []ExportMeta {
	dir, err := os.ReadDir(dirpath)
	if err != nil {
		log.Fatal(err)
	}
	var data []ExportMeta
	for _, f := range dir {
		name, _ := strings.CutSuffix(f.Name(), path.Ext(f.Name()))
		m := *NewMeta(name, "", path.Join(dirpath, f.Name()))
		data = append(data, m)
	}
	return data
}

func CreateMeta(instances []pkg.Instance) []ExportMeta {
	var data []ExportMeta
	for _, instance := range instances {
		m := *NewMeta(instance.Name, instance.Description, instance.Path)
		data = append(data, m)
	}
	return data
}

func metaMerge(m1 []ExportMeta, m2 []ExportMeta) []ExportMeta {
	var data []ExportMeta
	for idx, i := range m1 {
		contains(m2, i)
		data = append(data, m1[idx])
	}
	for idx, i := range m2 {
		if !i.In {
			data = append(data, m2[idx])
		}
	}
	return data
}

func contains(m []ExportMeta, i ExportMeta) {
	for idx, in := range m {
		if in.Name == i.Name || in.Path == i.Path {
			m[idx].In = true
			return
		}
	}
}
