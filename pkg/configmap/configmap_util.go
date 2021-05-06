package configmap

import (
	"context"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/equinor/radix-config-2-map/pkg/kubernetes"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// CreateFromFile Creates a configmap by name from file
func CreateFromFile(namespace, name, file string) error {
	content, err := readConfigFile(file)
	if err != nil {
		return fmt.Errorf("Could not find or read config yaml file \"%s\"", file)
	}

	client := kubernetes.GetKubernetesClient()
	_, err = client.CoreV1().ConfigMaps(namespace).Create(
		context.Background(),
		&corev1.ConfigMap{
			ObjectMeta: metav1.ObjectMeta{
				Name:      name,
				Namespace: namespace,
			},
			Data: map[string]string{
				"content": string(content),
			},
		},
		metav1.CreateOptions{})

	if err != nil {
		return err
	}

	return nil
}

func readConfigFile(filename string) ([]byte, error) {
	var content []byte
	var err error
	for _, filename := range filenameCandidates(filename) {
		content, err = ioutil.ReadFile(filename)
		if err == nil {
			break
		}
	}
	return content, err
}

func filenameCandidates(filename string) []string {
	if strings.HasSuffix(filename, ".yaml") {
		filename = filename[:len(filename)-5]
	} else if strings.HasSuffix(filename, ".yml") {
		filename = filename[:len(filename)-4]
	}

	return []string{
		filename + ".yaml",
		filename + ".yml",
		filename,
	}
}
