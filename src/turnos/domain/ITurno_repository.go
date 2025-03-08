package domain


type ITurnoRepository interface {
	CrearTurno(turno Turno) error
	// FinalizarTurno(id int, fin time.Time) error
	// ObtenerTurnoActivo(choferID int) (*Turno, error)
}
