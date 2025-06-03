package pg

import "errors"

var (
	errNotConnected = errors.New("Error al conectarse a la base de datos")
	errCheckStatus  = errors.New("Error al consultar el estado de la base de datos")
)
