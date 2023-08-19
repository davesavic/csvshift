grammar CsvShiftGrammar;

// Lexer rules
INPUT: 'Input Columns';
OUTPUT: 'Output Columns';
COLUMN: 'Column';
COLUMNS: 'Columns';

// Transformers
TRIM: '-> Trim';
REPLACE: '-> Replace';
JOIN: '-> Join with';
LOWER: '-> ToLower';
UPPER: '-> ToUpper';
SPLIT: '-> Split on';
REGEXREPLACE: '-> RegexReplace';

// Symbols
WITH: 'with';
AS: 'as';
COMMA: ',';
QUOTE: '"';
ARROW: '->';

IDENTIFIER: [a-zA-Z_][a-zA-Z_0-9]*;
STRING: QUOTE ~[\r\n"]* QUOTE;
WS: [ \t\r\n]+ -> skip;

// Parser rules
csvTransform: inputSection columnModifierSection* outputSection EOF;

inputSection: INPUT columns;
outputSection: OUTPUT columns;

columnModifierSection: singleColumnModifierSection | multipleColumnModifierSection;

singleColumnModifierSection: COLUMN IDENTIFIER singleColumnTransformation+;

multipleColumnModifierSection: COLUMNS columns multipleColumnTransformation+;

singleColumnTransformation: TRIM
                | REPLACE STRING WITH STRING
                | LOWER
                | UPPER
                | SPLIT STRING AS columns
                | REGEXREPLACE STRING WITH STRING
                ;

multipleColumnTransformation: TRIM
                | REPLACE STRING WITH STRING
                | JOIN STRING AS IDENTIFIER
                | LOWER
                | UPPER
                | REGEXREPLACE STRING WITH STRING
                ;

columns: IDENTIFIER (COMMA IDENTIFIER)*;
