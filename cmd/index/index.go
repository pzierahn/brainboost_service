package main

import (
	"context"
	"github.com/pzierahn/braingain/database"
	"github.com/pzierahn/braingain/index"
	"log"
)

const (
	baseDir    = "/Users/patrick/patrick.zierahn@gmail.com - Google Drive/My Drive/KIT/2023-SS/DeSys/"
	collection = "DeSys"
)

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	ctx := context.Background()
	conn, err := database.Connect(ctx, "postgresql://postgres:postgres@localhost:5432")
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	//conn.DeleteFile(ctx, collection, "sigma.pdf")

	source := index.NewIndex(conn, collection)

	//source.Files(baseDir)
	source.File(ctx, baseDir+"/Further Readings/IPTPS2002.pdf")
	//source.File(ctx, baseDir+"/Further Readings/358527.358537.pdf")
	//source.File(ctx, baseDir+"/Further Readings/shared_rsa.pdf")
	//source.File(ctx, baseDir+"/Further Readings/sigma.pdf")

	//source.File(ctx, baseDir+"/Further Readings/2102.08325.pdf")
	//source.File(ctx, baseDir+"/Further Readings/3558535.3559789.pdf")
	//source.File(ctx, baseDir+"/Further Readings/3149.214121.pdf")
	//source.File(ctx, baseDir+"/Further Readings/1-s2.0-089054018790054X-main.pdf")
	//source.File(ctx, baseDir+"/Further Readings/3465084.3467905.pdf")
	//source.File(ctx, baseDir+"/Further Readings/176429260X.pdf")
	//source.File(ctx, baseDir+"/Further Readings/Brewer_podc_keynote_2000.pdf")
	//source.File(ctx, baseDir+"/Further Readings/cap.pdf")
	//source.File(ctx, baseDir+"/Further Readings/Harvest_yield_and_scalable_tolerant_systems.pdf")
	//source.File(ctx, baseDir+"/Further Readings/Kademlia.pdf")
	//source.File(ctx, baseDir+"/Further Readings/Lamport - Time, Clocks, and the Ordering of Events in a Distributed System.pdf")
	//source.File(ctx, baseDir+"/Further Readings/RR-7687.pdf")
	//source.File(ctx, baseDir+"/Further Readings/The Byzantine Generals Problem.pdf")
	//source.File(ctx, baseDir+"/Further Readings/The Sybil Attack.pdf")
	//source.File(ctx, baseDir+"/Further Readings/Efficient_Byzantine_Fault-Tolerance.pdf")

	//conn.DeleteFile(ctx, collection, "2305.06123.pdf")

	//count, err := conn.Count(ctx, collection)
	//if err != nil {
	//	log.Fatalf("could not count points: %v", err)
	//}
	//
	//log.Printf("Documents in db: %v\n", count.Result.Count)
}
