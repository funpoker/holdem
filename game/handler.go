package game

import (
	"strings"

	"code.byted.org/gopkg/logs"
	"github.com/gogo/protobuf/proto"
	"github.com/sirupsen/logrus"

	"github.com/funpoker/holdem/models"
	pb "github.com/funpoker/holdem/proto"
)

type Context struct {
	Header    *pb.Header
	Body      []byte
	Send      chan []byte
	Broadcast chan []byte
}

type Handler func(c *Context)

var routers = map[pb.MessageType]Handler{
	pb.MessageType_JOIN_GAME_REQUEST:  JoinGame,
	pb.MessageType_JOIN_GAME_RESPONSE: nil,
	pb.MessageType_PLAYER_INFO_NOTIFY: nil,
	pb.MessageType_START_GAME_REQUEST: StartGame,
	pb.MessageType_SEND_CARD:          SendCard,
	pb.MessageType_BET_REQUEST:        nil,
	pb.MessageType_BET_NOTIFY:         nil,
	pb.MessageType_RESULT_NOTIFY:      nil,
	pb.MessageType_EXIT_GAME_REQUEST:  nil,
	pb.MessageType_EXIT_GAME_NOTIFY:   nil,
	pb.MessageType_ERROR:              nil,
}

func Handle(c *Context) {
	handler, ok := routers[c.Header.Type]
	if !ok {
		ErrorHandler(c)
		return
	}
	handler(c)
}

func ErrorHandler(c *Context) {
	logrus.Errorf("Unknown message type")
	msg := &pb.Error{Code: 1}
	bytes, err := proto.Marshal(msg)
	if err != nil {
		logrus.Errorf("Marshal %q err=%v", msg, err)
		return
	}
	c.Send <- bytes
}

func JoinGame(c *Context) {
	var req pb.JoinGameRequest
	if err := proto.Unmarshal(c.Body, &req); err != nil {
		logrus.Errorf("Unmarshal request err: %v", err)
		return
	}
	roomId := c.Header.RoomId
	playerId := req.PlayerId
	playerRole := req.Role

	// find gameId by roomId
	game, err := models.GetGame(uint64(roomId))
	if err != nil {
		logrus.Errorf("Get game by roomId %q err: %v", roomId, err)
		return
	}

	// create game record
	record := &models.RelationGamePlayer{
		GameId:     game.ID,
		PlayerId:   uint(playerId),
		PlayerRole: int(playerRole),
	}
	err = models.CreateRelationGamePlayer(record)
	if err != nil {
		logrus.Errorf("Create game record %v err: %v", record, err)
		return
	}

	players, err := models.ListGamePlayers(game.ID)
	if err != nil {
		logrus.Errorf("List game player err: %v", err)
		return
	}
	// broadcast all players info
	resp := &pb.JoinGameResponse{
		Players: make([]*pb.PlayerInfoNotify, len(players)),
	}
	for i, p := range players {
		resp.Players[i] = &pb.PlayerInfoNotify{
			Player: &pb.Player{
				Id:        int32(p.Id),
				Username:  p.Username,
				AvatarUrl: p.AvatarUrl,
				Amount:    int32(p.Amount),
			},
			Role:     pb.PlayerRole(p.PlayerRole),
			GameRole: pb.GameRole_NORMAL, // p.GameRole,
			Position: int32(p.PlayerPosition),
		}
	}
	respByte, err := proto.Marshal(resp)
	if err != nil {
		logrus.Errorf("Marshal %+v err: %v", resp, err)
		return
	}
	c.Broadcast <- respByte
}

func StartGame(c *Context) {
	var req pb.StartGameRequest
	if err := proto.Unmarshal(c.Body, &req); err != nil {
		logrus.Errorf("Unmarshal request err: %v", err)
		return
	}
	roomId := c.Header.RoomId
	playerId := req.PlayerId

	chMap.Add(int(playerId), c.Send)

	// find gameId by roomId
	game, err := models.GetGame(uint64(roomId))
	if err != nil {
		logrus.Errorf("Get game by roomId %q err: %v", roomId, err)
		return
	}

	// set the player status to start
	updates := map[string]interface{}{
		"game_id":   game.ID,
		"player_id": playerId,
		"status":    models.GameStatusStart,
	}
	if err = models.UpdateRelationGamePlayer(updates); err != nil {
		logrus.Errorf("update relation game player err: %v", err)
		return
	}

	players, err := models.ListGamePlayers(game.ID)
	if err != nil {
		logrus.Errorf("List game player err: %v", err)
		return
	}

	playerNum := 0
	for _, p := range players {
		if pb.PlayerRole(p.PlayerRole) == pb.PlayerRole_ACTOR && p.Status == models.GameStatusStart {
			playerNum++
		}
	}

	logrus.Infof("Room %q started player num:%d", roomId, playerNum)
	if playerNum < 5 {
		return
	}

	hp, err := NewHoldemPoker(playerNum)
	if err != nil {
		logrus.Errorf("Create holdem poker err: %v", err)
		return
	}

	// update db game
	updates = map[string]interface{}{
		"id":    game.ID,
		"flop":  hp.Flop.String(),
		"turn":  hp.Turn,
		"river": hp.River,
	}
	if err = models.UpdateGame(updates); err != nil {
		logrus.Errorf("update game err: %v", err)
		return
	}

	// update db relation_game_player
	i := 0
	for _, p := range players {
		if pb.PlayerRole(p.PlayerRole) == pb.PlayerRole_ACTOR {
			hands := hp.PlayerHands(i)
			updates = map[string]interface{}{
				"game_id":   game.ID,
				"player_id": p.Id,
				"hands":     hands,
			}
			if err = models.UpdateRelationGamePlayer(updates); err != nil {
				logrus.Errorf("update relation game player err: %v", err)
				return
			}
			// send hands to player
			ch, ok := chMap.Get(p.Id)
			if !ok {
				logs.Errorf("Find player %q 's send chan fail", p.Id)
			} else {
				resp := &pb.SendCardNotify{
					CardType: pb.CardType_PLAYER,
					Num:      2,
					Cards:    strings.Split(hands, " "),
				}
				respByte, err := proto.Marshal(resp)
				if err != nil {
					logrus.Errorf("Marshal %+v err: %v", resp, err)
					return
				}
				ch <- respByte
			}

			i++
		}
	}
}

func SendCard(c *Context) {
	msg := &pb.SendCardNotify{}
	msg.Cards = []string{"As", "Kc", "Qh"}
	msg.CardType = pb.CardType_FLOP
	msg.Num = 3

	bytes, err := proto.Marshal(msg)
	if err != nil {
		logrus.Errorf("Marshal %q err=%v", msg, err)
		return
	}
	c.Broadcast <- bytes
}
