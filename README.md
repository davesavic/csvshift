# CSVShift ![Tests](https://github.com/davesavic/csvshift/actions/workflows/go-tests.yml/badge.svg) ![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/davesavic/csvshift)
A powerful DSL tailored for seamless CSV file manipulation. With intuitive syntax, it handles everything from data ingestion to transformation, refinement, and final output. Elevate your CSV data processing with unparalleled ease and precision.

## Usage
### Installation
Download the latest [release](https://github.com/davesavic/csvshift/releases) for your operating system.

```bash
./csvshift --help
```

### Example
#### Command
```bash
csvshift --source large-file.csv --destination transformed-file.csv --path script.csvshift
```

#### Script.csvshift
```
Input Columns id, first_name, last_name, address, dob

Column last_name
-> Trim
-> Replace "S" with "K"

Columns first_name, last_name
-> Replace "K" with "L"
-> Join with " " as full_name

Column address
-> Split on "," as street, city, state

Columns street, city, state
-> Trim

Column state
-> Split on " " as state, postcode, country

Column dob
-> RegexReplace "(\d{4})-(\d{2})-(\d{2})" with "$3/$2/$1"

Column dob
-> RegexExtract "(\d{2})\/(\d{2})\/(\d{4})" as day, month, year

Output Columns id, full_name, street, city, state, postcode, country, dob, day, month, year
```

#### Source.csv
```csv
first_name,last_name,email,phone_number
John,  Doe ,JOHNDOE@company.com,1234567890
Jane,Smith ,JaNeSmItH@company.com,0987654321
```

#### Result.csv
```csv
full_name,handle,domain,phone_number
John Doe,johndoe,company.com,1234567890
Jane Smith,janesmith,company.com,0987654321
```

### Syntax
#### Input
```csv
Input Columns ... # list of columns to be used as input (required)
```

#### Output
```csv
Output Columns ... # list of columns to be used as output (required)
```

#### Single Column Transformation
```csv
Column ... # column to be transformed
-> ... # chained transformation to be applied
```

#### Multiple Column Transformation
```csv
Columns ... # list of columns to be transformed
-> ... # chained transformation to be applied
```

#### Available Transformations
```csv
-> ToLower # converts string to lowercase
-> ToUpper # converts string to uppercase
-> Trim # trims whitespace from string
-> Join with {separator} as {column_name} # combines columns into a single column with specified separator
-> Split on {separator} as {column_names}
-> Replace {old} with {new} # replaces old string with new string
-> RegexReplace {pattern} with {new} # replaces regex pattern with new string
-> RegexExtract {pattern} into {column_names} # extracts regex pattern into new columns if pattern matches
```

#### More Features To Come