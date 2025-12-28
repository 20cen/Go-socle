#!/bin/bash

# Script d'installation pour go-scaffold
# Usage: ./install.sh

set -e

# Couleurs
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
RED='\033[0;31m'
NC='\033[0m' # No Color

echo -e "${GREEN}"
echo "╔══════════════════════════════════════════════╗"
echo "║   Installation de go-scaffold                ║"
echo "║   Générateur de code automatique pour Go     ║"
echo "╚══════════════════════════════════════════════╝"
echo -e "${NC}"

# Vérifier si Go est installé
if ! command -v go &> /dev/null; then
    echo -e "${RED}❌ Go n'est pas installé. Veuillez installer Go 1.21 ou supérieur.${NC}"
    echo -e "${YELLOW}Téléchargez Go depuis: https://golang.org/dl/${NC}"
    exit 1
fi

GO_VERSION=$(go version | awk '{print $3}' | sed 's/go//')
echo -e "${GREEN}✓ Go ${GO_VERSION} détecté${NC}"

# Vérifier la version de Go (minimum 1.21)
REQUIRED_VERSION="1.21"
if [ "$(printf '%s\n' "$REQUIRED_VERSION" "$GO_VERSION" | sort -V | head -n1)" != "$REQUIRED_VERSION" ]; then
    echo -e "${RED}❌ Go 1.21 ou supérieur est requis${NC}"
    exit 1
fi

# Télécharger les dépendances
echo -e "\n${YELLOW}📦 Téléchargement des dépendances...${NC}"
go mod download
if [ $? -eq 0 ]; then
    echo -e "${GREEN}✓ Dépendances téléchargées${NC}"
else
    echo -e "${RED}❌ Échec du téléchargement des dépendances${NC}"
    exit 1
fi

# Compiler le projet
echo -e "\n${YELLOW}🔨 Compilation de go-scaffold...${NC}"
go build -o go-scaffold main.go
if [ $? -eq 0 ]; then
    echo -e "${GREEN}✓ Compilation réussie${NC}"
else
    echo -e "${RED}❌ Échec de la compilation${NC}"
    exit 1
fi

# Installation globale (optionnelle)
echo -e "\n${YELLOW}Voulez-vous installer go-scaffold globalement ? (y/n)${NC}"
read -r response

if [[ "$response" =~ ^([yY][eE][sS]|[yY])$ ]]; then
    INSTALL_PATH="/usr/local/bin"
    
    if [ ! -w "$INSTALL_PATH" ]; then
        echo -e "${YELLOW}⚠️  Privilèges administrateur requis pour l'installation globale${NC}"
        sudo mv go-scaffold "$INSTALL_PATH/"
    else
        mv go-scaffold "$INSTALL_PATH/"
    fi
    
    if [ $? -eq 0 ]; then
        echo -e "${GREEN}✓ go-scaffold installé dans $INSTALL_PATH${NC}"
        echo -e "${GREEN}✓ Vous pouvez maintenant utiliser 'go-scaffold' depuis n'importe où${NC}"
    else
        echo -e "${RED}❌ Échec de l'installation globale${NC}"
        exit 1
    fi
else
    echo -e "${YELLOW}ℹ️  go-scaffold n'a pas été installé globalement${NC}"
    echo -e "${YELLOW}ℹ️  Vous pouvez l'utiliser avec ./go-scaffold dans ce répertoire${NC}"
    echo -e "${YELLOW}ℹ️  Pour installer globalement plus tard, exécutez: sudo make install${NC}"
fi

# Afficher les informations d'utilisation
echo -e "\n${GREEN}╔══════════════════════════════════════════════╗${NC}"
echo -e "${GREEN}║  ✓ Installation terminée avec succès !      ║${NC}"
echo -e "${GREEN}╚══════════════════════════════════════════════╝${NC}"

echo -e "\n${YELLOW}📚 Pour commencer:${NC}"
echo -e "   1. Créer un nouveau projet:"
echo -e "      ${GREEN}go-scaffold init mon-projet${NC}"
echo -e ""
echo -e "   2. Créer un schéma:"
echo -e "      ${GREEN}cd mon-projet${NC}"
echo -e "      ${GREEN}go-scaffold make:schema user${NC}"
echo -e ""
echo -e "   3. Générer le code:"
echo -e "      ${GREEN}go-scaffold generate database/schemas/user.yaml${NC}"
echo -e ""
echo -e "   4. Lancer l'application:"
echo -e "      ${GREEN}go run main.go${NC}"
echo -e ""
echo -e "${YELLOW}📖 Pour plus d'informations:${NC}"
echo -e "   ${GREEN}go-scaffold --help${NC}"
echo -e "   Consultez le README.md et QUICKSTART.md"
echo -e ""
echo -e "${GREEN}Bon codage ! 🚀${NC}"
