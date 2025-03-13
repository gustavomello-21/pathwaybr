package destination

import (
	"encoding/json"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gustavomello-21/pathwaybr-backend/apps/pathwaybr/adapter/controllers/api/destination/dto"
)

type CategoryController struct{}

func NewCategoryController() *CategoryController {
	return &CategoryController{}
}

const categoriesFile = "/home/gustavo.melo/Documents/pathwaybr/pathwaybr-backend/internal/data/categories.json"

func (c *CategoryController) Index(httpContext *gin.Context) {
	file, err := os.Open(categoriesFile)
	if err != nil {
		httpContext.JSON(400, gin.H{"error": err.Error()})
		return
	}
	defer file.Close()

	var categories []dto.CategoryDto
	err = json.NewDecoder(file).Decode(&categories)
	if err != nil {
		httpContext.JSON(400, gin.H{"error": err.Error()})
		return
	}

	httpContext.JSON(200, categories)
}
