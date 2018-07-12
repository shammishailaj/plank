package config

// Spinnaker mirrors spinnaker.yaml files on disk
type Spinnaker struct {
	Services Services `json:"services" mapstructure:"services"`
}

// Services within Spinnaker.
type Services struct {
	Fiat    Service `json:"fiat" mapstructure:"fiat"`
	Front50 Front50 `json:"front50" mapstructure:"front50"`
	Jenkins Jenkins `json:"jenkins" mapstructure:"jenkins"`
}

// Jenkins service settings
type Jenkins struct {
	Enabled       bool `json:"enabled" mapstructure:"enabled"`
	DefaultMaster struct {
		Name     string `json:"name" mapstructure:"name"`
		BaseURL  string `json:"baseUrl" mapstructure:"baseUrl"`
		Username string `json:"username" mapstructure:"username"`
		Password string `json:"password" mapstructure:"password"`
	} `json:"defaultMaster" mapstructure:"defaultMaster"`
}

// Front50 service settings.
type Front50 struct {
	Service
	StorageBucket string `json:"storage_bucket" mapstructure:"storage_bucket"`
	StoragePrefix string `json:"rootFolder" mapstructure:"rootFolder"`
	S3            struct {
		Enabled bool `json:"enabled" mapstructure:"enabled"`
	} `json:"s3" mapstructure:"s3"`
	GCS struct {
		Enabled bool `json:"enabled" mapstructure:"enabled"`
	} `json:"gcs" mapstructure:"gcs"`
}

// Service such as Fiat, Orca, Deck, Gate, etc.
type Service struct {
	Enabled bool   `json:"enabled" mapstructure:"enabled"`
	BaseURL string `json:"baseUrl" mapstructure:"baseUrl"`
}
