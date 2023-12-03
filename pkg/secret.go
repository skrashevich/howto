package howto

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"strings"

	"github.com/zalando/go-keyring"
)

func SetOpenAiApiKey(apiKey string) error {
	err := keyring.Set(SERVICE_NAME, "openai_api_key", apiKey)
	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}

// GetOpenAiApiKey returns the OpenAI API key.
//
// It retrieves the API key from the environment variable "OPENAI_API_KEY".
// If the key does not start with "sk-", it checks if it exists in the keyring
// using the service name "openai_api_key" on macOS. If the key is not found,
// it prints an error message and returns an empty string.
// If the key is invalid, it prints an error message and returns the key.
// Otherwise, it returns the API key.
//
// Returns:
// - string: the OpenAI API key
// - error: any error that occurred during retrieval
func GetOpenAiApiKey() (string, error) {
	// many issues with keyring on Linux, use env var
	secret := os.Getenv("OPENAI_API_KEY")

	if !strings.HasPrefix(secret, "sk-") {
		if runtime.GOOS == "darwin" {
			var err error
			secret, err = keyring.Get(SERVICE_NAME, "openai_api_key")

			// check if it exists at all
			if err == keyring.ErrNotFound {
				fmt.Println("OpenAI API key not found. Please run `howto --setup` to set it in keyring.")
				return "", err
			}
		}
		fmt.Println("OpenAI API key is invalid. Please run `howto config` to set it.")
		return secret, nil
	}

	return secret, nil
}
