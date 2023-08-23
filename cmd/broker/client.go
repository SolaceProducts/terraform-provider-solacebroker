package broker

import (
	"os"
	"strings"
	"terraform-provider-solacebroker/cmd/command"
	"terraform-provider-solacebroker/internal/broker"
	"terraform-provider-solacebroker/internal/semp"
	"time"
)

func CliClient(url string) *semp.Client {
	username := terraform.StringWithDefaultFromEnv("username", true, "")
	password := terraform.StringWithDefaultFromEnv("password", false, "")
	bearerToken := terraform.StringWithDefaultFromEnv("bearer_token", false, "")
	retries, err := terraform.Int64WithDefaultFromEnv("retries", false, 10)
	if err != nil {
		terraform.LogCLIError("\nError: Unable to parse provider attribute. " + err.Error())
		os.Exit(1)
	}
	retryMinInterval, err := terraform.DurationWithDefaultFromEnv("retry_min_interval", false, 3*time.Second)
	if err != nil {
		terraform.LogCLIError("\nError: Unable to parse provider attribute. " + err.Error())
		os.Exit(1)
	}
	retryMaxInterval, err := terraform.DurationWithDefaultFromEnv("retry_max_interval", false, 30*time.Second)
	if err != nil {
		terraform.LogCLIError("\nError: Unable to parse provider attribute. " + err.Error())
		os.Exit(1)
	}
	requestTimeoutDuration, err := terraform.DurationWithDefaultFromEnv("request_timeout_duration", false, time.Minute)
	if err != nil {
		terraform.LogCLIError("\nError: Unable to parse provider attribute. " + err.Error())
		os.Exit(1)
	}
	requestMinInterval, err := terraform.DurationWithDefaultFromEnv("request_min_interval", false, 100*time.Millisecond)
	if err != nil {
		terraform.LogCLIError("\nError: Unable to parse provider attribute. " + err.Error())
		os.Exit(1)
	}
	insecure_skip_verify, err := terraform.BooleanWithDefaultFromEnv("insecure_skip_verify", false, false)
	if err != nil {
		terraform.LogCLIError("\nError: Unable to parse provider attribute. " + err.Error())
		os.Exit(1)
	}
	client := semp.NewClient(
		getFullSempAPIURL(url),
		insecure_skip_verify,
		semp.BasicAuth(username, password),
		semp.BearerToken(bearerToken),
		semp.Retries(uint(retries), retryMinInterval, retryMaxInterval),
		semp.RequestLimits(requestTimeoutDuration, requestMinInterval))
	return client
}

func getFullSempAPIURL(url string) string {
	url = strings.TrimSuffix(url, "/")
	baseBath := strings.TrimPrefix(broker.SempDetail.BasePath, "/")
	return url + "/" + baseBath
}
