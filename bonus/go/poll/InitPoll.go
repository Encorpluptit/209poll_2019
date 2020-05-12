package poll

import "strconv"

func InitPoll(args []string) (*Poll, *Error) {
	pSize, err := strconv.ParseInt(args[0], 10, 64)
	if err != nil || pSize < 0 {
		return nil, &IntError
	}
	sSize, err := strconv.ParseInt(args[1], 10, 64)
	if err != nil || sSize < 0 {
		return nil, &IntError
	}
	p, err := strconv.ParseFloat(args[2], 64)
	if err != nil || p < 0 || p > 100 {
		return nil, &FloatError
	}

	return &Poll{
		pSize: pSize,
		sSize: sSize,
		p:     p,
	}, nil
}
