package internal

type Character struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type DataStruct struct {
	Meta struct {
		ID  string `json:"id"`
		Map struct {
			ID   string `json:"id"`
			Name string `json:"name"`
		} `json:"map"`
		Mode string `json:"mode"`
	} `json:"meta"`
	Stats struct {
		Kills     int       `json:"kills"`
		Deaths    int       `json:"deaths"`
		Assists   int       `json:"assists"`
		Team      string    `json:"team"`
		Character Character `json:"character"`
	} `json:"stats"`
	Teams struct {
		Red  int `json:"red"`
		Blue int `json:"blue"`
	} `json:"teams"`
}

type DataMatchesFetcher interface {
	FetchMatchData(region string, puuid string) ([]int, map[string]string, []float64, string, error)
}

type DataMatchesFetcherFactory interface {
	Create() DataMatchesFetcher
}
