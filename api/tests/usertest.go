package tests

import (
	"github.com/revel/revel/testing"
	"strings"
	"net/http"
	"encoding/json"
	_ "fmt"
	"github.com/Zeloid/keeper-chat/api/app/models"
	"regexp"
	"fmt"
	"github.com/Zeloid/keeper-chat/api/app/components"
)

type UserTest struct {
	testing.TestSuite
}

/**
 * Test an array of errors for the presence of a certain error
 */
func (t *UserTest) AssertValidationError(response components.ValidationErrorResponse, key string, message string) {
	found := 0
	for i:=0; i < len(response.Errors); i++ {
		if response.Errors[i].Key == key && response.Errors[i].Message == message {
			found++
		}
	}

	if found == 0 {
		t.Assertf(false, "Error %v:%v not found in response", key, message)
	} else if found > 1 {
		t.Assertf(false, "Error %v:%v found multiple times in response (%d)", key, message, found)
	}
}

func (t *UserTest) AssertUUID (uuid string) {
	match, _ := regexp.MatchString("[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}", uuid)
	t.Assertf(match, "Malformed UUID: %v", uuid)
}

func (t *UserTest) AssertEntityUUID (uuid string, entityName string, tableId int8) {
	r := fmt.Sprintf("00000000-%04x-0000-[0-9a-f]{4}-[0-9a-f]{12}", tableId)
	match, _ := regexp.MatchString(r, uuid)
	t.Assertf(match, "Invalid %v UUID: %v", entityName, uuid)
}

func (t *UserTest) Before() {
	println("Set up")
}

func (t *UserTest) TestCorrectPost() {
	var responseJson models.Account

	requestBody := strings.NewReader(`{"name": "Peteco", "email": "peteco@mail.com"}`)

	t.Post("/users", "application/json", requestBody)
	t.AssertOk()
	t.AssertContentType("application/json; charset=utf-8")

	err := json.Unmarshal(t.ResponseBody, &responseJson)

	t.Assertf(err == nil, "Invalid Json error response: %q", err)

	t.AssertEqual(responseJson.Status, models.EnumAccountStatus.Pending)
	t.AssertEqual("peteco@mail.com", responseJson.Email)
	t.AssertEqual("Peteco", responseJson.Name)
	t.AssertUUID(responseJson.Id)
	t.AssertEntityUUID(responseJson.Id, "Account", 1)
}

func (t *UserTest) TestEmptyPost() {
	var responseJson components.ValidationErrorResponse
	const expectedErrorCount = 4

	requestBody := strings.NewReader(`{}`)

	t.Post("/users", "application/json", requestBody)
	t.AssertStatus(http.StatusBadRequest)
	t.AssertContentType("application/json; charset=utf-8")

	err := json.Unmarshal(t.ResponseBody, &responseJson)

	t.Assertf(err == nil, "Invalid Json error response: %q", err)

	t.Assertf(len(responseJson.Errors) == expectedErrorCount, "Invalid number errors reported(%d), %d expected", len(responseJson.Errors), expectedErrorCount)
	//fmt.Printf("Response of testing server is %q", t.ResponseBody)

	t.AssertValidationError(responseJson, "account.Name", "Required")
	t.AssertValidationError(responseJson, "account.Name", "Minimum size is 4\n")
	t.AssertValidationError(responseJson, "account.Email", "Required")
	t.AssertValidationError(responseJson, "account.Email", "Must be a valid email address\n")
}

func (t *UserTest) TestInvalidEmailPost() {
	var responseJson components.ValidationErrorResponse
	const expectedErrorCount = 1

	requestBody := strings.NewReader(`{"name": "Peteco", "email": "peteco@mailcom"}`)

	t.Post("/users", "application/json", requestBody)
	t.AssertStatus(http.StatusBadRequest)
	t.AssertContentType("application/json; charset=utf-8")

	err := json.Unmarshal(t.ResponseBody, &responseJson)

	t.Assertf(err == nil, "Invalid Json error response: %q", err)

	t.Assertf(len(responseJson.Errors) == expectedErrorCount, "Invalid number errors reported(%d), %d expected", len(responseJson.Errors), expectedErrorCount)
	fmt.Printf("Response of testing server is %q", t.ResponseBody)

	t.AssertValidationError(responseJson, "account.Email", "Must be a valid email address\n")
}

func (t *UserTest) After() {
	println("Tear down")
}
