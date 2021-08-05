package domain

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
	u.ID = urepo.ID
	u.Name = urepo.Name
	u.Email = urepo.Email
}

type UserRepository struct {
	ID string
	Name string
	Password string
	Email string
}

func (u *UserRepository) FromPOSTBody(body *UsersPOSTBody){
	u.Password = body.Password
	u.Email = body.Email
	u.Name = body.Name
}