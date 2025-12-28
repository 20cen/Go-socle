.PHONY: build install clean test run help

# Variables
BINARY_NAME=go-scaffold
INSTALL_PATH=/usr/local/bin

# Couleurs pour les messages
GREEN=\033[0;32m
YELLOW=\033[0;33m
NC=\033[0m # No Color

help: ## Afficher cette aide
	@echo "$(GREEN)Commandes disponibles:$(NC)"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "  $(YELLOW)%-15s$(NC) %s\n", $$1, $$2}'

build: ## Compiler le générateur
	@echo "$(GREEN)Compilation de $(BINARY_NAME)...$(NC)"
	go build -o $(BINARY_NAME) main.go
	@echo "$(GREEN)✓ Compilation terminée!$(NC)"

install: build ## Installer le générateur globalement
	@echo "$(GREEN)Installation de $(BINARY_NAME) dans $(INSTALL_PATH)...$(NC)"
	sudo mv $(BINARY_NAME) $(INSTALL_PATH)/
	@echo "$(GREEN)✓ Installation terminée! Vous pouvez maintenant utiliser '$(BINARY_NAME)' depuis n'importe où.$(NC)"

clean: ## Nettoyer les fichiers de build
	@echo "$(GREEN)Nettoyage...$(NC)"
	rm -f $(BINARY_NAME)
	go clean
	@echo "$(GREEN)✓ Nettoyage terminé!$(NC)"

test: ## Exécuter les tests
	@echo "$(GREEN)Exécution des tests...$(NC)"
	go test ./... -v
	@echo "$(GREEN)✓ Tests terminés!$(NC)"

run: build ## Compiler et exécuter
	@echo "$(GREEN)Exécution de $(BINARY_NAME)...$(NC)"
	./$(BINARY_NAME)

deps: ## Télécharger les dépendances
	@echo "$(GREEN)Téléchargement des dépendances...$(NC)"
	go mod download
	go mod tidy
	@echo "$(GREEN)✓ Dépendances téléchargées!$(NC)"

init: ## Initialiser un nouveau projet (usage: make init PROJECT=mon-projet)
	@if [ -z "$(PROJECT)" ]; then \
		echo "$(YELLOW)Usage: make init PROJECT=nom-du-projet$(NC)"; \
		exit 1; \
	fi
	@echo "$(GREEN)Initialisation du projet $(PROJECT)...$(NC)"
	./$(BINARY_NAME) init $(PROJECT)
	@echo "$(GREEN)✓ Projet $(PROJECT) initialisé!$(NC)"

generate: ## Générer le code (usage: make generate SCHEMA=path/to/schema.yaml)
	@if [ -z "$(SCHEMA)" ]; then \
		echo "$(YELLOW)Usage: make generate SCHEMA=path/to/schema.yaml$(NC)"; \
		exit 1; \
	fi
	@echo "$(GREEN)Génération du code depuis $(SCHEMA)...$(NC)"
	./$(BINARY_NAME) generate $(SCHEMA)
	@echo "$(GREEN)✓ Code généré!$(NC)"

generate-all: ## Générer le code pour tous les schémas
	@echo "$(GREEN)Génération du code pour tous les schémas...$(NC)"
	./$(BINARY_NAME) generate --all
	@echo "$(GREEN)✓ Code généré pour tous les schémas!$(NC)"

make-schema: ## Créer un nouveau schéma (usage: make make-schema NAME=user)
	@if [ -z "$(NAME)" ]; then \
		echo "$(YELLOW)Usage: make make-schema NAME=nom-du-schema$(NC)"; \
		exit 1; \
	fi
	@echo "$(GREEN)Création du schéma $(NAME)...$(NC)"
	./$(BINARY_NAME) make:schema $(NAME)
	@echo "$(GREEN)✓ Schéma $(NAME) créé!$(NC)"

make-migration: ## Créer une nouvelle migration (usage: make make-migration NAME=create_users_table)
	@if [ -z "$(NAME)" ]; then \
		echo "$(YELLOW)Usage: make make-migration NAME=nom-de-la-migration$(NC)"; \
		exit 1; \
	fi
	@echo "$(GREEN)Création de la migration $(NAME)...$(NC)"
	./$(BINARY_NAME) make:migration $(NAME)
	@echo "$(GREEN)✓ Migration $(NAME) créée!$(NC)"

fmt: ## Formater le code
	@echo "$(GREEN)Formatage du code...$(NC)"
	go fmt ./...
	@echo "$(GREEN)✓ Code formaté!$(NC)"

lint: ## Vérifier le code avec golangci-lint
	@echo "$(GREEN)Vérification du code...$(NC)"
	golangci-lint run
	@echo "$(GREEN)✓ Vérification terminée!$(NC)"

version: ## Afficher la version
	@echo "$(GREEN)go-scaffold version 1.0.0$(NC)"
