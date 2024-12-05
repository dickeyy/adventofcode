// get and format the data
async function loadData() {
    // get the data from the file
    const file = Bun.file("../../inputs/2023/day-7/input.txt");
    return await file.text();
}

// part one
function partOne(data: string) {
    const cardVals: any = {
        "2": 2,
        "3": 3,
        "4": 4,
        "5": 5,
        "6": 6,
        "7": 7,
        "8": 8,
        "9": 9,
        T: 10,
        J: 11,
        Q: 12,
        K: 13,
        A: 14
    }; // card values

    const hands = data.split("\n").map((h) => h.split(" ")); // split the data into hands

    const hwr: any = hands.map((p) => {
        // loop through the hands
        let cards: any = {}; // count the cards
        p[0].split("").forEach((c) => {
            // loop through the cards
            !cards[c] ? (cards[c] = 1) : cards[c]++; // count the cards
        });
        let ct = Object.values(cards).sort((a: any, b: any) => b - a); // sort the card counts
        if (ct[0] === 5) {
            // if the hand is a flush
            return [...p, 6]; // return the hand with the rank
        } else if (ct[0] === 4) {
            // if the hand is a four of a kind
            return [...p, 5]; // return the hand with the rank
        } else if (ct[0] === 3 && ct[1] === 2) {
            // if the hand is a full house
            return [...p, 4]; // return the hand with the rank
        } else if (ct[0] === 3 && ct[1] === 1) {
            // if the hand is a three of a kind
            return [...p, 3]; // return the hand with the rank
        } else if (ct[0] === 2 && ct[1] === 2) {
            // if the hand is a two pair
            return [...p, 2]; // return the hand with the rank
        } else if (ct[0] === 2 && ct[1] === 1) {
            // if the hand is a pair
            return [...p, 1]; // return the hand with the rank
        } else {
            // if the hand is a high card
            return [...p, 0]; // return the hand with the rank
        }
    }); // get the hands with ranks

    const sorted = hwr.sort((a: any, b: any) => {
        // sort the hands
        if (a[2] > b[2]) {
            // if the hand is higher
            return 1;
        } else if (a[2] < b[2]) {
            // if the hand is lower
            return -1;
        } else if (a[2] === b[2]) {
            // if the hands are the same
            for (let i = 0; i < a[0].length; i++) {
                // loop through the cards
                if (cardVals[a[0][i]] > cardVals[b[0][i]]) {
                    // if the card is higher
                    return 1;
                } else if (cardVals[a[0][i]] < cardVals[b[0][i]]) {
                    // if the card is lower
                    return -1;
                } else {
                    // if the cards are the same
                    continue;
                }
            }
        }
    }); // sort the hands

    return sorted.reduce((acc: number, card: (string | number)[], i: number) => acc + +card[1] * (i + 1), 0); // return the answer
}

// console.log(partOne(await loadData()));

// part two
function partTwo(data: string) {
    const cardVals: any = {
        // card values
        "2": 2,
        "3": 3,
        "4": 4,
        "5": 5,
        "6": 6,
        "7": 7,
        "8": 8,
        "9": 9,
        T: 10,
        J: 1,
        Q: 12,
        K: 13,
        A: 14
    };

    const hands = data.split("\n").map((h) => h.split(" ")); // split the data into hands

    const hwr: any = hands.map((p: any) => {
        // get the hands with ranks
        // loop through the hands
        let cards: any = {}; // count the cards
        let jokers: any = 0; // count the jokers

        p[0].split("").forEach((c: string) => {
            // loop through the cards
            if (c !== "J") {
                // if the card is not a joker
                !cards[c] ? (cards[c] = 1) : cards[c]++; // count the cards
            } else jokers++; // count the jokers
        });

        let ct = Object.values(cards).sort((a: any, b: any) => b - a); // sort the card counts
        ct[0] ? (ct[0] += jokers) : (ct[0] = jokers); // add jokers to the highest card count

        if (ct[0] === 5) {
            // if the hand is a flush
            return [...p, 6]; // return the hand with the rank
        } else if (ct[0] === 4) {
            // if the hand is a four of a kind
            return [...p, 5]; // return the hand with the rank
        } else if (ct[0] === 3 && ct[1] === 2) {
            // if the hand is a full house
            return [...p, 4]; // return the hand with the rank
        } else if (ct[0] === 3 && ct[1] === 1) {
            // if the hand is a three of a kind
            return [...p, 3]; // return the hand with the rank
        } else if (ct[0] === 2 && ct[1] === 2) {
            // if the hand is a two pair
            return [...p, 2]; // return the hand with the rank
        } else if (ct[0] === 2 && ct[1] === 1) {
            // if the hand is a pair
            return [...p, 1]; // return the hand with the rank
        } else {
            // if the hand is a high card
            return [...p, 0]; // return the hand with the rank
        }
    }); // get the hands with ranks

    const sorted = hwr.sort((a: any, b: any) => {
        // sort the hands
        if (a[2] > b[2]) {
            // if the hand is higher
            return 1;
        } else if (a[2] < b[2]) {
            // if the hand is lower
            return -1;
        } else if (a[2] === b[2]) {
            // if the hands are the same
            for (let i = 0; i < a[0].length; i++) {
                // loop through the cards
                if (cardVals[a[0][i]] > cardVals[b[0][i]]) {
                    // if the card is higher
                    return 1;
                } else if (cardVals[a[0][i]] < cardVals[b[0][i]]) {
                    // if the card is lower
                    return -1;
                } else {
                    // if the cards are the same
                    continue;
                }
            }
        }
    }); // sort the hands

    return sorted.reduce((acc: number, card: (string | number)[], i: number) => acc + +card[1] * (i + 1), 0); // return the answer
}

// console.log(partTwo(await loadData()));
