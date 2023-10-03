package tools

import "os"

func GetMyNamespace() string {
	namespace, _ := os.ReadFile("/var/run/secrets/kubernetes.io/serviceaccount/namespace")

	return string(namespace)
}
