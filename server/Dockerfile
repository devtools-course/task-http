FROM golang:1.16-alpine AS build

WORKDIR /home/go

COPY main.go go.mod ./
RUN go get gopkg.in/yaml.v2
RUN CGO_ENABLED=0 GOOS=linux go build -o /server main.go

FROM scratch
LABEL authors="lBAH, Kalinin Iwan <koefic.cien@gmail.com>"

ENV ADDR=1111

COPY --from=build /server /bin/server
COPY answer.txt /answer.txt

EXPOSE 1111
ENTRYPOINT ["/bin/server"]
