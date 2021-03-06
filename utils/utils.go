package utils

import (
	"net/http"
	"strconv"

	"github.com/Blac-Panda/Stardome-API/models"
	"github.com/gin-gonic/gin"
)

// CompareStrings :
func CompareStrings(first string, second string, err error) (bool, error) {
	if first != second {
		return true, nil
	}
	return false, err
}

// ParseQueryToInt ;
func ParseQueryToInt(queries ...string) ([]int, *models.ErrorParsing) {
	var list []int = make([]int, len(queries))

	for i, q := range queries {
		v, err := strconv.Atoi(q)

		if err != nil {
			return []int{}, &models.ErrorParsing{
				Type:       gin.ErrorTypePublic,
				Error:      ErrorInvalidQuery,
				Metadata:   http.StatusText(http.StatusBadRequest),
				StatusCode: http.StatusBadRequest,
			}
		}
		list[i] = v
	}

	return list, nil
}
