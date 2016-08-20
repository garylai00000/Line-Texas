package DB
import (
	"os"
	"strconv"
	"github.com/line/line-bot-sdk-go/linebot"
	"database/sql"
	_"github.com/go-sql-driver/mysql"
)

var bot *linebot.Client

func InRoomInst(MID string){
	strID := os.Getenv("ChannelID")
	numID, _ := strconv.ParseInt(strID, 10, 64) // string to integer
	bot, _ = linebot.NewClient(numID, os.Getenv("ChannelSecret"), os.Getenv("MID"))
	db,_ := sql.Open("mysql", os.Getenv("dbacc")+":"+os.Getenv("dbpass")+"@tcp("+os.Getenv("dbserver")+")/")
	var nn string
	db.QueryRow("SELECT UserNickName FROM sql6131889.User WHERE MID = ?", MID).Scan(&nn)
	var R string
	db.QueryRow("SELECT UserRoom FROM sql6131889.User WHERE MID = ?", MID).Scan(&R)
	bot.SendText([]string{MID}, "Hello! "+nn+"!\nYou are in chatroom "+R+"\nyou can use these instruction:\n!leavechatroom\n!joingame\n!startgame\n!newgame")
}
func InRoomNewGame(MID string){
	strID := os.Getenv("ChannelID")
	numID, _ := strconv.ParseInt(strID, 10, 64) // string to integer
	bot, _ = linebot.NewClient(numID, os.Getenv("ChannelSecret"), os.Getenv("MID"))
	db,_ := sql.Open("mysql", os.Getenv("dbacc")+":"+os.Getenv("dbpass")+"@tcp("+os.Getenv("dbserver")+")/")
	var haveGame string
	var RID int
	var R string
	var GID string
	db.QueryRow("SELECT UserRoom FROM sql6131889.User WHERE MID = ?", MID).Scan(&R)
	db.QueryRow("SELECT ID FROM sql6131889.Room WHERE  RoomName = ?", R).Scan(&RID)
	db.QueryRow("SELECT RoomID FROM sql6131889.Game WHERE RoomID = ?", RID).Scan(&haveGame)
	var gameCancel int
	row,_ := db.Query("SELECT Cancel FROM sql6131889.Game WHERE RoomID = ?", RID)
	for row.Next() { 
		row.Scan(&gameCancel)
	}
	if haveGame == "" || gameCancel == 1 {
		db.Exec("INSERT INTO sql6131889.Game (GameName, RoomID, GameStatus, GameTokens, GamePlayer1, GameMaster, PlayerNum, Cancel) VALUES (?, ?, ?, ?, ?, ?, ?, ?)", "TexasPoker", RID, 1, 0, MID, "0", 1, 0)	
		db.QueryRow("SELECT ID FROM sql6131889.Game WHERE RoomID = ?", RID).Scan(&GID)
		db.Exec("INSERT INTO sql6131889.GameAction (MID, GameID, PlayerX, Action, Cancel) VALUES (?, ?, ?, ?, ?)", MID, GID, 1, 0, 0)
		db.Exec("UPDATE sql6131889.Room SET RoomStatus = ? WHERE RoomName = ?", 101, R)
		bot.SendText([]string{MID}, "You build a new game")
		bot.SendText([]string{MID}, "You are Player1")
	}else{
		bot.SendText([]string{MID}, "There is already a game in this room")
	}
	db.Close()
}
func InRoomJoinGame(MID string){
	strID := os.Getenv("ChannelID")
	numID, _ := strconv.ParseInt(strID, 10, 64) // string to integer
	bot, _ = linebot.NewClient(numID, os.Getenv("ChannelSecret"), os.Getenv("MID"))
	db,_ := sql.Open("mysql", os.Getenv("dbacc")+":"+os.Getenv("dbpass")+"@tcp("+os.Getenv("dbserver")+")/")
	var haveGame int
	var RID int
	var R string
	var GID int
	db.QueryRow("SELECT UserRoom FROM sql6131889.User WHERE MID = ? AND Cancel = ?", MID, 0).Scan(&R)
	db.QueryRow("SELECT ID FROM sql6131889.Room WHERE  RoomName = ? AND Cancel = ?", R,  0).Scan(&RID)
	db.QueryRow("SELECT ID FROM sql6131889.Game WHERE RoomID = ? AND Cancel = ?", RID, 0).Scan(&GID)
	db.QueryRow("SELECT RoomStatus FROM sql6131889.Room WHERE RoomName = ? AND Cancel = ?", R, 0).Scan(&haveGame)
	if haveGame == 100 {
		bot.SendText([]string{MID}, "Please build a new game use !newgame")
	}else{
		var playerInGame string
		db.QueryRow("SELECT MID FROM sql6131889.GameAction WHERE MID = ? AND Cancel = ?", MID, 0).Scan(&playerInGame)
		var nextPlayer int
		
		if playerInGame == "" {
			db.QueryRow("SELECT PlayerNum FROM sql6131889.Game WHERE ID = ? AND Cancel = ?", GID, 0).Scan(&nextPlayer)
			nextPlayer = nextPlayer+1
		}else{
			nextPlayer = 50
		}
		if nextPlayer <= 10 {
			db.Exec("INSERT INTO sql6131889.GameAction (MID, GameID, PlayerX, Action, Cancel) VALUES (?, ?, ?, ?, ?)", MID, GID, nextPlayer, 0, 0)
			if nextPlayer == 1 {
				db.Exec("UPDATE sql6131889.Game SET GamePlayer1 = ? WHERE ID = ? AND Cancel = ?", MID, GID, 0)
				db.Exec("UPDATE sql6131889.GameAction SET PlayerX = ? WHERE MID = ? AND Cancel = ?", 1, MID, 0)
				bot.SendText([]string{MID}, "You are Player1")
			}else if nextPlayer == 2 {
				db.Exec("UPDATE sql6131889.Game SET GamePlayer2 = ? WHERE ID = ? AND Cancel = ?", MID, GID, 0)
				bot.SendText([]string{MID}, "You are Player2")
			}else if nextPlayer == 3 {
				db.Exec("UPDATE sql6131889.Game SET GamePlayer3 = ? WHERE ID = ? AND Cancel = ?", MID, GID, 0)
				bot.SendText([]string{MID}, "You are Player3")
			}else if nextPlayer == 4 {
				db.Exec("UPDATE sql6131889.Game SET GamePlayer4 = ? WHERE ID = ? AND Cancel = ?", MID, GID, 0)
				bot.SendText([]string{MID}, "You are Player4")
			}else if nextPlayer == 5 {
				db.Exec("UPDATE sql6131889.Game SET GamePlayer5 = ? WHERE ID = ? AND Cancel = ?", MID, GID, 0)
				bot.SendText([]string{MID}, "You are Player5")
			}else if nextPlayer == 6 {
				db.Exec("UPDATE sql6131889.Game SET GamePlayer6 = ? WHERE ID = ? AND Cancel = ?", MID, GID, 0)
				bot.SendText([]string{MID}, "You are Player6")
			}else if nextPlayer == 7 {
				db.Exec("UPDATE sql6131889.Game SET GamePlayer7 = ? WHERE ID = ? AND Cancel = ?", MID, GID, 0)
				bot.SendText([]string{MID}, "You are Player7")
			}else if nextPlayer == 8 {
				db.Exec("UPDATE sql6131889.Game SET GamePlayer8 = ? WHERE ID = ? AND Cancel = ?", MID, GID, 0)
				bot.SendText([]string{MID}, "You are Player8")
			}else if nextPlayer == 9 {
				db.Exec("UPDATE sql6131889.Game SET GamePlayer9 = ? WHERE ID = ? AND Cancel = ?", MID, GID, 0)
				bot.SendText([]string{MID}, "You are Player9")
			}else if nextPlayer == 10 {
				db.Exec("UPDATE sql6131889.Game SET GamePlayer10 = ? WHERE ID = ? AND Cancel = ?", MID, GID, 0)
				bot.SendText([]string{MID}, "You are Player10")
			}
			db.Exec("UPDATE sql6131889.Game SET PlayerNum = ? WHERE ID = ? AND Cancel = ?", nextPlayer, GID, 0)
		}else if nextPlayer == 50 {
			bot.SendText([]string{MID}, "You are playing game now!!")
		}else{
			bot.SendText([]string{MID}, "The player is full in this room!!")
		}
	}
	db.Close()
}
func InRoomStartGame(MID string){
	strID := os.Getenv("ChannelID")
	numID, _ := strconv.ParseInt(strID, 10, 64) // string to integer
	bot, _ = linebot.NewClient(numID, os.Getenv("ChannelSecret"), os.Getenv("MID"))
	db,_ := sql.Open("mysql", os.Getenv("dbacc")+":"+os.Getenv("dbpass")+"@tcp("+os.Getenv("dbserver")+")/")
	var RID string
	var R string
	var GID string
	var haveGame int
	db.QueryRow("SELECT UserRoom FROM sql6131889.User WHERE MID = ? AND Cancel = ?", MID, 0).Scan(&R)
	db.QueryRow("SELECT ID FROM sql6131889.Room WHERE  RoomName = ? AND Cancel = ?", R, 0).Scan(&RID)
	db.QueryRow("SELECT ID FROM sql6131889.Game WHERE RoomID = ? AND Cancel = ?", RID, 0).Scan(&GID)
	var playerInGame string
	db.QueryRow("SELECT MID FROM sql6131889.GameAction WHERE MID = ? AND Cancel = ?", MID, 0).Scan(&playerInGame)
	db.QueryRow("SELECT RoomStatus FROM sql6131889.Room WHERE RoomName = ?,  AND Cancel = ?", R, 0).Scan(&haveGame)
	if haveGame == 101 {
		if playerInGame != ""{
			if true{
					var waitingForStart int
					db.QueryRow("SELECT GameStatus FROM sql6131889.Game WHERE ID = ? AND Cancel = ?", GID, 0).Scan(&waitingForStart)
					if waitingForStart == 1 {
						var gamerNum int
						db.QueryRow("SELECT PlayerNum FROM sql6131889.Game WHERE ID = ?", GID).Scan(&gamerNum)
						if gamerNum > 1 {
							db.Exec("UPDATE sql6131889.Game SET GameStatus = ? WHERE ID = ? AND Cancel = ?", 2, GID, 0) //starting game now
							bot.SendText([]string{MID}, "== START THE GAME ==")
						}else{
							bot.SendText([]string{MID}, "the game can't start below 2 player")
						}
					}else{
						bot.SendText([]string{MID}, "game is already starting")
					}
			}else{
				bot.SendText([]string{MID}, "You are not in the game")
			}
		}else{
			bot.SendText([]string{MID}, "You are not in the game")
		}
	}else{
		bot.SendText([]string{MID}, "There isn't any game is playing in this room now!")
	} 
}
func CancelGameAction(MID string){
	strID := os.Getenv("ChannelID")
	numID, _ := strconv.ParseInt(strID, 10, 64) // string to integer
	bot, _ = linebot.NewClient(numID, os.Getenv("ChannelSecret"), os.Getenv("MID"))
	db,_ := sql.Open("mysql", os.Getenv("dbacc")+":"+os.Getenv("dbpass")+"@tcp("+os.Getenv("dbserver")+")/")
	var RID string
	var R string
	var GID string
	db.QueryRow("SELECT UserRoom FROM sql6131889.User WHERE MID = ?", MID).Scan(&R)
	db.QueryRow("SELECT ID FROM sql6131889.Room WHERE  RoomName = ?", R).Scan(&RID)
	db.QueryRow("SELECT ID FROM sql6131889.Game WHERE RoomID = ?", RID).Scan(&GID)
	db.Exec("UPDATE sql6131889.GameAction SET Cancel = ? WHERE GameID = ?", 1, GID)
}
func CancelGame(MID string){
	strID := os.Getenv("ChannelID")
	numID, _ := strconv.ParseInt(strID, 10, 64) // string to integer
	bot, _ = linebot.NewClient(numID, os.Getenv("ChannelSecret"), os.Getenv("MID"))
	db,_ := sql.Open("mysql", os.Getenv("dbacc")+":"+os.Getenv("dbpass")+"@tcp("+os.Getenv("dbserver")+")/")
	var RID string
	var R string
	var GID string
	db.QueryRow("SELECT UserRoom FROM sql6131889.User WHERE MID = ?", MID).Scan(&R)
	db.QueryRow("SELECT ID FROM sql6131889.Room WHERE  RoomName = ?", R).Scan(&RID)
	db.QueryRow("SELECT ID FROM sql6131889.Game WHERE RoomID = ?", RID).Scan(&GID)
	db.Exec("UPDATE sql6131889.Game SET Cancel = ? WHERE RoomID = ?", 1, RID)
}