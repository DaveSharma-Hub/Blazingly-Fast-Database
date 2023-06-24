package typeDeclaration

type databaseInput struct{
	InputCommandType string `json:"inputcommandtype"`
	Key string `json:"key"`
	TableName string `json:"tablename"`
	Data string `json:"data"`
}
// inputcommand eg. GET <header/key value> FROM <table name>
// INSERT <key value> INTO <table name> DATA

type databaseOutput struct{
	Output string `json:"output"`
	Data string `json:"data"`
}

type tableHeader struct{
	tableHeaderName string
	tableHeaderType string
}

var tmpData = []databaseOutput{
	{Output: "Blue Train", Data: "John Coltrane"},
	{Output: "Blue Train", Data: "John Coltrane"},
}