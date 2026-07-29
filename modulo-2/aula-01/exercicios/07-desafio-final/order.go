package order

import "errors"

type status string
type action string

const (
	statusOpen      status = "aberto"
	statusPaid      status = "pago"
	statusCancelled status = "cancelado"

	actionPay    action = "pagar"
	actionCancel action = "cancelar"
)

var (
	errAlreadyPaid       = errors.New("order already paid")
	errAlreadyCancelled  = errors.New("order already cancelled")
	errInvalidTransition = errors.New("invalid order transition")
	errInvalidAction     = errors.New("invalid order action")
)

func transitionOrder(current status, requested action) (status, error) {
	switch requested {
	case actionPay:
		switch current {
		case statusOpen:
			return statusPaid, nil
		case statusPaid:
			return statusPaid, nil
		default:
			return current, errInvalidTransition
		}

	case actionCancel:
		switch current {
		case statusOpen, statusPaid:
			return statusCancelled, nil
		case statusCancelled:
			return current, errAlreadyCancelled
		default:
			return current, errInvalidTransition
		}

	default:
		return current, errInvalidAction
	}
}
