package core

import (
	"strings"
	"strconv"
	"github.com/icrowley/fake"
)


// Array function argument catcher
var ArrayArgs = make(map[string]interface{})


// Random array generator for array datatypes
func ArrayGenerator(dt string) (string, error) {

	// Getting the value of itertors
	maxValues,_ := RandomInt(0, 6)
	maxIteration, _ := RandomInt(0, 3)

	// Collectors
	var value interface{}
	var resultArrayCollector []string

	for i := 0; i < maxIteration; i++ { // Max number of arrays
		var resultArray []string
		for j := 0; j < maxValues; j++ { // max number of values in a array.

			// Call appropriate function to generate a array
			switch dt {

				case "string": // strings
					value = RandomString(ArrayArgs["strlen"].(int))

				case "int":    // int
					intvalue, err := RandomInt(ArrayArgs["intmin"].(int), ArrayArgs["intmax"].(int))
					if err != nil {
						return "", err
					}
					value = strconv.Itoa(intvalue)

				case "float":  // float
					floatvalue, err := RandomFloat(ArrayArgs["floatmin"].(int), ArrayArgs["floatmax"].(int), ArrayArgs["floatprecision"].(int))
					if err != nil {
						return "", err
					}
					value = TruncateFloat(floatvalue, ArrayArgs["floatmax"].(int), ArrayArgs["floatprecision"].(int))
					value = strconv.FormatFloat(value.(float64), 'f', ArrayArgs["floatprecision"].(int), 64)

				case "bit":    // bit
					value = RandomBit(ArrayArgs["bitlen"].(int))

				case "text": // text
					value = fake.WordsN(1)

				case "date": // date
					dvalue, err := RandomDate(ArrayArgs["fromyear"].(int), ArrayArgs["toyear"].(int))
					if err != nil {
						return "", err
					}
					value = dvalue

				case "time": // timestamp
					tvalue, err := RandomTime(ArrayArgs["fromyear"].(int), ArrayArgs["toyear"].(int))
					if err != nil {
						return "", err
					}
					value = tvalue

				case "timetz": // timestamp
					ttzvalue, err := RandomTimetz(ArrayArgs["fromyear"].(int), ArrayArgs["toyear"].(int))
					if err != nil {
						return "", err
					}
					value = ttzvalue

				case "timestamp": // timestamp
					tsvalue, err := RandomTimestamp(ArrayArgs["fromyear"].(int), ArrayArgs["toyear"].(int))
					if err != nil {
						return "", err
					}
					value = tsvalue

				case "timestamptz": // timestamp
					tstzvalue, err := RandomTimestamptz(ArrayArgs["fromyear"].(int), ArrayArgs["toyear"].(int))
					if err != nil {
						return "", err
					}
					value = tstzvalue

				case "bool": // bool
					if RandomBoolean() {
						value = "true"
					} else {
						value = "false"
					}

				case "IP": // IP Address
					value = RandomIP()

				case "macaddr": // Mac Address
					value = RandomMacAddress()

				case "uuid": // UUID
					uvalue, err := RandomUUID()
					if err != nil {
						return "", err
					}
					value = uvalue

				case "txid_snapshot": // txid snapshot
					value = RandomTXID()

				case "pg_lsn": // pg lsn
					value = RandomLSN()

				case "tsquery": // TS Query
					value = RandomTSQuery()

				case "tsvector": // TS Vector
					value = RandomTSVector()

			}
			resultArray = append(resultArray, value.(string))
		}
		resultArrayCollector = append(resultArrayCollector, "{" + strings.Join(resultArray, ",") + "}")
	}
	return "{" + strings.Join(resultArrayCollector, ",") + "}", nil
}

// Random geometric array generators
func GeometricArrayGenerator(maxInt int, geometryType string) string {

	// Getting the value of iterators
	maxIterations,_ := RandomInt(0, 6)
	var resultArray []string

	if geometryType == "box" {
		value := RandomGeometricData(maxInt, geometryType, false)
		resultArray = append(resultArray, value)
	} else {
		for i := 0; i < maxIterations; i++ { // Max number of arrays
			value := RandomGeometricData(maxInt, geometryType, true)
			resultArray = append(resultArray, value)
		}
	}

	return "{" + strings.Join(resultArray, ",") + "}"
}

// Random XML & Json array generators.
func JsonXmlArrayGenerator(dt string) string {
	// Getting the value of iterators
	maxIterations,_ := RandomInt(0, 6)
	var resultArray []string
	var value string
	for i := 0; i < maxIterations; i++ { // Max number of arrays

		switch dt { // Choose the appropriate random data generators
			case "json":
				value = "\"" + RandomJson(true) + "\""
			case "xml":
				value = "\"" + RandomXML(true) + "\""
		}

		resultArray = append(resultArray, value)
	}
	return "{" + strings.Join(resultArray, ",")  + "}"
}