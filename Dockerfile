# Utilisation de golang comme image de base
# Le GOPATH par d�faut de cette image est /go.
FROM golang
 
# Copie des sources de notre projet dans le container,
# dans notre cas le main.go
ADD . /go/src/managIncident
 
# Lancement de la compilation avec go install
RUN go get -u github.com/astaxie/beego
RUN go get -u github.com/beego/bee
RUN go get -u github.com/garyburd/redigo/redis
RUN go get -u gopkg.in/gomail.v1
RUN go install managIncident
 
# D�finissions du point d'entr� de notre programme compil�
ENTRYPOINT /go/bin/managIncident
 
# Le port sur lequel notre serveur �coute
EXPOSE 8080