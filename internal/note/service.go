package note

import (
	"context"

	authpb "github.com/bekhuli/notes-service/api/note"
	pb "github.com/bekhuli/notes-service/api/note"
)

type Service struct {
	authpb.UnimplementedNoteServiceServer
	repo Repository
}

func NewNoteService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) CreateNote(ctx context.Context, req *pb.CreateNoteRequest) (*pb.NoteResponse, error) {
	id, err := s.repo.CreateNote(ctx, req.UserId, req.Title, req.Content)
	if err != nil {
		return nil, err
	}

	return &pb.NoteResponse{Id: id, UserId: req.UserId, Title: req.Title, Content: req.Content}, nil
}

func (s *Service) GetNotes(req *pb.UserIDRequest, stream pb.NoteService_GetNotesServer) error {
	ctx := stream.Context()

	notes, err := s.repo.GetNotes(ctx, req.UserId)
	if err != nil {
		return err
	}

	for _, n := range notes {
		resp := &pb.NoteResponse{
			Id:      n.ID,
			UserId:  n.UserID,
			Title:   n.Title,
			Content: n.Content,
		}
		if err = stream.Send(resp); err != nil {
			return err
		}
	}

	return nil
}

func (s *Service) UpdateNote(ctx context.Context, req *pb.UpdateNoteRequest) (*pb.NoteResponse, error) {
	note, err := s.repo.UpdateNote(ctx, req.NoteId, req.Title, req.Content)
	if err != nil {
		return nil, err
	}

	return &pb.NoteResponse{
		Id:      note.ID,
		UserId:  note.UserID,
		Title:   note.Title,
		Content: note.Content,
	}, nil
}

func (s *Service) DeleteNote(ctx context.Context, req *pb.NoteIDRequest) (*pb.Empty, error) {
	err := s.repo.DeleteNote(ctx, req.NoteId)
	if err != nil {
		return nil, err
	}

	return &pb.Empty{}, nil
}
