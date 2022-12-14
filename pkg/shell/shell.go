package shell

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

func ShellOutput(command string) (string, string, error) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer

	cmd := exec.Command("bash", "-c", command)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()

	return stdout.String(), stderr.String(), err
}

func ShellOutputWithPath(command string, path string) (string, string, error) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer

	// get the absolute path of the recieved path
	absPath, err := filepath.Abs(path)
	if err != nil {
		return "", "", err
	}

	cmd := exec.Command("bash", "-c", command)
	cmd.Dir = absPath
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err = cmd.Run()

	return stdout.String(), stderr.String(), err
}

func ShellInteractive(command string) error {
	cmd := exec.Command("bash", "-c", command)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	return cmd.Run()
}

func ShellInteractiveWithPath(command string, path string) error {
	// get the absolute path of the recieved path
	absPath, err := filepath.Abs(path)
	if err != nil {
		return err
	}

	cmd := exec.Command("bash", "-c", command)
	cmd.Dir = absPath
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	return cmd.Run()
}

func Exec(command string) {
	out, errout, err := ShellOutput(command)
	if err != nil {
		log.Printf("error: %v\n", err)
		fmt.Print(errout)
	}

	fmt.Print(out)
}

func ExecOutput(command string) string {
	out, errout, err := ShellOutput(command)
	if err != nil {
		log.Printf("error: %v\n", err)
		fmt.Print(errout)
	}

	return string(out)
}

func ExecWithPath(command string, path string) {
	out, errout, err := ShellOutputWithPath(command, path)
	if err != nil {
		log.Printf("error: %v\n", err)
		fmt.Print(errout)
	}

	fmt.Print(out)
}

func ExecInteractive(command string) {
	err := ShellInteractive(command)
	if err != nil {
		log.Printf("error: %v\n", err)
	}
}

func ExecInteractiveWithPath(command string, path string) {
  err := ShellInteractiveWithPath(command, path)
  if err != nil {
    log.Printf("error: %v\n", err)
  }
}
