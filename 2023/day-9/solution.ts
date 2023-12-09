// get and format the data
async function loadData() {
    // get the data from the file
    const file = Bun.file("./data.txt");
    return await file.text();
}

// part one
function partOne(data: string) {
    return 0;
}

console.log(partOne(await loadData()));
