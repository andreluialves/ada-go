package ports

func isValidPort(port int) bool {
	return port >= 1 && port <= 65535
}
