# Runbook de Operaciones - SkillPath

## Informacion General

| Campo | Valor |
|-------|-------|
| Servicio | SkillPath API |
| Puerto | 3005 |
| Health Check | `GET /health` |
| Metricas | `GET /metrics` |
| Repositorio | github.com/ronalc90/skill-path |

## Inicio Rapido

### Desarrollo Local

```bash
# Backend
cd backend
cp .env.example .env
go mod download
go run cmd/server/main.go

# Frontend
cd frontend
npm install
npm run dev
```

### Docker Compose

```bash
# Levantar todo
docker-compose up -d

# Verificar estado
docker-compose ps

# Ver logs
docker-compose logs -f backend
```

## Procedimientos Operativos

### Despliegue

#### Docker (Local/Staging)

```bash
# 1. Construir imagen
make docker

# 2. Levantar servicios
make docker-compose

# 3. Verificar health
curl http://localhost:3005/health
```

#### AWS ECS (Produccion)

```bash
# 1. Actualizar imagen
docker build -t ghcr.io/ronalc90/skill-path:v1.x.x -f backend/Dockerfile backend/
docker push ghcr.io/ronalc90/skill-path:v1.x.x

# 2. Actualizar task definition con nueva imagen
cd infra/terraform
terraform plan -var="container_image=ghcr.io/ronalc90/skill-path:v1.x.x"
terraform apply -var="container_image=ghcr.io/ronalc90/skill-path:v1.x.x"

# 3. Verificar despliegue
aws ecs describe-services --cluster skillpath-cluster --services skillpath-api
```

#### Kubernetes

```bash
# 1. Actualizar imagen en deployment
kubectl set image deployment/skillpath-api \
  skillpath-api=ghcr.io/ronalc90/skill-path:v1.x.x \
  -n skillpath

# 2. Verificar rollout
kubectl rollout status deployment/skillpath-api -n skillpath

# 3. Rollback si es necesario
kubectl rollout undo deployment/skillpath-api -n skillpath
```

## Troubleshooting

### La API no responde

1. Verificar que el contenedor esta corriendo:
   ```bash
   docker-compose ps
   # o
   kubectl get pods -n skillpath
   ```

2. Revisar logs:
   ```bash
   docker-compose logs backend
   # o
   kubectl logs -l app=skillpath -n skillpath
   ```

3. Verificar health check:
   ```bash
   curl -v http://localhost:3005/health
   ```

### Errores de base de datos

1. Verificar conectividad con PostgreSQL:
   ```bash
   docker-compose exec postgres pg_isready -U skillpath
   ```

2. Revisar logs de PostgreSQL:
   ```bash
   docker-compose logs postgres
   ```

3. Verificar que la base de datos existe:
   ```bash
   docker-compose exec postgres psql -U skillpath -d skillpath -c "\dt"
   ```

### Alto consumo de memoria/CPU

1. Verificar metricas en Prometheus:
   - Acceder a `http://localhost:9090`
   - Consulta: `skillpath_http_requests_in_flight`

2. Verificar en Grafana:
   - Acceder a `http://localhost:3001` (admin/admin)

3. Revisar HPA (Kubernetes):
   ```bash
   kubectl get hpa -n skillpath
   kubectl describe hpa skillpath-api-hpa -n skillpath
   ```

### Problemas de autenticacion JWT

1. Verificar que `JWT_SECRET` esta configurado:
   ```bash
   # Docker
   docker-compose exec backend env | grep JWT

   # Kubernetes
   kubectl get secret skillpath-secrets -n skillpath -o yaml
   ```

2. Verificar expiracion del token:
   - Los tokens expiran segun `JWT_EXPIRY_HOURS` (default: 72h)

## Escalamiento

### Horizontal (Kubernetes)

El HPA escala automaticamente basado en:
- CPU > 70% de utilizacion
- Memoria > 80% de utilizacion
- Rango: 2 a 10 replicas

### Horizontal (ECS)

Auto Scaling configurado:
- Target: 70% CPU
- Rango: 2 a 6 tareas

## Contacto

| Rol | Contacto |
|-----|----------|
| Autor | Ronald (@ronalc90) |
| Repositorio | github.com/ronalc90/skill-path |

## Autor

Desarrollado por **Ronald**.
