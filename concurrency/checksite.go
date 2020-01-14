package concurrency

//WebsiteChecker func type
type WebsiteChecker func(string) bool

type result struct {
	string
	bool
}

//CheckWebsites func
func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan result)

	for _, url := range urls {
		go func(u string) {
			resultChannel <- result{u, wc(u)}
		}(url)
	}
	for i := 0; i < len(urls); i++ {
		result := <-resultChannel
		results[result.string] = result.bool
	}
	// Linear way
	// for _, url := range urls {
	// 	results[url] = wc(url)
	// }

	return results
}
