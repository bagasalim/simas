package managelink

type GetLinkRequest struct {
	Link_Type string `json:"linktype" binding:"required"`
}

type UpdateLinkRequest struct {
	Link_Type  string `json:"linktype" binding:"required"`
	Link_Value string `json:"linkvalue" binding:"required"`
	Updated_By string `json:"updatedby" binding:"required"`
}
