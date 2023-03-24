package controller

import (
	"encoding/json"
	"net/http"
	"ryan-test-backend/app"
	"ryan-test-backend/helper"
	"ryan-test-backend/model/domain"
	"ryan-test-backend/model/web"
	"ryan-test-backend/service"
	"time"

	"github.com/julienschmidt/httprouter"
)

type ProductControllerImpl struct {
	ProductService service.ProductService
}

func NewProductController(productService service.ProductService) ProductController {
	return &ProductControllerImpl{
		ProductService: productService,
	}
}

func (controller *ProductControllerImpl) Create(writer http.ResponseWriter, request *http.Request, Params httprouter.Params) {
	productCreateRequest := web.ProductCreateRequest{}
	helper.ReadFromRequestBody(request, &productCreateRequest)
	productResponse := controller.ProductService.Create(request.Context(), productCreateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   productResponse,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *ProductControllerImpl) GetProductsSorted(writer http.ResponseWriter, request *http.Request, Params httprouter.Params) {
	// Check if data is in Redis cache
	redisClient := app.NewRedisClient()

	nameSort := request.URL.Query().Get("sort")
	redisKey := "products:" + nameSort
	redisValue, err := redisClient.Get(redisKey).Result()
	if err == nil {
		var products []domain.Product
		json.Unmarshal([]byte(redisValue), &products)
		webResponse := web.WebResponse{
			Code:   200,
			Status: "OK",
			Data:   products,
		}
		helper.WriteToResponseBody(writer, webResponse)
		return
	}
	productResponse := controller.ProductService.GetProductsSorted(request.Context(), nameSort)
	byteRedisValue, _ := json.Marshal(productResponse)
	redisClient.Set(redisKey, string(byteRedisValue), 10*time.Second)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   productResponse,
	}
	helper.WriteToResponseBody(writer, webResponse)
}
