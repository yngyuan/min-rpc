package rpcdemo

import (
	"errors"
)

// Service.Method

type DemoService struct {}

type Args struct {
	A, B int
}
func (DemoService) Div(
	args *Args, result *float64) error {
	if args.B == 0 {
		return errors.New("division by zero")
	}

	*result = float64(args.A)/ float64(args.B)
	return nil
}

func (DemoService) Add(
	args * Args, result *float64) error{

	*result = float64(args.A) + float64((args.B))
	return nil
}

func (DemoService) Sub(
	args * Args, result *float64) error{

	*result = float64(args.A) - float64((args.B))
	return nil
}

func (DemoService) Multiply(
	args * Args, result *float64) error{

	*result = float64(args.A) * float64((args.B))
	return nil
}
