package main

import (
    //"k8s.io/client-go/1.4/kubernetes"
	//"k8s.io/client-go/1.4/pkg/api"
	//"k8s.io/client-go/1.4/rest"
    log "github.com/Sirupsen/logrus"
)

func (mwm *MWManager) InstallApp() error {
    app := &MeshwalkerApp{
        Namespace: "my.awesome.app",
        Name: "My App Name",
        Publisher: "Richard Jentzsch",
    }

    if err := mwm.GenKubeNamespace(app); err != nil {
        log.Error(err)
    }

    log.Info("Application successfully installed!")
    return nil
}