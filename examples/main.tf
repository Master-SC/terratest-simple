terraform {
  required_version = ">= 0.12.26"
}

output "my_ip_address" {
    value="192.168.1.13"
}

output "demourl" {
  value = "https://reqres.in/api/users/2"
}
