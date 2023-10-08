package main

import (
	"fmt"
)

// Example of polymorphism
type DBConnectInterface interface {
	Connect() bool
}

type Session struct {
	connection DBConnectInterface
}

type PostgreSession struct {
}

type Mongosession struct {
}

func (pg *PostgreSession) Connect() bool {
	fmt.Println("Connected to postres!!")
	return true
}

func (mongo *Mongosession) Connect() bool {
	fmt.Println("Connected to mongodb!!")
	return true
}

func runPolyMorphism() {
	pgConnection := PostgreSession{}
	mongoConnection := Mongosession{}

	pgSession := Session{
		connection: &pgConnection,
	}
	mongoSession := Session{
		connection: &mongoConnection,
	}

	pgSession.connection.Connect()
	mongoSession.connection.Connect()
}

func main() {
	runPolyMorphism()
}
