run: api-gateway

api-gateway:
		go run api-gateway/main.go

swag: swagfmt swaginit

swaginit:
		swag init -g api-gateway/main.go -o ./api-gateway/docs

swagfmt:
		swag fmt api-gateway

grpc: grpc-api-gateway

grpc-api-gateway:
	protoc -Iapi-gateway/pkg/rpc/employee --go_out=api-gateway \
	--go_opt=module=github.com/karthikkalarikal/api-gateway \
	--go-grpc_out=api-gateway \
	--go-grpc_opt=module=github.com/karthikkalarikal/api-gateway \
	api-gateway/pkg/rpc/employee/employee.proto
