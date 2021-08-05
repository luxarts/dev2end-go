package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type UsersPOSTBody struct {
	Name string `json:"name"`
	Password string `json:"password"`
	Email string `json:"email"`
}

type UsersGETResponse struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
}

func (u *UsersGETResponse) FromUserRepository(urepo *UserRepository){
	u.ID = urepo.ID.Hex()
	u.Name = urepo.Name
	u.Email = urepo.Email
}

type UserRepository struct {
	ID primitive.ObjectID `bson:"_id,omitempty"`
	Name string `bson:"name"`
	Password string `bson:"password"`
	Email string `bson:"email"`
}

func (u *UserRepository) FromPOSTBody(body *UsersPOSTBody){
	u.Password = body.Password
	u.Email = body.Email
	u.Name = body.Name
}