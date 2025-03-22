package _structs

type Country struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}

type Category struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}

type Movie struct {
	Modified struct {
		Time string `json:"time"`
	} `json:"modified"`
	ID             string     `json:"_id"`
	Name           string     `json:"name"`
	Slug           string     `json:"slug"`
	OriginName     string     `json:"origin_name"`
	Type           string     `json:"type"`
	PosterURL      string     `json:"poster_url"`
	ThumbURL       string     `json:"thumb_url"`
	SubDocquyen    bool       `json:"sub_docquyen"`
	ChieuRap       bool       `json:"chieurap"`
	Time           string     `json:"time"`
	EpisodeCurrent string     `json:"episode_current"`
	Quality        string     `json:"quality"`
	Lang           string     `json:"lang"`
	Year           int        `json:"year"`
	Category       []Category `json:"category"`
	Country        []Country  `json:"country"`
}
