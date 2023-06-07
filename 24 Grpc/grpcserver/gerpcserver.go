package grpcserver

import (
	"TopLearn/Grpc/cmd"
	"TopLearn/Grpc/datalayer"
	"context"
)

type Grpcserver struct {
	dbHandler *datalayer.SQLHandler
}

func NewGrpcServer(connString string) (*Grpcserver, error) { //"user:1234@/people"
	db, err := datalayer.CreateDBConnection(connString)
	if err != nil {
		return nil, err
	}
	return &Grpcserver{
		dbHandler: db,
	}, err
}

func (server *Grpcserver) GetPerson(ctx context.Context, r *cmd.Request) (*cmd.Person, error) {
	person, err := server.dbHandler.GetPersonByName(r.GetName())
	if err != nil {
		return nil, err
	}
	return convertToGrpcPerson(person), nil
}

func (server *Grpcserver) GetPeople(r *cmd.Request, stream cmd.PersonService_GetPeopleServer) error {
	people, err := server.dbHandler.GetPeople()
	if err != nil {
		return err
	}

	for _, person := range people {
		grpcPerson := convertToGrpcPerson(person)
		err := stream.Send(grpcPerson)
		if err != nil {
			return err
		}
	}
	return nil
}

func convertToGrpcPerson(person datalayer.Person) *cmd.Person {
	return &cmd.Person{
		Id:     int32(person.ID),
		Name:   person.Name,
		Family: person.Family,
	}
}
