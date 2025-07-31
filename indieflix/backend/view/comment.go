package view

type CommentResponse struct {
	ID        int    `json:"id"`
	UserId    int    `json:"user_id"`
	MovieId   int    `json:"movie_id"`
	UserText  string `json:"user_text"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type CommentResponseList struct {
	Page    int               `json:"page"`
	Size    int               `json:"size"`
	Results []CommentResponse `json:"results"`
}
