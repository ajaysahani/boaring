package cassandra

import (
	"fmt"
	"log"

	"github.com/urfave/cli"
)

// setupSchema executes the setupSchemaTask
// using the given command line arguments
// as input
func setupSchema(cli *cli.Context) error {
	config, err := newSetupSchemaConfig(cli)
	if err != nil {
		return handleErr(newConfigError(err.Error()))
	}
	if err := handleSetupSchema(config); err != nil {
		return handleErr(err)
	}
	return nil
}

// updateSchema executes the updateSchemaTask
// using the given command lien args as input
func updateSchema(cli *cli.Context) error {
	config, err := newUpdateSchemaConfig(cli)
	if err != nil {
		return handleErr(newConfigError(err.Error()))
	}
	if err := handleUpdateSchema(config); err != nil {
		return handleErr(err)
	}
	return nil
}

// createKeyspace creates a cassandra keyspace
func createKeyspace(cli *cli.Context) error {
	config, err := newCreateKeyspaceConfig(cli)
	if err != nil {
		return handleErr(err)
	}
	client, err := newCQLClient(config.CassHosts, config.CassPort, config.CassUser, config.CassPassword, "system")
	if err != nil {
		return handleErr(fmt.Errorf("error creating cql client:%v", err))
	}
	err = client.CreateKeyspace(config.CassKeyspace, config.ReplicationFactor)
	if err != nil {
		return handleErr(fmt.Errorf("error creating keyspace:%v", err))
	}
	return nil
}

func handleUpdateSchema(config *UpdateSchemaConfig) error {
	task, err := NewUpdateSchemaTask(config)
	if err != nil {
		return fmt.Errorf("error creating task, err=%v", err)
	}
	if err := task.run(); err != nil {
		return fmt.Errorf("error setting up schema, err=%v", err)
	}
	return nil
}

func handleSetupSchema(config *SetupSchemaConfig) error {
	task, err := newSetupSchemaTask(config)
	if err != nil {
		return fmt.Errorf("error creating task, err=%v", err)
	}
	if err := task.run(); err != nil {
		return fmt.Errorf("error setting up schema, err=%v", err)
	}
	return nil
}

func validateSetupSchemaConfig(config *SetupSchemaConfig) error {
	if len(config.CassHosts) == 0 {
		return newConfigError("missing cassandra endpoint argument " + flag(cliOptEndpoint))
	}
	if config.CassPort == 0 {
		config.CassPort = defaultCassandraPort
	}
	if len(config.CassKeyspace) == 0 {
		return newConfigError("missing " + flag(cliOptKeyspace) + " argument ")
	}
	if len(config.SchemaFilePath) == 0 && config.DisableVersioning {
		return newConfigError("missing schemaFilePath " + flag(cliOptSchemaFile))
	}
	if (config.DisableVersioning && len(config.InitialVersion) > 0) ||
		(!config.DisableVersioning && len(config.InitialVersion) == 0) {
		return newConfigError("either " + flag(cliOptDisableVersioning) + " or " +
			flag(cliOptVersion) + " but not both must be specified")
	}
	if !config.DisableVersioning {
		ver, err := parseValidateVersion(config.InitialVersion)
		if err != nil {
			return newConfigError("invalid " + flag(cliOptVersion) + " argument:" + err.Error())
		}
		config.InitialVersion = ver
	}
	return nil
}

func newSetupSchemaConfig(cli *cli.Context) (*SetupSchemaConfig, error) {

	config := new(SetupSchemaConfig)
	config.CassHosts = cli.GlobalString(cliOptEndpoint)
	config.CassPort = cli.GlobalInt(cliOptPort)
	config.CassUser = cli.GlobalString(cliOptUser)
	config.CassPassword = cli.GlobalString(cliOptPassword)
	config.CassKeyspace = cli.GlobalString(cliOptKeyspace)
	config.SchemaFilePath = cli.String(cliOptSchemaFile)
	config.InitialVersion = cli.String(cliOptVersion)
	config.DisableVersioning = cli.Bool(cliOptDisableVersioning)
	config.Overwrite = cli.Bool(cliOptOverwrite)

	if err := validateSetupSchemaConfig(config); err != nil {
		return nil, err
	}

	return config, nil
}

func validateUpdateSchemaConfig(config *UpdateSchemaConfig) error {

	if len(config.CassHosts) == 0 {
		return newConfigError("missing cassandra endpoint argument " + flag(cliOptEndpoint))
	}
	if config.CassPort == 0 {
		config.CassPort = defaultCassandraPort
	}
	if len(config.CassKeyspace) == 0 {
		return newConfigError("missing " + flag(cliOptKeyspace) + " argument ")
	}
	if len(config.SchemaDir) == 0 {
		return newConfigError("missing " + flag(cliOptSchemaDir) + " argument ")
	}
	if len(config.TargetVersion) > 0 {
		ver, err := parseValidateVersion(config.TargetVersion)
		if err != nil {
			return newConfigError("invalid " + flag(cliOptTargetVersion) + " argument:" + err.Error())
		}
		config.TargetVersion = ver
	}
	return nil
}

func newUpdateSchemaConfig(cli *cli.Context) (*UpdateSchemaConfig, error) {

	config := new(UpdateSchemaConfig)
	config.CassHosts = cli.GlobalString(cliOptEndpoint)
	config.CassPort = cli.GlobalInt(cliOptPort)
	config.CassUser = cli.GlobalString(cliOptUser)
	config.CassPassword = cli.GlobalString(cliOptPassword)
	config.CassKeyspace = cli.GlobalString(cliOptKeyspace)
	config.SchemaDir = cli.String(cliOptSchemaDir)
	config.IsDryRun = cli.Bool(cliOptDryrun)
	config.TargetVersion = cli.String(cliOptTargetVersion)

	if err := validateUpdateSchemaConfig(config); err != nil {
		return nil, err
	}

	return config, nil
}

func newCreateKeyspaceConfig(cli *cli.Context) (*CreateKeyspaceConfig, error) {
	config := new(CreateKeyspaceConfig)
	config.CassHosts = cli.GlobalString(cliOptEndpoint)
	config.CassPort = cli.GlobalInt(cliOptPort)
	config.CassUser = cli.GlobalString(cliOptUser)
	config.CassPassword = cli.GlobalString(cliOptPassword)
	config.CassKeyspace = cli.String(cliOptKeyspace)
	config.ReplicationFactor = cli.Int(cliOptReplicationFactor)

	if err := validateCreateKeyspaceConfig(config); err != nil {
		return nil, err
	}
	return config, nil
}

func validateCreateKeyspaceConfig(config *CreateKeyspaceConfig) error {
	if len(config.CassHosts) == 0 {
		return newConfigError("missing cassandra endpoint argument " + flag(cliOptEndpoint))
	}
	if config.CassPort == 0 {
		config.CassPort = defaultCassandraPort
	}
	if len(config.CassKeyspace) == 0 {
		return newConfigError("missing " + flag(cliOptKeyspace) + " argument ")
	}
	return nil
}

func flag(opt string) string {
	return "(-" + opt + ")"
}

func handleErr(err error) error {
	log.Println(err)
	return err
}
