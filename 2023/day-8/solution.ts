// get and format the data
async function loadData() {
    // get the data from the file
    const file = Bun.file("./data.txt");
    return await file.text();
}

// part one
function partOne(data: string) {
    const inputArray = data.split("\n"); // split the data into an array
    const instructionsArray = inputArray[0].split(""); // split the instructions into an array

    const nodesMap: any = inputArray.slice(2).reduce((accumulator, current) => {
        // create a map of the nodes
        const [nodeName, nodeCoords] = current.replaceAll(" ", "").split("="); // split the node name and coords
        accumulator = { ...accumulator, [nodeName]: nodeCoords.slice(1, 8).split(",") }; // add the node to the map
        return accumulator;
    }, {});

    let totalSteps = 0; // total steps taken
    let instructionIndex = 0; // current instruction index
    let currentNode = "AAA"; // current node

    while (currentNode !== "ZZZ") {
        // while the current node is not the end node
        if (instructionIndex === instructionsArray.length) instructionIndex = 0; // if the instruction index is at the end, reset it
        const currentInstruction = instructionsArray[instructionIndex]; // get the current instruction
        currentNode = currentInstruction === "L" ? nodesMap[currentNode][0] : nodesMap[currentNode][1]; // get the next node
        totalSteps++; // increment the total steps
        instructionIndex++; // increment the instruction index
    }

    return totalSteps;
}

// console.log(partOne(await loadData()));

// part two
function partTwo(data: string) {
    const inputArray = data.split("\n"); // split the data into an array
    const instructionsArray = inputArray[0].split(""); // split the instructions into an array

    const nodesMap: any = inputArray.slice(2).reduce((accumulator, current) => {
        // create a map of the nodes
        const [nodeName, nodeCoords] = current.replaceAll(" ", "").split("="); // split the node name and coords
        accumulator = { ...accumulator, [nodeName]: nodeCoords.slice(1, 8).split(",") }; // add the node to the map
        return accumulator;
    }, {});

    const endingWithA = Object.keys(nodesMap).filter((nodeName) => nodeName.slice(-1) === "A"); // get all the nodes that end with A

    const endingWithALoops = endingWithA.reduce((accumulator, currentNode) => {
        // get the total steps for each node
        let totalSteps = 0; // total steps taken
        let counterIndex = 0; // current instruction index
        let currentElement = currentNode; // current node

        while (currentElement.slice(-1) !== "Z") {
            // while the current node is not the end node
            if (counterIndex === instructionsArray.length) counterIndex = 0; // if the instruction index is at the end, reset it
            const currentInstruction = instructionsArray[counterIndex]; // get the current instruction
            currentElement = currentInstruction === "L" ? nodesMap[currentElement][0] : nodesMap[currentElement][1]; // get the next node
            totalSteps++; // increment the total steps
            counterIndex++; // increment the instruction index
        }

        accumulator = { ...accumulator, [currentNode]: totalSteps }; // add the node to the map
        return accumulator;
    }, {});

    const calculateLCM = (num1: number, num2: number) => {
        // calculate the LCM of two numbers
        const findGCD: any = (x: number, y: number) => (y === 0 ? x : findGCD(y, x % y)); // find the GCD of two numbers
        return (num1 * num2) / findGCD(num1, num2); // calculate the LCM
    };

    const calculateLCMArray = (arrayOfNumbers: string | any[]) => {
        // calculate the LCM of an array of numbers
        let lcmValue = arrayOfNumbers[0]; // set the initial LCM value
        for (let i = 1; i < arrayOfNumbers.length; i++) {
            // loop through the array of numbers
            lcmValue = calculateLCM(lcmValue, arrayOfNumbers[i]); // calculate the LCM of the current number and the LCM value
        }
        return lcmValue;
    };

    return calculateLCMArray(Object.values(endingWithALoops));
}

// console.log(partTwo(await loadData()));
