# SkillPath - Alcance del MVP

**Autor:** Ronald
**Duracion estimada:** 12 sprints de 1 semana
**Equipo minimo:** 1 fullstack developer + 1 content curator

---

## Objetivo del MVP

Validar que los profesionales autodidactas adoptaran una plataforma de rutas de aprendizaje curadas con seguimiento de progreso y evaluaciones. El MVP debe permitir a un usuario:

1. Registrarse e iniciar sesion
2. Explorar y seleccionar una ruta de aprendizaje
3. Navegar el arbol de habilidades de la ruta
4. Marcar hitos como completados
5. Acceder a recursos curados por hito
6. Tomar evaluaciones basicas por hito
7. Ver su progreso general

---

## Plan de Sprints

### Sprint 1: Infraestructura Base
**Objetivo:** Tener el proyecto configurado y desplegado.

- Inicializar repositorio monorepo
- Configurar Next.js 14 con TypeScript y Tailwind CSS
- Configurar proyecto Go con Gin y estructura clean architecture
- Configurar PostgreSQL y Redis en Railway
- Configurar Vercel para frontend
- Implementar migraciones de base de datos con golang-migrate
- Configurar CI/CD basico (GitHub Actions: lint + test + deploy)
- Crear docker-compose para desarrollo local
- Documentar setup de desarrollo local

**Entregable:** Proyecto desplegado con health check funcional en ambos servicios.

### Sprint 2: Autenticacion y Usuarios
**Objetivo:** Sistema de autenticacion completo.

- Implementar Auth.js en frontend (GitHub OAuth + Google OAuth + Magic Link)
- Crear endpoints de usuario en backend Go:
  - POST /api/v1/auth/register
  - POST /api/v1/auth/login
  - GET /api/v1/auth/me
  - PUT /api/v1/users/profile
- Implementar middleware de autenticacion JWT en Go
- Crear modelo de usuario en PostgreSQL
- Implementar sesiones con Redis
- Disenar y construir paginas de login/registro en Next.js
- Crear layout base con navegacion
- Implementar proteccion de rutas (middleware Next.js)

**Entregable:** Usuarios pueden registrarse, iniciar sesion y ver su perfil.

### Sprint 3: Modelo de Datos de Rutas
**Objetivo:** Estructura de datos para rutas de aprendizaje.

- Disenar e implementar esquema de rutas, hitos y recursos
- Crear CRUD de rutas (admin):
  - POST /api/v1/admin/paths
  - PUT /api/v1/admin/paths/:id
  - POST /api/v1/admin/paths/:id/milestones
  - POST /api/v1/admin/milestones/:id/resources
- Crear endpoints publicos de lectura:
  - GET /api/v1/paths (listado con filtros)
  - GET /api/v1/paths/:slug (detalle con hitos)
  - GET /api/v1/milestones/:id/resources
- Implementar seeders con datos iniciales para las 4 rutas MVP
- Escribir tests de integracion para los endpoints

**Entregable:** API funcional para leer rutas con sus hitos y recursos.

### Sprint 4: Catalogo de Rutas (Frontend)
**Objetivo:** Paginas para explorar y ver rutas de aprendizaje.

- Pagina de catalogo de rutas (/paths):
  - Grid de tarjetas de rutas
  - Filtros por categoria, dificultad, duracion estimada
  - Busqueda por texto
- Pagina de detalle de ruta (/paths/[slug]):
  - Descripcion de la ruta
  - Lista de hitos en orden
  - Recursos por hito
  - Duracion estimada total
  - Boton "Iniciar esta ruta"
- Componente de tarjeta de ruta reutilizable
- Implementar skeleton loaders para carga
- Responsive design (mobile-first)

**Entregable:** Usuarios pueden explorar y ver todas las rutas disponibles.

### Sprint 5: Arbol de Habilidades (Skill Tree)
**Objetivo:** Visualizacion interactiva del arbol de habilidades.

- Disenar estructura de datos DAG (grafo dirigido aciclico) para dependencias entre hitos
- Implementar componente de arbol de habilidades con React Flow o D3.js:
  - Nodos representando hitos
  - Conexiones mostrando dependencias/prerrequisitos
  - Estados visuales: bloqueado, disponible, en progreso, completado
  - Zoom y pan
  - Tooltips con informacion del hito
- Integracion con datos de la API
- Vista mobile simplificada (lista lineal con indicadores)
- Animaciones de transicion al completar hitos

**Entregable:** Visualizacion funcional del arbol de habilidades en la pagina de ruta.

### Sprint 6: Seguimiento de Progreso
**Objetivo:** Los usuarios pueden rastrear su avance en las rutas.

- Endpoints de progreso:
  - POST /api/v1/paths/:id/start (iniciar ruta)
  - POST /api/v1/milestones/:id/complete (marcar completado)
  - GET /api/v1/users/me/progress (progreso general)
  - GET /api/v1/users/me/paths (rutas del usuario)
  - DELETE /api/v1/paths/:id/progress (abandonar ruta)
- Modelo de progreso en PostgreSQL (user_progress, milestone_completions)
- Logica de negocio:
  - Validar que prerrequisitos estan completados antes de marcar un hito
  - Calcular porcentaje de completacion de ruta
  - Registrar timestamps de inicio y completacion
- Actualizar arbol de habilidades para reflejar progreso real
- Pagina de dashboard del usuario (/dashboard):
  - Rutas activas con barra de progreso
  - Ultimo hito completado
  - Siguiente hito sugerido
  - Estadisticas generales

**Entregable:** Sistema de progreso funcional end-to-end.

### Sprint 7: Sistema de Evaluaciones
**Objetivo:** Quizzes basicos por hito para validar conocimiento.

- Disenar modelo de datos para evaluaciones:
  - Tabla skill_assessments (quiz metadata)
  - Tabla questions (preguntas con opciones)
  - Tabla user_assessment_results (resultados por usuario)
- Endpoints de evaluaciones:
  - GET /api/v1/milestones/:id/assessment (obtener quiz)
  - POST /api/v1/assessments/:id/submit (enviar respuestas)
  - GET /api/v1/users/me/assessments (historial de evaluaciones)
- Tipos de preguntas soportados en MVP:
  - Opcion multiple (una respuesta correcta)
  - Opcion multiple (multiples respuestas correctas)
  - Verdadero/Falso
- Pagina de quiz (/assessments/[id]):
  - Una pregunta a la vez con navegacion
  - Temporizador opcional
  - Pagina de resultados con respuestas correctas
- Logica: aprobar con 70%+ desbloquea el hito como "evaluado"
- Crear banco de 10-15 preguntas por hito para las 4 rutas MVP

**Entregable:** Usuarios pueden tomar quizzes y ver resultados.

### Sprint 8: Recursos y Calificaciones
**Objetivo:** Sistema de calificacion y bookmarks de recursos.

- Endpoints de recursos:
  - POST /api/v1/resources/:id/rate (calificar 1-5 estrellas)
  - POST /api/v1/resources/:id/bookmark (guardar recurso)
  - GET /api/v1/users/me/bookmarks (recursos guardados)
  - GET /api/v1/resources/:id/reviews (ver resenas)
  - POST /api/v1/resources/:id/reviews (escribir resena)
- Componente de tarjeta de recurso mejorado:
  - Tipo de recurso (video, articulo, curso, documentacion)
  - Calificacion promedio con estrellas
  - Etiqueta gratuito/de pago
  - Boton de bookmark
  - Link externo al recurso
- Pagina de recursos guardados (/bookmarks)
- Ordenar recursos por calificacion promedio en la vista de hito

**Entregable:** Usuarios pueden calificar, resenar y guardar recursos.

### Sprint 9: Perfil de Usuario
**Objetivo:** Perfil publico con progreso y habilidades.

- Pagina de perfil publico (/profile/[username]):
  - Avatar, nombre, bio, links sociales
  - Rutas completadas y en progreso
  - Grafico radar de habilidades (radar chart)
  - Estadisticas: hitos completados, evaluaciones aprobadas, dias activo
  - Proyectos de portafolio (links a GitHub/demos)
- Endpoints:
  - GET /api/v1/users/:username/profile (perfil publico)
  - PUT /api/v1/users/me/profile (actualizar perfil)
  - GET /api/v1/users/:username/skills (radar de habilidades)
- Pagina de configuracion de cuenta (/settings):
  - Editar perfil
  - Cambiar email
  - Preferencias de notificaciones
  - Eliminar cuenta
- Generar URL compartible del perfil

**Entregable:** Perfiles publicos funcionales con visualizacion de habilidades.

### Sprint 10: Landing Page y Onboarding
**Objetivo:** Pagina de aterrizaje para conversion y flujo de onboarding.

- Landing page (/):
  - Hero section con propuesta de valor clara
  - Como funciona (3 pasos)
  - Rutas destacadas
  - Testimonios (placeholder para beta)
  - CTA de registro
  - Footer con links
- Flujo de onboarding para nuevos usuarios:
  - Seleccionar objetivo profesional
  - Seleccionar nivel de experiencia
  - Recomendar ruta basada en respuestas
  - Tour rapido de la plataforma
- SEO basico: meta tags, Open Graph, sitemap.xml
- Pagina de "Acerca de" (/about)
- Pagina de preguntas frecuentes (/faq)

**Entregable:** Landing page lista para adquisicion de usuarios.

### Sprint 11: Busqueda y Optimizaciones
**Objetivo:** Busqueda funcional y mejoras de rendimiento.

- Integrar Algolia Search:
  - Indexar rutas de aprendizaje
  - Indexar recursos
  - Implementar componente de busqueda instantanea
  - Filtros facetados (categoria, dificultad, tipo de recurso)
- Optimizaciones de rendimiento:
  - Implementar cache en Redis para rutas populares
  - Lazy loading de imagenes
  - Optimizar queries SQL con indices
  - Implementar paginacion cursor-based
- Rate limiting en la API (Redis-based)
- Logging estructurado en el backend
- Error tracking con Sentry

**Entregable:** Busqueda instantanea y rendimiento optimizado.

### Sprint 12: Testing, Pulido y Beta
**Objetivo:** Preparar para lanzamiento beta.

- Tests end-to-end con Playwright (flujos criticos):
  - Registro -> Login -> Explorar ruta -> Iniciar ruta -> Completar hito -> Tomar quiz
- Tests unitarios para logica de negocio en Go (>80% coverage)
- Tests de componentes React con Testing Library
- Auditar accesibilidad (WCAG 2.1 AA)
- Revisar responsive design en dispositivos reales
- Preparar emails transaccionales (bienvenida, progreso semanal)
- Crear datos de demo para presentaciones
- Documentar API con Swagger/OpenAPI
- Preparar plan de beta cerrada (200 usuarios)
- Crear formulario de feedback en la aplicacion

**Entregable:** MVP listo para beta cerrada.

---

## Esquema de Base de Datos

### Tabla: users
```sql
CREATE TABLE users (
    id              UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    email           VARCHAR(255) NOT NULL UNIQUE,
    username        VARCHAR(50) NOT NULL UNIQUE,
    display_name    VARCHAR(100),
    avatar_url      TEXT,
    bio             TEXT,
    github_url      VARCHAR(255),
    linkedin_url    VARCHAR(255),
    website_url     VARCHAR(255),
    experience_level VARCHAR(20) CHECK (experience_level IN ('beginner', 'intermediate', 'advanced')),
    career_goal     VARCHAR(100),
    plan            VARCHAR(10) DEFAULT 'free' CHECK (plan IN ('free', 'pro')),
    plan_expires_at TIMESTAMPTZ,
    is_mentor       BOOLEAN DEFAULT FALSE,
    is_admin        BOOLEAN DEFAULT FALSE,
    auth_provider   VARCHAR(20) NOT NULL, -- 'github', 'google', 'email'
    auth_provider_id VARCHAR(255),
    created_at      TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    last_login_at   TIMESTAMPTZ
);

CREATE INDEX idx_users_email ON users(email);
CREATE INDEX idx_users_username ON users(username);
CREATE INDEX idx_users_plan ON users(plan);
```

### Tabla: paths
```sql
CREATE TABLE paths (
    id              UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    slug            VARCHAR(100) NOT NULL UNIQUE,
    title           VARCHAR(200) NOT NULL,
    description     TEXT NOT NULL,
    long_description TEXT,
    category        VARCHAR(50) NOT NULL, -- 'backend', 'frontend', 'data', 'mobile', 'devops'
    difficulty      VARCHAR(20) NOT NULL CHECK (difficulty IN ('beginner', 'intermediate', 'advanced')),
    estimated_hours INTEGER NOT NULL,
    icon_url        TEXT,
    cover_image_url TEXT,
    is_published    BOOLEAN DEFAULT FALSE,
    is_featured     BOOLEAN DEFAULT FALSE,
    tags            TEXT[], -- PostgreSQL array
    created_by      UUID REFERENCES users(id),
    created_at      TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_paths_slug ON paths(slug);
CREATE INDEX idx_paths_category ON paths(category);
CREATE INDEX idx_paths_difficulty ON paths(difficulty);
CREATE INDEX idx_paths_published ON paths(is_published);
CREATE INDEX idx_paths_tags ON paths USING GIN(tags);
```

### Tabla: milestones
```sql
CREATE TABLE milestones (
    id              UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    path_id         UUID NOT NULL REFERENCES paths(id) ON DELETE CASCADE,
    title           VARCHAR(200) NOT NULL,
    description     TEXT NOT NULL,
    learning_objectives TEXT[], -- que aprendera el usuario
    order_index     INTEGER NOT NULL,
    estimated_hours INTEGER NOT NULL DEFAULT 1,
    difficulty      VARCHAR(20) CHECK (difficulty IN ('beginner', 'intermediate', 'advanced')),
    has_assessment  BOOLEAN DEFAULT FALSE,
    icon            VARCHAR(50), -- nombre del icono
    created_at      TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    UNIQUE(path_id, order_index)
);

CREATE INDEX idx_milestones_path ON milestones(path_id);
CREATE INDEX idx_milestones_order ON milestones(path_id, order_index);
```

### Tabla: milestone_dependencies
```sql
CREATE TABLE milestone_dependencies (
    milestone_id    UUID NOT NULL REFERENCES milestones(id) ON DELETE CASCADE,
    depends_on_id   UUID NOT NULL REFERENCES milestones(id) ON DELETE CASCADE,
    PRIMARY KEY (milestone_id, depends_on_id),
    CHECK (milestone_id != depends_on_id)
);

CREATE INDEX idx_deps_milestone ON milestone_dependencies(milestone_id);
CREATE INDEX idx_deps_depends_on ON milestone_dependencies(depends_on_id);
```

### Tabla: resources
```sql
CREATE TABLE resources (
    id              UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    milestone_id    UUID NOT NULL REFERENCES milestones(id) ON DELETE CASCADE,
    title           VARCHAR(300) NOT NULL,
    description     TEXT,
    url             TEXT NOT NULL,
    resource_type   VARCHAR(20) NOT NULL CHECK (resource_type IN ('video', 'article', 'course', 'documentation', 'tutorial', 'book', 'tool', 'podcast')),
    provider        VARCHAR(50), -- 'youtube', 'udemy', 'coursera', 'medium', 'official_docs'
    is_free         BOOLEAN NOT NULL DEFAULT TRUE,
    price_usd       DECIMAL(8,2),
    language        VARCHAR(10) DEFAULT 'es', -- 'es', 'en'
    duration_minutes INTEGER,
    is_recommended  BOOLEAN DEFAULT FALSE, -- recurso principal vs alternativo
    order_index     INTEGER NOT NULL DEFAULT 0,
    avg_rating      DECIMAL(3,2) DEFAULT 0.00,
    rating_count    INTEGER DEFAULT 0,
    created_by      UUID REFERENCES users(id),
    created_at      TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_resources_milestone ON resources(milestone_id);
CREATE INDEX idx_resources_type ON resources(resource_type);
CREATE INDEX idx_resources_free ON resources(is_free);
CREATE INDEX idx_resources_rating ON resources(avg_rating DESC);
```

### Tabla: user_progress
```sql
CREATE TABLE user_progress (
    id              UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id         UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    path_id         UUID NOT NULL REFERENCES paths(id) ON DELETE CASCADE,
    status          VARCHAR(20) NOT NULL DEFAULT 'active' CHECK (status IN ('active', 'paused', 'completed', 'abandoned')),
    started_at      TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    completed_at    TIMESTAMPTZ,
    paused_at       TIMESTAMPTZ,
    completion_pct  DECIMAL(5,2) DEFAULT 0.00,
    UNIQUE(user_id, path_id)
);

CREATE INDEX idx_progress_user ON user_progress(user_id);
CREATE INDEX idx_progress_path ON user_progress(path_id);
CREATE INDEX idx_progress_status ON user_progress(status);
```

### Tabla: milestone_completions
```sql
CREATE TABLE milestone_completions (
    id              UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id         UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    milestone_id    UUID NOT NULL REFERENCES milestones(id) ON DELETE CASCADE,
    completed_at    TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    time_spent_mins INTEGER, -- tiempo reportado por el usuario
    notes           TEXT, -- notas personales del usuario
    UNIQUE(user_id, milestone_id)
);

CREATE INDEX idx_completions_user ON milestone_completions(user_id);
CREATE INDEX idx_completions_milestone ON milestone_completions(milestone_id);
```

### Tabla: skill_assessments
```sql
CREATE TABLE skill_assessments (
    id              UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    milestone_id    UUID NOT NULL REFERENCES milestones(id) ON DELETE CASCADE,
    title           VARCHAR(200) NOT NULL,
    description     TEXT,
    passing_score   INTEGER NOT NULL DEFAULT 70, -- porcentaje minimo para aprobar
    time_limit_mins INTEGER, -- NULL = sin limite
    max_attempts    INTEGER DEFAULT 3,
    created_at      TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_assessments_milestone ON skill_assessments(milestone_id);
```

### Tabla: questions
```sql
CREATE TABLE questions (
    id              UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    assessment_id   UUID NOT NULL REFERENCES skill_assessments(id) ON DELETE CASCADE,
    question_text   TEXT NOT NULL,
    question_type   VARCHAR(20) NOT NULL CHECK (question_type IN ('single_choice', 'multiple_choice', 'true_false')),
    options         JSONB NOT NULL, -- [{text: "...", is_correct: bool}]
    explanation     TEXT, -- explicacion de la respuesta correcta
    order_index     INTEGER NOT NULL DEFAULT 0,
    points          INTEGER NOT NULL DEFAULT 1,
    created_at      TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_questions_assessment ON questions(assessment_id);
```

### Tabla: user_assessment_results
```sql
CREATE TABLE user_assessment_results (
    id              UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id         UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    assessment_id   UUID NOT NULL REFERENCES skill_assessments(id) ON DELETE CASCADE,
    score           INTEGER NOT NULL, -- porcentaje obtenido
    passed          BOOLEAN NOT NULL,
    answers         JSONB NOT NULL, -- respuestas del usuario
    time_taken_secs INTEGER,
    attempt_number  INTEGER NOT NULL DEFAULT 1,
    taken_at        TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_results_user ON user_assessment_results(user_id);
CREATE INDEX idx_results_assessment ON user_assessment_results(assessment_id);
CREATE INDEX idx_results_user_assessment ON user_assessment_results(user_id, assessment_id);
```

### Tabla: resource_reviews
```sql
CREATE TABLE resource_reviews (
    id              UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id         UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    resource_id     UUID NOT NULL REFERENCES resources(id) ON DELETE CASCADE,
    rating          INTEGER NOT NULL CHECK (rating BETWEEN 1 AND 5),
    review_text     TEXT,
    created_at      TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    UNIQUE(user_id, resource_id)
);

CREATE INDEX idx_reviews_resource ON resource_reviews(resource_id);
CREATE INDEX idx_reviews_user ON resource_reviews(user_id);
```

### Tabla: bookmarks
```sql
CREATE TABLE bookmarks (
    id              UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id         UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    resource_id     UUID NOT NULL REFERENCES resources(id) ON DELETE CASCADE,
    created_at      TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    UNIQUE(user_id, resource_id)
);

CREATE INDEX idx_bookmarks_user ON bookmarks(user_id);
```

### Tabla: discussions
```sql
CREATE TABLE discussions (
    id              UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    milestone_id    UUID NOT NULL REFERENCES milestones(id) ON DELETE CASCADE,
    user_id         UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    title           VARCHAR(300) NOT NULL,
    body            TEXT NOT NULL,
    is_question     BOOLEAN DEFAULT FALSE,
    is_resolved     BOOLEAN DEFAULT FALSE,
    upvote_count    INTEGER DEFAULT 0,
    reply_count     INTEGER DEFAULT 0,
    created_at      TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_discussions_milestone ON discussions(milestone_id);
CREATE INDEX idx_discussions_user ON discussions(user_id);
CREATE INDEX idx_discussions_upvotes ON discussions(upvote_count DESC);
```

### Tabla: discussion_replies
```sql
CREATE TABLE discussion_replies (
    id              UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    discussion_id   UUID NOT NULL REFERENCES discussions(id) ON DELETE CASCADE,
    user_id         UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    body            TEXT NOT NULL,
    is_accepted     BOOLEAN DEFAULT FALSE,
    upvote_count    INTEGER DEFAULT 0,
    created_at      TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_replies_discussion ON discussion_replies(discussion_id);
CREATE INDEX idx_replies_user ON discussion_replies(user_id);
```

---

## Endpoints de la API

### Autenticacion
| Metodo | Endpoint | Descripcion |
|--------|----------|-------------|
| POST | /api/v1/auth/register | Registro con email |
| POST | /api/v1/auth/login | Login con email/password |
| POST | /api/v1/auth/oauth/callback | Callback OAuth (GitHub/Google) |
| POST | /api/v1/auth/magic-link | Enviar magic link |
| POST | /api/v1/auth/refresh | Refrescar token JWT |
| POST | /api/v1/auth/logout | Cerrar sesion |
| GET | /api/v1/auth/me | Obtener usuario actual |

### Usuarios
| Metodo | Endpoint | Descripcion |
|--------|----------|-------------|
| GET | /api/v1/users/:username/profile | Perfil publico |
| PUT | /api/v1/users/me/profile | Actualizar perfil |
| GET | /api/v1/users/:username/skills | Radar de habilidades |
| GET | /api/v1/users/me/dashboard | Dashboard del usuario |
| DELETE | /api/v1/users/me | Eliminar cuenta |

### Rutas de Aprendizaje
| Metodo | Endpoint | Descripcion |
|--------|----------|-------------|
| GET | /api/v1/paths | Listar rutas (con filtros y paginacion) |
| GET | /api/v1/paths/:slug | Detalle de ruta con hitos |
| GET | /api/v1/paths/:slug/tree | Arbol de habilidades (DAG) |
| GET | /api/v1/paths/featured | Rutas destacadas |
| GET | /api/v1/paths/categories | Categorias disponibles |

### Progreso
| Metodo | Endpoint | Descripcion |
|--------|----------|-------------|
| POST | /api/v1/paths/:id/start | Iniciar ruta |
| GET | /api/v1/users/me/paths | Rutas del usuario con progreso |
| GET | /api/v1/users/me/paths/:id/progress | Progreso detallado de una ruta |
| POST | /api/v1/milestones/:id/complete | Marcar hito completado |
| PUT | /api/v1/paths/:id/pause | Pausar ruta |
| PUT | /api/v1/paths/:id/resume | Reanudar ruta |
| DELETE | /api/v1/paths/:id/progress | Abandonar ruta |

### Recursos
| Metodo | Endpoint | Descripcion |
|--------|----------|-------------|
| GET | /api/v1/milestones/:id/resources | Recursos de un hito |
| POST | /api/v1/resources/:id/rate | Calificar recurso |
| GET | /api/v1/resources/:id/reviews | Ver resenas |
| POST | /api/v1/resources/:id/reviews | Escribir resena |
| POST | /api/v1/resources/:id/bookmark | Guardar recurso |
| DELETE | /api/v1/resources/:id/bookmark | Quitar bookmark |
| GET | /api/v1/users/me/bookmarks | Recursos guardados |

### Evaluaciones
| Metodo | Endpoint | Descripcion |
|--------|----------|-------------|
| GET | /api/v1/milestones/:id/assessment | Obtener quiz de un hito |
| POST | /api/v1/assessments/:id/submit | Enviar respuestas |
| GET | /api/v1/users/me/assessments | Historial de evaluaciones |
| GET | /api/v1/assessments/:id/results | Resultados de una evaluacion |

### Discusiones
| Metodo | Endpoint | Descripcion |
|--------|----------|-------------|
| GET | /api/v1/milestones/:id/discussions | Discusiones de un hito |
| POST | /api/v1/milestones/:id/discussions | Crear discusion |
| GET | /api/v1/discussions/:id | Detalle de discusion con respuestas |
| POST | /api/v1/discussions/:id/replies | Responder a discusion |
| POST | /api/v1/discussions/:id/upvote | Votar positivo |
| PUT | /api/v1/discussions/:id/resolve | Marcar como resuelta |

### Busqueda
| Metodo | Endpoint | Descripcion |
|--------|----------|-------------|
| GET | /api/v1/search | Busqueda global (rutas + recursos) |
| GET | /api/v1/search/paths | Buscar solo rutas |
| GET | /api/v1/search/resources | Buscar solo recursos |

### Admin
| Metodo | Endpoint | Descripcion |
|--------|----------|-------------|
| POST | /api/v1/admin/paths | Crear ruta |
| PUT | /api/v1/admin/paths/:id | Actualizar ruta |
| DELETE | /api/v1/admin/paths/:id | Eliminar ruta |
| POST | /api/v1/admin/paths/:id/milestones | Agregar hito |
| PUT | /api/v1/admin/milestones/:id | Actualizar hito |
| POST | /api/v1/admin/milestones/:id/resources | Agregar recurso |
| POST | /api/v1/admin/assessments | Crear evaluacion |
| POST | /api/v1/admin/assessments/:id/questions | Agregar pregunta |
| GET | /api/v1/admin/stats | Estadisticas generales |

---

## Paginas y Pantallas

### Paginas Publicas
| Pagina | Ruta | Descripcion |
|--------|------|-------------|
| Landing | / | Pagina de aterrizaje con propuesta de valor, rutas destacadas y CTA |
| Catalogo | /paths | Grid de rutas con filtros y busqueda |
| Detalle de Ruta | /paths/[slug] | Descripcion completa, arbol de habilidades, recursos, CTA iniciar |
| Perfil Publico | /profile/[username] | Habilidades, rutas completadas, estadisticas |
| Acerca de | /about | Historia, mision, equipo |
| FAQ | /faq | Preguntas frecuentes |
| Login | /login | Inicio de sesion (OAuth + magic link) |
| Registro | /register | Registro de cuenta nueva |

### Paginas de Usuario (autenticado)
| Pagina | Ruta | Descripcion |
|--------|------|-------------|
| Dashboard | /dashboard | Rutas activas, progreso, siguiente hito sugerido |
| Mi Ruta | /my-paths/[slug] | Vista de ruta con progreso personal, arbol interactivo |
| Hito | /milestones/[id] | Detalle del hito, recursos, discusiones, evaluacion |
| Quiz | /assessments/[id] | Interfaz de evaluacion con preguntas paso a paso |
| Resultados | /assessments/[id]/results | Resultados con explicaciones |
| Bookmarks | /bookmarks | Recursos guardados |
| Configuracion | /settings | Perfil, cuenta, notificaciones |
| Onboarding | /onboarding | Flujo guiado para nuevos usuarios |

### Paginas Admin
| Pagina | Ruta | Descripcion |
|--------|------|-------------|
| Panel Admin | /admin | Estadisticas generales |
| Gestionar Rutas | /admin/paths | CRUD de rutas |
| Editor de Ruta | /admin/paths/[id]/edit | Editar ruta, hitos y recursos |
| Editor de Quiz | /admin/assessments/[id]/edit | Crear y editar evaluaciones |

---

## Rutas Iniciales del MVP

Las 4 rutas de aprendizaje que se incluiran en el MVP:

### 1. Desarrollador Backend (Java)
- **Categoria:** Backend
- **Dificultad:** Intermedio
- **Duracion estimada:** 200 horas
- **Hitos principales:**
  1. Fundamentos de Java (sintaxis, OOP, colecciones)
  2. Java Avanzado (generics, streams, concurrencia)
  3. Spring Boot Fundamentos (IoC, DI, configuracion)
  4. APIs REST con Spring Boot (controllers, DTOs, validacion)
  5. Persistencia de Datos (JPA, Hibernate, PostgreSQL)
  6. Seguridad (Spring Security, JWT, OAuth2)
  7. Testing (JUnit, Mockito, tests de integracion)
  8. Arquitectura (Clean Architecture, patrones de diseno)
  9. Docker y Despliegue (contenedores, CI/CD)
  10. Proyecto Final: API REST completa con autenticacion

### 2. Desarrollador Frontend (React)
- **Categoria:** Frontend
- **Dificultad:** Intermedio
- **Duracion estimada:** 180 horas
- **Hitos principales:**
  1. HTML y CSS Modernos (semantica, Flexbox, Grid, responsive)
  2. JavaScript Moderno (ES6+, async/await, modulos)
  3. TypeScript Fundamentos (tipos, interfaces, generics)
  4. React Fundamentos (componentes, props, estado, hooks)
  5. React Avanzado (Context, reducers, custom hooks, rendimiento)
  6. Estado Global (Zustand o Redux Toolkit)
  7. Next.js (routing, SSR, SSG, API routes)
  8. Testing Frontend (Vitest, Testing Library, Playwright)
  9. Estilos Avanzados (Tailwind CSS, CSS-in-JS, animaciones)
  10. Proyecto Final: Aplicacion web completa con Next.js

### 3. Analista de Datos (Python)
- **Categoria:** Data
- **Dificultad:** Principiante
- **Duracion estimada:** 160 horas
- **Hitos principales:**
  1. Python Fundamentos (sintaxis, estructuras de datos, funciones)
  2. Python para Datos (NumPy, manejo de arrays)
  3. Pandas (DataFrames, limpieza, transformacion)
  4. Visualizacion (Matplotlib, Seaborn, graficos efectivos)
  5. Estadistica Descriptiva (medidas centrales, dispersion, distribuciones)
  6. SQL para Analistas (queries, joins, agregaciones, subqueries)
  7. Analisis Exploratorio (EDA, patrones, outliers, correlaciones)
  8. Dashboards (Plotly, Dash o Streamlit)
  9. Introduccion a Machine Learning (scikit-learn, regresion, clasificacion)
  10. Proyecto Final: Analisis completo de dataset real con dashboard

### 4. Desarrollador Movil (Flutter)
- **Categoria:** Mobile
- **Dificultad:** Intermedio
- **Duracion estimada:** 170 horas
- **Hitos principales:**
  1. Dart Fundamentos (sintaxis, OOP, async, null safety)
  2. Flutter Fundamentos (widgets, layouts, navegacion)
  3. Estado en Flutter (setState, Provider, Riverpod)
  4. UI Avanzada (animaciones, temas, responsive, widgets custom)
  5. Networking (HTTP, REST APIs, JSON, Dio)
  6. Persistencia Local (SharedPreferences, SQLite, Hive)
  7. Firebase (Auth, Firestore, Storage, Push Notifications)
  8. Testing Flutter (unit, widget, integration tests)
  9. Publicacion (Play Store, App Store, CI/CD)
  10. Proyecto Final: App completa con backend, auth y publicacion

---

## Criterios de Exito del MVP

### Metricas Cuantitativas (a 4 semanas de la beta)
- 200 usuarios registrados en beta cerrada
- 60% de usuarios inician al menos 1 ruta (tasa de activacion)
- 40% completan al menos 3 hitos (engagement temprano)
- 25% toman al menos 1 evaluacion
- NPS > 30 en encuesta de beta

### Metricas Cualitativas
- Los usuarios entienden como navegar el arbol de habilidades sin tutorial
- Los recursos curados son percibidos como utiles y de calidad
- Las evaluaciones son percibidas como relevantes (no triviales ni imposibles)
- Los usuarios expresan intencion de volver y continuar

### Criterio de Decision
- **Proceder a Fase 2** si: activacion > 50% Y engagement > 30% Y NPS > 20
- **Pivotar** si: activacion < 30% O engagement < 15%
- **Iterar MVP** si: metricas entre los umbrales anteriores
