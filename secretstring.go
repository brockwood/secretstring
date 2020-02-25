package secretstring

import (
	"errors"
	"os"
)

var (
	gcpProject = os.Getenv("GCP_PROJECT")
)

type SecretString string

func (s *SecretString) UnmarshalText(text []byte) error {
	if gcpProject == "" {
		return errors.New("the GCP_PROJECT environment variable is not set")
	}
	if secretManagerClient == nil {
		if err := setupSecretClient(); err != nil {
			return err
		}
	}
	secretName := string(text)
	value, err := secretManagerClient.getSecretString(gcpProject, secretName)
	if err != nil {
		return err
	}
	*s = SecretString(value)
	return nil
}
