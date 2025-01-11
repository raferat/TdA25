declare interface PromiseConstructor {
    after(timeout: number): Promise<Object>;
}

Promise.after = (timeout: number) => {
    return new Promise((resolve, reject) => {
        setTimeout(() => {
            resolve({});
        }, timeout);
    });
}

declare interface String {
    normalized(): string;
}

String.prototype.normalized = function () {
    let result = "";

    for (let i = 0; i < this.length; i++) {
        let ch = this.charAt(i).toLowerCase();

        if (ch == 'ě') ch = 'e';
        else if (ch == 'š') ch = 's';
        else if (ch == 'č') ch = 'c';
        else if (ch == 'ř') ch = 'r';
        else if (ch == 'ž') ch = 'z';
        else if (ch == 'ý') ch = 'y';
        else if (ch == 'á') ch = 'a';
        else if (ch == 'í') ch = 'i';
        else if (ch == 'é') ch = 'e';
        else if (ch == 'ú') ch = 'u';
        else if (ch == 'ů') ch = 'u';
        else if (ch == 'ď') ch = 'd';
        else if (ch == 'ť') ch = 't';
        else if (ch == 'ň') ch = 'n';
        else if (ch == 'ó') ch = 'o';

        ch = ch.trim();


        result = result + ch;
    }

    return result;
}