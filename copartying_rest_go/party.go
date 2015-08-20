package coparty

type Coparty struct {
    ID   int `json:"id" binding:required`
    Name string `json:"name" binding:required`
    Description string `json:"description"`
}
