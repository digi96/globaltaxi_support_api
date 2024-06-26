package controllers

import (
	"context"
	"database/sql"
	db "example/web-service-gin/db/sqlc"
	"example/web-service-gin/schemas"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type PayloadController struct {
	db  *db.Queries
	ctx context.Context
}

func NewPayloadController(db *db.Queries, ctx context.Context) *PayloadController {
	return &PayloadController{db, ctx}
}

// Create payload  handler
func (cc *PayloadController) CreatePayload(ctx *gin.Context) {
	var payloadTempalte *schemas.CreatePayload

	if err := ctx.ShouldBindJSON(&payloadTempalte); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "Failed payload", "error": err.Error()})
		return
	}

	now := time.Now()
	args := &db.CreatePayloadParams{
		Body:      payloadTempalte.Body,
		CreatedAt: now,
	}

	payload, err := cc.db.CreatePayload(ctx, *args)

	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "Failed retrieving payload", "error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "successfully created payload", "payload": payload})
}

// Update contact handler
func (cc *PayloadController) UpdatePayload(ctx *gin.Context) {
	payloadId := ctx.Param("payloadId")

	pid, err := strconv.ParseInt(payloadId, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "Failed retrieving payload", "error": err.Error()})
		return
	}

	payload, err := cc.db.UpdatePayload(ctx, int32(pid))

	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, gin.H{"status": "failed", "message": "Failed to retrieve payload with this ID"})
			return
		}
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "Failed retrieving payload", "error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "successfully updated payload", "payload": payload})
}

func (cc *PayloadController) GetUndoPayloads(ctx *gin.Context) {
	var page = ctx.DefaultQuery("page", "1")
	var limit = ctx.DefaultQuery("limit", "10")

	reqPageID, _ := strconv.Atoi(page)
	reqLimit, _ := strconv.Atoi(limit)
	offset := (reqPageID - 1) * reqLimit

	args := &db.ListUndoPayloadsParams{
		Limit:  int32(reqLimit),
		Offset: int32(offset),
	}

	payloads, err := cc.db.ListUndoPayloads(ctx, *args)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "Failed to retrieve undo payloads", "error": err.Error()})
		return
	}

	if payloads == nil {
		payloads = []db.Payload{}
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "Successfully retrieved undo payloads", "size": len(payloads), "payloads": payloads})
}
