package storage

import (
	"context"
	"freelance/admin_panel/models"
)

type StorageI interface {
	Branch() BranchI
	Group() GroupI
	Student() StudentI
}

type BranchI interface {
	Create(ctx context.Context, req *models.CreateBranchRequest) (res *models.BranchID, err error)
	Single(ctx context.Context, in string) (resp *models.Branch, err error)
	List(ctx context.Context, req *models.GetBranchListRequest) (res *models.GetBranchListResponse, err error)
	Update(ctx context.Context, req *models.BranchUpdate) (res *models.Response, err error)
	Delete(ctx context.Context, id string) (res *models.Response, err error)
}

type GroupI interface {
	Create(ctx context.Context, req *models.CreateGroupRequest) (res *models.GroupID, err error)
	Single(ctx context.Context, in string) (resp *models.Group, err error)
	List(ctx context.Context, req *models.GetGroupListRequest) (res *models.GetGroupListResponse, err error)
	Update(ctx context.Context, req *models.GroupUpdate) (res *models.Response, err error)
	Delete(ctx context.Context, id string) (res *models.Response, err error)
}

type StudentI interface {
	Create(ctx context.Context, req *models.CreateStudentRequest) (res *models.StudentID, err error)
	Single(ctx context.Context, in string) (resp *models.Student, err error)
	List(ctx context.Context, req *models.GetStudentListRequest) (res *models.GetStudentListResponse, err error)
	Update(ctx context.Context, req *models.StudentUpdate) (res *models.Response, err error)
	Delete(ctx context.Context, id string) (res *models.Response, err error)
}
