// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package domain

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	"sync"
)

var (
	lockAuthenticationServiceMockCreateToken  sync.RWMutex
	lockAuthenticationServiceMockParseToken   sync.RWMutex
	lockAuthenticationServiceMockPublicKey    sync.RWMutex
	lockAuthenticationServiceMockRefreshToken sync.RWMutex
	lockAuthenticationServiceMockVerifyToken  sync.RWMutex
)

// AuthenticationServiceMock is a mock implementation of AuthenticationService.
//
//     func TestSomethingThatUsesAuthenticationService(t *testing.T) {
//
//         // make and configure a mocked AuthenticationService
//         mockedAuthenticationService := &AuthenticationServiceMock{
//             CreateTokenFunc: func(ctx context.Context, userID uint64) (*string, error) {
// 	               panic("TODO: mock out the CreateToken method")
//             },
//             ParseTokenFunc: func(ctx context.Context, token string) (*jwt.Token, error) {
// 	               panic("TODO: mock out the ParseToken method")
//             },
//             PublicKeyFunc: func(ctx context.Context) (string, error) {
// 	               panic("TODO: mock out the PublicKey method")
//             },
//             RefreshTokenFunc: func(ctx context.Context, token string) (*string, error) {
// 	               panic("TODO: mock out the RefreshToken method")
//             },
//             VerifyTokenFunc: func(ctx context.Context, token string) (bool, *uint64, error) {
// 	               panic("TODO: mock out the VerifyToken method")
//             },
//         }
//
//         // TODO: use mockedAuthenticationService in code that requires AuthenticationService
//         //       and then make assertions.
//
//     }
type AuthenticationServiceMock struct {
	// CreateTokenFunc mocks the CreateToken method.
	CreateTokenFunc func(ctx context.Context, userID uint64) (*string, error)

	// ParseTokenFunc mocks the ParseToken method.
	ParseTokenFunc func(ctx context.Context, token string) (*jwt.Token, error)

	// PublicKeyFunc mocks the PublicKey method.
	PublicKeyFunc func(ctx context.Context) (string, error)

	// RefreshTokenFunc mocks the RefreshToken method.
	RefreshTokenFunc func(ctx context.Context, token string) (*string, error)

	// VerifyTokenFunc mocks the VerifyToken method.
	VerifyTokenFunc func(ctx context.Context, token string) (bool, *uint64, error)

	// calls tracks calls to the methods.
	calls struct {
		// CreateToken holds details about calls to the CreateToken method.
		CreateToken []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// UserID is the userID argument value.
			UserID uint64
		}
		// ParseToken holds details about calls to the ParseToken method.
		ParseToken []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Token is the token argument value.
			Token string
		}
		// PublicKey holds details about calls to the PublicKey method.
		PublicKey []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
		}
		// RefreshToken holds details about calls to the RefreshToken method.
		RefreshToken []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Token is the token argument value.
			Token string
		}
		// VerifyToken holds details about calls to the VerifyToken method.
		VerifyToken []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Token is the token argument value.
			Token string
		}
	}
}

// CreateToken calls CreateTokenFunc.
func (mock *AuthenticationServiceMock) CreateToken(ctx context.Context, userID uint64) (*string, error) {
	if mock.CreateTokenFunc == nil {
		panic("AuthenticationServiceMock.CreateTokenFunc: method is nil but AuthenticationService.CreateToken was just called")
	}
	callInfo := struct {
		Ctx    context.Context
		UserID uint64
	}{
		Ctx:    ctx,
		UserID: userID,
	}
	lockAuthenticationServiceMockCreateToken.Lock()
	mock.calls.CreateToken = append(mock.calls.CreateToken, callInfo)
	lockAuthenticationServiceMockCreateToken.Unlock()
	return mock.CreateTokenFunc(ctx, userID)
}

// CreateTokenCalls gets all the calls that were made to CreateToken.
// Check the length with:
//     len(mockedAuthenticationService.CreateTokenCalls())
func (mock *AuthenticationServiceMock) CreateTokenCalls() []struct {
	Ctx    context.Context
	UserID uint64
} {
	var calls []struct {
		Ctx    context.Context
		UserID uint64
	}
	lockAuthenticationServiceMockCreateToken.RLock()
	calls = mock.calls.CreateToken
	lockAuthenticationServiceMockCreateToken.RUnlock()
	return calls
}

// ParseToken calls ParseTokenFunc.
func (mock *AuthenticationServiceMock) ParseToken(ctx context.Context, token string) (*jwt.Token, error) {
	if mock.ParseTokenFunc == nil {
		panic("AuthenticationServiceMock.ParseTokenFunc: method is nil but AuthenticationService.ParseToken was just called")
	}
	callInfo := struct {
		Ctx   context.Context
		Token string
	}{
		Ctx:   ctx,
		Token: token,
	}
	lockAuthenticationServiceMockParseToken.Lock()
	mock.calls.ParseToken = append(mock.calls.ParseToken, callInfo)
	lockAuthenticationServiceMockParseToken.Unlock()
	return mock.ParseTokenFunc(ctx, token)
}

// ParseTokenCalls gets all the calls that were made to ParseToken.
// Check the length with:
//     len(mockedAuthenticationService.ParseTokenCalls())
func (mock *AuthenticationServiceMock) ParseTokenCalls() []struct {
	Ctx   context.Context
	Token string
} {
	var calls []struct {
		Ctx   context.Context
		Token string
	}
	lockAuthenticationServiceMockParseToken.RLock()
	calls = mock.calls.ParseToken
	lockAuthenticationServiceMockParseToken.RUnlock()
	return calls
}

// PublicKey calls PublicKeyFunc.
func (mock *AuthenticationServiceMock) PublicKey(ctx context.Context) (string, error) {
	if mock.PublicKeyFunc == nil {
		panic("AuthenticationServiceMock.PublicKeyFunc: method is nil but AuthenticationService.PublicKey was just called")
	}
	callInfo := struct {
		Ctx context.Context
	}{
		Ctx: ctx,
	}
	lockAuthenticationServiceMockPublicKey.Lock()
	mock.calls.PublicKey = append(mock.calls.PublicKey, callInfo)
	lockAuthenticationServiceMockPublicKey.Unlock()
	return mock.PublicKeyFunc(ctx)
}

// PublicKeyCalls gets all the calls that were made to PublicKey.
// Check the length with:
//     len(mockedAuthenticationService.PublicKeyCalls())
func (mock *AuthenticationServiceMock) PublicKeyCalls() []struct {
	Ctx context.Context
} {
	var calls []struct {
		Ctx context.Context
	}
	lockAuthenticationServiceMockPublicKey.RLock()
	calls = mock.calls.PublicKey
	lockAuthenticationServiceMockPublicKey.RUnlock()
	return calls
}

// RefreshToken calls RefreshTokenFunc.
func (mock *AuthenticationServiceMock) RefreshToken(ctx context.Context, token string) (*string, error) {
	if mock.RefreshTokenFunc == nil {
		panic("AuthenticationServiceMock.RefreshTokenFunc: method is nil but AuthenticationService.RefreshToken was just called")
	}
	callInfo := struct {
		Ctx   context.Context
		Token string
	}{
		Ctx:   ctx,
		Token: token,
	}
	lockAuthenticationServiceMockRefreshToken.Lock()
	mock.calls.RefreshToken = append(mock.calls.RefreshToken, callInfo)
	lockAuthenticationServiceMockRefreshToken.Unlock()
	return mock.RefreshTokenFunc(ctx, token)
}

// RefreshTokenCalls gets all the calls that were made to RefreshToken.
// Check the length with:
//     len(mockedAuthenticationService.RefreshTokenCalls())
func (mock *AuthenticationServiceMock) RefreshTokenCalls() []struct {
	Ctx   context.Context
	Token string
} {
	var calls []struct {
		Ctx   context.Context
		Token string
	}
	lockAuthenticationServiceMockRefreshToken.RLock()
	calls = mock.calls.RefreshToken
	lockAuthenticationServiceMockRefreshToken.RUnlock()
	return calls
}

// VerifyToken calls VerifyTokenFunc.
func (mock *AuthenticationServiceMock) VerifyToken(ctx context.Context, token string) (bool, *uint64, error) {
	if mock.VerifyTokenFunc == nil {
		panic("AuthenticationServiceMock.VerifyTokenFunc: method is nil but AuthenticationService.VerifyToken was just called")
	}
	callInfo := struct {
		Ctx   context.Context
		Token string
	}{
		Ctx:   ctx,
		Token: token,
	}
	lockAuthenticationServiceMockVerifyToken.Lock()
	mock.calls.VerifyToken = append(mock.calls.VerifyToken, callInfo)
	lockAuthenticationServiceMockVerifyToken.Unlock()
	return mock.VerifyTokenFunc(ctx, token)
}

// VerifyTokenCalls gets all the calls that were made to VerifyToken.
// Check the length with:
//     len(mockedAuthenticationService.VerifyTokenCalls())
func (mock *AuthenticationServiceMock) VerifyTokenCalls() []struct {
	Ctx   context.Context
	Token string
} {
	var calls []struct {
		Ctx   context.Context
		Token string
	}
	lockAuthenticationServiceMockVerifyToken.RLock()
	calls = mock.calls.VerifyToken
	lockAuthenticationServiceMockVerifyToken.RUnlock()
	return calls
}
