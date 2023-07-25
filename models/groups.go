package models

type Group struct {
	ID               string `json:"id" bson:"id"`
	Name             string `json:"name" bson:"name"`
	Teacher          string `json:"teacher" bson:"teacher"`
	StudentsQuantity int64  `json:"students_quantity" bson:"students_quantity"`
	BranchID         string `json:"branch_id" bson:"branch_id"`
	CreatedAt        string `json:"created_at" bson:"created_at"`
	ClosedAt         string `json:"closed_at" bson:"closed_at"`
}

type CreateGroupRequest struct {
	Name             string `json:"name" bson:"name"`
	Teacher          string `json:"teacher" bson:"teacher"`
	StudentsQuantity int64  `json:"students_quantity" bson:"students_quantity"`
	BranchID         string `json:"branch_id" bson:"branch_id"`
	CreatedAt        string `json:"created_at" bson:"created_at"`
	ClosedAt         string `json:"closed_at" bson:"closed_at"`
}

type GroupUpdate struct {
	ID               string `json:"id" bson:"id"`
	Name             string `json:"name" bson:"name"`
	Teacher          string `json:"teacher" bson:"teacher"`
	StudentsQuantity int64  `json:"students_quantity" bson:"students_quantity"`
	BranchID         string `json:"branch_id" bson:"branch_id"`
	CreatedAt        string `json:"created_at" bson:"created_at"`
	ClosedAt         string `json:"closed_at" bson:"closed_at"`
}

type GetGroupListRequest struct {
	Limit  int64 `json:"limit" bson:"limit"`
	Offset int64 `json:"offset" bson:"offset"`
	BranchID string `json:"branch_id" bson:"branch_id"`
}

type GetGroupListResponse struct {
	Count  int64    `json:"count" bson:"count"`
	Groups []*Group `json:"groups" bson:"groups"`
}

type GroupID struct {
	ID string `json:"id" bson:"id"`
}
