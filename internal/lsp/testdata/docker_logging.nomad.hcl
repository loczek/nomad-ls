job "docker-logging-test" {
  datacenters = ["dc1"]
  type        = "service"

  group "app" {
    count = 1

    task "server" {
      driver = "docker"

      config {
        image = "nginx:latest"

        logging {
          type = "fluentd"
          config {
            fluentd-address = "localhost:24224"
            tag             = "docker.nginx"
            fluentd-async   = "true"
          }
        }
      }

      resources {
        cpu    = 500
        memory = 256
      }
    }
  }
}
