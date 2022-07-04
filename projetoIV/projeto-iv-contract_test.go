/*
 * SPDX-License-Identifier: Apache-2.0
 */

package main

import (
	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	"github.com/stretchr/testify/mock"
)

const getStateError = "world state get error"

type MockStub struct {
	shim.ChaincodeStubInterface
	mock.Mock
}

func (ms *MockStub) GetState(key string) ([]byte, error) {
	args := ms.Called(key)

	return args.Get(0).([]byte), args.Error(1)
}

func (ms *MockStub) PutState(key string, value []byte) error {
	args := ms.Called(key, value)

	return args.Error(0)
}

func (ms *MockStub) DelState(key string) error {
	args := ms.Called(key)

	return args.Error(0)
}

type MockContext struct {
	contractapi.TransactionContextInterface
	mock.Mock
}

func (mc *MockContext) GetStub() shim.ChaincodeStubInterface {
	args := mc.Called()

	return args.Get(0).(*MockStub)
}

/*
func configureStub() (*MockContext, *MockStub) {
	var nilBytes []byte

	testProjetoIv := new(ProjetoIv)
	testProjetoIv.Value = "set value"
	projetoIvBytes, _ := json.Marshal(testProjetoIv)

	ms := new(MockStub)
	ms.On("GetState", "statebad").Return(nilBytes, errors.New(getStateError))
	ms.On("GetState", "missingkey").Return(nilBytes, nil)
	ms.On("GetState", "existingkey").Return([]byte("some value"), nil)
	ms.On("GetState", "projetoIvkey").Return(projetoIvBytes, nil)
	ms.On("PutState", mock.AnythingOfType("string"), mock.AnythingOfType("[]uint8")).Return(nil)
	ms.On("DelState", mock.AnythingOfType("string")).Return(nil)

	mc := new(MockContext)
	mc.On("GetStub").Return(ms)

	return mc, ms
}

func TestProjetoIvExists(t *testing.T) {
	var exists bool
	var err error

	ctx, _ := configureStub()
	c := new(ProjetoIvContract)

	exists, err = c.ProjetoIvExists(ctx, "statebad")
	assert.EqualError(t, err, getStateError)
	assert.False(t, exists, "should return false on error")

	exists, err = c.ProjetoIvExists(ctx, "missingkey")
	assert.Nil(t, err, "should not return error when can read from world state but no value for key")
	assert.False(t, exists, "should return false when no value for key in world state")

	exists, err = c.ProjetoIvExists(ctx, "existingkey")
	assert.Nil(t, err, "should not return error when can read from world state and value exists for key")
	assert.True(t, exists, "should return true when value for key in world state")
}

func TestCreateProjetoIv(t *testing.T) {
	var err error

	ctx, stub := configureStub()
	c := new(ProjetoIvContract)

	err = c.CreateProjetoIv(ctx, "statebad", "some value")
	assert.EqualError(t, err, fmt.Sprintf("Could not read from world state. %s", getStateError), "should error when exists errors")

	err = c.CreateProjetoIv(ctx, "existingkey", "some value")
	assert.EqualError(t, err, "The asset existingkey already exists", "should error when exists returns true")

	err = c.CreateProjetoIv(ctx, "missingkey", "some value")
	stub.AssertCalled(t, "PutState", "missingkey", []byte("{\"value\":\"some value\"}"))
}

func TestReadProjetoIv(t *testing.T) {
	var projetoIv *ProjetoIv
	var err error

	ctx, _ := configureStub()
	c := new(ProjetoIvContract)

	projetoIv, err = c.ReadProjetoIv(ctx, "statebad")
	assert.EqualError(t, err, fmt.Sprintf("Could not read from world state. %s", getStateError), "should error when exists errors when reading")
	assert.Nil(t, projetoIv, "should not return ProjetoIv when exists errors when reading")

	projetoIv, err = c.ReadProjetoIv(ctx, "missingkey")
	assert.EqualError(t, err, "The asset missingkey does not exist", "should error when exists returns true when reading")
	assert.Nil(t, projetoIv, "should not return ProjetoIv when key does not exist in world state when reading")

	projetoIv, err = c.ReadProjetoIv(ctx, "existingkey")
	assert.EqualError(t, err, "Could not unmarshal world state data to type ProjetoIv", "should error when data in key is not ProjetoIv")
	assert.Nil(t, projetoIv, "should not return ProjetoIv when data in key is not of type ProjetoIv")

	projetoIv, err = c.ReadProjetoIv(ctx, "projetoIvkey")
	expectedProjetoIv := new(ProjetoIv)
	expectedProjetoIv.Value = "set value"
	assert.Nil(t, err, "should not return error when ProjetoIv exists in world state when reading")
	assert.Equal(t, expectedProjetoIv, projetoIv, "should return deserialized ProjetoIv from world state")
}

func TestUpdateProjetoIv(t *testing.T) {
	var err error

	ctx, stub := configureStub()
	c := new(ProjetoIvContract)

	err = c.UpdateProjetoIv(ctx, "statebad", "new value")
	assert.EqualError(t, err, fmt.Sprintf("Could not read from world state. %s", getStateError), "should error when exists errors when updating")

	err = c.UpdateProjetoIv(ctx, "missingkey", "new value")
	assert.EqualError(t, err, "The asset missingkey does not exist", "should error when exists returns true when updating")

	err = c.UpdateProjetoIv(ctx, "projetoIvkey", "new value")
	expectedProjetoIv := new(ProjetoIv)
	expectedProjetoIv.Value = "new value"
	expectedProjetoIvBytes, _ := json.Marshal(expectedProjetoIv)
	assert.Nil(t, err, "should not return error when ProjetoIv exists in world state when updating")
	stub.AssertCalled(t, "PutState", "projetoIvkey", expectedProjetoIvBytes)
}

func TestDeleteProjetoIv(t *testing.T) {
	var err error

	ctx, stub := configureStub()
	c := new(ProjetoIvContract)

	err = c.DeleteProjetoIv(ctx, "statebad")
	assert.EqualError(t, err, fmt.Sprintf("Could not read from world state. %s", getStateError), "should error when exists errors")

	err = c.DeleteProjetoIv(ctx, "missingkey")
	assert.EqualError(t, err, "The asset missingkey does not exist", "should error when exists returns true when deleting")

	err = c.DeleteProjetoIv(ctx, "projetoIvkey")
	assert.Nil(t, err, "should not return error when ProjetoIv exists in world state when deleting")
	stub.AssertCalled(t, "DelState", "projetoIvkey")
}
*/
