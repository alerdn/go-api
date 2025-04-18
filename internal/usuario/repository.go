package usuario

import (
	"database/sql"

	"github.com/alerdn/go-api/config"
)

func Criar(usuario Usuario) (Usuario, error) {
	stmt, err := config.DB.Prepare("INSERT INTO usuarios (nome, email, senha) VALUES (?, ?, ?)")
	if err != nil {
		return Usuario{}, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(usuario.Nome, usuario.Email, usuario.Senha)
	if err != nil {
		return Usuario{}, err
	}

	id, _ := result.LastInsertId()
	usuario.ID = int(id)
	return usuario, nil
}

func Listar() ([]Usuario, error) {
	rows, err := config.DB.Query("SELECT id, nome, email FROM usuarios")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var usuarios []Usuario

	for rows.Next() {
		var u Usuario
		if err := rows.Scan(&u.ID, &u.Nome, &u.Email); err != nil {
			return nil, err
		}
		usuarios = append(usuarios, u)
	}

	return usuarios, nil
}

func BuscarPorID(id int) (*Usuario, error) {
	stmt, err := config.DB.Prepare("SELECT id, nome, email, senha FROM usuarios WHERE id = ? LIMIT 1")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var u Usuario
	err = stmt.QueryRow(id).Scan(&u.ID, &u.Nome, &u.Email, &u.Senha)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return &u, nil
}

func BuscarPorEmail(email string) (*Usuario, error) {
	stmt, err := config.DB.Prepare("SELECT id, nome, email, senha FROM usuarios WHERE email = ? LIMIT 1")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var u Usuario
	err = stmt.QueryRow(email).Scan(&u.ID, &u.Nome, &u.Email, &u.Senha)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return &u, nil
}
