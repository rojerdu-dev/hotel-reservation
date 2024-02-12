package api

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/rojerdu-dev/hotel-reservation/db"
	"github.com/rojerdu-dev/hotel-reservation/types"
)

func TestHandleAuthenticate(t *testing.T) {
	tdb := setup(t)
	defer tdb.teardown(t)
	insertTestUser(t, tdb.UserStore)

	app := fiber.New()
	authHandler := NewAuthHandler(tdb.UserStore)
	app.Post("/auth", authHandler.HandleAuthenticate)

	params := AuthParams{
		Email:    "ldicaprio@titanic.com",
		Password: "securepassword",
	}

	b, _ := json.Marshal(params)
	req := httptest.NewRequest("POST", "/auth", bytes.NewReader(b))
	req.Header.Set("Content-Type", "application/json")
	resp, err := app.Test(req)
	if err != nil {
		t.Error(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected http status of 200 but got %d", resp.StatusCode)
	}

	var authResp AuthResponse
	if err := json.NewDecoder(resp.Body).Decode(&authResp); err != nil {
		t.Error(err)
	}

	fmt.Println(resp)
}

func insertTestUser(t *testing.T, userStore db.UserStore) *types.User {
	user, err := types.NewUserFromParams(types.CreateUserParams{
		FirstName: "Leonardo",
		LastName:  "DiCaprio",
		Email:     "ldicaprio@titanic.com",
		Password:  "securepassword",
	})
	if err != nil {
		t.Fatal(err)
	}

	user, err = userStore.InsertUser(context.TODO(), user)
	if err != nil {
		t.Fatal(err)
	}

	return user
}
