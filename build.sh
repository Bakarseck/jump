#!/bin/bash

# Définir les plateformes cibles
platforms=("linux/amd64" "linux/arm64" "darwin/amd64" "darwin/arm64" "windows/amd64")

# Nom du programme
APP_NAME="jump"

# Dossier de sortie
OUTPUT_DIR="builds"
mkdir -p $OUTPUT_DIR

echo "📦 Compilation du projet Go pour plusieurs plateformes..."

# Boucle sur toutes les plateformes
for platform in "${platforms[@]}"
do
    GOOS=${platform%/*}    # Extraire OS (ex: linux, darwin, windows)
    GOARCH=${platform#*/}  # Extraire Architecture (ex: amd64, arm64)
    OUTPUT_NAME="${OUTPUT_DIR}/${APP_NAME}-${GOOS}-${GOARCH}"
    
    # Ajouter .exe pour Windows
    if [ "$GOOS" == "windows" ]; then
        OUTPUT_NAME+=".exe"
    fi

    echo "🚀 Compilation pour $GOOS/$GOARCH..."
    GOOS=$GOOS GOARCH=$GOARCH go build -o $OUTPUT_NAME

    if [ $? -ne 0 ]; then
        echo "❌ Erreur de compilation pour $GOOS/$GOARCH"
    else
        echo "✅ Build réussi : $OUTPUT_NAME"
    fi
done

echo "🎉 Tous les builds sont terminés !"
