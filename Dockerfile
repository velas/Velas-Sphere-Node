FROM golang AS build

WORKDIR /

ADD . .

RUN CGO_ENABLED=0 go build -v -o /velas-sphere .

FROM alpine as velas-sphere-node

COPY --from=build /velas-sphere /velas-sphere
COPY --from=build /config.json /config.json

ENTRYPOINT ["/velas-sphere", "node"]

# provider service
EXPOSE 8081/tcp

# storage service
EXPOSE 8083/tcp

# rest api
EXPOSE 3000/tcp

FROM alpine as velas-sphere-plugin

COPY --from=build /velas-sphere /velas-sphere

ENTRYPOINT ["/velas-sphere", "plugin"]

EXPOSE 8082/tcp
