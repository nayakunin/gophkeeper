package utils

import "google.golang.org/grpc/metadata"

// GetRequestMetadata returns metadata to be sent with every request.
func GetRequestMetadata(token string) metadata.MD {
	return metadata.Pairs("authorization", "Bearer "+token)
}
