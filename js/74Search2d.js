/**
 * @param {number[][]} matrix
 * @param {number} target
 * @return {boolean}
 */
var searchMatrix = function(matrix, target) {
    let m = matrix.length;
    let n = matrix[0].length;
    let start = 0;
    let right = m * n - 1
    while (start <= right) {
        const midPos = start + ((right - start / 2));
        const midValue = matrix[Math.floor(midPos / n)][midPos % n];
        if (midValue === target) {
            return true;
        } else if (midValue > target) {
            right = midPos - 1;
        } else {
            start = midPos + 1;
        }
    }
    return false;
};

module.exports = searchMatrix;