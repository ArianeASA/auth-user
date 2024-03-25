package domain

type Auth interface {
	NewUser(user User) error
	NewToken(cred CredentialsByUserName) (AuthResult, error)
}
