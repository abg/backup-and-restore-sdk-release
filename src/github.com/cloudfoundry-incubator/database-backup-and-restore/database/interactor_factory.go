package database

import (
	"github.com/cloudfoundry-incubator/database-backup-and-restore/config"
	"github.com/cloudfoundry-incubator/database-backup-and-restore/mysql"
	"github.com/cloudfoundry-incubator/database-backup-and-restore/postgres"
	"github.com/cloudfoundry-incubator/database-backup-and-restore/version"
)

type InteractorFactory struct {
	utilitiesConfig               config.UtilitiesConfig
	postgresServerVersionDetector ServerVersionDetector
}

func NewInteractorFactory(
	utilitiesConfig config.UtilitiesConfig,
	postgresServerVersionDetector ServerVersionDetector) InteractorFactory {

	return InteractorFactory{
		utilitiesConfig:               utilitiesConfig,
		postgresServerVersionDetector: postgresServerVersionDetector,
	}
}

func (f InteractorFactory) Make(action Action, config config.ConnectionConfig) Interactor {
	switch {
	case config.Adapter == "postgres" && action == "backup":
		return f.makePostgresBackuper(config)
	case config.Adapter == "mysql" && action == "backup":
		return f.makeMysqlBackuper(config)
	case config.Adapter == "postgres" && action == "restore":
		return postgres.NewRestorer(config, f.utilitiesConfig.Postgres94.Restore)
	case config.Adapter == "mysql" && action == "restore":
		return mysql.NewRestorer(config, f.utilitiesConfig.Mysql.Restore)
	}

	return nil
}

func (f InteractorFactory) makeMysqlBackuper(config config.ConnectionConfig) Interactor {
	return NewVersionSafeInteractor(
		mysql.NewBackuper(config, f.utilitiesConfig.Mysql.Dump),
		mysql.NewServerVersionDetector(f.utilitiesConfig.Mysql.Client),
		mysql.NewMysqlDumpUtilityVersionDetector(f.utilitiesConfig.Mysql.Dump),
		config)
}

func (f InteractorFactory) makePostgresBackuper(config config.ConnectionConfig) Interactor {
	// TODO: err
	postgresVersion, _ := f.postgresServerVersionDetector.GetVersion(config)

	postgres94Version := version.SemanticVersion{Major: "9", Minor: "4"}
	var pgDumpPath, psqlPath string
	if postgres94Version.MinorVersionMatches(postgresVersion) {
		psqlPath = f.utilitiesConfig.Postgres94.Client
		pgDumpPath = f.utilitiesConfig.Postgres94.Dump
	} else {
		psqlPath = f.utilitiesConfig.Postgres96.Client
		pgDumpPath = f.utilitiesConfig.Postgres96.Dump
	}

	postgresBackuper := postgres.NewBackuper(config, pgDumpPath)
	tableChecker := postgres.NewTableChecker(config, psqlPath)
	return NewTableCheckingInteractor(config, tableChecker, postgresBackuper)
}