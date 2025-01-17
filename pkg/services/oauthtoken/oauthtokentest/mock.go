package oauthtokentest

import (
	"context"

	"github.com/grafana/grafana/pkg/models"
	"github.com/grafana/grafana/pkg/services/datasources"
	"github.com/grafana/grafana/pkg/services/user"
	"golang.org/x/oauth2"
)

type MockOauthTokenService struct {
	GetCurrentOauthTokenFunc   func(ctx context.Context, usr *user.SignedInUser) *oauth2.Token
	IsOAuthPassThruEnabledFunc func(ds *datasources.DataSource) bool
	HasOAuthEntryFunc          func(ctx context.Context, usr *user.SignedInUser) (*models.UserAuth, bool, error)
	InvalidateOAuthTokensFunc  func(ctx context.Context, usr *models.UserAuth) error
	TryTokenRefreshFunc        func(ctx context.Context, usr *models.UserAuth) error
}

func (m *MockOauthTokenService) GetCurrentOAuthToken(ctx context.Context, usr *user.SignedInUser) *oauth2.Token {
	if m.GetCurrentOauthTokenFunc != nil {
		return m.GetCurrentOauthTokenFunc(ctx, usr)
	}
	return nil
}

func (m *MockOauthTokenService) IsOAuthPassThruEnabled(ds *datasources.DataSource) bool {
	if m.IsOAuthPassThruEnabledFunc != nil {
		return m.IsOAuthPassThruEnabledFunc(ds)
	}
	return false
}

func (m *MockOauthTokenService) HasOAuthEntry(ctx context.Context, usr *user.SignedInUser) (*models.UserAuth, bool, error) {
	if m.HasOAuthEntryFunc != nil {
		return m.HasOAuthEntryFunc(ctx, usr)
	}
	return nil, false, nil
}

func (m *MockOauthTokenService) InvalidateOAuthTokens(ctx context.Context, usr *models.UserAuth) error {
	if m.InvalidateOAuthTokensFunc != nil {
		return m.InvalidateOAuthTokensFunc(ctx, usr)
	}
	return nil
}

func (m *MockOauthTokenService) TryTokenRefresh(ctx context.Context, usr *models.UserAuth) error {
	if m.TryTokenRefreshFunc != nil {
		return m.TryTokenRefreshFunc(ctx, usr)
	}
	return nil
}
