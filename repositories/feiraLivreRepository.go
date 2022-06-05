package repositories

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/fsvxavier/unico/interfaces"
	"github.com/fsvxavier/unico/models"
	"github.com/fsvxavier/unico/utils"
)

const (
	defaultCacheExpireSeconds = 3600
	defaultCacheSizeMB        = 5
)

var (
	logger = new(utils.GenericLogger)
)

type mysqlFeiraLivreRepository struct {
	Conn *sql.DB
}

// NewMySQLFeiraLivreRepository exports an interface to arquivoRepository
func NewMySQLFeiraLivreRepository(Conn *sql.DB) interfaces.FeiraLivreRepository {
	return &mysqlFeiraLivreRepository{Conn: Conn}
}

func (r *mysqlFeiraLivreRepository) GetByID(id int64) ([]*models.FeiraLivre, error) {

	rows, err := r.Conn.Query(`SELECT id, longi, lat, setcens, areap, coddist, distrito, codsubpref, subprefe, regiao5, regiao8, nome_feira,
	registro, logradouro, numero, bairro, referencia
	FROM feira_livre fl
	WHERE fl.id = ?`, &id)
	if err != nil {
		return nil, err
	}
	list := make([]*models.FeiraLivre, 0)

	defer rows.Close()

	for rows.Next() {

		pa := new(models.FeiraLivre)
		var (
			id         sql.NullInt64
			longi      sql.NullInt64
			lat        sql.NullInt64
			setcens    sql.NullInt64
			areap      sql.NullInt64
			coddist    sql.NullInt64
			distrito   sql.NullString
			codsubpref sql.NullInt64
			subprefe   sql.NullString
			regiao5    sql.NullString
			regiao8    sql.NullString
			nomeFeira  sql.NullString
			registro   sql.NullString
			logradouro sql.NullString
			numero     sql.NullString
			bairro     sql.NullString
			referencia sql.NullString
		)

		err = rows.Scan(&id, &longi, &lat, &setcens, &areap, &coddist, &distrito, &codsubpref, &subprefe, &regiao5, &regiao8, &nomeFeira, &registro, &logradouro,
			&numero, &bairro, &referencia)
		if err != nil {
			return nil, err
		}

		pa.ID = id.Int64
		pa.Longi = longi.Int64
		pa.Lat = lat.Int64
		pa.Setcens = setcens.Int64
		pa.Areap = areap.Int64
		pa.Coddist = coddist.Int64
		pa.Distrito = distrito.String
		pa.Codsubpref = codsubpref.Int64
		pa.Subprefe = subprefe.String
		pa.Regiao5 = regiao5.String
		pa.Regiao8 = regiao8.String
		pa.NomeFeira = nomeFeira.String
		pa.Registro = registro.String
		pa.Logradouro = logradouro.String
		pa.Numero = numero.String
		pa.Bairro = bairro.String
		pa.Referencia = referencia.String

		list = append(list, pa)
	}
	return list, nil
}

//SearchFeiraLivre...
func (r *mysqlFeiraLivreRepository) SearchFeiraLivre(sFeiras *models.SearchFeiraLivre) ([]*models.FeiraLivre, error) {

	var where []string

	if sFeiras.Distrito != "" {
		where = append(where, "distrito LIKE '%"+sFeiras.Distrito+"%'")
	}

	if sFeiras.Regiao5 != "" {
		where = append(where, "regiao5 LIKE '%"+sFeiras.Regiao5+"%'")
	}

	if sFeiras.NomeFeira != "" {
		where = append(where, "nome_feira LIKE '%"+sFeiras.NomeFeira+"%'")
	}

	if sFeiras.Bairro != "" {
		where = append(where, "bairro LIKE '%"+sFeiras.Bairro+"%'")
	}

	rows, err := r.Conn.Query(`SELECT id, longi, lat, setcens, areap, coddist, distrito, codsubpref, subprefe, regiao5, regiao8, nome_feira,
		registro, logradouro, numero, bairro, referencia
		FROM feira_livre
		WHERE ` + strings.Join(where, " AND "))

	if err != nil {
		return nil, err
	}
	list := make([]*models.FeiraLivre, 0)

	defer rows.Close()

	for rows.Next() {

		pa := new(models.FeiraLivre)
		var (
			id         sql.NullInt64
			longi      sql.NullInt64
			lat        sql.NullInt64
			setcens    sql.NullInt64
			areap      sql.NullInt64
			coddist    sql.NullInt64
			distrito   sql.NullString
			codsubpref sql.NullInt64
			subprefe   sql.NullString
			regiao5    sql.NullString
			regiao8    sql.NullString
			nomeFeira  sql.NullString
			registro   sql.NullString
			logradouro sql.NullString
			numero     sql.NullString
			bairro     sql.NullString
			referencia sql.NullString
		)

		err = rows.Scan(&id, &longi, &lat, &setcens, &areap, &coddist, &distrito, &codsubpref, &subprefe, &regiao5, &regiao8, &nomeFeira, &registro, &logradouro,
			&numero, &bairro, &referencia)
		if err != nil {
			return nil, err
		}

		pa.ID = id.Int64
		pa.Longi = longi.Int64
		pa.Lat = lat.Int64
		pa.Setcens = setcens.Int64
		pa.Areap = areap.Int64
		pa.Coddist = coddist.Int64
		pa.Distrito = distrito.String
		pa.Codsubpref = codsubpref.Int64
		pa.Subprefe = subprefe.String
		pa.Regiao5 = regiao5.String
		pa.Regiao8 = regiao8.String
		pa.NomeFeira = nomeFeira.String
		pa.Registro = registro.String
		pa.Logradouro = logradouro.String
		pa.Numero = numero.String
		pa.Bairro = bairro.String
		pa.Referencia = referencia.String

		list = append(list, pa)
	}
	return list, nil
}

func (r *mysqlFeiraLivreRepository) GetAllByIds(ids string) ([]*models.FeiraLivre, error) {
	var values []interface{}
	var where []string

	sliceIds := strings.Split(ids, ",")
	if len(sliceIds) > 1 {
		where = append(where, "in (?"+strings.Repeat(",?", len(sliceIds)-1)+")")
		for _, status := range sliceIds {
			values = append(values, status)
		}
	} else if sliceIds[0] != "" {
		where = append(where, "= ?")
		values = append(values, sliceIds[0])
	}

	rows, err := r.Conn.Query(`SELECT id, longi, lat, setcens, areap, coddist, distrito, codsubpref, subprefe, regiao5, regiao8, nome_feira,
	registro, logradouro, numero, bairro, referencia
	from feira_livre 
   where id `+strings.Join(where, "  AND "), values...)

	if err != nil {
		return nil, err
	}

	list := make([]*models.FeiraLivre, 0)

	defer rows.Close()

	for rows.Next() {

		pa := new(models.FeiraLivre)
		var (
			id         sql.NullInt64
			longi      sql.NullInt64
			lat        sql.NullInt64
			setcens    sql.NullInt64
			areap      sql.NullInt64
			coddist    sql.NullInt64
			distrito   sql.NullString
			codsubpref sql.NullInt64
			subprefe   sql.NullString
			regiao5    sql.NullString
			regiao8    sql.NullString
			nomeFeira  sql.NullString
			registro   sql.NullString
			logradouro sql.NullString
			numero     sql.NullString
			bairro     sql.NullString
			referencia sql.NullString
		)

		err = rows.Scan(&id, &longi, &lat, &setcens, &areap, &coddist, &distrito, &codsubpref, &subprefe, &regiao5, &regiao8, &nomeFeira, &registro, &logradouro,
			&numero, &bairro, &referencia)
		if err != nil {
			return nil, err
		}

		pa.ID = id.Int64
		pa.Longi = longi.Int64
		pa.Lat = lat.Int64
		pa.Setcens = setcens.Int64
		pa.Areap = areap.Int64
		pa.Coddist = coddist.Int64
		pa.Distrito = distrito.String
		pa.Codsubpref = codsubpref.Int64
		pa.Subprefe = subprefe.String
		pa.Regiao5 = regiao5.String
		pa.Regiao8 = regiao8.String
		pa.NomeFeira = nomeFeira.String
		pa.Registro = registro.String
		pa.Logradouro = logradouro.String
		pa.Numero = numero.String
		pa.Bairro = bairro.String
		pa.Referencia = referencia.String

		list = append(list, pa)
	}
	return list, nil
}

// CreateFeiraLivre ...
func (r *mysqlFeiraLivreRepository) CreateFeiraLivre(iFeiras *models.FeiraLivre) (*models.Response, error) {

	statement := fmt.Sprintf(`insert into feira_livre 
	(	longi, lat, setcens, areap, coddist, distrito, codsubpref, subprefe, regiao5, regiao8, nome_feira, registro,
		logradouro, numero, bairro, referencia)
	values		
	(
		%d, %d, %d, %d, %d, '%s', %d, '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s'
	)`,
		iFeiras.Longi, iFeiras.Lat, iFeiras.Setcens, iFeiras.Areap, iFeiras.Coddist, iFeiras.Distrito, iFeiras.Codsubpref,
		iFeiras.Subprefe, iFeiras.Regiao5, iFeiras.Regiao8, iFeiras.NomeFeira, iFeiras.Registro, iFeiras.Logradouro, iFeiras.Numero,
		iFeiras.Bairro, iFeiras.Referencia)

	//fmt.Println(query)

	res, err := r.Conn.Exec(statement)
	if err != nil {
		return nil, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}
	iFeiras.ID = id

	ret := models.Response{
		Status:  "Success",
		Mensage: fmt.Sprintf("Registro inserido com sucesso com ID: %d", id),
	}

	return &ret, nil

}

// UpdateFeiraLivre ...
func (r *mysqlFeiraLivreRepository) UpdateFeiraLivre(iFeiras *models.FeiraLivre) (*models.Response, error) {

	idFeira := iFeiras.ID

	statement := fmt.Sprintf(`UPDATE feira_livre
	SET longi = %d, lat = %d, setcens = %d, areap = %d, coddist = %d, distrito = '%s', codsubpref = %d, 
	subprefe = '%s', regiao5 = '%s', regiao8 = '%s', nome_feira = '%s', registro = '%s',
	logradouro = '%s', numero = '%s', bairro = '%s', referencia = '%s'
	WHERE id = %d`,
		iFeiras.Longi, iFeiras.Lat, iFeiras.Setcens, iFeiras.Areap, iFeiras.Coddist, iFeiras.Distrito, iFeiras.Codsubpref,
		iFeiras.Subprefe, iFeiras.Regiao5, iFeiras.Regiao8, iFeiras.NomeFeira, iFeiras.Registro, iFeiras.Logradouro, iFeiras.Numero,
		iFeiras.Bairro, iFeiras.Referencia, idFeira)

	_, err := r.Conn.Exec(statement)
	if err != nil {
		return nil, err
	}

	ret := models.Response{
		Status:  "Success",
		Mensage: fmt.Sprintf("Registro atualizado com sucesso pelo ID: %d", idFeira),
	}

	return &ret, nil

}

// UpdateFeiraLivre ...
func (r *mysqlFeiraLivreRepository) DeleteFeiraLivre(id int64) (*models.Response, error) {

	qry := fmt.Sprintf(`DELETE FROM feira_livre WHERE id = %d`, id)

	ret := models.Response{
		Status:  "Success",
		Mensage: fmt.Sprintf("Registro apagado com sucesso pelo ID: %d", id),
	}

	_, err := r.Conn.Exec(qry)
	if err != nil {

		ret.Status = "Error"
		ret.Mensage = fmt.Sprintf("Erro ao tentar apagar o registro de ID: %d", id)

		return &ret, err
	}

	return &ret, nil

}
