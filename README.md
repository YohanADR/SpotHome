# SpotHome

## Archi

```bash
SpotHome/
├── bin/                               # Binaire de l'application
├── cmd/
│   └── main.go                        # Point d'entrée de l'application
├── infrastructure/                    # Couche Infrastructure
│   ├── config/                        # Gestion de la configuration (YAML, etc.)
│   │   └── config.go                  # Lecture et parsing des fichiers de configuration
│   ├── db/                            # Services de bases de données
│   │   ├── postgis/                   # Package PostGIS
│   │   │   └── postgis.go             # Initialisation de PostGIS avec interface database
│   │   └── redis/                     # Package Redis
│   │       └── redis.go               # Initialisation de Redis avec interface cache
│   ├── logger/                        # Gestion centralisée du logging
│   │   └── logger.go                  # Logger JSON avec Zap
│   ├── messaging/                     # Systèmes de messagerie
│   │   └── kafka/                     # Package Kafka
│   │       └── kafka.go               # Producteur Kafka avec interface message_producer
│   └── server/                        # Serveur HTTP et gestion des routes
│       ├── http_server.go             # Démarrage du serveur HTTP avec Gin
│       └── router/                    # Gestion des routes
│           └── router.go              # Initialisation du routeur et des routes
├── pkg/                               # Paquets partagés et interfaces
│   ├── cache/                         # Interface pour les caches
│   │   └── cache.go                   # Interface générique pour les systèmes de cache (Redis, etc.)
│   ├── database/                      # Interface pour les bases de données
│   │   └── database.go                # Interface pour les systèmes de bases de données (PostGIS, etc.)
│   ├── errors/                        # Gestion des erreurs
│   │   └── errors.go                  # Centralisation des erreurs de l'application
│   ├── events/                        # Gestion des événements
│   │   └── events.go                  # Centralisation de la gestion des événements
│   ├── messaging/                     # Interfaces pour la messagerie
│   │   └── message_producer.go        # Interface pour les producteurs de messages (Kafka, etc.)
│   ├── transport/                     # Gestion des transports
│   │   └── gin_transport.go           # Implémentation du transport HTTP avec Gin
│   └── app/                           # Initialisation de l'application
│       └── app.go                     # Initialisation globale des services
├── go.mod                             # Fichier des dépendances
└── go.sum                             # Hash des versions de modules
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