# Dossiers cibles
DOMAIN_DIR = internal/domain
INTERACTORS_DIR = internal/interactors
ADAPTERS_HTTP_DIR = internal/adapters/http
ADAPTERS_DB_DIR = internal/adapters/db
PORTS_DIR = internal/ports

# Commande pour créer des fichiers
.PHONY: go-createC
go-createC:
	@echo "Comment nommer la classe ? " && \
	read CLASS_NAME; \
	echo "Création des fichiers pour la classe $$CLASS_NAME"; \
	mkdir -p $(DOMAIN_DIR) $(INTERACTORS_DIR) $(ADAPTERS_HTTP_DIR) $(ADAPTERS_DB_DIR) $(PORTS_DIR); \
	echo "package domain\n\n// $$CLASS_NAME entity definition\n" > $(DOMAIN_DIR)/$$CLASS_NAME.go; \
	echo "$(DOMAIN_DIR)/$$CLASS_NAME.go créé"; \
	echo "package interactors\n\n// $$CLASS_NAMEInteractor contains the business logic for $$CLASS_NAME\n" > $(INTERACTORS_DIR)/$$CLASS_NAME\_interactor.go; \
	echo "$(INTERACTORS_DIR)/$$CLASS_NAME\_interactor.go créé"; \
	echo "package http\n\n// $$CLASS_NAMEHandler handles HTTP requests for $$CLASS_NAME\n" > $(ADAPTERS_HTTP_DIR)/$$CLASS_NAME\_handler.go; \
	echo "$(ADAPTERS_HTTP_DIR)/$$CLASS_NAME\_handler.go créé"; \
	echo "package db\n\n// $$CLASS_NAMERepository handles DB operations for $$CLASS_NAME\n" > $(ADAPTERS_DB_DIR)/$$CLASS_NAME\_repository.go; \
	echo "$(ADAPTERS_DB_DIR)/$$CLASS_NAME\_repository.go créé"; \
	echo "package ports\n\n// $$CLASS_NAMERepository is the interface for $$CLASS_NAME repository\n" > $(PORTS_DIR)/$$CLASS_NAME\_repository.go; \
	echo "$(PORTS_DIR)/$$CLASS_NAME\_repository.go créé"

# Nettoyage des fichiers générés
.PHONY: go-cleanC
go-cleanC:
	@echo "Quel est le nom de la classe à supprimer ? " && \
	read CLASS_NAME; \
	echo "Suppression des fichiers pour la classe $$CLASS_NAME"; \
	rm -f $(DOMAIN_DIR)/$$CLASS_NAME.go; \
	rm -f $(INTERACTORS_DIR)/$$CLASS_NAME\_interactor.go; \
	rm -f $(ADAPTERS_HTTP_DIR)/$$CLASS_NAME\_handler.go; \
	rm -f $(ADAPTERS_DB_DIR)/$$CLASS_NAME\_repository.go; \
	rm -f $(PORTS_DIR)/$$CLASS_NAME\_repository.go; \
	echo "Fichiers supprimés"
