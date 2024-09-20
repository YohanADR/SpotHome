# Commande pour créer des fichiers
.PHONY: go-createC
go-createC:
	@echo "Comment nommer la classe ? " && \
	read CLASS_NAME; \
	CLASS_NAME_UPPER=$$(echo $$CLASS_NAME | tr '[:lower:]' '[:upper:]'); \
	echo "Création des fichiers pour la classe $$CLASS_NAME"; \
	mkdir -p internal/$$CLASS_NAME_UPPER/domain internal/$$CLASS_NAME_UPPER/events internal/$$CLASS_NAME_UPPER/interactors internal/$$CLASS_NAME_UPPER/adapters/http internal/$$CLASS_NAME_UPPER/adapters/db internal/$$CLASS_NAME_UPPER/ports; \
	echo "package domain\n\n// $$CLASS_NAME entity definition\n" > internal/$$CLASS_NAME_UPPER/domain/$$CLASS_NAME.go; \
	echo "internal/$$CLASS_NAME_UPPER/domain/$$CLASS_NAME.go créé"; \
	echo "package interactors\n\n// $$CLASS_NAMEInteractor contains the business logic for $$CLASS_NAME\n" > internal/$$CLASS_NAME_UPPER/interactors/$$CLASS_NAME\_interactor.go; \
	echo "internal/$$CLASS_NAME_UPPER/interactors/$$CLASS_NAME\_interactor.go créé"; \
	echo "package http\n\n// $$CLASS_NAMEHandler handles HTTP requests for $$CLASS_NAME\n" > internal/$$CLASS_NAME_UPPER/adapters/http/$$CLASS_NAME\_handler.go; \
	echo "internal/$$CLASS_NAME_UPPER/adapters/http/$$CLASS_NAME\_handler.go créé"; \
	echo "package db\n\n// $$CLASS_NAMERepository handles DB operations for $$CLASS_NAME\n" > internal/$$CLASS_NAME_UPPER/adapters/db/$$CLASS_NAME\_repository.go; \
	echo "internal/$$CLASS_NAME_UPPER/adapters/db/$$CLASS_NAME\_repository.go créé"; \
	echo "package ports\n\n// $$CLASS_NAMERepository is the interface for $$CLASS_NAME repository\n" > internal/$$CLASS_NAME_UPPER/ports/$$CLASS_NAME\_repository.go; \
	echo "internal/$$CLASS_NAME_UPPER/ports/$$CLASS_NAME\_repository.go créé"

# Nettoyage des fichiers générés
.PHONY: go-cleanC
go-cleanC:
	@echo "Quel est le nom de la classe à supprimer ? " && \
	read CLASS_NAME; \
	CLASS_NAME_UPPER=$$(echo $$CLASS_NAME | tr '[:lower:]' '[:upper:]'); \
	echo "Suppression des fichiers pour la classe $$CLASS_NAME"; \
	rm -f internal/$$CLASS_NAME_UPPER/domain/$$CLASS_NAME.go; \
	rm -f internal/$$CLASS_NAME_UPPER/interactors/$$CLASS_NAME\_interactor.go; \
	rm -f internal/$$CLASS_NAME_UPPER/adapters/http/$$CLASS_NAME\_handler.go; \
	rm -f internal/$$CLASS_NAME_UPPER/adapters/db/$$CLASS_NAME\_repository.go; \
	rm -f internal/$$CLASS_NAME_UPPER/ports/$$CLASS_NAME\_repository.go; \
	echo "Suppression du dossier parent $$CLASS_NAME_UPPER"; \
	rm -rf internal/$$CLASS_NAME_UPPER; \
	echo "Fichiers et dossier supprimés"
