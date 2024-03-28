package auth

type AuthService interface {
	AuthenticateWithThirdParty(idToken string) (bool, error)
}
