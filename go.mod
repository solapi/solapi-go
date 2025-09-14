module github.com/solapi/solapi-go

go 1.18

// From v2 and above, use the module path github.com/solapi/solapi-go/v2
retract [v2.0.0, v999.0.0]

// v2+ is published under github.com/solapi/solapi-go/v2
// This retract ensures v2+ cannot be resolved via the root module path
