package poll

func CheckArgv(args []string) *Error {
	for _, arg := range args {
		if arg == "-h" || arg == "--help" {
			return &HelpError
		}
	}
	if len(args) != 3 {
		return &ArgvError
	}
	return nil
}
