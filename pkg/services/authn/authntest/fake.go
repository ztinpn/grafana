package authntest

import (
	"context"

	"github.com/grafana/grafana/pkg/services/authn"
)

type FakeService struct {
	authn.Service
}

var _ authn.Client = new(FakeClient)

type FakeClient struct {
	ExpectedErr      error
	ExpectedTest     bool
	ExpectedIdentity *authn.Identity
}

func (f *FakeClient) Authenticate(ctx context.Context, r *authn.Request) (*authn.Identity, error) {
	return f.ExpectedIdentity, f.ExpectedErr
}

func (f *FakeClient) Test(ctx context.Context, r *authn.Request) bool {
	return f.ExpectedTest
}

var _ authn.PasswordClient = new(FakePasswordClient)

type FakePasswordClient struct {
	ExpectedErr      error
	ExpectedIdentity *authn.Identity
}

func (f FakePasswordClient) AuthenticatePassword(ctx context.Context, r *authn.Request, username, password string) (*authn.Identity, error) {
	return f.ExpectedIdentity, f.ExpectedErr
}
