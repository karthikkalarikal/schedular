run: api-gateway employee

api-gateway:
		go run api-gateway/main.go

employee:
		go run employee-management/main.go

swag: swagfmt swaginit

swaginit:
		swag init -g api-gateway/main.go -o ./api-gateway/docs

swagfmt:
		swag fmt api-gateway 

grpc: grpc-api-gateway grpc-employee

grpc-api-gateway1:
	DEBUG=1 protoc -Iapi-gateway/pkg/rpc/employee --go_out=api-gateway \
	--go_opt=module=github.com/karthikkalarikal/api-gateway \
	--go-grpc_out=api-gateway \
	--go-grpc_opt=module=github.com/karthikkalarikal/api-gateway \
	api-gateway/pkg/rpc/employee/proto/employee.proto

grpc-api-gateway:
	protoc api-gateway/pkg/clients/employee/proto/employee.proto --go_out=./ --go-grpc_out=./

grpc-employee:
	protoc employee-management/pkg/pb/employee.proto --go_out=./ --go-grpc_out=./