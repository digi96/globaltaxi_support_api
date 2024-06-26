package schemas

type CreatePayload struct {
	Body string `json:"body" binding:"required"`
}
