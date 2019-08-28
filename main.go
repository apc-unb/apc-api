package main

import (
	"github.com/VerasThiago/plataforma-apc/server"
)

func main() {
	a := server.App{}

	//mongoHost := os.Getenv("CONN")
	mongoHost := "mongo"
	//mongoHost := "localhost"

	a.Initialize(mongoHost, "27017", "f3d968eea83ad8d5f21cad0365edcc200439c6f0", "b30c206b689d5ba004534c6780aa7be8e234a7f3")

	defer a.DB.Disconnect(nil)

	a.Run(":8080")

}
