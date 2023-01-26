package filereader

import (
	"io"
	"os"
)

type FileReader interface {
	GetFileContent(filePath string) (string, error)
	ListFilePathsRecursive() ([]string, error)
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

// ListFilePathsRecursive recursively lists all the files in the given directory
// and returns them as a slice of strings.
func (fr *FileReaderImpl) ListFilePathsRecursive() ([]string, error) {
	files, err := listFilePathsRecursive(fr.rootFolder)
	if err != nil {
		return nil, err
	}
	return removeRootFolder(fr.rootFolder, files), nil
}

func listFilePathsRecursive(dirPath string) ([]string, error) {
	var files []string
	file, err := os.Open(dirPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		return nil, err
	}

	if fileInfo.IsDir() {
		fileInfos, err := file.Readdir(-1)
		if err != nil {
			return nil, err
		}
		for _, fileInfo := range fileInfos {
			filePath := dirPath + "/" + fileInfo.Name()
			file, err := os.Open(filePath)
			if err != nil {
				return nil, err
			}
			defer file.Close()
			if fileInfo.IsDir() {
				subFiles, err := listFilePathsRecursive(filePath)
				if err != nil {
					return nil, err
				}
				files = append(files, subFiles...)
			} else {
				files = append(files, filePath)
			}
		}
	} else {
		files = append(files, dirPath)
	}

	return files, nil
}

func removeRootFolder(rootFolder string, files []string) []string {
	var filteredFiles []string
	for _, file := range files {
		filteredFiles = append(filteredFiles, file[len(rootFolder)+1:])
	}
	return filteredFiles
}
