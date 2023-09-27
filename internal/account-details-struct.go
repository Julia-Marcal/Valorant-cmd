package internal

type AccountResponse struct {
	Status int `json:"status"`
	Data   struct {
		Puuid        string `json:"puuid"`
		AccountLevel int    `json:"account_level"`
		Region       string `json:"region"`
		Name         string `json:"name"`
		Tag          string `json:"tag"`
		Card         struct {
			Large string `json:"large"`
		} `json:"card"`
	} `json:"data"`
}

type AccountInfo struct {
	Puuid        string
	AccountLevel int
	Region       string
	Name         string
	Tag          string
	Large        string
}

type AccountFetcher interface {
	FetchAccount(Name string, Tag string) (string, string, int, string, error)
}

type AccountFetcherFactory interface {
	Create() AccountFetcher
}
