package main

import (
	"github.com/ivandhitya/sinau/mongodb/mongodb"
)

func main() {
	_ = mongodb.ConnectMongoDB("mongodb://<username>:<password>@junction.proxy.rlwy.net:24268")
}
