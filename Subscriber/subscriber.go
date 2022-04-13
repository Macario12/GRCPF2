package main

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/segmentio/kafka-go"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Log struct {
	GameId   int    `json:"game_id"`
	Players  int    `json:"players"`
	Winner   string `json:"winner"`
	GameName string `json:"game_n"`
	Queue    string `json:"queue"`
	Fecha    string `json:"Fecha"`
}

var collection *mongo.Collection
var ctx = context.TODO()

func main() {

	cadenaCompara := ""
	coneccition(&cadenaCompara)

}

func coneccition(cadenaCompara *string) {

	var cadenaRecibida string
	conn, err := kafka.DialLeader(context.Background(), "tcp", "34.125.140.78:9092", "topicF2", 0)
	if err != nil {
		log.Fatal("failed to dial leader:", err)
	}

	conn.SetReadDeadline(time.Now().Add(time.Second * 3))

	batch := conn.ReadBatch(1e3, 1e9) // fetch 10KB min, 1MB max

	bytes := make([]byte, 1e3) // 10KB max per message
	for {
		_, err := batch.Read(bytes)
		if err != nil {
			break
		}
		cadenaRecibida = string(bytes)
		//fmt.Println(string(bytes))
	}

	if strings.Compare(*cadenaCompara, cadenaRecibida) != 0 {
		var Log Log
		arrayValores := strings.Split(cadenaRecibida, ",")

		fmt.Println("id", arrayValores[0])
		fmt.Println("players", arrayValores[1])
		fmt.Println("winner", arrayValores[2])
		t := time.Now()
		idgame, _ := strconv.Atoi(arrayValores[0])
		players, _ := strconv.Atoi(arrayValores[1])
		//winner, _ := strconv.Atoi(arrayValores[2])
		Log.GameId = idgame
		Log.Players = players
		Log.GameName = parseGame(idgame)
		Log.Winner = "Jugador" + arrayValores[2]
		Log.Queue = "KAFKA"
		Log.Fecha = t.Format("2006-01-02 15:04:05")

		//Almacenar Mongo//Conectar con mongodb
		clientOptions := options.Client().ApplyURI("mongodb://admin:pass123@34.125.197.46:27017")
		client, err := mongo.Connect(ctx, clientOptions)
		if err != nil {
			log.Fatal(err)
		}
		//Crear colleccion y base de datos si no existen y registrar en coleccion
		collection = client.Database("SO1_Proyecto1_Fase2").Collection("Game_Logs")
		respuesta, err := collection.InsertOne(context.TODO(), Log)
		if err != nil {
			fmt.Print("Logs No Registrado")
			panic(err)
		} else {
			fmt.Print(respuesta)
		}

		*cadenaCompara = cadenaRecibida
	}

	coneccition(*&cadenaCompara)
}

func parseGame(id int) string {
	var nameGame string
	switch id {
	case 1:
		nameGame = "Random"
		return nameGame

	case 2:
		nameGame = "Frist"
		return nameGame

	case 3:
		nameGame = "Last"
		return nameGame

	case 4:
		nameGame = "Mitad"
		return nameGame

	case 5:
		nameGame = "RandomPrimo"
		return nameGame
	}

	return nameGame
}