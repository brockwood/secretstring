package secretstring

import (
	"context"
	"fmt"

	secretmanager "cloud.google.com/go/secretmanager/apiv1beta1"
	secretmanagerpb "google.golang.org/genproto/googleapis/cloud/secretmanager/v1beta1"
)

type secretRetrieval interface {
	getSecretString(project, name string) (string, error)
}

const (
	secretResourceName = `projects/%s/secrets/%s/versions/latest`
)

var (
	secretManagerClient secretRetrieval
)

type secretClient struct {
	secretClient *secretmanager.Client
}

func (s *secretClient) getSecretString(project, desiredSecretName string) (string, error) {
	ctx := context.Background()
	name := fmt.Sprintf(secretResourceName, project, desiredSecretName)
	req := &secretmanagerpb.AccessSecretVersionRequest{
		Name: name,
	}
	resp, err := s.secretClient.AccessSecretVersion(ctx, req)
	if err != nil {
		return "", err
	}
	return string(resp.Payload.Data), nil
}

func setupSecretClient() (err error) {
	var sc secretClient
	ctx := context.Background()
	sc.secretClient, err = secretmanager.NewClient(ctx)
	secretManagerClient = &sc
	return
}
