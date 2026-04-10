variable "aws_region" {
  description = "Region de AWS para el despliegue"
  type        = string
  default     = "us-east-1"
}

variable "project_name" {
  description = "Nombre del proyecto"
  type        = string
  default     = "skillpath"
}

variable "environment" {
  description = "Entorno de despliegue (dev, staging, prod)"
  type        = string
  default     = "prod"
}

variable "container_image" {
  description = "URI de la imagen Docker del backend"
  type        = string
  default     = "ghcr.io/ronalc90/skill-path:latest"
}

variable "container_port" {
  description = "Puerto del contenedor"
  type        = number
  default     = 3005
}

variable "task_cpu" {
  description = "CPU units para la tarea ECS (1024 = 1 vCPU)"
  type        = string
  default     = "256"
}

variable "task_memory" {
  description = "Memoria en MB para la tarea ECS"
  type        = string
  default     = "512"
}

variable "desired_count" {
  description = "Numero deseado de instancias del servicio"
  type        = number
  default     = 2
}

variable "max_count" {
  description = "Numero maximo de instancias (auto scaling)"
  type        = number
  default     = 6
}

variable "jwt_secret" {
  description = "Secreto JWT para firmar tokens"
  type        = string
  sensitive   = true
}

variable "cors_origins" {
  description = "Origenes CORS permitidos"
  type        = string
  default     = "https://skillpath.example.com"
}
