package main

import (
	"flag"
	"os"
)


type Config struct {
	host string;
	throttle bool;
}

const HOST_DEFAULT string = "harrysprojects.com"
const THROTTLE_DEFAULT bool = true


func ConfigFromArgs() *Config {

	host := flag.String("host", HOST_DEFAULT, "The target host to crawl.")
	throttle := flag.Bool("throttle", THROTTLE_DEFAULT, "Reduce requests to the host.")

	flag.Parse()

	return &Config{
		host: *host,
		throttle: *throttle,
	}
}

func ConfigFromEnv() *Config {

	host := HOST_DEFAULT
	throttle := THROTTLE_DEFAULT

	if os.Getenv("WC_HOST") != "" {
		host = os.Getenv("WC_HOST")
	}

	if os.Getenv("WC_THROTTLE") == "true" {
		throttle = true
	}

	if os.Getenv("WC_THROTTLE") == "false" {
		throttle = false
	}

	return &Config{
		host: host,
		throttle: throttle,
	}
}