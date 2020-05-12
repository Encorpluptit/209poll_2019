package poll

type Error struct {
	Code    int
	Message string
}

var FloatError = Error{
	Code:    84,
	Message: "Argument must be a float superior to 0\n",
}

var IntError = Error{
	Code:    84,
	Message: "Argument must be a int superior to 0\n",
}

var ArgvError = Error{
	Code:    84,
	Message: "The program must have 3 arguments\n",
}

var HelpError = Error{
	Code:    0,
	Message: "",
}
