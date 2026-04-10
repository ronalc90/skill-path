package repository

import (
	"log"

	"github.com/ronalc90/skillpath/internal/model"
	"gorm.io/gorm"
)

// Seed populates the database with initial learning paths, milestones, resources, and assessments.
func Seed(db *gorm.DB) {
	var count int64
	db.Model(&model.LearningPath{}).Count(&count)
	if count > 0 {
		log.Println("Database already seeded, skipping...")
		return
	}

	log.Println("Seeding database...")

	paths := []model.LearningPath{
		backendJavaPath(),
		frontendReactPath(),
		dataAnalystPythonPath(),
		mobileFlutterPath(),
	}

	for i := range paths {
		if err := db.Create(&paths[i]).Error; err != nil {
			log.Printf("Error seeding path %s: %v", paths[i].Title, err)
		}
	}

	seedAssessments(db, paths)
	log.Println("Database seeded successfully!")
}

func backendJavaPath() model.LearningPath {
	return model.LearningPath{
		Title:          "Backend Developer con Java",
		Description:    "Domina el desarrollo backend con Java, Spring Boot y microservicios. Aprende desde los fundamentos de Java hasta arquitecturas distribuidas y despliegue en la nube.",
		Slug:           "backend-developer-java",
		Difficulty:     model.DifficultyIntermediate,
		EstimatedHours: 120,
		Category:       "Backend",
		IconURL:        "https://cdn.jsdelivr.net/gh/devicons/devicon/icons/java/java-original.svg",
		IsPublished:    true,
		Milestones: []model.Milestone{
			{
				Title:          "Fundamentos de Java",
				Description:    "Variables, tipos de datos, control de flujo, POO y colecciones.",
				OrderIndex:     1,
				EstimatedHours: 15,
				Resources: []model.Resource{
					{Title: "Java Programming Masterclass", URL: "https://www.udemy.com/course/java-the-complete-java-developer-course/", Type: model.ResourceTypeCourse, Provider: "Udemy", IsFree: false, EstimatedMinutes: 480, OrderIndex: 1},
					{Title: "Java Tutorial - W3Schools", URL: "https://www.w3schools.com/java/", Type: model.ResourceTypeTutorial, Provider: "W3Schools", IsFree: true, EstimatedMinutes: 120, OrderIndex: 2},
					{Title: "Effective Java - Joshua Bloch", URL: "https://www.oreilly.com/library/view/effective-java/9780134686097/", Type: model.ResourceTypeBook, Provider: "O'Reilly", IsFree: false, EstimatedMinutes: 600, OrderIndex: 3},
				},
			},
			{
				Title:          "Spring Boot Essentials",
				Description:    "Configuracion de proyectos, inyeccion de dependencias, REST APIs y testing.",
				OrderIndex:     2,
				EstimatedHours: 20,
				Resources: []model.Resource{
					{Title: "Spring Boot Reference Documentation", URL: "https://docs.spring.io/spring-boot/docs/current/reference/html/", Type: model.ResourceTypeArticle, Provider: "Spring.io", IsFree: true, EstimatedMinutes: 180, OrderIndex: 1},
					{Title: "Building REST APIs with Spring Boot", URL: "https://www.youtube.com/watch?v=9SGDpanrc8U", Type: model.ResourceTypeVideo, Provider: "YouTube", IsFree: true, EstimatedMinutes: 90, OrderIndex: 2},
					{Title: "Spring Boot REST API Project", URL: "https://github.com/spring-projects/spring-petclinic", Type: model.ResourceTypeExercise, Provider: "GitHub", IsFree: true, EstimatedMinutes: 240, OrderIndex: 3},
				},
			},
			{
				Title:          "Bases de Datos con JPA/Hibernate",
				Description:    "Mapeo objeto-relacional, consultas JPQL, transacciones y migraciones.",
				OrderIndex:     3,
				EstimatedHours: 15,
				Resources: []model.Resource{
					{Title: "JPA and Hibernate Tutorial", URL: "https://www.baeldung.com/learn-jpa-hibernate", Type: model.ResourceTypeTutorial, Provider: "Baeldung", IsFree: true, EstimatedMinutes: 120, OrderIndex: 1},
					{Title: "Database Design for Beginners", URL: "https://www.youtube.com/watch?v=ztHopE5Wnpc", Type: model.ResourceTypeVideo, Provider: "YouTube", IsFree: true, EstimatedMinutes: 120, OrderIndex: 2},
				},
			},
			{
				Title:          "Seguridad con Spring Security",
				Description:    "Autenticacion, autorizacion, JWT, OAuth2 y mejores practicas de seguridad.",
				OrderIndex:     4,
				EstimatedHours: 15,
				Resources: []model.Resource{
					{Title: "Spring Security in Action", URL: "https://www.manning.com/books/spring-security-in-action", Type: model.ResourceTypeBook, Provider: "Manning", IsFree: false, EstimatedMinutes: 480, OrderIndex: 1},
					{Title: "JWT Authentication Tutorial", URL: "https://www.baeldung.com/spring-security-oauth-jwt", Type: model.ResourceTypeTutorial, Provider: "Baeldung", IsFree: true, EstimatedMinutes: 90, OrderIndex: 2},
				},
			},
			{
				Title:          "Testing y Calidad de Codigo",
				Description:    "JUnit 5, Mockito, tests de integracion, cobertura y analisis estatico.",
				OrderIndex:     5,
				EstimatedHours: 12,
				Resources: []model.Resource{
					{Title: "Testing Spring Boot Applications", URL: "https://www.baeldung.com/spring-boot-testing", Type: model.ResourceTypeTutorial, Provider: "Baeldung", IsFree: true, EstimatedMinutes: 90, OrderIndex: 1},
					{Title: "JUnit 5 User Guide", URL: "https://junit.org/junit5/docs/current/user-guide/", Type: model.ResourceTypeArticle, Provider: "JUnit", IsFree: true, EstimatedMinutes: 120, OrderIndex: 2},
				},
			},
			{
				Title:          "Microservicios",
				Description:    "Patrones de microservicios, comunicacion entre servicios, service discovery y API Gateway.",
				OrderIndex:     6,
				EstimatedHours: 18,
				Resources: []model.Resource{
					{Title: "Microservices with Spring Cloud", URL: "https://spring.io/projects/spring-cloud", Type: model.ResourceTypeCourse, Provider: "Spring.io", IsFree: true, EstimatedMinutes: 240, OrderIndex: 1},
					{Title: "Building Microservices - Sam Newman", URL: "https://www.oreilly.com/library/view/building-microservices-2nd/9781492034018/", Type: model.ResourceTypeBook, Provider: "O'Reilly", IsFree: false, EstimatedMinutes: 600, OrderIndex: 2},
				},
			},
			{
				Title:          "Docker y Kubernetes",
				Description:    "Contenedores, imagenes, Docker Compose, orquestacion con Kubernetes.",
				OrderIndex:     7,
				EstimatedHours: 15,
				Resources: []model.Resource{
					{Title: "Docker for Java Developers", URL: "https://www.docker.com/blog/intro-guide-to-dockerfile-best-practices/", Type: model.ResourceTypeArticle, Provider: "Docker", IsFree: true, EstimatedMinutes: 60, OrderIndex: 1},
					{Title: "Kubernetes in Action", URL: "https://www.manning.com/books/kubernetes-in-action", Type: model.ResourceTypeBook, Provider: "Manning", IsFree: false, EstimatedMinutes: 480, OrderIndex: 2},
				},
			},
			{
				Title:          "CI/CD y Despliegue",
				Description:    "Pipelines de CI/CD, GitHub Actions, monitoreo y observabilidad.",
				OrderIndex:     8,
				EstimatedHours: 10,
				Resources: []model.Resource{
					{Title: "GitHub Actions for Java", URL: "https://docs.github.com/en/actions/automating-builds-and-tests/building-and-testing-java-with-maven", Type: model.ResourceTypeTutorial, Provider: "GitHub", IsFree: true, EstimatedMinutes: 90, OrderIndex: 1},
					{Title: "Deploying Spring Boot to AWS", URL: "https://www.youtube.com/watch?v=i7AHsN9C0Ao", Type: model.ResourceTypeVideo, Provider: "YouTube", IsFree: true, EstimatedMinutes: 60, OrderIndex: 2},
				},
			},
		},
	}
}

func frontendReactPath() model.LearningPath {
	return model.LearningPath{
		Title:          "Frontend Developer con React",
		Description:    "Conviertete en un desarrollador frontend profesional con React, TypeScript y el ecosistema moderno. Desde HTML/CSS hasta aplicaciones complejas con estado global.",
		Slug:           "frontend-developer-react",
		Difficulty:     model.DifficultyBeginner,
		EstimatedHours: 100,
		Category:       "Frontend",
		IconURL:        "https://cdn.jsdelivr.net/gh/devicons/devicon/icons/react/react-original.svg",
		IsPublished:    true,
		Milestones: []model.Milestone{
			{
				Title:          "HTML, CSS y JavaScript Moderno",
				Description:    "Fundamentos web: semantica HTML5, Flexbox, Grid, ES6+ y DOM.",
				OrderIndex:     1,
				EstimatedHours: 20,
				Resources: []model.Resource{
					{Title: "freeCodeCamp Responsive Web Design", URL: "https://www.freecodecamp.org/learn/2022/responsive-web-design/", Type: model.ResourceTypeCourse, Provider: "freeCodeCamp", IsFree: true, EstimatedMinutes: 300, OrderIndex: 1},
					{Title: "JavaScript.info", URL: "https://javascript.info/", Type: model.ResourceTypeTutorial, Provider: "javascript.info", IsFree: true, EstimatedMinutes: 480, OrderIndex: 2},
				},
			},
			{
				Title:          "TypeScript Fundamentals",
				Description:    "Tipos, interfaces, genericos, decoradores y configuracion de proyectos.",
				OrderIndex:     2,
				EstimatedHours: 12,
				Resources: []model.Resource{
					{Title: "TypeScript Handbook", URL: "https://www.typescriptlang.org/docs/handbook/intro.html", Type: model.ResourceTypeArticle, Provider: "TypeScript", IsFree: true, EstimatedMinutes: 180, OrderIndex: 1},
					{Title: "TypeScript Deep Dive", URL: "https://basarat.gitbook.io/typescript/", Type: model.ResourceTypeBook, Provider: "GitBook", IsFree: true, EstimatedMinutes: 360, OrderIndex: 2},
				},
			},
			{
				Title:          "React Core Concepts",
				Description:    "Components, JSX, props, state, hooks y ciclo de vida.",
				OrderIndex:     3,
				EstimatedHours: 18,
				Resources: []model.Resource{
					{Title: "React Official Tutorial", URL: "https://react.dev/learn", Type: model.ResourceTypeTutorial, Provider: "React", IsFree: true, EstimatedMinutes: 240, OrderIndex: 1},
					{Title: "React with TypeScript", URL: "https://www.youtube.com/watch?v=FJDVKeh7RJI", Type: model.ResourceTypeVideo, Provider: "YouTube", IsFree: true, EstimatedMinutes: 120, OrderIndex: 2},
				},
			},
			{
				Title:          "Estado Global y Data Fetching",
				Description:    "Context API, Zustand/Redux, React Query y patrones de estado.",
				OrderIndex:     4,
				EstimatedHours: 15,
				Resources: []model.Resource{
					{Title: "React Query Documentation", URL: "https://tanstack.com/query/latest", Type: model.ResourceTypeArticle, Provider: "TanStack", IsFree: true, EstimatedMinutes: 120, OrderIndex: 1},
					{Title: "Zustand State Management", URL: "https://github.com/pmndrs/zustand", Type: model.ResourceTypeTutorial, Provider: "GitHub", IsFree: true, EstimatedMinutes: 60, OrderIndex: 2},
				},
			},
			{
				Title:          "Next.js y SSR",
				Description:    "App Router, Server Components, SSR, SSG, API Routes e ISR.",
				OrderIndex:     5,
				EstimatedHours: 15,
				Resources: []model.Resource{
					{Title: "Next.js Documentation", URL: "https://nextjs.org/docs", Type: model.ResourceTypeArticle, Provider: "Vercel", IsFree: true, EstimatedMinutes: 180, OrderIndex: 1},
					{Title: "Next.js 14 Full Course", URL: "https://www.youtube.com/watch?v=wm5gMKuwSYk", Type: model.ResourceTypeVideo, Provider: "YouTube", IsFree: true, EstimatedMinutes: 240, OrderIndex: 2},
				},
			},
			{
				Title:          "Testing Frontend",
				Description:    "Jest, React Testing Library, Cypress, testing de componentes y E2E.",
				OrderIndex:     6,
				EstimatedHours: 10,
				Resources: []model.Resource{
					{Title: "Testing Library Documentation", URL: "https://testing-library.com/docs/react-testing-library/intro/", Type: model.ResourceTypeArticle, Provider: "Testing Library", IsFree: true, EstimatedMinutes: 90, OrderIndex: 1},
					{Title: "Cypress E2E Testing", URL: "https://www.cypress.io/", Type: model.ResourceTypeTutorial, Provider: "Cypress", IsFree: true, EstimatedMinutes: 120, OrderIndex: 2},
				},
			},
			{
				Title:          "Performance y Deploy",
				Description:    "Optimizacion, lazy loading, code splitting, Vercel y CI/CD.",
				OrderIndex:     7,
				EstimatedHours: 10,
				Resources: []model.Resource{
					{Title: "Web Vitals", URL: "https://web.dev/vitals/", Type: model.ResourceTypeArticle, Provider: "Google", IsFree: true, EstimatedMinutes: 60, OrderIndex: 1},
					{Title: "Deploy Next.js to Vercel", URL: "https://vercel.com/docs/frameworks/nextjs", Type: model.ResourceTypeTutorial, Provider: "Vercel", IsFree: true, EstimatedMinutes: 30, OrderIndex: 2},
				},
			},
		},
	}
}

func dataAnalystPythonPath() model.LearningPath {
	return model.LearningPath{
		Title:          "Data Analyst con Python",
		Description:    "Aprende analisis de datos con Python, desde los fundamentos de programacion hasta visualizacion avanzada y machine learning basico.",
		Slug:           "data-analyst-python",
		Difficulty:     model.DifficultyBeginner,
		EstimatedHours: 90,
		Category:       "Data",
		IconURL:        "https://cdn.jsdelivr.net/gh/devicons/devicon/icons/python/python-original.svg",
		IsPublished:    true,
		Milestones: []model.Milestone{
			{
				Title:          "Python Fundamentals",
				Description:    "Sintaxis, estructuras de datos, funciones, modulos y manejo de archivos.",
				OrderIndex:     1,
				EstimatedHours: 18,
				Resources: []model.Resource{
					{Title: "Python for Everybody", URL: "https://www.py4e.com/", Type: model.ResourceTypeCourse, Provider: "py4e", IsFree: true, EstimatedMinutes: 480, OrderIndex: 1},
					{Title: "Automate the Boring Stuff with Python", URL: "https://automatetheboringstuff.com/", Type: model.ResourceTypeBook, Provider: "No Starch Press", IsFree: true, EstimatedMinutes: 600, OrderIndex: 2},
				},
			},
			{
				Title:          "NumPy y Pandas",
				Description:    "Arrays, DataFrames, operaciones vectorizadas, limpieza y transformacion de datos.",
				OrderIndex:     2,
				EstimatedHours: 18,
				Resources: []model.Resource{
					{Title: "Pandas Documentation", URL: "https://pandas.pydata.org/docs/getting_started/index.html", Type: model.ResourceTypeArticle, Provider: "Pandas", IsFree: true, EstimatedMinutes: 120, OrderIndex: 1},
					{Title: "NumPy Tutorial", URL: "https://numpy.org/doc/stable/user/absolute_beginners.html", Type: model.ResourceTypeTutorial, Provider: "NumPy", IsFree: true, EstimatedMinutes: 90, OrderIndex: 2},
				},
			},
			{
				Title:          "Visualizacion de Datos",
				Description:    "Matplotlib, Seaborn, Plotly y principios de visualizacion efectiva.",
				OrderIndex:     3,
				EstimatedHours: 15,
				Resources: []model.Resource{
					{Title: "Matplotlib Tutorial", URL: "https://matplotlib.org/stable/tutorials/index.html", Type: model.ResourceTypeTutorial, Provider: "Matplotlib", IsFree: true, EstimatedMinutes: 120, OrderIndex: 1},
					{Title: "Data Visualization with Python", URL: "https://www.youtube.com/watch?v=UO98lJQ3QGI", Type: model.ResourceTypeVideo, Provider: "YouTube", IsFree: true, EstimatedMinutes: 180, OrderIndex: 2},
				},
			},
			{
				Title:          "SQL para Analistas",
				Description:    "Consultas, joins, agregaciones, subqueries y optimizacion.",
				OrderIndex:     4,
				EstimatedHours: 12,
				Resources: []model.Resource{
					{Title: "SQLBolt Interactive Tutorial", URL: "https://sqlbolt.com/", Type: model.ResourceTypeTutorial, Provider: "SQLBolt", IsFree: true, EstimatedMinutes: 120, OrderIndex: 1},
					{Title: "SQL for Data Analysis", URL: "https://mode.com/sql-tutorial/", Type: model.ResourceTypeTutorial, Provider: "Mode", IsFree: true, EstimatedMinutes: 180, OrderIndex: 2},
				},
			},
			{
				Title:          "Estadistica Aplicada",
				Description:    "Estadistica descriptiva, probabilidad, tests de hipotesis y regresion.",
				OrderIndex:     5,
				EstimatedHours: 15,
				Resources: []model.Resource{
					{Title: "Statistics with Python", URL: "https://www.coursera.org/specializations/statistics-with-python", Type: model.ResourceTypeCourse, Provider: "Coursera", IsFree: false, EstimatedMinutes: 480, OrderIndex: 1},
					{Title: "Think Stats", URL: "https://greenteapress.com/thinkstats2/html/index.html", Type: model.ResourceTypeBook, Provider: "Green Tea Press", IsFree: true, EstimatedMinutes: 360, OrderIndex: 2},
				},
			},
			{
				Title:          "Machine Learning Basico",
				Description:    "Scikit-learn, clasificacion, regresion, clustering y evaluacion de modelos.",
				OrderIndex:     6,
				EstimatedHours: 12,
				Resources: []model.Resource{
					{Title: "Scikit-learn Tutorial", URL: "https://scikit-learn.org/stable/tutorial/index.html", Type: model.ResourceTypeTutorial, Provider: "scikit-learn", IsFree: true, EstimatedMinutes: 180, OrderIndex: 1},
					{Title: "Machine Learning Crash Course", URL: "https://developers.google.com/machine-learning/crash-course", Type: model.ResourceTypeCourse, Provider: "Google", IsFree: true, EstimatedMinutes: 300, OrderIndex: 2},
				},
			},
		},
	}
}

func mobileFlutterPath() model.LearningPath {
	return model.LearningPath{
		Title:          "Mobile Developer con Flutter",
		Description:    "Desarrolla aplicaciones moviles multiplataforma con Flutter y Dart. Desde los fundamentos hasta la publicacion en App Store y Google Play.",
		Slug:           "mobile-developer-flutter",
		Difficulty:     model.DifficultyIntermediate,
		EstimatedHours: 85,
		Category:       "Mobile",
		IconURL:        "https://cdn.jsdelivr.net/gh/devicons/devicon/icons/flutter/flutter-original.svg",
		IsPublished:    true,
		Milestones: []model.Milestone{
			{
				Title:          "Dart Programming Language",
				Description:    "Sintaxis, tipos, clases, async/await, null safety y colecciones.",
				OrderIndex:     1,
				EstimatedHours: 12,
				Resources: []model.Resource{
					{Title: "Dart Language Tour", URL: "https://dart.dev/language", Type: model.ResourceTypeTutorial, Provider: "Dart", IsFree: true, EstimatedMinutes: 180, OrderIndex: 1},
					{Title: "Dart Programming Tutorial", URL: "https://www.youtube.com/watch?v=Ej_Pcr4uC2Q", Type: model.ResourceTypeVideo, Provider: "YouTube", IsFree: true, EstimatedMinutes: 240, OrderIndex: 2},
				},
			},
			{
				Title:          "Flutter UI Fundamentals",
				Description:    "Widgets, layouts, Material Design, navegacion y temas.",
				OrderIndex:     2,
				EstimatedHours: 18,
				Resources: []model.Resource{
					{Title: "Flutter Official Codelabs", URL: "https://docs.flutter.dev/codelabs", Type: model.ResourceTypeTutorial, Provider: "Flutter", IsFree: true, EstimatedMinutes: 240, OrderIndex: 1},
					{Title: "Flutter Widget Catalog", URL: "https://docs.flutter.dev/ui/widgets", Type: model.ResourceTypeArticle, Provider: "Flutter", IsFree: true, EstimatedMinutes: 120, OrderIndex: 2},
				},
			},
			{
				Title:          "State Management",
				Description:    "setState, Provider, Riverpod, BLoC y patrones de arquitectura.",
				OrderIndex:     3,
				EstimatedHours: 15,
				Resources: []model.Resource{
					{Title: "Flutter State Management Guide", URL: "https://docs.flutter.dev/data-and-backend/state-mgmt/intro", Type: model.ResourceTypeArticle, Provider: "Flutter", IsFree: true, EstimatedMinutes: 90, OrderIndex: 1},
					{Title: "Riverpod Documentation", URL: "https://riverpod.dev/", Type: model.ResourceTypeTutorial, Provider: "Riverpod", IsFree: true, EstimatedMinutes: 120, OrderIndex: 2},
				},
			},
			{
				Title:          "Networking y APIs",
				Description:    "HTTP requests, REST APIs, serialization JSON, manejo de errores.",
				OrderIndex:     4,
				EstimatedHours: 12,
				Resources: []model.Resource{
					{Title: "Flutter Networking Tutorial", URL: "https://docs.flutter.dev/data-and-backend/networking", Type: model.ResourceTypeTutorial, Provider: "Flutter", IsFree: true, EstimatedMinutes: 90, OrderIndex: 1},
					{Title: "Dio HTTP Client", URL: "https://pub.dev/packages/dio", Type: model.ResourceTypeArticle, Provider: "pub.dev", IsFree: true, EstimatedMinutes: 60, OrderIndex: 2},
				},
			},
			{
				Title:          "Persistencia y Base de Datos",
				Description:    "SharedPreferences, SQLite, Hive, Firebase Firestore.",
				OrderIndex:     5,
				EstimatedHours: 15,
				Resources: []model.Resource{
					{Title: "Flutter Firebase Tutorial", URL: "https://firebase.google.com/docs/flutter/setup", Type: model.ResourceTypeTutorial, Provider: "Firebase", IsFree: true, EstimatedMinutes: 120, OrderIndex: 1},
					{Title: "Local Data Persistence in Flutter", URL: "https://www.youtube.com/watch?v=fJtFDrjEvE8", Type: model.ResourceTypeVideo, Provider: "YouTube", IsFree: true, EstimatedMinutes: 90, OrderIndex: 2},
				},
			},
			{
				Title:          "Testing y Publicacion",
				Description:    "Unit tests, widget tests, integration tests, publicacion en stores.",
				OrderIndex:     6,
				EstimatedHours: 13,
				Resources: []model.Resource{
					{Title: "Flutter Testing Documentation", URL: "https://docs.flutter.dev/testing", Type: model.ResourceTypeArticle, Provider: "Flutter", IsFree: true, EstimatedMinutes: 120, OrderIndex: 1},
					{Title: "Publishing to App Stores", URL: "https://docs.flutter.dev/deployment", Type: model.ResourceTypeTutorial, Provider: "Flutter", IsFree: true, EstimatedMinutes: 60, OrderIndex: 2},
				},
			},
		},
	}
}

func seedAssessments(db *gorm.DB, paths []model.LearningPath) {
	assessments := []model.SkillAssessment{
		// Java Path - Milestone 1: Fundamentos de Java
		{MilestoneID: paths[0].Milestones[0].ID, Question: "Cual es la diferencia entre == y .equals() en Java?", Options: model.JSONSlice{"== compara referencias de objetos, .equals() compara contenido", "Son identicos en funcionalidad", "== solo funciona con primitivos", ".equals() solo funciona con Strings"}, CorrectAnswer: 0, Explanation: "El operador == compara referencias (direcciones de memoria), mientras que .equals() compara el contenido logico de los objetos."},
		{MilestoneID: paths[0].Milestones[0].ID, Question: "Que es el polimorfismo en Java?", Options: model.JSONSlice{"La capacidad de una variable de contener diferentes tipos", "Un tipo de herencia multiple", "La capacidad de un objeto de tomar muchas formas", "Un patron de diseno"}, CorrectAnswer: 2, Explanation: "El polimorfismo permite que un objeto se comporte de diferentes formas segun su tipo real en tiempo de ejecucion."},
		{MilestoneID: paths[0].Milestones[0].ID, Question: "Cual coleccion NO permite elementos duplicados?", Options: model.JSONSlice{"ArrayList", "LinkedList", "HashSet", "Vector"}, CorrectAnswer: 2, Explanation: "HashSet implementa la interfaz Set, que por definicion no permite elementos duplicados."},

		// Java Path - Milestone 2: Spring Boot
		{MilestoneID: paths[0].Milestones[1].ID, Question: "Que anotacion se usa para definir un controlador REST en Spring Boot?", Options: model.JSONSlice{"@Controller", "@RestController", "@Service", "@Component"}, CorrectAnswer: 1, Explanation: "@RestController combina @Controller y @ResponseBody, simplificando la creacion de APIs REST."},
		{MilestoneID: paths[0].Milestones[1].ID, Question: "Que hace la anotacion @Autowired?", Options: model.JSONSlice{"Crea un nuevo bean", "Inyecta automaticamente dependencias", "Define una ruta HTTP", "Marca un metodo como transaccional"}, CorrectAnswer: 1, Explanation: "@Autowired permite la inyeccion automatica de dependencias por el contenedor IoC de Spring."},
		{MilestoneID: paths[0].Milestones[1].ID, Question: "Que tipo de inyeccion de dependencias es recomendado en Spring?", Options: model.JSONSlice{"Por campo (@Autowired en atributo)", "Por setter", "Por constructor", "Por interfaz"}, CorrectAnswer: 2, Explanation: "La inyeccion por constructor es la recomendada porque garantiza inmutabilidad y facilita el testing."},

		// React Path - Milestone 3: React Core
		{MilestoneID: paths[1].Milestones[2].ID, Question: "Que hook se usa para manejar estado local en React?", Options: model.JSONSlice{"useEffect", "useState", "useContext", "useReducer"}, CorrectAnswer: 1, Explanation: "useState es el hook fundamental para declarar y actualizar estado local en componentes funcionales."},
		{MilestoneID: paths[1].Milestones[2].ID, Question: "Cual es el proposito principal de useEffect?", Options: model.JSONSlice{"Manejar estado", "Realizar efectos secundarios", "Optimizar renders", "Crear contextos"}, CorrectAnswer: 1, Explanation: "useEffect se utiliza para efectos secundarios como llamadas API, suscripciones y manipulacion del DOM."},
		{MilestoneID: paths[1].Milestones[2].ID, Question: "Que son las props en React?", Options: model.JSONSlice{"Variables globales", "Estado interno del componente", "Datos pasados de padre a hijo", "Metodos del ciclo de vida"}, CorrectAnswer: 2, Explanation: "Las props son datos inmutables que se pasan de un componente padre a un componente hijo."},

		// Python Path - Milestone 2: NumPy y Pandas
		{MilestoneID: paths[2].Milestones[1].ID, Question: "Cual es la estructura de datos principal de Pandas?", Options: model.JSONSlice{"Array", "DataFrame", "Dictionary", "Matrix"}, CorrectAnswer: 1, Explanation: "DataFrame es la estructura principal de Pandas, una tabla bidimensional con filas y columnas etiquetadas."},
		{MilestoneID: paths[2].Milestones[1].ID, Question: "Como se selecciona una columna en un DataFrame?", Options: model.JSONSlice{"df.column_name o df['column_name']", "df.get_column('name')", "df.select('column_name')", "df.col('column_name')"}, CorrectAnswer: 0, Explanation: "Se puede acceder a columnas usando notacion de punto o corchetes con el nombre de la columna."},
		{MilestoneID: paths[2].Milestones[1].ID, Question: "Que funcion de NumPy crea un array de ceros?", Options: model.JSONSlice{"np.empty()", "np.zeros()", "np.null()", "np.blank()"}, CorrectAnswer: 1, Explanation: "np.zeros() crea un array lleno de ceros con la forma especificada."},

		// Flutter Path - Milestone 2: Flutter UI
		{MilestoneID: paths[3].Milestones[1].ID, Question: "Cual es el widget basico para un layout en columna?", Options: model.JSONSlice{"Row", "Column", "Stack", "Container"}, CorrectAnswer: 1, Explanation: "Column organiza sus widgets hijos de forma vertical, uno debajo del otro."},
		{MilestoneID: paths[3].Milestones[1].ID, Question: "Que widget se usa para hacer scroll en Flutter?", Options: model.JSONSlice{"ScrollView", "SingleChildScrollView", "ListView.scrollable", "Scrollable"}, CorrectAnswer: 1, Explanation: "SingleChildScrollView envuelve un widget para hacerlo scrollable cuando su contenido excede el viewport."},
		{MilestoneID: paths[3].Milestones[1].ID, Question: "Cual es la diferencia entre StatelessWidget y StatefulWidget?", Options: model.JSONSlice{"No hay diferencia", "StatefulWidget tiene estado mutable, StatelessWidget no", "StatelessWidget es mas rapido", "StatefulWidget solo se usa con formularios"}, CorrectAnswer: 1, Explanation: "StatefulWidget puede mantener estado interno que cambia durante la vida del widget, mientras StatelessWidget es inmutable."},
	}

	for _, a := range assessments {
		if err := db.Create(&a).Error; err != nil {
			log.Printf("Error seeding assessment: %v", err)
		}
	}
}
