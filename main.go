package main

import (
  "bufio"
  "fmt"
  "github.com/mabiesen/go-tic_tac_toe/string_grid"
  "github.com/olekukonko/tablewriter"
  "os"
  "strconv"
  "strings"
)

type player struct{
  gameString string
  playerOrdinal int
}

func intro(){
  fmt.Println("This is a game of tic tac toe!")
  fmt.Println("Player 1 will be 'X', Player 2 will be 'O'")
  fmt.Println("Have fun!\n\n")
}

func outtro(winner player){
  if winner.playerOrdinal != 0{
     fmt.Println(fmt.Sprintf("Congrats player %d, you won!!", winner.playerOrdinal))
  }else{
    fmt.Println("Tie game!")
 }
 fmt.Println("Thanks for playing!")
}

func printGameBoard(g string_grid.Grid){
  table := tablewriter.NewWriter(os.Stdout)
  for _, v := range g.Matrix {
    table.Append(v)
  }
  table.Render()
}

func getInput(msg string)(string){
  fmt.Println(msg)
  reader := bufio.NewReader(os.Stdin)
  text, _ := reader.ReadString('\n')
  text = strings.Replace(text, "\n", "", -1)
  return text
}

func getCoordinate(coordchar string, maxval int)int{
  msg := fmt.Sprintf("Please provide a %s coordinate (0-%d), first row is zero)", coordchar, maxval - 1)
  resp, err := strconv.Atoi(getInput(msg))
  if resp > maxval {
    fmt.Println("Input higher than max val!")
    fmt.Println("Please try again")
    resp = getCoordinate(coordchar, maxval)
  }
  if err !=nil {
    fmt.Println("Unable to convert  input to string; please try again")
    resp = getCoordinate(coordchar, maxval)
  }
  return  resp
}

func getCoordinates(g string_grid.Grid)(int, int){
  xpos := getCoordinate("x", g.Size - 1)
  ypos := getCoordinate("y", g.Size - 1)
  if g.Matrix[ypos][xpos] != g.ZeroValue {
    fmt.Println("The grid already has a value here! Please try again")
    xpos, ypos = getCoordinates(g)
  }
  return xpos, ypos
}

func sliceAllSame(slc []string) bool {
  for i := 1; i < len(slc); i++ {
    if slc[i] != slc[0] {
      return false
    }
  }
  return true
}

func currentPlayer(players []player, turn int) player {
  if turn % 2 != 0 {
    return players[0]
  }
  return players[1]
}

func gameWinner(g string_grid.Grid, players []player) player {
  v_data := g.VerticalData()
  h_data := g.HorizontalData()
  d_data := g.DiagonalData()
  data := [][][]string{v_data, h_data, d_data}
  for i := 0; i < len(data); i++ {
     for j :=0; j < len(data[i]); j++ {
       if data[i][j][0] == g.ZeroValue{
         continue
       }
       if sliceAllSame(data[i][j]) {
         if data[i][j][0] == players[0].gameString{
           return  players[0]
         }else{
           return players[1]
         }
       }
     }
  }
  return player{}
}

func main() {
  intro()
  players := []player{ player{"X", 1}, player{"0", 2} }
  var current_player player
  var winner player
  mygrid := string_grid.MakeGrid(3, "-")
  for turn := 1; ;turn++ {
    printGameBoard(mygrid)
    winner = gameWinner(mygrid, players)
    if (winner != player{} || mygrid.CountBlankPositions() == 0){
      break
    }
    current_player = currentPlayer(players, turn)
    fmt.Println("You're up Player ", current_player.playerOrdinal)
    xpos, ypos := getCoordinates(mygrid)
    mygrid.Matrix[ypos][xpos] = current_player.gameString
  }
  outtro(winner)
}
