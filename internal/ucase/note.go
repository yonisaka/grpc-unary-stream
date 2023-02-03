package ucase

import (
	"fmt"
	"grpc-unary-stream/files/pb"
	"grpc-unary-stream/internal/repositories"
	"sync"
	"time"
)

type note struct {
	repo repositories.NoteRepository
	pb.UnimplementedNoteServiceServer
}

func NewNoteUsecase(repo repositories.NoteRepository) pb.NoteServiceServer {
	return &note{repo: repo}
}

func (u *note) FindLimit(req *pb.LimitRequest, stream pb.NoteService_FindLimitServer) error {
	if req.Limit == 0 {
		req.Limit = 5
	}

	var wg sync.WaitGroup
	for i := 0; i < int(req.Limit); i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			time.Sleep(time.Duration(i) * time.Second)

			id := int64(i + 1)
			ctx := stream.Context()
			dr, err := u.repo.FindById(ctx, id)
			if err != nil {
				fmt.Println(err)
				return
			}

			if err = stream.Send(&pb.NoteResponse{
				Id:          int64(dr.ID),
				Title:       dr.Title,
				Description: dr.Description,
			}); err != nil {
				fmt.Println(err)
				return
			}
		}(i)
	}
	wg.Wait()

	return nil
}

func (u *note) FindById(stream pb.NoteService_FindByIdServer) error {
	ctx := stream.Context()
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		req, err := stream.Recv()
		if err != nil {
			return err
		}

		dr, err := u.repo.FindById(ctx, req.Id)
		if err != nil {
			return err
		}

		if err = stream.Send(&pb.NoteResponse{
			Id:          int64(dr.ID),
			Title:       dr.Title,
			Description: dr.Description,
		}); err != nil {
			return err
		}
	}
}
