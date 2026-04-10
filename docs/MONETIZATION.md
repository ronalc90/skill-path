# SkillPath - Modelo de Negocio y Monetizacion

**Autor:** Ronald

---

## Resumen Ejecutivo

SkillPath opera bajo un modelo **freemium** con multiples fuentes de ingreso: suscripciones individuales (B2C), licencias empresariales (B2B), comisiones de afiliados y patrocinios. El objetivo es alcanzar $15,000 MRR en 12 meses con un mix saludable entre B2C y B2B.

---

## 1. Modelo Freemium (B2C)

### Plan Gratuito (Free)
El plan gratuito es la puerta de entrada. Debe ser lo suficientemente valioso para atraer usuarios, pero con limitaciones claras que motiven la conversion a Pro.

**Incluye:**
- Explorar todas las rutas de aprendizaje
- Seguir activamente 1 ruta de aprendizaje
- Arbol de habilidades interactivo (solo lectura para rutas no activas)
- Acceso a todos los recursos curados (gratuitos)
- Marcar hitos como completados
- 1 evaluacion por hito (sin reintentos)
- Perfil publico basico
- Acceso a discusiones (solo lectura)
- Anuncios discretos en la plataforma

**Limitaciones clave:**
- Solo 1 ruta activa simultaneamente
- Sin reintentos en evaluaciones
- Sin certificados de completacion
- Sin acceso a mentores
- Sin proyectos de portafolio sugeridos
- No puede participar en discusiones (solo leer)
- Sin grafico radar de habilidades

### Plan Pro ($7.99/mes o $59.99/ano)
El plan Pro elimina todas las restricciones y agrega funcionalidades premium.

**Incluye todo del plan Free, mas:**
- Rutas activas ilimitadas (seguir multiples rutas a la vez)
- Evaluaciones con reintentos ilimitados
- Explicaciones detalladas en respuestas de evaluaciones
- Certificado digital de completacion por ruta (compartible en LinkedIn)
- Acceso al sistema de mentoria
- Sugerencias personalizadas de proyectos para portafolio
- Participacion completa en discusiones (crear, responder, votar)
- Grafico radar de habilidades completo
- Prioridad en respuestas de la comunidad (badge "Pro")
- Sin anuncios
- Acceso anticipado a nuevas rutas
- Exportar progreso en PDF
- Estadisticas avanzadas de aprendizaje (tiempo por hito, comparativa con peers)

### Justificacion del Precio

| Referencia | Precio Mensual |
|------------|---------------|
| Duolingo Plus | $6.99/mes |
| Codecademy Pro | $17.99/mes |
| DataCamp | $12.42/mes |
| Pluralsight | $29/mes |
| Treehouse | $25/mes |

**$7.99/mes** se posiciona en el rango bajo, accesible para autodidactas y estudiantes en Latinoamerica. El descuento anual ($59.99/ano = $5.00/mes equivalente) incentiva la retencion a largo plazo con un ahorro del 37%.

### Estrategia de Conversion Free -> Pro

1. **Trigger natural**: cuando el usuario intenta iniciar una segunda ruta, se muestra upgrade suave
2. **Evaluaciones**: al fallar un quiz, se muestra "Con Pro puedes reintentar y ver explicaciones detalladas"
3. **Comunidad**: al intentar responder una pregunta, se muestra "Participa en la comunidad con Pro"
4. **Certificados**: al completar una ruta, "Obtiene tu certificado compartible con Pro"
5. **Trial de 14 dias**: prueba gratuita de Pro sin tarjeta de credito (activada despues de completar 3 hitos)
6. **Descuento estudiante**: 50% de descuento con email .edu verificado ($3.99/mes)

---

## 2. Licencias Empresariales (B2B)

### Plan Empresa ($15/usuario/mes)
Para empresas que quieren estructurar el upskilling de sus equipos.

**Incluye todo de Pro, mas:**
- Dashboard administrativo para managers
- Asignar rutas a empleados
- Seguimiento de progreso del equipo
- Reportes de completacion y evaluaciones
- Rutas personalizadas (crear rutas internas de la empresa)
- Integracion con SSO corporativo (SAML, OIDC)
- Soporte prioritario por email
- Onboarding asistido
- Facturacion centralizada
- SLA de 99.9% uptime

### Tiers B2B

| Tier | Usuarios | Precio/Usuario/Mes | Total Mensual | Beneficios Extra |
|------|----------|--------------------:|-------------:|-----------------|
| Startup | 5-25 | $15.00 | $75 - $375 | Dashboard basico |
| Growth | 26-100 | $12.00 | $312 - $1,200 | Rutas custom, SSO |
| Enterprise | 101-500 | $10.00 | $1,010 - $5,000 | API, account manager dedicado |
| Enterprise+ | 500+ | Personalizado | Negociable | Implementacion custom, SLA premium |

### Propuesta de Valor B2B
- **Para el Manager**: visibilidad clara del progreso de su equipo, puede asignar rutas alineadas con objetivos del negocio
- **Para HR**: programa de desarrollo profesional estructurado sin necesidad de crear contenido propio
- **Para el Empleado**: ruta clara de crecimiento, certificados para su carrera
- **ROI**: reduccion de tiempo de onboarding tecnico en 40%, reduccion de rotacion por desarrollo profesional

---

## 3. Comisiones de Afiliados

### Modelo
Cuando un recurso curado es un curso de pago (Udemy, Coursera, Pluralsight, etc.), el link incluye un codigo de afiliado. SkillPath gana una comision por cada compra realizada a traves de estos links.

### Programas de Afiliados Objetivo

| Plataforma | Comision Estimada | Precio Promedio Curso | Ingreso por Conversion |
|------------|------------------:|---------------------:|-----------------------:|
| Udemy | 15-20% | $12.99 (precio oferta) | $1.95 - $2.60 |
| Coursera | 10-45% | $49/mes | $4.90 - $22.05 |
| Pluralsight | 15-25% | $29/mes | $4.35 - $7.25 |
| Educative | 15% | $17.99/mes | $2.70 |
| Domestika | 10-15% | $9.99 | $1.00 - $1.50 |

### Implementacion
- Los recursos de pago se marcan claramente con icono y precio
- El link lleva al usuario a la plataforma con cookie de afiliado (30 dias)
- Se prioriza la transparencia: "Este link puede generar una comision para SkillPath sin costo adicional para ti"
- Los recursos gratuitos siempre se muestran primero
- Las calificaciones de la comunidad no se manipulan por comisiones

### Proyeccion Afiliados
- Estimado: 5% de usuarios hacen clic en recursos de pago por mes
- Estimado: 10% de clics convierten en compra
- Ingreso promedio por conversion: $3.00
- Con 5,000 MAU: 5,000 * 5% * 10% * $3.00 = $75/mes

---

## 4. Patrocinios de Rutas

### Concepto
Empresas tech pueden patrocinar rutas de aprendizaje relacionadas con sus productos. Por ejemplo: "Ruta: Cloud Computing con AWS" patrocinada por AWS, o "Ruta: DevOps con Azure" patrocinada por Microsoft.

### Beneficios para el Patrocinador
- Logo y marca en la ruta patrocinada
- Recursos del patrocinador incluidos como recomendados (e.g., AWS Free Tier labs)
- CTA hacia certificaciones oficiales del patrocinador
- Acceso a datos anonimizados de completacion (cuantos usuarios terminan la ruta)
- Mencion en comunicaciones ("Esta ruta fue posible gracias a [Patrocinador]")

### Precio de Patrocinios

| Tipo | Precio Trimestral | Incluye |
|------|-------------------:|---------|
| Ruta Patrocinada (Bronce) | $2,500 | Logo en la ruta + recursos recomendados |
| Ruta Patrocinada (Plata) | $5,000 | Bronce + CTA certificacion + datos anonimizados |
| Ruta Patrocinada (Oro) | $10,000 | Plata + ruta co-creada + webinar mensual |

### Patrocinadores Objetivo (Ano 1)
- AWS (rutas de cloud, backend)
- Google Cloud (rutas de data, DevOps)
- JetBrains (rutas de Java, Kotlin)
- Vercel (rutas de frontend, Next.js)
- MongoDB (rutas de backend, data)
- DigitalOcean (rutas de DevOps, infraestructura)

### Proyeccion Patrocinios
- Ano 1: 2-3 patrocinadores a nivel Bronce = $5,000 - $7,500/trimestre
- Ano 2: 5-8 patrocinadores mixtos = $20,000 - $40,000/trimestre

---

## 5. Proyecciones de Ingreso (12 Meses)

### Supuestos Base
- Crecimiento de usuarios registrados: 20% mensual despues del lanzamiento
- Tasa de conversion Free -> Pro: 5% (conservador)
- Precio Pro promedio: $6.50/mes (mix mensual/anual)
- Cuentas B2B: primeras ventas a partir del mes 6
- Afiliados: activos desde el mes 3
- Patrocinios: primer patrocinador en el mes 8

### Proyeccion Mes a Mes

| Mes | Usuarios Registrados | MAU | Suscriptores Pro | MRR Pro | MRR B2B | Afiliados | Patrocinios | MRR Total |
|-----|---------------------:|----:|------------------:|--------:|--------:|----------:|------------:|----------:|
| 1 | 200 | 120 | 0 | $0 | $0 | $0 | $0 | $0 |
| 2 | 400 | 240 | 5 | $33 | $0 | $0 | $0 | $33 |
| 3 | 650 | 390 | 15 | $98 | $0 | $25 | $0 | $123 |
| 4 | 1,000 | 600 | 35 | $228 | $0 | $45 | $0 | $273 |
| 5 | 1,500 | 900 | 60 | $390 | $0 | $68 | $0 | $458 |
| 6 | 2,200 | 1,320 | 100 | $650 | $300 | $99 | $0 | $1,049 |
| 7 | 3,200 | 1,920 | 155 | $1,008 | $600 | $144 | $0 | $1,752 |
| 8 | 4,500 | 2,700 | 230 | $1,495 | $1,050 | $203 | $833 | $3,581 |
| 9 | 6,200 | 3,720 | 330 | $2,145 | $1,500 | $279 | $833 | $4,757 |
| 10 | 8,500 | 5,100 | 460 | $2,990 | $2,100 | $383 | $833 | $6,306 |
| 11 | 11,500 | 6,900 | 630 | $4,095 | $3,000 | $518 | $1,667 | $9,280 |
| 12 | 15,000 | 9,000 | 850 | $5,525 | $4,200 | $675 | $1,667 | $12,067 |

### Resumen Anual

| Fuente de Ingreso | Ingreso Anual (Ano 1) | % del Total |
|-------------------|-----------------------:|------------:|
| Suscripciones Pro (B2C) | $18,657 | 46% |
| Licencias Empresa (B2B) | $12,750 | 31% |
| Comisiones Afiliados | $2,439 | 6% |
| Patrocinios de Rutas | $5,833 | 14% |
| **Total** | **$39,679** | **100%** |

---

## 6. Mix de Ingresos B2C vs B2B

### Evolucion del Mix

| Periodo | B2C | B2B | Otros |
|---------|----:|----:|------:|
| Meses 1-6 | 85% | 5% | 10% |
| Meses 7-12 | 50% | 35% | 15% |
| Ano 2 (proyeccion) | 35% | 50% | 15% |

### Estrategia de Balance
- **Corto plazo (Meses 1-6):** el foco esta en B2C para validar producto y construir base de usuarios. Los ingresos iniciales vendran de suscripciones Pro.
- **Mediano plazo (Meses 7-12):** se activa la venta B2B. Cada cuenta empresarial genera 10-50x mas ingreso que un usuario individual. Se inician patrocinios.
- **Largo plazo (Ano 2+):** B2B se convierte en la fuente principal de ingresos. Enterprise es mas predecible y tiene menor churn que B2C.

### Ventajas del Mix
- **B2C** proporciona volumen, datos de uso y validacion de mercado
- **B2B** proporciona ingresos predecibles y contratos mas grandes
- **Afiliados** son ingresos pasivos que escalan con el trafico
- **Patrocinios** son alto margen y refuerzan la marca

---

## 7. Estructura de Costos (Mes 12)

### Costos Fijos Mensuales

| Concepto | Costo Mensual |
|----------|-------------:|
| Railway (Backend + DB + Redis) | $50 |
| Vercel (Frontend) Pro | $20 |
| Algolia (Search) | $35 |
| Cloudflare Pro | $20 |
| Sentry (Error tracking) | $26 |
| Email transaccional (Resend) | $20 |
| Dominio + SSL | $3 |
| Herramientas desarrollo | $30 |
| **Total Infraestructura** | **$204** |

### Costos Variables

| Concepto | Costo Estimado |
|----------|-------------:|
| Procesamiento de pagos (Stripe 2.9% + $0.30) | ~$350 |
| Bandwidth adicional | ~$50 |
| **Total Variable** | **~$400** |

### Resumen Financiero (Mes 12)

| Metrica | Valor |
|---------|------:|
| MRR | $12,067 |
| Costos Totales | $604 |
| Margen Bruto | $11,463 |
| Margen Bruto % | 95% |

Nota: estos numeros no incluyen costo de equipo humano (desarrollo, contenido, marketing). El margen bruto alto es caracteristico de SaaS pero el ingreso debe cubrir salarios y marketing para ser un negocio viable.

---

## 8. Metricas Financieras Clave

| Metrica | Objetivo Mes 6 | Objetivo Mes 12 |
|---------|---------------:|----------------:|
| MRR | $1,049 | $12,067 |
| Suscriptores Pro | 100 | 850 |
| Tasa Conversion Free->Pro | 4% | 5.5% |
| Churn Mensual Pro | <8% | <5% |
| ARPU (Average Revenue Per User) | $0.80 | $1.34 |
| LTV Pro (12 meses) | $52 | $78 |
| CAC objetivo | <$15 | <$20 |
| LTV/CAC | >3.5 | >3.9 |
| Cuentas B2B | 2 | 10 |
| ARPA B2B (Avg Revenue Per Account) | $150 | $420 |

---

## 9. Riesgos y Mitigaciones

| Riesgo | Impacto | Probabilidad | Mitigacion |
|--------|---------|-------------|------------|
| Baja conversion Free->Pro | Alto | Media | Trial de 14 dias, nudges contextuales, precio accesible |
| Alto churn de Pro | Alto | Media | Contenido fresco mensual, engagement via comunidad, email de re-engagement |
| Dificultad vendiendo B2B | Alto | Alta | Empezar con startups pequenas, ofrecer piloto gratuito de 30 dias |
| Comisiones afiliados bajas | Bajo | Media | Diversificar plataformas afiliadas, negociar comisiones personalizadas |
| Patrocinadores no interesados | Medio | Media | Esperar a tener >5K MAU antes de buscar patrocinadores |
| Competidor con mas recursos | Alto | Media | Velocidad de ejecucion, enfoque en Latam, comunidad fuerte |

---

## 10. Hitos Financieros

| Hito | Objetivo | Significado |
|------|----------|-------------|
| Primer pago | Mes 2 | Validacion de que alguien pagara por el producto |
| $1K MRR | Mes 6 | Negocio viable en potencia, cubre infraestructura |
| 100 suscriptores Pro | Mes 6 | Base solida de usuarios de pago |
| Primer cliente B2B | Mes 6 | Validacion del canal enterprise |
| $5K MRR | Mes 10 | Momentum real, posible buscar inversion |
| $10K MRR | Mes 12 | Negocio sostenible (con equipo minimo) |
| $50K MRR | Mes 24 | Escala, posible equipo de 5+ personas |
