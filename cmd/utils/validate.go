package utils

// validate the url

func ValidateUrl(url string) bool {
	if url == "" {
		return false
	}
	if url[:7] == "http://" || url[:8] == "https://" {
		return true
	}
	return false
}

// InvalidUrlError
type InvalidUrlError struct {
	Url string
}

func (e *InvalidUrlError) Error() string {
	return "Invalid Url: " + e.Url
}

// to call this you can
