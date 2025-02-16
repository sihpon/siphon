package container

import "sigs.k8s.io/controller-runtime/pkg/client"

type Containerer interface {
	KubeClient() client.Client
}

type Container struct {
	client client.Client
}

func (c *Container) KubeClient() client.Client {
	return c.client
}

func NewContainer(client client.Client) Containerer {
	return &Container{
		client: client,
	}
}
