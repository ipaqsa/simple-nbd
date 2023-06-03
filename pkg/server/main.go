package server

import (
	"github.com/ipaqsa/simple-nbd/pkg"
	"github.com/pojntfx/go-nbd/pkg/server"
	"log"
	"net"
)

func NewServerWithConfig() *Server {
	metaFromDir := CreateMetaFromDir(pkg.Config.Dir)
	meta := CreateMeta(pkg.Config.Instances)
	return NewServer(pkg.Config.Port, pkg.Config.ReadOnly, pkg.Config.MinimumBlockSize, pkg.Config.PreferredBlockSize,
		pkg.Config.MaximumBlockSize, metaMerge(metaFromDir, meta))
}

func NewServer(port string, readOnly bool, minimumBlockSize, preferredBlockSize, maximumBlockSize int64, metasExport []ExportMeta) *Server {
	s := &Server{Port: port, MetasExport: metasExport, ReadOnly: readOnly, MaximumBlockSize: maximumBlockSize,
		PreferredBlockSize: preferredBlockSize, MinimumBlockSize: minimumBlockSize}

	//s.Agent = NewAgent()

	return s
}

func (s *Server) Serve() error {
	res, err := net.ResolveTCPAddr("tcp", s.Port)
	if err != nil {
		return err
	}
	listener, err := net.ListenTCP("tcp", res)
	if err != nil {
		return err
	}
	defer listener.Close()

	exports, err := s.createExports()
	if err != nil {
		return err
	}

	clients := 0
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Could not accept connection, continuing:", err)
			continue
		}
		clients++
		if !s.Agent.In(GetAddr(conn.RemoteAddr().String())) {
			log.Printf("Permission denied: %s\n", conn.RemoteAddr().String())
			closeHelper(conn)
			clients--
			continue
		}

		log.Printf("%d clients connected", clients)

		go func() {
			defer func() {
				closeHelper(conn)
				clients--
			}()

			if err = server.Handle(
				conn,
				exports,
				&server.Options{
					ReadOnly:           s.ReadOnly,
					MinimumBlockSize:   uint32(s.MinimumBlockSize),
					PreferredBlockSize: uint32(s.PreferredBlockSize),
					MaximumBlockSize:   uint32(s.MaximumBlockSize),
				}); err != nil {
				panic(err)
			}
		}()
	}
}

func closeHelper(conn net.Conn) {
	_ = conn.Close()
	if err := recover(); err != nil {
		log.Printf("Client disconnected with error: %v", err)
	}
}
