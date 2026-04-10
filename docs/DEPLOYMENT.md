# Despliegue

## Arquitectura de Produccion

```
Internet
   |
   v
[Vercel] ----------> Frontend (Next.js)
   |                      |
   |                      | API calls
   |                      v
[Railway] ----------> Backend (Go/Gin)
                          |
                          v
                    [PostgreSQL]
```

---

## Backend - Docker Multi-Stage Build

El backend usa un Dockerfile con build multi-stage para producir una imagen minima:

### Stage 1: Build

```dockerfile
FROM golang:1.22-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /skillpath-api ./cmd/server
```

- Imagen base `golang:1.22-alpine` para compilar
- `CGO_ENABLED=0` produce un binario estatico (sin dependencia de libc)
- El binario resultante es autonomo

### Stage 2: Runtime

```dockerfile
FROM alpine:3.19
RUN apk --no-cache add ca-certificates tzdata
WORKDIR /app
COPY --from=builder /skillpath-api .
EXPOSE 8080
CMD ["./skillpath-api"]
```

- Imagen final ~15MB (alpine + binario)
- Solo incluye certificados TLS y timezone data
- No contiene toolchain de Go ni codigo fuente

### Build y Ejecucion Local

```bash
cd backend

# Construir imagen
docker build -t skillpath-api .

# Ejecutar
docker run -p 8080:8080 \
  -e JWT_SECRET=mi-secreto-seguro \
  -e SERVER_PORT=8080 \
  skillpath-api
```

---

## Docker Compose (Desarrollo)

Para desarrollo local con PostgreSQL:

```bash
cd backend
docker compose up --build
```

Esto levanta:
- **PostgreSQL 16** en puerto 5432
- **SkillPath API** en puerto 8080

Variables de entorno configuradas en `docker-compose.yml`.

---

## Backend en Railway

### Prerequisitos
- Cuenta en [Railway](https://railway.app/)
- CLI de Railway instalado (`npm i -g @railway/cli`)

### Pasos

1. **Login**
   ```bash
   railway login
   ```

2. **Crear proyecto**
   ```bash
   cd backend
   railway init
   ```

3. **Agregar PostgreSQL**
   - Desde el dashboard de Railway, agregar un servicio PostgreSQL
   - Railway provee automaticamente las variables `DATABASE_URL`

4. **Configurar variables de entorno**
   ```bash
   railway variables set JWT_SECRET=tu-secreto-seguro-de-produccion
   railway variables set SERVER_PORT=8080
   railway variables set GIN_MODE=release
   railway variables set CORS_ORIGINS=https://tu-frontend.vercel.app
   ```

5. **Deploy**
   ```bash
   railway up
   ```

Railway detecta el `Dockerfile` automaticamente y construye la imagen.

### Variables de Entorno en Railway

| Variable | Valor |
|----------|-------|
| `SERVER_PORT` | `8080` |
| `GIN_MODE` | `release` |
| `JWT_SECRET` | (generar un secreto fuerte) |
| `CORS_ORIGINS` | URL del frontend en Vercel |
| `DB_HOST` | (provisto por Railway) |
| `DB_PORT` | (provisto por Railway) |
| `DB_USER` | (provisto por Railway) |
| `DB_PASSWORD` | (provisto por Railway) |
| `DB_NAME` | (provisto por Railway) |

---

## Frontend en Vercel

### Prerequisitos
- Cuenta en [Vercel](https://vercel.com/)
- CLI de Vercel (`npm i -g vercel`)

### Pasos

1. **Deploy desde CLI**
   ```bash
   cd frontend
   vercel
   ```

2. **O conectar repositorio en Vercel**
   - Ir a [vercel.com/new](https://vercel.com/new)
   - Importar el repositorio desde GitHub
   - Configurar Root Directory como `frontend`
   - Framework Preset: Next.js

3. **Variables de entorno**
   ```
   NEXT_PUBLIC_API_URL=https://tu-backend.railway.app/api/v1
   ```

### Configuracion de Vercel

```json
{
  "buildCommand": "npm run build",
  "outputDirectory": ".next",
  "installCommand": "npm ci",
  "framework": "nextjs"
}
```

---

## CI/CD con GitHub Actions

El pipeline en `.github/workflows/ci.yml` ejecuta automaticamente:

1. **Backend:** `go vet` + `go test` con cache de modulos
2. **Frontend:** `npm ci` + `npm run build`

Se ejecuta en cada push a `main` y en cada pull request.

---

## Checklist de Produccion

- [ ] Cambiar `JWT_SECRET` a un valor aleatorio seguro (minimo 32 caracteres)
- [ ] Configurar `GIN_MODE=release`
- [ ] Configurar `CORS_ORIGINS` con el dominio exacto del frontend
- [ ] Usar PostgreSQL en lugar de SQLite
- [ ] Configurar HTTPS (Railway y Vercel lo proveen automaticamente)
- [ ] Configurar dominio personalizado (opcional)
- [ ] Configurar monitoreo y alertas
- [ ] Revisar rate limiting en endpoints publicos
