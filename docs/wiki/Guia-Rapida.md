# Guia Rapida

Esta guia cubre la configuracion inicial y los primeros pasos para ejecutar SkillPath localmente.

## Prerequisitos

| Herramienta | Version | Verificar |
|-------------|---------|-----------|
| Go | 1.22+ | `go version` |
| Node.js | 18+ | `node --version` |
| npm | 9+ | `npm --version` |
| Git | 2.x | `git --version` |

## Paso 1: Clonar el Repositorio

```bash
git clone https://github.com/ronalc90/skill-path.git
cd skill-path
```

## Paso 2: Configurar el Backend

```bash
cd backend

# Copiar configuracion
cp .env.example .env

# Descargar dependencias de Go
go mod download

# Iniciar el servidor
go run cmd/server/main.go
```

Al iniciar por primera vez, el servidor:
1. Crea la base de datos SQLite (`skillpath.db`)
2. Ejecuta migraciones automaticas con GORM
3. Pobla la base de datos con 4 rutas de aprendizaje y sus assessments
4. Inicia en `http://localhost:3005`

## Paso 3: Verificar el Backend

```bash
# Health check
curl http://localhost:3005/health
# Respuesta: {"status":"ok"}

# Listar rutas
curl http://localhost:3005/api/v1/paths
```

## Paso 4: Configurar el Frontend

En otra terminal:

```bash
cd frontend

# Instalar dependencias
npm install

# Iniciar servidor de desarrollo
npm run dev
```

El frontend estara disponible en `http://localhost:3000`.

## Paso 5: Probar el Flujo Completo

### Registrar un usuario

```bash
curl -X POST http://localhost:3005/api/v1/auth/register \
  -H "Content-Type: application/json" \
  -d '{
    "email": "test@example.com",
    "password": "test123",
    "display_name": "Test User"
  }'
```

Guardar el `token` de la respuesta.

### Inscribirse en una ruta

```bash
curl -X POST http://localhost:3005/api/v1/paths/backend-developer-java/start \
  -H "Authorization: Bearer <tu-token>"
```

### Tomar un quiz

```bash
# Ver preguntas
curl http://localhost:3005/api/v1/milestones/1/assessment \
  -H "Authorization: Bearer <tu-token>"

# Enviar respuestas
curl -X POST http://localhost:3005/api/v1/milestones/1/assessment/submit \
  -H "Authorization: Bearer <tu-token>" \
  -H "Content-Type: application/json" \
  -d '{"answers": [{"question_id": 1, "selected_index": 0}, {"question_id": 2, "selected_index": 2}, {"question_id": 3, "selected_index": 2}]}'
```

## Alternativa: Docker Compose

Si prefieres usar Docker:

```bash
cd backend
docker compose up --build
```

Esto levanta PostgreSQL + el API en un solo comando.

## Comandos Make Utiles

```bash
cd backend
make help          # Ver todos los comandos
make run           # Ejecutar el servidor
make test          # Ejecutar tests
make test-coverage # Tests con reporte de cobertura
make lint          # Ejecutar linter
make build         # Compilar binario
make docker-up     # Levantar con Docker Compose
make docker-down   # Detener Docker Compose
```

## Problemas Comunes

### Error: "go: module not found"
Ejecutar `go mod download` o `go mod tidy` en el directorio `backend/`.

### Error: "port already in use"
Cambiar `SERVER_PORT` en el archivo `.env` o detener el proceso que usa el puerto.

### La base de datos esta vacia
Eliminar `skillpath.db` y reiniciar el servidor. El seed se ejecuta automaticamente si la base esta vacia.
