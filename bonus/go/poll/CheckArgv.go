package poll

func CheckArgv(args []string) *Error {
	for _, arg := range args[1:] {
		if arg == "-h" || arg == "--help" {
			return &HelpError
		}
	}
	if len(args[1:]) != 3 {
		return &ArgvError
	}
	return nil
}
