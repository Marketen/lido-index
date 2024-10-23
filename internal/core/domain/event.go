package domain

type Event struct {
    ID       string `json:"id"`
    Type     string `json:"type"`
    Data     string `json:"data"`
}
