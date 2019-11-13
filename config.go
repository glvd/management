package management

// DefaultConfig ...
func DefaultConfig() DBConfig {
	return DBConfig{
		ShowSQL:      true,
		ShowExecTime: true,
		UseCache:     false,
		Create:       false,
		DBType:       "sqlite3",
		Addr:         ".",
		Username:     "",
		Password:     "",
		Schema:       "deploy",
		Charset:      "",
		Prefix:       "",
		Location:     "",
	}
}
