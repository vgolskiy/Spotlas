package enteties

import (
	"fmt"
	"net/http"
)

type Spot struct {
	ID			string	`json:"id"`
	Name		string	`json:"name"`
	Website		string	`json:"website"`
	Coordinates	string	`json:"coordinates"`
	Description	string	`json:"description"`
	Rating		float32	`json:"rating"`
}

type Parameters struct {
	Latitude	float32	`json:"latitude"`
	Longitude	float32	`json:"longitude"`
	Radius		float32	`json:"radius"`
	Type		string	`json:"type"`
}

type PostgresConf struct {
	Username		string	`json:"username"`
	Password		string	`json:"password"`
	Host			string	`json:"host"`
	Port			string	`json:"port"`
	DatabaseName	string	`json:"database_name"`
}

func (obj *PostgresConf) GetConnectionString() string {
	return fmt.Sprintf(
		"postgres://%v:%v@%v:%v/%v",
		obj.Username,
		obj.Password,
		obj.Host,
		obj.Port,
		obj.DatabaseName)
}

type Route struct {
	Name		string
	Method		string
	Pattern		string
	Function	RoutingHandlerFunc
}

type RoutingHandlerFunc func(http.ResponseWriter, *http.Request)

func (f RoutingHandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f(w,r)
}