package service

import (
	"context"
	"github.com/mp3tags/file-repository-proto"
)

func (s Service) GetFile(context.Context, *file_repository.GetFileRequest) (*file_repository.File, error) {
	panic("implement me")
}
