package infraestructure

import (
	"bus-driver/src/turnos/domain"
	"database/sql"
	"errors"
	"fmt"
)

type MySQL struct {
	db *sql.DB
}

func NewMySQL(db *sql.DB) *MySQL {
	return &MySQL{db: db}
}

func (mysql *MySQL) CrearTurno(turno domain.Turno) error {
	var existe int
	err := mysql.db.QueryRow("SELECT COUNT(*) FROM turnos WHERE choferID = ? AND estado = 'activo'", turno.ChoferID).Scan(&existe)
	if err != nil {
		return fmt.Errorf("error al verificar turno activo: %w", err)
	}
	if existe > 0 {
		return errors.New("el chofer ya tiene un turno activo")
	}

	_, err = mysql.db.Exec("INSERT INTO turnos (choferID, busID, inicio, estado) VALUES (?, ?, ?, ?)",
		turno.ChoferID, turno.BusID, turno.Inicio, "activo")
	if err != nil {
		return fmt.Errorf("error al crear el turno: %w", err)
	}
	return nil
}

// func (mysql *MySQL) FinalizarTurno(id int, fin time.Time) error {
// 	// Verificar si el turno existe
// 	var estado string
// 	err := mysql.db.QueryRow("SELECT estado FROM turnos WHERE id = ?", id).Scan(&estado)
// 	if err != nil {
// 		if err == sql.ErrNoRows {
// 			return errors.New("el turno no existe")
// 		}
// 		return fmt.Errorf("error al obtener el turno: %w", err)
// 	}
// 	if estado == "finalizado" {
// 		return errors.New("el turno ya ha sido finalizado")
// 	}

// 	_, err = mysql.db.Exec("UPDATE turnos SET estado = 'finalizado', fin = ? WHERE id = ?",
// 		fin, id)
// 	if err != nil {
// 		return fmt.Errorf("error al finalizar el turno: %w", err)
// 	}
// 	return nil
// }

// func (mysql *MySQL ) ObtenerTurnoActivo(choferID int) (*domain.Turno, error) {
// 	var turno domain.Turno
// 	err := mysql.db.QueryRow("SELECT id, choferID, busID, inicio, fin, estado FROM turnos WHERE choferID = ? AND estado = 'activo'", choferID).
// 		Scan(&turno.ID, &turno.ChoferID, &turno.BusID, &turno.Inicio, &turno.Fin, &turno.Estado)
// 	if err != nil {
// 		if err == sql.ErrNoRows {
// 			return nil, errors.New("no se encontró un turno activo para el chofer")
// 		}
// 		return nil, fmt.Errorf("error al obtener el turno activo: %w", err)
// 	}
// 	return &turno, nil
// }
