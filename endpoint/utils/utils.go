package utils

import (
	"database/sql"
	"endpoint/enteties"
	_ "github.com/lib/pq"
	"log"
	"math"
	"os"
)

func GetEnvWithKey(key string, defaultValue string) string {
	if val := os.Getenv(key); len(val) > 0 {
		return val
	}
	log.Printf("Variable %v doesn't exist, using default value\n", key)
	return defaultValue
}

func GetPGConnection(conf *enteties.PostgresConf) *sql.DB {
	conn, err := sql.Open("postgres", conf.GetConnectionString())
	if err != nil {
		log.Println(err.Error())
	}
	return conn
}

func IsValidType(value string) bool {
	switch value {
	case "circle", "square":
		return true
	}
	return false
}

func radians(n float64) float64 {
	return n * math.Pi / 180
}

func degrees(n float64) float64 {
	return n * 180 / math.Pi
}

func calculateLatitudeLongitude(parameters *enteties.Parameters, bearing float64) (latitude float64, longitude float64) {
	aD := parameters.Radius/6378137
	rB := radians(bearing)
	rLat := radians(parameters.Latitude)
	rLng := radians(parameters.Longitude)
	newLatitude	:= math.Asin(math.Sin(rLat) * math.Cos(aD) + math.Cos(rLat) * math.Sin(aD) * math.Cos(rB))
	delta := math.Atan2(math.Sin(rB) * math.Sin(aD) * math.Cos(rLat),
		math.Cos(aD) - math.Sin(rLat) * math.Sin(newLatitude))
	newLongitude := math.Mod(rLng + delta + math.Pi, math.Pi * 2) - math.Pi
	return degrees(newLatitude), degrees(newLongitude)
}

func sortCoordinates(array [][]float64) (float64, float64, float64, float64) {
	minLat := array[0][0]
	minLng := array[0][1]
	maxLat := array[0][0]
	maxLng := array[0][1]
	for _, val := range array {
		if val[0] < minLat {
			minLat = val[0]
		}
		if val[0] > maxLat {
			maxLat = val[0]
		}
		if val[1] < minLng {
			minLng = val[1]
		}
		if val[1] > maxLng {
			maxLng = val[1]
		}
	}
	return minLat, minLng, maxLat, maxLng
}

func GetSquareCoordinates(parameters *enteties.Parameters) (minLat float64, minLng float64, maxLat float64, maxLng float64) {
	var coordinates [][]float64

	directions := []float64{0, 90, 180, 270}
	for _, bearing := range directions {
		lat, lng := calculateLatitudeLongitude(parameters, bearing)
		coordinates = append(coordinates, []float64{lat, lng})
	}
	return sortCoordinates(coordinates)
}
