// get and format the data
async function loadData() {
    // get the data from the file
    const file = Bun.file("../../inputs/2023/day-5/input.txt");
    return await file.text();
}

// part one
function partOne(data: string): number {
    const [rawSeeds, ...rawMaps] = data.split("\n\n");
    const seeds = [...rawSeeds.matchAll(/\d+/g)].map(Number);
    let maps: any[][] = [];

    for (const rawMap of rawMaps) {
        const ranges = [...rawMap.matchAll(/(\d+) (\d+) (\d+)/g)].map((range) => {
            const [dest, src, length] = range.slice(1).map(Number);
            return [
                [src, src + (length - 1)],
                [dest, dest + (length - 1)]
            ];
        });
        maps.push(ranges);
    }

    const inRange = (value: number, [min, max]: [number, number], offset = 0) => {
        return (value - (min - offset)) * (value - (max + offset)) <= 0;
    };

    const calc = (value: number, maps: any[][]) => {
        for (const map of maps) {
            const range = map.find(([input]) => inRange(value, input));
            if (!range) continue;
            const [[input], [output]] = range;
            value = output + value - input;
        }
        return value;
    };

    let location = Infinity;

    for (const seed of seeds) {
        location = Math.min(location, calc(seed, maps));
    }

    return location;
}

// console.log(partOne(await loadData()))

// part two
function partTwo(data: string) {
    const [rawSeeds, ...rawMaps] = data.split("\n\n");
    const seeds = [...rawSeeds.matchAll(/\d+/g)].map(Number);
    let maps: any[][] = [];

    for (const rawMap of rawMaps) {
        const ranges = [...rawMap.matchAll(/(\d+) (\d+) (\d+)/g)].map((range) => {
            const [dest, src, length] = range.slice(1).map(Number);
            return [
                [src, src + (length - 1)],
                [dest, dest + (length - 1)]
            ];
        });
        maps.push(ranges);
    }

    const seedRanges = Array.from({ length: seeds.length / 2 }, (_, index) => {
        const [start, length] = seeds.slice(index * 2, index * 2 + 2);
        return [start, start + (length - 1)];
    });

    maps = maps.toReversed().map((ranges) => ranges.map((range) => range.toReversed()));

    const inRange = (value: number, [min, max]: [number, number], offset = 0) => {
        return (value - (min - offset)) * (value - (max + offset)) <= 0;
    };

    const calc = (value: number, maps: any) => {
        for (const map of maps) {
            const range = map.find(([input]: any) => inRange(value, input));
            if (!range) continue;
            const [[input], [output]] = range;
            value = output + value - input;
        }
        return value;
    };

    let location = 0;

    while (true) {
        const seed = calc(location, maps);

        if (seedRanges.some((range) => inRange(seed, range as any))) break;

        location++;
    }

    return location;
}

// console.log(partTwo(await loadData()))
