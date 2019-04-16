package handlers

import (
	"fmt"
	"github.com/globalsign/mgo/bson"
	"github.com/labstack/echo"
	"gitlab.com/evzpav/documents-crud/server/db"
	"gitlab.com/evzpav/documents-crud/server/models"
	"log"
	"net/http"
	"time"
)

var queriesCounter int

type MongoSession struct {
	db         *db.Session
	collection *db.Collection
	uptime time.Time
}

func NewMongo(serverUp time.Time) *MongoSession {
	return &MongoSession{
		db: db.NewSession(),
		uptime: serverUp,
	}
}

func (m *MongoSession) CreateCollection(db, collection string) {
	m.collection = m.db.GetCollection(db, collection)
}

func (m *MongoSession) CreateDocument(c echo.Context) (err error) {
	var doc models.Document

	if err = c.Bind(&doc); err != nil {
		log.Printf("Could not unmarshal doc: %s", err.Error())
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	doc.CreatedAt = time.Now()
	doc.UpdatedAt = time.Now()
	doc.ID = bson.NewObjectId()

	err = m.collection.InsertDocument(&doc)
	queriesCounter++

	if err != nil {
		log.Printf("Could not insert doc: %s", err.Error())
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	log.Printf("Document created: %v", doc)

	return c.JSON(http.StatusOK, doc)
}

func (m *MongoSession) GetDocuments(c echo.Context) (err error) {
	var docs []models.Document
	err = m.collection.FindAllDocuments(&docs)
	queriesCounter++

	if err != nil {
		log.Printf("Could get docs: %s", err.Error())
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	log.Printf("Qty of docs found: %v", len(docs))
	return c.JSON(http.StatusOK, docs)
}

func (m *MongoSession) UpdateDocument(c echo.Context) (err error) {
	id, err := resolveID(c)
	if err != nil {
		log.Print(err)
		return err
	}

	var doc models.Document
	if err = c.Bind(&doc); err != nil {
		log.Printf("Could not unmarshal doc: %s", err.Error())
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	doc.UpdatedAt = time.Now()
	doc.ID = bson.ObjectIdHex(id)

	err = m.collection.UpdateDocument(id, &doc)
	queriesCounter++
	if err != nil {
		log.Printf("Could not update doc: %s", err.Error())
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, doc)
}

func resolveID(c echo.Context) (string, error) {
	id := c.Param("id")

	if id == "" {
		err := fmt.Errorf("invalid id")
		log.Print(err)
		return "", c.JSON(http.StatusBadRequest, err.Error())
	}
	return id, nil
}

func (m *MongoSession) DeleteDocument(c echo.Context) (err error) {
	id, err := resolveID(c)
	if err != nil {
		log.Print(err)
		return err
	}

	err = m.collection.RemoveDocument(id)
	queriesCounter++

	if err != nil {
		log.Printf("Could not delete doc: %s", err.Error())
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, "Deleted: "+id)

}

func (m *MongoSession) ServerStatus(c echo.Context) (err error) {
	var status models.ServerStatus
	uptime := time.Since(m.uptime)
	status.UpTime = uptime.Seconds()
	status.SessionQueries = queriesCounter
	return c.JSON(http.StatusOK, status)
}
