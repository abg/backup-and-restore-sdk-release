package config

import (
	"log"
	"os"
)

type UtilityPaths struct {
	Client  string
	Dump    string
	Restore string
}

type UtilitiesConfig struct {
	Postgres96  UtilityPaths
	Postgres94  UtilityPaths
	Postgres106 UtilityPaths
	Postgres11  UtilityPaths
	Mariadb     UtilityPaths
	Mysql55     UtilityPaths
	Mysql56     UtilityPaths
	Mysql57     UtilityPaths
}

func GetUtilitiesConfigFromEnv() UtilitiesConfig {
	return UtilitiesConfig{
		Postgres11: UtilityPaths{
			Client: lookupEnv("PG_CLIENT_PATH"),
			Dump:   lookupEnv("PG_DUMP_11_PATH"),
		},
		Postgres106: UtilityPaths{
			Client:  lookupEnv("PG_CLIENT_PATH"),
			Dump:    lookupEnv("PG_DUMP_10_6_PATH"),
			Restore: lookupEnv("PG_RESTORE_10_6_PATH"),
		},
		Postgres96: UtilityPaths{
			Client:  lookupEnv("PG_CLIENT_PATH"),
			Dump:    lookupEnv("PG_DUMP_9_6_PATH"),
			Restore: lookupEnv("PG_RESTORE_9_6_PATH"),
		},
		Postgres94: UtilityPaths{
			Client:  lookupEnv("PG_CLIENT_PATH"),
			Dump:    lookupEnv("PG_DUMP_9_4_PATH"),
			Restore: lookupEnv("PG_RESTORE_9_4_PATH"),
		},
		Mariadb: UtilityPaths{
			Client:  lookupEnv("MARIADB_CLIENT_PATH"),
			Dump:    lookupEnv("MARIADB_DUMP_PATH"),
			Restore: lookupEnv("MARIADB_CLIENT_PATH"),
		},
		Mysql55: UtilityPaths{
			Client:  lookupEnv("MYSQL_CLIENT_5_5_PATH"),
			Dump:    lookupEnv("MYSQL_DUMP_5_5_PATH"),
			Restore: lookupEnv("MYSQL_CLIENT_5_5_PATH"),
		},
		Mysql56: UtilityPaths{
			Client:  lookupEnv("MYSQL_CLIENT_5_6_PATH"),
			Dump:    lookupEnv("MYSQL_DUMP_5_6_PATH"),
			Restore: lookupEnv("MYSQL_CLIENT_5_6_PATH"),
		},
		Mysql57: UtilityPaths{
			Client:  lookupEnv("MYSQL_CLIENT_5_7_PATH"),
			Dump:    lookupEnv("MYSQL_DUMP_5_7_PATH"),
			Restore: lookupEnv("MYSQL_CLIENT_5_7_PATH"),
		},
	}
}

func lookupEnv(key string) string {
	value, valueSet := os.LookupEnv(key)
	if !valueSet {
		log.Fatalln(key + " must be set")
	}
	return value
}
