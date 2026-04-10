# Decisiones de Arquitectura (ADR)

Registro de las decisiones tecnicas mas relevantes del proyecto SkillPath, documentadas en formato ADR (Architecture Decision Record).

---

## ADR-001: Go sobre Node.js para el API

**Estado:** Aceptada
**Fecha:** 2025-01

### Contexto
Se necesitaba elegir el lenguaje para el backend. Las opciones principales eran Node.js (TypeScript) y Go.

### Decision
Se eligio Go como lenguaje para el API.

### Razones
- **Rendimiento:** Go compila a binarios nativos con rendimiento cercano a C. Node.js requiere un runtime (V8).
- **Concurrencia:** Goroutines y channels son nativos del lenguaje, ideales para manejar multiples requests simultaneos.
- **Binarios estaticos:** Un solo binario sin dependencias externas simplifica el despliegue enormemente.
- **Tipado estatico:** Errores detectados en compilacion en lugar de runtime.
- **Bajo consumo de memoria:** Crucial para reducir costos de hosting.

### Consecuencias
- Ecosistema mas pequeno que Node.js para web
- Curva de aprendizaje para desarrolladores acostumbrados a JavaScript
- Excelente rendimiento con minimo overhead

---

## ADR-002: Gin sobre Echo

**Estado:** Aceptada
**Fecha:** 2025-01

### Contexto
Se necesitaba un framework HTTP para Go. Las opciones principales eran Gin, Echo, Chi y Fiber.

### Decision
Se eligio Gin como framework HTTP.

### Razones
- **Ecosistema:** Es el framework Go mas popular con la comunidad mas grande
- **Documentacion:** Documentacion extensa y abundantes tutoriales
- **Rendimiento:** Basado en httprouter, uno de los routers mas rapidos
- **Middleware:** Ecosistema robusto de middleware (CORS, logging, recovery)
- **Validacion:** Integracion nativa con go-playground/validator para validacion de requests
- **Madurez:** Proyecto estable con versionado semantico

### Alternativas Consideradas
- **Echo:** Rendimiento similar pero comunidad mas pequena
- **Chi:** Mas ligero pero sin validacion integrada
- **Fiber:** API similar a Express pero basado en fasthttp (menos compatible con net/http)

---

## ADR-003: GORM sobre sqlx

**Estado:** Aceptada
**Fecha:** 2025-01

### Contexto
Se necesitaba una forma de interactuar con la base de datos. Las opciones eran un ORM completo (GORM) o un query builder/mapper (sqlx).

### Decision
Se eligio GORM como ORM.

### Razones
- **Productividad:** Migraciones automaticas con `AutoMigrate()` eliminan la necesidad de archivos de migracion SQL manuales
- **Relaciones:** Preloading de relaciones (path -> milestones -> resources) con una linea de codigo
- **Convenciones:** Naming conventions y timestamps automaticos reducen boilerplate
- **Hooks:** Posibilidad de agregar logica pre/post operaciones
- **Multi-driver:** Soporte transparente para SQLite y PostgreSQL

### Trade-offs
- Menos control sobre el SQL generado comparado con sqlx
- Consultas complejas pueden requerir SQL raw
- Overhead minimo de rendimiento por la capa de abstraccion

### Mitigacion
- Para consultas criticas de rendimiento se puede usar `db.Raw()` o `db.Exec()`
- El preloading genera queries separadas (N+1 controlado) en lugar de JOINs complejos

---

## ADR-004: SQLite para Desarrollo

**Estado:** Aceptada
**Fecha:** 2025-01

### Contexto
Se necesitaba una base de datos para desarrollo local que fuera simple de configurar.

### Decision
Usar SQLite como base de datos de desarrollo y PostgreSQL para produccion.

### Razones
- **Zero-config:** No requiere instalar ni configurar un servidor de base de datos
- **Archivo unico:** Toda la base de datos es un archivo `skillpath.db`
- **Seed automatico:** Al iniciar la aplicacion, si la BD esta vacia se pobla automaticamente
- **Testing:** Ideal para tests rapidos sin dependencias externas
- **GORM compatible:** Mismo codigo funciona con SQLite y PostgreSQL gracias a GORM

### Consecuencias
- Algunas features de PostgreSQL (como JSON nativo, full-text search) no estan disponibles en desarrollo
- El tipo `JSONSlice` personalizado maneja la serializacion JSON de forma compatible con ambos drivers

---

## ADR-005: Clean Architecture

**Estado:** Aceptada
**Fecha:** 2025-01

### Contexto
Se necesitaba definir la estructura del proyecto para que fuera mantenible a largo plazo.

### Decision
Implementar Clean Architecture con tres capas: Handler -> Service -> Repository.

### Razones
- **Separacion de responsabilidades:** Cada capa tiene una funcion clara
- **Testabilidad:** Los servicios se pueden testear sin HTTP, los repositorios sin logica de negocio
- **Mantenibilidad:** Cambiar la base de datos solo afecta la capa de repositorio
- **Convencion Go:** Sigue las convenciones del standard project layout de Go (`internal/`, `cmd/`, `pkg/`)

### Estructura

```
cmd/server/main.go          -> Punto de entrada, DI
internal/handler/            -> Capa HTTP (input/output)
internal/service/            -> Capa de negocio
internal/repository/         -> Capa de datos
internal/model/              -> Entidades de dominio
internal/dto/                -> Objetos de transferencia
internal/middleware/          -> Middleware HTTP
internal/config/             -> Configuracion
pkg/                         -> Paquetes reutilizables
```

### Inyeccion de Dependencias
Se eligio DI manual (sin frameworks como Wire o Dig) porque:
- El grafo de dependencias es pequeno y claro
- Importaciones circulares se detectan en compilacion
- No agrega complejidad innecesaria

---

## ADR-006: Slug-based Routing

**Estado:** Aceptada
**Fecha:** 2025-01

### Contexto
Se necesitaba definir como identificar las rutas de aprendizaje en las URLs.

### Decision
Usar slugs legibles en lugar de IDs numericos para las URLs de rutas.

### Ejemplo
- `/api/v1/paths/backend-developer-java` (slug)
- vs `/api/v1/paths/1` (ID numerico)

### Razones
- **SEO:** URLs legibles son mejor rankeadas por motores de busqueda
- **UX:** Los usuarios pueden entender el contenido de la URL
- **Compartibilidad:** Las URLs son auto-descriptivas al compartirlas
- **Estabilidad:** Los slugs no cambian si se reinicia la base de datos (los IDs auto-incrementales si)

### Implementacion
- Cada `LearningPath` tiene un campo `slug` con indice unico
- Los endpoints de progreso usan `:slug` como parametro y resuelven internamente al ID

---

## ADR-007: DAG para Skill Tree

**Estado:** Aceptada (parcialmente implementada)
**Fecha:** 2025-01

### Contexto
Las habilidades tecnologicas tienen dependencias naturales. Por ejemplo, no se puede aprender React sin saber JavaScript.

### Decision
Modelar las rutas de aprendizaje como un DAG (Directed Acyclic Graph) donde los hitos tienen un orden secuencial definido por `order_index`.

### Estado Actual
- Los hitos se ordenan por `order_index` (1, 2, 3...)
- No hay validacion de prerrequisitos todavia
- El usuario puede completar hitos en cualquier orden

### Roadmap
- Agregar campo `prerequisites` a milestones (array de milestone IDs)
- Validar que los prerrequisitos esten completados antes de permitir avanzar
- Visualizar el grafo de dependencias en el frontend
