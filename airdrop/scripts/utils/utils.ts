

export function writeError(err: any) {
    if (err) {
        console.log(err);
        throw err;
    }
}