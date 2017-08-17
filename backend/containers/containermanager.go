package containers

import (
	"os/exec"
	"os"
	"bytes"
	"fmt"
	"log"
)

func InstallImage(path string) error {
	cmd := exec.Command("docker", "load", "--input", path)
	cmdOutput := &bytes.Buffer{}
	cmd.Stdout = cmdOutput
	err := cmd.Run()
	if err != nil {
		return err
	}
	log.Println(string(cmdOutput.Bytes()))
	fmt.Println(string(cmdOutput.Bytes()))

	return nil
}

func StartContainer(pluginId string, imageName string, internalport int, externalport int, variables ...string) (containerId int, err error) {

	params := ""
	for _, variable := range variables {
		params += variable
	}

	containerName := "marcel_" + pluginId

	cmd := exec.Command("docker", "run", imageName, "-d", "-p", externalport+":"+internalport, "--name", containerName)
	cmdOutput := &bytes.Buffer{}
	cmd.Stdout = cmdOutput
	err = cmd.Run()
	if err != nil {
		os.Stderr.WriteString(err.Error())
	}
	fmt.Print(string(cmdOutput.Bytes()))

	return 1, nil
}

func StopContainer(containerId int) error {

	return nil
}
