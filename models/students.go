package models

type Student struct {
	ID          string `json:"id" bson:"id"`
	StudentID   string `json:"student_id" bson:"student_id"`
	FirstName   string `json:"first_name" bson:"first_name"`
	LastName    string `json:"last_name" bson:"last_name"`
	PhoneNumber string `json:"phone_number" bson:"phone_number"`
	Teacher     string `json:"teacher" bson:"teacher"`
	Coordinator string `json:"coordinator" bson:"coordinator"`
	BranchID    string `json:"branch_id" bson:"branch_id"`
	GroupID     string `json:"group_id" bson:"group_id"`
	CreatedAt   string `json:"created_at" bson:"created_at"`
	GraduatedAt string `json:"graduated_at" bson:"graduated_at"`
}

type CreateStudentRequest struct {
	StudentID   string `json:"student_id" bson:"student_id"`
	FirstName   string `json:"first_name" bson:"first_name"`
	LastName    string `json:"last_name" bson:"last_name"`
	PhoneNumber string `json:"phone_number" bson:"phone_number"`
	Teacher     string `json:"teacher" bson:"teacher"`
	Coordinator string `json:"coordinator" bson:"coordinator"`
	BranchID    string `json:"branch_id" bson:"branch_id"`
	GroupID     string `json:"group_id" bson:"group_id"`
	CreatedAt   string `json:"created_at" bson:"created_at"`
	GraduatedAt string `json:"graduated_at" bson:"graduated_at"`
}

type StudentUpdate struct {
	ID          string `json:"id" bson:"id"`
	StudentID   string `json:"student_id" bson:"student_id"`
	FirstName   string `json:"first_name" bson:"first_name"`
	LastName    string `json:"last_name" bson:"last_name"`
	PhoneNumber string `json:"phone_number" bson:"phone_number"`
	Teacher     string `json:"teacher" bson:"teacher"`
	Coordinator string `json:"coordinator" bson:"coordinator"`
	BranchID    string `json:"branch_id" bson:"branch_id"`
	GroupID     string `json:"group_id" bson:"group_id"`
	UpdatedAt   string `json:"updated_at" bson:"updated_at"`
	GraduatedAt string `json:"graduated_at" bson:"graduated_at"`
}

type GetStudentListRequest struct {
	Limit    int64  `json:"limit" bson:"limit"`
	Offset   int64  `json:"offset" bson:"offset"`
	BranchID string `json:"branch_id" bson:"branch_id"`
	GroupID  string `json:"group_id" bson:"group_id"`
}

type GetStudentListResponse struct {
	Count    int64      `json:"count" bson:"count"`
	Students []*Student `json:"students" bson:"students"`
}

type StudentID struct {
	ID string `json:"id" bson:"id"`
}
