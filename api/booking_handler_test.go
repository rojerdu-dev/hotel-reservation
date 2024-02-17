package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/rojerdu-dev/hotel-reservation/db/fixtures"
	"github.com/rojerdu-dev/hotel-reservation/types"
)

func TestUserGetBooking(t *testing.T) {
	db := setup(t)
	defer db.teardown(t)

	var (
		nonAuthUser    = fixtures.AddUser(db.Store, "jimmy", "connors", false)
		user           = fixtures.AddUser(db.Store, "james", "foo", false)
		hotel          = fixtures.AddHotel(db.Store, "Hotel Bar & Grill", "New York City", 5, nil)
		room           = fixtures.AddRoom(db.Store, "small", true, 4.4, hotel.ID)
		from           = time.Now()
		until          = from.AddDate(0, 0, 5)
		booking        = fixtures.AddBooking(db.Store, user.ID, room.ID, from, until)
		app            = fiber.New()
		route          = app.Group("/", JWTAuthentication(db.User))
		bookingHandler = NewBookingHandler(db.Store)
	)
	route.Get("/:id", bookingHandler.HandleGetBooking)
	req := httptest.NewRequest("GET", fmt.Sprintf("/%s", booking.ID.Hex()), nil)
	req.Header.Add("X-Api-Token", CreateTokenFromUser(user))
	resp, err := app.Test(req)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("non 200 code, got %d", resp.StatusCode)
	}
	var bookingResp *types.Booking
	if err = json.NewDecoder(resp.Body).Decode(&bookingResp); err != nil {
		t.Fatal(err)
	}
	if bookingResp.ID != booking.ID {
		t.Fatalf("expected %s got %s", booking.ID, bookingResp.ID)
	}
	if bookingResp.UserID != booking.UserID {
		t.Fatalf("expected %s got %s", booking.UserID, bookingResp.UserID)
	}
	req = httptest.NewRequest("GET", fmt.Sprintf("/%s", booking.ID.Hex()), nil)
	req.Header.Add("X-Api-Token", CreateTokenFromUser(nonAuthUser))
	resp, err = app.Test(req)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		t.Fatalf("expected non 200 status code but got %d", resp.StatusCode)
	}
}

func TestAdminGetBookings(t *testing.T) {
	db := setup(t)
	defer db.teardown(t)

	var (
		adminUser      = fixtures.AddUser(db.Store, "admin", "admin", true)
		user           = fixtures.AddUser(db.Store, "james", "foo", false)
		hotel          = fixtures.AddHotel(db.Store, "Hotel Bar & Grill", "New York City", 5, nil)
		room           = fixtures.AddRoom(db.Store, "small", true, 4.4, hotel.ID)
		from           = time.Now()
		until          = from.AddDate(0, 0, 5)
		booking        = fixtures.AddBooking(db.Store, user.ID, room.ID, from, until)
		app            = fiber.New(fiber.Config{ErrorHandler: ErrorHandler})
		admin          = app.Group("/", JWTAuthentication(db.User), AdminAuth)
		bookingHandler = NewBookingHandler(db.Store)
	)

	_ = booking
	admin.Get("/", bookingHandler.HandleGetBookings)
	req := httptest.NewRequest("GET", "/", nil)
	req.Header.Add("X-Api-Token", CreateTokenFromUser(adminUser))
	resp, err := app.Test(req)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("non 200 response got %d", resp.StatusCode)
	}
	var bookings []*types.Booking
	if err := json.NewDecoder(resp.Body).Decode(&bookings); err != nil {
		t.Fatal(err)
	}
	if len(bookings) != 1 {
		t.Fatalf("expected 1 booking got %d", len(bookings))
	}
	have := bookings[0]
	if have.ID != booking.ID {
		t.Fatalf("expected %s got %s", booking.ID, have.ID)
	}
	if have.UserID != booking.UserID {
		t.Fatalf("expected %s got %s", booking.UserID, have.UserID)
	}

	// test whether non-admin cannot access bookings
	req = httptest.NewRequest("GET", "/", nil)
	req.Header.Add("X-Api-Token", CreateTokenFromUser(user))
	resp, err = app.Test(req)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusUnauthorized {
		t.Fatalf("expected status unauthorized but got %d", resp.StatusCode)
	}
}
