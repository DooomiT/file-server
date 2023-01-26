package fileserver

import (
	"strings"

	filereader "github.com/dooomit/file-server/pkg/FileReader"

	"github.com/gin-gonic/gin"
)

type FileServer struct {
	fileReader filereader.FileReader
	router     *gin.Engine
}

func NewFileServer(root string) *FileServer {
	fileReader := filereader.NewFileReader(root)
	router := gin.Default()
	fs := FileServer{
		fileReader: fileReader,
		router:     router,
	}
	fs.router.GET("/files", fs.ListFiles)
	fs.router.GET("/files/file", fs.GetFileContent)
	return &fs
}

func (fs *FileServer) Run(addr string) error {
	return fs.router.Run(addr)
}

func (fs *FileServer) GetFileContent(c *gin.Context) {
	filePath := c.Query("path")
	if filePath == "" {
		c.JSON(400, gin.H{"error": "file path is required"})
		return
	}
	content, err := fs.fileReader.GetFileContent(filePath)
	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}
	c.Data(200, "text/plain", []byte(content))
}

func (fs *FileServer) ListFiles(c *gin.Context) {
	// Get url param "filter"
	filter := c.Query("filter")

	files, err := fs.fileReader.ListFilePathsRecursive()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	if filter == "" {
		c.JSON(200, gin.H{"files": files})
		return
	}

	filteredFiles := FilterFiles(filter, files)
	c.JSON(200, gin.H{"files": filteredFiles})
}

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
