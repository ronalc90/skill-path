.PHONY: build run test lint docker docker-compose clean security help

# --- Backend ---

## build: Compilar binario del backend
build:
	cd backend && go build -o bin/skillpath ./cmd/server/

## run: Ejecutar el backend localmente
run: build
	cd backend && ./bin/skillpath

## test: Ejecutar tests con cobertura
test:
	cd backend && go test ./... -v -race -cover

## lint: Ejecutar linter de Go
lint:
	cd backend && golangci-lint run ./...

## vet: Ejecutar go vet
vet:
	cd backend && go vet ./...

## tidy: Limpiar dependencias Go
tidy:
	cd backend && go mod tidy

# --- Frontend ---

## frontend-install: Instalar dependencias del frontend
frontend-install:
	cd frontend && npm ci

## frontend-build: Compilar el frontend
frontend-build:
	cd frontend && npm run build

## frontend-dev: Iniciar frontend en modo desarrollo
frontend-dev:
	cd frontend && npm run dev

## frontend-lint: Ejecutar linter del frontend
frontend-lint:
	cd frontend && npm run lint

# --- Docker ---

## docker: Construir imagen Docker del backend
docker:
	docker build -t skillpath -f backend/Dockerfile backend/

## docker-compose: Levantar todos los servicios
docker-compose:
	docker-compose up -d

## docker-down: Detener todos los servicios
docker-down:
	docker-compose down

## docker-logs: Ver logs de los contenedores
docker-logs:
	docker-compose logs -f

# --- Seguridad ---

## security-scan: Escanear imagen Docker con Trivy
security-scan:
	docker build -t skillpath:scan -f backend/Dockerfile backend/
	trivy image --severity HIGH,CRITICAL skillpath:scan

## gosec: Ejecutar analisis de seguridad del codigo Go
gosec:
	cd backend && gosec ./...

## security: Ejecutar todos los checks de seguridad
security: lint gosec security-scan

# --- Infraestructura ---

## k8s-apply: Aplicar manifiestos de Kubernetes
k8s-apply:
	kubectl apply -f infra/kubernetes/

## k8s-delete: Eliminar recursos de Kubernetes
k8s-delete:
	kubectl delete -f infra/kubernetes/

## tf-init: Inicializar Terraform
tf-init:
	cd infra/terraform && terraform init

## tf-plan: Planificar cambios de Terraform
tf-plan:
	cd infra/terraform && terraform plan

## tf-apply: Aplicar cambios de Terraform
tf-apply:
	cd infra/terraform && terraform apply

# --- Utilidades ---

## clean: Limpiar artefactos de compilacion
clean:
	rm -rf backend/bin/ backend/coverage.out backend/coverage.html
	rm -rf frontend/.next frontend/out

## help: Mostrar comandos disponibles
help:
	@echo "Comandos disponibles:"
	@sed -n 's/^##//p' $(MAKEFILE_LIST) | sort
