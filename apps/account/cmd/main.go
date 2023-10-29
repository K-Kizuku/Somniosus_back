package main

import (
	"os"

	"github.com/K-Kizuku/Somniosus_back/apps/account/app/ui/grpc"
	"github.com/K-Kizuku/Somniosus_back/apps/account/config"
)

func main() {
	config.LoadEnv()
	grpc.StartServer()
	os.Exit(0)
}
