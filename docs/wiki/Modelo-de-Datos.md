# Modelo de Datos

Esquema completo de la base de datos de SkillPath. GORM maneja las migraciones automaticamente con `AutoMigrate`.

## Diagrama de Entidades

```
+----------------+     +-------------------+     +--------------+
|     User       |     |  LearningPath     |     |  Milestone   |
|----------------|     |-------------------|     |--------------|
| id        (PK) |     | id           (PK) |     | id      (PK) |
| email     (UQ) |     | title             |     | path_id (FK) |
| password_hash  |     | slug         (UQ) |     | title        |
| display_name   |     | description       |     | description  |
| avatar_url     |     | difficulty        |     | order_index  |
| bio            |     | estimated_hours   |     | est_hours    |
| github_url     |     | category          |     +------+-------+
| linkedin_url   |     | icon_url          |            |
| created_at     |     | is_published      |            | 1:N
| updated_at     |     | created_at        |            v
| deleted_at     |     | updated_at        |     +--------------+
+-------+--------+     +--------+----------+     |  Resource    |
        |                        |                |--------------|
        |                        |                | id      (PK) |
        |                        |                | milestone_id |
        | 1:N                    | 1:N            | title        |
        v                        v                | url          |
+-------------------+   +-------------------+    | type         |
|  UserProgress     |   |  SkillAssessment  |    | provider     |
|-------------------|   |-------------------|    | is_free      |
| id           (PK) |   | id           (PK) |    | est_minutes  |
| user_id      (FK) |   | milestone_id (FK) |    | order_index  |
| path_id      (FK) |   | question          |    +--------------+
| started_at        |   | options     (JSON)|
| completed_at      |   | correct_answer    |
| last_activity_at  |   | explanation       |
+--------+----------+   +-------------------+
         |
         | 1:N
         v
+---------------------+     +---------------------+
| MilestoneProgress   |     | ResourceProgress    |
|---------------------|     |---------------------|
| id             (PK) |     | id             (PK) |
| user_progress_id(FK)|     | user_progress_id(FK)|
| milestone_id   (FK) |     | resource_id    (FK) |
| status              |     | is_completed        |
| completed_at        |     | completed_at        |
+---------------------+     | rating              |
                             +---------------------+

+---------------------+
| AssessmentResult    |
|---------------------|
| id             (PK) |
| user_id        (FK) |
| milestone_id   (FK) |
| score               |
| total_questions     |
| passed_at           |
| created_at          |
+---------------------+
```

## Detalle de Entidades

### User

| Campo | Tipo | Constraints | Descripcion |
|-------|------|-------------|-------------|
| id | uint | PK, auto-increment | Identificador |
| email | string(255) | UNIQUE, NOT NULL | Email del usuario |
| password_hash | string(255) | NOT NULL | Hash bcrypt del password |
| display_name | string(100) | NOT NULL | Nombre visible |
| avatar_url | string(500) | - | URL del avatar |
| bio | string(500) | - | Biografia |
| github_url | string(255) | - | URL de GitHub |
| linkedin_url | string(255) | - | URL de LinkedIn |
| created_at | timestamp | auto | Fecha de creacion |
| updated_at | timestamp | auto | Ultima modificacion |
| deleted_at | timestamp | INDEX, soft-delete | Para borrado logico |

### LearningPath

| Campo | Tipo | Constraints | Descripcion |
|-------|------|-------------|-------------|
| id | uint | PK | Identificador |
| title | string(200) | NOT NULL | Nombre de la ruta |
| slug | string(200) | UNIQUE, NOT NULL | Slug para URLs |
| description | text | NOT NULL | Descripcion detallada |
| difficulty | string(20) | NOT NULL, default: beginner | beginner/intermediate/advanced |
| estimated_hours | int | NOT NULL, default: 0 | Horas estimadas |
| category | string(100) | NOT NULL | Backend, Frontend, Data, Mobile |
| icon_url | string(500) | - | URL del icono |
| is_published | bool | NOT NULL, default: false | Si es visible publicamente |

### Milestone

| Campo | Tipo | Constraints | Descripcion |
|-------|------|-------------|-------------|
| id | uint | PK | Identificador |
| path_id | uint | FK, NOT NULL, INDEX | Ruta a la que pertenece |
| title | string(200) | NOT NULL | Nombre del hito |
| description | text | - | Descripcion |
| order_index | int | NOT NULL, default: 0 | Orden dentro de la ruta |
| estimated_hours | float64 | NOT NULL, default: 0 | Horas estimadas |

### Resource

| Campo | Tipo | Constraints | Descripcion |
|-------|------|-------------|-------------|
| id | uint | PK | Identificador |
| milestone_id | uint | FK, NOT NULL, INDEX | Hito al que pertenece |
| title | string(300) | NOT NULL | Nombre del recurso |
| url | string(500) | NOT NULL | Enlace al recurso |
| type | string(20) | NOT NULL | video/article/course/book/tutorial/exercise |
| provider | string(100) | - | Plataforma (Udemy, YouTube, etc.) |
| is_free | bool | NOT NULL, default: true | Si es gratuito |
| estimated_minutes | int | NOT NULL, default: 0 | Duracion estimada |
| order_index | int | NOT NULL, default: 0 | Orden dentro del hito |

### UserProgress

| Campo | Tipo | Constraints | Descripcion |
|-------|------|-------------|-------------|
| id | uint | PK | Identificador |
| user_id | uint | FK, NOT NULL, UNIQUE(user_id, path_id) | Usuario |
| path_id | uint | FK, NOT NULL | Ruta |
| started_at | timestamp | NOT NULL | Inicio de inscripcion |
| completed_at | timestamp | nullable | Completada (NULL si en progreso) |
| last_activity_at | timestamp | NOT NULL | Ultima actividad |

### MilestoneProgress

| Campo | Tipo | Constraints | Descripcion |
|-------|------|-------------|-------------|
| id | uint | PK | Identificador |
| user_progress_id | uint | FK, UNIQUE(up_id, milestone_id) | Progreso del usuario |
| milestone_id | uint | FK | Hito |
| status | string(20) | NOT NULL, default: not_started | not_started/in_progress/completed |
| completed_at | timestamp | nullable | Cuando se completo |

### ResourceProgress

| Campo | Tipo | Constraints | Descripcion |
|-------|------|-------------|-------------|
| id | uint | PK | Identificador |
| user_progress_id | uint | FK, UNIQUE(up_id, resource_id) | Progreso del usuario |
| resource_id | uint | FK | Recurso |
| is_completed | bool | NOT NULL, default: false | Si esta completado |
| completed_at | timestamp | nullable | Cuando se completo |
| rating | int | default: 0 | Calificacion del usuario (1-5) |

## Tipos Enumerados

### Difficulty
- `beginner`
- `intermediate`
- `advanced`

### ResourceType
- `video`
- `article`
- `course`
- `book`
- `tutorial`
- `exercise`

### MilestoneStatus
- `not_started`
- `in_progress`
- `completed`

## JSONSlice

Tipo personalizado para almacenar arrays JSON en SQLite/PostgreSQL:

```go
type JSONSlice []string

func (j JSONSlice) Value() (driver.Value, error)     // Serializa a JSON
func (j *JSONSlice) Scan(value interface{}) error     // Deserializa desde JSON
```

Se usa en `SkillAssessment.Options` para almacenar las opciones de respuesta.

## Indices

| Tabla | Indice | Tipo |
|-------|--------|------|
| users | email | UNIQUE |
| users | deleted_at | INDEX |
| learning_paths | slug | UNIQUE |
| milestones | path_id | INDEX |
| resources | milestone_id | INDEX |
| user_progresses | user_id + path_id | UNIQUE |
| milestone_progresses | user_progress_id + milestone_id | UNIQUE |
| resource_progresses | user_progress_id + resource_id | UNIQUE |
| skill_assessments | milestone_id | INDEX |
| assessment_results | user_id | INDEX |
| assessment_results | milestone_id | INDEX |
