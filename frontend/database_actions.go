package main

import "database/sql"

type Project struct {
	Title   string `json:"Title"`
	Desc    string `json:"Desc"`
	Tech    string `json:"Tech"`
	GitLink string `json:"GitLink"`
	WebLink string `json:"WebLink"`
	BlogId  int    `json:"BlogId"`
}

type Blog struct {
	BlogId int    `json:"BlogId"`
	Title  string `json:"Title"`
}

func GetProjects(db *sql.DB) ([]Project, error) {
	var projects []Project

	sqlQuery := `SELECT Title,Desc,Tech,GitLink,WebLink,Blog_id FROM Projects`
	rows, queryErr := db.Query(sqlQuery)
	if queryErr != nil {
		return nil, queryErr
	}
	defer rows.Close()

	for rows.Next() {
		p := &Project{}
		rowErr := rows.Scan(&p.Title, &p.Desc, &p.GitLink, &p.WebLink, &p.BlogId)
		if rowErr != nil {
			return nil, rowErr
		}
		projects = append(projects, *p)
	}
	return projects, nil
}

func GetBlogs(db *sql.DB) ([]Blog, error) {
	var blogs []Blog

	sqlQuery := `SELECT Blog_id,Title FROM Blogs`
	rows, queryErr := db.Query(sqlQuery)
	if queryErr != nil {
		return nil, queryErr
	}

	for rows.Next() {
		b := &Blog{}
		rowErr := rows.Scan(&b.BlogId, &b.Title)
		if rowErr != nil {
			return nil, rowErr
		}
		blogs = append(blogs, *b)
	}
	return blogs, nil
}
