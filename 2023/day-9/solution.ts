// get and format the data
async function loadData() {
    // get the data from the file
    const file = Bun.file("./data.txt");
    return await file.text();
}

// part one
function partOne(data: string): number {
    const lines = data.split("\n");

    function solveRow(data: number[]) {
        let result = data[data.length - 1];
        do {
            let prev = data[0];
            data = data.slice(1).map((next) => {
                const tmp = next - prev;
                prev = next;
                return tmp;
            });
            result += data.length > 0 ? data[data.length - 1] : 0;
        } while (!data.every((v) => v === 0));
        return result;
    }

    return lines
        .map((l) => l.split(" ").map((n) => parseInt(n)))
        .map(solveRow)
        .reduce((acc, cur) => acc + cur, 0);
}

// console.log(partOne(await loadData()));

// part two
function partTwo(data: string): number {
    const lines = data.split("\n");

    function solveRow(data: number[]) {
        let result = data[data.length - 1];
        do {
            let prev = data[0];
            data = data.slice(1).map((next) => {
                const tmp = next - prev;
                prev = next;
                return tmp;
            });
            result += data.length > 0 ? data[data.length - 1] : 0;
        } while (!data.every((v) => v === 0));
        return result;
    }

    return lines
        .map((l) =>
            l
                .split(" ")
                .map((n) => parseInt(n))
                .reverse()
        )
        .map(solveRow)
        .reduce((acc, cur) => acc + cur, 0);
}

// console.log(partTwo(await loadData()));
