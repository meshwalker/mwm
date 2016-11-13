package main

import (
    etcd "github.com/coreos/etcd/client"
	"k8s.io/client-go/1.4/kubernetes"
)

const (
    KubeAPIVersion  = "v1"
)

type Config struct {
    EtcdEndPoints string    `default:"etcd_nedpoints"`
}

type MWManager struct {
    Etcd    *etcd.KeysAPI
    K8s     *kubernetes.Clientset
}

type MeshwalkerApps struct {
    List    []MeshwalkerApp   `json:"mw-apps" etcd:"mw-apps"`
}

type MeshwalkerApp struct {
    Namespace    string     `json:"namespace" etcd:"namespace"`
    Name        string      `json:"name" etcd:"name"`
    Publisher   string      `json:"publisher" etcd:"publisher"`
    Description string      `json:"description" etcd:"description"`
    Version     int64       `json:"version" etcd:"version"`
    TargetVersion   int64   `json:"target_version" etcd:"target_version"`
    InternetAccess  bool    `json:"internet_access" etcd:"internet_access"`
    ExternalPort    int16   `json:"external_port" etcd:"external_port"`
    Containers   []Container `json:"containers" etcd:"containers"`
}

type Container struct {
    Name    string      `json:"name" etcd:"name"`
    Image   string      `json:"image" etcd:"image"`
    Secrets  []string    `json:"secrets" etcd:"secrets"`
    Configs  []string    `json:"configs" etcd:"configs"`
    Ports   []Port      `json:"ports" etcd:"ports"`
}

type Port struct {
    ContainerPort   int16   `json:"container_port" etcd:"container_port"`
    Protocol        string  `json:"protocol" etcd:"protocol"`
}

type Permissions struct {
    InternetAccess  bool
}