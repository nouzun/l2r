package service

import (
	"github.com/gin-gonic/gin"

	"github.com/nouzun/l2r/pkg/database"
	"github.com/nouzun/l2r/pkg/model"
)

type Service struct {
	DB *database.Database
}

func (s Service) GetWords(context *gin.Context) ([]model.Word, error) {

	words, err := s.DB.GetWords()
	if err != nil {
		return nil, err
	}

	return words, nil
}
