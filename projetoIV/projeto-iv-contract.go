/*
 * SPDX-License-Identifier: Apache-2.0
 */

package main

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// ProjetoIvContract contract for managing CRUD for Lot
type ProjetoIvContract struct {
	contractapi.Contract
}

// ProjetoIvExists returns true when asset with given ID exists in world state
func (c *ProjetoIvContract) LotExists(ctx contractapi.TransactionContextInterface, lotID string) (bool, error) {
	data, err := ctx.GetStub().GetState(lotID)

	if err != nil {
		return false, err
	}

	return data != nil, nil
}

// CreateProjetoIv creates a new instance of Lot
func (c *ProjetoIvContract) CreateLot(ctx contractapi.TransactionContextInterface, lotID, month, productID string, year, productQuantity int, co2e float64) (string, error) {
	exists, err := c.LotExists(ctx, lotID)
	if err != nil {
		return "", fmt.Errorf("could not read from world state. %s", err)
	} else if exists {
		return "", fmt.Errorf("the asset %s already exists", lotID)
	}
	/*
	 lot := new(Lot)
	 lot.Value = value */

	lot := &Lot{
		DocType:         "lot",
		ID:              lotID,
		CO2e:            co2e,
		Year:            year,
		Month:           month,
		ProductQuantity: productQuantity,
		ProductID:       productID,
	}

	bytes, _ := json.Marshal(lot)

	if err != nil {
		return "", err
	}

	err = ctx.GetStub().PutState(lotID, bytes)

	if err != nil {
		return "", fmt.Errorf("failed to put to world state: %v", err)
	}

	return fmt.Sprintf("%s created sucessfully", lotID), nil
}

// ReadLot retrieves an instance of Lot from the world state
func (c *ProjetoIvContract) ReadLot(ctx contractapi.TransactionContextInterface, lotID string) (*Lot, error) {
	exists, err := c.LotExists(ctx, lotID)
	if err != nil {
		return nil, fmt.Errorf("could not read from world state. %s", err)
	} else if !exists {
		return nil, fmt.Errorf("the asset %s does not exist", lotID)
	}

	bytes, _ := ctx.GetStub().GetState(lotID)

	lot := new(Lot)

	err = json.Unmarshal(bytes, lot)

	if err != nil {
		return nil, fmt.Errorf("could not unmarshal world state data to type Lot")
	}

	return lot, nil
}

// UpdateLot retrieves an instance of Lot from the world state and updates its value
func (c *ProjetoIvContract) UpdateLot(ctx contractapi.TransactionContextInterface, lotID string, newValue string) error {
	exists, err := c.LotExists(ctx, lotID)
	if err != nil {
		return fmt.Errorf("could not read from world state. %s", err)
	} else if !exists {
		return fmt.Errorf("the asset %s does not exist", lotID)
	}

	lot := new(Lot)
	//lot.Value = newValue

	bytes, _ := json.Marshal(lot)

	return ctx.GetStub().PutState(lotID, bytes)
}

func (c *ProjetoIvContract) UpdateLotCO2e(ctx contractapi.TransactionContextInterface, lotID string, co2e float64) (string, error) {
	exists, err := c.LotExists(ctx, lotID)
	if err != nil {
		return "", fmt.Errorf("could not read from world state. %s", err)
	} else if !exists {
		return "", fmt.Errorf("the asset %s does not exist", lotID)
	}

	outdatedLotBytes, _ := ctx.GetStub().GetState(lotID)

	outdatedLot := new(Lot)

	err = json.Unmarshal(outdatedLotBytes, outdatedLot)

	if err != nil {
		return "", fmt.Errorf("could not unmarshal world state data to type Lot")
	}

	lot := &Lot{
		DocType:         "lot",
		ID:              lotID,
		CO2e:            co2e,
		Year:            outdatedLot.Year,
		Month:           outdatedLot.Month,
		ProductQuantity: outdatedLot.ProductQuantity,
		ProductID:       outdatedLot.ProductID,
	}

	bytes, _ := json.Marshal(lot)
	if err != nil {
		return "", err
	}

	err = ctx.GetStub().PutState(lotID, bytes)

	if err != nil {
		return "", fmt.Errorf("failed to put world state: %v", err)

	}

	return fmt.Sprintf("%s product quantity updated successfully to %f", lotID, co2e), nil
}

func (c *ProjetoIvContract) UpdateLotPQ(ctx contractapi.TransactionContextInterface, lotID string, productQty int) (string, error) {
	exists, err := c.LotExists(ctx, lotID)
	if err != nil {
		return "", fmt.Errorf("could not read from world state. %s", err)
	} else if !exists {
		return "", fmt.Errorf("the asset %s does not exist", lotID)
	}

	outdatedLotBytes, _ := ctx.GetStub().GetState(lotID)

	outdatedLot := new(Lot)

	err = json.Unmarshal(outdatedLotBytes, outdatedLot)

	if err != nil {
		return "", fmt.Errorf("could not unmarshal world state data to type Lot")
	}

	lot := &Lot{
		DocType:         "lot",
		ID:              lotID,
		CO2e:            outdatedLot.CO2e,
		Year:            outdatedLot.Year,
		Month:           outdatedLot.Month,
		ProductQuantity: productQty,
		ProductID:       outdatedLot.ProductID,
	}

	bytes, _ := json.Marshal(lot)
	if err != nil {
		return "", err
	}

	err = ctx.GetStub().PutState(lotID, bytes)

	if err != nil {
		return "", fmt.Errorf("failed to put world state: %v", err)

	}

	return fmt.Sprintf("%s product quantity updated successfully to %d", lotID, productQty), nil
}
func (c *ProjetoIvContract) UpdateLotProductID(ctx contractapi.TransactionContextInterface, lotID, productID string) (string, error) {
	exists, err := c.LotExists(ctx, lotID)
	if err != nil {
		return "", fmt.Errorf("could not read from world state. %s", err)
	} else if !exists {
		return "", fmt.Errorf("the asset %s does not exist", lotID)
	}

	outdatedLotBytes, _ := ctx.GetStub().GetState(lotID)

	outdatedLot := new(Lot)

	err = json.Unmarshal(outdatedLotBytes, outdatedLot)

	if err != nil {
		return "", fmt.Errorf("could not unmarshal world state data to type Lot")
	}

	lot := &Lot{
		DocType:         "lot",
		ID:              lotID,
		CO2e:            outdatedLot.CO2e,
		Year:            outdatedLot.Year,
		Month:           outdatedLot.Month,
		ProductQuantity: outdatedLot.ProductQuantity,
		ProductID:       productID,
	}

	bytes, _ := json.Marshal(lot)
	if err != nil {
		return "", err
	}

	err = ctx.GetStub().PutState(lotID, bytes)

	if err != nil {
		return "", fmt.Errorf("failed to put world state: %v", err)

	}

	return fmt.Sprintf("%s product quantity updated successfully to %s", lotID, productID), nil
}

/*
 // UpdateLot retrieves an instance of Lot from the world state and updates its value
 func (c *ProjetoIvContract) TransferLot(ctx contractapi.TransactionContextInterface, lotID, newOwner string) (string,error) {
	 exists, err := c.LotExists(ctx, lotID)
	 if err != nil {
		 return "",fmt.Errorf("Could not read from world state. %s", err)
	 } else if !exists {
		 return "",fmt.Errorf("The asset %s does not exist", lotID)
	 }

	 lot := &Lot{
		 DocType:         "lot",
		 ID:              lotID,
		 CO2e:            co2e,
		 Year:            year,
		 Month:           month,
		 ProductQuantity: productQuantity,
	 }

	 bytes, _ := json.Marshal(lot)

	 if err != nil {
		 return "", err
	 }

	 err = ctx.GetStub().PutState(lotID, bytes)

	 if err != nil {
		 return "", fmt.Errorf("failed to put to world state: %v", err)
	 }

	 return fmt.Sprintf("%s created sucessfully", lotID), nil
 }
*/
// DeleteLot deletes an instance of Lot from the world state
func (c *ProjetoIvContract) DeleteLot(ctx contractapi.TransactionContextInterface, lotID string) error {
	exists, err := c.LotExists(ctx, lotID)
	if err != nil {
		return fmt.Errorf("could not read from world state. %s", err)
	} else if !exists {
		return fmt.Errorf("the asset %s does not exist", lotID)
	}

	return ctx.GetStub().DelState(lotID)
}

// ------------------------- PRODUCT LOT FOOTPRINT -----------------

// ProjetoIvExists returns true when asset with given ID exists in world state
func (c *ProjetoIvContract) PLFExists(ctx contractapi.TransactionContextInterface, plfID string) (bool, error) {
	data, err := ctx.GetStub().GetState(plfID)

	if err != nil {
		return false, err
	}

	return data != nil, nil
}

// CreateProjetoIv creates a new instance of Lot
func (c *ProjetoIvContract) CreatePLF(ctx contractapi.TransactionContextInterface, plfID, month string, year int, co2e float64, productLot Lot) (string, error) {
	exists, err := c.PLFExists(ctx, plfID)
	if err != nil {
		return "", fmt.Errorf("could not read from world state. %s", err)
	} else if exists {
		return "", fmt.Errorf("the asset %s already exists", plfID)
	}

	exists, err = c.LotExists(ctx, productLot.ID)
	if err != nil {
		return "", fmt.Errorf("could not read from world state. %s", err)
	} else if !exists {
		return "", fmt.Errorf("the lot [%s] doesn't exists", productLot.ID)
	}

	/*
	 lot := new(Lot)
	 lot.Value = value */

	plf := &ProductLotFootprint{
		DocType:    "plf",
		ID:         plfID,
		CO2e:       co2e,
		Year:       year,
		Month:      month,
		ProductLot: productLot,
	}

	bytes, _ := json.Marshal(plf)

	if err != nil {
		return "", err
	}

	err = ctx.GetStub().PutState(plfID, bytes)

	if err != nil {
		return "", fmt.Errorf("failed to put to world state: %v", err)
	}

	return fmt.Sprintf("%s created sucessfully", plfID), nil
}

// ReadLot retrieves an instance of Lot from the world state
func (c *ProjetoIvContract) ReadPLF(ctx contractapi.TransactionContextInterface, plfID string) (*ProductLotFootprint, error) {
	exists, err := c.PLFExists(ctx, plfID)
	if err != nil {
		return nil, fmt.Errorf("could not read from world state. %s", err)
	} else if !exists {
		return nil, fmt.Errorf("the asset %s does not exist", plfID)
	}

	bytes, _ := ctx.GetStub().GetState(plfID)

	plf := new(ProductLotFootprint)

	err = json.Unmarshal(bytes, plf)

	if err != nil {
		return nil, fmt.Errorf("could not unmarshal world state data to type Product Lot Footprint")
	}

	return plf, nil
}

// UpdateLot retrieves an instance of Lot from the world state and updates its value
func (c *ProjetoIvContract) UpdatePLF(ctx contractapi.TransactionContextInterface, plfID string, newValue string) error {
	exists, err := c.PLFExists(ctx, plfID)
	if err != nil {
		return fmt.Errorf("could not read from world state. %s", err)
	} else if !exists {
		return fmt.Errorf("the asset %s does not exist", plfID)
	}

	plf := new(ProductLotFootprint)
	//lot.Value = newValue

	bytes, _ := json.Marshal(plf)

	return ctx.GetStub().PutState(plfID, bytes)
}

func (c *ProjetoIvContract) UpdatePLFLot(ctx contractapi.TransactionContextInterface, plfID string, lot Lot) (string, error) {
	exists, err := c.PLFExists(ctx, plfID)
	if err != nil {
		return "", fmt.Errorf("could not read from world state. %s", err)
	} else if !exists {
		return "", fmt.Errorf("the asset %s does not exist", plfID)
	}

	exists, err = c.LotExists(ctx, lot.ID)
	if err != nil {
		return "", fmt.Errorf("could not read from world state. %s", err)
	} else if !exists {
		return "", fmt.Errorf("the lot [%s] doesn't exists", lot.ID)
	}

	outdatedPLFBytes, _ := ctx.GetStub().GetState(plfID)

	outdatedPLF := new(ProductLotFootprint)

	err = json.Unmarshal(outdatedPLFBytes, outdatedPLF)

	if err != nil {
		return "", fmt.Errorf("could not unmarshal world state data to type Product Lot Footprint")
	}

	plf := &ProductLotFootprint{
		DocType:    "plf",
		ID:         plfID,
		CO2e:       outdatedPLF.CO2e,
		Year:       outdatedPLF.Year,
		Month:      outdatedPLF.Month,
		ProductLot: lot,
	}

	bytes, _ := json.Marshal(lot)
	if err != nil {
		return "", err
	}

	err = ctx.GetStub().PutState(plfID, bytes)

	if err != nil {
		return "", fmt.Errorf("failed to put world state: %v", err)

	}

	return fmt.Sprintf("%s Product Lot Footprint Lot updated successfully to %s", plfID, plf.ID), nil
}

func (c *ProjetoIvContract) UpdatePLFCO2e(ctx contractapi.TransactionContextInterface, plfID string, co2e float64) (string, error) {
	exists, err := c.PLFExists(ctx, plfID)
	if err != nil {
		return "", fmt.Errorf("could not read from world state. %s", err)
	} else if !exists {
		return "", fmt.Errorf("the asset %s does not exist", plfID)
	}
	outdatedPLFBytes, _ := ctx.GetStub().GetState(plfID)

	outdatedPLF := new(ProductLotFootprint)

	err = json.Unmarshal(outdatedPLFBytes, outdatedPLF)

	if err != nil {
		return "", fmt.Errorf("could not unmarshal world state data to type Product Lot Footprint")
	}

	plf := &ProductLotFootprint{
		DocType:    "plf",
		ID:         plfID,
		CO2e:       co2e,
		Year:       outdatedPLF.Year,
		Month:      outdatedPLF.Month,
		ProductLot: outdatedPLF.ProductLot,
	}

	bytes, _ := json.Marshal(plf)
	if err != nil {
		return "", err
	}

	err = ctx.GetStub().PutState(plfID, bytes)

	if err != nil {
		return "", fmt.Errorf("failed to put world state: %v", err)

	}

	return fmt.Sprintf("%s Product Lot Footprint CO2 emissions updated successfully to %f", plfID, co2e), nil
}

// DeleteLot deletes an instance of Lot from the world state
func (c *ProjetoIvContract) DeletePLF(ctx contractapi.TransactionContextInterface, plfID string) error {
	exists, err := c.PLFExists(ctx, plfID)
	if err != nil {
		return fmt.Errorf("could not read from world state. %s", err)
	} else if !exists {
		return fmt.Errorf("the asset %s does not exist", plfID)
	}

	return ctx.GetStub().DelState(plfID)
}

// ------------------------- PRODUCT QUANTITY -----------------

// ProjetoIvExists returns true when asset with given ID exists in world state
func (c *ProjetoIvContract) PQExists(ctx contractapi.TransactionContextInterface, pqID string) (bool, error) {
	data, err := ctx.GetStub().GetState(pqID)

	if err != nil {
		return false, err
	}

	return data != nil, nil
}

// CreateProjetoIv creates a new instance of PQ
func (c *ProjetoIvContract) CreatePQ(ctx contractapi.TransactionContextInterface, pqID string, quantity int, activity ProductionActivity) (string, error) {
	exists, err := c.PQExists(ctx, pqID)
	if err != nil {
		return "", fmt.Errorf("could not read from world state. %s", err)
	} else if exists {
		return "", fmt.Errorf("the asset %s already exists", pqID)
	}

	exists, err = c.PAExists(ctx, activity.ID)
	if err != nil {
		return "", fmt.Errorf("could not read from world state. %s", err)
	} else if !exists {
		return "", fmt.Errorf("the lot [%s] doesn't exist", activity.ID)
	}
	/*
	 lot := new(Lot)
	 lot.Value = value */

	pq := &ProductQuantity{
		DocType:  "pq",
		ID:       pqID,
		Quantity: quantity,
		Activity: activity,
	}

	bytes, _ := json.Marshal(pq)

	if err != nil {
		return "", err
	}

	err = ctx.GetStub().PutState(pqID, bytes)

	if err != nil {
		return "", fmt.Errorf("failed to put to world state: %v", err)
	}

	return fmt.Sprintf("%s created sucessfully", pqID), nil
}

// ReadPQ retrieves an instance of PQ from the world state
func (c *ProjetoIvContract) ReadPQ(ctx contractapi.TransactionContextInterface, pqID string) (*ProductQuantity, error) {
	exists, err := c.PQExists(ctx, pqID)
	if err != nil {
		return nil, fmt.Errorf("could not read from world state. %s", err)
	} else if !exists {
		return nil, fmt.Errorf("the asset %s does not exist", pqID)
	}

	bytes, _ := ctx.GetStub().GetState(pqID)

	pq := new(ProductQuantity)

	err = json.Unmarshal(bytes, pq)

	if err != nil {
		return nil, fmt.Errorf("could not unmarshal world state data to type Lot")
	}

	return pq, nil
}

// UpdatePQ retrieves an instance of PQ from the world state and updates its value
func (c *ProjetoIvContract) UpdatePQ(ctx contractapi.TransactionContextInterface, pqID string, newValue string) error {
	exists, err := c.PQExists(ctx, pqID)
	if err != nil {
		return fmt.Errorf("could not read from world state. %s", err)
	} else if !exists {
		return fmt.Errorf("the asset %s does not exist", pqID)
	}

	pq := new(ProductQuantity)
	//pq.Value = newValue

	bytes, _ := json.Marshal(pq)

	return ctx.GetStub().PutState(pqID, bytes)
}

func (c *ProjetoIvContract) UpdatePQQty(ctx contractapi.TransactionContextInterface, pqID string, quantity int) (string, error) {
	exists, err := c.PQExists(ctx, pqID)
	if err != nil {
		return "", fmt.Errorf("could not read from world state. %s", err)
	} else if !exists {
		return "", fmt.Errorf("the asset %s does not exist", pqID)
	}

	outdatedPQBytes, _ := ctx.GetStub().GetState(pqID)

	outdatedPQ := new(ProductQuantity)

	err = json.Unmarshal(outdatedPQBytes, outdatedPQ)

	if err != nil {
		return "", fmt.Errorf("could not unmarshal world state data to type Product Quantity")
	}

	pq := &ProductQuantity{
		DocType:  "pq",
		ID:       outdatedPQ.ID,
		Quantity: quantity,
		Activity: outdatedPQ.Activity,
	}

	bytes, _ := json.Marshal(pq)
	if err != nil {
		return "", err
	}

	err = ctx.GetStub().PutState(pqID, bytes)

	if err != nil {
		return "", fmt.Errorf("failed to put world state: %v", err)

	}

	return fmt.Sprintf("%s Product Quantity quantity updated successfully to %d", pqID, quantity), nil
}

func (c *ProjetoIvContract) UpdatePQActivity(ctx contractapi.TransactionContextInterface, pqID string, activity ProductionActivity) (string, error) {
	exists, err := c.PQExists(ctx, pqID)
	if err != nil {
		return "", fmt.Errorf("could not read from world state. %s", err)
	} else if !exists {
		return "", fmt.Errorf("the asset %s does not exist", pqID)
	}

	exists, err = c.PAExists(ctx, activity.ID)
	if err != nil {
		return "", fmt.Errorf("could not read from world state. %s", err)
	} else if !exists {
		return "", fmt.Errorf("the lot [%s] doesn't exists", activity.ID)
	}

	outdatedPQBytes, _ := ctx.GetStub().GetState(pqID)

	outdatedPQ := new(ProductQuantity)

	err = json.Unmarshal(outdatedPQBytes, outdatedPQ)

	if err != nil {
		return "", fmt.Errorf("could not unmarshal world state data to type Product Quantity")
	}

	pq := &ProductQuantity{
		DocType:  "pq",
		ID:       outdatedPQ.ID,
		Quantity: outdatedPQ.Quantity,
		Activity: activity,
	}

	bytes, _ := json.Marshal(pq)
	if err != nil {
		return "", err
	}

	err = ctx.GetStub().PutState(pqID, bytes)

	if err != nil {
		return "", fmt.Errorf("failed to put world state: %v", err)

	}

	return fmt.Sprintf("%s Product Quantity activity updated successfully to %s", pqID, activity.ID), nil
}

// DeletePQ deletes an instance of PQ from the world state
func (c *ProjetoIvContract) DeletePQ(ctx contractapi.TransactionContextInterface, pqID string) error {
	exists, err := c.PQExists(ctx, pqID)
	if err != nil {
		return fmt.Errorf("could not read from world state. %s", err)
	} else if !exists {
		return fmt.Errorf("the asset %s does not exist", pqID)
	}

	return ctx.GetStub().DelState(pqID)
}

// ------------------------- PRODUCTION ACTIVITY -----------------

// ProjetoIvExists returns true when asset with given ID exists in world state
func (c *ProjetoIvContract) PAExists(ctx contractapi.TransactionContextInterface, paID string) (bool, error) {
	data, err := ctx.GetStub().GetState(paID)

	if err != nil {
		return false, err
	}

	return data != nil, nil
}

// CreateProjetoIv creates a new instance of PA
func (c *ProjetoIvContract) CreatePA(ctx contractapi.TransactionContextInterface, paID, description, date, userID, organizationID, costTypeID string, finalProductQty int, co2e float64, lot Lot, monthlyFixedCost MonthlyFixedCost) (string, error) {
	exists, err := c.PAExists(ctx, paID)
	if err != nil {
		return "", fmt.Errorf("could not read from world state. %s", err)
	} else if exists {
		return "", fmt.Errorf("the asset %s already exists", paID)
	}

	exists, err = c.LotExists(ctx, lot.ID)
	if err != nil {
		return "", fmt.Errorf("could not read from world state. %s", err)
	} else if !exists {
		return "", fmt.Errorf("the lot [%s] doesn't exist", lot.ID)
	}

	exists, err = c.MFCExists(ctx, monthlyFixedCost.ID)
	if err != nil {
		return "", fmt.Errorf("could not read from world state. %s", err)
	} else if !exists {
		return "", fmt.Errorf("the monthy fixed cost [%s] doesn't exist", monthlyFixedCost.ID)
	}
	/*
	 lot := new(Lot)
	 lot.Value = value */

	pa := &ProductionActivity{
		DocType:         "pa",
		ID:              paID,
		Description:     description,
		CO2e:            co2e,
		Date:            date,
		FinalProductQty: finalProductQty,
		ProductLot:      lot,
		FixedCost:       monthlyFixedCost,
		UserID:          userID,
		OrganizationID:  organizationID,
		CostTypeID:      costTypeID,
	}

	bytes, _ := json.Marshal(pa)

	if err != nil {
		return "", err
	}

	err = ctx.GetStub().PutState(paID, bytes)

	if err != nil {
		return "", fmt.Errorf("failed to put to world state: %v", err)
	}

	return fmt.Sprintf("%s created sucessfully", paID), nil
}

// ReadPA retrieves an instance of PA from the world state
func (c *ProjetoIvContract) ReadPA(ctx contractapi.TransactionContextInterface, paID string) (*ProductionActivity, error) {
	exists, err := c.PAExists(ctx, paID)
	if err != nil {
		return nil, fmt.Errorf("could not read from world state. %s", err)
	} else if !exists {
		return nil, fmt.Errorf("the asset %s does not exist", paID)
	}

	bytes, _ := ctx.GetStub().GetState(paID)

	pa := new(ProductionActivity)

	err = json.Unmarshal(bytes, pa)

	if err != nil {
		return nil, fmt.Errorf("could not unmarshal world state data to type PA")
	}

	return pa, nil
}

// UpdatePA retrieves an instance of PA from the world state and updates its value
func (c *ProjetoIvContract) UpdatePA(ctx contractapi.TransactionContextInterface, paID string, newValue string) error {
	exists, err := c.PAExists(ctx, paID)
	if err != nil {
		return fmt.Errorf("could not read from world state. %s", err)
	} else if !exists {
		return fmt.Errorf("the asset %s does not exist", paID)
	}

	pa := new(ProductionActivity)
	//lot.Value = newValue

	bytes, _ := json.Marshal(pa)

	return ctx.GetStub().PutState(paID, bytes)
}

func (c *ProjetoIvContract) UpdatePADescription(ctx contractapi.TransactionContextInterface, paID string, description string) (string, error) {
	exists, err := c.PAExists(ctx, paID)
	if err != nil {
		return "", fmt.Errorf("could not read from world state. %s", err)
	} else if !exists {
		return "", fmt.Errorf("the asset %s does not exist", paID)
	}
	/*
		 exists, err = c.PAExists(ctx, activity.ID)
		 if err != nil {
			 return "", fmt.Errorf("Could not read from world state. %s", err)
		 } else if !exists {
			 return "", fmt.Errorf("The lot [%s] doesn't exists", activity.ID)
		 } */

	outdatedPABytes, _ := ctx.GetStub().GetState(paID)

	outdatedPA := new(ProductionActivity)

	err = json.Unmarshal(outdatedPABytes, outdatedPA)

	if err != nil {
		return "", fmt.Errorf("could not unmarshal world state data to type Production Activity")
	}

	pa := &ProductionActivity{
		DocType:         "pa",
		ID:              paID,
		Description:     description,
		CO2e:            outdatedPA.CO2e,
		Date:            outdatedPA.Date,
		FinalProductQty: outdatedPA.FinalProductQty,
		ProductLot:      outdatedPA.ProductLot,
		FixedCost:       outdatedPA.FixedCost,
		UserID:          outdatedPA.UserID,
		OrganizationID:  outdatedPA.OrganizationID,
		CostTypeID:      outdatedPA.CostTypeID,
	}

	bytes, _ := json.Marshal(pa)
	if err != nil {
		return "", err
	}

	err = ctx.GetStub().PutState(paID, bytes)

	if err != nil {
		return "", fmt.Errorf("failed to put world state: %v", err)

	}

	return fmt.Sprintf("%s Production Activity description updated successfully to %s", paID, description), nil
}

func (c *ProjetoIvContract) UpdatePACO2e(ctx contractapi.TransactionContextInterface, paID string, co2e float64) (string, error) {
	exists, err := c.PAExists(ctx, paID)
	if err != nil {
		return "", fmt.Errorf("could not read from world state. %s", err)
	} else if !exists {
		return "", fmt.Errorf("the asset %s does not exist", paID)
	}

	/*
		 exists, err = c.PAExists(ctx, activity.ID)
		 if err != nil {
			 return "", fmt.Errorf("Could not read from world state. %s", err)
		 } else if !exists {
			 return "", fmt.Errorf("The lot [%s] doesn't exists", activity.ID)
		 } */

	outdatedPABytes, _ := ctx.GetStub().GetState(paID)

	outdatedPA := new(ProductionActivity)

	err = json.Unmarshal(outdatedPABytes, outdatedPA)

	if err != nil {
		return "", fmt.Errorf("could not unmarshal world state data to type Production Activity")
	}

	pa := &ProductionActivity{
		DocType:         "pa",
		ID:              paID,
		Description:     outdatedPA.Description,
		CO2e:            co2e,
		Date:            outdatedPA.Date,
		FinalProductQty: outdatedPA.FinalProductQty,
		ProductLot:      outdatedPA.ProductLot,
		FixedCost:       outdatedPA.FixedCost,
		UserID:          outdatedPA.UserID,
		OrganizationID:  outdatedPA.OrganizationID,
		CostTypeID:      outdatedPA.CostTypeID,
	}

	bytes, _ := json.Marshal(pa)
	if err != nil {
		return "", err
	}

	err = ctx.GetStub().PutState(paID, bytes)

	if err != nil {
		return "", fmt.Errorf("failed to put world state: %v", err)

	}

	return fmt.Sprintf("%s Production Activity CO2 emissions updated successfully to %f", paID, co2e), nil
}

func (c *ProjetoIvContract) UpdatePAFinalProductQty(ctx contractapi.TransactionContextInterface, paID string, finalProductQty int) (string, error) {
	exists, err := c.PAExists(ctx, paID)
	if err != nil {
		return "", fmt.Errorf("could not read from world state. %s", err)
	} else if !exists {
		return "", fmt.Errorf("the asset %s does not exist", paID)
	}

	/*
		 exists, err = c.PAExists(ctx, activity.ID)
		 if err != nil {
			 return "", fmt.Errorf("Could not read from world state. %s", err)
		 } else if !exists {
			 return "", fmt.Errorf("The lot [%s] doesn't exists", activity.ID)
		 } */

	outdatedPABytes, _ := ctx.GetStub().GetState(paID)

	outdatedPA := new(ProductionActivity)

	err = json.Unmarshal(outdatedPABytes, outdatedPA)

	if err != nil {
		return "", fmt.Errorf("could not unmarshal world state data to type Production Activity")
	}

	pa := &ProductionActivity{
		DocType:         "pa",
		ID:              paID,
		Description:     outdatedPA.Description,
		CO2e:            outdatedPA.CO2e,
		Date:            outdatedPA.Date,
		FinalProductQty: finalProductQty,
		ProductLot:      outdatedPA.ProductLot,
		FixedCost:       outdatedPA.FixedCost,
		UserID:          outdatedPA.UserID,
		OrganizationID:  outdatedPA.OrganizationID,
		CostTypeID:      outdatedPA.CostTypeID,
	}

	bytes, _ := json.Marshal(pa)
	if err != nil {
		return "", err
	}

	err = ctx.GetStub().PutState(paID, bytes)

	if err != nil {
		return "", fmt.Errorf("failed to put world state: %v", err)

	}

	return fmt.Sprintf("%s Production Activity Final Product Quantity updated successfully to %d", paID, finalProductQty), nil
}

func (c *ProjetoIvContract) UpdatePALot(ctx contractapi.TransactionContextInterface, paID string, lot Lot) (string, error) {
	exists, err := c.PAExists(ctx, paID)
	if err != nil {
		return "", fmt.Errorf("could not read from world state. %s", err)
	} else if !exists {
		return "", fmt.Errorf("the asset %s does not exist", paID)
	}

	exists, err = c.LotExists(ctx, lot.ID)
	if err != nil {
		return "", fmt.Errorf("could not read from world state. %s", err)
	} else if !exists {
		return "", fmt.Errorf("the lot [%s] doesn't exists", lot.ID)
	}

	outdatedPABytes, _ := ctx.GetStub().GetState(paID)

	outdatedPA := new(ProductionActivity)

	err = json.Unmarshal(outdatedPABytes, outdatedPA)

	if err != nil {
		return "", fmt.Errorf("could not unmarshal world state data to type Production Activity")
	}

	pa := &ProductionActivity{
		DocType:         "pa",
		ID:              paID,
		Description:     outdatedPA.Description,
		CO2e:            outdatedPA.CO2e,
		Date:            outdatedPA.Date,
		FinalProductQty: outdatedPA.FinalProductQty,
		ProductLot:      lot,
		FixedCost:       outdatedPA.FixedCost,
		UserID:          outdatedPA.UserID,
		OrganizationID:  outdatedPA.OrganizationID,
		CostTypeID:      outdatedPA.CostTypeID,
	}

	bytes, _ := json.Marshal(pa)
	if err != nil {
		return "", err
	}

	err = ctx.GetStub().PutState(paID, bytes)

	if err != nil {
		return "", fmt.Errorf("failed to put world state: %v", err)

	}

	return fmt.Sprintf("%s Production Activity Product Lot updated successfully to %s", paID, lot.ID), nil
}

func (c *ProjetoIvContract) UpdatePAFixedCost(ctx contractapi.TransactionContextInterface, paID string, fixedCost MonthlyFixedCost) (string, error) {
	exists, err := c.PAExists(ctx, paID)
	if err != nil {
		return "", fmt.Errorf("could not read from world state. %s", err)
	} else if !exists {
		return "", fmt.Errorf("the asset %s does not exist", paID)
	}

	exists, err = c.MFCExists(ctx, fixedCost.ID)
	if err != nil {
		return "", fmt.Errorf("could not read from world state. %s", err)
	} else if !exists {
		return "", fmt.Errorf("the fixed cost [%s] doesn't exists", fixedCost.ID)
	}

	outdatedPABytes, _ := ctx.GetStub().GetState(paID)

	outdatedPA := new(ProductionActivity)

	err = json.Unmarshal(outdatedPABytes, outdatedPA)

	if err != nil {
		return "", fmt.Errorf("could not unmarshal world state data to type Production Activity")
	}

	pa := &ProductionActivity{
		DocType:         "pa",
		ID:              paID,
		Description:     outdatedPA.Description,
		CO2e:            outdatedPA.CO2e,
		Date:            outdatedPA.Date,
		FinalProductQty: outdatedPA.FinalProductQty,
		ProductLot:      outdatedPA.ProductLot,
		FixedCost:       fixedCost,
		UserID:          outdatedPA.UserID,
		OrganizationID:  outdatedPA.OrganizationID,
		CostTypeID:      outdatedPA.CostTypeID,
	}

	bytes, _ := json.Marshal(pa)
	if err != nil {
		return "", err
	}

	err = ctx.GetStub().PutState(paID, bytes)

	if err != nil {
		return "", fmt.Errorf("failed to put world state: %v", err)

	}

	return fmt.Sprintf("%s Production Activity CO2 emissions updated successfully to %s", paID, fixedCost.ID), nil
}

func (c *ProjetoIvContract) UpdatePAUserID(ctx contractapi.TransactionContextInterface, paID, userID string) (string, error) {
	exists, err := c.PAExists(ctx, paID)
	if err != nil {
		return "", fmt.Errorf("could not read from world state. %s", err)
	} else if !exists {
		return "", fmt.Errorf("the asset %s does not exist", paID)
	}

	outdatedPABytes, _ := ctx.GetStub().GetState(paID)

	outdatedPA := new(ProductionActivity)

	err = json.Unmarshal(outdatedPABytes, outdatedPA)

	if err != nil {
		return "", fmt.Errorf("could not unmarshal world state data to type Production Activity")
	}

	pa := &ProductionActivity{
		DocType:         "pa",
		ID:              paID,
		Description:     outdatedPA.Description,
		CO2e:            outdatedPA.CO2e,
		Date:            outdatedPA.Date,
		FinalProductQty: outdatedPA.FinalProductQty,
		ProductLot:      outdatedPA.ProductLot,
		FixedCost:       outdatedPA.FixedCost,
		UserID:          userID,
		OrganizationID:  outdatedPA.OrganizationID,
		CostTypeID:      outdatedPA.CostTypeID,
	}

	bytes, _ := json.Marshal(pa)
	if err != nil {
		return "", err
	}

	err = ctx.GetStub().PutState(paID, bytes)

	if err != nil {
		return "", fmt.Errorf("failed to put world state: %v", err)

	}

	return fmt.Sprintf("%s Production Activity CO2 emissions updated successfully to %s", paID, userID), nil
}

func (c *ProjetoIvContract) UpdatePAOrganizationID(ctx contractapi.TransactionContextInterface, paID, organizationID string) (string, error) {
	exists, err := c.PAExists(ctx, paID)
	if err != nil {
		return "", fmt.Errorf("could not read from world state. %s", err)
	} else if !exists {
		return "", fmt.Errorf("the asset %s does not exist", paID)
	}

	outdatedPABytes, _ := ctx.GetStub().GetState(paID)

	outdatedPA := new(ProductionActivity)

	err = json.Unmarshal(outdatedPABytes, outdatedPA)

	if err != nil {
		return "", fmt.Errorf("could not unmarshal world state data to type Production Activity")
	}

	pa := &ProductionActivity{
		DocType:         "pa",
		ID:              paID,
		Description:     outdatedPA.Description,
		CO2e:            outdatedPA.CO2e,
		Date:            outdatedPA.Date,
		FinalProductQty: outdatedPA.FinalProductQty,
		ProductLot:      outdatedPA.ProductLot,
		FixedCost:       outdatedPA.FixedCost,
		UserID:          outdatedPA.UserID,
		OrganizationID:  organizationID,
		CostTypeID:      outdatedPA.CostTypeID,
	}

	bytes, _ := json.Marshal(pa)
	if err != nil {
		return "", err
	}

	err = ctx.GetStub().PutState(paID, bytes)

	if err != nil {
		return "", fmt.Errorf("failed to put world state: %v", err)

	}

	return fmt.Sprintf("%s Production Activity CO2 emissions updated successfully to %s", paID, organizationID), nil
}

func (c *ProjetoIvContract) UpdatePACostTypeID(ctx contractapi.TransactionContextInterface, paID, costTypeID string) (string, error) {
	exists, err := c.PAExists(ctx, paID)
	if err != nil {
		return "", fmt.Errorf("could not read from world state. %s", err)
	} else if !exists {
		return "", fmt.Errorf("the asset %s does not exist", paID)
	}

	outdatedPABytes, _ := ctx.GetStub().GetState(paID)

	outdatedPA := new(ProductionActivity)

	err = json.Unmarshal(outdatedPABytes, outdatedPA)

	if err != nil {
		return "", fmt.Errorf("could not unmarshal world state data to type Production Activity")
	}

	pa := &ProductionActivity{
		DocType:         "pa",
		ID:              paID,
		Description:     outdatedPA.Description,
		CO2e:            outdatedPA.CO2e,
		Date:            outdatedPA.Date,
		FinalProductQty: outdatedPA.FinalProductQty,
		ProductLot:      outdatedPA.ProductLot,
		FixedCost:       outdatedPA.FixedCost,
		UserID:          outdatedPA.UserID,
		OrganizationID:  outdatedPA.OrganizationID,
		CostTypeID:      costTypeID,
	}

	bytes, _ := json.Marshal(pa)
	if err != nil {
		return "", err
	}

	err = ctx.GetStub().PutState(paID, bytes)

	if err != nil {
		return "", fmt.Errorf("failed to put world state: %v", err)

	}

	return fmt.Sprintf("%s Production Activity CO2 emissions updated successfully to %s", paID, costTypeID), nil
}

// DeletePA deletes an instance of PA from the world state
func (c *ProjetoIvContract) DeletePA(ctx contractapi.TransactionContextInterface, paID string) error {
	exists, err := c.PAExists(ctx, paID)
	if err != nil {
		return fmt.Errorf("could not read from world state. %s", err)
	} else if !exists {
		return fmt.Errorf("the asset %s does not exist", paID)
	}

	return ctx.GetStub().DelState(paID)
}

// ------------------------- PRODUCTION COST -----------------

// ProjetoIvExists returns true when asset with given ID exists in world state
func (c *ProjetoIvContract) PCExists(ctx contractapi.TransactionContextInterface, pcID string) (bool, error) {
	data, err := ctx.GetStub().GetState(pcID)

	if err != nil {
		return false, err
	}

	return data != nil, nil
}

// CreateProjetoIv creates a new instance of PC
func (c *ProjetoIvContract) CreatePC(ctx contractapi.TransactionContextInterface, pcID, description, value, costTypeID string, co2e float64, activity ProductionActivity) (string, error) {
	exists, err := c.PCExists(ctx, pcID)
	if err != nil {
		return "", fmt.Errorf("could not read from world state. %s", err)
	} else if exists {
		return "", fmt.Errorf("the asset %s already exists", pcID)
	}

	exists, err = c.PAExists(ctx, activity.ID)
	if err != nil {
		return "", fmt.Errorf("could not read from world state. %s", err)
	} else if !exists {
		return "", fmt.Errorf("the production activity [%s] doesn't exist", activity.ID)
	}
	/*
	 lot := new(Lot)
	 lot.Value = value */

	pc := &ProductionCost{
		DocType:     "pc",
		ID:          pcID,
		CO2e:        co2e,
		Description: description,
		Value:       value,
		Activity:    activity,
		CostTypeID:  costTypeID,
	}

	bytes, _ := json.Marshal(pc)

	if err != nil {
		return "", err
	}

	err = ctx.GetStub().PutState(pcID, bytes)

	if err != nil {
		return "", fmt.Errorf("failed to put to world state: %v", err)
	}

	return fmt.Sprintf("%s created sucessfully", pcID), nil
}

// ReadPC retrieves an instance of PC from the world state
func (c *ProjetoIvContract) ReadPC(ctx contractapi.TransactionContextInterface, pcID string) (*ProductionCost, error) {
	exists, err := c.PCExists(ctx, pcID)
	if err != nil {
		return nil, fmt.Errorf("could not read from world state. %s", err)
	} else if !exists {
		return nil, fmt.Errorf("the asset %s does not exist", pcID)
	}

	bytes, _ := ctx.GetStub().GetState(pcID)

	pc := new(ProductionCost)

	err = json.Unmarshal(bytes, pc)

	if err != nil {
		return nil, fmt.Errorf("could not unmarshal world state data to type Lot")
	}

	return pc, nil
}

// UpdatePC retrieves an instance of PC from the world state and updates its value
func (c *ProjetoIvContract) UpdatePC(ctx contractapi.TransactionContextInterface, pcID string, newValue string) (string, error) {
	exists, err := c.PCExists(ctx, pcID)
	if err != nil {
		return "", fmt.Errorf("could not read from world state. %s", err)
	} else if !exists {
		return "", fmt.Errorf("the asset %s does not exist", pcID)
	}

	outdatedPCBytes, _ := ctx.GetStub().GetState(pcID)

	outdatedPC := new(ProductionCost)

	err = json.Unmarshal(outdatedPCBytes, outdatedPC)

	if err != nil {
		return "", fmt.Errorf("could not unmarshal world state data to type Production Activity")
	}

	pc := &ProductionCost{
		DocType:     "pc",
		ID:          outdatedPC.ID,
		CO2e:        outdatedPC.CO2e,
		Description: outdatedPC.Description,
		Value:       outdatedPC.Value,
		Activity:    outdatedPC.Activity,
		CostTypeID:  outdatedPC.CostTypeID,
	}

	bytes, _ := json.Marshal(pc)
	if err != nil {
		return "", err
	}

	err = ctx.GetStub().PutState(pcID, bytes)

	if err != nil {
		return "", fmt.Errorf("failed to put world state: %v", err)

	}

	return fmt.Sprintf("%s Production Activity CO2 emissions updated successfully to ", pcID), nil
}

func (c *ProjetoIvContract) UpdatePCCO2e(ctx contractapi.TransactionContextInterface, pcID string, co2e float64) (string, error) {
	exists, err := c.PCExists(ctx, pcID)
	if err != nil {
		return "", fmt.Errorf("could not read from world state. %s", err)
	} else if !exists {
		return "", fmt.Errorf("the asset %s does not exist", pcID)
	}

	/* exists, err = c.PAExists(ctx, fixedCost.ID)
	 if err != nil {
		 return "", fmt.Errorf("Could not read from world state. %s", err)
	 } else if !exists {
		 return "", fmt.Errorf("The fixed cost [%s] doesn't exists", fixedCost.ID)
	 }
	*/
	outdatedPCBytes, _ := ctx.GetStub().GetState(pcID)

	outdatedPC := new(ProductionCost)

	err = json.Unmarshal(outdatedPCBytes, outdatedPC)

	if err != nil {
		return "", fmt.Errorf("could not unmarshal world state data to type Production Activity")
	}

	pc := &ProductionCost{
		DocType:     "pc",
		ID:          outdatedPC.ID,
		CO2e:        co2e,
		Description: outdatedPC.Description,
		Value:       outdatedPC.Value,
		Activity:    outdatedPC.Activity,
		CostTypeID:  outdatedPC.CostTypeID,
	}

	bytes, _ := json.Marshal(pc)
	if err != nil {
		return "", err
	}

	err = ctx.GetStub().PutState(pcID, bytes)

	if err != nil {
		return "", fmt.Errorf("failed to put world state: %v", err)

	}

	return fmt.Sprintf("%s Production Activity CO2 emissions updated successfully to %f", pcID, co2e), nil
}
func (c *ProjetoIvContract) UpdatePCDescription(ctx contractapi.TransactionContextInterface, pcID, description string) (string, error) {
	exists, err := c.PCExists(ctx, pcID)
	if err != nil {
		return "", fmt.Errorf("could not read from world state. %s", err)
	} else if !exists {
		return "", fmt.Errorf("the asset %s does not exist", pcID)
	}

	outdatedPCBytes, _ := ctx.GetStub().GetState(pcID)

	outdatedPC := new(ProductionCost)

	err = json.Unmarshal(outdatedPCBytes, outdatedPC)

	if err != nil {
		return "", fmt.Errorf("could not unmarshal world state data to type Production Activity")
	}

	pc := &ProductionCost{
		DocType:     "pc",
		ID:          outdatedPC.ID,
		CO2e:        outdatedPC.CO2e,
		Description: description,
		Value:       outdatedPC.Value,
		Activity:    outdatedPC.Activity,
		CostTypeID:  outdatedPC.CostTypeID,
	}

	bytes, _ := json.Marshal(pc)
	if err != nil {
		return "", err
	}

	err = ctx.GetStub().PutState(pcID, bytes)

	if err != nil {
		return "", fmt.Errorf("failed to put world state: %v", err)
	}
	return fmt.Sprintf("%s Production Activity CO2 emissions updated successfully to %s", pcID, description), nil

}
func (c *ProjetoIvContract) UpdatePCValue(ctx contractapi.TransactionContextInterface, pcID, value string) (string, error) {
	exists, err := c.PCExists(ctx, pcID)
	if err != nil {
		return "", fmt.Errorf("could not read from world state. %s", err)
	} else if !exists {
		return "", fmt.Errorf("the asset %s does not exist", pcID)
	}

	outdatedPCBytes, _ := ctx.GetStub().GetState(pcID)

	outdatedPC := new(ProductionCost)

	err = json.Unmarshal(outdatedPCBytes, outdatedPC)

	if err != nil {
		return "", fmt.Errorf("could not unmarshal world state data to type Production Activity")
	}

	pc := &ProductionCost{
		DocType:     "pc",
		ID:          outdatedPC.ID,
		CO2e:        outdatedPC.CO2e,
		Description: outdatedPC.Description,
		Value:       value,
		Activity:    outdatedPC.Activity,
		CostTypeID:  outdatedPC.CostTypeID,
	}

	bytes, _ := json.Marshal(pc)
	if err != nil {
		return "", err
	}

	err = ctx.GetStub().PutState(pcID, bytes)

	if err != nil {
		return "", fmt.Errorf("failed to put world state: %v", err)
	}
	return fmt.Sprintf("%s Production Activity CO2 emissions updated successfully to ", pcID), nil

}
func (c *ProjetoIvContract) UpdatePCActivity(ctx contractapi.TransactionContextInterface, pcID string, activity ProductionActivity) (string, error) {
	exists, err := c.PCExists(ctx, pcID)
	if err != nil {
		return "", fmt.Errorf("could not read from world state. %s", err)
	} else if !exists {
		return "", fmt.Errorf("the asset %s does not exist", pcID)
	}

	exists, err = c.PAExists(ctx, activity.ID)
	if err != nil {
		return "", fmt.Errorf("could not read from world state. %s", err)
	} else if !exists {
		return "", fmt.Errorf("the fixed cost [%s] doesn't exists", activity.ID)
	}

	outdatedPCBytes, _ := ctx.GetStub().GetState(pcID)

	outdatedPC := new(ProductionCost)

	err = json.Unmarshal(outdatedPCBytes, outdatedPC)

	if err != nil {
		return "", fmt.Errorf("could not unmarshal world state data to type Production Activity")
	}

	pc := &ProductionCost{
		DocType:     "pc",
		ID:          outdatedPC.ID,
		CO2e:        outdatedPC.CO2e,
		Description: outdatedPC.Description,
		Value:       outdatedPC.Value,
		Activity:    activity,
		CostTypeID:  outdatedPC.CostTypeID,
	}

	bytes, _ := json.Marshal(pc)
	if err != nil {
		return "", err
	}

	err = ctx.GetStub().PutState(pcID, bytes)

	if err != nil {
		return "", fmt.Errorf("failed to put world state: %v", err)
	}

	return fmt.Sprintf("%s Production Activity CO2 emissions updated successfully to %s", pcID, activity.ID), nil

}

func (c *ProjetoIvContract) UpdatePCCostTypeID(ctx contractapi.TransactionContextInterface, pcID, costTypeID string) (string, error) {
	exists, err := c.PCExists(ctx, pcID)
	if err != nil {
		return "", fmt.Errorf("could not read from world state. %s", err)
	} else if !exists {
		return "", fmt.Errorf("the asset %s does not exist", pcID)
	}

	outdatedPCBytes, _ := ctx.GetStub().GetState(pcID)

	outdatedPC := new(ProductionCost)

	err = json.Unmarshal(outdatedPCBytes, outdatedPC)

	if err != nil {
		return "", fmt.Errorf("could not unmarshal world state data to type Production Activity")
	}

	pc := &ProductionCost{
		DocType:     "pc",
		ID:          outdatedPC.ID,
		CO2e:        outdatedPC.CO2e,
		Description: outdatedPC.Description,
		Value:       outdatedPC.Value,
		Activity:    outdatedPC.Activity,
		CostTypeID:  costTypeID,
	}

	bytes, _ := json.Marshal(pc)
	if err != nil {
		return "", err
	}

	err = ctx.GetStub().PutState(pcID, bytes)

	if err != nil {
		return "", fmt.Errorf("failed to put world state: %v", err)
	}

	return fmt.Sprintf("%s Production Activity CO2 emissions updated successfully to %s", pcID, costTypeID), nil

}

// DeletePC deletes an instance of PC from the world state
func (c *ProjetoIvContract) DeletePC(ctx contractapi.TransactionContextInterface, pcID string) error {
	exists, err := c.PCExists(ctx, pcID)
	if err != nil {
		return fmt.Errorf("could not read from world state. %s", err)
	} else if !exists {
		return fmt.Errorf("the asset %s does not exist", pcID)
	}

	return ctx.GetStub().DelState(pcID)
}

// ------------------------- MONTHLY FIXED COST -----------------

// ProjetoIvExists returns true when asset with given ID exists in world state
func (c *ProjetoIvContract) MFCExists(ctx contractapi.TransactionContextInterface, mfcID string) (bool, error) {
	data, err := ctx.GetStub().GetState(mfcID)

	if err != nil {
		return false, err
	}

	return data != nil, nil
}

// CreateProjetoIv creates a new instance of MFC
func (c *ProjetoIvContract) CreateMFC(ctx contractapi.TransactionContextInterface, mfcID, month, description, organizationID string, year, quantity int, co2e float64) (string, error) {
	exists, err := c.MFCExists(ctx, mfcID)
	if err != nil {
		return "", fmt.Errorf("could not read from world state. %s", err)
	} else if exists {
		return "", fmt.Errorf("the asset %s already exists", mfcID)
	}
	/*
	 lot := new(Lot)
	 lot.Value = value */

	mfc := &MonthlyFixedCost{
		DocType:        "mfc",
		ID:             mfcID,
		CO2e:           co2e,
		Description:    description,
		Year:           year,
		Month:          month,
		Quantity:       quantity,
		OrganizationID: organizationID,
	}

	bytes, _ := json.Marshal(mfc)

	if err != nil {
		return "", err
	}

	err = ctx.GetStub().PutState(mfcID, bytes)

	if err != nil {
		return "", fmt.Errorf("failed to put to world state: %v", err)
	}

	return fmt.Sprintf("%s created sucessfully", mfcID), nil

}

// ReadMFC retrieves an instance of MFC from the world state
func (c *ProjetoIvContract) ReadMFC(ctx contractapi.TransactionContextInterface, mfcID string) (*MonthlyFixedCost, error) {
	exists, err := c.MFCExists(ctx, mfcID)
	if err != nil {
		return nil, fmt.Errorf("could not read from world state. %s", err)
	} else if !exists {
		return nil, fmt.Errorf("the asset %s does not exist", mfcID)
	}

	bytes, _ := ctx.GetStub().GetState(mfcID)

	mfc := new(MonthlyFixedCost)

	err = json.Unmarshal(bytes, mfc)

	if err != nil {
		return nil, fmt.Errorf("could not unmarshal world state data to type MFC")
	}

	return mfc, nil
}

// UpdateMFC retrieves an instance of MFC from the world state and updates its value
func (c *ProjetoIvContract) UpdateMFC(ctx contractapi.TransactionContextInterface, mfcID string, newValue string) (string, error) {
	exists, err := c.MFCExists(ctx, mfcID)
	if err != nil {
		return "", fmt.Errorf("could not read from world state. %s", err)
	} else if !exists {
		return "", fmt.Errorf("the asset %s does not exist", mfcID)
	}

	/* exists, err = c.PAExists(ctx, activity.ID)
	 if err != nil {
		 return "", fmt.Errorf("Could not read from world state. %s", err)
	 } else if !exists {
		 return "", fmt.Errorf("The fixed cost [%s] doesn't exists", activity.ID)
	 }
	*/
	outdatedMFCBytes, _ := ctx.GetStub().GetState(mfcID)

	outdatedMFC := new(MonthlyFixedCost)

	err = json.Unmarshal(outdatedMFCBytes, outdatedMFC)

	if err != nil {
		return "", fmt.Errorf("could not unmarshal world state data to type Production Activity")
	}

	mfc := &MonthlyFixedCost{
		DocType:        "mfc",
		ID:             mfcID,
		CO2e:           outdatedMFC.CO2e,
		Description:    outdatedMFC.Description,
		Year:           outdatedMFC.Year,
		Month:          outdatedMFC.Month,
		Quantity:       outdatedMFC.Quantity,
		OrganizationID: outdatedMFC.OrganizationID,
	}

	bytes, _ := json.Marshal(mfc)
	if err != nil {
		return "", err
	}

	err = ctx.GetStub().PutState(mfcID, bytes)

	if err != nil {
		return "", fmt.Errorf("failed to put world state: %v", err)
	}

	return fmt.Sprintf("%s Production Activity CO2 emissions updated successfully to ", mfcID), nil
}

func (c *ProjetoIvContract) UpdateMFCCO2e(ctx contractapi.TransactionContextInterface, mfcID string, co2e float64) (string, error) {
	exists, err := c.MFCExists(ctx, mfcID)
	if err != nil {
		return "", fmt.Errorf("could not read from world state. %s", err)
	} else if !exists {
		return "", fmt.Errorf("the asset %s does not exist", mfcID)
	}

	/* exists, err = c.PAExists(ctx, activity.ID)
	 if err != nil {
		 return "", fmt.Errorf("Could not read from world state. %s", err)
	 } else if !exists {
		 return "", fmt.Errorf("The fixed cost [%s] doesn't exists", activity.ID)
	 }
	*/
	outdatedMFCBytes, _ := ctx.GetStub().GetState(mfcID)

	outdatedMFC := new(MonthlyFixedCost)

	err = json.Unmarshal(outdatedMFCBytes, outdatedMFC)

	if err != nil {
		return "", fmt.Errorf("could not unmarshal world state data to type Production Activity")
	}

	mfc := &MonthlyFixedCost{
		DocType:     "mfc",
		ID:          outdatedMFC.ID,
		CO2e:        co2e,
		Description: outdatedMFC.Description,
		Year:        outdatedMFC.Year,
		Month:       outdatedMFC.Month,
		Quantity:    outdatedMFC.Quantity,
	}

	bytes, _ := json.Marshal(mfc)
	if err != nil {
		return "", err
	}

	err = ctx.GetStub().PutState(mfcID, bytes)

	if err != nil {
		return "", fmt.Errorf("failed to put world state: %v", err)
	}

	return fmt.Sprintf("%s Production Activity CO2 emissions updated successfully to ", mfcID), nil
}

func (c *ProjetoIvContract) UpdateMFCQuantity(ctx contractapi.TransactionContextInterface, mfcID string, quantity int) (string, error) {
	exists, err := c.MFCExists(ctx, mfcID)
	if err != nil {
		return "", fmt.Errorf("could not read from world state. %s", err)
	} else if !exists {
		return "", fmt.Errorf("the asset %s does not exist", mfcID)
	}

	/* exists, err = c.PAExists(ctx, activity.ID)
	 if err != nil {
		 return "", fmt.Errorf("Could not read from world state. %s", err)
	 } else if !exists {
		 return "", fmt.Errorf("The fixed cost [%s] doesn't exists", activity.ID)
	 }
	*/
	outdatedMFCBytes, _ := ctx.GetStub().GetState(mfcID)

	outdatedMFC := new(MonthlyFixedCost)

	err = json.Unmarshal(outdatedMFCBytes, outdatedMFC)

	if err != nil {
		return "", fmt.Errorf("could not unmarshal world state data to type Production Activity")
	}

	mfc := &MonthlyFixedCost{
		DocType:        "mfc",
		ID:             outdatedMFC.ID,
		CO2e:           outdatedMFC.CO2e,
		Description:    outdatedMFC.Description,
		Year:           outdatedMFC.Year,
		Month:          outdatedMFC.Month,
		Quantity:       quantity,
		OrganizationID: outdatedMFC.OrganizationID,
	}

	bytes, _ := json.Marshal(mfc)
	if err != nil {
		return "", err
	}

	err = ctx.GetStub().PutState(mfcID, bytes)

	if err != nil {
		return "", fmt.Errorf("failed to put world state: %v", err)
	}

	return fmt.Sprintf("%s Production Activity CO2 emissions updated successfully to ", mfcID), nil
}

func (c *ProjetoIvContract) UpdateMFCDescription(ctx contractapi.TransactionContextInterface, mfcID, description string) (string, error) {
	exists, err := c.MFCExists(ctx, mfcID)
	if err != nil {
		return "", fmt.Errorf("could not read from world state. %s", err)
	} else if !exists {
		return "", fmt.Errorf("the asset %s does not exist", mfcID)
	}

	/* exists, err = c.PAExists(ctx, activity.ID)
	 if err != nil {
		 return "", fmt.Errorf("Could not read from world state. %s", err)
	 } else if !exists {
		 return "", fmt.Errorf("The fixed cost [%s] doesn't exists", activity.ID)
	 }
	*/
	outdatedMFCBytes, _ := ctx.GetStub().GetState(mfcID)

	outdatedMFC := new(MonthlyFixedCost)

	err = json.Unmarshal(outdatedMFCBytes, outdatedMFC)

	if err != nil {
		return "", fmt.Errorf("could not unmarshal world state data to type Production Activity")
	}

	mfc := &MonthlyFixedCost{
		DocType:        "mfc",
		ID:             outdatedMFC.ID,
		CO2e:           outdatedMFC.CO2e,
		Description:    description,
		Year:           outdatedMFC.Year,
		Month:          outdatedMFC.Month,
		Quantity:       outdatedMFC.Quantity,
		OrganizationID: outdatedMFC.OrganizationID,
	}

	bytes, _ := json.Marshal(mfc)
	if err != nil {
		return "", err
	}

	err = ctx.GetStub().PutState(mfcID, bytes)

	if err != nil {
		return "", fmt.Errorf("failed to put world state: %v", err)
	}

	return fmt.Sprintf("%s Production Activity CO2 emissions updated successfully to ", mfcID), nil
}
func (c *ProjetoIvContract) UpdateMFCorganizationID(ctx contractapi.TransactionContextInterface, mfcID, organizationID string) (string, error) {
	exists, err := c.MFCExists(ctx, mfcID)
	if err != nil {
		return "", fmt.Errorf("could not read from world state. %s", err)
	} else if !exists {
		return "", fmt.Errorf("the asset %s does not exist", mfcID)
	}

	/* exists, err = c.PAExists(ctx, activity.ID)
	 if err != nil {
		 return "", fmt.Errorf("Could not read from world state. %s", err)
	 } else if !exists {
		 return "", fmt.Errorf("The fixed cost [%s] doesn't exists", activity.ID)
	 }
	*/
	outdatedMFCBytes, _ := ctx.GetStub().GetState(mfcID)

	outdatedMFC := new(MonthlyFixedCost)

	err = json.Unmarshal(outdatedMFCBytes, outdatedMFC)

	if err != nil {
		return "", fmt.Errorf("could not unmarshal world state data to type Production Activity")
	}

	mfc := &MonthlyFixedCost{
		DocType:        "mfc",
		ID:             outdatedMFC.ID,
		CO2e:           outdatedMFC.CO2e,
		Description:    outdatedMFC.Description,
		Year:           outdatedMFC.Year,
		Month:          outdatedMFC.Month,
		Quantity:       outdatedMFC.Quantity,
		OrganizationID: organizationID,
	}

	bytes, _ := json.Marshal(mfc)
	if err != nil {
		return "", err
	}

	err = ctx.GetStub().PutState(mfcID, bytes)

	if err != nil {
		return "", fmt.Errorf("failed to put world state: %v", err)
	}

	return fmt.Sprintf("%s Production Activity CO2 emissions updated successfully to ", mfcID), nil
}

// DeleteMFC deletes an instance of MFC from the world state
func (c *ProjetoIvContract) DeleteMFC(ctx contractapi.TransactionContextInterface, mfcID string) error {
	exists, err := c.MFCExists(ctx, mfcID)
	if err != nil {
		return fmt.Errorf("could not read from world state. %s", err)
	} else if !exists {
		return fmt.Errorf("the asset %s does not exist", mfcID)
	}

	return ctx.GetStub().DelState(mfcID)
}
