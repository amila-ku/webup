package bootstrap

import (
	"errors"
	"fmt"
	"os/exec"
)

func commandExists(cmd string) bool {
	_, err := exec.LookPath(cmd)
	if err != nil {
		return false
	}
	return true
}

func isHugoInstalled() bool {
	return commandExists("hugo")
}

func initiateHugoWebSite(websitename string) error {
	if !isHugoInstalled() {
		return errors.New("Hugo not installed, please install hugo and retry https://gohugo.io/getting-started/quick-start/#step-1-install-hugo")
	}

	cmd := exec.Command("hugo", "new", "site", websitename)
	err := cmd.Run()
	if err != nil {
		return err
	}
	fmt.Println("Created hugo website!")

	return nil
}
