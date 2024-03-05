#!/bin/bash

# Emplacement du fichier JSON pour stocker les répertoires
homeDir=$(eval echo ~$USER)
pathJson="$homeDir/jump.json"

# Fonction pour ajouter un répertoire
add_directory() {
    dir=""
    path=""

    if [ "$1" == "." ]; then
        dir=$(basename "$PWD")
        path="$PWD"
    elif [ $# -eq 2 ]; then
        dir="$1"
        path="$2"
    else
        echo "Usage: jump add [name] [path] or jump add ."
        return
    fi

    if [ ! -f "$pathJson" ]; then
        touch "$pathJson"
        echo "[]" > "$pathJson"
    fi

    # Ajouter le nouveau répertoire au fichier JSON
    jq --arg dir "$dir" --arg path "$path" '. += [{"Dir": $dir, "Path": $path}]' "$pathJson" > tmp.$$.json && mv tmp.$$.json "$pathJson"
}

# Fonction pour changer de répertoire
change_directory() {
    if [ $# -ne 1 ]; then
        echo "Usage: jump to [name]"
        return
    fi

    dir="$1"
    path=$(jq -r --arg dir "$dir" '.[] | select(.Dir==$dir) | .Path' "$pathJson")

    if [ -n "$path" ] && [ -d "$path" ]; then
        cd "$path" || return
    else
        echo "Répertoire non trouvé."
    fi
}

# Fonction pour cloner un dépôt Git
clone_repository() {
    if [ $# -ne 1 ]; then
        echo "Usage: jump clone [repository-url]"
        return
    fi

    repositoryURL="$1"
    git clone "$repositoryURL"
}

# Analyse des arguments de la ligne de commande
case "$1" in
    add)
        add_directory "${@:2}"
        ;;
    to)
        change_directory "${@:2}"
        ;;
    clone)
        clone_repository "${@:2}"
        ;;
    *)
        echo "Commandes disponibles :"
        echo "  jump add [name] [path] - Ajouter un répertoire"
        echo "  jump add . - Ajouter le répertoire courant"
        echo "  jump to [name] - Aller au répertoire"
        echo "  jump clone [repository-url] - Cloner un dépôt Git"
        ;;
esac
