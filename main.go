package main

import (
	"./rcdaemon"
	"os"
	"strconv"
	"log"
)

func main() {
	redisHost, exists := os.LookupEnv("REDIS_HOST")
	if !exists {
		panic("REDIS_HOST env should be provided")
	}

	redisPortStr, exists := os.LookupEnv("REDIS_PORT")
	if !exists {
		redisPortStr = "6379"
	}

	redisPort, err := strconv.Atoi(redisPortStr)
	log.Printf("Using redis connection to %s:%d", redisHost, redisPort)

	server, err := rcdaemon.NewServer("", nil)
	if err != nil {
		panic(err)
	}

	// register handler
	server.RegisterFunc("get", rcdaemon.GetHandler)
	server.RegisterFunc("gets", rcdaemon.GetHandler)
	server.RegisterFunc("add", rcdaemon.AddHandler)
	server.RegisterFunc("set", rcdaemon.SetHandler)
	server.RegisterFunc("delete", rcdaemon.DeleteHandler)
	server.RegisterFunc("incr", rcdaemon.IncrHandler)
	server.RegisterFunc("flush_all", rcdaemon.FlushAllHandler)
	server.RegisterFunc("version", rcdaemon.VersionHandler)

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
