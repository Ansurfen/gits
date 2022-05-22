package utils

import (
	"bytes"
	"os"
	"os/exec"
)

func Git(args ...string) string {
	git := getConf().GetString("env.git") + "/git.exe"
	cmd := exec.Command(git, args...)
	var stdout, stderr bytes.Buffer
	cmd.Stderr, cmd.Stdout = &stderr, &stdout
	if err := cmd.Run(); err != nil {
		return stderr.String()
	}
	return ""
}

func MkDirs(dirPath string) {
	err := os.MkdirAll(dirPath, 0666)
	Panic(err)
}

func GetArgs(command string, args ...string) (fargs []string) {
	fargs = append(fargs, command)
	fargs = append(fargs, args...)
	return fargs
}
