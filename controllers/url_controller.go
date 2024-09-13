package controllers

import (
	"SWOYO/models"
	"SWOYO/store"
	"net/http"

	"github.com/gin-gonic/gin"
)

type URLController struct {
	storage store.Store
}

func NewURLController(storage store.Store) *URLController {
	return &URLController{storage: storage}
}

func (c *URLController) HandlePost(ctx *gin.Context) {
	originalURL := ctx.PostForm("originalURL")

	if originalURL == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "originalURL is required"})
		return
	}

	// Генерация короткого URL
	shortURL, err := models.GenerateShortURL(originalURL, c.storage)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate short URL"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"shortURL": shortURL})
}

func (c *URLController) HandleGet(ctx *gin.Context) {
	shortURL := ctx.Param("shortURL")

	originalURL, err := c.storage.GetOriginalURL(shortURL)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "URL not found"})
		return
	}

	ctx.Redirect(http.StatusMovedPermanently, originalURL)
}
