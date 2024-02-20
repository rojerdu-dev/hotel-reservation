package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rojerdu-dev/hotel-reservation/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type HotelHandler struct {
	store      *db.Store
	hotelStore db.HotelStore
	roomStore  db.RoomStore
}

func NewHotelHandler(store *db.Store) *HotelHandler {
	return &HotelHandler{
		store: store,
	}
}

func (h *HotelHandler) HandleGetRooms(c *fiber.Ctx) error {
	id := c.Params("id")
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return ErrInvalidID()
	}
	filter := bson.M{"hotelID": oid}
	rooms, err := h.store.Room.GetRooms(c.Context(), filter)
	if err != nil {
		return ErrorResourceNotFound("hotel")
	}
	return c.JSON(rooms)
}

func (h *HotelHandler) HandleGetHotel(c *fiber.Ctx) error {
	id := c.Params("id")
	hotel, err := h.store.Hotel.GetHotelByID(c.Context(), id)
	if err != nil {
		return ErrorResourceNotFound("hotel")
	}
	return c.JSON(hotel)
}

type ResourceResp struct {
	Results int `json:"results"`
	Data    any `json:"data"`
	Page    int `json:"page"`
}

type HotelQueryParams struct {
	db.Pagination
	Rating int
}

func (h *HotelHandler) HandleGetHotels(c *fiber.Ctx) error {
	var params HotelQueryParams
	if err := c.QueryParser(&params); err != nil {
		return ErrorBadRequest()
	}
	filter := db.Map{}
	if params.Rating != 0 {
		filter["rating"] = params.Rating
	}

	hotels, err := h.store.Hotel.GetHotels(c.Context(), filter, &params.Pagination)
	if err != nil {
		return ErrorResourceNotFound("hotels")
	}
	resp := ResourceResp{
		Results: len(hotels),
		Data:    hotels,
		Page:    int(params.Page),
	}

	return c.JSON(resp)
}
