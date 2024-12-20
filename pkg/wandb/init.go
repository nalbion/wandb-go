package wandb

import (
	"context"
	"fmt"
	"net/url"
	"sync"

	"github.com/lukasbm/wandb-go/api"
	"github.com/lukasbm/wandb-go/internal/observability"
	"github.com/lukasbm/wandb-go/internal/utils"
	"github.com/shurcooL/graphql"
	settings "google.golang.org/genproto/googleapis/cloud/securitycenter/settings/v1beta1"

	"github.com/lukasbm/wandb-go/internal/api/credentials"
	proto "github.com/lukasbm/wandb-go/internal/api/proto/api"
	"github.com/lukasbm/wandb-go/internal/api/proto/filestream"
	"github.com/lukasbm/wandb-go/internal/api/wboperation"
)

type InitOptions struct {
	Entity   string
	Project  string
	Dir      string
	ID       string
	Name     string
	Notes    string
	Tags     []string
	Config   map[string]any
	Settings Settings
	// ConfigExcludeKeys []string
	// ConfigIncludeKeys []string
	// AllowValChange    bool
	Group   string
	JobType string
	Mode    string
	// Force             bool
	// Anonymous         *string
	// Reinit            bool
	Resume bool
	// ResumeFrom        *string
	// FormFrom          *string
	// SaveCode          bool
	// SyncTensorBoard   bool
	// MonitorGym        bool
}

var (
	run   *Run
	mutex = sync.Mutex{}
)

func Init(ctx context.Context, entity string, project string, options InitOptions) (*Run, error) {
	mutex.Lock()
	defer mutex.Unlock()

	if run != nil {
		return run, nil
	}

	if options.ID == "" {
		options.ID = generateRunId()
	}

	cfg := NewConfig()
	settings := SettingsWithOverrides(cfg, &options.Settings)
	if settings.APIKey == "" {
		apiKey, err := Login(LoginOptions{
			APIURL: settings.BaseURL,
		})
		if err != nil {
			return nil, err
		}

		settings.APIKey = apiKey
	}
	options.Settings = *settings

	client := graphql.Client{}
	runResponse, err := api.UpsertRun(ctx, client, project, entity, options.Name)
	if err != nil {
		return nil, err
	}
	fmt.Printf("Run response: %v\n", runResponse)

	protoClient, err := createProtoClient(settings)
	if err != nil {
		return nil, err
	}

	fs := filestream.NewFileStream(filestream.FileStreamParams{
		Settings:          settings,
		Logger:            &observability.CoreLogger{},
		Operations:        &wboperation.WandbOperations{},
		Printer:           &observability.Printer{},
		ApiClient:         protoClient,
		TransmitRateLimit: nil,
	})

	run = NewRun(entity, project, protoClient, options)
	return run, nil
}

func createProtoClient(settings *settings.Settings) (proto.Client, error) {
	baseUrl, err := url.Parse(settings.BaseURL)
	if err != nil {
		return nil, err
	}

	credentialProvider := credentials.NewAPIKeyCredentialProvider(settings.APIKey)

	backendOpts := proto.BackendOptions{
		BaseURL:            baseUrl,
		CredentialProvider: credentialProvider,
	}
	backend := proto.New(backendOpts)

	clientOpts := proto.ClientOptions{
		CredentialProvider: credentialProvider,
	}
	client := backend.NewClient(clientOpts)

	return client, nil
}

func generateRunId() string {
	return utils.RandStringBytes(10)
}
