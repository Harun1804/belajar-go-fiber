package dtos

type BookResponse struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Author    string `json:"author"`
	Year      int    `json:"year"`
	Publisher string `json:"publisher"`
	Cover     string `json:"cover,omitempty"`
	CoverURL  string `json:"coverUrl,omitempty"`
}