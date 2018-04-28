//go:generate go install github.com/go-fed/activity/tools/vocab
//go:generate vocab
package vocab

import (
	"encoding/json"
	"github.com/go-test/deep"
	"net/url"
	"testing"
)

func GetJSONDiff(str1, str2 []byte) ([]string, error) {
	var i1 interface{}
	var i2 interface{}

	err := json.Unmarshal(str1, &i1)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(str2, &i2)
	if err != nil {
		return nil, err
	}
	return deep.Equal(i1, i2), nil
}

func MustParseURL(s string) url.URL {
	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}
	return *u
}

var (
	tables = []struct {
		name           string
		expectedJSON   string
		expectedStruct Serializer
		deserializer   func() Deserializer
	}{
		{
			name:           "Example 1",
			expectedJSON:   example1,
			expectedStruct: example1Type,
			deserializer:   func() Deserializer { return &Object{} },
		},
		{
			name:           "Example 2",
			expectedJSON:   example2,
			expectedStruct: example2Type,
			deserializer:   func() Deserializer { return &Link{} },
		},
		{
			name:           "Example 3",
			expectedJSON:   example3,
			expectedStruct: example3Type,
			deserializer:   func() Deserializer { return &Activity{} },
		},
		{
			name:           "Example 4",
			expectedJSON:   example4,
			expectedStruct: example4Type,
			deserializer:   func() Deserializer { return &Travel{} },
		},
		{
			name:           "Example 5",
			expectedJSON:   example5,
			expectedStruct: example5Type,
			deserializer:   func() Deserializer { return &Collection{} },
		},
		{
			name:           "Example 6",
			expectedJSON:   example6,
			expectedStruct: example6Type,
			deserializer:   func() Deserializer { return &OrderedCollection{} },
		},
		{
			name:           "Example 7",
			expectedJSON:   example7,
			expectedStruct: example7Type,
			deserializer:   func() Deserializer { return &CollectionPage{} },
		},
		{
			name:           "Example 8",
			expectedJSON:   example8,
			expectedStruct: example8Type,
			deserializer:   func() Deserializer { return &OrderedCollectionPage{} },
		},
		{
			name:           "Example 9",
			expectedJSON:   example9,
			expectedStruct: example9Type,
			deserializer:   func() Deserializer { return &Accept{} },
		},
		{
			name:           "Example 10",
			expectedJSON:   example10,
			expectedStruct: example10Type,
			deserializer:   func() Deserializer { return &Accept{} },
		},
		{
			name:           "Example 11",
			expectedJSON:   example11,
			expectedStruct: example11Type,
			deserializer:   func() Deserializer { return &TentativeAccept{} },
		},
		{
			name:           "Example 12",
			expectedJSON:   example12,
			expectedStruct: example12Type,
			deserializer:   func() Deserializer { return &Add{} },
		},
		{
			name:           "Example 13",
			expectedJSON:   example13,
			expectedStruct: example13Type,
			deserializer:   func() Deserializer { return &Add{} },
		},
		{
			name:           "Example 14",
			expectedJSON:   example14,
			expectedStruct: example14Type,
			deserializer:   func() Deserializer { return &Arrive{} },
		},
		{
			name:           "Example 15",
			expectedJSON:   example15,
			expectedStruct: example15Type,
			deserializer:   func() Deserializer { return &Create{} },
		},
		{
			name:           "Example 16",
			expectedJSON:   example16,
			expectedStruct: example16Type,
			deserializer:   func() Deserializer { return &Delete{} },
		},
		{
			name:           "Example 17",
			expectedJSON:   example17,
			expectedStruct: example17Type,
			deserializer:   func() Deserializer { return &Follow{} },
		},
		{
			name:           "Example 18",
			expectedJSON:   example18,
			expectedStruct: example18Type,
			deserializer:   func() Deserializer { return &Ignore{} },
		},
		{
			name:           "Example 19",
			expectedJSON:   example19,
			expectedStruct: example19Type,
			deserializer:   func() Deserializer { return &Join{} },
		},
		{
			name:           "Example 20",
			expectedJSON:   example20,
			expectedStruct: example20Type,
			deserializer:   func() Deserializer { return &Leave{} },
		},
		{
			name:           "Example 21",
			expectedJSON:   example21,
			expectedStruct: example21Type,
			deserializer:   func() Deserializer { return &Leave{} },
		},
		{
			name:           "Example 22",
			expectedJSON:   example22,
			expectedStruct: example22Type,
			deserializer:   func() Deserializer { return &Like{} },
		},
		{
			name:           "Example 23",
			expectedJSON:   example23,
			expectedStruct: example23Type,
			deserializer:   func() Deserializer { return &Offer{} },
		},
		{
			name:           "Example 24",
			expectedJSON:   example24,
			expectedStruct: example24Type,
			deserializer:   func() Deserializer { return &Invite{} },
		},
		{
			name:           "Example 25",
			expectedJSON:   example25,
			expectedStruct: example25Type,
			deserializer:   func() Deserializer { return &Reject{} },
		},
		{
			name:           "Example 26",
			expectedJSON:   example26,
			expectedStruct: example26Type,
			deserializer:   func() Deserializer { return &TentativeReject{} },
		},
		{
			name:           "Example 27",
			expectedJSON:   example27,
			expectedStruct: example27Type,
			deserializer:   func() Deserializer { return &Remove{} },
		},
		{
			name:           "Example 28",
			expectedJSON:   example28,
			expectedStruct: example28Type,
			deserializer:   func() Deserializer { return &Remove{} },
		},
		{
			name:           "Example 29",
			expectedJSON:   example29,
			expectedStruct: example29Type,
			deserializer:   func() Deserializer { return &Undo{} },
		},
		{
			name:           "Example 30",
			expectedJSON:   example30,
			expectedStruct: example30Type,
			deserializer:   func() Deserializer { return &Update{} },
		},
		{
			name:           "Example 31",
			expectedJSON:   example31,
			expectedStruct: example31Type,
			deserializer:   func() Deserializer { return &View{} },
		},
		{
			name:           "Example 32",
			expectedJSON:   example32,
			expectedStruct: example32Type,
			deserializer:   func() Deserializer { return &Listen{} },
		},
		{
			name:           "Example 33",
			expectedJSON:   example33,
			expectedStruct: example33Type,
			deserializer:   func() Deserializer { return &Read{} },
		},
		{
			name:           "Example 34",
			expectedJSON:   example34,
			expectedStruct: example34Type,
			deserializer:   func() Deserializer { return &Move{} },
		},
		{
			name:           "Example 35",
			expectedJSON:   example35,
			expectedStruct: example35Type,
			deserializer:   func() Deserializer { return &Travel{} },
		},
		{
			name:           "Example 36",
			expectedJSON:   example36,
			expectedStruct: example36Type,
			deserializer:   func() Deserializer { return &Announce{} },
		},
		{
			name:           "Example 37",
			expectedJSON:   example37,
			expectedStruct: example37Type,
			deserializer:   func() Deserializer { return &Block{} },
		},
		{
			name:           "Example 38",
			expectedJSON:   example38,
			expectedStruct: example38Type,
			deserializer:   func() Deserializer { return &Flag{} },
		},
		{
			name:           "Example 39",
			expectedJSON:   example39,
			expectedStruct: example39Type,
			deserializer:   func() Deserializer { return &Dislike{} },
		},
		{
			name:           "Example 40",
			expectedJSON:   example40,
			expectedStruct: example40Type,
			deserializer:   func() Deserializer { return &Question{} },
		},
		{
			name:           "Example 41",
			expectedJSON:   example41,
			expectedStruct: example41Type,
			deserializer:   func() Deserializer { return &Question{} },
		},
		{
			name:           "Example 42",
			expectedJSON:   example42,
			expectedStruct: example42Type,
			deserializer:   func() Deserializer { return &Application{} },
		},
		{
			name:           "Example 43",
			expectedJSON:   example43,
			expectedStruct: example43Type,
			deserializer:   func() Deserializer { return &Group{} },
		},
		{
			name:           "Example 44",
			expectedJSON:   example44,
			expectedStruct: example44Type,
			deserializer:   func() Deserializer { return &Organization{} },
		},
		{
			name:           "Example 45",
			expectedJSON:   example45,
			expectedStruct: example45Type,
			deserializer:   func() Deserializer { return &Person{} },
		},
		{
			name:           "Example 46",
			expectedJSON:   example46,
			expectedStruct: example46Type,
			deserializer:   func() Deserializer { return &Service{} },
		},
		{
			name:           "Example 47",
			expectedJSON:   example47,
			expectedStruct: example47Type,
			deserializer:   func() Deserializer { return &Relationship{} },
		},
		{
			name:           "Example 48",
			expectedJSON:   example48,
			expectedStruct: example48Type,
			deserializer:   func() Deserializer { return &Article{} },
		},
		{
			name:           "Example 49",
			expectedJSON:   example49,
			expectedStruct: example49Type,
			deserializer:   func() Deserializer { return &Document{} },
		},
		{
			name:           "Example 50",
			expectedJSON:   example50,
			expectedStruct: example50Type,
			deserializer:   func() Deserializer { return &Audio{} },
		},
		{
			name:           "Example 51",
			expectedJSON:   example51,
			expectedStruct: example51Type,
			deserializer:   func() Deserializer { return &Image{} },
		},
		{
			name:           "Example 52",
			expectedJSON:   example52,
			expectedStruct: example52Type,
			deserializer:   func() Deserializer { return &Video{} },
		},
		{
			name:           "Example 53",
			expectedJSON:   example53,
			expectedStruct: example53Type,
			deserializer:   func() Deserializer { return &Note{} },
		},
		{
			name:           "Example 54",
			expectedJSON:   example54,
			expectedStruct: example54Type,
			deserializer:   func() Deserializer { return &Page{} },
		},
		{
			name:           "Example 55",
			expectedJSON:   example55,
			expectedStruct: example55Type,
			deserializer:   func() Deserializer { return &Event{} },
		},
		{
			name:           "Example 56",
			expectedJSON:   example56,
			expectedStruct: example56Type,
			deserializer:   func() Deserializer { return &Place{} },
		},
		{
			name:           "Example 57",
			expectedJSON:   example57,
			expectedStruct: example57Type,
			deserializer:   func() Deserializer { return &Place{} },
		},
		{
			name:           "Example 58",
			expectedJSON:   example58,
			expectedStruct: example58Type,
			deserializer:   func() Deserializer { return &Mention{} },
		},
		{
			name:           "Example 59",
			expectedJSON:   example59,
			expectedStruct: example59Type,
			deserializer:   func() Deserializer { return &Profile{} },
		},
		{
			name:           "Example 60",
			expectedJSON:   example60,
			expectedStruct: example60Type,
			deserializer:   func() Deserializer { return &OrderedCollection{} },
		},
		{
			name:           "Example 61",
			expectedJSON:   example61,
			expectedStruct: example61Type,
			deserializer:   func() Deserializer { return &Unknown{} },
		},
		{
			name:           "Example 62",
			expectedJSON:   example62,
			expectedStruct: example62Type,
			deserializer:   func() Deserializer { return &Unknown{} },
		},
		{
			name:           "Example 63",
			expectedJSON:   example63,
			expectedStruct: example63Type,
			deserializer:   func() Deserializer { return &Offer{} },
		},
		{
			name:           "Example 64",
			expectedJSON:   example64,
			expectedStruct: example64Type,
			deserializer:   func() Deserializer { return &Offer{} },
		},
		{
			name:           "Example 65",
			expectedJSON:   example65,
			expectedStruct: example65Type,
			deserializer:   func() Deserializer { return &Offer{} },
		},
		{
			name:           "Example 66",
			expectedJSON:   example66,
			expectedStruct: example66Type,
			deserializer:   func() Deserializer { return &Note{} },
		},
		{
			name:           "Example 67",
			expectedJSON:   example67,
			expectedStruct: example67Type,
			deserializer:   func() Deserializer { return &Image{} },
		},
		{
			name:           "Example 68",
			expectedJSON:   example68,
			expectedStruct: example68Type,
			deserializer:   func() Deserializer { return &Image{} },
		},
		{
			name:           "Example 69",
			expectedJSON:   example69,
			expectedStruct: example69Type,
			deserializer:   func() Deserializer { return &Note{} },
		},
		{
			name:           "Example 70",
			expectedJSON:   example70,
			expectedStruct: example70Type,
			deserializer:   func() Deserializer { return &Offer{} },
		},
		{
			name:           "Example 71",
			expectedJSON:   example71,
			expectedStruct: example71Type,
			deserializer:   func() Deserializer { return &Offer{} },
		},
		{
			name:           "Example 72",
			expectedJSON:   example72,
			expectedStruct: example72Type,
			deserializer:   func() Deserializer { return &Offer{} },
		},
		{
			name:           "Example 73",
			expectedJSON:   example73,
			expectedStruct: example73Type,
			deserializer:   func() Deserializer { return &Collection{} },
		},
		{
			name:           "Example 74",
			expectedJSON:   example74,
			expectedStruct: example74Type,
			deserializer:   func() Deserializer { return &Collection{} },
		},
		{
			name:           "Example 75",
			expectedJSON:   example75,
			expectedStruct: example75Type,
			deserializer:   func() Deserializer { return &Collection{} },
		},
		{
			name:           "Example 76",
			expectedJSON:   example76,
			expectedStruct: example76Type,
			deserializer:   func() Deserializer { return &Collection{} },
		},
		{
			name:           "Example 77",
			expectedJSON:   example77,
			expectedStruct: example77Type,
			deserializer:   func() Deserializer { return &Collection{} },
		},
		{
			name:           "Example 78",
			expectedJSON:   example78,
			expectedStruct: example78Type,
			deserializer:   func() Deserializer { return &Note{} },
		},
		{
			name:           "Example 79",
			expectedJSON:   example79,
			expectedStruct: example79Type,
			deserializer:   func() Deserializer { return &Note{} },
		},
		{
			name:           "Example 80",
			expectedJSON:   example80,
			expectedStruct: example80Type,
			deserializer:   func() Deserializer { return &Note{} },
		},
		{
			name:           "Example 81",
			expectedJSON:   example81,
			expectedStruct: example81Type,
			deserializer:   func() Deserializer { return &Note{} },
		},
		{
			name:           "Example 82",
			expectedJSON:   example82,
			expectedStruct: example82Type,
			deserializer:   func() Deserializer { return &Note{} },
		},
		{
			name:           "Example 83",
			expectedJSON:   example83,
			expectedStruct: example83Type,
			deserializer:   func() Deserializer { return &Note{} },
		},
		{
			name:           "Example 84",
			expectedJSON:   example84,
			expectedStruct: example84Type,
			deserializer:   func() Deserializer { return &Note{} },
		},
		{
			name:           "Example 85",
			expectedJSON:   example85,
			expectedStruct: example85Type,
			deserializer:   func() Deserializer { return &Listen{} },
		},
		{
			name:           "Example 86",
			expectedJSON:   example86,
			expectedStruct: example86Type,
			deserializer:   func() Deserializer { return &Collection{} },
		},
		{
			name:           "Example 87",
			expectedJSON:   example87,
			expectedStruct: example87Type,
			deserializer:   func() Deserializer { return &Collection{} },
		},
		{
			name:           "Example 88",
			expectedJSON:   example88,
			expectedStruct: example88Type,
			deserializer:   func() Deserializer { return &Person{} },
		},
		{
			name:           "Example 89",
			expectedJSON:   example89,
			expectedStruct: example89Type,
			deserializer:   func() Deserializer { return &Collection{} },
		},
		{
			name:           "Example 90",
			expectedJSON:   example90,
			expectedStruct: example90Type,
			deserializer:   func() Deserializer { return &OrderedCollection{} },
		},
		{
			name:           "Example 91",
			expectedJSON:   example91,
			expectedStruct: example91Type,
			deserializer:   func() Deserializer { return &Question{} },
		},
		{
			name:           "Example 92",
			expectedJSON:   example92,
			expectedStruct: example92Type,
			deserializer:   func() Deserializer { return &Question{} },
		},
		{
			name:           "Example 93",
			expectedJSON:   example93,
			expectedStruct: example93Type,
			deserializer:   func() Deserializer { return &Question{} },
		},
		{
			name:           "Example 94",
			expectedJSON:   example94,
			expectedStruct: example94Type,
			deserializer:   func() Deserializer { return &Move{} },
		},
		{
			name:           "Example 95",
			expectedJSON:   example95,
			expectedStruct: example95Type,
			deserializer:   func() Deserializer { return &CollectionPage{} },
		},
		{
			name:           "Example 96",
			expectedJSON:   example96,
			expectedStruct: example96Type,
			deserializer:   func() Deserializer { return &CollectionPage{} },
		},
		{
			name:           "Example 97",
			expectedJSON:   example97,
			expectedStruct: example97Type,
			deserializer:   func() Deserializer { return &Like{} },
		},
		{
			name:           "Example 98",
			expectedJSON:   example98,
			expectedStruct: example98Type,
			deserializer:   func() Deserializer { return &Like{} },
		},
		{
			name:           "Example 99",
			expectedJSON:   example99,
			expectedStruct: example99Type,
			deserializer:   func() Deserializer { return &Like{} },
		},
		{
			name:           "Example 100",
			expectedJSON:   example100,
			expectedStruct: example100Type,
			deserializer:   func() Deserializer { return &CollectionPage{} },
		},
		{
			name:           "Example 101",
			expectedJSON:   example101,
			expectedStruct: example101Type,
			deserializer:   func() Deserializer { return &CollectionPage{} },
		},
		{
			name:           "Example 102",
			expectedJSON:   example102,
			expectedStruct: example102Type,
			deserializer:   func() Deserializer { return &Video{} },
		},
		{
			name:           "Example 103",
			expectedJSON:   example103,
			expectedStruct: example103Type,
			deserializer:   func() Deserializer { return &Activity{} },
		},
		{
			name:           "Example 104",
			expectedJSON:   example104,
			expectedStruct: example104Type,
			deserializer:   func() Deserializer { return &Note{} },
		},
		{
			name:           "Example 105",
			expectedJSON:   example105,
			expectedStruct: example105Type,
			deserializer:   func() Deserializer { return &Image{} },
		},
		{
			name:           "Example 106",
			expectedJSON:   example106,
			expectedStruct: example106Type,
			deserializer:   func() Deserializer { return &Offer{} },
		},
		{
			name:           "Example 107",
			expectedJSON:   example107,
			expectedStruct: example107Type,
			deserializer:   func() Deserializer { return &Offer{} },
		},
		{
			name:           "Example 108",
			expectedJSON:   example108,
			expectedStruct: example108Type,
			deserializer:   func() Deserializer { return &Offer{} },
		},
		{
			name:           "Example 109",
			expectedJSON:   example109,
			expectedStruct: example109Type,
			deserializer:   func() Deserializer { return &Document{} },
		},
		{
			name:           "Example 110",
			expectedJSON:   example110,
			expectedStruct: example110Type,
			deserializer:   func() Deserializer { return &Document{} },
		},
		{
			name:           "Example 111",
			expectedJSON:   example111,
			expectedStruct: example111Type,
			deserializer:   func() Deserializer { return &Document{} },
		},
		{
			name:           "Example 112",
			expectedJSON:   example112,
			expectedStruct: example112Type,
			deserializer:   func() Deserializer { return &Place{} },
		},
		{
			name:           "Example 113",
			expectedJSON:   example113,
			expectedStruct: example113Type,
			deserializer:   func() Deserializer { return &Place{} },
		},
		{
			name:           "Example 114",
			expectedJSON:   example114,
			expectedStruct: example114Type,
			deserializer:   func() Deserializer { return &Note{} },
		},
		{
			name:           "Example 115",
			expectedJSON:   example115,
			expectedStruct: example115Type,
			deserializer:   func() Deserializer { return &Note{} },
		},
		{
			name:           "Example 116",
			expectedJSON:   example116,
			expectedStruct: example116Type,
			deserializer:   func() Deserializer { return &Note{} },
		},
		{
			name:           "Example 117",
			expectedJSON:   example117,
			expectedStruct: example117Type,
			deserializer:   func() Deserializer { return &Note{} },
		},
		{
			name:           "Example 118",
			expectedJSON:   example118,
			expectedStruct: example118Type,
			deserializer:   func() Deserializer { return &Note{} },
		},
		{
			name:           "Example 119",
			expectedJSON:   example119,
			expectedStruct: example119Type,
			deserializer:   func() Deserializer { return &Video{} },
		},
		{
			name:           "Example 120",
			expectedJSON:   example120,
			expectedStruct: example120Type,
			deserializer:   func() Deserializer { return &Link{} },
		},
		{
			name:           "Example 121",
			expectedJSON:   example121,
			expectedStruct: example121Type,
			deserializer:   func() Deserializer { return &Link{} },
		},
		{
			name:           "Example 122",
			expectedJSON:   example122,
			expectedStruct: example122Type,
			deserializer:   func() Deserializer { return &Link{} },
		},
		{
			name:           "Example 123",
			expectedJSON:   example123,
			expectedStruct: example123Type,
			deserializer:   func() Deserializer { return &CollectionPage{} },
		},
		{
			name:           "Example 124",
			expectedJSON:   example124,
			expectedStruct: example124Type,
			deserializer:   func() Deserializer { return &Place{} },
		},
		{
			name:           "Example 125",
			expectedJSON:   example125,
			expectedStruct: example125Type,
			deserializer:   func() Deserializer { return &Place{} },
		},
		{
			name:           "Example 126",
			expectedJSON:   example126,
			expectedStruct: example126Type,
			deserializer:   func() Deserializer { return &Link{} },
		},
		{
			name:           "Example 127",
			expectedJSON:   example127,
			expectedStruct: example127Type,
			deserializer:   func() Deserializer { return &Event{} },
		},
		{
			name:           "Example 128",
			expectedJSON:   example128,
			expectedStruct: example128Type,
			deserializer:   func() Deserializer { return &Note{} },
		},
		{
			name:           "Example 129",
			expectedJSON:   example129,
			expectedStruct: example129Type,
			deserializer:   func() Deserializer { return &Event{} },
		},
		{
			name:           "Example 130",
			expectedJSON:   example130,
			expectedStruct: example130Type,
			deserializer:   func() Deserializer { return &Place{} },
		},
		{
			name:           "Example 131",
			expectedJSON:   example131,
			expectedStruct: example131Type,
			deserializer:   func() Deserializer { return &Link{} },
		},
		{
			name:           "Example 132",
			expectedJSON:   example132,
			expectedStruct: example132Type,
			deserializer:   func() Deserializer { return &OrderedCollectionPage{} },
		},
		{
			name:           "Example 133",
			expectedJSON:   example133,
			expectedStruct: example133Type,
			deserializer:   func() Deserializer { return &Note{} },
		},
		{
			name:           "Example 134",
			expectedJSON:   example134,
			expectedStruct: example134Type,
			deserializer:   func() Deserializer { return &Note{} },
		},
		{
			name:           "Example 135",
			expectedJSON:   example135,
			expectedStruct: example135Type,
			deserializer:   func() Deserializer { return &Collection{} },
		},
		{
			name:           "Example 136",
			expectedJSON:   example136,
			expectedStruct: example136Type,
			deserializer:   func() Deserializer { return &Place{} },
		},
		{
			name:           "Example 137",
			expectedJSON:   example137,
			expectedStruct: example137Type,
			deserializer:   func() Deserializer { return &Note{} },
		},
		{
			name:           "Example 138",
			expectedJSON:   example138,
			expectedStruct: example138Type,
			deserializer:   func() Deserializer { return &Link{} },
		},
		{
			name:           "Example 139",
			expectedJSON:   example139,
			expectedStruct: example139Type,
			deserializer:   func() Deserializer { return &Relationship{} },
		},
		{
			name:           "Example 140",
			expectedJSON:   example140,
			expectedStruct: example140Type,
			deserializer:   func() Deserializer { return &Relationship{} },
		},
		{
			name:           "Example 141",
			expectedJSON:   example141,
			expectedStruct: example141Type,
			deserializer:   func() Deserializer { return &Profile{} },
		},
		{
			name:           "Example 142",
			expectedJSON:   example142,
			expectedStruct: example142Type,
			deserializer:   func() Deserializer { return &Tombstone{} },
		},
		{
			name:           "Example 143",
			expectedJSON:   example143,
			expectedStruct: example143Type,
			deserializer:   func() Deserializer { return &Tombstone{} },
		},
		{
			name:           "Example 144",
			expectedJSON:   example144,
			expectedStruct: example144Type,
			deserializer:   func() Deserializer { return &Collection{} },
		},
		{
			name:           "Example 145",
			expectedJSON:   example145,
			expectedStruct: example145Type,
			deserializer:   func() Deserializer { return &Collection{} },
		},
		{
			name:           "Example 146",
			expectedJSON:   example146,
			expectedStruct: example146Type,
			deserializer:   func() Deserializer { return &Create{} },
		},
		{
			name:           "Example 147",
			expectedJSON:   example147,
			expectedStruct: example147Type,
			deserializer:   func() Deserializer { return &Offer{} },
		},
		{
			name:           "Example 148",
			expectedJSON:   example148,
			expectedStruct: example148Type,
			deserializer:   func() Deserializer { return &Collection{} },
		},
		{
			name:           "Example 149",
			expectedJSON:   example149,
			expectedStruct: example149Type,
			deserializer:   func() Deserializer { return &Place{} },
		},
		{
			name:           "Example 150",
			expectedJSON:   example150,
			expectedStruct: example150Type,
			deserializer:   func() Deserializer { return &Place{} },
		},
		{
			name:           "Example 151",
			expectedJSON:   example151,
			expectedStruct: example151Type,
			deserializer:   func() Deserializer { return &Question{} },
		},
		{
			name:           "Example 152",
			expectedJSON:   example152,
			expectedStruct: example152Type,
			deserializer:   func() Deserializer { return &Question{} },
		},
		{
			name:           "Example 153",
			expectedJSON:   example153,
			expectedStruct: example153Type,
			deserializer:   func() Deserializer { return &Unknown{} },
		},
		{
			name:           "Example 154",
			expectedJSON:   example154,
			expectedStruct: example154Type,
			deserializer:   func() Deserializer { return &Question{} },
		},
		{
			name:           "Example 155",
			expectedJSON:   example155,
			expectedStruct: example155Type,
			deserializer:   func() Deserializer { return &Collection{} },
		},
		{
			name:           "Example 156",
			expectedJSON:   example156,
			expectedStruct: example156Type,
			deserializer:   func() Deserializer { return &Collection{} },
		},
		{
			name:           "Example 157",
			expectedJSON:   example157,
			expectedStruct: example157Type,
			deserializer:   func() Deserializer { return &Note{} },
		},
		{
			name:           "Example 158",
			expectedJSON:   example158,
			expectedStruct: example158Type,
			deserializer:   func() Deserializer { return &Note{} },
		},
		{
			name:           "Example 159",
			expectedJSON:   example159,
			expectedStruct: example159Type,
			deserializer:   func() Deserializer { return &Move{} },
		},
	}
)

func TestDeserialization(t *testing.T) {
	for _, r := range tables {
		// Test Deserialize
		m := make(map[string]interface{})
		err := json.Unmarshal([]byte(r.expectedJSON), &m)
		if err != nil {
			t.Errorf("%s: Cannot json.Unmarshal: %s", r.name, err)
			continue
		}
		actual := r.deserializer()
		err = actual.Deserialize(m)
		if err != nil {
			t.Errorf("%s: Cannot Deserialize: %s", r.name, err)
			continue
		}
		if diff := deep.Equal(actual, r.expectedStruct); diff != nil {
			t.Errorf("%s: Deserialize deep equal is false: %s", diff)
		}
	}
}

func TestSerialization(t *testing.T) {
	for _, r := range tables {
		m, err := r.expectedStruct.Serialize()
		if err != nil {
			t.Errorf("%s: Cannot Serialize: %s", r.name, err)
			continue
		}
		m["@context"] = "https://www.w3.org/ns/activitystreams"
		b, err := json.Marshal(m)
		if err != nil {
			t.Errorf("%s: Cannot json.Marshal: %s", r.name, err)
			continue
		}
		if diff, err := GetJSONDiff(b, []byte(r.expectedJSON)); err == nil && diff != nil {
			t.Errorf("%s: Serialize JSON equality is false:\n%s", r.name, diff)
		} else if err != nil {
			t.Errorf("%s: GetJSONDiff returned error: %s", r.name, err)
		}
	}
}

func TestSerializationWithoutTypeSet(t *testing.T) {
	obj := &Object{}
	obj.SetId(MustParseURL("http://www.test.example/object/1"))
	obj.AddNameString("A Simple, non-specific object")

	person := &Person{}
	person.AddNameString("Sally")
	place := &Place{}
	place.AddNameString("Work")
	obj4 := &Travel{}
	obj4.AddSummaryString("Sally went to work")
	obj4.AddActorObject(person)
	obj4.AddTargetObject(place)

	tables := []struct {
		name         string
		expectedJSON string
		input        Serializer
	}{
		{
			name:         "Example 1",
			expectedJSON: example1,
			input:        obj,
		},
		{
			name:         "Example 4",
			expectedJSON: example4,
			input:        obj4,
		},
	}
	for _, r := range tables {
		m, err := r.input.Serialize()
		if err != nil {
			t.Errorf("%s: Cannot Serialize: %s", r.name, err)
			continue
		}
		m["@context"] = "https://www.w3.org/ns/activitystreams"
		b, err := json.Marshal(m)
		if err != nil {
			t.Errorf("%s: Cannot json.Marshal: %s", r.name, err)
			continue
		}
		if diff, err := GetJSONDiff(b, []byte(r.expectedJSON)); err == nil && diff != nil {
			t.Errorf("%s: Serialize JSON equality is false:\n%s", r.name, diff)
		} else if err != nil {
			t.Errorf("%s: GetJSON Diff returned error: %s", r.name, err)
		}
	}
}

func TestReserializationAbility(t *testing.T) {
	const (
		samIRIInboxString = "https://example.com/sam/inbox"
		samIRIString      = "https://example.com/sam"
		noteURIString     = "https://example.com/note/123"
		sallyIRIString    = "https://example.com/sally"
		activityIRIString = "https://example.com/test/new/iri"
	)
	samIRIInbox, err := url.Parse(samIRIInboxString)
	if err != nil {
		t.Fatal(err)
	}
	samIRI, err := url.Parse(samIRIString)
	if err != nil {
		t.Fatal(err)
	}
	noteIRI, err := url.Parse(noteURIString)
	if err != nil {
		t.Fatal(err)
	}
	sallyIRI, err := url.Parse(sallyIRIString)
	if err != nil {
		t.Fatal(err)
	}
	activityIRI, err := url.Parse(activityIRIString)
	if err != nil {
		t.Fatal(err)
	}
	expectedSamActor := &Person{}
	expectedSamActor.SetInboxAnyURI(*samIRIInbox)
	expectedSamActor.SetId(*samIRI)
	expectedNote := &Note{}
	expectedNote.SetId(*noteIRI)
	expectedNote.AddNameString("A Note")
	expectedNote.AddContentString("This is a simple note")
	expectedNote.AddToObject(expectedSamActor)
	expectedUpdate := &Update{}
	expectedUpdate.AddActorIRI(*sallyIRI)
	expectedUpdate.AddSummaryString("Sally updated her note")
	expectedUpdate.SetId(*activityIRI)
	expectedUpdate.AddObject(expectedNote)
	tables := []struct {
		name         string
		expected     Serializer
		deserializer func() Deserializer
		input        string
	}{
		{
			name:         "JSON with null",
			expected:     expectedUpdate,
			deserializer: func() Deserializer { return &Update{} },
			input: `
                        {
                          "@context": "https://www.w3.org/ns/activitystreams",
                          "summary": "Sally updated her note",
                          "type": "Update",
                          "actor": "https://example.com/sally",
	                  "id": "https://example.com/test/new/iri",
                          "object": {
                            "id": "https://example.com/note/123",
	                    "type": "Note",
                            "to": {
                              "id": "https://example.com/sam",
                              "inbox": "https://example.com/sam/inbox",
	                      "type": "Person",
	                      "name": null
                            }
	                  }
                        }
			`,
		},
	}
	for _, r := range tables {
		m := make(map[string]interface{})
		err := json.Unmarshal([]byte(r.input), &m)
		if err != nil {
			t.Errorf("%s: Cannot json.Unmarshal: %s", r.name, err)
			continue
		}
		actual := r.deserializer()
		err = actual.Deserialize(m)
		if err != nil {
			t.Errorf("%s: Cannot Deserialize: %s", r.name, err)
			continue
		}
		if diff := deep.Equal(actual, r.expected); diff != nil {
			t.Errorf("%s: Deserialize deep equal is false: %s", r.name, diff)
		}
		// Only test ability to reserialize
		serializer, ok := actual.(Serializer)
		if !ok {
			t.Errorf("Deserializer is not also Serializer")
		}
		m, err = serializer.Serialize()
		if err != nil {
			t.Errorf("%s: Cannot Serialize: %s", r.name, err)
			continue
		}
		m["@context"] = "https://www.w3.org/ns/activitystreams"
		_, err = json.Marshal(m)
		if err != nil {
			t.Errorf("%s: Cannot json.Marshal: %s", r.name, err)
			continue
		}
	}
}

func TestBooleanOperationsOnEmptyTypes(t *testing.T) {
	o := &Object{}
	// Functional multi-type property
	if o.IsAltitude() {
		t.Errorf("IsAltitude returned true; expected false")
	}
	if o.IsAltitudeIRI() {
		t.Errorf("IsAltitudeIRI returned true; expected false")
	}
	if o.HasUnknownAltitude() {
		t.Errorf("HasUnknownAltitude returned true; expected false")
	}
	// Non-functional multi-type property
	if o.AttachmentLen() > 0 {
		t.Errorf("AttachentLen returned >0; expected ==0")
	}
	if o.HasUnknownAttachment() {
		t.Errorf("HasUnknownAttachment returned true; expected false")
	}
	// Functional any property
	// None exist
	// Non-functional any property
	if o.TypeLen() > 0 {
		t.Errorf("TypeLen returned >0; expected ==0")
	}

	l := &Link{}
	// Functional single-type property
	if l.HasHref() {
		t.Errorf("HasHref returned true; expected false")
	}
	if l.HasUnknownHref() {
		t.Errorf("HasUnknownHref returned true; expected false")
	}
}
