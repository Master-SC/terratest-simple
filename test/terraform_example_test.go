package test

import (
	"fmt"
	"os"
	"testing"

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
	v := p.GetString("data", "null")

	// retryable errors in terraform testing.
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../examples",
	})

	defer terraform.Destroy(t, terraformOptions)

	terraform.InitAndApply(t, terraformOptions)

	output := terraform.Output(t, terraformOptions, "my_ip_address")
	assert.Equal(t, v, output)
}
