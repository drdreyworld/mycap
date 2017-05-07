package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"mycap/libs/config"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"syscall"
)

var (
	process    = flag.String("process", "", "Process name (agent, server or web)")
	command    = flag.String("command", "", "Command name (start, stop, restart, status)")
	configFile = flag.String("config", "./../etc/tool.json", "Config filename")
)

type ProcessInfo struct {
	Config string `json:"config"`
	Exec   string `json:"exec"`
	Pidf   string `json:"pidf"`
	Logf   string `json:"logf"`
}

type ProcessesInfo map[string]ProcessInfo

type Tool struct {
	Processes ProcessesInfo `json:"processes"`
}

func main() {
	flag.Parse()

	tool := Tool{}

	if err := config.ReadConfig(*configFile, &tool); err != nil {
		log.Fatal(err)
	}

	tool.ExecCommand(*process, *command)
}

func (self *Tool) ExecCommand(process string, command string) {
	if pinfo, ok := self.Processes[process]; !ok {
		log.Fatalf("process with name %s not found\n", process)
	} else {
		switch command {
		case "status":
			self.CheckProcess(process, pinfo)
			break

		case "start":
			self.StartProcess(process, pinfo)
			break

		case "stop":
			self.StopProcess(process, pinfo)
			break

		case "restart":
			self.RestartProcess(process, pinfo)
			break
		}
	}
}

func (self *Tool) CheckProcess(name string, pinfo ProcessInfo) {
	if isProcessRunning(pinfo.Pidf) {
		log.Printf("process %s status is running", name)
	} else {
		log.Printf("process %s status is stopped", name)
	}
}

func (self *Tool) StartProcess(process string, pinfo ProcessInfo) {
	if isProcessRunning(pinfo.Pidf) {
		log.Printf("%s is already started\n", process)
	} else {
		log.Printf("start %s\n", process)

		logfile, err := os.Create(pinfo.Logf)
		if err != nil {
			panic(err)
		}
		defer logfile.Close()

		cmd := exec.Command(pinfo.Exec, "-config", pinfo.Config)
		cmd.Stdout = logfile
		cmd.Stderr = logfile

		if err := cmd.Start(); err != nil {
			log.Fatal(err)
		}

		log.Printf("%s started with pid %d\n", process, cmd.Process.Pid)

		ioutil.WriteFile(pinfo.Pidf, ([]byte)(fmt.Sprintf("%d", cmd.Process.Pid)), 0644)
	}
}

func (self *Tool) StopProcess(process string, pinfo ProcessInfo) {
	if !isProcessRunning(pinfo.Pidf) {
		log.Printf("%s is not runningn\n", process)
	} else {
		log.Printf("start %s\n", process)

		if err := getProcessByPIDFileName(pinfo.Pidf).Kill(); err != nil {
			log.Fatal(err)
		}

		log.Printf("%s stopped\n", process)

		if err := os.Remove(pinfo.Pidf); err != nil {
			log.Fatal(err)
		}

		log.Printf("pidfile %s removed\n", pinfo.Pidf)
	}
}

func (self *Tool) RestartProcess(process string, pinfo ProcessInfo) {
	if isProcessRunning(pinfo.Pidf) {
		self.StopProcess(process, pinfo)
	}

	self.StartProcess(process, pinfo)
}

func isProcessRunning(pidfile string) bool {
	var err error

	if _, err = os.Stat(pidfile); err != nil {
		return false
	}

	content, err := ioutil.ReadFile(pidfile)

	if err != nil {
		log.Fatal(err)
	}

	pid, err := strconv.Atoi(strings.TrimSpace(string(content)))

	if err != nil {
		log.Fatal(err)
	}

	process, err := os.FindProcess(pid)

	if err != nil {
		log.Fatal(err)
	}

	err = process.Signal(syscall.Signal(0))

	if err != nil && err.Error() == "os: process already finished" {
		return false
	} else if err != nil {
		log.Fatal(err)
	}

	return err == nil
}

func getProcessByPIDFileName(pidfile string) *os.Process {
	var err error

	content, err := ioutil.ReadFile(pidfile)

	if err != nil {
		log.Fatal(err)
	}

	pid, err := strconv.Atoi(strings.TrimSpace(string(content)))

	if err != nil {
		log.Fatal(err)
	}

	process, err := os.FindProcess(pid)

	if err != nil {
		log.Fatal(err)
	}

	return process
}
