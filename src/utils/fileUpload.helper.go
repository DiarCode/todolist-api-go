package utils

import (
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"os"
	"path/filepath"
)

func UploadImage(file *multipart.FileHeader, filename string) (avatarPath string, err error) {
	path := fmt.Sprintf("src/storage/avatars/%v", filename)
	abs, _ := filepath.Abs(path)

	f, err := file.Open()
	if err != nil {
		return "", err
	}

	out, err := os.Create(abs)
	if err != nil {
		return "", err
	}

	defer out.Close()
	_, err = io.Copy(out, f)

	if err != nil {
		return "", err
	}

	resPath := "https://todolist-api-go.up.railway.app/api/v1/users/avatars/" + filename

	return resPath, nil
}

func GetImageByName(filename string) (fb []byte, err error) {
	emptyBytesSlice := make([]byte, 0)
	path := fmt.Sprintf("src/storage/avatars/%v", filename)
	abs, _ := filepath.Abs(path)

	fileBytes, err := ioutil.ReadFile(abs)
	if err != nil {
		return emptyBytesSlice, err
	}

	return fileBytes, nil
}
