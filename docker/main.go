package main

import (
	"bufio"
	"encoding/json"
	"io"
	"os"
	"os/exec"

	"github.com/charmbracelet/log"
)

type Volume struct {
	Availability string `json:"Availability"`
	Driver       string `json:"Driver"`
	Group        string `json:"Group"`
	Labels       string `json:"Labels"`
	Links        string `json:"Links"`
	Mountpoint   string `json:"Mountpoint"`
	Name         string `json:"Name"`
	Scope        string `json:"Scope"`
	Size         string `json:"Size"`
	Status       string `json:"Status"`
}

func main() {
	cmd := exec.Command("docker", "volume", "ls", "-f", "dangling=true", "--format", "json")

	stdout, err := cmd.StdoutPipe()

	if err != nil {
		log.Error("error getting StdoutPipe: %v", err)
	}

	if err := cmd.Start(); err != nil {
		log.Error("error getting docker volumes: %v", err)
	}

	reader := bufio.NewReader(stdout)

	log.Info("starting stream...")

	for {
		line, err := reader.ReadString('\n')

		if err != nil {
			if err == io.EOF {
				log.Info("end of stream...")
				os.Exit(0)
			}
		}

		jsonData := []byte(line)

		var volume Volume
		er := json.Unmarshal(jsonData, &volume)

		if er != nil {
			log.Error("error unmarshaling JSON:", er)
		}

		if volume.Labels == "com.docker.volume.anonymous=" {
			log.Info("removing volume ", volume.Name)

			cmd := exec.Command("docker", "volume", "rm", volume.Name)

			er := cmd.Run()

			if er != nil {
				log.Error("error removing volume ", er)
			} else {
				log.Info("successfully removed ", volume.Name)
			}
		}
	}
}
