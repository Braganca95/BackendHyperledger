/*
 * SPDX-License-Identifier: Apache-2.0
 */

package main

type Lot struct {
	DocType         string  `json:"docType"`
	ID              string  `json:"ID"`
	CO2e            float64 `json:"co2e"`
	Year            int     `json:"year"`
	Month           string  `json:"month"`
	ProductQuantity int     `json:"productQuantity"`
	ProductID       string  `json:"productID"`
}

type ProductLotFootprint struct {
	DocType    string  `json:"docType"`
	ID         string  `json:"ID"`
	CO2e       float64 `json:"co2e"`
	Year       int     `json:"year"`
	Month      string  `json:"month"`
	ProductLot Lot     `json:"lot"`
}

type ProductQuantity struct {
	DocType  string             `json:"docType"`
	ID       string             `json:"ID"`
	Quantity int                `json:"quantity"`
	Activity ProductionActivity `json:"activity"`
}

type ProductionActivity struct {
	DocType         string           `json:"docType"`
	ID              string           `json:"ID"`
	Description     string           `json:"description"`
	CO2e            float64          `json:"co2e"`
	Date            string           `json:"date"`
	FinalProductQty int              `json:"finalProductQty"`
	ProductLot      Lot              `json:"lot"`
	FixedCost       MonthlyFixedCost `json:"fixedCost"`
	UserID          string           `json:"user"`
	OrganizationID  string           `json:"organization"`
	CostTypeID      string           `json:"costType"`
}

type ProductionCost struct {
	DocType     string             `json:"docType"`
	ID          string             `json:"ID"`
	CO2e        float64            `json:"co2e"`
	Description string             `json:"description"`
	Value       string             `json:"value"`
	Activity    ProductionActivity `json:"actiivity"`
	CostTypeID  string             `json:"costType"`
}

type MonthlyFixedCost struct {
	DocType        string  `json:"docType"`
	ID             string  `json:"ID"`
	CO2e           float64 `json:"co2e"`
	Year           int     `json:"year"`
	Quantity       int     `json:"quantity"`
	Description    string  `json:"description"`
	Month          string  `json:"month"`
	OrganizationID string  `json:"organization"`
}
