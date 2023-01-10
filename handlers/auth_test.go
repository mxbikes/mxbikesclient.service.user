package handlers

import (
	"testing"
	"time"

	"github.com/mxbikes/mxbikesclient.service.user/models"
	"github.com/stretchr/testify/assert"
)

func TestIsValidNotExpired(t *testing.T) {
	// Arrange
	var expiresAt = time.Now()

	// Act
	isValid := IsValid(models.AuthResponse{ExpiresAt: expiresAt})

	// Assert
	assert.Equal(t, isValid, true)
}

func TestIsValidExpired(t *testing.T) {
	// Arrange
	var expiresAt = time.Date(2021, 8, 15, 14, 30, 45, 100, time.Local)

	// Act
	isValid := IsValid(models.AuthResponse{ExpiresAt: expiresAt})

	// Assert
	assert.Equal(t, isValid, false)
}
