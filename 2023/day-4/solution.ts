// get and format the data
async function loadData() {
    // get the data from the file
    const file = Bun.file("../../inputs/2023/day-4/input.txt");
    return await file.text();
}

// part 1
function partOne(data: string) {
    let total = 0;

    // split each line
    const lines = data.split("\n");

    for (const line of lines) {
        let cardValue = 0;

        const split = line.split(":");
        const winningNums = split[1]
            .split("|")[0]
            .trim()
            .split(" ")
            .filter((str) => str.trim());
        const ourNums = split[1]
            .split("|")[1]
            .trim()
            .split(" ")
            .filter((str) => str.trim());

        for (let i = 0; i < ourNums.length; i++) {
            // check if our num is in the winningnums array
            if (winningNums.includes(ourNums[i])) {
                if (cardValue == 0) cardValue++;
                else {
                    cardValue *= 2;
                }
            }
        }

        total += cardValue;
    }

    return total;
}

// partOne(await loadData())

// part 2
function partTwo(data: string) {
    const copies: any = {};

    // split each line
    const lines = data.split("\n");

    for (let i = lines.length - 1; i >= 0; i--) {
        const split = lines[i].split(":");
        const winningNums = split[1]
            .split("|")[0]
            .trim()
            .split(" ")
            .filter((str) => str.trim());
        const ourNums = split[1]
            .split("|")[1]
            .trim()
            .split(" ")
            .filter((str) => str.trim());

        let winners = 0;
        for (let i = 0; i < ourNums.length; i++) {
            // check if our num is in the winningnums array
            if (winningNums.includes(ourNums[i])) {
                winners++;
            }
        }

        let newCards = winners;
        for (let w = 1; w <= winners; w++) {
            if (i + w < lines.length) {
                newCards += copies[i + w];
            }
        }

        copies[i] = newCards;
    }

    return (
        (Object.values(copies) as number[]).reduce((partialSum, a) => partialSum + (parseInt(String(a)) || 0), 0) +
        lines.length
    );
}

// partTwo(await loadData())
