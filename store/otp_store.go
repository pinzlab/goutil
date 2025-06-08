package store

import (
	"fmt"
	"math/rand"
	"time"
)

// OTPStore is a specialized key-value store for handling One-Time Passwords (OTPs).
// It uses a generic Store to store OTPs keyed by a prefix + OTP string,
// and supports automatic expiration using TTL.
//
// The type parameter P is constrained to string and represents a customizable prefix type.
type OTPStore[P ~string] struct {
	store Store[string]
	ttl   time.Duration
}

// NewOTPStore creates and returns a new instance of OTPStore with the given TTL.
// It initializes the internal key-value store and configures the expiration time for each OTP.
func NewOTPStore[P ~string](ttl time.Duration) *OTPStore[P] {
	return &OTPStore[P]{
		store: *NewStore[string](),
		ttl:   ttl,
	}
}

// generateOTP creates a random 6-digit numeric OTP as a string.
//
// Returns:
// - a string representing the generated OTP (e.g., "482913")
func (s *OTPStore[P]) generateOTP() string {
	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)
	return fmt.Sprintf("%06d", r.Intn(1000000))
}

// Add stores a new OTP value associated with the given prefix.
// If the value is empty, a new OTP is generated automatically.
// The OTP is stored with a TTL, after which it expires automatically.
//
// Parameters:
// - prefix: the namespace or identifier prefix for the OTP
// - value: the OTP value to store (optional; if empty, a new OTP is generated)
//
// Returns:
//   - The OTP that was stored (either the one provided or the one generated)
func (s *OTPStore[P]) Add(prefix P, value string) string {
	otp := s.generateOTP()
	if value == "" {
		value = otp
	}
	key := string(prefix) + otp
	s.store.Add(key, value, s.ttl)

	return otp
}

// Get retrieves the stored value associated with a given prefix and OTP.
//
// Parameters:
// - prefix: the prefix used when storing the OTP
// - otpValue: the OTP value to look up
//
// Returns:
// - the stored value
// - a boolean indicating if the key was found
func (s *OTPStore[P]) Get(prefix P, otpValue string) (Value[string], bool) {
	key := string(prefix) + otpValue
	return s.store.Get(key)
}

// Delete removes an OTP entry associated with the given prefix and OTP value.
//
// Parameters:
// - prefix: the prefix used when storing the OTP
// - otpValue: the OTP value to delete
func (s *OTPStore[P]) Delete(prefix P, otpValue string) {
	key := string(prefix) + otpValue
	s.store.Delete(key)
}
