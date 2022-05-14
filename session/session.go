package cognitosession

var sessionid string


type UserInfo struct {
	Name string
	BirthDay string
	Age string
}

func (u *UserInfo) set_Name(n string){
	u.Name = n
}
func (u *UserInfo) get_Name() string{
	return u.Name
}

type Credential struct {
	// encript
	AccessToken string
	RefreshToken string
}

type SampleInfo struct {
	UserInfo UserInfo
	Newuser string 
}

// the temporary solution
type SessionStruct struct {
	Credential Credential
	UserInfo UserInfo
}


// later
type SessionInteface interface {
	Set_Name(a string)
	Get_Name() string


}
