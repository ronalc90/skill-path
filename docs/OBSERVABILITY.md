# Observabilidad - SkillPath

## Arquitectura de Monitoreo

```
+------------------+     +------------------+     +------------------+
|  SkillPath API   | --> |   Prometheus     | --> |   Grafana        |
|  (metricas HTTP) |     |   (recoleccion)  |     |   (dashboards)   |
|  :3005/metrics   |     |   :9090          |     |   :3001          |
+------------------+     +------------------+     +------------------+
```

## Metricas Expuestas

El middleware de Prometheus (`internal/middleware/metrics.go`) expone las siguientes metricas en `/metrics`:

### Contadores

| Metrica | Tipo | Labels | Descripcion |
|---------|------|--------|-------------|
| `skillpath_http_requests_total` | Counter | method, path, status | Total de requests HTTP |

### Histogramas

| Metrica | Tipo | Labels | Descripcion |
|---------|------|--------|-------------|
| `skillpath_http_request_duration_seconds` | Histogram | method, path | Duracion de requests en segundos |

### Gauges

| Metrica | Tipo | Descripcion |
|---------|------|-------------|
| `skillpath_http_requests_in_flight` | Gauge | Requests siendo procesados actualmente |

## Configuracion

### Prometheus

Archivo de configuracion: `monitoring/prometheus.yml`

- Intervalo de scraping: 15 segundos
- Endpoint: `http://backend:3005/metrics`
- Retencion de datos: 15 dias

### Grafana

- URL: `http://localhost:3001`
- Usuario por defecto: `admin`
- Password por defecto: `admin` (cambiar en produccion via `GRAFANA_PASSWORD`)

## Inicio Rapido

```bash
# Levantar el stack completo de monitoreo
docker-compose up -d

# Verificar que las metricas estan disponibles
curl http://localhost:3005/metrics

# Acceder a Prometheus
open http://localhost:9090

# Acceder a Grafana
open http://localhost:3001
```

## Consultas Utiles en Prometheus

### Tasa de requests por segundo

```promql
rate(skillpath_http_requests_total[5m])
```

### Latencia P95 por endpoint

```promql
histogram_quantile(0.95, rate(skillpath_http_request_duration_seconds_bucket[5m]))
```

### Tasa de errores (5xx)

```promql
sum(rate(skillpath_http_requests_total{status=~"5.."}[5m]))
/
sum(rate(skillpath_http_requests_total[5m]))
```

### Requests en vuelo

```promql
skillpath_http_requests_in_flight
```

### Top 5 endpoints mas lentos

```promql
topk(5, histogram_quantile(0.95, rate(skillpath_http_request_duration_seconds_bucket[5m])))
```

## Alertas Sugeridas

Para configurar en Prometheus/Alertmanager:

| Alerta | Condicion | Severidad |
|--------|-----------|-----------|
| Alta latencia | P95 > 2s por 5 min | Warning |
| Tasa de errores alta | 5xx > 5% por 5 min | Critical |
| Servicio caido | up == 0 por 1 min | Critical |
| Muchos requests en vuelo | in_flight > 100 por 2 min | Warning |

## Dashboards Recomendados de Grafana

1. **Vista general del API**: requests/s, latencia, errores, uptime
2. **Detalle por endpoint**: desglose por ruta y metodo HTTP
3. **Recursos del sistema**: CPU, memoria, red (metricas del contenedor)

## Logs

### Docker Compose

```bash
# Todos los servicios
docker-compose logs -f

# Solo el backend
docker-compose logs -f backend

# Solo Prometheus
docker-compose logs -f prometheus
```

### AWS ECS

Los logs se envian a CloudWatch Logs:
- Log group: `/ecs/skillpath`
- Stream prefix: `ecs`

### Kubernetes

```bash
# Logs del pod
kubectl logs -l app=skillpath -n skillpath -f

# Logs de un pod especifico
kubectl logs skillpath-api-xxxxx -n skillpath
```

## Autor

Desarrollado por **Ronald**.
