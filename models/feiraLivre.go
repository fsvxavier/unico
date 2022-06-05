package models

type FeiraLivre struct {
	ID         int64  `json:"id"`
	Longi      int64  `json:"longi"`
	Lat        int64  `json:"lat"`
	Setcens    int64  `json:"setcens"`
	Areap      int64  `json:"areap"`
	Coddist    int64  `json:"coddist"`
	Distrito   string `json:"distrito"`
	Codsubpref int64  `json:"codsubpref"`
	Subprefe   string `json:"subprefe"`
	Regiao5    string `json:"regiao5"`
	Regiao8    string `json:"regiao8"`
	NomeFeira  string `json:"nome_feira"`
	Registro   string `json:"registro"`
	Logradouro string `json:"logradouro"`
	Numero     string `json:"numero"`
	Bairro     string `json:"bairro"`
	Referencia string `json:"referencia"`
}
