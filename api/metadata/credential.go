package metadata

import (
	"bitbucket.pearson.com/apseng/tensor/models"
	"bitbucket.pearson.com/apseng/tensor/db"
	"github.com/gin-gonic/gin"
)

func CredentialMetadata(cred *models.Credential) error {

	ID := cred.ID.Hex()
	cred.Type = "credential"
	cred.Url = "/v1/credentials/" + ID + "/"
	related := gin.H{
		"created_by": "/v1/users/" + cred.CreatedByID.Hex() + "/",
		"modified_by": "/v1/users/" + cred.ModifiedByID.Hex() + "/",
		"owner_teams": "/v1/organizations/" + ID + "/users/",
		"owner_users": "/v1/organizations/" + ID + "/object_roles/",
		"activity_stream": "/v1/organizations/" + ID + "/activity_stream/",
		"access_list": "/v1/organizations/" + ID + "/access_list/",
		"object_roles": "/api/v1/credentials/" + ID + "/object_roles/",
		"user": "/api/v1/users/" + cred.CreatedByID.Hex() + "/",
	}

	if cred.OrganizationID != nil {
		related["organization"] = "/api/v1/organizations/" + *cred.OrganizationID + "/"
	}

	cred.Related = related

	if err := setCredentialSummary(cred); err != nil {
		return err
	}

	return nil
}

func setCredentialSummary(cred *models.Credential) error {
	dbu := db.C(db.USERS)

	var modified models.User
	var created models.User
	var org models.Organization
	var owners []models.User

	if err := dbu.FindId(cred.CreatedByID).One(&created); err != nil {
		return err
	}

	if err := dbu.FindId(cred.ModifiedByID).One(&modified); err != nil {
		return err
	}

	//TODO: include teams to owners list

	summary := gin.H{
		"object_roles": []gin.H{
			{
				"Description": "Can manage all aspects of the credential",
				"Name":"admin",
			},
			{
				"Description":"Can use the credential in a job template",
				"Name":"use",
			},
			{
				"Description":"May view settings for the credential",
				"Name":"read",
			},
		},
		"created_by": gin.H{
			"id":         created.ID,
			"username":   created.Username,
			"first_name": created.FirstName,
			"last_name":  created.LastName,
		},
		"modified_by": gin.H{
			"id":         modified.ID,
			"username":   modified.Username,
			"first_name": modified.FirstName,
			"last_name":  modified.LastName,
		},
		"owners": owners,
	}

	if cred.OrganizationID != nil {

		if err := dbu.FindId(cred.OrganizationID).One(&org); err != nil {
			return err
		}

		summary["organization"] = gin.H{
			"id": org.ID,
			"name": org.Name,
			"description": org.Description,
		}
	}

	cred.Summary = summary;

	return nil
}