package src

// Global structs
type Payload struct {
    Targetdb string
    RowName string
    Set string

    Args []string
}


type DBConfig struct {
    Name string
    Type string
}

type DBStruct {
    Config *DBConfig

}
