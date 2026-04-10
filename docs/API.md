# API Documentation

Base URL: `http://localhost:3005/api/v1`

Todas las respuestas usan formato JSON. Los endpoints protegidos requieren un header `Authorization: Bearer <token>`.

---

## Health Check

```bash
curl http://localhost:3005/health
```

Respuesta:
```json
{"status": "ok"}
```

---

## Autenticacion

### Registro

```bash
curl -X POST http://localhost:3005/api/v1/auth/register \
  -H "Content-Type: application/json" \
  -d '{
    "email": "dev@example.com",
    "password": "secret123",
    "display_name": "Dev User"
  }'
```

Respuesta `201 Created`:
```json
{
  "token": "eyJhbGciOiJIUzI1NiIs...",
  "user": {
    "id": 1,
    "email": "dev@example.com",
    "display_name": "Dev User"
  }
}
```

Errores:
- `400` - Validacion fallida (email invalido, password < 6 chars)
- `409` - Email ya registrado

### Login

```bash
curl -X POST http://localhost:3005/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "dev@example.com",
    "password": "secret123"
  }'
```

Respuesta `200 OK`:
```json
{
  "token": "eyJhbGciOiJIUzI1NiIs...",
  "user": {
    "id": 1,
    "email": "dev@example.com",
    "display_name": "Dev User"
  }
}
```

Errores:
- `401` - Credenciales invalidas

### Perfil del Usuario

```bash
curl http://localhost:3005/api/v1/auth/me \
  -H "Authorization: Bearer <token>"
```

Respuesta `200 OK`:
```json
{
  "id": 1,
  "email": "dev@example.com",
  "display_name": "Dev User",
  "avatar_url": "",
  "bio": "",
  "github_url": "",
  "linkedin_url": ""
}
```

### Actualizar Perfil

```bash
curl -X PUT http://localhost:3005/api/v1/auth/profile \
  -H "Authorization: Bearer <token>" \
  -H "Content-Type: application/json" \
  -d '{
    "display_name": "Ronald Dev",
    "bio": "Backend developer",
    "github_url": "https://github.com/ronalc90"
  }'
```

Respuesta `200 OK`: objeto `UserResponse` actualizado.

---

## Rutas de Aprendizaje

### Listar Rutas

Soporta paginacion y filtros por categoria, dificultad y busqueda.

```bash
# Todas las rutas
curl "http://localhost:3005/api/v1/paths"

# Filtrar por categoria
curl "http://localhost:3005/api/v1/paths?category=Backend"

# Filtrar por dificultad
curl "http://localhost:3005/api/v1/paths?difficulty=beginner"

# Busqueda por texto
curl "http://localhost:3005/api/v1/paths?search=java"

# Paginacion
curl "http://localhost:3005/api/v1/paths?page=1&page_size=10"
```

Respuesta `200 OK`:
```json
{
  "paths": [
    {
      "id": 1,
      "title": "Backend Developer con Java",
      "slug": "backend-developer-java",
      "description": "Domina el desarrollo backend con Java...",
      "difficulty": "intermediate",
      "estimated_hours": 120,
      "category": "Backend",
      "icon_url": "https://cdn.jsdelivr.net/...",
      "milestone_count": 8
    }
  ],
  "total": 4,
  "page": 1,
  "page_size": 12,
  "total_pages": 1
}
```

### Listar Categorias

```bash
curl http://localhost:3005/api/v1/paths/categories
```

Respuesta `200 OK`:
```json
[
  {"name": "Backend", "count": 1},
  {"name": "Frontend", "count": 1},
  {"name": "Data", "count": 1},
  {"name": "Mobile", "count": 1}
]
```

### Detalle de Ruta

```bash
curl http://localhost:3005/api/v1/paths/backend-developer-java
```

Respuesta `200 OK`:
```json
{
  "id": 1,
  "title": "Backend Developer con Java",
  "slug": "backend-developer-java",
  "description": "...",
  "difficulty": "intermediate",
  "estimated_hours": 120,
  "category": "Backend",
  "is_published": true,
  "milestones": [
    {
      "id": 1,
      "title": "Fundamentos de Java",
      "description": "Variables, tipos de datos, control de flujo, POO y colecciones.",
      "order_index": 1,
      "estimated_hours": 15,
      "resources": [
        {
          "id": 1,
          "title": "Java Programming Masterclass",
          "url": "https://www.udemy.com/course/java-the-complete-java-developer-course/",
          "type": "course",
          "provider": "Udemy",
          "is_free": false,
          "estimated_minutes": 480,
          "order_index": 1
        }
      ]
    }
  ]
}
```

---

## Progreso

### Inscribirse en una Ruta

```bash
curl -X POST http://localhost:3005/api/v1/paths/backend-developer-java/start \
  -H "Authorization: Bearer <token>"
```

Respuesta `201 Created`:
```json
{
  "path_id": 1,
  "path": {"id": 1, "title": "Backend Developer con Java", "...": "..."},
  "started_at": "2025-01-15T10:30:00Z",
  "last_activity_at": "2025-01-15T10:30:00Z",
  "progress_pct": 0
}
```

Errores:
- `404` - Ruta no encontrada
- `409` - Ya inscrito en esta ruta

### Ver Progreso en una Ruta

```bash
curl http://localhost:3005/api/v1/paths/backend-developer-java/progress \
  -H "Authorization: Bearer <token>"
```

Respuesta `200 OK`:
```json
{
  "path_id": 1,
  "path": {"...": "..."},
  "started_at": "2025-01-15T10:30:00Z",
  "last_activity_at": "2025-01-16T14:00:00Z",
  "progress_pct": 12.5,
  "milestones": [
    {
      "milestone_id": 1,
      "title": "Fundamentos de Java",
      "status": "completed",
      "order_index": 1
    }
  ]
}
```

### Completar un Hito

```bash
curl -X PUT http://localhost:3005/api/v1/paths/backend-developer-java/milestones/1/complete \
  -H "Authorization: Bearer <token>"
```

Respuesta `200 OK`:
```json
{"message": "Milestone completed successfully"}
```

### Completar un Recurso

```bash
curl -X PUT http://localhost:3005/api/v1/paths/backend-developer-java/resources/1/complete \
  -H "Authorization: Bearer <token>" \
  -H "Content-Type: application/json" \
  -d '{"rating": 5}'
```

Respuesta `200 OK`:
```json
{"message": "Resource completed successfully"}
```

### Mis Rutas

```bash
curl http://localhost:3005/api/v1/my-paths \
  -H "Authorization: Bearer <token>"
```

Respuesta `200 OK`: array de `UserPathProgress`.

---

## Evaluaciones (Assessment)

### Obtener Quiz de un Hito

```bash
curl http://localhost:3005/api/v1/milestones/1/assessment \
  -H "Authorization: Bearer <token>"
```

Respuesta `200 OK`:
```json
[
  {
    "id": 1,
    "question": "Cual es la diferencia entre == y .equals() en Java?",
    "options": [
      "== compara referencias de objetos, .equals() compara contenido",
      "Son identicos en funcionalidad",
      "== solo funciona con primitivos",
      ".equals() solo funciona con Strings"
    ]
  }
]
```

Las respuestas correctas NO se incluyen en esta respuesta.

### Enviar Respuestas del Quiz

```bash
curl -X POST http://localhost:3005/api/v1/milestones/1/assessment/submit \
  -H "Authorization: Bearer <token>" \
  -H "Content-Type: application/json" \
  -d '{
    "answers": [
      {"question_id": 1, "selected_index": 0},
      {"question_id": 2, "selected_index": 2},
      {"question_id": 3, "selected_index": 2}
    ]
  }'
```

Respuesta `200 OK`:
```json
{
  "score": 3,
  "total_questions": 3,
  "passed": true,
  "details": [
    {
      "question_id": 1,
      "is_correct": true,
      "correct_answer": 0,
      "explanation": "El operador == compara referencias..."
    }
  ]
}
```

---

## Dashboard

```bash
curl http://localhost:3005/api/v1/dashboard \
  -H "Authorization: Bearer <token>"
```

Respuesta `200 OK`:
```json
{
  "paths_in_progress": 2,
  "paths_completed": 1,
  "total_hours_learned": 45.5,
  "completion_rate": 33.33,
  "recent_activity": [
    {
      "type": "progress",
      "title": "Backend Developer con Java",
      "path_title": "Backend Developer con Java",
      "timestamp": "2025-01-16T14:00:00Z"
    }
  ],
  "skill_radar": [
    {"category": "Backend", "score": 75, "max_score": 100}
  ],
  "enrolled_paths": []
}
```

---

## Busqueda

```bash
curl "http://localhost:3005/api/v1/search?q=spring"
```

Respuesta `200 OK`:
```json
{
  "paths": [
    {"id": 1, "title": "Backend Developer con Java", "...": "..."}
  ],
  "resources": [
    {"id": 4, "title": "Spring Boot Reference Documentation", "...": "..."}
  ]
}
```

---

## Errores

Todas las respuestas de error siguen este formato:

```json
{
  "error": "codigo_error",
  "message": "Descripcion legible del error"
}
```

Codigos HTTP comunes:
- `400` - Error de validacion
- `401` - No autenticado o token invalido
- `404` - Recurso no encontrado
- `409` - Conflicto (duplicado)
- `500` - Error interno del servidor
