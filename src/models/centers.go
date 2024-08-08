package models

import (
	"errors"
)

var (
	ErrorCenterFound = errors.New("centro asistencial no encontrado")
    ErrorCenterNotCreated = errors.New("centro asistencial no creado")
    ErrorCenterNotUpdated= errors.New("centro asistencial no actualizado")
    ErrorCenterNotDeleted= errors.New("centro asistencial no actualizado")
)

type Center struct {
	Id           *int `json:"id" db:"id"`
	UrbId		 *int `json:"urb_id"`
	Type    	 *string `json:"type"`
	Name         *string `json:"name"`
	Address      *string `json:"address"`
	Phone        *string `json:"phone"`
}
