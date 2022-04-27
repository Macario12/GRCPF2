package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type LogStruct struct {
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

	coneccition()

}

func coneccition() {
	var cadenaRecibida string
	var Log LogStruct //Log players

	var arrarLogs []LogStruct

	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": os.Getenv("ADD_KAFKA"),
		"group.id":          "myGroup",
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		panic(err)
	}

	c.SubscribeTopics([]string{"topicF2", "^aRegex.*[Tt]opic"}, nil)

	for {
		msg, err := c.ReadMessage(-1)
		if err == nil {
			cadenaRecibida = string(msg.Value)
			//fmt.Println(string(bytes))

			//array log
			arrayValores := strings.Split(cadenaRecibida, ",")

			fmt.Println("id", arrayValores[0])
			fmt.Println("players", arrayValores[1])
			fmt.Println("winner", arrayValores[2])
			t := time.Now()
			idgame, _ := strconv.Atoi(arrayValores[0])
			players, _ := strconv.Atoi(arrayValores[1])
			Log.GameId = idgame
			Log.Players = players
			Log.GameName = parseGame(idgame)
			Log.Winner = "Jugador" + arrayValores[2]
			Log.Queue = "KAFKA"
			Log.Fecha = t.Format("2006-01-02 15:04:05")

			arrarLogs = append(arrarLogs, Log)
			//Almacenar Mongo//Conectar con mongodb
			clientOptions := options.Client().ApplyURI("mongodb://admin:123@34.125.159.79:27017")
			client, errmongo := mongo.Connect(ctx, clientOptions)
			if errmongo != nil {
				log.Fatal(errmongo)
			}
			/*	//cliente redis
				clientRedis := redis.NewClient(&redis.Options{
					Addr:     "34.125.235.244:6379",
					Password: "",
					DB:       0,
				})
				//cliente tidis
				clienttidis := redis.NewClient(&redis.Options{
					Addr:     "34.125.12.54:5379",
					Password: "",
					DB:       0,
				})

				//Crear colleccion y base de datos si no existen y registrar en coleccion

			*/
			collection = client.Database("SO1_Proyecto1_Fase2").Collection("Game_Logs")
			respuesta, errcoleccion := collection.InsertOne(context.TODO(), Log)
			if errcoleccion != nil {
				fmt.Print("Logs No Registrado")
				panic(errcoleccion)
			} else {
				fmt.Print(respuesta)
			}
			/*datosTiemporeal, err := json.Marshal(Log)

			//guardar dato en tiempo real en redis
			errorRedis := clientRedis.Set("tiempoReal", datosTiemporeal, 0).Err()
			if errorRedis != nil {
				panic(errorRedis)
			}

			errorTidis := clienttidis.Set("tiempoReal", datosTiemporeal, 0).Err()
			if errorTidis != nil {
				panic(errorTidis)
			}
			*/

			fmt.Println("Redis: Valor agregado Correctamente")

			fmt.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))

		} else {
			// The client will automatically try to recover from all errors.
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
		}
	}
	c.Close()
	coneccition()
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
