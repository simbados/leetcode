var diamondKata = function (letter) {
    switch (letter) {
        case 'A':
            return 'A';
        default:
            return printPreviousBlock(letter) + middleLine(letter) + printNextBlock(letter);
    }
}

var printNextBlock = function(letter) {
    return printPreviousBlock(letter).split('').reverse().join('')
}

var printPreviousBlock = function (letter) {
    let previousBlock = ''
    for (let i = indexOfLetter(letter); i > 0; i--) {
        let index = indexOfLetter(letter) - i;
        previousBlock += printLetterWithPoints(letterForIndex(index), i) + '\n'
    }
    return previousBlock
}

var printLetterWithPoints = function (letter, points) {
    if (letter === 'A') {
        return generatePoints(points) + letterForIndex(indexOfLetter(letter)) + generatePoints(points)
    } else {
        return generatePoints(points) + middleLine(letterForIndex(indexOfLetter(letter))) + generatePoints(points)
    }
}

var middleLine = function (letter) {
    return letter + generatePoints(1 + (2 * (indexOfLetter(letter) - 1))) + letter
}

const alphabet = 'ABCDEFGHIJKLMNOPQRSTUVWXYZ';
var indexOfLetter = function (letter) {
    return alphabet.indexOf(letter)
}

var letterForIndex = function (index) {
    return alphabet[index]
}

function generatePoints(n) {
    return ".".repeat(n)
}

module.exports = {diamondKata, generatePoints, middleLine, printLetterWithPoints,printPreviousBlock, letterForIndex, printNextBlock}