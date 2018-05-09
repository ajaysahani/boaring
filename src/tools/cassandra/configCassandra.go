package cassandra

// Cassandra contains configuration to connect to Cassandra cluster
type Cassandra struct {
	// Hosts is a csv of cassandra endpoints
	Hosts string `yaml:"hosts" validate:"nonzero"`
	// Port is the cassandra port used for connection by gocql client
	Port int `yaml:"port"`
	// User is the cassandra user used for authentication by gocql client
	User string `yaml:"user"`
	// Password is the cassandra password used for authentication by gocql client
	Password string `yaml:"password"`
	// Keyspace is the cassandra keyspace
	Keyspace string `yaml:"keyspace" validate:"nonzero"`
	// VisibilityKeyspace is the cassandra keyspace for visibility store
	VisibilityKeyspace string `yaml:"visibilityKeyspace" validate:"nonzero"`
	// Consistency is the default cassandra consistency level
	Consistency string `yaml:"consistency"`
	// Datacenter is the data center filter arg for cassandra
	Datacenter string `yaml:"datacenter"`
	// NumHistoryShards is the desired number of history shards
	NumHistoryShards int `yaml:"numHistoryShards" validate:"nonzero"`
}
