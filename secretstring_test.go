package secretstring

import (
	"fmt"
	"testing"
)

type getSecretStringFunc func(project, name string) (string, error)

func newTestClient(get getSecretStringFunc) *testClient {
	return &testClient{getsecretstring: get}
}

type testClient struct {
	getsecretstring func(project, name string) (string, error)
}

func (t *testClient) getSecretString(project, secretName string) (string, error) {
	if t.getsecretstring != nil {
		return t.getsecretstring(project, secretName)
	}
	return "I'm an important string", nil
}

func TestSecretString_UnmarshalText(t *testing.T) {
	t.Run("test successful retrieval", func(t *testing.T) {
		tc := newTestClient(func(project, name string) (string, error) {
			return "IAMASECUREPASWORD", nil
		})
		secretManagerClient = tc
		secretText := SecretString("")
		err := secretText.UnmarshalText([]byte("ENCRYPTEDPASSWORD"))
		if err != nil {
			t.Error("Error retrieving secret:", err)
		}
		t.Log("Retrieved secret ", secretText)
	})
	t.Run("test failed retrieval", func(t *testing.T) {
		tc := newTestClient(func(project, name string) (string, error) {
			return "", fmt.Errorf("unable to retrieve secret for project %s, name %s", project, name)
		})
		secretManagerClient = tc
		secretText := SecretString("")
		err := secretText.UnmarshalText([]byte("ENCRYPTEDPASSWORD"))
		if err == nil {
			t.Error("Should of received an error but did not.")
		}
	})
}
