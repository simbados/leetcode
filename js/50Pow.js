/**
 * @param {number} x
 * @param {number} n
 * @return {number}
 */
var myPow = function(x, n) {
    if (n < 0) {
        n = -n;
        x = 1 / x
    }
    return toPrecision(helper(x, n))
};

function helper(x, n) {
    if (n === 0) {
        return 1;
    }
    if (n % 2 === 0) {
        let square = helper(x, n / 2)
        return square * square
    } else {
        return x * helper(x, n -1)
    }
}

function toPrecision(value) {
    const factor = Math.pow(10, 5)
    return Math.round(value * factor) / factor
}

module.exports = myPow;