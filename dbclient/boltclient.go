package dbclient

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	"github.com/boltdb/bolt"
	"github.com/leomarquezani/rest-api/model"
)

type IBoltClient interface {
	OpenBoltDb()
	QueryAccount(accountId string) (model.Account, error)
	Seed()
	Check() bool
}

type BoltClient struct {
	boltDB *bolt.DB
}

func (bc *BoltClient) OpenBoltDb() {
	var err error
	bc.boltDB, err = bolt.Open("accounts.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func (bc *BoltClient) Seed() {
	bc.initializeBucket()
	bc.seedAccounts()
}

func (bc *BoltClient) initializeBucket() {
	bc.boltDB.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucket([]byte("AccountBucket"))
		if err != nil {
			return fmt.Errorf("Create Bucket Failed %s", err)
		}
		return nil
	})
}
func (bc *BoltClient) seedAccounts() {
	total := 100

	for i := 0; i < total; i++ {
		//Generate a key
		key := strconv.Itoa(10000 + i)

		//create instace of account struct
		acc := model.Account{
			Id:   key,
			Name: "Person_" + strconv.Itoa(i),
		}

		//Serialize JSON
		jsonBytes, _ := json.Marshal(acc)

		//Write data to AccountBucket
		bc.boltDB.Update(func(tx *bolt.Tx) error {
			b := tx.Bucket([]byte("AccountBucket"))
			err := b.Put([]byte(key), jsonBytes)
			return err
		})
	}
	fmt.Printf("Seeded %v fake Accounts...", total)
}

func (bc *BoltClient) QueryAccount(accountId string) (model.Account, error) {

	acc := model.Account{}

	err := bc.boltDB.View(func(tx *bolt.Tx) error {

		b := tx.Bucket([]byte("AccountBucket"))

		accountBytes := b.Get([]byte(accountId))

		if accountBytes == nil {
			return fmt.Errorf("No account found for " + accountId)
		}
		json.Unmarshal(accountBytes, &acc)

		return nil
	})

	if err != nil {
		return model.Account{}, err
	}
	return acc, nil
}

func (bc *BoltClient) Check() bool {
	return bc.boltDB != nil
}
