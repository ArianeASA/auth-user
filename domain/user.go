package domain

import (
	"auth-user/domain/dto"
)

type User struct {
	Name               string
	UserName           string
	RegistrationNumber string
	Email              string
	Password           string
}

type CredentialsByUserName struct {
	Password string
	UserName string
}

type CredentialsByRegistrationNumber struct {
	Password           string
	RegistrationNumber string
}

type AuthResult struct {
	AccessToken *string
	ExpiresIn   *int64
	TokenType   *string
}

func (user *User) userFromDomain() User {
	return User{
		UserName:           user.UserName,
		RegistrationNumber: user.RegistrationNumber,
		Name:               user.Name,
		Email:              user.Email,
		Password:           user.Password,
	}
}

func (cred *CredentialsByUserName) credentialsFromDomain() dto.Credentials {
	return dto.Credentials{
		Password: cred.Password,
		Username: cred.UserName,
	}
}

func AuthResultToDomain(result dto.AuthResult) AuthResult {
	return AuthResult{
		AccessToken: result.AccessToken,
		ExpiresIn:   result.ExpiresIn,
		TokenType:   result.TokenType,
	}
}

func UserToDomain(user dto.User) User {
	return User{
		UserName:           user.Username,
		RegistrationNumber: user.RegistrationNumber,
		Name:               user.Name,
		Email:              user.Email,
		Password:           user.Password,
	}
}

func CredentialsToDomain(cred dto.Credentials) CredentialsByUserName {
	return CredentialsByUserName{
		Password: cred.Password,
		UserName: cred.Username,
	}
}

func (a *AuthResult) DomainToResponse() dto.AuthResult {
	return dto.AuthResult{
		AccessToken: a.AccessToken,
		ExpiresIn:   a.ExpiresIn,
		TokenType:   a.TokenType,
	}
}
