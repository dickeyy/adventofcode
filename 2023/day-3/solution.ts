// get and format the data
async function loadData() {
    // get the data from the file
    const file = Bun.file('./day-3/data.txt')
    return await file.text()
}

// part one
function partOne(data:string) {

    let total:any = 0
    
    let symbolIndices = [] // array of the x,y position of each symbol
    let numberIndices = [] // array of the x,y position of each number
    let numberIndices2 = []

    let partsIndices: number[][][] = []

    let lines = data.split("\n")

    // find x,y positions of each number and symbol
    for (let y=0; y<lines.length; y++) {

        let pois = lines[y].split("")

        for (let x=0; x<pois.length; x++) {

            if (pois[x].match(/[0-9][0-9]*/)) { // is a number 
                numberIndices.push([x,y])
            } else if (pois[x] !== ".") {
                symbolIndices.push([x,y])
            }

        }

    }

    // now combine the numbers into arrays of coords for each number [ number: [xy],[xy]... ]
    for (let i=0;i<numberIndices.length;i++) {
        let x = numberIndices[i][0]

        if (numberIndices[i+1] !== undefined && numberIndices[i+1][0] == (x+1)) { 
            // the next number is touching me (ew)
            if (numberIndices[i+2] !== undefined && numberIndices[i+2][0] == (x+2)) {
                // there ANOTHER number touching me
                numberIndices2.push([
                    [x,numberIndices[i][1]],
                    [numberIndices[i+1][0], numberIndices[i+1][1]],
                    [numberIndices[i+2][0],numberIndices[i+2][1]]
                ])

                numberIndices.splice(i,2)
            } else {
                // ok only one is touching me
                numberIndices2.push([
                    [x,numberIndices[i][1]],
                    [numberIndices[i+1][0], numberIndices[i+1][1]],
                ])

                numberIndices.splice(i,1)
            }
        } else {
            // no one is touching me
            numberIndices2.push([
                [x,numberIndices[i][1]],
            ])
        }
    }


    // for each number check if in the symbolIndices array, theres is symbol +1x -1x, +1y -1y, or +1y+1x-1x -1y+1x-1x
    for (let ni=0; ni<numberIndices2.length; ni++) { // for each number
        for (let di=0; di<numberIndices2[ni].length; di++) { // for each digit
            let dx = numberIndices2[ni][di][0]
            let dy = numberIndices2[ni][di][1]

            for (let si=0; si<symbolIndices.length; si++) { // check each symbol
                let sx = symbolIndices[si][0]
                let sy = symbolIndices[si][1]

                if (dx+1 == sx && dy == sy || dx-1 == sx && dy == sy) { // look right, look left on the same y
                    partsIndices.push(numberIndices2[ni])
                }

                else if (dy+1 == sy && dx == sx || dy-1 == sy && dx == sx) { // look up and down on the same x
                    partsIndices.push(numberIndices2[ni])
                } 

                else if (dx+1 == sx && dy+1 == sy || dx-1 == sx && dy+1 == sy || dx+1 == sx && dy-1 == sy || dx-1 == sx && dy-1 == sy) { // move up or down, look left or right 
                    partsIndices.push(numberIndices2[ni])
                }

            }
        }
    }

    let uniquePartsIndices = new Set(partsIndices)

    // now get the real numbers for each part
    uniquePartsIndices.forEach((value) => {
        let real = lines[value[0][1]].substring(value[0][0], value[value.length-1][0]+1)
        total += Number(real)
    })

    return total

}

// partOne(await loadData())

// part two 
function partTwo(data:string) {

    let total:any = 0
    
    let symbolIndices = [] // array of the x,y position of each symbol
    let numberIndices = [] // array of the x,y position of each number
    let numberIndices2 = []

    let partsIndices:number[][][] = []
    let correlation = []

    let lines = data.split("\n")

    // find x,y positions of each number and symbol
    for (let y=0; y<lines.length; y++) {

        let pois = lines[y].split("")

        for (let x=0; x<pois.length; x++) {

            if (pois[x].match(/[0-9][0-9]*/)) { // is a number 
                numberIndices.push([x,y])
            } else if (pois[x] !== "." && pois[x] == "*") {
                symbolIndices.push([x,y])
            }

        }

    }

    // now combine the numbers into arrays of coords for each number [ number: [xy],[xy]... ]
    for (let i=0;i<numberIndices.length;i++) {
        let x = numberIndices[i][0]

        if (numberIndices[i+1] !== undefined && numberIndices[i+1][0] == (x+1)) { 
            // the next number is touching me (ew)
            if (numberIndices[i+2] !== undefined && numberIndices[i+2][0] == (x+2)) {
                // there ANOTHER number touching me
                numberIndices2.push([
                    [x,numberIndices[i][1]],
                    [numberIndices[i+1][0], numberIndices[i+1][1]],
                    [numberIndices[i+2][0],numberIndices[i+2][1]]
                ])

                numberIndices.splice(i,2)
            } else {
                // ok only one is touching me
                numberIndices2.push([
                    [x,numberIndices[i][1]],
                    [numberIndices[i+1][0], numberIndices[i+1][1]],
                ])

                numberIndices.splice(i,1)
            }
        } else {
            // no one is touching me
            numberIndices2.push([
                [x,numberIndices[i][1]],
            ])
        }
    }


    // for each number check if in the symbolIndices array, theres is symbol +1x -1x, +1y -1y, or +1y+1x-1x -1y+1x-1x
    for (let ni = 0; ni < numberIndices2.length; ni++) { // for each number
        let numberAdded = false; // Flag to check if the current number has been added
    
        for (let di = 0; di < numberIndices2[ni].length; di++) { // for each digit
            let dx = numberIndices2[ni][di][0];
            let dy = numberIndices2[ni][di][1];
    
            // Check if the number has already been added
            if (numberAdded) {
                break;
            }
    
            for (let si = 0; si < symbolIndices.length; si++) { // check each symbol
                let sx = symbolIndices[si][0];
                let sy = symbolIndices[si][1];
    
                if (
                    (dx + 1 === sx && dy === sy) ||
                    (dx - 1 === sx && dy === sy) ||
                    (dy + 1 === sy && dx === sx) ||
                    (dy - 1 === sy && dx === sx) ||
                    (dx + 1 === sx && dy + 1 === sy) ||
                    (dx - 1 === sx && dy + 1 === sy) ||
                    (dx + 1 === sx && dy - 1 === sy) ||
                    (dx - 1 === sx && dy - 1 === sy)
                ) { // look right, look left, up, down, or diagonally
                    if (!numberAdded) {
                        partsIndices.push([...numberIndices2[ni]]);
                        numberAdded = true;
                    }
                    partsIndices[partsIndices.length - 1].push([si]);
                    break;
                }
            }
        }
    }

    let uniquePartsIndices = new Set(partsIndices)

    // Create a Map to store number arrays based on their index
    const indexMap = new Map();

    // Iterate through the data to populate the Map
    for (const numberArray of uniquePartsIndices) {
        const index = numberArray[numberArray.length - 1][0];
        if (indexMap.has(index)) {
            indexMap.get(index).push(numberArray);
        } else {
            indexMap.set(index, [numberArray]);
        }
    }

    // Filter out number arrays that do not have another number array with the same index
    const filteredData = Array.from(indexMap.values())
    .filter(numberArrays => numberArrays.length > 1)
    
    filteredData.forEach(group => {
        let reals: number[] = []
        group.forEach((value:any) => {
            let real = lines[value[0][1]].substring(value[0][0], value[value.length-2][0]+1)
            reals.push(Number(real))
        })
        total += reals[0] * reals[1]
    });

    return total

}

// partTwo(await loadData())