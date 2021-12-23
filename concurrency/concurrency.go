package concurrency

type WebsiteChecker func(string) bool
type result struct {
	Url string
	Contains bool
} 

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
    results := make(map[string]bool)
	// resChannel := make(chan result)

    // for _, url := range urls {
	// 	go func(url_ string) {
	// 		resChannel <- result{url_, wc(url_)}
	// 	}(url)
    // }
	for i := 0; i < len(urls); i++ {
		// res := <- resChannel
		url_ := urls[i]
		results[url_] = wc(url_)
    }

    return results
}