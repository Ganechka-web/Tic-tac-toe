package main

import (
	"fmt"
	"strings"
)


type TTTGame struct {
    square map[string]string
    player1 string
    player2 string
    player1_sign string 
    player2_sign string
}

func (g *TTTGame) Init(player1_name string, 
                       player2_name string) {
    // players definition 
    g.player1 = player1_name
    g.player1_sign = "X"
    g.player2 = player2_name
    g.player2_sign = "O"
                        
    // square definition 
    g.square = map[string]string{}
    for i := 0; i < 3; i++ {
        for j := 0; j < 3; j++ {
            g.square[fmt.Sprint(i) + fmt.Sprint(j)] = "       "
        }
    }
}

func (g *TTTGame) PrintCurSquare() {
    fmt.Println("-------------------------")
    for i := 0; i < 3; i++ {
        fmt.Print("|")
        for j := 0; j < 3; j++ {
            var key string = fmt.Sprint(i) + fmt.Sprint(j)

            if elem, ok := g.square[key]; ok && strings.TrimSpace(elem) != "" {
                fmt.Printf("%s|", elem)
            } else {
                fmt.Printf("   %s  |", key)
            }
        }
        fmt.Println("\n-------------------------")
    }
}

func (g *TTTGame) PlayerTurn(target string, player_sybmol string) {
    if elem, ok := g.square[target]; ok {
        if strings.ContainsAny(elem, "XO") {
            fmt.Println("This place is busy!!! You are missing your turn")
        } else {
            g.square[target] = fmt.Sprintf("   %s   ", player_sybmol)
        }
    } else {
        fmt.Println("This place is unreachable!!! You are missing your turn")
    }
}
 
func IsEqual(item string, item2 string, item3 string) (bool, string) {
    var row string = strings.ReplaceAll(item + item2 + item3, " ", "")

    if strings.Count(row, "X") == 3 {
        return true, "X"
    } else if strings.Count(row, "O") == 3 {
        return true, "O"
    } 
    return false, ""
}

func (g *TTTGame) FindWinner() string {
    // check rows
    for i := 0; i < 3; i++ {
        result, symbol := IsEqual(g.square[fmt.Sprintf("%d0", i)], 
                                  g.square[fmt.Sprintf("%d1", i)],
                                  g.square[fmt.Sprintf("%d2", i)])
        if result {
            return symbol
        }
    }

    // check colls 
    for i := 0; i < 3; i++ {
        result, symbol := IsEqual(g.square[fmt.Sprintf("0%d", i)], 
                                  g.square[fmt.Sprintf("1%d", i)],
                                  g.square[fmt.Sprintf("2%d", i)])
        if result {
            return symbol
        }
    }

    // check diagonals
    result_right, symbol_right := IsEqual(g.square["02"], g.square["11"], g.square["20"])
    result_left, sybmol_left := IsEqual(g.square["00"], g.square["11"], g.square["22"])
    if result_right {
        return symbol_right
    }
    if result_left {
        return sybmol_left
    }
    return ""
}

func main() {
    var player1_name string
    var player2_name string

    fmt.Println("Enter the first player name: ")
    fmt.Scan(&player1_name)
    fmt.Println("Enter the second player name: ")
    fmt.Scan(&player2_name)

    var Game TTTGame = TTTGame{}
    Game.Init(player1_name, player2_name)

    var i int = 0
    for {
        Game.PrintCurSquare()
        
        var target string;

        if i % 2 == 0 {
            fmt.Printf("%s turn (X). Select and type place on the board\n", player1_name)
            fmt.Scan(&target)

            Game.PlayerTurn(target, "X")
        } else {
            fmt.Printf("%s turn (O). Select and type place on the board\n", player2_name)
            fmt.Scan(&target)

            Game.PlayerTurn(target, "O")        
        }

        if winnerSign := Game.FindWinner(); winnerSign != "" {
            Game.PrintCurSquare()
            if winnerSign == "X" {
                fmt.Printf("Game over! Winner is %s\n", player1_name)
            } else {
                fmt.Printf("Game over! Winner is %s\n", player2_name)
            }
            break
        }
        i++
    }
}