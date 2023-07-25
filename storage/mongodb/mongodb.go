package mongodb

import (
	"freelance/admin_panel/storage"

	"go.mongodb.org/mongo-driver/mongo"
)

type storagePg struct {
	branchRepo storage.BranchI
	groupRepo storage.GroupI
	studentRepo storage.StudentI
}

func NewStoragePg(db *mongo.Database) storage.StorageI {
	return &storagePg{
		branchRepo: NewBranchRepo(db),
		groupRepo: NewGroupRepo(db),
		studentRepo: NewStudentRepo(db),
	}
}

func (s *storagePg) Branch() storage.BranchI {
	return s.branchRepo
}

func (s * storagePg) Group() storage.GroupI {
	return s.groupRepo
}

func (s * storagePg) Student() storage.StudentI {
	return s.studentRepo
}