package main

import (
	"fmt"
	"github.com/bryanstephens/go-eureka-client/eureka"
	"net/url"
	"strconv"
)

type EurekaRegistry struct {
	client *eureka.Client
	path   string
}

func NewEurekaRegistry(uri *url.URL) ServiceRegistry {
	urls := make([]string, 0)
	if uri.Host != "" {
		urls = append(urls, "http://"+uri.Host)
	}
	return &EurekaRegistry{client: eureka.NewClient(urls), path: uri.Path}
}

func (registry *EurekaRegistry) Register(service *Service) error {
	port, err := strconv.ParseUint(service.pp.HostPort, 10, 0)
	if err != nil {
		return err
	}
	instanceInfo := eureka.NewInstanceInfo(service.pp.HostName, service.Name, service.pp.HostIP, uint(port), uint(service.TTL))
	fmt.Println(instanceInfo)
	return registry.client.RegisterInstance(service.Name, instanceInfo)
}

func (registry *EurekaRegistry) Deregister(service *Service) error {
	return registry.client.UnregisterInstance(service.Name, service.ID)
}

func (registry *EurekaRegistry) Refresh(service *Service) error {
	return registry.client.SendHeartbeat(service.Name, service.ID)
}
