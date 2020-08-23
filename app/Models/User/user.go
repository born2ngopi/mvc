package User

type UserModelRepo interface {
	GetAllUser()
}

type user struct {}

func Init() UserModelRepo {
	return &user{}
}

func (*user) GetAllUser(){

}

