

export function writeError(err: any) {
    if (err) {
        console.log(err);
        throw err;
    }
}

export function isBaseTenNumber(x:string):boolean {
    return /^\d+$/.test(x)
}