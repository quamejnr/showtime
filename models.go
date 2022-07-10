package main

type season struct {
	Season   int      `json:"season"`
	Episodes []string `json:"episodes"`	
}

type series struct {
	ID		int 	 `json:"id"`
	Title   string   `json:"title"`
	Genre   string   `json:"genre"`
	Seasons []season `json:"seasons"`
}

var tvShows = []series{
	{ID: 1, Title: "Ted Lasso", Genre: "Comedy", Seasons: []season{
		{Season: 1, Episodes: []string{"s01e01", "s01e02", "s01e03", "s01e04", "s01e05"}},
		{Season: 2, Episodes: []string{"s02e01", "s02e02", "s02e03", "s02e04", "s02e05"}},
	}},
	{ID: 2, Title: "Game of Thrones", Genre: "Action", Seasons: []season{
		{Season: 1, Episodes: []string{"s01e01", "s01e02", "s01e03", "s01e04", "s01e05"}},
		{Season: 2, Episodes: []string{"s02e01", "s02e02", "s02e03", "s02e04", "s02e05"}},
	}},
}
