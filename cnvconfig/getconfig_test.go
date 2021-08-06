package cnvconfig

import (
	"context"
	"os"
	"testing"

	converter "github.com/GoogleCloudPlatform/terraform-google-conversion/google"
	"github.com/stretchr/testify/assert"
)

type configAttrGetter func(cfg *converter.Config) string

func getCredentials(cfg *converter.Config) string {
	return cfg.Credentials
}
func getAccessToken(cfg *converter.Config) string {
	return cfg.AccessToken
}

func TestNewConverterCredentials(t *testing.T) {
	ctx := context.Background()
	offline := true
	cases := []struct {
		name           string
		envKey         string
		envValue       string
		getConfigValue configAttrGetter
	}{
		{
			name:           "GOOGLE_CREDENTIALS",
			envKey:         "GOOGLE_CREDENTIALS",
			envValue:       "whatever",
			getConfigValue: getCredentials,
		},
		{
			name:           "GOOGLE_CLOUD_KEYFILE_JSON",
			envKey:         "GOOGLE_CLOUD_KEYFILE_JSON",
			envValue:       "whatever",
			getConfigValue: getCredentials,
		},
		{
			name:           "GCLOUD_KEYFILE_JSON",
			envKey:         "GCLOUD_KEYFILE_JSON",
			envValue:       "whatever",
			getConfigValue: getCredentials,
		},
		{
			name:           "GOOGLE_OAUTH_ACCESS_TOKEN",
			envKey:         "GOOGLE_OAUTH_ACCESS_TOKEN",
			envValue:       "whatever",
			getConfigValue: getAccessToken,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			originalValue, isSet := os.LookupEnv(c.envKey)
			err := os.Setenv(c.envKey, c.envValue)
			if err != nil {
				t.Fatalf("error setting env var %s=%s: %s", c.envKey, c.envValue, err)
			}

			cfg, err := GetConfig(ctx, "project", offline)
			if err != nil {
				t.Fatalf("error building converter: %s", err)
			}

			assert.EqualValues(t, c.getConfigValue(cfg), c.envValue)

			if isSet {
				err = os.Setenv(c.envKey, originalValue)
				if err != nil {
					t.Fatalf("error setting env var %s=%s: %s", c.envKey, originalValue, err)
				}
			} else {
				err = os.Unsetenv(c.envKey)
				if err != nil {
					t.Fatalf("error unsetting env var %s: %s", c.envKey, err)
				}
			}
		})
	}
}
