package main

import (
	"github.com/labstack/echo/v4"
	"github.com/miguelpragier/pgkebab"
	"net/http"
	"time"
)

type earthquake struct {
	ID                int       `json:"id"`
	OccurredOn        time.Time `json:"ocurred_on"`
	Latitude          float64   `json:"latitude"`
	Longitude         float64   `json:"longitude"`
	Depth             float64   `json:"depth"`
	Magnitude         float64   `json:"magnitude"`
	CalculationMethod string    `json:"calculation_method"`
	NetworkID         string    `json:"network_id"`
	Place             string    `json:"place"`
	Cause             string    `json:"cause"`
}

func earthquakeList(c echo.Context, db *pgkebab.DBLink) error {

	eq, err := db.GetMany("SELECT * FROM earthquake")

	if err != nil {
		return c.String(http.StatusInternalServerError, "erro técnico")
	}

	var earthquakeList []earthquake

	for eq.Next() {
		row, _ := eq.Row()

		li := earthquake{
			ID:                row.Int("earthquake_id"),
			OccurredOn:        row.Time("ocurred_on"),
			Latitude:          row.Float64("latitude"),
			Longitude:         row.Float64("longitude"),
			Depth:             row.Float64("depth"),
			Magnitude:         row.Float64("magnitude"),
			CalculationMethod: row.String("calculation_method"),
			NetworkID:         row.String("network_id"),
			Place:             row.String("place"),
			Cause:             row.String("cause"),
		}

		earthquakeList = append(earthquakeList, li)
	}

	return c.JSON(http.StatusOK, earthquakeList)
}

func earthquakeByID(c echo.Context, db *pgkebab.DBLink) error {
	eqID := c.Param("id")

	eq, err := db.GetOne("SELECT * FROM earthquake WHERE earthquake_id=$1", eqID)

	if err != nil {
		return c.String(http.StatusBadRequest, "Earthquake não cadastrado")
	}

	eqResp := earthquake{
		ID:                eq.Int("earthquake_id"),
		OccurredOn:        eq.Time("ocurred_on"),
		Latitude:          eq.Float64("latitude"),
		Longitude:         eq.Float64("longitude"),
		Depth:             eq.Float64("depth"),
		Magnitude:         eq.Float64("magnitude"),
		CalculationMethod: eq.String("calculation_method"),
		NetworkID:         eq.String("network_id"),
		Place:             eq.String("place"),
		Cause:             eq.String("cause"),
	}

	return c.JSON(http.StatusOK, eqResp)
}

func createEarthquake(c echo.Context, db *pgkebab.DBLink) error {
	var eq earthquake

	err := c.Bind(eq)

	if err != nil {
		return c.String(http.StatusUnprocessableEntity, "credenciais inválidas")
	}

	return nil
}
