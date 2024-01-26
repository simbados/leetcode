const majorityElement = require("./169MajorityElement");

test('Majority Element', () => {
    expect(majorityElement([1, 2, 1, 2, 1])).toBe(1);
    expect(majorityElement([2,3,4,5,5,5])).toBe(5);
})
