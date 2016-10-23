package main

import (
	"encoding/json"
	"flag"
	"log"
	"os"
)

var (
	env    = flag.String("environment", "development", "the environment you are using")
	target = flag.String("target", "", `the config target you are planning to generate ("AWS")`)
)

func main() {
	//variables
	var cloudConfig CloudConfiguration

	//init
	flag.Parse()

	//input
	stat, _ := os.Stdin.Stat()

	if (stat.Mode() & os.ModeCharDevice) == 0 {
		jsonDecoder := json.NewDecoder(os.Stdin)
		handleError(jsonDecoder.Decode(&cloudConfig))
	} else {
		handleInputFromTerminal(&cloudConfig)
	}

	//output
	jsonEncoder := json.NewEncoder(os.Stdout)
	switch *target {
	case "aws", "AWS":
		handleError(jsonEncoder.Encode(cloudConfig.AWS))
	default:
		handleError(jsonEncoder.Encode(cloudConfig))
	}
}

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func handleInputFromTerminal(cloudConfig *CloudConfiguration) {
	var file string

	switch *env {
	case "staging":
		file = "configurations/configuration.staging.json"
		break
	case "production":
		file = "configurations/configuration.production.json"
		break
	default:
		file = "configurations/configuration.json"
	}

	//open/close file
	configFile, err := os.Open(file)
	defer configFile.Close()
	handleError(err)

	jsonDecoder := json.NewDecoder(configFile)
	handleError(jsonDecoder.Decode(cloudConfig))
}
