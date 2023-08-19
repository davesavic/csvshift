# CSVShift
A powerful DSL tailored for seamless CSV file manipulation. With intuitive syntax, it handles everything from data ingestion to transformation, refinement, and final output. Elevate your CSV data processing with unparalleled ease and precision.

## Usage
### Installation
```bash
git clone git@github.com:davesavic/csvshift.git
cd csvshift
make build # builds binary depending on your OS to ./bin/csvshift
./bin/csvshift --help
```

### Example
#### Command
```bash
csvshift --source large-file.csv --destination transformed-file.csv --path script.csvshift
```

#### Script.csvshift
```
Input Columns first_name, last_name, email, phone_number

Column email
-> ToLower

Columns first_name, last_name
-> Trim
-> Join with " " as full_name

Column email
-> Split on "@" as handle, domain

Column domain
-> Replace "@" with ""

Output Columns full_name, handle, domain, phone_number
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
```

#### More Features To Come