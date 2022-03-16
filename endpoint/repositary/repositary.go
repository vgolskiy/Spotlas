package repositary

import (
	"endpoint/enteties"
	"endpoint/utils"
	"log"
)

var (
	PostgresConf = enteties.PostgresConf {
		Username:    	"postgres",
		Password:    	"password",
		Host:        	"localhost",
		Port:        	"5432",
		DatabaseName:	"postgres",
		SSLMode:		"disable",
	}
)

func GetSpotByID(id string) (error, enteties.Spot) {
	var spot enteties.Spot

	conn := utils.GetPGConnection(&PostgresConf)
	defer conn.Close()
	err := conn.QueryRow(`select id, name, website, coordinates, description, rating
from "MY_TABLE"
where id = $1`, id).Scan(&spot.ID, &spot.Name, &spot.Website, &spot.Coordinates, &spot.Description, &spot.Rating)
	if err != nil {
		log.Println(err.Error())
		return err, spot
	}
	return nil, spot
}

func GetSpotsByParameters(parameters *enteties.Parameters) (error, []enteties.Spot) {
	var spots []enteties.Spot

	conn := utils.GetPGConnection(&PostgresConf)
	defer conn.Close()
	conn.Query(``)
	return nil, spots
}
