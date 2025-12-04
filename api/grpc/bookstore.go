package grpc

// type bookstoreServer struct {
// 	svc service.BookstoreService
// 	pb.UnimplementedBookstoreServiceServer
// }

// func (b *bookstoreServer) mustEmbedUnimplementedBookstoreServiceServer() {
// 	panic("implement me")
// }

// func (b *bookstoreServer) ListBooks(ctx context.Context, req *pb.ListBooksRequest) (*pb.ListBooksResponse, error) {
// 	return b.svc.ListBooks(ctx, req)
// }

// func newBookstoreServer(bookstoreService service.BookstoreService) pb.BookstoreServiceServer {
// 	return &bookstoreServer{
// 		svc: bookstoreService,
// 	}
// }
