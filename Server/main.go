package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"net"
	"strconv"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	pb "github.com/macario12/GRCPF2/protos"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

type server struct {
	pb.UnimplementedGameServiceServer
}

// SayHello implements helloworld.GreeterServer
func (s *server) AddGame(ctx context.Context, in *pb.GameRequest) (*pb.GameResponse, error) {
	log.Printf("Received: %v", in.GetGameId())
	winner := strconv.Itoa(selectGame(int(in.GetGameId()), int(in.GetPlayers())))
	//conexion kafka

	//conn, err := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", "topicF2", 0)

	/*if err != nil {
		log.Fatal("failed to dial leader:", err)
	}
	conn.SetWriteDeadline(time.Now().Add(time.Second * 10))*/
	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "localhost"})
	if err != nil {
		panic(err)
	}

	defer p.Close()

	cadenaKafka := strconv.Itoa(int(in.GetGameId())) + "," + strconv.Itoa(int(in.GetPlayers())) + "," + winner + "," + "fin"

	// Delivery report handler for produced messages
	go func() {
		for e := range p.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					fmt.Printf("Delivery failed: %v\n", ev.TopicPartition)
				} else {
					fmt.Printf("Delivered message to %v\n", ev.TopicPartition)
				}
			}
		}
	}()

	// Produce messages to topic (asynchronously)
	topic := "topicF2"
	//for _, word := range []string{cadenaKafka} {
	p.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          []byte(cadenaKafka),
	}, nil)
	///}

	// Wait for message deliveries before shutting down
	p.Flush(15 * 1000)
	//conn.WriteMessages(kafka.Message{Value: []byte(cadenaKafka)})

	return &pb.GameResponse{Message: "Winner: " + winner + " in Game: " + strconv.Itoa(int(in.GetGameId()))}, nil
}

func selectGame(numGame, players int) int {
	switch numGame {
	case 1:
		fmt.Println("Ejecutando Juego 1")
		return GameRadmon(players)
	case 2:
		fmt.Println("Ejecutando Juego 2")
		return GameFrist(players)
	case 3:
		fmt.Println("Ejecutando Juego 3")
		return GameLast(players)
	case 4:
		fmt.Println("Ejecutando Juego 4")
		return GameMedium(players)
	case 5:
		fmt.Println("Ejecutando Juego 5")
		return GamePrimoRadmon(players)

	}
	return 0
}

func GameRadmon(players int) int {
	winner := rand.Intn(players)
	return winner
}

func GameFrist(players int) int {
	winner := 1
	return winner
}

func GameLast(players int) int {
	winner := players
	return winner
}

func GameMedium(players int) int {
	winner := int(players / 2)
	return winner
}

func GamePrimoRadmon(players int) int {
	tamañoPrimos := 0
	for i := 2; i < players; i++ {
		if Esprimo(i) {
			tamañoPrimos = tamañoPrimos + 1
			//fmt.Println(i)
		}
	}
	//fmt.Println(tamañoPrimos)
	arrayPrimos := make([]int, tamañoPrimos)

	contadorArray := 0
	for i := 2; i < players; i++ {
		if Esprimo(i) {
			arrayPrimos[contadorArray] = i
			contadorArray++
		}

	}
	winnerPos := rand.Intn(tamañoPrimos - 1)
	fmt.Println(winnerPos)
	winner := arrayPrimos[winnerPos]

	return winner
}

func Esprimo(numero int) bool {
	contador := 2
	primo := true
	for (primo) && (contador != numero) {
		if numero%contador == 0 {
			primo = false
		}
		contador = contador + 1
	}

	return primo
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGameServiceServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
