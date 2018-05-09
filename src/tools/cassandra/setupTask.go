package cassandra

import "log"

// SetupSchemaTask represents a task
// that sets up cassandra schema on
// a specified keyspace
type SetupSchemaTask struct {
	client CQLClient
	config *SetupSchemaConfig
}

func newSetupSchemaTask(config *SetupSchemaConfig) (*SetupSchemaTask, error) {
	client, err := newCQLClient(config.CassHosts, config.CassPort, config.CassUser, config.CassPassword,
		config.CassKeyspace)
	if err != nil {
		return nil, err
	}
	return &SetupSchemaTask{
		config: config,
		client: client,
	}, nil
}

// run executes the task
func (task *SetupSchemaTask) run() error {

	config := task.config

	defer func() {
		task.client.Close()
	}()

	log.Printf("Starting schema setup, config=%+v\n", config)

	if config.Overwrite {
		dropAllTablesTypes(task.client)
	}

	if !config.DisableVersioning {
		log.Printf("Setting up version tables\n")
		if err := task.client.CreateSchemaVersionTables(); err != nil {
			return err
		}
	}

	if len(config.SchemaFilePath) > 0 {
		stmts, err := ParseCQLFile(config.SchemaFilePath)
		if err != nil {
			return err
		}

		log.Println("----- Creating types and tables -----")
		for _, stmt := range stmts {
			log.Println(rmspaceRegex.ReplaceAllString(stmt, " "))
			if err := task.client.Exec(stmt); err != nil {
				return err
			}
		}
		log.Println("----- Done -----")
	}

	if !config.DisableVersioning {
		log.Printf("Setting initial schema version to %v\n", config.InitialVersion)
		err := task.client.UpdateSchemaVersion(config.InitialVersion, config.InitialVersion)
		if err != nil {
			return err
		}
		log.Printf("Updating schema update log\n")
		err = task.client.WriteSchemaUpdateLog("0", config.InitialVersion, "", "initial version")
		if err != nil {
			return err
		}
	}

	log.Println("Schema setup complete")

	return nil
}
