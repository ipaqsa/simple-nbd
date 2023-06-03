package server

import "sync"

type Server struct {
	Agent              *Agent
	MetasExport        []ExportMeta
	Port               string
	ReadOnly           bool
	MinimumBlockSize   int64
	PreferredBlockSize int64
	MaximumBlockSize   int64
}

type ExportMeta struct {
	Path        string
	Name        string
	Description string
	In          bool
}

type Agent struct {
	Storage *Storage
}

type Storage struct {
	Localy map[string]bool
	mtx    sync.Mutex
}
