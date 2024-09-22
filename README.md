# SpotHome

## Archi

```bash
SpotHome/
├── bin/                             # Binaire de l'application
├── cmd/
│   └── main.go                      # Point d'entrée de l'application
├── internal/
│   ├── listing/                     # Contexte ou module métier (ex : gestion des annonces)
│   │   ├── domain/                  # Couche Domaine (Business Logic)
│   │   │   ├── listing.go           # Entité "Listing" et interfaces des ports
│   │   │   ├── listing_service.go   # Interface du service métier (Ports côté application)
│   │   │
│   │   ├── events/                  # Gestion des événements (optionnel, si on implémente des Event Sourcing ou CQRS)
│   │   │   └── listing_events.go    # Définition des événements
│   │   │
│   │   ├── usecases/                # Cas d'utilisation (Application Logic)
│   │   │   └── listing_interactor.go # Logique métier, appels aux ports côté domaine
│   │   │
│   │   ├── adapters/                # Adaptateurs (Ports implémentés côté infrastructure)
│   │   │   ├── http/
│   │   │   │   └── listing_handler.go  # Gestion des routes HTTP avec Gin
│   │   │   ├── db/
│   │   │   │   └── listing_repository.go # Gestion des accès à la base de données (Ex : via GORM)
│   │   │   └── messaging/           # Adaptateurs pour la communication (ex : message queue)
│   │   │       └── kafka_publisher.go
│   │   │
│   │   └── ports/                   # Interfaces pour définir les contrats
│   │       ├── repository.go        # Interface du repository pour les annonces
│   │       └── service.go           # Interface pour les services métiers (implémentés dans les "usecases")
│   │
│   ├── infrastructure/              # Couche Infrastructure
│   │   ├── config/                  # Gestion de la configuration (ex: fichiers .env, YAML, etc.)
│   │   │   └── config.go            # Lecture et parsing des fichiers de configuration
│   │   ├── database/                # Initialisation de la base de données (ex : connection pool)
│   │   │   └── db.go                # Configuration et initialisation de la DB (ex : GORM)
│   │   ├── server/                  # Configuration du serveur HTTP
│   │   │   └── http_server.go       # Démarrage du serveur Gin
│   │   └── messaging/               # Infrastructure pour la gestion des messages (ex: Kafka, RabbitMQ)
│           └── kafka.go
│
├── pkg/                             # Paquets partagés entre différents contextes métier
│   └── errors/                      # Gestion des erreurs partagées
│       └── custom_errors.go
│
├── go.mod                           # Fichier des dépendances
└── go.sum                           # Hash des versions de modules

```

## Commands line utils


General command
```bash
git clone git@github.com:YohanADR/SpotHome.git
```

Make commands 
```bash
make create     # Création des différentes class nécessaire à l'architecture
```

```bash
make clean     # Suppréssion des class dans l'architecture
```