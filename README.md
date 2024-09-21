# SpotHome

## Archi

```bash
SpotHome/
├── cmd/
│   └── main.go         # Point d'entrée de l'application
├── internal/
│   │
│   ├── PARENTCLASS/           # Dossier parent pour la classe 
│   │   ├── domain/     # Entités et règles métier
│   │   │   └── example.go # Entité example et interfaces des ports
│   │   │
│   │   ├── events/     # Définition des événements (bus de message/connection avec d'autres contexte )
│   │   │
│   │   ├── interactors/ # Cas d'utilisation : logique métier
│   │   │   └── example_interactor.go
│   │   │
│   │   ├── adapters/   # Adaptateurs : implémentations des ports
│   │   │   ├── http/
│   │   │   │   └── example_handler.go       # Définition des routes http(s)
│   │   │   └── db/
│   │   │       └── example_repository.go     # Repo des contrats en DB
│   │   │
│   │   └── ports/    # Ports : interfaces définissant les contrats
│   │       └── example_repository.go
│   │
│   ├── infrastructure/     # Infrastructure : fichier de configuration de l'application (.env/config.yaml/services externes/etc..)
└── go.mod

```

## Commands line utils


General command
```bash
git clone git@github.com:YohanADR/SpotHome.git
```

Make commands 
```bash
make go-createC     # Création des différentes class nécessaire à l'architecture
```

```bash
make go-cleanC      # Suppréssion des class dans l'architecture
```