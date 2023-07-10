package structure

// var Client []*websocket.Conn
type Message struct {
	Request string `json:"Request"`
}

type Message2 struct {
	Message string `json:"message"`
}

type UserRegister struct {
	FirstName      string `json:"firstName"`
	LastName       string `json:"lastName"`
	Username       string `json:"username"`
	Gender         string `json:"gender"`
	Email          string `json:"email"`
	Password       string `json:"password"`
	RepeatPassword string `json:"repeatPassword"`
	Age            string `json:"age"`
}

type UserLogin struct {
	UsernameMail string `json:"usernameMail"`
	Password     string `json:"passwordLogin"`
}

type Posts struct {
	ContentPost string `json:"contentPost"`
}

type Cookie struct {
	Message string
	Value string
}
