FROM golang:1.9

ADD . /go/src/github.com/anabiozz/yotunheim/backend
ADD ./default-config.toml /go/bin/

RUN go get -u github.com/kardianos/govendor
RUN cd /go/src/github.com/anabiozz/yotunheim/backend govendor init && govendor add +external && govendor sync
RUN go install github.com/anabiozz/yotunheim/backend/main

EXPOSE 8888

CMD [ "/go/bin/main", "/go/bin/default-config.toml" ]

