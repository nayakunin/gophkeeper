package utils

import "google.golang.org/grpc/metadata"

func GetRequestMetadata(token string) metadata.MD {
	return metadata.Pairs("authorization", "Bearer "+token)
}
