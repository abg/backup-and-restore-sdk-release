package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

var supportedAdapters = []string{"postgres"}

func isSupported(adapter string) bool {
	for _, el := range supportedAdapters {
		if el == adapter {
			return true
		}
	}
	return false
}

type Config struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Port     int    `json:"port"`
	Adapter  string `json:"adapter"`
	Host     string `json:"host"`
	Database string `json:"database"`
}

func main() {
	var configPath = flag.String("config", "", "Path to JSON config file")
	var backupAction = flag.Bool("backup", false, "Run database backup")
	var restoreAction = flag.Bool("restore", false, "Run database restore")
	var artifactFilePath = flag.String("artifact-file", "", "Path to output file")

	flag.Parse()

	if *backupAction && *restoreAction {
		failAndPrintUsage("Only one of: --backup or --restore can be provided")
	}

	if *configPath == "" {
		failAndPrintUsage("Missing --config flag")
	}

	if !*backupAction && !*restoreAction {
		failAndPrintUsage("Missing --backup or --restore flag")
	}

	if *artifactFilePath == "" {
		failAndPrintUsage("Missing --artifact-file flag")
	}

	configString, err := ioutil.ReadFile(*configPath)
	if err != nil {
		log.Fatalf("Fail reading config file: %s\n", err)
	}

	var config Config
	if err := json.Unmarshal(configString, &config); err != nil {
		log.Fatalf("Could not parse config json: %s\n", err)
	}

	if !isSupported(config.Adapter) {
		log.Fatalf("Unsupported adapter %s\n", config.Adapter)
	}

	var cmd *exec.Cmd
	if *restoreAction {
		cmd = restore(config, *artifactFilePath)
	} else {
		cmd = backup(config, *artifactFilePath)
	}

	fmt.Println(cmd.Args)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func failAndPrintUsage(message string) {
	log.Fatalf("%s\nUsage: database-backuper [--backup|--restore] --config <config-file> --artifact-file <artifact-file>\n", message)
}

func restore(config Config, artifactFilePath string) *exec.Cmd {
	pgRestorePath, pgRestorePathVariableSet := os.LookupEnv("PG_RESTORE_PATH")

	if !pgRestorePathVariableSet {
		log.Fatalln("PG_RESTORE_PATH must be set")
	}

	cmd := exec.Command(pgRestorePath,
		"-v",
		"--user="+config.Username,
		"--host="+config.Host,
		fmt.Sprintf("--port=%d", config.Port),
		"--format=custom",
		"--dbname="+config.Database,
		"--clean",
		artifactFilePath,
	)
	cmd.Env = append(cmd.Env, "PGPASSWORD="+config.Password)

	return cmd
}

func backup(config Config, artifactFilePath string) *exec.Cmd {
	pgDumpPath, pgDumpPathVariableSet := os.LookupEnv("PG_DUMP_PATH")

	if !pgDumpPathVariableSet {
		log.Fatalln("PG_DUMP_PATH must be set")
	}

	cmd := exec.Command(pgDumpPath,
		"-v",
		"--user="+config.Username,
		"--host="+config.Host,
		fmt.Sprintf("--port=%d", config.Port),
		"--format=custom",
		"--file="+artifactFilePath,
		config.Database,
	)
	cmd.Env = append(cmd.Env, "PGPASSWORD="+config.Password)

	return cmd
}