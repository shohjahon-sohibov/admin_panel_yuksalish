package mongodb

import (
	"context"
	"fmt"
	"freelance/admin_panel/models"
	"freelance/admin_panel/storage"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type studentRepo struct {
	studentCollection *mongo.Collection
}

func NewStudentRepo(db *mongo.Database) storage.StudentI {
	return &studentRepo{
		studentCollection: db.Collection("students"),
	}
}

func (s studentRepo) Create(ctx context.Context, req *models.CreateStudentRequest) (res *models.StudentID, err error) {
	id := primitive.NewObjectID()

	con := bson.M{
		"id":           id.Hex(),
		"student_id":   req.StudentID,
		"first_name":   req.FirstName,
		"last_name":    req.LastName,
		"phone_number": req.PhoneNumber,
		"teacher":      req.Teacher,
		"coordinator":  req.Coordinator,
		"branch_id":    req.BranchID,
		"group_id":     req.GroupID,
		"created_at":   time.Now().Format("02.01.2006 15:04"),
		"graduated_at": req.GraduatedAt,
	}

	_, err = s.studentCollection.InsertOne(
		ctx,
		con,
	)
	if err != nil {
		fmt.Println("insert Student err: ", err.Error())
		return nil, err
	}
	Id := con["id"].(string)
	return &models.StudentID{
		ID: Id,
	}, nil
}

func (s studentRepo) Single(ctx context.Context, in string) (*models.Student, error) {
	filter := bson.M{}
	if in != "" {
		filter["id"] = in
	}

	pipeline := []bson.M{
		{
			"$lookup": bson.M{
				"from":         "branches",
				"localField":   "branch_id",
				"foreignField": "id",
				"as":           "branch",
			},
		},
		{
			"$unwind": "$branch",
		},
		{
			"$lookup": bson.M{
				"from":         "groups",
				"localField":   "group_id",
				"foreignField": "id",
				"as":           "group",
			},
		},
		{
			"$unwind": "$group",
		},
		{
			"$match": filter,
		},
		{
			"$project": bson.M{
				"_id":          0,
				"id":           1,
				"student_id":   1,
				"first_name":   1,
				"last_name":    1,
				"phone_number": 1,
				"teacher":      1,
				"coordinator":  1,
				"branchName":   "$branch.name",
				"groupName":    "$group.name",
				"created_at":   1,
				"graduated_at": 1,
			},
		},
	}

	cursor, err := s.studentCollection.Aggregate(ctx, pipeline)
	if err == mongo.ErrNoDocuments {
		return &models.Student{}, nil
	} else if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var results []models.Student
	for cursor.Next(ctx) {
		var student models.Student
		err = cursor.Decode(&student)
		if err != nil {
			return nil, err
		}
		results = append(results, student)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	resp := &models.Student{}
	if len(results) > 0 {
		resp = &results[0]
	}

	return resp, nil
}

func (s studentRepo) List(ctx context.Context, req *models.GetStudentListRequest) (res *models.GetStudentListResponse, err error) {
	var filter bson.M

	if req.Limit == 0 {
		req.Limit = 10
	}

	if req.BranchID != "" {
		filter = bson.M{"branch_id": req.BranchID}
	}

	if req.GroupID != "" {
		filter = bson.M{"group_id": req.GroupID}
	}

	opts := options.Find().SetSort(bson.M{
		"created_at": -1,
	},
	).SetSkip(req.Offset).SetLimit(req.Limit)

	count, err := s.studentCollection.CountDocuments(context.Background(), filter)
	if err != nil {
		fmt.Printf("count Documents error: %v", err.Error())
		return nil, err
	}

	rows, err := s.studentCollection.Find(
		context.Background(),
		filter,
		opts,
	)
	var students = []*models.Student{}

	if err != nil {
		fmt.Printf("Find Student List err: %v", err.Error())
		return nil, err
	}

	if err = rows.All(context.Background(), &students); err != nil {
		fmt.Printf("get all student error: %s", err.Error())
		return nil, err
	}

	return &models.GetStudentListResponse{
		Count:    count,
		Students: students,
	}, nil
}

func (s studentRepo) Update(ctx context.Context, req *models.StudentUpdate) (res *models.Response, err error) {
	updateReq := bson.M{
		"$set": bson.M{
			"student_id":   req.StudentID,
			"first_name":   req.FirstName,
			"last_name":    req.LastName,
			"phone_number": req.PhoneNumber,
			"teacher":      req.Teacher,
			"coordinator":  req.Coordinator,
			"branch_id":    req.BranchID,
			"group_id":     req.GroupID,
			"updated_at":   time.Now().Format("02.01.2006 15:04"),
			"graduated_at": req.GraduatedAt,
		},
	}
	_, err = s.studentCollection.UpdateOne(ctx, bson.M{"id": req.ID}, updateReq)
	if err != nil {
		fmt.Printf("Error while updating student: %v", err.Error())
		return nil, err
	}

	return &models.Response{
		Text: "Succesfully Updated !",
	}, nil
}

func (s studentRepo) Delete(ctx context.Context, id string) (res *models.Response, err error) {
	_, err = s.studentCollection.DeleteOne(ctx, bson.M{"id": id})
	if err != nil {
		fmt.Printf("Delete Student err: %v", err)
		return nil, err
	}

	return &models.Response{
		Text: "Successfully Deleted !",
	}, nil
}
