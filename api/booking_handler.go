package api

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/rojerdu-dev/hotel-reservation/db"
	"github.com/rojerdu-dev/hotel-reservation/types"
	"go.mongodb.org/mongo-driver/bson"
)

type BookingHandler struct {
	store *db.Store
}

func NewBookingHandler(store *db.Store) *BookingHandler {
	return &BookingHandler{
		store: store,
	}
}

func (h *BookingHandler) HandleCancelBooking(c *fiber.Ctx) error {
	id := c.Params("id")
	booking, err := h.store.Booking.GetBookingByID(c.Context(), id)
	if err != nil {
		return err
	}
	user, err := getAuthUser(c)
	if err != nil {
		return err
	}
	if booking.UserID != user.ID {
		return c.Status(http.StatusUnauthorized).JSON(genericResp{
			Type: "error",
			Msg:  "unauthorized",
		})
	}
	if err = h.store.Booking.UpdateBooking(c.Context(), c.Params("id"), bson.M{"canceled": true}); err != nil {
		return err
	}
	return c.JSON(genericResp{
		Type: "message",
		Msg:  "updated",
	})
}

func (h *BookingHandler) HandleGetBookings(c *fiber.Ctx) error {
	bookings, err := h.store.Booking.GetBookings(c.Context(), bson.M{})
	if err != nil {
		return err
	}
	return c.JSON(bookings)
}

func (h *BookingHandler) HandleGetBooking(c *fiber.Ctx) error {
	id := c.Params("id")
	booking, err := h.store.Booking.GetBookingByID(c.Context(), id)
	if err != nil {
		return nil
	}
	user, ok := c.Context().UserValue("user").(*types.User)
	if !ok {
		return err
	}
	if booking.UserID != user.ID {
		return c.Status(http.StatusUnauthorized).JSON(genericResp{
			Type: "error",
			Msg:  "unauthorized",
		})
	}
	return c.JSON(booking)
}
