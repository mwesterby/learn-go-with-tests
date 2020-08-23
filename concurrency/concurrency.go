package concurrency

type WebsiteChecker func(string) bool
type result struct {
	string
	bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultsChannel := make(chan result)

	for _, url := range urls {
		go func(u string) {
			// Send statement - sends result to the resultsChannel
			resultsChannel <- result{u, wc(u)}
		}(url) // The () at the end of the anonymous function executes it

	}

	for i := 0; i < len(urls); i++ {
		// Receive expression
		result := <-resultsChannel
		results[result.string] = result.bool
	}

	return results
}
