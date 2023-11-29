package gogit

import (
	"compress/zlib"
	"io"
	"os"
)

func ObjectFromFile(filename string) (string, error) {

	f, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer f.Close()
	r, err := zlib.NewReader(f)
	if err != nil {
		return "", err
	}
	// Читаем данные из zlib.Reader в виде среза байтов
	compressedData, err := io.ReadAll(r)
	if err != nil {
		return "", err
	}

	// Преобразуем срез байтов в строку
	resultString := string(compressedData)

	return resultString, nil
}
