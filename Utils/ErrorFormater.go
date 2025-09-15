package Utils

import (
	"SSO_BE_API/Model/DTO/Response"
	"errors"
	"gorm.io/gorm"
	"net/http"
)

func ErrorFormater(err error) DTO.ResponseError[string] {
	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		return DTO.ResponseError[string]{
			Status:  http.StatusNotFound,
			Message: http.StatusText(http.StatusNotFound),
		}

	case errors.Is(err, gorm.ErrInvalidTransaction):
		return DTO.ResponseError[string]{
			Status:  http.StatusBadRequest,
			Message: http.StatusText(http.StatusBadRequest),
		}

	case errors.Is(err, gorm.ErrNotImplemented):
		return DTO.ResponseError[string]{
			Status:  http.StatusNotImplemented,
			Message: http.StatusText(http.StatusNotImplemented),
		}

	case errors.Is(err, gorm.ErrMissingWhereClause):
		return DTO.ResponseError[string]{
			Status:  http.StatusBadRequest,
			Message: http.StatusText(http.StatusBadRequest),
		}

	case errors.Is(err, gorm.ErrUnsupportedRelation):
		return DTO.ResponseError[string]{
			Status:  http.StatusBadRequest,
			Message: http.StatusText(http.StatusBadRequest),
		}

	case errors.Is(err, gorm.ErrPrimaryKeyRequired):
		return DTO.ResponseError[string]{
			Status:  http.StatusBadRequest,
			Message: http.StatusText(http.StatusBadRequest),
		}

	case errors.Is(err, gorm.ErrModelValueRequired):
		return DTO.ResponseError[string]{
			Status:  http.StatusBadRequest,
			Message: http.StatusText(http.StatusBadRequest),
		}

	case errors.Is(err, gorm.ErrInvalidData):
		return DTO.ResponseError[string]{
			Status:  http.StatusBadRequest,
			Message: http.StatusText(http.StatusBadRequest),
		}

	case errors.Is(err, gorm.ErrUnsupportedDriver):
		return DTO.ResponseError[string]{
			Status:  http.StatusBadRequest,
			Message: http.StatusText(http.StatusBadRequest),
		}

	case errors.Is(err, gorm.ErrRegistered):
		return DTO.ResponseError[string]{
			Status:  http.StatusConflict,
			Message: http.StatusText(http.StatusConflict),
		}

	case errors.Is(err, gorm.ErrEmptySlice):
		return DTO.ResponseError[string]{
			Status:  http.StatusBadRequest,
			Message: http.StatusText(http.StatusBadRequest),
		}

	case errors.Is(err, gorm.ErrDryRunModeUnsupported):
		return DTO.ResponseError[string]{
			Status:  http.StatusBadRequest,
			Message: http.StatusText(http.StatusBadRequest),
		}

	default:
		return DTO.ResponseError[string]{
			Status:  http.StatusInternalServerError,
			Message: http.StatusText(http.StatusInternalServerError),
			Error:   err.Error(),
		}
	}
}
