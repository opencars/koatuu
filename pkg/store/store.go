package store

type Store interface {
	Level1() Level1Repository
	Level2() Level2Repository
	Level3() Level3Repository
	Level4() Level4Repository
}
