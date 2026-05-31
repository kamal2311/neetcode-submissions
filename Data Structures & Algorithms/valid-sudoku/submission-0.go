func isValidSudoku(board [][]byte) bool {
    
    for i := 0; i < 9; i++ {

        if !isValidRow(board, i) {
            return false
        }

        if !isValidCol(board,i) {
            return false
        }
    } 

    cellMarkers := [3]int{0,3,6}

    for i := 0; i < 3; i ++ {
        for j := 0; j < 3; j++ {
            if !isValidCell(board,cellMarkers[i],cellMarkers[j]) {
                return false
            }
        }             
    }

    return true

}

func isValidRow(board [][]byte, row int) bool {
    ur := make(map[byte]struct{})

    for i := 0; i < 9; i++ {

        if board[row][i] == '.'{
            continue
        }

        if _, ok := ur[board[row][i]]; ok {
            return false
        }

        ur[board[row][i]] = struct{}{}
    }

    return true

}

func isValidCol(board [][]byte, col int) bool {
    
    uc := make(map[byte]struct{})

    for i := 0; i < 9; i++ {

        if board[i][col] == '.'{
            continue
        }

        if _, ok := uc[board[i][col]]; ok {
            return false
        }

        uc[board[i][col]] = struct{}{}
    }

    return true
}

func isValidCell(board [][]byte, leftRow int, leftCol int) bool {
    
    ucell := make(map[byte]struct{})

    for i := 0; i < 3; i++ {

        for j := 0; j < 3; j++ {

            if board[leftRow+i][leftCol+j] == '.'{
                continue
            }

            if _, ok := ucell[board[leftRow+i][leftCol+j]]; ok {
                return false
            }

            ucell[board[leftRow+i][leftCol+j]] = struct{}{}
                    
        }     

    }

    return true
}
