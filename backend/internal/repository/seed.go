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
		pythonEssentialPath(),
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

func pythonEssentialPath() model.LearningPath {
	return model.LearningPath{
		Title:          "Python Esencial a Avanzado - Guia Completa",
		Description:    "Domina Python desde cero hasta nivel avanzado. Cubre todos los fundamentos, estructuras de datos, POO, modulos, testing, y temas avanzados como decoradores, generadores, asyncio y mas. Incluye 3+ ejemplos y 3+ ejercicios por tema.",
		Slug:           "python-esencial-avanzado",
		Difficulty:     model.DifficultyBeginner,
		EstimatedHours: 120,
		Category:       "Python",
		IconURL:        "https://cdn.jsdelivr.net/gh/devicons/devicon/icons/python/python-original.svg",
		IsPublished:    true,
		Milestones: []model.Milestone{
			// ── Milestone 1: Introduccion y Entorno ──
			{
				Title:          "Introduccion y Entorno",
				Description:    "Que es Python, historia y versiones. Instalacion en Windows, Mac y Linux. IDEs: VS Code, PyCharm, Jupyter Notebook. Tu primer programa: print('Hola Mundo'). Variables y tipos basicos (int, float, str, bool). Operadores aritmeticos, de comparacion y logicos. Funcion input() y type(). Incluye 3 ejemplos practicos y 3 ejercicios para afianzar conceptos.",
				OrderIndex:     1,
				EstimatedHours: 8,
				Resources: []model.Resource{
					{Title: "Tutorial oficial de Python - Documentacion", URL: "https://docs.python.org/es/3/tutorial/index.html", Type: model.ResourceTypeTutorial, Provider: "docs.python.org", IsFree: true, EstimatedMinutes: 120, OrderIndex: 1},
					{Title: "Python para principiantes - W3Schools", URL: "https://www.w3schools.com/python/python_intro.asp", Type: model.ResourceTypeTutorial, Provider: "W3Schools", IsFree: true, EstimatedMinutes: 60, OrderIndex: 2},
					{Title: "Instalar Python y configurar VS Code", URL: "https://realpython.com/python-first-steps/", Type: model.ResourceTypeArticle, Provider: "Real Python", IsFree: true, EstimatedMinutes: 45, OrderIndex: 3},
					{Title: "Python Getting Started - Ejercicios interactivos", URL: "https://www.w3schools.com/python/python_getstarted.asp", Type: model.ResourceTypeExercise, Provider: "W3Schools", IsFree: true, EstimatedMinutes: 30, OrderIndex: 4},
					{Title: "Historia y filosofia de Python", URL: "https://docs.python.org/3/faq/general.html", Type: model.ResourceTypeArticle, Provider: "docs.python.org", IsFree: true, EstimatedMinutes: 20, OrderIndex: 5},
				},
			},
			// ── Milestone 2: Variables, Tipos de Datos y Operadores ──
			{
				Title:          "Variables, Tipos de Datos y Operadores",
				Description:    "Variables: reglas de nombrado y snake_case. Tipos: int, float, str, bool, None. Conversion de tipos (casting): int(), float(), str(). Operadores aritmeticos: +, -, *, /, //, %%, **. Operadores de comparacion: ==, !=, >, <, >=, <=. Operadores logicos: and, or, not. Operadores de asignacion: +=, -=, *=. f-strings y format(). Incluye 3 ejemplos practicos y 3 ejercicios.",
				OrderIndex:     2,
				EstimatedHours: 8,
				Resources: []model.Resource{
					{Title: "Variables y tipos de datos en Python", URL: "https://realpython.com/python-variables/", Type: model.ResourceTypeArticle, Provider: "Real Python", IsFree: true, EstimatedMinutes: 45, OrderIndex: 1},
					{Title: "Python Variables - W3Schools", URL: "https://www.w3schools.com/python/python_variables.asp", Type: model.ResourceTypeTutorial, Provider: "W3Schools", IsFree: true, EstimatedMinutes: 30, OrderIndex: 2},
					{Title: "Operadores en Python - Documentacion oficial", URL: "https://docs.python.org/3/library/operator.html", Type: model.ResourceTypeArticle, Provider: "docs.python.org", IsFree: true, EstimatedMinutes: 40, OrderIndex: 3},
					{Title: "f-strings en Python: guia completa", URL: "https://realpython.com/python-f-strings/", Type: model.ResourceTypeArticle, Provider: "Real Python", IsFree: true, EstimatedMinutes: 30, OrderIndex: 4},
					{Title: "Ejercicios de variables y tipos", URL: "https://www.w3schools.com/python/python_variables_exercises.asp", Type: model.ResourceTypeExercise, Provider: "W3Schools", IsFree: true, EstimatedMinutes: 40, OrderIndex: 5},
				},
			},
			// ── Milestone 3: Estructuras de Control ──
			{
				Title:          "Estructuras de Control",
				Description:    "Condicionales: if, elif, else. Operador ternario. Bucles while y for con range(). Sentencias break, continue y pass. Loops anidados. match/case (Python 3.10+). Patrones de control de flujo comunes. Incluye 3 ejemplos practicos y 3 ejercicios.",
				OrderIndex:     3,
				EstimatedHours: 8,
				Resources: []model.Resource{
					{Title: "Control de flujo en Python", URL: "https://docs.python.org/3/tutorial/controlflow.html", Type: model.ResourceTypeTutorial, Provider: "docs.python.org", IsFree: true, EstimatedMinutes: 60, OrderIndex: 1},
					{Title: "Python If...Else - W3Schools", URL: "https://www.w3schools.com/python/python_conditions.asp", Type: model.ResourceTypeTutorial, Provider: "W3Schools", IsFree: true, EstimatedMinutes: 30, OrderIndex: 2},
					{Title: "Structural Pattern Matching en Python", URL: "https://realpython.com/python310-new-features/#structural-pattern-matching", Type: model.ResourceTypeArticle, Provider: "Real Python", IsFree: true, EstimatedMinutes: 40, OrderIndex: 3},
					{Title: "Python Loops - W3Schools", URL: "https://www.w3schools.com/python/python_for_loops.asp", Type: model.ResourceTypeTutorial, Provider: "W3Schools", IsFree: true, EstimatedMinutes: 30, OrderIndex: 4},
					{Title: "Ejercicios de estructuras de control", URL: "https://www.w3schools.com/python/python_conditions_exercises.asp", Type: model.ResourceTypeExercise, Provider: "W3Schools", IsFree: true, EstimatedMinutes: 45, OrderIndex: 5},
				},
			},
			// ── Milestone 4: Listas y Tuplas ──
			{
				Title:          "Listas y Tuplas",
				Description:    "Listas: creacion, acceso por indice, slicing. Metodos: append, insert, remove, pop, sort, reverse. List comprehensions. Tuplas: creacion e inmutabilidad. Diferencias lista vs tupla. Desempaquetado (unpacking). Funciones con listas: len, min, max, sum, sorted. enumerate() y zip(). Incluye 3 ejemplos practicos y 3 ejercicios.",
				OrderIndex:     4,
				EstimatedHours: 8,
				Resources: []model.Resource{
					{Title: "Listas en Python - Tutorial oficial", URL: "https://docs.python.org/3/tutorial/datastructures.html#more-on-lists", Type: model.ResourceTypeTutorial, Provider: "docs.python.org", IsFree: true, EstimatedMinutes: 60, OrderIndex: 1},
					{Title: "Python Lists - W3Schools", URL: "https://www.w3schools.com/python/python_lists.asp", Type: model.ResourceTypeTutorial, Provider: "W3Schools", IsFree: true, EstimatedMinutes: 40, OrderIndex: 2},
					{Title: "List Comprehensions en Python", URL: "https://realpython.com/list-comprehension-python/", Type: model.ResourceTypeArticle, Provider: "Real Python", IsFree: true, EstimatedMinutes: 35, OrderIndex: 3},
					{Title: "Tuplas en Python", URL: "https://www.w3schools.com/python/python_tuples.asp", Type: model.ResourceTypeTutorial, Provider: "W3Schools", IsFree: true, EstimatedMinutes: 25, OrderIndex: 4},
					{Title: "Ejercicios de listas y tuplas", URL: "https://www.w3schools.com/python/python_lists_exercises.asp", Type: model.ResourceTypeExercise, Provider: "W3Schools", IsFree: true, EstimatedMinutes: 45, OrderIndex: 5},
				},
			},
			// ── Milestone 5: Diccionarios y Sets ──
			{
				Title:          "Diccionarios y Sets",
				Description:    "Diccionarios: creacion, acceso, metodos. get(), keys(), values(), items(). Dictionary comprehensions. Sets: creacion, operaciones (union, interseccion, diferencia). Frozensets. Cuando usar dict vs list vs set vs tuple. defaultdict y OrderedDict del modulo collections. Incluye 3 ejemplos practicos y 3 ejercicios.",
				OrderIndex:     5,
				EstimatedHours: 8,
				Resources: []model.Resource{
					{Title: "Diccionarios en Python - Tutorial oficial", URL: "https://docs.python.org/3/tutorial/datastructures.html#dictionaries", Type: model.ResourceTypeTutorial, Provider: "docs.python.org", IsFree: true, EstimatedMinutes: 50, OrderIndex: 1},
					{Title: "Python Dictionaries - W3Schools", URL: "https://www.w3schools.com/python/python_dictionaries.asp", Type: model.ResourceTypeTutorial, Provider: "W3Schools", IsFree: true, EstimatedMinutes: 35, OrderIndex: 2},
					{Title: "Sets en Python - Real Python", URL: "https://realpython.com/python-sets/", Type: model.ResourceTypeArticle, Provider: "Real Python", IsFree: true, EstimatedMinutes: 40, OrderIndex: 3},
					{Title: "Modulo collections en Python", URL: "https://docs.python.org/3/library/collections.html", Type: model.ResourceTypeArticle, Provider: "docs.python.org", IsFree: true, EstimatedMinutes: 45, OrderIndex: 4},
					{Title: "Ejercicios de diccionarios y sets", URL: "https://www.w3schools.com/python/python_dictionaries_exercises.asp", Type: model.ResourceTypeExercise, Provider: "W3Schools", IsFree: true, EstimatedMinutes: 40, OrderIndex: 5},
				},
			},
			// ── Milestone 6: Cadenas de Texto (Strings) ──
			{
				Title:          "Cadenas de Texto (Strings)",
				Description:    "Metodos de strings: upper, lower, strip, split, join, replace. Slicing de strings. f-strings avanzados con formato numerico y alineacion. Expresiones regulares con el modulo re: search, match, findall, sub. Encode/decode (UTF-8). Template strings. Incluye 3 ejemplos practicos y 3 ejercicios.",
				OrderIndex:     6,
				EstimatedHours: 7,
				Resources: []model.Resource{
					{Title: "Metodos de String en Python", URL: "https://www.w3schools.com/python/python_strings_methods.asp", Type: model.ResourceTypeTutorial, Provider: "W3Schools", IsFree: true, EstimatedMinutes: 40, OrderIndex: 1},
					{Title: "Expresiones Regulares en Python", URL: "https://docs.python.org/3/howto/regex.html", Type: model.ResourceTypeTutorial, Provider: "docs.python.org", IsFree: true, EstimatedMinutes: 60, OrderIndex: 2},
					{Title: "Python Regex - Real Python", URL: "https://realpython.com/regex-python/", Type: model.ResourceTypeArticle, Provider: "Real Python", IsFree: true, EstimatedMinutes: 50, OrderIndex: 3},
					{Title: "String Formatting en Python", URL: "https://realpython.com/python-string-formatting/", Type: model.ResourceTypeArticle, Provider: "Real Python", IsFree: true, EstimatedMinutes: 35, OrderIndex: 4},
					{Title: "Ejercicios de strings", URL: "https://www.w3schools.com/python/python_strings_exercises.asp", Type: model.ResourceTypeExercise, Provider: "W3Schools", IsFree: true, EstimatedMinutes: 35, OrderIndex: 5},
				},
			},
			// ── Milestone 7: Funciones ──
			{
				Title:          "Funciones",
				Description:    "Definicion con def. Parametros y argumentos posicionales y por nombre. return vs print. Diferencia entre funcion, metodo y procedimiento. *args y **kwargs explicados a fondo con ejemplos reales. Valores por defecto. Scope: local, global, nonlocal. Funciones como objetos de primera clase. Funciones lambda. Funciones built-in: map, filter, reduce. Type hints en funciones. Incluye 3 ejemplos practicos y 3 ejercicios.",
				OrderIndex:     7,
				EstimatedHours: 10,
				Resources: []model.Resource{
					{Title: "Definir funciones en Python - Tutorial oficial", URL: "https://docs.python.org/3/tutorial/controlflow.html#defining-functions", Type: model.ResourceTypeTutorial, Provider: "docs.python.org", IsFree: true, EstimatedMinutes: 50, OrderIndex: 1},
					{Title: "Python Functions - W3Schools", URL: "https://www.w3schools.com/python/python_functions.asp", Type: model.ResourceTypeTutorial, Provider: "W3Schools", IsFree: true, EstimatedMinutes: 35, OrderIndex: 2},
					{Title: "args y kwargs en Python", URL: "https://realpython.com/python-kwargs-and-args/", Type: model.ResourceTypeArticle, Provider: "Real Python", IsFree: true, EstimatedMinutes: 40, OrderIndex: 3},
					{Title: "Funciones Lambda en Python", URL: "https://realpython.com/python-lambda/", Type: model.ResourceTypeArticle, Provider: "Real Python", IsFree: true, EstimatedMinutes: 35, OrderIndex: 4},
					{Title: "Ejercicios de funciones", URL: "https://www.w3schools.com/python/python_functions_exercises.asp", Type: model.ResourceTypeExercise, Provider: "W3Schools", IsFree: true, EstimatedMinutes: 45, OrderIndex: 5},
				},
			},
			// ── Milestone 8: Programacion Orientada a Objetos (POO) ──
			{
				Title:          "Programacion Orientada a Objetos (POO)",
				Description:    "Que es una clase vs un objeto. __init__ y self. Atributos de instancia vs de clase. Metodos de instancia, de clase (@classmethod) y estaticos (@staticmethod). Encapsulamiento: public, _protected, __private (name mangling). Herencia simple y multiple. super(). Polimorfismo. Clases abstractas (ABC). Dataclasses con @dataclass. Incluye 3 ejemplos practicos y 3 ejercicios.",
				OrderIndex:     8,
				EstimatedHours: 10,
				Resources: []model.Resource{
					{Title: "Clases en Python - Tutorial oficial", URL: "https://docs.python.org/3/tutorial/classes.html", Type: model.ResourceTypeTutorial, Provider: "docs.python.org", IsFree: true, EstimatedMinutes: 60, OrderIndex: 1},
					{Title: "OOP en Python - Real Python", URL: "https://realpython.com/python3-object-oriented-programming/", Type: model.ResourceTypeArticle, Provider: "Real Python", IsFree: true, EstimatedMinutes: 55, OrderIndex: 2},
					{Title: "Python Classes - W3Schools", URL: "https://www.w3schools.com/python/python_classes.asp", Type: model.ResourceTypeTutorial, Provider: "W3Schools", IsFree: true, EstimatedMinutes: 40, OrderIndex: 3},
					{Title: "Dataclasses en Python", URL: "https://realpython.com/python-data-classes/", Type: model.ResourceTypeArticle, Provider: "Real Python", IsFree: true, EstimatedMinutes: 40, OrderIndex: 4},
					{Title: "Ejercicios de POO en Python", URL: "https://www.w3schools.com/python/python_classes_exercises.asp", Type: model.ResourceTypeExercise, Provider: "W3Schools", IsFree: true, EstimatedMinutes: 50, OrderIndex: 5},
				},
			},
			// ── Milestone 9: Manejo de Errores y Excepciones ──
			{
				Title:          "Manejo de Errores y Excepciones",
				Description:    "try, except, else, finally. Tipos de excepciones: ValueError, TypeError, KeyError, IndexError, FileNotFoundError y mas. Crear excepciones personalizadas con herencia. raise. assert para validaciones. Context managers (sentencia with). Buenas practicas de manejo de errores. Incluye 3 ejemplos practicos y 3 ejercicios.",
				OrderIndex:     9,
				EstimatedHours: 7,
				Resources: []model.Resource{
					{Title: "Errores y Excepciones - Tutorial oficial", URL: "https://docs.python.org/3/tutorial/errors.html", Type: model.ResourceTypeTutorial, Provider: "docs.python.org", IsFree: true, EstimatedMinutes: 45, OrderIndex: 1},
					{Title: "Excepciones en Python - Real Python", URL: "https://realpython.com/python-exceptions/", Type: model.ResourceTypeArticle, Provider: "Real Python", IsFree: true, EstimatedMinutes: 40, OrderIndex: 2},
					{Title: "Python Try Except - W3Schools", URL: "https://www.w3schools.com/python/python_try_except.asp", Type: model.ResourceTypeTutorial, Provider: "W3Schools", IsFree: true, EstimatedMinutes: 25, OrderIndex: 3},
					{Title: "Excepciones personalizadas en Python", URL: "https://realpython.com/python-custom-exceptions/", Type: model.ResourceTypeArticle, Provider: "Real Python", IsFree: true, EstimatedMinutes: 35, OrderIndex: 4},
					{Title: "Ejercicios de excepciones", URL: "https://www.w3schools.com/python/python_try_except_exercises.asp", Type: model.ResourceTypeExercise, Provider: "W3Schools", IsFree: true, EstimatedMinutes: 30, OrderIndex: 5},
				},
			},
			// ── Milestone 10: Modulos y Paquetes ──
			{
				Title:          "Modulos y Paquetes",
				Description:    "import, from...import, as. Crear modulos propios. __name__ == '__main__'. Paquetes y __init__.py. pip y requirements.txt. Entornos virtuales con venv. Modulos estandar: os, sys, json, datetime, math, random, collections. Explorando la libreria estandar. Incluye 3 ejemplos practicos y 3 ejercicios.",
				OrderIndex:     10,
				EstimatedHours: 8,
				Resources: []model.Resource{
					{Title: "Modulos en Python - Tutorial oficial", URL: "https://docs.python.org/3/tutorial/modules.html", Type: model.ResourceTypeTutorial, Provider: "docs.python.org", IsFree: true, EstimatedMinutes: 50, OrderIndex: 1},
					{Title: "Python Modules - W3Schools", URL: "https://www.w3schools.com/python/python_modules.asp", Type: model.ResourceTypeTutorial, Provider: "W3Schools", IsFree: true, EstimatedMinutes: 30, OrderIndex: 2},
					{Title: "Entornos virtuales con venv", URL: "https://realpython.com/python-virtual-environments-a-primer/", Type: model.ResourceTypeArticle, Provider: "Real Python", IsFree: true, EstimatedMinutes: 40, OrderIndex: 3},
					{Title: "pip y gestion de dependencias", URL: "https://realpython.com/what-is-pip/", Type: model.ResourceTypeArticle, Provider: "Real Python", IsFree: true, EstimatedMinutes: 35, OrderIndex: 4},
					{Title: "Indice de la Libreria Estandar", URL: "https://docs.python.org/3/library/index.html", Type: model.ResourceTypeArticle, Provider: "docs.python.org", IsFree: true, EstimatedMinutes: 30, OrderIndex: 5},
				},
			},
			// ── Milestone 11: Archivos y E/S ──
			{
				Title:          "Archivos y Entrada/Salida",
				Description:    "open(), read(), write(), close(). Modos de apertura: r, w, a, rb, wb. Context manager con with para archivos. Modulo csv: leer y escribir archivos CSV. Modulo json: json.load, json.dump, json.loads, json.dumps. Paths con pathlib. Manejo de archivos grandes con iteradores. Incluye 3 ejemplos practicos y 3 ejercicios.",
				OrderIndex:     11,
				EstimatedHours: 7,
				Resources: []model.Resource{
					{Title: "Lectura y escritura de archivos - Tutorial oficial", URL: "https://docs.python.org/3/tutorial/inputoutput.html#reading-and-writing-files", Type: model.ResourceTypeTutorial, Provider: "docs.python.org", IsFree: true, EstimatedMinutes: 40, OrderIndex: 1},
					{Title: "Python File Handling - W3Schools", URL: "https://www.w3schools.com/python/python_file_handling.asp", Type: model.ResourceTypeTutorial, Provider: "W3Schools", IsFree: true, EstimatedMinutes: 30, OrderIndex: 2},
					{Title: "Trabajar con JSON en Python", URL: "https://realpython.com/python-json/", Type: model.ResourceTypeArticle, Provider: "Real Python", IsFree: true, EstimatedMinutes: 35, OrderIndex: 3},
					{Title: "pathlib: rutas orientadas a objetos", URL: "https://realpython.com/python-pathlib/", Type: model.ResourceTypeArticle, Provider: "Real Python", IsFree: true, EstimatedMinutes: 35, OrderIndex: 4},
					{Title: "Ejercicios de manejo de archivos", URL: "https://www.w3schools.com/python/python_file_handling_exercises.asp", Type: model.ResourceTypeExercise, Provider: "W3Schools", IsFree: true, EstimatedMinutes: 35, OrderIndex: 5},
				},
			},
			// ── Milestone 12: Decoradores y Generadores ──
			{
				Title:          "Decoradores y Generadores",
				Description:    "Funciones como argumentos y closures. Decoradores: sintaxis @decorator y como funcionan internamente. Decoradores con argumentos. Decoradores de clase. @property, @staticmethod, @classmethod. Generadores con yield. Generator expressions. Modulo itertools: chain, product, permutations, combinations. Incluye 3 ejemplos practicos y 3 ejercicios.",
				OrderIndex:     12,
				EstimatedHours: 9,
				Resources: []model.Resource{
					{Title: "Decoradores en Python - Primer paso", URL: "https://realpython.com/primer-on-python-decorators/", Type: model.ResourceTypeArticle, Provider: "Real Python", IsFree: true, EstimatedMinutes: 50, OrderIndex: 1},
					{Title: "Generadores en Python", URL: "https://realpython.com/introduction-to-python-generators/", Type: model.ResourceTypeArticle, Provider: "Real Python", IsFree: true, EstimatedMinutes: 40, OrderIndex: 2},
					{Title: "itertools - Documentacion oficial", URL: "https://docs.python.org/3/library/itertools.html", Type: model.ResourceTypeArticle, Provider: "docs.python.org", IsFree: true, EstimatedMinutes: 45, OrderIndex: 3},
					{Title: "Closures en Python", URL: "https://realpython.com/python-closure/", Type: model.ResourceTypeArticle, Provider: "Real Python", IsFree: true, EstimatedMinutes: 30, OrderIndex: 4},
					{Title: "Ejercicios de decoradores y generadores", URL: "https://www.w3schools.com/python/python_iterators.asp", Type: model.ResourceTypeExercise, Provider: "W3Schools", IsFree: true, EstimatedMinutes: 40, OrderIndex: 5},
				},
			},
			// ── Milestone 13: Testing y Debugging ──
			{
				Title:          "Testing y Debugging",
				Description:    "Modulo unittest: TestCase, setUp, tearDown. pytest: funciones de test, fixtures, parametrize. Assertions y matchers. Mocking con unittest.mock y patch. Coverage: medir cobertura de tests. TDD basico: Red-Green-Refactor. pdb debugger: breakpoints e inspeccion. Modulo logging: niveles, formatos y handlers. Incluye 3 ejemplos practicos y 3 ejercicios.",
				OrderIndex:     13,
				EstimatedHours: 9,
				Resources: []model.Resource{
					{Title: "unittest - Documentacion oficial", URL: "https://docs.python.org/3/library/unittest.html", Type: model.ResourceTypeArticle, Provider: "docs.python.org", IsFree: true, EstimatedMinutes: 50, OrderIndex: 1},
					{Title: "Comenzar con pytest", URL: "https://realpython.com/pytest-python-testing/", Type: model.ResourceTypeArticle, Provider: "Real Python", IsFree: true, EstimatedMinutes: 50, OrderIndex: 2},
					{Title: "Mocking en Python", URL: "https://realpython.com/python-mock-library/", Type: model.ResourceTypeArticle, Provider: "Real Python", IsFree: true, EstimatedMinutes: 40, OrderIndex: 3},
					{Title: "Logging en Python", URL: "https://realpython.com/python-logging/", Type: model.ResourceTypeArticle, Provider: "Real Python", IsFree: true, EstimatedMinutes: 35, OrderIndex: 4},
					{Title: "pdb - El debugger de Python", URL: "https://docs.python.org/3/library/pdb.html", Type: model.ResourceTypeArticle, Provider: "docs.python.org", IsFree: true, EstimatedMinutes: 30, OrderIndex: 5},
				},
			},
			// ── Milestone 14: Programacion Asincrona ──
			{
				Title:          "Programacion Asincrona",
				Description:    "Concurrencia vs paralelismo: conceptos clave. Threading: modulo threading y GIL. Multiprocessing: procesos independientes con multiprocessing. asyncio: async/await, event loop, coroutines. aiohttp para peticiones HTTP asincronas. Futures y Tasks. concurrent.futures: ThreadPoolExecutor y ProcessPoolExecutor. Incluye 3 ejemplos practicos y 3 ejercicios.",
				OrderIndex:     14,
				EstimatedHours: 9,
				Resources: []model.Resource{
					{Title: "asyncio - Documentacion oficial", URL: "https://docs.python.org/3/library/asyncio.html", Type: model.ResourceTypeArticle, Provider: "docs.python.org", IsFree: true, EstimatedMinutes: 60, OrderIndex: 1},
					{Title: "Async IO en Python: guia completa", URL: "https://realpython.com/async-io-python/", Type: model.ResourceTypeArticle, Provider: "Real Python", IsFree: true, EstimatedMinutes: 55, OrderIndex: 2},
					{Title: "Threading en Python", URL: "https://realpython.com/intro-to-python-threading/", Type: model.ResourceTypeArticle, Provider: "Real Python", IsFree: true, EstimatedMinutes: 45, OrderIndex: 3},
					{Title: "Multiprocessing en Python", URL: "https://docs.python.org/3/library/multiprocessing.html", Type: model.ResourceTypeArticle, Provider: "docs.python.org", IsFree: true, EstimatedMinutes: 50, OrderIndex: 4},
					{Title: "concurrent.futures", URL: "https://docs.python.org/3/library/concurrent.futures.html", Type: model.ResourceTypeArticle, Provider: "docs.python.org", IsFree: true, EstimatedMinutes: 35, OrderIndex: 5},
				},
			},
			// ── Milestone 15: Temas Avanzados ──
			{
				Title:          "Temas Avanzados",
				Description:    "Metaclases: type, __new__, __init_subclass__. Descriptores: __get__, __set__, __delete__. Context managers avanzados con __enter__ y __exit__. Type hints avanzados: Generic, Protocol, TypeVar, Union, Optional. Walrus operator :=. Pattern matching (match/case). Memory management y garbage collection. Performance: profiling con cProfile y timeit. Comprensiones avanzadas anidadas. __slots__ para optimizar memoria. Incluye 3 ejemplos practicos y 3 ejercicios.",
				OrderIndex:     15,
				EstimatedHours: 12,
				Resources: []model.Resource{
					{Title: "Metaclases en Python", URL: "https://realpython.com/python-metaclasses/", Type: model.ResourceTypeArticle, Provider: "Real Python", IsFree: true, EstimatedMinutes: 55, OrderIndex: 1},
					{Title: "Descriptores en Python", URL: "https://docs.python.org/3/howto/descriptor.html", Type: model.ResourceTypeTutorial, Provider: "docs.python.org", IsFree: true, EstimatedMinutes: 50, OrderIndex: 2},
					{Title: "Type Hints avanzados", URL: "https://realpython.com/python-type-checking/", Type: model.ResourceTypeArticle, Provider: "Real Python", IsFree: true, EstimatedMinutes: 50, OrderIndex: 3},
					{Title: "Performance Tips en Python", URL: "https://docs.python.org/3/faq/programming.html", Type: model.ResourceTypeArticle, Provider: "docs.python.org", IsFree: true, EstimatedMinutes: 40, OrderIndex: 4},
					{Title: "Novedades de Python 3.10+", URL: "https://realpython.com/python310-new-features/", Type: model.ResourceTypeArticle, Provider: "Real Python", IsFree: true, EstimatedMinutes: 35, OrderIndex: 5},
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

		// ═══════════════════════════════════════════════════════════════
		// Python Esencial a Avanzado - paths[4]
		// ═══════════════════════════════════════════════════════════════

		// Python Path - Milestone 1: Introduccion y Entorno
		{MilestoneID: paths[4].Milestones[0].ID, Question: "Cual es la funcion que imprime texto en la consola en Python?", Options: model.JSONSlice{"echo()", "console.log()", "print()", "write()"}, CorrectAnswer: 2, Explanation: "La funcion print() es la forma estandar de mostrar texto en la consola de Python."},
		{MilestoneID: paths[4].Milestones[0].ID, Question: "Que funcion se usa para conocer el tipo de una variable?", Options: model.JSONSlice{"typeof()", "type()", "class()", "isinstance()"}, CorrectAnswer: 1, Explanation: "type() devuelve el tipo del objeto pasado como argumento, por ejemplo type(42) retorna <class 'int'>."},
		{MilestoneID: paths[4].Milestones[0].ID, Question: "Que funcion permite leer datos del usuario por teclado?", Options: model.JSONSlice{"read()", "scan()", "input()", "get()"}, CorrectAnswer: 2, Explanation: "input() lee una linea de texto ingresada por el usuario y la retorna como string."},
		{MilestoneID: paths[4].Milestones[0].ID, Question: "Python es un lenguaje de tipado:", Options: model.JSONSlice{"Estatico y fuerte", "Dinamico y fuerte", "Estatico y debil", "Dinamico y debil"}, CorrectAnswer: 1, Explanation: "Python es dinamico (los tipos se verifican en tiempo de ejecucion) y fuerte (no permite operaciones implicitas entre tipos incompatibles)."},
		{MilestoneID: paths[4].Milestones[0].ID, Question: "Cual de estos NO es un tipo basico de Python?", Options: model.JSONSlice{"int", "float", "char", "bool"}, CorrectAnswer: 2, Explanation: "Python no tiene un tipo char. Los caracteres individuales se representan como strings de longitud 1."},

		// Python Path - Milestone 2: Variables, Tipos de Datos y Operadores
		{MilestoneID: paths[4].Milestones[1].ID, Question: "Cual es la convencion de nombrado de variables en Python?", Options: model.JSONSlice{"camelCase", "PascalCase", "snake_case", "kebab-case"}, CorrectAnswer: 2, Explanation: "Python sigue PEP 8 que recomienda snake_case para variables y funciones (ejemplo: mi_variable)."},
		{MilestoneID: paths[4].Milestones[1].ID, Question: "Que operador realiza la division entera en Python?", Options: model.JSONSlice{"/", "//", "%", "**"}, CorrectAnswer: 1, Explanation: "El operador // realiza la division entera (floor division), descartando la parte decimal. 7 // 2 retorna 3."},
		{MilestoneID: paths[4].Milestones[1].ID, Question: "Que resultado da: int('42.5')?", Options: model.JSONSlice{"42", "42.5", "Error: ValueError", "43"}, CorrectAnswer: 2, Explanation: "int() no puede convertir directamente un string con punto decimal. Se debe usar int(float('42.5')) para obtener 42."},
		{MilestoneID: paths[4].Milestones[1].ID, Question: "Cual es la forma correcta de usar f-strings?", Options: model.JSONSlice{"f'Hola {nombre}'", "'Hola' + nombre", "format('Hola', nombre)", "'Hola %s' nombre"}, CorrectAnswer: 0, Explanation: "f-strings (formatted string literals) usan el prefijo f y llaves {} para interpolar variables directamente."},
		{MilestoneID: paths[4].Milestones[1].ID, Question: "Que operador calcula la potencia en Python?", Options: model.JSONSlice{"^", "**", "pow", "^^"}, CorrectAnswer: 1, Explanation: "El operador ** calcula la potencia. 2 ** 3 retorna 8. Tambien se puede usar la funcion pow(2, 3)."},

		// Python Path - Milestone 3: Estructuras de Control
		{MilestoneID: paths[4].Milestones[2].ID, Question: "Cual es la sintaxis correcta de un condicional en Python?", Options: model.JSONSlice{"if (x > 5) {}", "if x > 5:", "if x > 5 then", "if x > 5 do"}, CorrectAnswer: 1, Explanation: "Python usa la sintaxis 'if condicion:' seguida de un bloque indentado. No usa llaves ni parentesis obligatorios."},
		{MilestoneID: paths[4].Milestones[2].ID, Question: "Que hace la sentencia 'pass' en Python?", Options: model.JSONSlice{"Salta a la siguiente iteracion", "Termina el bucle", "No hace nada (placeholder)", "Sale de la funcion"}, CorrectAnswer: 2, Explanation: "pass es una sentencia nula que no hace nada. Se usa como placeholder cuando se requiere una sentencia sintacticamente."},
		{MilestoneID: paths[4].Milestones[2].ID, Question: "Cual es la diferencia entre 'break' y 'continue'?", Options: model.JSONSlice{"Son iguales", "break salta una iteracion, continue termina el bucle", "break termina el bucle, continue salta a la siguiente iteracion", "Ambos terminan el bucle"}, CorrectAnswer: 2, Explanation: "break sale completamente del bucle, mientras que continue salta el resto del cuerpo y continua con la siguiente iteracion."},
		{MilestoneID: paths[4].Milestones[2].ID, Question: "Que genera range(3)?", Options: model.JSONSlice{"[1, 2, 3]", "[0, 1, 2]", "[0, 1, 2, 3]", "[3]"}, CorrectAnswer: 1, Explanation: "range(3) genera la secuencia 0, 1, 2. range(n) produce n numeros empezando desde 0."},
		{MilestoneID: paths[4].Milestones[2].ID, Question: "En que version de Python se introdujo match/case?", Options: model.JSONSlice{"Python 3.8", "Python 3.9", "Python 3.10", "Python 3.11"}, CorrectAnswer: 2, Explanation: "Structural Pattern Matching (match/case) fue introducido en Python 3.10 como PEP 634."},

		// Python Path - Milestone 4: Listas y Tuplas
		{MilestoneID: paths[4].Milestones[3].ID, Question: "Que retorna lista[1:4] si lista = [10, 20, 30, 40, 50]?", Options: model.JSONSlice{"[10, 20, 30]", "[20, 30, 40]", "[20, 30, 40, 50]", "[10, 20, 30, 40]"}, CorrectAnswer: 1, Explanation: "El slicing lista[1:4] retorna los elementos de indice 1, 2 y 3 (el indice final es exclusivo)."},
		{MilestoneID: paths[4].Milestones[3].ID, Question: "Cual es la diferencia principal entre una lista y una tupla?", Options: model.JSONSlice{"Las listas son mas rapidas", "Las tuplas pueden contener mas elementos", "Las listas son mutables, las tuplas son inmutables", "Las tuplas solo contienen numeros"}, CorrectAnswer: 2, Explanation: "Las listas son mutables (se pueden modificar), las tuplas son inmutables (no se pueden cambiar despues de crearlas)."},
		{MilestoneID: paths[4].Milestones[3].ID, Question: "Que hace [x**2 for x in range(5)]?", Options: model.JSONSlice{"[0, 1, 2, 3, 4]", "[1, 4, 9, 16, 25]", "[0, 1, 4, 9, 16]", "[0, 2, 4, 6, 8]"}, CorrectAnswer: 2, Explanation: "Esta list comprehension genera los cuadrados de 0 a 4: [0, 1, 4, 9, 16]."},
		{MilestoneID: paths[4].Milestones[3].ID, Question: "Que funcion combina dos listas en pares de tuplas?", Options: model.JSONSlice{"combine()", "merge()", "zip()", "pair()"}, CorrectAnswer: 2, Explanation: "zip() combina iterables elemento por elemento. zip([1,2], ['a','b']) produce [(1,'a'), (2,'b')]."},
		{MilestoneID: paths[4].Milestones[3].ID, Question: "Que metodo agrega un elemento al final de una lista?", Options: model.JSONSlice{"add()", "push()", "append()", "insert()"}, CorrectAnswer: 2, Explanation: "append() agrega un elemento al final de la lista. insert() permite especificar la posicion."},

		// Python Path - Milestone 5: Diccionarios y Sets
		{MilestoneID: paths[4].Milestones[4].ID, Question: "Que retorna dict.get('clave', 'default') si la clave no existe?", Options: model.JSONSlice{"None", "KeyError", "'default'", "False"}, CorrectAnswer: 2, Explanation: "El metodo get() retorna el valor por defecto especificado si la clave no existe, en lugar de lanzar KeyError."},
		{MilestoneID: paths[4].Milestones[4].ID, Question: "Que estructura de datos no permite elementos duplicados?", Options: model.JSONSlice{"list", "tuple", "dict", "set"}, CorrectAnswer: 3, Explanation: "Los sets (conjuntos) no permiten elementos duplicados. Al agregar un duplicado, simplemente se ignora."},
		{MilestoneID: paths[4].Milestones[4].ID, Question: "Que retorna {1, 2, 3} & {2, 3, 4}?", Options: model.JSONSlice{"{1, 2, 3, 4}", "{2, 3}", "{1, 4}", "Error"}, CorrectAnswer: 1, Explanation: "El operador & realiza la interseccion de sets, retornando los elementos comunes: {2, 3}."},
		{MilestoneID: paths[4].Milestones[4].ID, Question: "Que es un frozenset?", Options: model.JSONSlice{"Un set ordenado", "Un set inmutable", "Un set vacio", "Un set con tipo fijo"}, CorrectAnswer: 1, Explanation: "frozenset es una version inmutable de set. No se puede modificar despues de crearlo y puede usarse como clave de diccionario."},
		{MilestoneID: paths[4].Milestones[4].ID, Question: "Cual es la sintaxis de una dictionary comprehension?", Options: model.JSONSlice{"{k: v for k, v in items}", "[k: v for k, v in items]", "(k: v for k, v in items)", "dict(k: v for k, v in items)"}, CorrectAnswer: 0, Explanation: "Las dictionary comprehensions usan llaves con la sintaxis {clave: valor for elemento in iterable}."},

		// Python Path - Milestone 6: Cadenas de Texto
		{MilestoneID: paths[4].Milestones[5].ID, Question: "Que metodo divide un string en una lista por un separador?", Options: model.JSONSlice{"divide()", "split()", "cut()", "separate()"}, CorrectAnswer: 1, Explanation: "split() divide un string por el separador indicado (espacio por defecto) y retorna una lista."},
		{MilestoneID: paths[4].Milestones[5].ID, Question: "Que modulo se usa para expresiones regulares en Python?", Options: model.JSONSlice{"regex", "re", "regexp", "pattern"}, CorrectAnswer: 1, Explanation: "El modulo 're' de la libreria estandar proporciona operaciones con expresiones regulares."},
		{MilestoneID: paths[4].Milestones[5].ID, Question: "Que retorna 'Python'[1:4]?", Options: model.JSONSlice{"'Pyt'", "'yth'", "'ytho'", "'Pyth'"}, CorrectAnswer: 1, Explanation: "El slicing [1:4] toma los caracteres en los indices 1, 2 y 3: 'y', 't', 'h' = 'yth'."},
		{MilestoneID: paths[4].Milestones[5].ID, Question: "Que hace el metodo strip()?", Options: model.JSONSlice{"Elimina la primera letra", "Elimina espacios al inicio y final", "Convierte a minusculas", "Revierte el string"}, CorrectAnswer: 1, Explanation: "strip() elimina los espacios en blanco (y otros caracteres especificados) al inicio y final del string."},
		{MilestoneID: paths[4].Milestones[5].ID, Question: "Que hace ' '.join(['a', 'b', 'c'])?", Options: model.JSONSlice{"' abc'", "'a b c'", "'abc'", "['a', 'b', 'c']"}, CorrectAnswer: 1, Explanation: "join() une los elementos de un iterable usando el string como separador: 'a b c'."},

		// Python Path - Milestone 7: Funciones
		{MilestoneID: paths[4].Milestones[6].ID, Question: "Que permite *args en una funcion?", Options: model.JSONSlice{"Recibir un numero variable de argumentos posicionales como tupla", "Recibir un solo argumento obligatorio", "Recibir argumentos con nombre", "Definir valores por defecto"}, CorrectAnswer: 0, Explanation: "*args permite recibir un numero variable de argumentos posicionales, empaquetados en una tupla."},
		{MilestoneID: paths[4].Milestones[6].ID, Question: "Que permite **kwargs en una funcion?", Options: model.JSONSlice{"Argumentos posicionales variables", "Argumentos con nombre variables como diccionario", "Dos argumentos obligatorios", "Argumentos de solo lectura"}, CorrectAnswer: 1, Explanation: "**kwargs permite recibir un numero variable de argumentos con nombre (keyword arguments), empaquetados en un diccionario."},
		{MilestoneID: paths[4].Milestones[6].ID, Question: "Que es una funcion lambda?", Options: model.JSONSlice{"Una funcion que no retorna nada", "Una funcion anonima de una sola expresion", "Una funcion recursiva", "Una funcion asincrona"}, CorrectAnswer: 1, Explanation: "lambda crea funciones anonimas pequenas con una sola expresion. Ejemplo: lambda x: x * 2."},
		{MilestoneID: paths[4].Milestones[6].ID, Question: "Que hace la funcion map()?", Options: model.JSONSlice{"Crea un diccionario", "Aplica una funcion a cada elemento de un iterable", "Filtra elementos de una lista", "Ordena una lista"}, CorrectAnswer: 1, Explanation: "map() aplica una funcion a cada elemento de un iterable y retorna un iterador con los resultados."},
		{MilestoneID: paths[4].Milestones[6].ID, Question: "Cual es la diferencia entre 'return' y 'print' en una funcion?", Options: model.JSONSlice{"Son equivalentes", "return devuelve un valor al llamador, print solo muestra en consola", "print devuelve un valor, return no", "return solo funciona con numeros"}, CorrectAnswer: 1, Explanation: "return devuelve un valor que puede ser usado por el codigo que llamo a la funcion. print solo muestra texto en consola sin devolver nada util."},

		// Python Path - Milestone 8: POO
		{MilestoneID: paths[4].Milestones[7].ID, Question: "Que es self en una clase de Python?", Options: model.JSONSlice{"Una palabra reservada obligatoria", "Una referencia a la instancia actual del objeto", "El nombre de la clase", "Un tipo de dato especial"}, CorrectAnswer: 1, Explanation: "self es una referencia a la instancia actual de la clase. Se usa para acceder a atributos y metodos del objeto."},
		{MilestoneID: paths[4].Milestones[7].ID, Question: "Que metodo se ejecuta al crear una nueva instancia?", Options: model.JSONSlice{"__start__()", "__create__()", "__init__()", "__new__()"}, CorrectAnswer: 2, Explanation: "__init__() es el metodo inicializador que se ejecuta automaticamente al crear una nueva instancia de la clase."},
		{MilestoneID: paths[4].Milestones[7].ID, Question: "Que indica el prefijo __ (doble guion bajo) en un atributo?", Options: model.JSONSlice{"Es publico", "Es protegido", "Activa name mangling (pseudo-privado)", "Es una constante"}, CorrectAnswer: 2, Explanation: "El doble guion bajo activa name mangling: Python transforma __atributo a _Clase__atributo para dificultar el acceso accidental."},
		{MilestoneID: paths[4].Milestones[7].ID, Question: "Que hace @classmethod?", Options: model.JSONSlice{"Define un metodo que no recibe ningun argumento", "Define un metodo que recibe la clase como primer argumento (cls)", "Define un metodo privado", "Define un metodo estatico"}, CorrectAnswer: 1, Explanation: "@classmethod hace que el metodo reciba la clase (cls) como primer argumento en lugar de la instancia (self)."},
		{MilestoneID: paths[4].Milestones[7].ID, Question: "Que modulo se usa para crear clases abstractas?", Options: model.JSONSlice{"abstract", "abc", "interface", "base"}, CorrectAnswer: 1, Explanation: "El modulo abc (Abstract Base Classes) proporciona ABC y @abstractmethod para crear clases abstractas."},

		// Python Path - Milestone 9: Manejo de Errores
		{MilestoneID: paths[4].Milestones[8].ID, Question: "En que orden se ejecutan los bloques try/except/else/finally?", Options: model.JSONSlice{"try -> except -> else -> finally", "try -> else (si no hay error) o except (si hay error) -> finally", "try -> finally -> except -> else", "except -> try -> else -> finally"}, CorrectAnswer: 1, Explanation: "Se ejecuta try, luego else si no hubo excepcion o except si hubo una, y finally siempre se ejecuta al final."},
		{MilestoneID: paths[4].Milestones[8].ID, Question: "Que excepcion se lanza al acceder a una clave inexistente en un diccionario?", Options: model.JSONSlice{"ValueError", "IndexError", "KeyError", "AttributeError"}, CorrectAnswer: 2, Explanation: "KeyError se lanza cuando se intenta acceder a una clave que no existe en un diccionario con la notacion dict[clave]."},
		{MilestoneID: paths[4].Milestones[8].ID, Question: "Como se crea una excepcion personalizada?", Options: model.JSONSlice{"Usando la funcion error()", "Creando una clase que hereda de Exception", "Usando @exception decorator", "Modificando el modulo errors"}, CorrectAnswer: 1, Explanation: "Las excepciones personalizadas se crean como clases que heredan de Exception o de alguna subclase existente."},
		{MilestoneID: paths[4].Milestones[8].ID, Question: "Que bloque se ejecuta SIEMPRE, haya o no excepcion?", Options: model.JSONSlice{"try", "except", "else", "finally"}, CorrectAnswer: 3, Explanation: "El bloque finally se ejecuta siempre, independientemente de si hubo excepcion o no. Es ideal para liberar recursos."},
		{MilestoneID: paths[4].Milestones[8].ID, Question: "Que hace la sentencia 'raise'?", Options: model.JSONSlice{"Captura una excepcion", "Lanza (dispara) una excepcion manualmente", "Ignora una excepcion", "Registra una excepcion en log"}, CorrectAnswer: 1, Explanation: "raise permite lanzar una excepcion de forma manual. Se puede usar con o sin un objeto de excepcion."},

		// Python Path - Milestone 10: Modulos y Paquetes
		{MilestoneID: paths[4].Milestones[9].ID, Question: "Que verifica la condicion if __name__ == '__main__'?", Options: model.JSONSlice{"Si el modulo esta instalado", "Si el archivo se ejecuta directamente (no importado)", "Si Python esta actualizado", "Si hay errores de sintaxis"}, CorrectAnswer: 1, Explanation: "Cuando un archivo se ejecuta directamente, __name__ vale '__main__'. Si es importado como modulo, __name__ tiene el nombre del modulo."},
		{MilestoneID: paths[4].Milestones[9].ID, Question: "Que archivo convierte un directorio en un paquete Python?", Options: model.JSONSlice{"__main__.py", "__init__.py", "setup.py", "package.json"}, CorrectAnswer: 1, Explanation: "__init__.py indica que el directorio es un paquete Python. Puede estar vacio o contener codigo de inicializacion."},
		{MilestoneID: paths[4].Milestones[9].ID, Question: "Que comando crea un entorno virtual en Python?", Options: model.JSONSlice{"pip create env", "python -m venv nombre", "virtualenv --create", "python --new-env"}, CorrectAnswer: 1, Explanation: "python -m venv nombre_del_entorno crea un entorno virtual con su propia copia del interprete y pip."},
		{MilestoneID: paths[4].Milestones[9].ID, Question: "Para que sirve requirements.txt?", Options: model.JSONSlice{"Para documentar el proyecto", "Para listar las dependencias del proyecto con sus versiones", "Para configurar Python", "Para definir variables de entorno"}, CorrectAnswer: 1, Explanation: "requirements.txt lista las dependencias del proyecto. Se usa con pip install -r requirements.txt para instalar todas."},
		{MilestoneID: paths[4].Milestones[9].ID, Question: "Que hace 'from os import path'?", Options: model.JSONSlice{"Importa todo el modulo os", "Importa solo el submodulo path del modulo os", "Crea un alias para os", "Instala el modulo os"}, CorrectAnswer: 1, Explanation: "from...import permite importar elementos especificos de un modulo sin necesidad de usar el prefijo del modulo."},

		// Python Path - Milestone 11: Archivos y E/S
		{MilestoneID: paths[4].Milestones[10].ID, Question: "Cual es la ventaja de usar 'with open()' en lugar de open()/close()?", Options: model.JSONSlice{"Es mas rapido", "Cierra el archivo automaticamente, incluso si hay error", "Permite leer archivos mas grandes", "No hay diferencia"}, CorrectAnswer: 1, Explanation: "El context manager 'with' garantiza que el archivo se cierre correctamente, incluso si ocurre una excepcion."},
		{MilestoneID: paths[4].Milestones[10].ID, Question: "Que modo de apertura se usa para agregar contenido sin borrar lo existente?", Options: model.JSONSlice{"'r'", "'w'", "'a'", "'x'"}, CorrectAnswer: 2, Explanation: "El modo 'a' (append) agrega contenido al final del archivo sin borrar el contenido existente."},
		{MilestoneID: paths[4].Milestones[10].ID, Question: "Que funcion de json convierte un diccionario a JSON y lo guarda en archivo?", Options: model.JSONSlice{"json.write()", "json.dump()", "json.save()", "json.export()"}, CorrectAnswer: 1, Explanation: "json.dump() serializa un objeto Python y lo escribe en un archivo. json.dumps() hace lo mismo pero retorna un string."},
		{MilestoneID: paths[4].Milestones[10].ID, Question: "Que modulo de Python proporciona paths orientados a objetos?", Options: model.JSONSlice{"os.path", "pathlib", "filepath", "sys.path"}, CorrectAnswer: 1, Explanation: "pathlib ofrece clases Path que representan rutas del sistema de archivos de forma orientada a objetos, mas moderna que os.path."},
		{MilestoneID: paths[4].Milestones[10].ID, Question: "Que modo se usa para leer archivos binarios?", Options: model.JSONSlice{"'r'", "'b'", "'rb'", "'bin'"}, CorrectAnswer: 2, Explanation: "El modo 'rb' (read binary) abre un archivo en modo lectura binaria, necesario para archivos como imagenes o PDFs."},

		// Python Path - Milestone 12: Decoradores y Generadores
		{MilestoneID: paths[4].Milestones[11].ID, Question: "Que es un closure en Python?", Options: model.JSONSlice{"Una funcion que se cierra sola", "Una funcion interna que recuerda variables del scope externo", "Un tipo de clase", "Un metodo de cierre de archivos"}, CorrectAnswer: 1, Explanation: "Un closure es una funcion interna que recuerda y tiene acceso a variables del scope de la funcion que la contiene, incluso despues de que esta haya terminado."},
		{MilestoneID: paths[4].Milestones[11].ID, Question: "Que hace la sentencia 'yield' en una funcion?", Options: model.JSONSlice{"Termina la funcion", "Pausa la funcion y produce un valor, convirtiendola en generador", "Retorna un error", "Importa un modulo"}, CorrectAnswer: 1, Explanation: "yield pausa la ejecucion de la funcion y produce un valor. La proxima vez que se llame next(), continua desde donde se pauso."},
		{MilestoneID: paths[4].Milestones[11].ID, Question: "Que es un decorador en Python?", Options: model.JSONSlice{"Un comentario especial", "Una funcion que modifica el comportamiento de otra funcion", "Un tipo de clase", "Una variable global"}, CorrectAnswer: 1, Explanation: "Un decorador es una funcion que recibe otra funcion como argumento y retorna una version modificada de ella, usando la sintaxis @decorador."},
		{MilestoneID: paths[4].Milestones[11].ID, Question: "Que hace @property en una clase?", Options: model.JSONSlice{"Hace un atributo publico", "Convierte un metodo en un atributo de solo lectura", "Crea una propiedad estatica", "Define una constante"}, CorrectAnswer: 1, Explanation: "@property permite acceder a un metodo como si fuera un atributo, encapsulando la logica de acceso."},
		{MilestoneID: paths[4].Milestones[11].ID, Question: "Que diferencia hay entre un generador y una lista?", Options: model.JSONSlice{"No hay diferencia", "El generador produce valores bajo demanda (lazy), la lista los almacena todos en memoria", "La lista es mas lenta", "El generador no puede iterarse"}, CorrectAnswer: 1, Explanation: "Los generadores producen valores uno a la vez (lazy evaluation), usando mucha menos memoria que una lista que almacena todos los elementos."},

		// Python Path - Milestone 13: Testing y Debugging
		{MilestoneID: paths[4].Milestones[12].ID, Question: "Que framework de testing es mas popular en el ecosistema Python?", Options: model.JSONSlice{"unittest", "pytest", "nose", "doctest"}, CorrectAnswer: 1, Explanation: "pytest es el framework de testing mas popular por su sintaxis simple, autodescubrimiento de tests y sistema de fixtures."},
		{MilestoneID: paths[4].Milestones[12].ID, Question: "Que es un fixture en pytest?", Options: model.JSONSlice{"Un tipo de assertion", "Una funcion que proporciona datos o configuracion reutilizable para tests", "Un reporte de tests", "Un tipo de mock"}, CorrectAnswer: 1, Explanation: "Un fixture es una funcion decorada con @pytest.fixture que prepara datos, objetos o configuracion que los tests necesitan."},
		{MilestoneID: paths[4].Milestones[12].ID, Question: "Que hace unittest.mock.patch()?", Options: model.JSONSlice{"Corrige errores automaticamente", "Reemplaza un objeto con un mock durante el test", "Actualiza el codigo fuente", "Parchea el interprete de Python"}, CorrectAnswer: 1, Explanation: "patch() reemplaza temporalmente un objeto con un Mock, permitiendo aislar la unidad bajo test de sus dependencias."},
		{MilestoneID: paths[4].Milestones[12].ID, Question: "Que principio sigue TDD?", Options: model.JSONSlice{"Write-Test-Deploy", "Red-Green-Refactor", "Plan-Code-Test", "Design-Build-Test"}, CorrectAnswer: 1, Explanation: "TDD sigue el ciclo Red (escribir test que falla) -> Green (escribir codigo minimo para que pase) -> Refactor (mejorar el codigo)."},
		{MilestoneID: paths[4].Milestones[12].ID, Question: "Que modulo de Python se usa para debugging interactivo?", Options: model.JSONSlice{"debug", "pdb", "dbg", "inspector"}, CorrectAnswer: 1, Explanation: "pdb (Python Debugger) es el debugger interactivo de la libreria estandar. Se puede activar con breakpoint() o import pdb; pdb.set_trace()."},

		// Python Path - Milestone 14: Programacion Asincrona
		{MilestoneID: paths[4].Milestones[13].ID, Question: "Que es el GIL (Global Interpreter Lock) en Python?", Options: model.JSONSlice{"Un tipo de lock para archivos", "Un mecanismo que permite que solo un thread ejecute bytecode Python a la vez", "Un sistema de permisos", "Un tipo de excepcion"}, CorrectAnswer: 1, Explanation: "El GIL es un mutex que protege el acceso a objetos Python, impidiendo que multiples threads ejecuten bytecode Python simultaneamente."},
		{MilestoneID: paths[4].Milestones[13].ID, Question: "Cuando es preferible usar multiprocessing en lugar de threading?", Options: model.JSONSlice{"Para operaciones de I/O", "Para tareas CPU-bound que necesitan paralelismo real", "Para tareas simples", "Nunca, threading siempre es mejor"}, CorrectAnswer: 1, Explanation: "multiprocessing es preferible para tareas CPU-bound porque cada proceso tiene su propio GIL, permitiendo paralelismo real."},
		{MilestoneID: paths[4].Milestones[13].ID, Question: "Que palabra clave define una funcion asincrona en Python?", Options: model.JSONSlice{"async def", "await def", "coroutine def", "parallel def"}, CorrectAnswer: 0, Explanation: "async def define una coroutine function. Dentro de ella se puede usar await para esperar otras coroutines."},
		{MilestoneID: paths[4].Milestones[13].ID, Question: "Que hace 'await' en una coroutine?", Options: model.JSONSlice{"Detiene el programa completamente", "Pausa la coroutine actual y permite que otras se ejecuten", "Crea un nuevo thread", "Termina la coroutine"}, CorrectAnswer: 1, Explanation: "await pausa la coroutine actual, devolviendo el control al event loop para que otras coroutines puedan ejecutarse mientras se espera el resultado."},
		{MilestoneID: paths[4].Milestones[13].ID, Question: "Que funcion inicia el event loop de asyncio?", Options: model.JSONSlice{"asyncio.start()", "asyncio.run()", "asyncio.begin()", "asyncio.execute()"}, CorrectAnswer: 1, Explanation: "asyncio.run() es la forma recomendada de ejecutar la coroutine principal. Crea un event loop, ejecuta la coroutine y lo cierra."},

		// Python Path - Milestone 15: Temas Avanzados
		{MilestoneID: paths[4].Milestones[14].ID, Question: "Que es una metaclase en Python?", Options: model.JSONSlice{"Una clase muy grande", "Una clase cuyas instancias son clases", "Un tipo de herencia", "Una clase abstracta"}, CorrectAnswer: 1, Explanation: "Una metaclase es una clase cuyas instancias son clases. type es la metaclase por defecto de todas las clases en Python."},
		{MilestoneID: paths[4].Milestones[14].ID, Question: "Que hace el walrus operator (:=)?", Options: model.JSONSlice{"Compara dos valores", "Asigna un valor a una variable como parte de una expresion", "Define un tipo", "Crea una lambda"}, CorrectAnswer: 1, Explanation: "El walrus operator := (PEP 572) permite asignar un valor a una variable dentro de una expresion, por ejemplo: if (n := len(a)) > 10."},
		{MilestoneID: paths[4].Milestones[14].ID, Question: "Para que sirve __slots__ en una clase?", Options: model.JSONSlice{"Para definir metodos especiales", "Para restringir atributos y optimizar uso de memoria", "Para crear slots de herencia", "Para definir constantes"}, CorrectAnswer: 1, Explanation: "__slots__ restringe los atributos de una clase a los declarados, eliminando __dict__ y ahorrando memoria significativamente."},
		{MilestoneID: paths[4].Milestones[14].ID, Question: "Que es un descriptor en Python?", Options: model.JSONSlice{"Un tipo de comentario", "Un objeto que define __get__, __set__ o __delete__ para controlar acceso a atributos", "Un tipo de iterador", "Un tipo de decorador"}, CorrectAnswer: 1, Explanation: "Un descriptor es un objeto que implementa al menos uno de __get__, __set__ o __delete__, permitiendo personalizar el acceso a atributos de clase."},
		{MilestoneID: paths[4].Milestones[14].ID, Question: "Que herramienta se usa para medir performance de codigo Python?", Options: model.JSONSlice{"speedtest", "cProfile", "benchmark", "optimizer"}, CorrectAnswer: 1, Explanation: "cProfile es el profiler de la libreria estandar que mide el tiempo de ejecucion de cada funcion. timeit se usa para medir fragmentos pequenos de codigo."},
	}

	for _, a := range assessments {
		if err := db.Create(&a).Error; err != nil {
			log.Printf("Error seeding assessment: %v", err)
		}
	}
}
