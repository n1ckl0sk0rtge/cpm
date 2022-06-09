package cruntime

import (
	"fmt"
	"github.com/n1ckl0sk0rtge/cpm/config"
	"net"
)

const (
	Docker string = "docker"
	Podman string = "podman"
)

func Available() error {
	socket := config.Instance.GetString(config.Socket)
	connection, err := net.Dial("unix", socket)
	defer func(connection net.Conn) {
		_ = connection.Close()
	}(connection)

	if err != nil {
		return fmt.Errorf("container cruntime is not availabe. Please check and provide a valid cruntime")
	}

	return nil
}
