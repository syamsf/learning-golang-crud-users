package helper

func PanicIfError(err error, errorType string) {
  if err != nil {
    if errorType == "database" {
      panic("Could not connect to the database")
    }

    panic(err)
  }
}

func PanicIfErrorDefault(err error) {
  PanicIfError(err, "default")
}
