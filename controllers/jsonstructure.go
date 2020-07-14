package controllers

type DecodeData struct {
	End_Year   int    `json:"end_year"`
	Intensity  int    `json:"intensity"`
	Sector     string `json:"sector"`
	Topic      string `json:"topic"`
	Insight    string `json:"insight"`
	Url        string `json:"url"`
	Region     string `json:"region"`
	Start_Year int    `json:"start_year"`
	Impact     int    `json:"impact"`
	Added      string `json:"added"`
	Published  string `json:"published"`
	Country    string `json:"country"`
	Relevance  int    `json:"relevance"`
	Pestle     string `json:"pestle"`
	Source     string `json:"source"`
	Title      string `json:"title"`
	Likelihood int    `json:"likelihood"`
}
