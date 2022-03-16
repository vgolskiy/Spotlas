package repositary

import (
	"endpoint/enteties"
	"endpoint/utils"
)

var (
	PostgresConf = enteties.PostgresConf {
		Username:    	"postgres",
		Password:    	"password",
		Host:        	"localhost",
		Port:        	"5432",
		DatabaseName:	"asset_model",
	}
)

func GetSpotByID(id string) (error, enteties.Spot) {
	conn := utils.GetPGConnection(&PostgresConf)
	defer conn.Close()
	return
}

func GetSpotsByParameters(parameters *enteties.Parameters) (error, []enteties.Spot) {
	conn := utils.GetPGConnection(&PostgresConf)
	defer conn.Close()
	conn.Query()
}
