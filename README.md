# Jump CLI

## À propos de Jump

Jump est un gestionnaire de répertoires flexible et rapide développé en Go. Il offre une interface en ligne de commande pour ajouter et naviguer facilement entre différents répertoires.

## Installation

(Instructions d'installation pour votre application.)

## Utilisation

### Commandes disponibles

- **Ajouter un répertoire** :
  - `jump add [nom] [chemin]` : Ajoute un nouveau répertoire à la liste.
  - Utiliser `jump add .` : Ajoute le répertoire courant.

- **Changer de répertoire** :
  - `jump to [nom]` : Change le répertoire courant en celui spécifié.
  - Avant d'utiliser la commande jump to il faut s'assurer qu'on a bien ajouter le répertoire

### Exemples

- Ajouter un répertoire nommé "projet" avec son chemin :

### Exemples

- Ajouter un répertoire nommé "projet" avec son chemin :

    jump add projet /chemin/vers/projet

- Ajouter le répertoire courant :

jump add .

- Aller au répertoire "projet" :

jump to projet

## Configuration

- **Fichier de configuration** :
- `jump.json` est stocké dans le répertoire personnel de l'utilisateur.
- Contient la liste des répertoires ajoutés.

## Contribution

(Instructions pour contribuer au projet.)

## Licence

(Spécifiez ici la licence sous laquelle votre projet est distribué, par exemple, MIT, GPL, etc.)

## Contact

(Informations de contact pour les utilisateurs qui souhaitent vous atteindre.)
