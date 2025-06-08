package store

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestOTPStore_AddGetDelete(t *testing.T) {
	store := NewOTPStore[string](time.Second)

	prefix := "test:"
	otp := store.Add(prefix, "") // Let it auto-generate an OTP

	// Ensure the OTP is immediately retrievable
	value, ok := store.Get(prefix, otp)
	assert.True(t, ok, "OTP should be retrievable after Add")
	assert.Equal(t, otp, value.Value, "Stored and retrieved OTP should match")

	// Delete the OTP and verify it's no longer retrievable
	store.Delete(prefix, otp)
	_, ok = store.Get(prefix, otp)
	assert.False(t, ok, "OTP should be deleted")
}

func TestOTPStore_Expiration(t *testing.T) {
	store := NewOTPStore[string](time.Second)

	prefix := "expire:"
	otp := store.Add(prefix, "")

	// Ensure it's initially present
	_, ok := store.Get(prefix, otp)
	assert.True(t, ok, "OTP should exist immediately after Add")

	// Wait for TTL to expire
	time.Sleep(2 * time.Second)

	// Check that OTP has expired
	_, ok = store.Get(prefix, otp)
	assert.False(t, ok, "OTP should be expired and no longer available")
}
