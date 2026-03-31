package storage

import (
	"sqliteMoney/internal/reports/models"
	"sqliteMoney/pkg/sqlite3"
)


type ReportRepository struct{
	store *sqlite3.Store
}

func NewRepository(store *sqlite3.Store) *ReportRepository{
	return &ReportRepository{
		store : store,
	}
}

func (rr *ReportRepository) Report(beginDate, endDate string) []*models.Report{
	stmt := `
		with I as (
			SELECT
				strftime('%Y-%m', date) AS year_month,
				sum(amount) as incom
			FROM
				transactions T
			inner JOIN
				categories C
				on
					T.category_id = C.id
			WHERE
				date >= ? and date <= ? and C.type_of_category = "Доход" and C.is_public = True
			group by
				year_month),
		E as (
			SELECT
				strftime('%Y-%m', date) AS year_month,
				sum(amount) as expense
			FROM
				transactions T
			inner JOIN
				categories C
				on
					T.category_id = C.id
			WHERE
				date >= ? and date <= ? and C.type_of_category = "Расход" and C.is_public = True
			group by
				year_month)
	
		select 
			I.year_month,
			I.incom,
			E.expense,
			I.incom - E.expense AS profit
		from
			I
		INNER JOIN
			E on I.year_month = E.year_month`
	
	res, err :=rr.store.DB.Query(stmt,
		beginDate,
		endDate,
		beginDate,
		endDate)
	if err != nil{
		panic(err)
	}

	var reps []*models.Report

	for res.Next(){
		month := &models.Report{}

		err = res.Scan(&month.Month,
			&month.Incom,
			&month.Expense,
			&month.Profit)

		if err != nil{
			panic(err)
		}

		reps = append(reps, month)
	}

	return reps	
}

func (rr *ReportRepository) GetFinalReport(beginDate, endDate string) []*models.FinalReport{
	stmt := `
	WITH SI AS (
		SELECT
			SUM(amount) as sum_incom
		FROM
			transactions AS t
		INNER JOIN
			categories AS C
			ON
				t.category_id = c.id
		WHERE
			date >= ?
				AND
			date <= ?
				AND
			type_of_category == "Доход"
				AND
			is_public = true),

	SE AS (
		SELECT
			SUM(amount) as sum_expense
		FROM
			transactions AS t
		INNER JOIN
			categories AS C
			ON
				t.category_id = c.id
		WHERE
			date >= ?
				AND
			date <= ?
				AND
			type_of_category == "Расход"
				AND
			is_public = true),
	
	total_table as (
		SELECT
			1 as num,
			"Суммарный доход за период:" as type,
			SI.sum_incom
		FROM
			SI
		UNION
		SELECT
			2 as num,
			"Суммарный расход за период:" as type,
			SE.sum_expense
		FROM
			SE
		UNION
		SELECT
			3 as num,
			"Прибыль за период:" as type,
			round(SI.sum_incom - SE.sum_expense, 2)
		FROM
			SI, SE)
	
		select 
			tb.type
			,tb.sum_incom AS amount
			,round((tb.sum_incom / SI.sum_incom)*100, 2) as persent	
		from total_table as tb, SI`

	res, err :=rr.store.DB.Query(stmt,
		beginDate,
		endDate,
		beginDate,
		endDate)
	if err != nil{
		panic(err)
	}

	var freps []*models.FinalReport

	for res.Next(){
		frep := &models.FinalReport{}

		err = res.Scan(&frep.Type,
					&frep.Amount,
					&frep.Persent)

		if err != nil{
			panic(err)
		}

		freps = append(freps, frep)		
	}
	return freps
}

func (rr *ReportRepository) GetIncoms(beginDate, endDate string) []*models.CategoryReport{
	stmt := `
	SELECT
		c.name
		,SUM(amount)
	FROM
		transactions as t
	INNER JOIN
		categories as c
		on
			t.category_id = c.id
	WHERE
		date >= ?
			AND
		date <= ?
			AND
		type_of_category = "Доход"
			AND
		is_public=true
	GROUP BY
		C.name
	ORDER BY
		SUM(amount) DESC`

	res, err :=rr.store.DB.Query(stmt,
		beginDate,
		endDate)
	if err != nil{
		panic(err)
	}

	var creps []*models.CategoryReport

	for res.Next(){
		crep := &models.CategoryReport{}

		err = res.Scan(&crep.Type,
						&crep.Amount)

		if err != nil{
			panic(err)
		}
		
		creps = append(creps, crep)
	}

	return creps
}

func (r *ReportRepository) GetExpenses(beginDate, endDate string) []*models.CategoryReport{
	stmt := `
	SELECT
		c.name
		,SUM(amount)
	FROM
		transactions as t
	INNER JOIN
		categories as c
		on
			t.category_id = c.id
	WHERE
		date >= ?
			AND
		date <= ?
			AND
		type_of_category = "Расход"
			AND
		is_public=true
	GROUP BY
		C.name
	ORDER BY
		SUM(amount) DESC`

	res, err :=r.store.DB.Query(stmt,
		beginDate,
		endDate)
	if err != nil{
		panic(err)
	}

	var creps []*models.CategoryReport

	for res.Next(){
		crep := &models.CategoryReport{}

		err = res.Scan(&crep.Type,
						&crep.Amount)

		if err != nil{
			panic(err)
		}
		
		creps = append(creps, crep)
	}

	return creps
}