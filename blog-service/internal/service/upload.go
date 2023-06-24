package service

import (
	"errors"
	"github.com/go-programming-tour-book/blog-service/global"
	"github.com/go-programming-tour-book/blog-service/pkg/upload"
	"mime/multipart"
	"os"
)

type FileInfo struct {
	Name       string
	AccessPath string
}

func (svc *Service) UploadFile(fileType upload.FileType, file multipart.File, fileHeader *multipart.FileHeader) (*FileInfo, error) {
	//get file name(encoded)
	fileName := upload.GetFileName(fileHeader.Filename)
	if !upload.CheckContainExt(fileType, fileName) {
		return nil, errors.New("file suffix is not supported.")
	}
	if !upload.CheckMaxSize(fileType, file) {
		return nil, errors.New("exceeded maximum file limit.")
	}
	uploadSavePath := upload.GetSavePath()
	//save path check
	if upload.CheckSavePath(uploadSavePath) {
		//not exist -> create
		if err := upload.CreateSavePath(uploadSavePath, os.ModePerm); err != nil {
			return nil, errors.New("failed to create save directory.")
		}
	}

	if upload.CheckPermission(uploadSavePath) {
		return nil, errors.New("insufficient file permissions.")
	}

	//destination path
	dst := uploadSavePath + "/" + fileName
	if err := upload.SaveFile(fileHeader, dst); err != nil {
		return nil, errors.New("failed to save file.")
	}
	accessUrl := global.AppSetting.UploadServerUrl + "/" + fileName

	return &FileInfo{Name: fileName, AccessPath: accessUrl}, nil
}
