func numIslands(grid [][]byte) int {
    result := 0

    for y := 0; y < len(grid); y++ {
        for x := 0; x < len(grid[y]); x++ {
            if grid[y][x] == '1' {
                result++
                checkIsland(grid, y, x)
            }
        }
    }
    return result
}

func checkIsland(grid [][]byte, y, x int) {
    if y < 0 || y >= len(grid) || x < 0 || x >= len(grid[y]) || grid[y][x] != '1' {
        return
    }

    grid[y][x] = '2'
    checkIsland(grid, y + 1, x)
    checkIsland(grid, y - 1, x)
    checkIsland(grid, y, x + 1)
    checkIsland(grid, y, x - 1)
}
