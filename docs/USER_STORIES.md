# SkillPath - Historias de Usuario

**Autor:** Ronald

---

## Formato

Cada historia de usuario sigue el formato:

```
Como [tipo de usuario]
Quiero [accion]
Para [beneficio]

Criterios de Aceptacion:
- [criterio 1]
- [criterio 2]
- ...

Prioridad: Alta | Media | Baja
Sprint estimado: N
```

---

## Epic 1: Descubrimiento de Rutas

### US-1.1: Explorar catalogo de rutas
**Como** visitante o usuario registrado
**Quiero** ver todas las rutas de aprendizaje disponibles en un catalogo organizado
**Para** encontrar una ruta que se alinee con mis objetivos profesionales

**Criterios de Aceptacion:**
- Se muestra un grid de tarjetas de rutas con titulo, descripcion corta, categoria, dificultad y duracion estimada
- Las rutas se cargan con paginacion (20 por pagina)
- Se muestra un skeleton loader mientras cargan los datos
- Las tarjetas son clickeables y llevan al detalle de la ruta
- La pagina es responsive (1 columna en mobile, 2 en tablet, 3 en desktop)
- Si no hay rutas que coincidan con los filtros, se muestra un mensaje informativo

**Prioridad:** Alta
**Sprint:** 4

---

### US-1.2: Filtrar rutas por categoria
**Como** visitante o usuario registrado
**Quiero** filtrar las rutas por categoria (Backend, Frontend, Data, Mobile, DevOps)
**Para** ver solo las rutas relevantes a mi area de interes

**Criterios de Aceptacion:**
- Se muestran botones o tabs de filtro por categoria
- Al seleccionar una categoria, el catalogo se actualiza inmediatamente (sin recarga de pagina)
- Se puede seleccionar solo una categoria a la vez
- Existe una opcion "Todas" para quitar el filtro
- El filtro activo se refleja en la URL (query param) para poder compartir
- El conteo de rutas por categoria se muestra junto al nombre

**Prioridad:** Alta
**Sprint:** 4

---

### US-1.3: Filtrar rutas por dificultad
**Como** visitante o usuario registrado
**Quiero** filtrar las rutas por nivel de dificultad (Principiante, Intermedio, Avanzado)
**Para** encontrar rutas apropiadas para mi nivel actual

**Criterios de Aceptacion:**
- Se muestran opciones de filtro por dificultad
- Los filtros de dificultad se pueden combinar con el filtro de categoria
- El nivel de dificultad se muestra con un indicador visual (color o icono)
- Los filtros se pueden limpiar con un boton "Limpiar filtros"

**Prioridad:** Media
**Sprint:** 4

---

### US-1.4: Buscar rutas por texto
**Como** visitante o usuario registrado
**Quiero** buscar rutas de aprendizaje por palabras clave
**Para** encontrar rapidamente una ruta especifica que tengo en mente

**Criterios de Aceptacion:**
- Existe un campo de busqueda prominente en la parte superior del catalogo
- La busqueda es instantanea (resultados mientras se escribe, con debounce de 300ms)
- Se busca en titulo, descripcion y tags de las rutas
- Los terminos de busqueda se resaltan en los resultados
- Si no hay resultados, se muestra un mensaje con sugerencias
- La busqueda funciona con Algolia para resultados rapidos y relevantes
- Se puede buscar en espanol e ingles

**Prioridad:** Media
**Sprint:** 11

---

### US-1.5: Ver detalle de una ruta
**Como** visitante o usuario registrado
**Quiero** ver la informacion completa de una ruta de aprendizaje
**Para** decidir si quiero iniciarla

**Criterios de Aceptacion:**
- Se muestra titulo, descripcion larga, categoria, dificultad, duracion estimada
- Se muestra la lista completa de hitos en orden
- Cada hito muestra titulo, descripcion corta, duracion estimada y si tiene evaluacion
- Se muestra el arbol de habilidades visual (skill tree)
- Se muestra un boton "Iniciar esta ruta" (lleva a registro si no esta autenticado)
- Se muestran estadisticas: numero de estudiantes, calificacion promedio
- La URL usa slug amigable (/paths/backend-developer-java)
- Los meta tags de SEO se generan dinamicamente

**Prioridad:** Alta
**Sprint:** 4

---

### US-1.6: Ver rutas destacadas
**Como** visitante que llega a la landing page
**Quiero** ver las rutas de aprendizaje mas populares o recomendadas
**Para** tener un punto de partida si no se que elegir

**Criterios de Aceptacion:**
- La landing page muestra una seccion "Rutas Destacadas" con 3-4 rutas
- Las rutas destacadas se seleccionan manualmente por el admin (flag is_featured)
- Cada tarjeta destacada muestra un diseno visual mas prominente que el catalogo regular
- Al hacer clic se navega al detalle de la ruta

**Prioridad:** Media
**Sprint:** 10

---

## Epic 2: Jornada de Aprendizaje

### US-2.1: Iniciar una ruta de aprendizaje
**Como** usuario registrado
**Quiero** iniciar una ruta de aprendizaje
**Para** comenzar a seguir un camino estructurado hacia mi objetivo profesional

**Criterios de Aceptacion:**
- Al hacer clic en "Iniciar ruta", se crea un registro de progreso para el usuario
- El usuario es redirigido al dashboard con su nueva ruta activa
- Los hitos sin dependencias se marcan como "disponibles" inmediatamente
- El arbol de habilidades se actualiza mostrando los hitos disponibles
- Si el usuario es Free y ya tiene una ruta activa, se muestra mensaje de upgrade a Pro
- Se envia un email de bienvenida a la ruta con consejos para empezar
- La fecha de inicio se registra para estadisticas

**Prioridad:** Alta
**Sprint:** 6

---

### US-2.2: Ver mi dashboard de aprendizaje
**Como** usuario registrado con rutas activas
**Quiero** ver un resumen de mi progreso en todas mis rutas
**Para** saber en que punto estoy y que hacer a continuacion

**Criterios de Aceptacion:**
- Se muestran todas las rutas activas del usuario
- Cada ruta muestra barra de progreso con porcentaje de completacion
- Se muestra el ultimo hito completado y la fecha
- Se muestra el siguiente hito sugerido (proximo disponible)
- Se muestran estadisticas generales: total hitos completados, dias activo, evaluaciones aprobadas
- Si el usuario no tiene rutas activas, se muestra CTA para explorar el catalogo
- La pagina se carga en menos de 2 segundos

**Prioridad:** Alta
**Sprint:** 6

---

### US-2.3: Navegar el arbol de habilidades interactivo
**Como** usuario con una ruta activa
**Quiero** ver mi progreso en el arbol de habilidades interactivo
**Para** visualizar que he aprendido, donde estoy y que sigue

**Criterios de Aceptacion:**
- Los hitos se muestran como nodos conectados en un grafo
- Los nodos tienen 4 estados visuales distintos: bloqueado (gris), disponible (azul), en progreso (amarillo), completado (verde)
- Las conexiones muestran dependencias entre hitos
- Se puede hacer zoom in/out y pan (arrastrar) el arbol
- Al hacer hover sobre un nodo se muestra tooltip con titulo, descripcion y duracion
- Al hacer clic en un nodo disponible se navega al detalle del hito
- Los nodos bloqueados muestran icono de candado y no son clickeables
- En mobile se muestra una vista simplificada (lista lineal con indicadores de estado)
- Al completar un hito, los nodos que se desbloquean tienen una animacion de transicion

**Prioridad:** Alta
**Sprint:** 5

---

### US-2.4: Marcar un hito como completado
**Como** usuario con una ruta activa
**Quiero** marcar un hito como completado cuando termine de estudiarlo
**Para** registrar mi progreso y desbloquear los siguientes hitos

**Criterios de Aceptacion:**
- Existe un boton "Marcar como completado" en la pagina del hito
- El sistema valida que todos los prerrequisitos estan completados
- Si el hito tiene evaluacion, el boton se habilita solo si la evaluacion fue aprobada
- Al completar, la barra de progreso de la ruta se actualiza inmediatamente
- Los hitos que dependian de este se desbloquean automaticamente
- Se muestra una animacion o mensaje de felicitacion
- El arbol de habilidades se actualiza en tiempo real
- Se registra la fecha y hora de completacion
- Si la ruta llega al 100%, se muestra mensaje de completacion de ruta

**Prioridad:** Alta
**Sprint:** 6

---

### US-2.5: Pausar una ruta de aprendizaje
**Como** usuario con una ruta activa
**Quiero** pausar mi progreso en una ruta
**Para** poder enfocarme en otra cosa sin perder mi avance

**Criterios de Aceptacion:**
- Existe una opcion "Pausar ruta" en el dashboard o detalle de la ruta
- Al pausar, el estado cambia a "paused" y se registra la fecha
- La ruta pausada aparece en una seccion separada del dashboard ("Rutas Pausadas")
- El progreso se conserva intacto
- Se puede reanudar en cualquier momento con un boton "Reanudar"
- Al reanudar, el estado vuelve a "active" y la fecha se actualiza
- Las rutas pausadas no cuentan contra el limite de rutas activas (para usuarios Free)

**Prioridad:** Media
**Sprint:** 6

---

### US-2.6: Abandonar una ruta
**Como** usuario con una ruta activa
**Quiero** abandonar una ruta que ya no me interesa
**Para** liberar mi cupo de ruta activa (si soy usuario Free)

**Criterios de Aceptacion:**
- Existe una opcion "Abandonar ruta" en la configuracion de la ruta
- Se muestra una confirmacion antes de abandonar ("Perderas tu progreso en esta ruta")
- Al abandonar, el estado cambia a "abandoned"
- El usuario puede reiniciar la misma ruta desde cero en el futuro
- Las estadisticas de la ruta abandonada se conservan para analytics internos
- No se elimina de la base de datos, solo cambia de estado

**Prioridad:** Baja
**Sprint:** 6

---

### US-2.7: Recibir sugerencia del siguiente paso
**Como** usuario con una ruta activa
**Quiero** que la plataforma me sugiera que recurso o hito abordar a continuacion
**Para** no perder tiempo decidiendo y mantener el momentum

**Criterios de Aceptacion:**
- En el dashboard se muestra un banner "Tu siguiente paso" con el recurso recomendado
- La recomendacion prioriza: hitos disponibles con mayor relevancia, luego recursos mejor calificados
- Se puede descartar la sugerencia y ver la siguiente
- La sugerencia incluye tipo de recurso, duracion estimada y calificacion
- Al hacer clic se navega directamente al recurso o hito

**Prioridad:** Media
**Sprint:** 11

---

## Epic 3: Recursos

### US-3.1: Ver recursos curados de un hito
**Como** usuario que esta estudiando un hito
**Quiero** ver los recursos de aprendizaje curados para ese hito
**Para** acceder a materiales de calidad sin perder tiempo buscando

**Criterios de Aceptacion:**
- Los recursos se muestran en la pagina del hito, ordenados por relevancia (recomendados primero, luego por calificacion)
- Cada recurso muestra: titulo, tipo (icono de video/articulo/curso), proveedor, duracion, calificacion promedio, etiqueta gratuito/de pago
- Los recursos gratuitos se muestran antes que los de pago
- El recurso recomendado tiene un badge visual "Recomendado"
- Al hacer clic en un recurso se abre en nueva pestana (link externo)
- Se muestra un contador de cuantos estudiantes usaron cada recurso

**Prioridad:** Alta
**Sprint:** 3-4

---

### US-3.2: Calificar un recurso
**Como** usuario que consumio un recurso
**Quiero** calificarlo con estrellas (1-5)
**Para** ayudar a otros estudiantes a elegir los mejores recursos

**Criterios de Aceptacion:**
- Se muestra un componente de 5 estrellas debajo del recurso
- El usuario puede seleccionar de 1 a 5 estrellas
- Se puede cambiar la calificacion despues de enviarla
- La calificacion promedio se actualiza inmediatamente (optimistic update)
- Solo usuarios autenticados pueden calificar
- Un usuario puede calificar cada recurso solo una vez (puede actualizar)

**Prioridad:** Media
**Sprint:** 8

---

### US-3.3: Escribir una resena de un recurso
**Como** usuario que consumio un recurso
**Quiero** escribir una resena explicando mi opinion
**Para** dar feedback detallado que ayude a otros estudiantes

**Criterios de Aceptacion:**
- Existe un area de texto para escribir la resena (maximo 1000 caracteres)
- La resena se asocia a la calificacion (no se puede resenar sin calificar)
- Se muestra la lista de resenas debajo del recurso, mas recientes primero
- Cada resena muestra avatar del usuario, nombre, calificacion y fecha
- Se puede editar la resena propia
- El contenido se sanitiza para prevenir XSS

**Prioridad:** Baja
**Sprint:** 8

---

### US-3.4: Sugerir un nuevo recurso
**Como** usuario que encontro un buen recurso no incluido en la plataforma
**Quiero** sugerirlo para que sea agregado a un hito
**Para** contribuir a la comunidad y mejorar el contenido

**Criterios de Aceptacion:**
- Existe un boton "Sugerir recurso" en la pagina del hito
- El formulario pide: URL, titulo, tipo, idioma, si es gratuito
- La sugerencia se envia para revision del admin (no se publica automaticamente)
- El usuario recibe notificacion cuando su sugerencia es aprobada o rechazada
- Solo usuarios Pro pueden sugerir recursos (incentivo para upgrade)
- Se valida que la URL no esta ya incluida en el hito

**Prioridad:** Baja
**Sprint:** Fase 2

---

### US-3.5: Guardar un recurso como bookmark
**Como** usuario que quiere volver a un recurso mas tarde
**Quiero** guardarlo en mis bookmarks
**Para** acceder rapidamente sin buscarlo de nuevo

**Criterios de Aceptacion:**
- Existe un icono de bookmark en cada tarjeta de recurso
- Al hacer clic, el recurso se agrega a la lista de bookmarks (toggle)
- El icono cambia de estado (lleno/vacio) al agregar/quitar
- Los bookmarks se guardan por usuario, persistentes entre sesiones
- Existe una pagina /bookmarks con todos los recursos guardados
- En la pagina de bookmarks se puede filtrar por tipo y ordenar por fecha
- Se puede quitar un bookmark desde la pagina de bookmarks

**Prioridad:** Media
**Sprint:** 8

---

## Epic 4: Evaluacion de Habilidades

### US-4.1: Tomar un quiz de evaluacion
**Como** usuario que completo los recursos de un hito
**Quiero** tomar un quiz para evaluar mi comprension
**Para** validar que realmente aprendi el tema antes de avanzar

**Criterios de Aceptacion:**
- Al acceder a la evaluacion se muestran las instrucciones: numero de preguntas, tiempo limite (si aplica), score minimo para aprobar
- Las preguntas se muestran una a la vez
- Para preguntas de opcion unica: radio buttons
- Para preguntas de opcion multiple: checkboxes
- Para verdadero/falso: dos botones
- Se puede navegar entre preguntas (siguiente/anterior)
- Se muestra indicador de progreso (pregunta 3 de 10)
- Si hay tiempo limite, se muestra un contador regresivo
- Existe un boton "Enviar evaluacion" que requiere confirmacion
- No se pueden ver las respuestas correctas antes de enviar
- El orden de las preguntas y opciones se aleatoriza

**Prioridad:** Alta
**Sprint:** 7

---

### US-4.2: Ver resultados de una evaluacion
**Como** usuario que termino un quiz
**Quiero** ver mis resultados detallados
**Para** saber en que acerte, en que falle y aprender de mis errores

**Criterios de Aceptacion:**
- Se muestra el score total (porcentaje)
- Se muestra si aprobo o no (score >= 70%)
- Se muestra cada pregunta con:
  - La respuesta que selecciono el usuario
  - La respuesta correcta (resaltada en verde)
  - Si la respuesta del usuario fue correcta o incorrecta
  - Explicacion de la respuesta correcta (solo para usuarios Pro)
- Si no aprobo: mensaje alentador + boton "Reintentar" (solo Pro, max 3 intentos)
- Si aprobo: mensaje de felicitacion + boton "Completar hito"
- Se muestra el numero de intento (1/3, 2/3, etc.)
- Se registra el tiempo que tardo en completar el quiz

**Prioridad:** Alta
**Sprint:** 7

---

### US-4.3: Reintentar una evaluacion fallida
**Como** usuario Pro que no aprobo un quiz
**Quiero** volver a intentar la evaluacion
**Para** tener otra oportunidad de demostrar que aprendi

**Criterios de Aceptacion:**
- Solo usuarios Pro pueden reintentar (Free users ven mensaje de upgrade)
- Maximo 3 intentos por evaluacion
- El orden de preguntas y opciones se aleatoriza en cada intento
- Se muestra cuantos intentos quedan (ej: "Intento 2 de 3")
- El mejor score se conserva (no se sobreescribe con uno peor)
- Si se agotan los intentos sin aprobar, se sugiere revisar los recursos del hito
- Hay un cooldown de 30 minutos entre intentos (para fomentar estudio)

**Prioridad:** Media
**Sprint:** 7

---

### US-4.4: Ver grafico radar de habilidades
**Como** usuario registrado
**Quiero** ver un grafico radar que muestre mis habilidades por area
**Para** tener una vision clara de mis fortalezas y areas de mejora

**Criterios de Aceptacion:**
- Se muestra un radar chart en el perfil del usuario
- Los ejes del radar representan categorias de habilidades (Backend, Frontend, Data, Mobile, DevOps)
- El valor de cada eje se calcula como: (hitos completados en esa categoria / total hitos en rutas iniciadas de esa categoria) * 100
- Solo usuarios Pro pueden ver su radar completo (Free ven version limitada)
- El radar se actualiza automaticamente al completar hitos
- Se puede compartir el radar como imagen (para LinkedIn, etc.)
- En mobile se muestra una version simplificada (barras horizontales)

**Prioridad:** Media
**Sprint:** 9

---

### US-4.5: Ver historial de evaluaciones
**Como** usuario registrado
**Quiero** ver el historial de todas mis evaluaciones
**Para** rastrear mi rendimiento a lo largo del tiempo

**Criterios de Aceptacion:**
- Se muestra una lista de todas las evaluaciones tomadas, ordenadas por fecha
- Cada entrada muestra: nombre del hito, score, aprobado/reprobado, fecha, numero de intento
- Se puede filtrar por ruta de aprendizaje
- Se puede filtrar por estado (aprobadas/reprobadas/todas)
- Al hacer clic en una entrada se ven los resultados detallados
- Se muestra un promedio general de todas las evaluaciones

**Prioridad:** Baja
**Sprint:** 9

---

## Epic 5: Comunidad

### US-5.1: Ver discusiones de un hito
**Como** usuario que esta estudiando un hito
**Quiero** ver las discusiones existentes sobre ese tema
**Para** aprender de las preguntas y respuestas de otros estudiantes

**Criterios de Aceptacion:**
- Se muestra una seccion "Discusiones" en la pagina del hito
- Las discusiones se muestran en una lista con titulo, autor, fecha, numero de respuestas y votos
- Se puede ordenar por: mas recientes, mas votadas, sin responder
- Las discusiones marcadas como "pregunta" tienen un icono especial
- Las preguntas resueltas tienen un check verde
- Los usuarios Free pueden leer pero no participar (mensaje de upgrade)

**Prioridad:** Media
**Sprint:** Fase 2

---

### US-5.2: Crear una discusion
**Como** usuario Pro que tiene una duda o quiere compartir algo
**Quiero** crear una nueva discusion en un hito
**Para** obtener ayuda de la comunidad o contribuir con mi conocimiento

**Criterios de Aceptacion:**
- Existe un boton "Nueva Discusion" en la seccion de discusiones
- El formulario pide: titulo (obligatorio), cuerpo (obligatorio, con formato Markdown basico), checkbox "Es una pregunta"
- El titulo debe tener entre 10 y 300 caracteres
- El cuerpo debe tener entre 20 y 5000 caracteres
- Solo usuarios Pro pueden crear discusiones
- Limite de 20 discusiones por dia (prevencion de spam)
- Al crear, se redirige a la pagina de la discusion
- El contenido se sanitiza contra XSS

**Prioridad:** Media
**Sprint:** Fase 2

---

### US-5.3: Responder a una discusion
**Como** usuario Pro que quiere ayudar a otro estudiante
**Quiero** responder a una discusion existente
**Para** compartir mi conocimiento o resolver la duda de alguien

**Criterios de Aceptacion:**
- Existe un area de texto para responder debajo de la discusion
- Las respuestas soportan formato Markdown basico (negrita, codigo, listas)
- Solo usuarios Pro pueden responder
- Las respuestas se muestran en orden cronologico
- El autor de la discusion (si fue marcada como pregunta) puede marcar una respuesta como "aceptada"
- La respuesta aceptada se muestra primero con un indicador visual
- Se envia notificacion al autor de la discusion cuando alguien responde

**Prioridad:** Media
**Sprint:** Fase 2

---

### US-5.4: Votar positivamente una discusion o respuesta
**Como** usuario Pro
**Quiero** votar positivamente discusiones o respuestas utiles
**Para** destacar el contenido de mayor calidad

**Criterios de Aceptacion:**
- Existe un boton de upvote (flecha arriba) en cada discusion y respuesta
- Un usuario puede votar una vez por item (toggle: clic para votar, clic de nuevo para quitar voto)
- El conteo de votos se actualiza inmediatamente (optimistic update)
- Las discusiones se pueden ordenar por numero de votos
- No existe downvote (solo upvote) para mantener la comunidad positiva
- Solo usuarios Pro pueden votar

**Prioridad:** Baja
**Sprint:** Fase 2

---

### US-5.5: Hacer una pregunta en un hito
**Como** usuario Pro que esta atascado en un tema
**Quiero** hacer una pregunta especifica sobre el hito actual
**Para** obtener ayuda de la comunidad y desbloquear mi aprendizaje

**Criterios de Aceptacion:**
- Al crear una discusion y marcar "Es una pregunta", se muestra con icono de interrogacion
- Las preguntas sin responder aparecen primero en la lista
- El autor puede marcar la pregunta como "resuelta" cuando obtiene una respuesta satisfactoria
- Las preguntas resueltas se muestran con un check verde
- Se puede filtrar para ver solo preguntas (vs discusiones generales)
- Se muestra un contador de preguntas sin responder por hito

**Prioridad:** Media
**Sprint:** Fase 2

---

## Epic 6: Perfil y Portafolio

### US-6.1: Ver mi perfil publico
**Como** usuario registrado
**Quiero** ver como se ve mi perfil publico
**Para** saber que informacion otros pueden ver de mi

**Criterios de Aceptacion:**
- Se muestra: avatar, nombre, bio, links (GitHub, LinkedIn, web personal)
- Se muestran rutas completadas con fecha de completacion
- Se muestran rutas activas con porcentaje de progreso
- Se muestra el grafico radar de habilidades (Pro)
- Se muestran estadisticas: hitos completados, evaluaciones aprobadas, dias activo
- Se muestra la seccion de proyectos de portafolio (si hay)
- La URL es amigable: /profile/username
- Se generan meta tags OG para compartir en redes sociales

**Prioridad:** Media
**Sprint:** 9

---

### US-6.2: Editar mi perfil
**Como** usuario registrado
**Quiero** editar mi informacion de perfil
**Para** mantener mi informacion actualizada y personalizar mi presencia

**Criterios de Aceptacion:**
- Se pueden editar: nombre para mostrar, bio (max 500 caracteres), avatar (upload o URL), links sociales (GitHub, LinkedIn, web)
- Se puede cambiar el username (unico, alfanumerico + guiones, 3-30 caracteres)
- Los cambios se guardan al hacer clic en "Guardar cambios"
- Se muestra confirmacion visual de que los cambios fueron guardados
- Se valida formato de URLs de redes sociales
- El avatar se redimensiona automaticamente a 200x200px

**Prioridad:** Media
**Sprint:** 9

---

### US-6.3: Compartir mi perfil
**Como** usuario registrado que quiere mostrar sus habilidades
**Quiero** compartir mi perfil de SkillPath
**Para** demostrar mi progreso y habilidades a empleadores o colegas

**Criterios de Aceptacion:**
- Existe un boton "Compartir perfil" que copia la URL al portapapeles
- La URL es publica y accesible sin autenticacion
- Al compartir en redes sociales, se muestra una preview con OG tags (imagen con avatar, nombre, estadisticas)
- Se puede generar una imagen del grafico radar para descargar
- Se incluye un boton para compartir directamente en LinkedIn y Twitter

**Prioridad:** Baja
**Sprint:** 9

---

### US-6.4: Ver sugerencias de proyectos de portafolio
**Como** usuario Pro que esta avanzando en una ruta
**Quiero** ver proyectos sugeridos que puedo construir
**Para** aplicar lo aprendido y tener proyectos concretos para mostrar a empleadores

**Criterios de Aceptacion:**
- Cada ruta incluye 2-3 proyectos sugeridos alineados con los hitos
- Cada proyecto muestra: titulo, descripcion, nivel de dificultad, habilidades que ejercita, tiempo estimado
- Los proyectos se desbloquean progresivamente al completar hitos
- Se pueden agregar links a los proyectos completados (URL de GitHub o demo)
- Los proyectos aparecen en el perfil publico
- Solo usuarios Pro pueden ver las sugerencias de proyectos

**Prioridad:** Baja
**Sprint:** Fase 2

---

### US-6.5: Agregar proyecto al portafolio
**Como** usuario Pro que termino un proyecto
**Quiero** agregarlo a mi perfil
**Para** mostrar mi trabajo practico junto a mi progreso de aprendizaje

**Criterios de Aceptacion:**
- Se puede agregar un proyecto con: titulo, descripcion, URL de repositorio, URL de demo, imagen de captura
- Los proyectos se muestran en una seccion dedicada del perfil
- Se puede editar o eliminar proyectos propios
- Maximo 10 proyectos en el portafolio
- Los proyectos pueden estar asociados a una ruta o ser independientes

**Prioridad:** Baja
**Sprint:** Fase 2

---

### US-6.6: Ver grafico radar de habilidades en el perfil
**Como** visitante que ve el perfil de otro usuario
**Quiero** ver su grafico radar de habilidades
**Para** entender sus fortalezas y areas de conocimiento

**Criterios de Aceptacion:**
- El radar chart se muestra en el perfil publico
- Los ejes muestran las categorias de habilidades con valores 0-100%
- El grafico es interactivo (tooltip al hover con valor exacto)
- Se puede comparar con el promedio de la plataforma (linea punteada)
- Solo se muestra si el usuario tiene al menos 1 ruta activa
- El diseno es visualmente atractivo y profesional

**Prioridad:** Baja
**Sprint:** 9

---

## Epic 7: Mentoria

### US-7.1: Solicitar un mentor
**Como** usuario Pro que necesita orientacion personalizada
**Quiero** solicitar un mentor en mi area de interes
**Para** recibir consejos de alguien con experiencia real en la industria

**Criterios de Aceptacion:**
- Existe una seccion "Mentoria" accesible desde el menu principal
- Se muestra un formulario con: area de interes (seleccionar de las categorias), nivel actual, objetivos especificos, disponibilidad horaria
- Solo usuarios Pro pueden solicitar mentoria
- Al enviar, el sistema busca mentores compatibles y muestra opciones
- Si no hay mentores disponibles, se coloca en lista de espera con notificacion
- La solicitud tiene un estado visible: pendiente, matched, activa, completada

**Prioridad:** Baja
**Sprint:** Fase 3

---

### US-7.2: Explorar mentores disponibles
**Como** usuario Pro
**Quiero** ver los mentores disponibles y sus perfiles
**Para** elegir al mentor que mejor se ajuste a mis necesidades

**Criterios de Aceptacion:**
- Se muestra una lista de mentores con: foto, nombre, especialidades, anos de experiencia, calificacion promedio, numero de sesiones realizadas
- Se puede filtrar por especialidad (categoria)
- Se puede ordenar por calificacion o numero de sesiones
- Al hacer clic se ve el perfil completo del mentor con bio, experiencia, reviews
- Cada mentor tiene un boton "Solicitar sesion"
- Se muestra disponibilidad del mentor (dias/horas)

**Prioridad:** Baja
**Sprint:** Fase 3

---

### US-7.3: Programar una sesion de mentoria
**Como** usuario Pro que fue emparejado con un mentor
**Quiero** programar una sesion de mentoria
**Para** tener una conversacion estructurada y productiva

**Criterios de Aceptacion:**
- Se muestra un calendario con la disponibilidad del mentor
- El usuario selecciona fecha y hora
- Se puede agregar un mensaje de contexto para la sesion (que quiere discutir)
- Las sesiones duran 30 o 60 minutos
- Se envia notificacion al mentor para confirmar
- Se genera un link de videoconferencia (integracion con Google Meet o Zoom)
- Se envian recordatorios por email 24h y 1h antes de la sesion
- Tanto mentor como mentee pueden cancelar con al menos 12h de anticipacion

**Prioridad:** Baja
**Sprint:** Fase 3

---

### US-7.4: Calificar una sesion de mentoria
**Como** usuario que tuvo una sesion de mentoria
**Quiero** calificar la sesion y al mentor
**Para** dar feedback que ayude a mejorar el servicio y guiar a futuros mentees

**Criterios de Aceptacion:**
- Despues de cada sesion, se muestra un formulario de calificacion
- Se califica con estrellas (1-5) en: conocimiento tecnico, comunicacion, utilidad practica
- Se puede escribir un comentario (opcional, max 500 caracteres)
- La calificacion es anonima por defecto (con opcion de mostrar nombre)
- El promedio de calificaciones se actualiza en el perfil del mentor
- Mentores con calificacion promedio < 3.0 despues de 10 sesiones son revisados

**Prioridad:** Baja
**Sprint:** Fase 3

---

### US-7.5: Registrarse como mentor
**Como** profesional con experiencia que quiere ayudar a otros
**Quiero** registrarme como mentor en la plataforma
**Para** compartir mi conocimiento y recibir compensacion

**Criterios de Aceptacion:**
- Existe un formulario de aplicacion para mentores
- Se solicita: anos de experiencia, areas de especialidad, links de LinkedIn/GitHub, bio profesional, disponibilidad, motivacion
- Las aplicaciones son revisadas manualmente por el equipo de SkillPath
- Se requiere minimo 3 anos de experiencia profesional
- Se notifica por email si la aplicacion fue aprobada o rechazada
- Los mentores aprobados reciben onboarding y guias de mejores practicas
- Los mentores reciben el 70% del valor de cada sesion (el 30% es para la plataforma)

**Prioridad:** Baja
**Sprint:** Fase 3

---

## Resumen de Prioridades por Sprint

| Sprint | Historias de Usuario |
|--------|---------------------|
| 3 | US-3.1 (parcial: backend) |
| 4 | US-1.1, US-1.2, US-1.3, US-1.5, US-3.1 (frontend) |
| 5 | US-2.3 |
| 6 | US-2.1, US-2.2, US-2.4, US-2.5, US-2.6 |
| 7 | US-4.1, US-4.2, US-4.3 |
| 8 | US-3.2, US-3.3, US-3.5 |
| 9 | US-4.4, US-4.5, US-6.1, US-6.2, US-6.3, US-6.6 |
| 10 | US-1.6 |
| 11 | US-1.4, US-2.7 |
| Fase 2 | US-3.4, US-5.1, US-5.2, US-5.3, US-5.4, US-5.5, US-6.4, US-6.5 |
| Fase 3 | US-7.1, US-7.2, US-7.3, US-7.4, US-7.5 |

---

## Conteo Total

| Epic | Historias | Alta | Media | Baja |
|------|----------|------|-------|------|
| Epic 1: Descubrimiento | 6 | 2 | 3 | 1 |
| Epic 2: Jornada de Aprendizaje | 7 | 3 | 3 | 1 |
| Epic 3: Recursos | 5 | 1 | 2 | 2 |
| Epic 4: Evaluacion | 5 | 2 | 1 | 2 |
| Epic 5: Comunidad | 5 | 0 | 3 | 2 |
| Epic 6: Perfil y Portafolio | 6 | 0 | 2 | 4 |
| Epic 7: Mentoria | 5 | 0 | 0 | 5 |
| **Total** | **39** | **8** | **14** | **17** |
