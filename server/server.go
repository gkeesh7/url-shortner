package server

import "url-shortner/utils/concurrencyutils"

const (
	logTag = "server"
)

// Start starts the server
func Start() {
	go StartServer()
}

// Stop stops all servers.
func Stop() <-chan struct{} {
	return concurrencyutils.WaitChannels(stopHTTPServer())
}
