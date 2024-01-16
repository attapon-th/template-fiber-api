package server

func StartServer() {
	// init zerolog
	InitZeroLog()

	// start fiber server
	Listen()
}
