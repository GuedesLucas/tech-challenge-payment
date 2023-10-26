package types

import "github.com/go-playground/validator"

type Payment struct {
	ID string `json:"id" validate:"required"`
}

func PayData(id string) (*Payment, error) {
	payData := &Payment{
		ID: id,
	}

	validate := validator.New()
	if err := validate.Struct(payData); err != nil {
		return nil, err
	}

	return payData, nil
}
