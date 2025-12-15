variables {
  app_name = "example-app"
  version  = "1.0.0"
}

job "example" {
  datacenters = ["dc1"]
  type        = "service"

  meta {
		owner = "dev-team"
  }

  group "app" {
    count = 1

    task "server" {
      driver = "docker"

      config {
        image = "${var.app_name}:${var.version}"
      }

      resources {
        cpu    = 500
        memory = 256
      }
    }
  }
}
