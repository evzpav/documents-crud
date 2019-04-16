package db

import (
	"fmt"
	"github.com/globalsign/mgo/bson"
	"gitlab.com/evzpav/documents-crud/server/models"
	"log"
	"github.com/globalsign/mgo"
	"os"
)

type Session struct {
	session *mgo.Session
}

type Collection struct {
	collection *mgo.Collection
}

func NewSession() *Session {
	mongoInfo := fmt.Sprintf("mongodb://%s:%s", os.Getenv("MONGO_HOST"), os.Getenv("MONGO_PORT"))
	log.Println("Mongo URL: " + mongoInfo)
	s, err := mgo.Dial(mongoInfo)

	if err != nil {
		log.Println(err)
	}
	return &Session{
		session: s,
	}
}

func (s *Session) GetCollection(db string, col string) *Collection {
	return &Collection{
		collection: s.session.DB(db).C(col),
	}
}

func (c *Collection) FindDocumentByID(id string, doc *models.Document) error {
	return c.collection.FindId(bson.ObjectIdHex(id)).One(&doc)
}

func (c *Collection) FindAllDocuments(docs *[]models.Document) error {
	return c.collection.Find(bson.M{}).Sort("-updated_at").All(docs)
}

func (c *Collection) InsertDocument(doc *models.Document) error {
	return c.collection.Insert(doc)
}

func (c *Collection) UpdateDocument(id string, doc *models.Document) error {
	return c.collection.UpdateId(bson.ObjectIdHex(id), doc)
}

func (c *Collection) RemoveDocument(id string) error {
	return c.collection.RemoveId(bson.ObjectIdHex(id))
}
