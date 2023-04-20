terraform {
  required_version = ">= 0.12.26"
}

output "my_bucket_name" {
    value="test_bucket_123"
}

output "demourl" {
  value = "https://reqres.in/api/users/2"
}
