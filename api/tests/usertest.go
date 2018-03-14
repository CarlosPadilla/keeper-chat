package tests

import (
	"github.com/revel/revel/testing"
	"strings"
)

type UserTest struct {
	testing.TestSuite
}

func (t *UserTest) Before() {
	println("Set up")
}

func (t *UserTest) TheRightPost() {
	json := strings.NewReader(`{"name": "Peteco", "email": "peteco@mail.com"}`)

	t.Post("/users", "application/json", json)
	t.AssertOk()
	t.AssertContentType("application/json; charset=utf-8")

}

func (t *UserTest) TheIncompletePost() {
	json := strings.NewReader(`{"email": "peteco@mail.com"}`)

	t.Post("/users", "application/json", json)
	t.AssertOk()
	t.AssertContentType("application/json; charset=utf-8")

}

func (t *UserTest) After() {
	println("Tear down")
}
