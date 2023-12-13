https://medium.com/@bnprashanth256/reading-configuration-files-and-environment-variables-in-go-golang-c2607f912b63

https://dev.to/koddr/let-s-write-config-for-your-golang-web-app-on-right-way-yaml-5ggp


https://github.com/rahmanfadhil/gin-bookstore/blob/master/main.go

Creation avec 2 frameworks => Mux pour l'exposition api et gorm pour la base de données


==========================-
Parcours utilisateurs
    Table principale (admin) :
        -> Création d'un tournoi
            -> Création des catégories
                -> Ouverture des inscriptions d'une catégorie
                -> Fermeture des inscriptions de la catégorie
                -> Création des tableaux
            -> Fermeture de la catégorie (Résumé dispo)
        -> Fermeture du tournoi

    Table de pesé :
        -> Inscrire un judoka


==========================-
Init env de dev
    go get github.com/gin-gonic/gin
    go get gorm.io/gorm
    go get gorm.io/driver/sqlite
    go get gorm.io/datatypes


==========================-
Lancement du dev ()
    go run . 
ou
    find -name '*.go' | entr -rd go run .  # Permet de relancer en cas de modif