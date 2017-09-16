// scantaskschedule_test
package scanengine

import (
	"fmt"
	"testing"
	"unicontract/src/config"
)

func Test_Start(t *testing.T) {
	config.Init()

	Start()
}

func Test_WriteFile(t *testing.T) {
	slID := []string{
		"asdfasdfasdfasdfasdfasdfasdf",
		"wertwertwertwertwerwertwerwe",
		"dfghdfghdfghdfghdfghdfghfgjg",
		"qwerqwertetwersdfsdfgsdfgrer",
		"dfghdrthwergsdgsdfgsdfgssdfg",
		"asdfasdfasdfasdfasdfasdfasdf",
		"sdfgweryeyur6ktuihku1212yuil",
		"asdfasdfasdfasdfasdfasdfasdf",
		"dfjhfghjfghjk123123412341234",
		"shw5yerdrjrtkgyjkgj3ryukrtry",
		"asdfasdfasdfasdfasdfasdfasdf",
		"cndghdrhey345y35y345y345y111",
	}

	var strID string
	for index:= range slID {
		strID = fmt.Sprintf("%s\n%s", strID, slID[index])
	}
	fmt.Println(strID)

	writeCount, err := _WriteFile("./test", strID)
	if err != nil {
		t.Error(err)
	}
	if writeCount != len(strID) {
		t.Error("write count is error")
	}
}
