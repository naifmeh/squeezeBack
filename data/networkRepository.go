package data

import (
	"squeezecnn/models"
	"os/exec"
	"errors"
)

func TrainNetwork(network models.Network) error {
	if network.Train == true {
		cmd := "./scripts/trainNet.sh"
		if err := exec.Command("/bin/sh", cmd).Run(); err == nil {
			cmd = "./scripts/executeProducer.sh"

			if err = exec.Command("/bin/sh", cmd).Run(); err != nil {
				return errors.New("could not execute producer")
			}
		} else {
			return err
		}
	}

	return nil
}

func RemovePictures() error {
	cmd := "./scripts/removePics.sh"
	if err := exec.Command("/bin/sh", cmd).Run(); err != nil {
		return errors.New("could not remove pictures")
	}
	return nil
}