package record

import (
	"Projects/hangle_server/database"
	"Projects/hangle_server/types"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// Delete deletes entry for the given record
func Delete(c *gin.Context) {
	logrus.Info("Deleting entry for the given record")
	db, err := database.Open()
	if err != nil {
		retErr := fmt.Errorf("unable to open database: %w", err)
		c.Error(retErr)
		c.AbortWithStatusJSON(http.StatusBadRequest, retErr.Error())
		return
	}

	id := c.Param("id")

	logrus.Debugf("delete: %+v\n", db.Where("id = ?", id).Delete(&types.Record{}))

	resp := fmt.Sprintf("deleted entry %+v", id)
	c.JSON(http.StatusOK, resp)
}