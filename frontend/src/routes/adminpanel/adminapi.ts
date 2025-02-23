interface User {
    uuid: string,
    createdAt: string,
    username: string,
    email: string,
    elo: number,
    wins: number,
    draws: number,
    losses: number,
}

export function getUserData(): User[] {
    
}