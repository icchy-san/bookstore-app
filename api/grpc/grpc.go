package grpc

// type Config struct {
// 	Service service.BookstoreService
// }

// func NewGrpcServer(config *Config) (*grpc.Server, error) {
// 	opts := []grpc.ServerOption{}

// 	return newGrpcServer(config, opts...)
// }

// func newGrpcServer(config *Config, opts ...grpc.ServerOption) (*grpc.Server, error) {
// 	grpcServer := grpc.NewServer(opts...)

// 	bookstoreServer := newBookstoreServer(config.Service)
// 	pb.RegisterBookstoreServiceServer(grpcServer, bookstoreServer)

// 	return grpcServer, nil
// }
