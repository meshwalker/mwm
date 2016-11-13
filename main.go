package main

import (
    //"github.com/kelseyhightower/envconfig"
    log "github.com/Sirupsen/logrus"
    //etcd "github.com/coreos/etcd/client"
    //"github.com/mickep76/etcdmap"
    //"golang.org/x/net/context"
    "k8s.io/client-go/1.4/kubernetes"
	"k8s.io/client-go/1.4/rest"
)

// ExampleNestedStruct creates a Etcd directory using a nested Go struct and then gets the directory as JSON.
func main() {
    //var mwmConfig Config
    mwManager := &MWManager{}

  /*  err := envconfig.Process("mwm", &mwmConfig)
    if err != nil {
        log.Fatal(err.Error())
    }

    // Connect to etcd.
    cfg := etcd.Config{
        Endpoints:               strings.Split(*mwmConfig.EtcdEndPoints, ","),
        Transport:               etcd.DefaultTransport,
        HeaderTimeoutPerRequest: time.Second,
    }

    client, err := etcd.New(cfg)
    if err != nil {
        log.Fatal(err)
    }

    mwManager.Etcd = etcd.NewKeysAPI(client)*/
/*
    // Create directory structure based on struct.
    err2 := etcdmap.Create(kapi, "/example", reflect.ValueOf(g))
    if err2 != nil {
        log.Fatal(err2.Error())
    }

    // Get directory structure from Etcd.
    res, err3 := kapi.Get(context.TODO(), "/example", &etcd.GetOptions{Recursive: true})
    if err3 != nil {
        log.Fatal(err3.Error())
    }

    j, err4 := etcdmap.JSON(res.Node)
    if err4 != nil {
        log.Fatal(err4.Error())
    }*/


    // creates the in-cluster config
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}
	// creates the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
    mwManager.K8s = clientset

    if err := mwManager.InstallApp(); err != nil {
        log.Error(err)
    }

    // Get directory structure from Etcd.
    /*
    	res, err5 := kapi.Get(context.TODO(), "/example/users/0", &etcd.GetOptions{Recursive: true})
    	if err5 != nil {
    		log.Fatal(err5.Error())
    	}

    	s := User{}
    	err6 := etcdmap.Struct(res.Node, reflect.ValueOf(&s))
    	if err6 != nil {
    		log.Fatal(err6.Error())
    	}

    	fmt.Println(s)
    */

}