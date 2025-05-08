package do

type In struct {
	UserId    int
	Role      string
	ExpiresIn string
}

type Out struct {
	Result string
}

func Do(in In) (out Out) {
	return Out{Result: "result ok"}
}
