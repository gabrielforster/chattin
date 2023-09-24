build:
	env GOOS=linux go build -ldflags="-s -w" -o bin/main cmd/main.go

protogen:
	protoc \
		--go_out=. \
		--go_opt=paths=source_relative \
		--go-grpc_out=. \
		--go-grpc_opt=paths=source_relative \
		./proto/auth.proto
