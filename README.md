# SpotHome

## Archi

```bash
SpotHome/
├── cmd/
│   └── main.go         # Point d'entrée de l'application
├── internal/
│   │
│   ├── domain/         # Entités et règles métier
│   │   └── product.go  # Entité Product et interfaces des ports
│   │
│   ├── interactors/        # Cas d'utilisation : logique métier
│   │   └── product_interactors.go
│   │
│   ├── adapters/       # Adaptateurs : implémentations des ports
│   │   ├── http/
│   │   │   └── product_handler.go       # Définition des routes http(s)
│   │   └── db/
│   │       └── product_repository.go       # Repo des contrats en DB
│   │
│   └── ports/          # Ports : interfaces définissant les contrats
│       └── product_repository.go
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