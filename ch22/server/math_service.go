package server

type MathService struct {
}

type Args struct {
	A, B int
}

func (m *MathService) Add(args Args, reply *int) error {
	*reply = args.A + args.B
	return nil
}

