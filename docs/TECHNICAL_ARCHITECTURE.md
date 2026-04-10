# SkillPath - Arquitectura Tecnica

**Autor:** Ronald

---

## 1. Arquitectura General del Backend (Go)

### Clean Architecture

El backend sigue los principios de Clean Architecture, separando responsabilidades en capas con dependencias que siempre apuntan hacia el centro (dominio).

```
+-------------------------------------------------------+
|                     cmd/api/                           |
|                  (Entry Point)                         |
|                                                       |
|  +--------------------------------------------------+ |
|  |              internal/handler/                    | |
|  |           (HTTP Handlers / Controllers)           | |
|  |  - Recibe HTTP requests                          | |
|  |  - Valida input (binding, formato)               | |
|  |  - Llama al servicio correspondiente             | |
|  |  - Serializa la respuesta                        | |
|  |                                                  | |
|  |  +---------------------------------------------+ | |
|  |  |           internal/service/                  | | |
|  |  |        (Logica de Negocio)                   | | |
|  |  |  - Reglas de negocio                        | | |
|  |  |  - Orquestacion de operaciones              | | |
|  |  |  - Validaciones de dominio                  | | |
|  |  |  - Transacciones                            | | |
|  |  |                                             | | |
|  |  |  +----------------------------------------+ | | |
|  |  |  |      internal/repository/               | | | |
|  |  |  |     (Acceso a Datos)                    | | | |
|  |  |  |  - Queries SQL                         | | | |
|  |  |  |  - Cache Redis                         | | | |
|  |  |  |  - Algolia indexing                    | | | |
|  |  |  +----------------------------------------+ | | |
|  |  +---------------------------------------------+ | |
|  +--------------------------------------------------+ |
|                                                       |
|  internal/model/      (Modelos de Dominio)            |
|  internal/middleware/  (Auth, Logging, CORS, Rate)    |
|  internal/config/     (Configuracion)                 |
|  pkg/                 (Paquetes reutilizables)        |
+-------------------------------------------------------+
```

### Estructura de Directorios del Backend

```
backend/
  cmd/
    api/
      main.go                 # Entry point, inicializa servidor
  internal/
    config/
      config.go               # Carga de env vars, defaults
    handler/
      auth_handler.go         # Endpoints de autenticacion
      path_handler.go         # Endpoints de rutas
      milestone_handler.go    # Endpoints de hitos
      resource_handler.go     # Endpoints de recursos
      progress_handler.go     # Endpoints de progreso
      assessment_handler.go   # Endpoints de evaluaciones
      discussion_handler.go   # Endpoints de discusiones
      user_handler.go         # Endpoints de usuarios
      admin_handler.go        # Endpoints administrativos
      search_handler.go       # Endpoints de busqueda
    service/
      auth_service.go
      path_service.go
      milestone_service.go
      resource_service.go
      progress_service.go
      assessment_service.go
      discussion_service.go
      user_service.go
      search_service.go
    repository/
      user_repository.go
      path_repository.go
      milestone_repository.go
      resource_repository.go
      progress_repository.go
      assessment_repository.go
      discussion_repository.go
      interfaces.go           # Interfaces de repositorios
    model/
      user.go
      path.go
      milestone.go
      resource.go
      progress.go
      assessment.go
      discussion.go
      errors.go               # Errores de dominio
    middleware/
      auth.go                 # JWT validation
      cors.go                 # CORS configuration
      logger.go               # Request logging
      ratelimit.go            # Rate limiting
      recovery.go             # Panic recovery
    dto/
      request/                # DTOs de entrada
        auth_request.go
        path_request.go
        progress_request.go
        assessment_request.go
      response/               # DTOs de salida
        auth_response.go
        path_response.go
        progress_response.go
        error_response.go
  migrations/
    001_create_users.up.sql
    001_create_users.down.sql
    002_create_paths.up.sql
    ...
  pkg/
    jwt/
      jwt.go                  # Generacion y validacion de tokens
    hash/
      hash.go                 # Hashing de passwords
    validator/
      validator.go            # Validaciones customizadas
    pagination/
      cursor.go               # Paginacion cursor-based
  go.mod
  go.sum
  Makefile
  Dockerfile
```

### Patron de Inyeccion de Dependencias

```go
// internal/service/path_service.go
type PathService struct {
    pathRepo      repository.PathRepository
    milestoneRepo repository.MilestoneRepository
    cache         repository.CacheRepository
    clock         func() time.Time
}

func NewPathService(
    pathRepo repository.PathRepository,
    milestoneRepo repository.MilestoneRepository,
    cache repository.CacheRepository,
    clock func() time.Time,
) *PathService {
    return &PathService{
        pathRepo:      pathRepo,
        milestoneRepo: milestoneRepo,
        cache:         cache,
        clock:         clock,
    }
}
```

```go
// internal/repository/interfaces.go
type PathRepository interface {
    FindBySlug(ctx context.Context, slug string) (*model.Path, error)
    FindAll(ctx context.Context, filters PathFilters) ([]model.Path, string, error)
    Create(ctx context.Context, path *model.Path) error
    Update(ctx context.Context, path *model.Path) error
    Delete(ctx context.Context, id uuid.UUID) error
}

type CacheRepository interface {
    Get(ctx context.Context, key string) ([]byte, error)
    Set(ctx context.Context, key string, value []byte, ttl time.Duration) error
    Delete(ctx context.Context, key string) error
}
```

### Manejo de Errores Centralizado

```go
// internal/model/errors.go
type AppError struct {
    Code    int    `json:"-"`
    Type    string `json:"type"`
    Message string `json:"message"`
    Details any    `json:"details,omitempty"`
}

func (e *AppError) Error() string {
    return e.Message
}

var (
    ErrNotFound       = &AppError{Code: 404, Type: "not_found", Message: "Recurso no encontrado"}
    ErrUnauthorized   = &AppError{Code: 401, Type: "unauthorized", Message: "No autorizado"}
    ErrForbidden      = &AppError{Code: 403, Type: "forbidden", Message: "Acceso denegado"}
    ErrBadRequest     = &AppError{Code: 400, Type: "bad_request", Message: "Solicitud invalida"}
    ErrConflict       = &AppError{Code: 409, Type: "conflict", Message: "Conflicto con estado actual"}
    ErrInternal       = &AppError{Code: 500, Type: "internal", Message: "Error interno"}
)
```

```go
// internal/middleware/recovery.go - Error handler centralizado
func ErrorHandler() gin.HandlerFunc {
    return func(c *gin.Context) {
        c.Next()
        if len(c.Errors) > 0 {
            err := c.Errors.Last().Err
            if appErr, ok := err.(*model.AppError); ok {
                c.JSON(appErr.Code, appErr)
                return
            }
            // Error no controlado -> log + 500
            log.Error("unhandled error", "error", err)
            c.JSON(500, ErrInternal)
        }
    }
}
```

---

## 2. Arbol de Habilidades (Skill Tree) - Estructura DAG

### Modelo de Datos

El arbol de habilidades se modela como un **DAG (Directed Acyclic Graph)** -- grafo dirigido aciclico. Cada hito es un nodo, y las aristas representan dependencias (prerrequisitos).

```
Ejemplo: Ruta "Backend Developer (Java)"

    [Java Fundamentos]
           |
    [Java Avanzado]
        /       \
[Spring Boot]  [Testing JUnit]
       |            |
[APIs REST]    [Testing Integracion]
       |            |
[Persistencia JPA]  |
       \           /
    [Seguridad Spring]
           |
    [Docker y Deploy]
           |
    [Proyecto Final]
```

### Estructura de Datos para el DAG

```go
// internal/model/skill_tree.go

type SkillTreeNode struct {
    ID              uuid.UUID        `json:"id"`
    Title           string           `json:"title"`
    Description     string           `json:"description"`
    OrderIndex      int              `json:"order_index"`
    EstimatedHours  int              `json:"estimated_hours"`
    Difficulty      string           `json:"difficulty"`
    HasAssessment   bool             `json:"has_assessment"`
    Status          NodeStatus       `json:"status"`  // locked, available, in_progress, completed
    Dependencies    []uuid.UUID      `json:"dependencies"`
    Position        Position         `json:"position"` // coordenadas para renderizado
}

type NodeStatus string

const (
    NodeLocked      NodeStatus = "locked"
    NodeAvailable   NodeStatus = "available"
    NodeInProgress  NodeStatus = "in_progress"
    NodeCompleted   NodeStatus = "completed"
)

type Position struct {
    X int `json:"x"`
    Y int `json:"y"`
}

type SkillTree struct {
    PathID    uuid.UUID         `json:"path_id"`
    PathTitle string            `json:"path_title"`
    Nodes     []SkillTreeNode   `json:"nodes"`
    Edges     []SkillTreeEdge   `json:"edges"`
}

type SkillTreeEdge struct {
    From uuid.UUID `json:"from"`
    To   uuid.UUID `json:"to"`
}
```

### Algoritmo de Renderizado del DAG

El arbol de habilidades se renderiza usando un algoritmo de layout en capas (Sugiyama):

```
1. Ordenamiento Topologico (Topological Sort)
   - Determinar el orden de procesamiento de nodos respetando dependencias
   - Detectar ciclos (error si se encuentran)

2. Asignacion de Capas (Layer Assignment)
   - Cada nodo se asigna a una capa (fila) basada en la distancia
     maxima desde cualquier nodo raiz
   - Nodos sin dependencias van en la capa 0
   - Un nodo va en la capa max(capa de cada dependencia) + 1

3. Reduccion de Cruces (Crossing Minimization)
   - Dentro de cada capa, ordenar nodos para minimizar
     cruces de aristas con la capa anterior
   - Heuristica: mediana o baricentro de las posiciones
     de los nodos conectados

4. Asignacion de Coordenadas (Coordinate Assignment)
   - X: posicion horizontal dentro de la capa (equidistante)
   - Y: numero de capa * espaciado vertical
   - Centrar nodos con pocos hermanos
```

```go
// internal/service/skill_tree_service.go

func (s *SkillTreeService) BuildSkillTree(
    ctx context.Context,
    pathID uuid.UUID,
    userID *uuid.UUID,
) (*model.SkillTree, error) {
    milestones, err := s.milestoneRepo.FindByPath(ctx, pathID)
    if err != nil {
        return nil, err
    }

    deps, err := s.milestoneRepo.FindDependencies(ctx, pathID)
    if err != nil {
        return nil, err
    }

    // 1. Construir grafo de adyacencia
    graph := buildAdjacencyList(milestones, deps)

    // 2. Verificar que no hay ciclos
    if hasCycle(graph) {
        return nil, model.ErrInvalidDAG
    }

    // 3. Ordenamiento topologico
    sorted := topologicalSort(graph)

    // 4. Asignar capas
    layers := assignLayers(sorted, graph)

    // 5. Calcular posiciones
    positions := calculatePositions(layers)

    // 6. Si hay usuario, calcular estados
    var completions map[uuid.UUID]bool
    if userID != nil {
        completions, err = s.progressRepo.GetCompletions(ctx, *userID, pathID)
        if err != nil {
            return nil, err
        }
    }

    // 7. Construir arbol con estados
    tree := assembleSkillTree(milestones, deps, positions, completions)
    return tree, nil
}

// Logica de estado de nodos
func determineNodeStatus(
    nodeID uuid.UUID,
    deps []uuid.UUID,
    completions map[uuid.UUID]bool,
) model.NodeStatus {
    if completions[nodeID] {
        return model.NodeCompleted
    }

    // Un nodo esta disponible si todas sus dependencias estan completadas
    allDepsCompleted := true
    for _, depID := range deps {
        if !completions[depID] {
            allDepsCompleted = false
            break
        }
    }

    if allDepsCompleted {
        return model.NodeAvailable
    }

    return model.NodeLocked
}
```

### Renderizado en Frontend (React Flow)

```
Componente SkillTree (React):
  - Recibe datos del DAG desde la API
  - Usa React Flow para renderizar nodos y aristas
  - Cada nodo es un componente custom con estados visuales:
    - Locked:      gris, icono de candado, no clickeable
    - Available:   azul, borde brillante, clickeable
    - InProgress:  amarillo, barra de progreso animada
    - Completed:   verde, icono de check, efecto de brillo
  - Las aristas cambian de color segun el estado
  - Click en nodo disponible -> navega a detalle del hito
  - Tooltip al hover muestra titulo, descripcion, duracion estimada
  - Controles de zoom y pan
  - Vista mobile: lista lineal con indicadores de estado
```

---

## 3. Integracion con Algolia (Busqueda)

### Indices

```
Indice: skillpath_paths
{
  objectID:        "uuid-de-la-ruta",
  title:           "Desarrollador Backend (Java)",
  description:     "Ruta completa para convertirte en...",
  category:        "backend",
  difficulty:      "intermediate",
  estimated_hours: 200,
  tags:            ["java", "spring", "backend", "api"],
  milestone_count: 10,
  is_featured:     true
}

Indice: skillpath_resources
{
  objectID:        "uuid-del-recurso",
  title:           "Spring Boot Tutorial Completo",
  description:     "Tutorial paso a paso de Spring Boot...",
  resource_type:   "video",
  provider:        "youtube",
  is_free:         true,
  avg_rating:      4.5,
  path_title:      "Desarrollador Backend (Java)",
  milestone_title: "Spring Boot Fundamentos"
}

Indice: skillpath_mentors (Fase 3)
{
  objectID:        "uuid-del-mentor",
  display_name:    "Maria Garcia",
  specialties:     ["java", "spring", "microservicios"],
  experience_years: 8,
  avg_rating:      4.8,
  sessions_count:  45
}
```

### Configuracion de Busqueda

```
Atributos buscables (searchableAttributes):
  paths:     title, description, tags, category
  resources: title, description, provider, path_title

Facetas para filtrado (attributesForFaceting):
  paths:     category, difficulty, tags
  resources: resource_type, provider, is_free, milestone_title

Ranking personalizado (customRanking):
  paths:     desc(is_featured), desc(milestone_count)
  resources: desc(avg_rating), desc(is_recommended)
```

### Pipeline de Sincronizacion

```
1. Escritura en PostgreSQL (fuente de verdad)
       |
2. Evento de cambio capturado en el servicio
       |
3. Serializacion a formato Algolia
       |
4. Envio asincrono a Algolia via goroutine
       |
5. Si falla -> cola de reintentos (max 3 intentos con backoff)
       |
6. Log de sincronizacion para auditing
```

```go
// internal/service/search_service.go

type SearchService struct {
    algoliaClient *search.APIClient
    pathIndex     string
    resourceIndex string
}

func (s *SearchService) IndexPath(ctx context.Context, path *model.Path) error {
    record := map[string]any{
        "objectID":        path.ID.String(),
        "title":           path.Title,
        "description":     path.Description,
        "category":        path.Category,
        "difficulty":      path.Difficulty,
        "estimated_hours": path.EstimatedHours,
        "tags":            path.Tags,
        "is_featured":     path.IsFeatured,
    }

    _, err := s.algoliaClient.SaveObject(
        s.algoliaClient.NewApiSaveObjectRequest(s.pathIndex, record),
    )
    return err
}
```

---

## 4. Algoritmo de Seguimiento de Progreso

### Calculo de Porcentaje de Completacion

```
Porcentaje = (Hitos Completados / Total Hitos en la Ruta) * 100

Nota: Todos los hitos tienen el mismo peso en el MVP.
      En versiones futuras se podria ponderar por estimated_hours.
```

### Flujo de Completacion de Hito

```
1. Usuario marca hito como completado
       |
2. Backend valida:
   a. El usuario tiene una ruta activa que contiene este hito
   b. Todas las dependencias del hito estan completadas
   c. El hito no estaba ya completado
       |
3. Si el hito tiene evaluacion:
   a. Verificar que el usuario aprobo la evaluacion (score >= passing_score)
   b. Si no aprobo -> rechazar completacion
       |
4. Crear registro en milestone_completions
       |
5. Recalcular completion_pct en user_progress
       |
6. Si completion_pct == 100:
   a. Marcar ruta como 'completed'
   b. Registrar completed_at
   c. Generar evento de completacion (para notificaciones, stats)
       |
7. Invalidar cache de progreso del usuario en Redis
       |
8. Actualizar nodos del arbol de habilidades:
   a. Hito actual -> 'completed'
   b. Hitos que dependian de este -> verificar si ahora son 'available'
       |
9. Retornar estado actualizado al frontend
```

```go
// internal/service/progress_service.go

func (s *ProgressService) CompleteMilestone(
    ctx context.Context,
    userID uuid.UUID,
    milestoneID uuid.UUID,
) (*model.ProgressUpdate, error) {
    // 1. Obtener progreso activo del usuario
    milestone, err := s.milestoneRepo.FindByID(ctx, milestoneID)
    if err != nil {
        return nil, err
    }

    progress, err := s.progressRepo.FindActiveByUserAndPath(ctx, userID, milestone.PathID)
    if err != nil {
        return nil, model.NewAppError(400, "no_active_path", "No tienes esta ruta activa")
    }

    // 2. Verificar que no esta ya completado
    if s.progressRepo.IsMilestoneCompleted(ctx, userID, milestoneID) {
        return nil, model.NewAppError(409, "already_completed", "Este hito ya esta completado")
    }

    // 3. Verificar dependencias
    deps, err := s.milestoneRepo.FindDependencies(ctx, milestone.PathID)
    if err != nil {
        return nil, err
    }

    milestoneDeps := filterDepsForMilestone(deps, milestoneID)
    for _, depID := range milestoneDeps {
        if !s.progressRepo.IsMilestoneCompleted(ctx, userID, depID) {
            return nil, model.NewAppError(400, "deps_not_met",
                "Debes completar los prerrequisitos primero")
        }
    }

    // 4. Verificar evaluacion si es requerida
    if milestone.HasAssessment {
        passed, err := s.assessmentRepo.HasPassedAssessment(ctx, userID, milestoneID)
        if err != nil {
            return nil, err
        }
        if !passed {
            return nil, model.NewAppError(400, "assessment_required",
                "Debes aprobar la evaluacion para completar este hito")
        }
    }

    // 5. Registrar completacion (en transaccion)
    now := s.clock()
    err = s.progressRepo.WithTransaction(ctx, func(tx repository.Transaction) error {
        if err := tx.CreateMilestoneCompletion(ctx, userID, milestoneID, now); err != nil {
            return err
        }

        totalMilestones, err := tx.CountMilestonesInPath(ctx, milestone.PathID)
        if err != nil {
            return err
        }
        completedMilestones, err := tx.CountCompletedMilestones(ctx, userID, milestone.PathID)
        if err != nil {
            return err
        }

        pct := float64(completedMilestones) / float64(totalMilestones) * 100
        progress.CompletionPct = pct

        if pct >= 100 {
            progress.Status = "completed"
            progress.CompletedAt = &now
        }

        return tx.UpdateProgress(ctx, progress)
    })

    if err != nil {
        return nil, err
    }

    // 6. Invalidar cache
    s.cache.Delete(ctx, fmt.Sprintf("progress:%s:%s", userID, milestone.PathID))

    return &model.ProgressUpdate{
        MilestoneID:   milestoneID,
        CompletionPct: progress.CompletionPct,
        PathCompleted: progress.Status == "completed",
        // Nodos que se desbloquean
        UnlockedNodes: s.findNewlyUnlockedNodes(ctx, userID, milestone.PathID, milestoneID),
    }, nil
}
```

---

## 5. Motor de Evaluaciones (Quiz System)

### Arquitectura del Sistema de Quizzes

```
+-------------------+     +------------------+     +------------------+
|  Frontend Quiz    |     |  Assessment      |     |  Question Bank   |
|  Component        | --> |  Handler         | --> |  (PostgreSQL)    |
|                   |     |                  |     |                  |
| - Renderiza       |     | - Valida input   |     | - questions      |
|   preguntas       |     | - Verifica       |     | - opciones       |
| - Controla timer  |     |   intentos       |     | - respuestas     |
| - Recoge          |     | - Calcula score  |     |   correctas      |
|   respuestas      |     | - Registra       |     | - explicaciones  |
| - Muestra         |     |   resultado      |     |                  |
|   resultados      |     +------------------+     +------------------+
+-------------------+
```

### Flujo de una Evaluacion

```
1. Usuario solicita evaluacion de un hito
   GET /api/v1/milestones/:id/assessment
       |
2. Backend:
   a. Verificar que el usuario tiene la ruta activa
   b. Verificar que no excedio max_attempts
   c. Seleccionar preguntas (todas en MVP, random subset en futuro)
   d. Retornar preguntas SIN respuestas correctas
       |
3. Frontend:
   a. Renderizar preguntas una a la vez
   b. Mostrar timer si aplica
   c. Permitir navegar entre preguntas
   d. Boton "Enviar evaluacion"
       |
4. Usuario envia respuestas
   POST /api/v1/assessments/:id/submit
   Body: { answers: [{ question_id, selected_options: [0, 2] }] }
       |
5. Backend:
   a. Validar que todas las preguntas fueron respondidas
   b. Comparar respuestas con opciones correctas
   c. Calcular score: (respuestas_correctas / total_preguntas) * 100
   d. Determinar si aprobo: score >= passing_score (default 70%)
   e. Guardar resultado con attempt_number
   f. Retornar resultado con explicaciones
       |
6. Frontend muestra resultados:
   a. Score total
   b. Aprobado / No aprobado
   c. Cada pregunta con respuesta del usuario vs correcta
   d. Explicacion de cada respuesta (solo para Pro)
   e. Boton "Reintentar" si no aprobo (solo para Pro)
```

### Calculo de Score

```go
func calculateScore(questions []model.Question, answers []SubmittedAnswer) int {
    if len(questions) == 0 {
        return 0
    }

    totalPoints := 0
    earnedPoints := 0

    for _, q := range questions {
        totalPoints += q.Points
        answer := findAnswer(answers, q.ID)
        if answer == nil {
            continue
        }

        correctOptions := getCorrectOptionIndices(q.Options)

        switch q.QuestionType {
        case "single_choice":
            // Una sola opcion: correcto si coincide exactamente
            if len(answer.SelectedOptions) == 1 &&
               contains(correctOptions, answer.SelectedOptions[0]) {
                earnedPoints += q.Points
            }

        case "multiple_choice":
            // Multiples opciones: correcto si selecciono exactamente las correctas
            if sameElements(answer.SelectedOptions, correctOptions) {
                earnedPoints += q.Points
            }

        case "true_false":
            // Igual que single choice
            if len(answer.SelectedOptions) == 1 &&
               contains(correctOptions, answer.SelectedOptions[0]) {
                earnedPoints += q.Points
            }
        }
    }

    return (earnedPoints * 100) / totalPoints
}
```

---

## 6. Pipeline de Curacion de Contenido

### Proceso para Agregar Recursos

```
Fase 1: Descubrimiento
  - Busqueda manual de recursos por expertos en el tema
  - Sugerencias de la comunidad (Fase 2+)
  - Alertas de nuevos cursos en plataformas populares
       |
Fase 2: Evaluacion
  - Revisar contenido completo del recurso
  - Criterios de evaluacion:
    * Precision tecnica (sin errores conceptuales)
    * Actualizacion (publicado o actualizado en los ultimos 2 anos)
    * Calidad de produccion (audio, video, formato legible)
    * Profundidad apropiada para el nivel del hito
    * Idioma (espanol e ingles)
    * Accesibilidad (subtitulos, texto alternativo)
  - Calificacion interna: 1-5 (solo se publican >= 3)
       |
Fase 3: Clasificacion
  - Asignar a ruta y hito especifico
  - Clasificar tipo: video, articulo, curso, documentacion, etc.
  - Marcar como recomendado (1 por hito) o alternativo
  - Determinar si es gratuito o de pago
  - Estimar duracion en minutos
  - Agregar tags relevantes
       |
Fase 4: Publicacion
  - Cargar en la base de datos via admin panel
  - Sincronizar con Algolia para busqueda
  - Verificar que los links funcionan (link check automatizado)
       |
Fase 5: Mantenimiento
  - Verificacion trimestral de links activos
  - Reemplazo de recursos obsoletos
  - Incorporacion de calificaciones de la comunidad
  - Promover/degradar recursos basado en ratings
```

### Validacion Automatica de Recursos

```go
// pkg/resourcechecker/checker.go

type ResourceChecker struct {
    httpClient *http.Client
}

type CheckResult struct {
    URL        string
    IsActive   bool
    StatusCode int
    Error      string
    CheckedAt  time.Time
}

func (c *ResourceChecker) CheckResource(ctx context.Context, url string) CheckResult {
    result := CheckResult{
        URL:       url,
        CheckedAt: time.Now(),
    }

    req, err := http.NewRequestWithContext(ctx, http.MethodHead, url, nil)
    if err != nil {
        result.Error = err.Error()
        return result
    }

    resp, err := c.httpClient.Do(req)
    if err != nil {
        result.Error = err.Error()
        return result
    }
    defer resp.Body.Close()

    result.StatusCode = resp.StatusCode
    result.IsActive = resp.StatusCode >= 200 && resp.StatusCode < 400

    return result
}

// Job programado: verificar todos los recursos cada semana
func (c *ResourceChecker) CheckAllResources(ctx context.Context, repo repository.ResourceRepository) {
    resources, err := repo.FindAll(ctx)
    if err != nil {
        log.Error("failed to load resources for checking", "error", err)
        return
    }

    for _, r := range resources {
        result := c.CheckResource(ctx, r.URL)
        if !result.IsActive {
            log.Warn("resource link broken",
                "resource_id", r.ID,
                "url", r.URL,
                "status", result.StatusCode,
            )
            // Marcar recurso como inactivo para revision manual
            repo.MarkInactive(ctx, r.ID)
        }
    }
}
```

---

## 7. Motor de Recomendaciones

### Algoritmo de Recomendacion de Siguiente Recurso

El motor de recomendaciones sugiere el siguiente recurso mas relevante basandose en tres senales:

```
Score de Recomendacion = (w1 * RelevanciaDeProgreso) +
                         (w2 * CalificacionDePeers) +
                         (w3 * PreferenciaDeFormato)

Donde:
  w1 = 0.5 (peso de relevancia por progreso)
  w2 = 0.3 (peso de calificacion de peers similares)
  w3 = 0.2 (peso de preferencia de formato del usuario)
```

### Relevancia de Progreso (w1 = 0.5)

```
- Recurso del hito actual disponible:           score = 1.0
- Recurso del siguiente hito disponible:         score = 0.7
- Recurso del hito dos posiciones adelante:      score = 0.3
- Recurso de hitos ya completados:               score = 0.0
- Recurso recomendado (is_recommended = true):   bonus + 0.2
```

### Calificacion de Peers (w2 = 0.3)

```
- Usuarios con perfil similar (misma experiencia, mismo objetivo):
  - Calcular calificacion promedio del recurso entre peers
  - Normalizar a 0-1 (rating / 5)
  - Si menos de 5 ratings -> usar calificacion global
```

### Preferencia de Formato (w3 = 0.2)

```
- Trackear que tipo de recursos completa mas el usuario:
  - Si el usuario consume mas videos -> bonus para videos
  - Si prefiere articulos -> bonus para articulos
  - Calculado como: completaciones_tipo / total_completaciones
```

```go
// internal/service/recommendation_service.go

type RecommendationService struct {
    resourceRepo  repository.ResourceRepository
    progressRepo  repository.ProgressRepository
    reviewRepo    repository.ReviewRepository
}

func (s *RecommendationService) GetNextResources(
    ctx context.Context,
    userID uuid.UUID,
    pathID uuid.UUID,
    limit int,
) ([]model.ScoredResource, error) {
    // 1. Obtener hitos disponibles del usuario
    availableHitos, err := s.getAvailableMilestones(ctx, userID, pathID)
    if err != nil {
        return nil, err
    }

    // 2. Obtener recursos de esos hitos
    resources, err := s.resourceRepo.FindByMilestones(ctx, milestoneIDs(availableHitos))
    if err != nil {
        return nil, err
    }

    // 3. Obtener preferencias del usuario
    prefs, err := s.getUserPreferences(ctx, userID)
    if err != nil {
        return nil, err
    }

    // 4. Calcular score para cada recurso
    scored := make([]model.ScoredResource, 0, len(resources))
    for _, r := range resources {
        progressScore := s.calcProgressScore(r, availableHitos)
        peerScore := s.calcPeerScore(ctx, r.ID, userID)
        formatScore := s.calcFormatScore(r.ResourceType, prefs)

        totalScore := 0.5*progressScore + 0.3*peerScore + 0.2*formatScore

        scored = append(scored, model.ScoredResource{
            Resource: r,
            Score:    totalScore,
        })
    }

    // 5. Ordenar por score descendente
    sort.Slice(scored, func(i, j int) bool {
        return scored[i].Score > scored[j].Score
    })

    // 6. Retornar top N
    if len(scored) > limit {
        scored = scored[:limit]
    }

    return scored, nil
}
```

---

## 8. Estrategia de CDN (Cloudflare)

### Capas de Cache

```
Capa 1: Cache del Navegador (Browser Cache)
  - Assets estaticos (JS, CSS, imagenes): max-age=31536000 (1 ano) + inmutable
  - HTML: no-cache (siempre revalidar con servidor)
  - API responses: no-store (nunca cachear en browser)

Capa 2: Cloudflare Edge Cache
  - Assets estaticos: cache forever (invalidar con hash en filename)
  - Paginas SSR publicas (landing, catalogo): cache 5 minutos
  - API GET publicas (rutas, recursos): cache 1 minuto
  - API autenticadas: bypass de cache (pass-through)

Capa 3: Redis Application Cache
  - Rutas populares (top 20): TTL 10 minutos
  - Arbol de habilidades por ruta: TTL 30 minutos
  - Perfil de usuario: TTL 5 minutos
  - Resultados de busqueda frecuentes: TTL 2 minutos

Capa 4: PostgreSQL (fuente de verdad)
  - Sin cache, siempre consistente
```

### Configuracion de Headers

```
Assets estaticos (Next.js _next/static/):
  Cache-Control: public, max-age=31536000, immutable

Paginas SSR publicas:
  Cache-Control: public, s-maxage=300, stale-while-revalidate=60

API publica (GET /api/v1/paths):
  Cache-Control: public, s-maxage=60, stale-while-revalidate=30

API autenticada:
  Cache-Control: private, no-store

Imagenes de usuario (avatars):
  Cache-Control: public, max-age=86400
```

### Invalidacion de Cache

```
Evento: Se actualiza una ruta de aprendizaje
  1. Actualizar en PostgreSQL
  2. Invalidar cache de Redis: "path:{slug}", "skill_tree:{pathID}"
  3. Purge Cloudflare: POST /zones/{zone}/purge_cache
     con URLs: /paths/{slug}, /api/v1/paths/{slug}
  4. Reindexar en Algolia
```

---

## 9. Rate Limiting y Prevencion de Abuso

### Limites por Tier

```
Usuarios no autenticados:
  - 30 requests/minuto por IP
  - 5 requests/minuto a endpoints de busqueda
  - 3 intentos de login por minuto por IP

Usuarios Free:
  - 60 requests/minuto por usuario
  - 10 requests/minuto a endpoints de busqueda
  - 3 envios de evaluacion por hora por assessment
  - 5 discusiones por dia
  - 20 calificaciones por hora

Usuarios Pro:
  - 120 requests/minuto por usuario
  - 30 requests/minuto a endpoints de busqueda
  - Envios de evaluacion ilimitados
  - 20 discusiones por dia
  - Calificaciones ilimitadas

Admin:
  - 300 requests/minuto
  - Sin limites en operaciones CRUD
```

### Implementacion con Redis

```go
// internal/middleware/ratelimit.go

type RateLimiter struct {
    redis  *redis.Client
    config RateLimitConfig
}

type RateLimitConfig struct {
    Window   time.Duration
    MaxReqs  int
    KeyPrefix string
}

func (rl *RateLimiter) Middleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        var key string

        // Determinar key: userID si autenticado, IP si no
        userID := c.GetString("userID")
        if userID != "" {
            key = fmt.Sprintf("%s:user:%s", rl.config.KeyPrefix, userID)
        } else {
            key = fmt.Sprintf("%s:ip:%s", rl.config.KeyPrefix, c.ClientIP())
        }

        // Sliding window counter con Redis
        now := time.Now().UnixMilli()
        windowStart := now - rl.config.Window.Milliseconds()

        pipe := rl.redis.Pipeline()
        // Remover entradas fuera de la ventana
        pipe.ZRemRangeByScore(c, key, "0", fmt.Sprintf("%d", windowStart))
        // Agregar request actual
        pipe.ZAdd(c, key, redis.Z{Score: float64(now), Member: now})
        // Contar requests en ventana
        pipe.ZCard(c, key)
        // Establecer expiracion
        pipe.Expire(c, key, rl.config.Window)

        results, err := pipe.Exec(c)
        if err != nil {
            c.Next() // Si Redis falla, permitir el request
            return
        }

        count := results[2].(*redis.IntCmd).Val()

        // Agregar headers informativos
        c.Header("X-RateLimit-Limit", fmt.Sprintf("%d", rl.config.MaxReqs))
        c.Header("X-RateLimit-Remaining", fmt.Sprintf("%d", max(0, int64(rl.config.MaxReqs)-count)))
        c.Header("X-RateLimit-Reset", fmt.Sprintf("%d", now+rl.config.Window.Milliseconds()))

        if count > int64(rl.config.MaxReqs) {
            c.JSON(429, model.AppError{
                Code:    429,
                Type:    "rate_limited",
                Message: "Demasiadas solicitudes. Intenta de nuevo en un momento.",
            })
            c.Abort()
            return
        }

        c.Next()
    }
}
```

### Prevencion de Abuso Adicional

```
1. Proteccion contra bots:
   - Cloudflare Bot Management
   - Challenge pages para trafico sospechoso
   - Honeypot fields en formularios

2. Proteccion de contenido:
   - Validar que URLs de recursos son de dominios permitidos
   - Sanitizar HTML en discusiones y resenas (bluemonday)
   - Limitar tamano de texto en campos de input

3. Proteccion de evaluaciones:
   - Las respuestas correctas nunca se envian al frontend antes del submit
   - Tiempo minimo entre preguntas (anti-script)
   - Variacion aleatoria en orden de preguntas y opciones

4. Proteccion de cuentas:
   - Lockout temporal despues de 5 intentos fallidos de login
   - Verificacion de email para operaciones sensibles
   - Deteccion de cuentas multiples por IP (alertar, no bloquear)
```

---

## 10. Estrategia de Indexacion de Base de Datos

### Principios

```
1. Indexar columnas usadas en WHERE, JOIN y ORDER BY frecuentes
2. Usar indices parciales cuando filtramos por subconjunto (is_published = true)
3. Usar indices GIN para arrays (tags) y JSONB (options)
4. Monitorear con pg_stat_user_indexes para detectar indices no usados
5. Usar EXPLAIN ANALYZE en queries criticas antes de agregar indices
```

### Indices Criticos por Tabla

```sql
-- PATHS: consultas mas frecuentes
CREATE INDEX idx_paths_published_category ON paths(category) WHERE is_published = TRUE;
CREATE INDEX idx_paths_published_difficulty ON paths(difficulty) WHERE is_published = TRUE;
CREATE INDEX idx_paths_featured ON paths(created_at DESC) WHERE is_featured = TRUE AND is_published = TRUE;
CREATE INDEX idx_paths_fulltext ON paths USING GIN(to_tsvector('spanish', title || ' ' || description));

-- MILESTONES: siempre se consultan por path
CREATE INDEX idx_milestones_path_order ON milestones(path_id, order_index);

-- RESOURCES: se consultan por milestone y se ordenan por rating
CREATE INDEX idx_resources_milestone_recommended ON resources(milestone_id, is_recommended DESC, avg_rating DESC);
CREATE INDEX idx_resources_milestone_type ON resources(milestone_id, resource_type);

-- USER_PROGRESS: dashboard del usuario
CREATE INDEX idx_progress_user_active ON user_progress(user_id, started_at DESC) WHERE status = 'active';
CREATE INDEX idx_progress_user_path ON user_progress(user_id, path_id);

-- MILESTONE_COMPLETIONS: verificar si un hito esta completado
CREATE INDEX idx_completions_user_milestone ON milestone_completions(user_id, milestone_id);

-- DISCUSSIONS: listar por hito, ordenar por actividad
CREATE INDEX idx_discussions_milestone_recent ON discussions(milestone_id, created_at DESC);
CREATE INDEX idx_discussions_milestone_popular ON discussions(milestone_id, upvote_count DESC);

-- ASSESSMENT_RESULTS: verificar si aprobo, historial
CREATE INDEX idx_results_user_assessment_score ON user_assessment_results(user_id, assessment_id, passed);

-- BOOKMARKS: recursos guardados del usuario
CREATE INDEX idx_bookmarks_user_recent ON bookmarks(user_id, created_at DESC);
```

### Queries Criticas y Sus Planes

```sql
-- Query 1: Listar rutas publicadas por categoria (pagina de catalogo)
-- Usa: idx_paths_published_category
SELECT id, slug, title, description, category, difficulty, estimated_hours
FROM paths
WHERE is_published = TRUE AND category = $1
ORDER BY created_at DESC
LIMIT 20 OFFSET $2;

-- Query 2: Arbol de habilidades (todos los hitos de una ruta con dependencias)
-- Usa: idx_milestones_path_order + join con milestone_dependencies
SELECT m.*, array_agg(md.depends_on_id) as dependencies
FROM milestones m
LEFT JOIN milestone_dependencies md ON md.milestone_id = m.id
WHERE m.path_id = $1
GROUP BY m.id
ORDER BY m.order_index;

-- Query 3: Dashboard del usuario (rutas activas con progreso)
-- Usa: idx_progress_user_active + join con paths
SELECT up.*, p.title, p.slug, p.icon_url
FROM user_progress up
JOIN paths p ON p.id = up.path_id
WHERE up.user_id = $1 AND up.status = 'active'
ORDER BY up.started_at DESC;

-- Query 4: Verificar si puede completar un hito (dependencias completadas)
-- Usa: idx_deps_milestone + idx_completions_user_milestone
SELECT md.depends_on_id,
       EXISTS(SELECT 1 FROM milestone_completions mc
              WHERE mc.user_id = $1 AND mc.milestone_id = md.depends_on_id) as is_completed
FROM milestone_dependencies md
WHERE md.milestone_id = $2;

-- Query 5: Recursos del hito ordenados por relevancia
-- Usa: idx_resources_milestone_recommended
SELECT r.*, COALESCE(b.id IS NOT NULL, FALSE) as is_bookmarked
FROM resources r
LEFT JOIN bookmarks b ON b.resource_id = r.id AND b.user_id = $1
WHERE r.milestone_id = $2
ORDER BY r.is_recommended DESC, r.avg_rating DESC, r.order_index;
```

### Monitoreo de Performance

```sql
-- Verificar indices no usados (ejecutar mensualmente)
SELECT schemaname, tablename, indexname, idx_scan
FROM pg_stat_user_indexes
WHERE idx_scan = 0
AND indexname NOT LIKE '%_pkey'
ORDER BY pg_relation_size(indexrelid) DESC;

-- Verificar tablas con sequential scans frecuentes (posible indice faltante)
SELECT schemaname, relname, seq_scan, seq_tup_read, idx_scan
FROM pg_stat_user_tables
WHERE seq_scan > 100
ORDER BY seq_tup_read DESC;

-- Identificar queries lentas (configurar en postgresql.conf)
-- log_min_duration_statement = 200  (logear queries >200ms)
```

---

## 11. Diagrama de Despliegue

```
                    +---------------------------+
                    |       GitHub Actions       |
                    |  (CI/CD Pipeline)          |
                    |                            |
                    |  1. Lint + Format           |
                    |  2. Unit Tests              |
                    |  3. Integration Tests       |
                    |  4. Build                   |
                    |  5. Deploy                  |
                    +--------+--------+----------+
                             |        |
                    +--------v--+  +--v----------+
                    |  Vercel    |  |  Railway     |
                    |            |  |              |
                    | Next.js    |  | Go API       |
                    | Frontend   |  | PostgreSQL   |
                    |            |  | Redis        |
                    +--------+--+  +--+-----------+
                             |        |
                    +--------v--------v----------+
                    |       Cloudflare            |
                    |  CDN + WAF + DNS             |
                    +----------------------------+
                                 |
                    +------------v---------------+
                    |        Internet             |
                    |     (Usuarios finales)       |
                    +----------------------------+
```

### Ambientes

| Ambiente | Proposito | URL |
|----------|----------|-----|
| Local | Desarrollo | localhost:3000 (front) / localhost:8080 (back) |
| Staging | QA y preview | staging.skillpath.dev |
| Produccion | Usuarios reales | app.skillpath.dev |

### CI/CD Pipeline (GitHub Actions)

```yaml
# Trigger: push a main o PR
1. checkout
2. setup-go + setup-node
3. Backend:
   - go vet ./...
   - go test ./... -race -coverprofile=coverage.out
   - go build -o bin/api ./cmd/api
4. Frontend:
   - npm ci
   - npm run lint
   - npm run type-check
   - npm run test
   - npm run build
5. Deploy (solo en main):
   - Vercel: deploy automatico via integracion
   - Railway: deploy via CLI con railway up
6. Post-deploy:
   - Health check en produccion
   - Notificar en Slack/Discord
```
