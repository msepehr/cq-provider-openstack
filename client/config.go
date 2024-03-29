package client

// Provider Configuration

type Config struct {
	// here goes top level configuration for your provider
	// This object will be pass filled in depending on user's configuration
	ExampleConfig bool `yaml:"example_config"`

	// resources that user asked to fetch
	// each resource can have optional additional configurations
	Resources []struct {
		Name  string
		Other map[string]interface{} `yaml:",inline"`
	}
}

func (c Config) Example() string {
	return `configuration {
	// Optional. if you want to assume role to multiple account and fetch data from them
    //accounts "<YOUR ID>" {
	// Optional. Role ARN we want to assume when accessing this account
	// role_arn = <YOUR_ROLE_ARN>
	// }
	// Optional. by default assumes all regions
	// regions = ["us-east-1", "us-west-2"]
	// Optional. Enable AWS SDK debug logging.
       aws_debug = false  
	// The maximum number of times that a request will be retried for failures. Defaults to 5 retry attempts.
	// max_retries = 5
	// The maximum back off delay between attempts. The backoff delays exponentially with a jitter based on the number of attempts. Defaults to 60 seconds.
	// max_backoff = 30 
}
`
}
