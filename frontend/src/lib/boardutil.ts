import { type Board } from "./api";

export function isWinMove(board: Board, x: number, y: number, symbol: "X" | "O"): boolean {
    let horizontal = getSymbolsInDirection({x: -1, y: 0}, board, x, y, symbol) + getSymbolsInDirection({x: 1, y: 0}, board, x, y, symbol);
    let vertical = getSymbolsInDirection({x: 0, y: -1}, board, x, y, symbol) + getSymbolsInDirection({x: 0, y: 1}, board, x, y, symbol);
    let diagonal1 = getSymbolsInDirection({x: -1, y: -1}, board, x, y, symbol) + getSymbolsInDirection({x: 1, y: 1}, board, x, y, symbol);
    let diagonal2 = getSymbolsInDirection({x: -1, y: 1}, board, x, y, symbol) + getSymbolsInDirection({x: 1, y: -1}, board, x, y, symbol);



    return horizontal >= 4 || vertical >= 4 || diagonal1 >= 4 || diagonal2 >= 4;
}

function getSymbolsInDirection(direction: {x: number, y: number}, board: Board, startX: number, startY: number, symbol: "X" | "O"): number {
    let x = startX + direction.x;
    let y = startY + direction.y;

    let cnt = 0;
    while (x >= 0 && y >= 0 && y < board.length && x < board[y].length) {
        if (board[y][x] == symbol) {
            cnt ++;
        } else {
            break;
        }


        x += direction.x;
        y += direction.y;
    }
    return cnt;
}

export function calculateNextSymbol(board: Board): "X" | "O" {
    let cntX = 0;
    let cntO = 0;

    for (const row of board) {
        for (const element of row) {
            if (element == "X")
                cntX ++;
            else if (element == "O")
                cntO ++;
        }
    }

    if (cntX > cntO)
        return "O";
    return "X";
}