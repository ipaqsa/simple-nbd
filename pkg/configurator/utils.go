package configurator

import (
	"errors"
	"net"
	"os"
	"path"
	"strings"
)

func pathToConfIsValid(path string) error {
	splits := strings.Split(path, ".")
	if len(splits) < 2 {
		return errors.New("path invalid")
	}
	format := splits[len(splits)-1]
	if format != "yml" {
		return errors.New("format is not yml")
	}
	return nil
}

func getOutboundIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return nil
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP
}

func splitPath(cpath string) (string, string, string) {
	ext := path.Ext(cpath)
	splits := strings.Split(cpath, "/")
	filename := splits[len(splits)-1]
	cpath = strings.TrimSuffix(cpath, filename)
	filename = strings.TrimSuffix(filename, ext)
	return cpath, filename, ext
}

func exists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}
