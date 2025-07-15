package main

import (
	"log"
	"net"

	notepb "github.com/bekhuli/notes-service/api/note"
	"github.com/bekhuli/notes-service/internal/note"
	"github.com/bekhuli/notes-service/pkg/db"

	"google.golang.org/grpc"
)

func main() {
	db.Connect()
	defer db.DB.Close()

	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	noteRepository := note.NewNoteRepository(db.DB)
	noteService := note.NewNoteService(noteRepository)
	notepb.RegisterNoteServiceServer(grpcServer, noteService)

	log.Println("NoteService listening on :50052")
	if err = grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
