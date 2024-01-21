const canConstruct = require('./383RansomNote');

test('Remove duplicate', () => {
    expect(canConstruct("a", "ab")).toEqual(true);
    expect(canConstruct("ac", "ab")).toEqual(false);
})