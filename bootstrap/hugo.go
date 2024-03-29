package bootstrap

import (
	"errors"
	"fmt"
	"os/exec"
)

func commandExists(cmd string) bool {
	_, err := exec.LookPath(cmd)
	return err == nil
}

func isHugoInstalled() bool {
	return commandExists("hugo")
}

func initiateHugoWebSite(websitename string) error {
	if !isHugoInstalled() {
		return errors.New("hugo not installed, please install hugo and retry https://gohugo.io/getting-started/quick-start/#step-1-install-hugo")
	}

	cmd := exec.Command("hugo", "new", "site", websitename)
	err := cmd.Run()
	if err != nil {
		return err
	}
	fmt.Println("Created hugo website!")

	return nil
}
