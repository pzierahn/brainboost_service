package main

import (
	"context"
	"github.com/google/uuid"
	"github.com/pzierahn/braingain/database"
	"github.com/pzierahn/braingain/database_pg"
	"log"
	"os"
)

type QdrantExport struct {
	Id        string
	Embedding []float32
	Filename  string
	Page      int
	Text      string
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	ctx := context.Background()

	old, err := database.Connect("localhost:6334")
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer func() { _ = old.Close() }()

	results, err := old.GetAll(ctx, "DeSys")
	if err != nil {
		log.Fatal(err)
	}

	exports := make([]QdrantExport, len(results))

	log.Printf("results=%v", len(results))
	for inx, result := range results {
		text := result.Payload["content"].GetStringValue()
		filename := result.Payload["filename"].GetStringValue()
		page := result.Payload["page"].GetIntegerValue()

		// log.Printf("%v --> %d", filename, page)

		exports[inx] = QdrantExport{
			Id:        result.Id.GetUuid(),
			Embedding: result.Vectors.GetVector().Data,
			Filename:  filename,
			Page:      int(page),
			Text:      text,
		}
	}

	//byt, err := json.MarshalIndent(exports, "", "  ")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//err = os.WriteFile("qdrant_export.json", byt, 0644)
	//if err != nil {
	//	log.Fatal(err)
	//}

	//db, err := database_pg.Connect(ctx, os.Getenv("NEON_DB"))
	db, err := database_pg.Connect(ctx, os.Getenv("SUPABASE_DB"))
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	_, err = db.Upsert(ctx, database_pg.Point{
		Source:    uuid.New(),
		Page:      0,
		Text:      "xxx",
		Embedding: make([]float32, 1536),
	})
	if err != nil {
		log.Fatal(err)
	}

	//sources := make(map[string]uuid.UUID)
	//for _, result := range results {
	//	filename := result.Payload["filename"].GetStringValue()
	//	if _, ok := sources[filename]; ok {
	//		continue
	//	}
	//
	//	source := database_pg.Document{
	//		Filename: filename,
	//	}
	//
	//	id, err := db.CreateSource(ctx, source)
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//
	//	sources[filename] = id
	//	log.Printf("%v --> %v", filename, id)
	//}
	//
	//for _, result := range exports {
	//	filename := result.Filename
	//	sourceID := sources[filename]
	//
	//	point := database_pg.Point{
	//		Source:    sourceID,
	//		Page:      result.Page,
	//		Text:      result.Text,
	//		Embedding: result.Embedding,
	//	}
	//
	//	log.Printf("Upsert: %v --> %v", filename, point.Page)
	//	_, err := db.Upsert(ctx, point)
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//}
}
