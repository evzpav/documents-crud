package models

import (
	"github.com/globalsign/mgo/bson"
	"time"
)

type Document struct {
	ID            bson.ObjectId `json:"id" bson:"_id,omitempty"`
	DocType       string        `json:"docType" bson:"doc_type"`
	IsBlacklisted bool          `json:"isBlacklisted" bson:"is_blacklisted"`
	Value         string        `json:"value" bson:"value"`
	CreatedAt     time.Time     `json:"createdAt" bson:"created_at"`
	UpdatedAt     time.Time     `json:"updatedAt" bson:"updated_at"`
}

type ServerStatus struct {
	UpTime         float64 `json:"uptime"`
	SessionQueries int     `json:"sessionQueries"`
}
