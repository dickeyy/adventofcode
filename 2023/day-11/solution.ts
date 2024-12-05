// get and format the data
async function loadData() {
    // get the data from the file
    const file = Bun.file("../../inputs/2023/day-11/input.txt");
    return await file.text();
}

// part one
function partOne(data: string): number {
    const lines = data.split("\n");

    const addLines = (mat: string[][]): string[][] => {
        const size = mat.length;
        const nm: string[][] = [];

        for (let i = 0; i < size; i++) {
            let r = true;
            for (let j = 0; j < size; j++) {
                r = r && mat[i][j] !== "#";
            }
            nm.push([...mat[i]]);
            if (r) {
                nm.push(Array(size).fill("."));
            }
        }
        return nm;
    };

    const addCols = (mat: string[][]): string[][] => {
        const size = mat.length;
        const nm: string[][] = [];

        for (let i = 0; i < size; i++) {
            nm.push([]);
        }

        for (let i = 0; i < mat[0].length; i++) {
            let c = true;
            for (let j = 0; j < size; j++) {
                c = c && mat[j][i] !== "#";
                nm[j].push(mat[j][i]);
            }
            if (c) {
                for (let j = 0; j < size; j++) {
                    nm[j].push(".");
                }
            }
        }

        return nm;
    };

    const findPoints = (mat: string[][]): number[][] => {
        const p: any[] = [];

        for (let i = 0; i < mat.length; i++) {
            for (let j = 0; j < mat[i].length; j++) {
                if (mat[i][j] === "#") {
                    p.push({ x: i, y: j });
                }
            }
        }

        return p;
    };

    let mat: string[][] = lines.map((line) => line.split(""));
    mat = addLines(mat);
    mat = addCols(mat);
    const points = findPoints(mat);

    const dists: number[] = [];

    for (let i = 0; i < points.length; i++) {
        const p1: any = points[i];
        for (let j = i + 1; j < points.length; j++) {
            const p2: any = points[j];
            const x = Math.abs(p1.x - p2.x) + Math.abs(p1.y - p2.y);
            dists.push(x);
        }
    }

    return dists.reduce((sum, dist) => sum + dist, 0);
}

// console.log(partOne(await loadData()));

// part two
function partTwo(data: string): number {
    const lines = data.split("\n");

    const empty = (mat: string[][]): { er: { [key: number]: {} }; ec: { [key: number]: {} } } => {
        const size = mat.length;
        const er: { [key: number]: {} } = {};
        const ec: { [key: number]: {} } = {};

        for (let i = 0; i < size; i++) {
            let r = true;
            let c = true;
            for (let j = 0; j < size; j++) {
                r = r && mat[i][j] !== "#";
                c = c && mat[j][i] !== "#";
            }
            if (r) {
                er[i] = {};
            }
            if (c) {
                ec[i] = {};
            }
        }
        return { er, ec };
    };

    const findPoints = (mat: string[][]): number[][] => {
        const p: any[] = [];

        for (let i = 0; i < mat.length; i++) {
            for (let j = 0; j < mat[i].length; j++) {
                if (mat[i][j] === "#") {
                    p.push({ x: i, y: j });
                }
            }
        }

        return p;
    };

    let mat: string[][] = lines.map((line) => line.split(""));
    const result = empty(mat);
    const el: { [key: number]: {} } = result.er;
    const ec: { [key: number]: {} } = result.ec;
    const points = findPoints(mat);
    const distS: number[] = [];
    const weight = 1000000;

    for (let i = 0; i < points.length; i++) {
        const p1: any = points[i];
        for (let j = i + 1; j < points.length; j++) {
            const p2: any = points[j];
            let x = 0;

            for (let k = Math.min(p1.x, p2.x); k < Math.max(p1.x, p2.x); k++) {
                if (el[k] !== undefined) {
                    x += weight;
                } else {
                    x++;
                }
            }

            for (let k = Math.min(p1.y, p2.y); k < Math.max(p1.y, p2.y); k++) {
                if (ec[k] !== undefined) {
                    x += weight;
                } else {
                    x++;
                }
            }

            distS.push(x);
        }
    }

    return distS.reduce((sum, dist) => sum + dist, 0);
}

// console.log(partTwo(await loadData()));
