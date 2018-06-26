package data

import (
	"squeezecnn/models"
	"os/exec"
	"errors"
)

func TrainNetwork(network models.Network) error {
	if network.Train == true {
		cmd := "./trainNet.sh"
		if err := exec.Command("/bin/sh", cmd).Run(); err == nil {
			cmd = "./executeProducer.sh"

			if err = exec.Command("/bin/sh", cmd).Run(); err != nil {
				return errors.New("could not execute producer")
			}
		} else {
			return err
		}
	}

	return nil
}