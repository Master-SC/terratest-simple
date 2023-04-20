package test

import (
	"fmt"
	"os"
	"testing"
	"time"

	http_helper "github.com/gruntwork-io/terratest/modules/http-helper"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/magiconair/properties"
	"github.com/stretchr/testify/assert"
)

func TestTerraformExample(t *testing.T) {

	//get the current directory path and concate with external properties file path
	mydir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}

	propfileName := "/resource/data.properties"
	path := mydir + propfileName

	// Load the properties File & get the value
	p := properties.MustLoadFile(path, properties.UTF8)
	expectedBucketName := p.GetString("bucket.name", "null") // Value for Test1 Assertion
	body := p.GetString("response.body", "null")             //Value for Test2 Response Body
	responseCode := p.GetInt("response.status", 0)           //Value for Test2 status code

	t.Parallel()
	// retryable errors in terraform testing.
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../examples",
	})

	defer terraform.Destroy(t, terraformOptions)

	terraform.InitAndApply(t, terraformOptions)

	//Test#1 Baics Test to Assert a value.

	actualBucketName := terraform.Output(t, terraformOptions, "my_bucket_name")
	assert.Equal(t, expectedBucketName, actualBucketName)

	//Test#2 To hit a url and check the status code along with response body.

	//func HttpGetWithRetry
	//func HttpGetWithRetry(t testing.TestingT, url string, tlsConfig *tls.Config, expectedStatus int, expectedBody string, retries int, sleepBetweenRetries time.Duration)
	url := terraform.Output(t, terraformOptions, "demourl")
	http_helper.HttpGetWithRetry(t, url, nil, responseCode, body, 5, 5*time.Second)
}
