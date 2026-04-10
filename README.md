# SkillPath

[![Go](https://img.shields.io/badge/Go-1.22-00ADD8?style=flat&logo=go&logoColor=white)](https://go.dev/)
[![Gin](https://img.shields.io/badge/Gin-Framework-00ADD8?style=flat&logo=go&logoColor=white)](https://gin-gonic.com/)
[![GORM](https://img.shields.io/badge/GORM-ORM-00ADD8?style=flat&logo=go&logoColor=white)](https://gorm.io/)
[![Next.js](https://img.shields.io/badge/Next.js-14-000000?style=flat&logo=next.js&logoColor=white)](https://nextjs.org/)
[![TypeScript](https://img.shields.io/badge/TypeScript-5-3178C6?style=flat&logo=typescript&logoColor=white)](https://www.typescriptlang.org/)
[![License](https://img.shields.io/badge/License-MIT-green.svg)](LICENSE)

> Rutas de aprendizaje personalizadas para developers. Planifica tu carrera tech con caminos estructurados, hitos medibles y evaluaciones de conocimiento.

---

## Sobre el Proyecto

**SkillPath** es una plataforma de aprendizaje que ofrece rutas estructuradas para carreras tecnologicas. Cada ruta esta compuesta por hitos (milestones) con recursos curados, y un motor de evaluacion que permite validar el conocimiento adquirido mediante quizzes interactivos.

El objetivo es resolver un problema comun: la falta de estructura al aprender a programar. En lugar de saltar entre tutoriales aleatorios, SkillPath ofrece un camino claro con progreso medible.

## Caracteristicas

- **Rutas de aprendizaje curadas** con recursos de alta calidad (cursos, libros, videos, tutoriales)
- **Sistema de hitos** que divide cada ruta en etapas manejables
- **Motor de evaluacion** con quizzes por hito y algoritmo de scoring (umbral 70%)
- **Dashboard personalizado** con metricas de progreso, radar de habilidades y actividad reciente
- **Autenticacion JWT** con registro, login y perfiles de usuario
- **Busqueda global** de rutas y recursos
- **API RESTful** versionada con paginacion y filtros
- **Arquitectura limpia** en Go con separacion de capas

## Rutas Disponibles

| Ruta | Categoria | Dificultad | Horas Est. | Hitos |
|------|-----------|------------|------------|-------|
| Backend Developer con Java | Backend | Intermedio | 120h | 8 (Fundamentos Java, Spring Boot, JPA/Hibernate, Spring Security, Testing, Microservicios, Docker/K8s, CI/CD) |
| Frontend Developer con React | Frontend | Principiante | 100h | 7 (HTML/CSS/JS, TypeScript, React Core, Estado Global, Next.js, Testing, Performance) |
| Data Analyst con Python | Data | Principiante | 90h | 6 (Python, NumPy/Pandas, Visualizacion, SQL, Estadistica, ML Basico) |
| Mobile Developer con Flutter | Mobile | Intermedio | 85h | 6 (Dart, Flutter UI, State Management, Networking, Persistencia, Testing/Deploy) |

## Arquitectura

El backend sigue **Clean Architecture** en Go con inyeccion de dependencias manual:

```
HTTP Request
    |
    v
+-------------------+
|    Handler         |  <-- Validacion de input, HTTP responses
+-------------------+
    |
    v
+-------------------+
|    Service         |  <-- Logica de negocio, reglas de dominio
+-------------------+
    |
    v
+-------------------+
|    Repository      |  <-- Acceso a datos, queries GORM
+-------------------+
    |
    v
+-------------------+
|    Database        |  <-- SQLite (dev) / PostgreSQL (prod)
+-------------------+
```

Cada capa depende solo de la capa inmediatamente inferior. Los handlers no conocen la base de datos, y los repositorios no conocen HTTP.

## Stack Tecnologico

### Backend
- **Go 1.22** - Lenguaje de programacion
- **Gin** - Framework HTTP
- **GORM** - ORM para Go
- **SQLite** - Base de datos en desarrollo
- **PostgreSQL** - Base de datos en produccion
- **JWT** (golang-jwt) - Autenticacion
- **bcrypt** - Hash de contrasenas

### Frontend
- **Next.js 14** - Framework React con App Router
- **TypeScript 5** - Tipado estatico
- **Tailwind CSS 3** - Estilos utilitarios
- **React 18** - Biblioteca de UI

### Infraestructura
- **Docker** - Contenedorizacion (multi-stage build)
- **Docker Compose** - Orquestacion local
- **GitHub Actions** - CI/CD
- **Railway** - Deploy del backend
- **Vercel** - Deploy del frontend

## Instalacion

### Prerequisitos

- Go 1.22+
- Node.js 18+
- npm 9+

### Backend

```bash
cd backend

# Copiar variables de entorno
cp .env.example .env

# Descargar dependencias
go mod download

# Ejecutar (crea SQLite automaticamente + seed de datos)
go run cmd/server/main.go

# O usando Make
make run
```

El servidor inicia en `http://localhost:3005` con datos pre-cargados.

### Frontend

```bash
cd frontend

# Instalar dependencias
npm install

# Iniciar servidor de desarrollo
npm run dev
```

El frontend inicia en `http://localhost:3000`.

### Docker (todo junto)

```bash
cd backend
docker compose up --build
```

## Variables de Entorno

| Variable | Descripcion | Valor por defecto |
|----------|-------------|-------------------|
| `SERVER_PORT` | Puerto del servidor | `3005` |
| `GIN_MODE` | Modo de Gin (debug/release) | `debug` |
| `DB_PATH` | Ruta de la base de datos SQLite | `skillpath.db` |
| `JWT_SECRET` | Secreto para firmar tokens JWT | `default-secret-change-me` |
| `JWT_EXPIRY_HOURS` | Horas de validez del token | `72` |
| `CORS_ORIGINS` | Origenes permitidos (separados por coma) | `http://localhost:3000` |

## API Documentation

Base URL: `http://localhost:3005/api/v1`

### Endpoints Publicos

| Metodo | Ruta | Descripcion |
|--------|------|-------------|
| `POST` | `/auth/register` | Registro de usuario |
| `POST` | `/auth/login` | Inicio de sesion |
| `GET` | `/paths` | Listar rutas (paginado, filtros) |
| `GET` | `/paths/categories` | Listar categorias |
| `GET` | `/paths/:slug` | Detalle de ruta con hitos y recursos |
| `GET` | `/search?q=` | Busqueda global |
| `GET` | `/health` | Health check |

### Endpoints Protegidos (requieren Bearer Token)

| Metodo | Ruta | Descripcion |
|--------|------|-------------|
| `GET` | `/auth/me` | Perfil del usuario |
| `PUT` | `/auth/profile` | Actualizar perfil |
| `POST` | `/paths/:slug/start` | Inscribirse en una ruta |
| `GET` | `/paths/:slug/progress` | Ver progreso en una ruta |
| `PUT` | `/paths/:slug/milestones/:id/complete` | Completar hito |
| `PUT` | `/paths/:slug/resources/:id/complete` | Completar recurso |
| `GET` | `/my-paths` | Rutas inscritas del usuario |
| `GET` | `/milestones/:id/assessment` | Obtener quiz de un hito |
| `POST` | `/milestones/:id/assessment/submit` | Enviar respuestas del quiz |
| `GET` | `/dashboard` | Dashboard con metricas |

Documentacion completa con ejemplos curl en [`docs/API.md`](docs/API.md).

## Estructura del Proyecto

```
skill-path/
|-- backend/
|   |-- cmd/
|   |   +-- server/
|   |       +-- main.go              # Punto de entrada, DI manual
|   |-- internal/
|   |   |-- config/
|   |   |   +-- config.go            # Configuracion desde env vars
|   |   |-- dto/
|   |   |   +-- dto.go               # Data Transfer Objects
|   |   |-- handler/
|   |   |   |-- auth_handler.go      # Endpoints de autenticacion
|   |   |   |-- path_handler.go      # Endpoints de rutas
|   |   |   |-- progress_handler.go  # Endpoints de progreso
|   |   |   |-- assessment_handler.go# Endpoints de evaluaciones
|   |   |   |-- dashboard_handler.go # Endpoint de dashboard
|   |   |   +-- search_handler.go    # Endpoint de busqueda
|   |   |-- middleware/
|   |   |   |-- auth.go              # JWT middleware
|   |   |   |-- logging.go           # Request logging
|   |   |   +-- recovery.go          # Panic recovery
|   |   |-- model/
|   |   |   +-- models.go            # Modelos GORM
|   |   |-- repository/
|   |   |   |-- repository.go        # Repositorios de datos
|   |   |   +-- seed.go              # Datos iniciales
|   |   |-- router/
|   |   |   +-- router.go            # Configuracion de rutas
|   |   +-- service/
|   |       |-- auth_service.go      # Logica de autenticacion
|   |       |-- path_service.go      # Logica de rutas
|   |       |-- progress_service.go  # Logica de progreso
|   |       |-- assessment_service.go# Logica de evaluaciones
|   |       |-- dashboard_service.go # Logica de dashboard
|   |       +-- search_service.go    # Logica de busqueda
|   |-- pkg/
|   |   |-- hash/
|   |   |   +-- hash.go              # bcrypt helpers
|   |   +-- jwt/
|   |       +-- jwt.go               # JWT helpers
|   |-- Dockerfile
|   |-- docker-compose.yml
|   |-- Makefile
|   +-- go.mod
|-- frontend/
|   |-- src/
|   |   +-- app/
|   |-- package.json
|   |-- next.config.mjs
|   |-- tailwind.config.ts
|   +-- tsconfig.json
|-- docs/
|   |-- ARCHITECTURE.md
|   |-- API.md
|   |-- LEARNING-PATHS.md
|   |-- ASSESSMENT-ENGINE.md
|   |-- DEPLOYMENT.md
|   |-- DECISIONS.md
|   +-- wiki/
+-- README.md
```

## Decisiones Tecnicas

| Decision | Razon |
|----------|-------|
| **Go sobre Node.js** | Rendimiento superior, binarios estaticos, concurrencia nativa con goroutines |
| **Clean Architecture** | Separacion de responsabilidades, testabilidad, mantenibilidad a largo plazo |
| **GORM sobre sqlx** | Productividad: migraciones automaticas, relaciones, preloading. Trade-off aceptable en control SQL |
| **Gin sobre Echo** | Ecosistema mas grande, mejor documentacion, rendimiento comparable |
| **SQLite para desarrollo** | Zero-config, archivo unico, ideal para desarrollo local y testing |
| **Slug-based routing** | URLs legibles (`/paths/backend-developer-java`) en lugar de IDs numericos |
| **DAG para skill tree** | Los hitos siguen un orden (Directed Acyclic Graph), permitiendo prerrequisitos futuros |

Documentacion detallada de ADRs en [`docs/DECISIONS.md`](docs/DECISIONS.md).

## Testing

```bash
cd backend

# Ejecutar todos los tests
make test

# Tests con cobertura
make test-coverage

# Tests especificos
go test -v ./internal/handler/...
go test -v ./internal/service/...
```

Los tests incluyen:
- Tests unitarios de handlers con mocks
- Tests de servicios con logica de negocio
- Validacion de DTOs y reglas de scoring

## Despliegue

### Backend (Railway)
```bash
# Railway detecta el Dockerfile automaticamente
railway up
```

### Frontend (Vercel)
```bash
cd frontend
vercel
```

Documentacion completa en [`docs/DEPLOYMENT.md`](docs/DEPLOYMENT.md).

## Roadmap

- [ ] Sistema de prerrequisitos entre hitos (DAG completo)
- [ ] Gamificacion: badges y puntos de experiencia
- [ ] Recomendaciones personalizadas basadas en progreso
- [ ] Creacion de rutas por la comunidad
- [ ] Integracion con GitHub para validar proyectos
- [ ] App movil con React Native
- [ ] Sistema de mentoria entre usuarios
- [ ] Exportacion de certificados PDF

## Autor

Desarrollado por **Ronald**.

- GitHub: [@ronalc90](https://github.com/ronalc90)
