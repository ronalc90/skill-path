# API Reference

Referencia completa de la API REST de SkillPath.

Base URL: `http://localhost:3005/api/v1`

## Autenticacion

Los endpoints protegidos requieren un header:
```
Authorization: Bearer <jwt_token>
```

El token se obtiene al registrarse o hacer login. Expira en 72 horas (configurable).

## Endpoints

### Health

| Metodo | Ruta | Auth | Descripcion |
|--------|------|------|-------------|
| GET | `/health` | No | Verificar estado del servidor |

### Auth

| Metodo | Ruta | Auth | Descripcion |
|--------|------|------|-------------|
| POST | `/auth/register` | No | Crear cuenta |
| POST | `/auth/login` | No | Iniciar sesion |
| GET | `/auth/me` | Si | Ver perfil |
| PUT | `/auth/profile` | Si | Actualizar perfil |

### Rutas de Aprendizaje

| Metodo | Ruta | Auth | Descripcion |
|--------|------|------|-------------|
| GET | `/paths` | No | Listar rutas (paginado) |
| GET | `/paths/categories` | No | Listar categorias |
| GET | `/paths/:slug` | No | Detalle de ruta |

Query params para listar rutas:
- `category` - Filtrar por categoria
- `difficulty` - Filtrar por dificultad (beginner, intermediate, advanced)
- `search` - Buscar por titulo o descripcion
- `page` - Numero de pagina (default: 1)
- `page_size` - Elementos por pagina (default: 12, max: 50)

### Progreso

| Metodo | Ruta | Auth | Descripcion |
|--------|------|------|-------------|
| POST | `/paths/:slug/start` | Si | Inscribirse |
| GET | `/paths/:slug/progress` | Si | Ver progreso |
| PUT | `/paths/:slug/milestones/:id/complete` | Si | Completar hito |
| PUT | `/paths/:slug/resources/:id/complete` | Si | Completar recurso |
| GET | `/my-paths` | Si | Mis rutas inscritas |

### Evaluaciones

| Metodo | Ruta | Auth | Descripcion |
|--------|------|------|-------------|
| GET | `/milestones/:id/assessment` | Si | Obtener quiz |
| POST | `/milestones/:id/assessment/submit` | Si | Enviar respuestas |

### Dashboard

| Metodo | Ruta | Auth | Descripcion |
|--------|------|------|-------------|
| GET | `/dashboard` | Si | Metricas del usuario |

### Busqueda

| Metodo | Ruta | Auth | Descripcion |
|--------|------|------|-------------|
| GET | `/search?q=` | No | Busqueda global |

## Formato de Errores

```json
{
  "error": "codigo_de_error",
  "message": "Descripcion legible del error"
}
```

### Codigos de Error

| Codigo HTTP | Error | Significado |
|-------------|-------|-------------|
| 400 | `validation_error` | Datos de entrada invalidos |
| 400 | `invalid_id` | ID no es un numero valido |
| 401 | `unauthorized` | Token faltante o invalido |
| 401 | `invalid_credentials` | Email o password incorrectos |
| 404 | `not_found` | Recurso no encontrado |
| 404 | `not_enrolled` | No inscrito en la ruta |
| 409 | `email_exists` | Email ya registrado |
| 409 | `already_enrolled` | Ya inscrito en la ruta |
| 500 | `internal_error` | Error interno del servidor |

## Request/Response Bodies

### RegisterRequest
```json
{
  "email": "string (required, valid email)",
  "password": "string (required, min 6 chars)",
  "display_name": "string (required, 2-100 chars)"
}
```

### LoginRequest
```json
{
  "email": "string (required)",
  "password": "string (required)"
}
```

### UpdateProfileRequest
```json
{
  "display_name": "string (optional, 2-100 chars)",
  "avatar_url": "string (optional, valid URL)",
  "bio": "string (optional, max 500 chars)",
  "github_url": "string (optional, valid URL)",
  "linkedin_url": "string (optional, valid URL)"
}
```

### SubmitAssessmentRequest
```json
{
  "answers": [
    {
      "question_id": "uint (required)",
      "selected_index": "int (required, >= 0)"
    }
  ]
}
```

### CompleteResourceRequest
```json
{
  "rating": "int (optional, 1-5)"
}
```

Documentacion detallada con ejemplos curl en [docs/API.md](../API.md).
