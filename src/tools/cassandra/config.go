package cassandra

import (
	"fmt"
	"regexp"
)

type (
	// BaseConfig is the common config
	// for all of the tasks that work
	// with cassandra
	BaseConfig struct {
		CassHosts    string
		CassPort     int
		CassUser     string
		CassPassword string
		CassKeyspace string
	}

	// UpdateSchemaConfig holds the config
	// params for executing a UpdateSchemaTask
	UpdateSchemaConfig struct {
		BaseConfig
		TargetVersion string
		SchemaDir     string
		IsDryRun      bool
	}

	// SetupSchemaConfig holds the config
	// params need by the SetupSchemaTask
	SetupSchemaConfig struct {
		BaseConfig
		SchemaFilePath    string
		InitialVersion    string
		Overwrite         bool // overwrite previous data
		DisableVersioning bool // do not use schema versioning
	}

	// CreateKeyspaceConfig holds the config
	// params needed to create a cassandra
	// keyspace
	CreateKeyspaceConfig struct {
		BaseConfig
		ReplicationFactor int
	}

	// ConfigError is an error type that
	// represents a problem with the config
	ConfigError struct {
		msg string
	}
)

const (
	cliOptEndpoint          = "endpoint"
	cliOptPort              = "port"
	cliOptUser              = "user"
	cliOptPassword          = "password"
	cliOptKeyspace          = "keyspace"
	cliOptVersion           = "version"
	cliOptSchemaFile        = "schema-file"
	cliOptOverwrite         = "overwrite"
	cliOptDisableVersioning = "disable-versioning"
	cliOptTargetVersion     = "version"
	cliOptDryrun            = "dryrun"
	cliOptSchemaDir         = "schema-dir"
	cliOptReplicationFactor = "replication-factor"
	cliOptQuiet             = "quiet"

	cliFlagEndpoint          = cliOptEndpoint + ", ep"
	cliFlagPort              = cliOptPort + ", p"
	cliFlagUser              = cliOptUser + ", u"
	cliFlagPassword          = cliOptPassword + ", pw"
	cliFlagKeyspace          = cliOptKeyspace + ", k"
	cliFlagVersion           = cliOptVersion + ", v"
	cliFlagSchemaFile        = cliOptSchemaFile + ", f"
	cliFlagOverwrite         = cliOptOverwrite + ", o"
	cliFlagDisableVersioning = cliOptDisableVersioning + ", d"
	cliFlagTargetVersion     = cliOptTargetVersion + ", v"
	cliFlagDryrun            = cliOptDryrun + ", y"
	cliFlagSchemaDir         = cliOptSchemaDir + ", d"
	cliFlagReplicationFactor = cliOptReplicationFactor + ", rf"
	cliFlagQuiet             = cliOptQuiet + ", q"
)

var rmspaceRegex = regexp.MustCompile("\\s+")

func newConfigError(msg string) error {
	return &ConfigError{msg: msg}
}

func (e *ConfigError) Error() string {
	return fmt.Sprintf("Config Error:%v", e.msg)
}
