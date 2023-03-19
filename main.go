package main

type Story struct {
	Title string `json:"title"`
	URL   string `json:"url"`
}

func main() {
	// Declare a variable of type Story
	var story Story = Story{
		Title: "Hello World",
		URL:   "https://github.com",
	}

	println(story.Title)
	println(story.URL)

}
