package containers

import (
	"os/exec"
	"bytes"
	"log"
	"strconv"
)

func InstallImage(path string) (string, error) {
	cmd := exec.Command("docker", "load", "--input", path)
	cmdOutput := &bytes.Buffer{}
	cmd.Stdout = cmdOutput
	err := cmd.Run()
	if err != nil {
		return "", err
	}
	log.Println(string(cmdOutput.Bytes()))

	return string(cmdOutput.Bytes()), nil
}

func StartContainer(imageName string, internalport int, externalport int, props map[string]interface{}) (containerId string, err error) {

	params := []string{}
	for k, v := range props {
		if value, ok := v.(string); ok {
			params = append(params, "-e", string(k)+"="+string(value))
		}
	}

	hostVolume := "/tmp"
	containerVolume := "./data"

	req := []string{}
	req = append(req, "run",
		"-d",
		"-p", strconv.Itoa(externalport)+":"+strconv.Itoa(internalport),
		"-v", hostVolume+":"+containerVolume,
		imageName,
	)
	//req = append(req, params...)
	//req = append(req, imageName)

	cmd := exec.Command("docker", req...)

	cmdOutput := &bytes.Buffer{}
	cmd.Stdout = cmdOutput
	err = cmd.Run()
	if err != nil {
		log.Println(err.Error())
		return "", err
	}

	return string(cmdOutput.Bytes()), nil
}

func StopContainer(containerId string) error {
	cmd := exec.Command("docker", "stop", containerId)
	cmd = exec.Command("docker", "rm", "-f", containerId)
	cmdOutput := &bytes.Buffer{}
	cmd.Stdout = cmdOutput
	err := cmd.Run()
	if err != nil {
		return err
	}
	log.Println(string(cmdOutput.Bytes()))

	return nil
}
