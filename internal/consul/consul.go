// internal/consul/consul.go
package consul

import (
	"fmt"
	"loan-admin/internal/config"
	"loan-admin/internal/pkg/logger"

	"github.com/hashicorp/consul/api"
	"go.uber.org/zap"
)

func NewConsulClient(cfg *config.Config) (*api.Client, error) {
	if cfg.Consul.Addr == "" {
		return nil, nil
	}

	config := api.DefaultConfig()
	config.Address = cfg.Consul.Addr

	client, err := api.NewClient(config)
	if err != nil {
		return nil, fmt.Errorf("error creating consul client: %w", err)
	}

	return client, nil
}

func RegisterService(cfg *config.Config, logger *logger.Logger) {
	if cfg.Consul.Addr == "" || cfg.Env != "prod" {
		return
	}

	client, err := NewConsulClient(cfg)
	if err != nil {
		logger.Error("Failed to create consul client", zap.Error(err))
		return
	}

	registration := &api.AgentServiceRegistration{
		ID:      cfg.Consul.ServiceID,
		Name:    cfg.Consul.ServiceName,
		Address: cfg.HTTP.Host,
		Port:    cfg.HTTP.Port,
		Check: &api.AgentServiceCheck{
			HTTP:                           fmt.Sprintf("http://%s:%d/health", cfg.HTTP.Host, cfg.HTTP.Port),
			Interval:                       cfg.Consul.CheckInterval.String(),
			Timeout:                        "5s",
			DeregisterCriticalServiceAfter: "30s",
		},
	}

	if err := client.Agent().ServiceRegister(registration); err != nil {
		logger.Error("Failed to register service with consul", zap.Error(err))
	}
}

func DeregisterService(cfg *config.Config, logger *logger.Logger) {
	if cfg.Consul.Addr == "" || cfg.Env != "prod" {
		return
	}

	client, err := NewConsulClient(cfg)
	if err != nil {
		logger.Error("Failed to create consul client", zap.Error(err))
		return
	}

	if err := client.Agent().ServiceDeregister(cfg.Consul.ServiceID); err != nil {
		logger.Error("Failed to deregister service from consul", zap.Error(err))
	}
}
