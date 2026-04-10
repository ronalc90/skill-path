```
  ____  _    _ _ _ ____       _   _     
 / ___|| | _(_) | |  _ \ __ _| |_| |__  
 \___ \| |/ / | | | |_) / _` | __| '_ \ 
  ___) |   <| | | |  __/ (_| | |_| | | |
 |____/|_|\_\_|_|_|_|   \__,_|\__|_| |_|
```

# SkillPath

![Go](https://img.shields.io/badge/Go-1.22-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![Gin](https://img.shields.io/badge/Gin-Framework-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![GORM](https://img.shields.io/badge/GORM-ORM-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![Next.js](https://img.shields.io/badge/Next.js-14-000000?style=for-the-badge&logo=nextdotjs&logoColor=white)
![TypeScript](https://img.shields.io/badge/TypeScript-5-3178C6?style=for-the-badge&logo=typescript&logoColor=white)
![License](https://img.shields.io/badge/Licencia-MIT-green?style=for-the-badge)

> **Rutas de aprendizaje personalizadas para tu carrera en tecnologia -- como Duolingo para developers.**

SkillPath es una plataforma que guia a desarrolladores a traves de rutas de aprendizaje curadas, con seguimiento de progreso, evaluaciones de habilidades y visualizacion de avance en tiempo real.

---

## Caracteristicas

- **Rutas de aprendizaje curadas** -- Paths estructurados con milestones, recursos y orden logico de estudio.
- **Seguimiento de progreso** -- Trackea tu avance por milestone, recurso completado y porcentaje general.
- **Evaluaciones de habilidades** -- Quizzes al final de cada milestone para validar conocimientos.
- **Visualizacion de progreso** -- Dashboard con estadisticas, skill radar y graficos de avance.
- **Busqueda full-text** -- Encuentra paths y recursos por palabra clave.
- **4 rutas pre-cargadas** -- Backend Java, Frontend React, Data Python y Mobile Flutter.

---

## Rutas Pre-cargadas

| Ruta | Categoria | Dificultad | Descripcion |
|------|-----------|------------|-------------|
| Backend Java | Backend | Intermedio | Domina Java, Spring Boot y microservicios |
| Frontend React | Frontend | Principiante | De cero a React profesional con TypeScript |
| Data Python | Data | Intermedio | Analisis de datos, pandas, ML basico |
| Mobile Flutter | Mobile | Principiante | Apps nativas con Dart y Flutter |

---

## Tech Stack

| Capa | Tecnologia | Descripcion |
|------|------------|-------------|
| **Backend** | Go 1.22 + Gin | API REST de alto rendimiento |
| **ORM** | GORM | Mapeo objeto-relacional con migraciones automaticas |
| **Base de datos** | SQLite / PostgreSQL | SQLite para desarrollo, PostgreSQL para produccion |
| **Autenticacion** | JWT + bcrypt | Tokens seguros con hashing de passwords |
| **Frontend** | Next.js 14 + TypeScript | SSR, App Router, componentes React |
| **Estilos** | Tailwind CSS 3 | Utility-first CSS framework |

---

## Arquitectura

El backend sigue **Clean Architecture** con capas bien definidas:

```
Handler (HTTP) --> Service (Logica de negocio) --> Repository (Datos)
```

```
backend/
  cmd/server/          # Punto de entrada
  internal/
    config/            # Configuracion y variables de entorno
    dto/               # Data Transfer Objects
    handler/           # Controladores HTTP (+ tests)
    middleware/         # Auth JWT, logging, recovery
    model/             # Entidades GORM
    repository/        # Acceso a datos + seed
    router/            # Definicion de rutas
    service/           # Logica de negocio (+ tests)
  pkg/
    hash/              # Utilidades de hashing
    jwt/               # Generacion y validacion de tokens
```

---

## Como ejecutar localmente

### Prerequisitos

- Go 1.22+
- Node.js 18+
- npm o yarn

### Backend

```bash
cd backend

# Configurar variables de entorno
cp .env.example .env

# Descargar dependencias
go mod download

# Ejecutar servidor (puerto 8080)
go run cmd/server/main.go
```

### Frontend

```bash
cd frontend

# Instalar dependencias
npm install

# Ejecutar en modo desarrollo (puerto 3000)
npm run dev
```

Abre [http://localhost:3000](http://localhost:3000) en tu navegador.

---

## API Endpoints

### Autenticacion

| Metodo | Endpoint | Descripcion | Auth |
|--------|----------|-------------|------|
| `POST` | `/api/v1/auth/register` | Registrar nuevo usuario | No |
| `POST` | `/api/v1/auth/login` | Iniciar sesion | No |
| `GET` | `/api/v1/auth/me` | Obtener perfil actual | Si |
| `PUT` | `/api/v1/auth/profile` | Actualizar perfil | Si |

### Rutas de Aprendizaje

| Metodo | Endpoint | Descripcion | Auth |
|--------|----------|-------------|------|
| `GET` | `/api/v1/paths` | Listar todas las rutas | No |
| `GET` | `/api/v1/paths/categories` | Obtener categorias | No |
| `GET` | `/api/v1/paths/:slug` | Detalle de una ruta | No |

### Progreso

| Metodo | Endpoint | Descripcion | Auth |
|--------|----------|-------------|------|
| `POST` | `/api/v1/paths/:slug/start` | Iniciar una ruta | Si |
| `GET` | `/api/v1/paths/:slug/progress` | Ver progreso en una ruta | Si |
| `PUT` | `/api/v1/paths/:slug/milestones/:id/complete` | Completar milestone | Si |
| `PUT` | `/api/v1/paths/:slug/resources/:id/complete` | Completar recurso | Si |
| `GET` | `/api/v1/my-paths` | Mis rutas activas | Si |

### Evaluaciones

| Metodo | Endpoint | Descripcion | Auth |
|--------|----------|-------------|------|
| `GET` | `/api/v1/milestones/:id/assessment` | Obtener quiz | Si |
| `POST` | `/api/v1/milestones/:id/assessment/submit` | Enviar respuestas | Si |

### Dashboard y Busqueda

| Metodo | Endpoint | Descripcion | Auth |
|--------|----------|-------------|------|
| `GET` | `/api/v1/dashboard` | Dashboard con estadisticas | Si |
| `GET` | `/api/v1/search?q=term` | Buscar paths y recursos | No |
| `GET` | `/health` | Health check | No |

---

## Estructura del Proyecto

```
skill-path/
|-- backend/
|   |-- cmd/server/main.go       # Entry point
|   |-- internal/
|   |   |-- config/              # App configuration
|   |   |-- dto/                 # Request/Response DTOs
|   |   |-- handler/             # HTTP handlers + tests
|   |   |-- middleware/          # JWT auth, logging, recovery
|   |   |-- model/               # GORM entities
|   |   |-- repository/          # Data access layer + seed
|   |   |-- router/              # Route definitions
|   |   |-- service/             # Business logic + tests
|   |-- pkg/                     # Shared utilities
|   |-- Dockerfile
|   |-- docker-compose.yml
|   |-- Makefile
|-- frontend/
|   |-- src/
|   |   |-- app/                 # Next.js App Router pages
|   |   |-- components/          # React components
|   |   |-- lib/                 # API client, auth context
|   |-- index.html               # Landing page
|   |-- tailwind.config.ts
|-- docs/                        # GitHub Pages
|-- README.md
|-- LICENSE
```

---

## Autor

Desarrollado por **Ronald**.

## Licencia

Este proyecto esta bajo la licencia [MIT](LICENSE).
