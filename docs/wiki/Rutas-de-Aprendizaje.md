# Rutas de Aprendizaje

SkillPath incluye 4 rutas de aprendizaje pre-cargadas que cubren las carreras tecnologicas mas demandadas.

## Resumen

| Ruta | Slug | Categoria | Dificultad | Horas | Hitos |
|------|------|-----------|------------|-------|-------|
| Backend Developer con Java | `backend-developer-java` | Backend | Intermedio | 120 | 8 |
| Frontend Developer con React | `frontend-developer-react` | Frontend | Principiante | 100 | 7 |
| Data Analyst con Python | `data-analyst-python` | Data | Principiante | 90 | 6 |
| Mobile Developer con Flutter | `mobile-developer-flutter` | Mobile | Intermedio | 85 | 6 |

**Total:** 395 horas de contenido, 27 hitos, 50+ recursos curados.

## Backend Developer con Java

Ruta completa desde los fundamentos del lenguaje hasta microservicios y despliegue.

**Hitos:**
1. Fundamentos de Java (15h) - POO, colecciones, tipos
2. Spring Boot Essentials (20h) - REST APIs, DI, testing
3. Bases de Datos con JPA/Hibernate (15h) - ORM, JPQL, transacciones
4. Seguridad con Spring Security (15h) - JWT, OAuth2
5. Testing y Calidad de Codigo (12h) - JUnit 5, Mockito
6. Microservicios (18h) - Patrones, service discovery
7. Docker y Kubernetes (15h) - Contenedores, orquestacion
8. CI/CD y Despliegue (10h) - GitHub Actions, monitoreo

## Frontend Developer con React

Desde HTML/CSS basico hasta aplicaciones complejas con Next.js.

**Hitos:**
1. HTML, CSS y JavaScript Moderno (20h) - Flexbox, Grid, ES6+
2. TypeScript Fundamentals (12h) - Tipos, interfaces, genericos
3. React Core Concepts (18h) - Components, hooks, state
4. Estado Global y Data Fetching (15h) - Context, Zustand, React Query
5. Next.js y SSR (15h) - App Router, Server Components
6. Testing Frontend (10h) - Jest, Testing Library, Cypress
7. Performance y Deploy (10h) - Code splitting, Vercel

## Data Analyst con Python

Analisis de datos desde los fundamentos hasta machine learning basico.

**Hitos:**
1. Python Fundamentals (18h) - Sintaxis, funciones, modulos
2. NumPy y Pandas (18h) - DataFrames, limpieza de datos
3. Visualizacion de Datos (15h) - Matplotlib, Seaborn, Plotly
4. SQL para Analistas (12h) - Joins, agregaciones, subqueries
5. Estadistica Aplicada (15h) - Tests de hipotesis, regresion
6. Machine Learning Basico (12h) - Scikit-learn, clasificacion

## Mobile Developer con Flutter

Desarrollo multiplataforma desde Dart hasta publicacion en stores.

**Hitos:**
1. Dart Programming Language (12h) - Async/await, null safety
2. Flutter UI Fundamentals (18h) - Widgets, Material Design
3. State Management (15h) - Provider, Riverpod, BLoC
4. Networking y APIs (12h) - HTTP, REST, JSON serialization
5. Persistencia y Base de Datos (15h) - SQLite, Firebase
6. Testing y Publicacion (13h) - Unit/widget/integration tests

## Tipos de Recursos

Cada hito contiene recursos de diferentes tipos:

| Tipo | Descripcion |
|------|-------------|
| `course` | Curso completo (Udemy, Coursera, freeCodeCamp) |
| `book` | Libro o e-book |
| `video` | Video tutorial (YouTube) |
| `tutorial` | Tutorial interactivo o guia paso a paso |
| `article` | Articulo o documentacion oficial |
| `exercise` | Ejercicio practico o proyecto |

## API

```bash
# Listar todas las rutas
curl http://localhost:3005/api/v1/paths

# Filtrar por categoria
curl "http://localhost:3005/api/v1/paths?category=Backend"

# Ver detalle con hitos y recursos
curl http://localhost:3005/api/v1/paths/backend-developer-java
```

Documentacion detallada en [docs/LEARNING-PATHS.md](../LEARNING-PATHS.md).
