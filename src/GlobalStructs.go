package src

// Global structs

// payload
type Payload struct {
    Coldef interface{}

    ColName string

    Values []interface{}
    Err error
}

// columns defintion struct
type ColumnsDef struct {
    Type string

    Limit int
}

// the cofing struct
type ConfigStruct struct {
    TableName string

    Size int

    ColsStruct map[string]ColumnsDef
}

// the tbale struct
type Table struct {
    Config *ConfigStruct

    Cols map[string] []interface{}
}
