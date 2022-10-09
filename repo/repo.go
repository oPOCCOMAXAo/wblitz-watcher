package repo

import (
	"context"

	"github.com/opoccomaxao/wblitz-watcher/wg/types"

	"github.com/opoccomaxao-go/generic-collection/slice"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Repo struct {
	client *mongo.Client
	db     *mongo.Database
	cfg    Config
}

type Config struct {
	ConnectURL string `env:"URL,required"`  // mongodb connect url
	DBName     string `env:"NAME,required"` // used db, default: wotclan
}

func New(cfg Config) (*Repo, error) {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(cfg.ConnectURL))
	if err != nil {
		return nil, errors.WithStack(err)
	}

	if err := client.Ping(context.Background(), readpref.Primary()); err != nil {
		return nil, errors.WithStack(err)
	}

	return &Repo{
		client: client,
		db:     client.Database(cfg.DBName),
		cfg:    cfg,
	}, nil
}

func (r *Repo) GetClans(
	ctx context.Context,
	regionName string,
	ids []int,
) ([]*types.ClanInfo, error) {
	bsonIDs := slice.Map(ids, func(id int) ClanID {
		return ClanID{ID: id, Region: regionName}
	})

	cur, err := r.db.Collection("clans").
		Find(ctx, bson.M{"_id": bson.M{"$in": bsonIDs}})
	if err != nil {
		return nil, errors.WithStack(err)
	}

	var res []*Clan

	err = cur.All(ctx, &res)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return slice.Map(res, Clan{}.ToType), nil
}

func (r *Repo) SaveClan(ctx context.Context, clan *types.ClanInfo) error {
	toSave := Clan{}.FromType(clan)

	typ, data, err := bson.MarshalValue(toSave)
	if err != nil {
		return errors.WithStack(err)
	}

	_, err = r.db.Collection("clans").
		UpdateByID(ctx,
			toSave.ID,
			bson.M{"$set": bson.RawValue{Type: typ, Value: data}},
			options.Update().SetUpsert(true))
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}
