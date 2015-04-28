package templeFights

import e "github.com/daniloanp/IA/environment"

func TemplesSolution(saints []e.Saint, avTime float64, temples []e.Temple) (bool, []*GameState) {

	if mapToPossibility == nil {
		mapToPossibility.init(saints)
	}
	optimalTime = avTime/float64(len(temples)) + 1
	ctx := &context{
		Saints:  saints,
		Temples: temples,
		MaxTime: avTime,
	}
	//	lifes := make([]int, len(ctx.Saints))

	//	for inx, saint := range ctx.Saints {
	//		lifes[inx] = saint.Lives
	//	}

	//	current := &GameState{
	//		Context:   ctx,
	//		InxTemple: 0,
	//		TimeLeft:  avTime,
	//		Fighters:  0,
	//		Previous:  nil,
	//		Lifes:     lifes,
	//		Neighbors: make([]*GameState, 0, 1<<uint(len(ctx.Saints))),
	//	}
	optimalTime = avTime / float64(len(temples)+1)
	//_ = buildGraph(current, avTime)
	return backtrackedAStar(ctx, avTime)
	//return false, nil
}
