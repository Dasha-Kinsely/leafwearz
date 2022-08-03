## In every directory, if default.go is present, it is always used to congregate the entry points to every other files. 

# common workflow: routers -> middlewares -> handlers -> 1st validator call for dto -> services -> repos or pkgs -> entities -> 2nd validator call for service results -> serializers