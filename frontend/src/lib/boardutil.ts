import { type Board } from "./api";

export function isWinMove(board: Board, x: number, y: number, symbol: "X" | "O"): boolean {
    let horizontal = getSymbolsInDirection({x: -1, y: 0}, board, x, y, symbol) + getSymbolsInDirection({x: 1, y: 0}, board, x, y, symbol);
    let vertical = getSymbolsInDirection({x: 0, y: -1}, board, x, y, symbol) + getSymbolsInDirection({x: 0, y: 1}, board, x, y, symbol);
    let diagonal1 = getSymbolsInDirection({x: -1, y: -1}, board, x, y, symbol) + getSymbolsInDirection({x: 1, y: 1}, board, x, y, symbol);
    let diagonal2 = getSymbolsInDirection({x: -1, y: 1}, board, x, y, symbol) + getSymbolsInDirection({x: 1, y: -1}, board, x, y, symbol);



    return horizontal >= 4 || vertical >= 4 || diagonal1 >= 4 || diagonal2 >= 4;
}

export function getWinBounds(board: Board, x: number, y: number, symbol: "X" | "O"): {x1: number, y1: number, x2: number, y2: number} {
    let horizontal = getSymbolsInDirection({x: -1, y: 0}, board, x, y, symbol) + getSymbolsInDirection({x: 1, y: 0}, board, x, y, symbol);
    let vertical = getSymbolsInDirection({x: 0, y: -1}, board, x, y, symbol) + getSymbolsInDirection({x: 0, y: 1}, board, x, y, symbol);
    let diagonal1 = getSymbolsInDirection({x: -1, y: -1}, board, x, y, symbol) + getSymbolsInDirection({x: 1, y: 1}, board, x, y, symbol);
    let diagonal2 = getSymbolsInDirection({x: -1, y: 1}, board, x, y, symbol) + getSymbolsInDirection({x: 1, y: -1}, board, x, y, symbol);

    if ( horizontal >= 4 ) {
        const y1 = y;
        const y2 = y;
        const x1 = x - getSymbolsInDirection({x: -1, y: 0}, board, x, y, symbol);
        const x2 = x + getSymbolsInDirection({x: 1, y: 0}, board, x, y, symbol);
        return {x1, y1, x2, y2};
    }
    if ( vertical >= 4 ) {
        const x1 = x;
        const x2 = x;
        const y1 = y - getSymbolsInDirection({x: 0, y: -1}, board, x, y, symbol);
        const y2 = y + getSymbolsInDirection({x: 0, y: 1}, board, x, y, symbol);
        return {x1, y1, x2, y2};
    }
    if ( diagonal1 >= 4 ) {
        const lu = getSymbolsInDirection({x: -1, y: -1}, board, x, y, symbol);
        const rd = getSymbolsInDirection({x: 1, y: 1}, board, x, y, symbol);

        const x1 = x - lu;
        const y1 = y - lu;
        const x2 = x + rd;
        const y2 = y + rd;

        return {x1, y1, x2, y2};
    }
    if ( diagonal2 >= 4 ) {
        const ru = getSymbolsInDirection({x: 1, y: -1}, board, x, y, symbol);
        const ld = getSymbolsInDirection({x: -1, y: 1}, board, x, y, symbol);

        const x1 = x - ld;
        const y1 = y + ld;
        const x2 = x + ru;
        const y2 = y - ru;

        return {x1, y1, x2, y2};
    }

    return {x1: 0,x2: 0,y2: 0,y1: 0};
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

export function isBoardCorrect(board: Board | undefined): boolean {
    if (board == undefined) return false;
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

    return cntX == cntO+1 || cntX == cntO;
}