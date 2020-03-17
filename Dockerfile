FROM golang AS build

WORKDIR /

ADD . .

RUN CGO_ENABLED=0 go build -v -o /velas-sphere .

FROM alpine as velas-sphere-requester

COPY --from=build /velas-sphere /velas-sphere

ENTRYPOINT ["/velas-sphere", "requester"]

FROM alpine as velas-sphere-provider

COPY --from=build /velas-sphere /velas-sphere

ENTRYPOINT ["/velas-sphere", "provider"]

EXPOSE 8082/tcp

FROM alpine as velas-sphere-plugin

COPY --from=build /velas-sphere /velas-sphere

ENTRYPOINT ["/velas-sphere", "plugin"]

EXPOSE 8082/tcp