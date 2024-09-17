package config

var (
	// These variables are populated at build time with the Makefile and config files.
	Version                       string
	Server_addr                   string
	Server_port                   string
	Default_journal_dir_from_home string
	Server_key                    string
)
