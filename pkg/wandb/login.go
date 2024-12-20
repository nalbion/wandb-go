package wandb

import (
	"errors"
	"fmt"
	"os/user"
	"path/filepath"
	"strings"

	"github.com/jdxcode/netrc"
	"github.com/lukasbm/wandb-go/internal/settings"
)

type LoginOptions struct {
	Key     string
	APIURL  string
	ReLogin bool
}

func Login(opts LoginOptions) (string, error) {
	if opts.Key != "" && !opts.ReLogin {
		return opts.Key, nil
	}

	apiUrl := opts.APIURL
	if apiUrl == "" {
		apiUrl = "https://api.wandb.ai"
	}

	relogin := opts.Key != "" && opts.ReLogin

	if relogin {
		cfg := settings.NewConfig()
		appHost := cfg.AppURL
		if appHost == "" {
			appHost = strings.Replace(apiUrl, "api.wandb", "wandb", 1)
		}

		return "", fmt.Errorf("need to provide API key from %s/authorize", appHost)
	}

	usr, err := user.Current()
	if err != nil {
		return "", err
	}

	n, err := netrc.Parse(filepath.Join(usr.HomeDir, ".netrc"))
	if err != nil {
		return "", err
	}

	domain := strings.Split(apiUrl, "://")[1]
	key := n.Machine(domain).Get("password")

	if key == "" {
		return "", errors.New("no API key found")
	}
	return key, nil
}
