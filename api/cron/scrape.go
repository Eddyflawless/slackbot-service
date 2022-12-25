package cron

type CurrencyResponse struct {
	Pagination interface{} `json:"pagination"`
	Data       interface{} `json:"data"`
}

func fetchData(endpoint string, options map[string]string) {
	//api_url := fmt.Sprintf("http://api.aviationstack.com/v1/%s", endpoint)

}
