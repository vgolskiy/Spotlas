package repositary

import (
	"database/sql"
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

func rowsConversion(rows *sql.Rows) (error, []enteties.Spot) {
	var spots []enteties.Spot

	for rows.Next() {
		var spot enteties.Spot
		err := rows.Scan(&spot.ID, &spot.Name, &spot.Website, &spot.Coordinates, &spot.Description, &spot.Rating)
		if err != nil {
			log.Println(err.Error())
			return err, spots
		}
		spots = append(spots, spot)
	}
	return nil, spots
}

func GetSpotsByParameters(parameters *enteties.Parameters) (error, []enteties.Spot) {
	var spots []enteties.Spot

	conn := utils.GetPGConnection(&PostgresConf)
	defer conn.Close()
	if parameters.Type == "square" {
		minLat, minLng, maxLat, maxLng := utils.GetSquareCoordinates(parameters)
		rows, err := conn.Query(`select id, name, website, coordinates, description, rating
from "MY_TABLE"
where coordinates::geometry && ST_Envelope(ST_GeomFromText(CONCAT('LINESTRING(', $1::float8, ' ', $2::float8, ',', $3::float8, ' ', $4::float8, ')'), 4326))`, minLng, minLat, maxLng, maxLat)
		if err != nil {
			log.Println(err.Error())
			return err, spots
		}
		defer rows.Close()
		err, spots = rowsConversion(rows)
		if err != nil {
			return err, spots
		}
	} else {
		rows, err := conn.Query(`select id, name, website, coordinates, description, rating
from "MY_TABLE"
where coordinates::geometry && st_buffer(st_makepoint($1, $2)::geography, $3, 'quad_segs=8');`,
parameters.Longitude, parameters.Latitude, parameters.Radius)
		if err != nil {
			log.Println(err.Error())
			return err, spots
		}
		defer rows.Close()
		err, spots = rowsConversion(rows)
		if err != nil {
			return err, spots
		}
	}
	return nil, spots
}
