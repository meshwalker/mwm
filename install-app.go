package main

import (
    //"k8s.io/client-go/1.4/kubernetes"
	//"k8s.io/client-go/1.4/pkg/api"
	//"k8s.io/client-go/1.4/rest"
    log "github.com/Sirupsen/logrus"
	"strings"
)

func (mwm *MWManager) InstallApp() error {
    app := &MeshwalkerApp{
        Namespace: strings.Replace("my_awesome_app", ".", "_", -1),
        Name: strings.Replace("My APP", " ", "-", -1),
        Publisher: strings.Replace("Richard Jentzsch", " ", "-", -1),
    }

    if err := mwm.GenKubeNamespace(app); err != nil {
        log.Error(err)
    }

    log.Info("Application successfully installed!")
    return nil
}