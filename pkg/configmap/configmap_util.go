package configmap

import (
	"io/ioutil"

	"github.com/equinor/radix-config-2-map/pkg/kubernetes"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// CreateFromFile Creates a configmap by name from file
func CreateFromFile(namespace, name, file string) error {
	client := kubernetes.GetKubernetesClient()
	content, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}

	_, err = client.CoreV1().ConfigMaps(namespace).Create(
		&corev1.ConfigMap{
			ObjectMeta: metav1.ObjectMeta{
				Name:      name,
				Namespace: namespace,
			},
			Data: map[string]string{
				"content": string(content),
			},
		})

	if err != nil {
		return err
	}

	return nil
}
