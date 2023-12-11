package entitys

type Table struct {
	TableName string
	Players   []Player
}

type Player struct {
	Name         string
	ActualWinner bool
}
