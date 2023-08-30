package main

import (
	"context"
	"encoding/json"
	"github.com/pzierahn/braingain/braingain"
	"github.com/pzierahn/braingain/database"
	"github.com/sashabaranov/go-openai"
	"log"
	"os"
	"sort"
)

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	conn, err := database.Connect("localhost:6334")
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer func() { _ = conn.Close() }()

	ctx := context.Background()

	//search := "Explain the Practical Byzantine Fault Tolerance (PBFT) algorithm in detail"
	//search := "What is the difference between the Practical Byzantine Fault Tolerance (PBFT) algorithm and the DAG-Rider algorithm?"
	//search := "How does the DAG-Rider algorithm work?"
	//search := "Why does the DAG-Rider algorithm needs waves?"
	//search := "Explain the DAG-Rider algorithm in detail"
	//search := "Which properties of an Operation-Based CRDT have to be shown?"
	//search := "How does a Sybil Attack with Distributed Secret Sharing work?"
	//search := "How is consensus archived in total order broadcasts?"
	//search := "What is the differance between consistency and consensus?"
	//search := "How does the TEE-Rider algorithm work in detail?"
	//search := "What is the TEE-Rider algorithm and how does it work in detail?"
	//search := "What is a Byzantine Atomic Broadcast?"
	//search := "What is the difference between a Byzantine Atomic Broadcast and a Byzantine Reliable Broadcast?"
	//search := "How can total order be archived with a Byzantine Atomic Broadcast?"
	//search := "What is the definition of Byzantine Broadcast Channel? What is the difference between RB-Agreement and RB-Validity?"
	//search := "Explain the TEE-based Reliable Broadcast setting"
	//search := "Explain the TEE-based Reliable Broadcast in detail"
	//search := "Explain the TEE-Rider algorithm in detail"
	//search := "What are the failure models?"
	//search := "How can correct processors verify that a number was generated only by a Unique Sequential Identifier Generator (USIG)?"
	//search := "How is a message signed with USIG?"
	//search := "What is the difference between DAG-Rider and TEE-Rider algorithms?"
	//search := "What is MinBFT?"
	//search := "In how far does the TEE-based Reliable Broadcast differ from other reliable broadcasts?"
	//search := "How is partial synchrony, synchrony and asynchrony defined?"
	//search := "What is the difference between partial synchrony and asynchrony defined?"
	//search := "What is the advantage of knowing that \"Δ is arbitrary but fix and unknown\" in partial synchrony?"
	//search := "How does the TEE-Rider algorithm uses four waves?"
	//search := "What is the advantage of partial synchrony against Asynchrony?"
	//search := "Why can partial synchrony guarantee safety and liveness?"
	//search := "Explain the flp impossibility in detail"
	//search := "Explain the impossibility of Distributed Consensus with One Faulty Process"
	//search := "How can a Total Order Broadcast be archived?"
	//search := "How can consensus be derived in TEE-Rider?"
	//search := "How is a Total Order in TEE-Rider archived?"
	//search := "What makes a distributed system a decentralized system? → The Meaning of “Decentralization”"
	//search := "What are the reasons for decentralization?"
	//search := "What is the CAP Theorem? Why can only two of the three properties be archived?"
	//search := "What is the CAP Theorem? Why can only two of the three properties be archived? Give an examples"
	//search := "What is the differance between Consistency and Partition Tolerance in CAP?"
	//search := "Why is there no terminating algorithm for consensus in the asynchronous model?"
	//search := "Why it is impossible for processes in an asynchronous system to unanimously agree on a consensus value if even a single process could fail?"
	//search := "Explain the FLP Impossibility in detail"
	//search := "What is the difference between synchrony, partial synchrony and asynchrony? Is the delta in partial synchrony approximated?"
	//search := "What are the implications if the delta in partial synchrony approximated wrongly?"
	//search := "Which properties need to be shown to proof eventual consistency and strong eventual consistency?"
	//search := "How does leader election work in PBFT?"
	//search := "Why are 2f+1 needed to deal with f faulty processes?"
	//search := "How can f be determined in DAG-Rider?"
	//search := "What is deterministic threshold signatures? How do they work?"
	//search := "How does deterministic threshold signing work in DAG-Rider?"
	//search := "How does deterministic threshold signing work?"
	//search := "How does Distributed Key Generation (DKG) for RSA work?"
	//search := "How is safety and liveliness archived in DAG-Rider?"
	//search := "How does leader election in DAG-Rider work?"
	//search := "Why are 4 rounds a wave in DAG-Rider?"
	//search := "How can safety and liveness be proofen in PBFT?"
	//search := "Explain the three phases of Bracha’s Reliable Broadcast"
	//search := "What is strong clock consistency?"
	//search := "Explain why a vector clock is a Conflict-Free Replicated Data Type in itself. Is it state-based or operation-based?"
	//search := "Discuss whether you can build a state-based Conflict-Free Replicated Data type on top of a vector clock. How can you utilize the vector information in the design?"
	//search := "Explain termination, agreement and validity in consensus"
	//search := "What is the definition of consensus? What properties does it have?"
	//search := "Why do you need n > 3f processors to solve the Byzantine Generals Problem?"
	//search := "What is Nakamoto Consensus?"
	//search := "What legal aspects need to be considered in decentralized systems?"
	//search := "What is the Oral Messages Algorithm and how does it work?"
	//search := "What is the Byzantine generals' problem? What are the Communication Channels assumptions? Answer in bullet points"
	//search := "What is the difference between a the Oral Messages Algorithms and Bracha’s Reliable Broadcast?"
	//search := "How does Bracha’s Consensus Protocol work?"
	//search := "How does Bracha’s Consensus Protocol need n > 3f nodes? Why not n > 2f+1 nodes?"
	//search := "What is the difference between consistency and consensus?"
	//search := "Explain the k-neighborhood algorithm in detail"
	//search := "What is the problem with mutual exclusion in decentralized systems?"
	//search := "How does the Lamport Mutual Exclusion algorithm work?"
	//search := "Explain the required properties of the Lamport Mutual Exclusion algorithm"
	//search := "Explain Lamport Mutual Exclusion in detail"
	//search := "Explain assumptions of Lamport Mutual Exclusion"
	//search := "How does Lamport Mutual Exclusion handle concurrent requests?"
	//search := "How does Lamport Mutual Exclusion ensure total ordering?"
	//search := "Explain Distributed Algorithm by Ricart and Agrawala"
	//search := "Explain the Crash Fault-Tolerant Algorithm by Lin et al."
	//search := "Explain the Crash Fault-Tolerant Algorithm by Lin et al."
	//search := "Explain the Practical Distributed Mutual Exclusion Protocol in Dynamic Peer-to-Peer Systems in detail"
	//search := "How do Lamport Clocks guarantee a total order?"
	//search := "How are safety and liveness defined?"
	//search := "Discuss the safety and liveness guarantees of token-ring algorithms"
	//search := "How does OrbitDB work?"
	//search := "Does IPFS Pub-Sub mechanism has a way of dealing with byzantine nodes?"
	//search := "Explain Casper FFG"
	//search := "Explain LMD GHOST Fork Choice"
	//search := "Can forks happen with a Proof of Stake (PoS) blockchain system?"
	//search := "Explain the attestations process in Proof of Stake (PoS) blockchain systems"
	//search := "How does the data retrieval process work in IPFS?"
	//search := "How does IPFS handle byzantine nodes?"
	//search := "Explain sharding in detail"
	//search := "Explain a beacon chain in detail"
	//search := "How can Ethereum be scaled to handel more transactions?"

	// TODO:
	//search := "Explain Rollups in Ethereum in detail"
	search := "How exactly do sidechains work in blockchains?"
	//search := "What is the differance between side chains and payment channel networks?"
	//search := "What is the differance between gas and ether?"
	//search := "Explain zk-Rollups in detail. Explain how zero-knowledge proofs work in this context"

	token := os.Getenv("OPENAI_API_KEY")
	ai := openai.NewClient(token)

	chat := braingain.NewChat(conn, ai)
	//chat.Model = openai.GPT4

	response, err := chat.RAG(ctx, search)
	if err != nil {
		log.Fatalf("ChatCompletion error: %v\n", err)
	}

	sources := make(map[string][]int)
	for _, source := range response.Sources {
		sources[source.Filename] = append(sources[source.Filename], source.Page)
	}

	keys := make([]string, 0, len(sources))
	for k := range sources {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, filename := range keys {
		pages := sources[filename]
		log.Printf("%s --> %v\n", filename, pages)
	}

	byt, _ := json.MarshalIndent(response.Costs, "", "  ")
	log.Printf("Costs: %s\n", string(byt))

	log.Println(response.Completion)
	_ = os.WriteFile("output.txt", []byte(response.Completion), 0644)

	byt, _ = json.MarshalIndent(response.Sources, "", "  ")
	_ = os.WriteFile("sources.json", byt, 0644)
}
