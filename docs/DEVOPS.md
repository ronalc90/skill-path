# DevOps - SkillPath

## Descripcion General

Este documento describe la infraestructura de DevOps para SkillPath, incluyendo CI/CD, contenedorizacion, orquestacion, monitoreo y seguridad.

## Arquitectura de Infraestructura

```
                    +-------------------+
                    |   GitHub Actions  |
                    |   (CI/CD)         |
                    +--------+----------+
                             |
                    +--------v----------+
                    |   GHCR             |
                    |   (Container       |
                    |    Registry)       |
                    +--------+----------+
                             |
              +--------------+--------------+
              |                             |
     +--------v----------+       +----------v--------+
     |   AWS ECS          |       |   Kubernetes      |
     |   (Fargate)        |       |   (Alternativa)   |
     +--------+----------+       +-------------------+
              |
     +--------v----------+
     |   ALB              |
     |   (Load Balancer)  |
     +-------------------+
```

## CI/CD (GitHub Actions)

### Pipeline de CI (`ci.yml`)

Se ejecuta en push a `main` y en pull requests:

| Job | Descripcion |
|-----|-------------|
| `backend-lint` | golangci-lint sobre el codigo Go |
| `backend-test` | Tests con race detector y cobertura |
| `backend-build` | Compilacion del binario |
| `frontend` | npm ci, lint y build de Next.js |
| `docker` | Build multi-arch (amd64/arm64) y push a GHCR |
| `trivy-scan` | Escaneo de vulnerabilidades de la imagen Docker |

### Pipeline de Seguridad (`security.yml`)

Se ejecuta en push, PRs y semanalmente (lunes 6:00 AM):

| Job | Descripcion |
|-----|-------------|
| `gosec` | Analisis estatico de seguridad para Go |
| `trivy-fs` | Escaneo de vulnerabilidades del filesystem |
| `golangci-lint` | Lint con reglas de seguridad habilitadas |
| `dependency-review` | Revision de dependencias en PRs |

## Docker

### Imagen del Backend

Build multi-stage con Alpine Linux:

```bash
# Construir imagen
make docker

# Levantar todos los servicios
make docker-compose

# Ver logs
make docker-logs
```

### Servicios (docker-compose.yml)

| Servicio | Puerto | Descripcion |
|----------|--------|-------------|
| `backend` | 3005 | API Go/Gin |
| `postgres` | 5432 | Base de datos PostgreSQL 16 |
| `prometheus` | 9090 | Recoleccion de metricas |
| `grafana` | 3001 | Dashboards de monitoreo |

## Kubernetes

Los manifiestos se encuentran en `infra/kubernetes/`:

| Archivo | Descripcion |
|---------|-------------|
| `namespace.yml` | Namespace dedicado `skillpath` |
| `configmap.yml` | Configuracion no sensible |
| `secret.yml` | Secretos (JWT, etc.) |
| `deployment.yml` | Deployment con 2 replicas, probes, limites |
| `service.yml` | Service ClusterIP |
| `ingress.yml` | Ingress con TLS y rate limiting |
| `hpa.yml` | Autoscaling basado en CPU/memoria |
| `pvc.yml` | Persistent Volume Claim para datos |

### Despliegue en Kubernetes

```bash
# Aplicar todos los manifiestos
make k8s-apply

# Verificar el estado
kubectl get all -n skillpath

# Eliminar recursos
make k8s-delete
```

## Terraform (AWS ECS)

Infraestructura como codigo en `infra/terraform/`:

### Recursos creados

- VPC con subnets publicas
- ECS Cluster con Fargate
- Application Load Balancer
- Auto Scaling (2-6 instancias)
- CloudWatch Logs
- SSM Parameter Store para secretos
- Security Groups

### Uso

```bash
# Configurar variables
cp infra/terraform/terraform.tfvars.example infra/terraform/terraform.tfvars
# Editar terraform.tfvars con valores reales

# Inicializar
make tf-init

# Planificar
make tf-plan

# Aplicar
make tf-apply
```

## Seguridad

### Escaneos automatizados

```bash
# Ejecutar todos los checks de seguridad
make security

# Solo escaneo de imagen Docker
make security-scan

# Solo analisis de codigo Go
make gosec
```

### Practicas implementadas

- Imagen Docker con usuario no-root
- Health checks en Docker y Kubernetes
- Secretos manejados via SSM Parameter Store / K8s Secrets
- Rate limiting en ingress
- Escaneo semanal automatico de vulnerabilidades
- Revision de dependencias en PRs

## Autor

Desarrollado por **Ronald**.
