package security

import (
	"net/http"

	"crypto/x509"

	"encoding/pem"

	"github.com/fokal/fokal/pkg/handler"
	"github.com/fokal/fokal/pkg/tokens"
)

func PublicKeyHandler(state *handler.State, w http.ResponseWriter, r *http.Request) (handler.Response, error) {
	key := make(map[string]string)
	keyBytes, err := x509.MarshalPKIXPublicKey(state.PublicKeys[state.KeyHash])
	if err != nil {
		return handler.Response{}, handler.StatusError{Code: http.StatusInternalServerError}
	}

	pemBytes := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: keyBytes,
	})

	key[state.KeyHash] = string(pemBytes)
	return handler.Response{Code: http.StatusOK, Data: key}, nil
}

func RefreshHandler(state *handler.State, w http.ResponseWriter, r *http.Request) (handler.Response, error) {
	user, err := tokens.Verify(state, r)
	if err != nil {
		return handler.Response{}, handler.StatusError{Code: http.StatusBadRequest, Err: err}
	}
	jwt, err := tokens.Create(state, user)
	if err != nil {
		return handler.Response{}, handler.StatusError{Code: http.StatusBadRequest, Err: err}
	}
	return handler.Response{Code: http.StatusOK, Data: map[string]string{"token": jwt}}, nil
}
