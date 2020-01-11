package config

import "strconv"

// GetMQHost get the message queue host
func GetMQHost() string {
	return getenv("MQ_HOST", "127.0.0.1")
}

// GetMQPort get the message queue port
func GetMQPort() int {
	portStr := getenv("MQ_PORT", "11300")
	port, _ := strconv.Atoi(portStr)
	return port
}
