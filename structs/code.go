package structs

type SignUpStruct struct {
	Id       string `bson:"_id"`
	Name     string
	Surname  string
	Email    string
	Balance  string
	Login    string
	Password string
}
