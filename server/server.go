package server

import (
	"golang_in_memory_key_value/cache"
	"golang_in_memory_key_value/handlers"

	"github.com/gin-gonic/gin"
)

func NewServer() *gin.Engine {
	cache.GetCache()

	router := gin.Default()

	// swagger:route GET /keys/{id} get-key
	// Verilen key'e ait değeri döndürür.
	//
	// responses:
	//   200: ValueResponse
	//   400: ErrorResponse400
	//   default: GenericErrorResponse
	router.GET("/keys/:id", func(ctx *gin.Context) {
		keyHandler := handlers.NewKeyHandler(cache.GetCache())
		keyHandler.HandleGetValueFromKey(ctx)
	})

	// swagger:route POST /keys create-key
	// Yeni bir key oluşturur ya da günceller.
	//
	// responses:
	//   201: KeyValueItemResponse
	//   400: ErrorResponse400
	//   default: GenericErrorResponse
	router.POST("/keys", func(ctx *gin.Context) {
		keyHandler := handlers.NewKeyHandler(cache.GetCache())
		keyHandler.HandleAddOrUpdateNewItem(ctx)
	})

	// swagger:route DELETE /keys flush-keys
	// Bütün key-value itemlarini siler
	//
	// responses:
	//   200: EmptyResponse
	//   400: ErrorResponse400
	//   default: GenericErrorResponse
	router.DELETE("/keys", func(ctx *gin.Context) {
		keyHandler := handlers.NewKeyHandler(cache.GetCache())
		keyHandler.HandleFlush(ctx)

	})

	return router
}
