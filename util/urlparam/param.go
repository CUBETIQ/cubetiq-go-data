package param

import (
	"regexp"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Param struct {
	Page  int64  `query:"page"`
	Paged bool   `query:"paged"`
	Q     string `query:"q"`
	Size  int64  `query:"size"`
	Sort  string `query:"sort"`
}

func (p *Param) GetDefaultParam() *Param {
	return &Param{
		Page:  0,
		Paged: false,
		Q:     "",
		Size:  20,
		Sort:  "_id,desc",
	}
}

type SortByKeyValue struct {
	Key   string
	Value int
}

func GetSortBy(s string) []SortByKeyValue {
	// regex pattern
	regexPattern := `[A-Za-z0-9\,\;\_]+$`
	match, _ := regexp.MatchString(regexPattern, s)

	sortByKeyValue := make([]SortByKeyValue, 0)

	// if string match regex pattern
	if match {
		// split into array by ';'
		data := strings.Split(s, ";")
		// loop after splitting
		for _, s := range data {
			// check if string containes ','
			if strings.Contains(s, ",") {
				// split to get key and value
				_keyValue := strings.Split(s, ",")
				_key := _keyValue[0]
				_v := _keyValue[1]
				// convert value to lowercase
				_value := strings.ToLower(_v)
				// define new variable to get -1 or 1
				var resultValue int

				// asc = 1
				// desc = -1
				if _value == "asc" {
					resultValue = 1
				} else {
					resultValue = -1
				}

				// append a slice
				sortByKeyValue = append(sortByKeyValue, SortByKeyValue{_key, resultValue})
			}
		}
	} else {
		// append default value
		sortByKeyValue = append(sortByKeyValue, SortByKeyValue{"_id", -1})
	}

	return sortByKeyValue
}

func GetParam(p Param, q []string) (_filter primitive.D, _options *options.FindOptions) {
	var filter bson.D
	for _, query := range q {
		filter = append(filter, bson.E{Key: query, Value: primitive.Regex{Pattern: p.Q, Options: "i"}})
	}
	findOption := options.Find()
	findOption.SetLimit(p.Size)
	if p.Page > 0 {
		findOption.SetSkip(p.Page * p.Size)
	} else {
		findOption.SetSkip(0)
	}
	// get sort by
	keyValue := GetSortBy(p.Sort)
	var sortOpt bson.D
	for _, kv := range keyValue {
		sortOpt = append(sortOpt, bson.E{Key: kv.Key, Value: kv.Value})
	}
	findOption.SetSort(sortOpt)

	_filter = filter
	_options = findOption

	return
}
