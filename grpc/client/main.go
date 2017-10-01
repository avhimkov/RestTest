package main

import (
	"io"
	"fmt"
	"google.golang.org/grpc"
	"golang.org/x/net/context"
	prsn "github.com/agiratech/golang-rpc/person"
)

const (
	addres = "localhost:3333"
)

func createPerson(client prsn.PersonClient, person *prsn.PersonRequest) {
	resp, err := client.CreatePerson(context.Background(), person)
	if err != nil {
		fmt.Println("Could not create Person: ", err)
		return
	}
	if resp.Success {
		fmt.Println("A new Person has been added with id: ", resp.Id)
	}
}

func getPersons(client prsn.PersonClient, filter *prsn.PersonFilter) {
	stream, err := client.GetPersons(context.Background(), filter)
	if err != nil {
		fmt.Println("Error on get persons: ", err)
		return
	}
	for {
		person, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Printf("%v.GetPersons(_) = _, %v", client, err)
		}
		fmt.Println("Person: ", person)
	}
}

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		fmt.Println("did not connect: ", err)
		return
	}
	defer conn.Close()
	client := prsn.NewPersonClient(conn)
	person := &prsn.PersonRequest{
		Id:    1001,
		Name:  "Reddy",
		Email: "reddy@xyz.com",
		Phone: "9898982929",
		Addresses: []*prsn.PersonRequest_Address{
			&prsn.PersonRequest_Address{
				Street:            "Tripilcane",
				City:              "Chennai",
				State:             "TN",
				Zip:               "600002",
				IsShippingAddress: false,
			},
			&prsn.PersonRequest_Address{
				Street:            "Balaji colony",
				City:              "Tirupati",
				State:             "AP",
				Zip:               "517501",
				IsShippingAddress: true,
			},
		},
	}

	createPerson(client, person)
	person = &prsn.PersonRequest{
		Id:    1002,
		Name:  "Raj",
		Email: "raj@xyz.com",
		Phone: "5000510001",
		Addresses: []*prsn.PersonRequest_Address{
			&prsn.PersonRequest_Address{
				Street:            "Marathahalli",
				City:              "Bangalore",
				State:             "KS",
				Zip:               "560037",
				IsShippingAddress: true,
			},
		},
	}

	createPerson(client, person)
	filter := &prsn.PersonFilter{Keyword: ""}
	getPersons(client, filter)
}