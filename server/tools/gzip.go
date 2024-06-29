package tools

import (
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

func Gzip(paths []string) gin.HandlerFunc {
	return gzip.Gzip(gzip.DefaultCompression, gzip.WithExcludedPaths(paths))
}
