package config

type NatsStreaming struct {
	Address   string `yaml:"address" validate:"required"`
	ClusterID string `yaml:"cluster_id" validate:"required"`
	ClientID  string `yaml:"client_id" validate:"required"`
}
