package swagger

	type BidRequest struct {
		Id int64 `json:"id" validate:"required" example:"1"`
		Ttl  int `json:"ttl" validate:"required" example:"1200"`
		Price  int `json:"price" validate:"required" example:"500"`
	}
	
