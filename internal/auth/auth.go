package auth

type Authenticator interface {
	IsAuthenticated() (bool, error)
	Login() error
}
