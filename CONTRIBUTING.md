# Contribuir a SkillPath

Gracias por tu interes en contribuir a SkillPath. Esta guia describe el proceso para colaborar en el proyecto.

## Requisitos Previos

- Go 1.22+
- Node.js 18+
- Git

## Configuracion del Entorno

1. Fork del repositorio
2. Clonar tu fork:
   ```bash
   git clone https://github.com/tu-usuario/skill-path.git
   cd skill-path
   ```
3. Configurar el backend:
   ```bash
   cd backend
   cp .env.example .env
   go mod download
   make run
   ```
4. Configurar el frontend:
   ```bash
   cd frontend
   npm install
   npm run dev
   ```

## Flujo de Trabajo

1. Crear una rama desde `main`:
   ```bash
   git checkout -b feature/mi-feature
   ```
2. Implementar los cambios
3. Ejecutar tests:
   ```bash
   cd backend && make test
   ```
4. Commit con mensajes descriptivos:
   ```bash
   git commit -m "feat: agregar filtro por dificultad en rutas"
   ```
5. Push y crear Pull Request

## Convenciones de Commits

Usamos [Conventional Commits](https://www.conventionalcommits.org/):

- `feat:` - Nueva funcionalidad
- `fix:` - Correccion de bug
- `docs:` - Cambios en documentacion
- `refactor:` - Refactorizacion sin cambio funcional
- `test:` - Agregar o modificar tests
- `chore:` - Tareas de mantenimiento

## Estructura del Codigo

### Backend (Go)
- Handlers en `internal/handler/` - Solo parseo HTTP y respuestas
- Servicios en `internal/service/` - Toda la logica de negocio
- Repositorios en `internal/repository/` - Queries a la base de datos
- Modelos en `internal/model/` - Estructuras de datos GORM
- DTOs en `internal/dto/` - Objetos de transferencia HTTP

### Frontend (Next.js)
- Paginas en `src/app/`
- Componentes reutilizables separados de la logica de pagina

## Guia de Estilo

### Go
- Seguir las convenciones estandar de Go (`gofmt`, `go vet`)
- Documentar funciones publicas con comentarios godoc
- Manejar errores explicitamente (nunca ignorar con `_`)
- Tests en archivos `_test.go` junto al codigo que testean

### TypeScript
- Tipado estricto, evitar `any`
- Componentes funcionales con hooks
- ESLint + Prettier para formateo

## Reportar Bugs

Abrir un issue con:
- Descripcion del problema
- Pasos para reproducir
- Comportamiento esperado vs actual
- Version de Go/Node.js

## Agregar Rutas de Aprendizaje

Para agregar una nueva ruta:
1. Crear la funcion en `backend/internal/repository/seed.go`
2. Agregar hitos con recursos curados
3. Agregar assessments para al menos un hito
4. Documentar en `docs/LEARNING-PATHS.md`
5. Crear PR con la nueva ruta

## Codigo de Conducta

Se espera un trato respetuoso y profesional en todas las interacciones del proyecto.
