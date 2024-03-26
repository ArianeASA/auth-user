package domain

type Auth interface {
	NewUser(user User) error
	NewToken(cred Credentials) (AuthResult, error)
}
