# Motor de Evaluacion

El motor de evaluacion permite validar el conocimiento adquirido en cada hito mediante quizzes de opcion multiple.

## Flujo

```
1. Usuario solicita quiz       GET  /milestones/:id/assessment
                                    -> Retorna preguntas SIN respuestas correctas
                                    
2. Usuario responde            POST /milestones/:id/assessment/submit
                                    -> Recibe score, passed, y detalle por pregunta
                                    
3. Sistema almacena resultado  -> AssessmentResult en la base de datos
```

## Modelo de Datos

### SkillAssessment (Pregunta)

| Campo | Tipo | Descripcion |
|-------|------|-------------|
| `id` | uint | PK auto-incremental |
| `milestone_id` | uint | FK al hito al que pertenece |
| `question` | text | Texto de la pregunta |
| `options` | JSON | Array de opciones (tipicamente 4) |
| `correct_answer` | int | Indice de la respuesta correcta (0-based) |
| `explanation` | text | Por que esta respuesta es correcta |

### AssessmentResult (Resultado)

| Campo | Tipo | Descripcion |
|-------|------|-------------|
| `id` | uint | PK auto-incremental |
| `user_id` | uint | FK al usuario |
| `milestone_id` | uint | FK al hito evaluado |
| `score` | int | Respuestas correctas |
| `total_questions` | int | Total de preguntas |
| `passed_at` | timestamp | NULL si no aprobo, timestamp si aprobo |
| `created_at` | timestamp | Cuando se tomo el quiz |

## Algoritmo de Scoring

### Paso 1: Recopilar preguntas del hito
```go
assessments, err := s.assessmentRepo.FindByMilestoneID(milestoneID)
```

### Paso 2: Comparar respuestas
```go
for _, answer := range req.Answers {
    question := questionMap[answer.QuestionID]
    isCorrect := answer.SelectedIndex == question.CorrectAnswer
    if isCorrect {
        score++
    }
}
```

### Paso 3: Calcular aprobacion
```go
const passingScorePercentage = 70
passed := float64(score)/float64(totalQuestions)*100 >= float64(passingScorePercentage)
```

### Paso 4: Registrar timestamp si aprobo
```go
if passed {
    now := time.Now()
    result.PassedAt = &now
}
```

## Tabla de Aprobacion

| Preguntas | Minimo para aprobar | Porcentaje |
|-----------|---------------------|------------|
| 3 | 3/3 | 100% (2/3 = 66.7% < 70%) |
| 5 | 4/5 | 80% |
| 10 | 7/10 | 70% |
| 15 | 11/15 | 73.3% |
| 20 | 14/20 | 70% |

## Seguridad

1. **GET** del assessment: solo retorna `id`, `question`, `options`. No incluye `correct_answer` ni `explanation`.
2. **POST** submit: retorna `correct_answer` y `explanation` solo despues de que el usuario envie sus respuestas.
3. Ambos endpoints requieren autenticacion JWT.

## Quizzes Disponibles

Hay 15 preguntas distribuidas en 5 hitos (3 preguntas cada uno):

- Fundamentos de Java (Ruta Backend)
- Spring Boot Essentials (Ruta Backend)
- React Core Concepts (Ruta Frontend)
- NumPy y Pandas (Ruta Data)
- Flutter UI Fundamentals (Ruta Mobile)

## Reintentos

No hay limite de intentos. Cada envio crea un nuevo `AssessmentResult`. Esto permite:
- Ver el historial de intentos
- Identificar areas de mejora
- Medir la velocidad de aprendizaje

Documentacion detallada en [docs/ASSESSMENT-ENGINE.md](../ASSESSMENT-ENGINE.md).
