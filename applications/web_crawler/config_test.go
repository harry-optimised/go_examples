package main

import (
	"os"
	"testing"
)

func TestConfig(t *testing.T) {

	t.Run("test config struct can be instantiated", func(t *testing.T) {
		config := Config{host: "monzo.com", throttle: true}

		assertEqual(t, config.host,"monzo.com")
		assertEqual(t, config.throttle, true)
	})

	t.Run("test config is created from env vars", func(t *testing.T) {
		// Given - Two set env vars.
		os.Setenv("WC_HOST", "google.com")
		os.Setenv("WC_THROTTLE", "false")

		// When - We create config from env vars.
		config := ConfigFromEnv()

		// Then - It contains the expected values.
		assertEqual(t, config.host, "google.com")
		assertEqual(t, config.throttle, false)
	})

	t.Run("test config from env vars has default values", func(t *testing.T) {
		// Given - No environment variables.
		os.Unsetenv("WC_HOST")
		os.Unsetenv("WC_THROTTLE")

		// When - We create config from env vars.
		config := ConfigFromEnv()

		// Then - It contains the default values.
		assertEqual(t, config.host, HOST_DEFAULT)
		assertEqual(t, config.throttle, THROTTLE_DEFAULT)
	})


}

func assertEqual[T comparable](t *testing.T, got, want T) {
	t.Helper()
	if got != want {
		t.Errorf("want %v got %v", want, got)
	}
}
