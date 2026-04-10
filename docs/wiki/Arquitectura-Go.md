# Arquitectura Go

SkillPath sigue Clean Architecture en Go con tres capas principales y inyeccion de dependencias manual.

## Diagrama de Capas

```
                    HTTP Request
                         |
                         v
+----------------------------------------------------+
|  Router (Gin)                                       |
|  - Configura rutas y middleware                     |
|  - CORS, logging, recovery, auth                   |
+----------------------------------------------------+
                         |
                         v
+----------------------------------------------------+
|  Handler Layer                                      |
|  - Parsea input (JSON body, query params, headers) |
|  - Valida con tags de Gin (binding)                |
|  - Llama al servicio                               |
|  - Formatea respuesta HTTP                         |
+----------------------------------------------------+
                         |
                         v
+----------------------------------------------------+
|  Service Layer                                      |
|  - Logica de negocio                               |
|  - Reglas de dominio (scoring 70%, etc.)           |
|  - Orquesta llamadas a repositorios               |
|  - Define errores de dominio                       |
+----------------------------------------------------+
                         |
                         v
+----------------------------------------------------+
|  Repository Layer                                   |
|  - Queries GORM                                    |
|  - Preloading de relaciones                        |
|  - Paginacion                                      |
+----------------------------------------------------+
                         |
                         v
+----------------------------------------------------+
|  Database (GORM)                                    |
|  - SQLite (dev) / PostgreSQL (prod)                |
|  - AutoMigrate para esquema                        |
+----------------------------------------------------+
```

## Regla de Dependencia

Las dependencias van **siempre hacia adentro**:

- Handler depende de Service
- Service depende de Repository
- Repository depende de GORM/Database
- **Ninguna capa inferior conoce a una superior**

## Inyeccion de Dependencias

Se realiza manualmente en `cmd/server/main.go`:

```
DB (GORM)
  |
  +-- UserRepository
  |     +-- AuthService
  |           +-- AuthHandler
  |
  +-- PathRepository
  |     +-- PathService
  |     |     +-- PathHandler
  |     +-- SearchService
  |     |     +-- SearchHandler
  |     +-- ProgressService (+ ProgressRepository)
  |           +-- ProgressHandler
  |
  +-- ProgressRepository
  |     +-- DashboardService (+ AssessmentRepository)
  |           +-- DashboardHandler
  |
  +-- AssessmentRepository
        +-- AssessmentService
              +-- AssessmentHandler
```

## Paquetes

### internal/

Codigo privado de la aplicacion (Go no permite importar `internal/` desde fuera):

| Paquete | Responsabilidad |
|---------|-----------------|
| `config` | Carga configuracion desde variables de entorno |
| `dto` | Data Transfer Objects para requests/responses HTTP |
| `handler` | Handlers HTTP (capa de presentacion) |
| `middleware` | JWT auth, logging, panic recovery |
| `model` | Entidades GORM (tablas de la BD) |
| `repository` | Acceso a datos y seed inicial |
| `router` | Configuracion de rutas Gin |
| `service` | Logica de negocio |

### pkg/

Paquetes reutilizables que no dependen de la aplicacion:

| Paquete | Funcion |
|---------|---------|
| `hash` | bcrypt: HashPassword, CheckPassword |
| `jwt` | GenerateToken, ValidateToken con HS256 |

### cmd/

Punto de entrada de la aplicacion:

| Archivo | Funcion |
|---------|---------|
| `cmd/server/main.go` | Bootstrap: config, DB, repos, services, handlers, router |

## Middleware

### AuthMiddleware
- Extrae token del header `Authorization: Bearer <token>`
- Valida el JWT con el secreto configurado
- Inyecta `userID` y `email` en el contexto de Gin
- Retorna 401 si el token es invalido o falta

### Logging
- Registra metodo HTTP, ruta, codigo de respuesta y duracion

### Recovery
- Captura panics y responde con 500 en lugar de crashear el servidor

## DTOs vs Models

Los **Models** representan tablas de la base de datos:
```
model.User -> tabla users (incluye password_hash)
```

Los **DTOs** controlan lo que entra y sale por HTTP:
```
dto.UserResponse -> respuesta sin password_hash
dto.RegisterRequest -> request con validacion
```

Esta separacion evita exponer datos sensibles y permite validacion declarativa independiente del esquema de BD.

## Testing

Los tests se organizan junto al codigo que testean:

```
internal/handler/path_handler.go
internal/handler/path_handler_test.go

internal/service/progress_service.go
internal/service/progress_service_test.go
```

Patron de testing:
- **Handlers:** se testean los codigos de respuesta HTTP y el formato JSON
- **Services:** se testea la logica de negocio con datos controlados
