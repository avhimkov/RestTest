package main

import (
	"fmt"
	"net"
	"strings"
	"google.golang.org/grpc"
	"golang.org/x/net/context"
	prsn "github.com/agiratech/golang-rpc/person"
)

const (
	port = ":3333"
)

type Person struct {
	savedPersons []*prsn.PersonRequest
}

func (p *Person) GetPersons(fltr *PersonFilter, stream prsn.Person_GetPersonServer) error {
	for _, person :=range p.savedPersons{
		if fltr.Keyword != ""{
			if !strings.Contains(person.Name, fltr.Keyword){
				continue
			}
		}
		err := stream.Send(person)
		if err != nil{
			return err
		}
	}
	return nil
}

func main()  {
	lis, err := net.Listen("tcp", port)
		if err != nil{
			fmt.Println("failed to listen", err)
			return
		}
		s := grpc.NewServer()
		prsn.RegisterPersonServer(s, &Person{})
		s.Serve(lis)
}