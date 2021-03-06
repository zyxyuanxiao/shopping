package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/micro/go-micro/client"
	proto "github.com/ZipoChan/shopping/product/proto/product"
)

func SearchByID(w http.ResponseWriter, r *http.Request) {
	// decode the incoming request as json
	var request map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// 从前端请求获取id
	id := 0

	// call the backend service
	webClient := proto.NewProductService("go.micro.service.product", client.DefaultClient)
	rsp, err := webClient.SearchByID(context.TODO(), &proto.SearchByIDReq{
		Id: int32(id),
	})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// we want to augment the response
	response := map[string]interface{}{
		"msg": rsp.Msg,
		"ref": time.Now().UnixNano(),
	}

	// encode and write the response as json
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

func SearchByName(w http.ResponseWriter, r *http.Request) {
	// decode the incoming request as json
	var request map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	name := ""

	// call the backend service
	webClient := proto.NewProductService("go.micro.service.product", client.DefaultClient)
	rsp, err := webClient.SearchByName(context.TODO(), &proto.SearchByNameReq{
		Name: name,
	})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// we want to augment the response
	response := map[string]interface{}{
		"msg": rsp.Msg,
		"ref": time.Now().UnixNano(),
	}

	// encode and write the response as json
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

func SearchByClassify(w http.ResponseWriter, r *http.Request) {
	// decode the incoming request as json
	var request map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// 从前端请求获取classify
	classify := ""

	// call the backend service
	webClient := proto.NewProductService("go.micro.service.product", client.DefaultClient)
	rsp, err := webClient.SearchByClassify(context.TODO(), &proto.SearchByClassifyReq{
		Classify: classify,
	})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// we want to augment the response
	response := map[string]interface{}{
		"msg": rsp.Msg,
		"ref": time.Now().UnixNano(),
	}

	// encode and write the response as json
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

func SearchByTag(w http.ResponseWriter, r *http.Request) {
	// decode the incoming request as json
	var request map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// 从前端请求获取Tag
	tag := ""

	// call the backend service
	webClient := proto.NewProductService("go.micro.service.product", client.DefaultClient)
	rsp, err := webClient.SearchByTag(context.TODO(), &proto.SearchByTagReq{
		Tag: tag,
	})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// we want to augment the response
	response := map[string]interface{}{
		"msg": rsp.Msg,
		"ref": time.Now().UnixNano(),
	}

	// encode and write the response as json
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

func SortByPrice(w http.ResponseWriter, r *http.Request) {
	// decode the incoming request as json
	var request map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// 从前端请求处获取商品name
	name := ""

	// call the backend service
	webClient := proto.NewProductService("go.micro.service.product", client.DefaultClient)
	rsp, err := webClient.SortByPrice(context.TODO(), &proto.SortByPriceReq{
		Name: name,
	})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// we want to augment the response
	response := map[string]interface{}{
		"msg": rsp.Msg,
		"ref": time.Now().UnixNano(),
	}

	// encode and write the response as json
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

func SortBySalesVolume(w http.ResponseWriter, r *http.Request) {
	// decode the incoming request as json
	var request map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// 从前端请求处获取商品name
	name := ""

	// call the backend service
	webClient := proto.NewProductService("go.micro.service.product", client.DefaultClient)
	rsp, err := webClient.SortBySalesVolume(context.TODO(), &proto.SortBySalesVolumeReq{
		Name: name,
	})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// we want to augment the response
	response := map[string]interface{}{
		"msg": rsp.Msg,
		"ref": time.Now().UnixNano(),
	}

	// encode and write the response as json
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

func SortByCommentsNum(w http.ResponseWriter, r *http.Request) {
	// decode the incoming request as json
	var request map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// 从前端请求处获取商品name
	name := ""

	// call the backend service
	webClient := proto.NewProductService("go.micro.service.product", client.DefaultClient)
	rsp, err := webClient.SortByCommentsNum(context.TODO(), &proto.SortByCommentsNumReq{
		Name: name,
	})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// we want to augment the response
	response := map[string]interface{}{
		"msg": rsp.Msg,
		"ref": time.Now().UnixNano(),
	}

	// encode and write the response as json
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

func AddProduct(w http.ResponseWriter, r *http.Request) {
	// decode the incoming request as json
	var request map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// 从前端请求，初始化product
	product := &proto.Product{
		Id:           0,
		Name:         "",
		Classify:     "",
		Tag:          "",
		Price:        "",
		SalesVolume:  0,
		CommentsNum:  0,
		Inventory:    0,
		Introduction: "",
	}

	// call the backend service
	webClient := proto.NewProductService("go.micro.service.product", client.DefaultClient)
	rsp, err := webClient.AddProduct(context.TODO(), &proto.AddProductReq{
		Product: product,
	})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// we want to augment the response
	response := map[string]interface{}{
		"msg": rsp.Msg,
		"ref": time.Now().UnixNano(),
	}

	// encode and write the response as json
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	// decode the incoming request as json
	var request map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// 从前端请求，初始化product
	product := &proto.Product{
		Id:           0,
		Name:         "",
		Classify:     "",
		Tag:          "",
		Price:        "",
		SalesVolume:  0,
		CommentsNum:  0,
		Inventory:    0,
		Introduction: "",
	}

	// call the backend service
	webClient := proto.NewProductService("go.micro.service.product", client.DefaultClient)
	rsp, err := webClient.UpdateProduct(context.TODO(), &proto.UpdateProductReq{
		Product: product,
	})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// we want to augment the response
	response := map[string]interface{}{
		"msg": rsp.Msg,
		"ref": time.Now().UnixNano(),
	}

	// encode and write the response as json
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

func DelProduct(w http.ResponseWriter, r *http.Request) {
	// decode the incoming request as json
	var request map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	id := 0

	// call the backend service
	webClient := proto.NewProductService("go.micro.service.product", client.DefaultClient)
	rsp, err := webClient.DelProduct(context.TODO(), &proto.DelProductReq{
		Id: int32(id),
	})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// we want to augment the response
	response := map[string]interface{}{
		"msg": rsp.Msg,
		"ref": time.Now().UnixNano(),
	}

	// encode and write the response as json
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

