#!/bin/bash

# Vérification si un nom de classe est fourni
if [ -z "$1" ]; then
  echo "Le nom de la classe ne peut pas être vide"
  exit 1
fi

CLASS_NAME=$1

echo "Création des fichiers pour la classe $CLASS_NAME"

# Création des dossiers nécessaires
mkdir -p internal/$CLASS_NAME/domain internal/$CLASS_NAME/events internal/$CLASS_NAME/usecases internal/$CLASS_NAME/adapters/http internal/$CLASS_NAME/adapters/db internal/$CLASS_NAME/adapters/messaging internal/$CLASS_NAME/ports

# Création des fichiers dans la couche Domain
echo "package domain\n\n// $CLASS_NAME entity definition\n" > internal/$CLASS_NAME/domain/$CLASS_NAME.go
echo "internal/$CLASS_NAME/domain/$CLASS_NAME.go créé"
echo "package domain\n\n// $CLASS_NAME service definition\n" > internal/$CLASS_NAME/domain/${CLASS_NAME}_service.go
echo "internal/$CLASS_NAME/domain/${CLASS_NAME}_service.go créé"

# Création des fichiers dans la couche Events
echo "package events\n\n// $CLASS_NAME events definition\n" > internal/$CLASS_NAME/events/${CLASS_NAME}_events.go
echo "internal/$CLASS_NAME/events/${CLASS_NAME}_events.go créé"

# Création des fichiers dans la couche Usecases
echo "package usecases\n\n// ${CLASS_NAME}Interactor contains the business logic for $CLASS_NAME\n" > internal/$CLASS_NAME/usecases/${CLASS_NAME}_interactor.go
echo "internal/$CLASS_NAME/usecases/${CLASS_NAME}_interactor.go créé"

# Création des fichiers dans la couche HTTP
echo "package http\n\n// ${CLASS_NAME}Handler handles HTTP requests for $CLASS_NAME\n" > internal/$CLASS_NAME/adapters/http/${CLASS_NAME}_handler.go
echo "internal/$CLASS_NAME/adapters/http/${CLASS_NAME}_handler.go créé"

# Création des fichiers dans la couche DB
echo "package db\n\n// ${CLASS_NAME}Repository handles DB operations for $CLASS_NAME\n" > internal/$CLASS_NAME/adapters/db/${CLASS_NAME}_repository.go
echo "internal/$CLASS_NAME/adapters/db/${CLASS_NAME}_repository.go créé"

# Création des fichiers dans la couche Messaging
echo "package messaging\n\n// ${CLASS_NAME}Publisher publishes messages for $CLASS_NAME\n" > internal/$CLASS_NAME/adapters/messaging/${CLASS_NAME}_kafka_publisher.go
echo "internal/$CLASS_NAME/adapters/messaging/${CLASS_NAME}_kafka_publisher.go créé"

# Création des fichiers dans la couche Ports
echo "package ports\n\n// ${CLASS_NAME}Repository is the interface for $CLASS_NAME repository\n" > internal/$CLASS_NAME/ports/${CLASS_NAME}_repository.go
echo "internal/$CLASS_NAME/ports/${CLASS_NAME}_repository.go créé"
echo "package ports\n\n// ${CLASS_NAME}Service is the interface for $CLASS_NAME services\n" > internal/$CLASS_NAME/ports/${CLASS_NAME}_service.go
echo "internal/$CLASS_NAME/ports/${CLASS_NAME}_service.go créé"

echo "Fichiers et répertoires pour $CLASS_NAME créés avec succès."
