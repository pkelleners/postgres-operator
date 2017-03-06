package util

import (
	"math/rand"
	"time"

	"github.bus.zalan.do/acid/postgres-operator/pkg/spec"
	"k8s.io/client-go/pkg/api/v1"
	"k8s.io/client-go/pkg/types"
)

var passwordChars = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func init() {
	rand.Seed(int64(time.Now().Unix()))
}

func RandomPassword(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = passwordChars[rand.Intn(len(passwordChars))]
	}

	return string(b)
}

func NameFromMeta(meta v1.ObjectMeta) types.NamespacedName {
	obj := types.NamespacedName{
		Namespace: meta.Namespace,
		Name:      meta.Name,
	}

	return obj
}

func PodClusterName(pod *v1.Pod) spec.ClusterName {
	if name, ok := pod.Labels["spilo-cluster"]; ok {
		return spec.ClusterName{
			Namespace: pod.Namespace,
			Name:      name,
		}
	}

	return spec.ClusterName{}
}