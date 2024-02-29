#!/bin/bash

# Installation de zsh
sudo apt-get update
sudo apt-get install zsh -y

# Changement du shell par défaut vers zsh
chsh -s $(which zsh)

# Installation d'Oh My Zsh
sh -c "$(curl -fsSL https://raw.githubusercontent.com/ohmyzsh/ohmyzsh/master/tools/install.sh)"


# Configuration du thème Zsh aléatoire
sed -i 's/ZSH_THEME=".*"/ZSH_THEME="rkj"/' ~/.zshrc

# Installation de zsh-syntax-highlighting
git clone https://github.com/zsh-users/zsh-syntax-highlighting.git ~/.oh-my-zsh/custom/plugins/zsh-syntax-highlighting

# Installation de zsh-autosuggestions
git clone https://github.com/zsh-users/zsh-autosuggestions.git ~/.oh-my-zsh/custom/plugins/zsh-autosuggestions

# Installation des plugins
plugins=(
    git
    docker
    zsh-autosuggestions
    zsh-syntax-highlighting
)

# Installation des plugins Oh My Zsh
sed -i 's/^plugins=(.*)/plugins=('"${plugins[*]}"')/' ~/.zshrc

# Installation de z (outil de navigation rapide)
git clone https://github.com/rupa/z.git ~/z
echo ". ~/z/z.sh" >> ~/.zshrc

# Mise en place de la complétion automatique pour zsh
echo "autoload -U compinit && compinit" >> ~/.zshrc

# Mise à jour du fichier de configuration
source ~/.zshrc

# Explication des outils installés
echo "Zsh est un shell amélioré qui offre de nombreuses fonctionnalités. Oh My Zsh est un gestionnaire de configuration pour zsh qui facilite la personnalisation. Les plugins git, docker, zsh-autosuggestions et zsh-syntax-highlighting ajoutent des fonctionnalités pour une meilleure utilisation de ces outils. 'z' est un utilitaire pour naviguer rapidement entre les répertoires. La complétion automatique est améliorée avec 'autoload -U compinit && compinit'." 

# Demande de redémarrer le terminal pour appliquer les changements
echo "Veuillez redémarrer votre terminal pour appliquer les changements."