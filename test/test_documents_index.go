package test

import (
	"bytes"
	_ "embed"
	"errors"
	"fmt"
	"github.com/google/uuid"
	pb "github.com/pzierahn/brainboost/proto"
	"io"
	"log"
)

//go:embed test.pdf
var testPdf []byte

func (setup *Setup) DocumentsIndex() {

	ctx, userId := setup.createRandomSignIn()
	defer setup.DeleteUser(userId)

	collection, err := setup.collections.Create(ctx, &pb.Collection{Name: "test"})
	if err != nil {
		log.Fatal(err)
	}

	docId := uuid.NewString()
	path := fmt.Sprintf("%s/%s/%s.pdf", userId, collection.Id, docId)

	resp := setup.storage.UploadFile(bucket, path, bytes.NewReader(testPdf))
	if resp.Error != "" {
		log.Fatalf("upload failed: %v", resp.Error)
	}

	doc := &pb.Document{
		Id:           docId,
		CollectionId: collection.Id,
		Filename:     "test.pdf",
		Path:         path,
	}

	stream, err := setup.documents.Index(ctx, doc)
	if err != nil {
		log.Fatal(err)
	}

	for {
		update, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			break
		} else if err != nil {
			log.Fatal(err)
		}

		if update.TotalPages <= 0 {
			log.Fatalf("invalid total pages: %v", update.TotalPages)
		}

		log.Printf("update: %v", update)
	}

	log.Printf("Uploaded and processed: %v", doc)
}