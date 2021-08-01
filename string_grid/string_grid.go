package string_grid

// Grid stores a size and a Matrix
// The matrix is a slice of slices used for storing values
type Grid struct{
  Size int
  Matrix [][]string
  ZeroValue string
}

// MakeGrid makes a grid of variable size, returns grid
func MakeGrid (size int, zeroValue string) Grid {
  g := Grid{}
  g.Size = size
  g.ZeroValue = zeroValue 
  g.Matrix = make([][]string, size)
  for i := 0; i < size; i++ {
    g.Matrix[i] = make([]string, size)
    for j := 0; j < size; j ++ {
      g.Matrix[i][j]  = zeroValue
    }
  }
  return g
}

// ResetMatrix is a grid method that sets matrix elements to zero value
func (g *Grid) Reset() {
  for i := 0; i < g.Size; i++ {
    for j :=0; j < g.Size; j++ {
      g.Matrix[i][j] = g.ZeroValue
    }
  }
}

// CountBlankPositions counts all zero values in the grid matrix
// returns count
func (g *Grid) CountBlankPositions() int{
  ctr := 0
  for i := 0; i < g.Size; i++ {
    for j :=0; j < g.Size; j++ {
      if g.Matrix[i][j] == g.ZeroValue {
       ctr ++ 
      }
    }
  }
  return ctr
}

// PositionIsOccupied is a boolean method to detect whether position has a ZeroValue
// returns true/false
func (g *Grid) PositionIsOccupied(x, y int) bool{
  return g.Matrix[x][y] != g.ZeroValue 
}

// HorizontalData provides array of row data (i.e. our matrix)
// returns array of column[] data
func (g *Grid) VerticalData() [][]string {
  retVar := make([][]string, g.Size)
  for i := 0; i < g.Size; i++ {
    retVar[i] = make([]string, g.Size)
  }
  for i := 0; i < g.Size; i++ {
    for j :=0; j < g.Size; j++ {
      retVar[j][i] = g.Matrix[i][j]
    }
  }
  return retVar
}

// HorizontalData provides array of row data (i.e. our matrix)
// function is for readability
// returns array of row[] data
func (g *Grid) HorizontalData() [][]string {
  return g.Matrix
}

// DiagonalData provides data from the two full length diagonals present in the matrix
func (g *Grid) DiagonalData() [][]string {
  retVar := make([][]string, 2)
  retVar[0] = make([]string, g.Size)
  retVar[1] = make([]string, g.Size)
  for i := 0; i < g.Size; i++ {
    retVar[0][i] = g.Matrix[i][i]
    retVar[1][i] = g.Matrix[g.Size - 1 - i][i]
  }
  return retVar
}
