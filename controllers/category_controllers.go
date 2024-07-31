package controllers

import (
	"crud-api/configs"
	"crud-api/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)
// Read All Categories
func ReadAllCategories(c echo.Context) (err error) {
	var responses []models.CategoryResponse

	// Buat koneksi ke database
	db, err := configs.ConnectDatabase()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Error connecting to database!",
			"page":    nil,
			"data":    nil,
			"error":   err.Error(),
		})
	}
	defer db.Close()

	const readAllCategoriesQuery = `
	SELECT
		id, category_name
	FROM
		categories
	`

	rows, err := db.QueryContext(c.Request().Context(), readAllCategoriesQuery)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Error reading all categories!",
			"page":    nil,
			"data":    nil,
			"error":   err.Error(),
		})
	}

	for rows.Next() {
		var response models.CategoryResponse

		err = rows.Scan(
			&response.ID,
			&response.CategoryName,
		)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{
				"message": "Error reading all categories!",
				"page":    nil,
				"data":    nil,
				"error":   err.Error(),
			})
		}

		responses = append(responses, response)
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "Success reading all categories!",
		"page":    nil,
		"data":    responses,
		"error":   nil,
	})
}


// Read Detail Category
func ReadDetailCategories(c echo.Context) (err error) {
	var response models.CategoryResponse

	// Buat koneksi ke database
	db, err := configs.ConnectDatabase()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Error connecting to database!",
			"page":    nil,
			"data":    nil,
			"error":   err.Error(),
		})
	}
	defer db.Close()

	idParam := c.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Error parsing parameter to integer!",
			"page":    nil,
			"data":    nil,
			"error":   err.Error(),
		})
	}

	const readDetailCategoryQuery = `
	SELECT
		categories.id, categories.category_name
	FROM
		categories
	INNER JOIN products ON categories.product_id = products.id 		
	WHERE
		categories.id = ? 
	`

	row := db.QueryRowContext(c.Request().Context(), readDetailCategoryQuery, id)

	err = row.Scan(
		&response.ID,
		&response.CategoryName,
	
	)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Error reading detail category!",
			"page":    nil,
			"data":    nil,
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "Success reading all categories!",
		"page":    nil,
		"data":    response,
		"error":   nil,
	})
}

// Create Category
func CreateCategory(c echo.Context) (err error) {
	var request models.CategoryRequest

	err = c.Bind(&request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Error binding request!",
			"page":    nil,
			"data":    nil,
			"error":   err.Error(),
		})
	}

	db, err := configs.ConnectDatabase()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Error connecting to database!",
			"page":    nil,
			"data":    nil,
			"error":   err.Error(),
		})
	}
	defer db.Close()

	const createCategoryQuery = `
	INSERT INTO categories
	(category_name)
	VALUES
	(?, ?, ?, ?)
	`

	_, err = db.ExecContext(c.Request().Context(), createCategoryQuery,
		request.CategoryName,
	
	)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Error creating data category!",
			"page":    nil,
			"data":    nil,
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, echo.Map{
		"message": "Success creating data category!",
		"page":    nil,
		"data":    nil,
		"error":   nil,
	})
}

// Update Category

// Delete Category
func DeleteCategory(c echo.Context) (err error) {
	db, err := configs.ConnectDatabase()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "failed connect to database",
			"page":    nil,
			"data":    nil,
			"error":   nil,
		})
	}
	defer db.Close()

	idParam := c.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "failed converting id",
			"page":    nil,
			"data":    nil,
			"error":   nil,
		})
	}

	const deleteCategoryQuery = `
	DELETE
	FROM
		categories
	WHERE
		id = ?
	`

	_, err = db.ExecContext(c.Request().Context(), deleteCategoryQuery, id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "failed delete category",
			"page":    nil,
			"data":    nil,
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "Successfully delete category",
		"page":    nil,
		"data":    nil,
		"error":   nil,
	})
}