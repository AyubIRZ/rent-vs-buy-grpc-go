package rent_vs_buy_grpc_go

// Following directives generate respective .pb.go files. Add another go:generate comment if you needed new gRPC
// services or new API versions.

//go:generate protoc --go_out=plugins=grpc:. internal/protos/v1/breakeven/breakeven.proto
