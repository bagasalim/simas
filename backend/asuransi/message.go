package asuransi

type GetAsuransiRequest struct {
	Judul string `json:"judul"`
}

type UpdateAsuransiRequest struct {
	Judul             string `json:"judul"`
	Premi             int32  `json:"premi" binding:"required"`
	UangPertanggungan int64  `json:"uangpertanggungan" binding:"required"`
	Deskripsi         string `json:"deskripsi" binding:"required"`
	Syarat            string `json:"syarat" binding:"required"`
	Foto              string `json:"foto" binding:"required"`
}
