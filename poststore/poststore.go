package poststore

import (
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
	"github.com/hashicorp/consul/api"
	"github.com/veljkomaksimovic/nginx-example/model"
)

type PostStore struct {
	cli *api.Client
}

func New() (*PostStore, error) {
	//db := os.Getenv("DB")
	db := "consul"
	//dbport := os.Getenv("DBPORT")
	dbport := "8500"
	db_config := api.DefaultConfig()
	db_config.Address = fmt.Sprintf("%s:%s", db, dbport)
	client, err := api.NewClient(db_config)
	if err != nil {
		return nil, err
	}

	return &PostStore{
		cli: client,
	}, nil
}

func (ps *PostStore) Get(id string) (*model.Consumer, error) {

	kv := ps.cli.KV()

	//drugi parametar su napredne opcije za upit, a druga povratna vrednost su metapodaci upita
	pair, _, err := kv.Get(id, nil)
	if err != nil {
		return nil, err
	}

	consumer := &model.Consumer{}
	err = json.Unmarshal(pair.Value, consumer)
	if err != nil {
		return nil, err
	}

	return consumer, nil
}

func (ps *PostStore) Post(consumer *model.Consumer) (*model.Consumer, error) {
	kv := ps.cli.KV()
	id := generateKey()
	consumer.Id = id

	data, err := json.Marshal(consumer)
	if err != nil {
		return nil, err
	}

	kv_pair := &api.KVPair{Key: id, Value: data}
	fmt.Println(*kv_pair)
	_, err = kv.Put(kv_pair, nil)
	if err != nil {
		fmt.Println("test3")
		return nil, err
	}

	return consumer, nil
}

func generateKey() string {
	id := uuid.New().String()
	fmt.Println(id)
	return id
}
