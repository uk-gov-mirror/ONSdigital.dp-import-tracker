package config

import (
	"time"

	"encoding/json"

	"github.com/kelseyhightower/envconfig"
)

// Config represents the app configuration
type Config struct {
	NewInstanceTopic                  string        `envconfig:"INPUT_FILE_AVAILABLE_TOPIC"`
	NewInstanceConsumerGroup          string        `envconfig:"INPUT_FILE_AVAILABLE_CONSUMER_GROUP"`
	ObservationsInsertedTopic         string        `envconfig:"IMPORT_OBSERVATIONS_INSERTED_TOPIC"`
	ObservationsInsertedConsumerGroup string        `envconfig:"IMPORT_OBSERVATIONS_INSERTED_CONSUMER_GROUP"`
	HierarchyBuiltTopic               string        `envconfig:"HIERARCHY_BUILT_TOPIC"`
	HierarchyBuiltConsumerGroup       string        `envconfig:"HIERARCHY_BUILT_CONSUMER_GROUP"`
	SearchBuiltTopic                  string        `envconfig:"SEARCH_BUILT_TOPIC"`
	SearchBuiltConsumerGroup          string        `envconfig:"SEARCH_BUILT_CONSUMER_GROUP"`
	DataImportCompleteTopic           string        `envconfig:"DATA_IMPORT_COMPLETE_TOPIC"`
	Brokers                           []string      `envconfig:"KAFKA_ADDR"`
	ImportAPIAddr                     string        `envconfig:"IMPORT_API_ADDR"`
	DatasetAPIAddr                    string        `envconfig:"DATASET_API_ADDR"`
	ShutdownTimeout                   time.Duration `envconfig:"GRACEFUL_SHUTDOWN_TIMEOUT"`
	BindAddr                          string        `envconfig:"BIND_ADDR"`
	ServiceAuthToken                  string        `envconfig:"SERVICE_AUTH_TOKEN"                   json:"-"`
	ZebedeeURL                        string        `envconfig:"ZEBEDEE_URL"`
	DatabaseAddress                   string        `envconfig:"DATABASE_ADDRESS"                     json:"-"`
	DatabasePoolSize                  int           `envconfig:"DATABASE_POOL_SIZE"`
	CheckCompleteInterval             time.Duration `envconfig:"CHECK_COMPLETE_INTERVAL"`
	InitialiseListInterval            time.Duration `envconfig:"INITIALISE_LIST_INTERVAL"`
	InitialiseListAttempts            int           `envconfig:"INITIALISE_LIST_ATTEMPTS"`
}

// NewConfig creates the config object
func NewConfig() (*Config, error) {
	cfg := Config{
		BindAddr:                          ":21300",
		ServiceAuthToken:                  "AB0A5CFA-3C55-4FA8-AACC-F98039BED0AC",
		NewInstanceTopic:                  "input-file-available",
		NewInstanceConsumerGroup:          "dp-import-tracker",
		ObservationsInsertedTopic:         "import-observations-inserted",
		ObservationsInsertedConsumerGroup: "dp-import-tracker",
		HierarchyBuiltTopic:               "hierarchy-built",
		HierarchyBuiltConsumerGroup:       "dp-import-tracker",
		SearchBuiltTopic:                  "search-built",
		SearchBuiltConsumerGroup:          "dp-import-tracker",
		DataImportCompleteTopic:           "data-import-complete",
		Brokers:                           []string{"localhost:9092"},
		ImportAPIAddr:                     "http://localhost:21800",
		DatasetAPIAddr:                    "http://localhost:22000",
		ZebedeeURL:                        "http://localhost:8082",
		ShutdownTimeout:                   5 * time.Second,
		DatabaseAddress:                   "bolt://localhost:7687",
		DatabasePoolSize:                  30,
		CheckCompleteInterval:             2000 * time.Millisecond,
		InitialiseListInterval:            4 * time.Second,
		InitialiseListAttempts:            20,
	}
	if err := envconfig.Process("", &cfg); err != nil {
		return nil, err
	}

	cfg.ServiceAuthToken = "Bearer " + cfg.ServiceAuthToken

	return &cfg, nil
}

// String is implemented to prevent sensitive fields being logged.
// The config is returned as JSON with sensitive fields omitted.
func (config Config) String() string {
	json, _ := json.Marshal(config)
	return string(json)
}
