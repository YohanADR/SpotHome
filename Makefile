# Commande pour créer des fichiers selon l'architecture hexagonale
.PHONY: create
create:
	@echo "Comment nommer la classe ? " && \
	read CLASS_NAME; \
	./create_class.sh $$CLASS_NAME

# Nettoyage des fichiers générés selon l'architecture hexagonale
.PHONY: clean
clean:
	@echo "Quel est le nom de la classe à supprimer ? " && \
	read CLASS_NAME; \
	if [ -z "$$CLASS_NAME" ]; then echo "Le nom de la classe ne peut pas être vide"; exit 1; fi; \
	echo "Suppression des fichiers pour la classe $$CLASS_NAME"; \
	rm -f internal/$$CLASS_NAME/domain/$$CLASS_NAME.go; \
	rm -f internal/$$CLASS_NAME/domain/${CLASS_NAME}_service.go; \
	rm -f internal/$$CLASS_NAME/events/${CLASS_NAME}_events.go; \
	rm -f internal/$$CLASS_NAME/usecases/${CLASS_NAME}_interactor.go; \
	rm -f internal/$$CLASS_NAME/adapters/http/${CLASS_NAME}_handler.go; \
	rm -f internal/$$CLASS_NAME/adapters/db/${CLASS_NAME}_repository.go; \
	rm -f internal/$$CLASS_NAME/adapters/messaging/${CLASS_NAME}_kafka_publisher.go; \
	rm -f internal/$$CLASS_NAME/ports/${CLASS_NAME}_repository.go; \
	rm -f internal/$$CLASS_NAME/ports/${CLASS_NAME}_service.go; \
	echo "Suppression du dossier parent $$CLASS_NAME"; \
	rm -rf internal/$$CLASS_NAME; \
	echo "Fichiers et dossier supprimés"

# Build de l'application
.PHONY: build
build:
	go build -o bin/spotHome cmd/main.go
