package api

import (
	"log"
	"net/http"

	"github.com/BerkCicekler/shoe-api/repository"
	"github.com/BerkCicekler/shoe-api/service/user"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

type APIServer struct {
	addr string
}

func NewAPIServer(addr string) *APIServer{
	return &APIServer{
		addr: addr,
	}
}

func (s *APIServer) Run(mongoDatabase *mongo.Database) error {
	router:= mux.NewRouter()
	subRouter:= router.PathPrefix("/api/v1").Subrouter()


	userRepository := repository.UsersRepo{
		MongoCollection: mongoDatabase.Collection("users"),
	}
	userHandler := user.UserServiceNewHandler(userRepository)
	userHandler.RegisterRoutes(subRouter)

	log.Println("Listening on", s.addr)


	return http.ListenAndServe(s.addr, router)
}