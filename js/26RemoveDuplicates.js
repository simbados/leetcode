const removeDuplicates = function(nums) {
    let lastVal = null;
    let unique = nums.length;
    for (const [index, val] of nums.entries()) {
        if (lastVal != null && val === lastVal) {
            nums[index] = 101;
            unique--;
        } else {
            lastVal = val;
        }
    }
    nums.sort((a, b) => a < b ? -1 : 1);
    return unique;
};
module.exports = removeDuplicates;
