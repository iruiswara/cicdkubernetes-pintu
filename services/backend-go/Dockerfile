FROM golang:alpine


WORKDIR /app


COPY . .


RUN go build -o goService .


EXPOSE 8080


CMD ["./goService"]
