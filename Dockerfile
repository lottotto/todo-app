FROM golang:1.14 as builder
WORKDIR /my-todo-app
COPY . .
RUN GOOS=linux GOARCH=amd64 go build main.go
#This is the final stage, and we copy artifacts from "builder"
FROM gcr.io/distroless/base
COPY --from=builder /my-todo-app/main /bin/main
EXPOSE 1323
ENTRYPOINT ["/bin/main"]