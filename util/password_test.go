package util

import (
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

func TestPassword(t *testing.T) {
	password := RandomString(6)

	hashedPassWord1, err := HashedPassword(password)
	require.NoError(t, err)
	require.NotEmpty(t, hashedPassWord1)

	err = CheckPassword(password, hashedPassWord1)
	require.NoError(t, err)

	wrongPassword := RandomString(6)
	err = CheckPassword(wrongPassword, hashedPassWord1)
	require.EqualError(t, err, bcrypt.ErrMismatchedHashAndPassword.Error())

	hashedPassWord2, err := HashedPassword(password)
	require.NoError(t, err)
	require.NotEmpty(t, hashedPassWord2)
	require.NotEqual(t, hashedPassWord1, hashedPassWord2)
}