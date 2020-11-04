package routes

import (
	"S3_FriendManagement_ThinhNguyen/handlers"
	"S3_FriendManagement_ThinhNguyen/repositories"
	"S3_FriendManagement_ThinhNguyen/service"
	"database/sql"
	"github.com/go-chi/chi"
	"net/http"
)

func CreateRoutes(db *sql.DB) *chi.Mux {
	r := chi.NewRouter()

	//Routes for user
	r.Route("/user", func(r chi.Router) {
		UserHandler := handlers.UserHandler{
			IUserService: service.UserService{
				IUserRepo: repositories.UserRepo{
					Db: db,
				},
			},
		}
		r.MethodFunc(http.MethodPost, "/", UserHandler.CreateUser)
	})

	//Routes for Friend
	r.Route("/friend", func(r chi.Router) {
		FriendHandler := handlers.FriendHandler{
			IUserService: service.UserService{
				IUserRepo: repositories.UserRepo{
					Db: db,
				},
			},
			IFriendServices: service.FriendService{
				IFriendRepo: repositories.FriendRepo{
					Db: db,
				},
			},
		}
		r.MethodFunc(http.MethodPost, "/", FriendHandler.CreateFriend)
	})
	return r
}
