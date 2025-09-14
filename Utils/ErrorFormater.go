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
			Message: "Record Not Found",
		}

	case errors.Is(err, gorm.ErrInvalidTransaction):
		return DTO.ResponseError[string]{
			Status:  http.StatusBadRequest,
			Message: "Invalid Transaction",
		}

	case errors.Is(err, gorm.ErrNotImplemented):
		return DTO.ResponseError[string]{
			Status:  http.StatusNotImplemented,
			Message: "Not Implemented",
		}

	case errors.Is(err, gorm.ErrMissingWhereClause):
		return DTO.ResponseError[string]{
			Status:  http.StatusBadRequest,
			Message: "Missing WHERE clause",
		}

	case errors.Is(err, gorm.ErrUnsupportedRelation):
		return DTO.ResponseError[string]{
			Status:  http.StatusBadRequest,
			Message: "Unsupported Relation",
		}

	case errors.Is(err, gorm.ErrPrimaryKeyRequired):
		return DTO.ResponseError[string]{
			Status:  http.StatusBadRequest,
			Message: "Primary Key Required",
		}

	case errors.Is(err, gorm.ErrModelValueRequired):
		return DTO.ResponseError[string]{
			Status:  http.StatusBadRequest,
			Message: "Model Value Required",
		}

	case errors.Is(err, gorm.ErrInvalidData):
		return DTO.ResponseError[string]{
			Status:  http.StatusBadRequest,
			Message: "Invalid Data",
		}

	case errors.Is(err, gorm.ErrUnsupportedDriver):
		return DTO.ResponseError[string]{
			Status:  http.StatusBadRequest,
			Message: "Unsupported Driver",
		}

	case errors.Is(err, gorm.ErrRegistered):
		return DTO.ResponseError[string]{
			Status:  http.StatusConflict,
			Message: "Already Registered",
		}

	case errors.Is(err, gorm.ErrEmptySlice):
		return DTO.ResponseError[string]{
			Status:  http.StatusBadRequest,
			Message: "Empty Slice Provided",
		}

	case errors.Is(err, gorm.ErrDryRunModeUnsupported):
		return DTO.ResponseError[string]{
			Status:  http.StatusBadRequest,
			Message: "DryRun Mode Unsupported",
		}

	default:
		return DTO.ResponseError[string]{
			Status:  http.StatusInternalServerError,
			Message: http.StatusText(http.StatusInternalServerError),
			Error:   err.Error(),
		}
	}
}
