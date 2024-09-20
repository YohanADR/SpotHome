# SpotHome

## Archi

SpotHome/ ├── cmd/ │ └── main.go # Point d'entrée de l'application ├── internal/ │ ├── domain/ # Entités et règles métier │ │ └── product.go # Entité Product et interfaces des ports │ ├── usecase/ # Cas d'utilisation : logique métier │ │ └── product_service.go │ ├── adapters/ # Adaptateurs : implémentations des ports │ │ ├── http/ │ │ │ └── product_handler.go │ │ └── db/ │ │ └── product_repository.go │ └── ports/ # Ports : interfaces définissant les contrats │ └── product_repository.go └── go.mod


## Commands line utils

```bash
git clone git@github.com:YohanADR/SpotHome.git