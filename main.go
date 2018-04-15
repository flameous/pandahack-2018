package main

import (
	_ "github.com/lib/pq"
	"log"
	"github.com/flameous/pandahack-2018/db"
	"github.com/flameous/pandahack-2018/server"
	"os"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func main() {
	d, err := db.NewDatabase(
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_ADDR"),
		os.Getenv("POSTGRES_DB"),
	)
	if err != nil {
		log.Fatal(err)
	}
	s := server.NewServer(d, os.Getenv("APP_PORT"))
	s.Serve()
}
