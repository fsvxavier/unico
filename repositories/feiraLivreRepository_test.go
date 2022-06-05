package repositories_test

import (
	"errors"
	"fmt"
	"strings"
	"testing"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/fsvxavier/unico/models"
	"github.com/fsvxavier/unico/repositories"
	"github.com/stretchr/testify/assert"
)

//TestSearchFeiraLivre...
func TestSearchFeiraLivre(t *testing.T) {

	sFeiras := models.SearchFeiraLivre{
		Distrito:  "Teste",
		Regiao5:   "Teste",
		NomeFeira: "Teste",
		Bairro:    "Teste",
	}

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

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	rows := sqlmock.NewRows([]string{"id", "longi", "lat", "setcens", "areap", "coddist", "distrito",
		"codsubpref", "subprefe", "regiao5", "regiao8", "nome_feira", "registro", "logradouro", "numero", "bairro", "referencia"}).
		AddRow(1, 0, 0, 0, 0, 0, "", 0, "", "", "", "", "", "", "", "", "")

	query := fmt.Sprint(`SELECT id, longi, lat, setcens, areap, coddist, distrito, codsubpref, subprefe, regiao5, regiao8, nome_feira,
	registro, logradouro, numero, bairro, referencia
	FROM feira_livre
	WHERE ` + strings.Join(where, " AND "))

	fmt.Println(query)

	mock.ExpectQuery(query).WillReturnRows(rows)
	r := repositories.NewMySQLFeiraLivreRepository(db)
	p, err := r.SearchFeiraLivre(&sFeiras)
	assert.NoError(t, err)
	assert.NotNil(t, p)
}

//TestGetByIDFeiraLivre...
func TestGetByIDFeiraLivre(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	rows := sqlmock.NewRows([]string{"id", "longi", "lat", "setcens", "areap", "coddist", "distrito",
		"codsubpref", "subprefe", "regiao5", "regiao8", "nome_feira", "registro", "logradouro", "numero", "bairro", "referencia"}).
		AddRow(1, 0, 0, 0, 0, 0, "", 0, "", "", "", "", "", "", "", "", "")

	query := `SELECT id, longi, lat, setcens, areap, coddist, distrito, codsubpref, subprefe, regiao5, regiao8, nome_feira,
	registro, logradouro, numero, bairro, referencia
	FROM feira_livre fl
	WHERE fl.id = ?`

	mock.ExpectQuery(query).WillReturnRows(rows)
	r := repositories.NewMySQLFeiraLivreRepository(db)
	id := int64(5)
	p, err := r.GetByID(id)
	assert.NoError(t, err)
	assert.NotNil(t, p)
}

//TestGetAllByIdsFeiraLivre...
func TestGetAllByIdsFeiraLivre(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mockLP := []*models.FeiraLivre{
		{
			ID:         1,
			Longi:      0,
			Lat:        0,
			Setcens:    0,
			Areap:      0,
			Coddist:    0,
			Distrito:   "",
			Codsubpref: 0,
			Subprefe:   "",
			Regiao5:    "",
			Regiao8:    "",
			NomeFeira:  "",
			Registro:   "",
			Logradouro: "",
			Numero:     "",
			Bairro:     "",
			Referencia: "",
		},
		{
			ID:         2,
			Longi:      0,
			Lat:        0,
			Setcens:    0,
			Areap:      0,
			Coddist:    0,
			Distrito:   "",
			Codsubpref: 0,
			Subprefe:   "",
			Regiao5:    "",
			Regiao8:    "",
			NomeFeira:  "",
			Registro:   "",
			Logradouro: "",
			Numero:     "",
			Bairro:     "",
			Referencia: "",
		},
	}

	t.Run("success-one-id", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"id", "longi", "lat", "setcens", "areap", "coddist", "distrito",
			"codsubpref", "subprefe", "regiao5", "regiao8", "nome_feira", "registro", "logradouro", "numero", "bairro", "referencia"}).
			AddRow(mockLP[0].ID, mockLP[0].Longi, mockLP[0].Lat, mockLP[0].Setcens, mockLP[0].Areap, mockLP[0].Coddist, mockLP[0].Distrito, mockLP[0].Codsubpref, mockLP[0].Subprefe,
				mockLP[0].Regiao5, mockLP[0].Regiao8, mockLP[0].NomeFeira, mockLP[0].Registro, mockLP[0].Logradouro, mockLP[0].Numero, mockLP[0].Bairro, mockLP[0].Referencia,
			)

		query := `SELECT id, longi, lat, setcens, areap, coddist, distrito, codsubpref, subprefe, regiao5, regiao8, nome_feira,
		registro, logradouro, numero, bairro, referencia
		from feira_livre 
	   where id = ?`

		mock.ExpectQuery(query).WillReturnRows(rows)
		r := repositories.NewMySQLFeiraLivreRepository(db)

		ids := "1"
		list, err := r.GetAllByIds(ids)
		ret := list
		if len(list) <= 0 {
			ret = nil
		}

		assert.NotEmpty(t, ret)
		assert.NoError(t, err)
		assert.Len(t, ret, 1)
	})
	t.Run("success-two-id", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"id", "longi", "lat", "setcens", "areap", "coddist", "distrito",
			"codsubpref", "subprefe", "regiao5", "regiao8", "nome_feira", "registro", "logradouro", "numero", "bairro", "referencia"}).
			AddRow(mockLP[0].ID, mockLP[0].Longi, mockLP[0].Lat, mockLP[0].Setcens, mockLP[0].Areap, mockLP[0].Coddist, mockLP[0].Distrito, mockLP[0].Codsubpref, mockLP[0].Subprefe,
				mockLP[0].Regiao5, mockLP[0].Regiao8, mockLP[0].NomeFeira, mockLP[0].Registro, mockLP[0].Logradouro, mockLP[0].Numero, mockLP[0].Bairro, mockLP[0].Referencia,
			).
			AddRow(mockLP[1].ID, mockLP[1].Longi, mockLP[1].Lat, mockLP[1].Setcens, mockLP[1].Areap, mockLP[1].Coddist, mockLP[1].Distrito, mockLP[1].Codsubpref, mockLP[1].Subprefe,
				mockLP[1].Regiao5, mockLP[1].Regiao8, mockLP[1].NomeFeira, mockLP[1].Registro, mockLP[1].Logradouro, mockLP[1].Numero, mockLP[1].Bairro, mockLP[1].Referencia,
			)
		query := `SELECT id, longi, lat, setcens, areap, coddist, distrito, codsubpref, subprefe, regiao5, regiao8, nome_feira,
		registro, logradouro, numero, bairro, referencia
		from feira_livre 
	   where id in (?,?)`

		mock.ExpectQuery(query).WillReturnRows(rows)
		r := repositories.NewMySQLFeiraLivreRepository(db)

		ids := "1,2"
		list, err := r.GetAllByIds(ids)
		assert.NotEmpty(t, list)
		assert.NoError(t, err)
		assert.Len(t, list, 2)
	})
	t.Run("error", func(t *testing.T) {

		query := `SELECT id, longi, lat, setcens, areap, coddist, distrito, codsubpref, subprefe, regiao5, regiao8, nome_feira,
		registro, logradouro, numero, bairro, referencia
		from feira_livre 
	   where id = ?`

		mock.ExpectQuery(query).WillReturnError(errors.New(""))
		r := repositories.NewMySQLFeiraLivreRepository(db)

		ids := "1,2"
		_, err := r.GetAllByIds(ids)
		assert.Error(t, err)
		//assert.NotEmpty(t, list)
		//assert.Len(t, list, 1)
	})

	t.Run("error-scan", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"id", "descricao"}).
			AddRow(mockLP[0].ID, mockLP[0].Distrito).
			AddRow(mockLP[1].ID, 1)
		query := `SELECT id, longi, lat, setcens, areap, coddist, distrito, codsubpref, subprefe, regiao5, regiao8, nome_feira,
		registro, logradouro, numero, bairro, referencia
		from feira_livre 
	   where id in (?,?)`

		mock.ExpectQuery(query).WillReturnRows(rows)
		r := repositories.NewMySQLFeiraLivreRepository(db)

		ids := "1,2"
		_, err := r.GetAllByIds(ids)
		assert.Error(t, err)
	})
}

//TestCreateFeiraLivre...
func TestCreateFeiraLivre(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockLPC := models.FeiraLivre{}

		statement := fmt.Sprintf(`insert into feira_livre 
	(	longi, lat, setcens, areap, coddist, distrito, codsubpref, subprefe, regiao5, regiao8, nome_feira, registro,
		logradouro, numero, bairro, referencia)
	values		
	(
		%d, %d, %d, %d, %d, '%s', %d, '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s'
	)`,
			mockLPC.Longi, mockLPC.Lat, mockLPC.Setcens, mockLPC.Areap, mockLPC.Coddist, mockLPC.Distrito, mockLPC.Codsubpref, mockLPC.Subprefe,
			mockLPC.Regiao5, mockLPC.Regiao8, mockLPC.NomeFeira, mockLPC.Registro, mockLPC.Logradouro, mockLPC.Numero, mockLPC.Bairro, mockLPC.Referencia)

		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}
		defer db.Close()

		mock.ExpectExec(statement).WillReturnResult(sqlmock.NewResult(1, 1))
		repository := repositories.NewMySQLFeiraLivreRepository(db)

		_, err = repository.CreateFeiraLivre(&mockLPC)
		assert.NoError(t, err)
	})

}

//TestUpdateFeiraLivre...
func TestUpdateFeiraLivre(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockLPC := models.FeiraLivre{}

		statement := fmt.Sprintf(`UPDATE feira_livre
	SET longi = %d, lat = %d, setcens = %d, areap = %d, coddist = %d, distrito = '%s', codsubpref = %d, 
	subprefe = '%s', regiao5 = '%s', regiao8 = '%s', nome_feira = '%s', registro = '%s',
	logradouro = '%s', numero = '%s', bairro = '%s', referencia = '%s'
	WHERE id = %d`,
			mockLPC.Longi, mockLPC.Lat, mockLPC.Setcens, mockLPC.Areap, mockLPC.Coddist, mockLPC.Distrito, mockLPC.Codsubpref, mockLPC.Subprefe,
			mockLPC.Regiao5, mockLPC.Regiao8, mockLPC.NomeFeira, mockLPC.Registro, mockLPC.Logradouro, mockLPC.Numero, mockLPC.Bairro, mockLPC.Referencia,
			mockLPC.ID)

		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}
		defer db.Close()

		mock.ExpectExec(statement).WillReturnResult(sqlmock.NewResult(1, 1))
		repository := repositories.NewMySQLFeiraLivreRepository(db)

		_, err = repository.UpdateFeiraLivre(&mockLPC)
		assert.NoError(t, err)
	})

}

//TestDeleteFeiraLivre...
func TestDeleteFeiraLivre(t *testing.T) {
	t.Run("success", func(t *testing.T) {

		id := int64(1)

		qry := fmt.Sprintf(`DELETE FROM feira_livre WHERE id = %d`, id)

		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}
		defer db.Close()

		mock.ExpectExec(qry).WillReturnResult(sqlmock.NewResult(1, 1))
		repository := repositories.NewMySQLFeiraLivreRepository(db)

		_, err = repository.DeleteFeiraLivre(id)
		assert.NoError(t, err)
	})

}
