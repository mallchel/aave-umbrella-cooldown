package appconfig

import "1-task/internal/envutil"

// ServiceConfig contains shared runtime configuration for services.
type ServiceConfig struct {
	HTTPAddr    string
	SwaggerAddr string
	PostgresDSN string

	DaemonPidFileName  string
	DaemonLogFileName  string
	MigrationsPath     string
	RPCURL             string
	UmbrellaConfigPath string
}

// LoadServiceConfig returns environment-driven runtime configuration.
func LoadServiceConfig() ServiceConfig {
	return ServiceConfig{
		HTTPAddr:    envutil.Get("HTTP_ADDR", ":8888"),
		SwaggerAddr: envutil.Get("SWAGGER_ADDR", ":9090"),
		PostgresDSN: envutil.Get("POSTGRES_DSN", "postgresql://umbrella_user:umbrella_pass@localhost:5432/umbrella_db?sslmode=disable"),

		DaemonPidFileName:  envutil.Get("DAEMON_PID_FILE", "./tmp/umbrella-daemon.pid"),
		DaemonLogFileName:  envutil.Get("DAEMON_LOG_FILE", "./tmp/umbrella-daemon.log"),
		MigrationsPath:     envutil.Get("MIGRATIONS_PATH", "./migrations"),
		RPCURL:             envutil.Get("RPC_URL", "https://ethereum-rpc.publicnode.com"),
		UmbrellaConfigPath: envutil.Get("UMBRELLA_CONFIG_PATH", "./configs/umbrella/mainnet.json"),
	}
}
