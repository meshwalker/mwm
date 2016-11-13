package main

import (
    //"k8s.io/client-go/1.4/kubernetes"
	//api "k8s.io/client-go/1.4/pkg/api"
	v1 "k8s.io/client-go/1.4/pkg/api/v1"
	//"k8s.io/client-go/1.4/rest"
    log "github.com/Sirupsen/logrus"
	"strconv"
	"k8s.io/client-go/1.4/pkg/api/unversioned"
)

func (mwm *MWManager) GenKubeNamespace(app *MeshwalkerApp) error {
    log.Info("Generate kubernetes namespace object")

    labels := make(map[string]string)
    labels["publisher"] = app.Publisher
    labels["app-name"]  = app.Name
    labels["version"]   = strconv.Itoa(int(app.Version))

    namespace := &v1.Namespace{
        TypeMeta:   unversioned.TypeMeta{
            Kind:       "Namespace",
            APIVersion: KubeAPIVersion,
        },
        ObjectMeta: v1.ObjectMeta{
            Name:   app.Namespace,
            Labels: labels,
        },
    }

    resp, err := mwm.K8s.Core().Namespaces().Create(namespace)
    if err != nil {
        log.Fatal(err)
    }
    
    log.Info("Namespace created: \n", resp)
    return nil
}