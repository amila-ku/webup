package bootstrap

func SetUpBasicWebSite(wsn string) error {
	err := initiateHugoWebSite(wsn)
	if err != nil {
		return err
	}

	return nil
}
