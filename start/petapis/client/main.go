package main

import (
	"context"
	"log"

	petv1 "go.buf.build/grpc/go/karlmutch/petapis/pet/v1"
	"google.golang.org/grpc"
)

func main() {
	connectTo := "127.0.0.1:8080"
	conn, err := grpc.Dial(connectTo, grpc.WithBlock(), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect to PetStoreService on %s due to %s", connectTo, err.Error())
	}
	log.Println("Connected to", connectTo)

	petStore := petv1.NewPetStoreServiceClient(conn)
	if _, err := petStore.PutPet(context.Background(), &petv1.PutPetRequest{
		PetType: petv1.PetType_PET_TYPE_SNAKE,
		Name:    "Ekans",
	}); err != nil {
		log.Fatalf("failed to PutPet %s", err.Error())
	}
}
