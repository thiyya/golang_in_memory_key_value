package handlers

import (
	"golang_in_memory_key_value/cache"
	"golang_in_memory_key_value/constants"
	"golang_in_memory_key_value/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type KeyHandler struct {
	cache *cache.Cache
}

func NewKeyHandler(cache *cache.Cache) KeyHandler {
	return KeyHandler{cache}
}

func (k *KeyHandler) HandleGetValueFromKey(ctx *gin.Context) {
	id := ctx.Param("id")
	value, found := cache.GetCache().Get(id)
	if !found {
		ctx.IndentedJSON(http.StatusNotFound, models.ErrorResponse400{Message: value})
		return
	}
	ctx.IndentedJSON(http.StatusOK, models.ValueResponse{Value: value})
}

func (k *KeyHandler) HandleAddOrUpdateNewItem(ctx *gin.Context) {
	var newItem models.KeyValueItem
	if err := ctx.BindJSON(&newItem); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}
	ttl := constants.DEFAULT_TTL
	if newItem.Ttl > 0 {
		ttl = time.Second * time.Duration(newItem.Ttl)
	}
	cache.GetCache().Set(newItem.Key, newItem.Value, ttl)
	ctx.IndentedJSON(http.StatusCreated, newItem)
}

func (k *KeyHandler) HandleFlush(ctx *gin.Context) {
	cache.GetCache().Flush()
	ctx.IndentedJSON(http.StatusCreated, nil)
}
