// get and format the data
async function loadData() {
    // get the data from the file
    const file = Bun.file("../../inputs/2023/day-6/input.txt");
    return await file.text();
}

// part one
function partOne(data: string) {
    const times: number[] = data
        .split("\n")[0]
        .split(" ")
        .filter((value) => value)
        .filter((value) => value !== "Time:")
        .map(Number); // the times

    const records = data
        .split("\n")[1]
        .split(" ")
        .filter((value) => value)
        .filter((value) => value !== "Distance:")
        .map(Number); // the records to beat

    let viableWays = 1; // the amount of viable ways

    const calcWays = (time: number, record: number) => {
        let ways = 0;

        for (let i = 1; i < time; i++) {
            // i represents the time spent pressing the button up until the time alloted
            const nav = time - i; // represents the amout of time it has after pressing the button to travel
            let dist = nav * i; // represents the distance it travels
            if (dist > record) {
                // if the distance it went is greater than the record, this is viable
                ways += 1;
            }
        }
        return ways;
    };

    // loop through each time slot
    for (let i = 0; i < times.length; i++) {
        viableWays *= calcWays(times[i], records[i]); // multiply the ways together
    }

    return viableWays;
}

// console.log(partOne(await loadData()));

// part two
function partTwo(data: string) {
    const time = Number(
        data
            .split("\n")[0]
            .split(" ")
            .filter((value) => value)
            .filter((value) => value !== "Time:")
            .join("")
    ); // the time

    const record = Number(
        data
            .split("\n")[1]
            .split(" ")
            .filter((value) => value)
            .filter((value) => value !== "Distance:")
            .join("")
    ); // the record to beat

    let viableWays = 1; // the amount of viable ways

    const calcWays = (time: number, record: number) => {
        let ways = 0;

        for (let i = 1; i < time; i++) {
            // i represents the time spent pressing the button up until the time alloted
            const nav = time - i; // represents the amout of time it has after pressing the button to travel
            let dist = nav * i; // represents the distance it travels
            if (dist > record) {
                // if the distance it went is greater than the record, this is viable
                ways += 1;
            }
        }
        return ways;
    };

    // now theres only one time slot
    viableWays *= calcWays(time, record); // multiply the ways together

    return viableWays;
}

// console.log(partTwo(await loadData()));
