var majorityElement = function(nums) {
    if (nums.length === 1) {
        return nums[0]
    }
    const mapping = new Map();
    for (const num of nums) {
        mapping.set(num, mapping.get(num) != null ? mapping.get(num) + 1 : 1)
    }
    let majority = 0;
    let previousVal = 0;
    for (const [index, val] of mapping) {
        if (val > previousVal) {
            previousVal = val;
            majority = index;
        }
    }
    return majority;
};

module.exports = majorityElement;