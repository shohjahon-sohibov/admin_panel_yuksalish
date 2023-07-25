package models

type Branch struct {
	ID             string `json:"id" bson:"id"`
	Name           string `json:"name" bson:"name"`
	Director       string `json:"director" bson:"director"`
	GroupsQuantity int64  `json:"groups_quantity" bson:"groups_quantity"`
	Address        string `json:"address" bson:"address"`
	CreatedAt      string `json:"created_at" bson:"created_at"`
}

type CreateBranchRequest struct {
	Name           string `json:"name" bson:"name"`
	Director       string `json:"director" bson:"director"`
	GroupsQuantity int64  `json:"groups_quantity" bson:"groups_quantity"`
	Address        string `json:"address" bson:"address"`
	CreatedAt      string `json:"created_at" bson:"created_at"`
}

type BranchUpdate struct {
	ID             string `json:"id" bson:"id"`
	Name           string `json:"name" bson:"name"`
	Director       string `json:"director" bson:"director"`
	GroupsQuantity int64  `json:"groups_quantity" bson:"groups_quantity"`
	Address        string `json:"address" bson:"address"`
	CreatedAt      string `json:"created_at" bson:"created_at"`
}

type GetBranchListRequest struct {
	Limit  int64 `json:"limit" bson:"limit"`
	Offset int64 `json:"offset" bson:"offset"`
}

type GetBranchListResponse struct {
	Count    int64     `json:"count" bson:"count"`
	Branches []*Branch `json:"branches" bson:"branches"`
}

type BranchID struct {
	ID string `json:"id" bson:"id"`
}

type Response struct {
	Text string `json:"text" bson:"text"`
}
