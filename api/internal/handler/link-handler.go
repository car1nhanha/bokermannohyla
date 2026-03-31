package handler

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"thoropa/internal/model"
	"thoropa/internal/service"
	"time"

	"github.com/gin-gonic/gin"
	gonanoid "github.com/matoous/go-nanoid/v2"
)

type LinkHandler struct {
	service *service.LinkService
}

type DataBody struct {
	Original string `json:"original"`
}

func NewLinkHandler(s *service.LinkService) *LinkHandler {
	return &LinkHandler{service: s}
}

func normalizeClientIP(raw string) string {
	ip := net.ParseIP(raw)
	if ip == nil {
		return raw
	}

	if ip.IsLoopback() {
		return "127.0.0.1"
	}

	if v4 := ip.To4(); v4 != nil {
		return v4.String()
	}

	return ip.String()
}

func (h *LinkHandler) Create(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), 3*time.Second)
	defer cancel()
	id, err := gonanoid.New(10)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Falha ao gerar identificador"})
		return
	}

	var l DataBody

	if err := c.ShouldBind(&l); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	link := model.Link{
		Id:        id,
		Ip:        normalizeClientIP(c.ClientIP()),
		CreatedAt: time.Now().Unix(),
		Accesses:  0,
		Original:  l.Original,
	}

	err = h.service.Create(ctx, &link)

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, link)
}

func (h *LinkHandler) GetById(c *gin.Context) {
	id := c.Param("id")

	link, err := h.service.FindByID(c, id)

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if link == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Link não encontrado"})
		return
	}

	c.JSON(http.StatusOK, link)
}

func (h *LinkHandler) GetByIP(c *gin.Context) {
	ip := normalizeClientIP(c.ClientIP())

	fmt.Printf("Buscando links para IP: %s\n", ip)

	links, err := h.service.FindByIP(c, ip)

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, links)
}

func (h *LinkHandler) DeleteById(c *gin.Context) {
	id := c.Param("id")

	link, err := h.service.FindByID(c, id)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if link == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Link não encontrado"})
		return
	}

	err = h.service.DeleteByID(c, id)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Link deletado com sucesso"})
}
