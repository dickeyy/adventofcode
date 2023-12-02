// get and format the data
async function loadData() {

    // get the data from the file
    const file = Bun.file('./day-2/data.txt')
    return await file.text()

}

// part one
function partOne(data:string) {

    // set our possible game
    const possibleGame:any = { red: 12, green: 13, blue: 14 }


    // initiate a total variable
    let total = 0;

    // split the data by linebreak
    data.split("\n").forEach((line) => { // for each line

        let isPossible = true // default to this game being possible

        let split = line.split(":").map(item => item.trim()) // get the gameInfo and gameData
        let gameInfo = split[0] // left side of the :
        let gameData = split[1] // right side of the :

        let gameId = gameInfo.split(" ")[1] // get the game id by splitting "Game ID" (right side of the space)
        let gameSets = gameData.split(";").map((item) => item.trim()) // get each set (each set is separated by ;)

        gameSets.forEach((item) => { // for each set (ex: 1 red, 3 blue, 4 green)
            let cubes = item.split(",").map((item2) => item2.trim()) // get each individual cube (ex: 1 red)
            
            cubes.forEach((cube) => { // for each of those cubes
                let c = cube.split(" ").map((item3) => item3.trim()) // split by the space
                let color = c[1] // the color is on the right
                let count = c[0] // the count is on the left

                if (Number(count) > possibleGame[color]) { // determine if the count for that color is greater than what is possible
                    isPossible = false // if so set to false
                }
            })

        })

        if (isPossible) { // if nothing made it impossible
            total += Number(gameId) // add the gameid to the total sum
        }

    })

    return total

}

// partOne(await (loadData()))

// part two
function partTwo(data:string) {

    // initiate a total variable
    let total = 0;

    // split the data by linebreak
    data.split("\n").forEach((line) => { // for each line

        let highestValues:any = [0,0,0]

        let split = line.split(":").map(item => item.trim()) // get the gameInfo and gameData
        let gameData = split[1] // right side of the :

        let gameSets = gameData.split(";").map((item) => item.trim()) // get each set (each set is separated by ;)

        gameSets.forEach((item) => { // for each set (ex: 1 red, 3 blue, 4 green)
            let cubes = item.split(",").map((item2) => item2.trim()) // get each individual cube (ex: 1 red)
            
            cubes.forEach((cube) => { // for each of those cubes
                let c = cube.split(" ").map((item3) => item3.trim()) // split by the space
                let color = c[1] // the color is on the right
                let count = c[0] // the count is on the left

                if (color == "red") {
                    if (highestValues[0] < Number(count)) {
                        highestValues[0] = Number(count)
                    }
                } else if (color == "green") {
                    if (highestValues[1] < Number(count)) {
                        highestValues[1] = Number(count)
                    }
                } else if (color == "blue") {
                    if (highestValues[2] < Number(count)) {
                        highestValues[2] = Number(count)
                    }
                }
            })

        })

        let power = highestValues[0] * highestValues[1] * highestValues[2]
        total += power

    })

    return total

}

// console.log(partTwo(await loadData()))