module github.com/c-wire/fastrtb/examples/go

go 1.22.4

replace github.com/c-wire/fastrtb => ../../go/fastrtb

require (
	github.com/c-wire/fastrtb v0.0.0-00010101000000-000000000000
	github.com/google/uuid v1.6.0
	google.golang.org/grpc v1.71.0
)

require (
	golang.org/x/net v0.34.0 // indirect
	golang.org/x/sys v0.29.0 // indirect
	golang.org/x/text v0.21.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250115164207-1a7da9e5054f // indirect
	google.golang.org/protobuf v1.36.5 // indirect
)
