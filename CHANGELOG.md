# Changelog

Todos los cambios notables del proyecto se documentan en este archivo.

El formato esta basado en [Keep a Changelog](https://keepachangelog.com/es/1.1.0/).

## [0.1.0] - 2025-01-15

### Agregado

#### Backend (Go/Gin)
- API RESTful con versionado `/api/v1`
- Sistema de autenticacion con JWT (registro, login, perfil)
- CRUD de rutas de aprendizaje con slug-based routing
- Sistema de hitos y recursos por ruta
- Motor de evaluacion con quizzes de opcion multiple
- Algoritmo de scoring con umbral de aprobacion del 70%
- Tracking de progreso por usuario (rutas, hitos, recursos)
- Dashboard con metricas agregadas y radar de habilidades
- Busqueda global de rutas y recursos
- Paginacion y filtros en listado de rutas
- Middleware de autenticacion JWT
- Middleware de logging y recovery
- Configuracion CORS
- Seed automatico con 4 rutas de aprendizaje pre-cargadas
- Docker multi-stage build
- Docker Compose con PostgreSQL
- Makefile con comandos utiles

#### Rutas Pre-cargadas
- Backend Developer con Java (8 hitos, 120h)
- Frontend Developer con React (7 hitos, 100h)
- Data Analyst con Python (6 hitos, 90h)
- Mobile Developer con Flutter (6 hitos, 85h)

#### Frontend (Next.js)
- Proyecto Next.js 14 con App Router
- TypeScript configurado
- Tailwind CSS integrado

#### Infraestructura
- GitHub Actions CI pipeline
- Dockerfile multi-stage optimizado
- Documentacion completa del proyecto
