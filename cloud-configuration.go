package main

import awsConfiguration "github.com/raymondsquared/common-utilities/go/cloud/aws/configuration"

//CloudConfiguration represents Cloud configuration file
type CloudConfiguration struct {
	AWS awsConfiguration.Configuration
}
