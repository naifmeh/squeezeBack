package data

import (
	"squeezecnn/models"
	"os/exec"
	"errors"
	"os"
	"path/filepath"
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

func RemoveContents(dir string) error {
	d, err := os.Open(dir)
	if err != nil {
		return err
	}
	defer d.Close()
	names, err := d.Readdirnames(-1)
	if err != nil {
		return err
	}
	for _, name := range names {
		err = os.RemoveAll(filepath.Join(dir, name))
		if err != nil {
			return err
		}
	}
	return nil
}
