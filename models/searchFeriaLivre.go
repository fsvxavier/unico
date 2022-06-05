package models

type SearchFeiraLivre struct {
	Bairro    string `json:"bairro"`
	Distrito  string `json:"distrito"`
	NomeFeira string `json:"nome_feira"`
	Regiao5   string `json:"regiao5"`
}
