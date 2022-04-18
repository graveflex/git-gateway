package api

import (
	"context"
	"math/rand"
	"net/url"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/netlify/git-gateway/conf"
	"github.com/netlify/git-gateway/models"
)

type Role struct {
	Name string
}

type contextKey string

func (c contextKey) String() string {
	return "git-gateway api context key " + string(c)
}

const (
	accessTokenKey = contextKey("access_token")
	tokenKey       = contextKey("jwt")
	requestIDKey   = contextKey("request_id")
	configKey      = contextKey("config")
	instanceIDKey  = contextKey("instance_id")
	instanceKey    = contextKey("instance")
	proxyTargetKey = contextKey("target")
	signatureKey   = contextKey("signature")
	netlifyIDKey   = contextKey("netlify_id")
)

// withToken adds the JWT token to the context.
func withToken(ctx context.Context, token *jwt.Token) context.Context {
	return context.WithValue(ctx, tokenKey, token)
}

// getToken reads the JWT token from the context.
func getToken(ctx context.Context) *jwt.Token {
	obj := ctx.Value(tokenKey)
	if obj == nil {
		return nil
	}

	return obj.(*jwt.Token)
}

func getClaims(ctx context.Context) *GatewayClaims {
	token := getToken(ctx)
	if token == nil {
		return nil
	}
	return token.Claims.(*GatewayClaims)
}

func withRequestID(ctx context.Context, id string) context.Context {
	return context.WithValue(ctx, requestIDKey, id)
}

func getRequestID(ctx context.Context) string {
	obj := ctx.Value(requestIDKey)
	if obj == nil {
		return ""
	}

	return obj.(string)
}

func getConfig(ctx context.Context) *conf.Configuration {
	obj := ctx.Value(configKey)
	if obj == nil {
		return nil
	}

	config := obj.(*conf.Configuration)
	return config
}

func withConfig(ctx context.Context, config *conf.Configuration) context.Context {
	return context.WithValue(ctx, configKey, config)
}

func withInstanceID(ctx context.Context, id string) context.Context {
	return context.WithValue(ctx, instanceIDKey, id)
}

func withProxyTarget(ctx context.Context, target *url.URL) context.Context {
	return context.WithValue(ctx, proxyTargetKey, target)
}

func getProxyTarget(ctx context.Context) *url.URL {
	obj := ctx.Value(proxyTargetKey)
	if obj == nil {
		return nil
	}
	return obj.(*url.URL)
}

func withAccessToken(ctx context.Context, token string) context.Context {
	return context.WithValue(ctx, accessTokenKey, token)
}

func getAccessToken(ctx context.Context) string {
	obj := ctx.Value(accessTokenKey)
	if obj == nil {
		return ""
	}

	tokens := strings.Split(obj.(string), ",")

	return tokens[rand.Intn(len(tokens))]
}

func getInstanceID(ctx context.Context) string {
	obj := ctx.Value(instanceIDKey)
	if obj == nil {
		return ""
	}
	return obj.(string)
}

func withInstance(ctx context.Context, i *models.Instance) context.Context {
	return context.WithValue(ctx, instanceKey, i)
}

func getInstance(ctx context.Context) *models.Instance {
	obj := ctx.Value(instanceKey)
	if obj == nil {
		return nil
	}
	return obj.(*models.Instance)
}

func withSignature(ctx context.Context, id string) context.Context {
	return context.WithValue(ctx, signatureKey, id)
}

func getSignature(ctx context.Context) string {
	obj := ctx.Value(signatureKey)
	if obj == nil {
		return ""
	}

	return obj.(string)
}

func withNetlifyID(ctx context.Context, id string) context.Context {
	return context.WithValue(ctx, netlifyIDKey, id)
}

func getNetlifyID(ctx context.Context) string {
	obj := ctx.Value(netlifyIDKey)
	if obj == nil {
		return ""
	}

	return obj.(string)
}
