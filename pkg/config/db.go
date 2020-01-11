package config

// GetDBDialect get the database dialect
func GetDBDialect() string {
	return getenv("DB_DIALECT", "sqlite3")
}

// GetDBConnectArgs get the database connect args
func GetDBConnectArgs() string {
	return getenv("DB_CONNECTARGS", "ally.db")
}
