# TERRATEST-SIMPLE

This repo contains a simple Terraform module to demonstrate how to use Terratest to write automated tests for Terraform code.

# examples Folder

This is this folder all the terraform related infra codes are present

# test Folder

This is the folder where all the terratest test scriprs are present with _test.go extention. As all the tests is written in GO. 

## Prerequisite 

1. Go should be installed (https://go.dev/doc/install)
1. Terraform should be present in system environment path (https://developer.hashicorp.com/terraform/downloads)
1. IDE should be present preferred VSCODE (https://code.visualstudio.com/)

## Running Terraform module manually

1. Go to the `examples` Folder
1. Check and verify Terraform is Correctly Present in System Run `terraform -version`.
1. Run `terraform init`.
1. Run `terraform apply`.
1. When you're done, run `terraform destroy`.

##  configure GO dependencies

1. Go to the `test` Folder
1. Check GO is present in your system `go version`
1. Initialize Go dependencies run `go mod init "test/terraform_example_test"` it will initiate all the required dependencies which are required to test our Terratest Test. After executing successfully `go.mod` file should be generated. 
1. For chcek sum the modules run `go mod tidy`, A `go.sum` File should be generated. 

## to run Terratest Tests
1.  To run the Tests run `go test -v -timeout 30m`. Or if you want to run a specific test then run `go test terraform_example_test.go`

## Integration with Github CI

This repo is integrated with the Github CI (Github Actions). Any commit made to this repo will result a Terratest test to be trigger. The yml configaration file is present in the below location:
.github/workflows/ci.yml




