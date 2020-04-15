package springparse

import (
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"strings"
)

type kubeInfo struct {
	PodName       string `json:"podname"`   // Podname
	NameSpace     string `json:"namespace"` // Namespace
	App           string `json:"app"`       // App
	AppType       string `json:"apptype"`   // AppType
	ContainerName string `json:"containername"`
}

type labels struct {
	AppName string `json:"app_kubernetes_io/name"`
}

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
		return fmt.Errorf("Unable to get in cluster config in pod %v", err)
	}
	// creates the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return fmt.Errorf("Unable to set new config %v", err)
	}
	result, err := clientset.CoreV1().Pods(c.nameSpace).Get(c.podName, metav1.GetOptions{})
	if err != nil {
		return fmt.Errorf("Unable to get pod specs %v", err)
	}
	s.es.KubeInfo = kubeInfo{
		PodName:       c.podName,
		NameSpace:     c.nameSpace,
		ContainerName: c.containerName,
		App:           result.ObjectMeta.Labels["app"],
		AppType:       result.ObjectMeta.Labels["app.type"],
	}
	s.es.Labels = labels{
		AppName: result.ObjectMeta.Labels["app"],
	}
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
