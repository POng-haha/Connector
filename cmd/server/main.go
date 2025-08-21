{
  "destinations": {
    "dopa": {
      "type": "https",
      "host": "https://idcard.bora.dopa.go.th",
      "timeout": 20
    }
  },
  "routes": {
    "POST:/api/Customer/CheckDOPA": {
      "destinationName": "dopa",
      "path": "/CheckCardStatus/CheckCardService.asmx",
      "apiKey": ""
    }
  }
}
