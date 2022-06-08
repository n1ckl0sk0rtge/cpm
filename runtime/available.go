package runtime

import (
	"fmt"
	"github.com/n1ckl0sk0rtge/cpm/config"
	"net"
)

func Available() error {
	socket := config.Instance.GetString(config.Socket)
	connection, err := net.Dial("unix", socket)
	defer func(connection net.Conn) {
		_ = connection.Close()
	}(connection)

	if err != nil {
		return fmt.Errorf("container runtime is not availabe. Please check and provide a valid runtime")
	}

	return nil
}
