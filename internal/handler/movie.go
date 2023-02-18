package handler

import (
	"fmt"
	"movie_restfull/internal/entity"
	"movie_restfull/internal/helper"
	"time"

	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (handler *moviehandler) GetAllMovie(context *gin.Context) {
	movies, err := handler.movieusecase.GetAllMovie(context)
	if err != nil {
		res := helper.BuildErrorResponse("Internal Server Error", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}

	res := helper.BuildResponse(true, "OK!", movies)
	context.JSON(http.StatusOK, res)
}

func (handler *moviehandler) GetMovie(context *gin.Context) {
	// Get id from request param
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		res := helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	// Get all movie from usecase
	movie, err := handler.movieusecase.GetMovie(context, id)
	if err != nil {
		res := helper.BuildErrorResponse("Internal Server Error", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}

	res := helper.BuildResponse(true, "OK!", movie)
	context.JSON(http.StatusOK, res)
}

func (handler *moviehandler) CreateMovie(context *gin.Context) {
	var requestBody entity.Movies

	// Get request body from user
	err := context.BindJSON(&requestBody)
	if err != nil {
		res := helper.BuildErrorResponse("Bad request", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	nowTime, err := time.Parse("2006-01-02 15:04:05", time.Now().Format("2006-01-02 15:04:05"))
	if err != nil {
		fmt.Println(nowTime)
		return
	}
	requestBody.CreatedAt = nowTime
	requestBody.UpdatedAt = nowTime

	// Call the usecase
	movie, err := handler.movieusecase.CreateMovie(context, &requestBody)
	if err != nil {
		res := helper.BuildErrorResponse("Internal Server Error", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}

	res := helper.BuildResponse(true, "OK!", movie)
	context.JSON(http.StatusOK, res)
}

func (handler *moviehandler) UpdateMovie(context *gin.Context) {
	var requestBody entity.Movies

	// Get id from request param
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		res := helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	// Get request body from user
	err = context.BindJSON(&requestBody)
	if err != nil {
		res := helper.BuildErrorResponse("Bad request", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	// Set id from params json
	requestBody.ID = id

	setTime := time.Now().Format("2006-01-02 15:04:05")
	datetime, err := time.Parse("2006-01-02 15:04:05", setTime)
	if err != nil {
		fmt.Println(datetime)
		return
	}
	requestBody.UpdatedAt = datetime

	// Call the usecase
	movie, err := handler.movieusecase.UpdateMovie(context, requestBody)
	if err != nil {
		res := helper.BuildErrorResponse("Internal Server Error", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}

	res := helper.BuildResponse(true, "OK!", movie)
	context.JSON(http.StatusOK, res)
}

func (handler *moviehandler) DeleteMovie(context *gin.Context) {
	// Get id from request param json
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		res := helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	err = handler.movieusecase.DeleteMovie(context, id)
	if err != nil {
		res := helper.BuildErrorResponse("Internal Server Error", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}

	res := helper.BuildResponse(true, "OK!", helper.EmptyObj{})
	context.JSON(http.StatusOK, res)
}
