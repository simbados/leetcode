const removeDuplicates = require('./80RemoveDuplicatesII');

test('Remove duplicate', () => {
    const numbers1 = [1, 1, 1, 2, 2, 3];
    const numbers2 = [0, 0, 1, 1, 1, 1, 2, 3, 3];
    expect(removeDuplicates(numbers1)).toEqual(5);
    expect(numbers1).toEqual([1, 1, 2, 2, 3, Math.pow(10, 5)]);
    expect(removeDuplicates(numbers2)).toEqual(7);
    expect(numbers2).toEqual([0, 0, 1, 1, 2, 3, 3, Math.pow(10, 5), Math.pow(10, 5)]);
})