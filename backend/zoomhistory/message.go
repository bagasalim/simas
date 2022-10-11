package zoomhistory

type ZoomHistoryRequest struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Kategori string `json:"kategori" binding:"required"`
	Keluhan  string `json:"keluhan" binding:"required"`
}
