package api

import (
	"fmt"
	"github.com/rojerdu-dev/hotel-reservation/db/fixtures"
	"testing"
	"time"
)

func TestAdminGetBookings(t *testing.T) {
	db := setup(t)
	defer db.teardown(t)

	user := fixtures.AddUser(db.Store, "james", "foo", false)
	hotel := fixtures.AddHotel(db.Store, "Hotel Bar & Grill", "New York City", 5, nil)
	room := fixtures.AddRoom(db.Store, "small", true, 4.4, hotel.ID)

	from := time.Now()
	until := from.AddDate(0, 0, 5)
	booking := fixtures.AddBooking(db.Store, user.ID, room.ID, from, until)
	fmt.Println(booking)
}
