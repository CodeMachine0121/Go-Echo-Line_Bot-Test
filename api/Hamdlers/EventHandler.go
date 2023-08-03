package handlers

import models "go-line/Models"

type LineEventHandler interface {
	Handle(*models.HandleDto)
}
