package app

import (
	"context"
	"log"
	"wblitz-watcher/wg/types"
)

func (app *App) TestEvent() {
	testClan := types.ClanInfo{
		Name: "Test",
		Tag:  "TEST",
	}

	testPlayer := types.AccountInfo{
		AccountID: 1,
		Nickname:  "Player_1488",
	}

	testMSG := app.buildClanEvent(&testClan, &testPlayer, DiffLeave)

	err := app.SendDiscordMessage(context.Background(), testMSG)
	if err != nil {
		log.Printf("%+v\n", err)
	}
}
