job "example" {
  datacenters = ["dc1"]
  type        = "service"

  # This attribute doesn't exist in the job schema
  invalid_attribute_that_should_error = "test"

  group "app" {
    count = 1

    task "server" {
      driver = "exec"

      config {
        command = "/bin/echo"
      }
    }
  }
}
