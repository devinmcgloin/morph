package routes

import (
	"github.com/fokal/fokal-core/pkg/create"
	"github.com/fokal/fokal-core/pkg/handler"
	"github.com/fokal/fokal-core/pkg/security"
	"github.com/gorilla/mux"
	"github.com/justinas/alice"
)

func RegisterCreateRoutes(state *handler.State, api *mux.Router, chain alice.Chain) {
	post := api.Methods("POST").Subrouter()
	opts := api.Methods("OPTIONS").Subrouter()

	post.Handle("/images", chain.Append(handler.Middleware{
		State: state,
		M:     security.Authenticate}.Handler).Then(handler.Handler{State: state, H: create.ImageHandler}))
	opts.Handle("/images", chain.Then(handler.Options("POST")))

	post.Handle("/users", chain.Then(handler.Handler{
		State: state,
		H:     create.UserHandler,
	}))

	opts.Handle("/users", chain.Then(handler.Options("POST")))

	put := api.Methods("PUT").Subrouter()
	put.Handle("/users/me/avatar", chain.Append(
		handler.Middleware{
			State: state,
			M:     security.Authenticate,
		}.Handler).Then(handler.Handler{
		State: state,
		H:     create.AvatarHandler,
	}))
	opts.Handle("/users/me/avatar", chain.Then(handler.Options("PUT")))

}
