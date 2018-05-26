package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/calc_service/src/api"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":12345", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %s", err)
	}
	c := api.NewCalcClient(conn)
	defer conn.Close()

	var x int64
	var y int64
	var operator string
	for {

		fmt.Print("Insert the expresion ")
		reader := bufio.NewReader(os.Stdin)
		exp, inErr := reader.ReadString('\n')
		if inErr != nil {
			log.Fatal(inErr)
			continue
		}
		x, y, operator, err = parse(&exp)
		if err != nil {
			log.Println("Wrong expression")
			continue
		}

		switch operator {
		case "+":
			doSummation(c, x, y)

		case "-":
			doSubtraction(c, x, y)

		case "*":
			doMultiplication(c, x, y)

		case "/":
			doDivision(c, x, y)

		default:
			log.Println("Wrong Opernad")

		}

	}
}

// Parse the expresion, returns the operands and the opertor
func parse(str *string) (int64, int64, string, error) {
	re := regexp.MustCompile(`[-+*/]`)
	s := re.Split((*str), -1)

	if len(s) == 2 {

		x, err := strconv.ParseInt(strings.TrimSpace(s[0]), 10, 64)
		if err != nil {
			return 0, 0, "", err
		}

		y, err := strconv.ParseInt(strings.TrimSpace(s[1]), 10, 64)
		if err != nil {
			return 0, 0, "", err
		}

		operator := getOperator(str, &s[0])

		return x, y, operator, nil
	}
	return 0, 0, "", errors.New("Failed to parse")
}

// Get the operator from the expression.
// By knowing the prefix length preceding the operator,
// we can access the operator from the expression as an index.
func getOperator(str *string, prefix *string) string {
	return string((*str)[len((*prefix))])
}

func doMultiplication(c api.CalcClient, x int64, y int64) {
	mulResponse, err := c.Mul(context.Background(), &api.Args{Arg1: x, Arg2: y})
	if err != nil {
		log.Fatalf("Error when calling Mul: %s", err)
	}
	log.Printf("Response from server: %d", mulResponse.Value)
}

func doSummation(c api.CalcClient, x int64, y int64) {
	sumResponse, err := c.Sum(context.Background(), &api.Args{Arg1: x, Arg2: y})
	if err != nil {
		log.Fatalf("Error when calling Sum: %s", err)
	}
	log.Printf("Response from server: %d", sumResponse.Value)
}

func doDivision(c api.CalcClient, x int64, y int64) {
	divResponse, err := c.Div(context.Background(), &api.Args{Arg1: x, Arg2: y})
	if err != nil {
		log.Fatalf("Error when calling Div: %s", err)
	}
	log.Printf("Response from server: %d", divResponse.Value)
}

func doSubtraction(c api.CalcClient, x int64, y int64) {
	subResponse, err := c.Sub(context.Background(), &api.Args{Arg1: x, Arg2: y})
	if err != nil {
		log.Fatalf("Error when calling Sub: %s", err)
	}
	log.Printf("Response from server: %d", subResponse.Value)
}
