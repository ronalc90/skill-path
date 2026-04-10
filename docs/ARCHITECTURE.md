# Arquitectura

## Vision General

SkillPath sigue el patron de **Clean Architecture** adaptado a Go. La idea central es que las dependencias siempre apuntan hacia adentro: las capas externas conocen a las internas, pero nunca al reves.

## Diagrama de Capas

```
+-----------------------------------------------------------+
|                        HTTP Layer                          |
|  Router (Gin) --> Middleware --> Handler                   |
+-----------------------------------------------------------+
                         |
                         v
+-----------------------------------------------------------+
|                     Service Layer                          |
|  AuthService, PathService, ProgressService,               |
|  AssessmentService, DashboardService, SearchService       |
+-----------------------------------------------------------+
                         |
                         v
+-----------------------------------------------------------+
|                    Repository Layer                        |
|  UserRepository, PathRepository, ProgressRepository,      |
|  AssessmentRepository                                     |
+-----------------------------------------------------------+
                         |
                         v
+-----------------------------------------------------------+
|                    Database Layer                          |
|  GORM + SQLite (dev) / PostgreSQL (prod)                  |
+-----------------------------------------------------------+
```

## Capas en Detalle

### Handler (Capa de Presentacion)

Ubicacion: `internal/handler/`

Responsabilidades:
- Parsear y validar el input HTTP (JSON body, query params, path params)
- Llamar al servicio correspondiente
- Formatear la respuesta HTTP con codigos de estado apropiados
- Manejar errores de dominio y traducirlos a errores HTTP

Los handlers NO contienen logica de negocio. Si un handler necesita tomar una decision de negocio, esa logica debe estar en el servicio.

### Service (Capa de Negocio)

Ubicacion: `internal/service/`

Responsabilidades:
- Implementar toda la logica de negocio
- Orquestar llamadas a uno o mas repositorios
- Validar reglas de dominio (ej: umbral de aprobacion del 70%)
- Definir errores de dominio (`ErrEmailAlreadyExists`, `ErrNotEnrolled`, etc.)

Cada servicio recibe sus dependencias via constructor (inyeccion por constructor).

### Repository (Capa de Datos)

Ubicacion: `internal/repository/`

Responsabilidades:
- Encapsular todas las consultas a la base de datos
- Usar GORM para construir queries
- Manejar preloading de relaciones
- No contener logica de negocio

### Model (Entidades de Dominio)

Ubicacion: `internal/model/`

Define las estructuras de datos que representan las tablas de la base de datos:

```
User
  |-- id, email, password_hash, display_name, avatar_url, bio, github_url, linkedin_url

LearningPath
  |-- id, title, slug, description, difficulty, estimated_hours, category, icon_url, is_published
  +-- Milestones[] (1:N)
        |-- id, path_id, title, description, order_index, estimated_hours
        +-- Resources[] (1:N)
              |-- id, milestone_id, title, url, type, provider, is_free, estimated_minutes, order_index

UserProgress
  |-- id, user_id, path_id, started_at, completed_at, last_activity_at
  |-- MilestoneProgresses[] (1:N)
  +-- ResourceProgresses[] (1:N)

SkillAssessment
  |-- id, milestone_id, question, options (JSON), correct_answer, explanation

AssessmentResult
  |-- id, user_id, milestone_id, score, total_questions, passed_at
```

### DTO (Data Transfer Objects)

Ubicacion: `internal/dto/`

Estructuras separadas de los modelos para controlar exactamente que datos se envian y reciben por HTTP. Esto evita exponer campos internos como `password_hash` y permite validacion declarativa con tags de Gin.

### Middleware

Ubicacion: `internal/middleware/`

- **auth.go** - Valida tokens JWT del header `Authorization: Bearer <token>`
- **logging.go** - Registra metodo, ruta, codigo de respuesta y duracion de cada request
- **recovery.go** - Captura panics y responde con 500 en lugar de crashear

## Inyeccion de Dependencias

La DI se realiza manualmente en `cmd/server/main.go`:

```go
// Repositorios
userRepo := repository.NewUserRepository(db)
pathRepo := repository.NewPathRepository(db)
progressRepo := repository.NewProgressRepository(db)
assessmentRepo := repository.NewAssessmentRepository(db)

// Servicios (reciben repositorios)
authService := service.NewAuthService(userRepo, cfg.JWTSecret, cfg.JWTExpiryHrs)
pathService := service.NewPathService(pathRepo)
progressService := service.NewProgressService(progressRepo, pathRepo)

// Handlers (reciben servicios)
authHandler := handler.NewAuthHandler(authService)
pathHandler := handler.NewPathHandler(pathService)
```

Esta aproximacion es intencional: en Go es preferible la DI explicita sobre frameworks de inyeccion como Wire o Dig. El grafo de dependencias es claro y las importaciones circulares se detectan en compilacion.

## Modelos GORM

GORM maneja las migraciones automaticamente con `db.AutoMigrate()`. Los modelos usan:

- `gorm:"primaryKey"` para claves primarias auto-incrementales
- `gorm:"uniqueIndex"` para constraints de unicidad (email, slug)
- `gorm:"foreignKey:X"` para relaciones entre tablas
- `JSONSlice` tipo personalizado que serializa `[]string` como JSON en la columna
- `gorm.DeletedAt` para soft-delete en usuarios

## Patron de Preloading

Para consultas que necesitan datos relacionados, se usa el preloading de GORM:

```go
db.Where("slug = ?", slug).
    Preload("Milestones", func(db *gorm.DB) *gorm.DB {
        return db.Order("order_index ASC")
    }).
    Preload("Milestones.Resources", func(db *gorm.DB) *gorm.DB {
        return db.Order("order_index ASC")
    }).
    First(&path)
```

Esto genera queries separadas (no JOINs) pero mantiene la API de consulta simple y predecible.

## Paquetes Compartidos (pkg/)

- `pkg/hash` - Funciones de hashing con bcrypt
- `pkg/jwt` - Generacion y validacion de tokens JWT con HS256

Estos paquetes son reutilizables y no dependen de ninguna capa interna.
