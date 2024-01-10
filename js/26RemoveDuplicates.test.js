const removeDuplicates = require('./26RemoveDuplicates');

test('Remove duplicate', () => {
    const numbers1 = [1, 2, 2, 3, 4, 5, 5];
    const numbers2 = [1];
    const numbers3 = [1, 1];
    const numbers4 = [0, 0, 1, 2, 3, 4, 5];
    expect(removeDuplicates(numbers1)).toEqual(5);
    expect(numbers1).toEqual([1, 2, 3, 4, 5, 101, 101]);
    expect(removeDuplicates(numbers2)).toEqual(1);
    expect(numbers2).toEqual([1]);
    expect(removeDuplicates(numbers3)).toEqual(1);
    expect(numbers3).toEqual([1, 101]);
    expect(removeDuplicates(numbers4)).toEqual(6);
    expect(numbers4).toEqual([0, 1, 2, 3, 4, 5, 101]);
})