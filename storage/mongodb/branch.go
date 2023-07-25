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

type branchRepo struct {
	branchCollection *mongo.Collection
}

func NewBranchRepo(db *mongo.Database) storage.BranchI {
	return &branchRepo{
		branchCollection: db.Collection("branches"),
	}
}

func (b branchRepo) Create(ctx context.Context, req *models.CreateBranchRequest) (res *models.BranchID, err error) {
	id := primitive.NewObjectID()

	con := bson.M{
		"id":             id.Hex(),
		"name":           req.Name,
		"director":       req.Director,
		"group_quantity": req.GroupsQuantity,
		"address":        req.Address,
		"created_at":     time.Now().Format("02.01.2006 15:04"),
	}

	_, err = b.branchCollection.InsertOne(
		ctx,
		con,
	)
	if err != nil {
		fmt.Println("insert Branch err: ", err.Error())
		return nil, err
	}
	Id := con["id"].(string)
	return &models.BranchID{
		ID: Id,
	}, nil
}

func (b branchRepo) Single(ctx context.Context, in string) (resp *models.Branch, err error) {
	err = b.branchCollection.FindOne(ctx, bson.M{"id": &in}).Decode(&resp)
	if err == mongo.ErrNoDocuments {
		return &models.Branch{}, nil
	} else if err != nil {
		return nil, err
	}
	return
}

func (b branchRepo) List(ctx context.Context, req *models.GetBranchListRequest) (res *models.GetBranchListResponse, err error) {
	if req.Limit == 0 {
		req.Limit = 10
	}
	opts := options.Find().SetSort(bson.M{
		"created_at": -1,
	},
	).SetSkip(req.Offset).SetLimit(req.Limit)

	var filter bson.M

	count, err := b.branchCollection.CountDocuments(context.Background(), filter)
	if err != nil {
		fmt.Printf("count Documents error: %v", err.Error())
		return nil, err
	}

	rows, err := b.branchCollection.Find(
		context.Background(),
		filter,
		opts,
	)
	var branches = []*models.Branch{}

	if err != nil {
		fmt.Printf("Find Branch List err: %v", err.Error())
		return nil, err
	}

	if err = rows.All(context.Background(), &branches); err != nil {
		fmt.Printf("get all branch error: %s", err.Error())
		return nil, err
	}

	return &models.GetBranchListResponse{
		Count: count,
		Branches: branches,
	}, nil
}

func (b branchRepo) Update(ctx context.Context, req *models.BranchUpdate) (res *models.Response, err error) {
	updateReq := bson.M{
		"$set": bson.M{
			"name":           req.Name,
			"director":       req.Director,
			"group_quantity": req.GroupsQuantity,
			"address":        req.Address,
			"updated_at":     time.Now().Format("02.01.2006 15:04"),
		},
	}
	_, err = b.branchCollection.UpdateOne(ctx, bson.M{"id": req.ID}, updateReq)
	if err != nil {
		fmt.Printf("Error while updating branch: %v", err.Error())
		return  nil, err
	}

	return &models.Response{
		Text: "Succesfully Updated !",
	}, nil
}

func (b branchRepo) Delete(ctx context.Context, id string) (res *models.Response, err error) {
	_, err = b.branchCollection.DeleteOne(ctx, bson.M{"id": id})
	if err != nil {
		fmt.Printf("Delete Branch err: %v", err)
		return nil, err
	}
	
	return &models.Response{
		Text: "Successfully Deleted !",
	}, nil
}