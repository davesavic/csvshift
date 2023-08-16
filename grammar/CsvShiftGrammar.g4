grammar CsvShiftGrammar;

// Lexer rules
INPUT: 'Input Columns';
OUTPUT: 'Output Columns';
COLUMN: 'Column';
COLUMNS: 'Columns';

// Transformers
TRIM: '-> Trim';
REPLACE: '-> Replace';
COMBINE: '-> Combine with';

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
              ;

multipleColumnTransformation: TRIM
              | REPLACE STRING WITH STRING
              | COMBINE STRING AS IDENTIFIER
              ;

columns: IDENTIFIER (COMMA IDENTIFIER)*;
