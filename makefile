run:
	go run server.go
run-client:
	go run client/client.go
run-gen:
	protoc calculatorpb/calculator.proto --go_out=.
