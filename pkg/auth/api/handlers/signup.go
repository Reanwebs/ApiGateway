package handlers

import "gateway/pkg/auth/client/interfaces"

type SignupHandler struct {
	Client interfaces.AuthClient
}
