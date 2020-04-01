package springparse

import (
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"strings"
)

type getkubeInfoInput struct {
	fileName string
	es       *elasticOut
}

func getkubeInfo(s getkubeInfoInput) error {
	c, err := getContainerInfo(s.fileName)
	if err != nil {
		return err
	}
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}
	// creates the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	result, err := clientset.CoreV1().Pods("api").Get("banking-649df48fdb-q9m8l", metav1.GetOptions{})
	if err != nil {
		panic(err)
	}
	s.es.PodName = c.podName
	s.es.NameSpace = c.nameSpace
	s.es.ContainerName = c.containerName
	s.es.App = result.ObjectMeta.Labels["app"]
	s.es.AppType = result.ObjectMeta.Labels["app.type"]
	return nil
}

type getContainerInfoOutput struct {
	podName       string
	nameSpace     string
	containerName string
	dockerId      string
}

func getContainerInfo(fileName string) (getContainerInfoOutput, error) {
	splitslash := strings.Split(fileName, "/")
	if len(splitslash) != 5 {
		return getContainerInfoOutput{}, fmt.Errorf("fileName error " + fileName)
	}
	splitUnderScore := strings.Split(splitslash[4], "_")
	if len(splitUnderScore) != 3 {
		return getContainerInfoOutput{}, fmt.Errorf("logFile name error" + splitslash[3])
	}
	splitDash := strings.Split(splitUnderScore[2], "-")
	if len(splitDash) != 2 {
		return getContainerInfoOutput{}, fmt.Errorf("logfile name error in dash" + splitUnderScore[2])
	}
	return getContainerInfoOutput{
		podName:       splitUnderScore[0],
		nameSpace:     splitUnderScore[1],
		containerName: splitDash[0],
		dockerId:      strings.Trim(splitDash[1], ".log"),
	}, nil
}
