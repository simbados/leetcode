const {diamondKata, generatePoints, middleLine, printPreviousLetter, letterForIndex,
    printPreviousBlock, printLetterWithPoints, printNextBlock
} = require('./DiamondKata')

/*
    A
 */
/*
    .A.
    B.B
    .A.
 */
/*
    ..A..
    .B.B.
    C...C
    .B.B.
    ..A..
 */
/*
    ...A...
    ..B.B..
    .C...C.
    D.....D
    .C...C.
    ..B.B..
    ...A...
 */
describe('Diamond test', () => {
    test('A to return A', () => {
        expect(diamondKata('A')).toBe('A')
    })

    test('B to return A.B.A', () => {
        expect(diamondKata('B')).toBe('.A.\nB.B\n.A.')
    })

    test('C to return c diamond', () => {
        expect(diamondKata('C')).toBe('..A..\n.B.B.\nC...C\n.B.B.\n..A..')
    })
})

describe('generate points', () => {
    test('Generate Points', () => {
        expect(generatePoints(2)).toBe('..')
    });
});
describe('middle line', () => {
    test('middle line for B', () => {
        expect(middleLine('B')).toEqual('B.B')
    });
    test('middle line for C', () => {
        expect(middleLine('C')).toEqual('C...C')
    })
    test('middle line for D', () => {
        expect(middleLine('D')).toEqual('D.....D')
    })
})

test('letter for index', () => {
    expect(letterForIndex(2)).toEqual('C')
});

describe('print letter with points', () => {
    test('for B', () => {
        expect(printLetterWithPoints('A', 1)).toEqual('.A.')
    })
    test('for C', () => {
        expect(printLetterWithPoints('B', 1)).toEqual('.B.B.')
    })
    test('for D', () => {
        expect(printLetterWithPoints('C', 2)).toEqual('..C...C..')
    })
})

describe('print previous block', () => {
    test('for B', () => {
        expect(printPreviousBlock('B')).toEqual('.A.\n')
    })
    test('for C', () => {
        expect(printPreviousBlock('C')).toEqual('..A..\n.B.B.\n')
    })
    test('for D', () => {
        expect(printPreviousBlock('D')).toEqual('...A...\n..B.B..\n.C...C.\n')
    })
})

describe('print next block', () => {
    test('for B', () => {
        expect(printNextBlock('B')).toEqual('\n.A.')
    })
    test('for C', () => {
        expect(printNextBlock('C')).toEqual('\n.B.B.\n..A..')
    })
    test('for D', () => {
        expect(printNextBlock('D')).toEqual('\n.C...C.\n..B.B..\n...A...')
    })
})
