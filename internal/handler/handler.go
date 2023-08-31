package handler

import (
	"UserSegmentationService/internal/repository"
	"UserSegmentationService/internal/service"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"net/http"
	"strconv"
)

type Handler struct {
	serv *service.Service
}

func NewHandler(serv *service.Service) *Handler {
	return &Handler{serv: serv}
}

func (h *Handler) CreateSegment(ctx *gin.Context) {
	segment, err := GetSegmentFromContext(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = h.serv.CreateSegment(repository.Segment{Name: segment})

	if err != nil {
		if errors.Is(err, service.ErrSegmentExists) {
			ctx.JSON(http.StatusConflict, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

}

func (h *Handler) DeleteSegment(ctx *gin.Context) {
	segment, err := GetSegmentFromContext(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = h.serv.DeleteSegment(repository.Segment{Name: segment})

	if err != nil {
		if errors.Is(err, service.ErrNoSegment) {
			ctx.JSON(http.StatusConflict, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
}

func (h *Handler) ChangeUserSegments(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "id is empty"})
		return
	}

	var req RequestChange
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	h.serv.ChangeUserSegments(
		repository.User{Id: id},
		req.ToAddSegment(),
		req.ToDeleteSegment(),
	)

	if err != nil {
		if errors.Is(err, service.ErrNoUser) {
			ctx.JSON(http.StatusConflict, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
}

func (h *Handler) ShowUserSegments(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "id is empty"})
		return
	}

	segments, err := h.serv.ShowUserSegments(repository.User{Id: id})
	if err != nil {
		if errors.Is(err, service.ErrNoUser) {
			ctx.JSON(http.StatusConflict, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "id is empty"})
		return
	}

	result, _ := json.Marshal(segments)
	ctx.JSON(http.StatusOK, string(result))
}

func (h *Handler) InitRouter() *gin.Engine {
	router := gin.New()

	router.POST("/segments", h.CreateSegment)
	router.DELETE("/segments", h.DeleteSegment)
	router.PUT("/users/:id", h.ChangeUserSegments)
	router.GET("/users/:id", h.ShowUserSegments)

	return router
}
