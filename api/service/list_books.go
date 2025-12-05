package service

import (
	"context"

	"github.com/bufbuild/connect-go"
	pb "github.com/icchy-san/bookstore-app/api/gen/go/bookstore/v1"
)

func (b *bookstoreService) ListBooks(ctx context.Context, req *connect.Request[pb.ListBooksRequest]) (*connect.Response[pb.ListBooksResponse], error) {
	res := connect.NewResponse(&pb.ListBooksResponse{
		Books: []*pb.Book{
			{
				Id:     "1",
				Name:   "Effective Go",
				Author: "Alan A. A. Donovan & Brian W. Kernighan",
			},
			{
				Id:     "2",
				Name:   "Learning React Native",
				Author: "Bonnie Eisenman",
			},
			{
				Id:     "3",
				Name:   "Learning gRPC",
				Author: "gRPC Master",
			},
			{
				Id:     "4",
				Name:   "Designing Data-Intensive Applications",
				Author: "Martin Kleppmann",
			},
			{
				Id:     "5",
				Name:   "Clean Code",
				Author: "Robert C. Martin",
			},
			{
				Id:     "6",
				Name:   "The Pragmatic Programmer",
				Author: "Andrew Hunt & David Thomas",
			},
			{
				Id:     "7",
				Name:   "Introduction to Algorithms",
				Author: "Thomas H. Cormen",
			},
			{
				Id:     "8",
				Name:   "Design Patterns",
				Author: "Erich Gamma",
			},
			{
				Id:     "9",
				Name:   "Refactoring",
				Author: "Martin Fowler",
			},
			{
				Id:     "10",
				Name:   "You Don't Know JS",
				Author: "Kyle Simpson",
			},
		},
	})
	return res, nil
}
