// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package azidentity

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
)

var (
	mockCLITokenProviderSuccess = func(ctx context.Context, resource string) ([]byte, error) {
		return []byte(" {\"accessToken\":\"mocktoken\" , " +
			"\"expiresOn\": \"2007-01-01 01:01:01.079627\"," +
			"\"subscription\": \"mocksub\"," +
			"\"tenant\": \"mocktenant\"," +
			"\"tokenType\": \"mocktype\"}"), nil
	}
	mockCLITokenProviderFailure = func(ctx context.Context, resource string) ([]byte, error) {
		return nil, errors.New("provider failure message")
	}
)

func TestAzureCLICredential_GetTokenSuccess(t *testing.T) {
	cred, err := NewAzureCLICredential(&AzureCLICredentialOptions{TokenProvider: mockCLITokenProviderSuccess})
	if err != nil {
		t.Fatalf("Unable to create credential. Received: %v", err)
	}
	at, err := cred.GetToken(context.Background(), azcore.TokenRequestOptions{Scopes: []string{scope}})
	if err != nil {
		fmt.Println(err.Error())
		t.Fatalf("Expected an empty error but received: %v", err)
	}
	if len(at.Token) == 0 {
		t.Fatalf(("Did not receive a token"))
	}
	if at.Token != "mocktoken" {
		t.Fatalf(("Did not receive the correct access token"))
	}
}

func TestAzureCLICredential_GetTokenInvalidToken(t *testing.T) {
	cred, err := NewAzureCLICredential(&AzureCLICredentialOptions{TokenProvider: mockCLITokenProviderFailure})
	if err != nil {
		t.Fatalf("Unable to create credential. Received: %v", err)
	}
	_, err = cred.GetToken(context.Background(), azcore.TokenRequestOptions{Scopes: []string{scope}})
	if err == nil {
		t.Fatalf("Expected an error but did not receive one.")
	}
}