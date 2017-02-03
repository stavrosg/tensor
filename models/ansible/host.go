package ansible

import (
	"time"

	"gopkg.in/gin-gonic/gin.v1"
	"gopkg.in/mgo.v2/bson"
)

type Host struct {
	ID bson.ObjectId `bson:"_id" json:"id"`

	// required
	Name        string         `bson:"name" json:"name" binding:"required,iphost"`
	InventoryID bson.ObjectId  `bson:"inventory_id" json:"inventory" binding:"required"`
	Description string         `bson:"description,omitempty" json:"description"`
	GroupID     *bson.ObjectId `bson:"group_id,omitempty" json:"group"`
	InstanceID  string         `bson:"instance_id,omitempty" json:"instance_id"`
	Variables   string         `bson:"variables,omitempty" json:"variables"`
	Enabled     bool           `bson:"enabled,omitempty" json:"enabled"`

	LastJobID            *bson.ObjectId `bson:"last_job_id,omitempty" json:"last_job" binding:"omitempty,naproperty"`
	LastJobHostSummaryID *bson.ObjectId `bson:"last_job_host_summary_id,omitempty" json:"last_job_host_summary" binding:"omitempty,naproperty"`

	HasActiveFailures   bool          `bson:"has_active_failures,omitempty" json:"has_active_failures" binding:"omitempty,naproperty"`
	HasInventorySources bool          `bson:"has_inventory_sources,omitempty" json:"has_inventory_sources" binding:"omitempty,naproperty"`
	CreatedByID         bson.ObjectId `bson:"created_by_id" json:"-"`
	ModifiedByID        bson.ObjectId `bson:"modified_by_id" json:"-"`
	Created             time.Time     `bson:"created" json:"created" binding:"omitempty,naproperty"`
	Modified            time.Time     `bson:"modified" json:"modified" binding:"omitempty,naproperty"`

	Type    string `bson:"-" json:"type"`
	URL     string `bson:"-" json:"url"`
	Related gin.H  `bson:"-" json:"related"`
	Summary gin.H  `bson:"-" json:"summary_fields"`
}

type PatchHost struct {
	Name        *string        `json:"name" binding:"omitempty,iphost"`
	InventoryID *bson.ObjectId `json:"inventory"`
	Description *string        `json:"description"`
	GroupID     *bson.ObjectId `json:"group"`
	InstanceID  *string        `json:"instance_id"`
	Variables   *string        `json:"variables"`
	Enabled     *bool          `json:"enabled"`
}