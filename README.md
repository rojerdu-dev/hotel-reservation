# Hotel Reservations Backend
Welcome to my Hotel Reservations project!

This project is a backend system for a hotel's reservation booking system developed in Go (Fiber framework) and using MongoDB as the datastore.

The RESTful API handles Create, Read, Update, and Delete (CRUD) operations associated with users booking room reservations. Final application is Dockerized.


## Project outline
- users -> book room from an hotel 
- admins -> going to check reservation/bookings 
- Authentication and authorization -> JWT tokens
- Hotels -> CRUD API -> JSON
- Rooms -> CRUD API -> JSON
- Scripts -> database management -> seeding, migration

## Resources
### Mongodb driver 
Documentation
```
https://mongodb.com/docs/drivers/go/current/quick-start
```

Installing mongodb client
```
go get go.mongodb.org/mongo-driver/mongo
```

### gofiber 
Documentation
```
https://gofiber.io
```

Installing gofiber
```
go get github.com/gofiber/fiber/v2
```

## Docker
### Installing mongodb as a Docker container
```
docker run --name mongodb -d mongo:latest -p 27017:27017
```
