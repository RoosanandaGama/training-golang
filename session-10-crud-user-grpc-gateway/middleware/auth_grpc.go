package middleware

import (
	"context"
	"encoding/base64"
	"strings"
	"training-golang/session-10-crud-user-grpc-gateway/config"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func UnaryAuthInterceptor() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler) (interface{}, error) {

		// Check for public methods that do not require authentication
		publicMethods := []string{
			"/proto.user_service.v1.UserService/GetUsers",
			"/proto.user_service.v1.UserService/GetUserByID",
		}

		for _, method := range publicMethods {
			if info.FullMethod == method {
				return handler(ctx, req)
			}
		}

		// Extract metadata from context
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return nil, status.Errorf(codes.Unauthenticated, "metadata is not provided")
		}

		// Get Authorization header
		authHeader, ok := md["authorization"]
		if !ok || len(authHeader) == 0 {
			return nil, status.Errorf(codes.Unauthenticated, "authorization header is missing")
		}

		// Check for Basic Auth Scheme
		if !strings.HasPrefix(authHeader[0], "Basic ") {
			return nil, status.Errorf(codes.Unauthenticated, "invalid authorization scheme")
		}

		// Decode Base64 encoded credentials
		decoded, err := base64.StdEncoding.DecodeString(authHeader[0][6:])
		if err != nil {
			return nil, status.Errorf(codes.Unauthenticated, "invalid authorization token")
		}

		// Split credentials into username and password
		creds := strings.SplitN(string(decoded), ":", 2)
		if len(creds) != 2 {
			return nil, status.Errorf(codes.Unauthenticated, "invalid authorization token")
		}

		username, password := creds[0], creds[1]

		//validate the credentials
		if username != config.AuthBasicUsername || password != config.AuthBasicPassword {
			return nil, status.Errorf(codes.Unauthenticated, "invalid username or password")
		}

		return handler(ctx, req)

	}
}
