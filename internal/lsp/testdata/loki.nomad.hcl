variable "image" {
  type    = string
  default = "grafana/loki:latest"
}

variable "s3_bucket" {
  type = string
}

job "loki" {
  meta {
    s3_bucket = var.s3_bucket
  }

  group "loki" {
    constraint {
      attribute = "${meta.role}"
      operator  = "set_contains"
      value     = "monitoring"
    }

    network {
      port "http" {
        static = 3100
        to     = 3100
      }
    }

    service {
      name     = "loki"
      port     = "http"
      provider = "nomad"
    }

    ephemeral_disk {
      size = 200
    }

    volume "loki-data" {
      type            = "csi"
      source          = "loki-volume"
      access_mode     = "single-node-writer"
      attachment_mode = "file-system"
    }

    task "prep-disk" {
      driver = "docker"

      volume_mount {
        volume      = "loki-data"
        destination = "/loki-data/"
      }

      config {
        image   = "busybox:latest"
        command = "sh"
        args    = ["-c", "chown -R 1000:1000 /loki-data/"]
      }

      resources {
        cpu    = 200
        memory = 128
      }

      lifecycle {
        hook    = "prestart"
        sidecar = false
      }
    }

    task "loki" {
      driver = "docker"
      user   = "1000"

      config {
        image = var.image
        ports = ["http"]
        args  = ["--config.file=/etc/loki/config/loki.yml"]
        volumes = [
          "local/config/loki.yaml:/etc/loki/config/loki.yml",
        ]
      }

      volume_mount {
        volume      = "loki-data"
        destination = "/loki"
      }

      template {
        data        = file("./deployment/jobs/configs/loki.yml.tmpl")
        destination = "${NOMAD_TASK_DIR}/config/loki.yaml"
      }

      resources {
        cpu    = 200
        memory = 256
      }
    }
  }
}
