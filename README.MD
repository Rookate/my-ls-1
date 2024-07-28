# `my-ls-project`

`my-ls-project` est une implémentation personnalisée de la commande `ls` en Go. Ce projet permet de lister les fichiers et répertoires avec diverses options de tri et d'affichage, similaires à celles que l'on trouve dans la commande `ls` de Unix/Linux.

## Fonctionnalités

- Liste les fichiers et répertoires dans un répertoire spécifié.
- Affiche les fichiers cachés lorsqu'`-a` est activé.
- Supporte le mode long `-l` pour afficher des détails supplémentaires sur les fichiers.
- Tri des fichiers par date de modification avec l'option `-t`.
- Tri des fichiers par ordre alphabétique de manière croissante ou décroissant en fonction de `r`.
- Support de l'affichage des fichiers dans plusieurs colonnes.
- Support du mode récursif `-R` pour afficher les contenus des sous-répertoires.

## Installation

Pour utiliser ce projet, clonez le dépôt et compilez le programme :

```bash
git clone https://zone01normandie.org/git/gkopoin/my-ls-1.git
cd my-ls-project
```

## Utilisation

```bash
go run . ls [Options] [Path...]
```

### Options

- `-a` : Affiche tous les fichiers, y compris les fichiers cachés (ceux qui commencent par un point).
- `-l` : Affiche les détails des fichiers dans le format long.
- `-t` : Trie les fichiers par date de modification, les plus récents en premier.
- `-R` : Affiche les fichiers de manière récursive dans les sous-répertoires.
- `-a` : Affiche le répertoire actuel (`.`) et le répertoire parent (`..`).

### Exemples

1. **Lister les fichiers dans le répertoire actuel** :

    ```bash
    go run . ls
    ```

2. **Lister les fichiers avec détails longs** :

    ```bash
    go run . ls -l
    ```

3. **Afficher tous les fichiers, y compris les cachés** :

    ```bash
    go run . ls -a
    ```

4. **Lister les fichiers triés par date de modification** :

    ```bash
    go run . ls -t
    ```

5. **Lister les fichiers de manière récursive** :

    ```bash
    go run . ls -R
    ```

6. **Combiner les options** :

    ```bash
    go run . ls -alRrt
    ```
---