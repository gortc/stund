package config

import "os"

type EnvironmentVariable string

const (
	//ListenAddres - Address listen.
	ListenAddres EnvironmentVariable = "LISTEN_ADDRESS"
	//ListenPort - Port listen.
	ListenPort EnvironmentVariable = "LISTEN_PORT"
)

func getVariable(name EnvironmentVariable, defaultValue string) string {
	if variable := os.Getenv(string(name)); variable != "" {
		return variable
	}
	return defaultValue
}

// EnvironmentVariableValue - Find to environment variable value
// or return default value of variable.
func EnvironmentVariableValue(variable EnvironmentVariable) string {
	switch variable {
	case ListenAddres:
		return getVariable(ListenAddres, "0.0.0.0")
	case ListenPort:
		return getVariable(ListenPort, "3478")
	}
	return ""
}
