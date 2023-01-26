package fileserver

import (
	"strings"

	"github.com/dooomit/file-server/pkg/filereader"
)

type FileServer struct {
	fileReader *filereader.FileReader
}

func NewFileServer(root string) *FileServer {
	return &FileServer{
		fileReader: filereader.NewFileReader(root),
	}
}

func (fs *FileServer) GetFileContent(filePath string) (string, error) {}

// FilterFiles returns a slice of strings containing only the files that
// contain the given pattern.
func FilterFiles(pattern string, files []string) []string {
	var filteredFiles []string
	for _, file := range files {
		if strings.Contains(file, pattern) {
			filteredFiles = append(filteredFiles, file)
		}
	}
	return filteredFiles
}
