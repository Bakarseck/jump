#!/bin/bash

# Alias Git
alias ga='git add'
alias gc='git commit -m'
alias gcl='git clone'
alias gs='git status'
alias gp='git push'
alias gpl='git pull'
alias gb='git branch'
alias gco='git checkout'
alias gm='git merge'
alias gd='git diff'
alias gl='git log'
alias gsh='git stash'
alias grh='git reset --hard'
alias grs='git reset --soft'
alias gcp='git cherry-pick'

# Alias système
alias cls='clear'
alias ll='ls -la'
alias df='df -h' # Affiche l'utilisation du disque en format lisible
alias free='free -m' # Affiche la mémoire libre et utilisée en MB

# Alias Docker
alias d='docker'
alias dc='docker-compose'
alias di='docker images'
alias dps='docker ps'
alias dpsa='docker ps -a'
alias drm='docker rm'
alias drmi='docker rmi'
alias drun='docker run'
alias dex='docker exec -it'
alias db='docker build'
alias dlogs='docker logs'
alias dv='docker volume'
alias dn='docker network'

# Ajouts supplémentaires pour Docker
alias dprune='docker system prune -a' # Supprime tous les conteneurs, réseaux, et images non utilisés
alias dstopall='docker stop $(docker ps -aq)' # Arrête tous les conteneurs en cours
alias drmiall='docker rmi $(docker images -q)' # Supprime toutes les images Docker

# Utilisation de 'source' pour recharger le fichier .bashrc ou .zshrc après modification
# alias reload='source ~/.bashrc' # Si vous utilisez bash
alias reload='source ~/.zshrc' # Si vous utilisez zsh
alias config='code ~/.zshrc'
