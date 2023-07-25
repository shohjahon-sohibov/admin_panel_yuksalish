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

type groupRepo struct {
	groupCollection *mongo.Collection
}

func NewGroupRepo(db *mongo.Database) storage.GroupI {
	return &groupRepo{
		groupCollection: db.Collection("groups"),
	}
}

func (g groupRepo) Create(ctx context.Context, req *models.CreateGroupRequest) (res *models.GroupID, err error) {
	id := primitive.NewObjectID()

	con := bson.M{
		"id":               id.Hex(),
		"name":             req.Name,
		"teacher":          req.Teacher,
		"student_quantity": req.StudentsQuantity,
		"branch_id":        req.BranchID,
		"created_at":       time.Now().Format("02.01.2006 15:04"),
		"closed_at":        req.ClosedAt,
	}

	_, err = g.groupCollection.InsertOne(
		ctx,
		con,
	)
	if err != nil {
		fmt.Println("insert Group err: ", err.Error())
		return nil, err
	}
	Id := con["id"].(string)
	return &models.GroupID{
		ID: Id,
	}, nil
}

func (g groupRepo) Single(ctx context.Context, in string) (resp *models.Group, err error) {
	err = g.groupCollection.FindOne(ctx, bson.M{"id": &in}).Decode(&resp)
	if err == mongo.ErrNoDocuments {
		return &models.Group{}, nil
	} else if err != nil {
		return nil, err
	}
	return
}

func (g groupRepo) List(ctx context.Context, req *models.GetGroupListRequest) (res *models.GetGroupListResponse, err error) {
	var filter bson.M
	
	if req.Limit == 0 {
		req.Limit = 10
	}
	if req.BranchID != "" {
		filter = bson.M{"branch_id": req.BranchID}
	}
	opts := options.Find().SetSort(bson.M{
		"created_at": -1,
	},
	).SetSkip(req.Offset).SetLimit(req.Limit)


	count, err := g.groupCollection.CountDocuments(context.Background(), filter)
	if err != nil {
		fmt.Printf("count Documents error: %v", err.Error())
		return nil, err
	}

	rows, err := g.groupCollection.Find(
		context.Background(),
		filter,
		opts,
	)
	var groups = []*models.Group{}

	if err != nil {
		fmt.Printf("Find Group List err: %v", err.Error())
		return nil, err
	}

	if err = rows.All(context.Background(), &groups); err != nil {
		fmt.Printf("get all group error: %s", err.Error())
		return nil, err
	}

	return &models.GetGroupListResponse{
		Count:    count,
		Groups: groups,
	}, nil
}

func (g groupRepo) Update(ctx context.Context, req *models.GroupUpdate) (res *models.Response, err error) {
	updateReq := bson.M{
		"$set": bson.M{
			"name":             req.Name,
			"teacher":          req.Teacher,
			"student_quantity": req.StudentsQuantity,
			"branch_id":        req.BranchID,
			"updated_at":       time.Now().Format("02.01.2006 15:04"),
			"closed_at":        req.ClosedAt,
		},
	}
	_, err = g.groupCollection.UpdateOne(ctx, bson.M{"id": req.ID}, updateReq)
	if err != nil {
		fmt.Printf("Error while updating group: %v", err.Error())
		return nil, err
	}

	return &models.Response{
		Text: "Succesfully Updated !",
	}, nil
}

func (g groupRepo) Delete(ctx context.Context, id string) (res *models.Response, err error) {
	_, err = g.groupCollection.DeleteOne(ctx, bson.M{"id": id})
	if err != nil {
		fmt.Printf("Delete Group err: %v", err)
		return nil, err
	}

	return &models.Response{
		Text: "Successfully Deleted !",
	}, nil
}
