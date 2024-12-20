package wandb

import (
	proto "github.com/lukasbm/wandb-go/internal/api/proto/api"
	settings "google.golang.org/genproto/googleapis/cloud/securitycenter/settings/v1beta1"
)

func NewRun(entity string, project string, client proto.Client, options InitOptions) *Run {
	return &Run{
		client:   client,
		Entity:   entity,
		Project:  project,
		Name:     options.Name,
		ID:       options.ID,
		Settings: &options.Settings,
		Config:   options.Config,
		// SweepConfig: options.SweepConfig,
		// LaunchConfig: options.LaunchConfig,
	}
}

type Run struct {
	client   proto.Client
	Entity   string
	Project  string
	ID       string
	Name     string
	Settings *settings.Settings
	// Step         int
	Config       map[string]any
	SweepConfig  map[string]any
	LaunchConfig map[string]any
}

func (run *Run) Log(data map[string]any) {
	req := &proto.Request{}
	run.client.Send(req)
}

func (run *Run) Finish() {
}
