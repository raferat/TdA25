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