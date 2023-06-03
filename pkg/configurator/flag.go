package configurator

import (
	"errors"
	"flag"
)

var PathToConfig *string

func initConfig() error {
	PathToConfig = flag.String("c", "", "Path to config file")
	flag.Parse()
	if *PathToConfig != "" {
		err := pathToConfIsValid(*PathToConfig)
		if err != nil {
			return err
		}
		return nil
	}
	return errors.New("enter path to config")
}
