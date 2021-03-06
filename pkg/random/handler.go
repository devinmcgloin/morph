package random

import (
	"net/http"

	"github.com/fokal/fokal-core/pkg/handler"
	"github.com/fokal/fokal-core/pkg/retrieval"
)

func ImageHandler(store *handler.State, w http.ResponseWriter, r *http.Request) (handler.Response, error) {
	var userID *int64
	username, ok := r.URL.Query()["username"]
	if ok {
		ref, err := retrieval.GetUserRef(store.DB, username[0])
		if err != nil {
			return handler.Response{}, err
		}
		userID = &ref.Id
	}

	image, err := Image(store, userID)
	if err != nil {
		return handler.Response{}, err
	}

	return handler.Response{Code: http.StatusOK, Data: image}, nil

}
