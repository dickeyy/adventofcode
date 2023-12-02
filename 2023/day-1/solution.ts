function unscramble(str:string) {

    // part 1
    // create a new variable -> value
    // loop through the str, if the character is a digit, add it to an array of digits in that string
    // return the first digit in the array and the last

    // part 2
    // we have an array of number digits.
    // now we need to loop through and see if we have string digits (one, two, three...), and translate then add them

    let arr = []; // array of numbers found in the string

    const wordNums = ["one","two","three","four","five","six","seven","eight","nine"] // array of word number values
    const letters = ["o","t","f","s","e","n"] // array of letters that the nums start with

    for (let i=0; i<str.length; i++) { // loop through each character in the string

        if (letters.includes(str[i])) {  // if the character is one of the letters
            for (let x=0; x<wordNums.length; x++) { // then loop through the number words
                if (wordNums[x][0] == str[i]) { // if the number word starts with the character
                    let sub = str.substring(i, i + wordNums[x].length) // create a substring the length of the number word
                    let translated = translate(sub) // attempt to translate the substring to a number
                    if (translated) { // if it translates, then it is good to go 
                        arr.push(translated.toString()) // add it to the array of numbers found
                    }
                }
            }
        } else { // character is not one of the letters for numbers
            if (isDigit(str[i])) { // check if the character is a digit
                arr.push(str[i]) // if it is, push it to the array
            }
        }

    }

    return arr[0] + arr[arr.length - 1] // return the first and last values of the array

}

function isDigit(str:string) {
    const digits = ["0","1","2","3","4","5","6","7","8","9"] // possible digits
    if (digits.includes(str)) {
        return true
    }
    return false
}

function translate(str:string) {
    // take a string number (one, two, three, ...)
    // turn it into a digit and return

    // create a list of all the numbers
    // loop through the list
    // if the string is in the list, return the index + 1

    const numbers = ["one","two","three","four","five","six","seven","eight","nine"] // possible numbers

    if (numbers.includes(str)) { // if the word given is a number word
        return numbers.indexOf(str) + 1 // return the value of the word (index in the array + 1)
    } else {
        return undefined
    }
}

function main(list:string[]) {
    
    let total = 0;

    for (let i=0; i<list.length; i++) {
        total += Number(unscramble(list[i]))
    }

    return total

}

function loadFromFile(filename:string) {
    // loads our data from a file
    let file = Bun.file("./day-1/" + filename)
    return file.text()
}

function toArray(data:string) {
    // turn a long list of text to an array
    return data.split("\n")
}

console.log(main(toArray(await loadFromFile("data.txt")))) // outputs our solution