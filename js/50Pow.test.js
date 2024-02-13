const myPow = require('./50Pow')

test('pow test', () => {
    expect(myPow(2.0, 10)).toBe(1024)
    expect(myPow(2.1, 3)).toBe(9.261)
    expect(myPow(2.0, -2)).toBe(0.25)
    expect(myPow(8.88023, 3)).toBe(700.28148)
})
