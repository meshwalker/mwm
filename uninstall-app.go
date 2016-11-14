package main

import (
	api "k8s.io/client-go/1.4/pkg/api"
    log "github.com/Sirupsen/logrus"
)

func (mwm *MWManager) UninstallApp(namespace string) error {
    log.Info("Uninstall App: ",namespace)

    opts := new(api.DeleteOptions)
    err := mwm.K8s.Core().Namespaces().Delete(namespace, opts)
    if err != nil {
        log.Error(err)
        return err
    }
    
    log.Info("Successfully removed Meshwalker App: ", namespace)
    return nil
}