package managelink

type GetLinkRequest struct {
	Link_Type string `json:"linktype" binding:"required"`
}
