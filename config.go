package management

// DefaultConfig ...
func DefaultConfig() DBConfig {
	return DBConfig{
		DBType:   "sqlite3",
		Schema:   "deploy",
		Username: "",
		Password: "",
		Addr:     ".",
	}
}
