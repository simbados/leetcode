var removeDuplicates = function (nums) {
    let count = nums.length;
    if (nums.length < 3) {
        return nums.length;
    }
    for (let i = 2; i < nums.length; i++) {
        if (nums[i - 2] === nums[i - 1] && nums[i - 2] === nums[i]) {
            nums[i - 2] = Math.pow(10, 5);
            count--;
        }
    }
    nums.sort((a, b) => a < b ? -1 : 1);
    return count;
};

module.exports = removeDuplicates;