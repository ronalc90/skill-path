# Motor de Evaluacion

El motor de evaluacion de SkillPath permite a los usuarios validar su conocimiento mediante quizzes asociados a hitos especificos de cada ruta de aprendizaje.

## Como Funciona

### Flujo del Usuario

1. El usuario avanza en una ruta de aprendizaje completando recursos
2. Al llegar a un hito que tiene assessment, puede tomar el quiz
3. El sistema presenta preguntas de opcion multiple (sin revelar respuestas correctas)
4. El usuario envia sus respuestas
5. El sistema calcula el score y determina si aprobo
6. Se muestra el resultado detallado con explicaciones por pregunta

### Estructura de Preguntas

Cada pregunta (`SkillAssessment`) contiene:

```
SkillAssessment
  |-- id                # Identificador unico
  |-- milestone_id      # Hito al que pertenece
  |-- question          # Texto de la pregunta
  |-- options           # Array JSON de opciones (4 opciones tipicamente)
  |-- correct_answer    # Indice de la respuesta correcta (0-based)
  |-- explanation       # Explicacion de por que es correcta
```

### Ejemplo de Pregunta

```json
{
  "question": "Cual es la diferencia entre == y .equals() en Java?",
  "options": [
    "== compara referencias de objetos, .equals() compara contenido",
    "Son identicos en funcionalidad",
    "== solo funciona con primitivos",
    ".equals() solo funciona con Strings"
  ],
  "correct_answer": 0,
  "explanation": "El operador == compara referencias (direcciones de memoria), mientras que .equals() compara el contenido logico de los objetos."
}
```

## Algoritmo de Scoring

### Calculo del Score

```
score = cantidad_de_respuestas_correctas
total_questions = cantidad_total_de_preguntas_del_hito
porcentaje = (score / total_questions) * 100
```

### Umbral de Aprobacion

```
PASSING_SCORE_PERCENTAGE = 70%
passed = porcentaje >= 70
```

El umbral del 70% esta definido como constante en `assessment_service.go`:

```go
const passingScorePercentage = 70
```

### Ejemplo de Calculo

Para un quiz de 3 preguntas:
- 3/3 correctas = 100% -> APROBADO
- 2/3 correctas = 66.7% -> NO APROBADO
- 3/3 es el minimo para aprobar con 3 preguntas (dado que 2/3 = 66.7% < 70%)

Para un quiz de 10 preguntas:
- 7/10 correctas = 70% -> APROBADO
- 6/10 correctas = 60% -> NO APROBADO

## Almacenamiento de Resultados

Cada intento se almacena en `AssessmentResult`:

```
AssessmentResult
  |-- id
  |-- user_id          # Usuario que tomo el quiz
  |-- milestone_id     # Hito evaluado
  |-- score            # Respuestas correctas
  |-- total_questions  # Total de preguntas
  |-- passed_at        # Timestamp si aprobo (NULL si no)
  |-- created_at       # Cuando se tomo el quiz
```

Los usuarios pueden tomar el quiz multiples veces. Cada intento genera un nuevo registro.

## Seguridad

- **Las respuestas correctas nunca se envian al obtener el quiz.** El endpoint `GET /milestones/:id/assessment` solo retorna `id`, `question` y `options`.
- **Las respuestas correctas se revelan despues del envio.** El endpoint `POST /milestones/:id/assessment/submit` retorna `correct_answer` y `explanation` por cada pregunta.
- **Autenticacion requerida.** Ambos endpoints requieren un token JWT valido.

## Respuesta del Quiz

Al enviar las respuestas, el usuario recibe:

```json
{
  "score": 2,
  "total_questions": 3,
  "passed": false,
  "details": [
    {
      "question_id": 1,
      "is_correct": true,
      "correct_answer": 0,
      "explanation": "Explicacion detallada..."
    },
    {
      "question_id": 2,
      "is_correct": false,
      "correct_answer": 2,
      "explanation": "Explicacion detallada..."
    },
    {
      "question_id": 3,
      "is_correct": true,
      "correct_answer": 2,
      "explanation": "Explicacion detallada..."
    }
  ]
}
```

## Assessments Disponibles

Actualmente hay assessments para los siguientes hitos:

| Ruta | Hito | Preguntas |
|------|------|-----------|
| Backend Developer con Java | Fundamentos de Java | 3 |
| Backend Developer con Java | Spring Boot Essentials | 3 |
| Frontend Developer con React | React Core Concepts | 3 |
| Data Analyst con Python | NumPy y Pandas | 3 |
| Mobile Developer con Flutter | Flutter UI Fundamentals | 3 |

## Integracion con Dashboard

Los resultados de assessments aparecen en:
- **Actividad reciente** del dashboard como tipo `"assessment"`
- El sistema registra cada intento para tracking historico

## Consideraciones de Diseno

1. **Opcion multiple** se eligio por su facilidad de evaluacion automatica y experiencia de usuario simple
2. **Explicaciones** en cada pregunta transforman el quiz en una herramienta de aprendizaje, no solo de evaluacion
3. **Reintentos ilimitados** permiten al usuario volver a intentar despues de estudiar mas
4. **70% como umbral** es un estandar comun en certificaciones tecnologicas
