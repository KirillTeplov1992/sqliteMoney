package storage

import (
	"sqliteMoney/internal/category/models"
	"sqliteMoney/pkg/sqlite3"
)


type CategoryRepository struct{
	store *sqlite3.Store
}

func NewRepository(store *sqlite3.Store) *CategoryRepository{
	return &CategoryRepository{
		store : store,
	}
}

func (cr *CategoryRepository) GetExpenses() []*models.Category{
	stmt := `
	SELECT
		id,
		name
 	FROM 
		categories
	WHERE
		is_public and  type_of_category = "Расход"`
	
		res, err := cr.store.DB.Query(stmt)
		if err != nil{
		panic(err)
		}

		var categoryList []*models.Category

		for res.Next(){
			category := &models.Category{}
			err = res.Scan(&category.ID,
							&category.Name)

			if err != nil{
				panic(err)
			}

			categoryList = append(categoryList, category)
		}

		return categoryList
}

func (cr *CategoryRepository) GetIncoms() []*models.Category{
	stmt := `
	SELECT
		id,
		name
	FROM
		categories
	WHERE
		is_public and type_of_category = "Доход"`

	res, err := cr.store.DB.Query(stmt)
	if err != nil{
		panic(err)
	}

	var categoryList []*models.Category

	for res.Next(){
		category := &models.Category{}
		err = res.Scan(&category.ID,
					   &category.Name)

		if err != nil{
			panic(err)
		}

		categoryList = append(categoryList, category)
	}

	return categoryList
}

func (cr *CategoryRepository) GetTypeOfCategory (id int) *models.Category {
	stmt := `
	SELECT
		name,
		type_of_category
	FROM
		categories
	WHERE
		id = ?`
	
	row := cr.store.DB.QueryRow(stmt, id)

	category := &models.Category{}
	err := row.Scan(&category.Name,
					&category.Type_of_category)
	if err != nil {
		panic(err)
	}
	
	return category
}

