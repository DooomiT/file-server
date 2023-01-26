package filehandler

import (
	"io"
	"os"
)

type FileReader interface {
	GetFileContent(filePath string) (string, error)
	ListFilesRecursive(dirPath string) ([]string, error)
}

type FileReaderImpl struct {
	rootFolder string
}

func NewFileReader(rootFolder string) FileReader {
	fileReader := &FileReaderImpl{
		rootFolder: rootFolder,
	}
	return fileReader
}

// GetFileContent reads the content of filePath, and returns it as a string.
// If the file does not exist, or if the file is empty, an error is returned.
func (fr *FileReaderImpl) GetFileContent(filePath string) (string, error) {
	absoluteFilePath := fr.rootFolder + "/" + filePath
	content, err := getFileContent(absoluteFilePath)
	if err != nil {
		return "", err
	}
	return content, nil
}

func getFileContent(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	var content []byte
	buf := make([]byte, 1024)
	for {
		n, err := file.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			return "", err
		}
		content = append(content, buf[:n]...)
	}

	return string(content), nil
}

// ListFilesRecursive recursively lists all the files in the given directory
// and returns them as a slice of strings.
func (fr *FileReaderImpl) ListFilesRecursive(dirPath string) ([]string, error) {
	absoluteDirPath := fr.rootFolder + "/" + dirPath
	files, err := listFilesRecursive(absoluteDirPath)
	if err != nil {
		return nil, err
	}
	return files, nil
}

func listFilesRecursive(dirPath string) ([]string, error) {
	var files []string

	file, err := os.Open(dirPath)
	if err != nil {
		return files, err
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		return files, err
	}

	if !fileInfo.IsDir() {
		files = append(files, fileInfo.Name())
		return files, nil
	}

	fileInfos, err := file.Readdir(0)
	if err != nil {
		return files, err
	}

	for _, fileInfo := range fileInfos {
		if fileInfo.IsDir() {
			subDirFiles, err := listFilesRecursive(dirPath + "/" + fileInfo.Name())
			if err != nil {
				return files, err
			}
			files = append(files, subDirFiles...)
		} else {
			files = append(files, fileInfo.Name())
		}
	}

	return files, nil
}
