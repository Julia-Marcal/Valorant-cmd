package internal

// Account data fetching

type ConcreteAccountFetcherFactory struct{}

func (f *ConcreteAccountFetcherFactory) Create() AccountFetcher {
	return &ConcreteAccountFetcher{}
}

type ConcreteAccountFetcher struct{}

func (c *ConcreteAccountFetcher) FetchAccount(Name string, Tag string) (string, string, int, string, error) {
	accountInfo, err := AccountInformation(Name, Tag)
	if err != nil {
		return "", "", 0, "", err
	}
	return accountInfo.Region, accountInfo.Puuid, accountInfo.AccountLevel, accountInfo.Large, nil
}

// Match data fetching
type ConcreteMatchesFetcherFactory struct{}

func (f *ConcreteMatchesFetcherFactory) Create() DataMatchesFetcher {
	return &ConcreteMatchDataFetcher{}
}

type ConcreteMatchDataFetcher struct{}

func (c *ConcreteMatchDataFetcher) FetchMatchData(region string, puuid string) ([]int, map[string]string, []float64, string, error) {
	intInfo, mapInfo, averages, bestCharacter, err := FetchMatches(region, puuid)
	if err != nil {
		return nil, nil, nil, "", err
	}
	return intInfo, mapInfo, averages, bestCharacter, nil
}
