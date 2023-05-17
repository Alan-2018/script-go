package isyntax

import (
	"fmt"
	"log"
	"sort"
	"time"
)

/*
	sort
*/

type MapsSort struct {
	Key     string
	MapList []map[string]interface{}
}

func (m *MapsSort) Len() int {
	return len(m.MapList)
}

func (m *MapsSort) Less(i, j int) bool {
	return m.MapList[i][m.Key].(int) > m.MapList[j][m.Key].(int)
}

func (m *MapsSort) Swap(i, j int) {
	m.MapList[i], m.MapList[j] = m.MapList[j], m.MapList[i]
}

func (m *MapsSort) Sort() []map[string]interface{} {

	sort.Sort(m)

	return m.MapList
}

func TestMapsSort() {

	var jarr []jmap = make([]jmap, 0, 10)

	for i := 0; i < 10; i++ {
		j := make(jmap)

		j["id"] = string(i)
		j["num"] = i
		j["ctime"] = int(time.Now().UnixNano() / 1e6)

		jarr = append(jarr, j)
	}

	for idx, _ := range jarr {
		log.Println(jarr[idx])
	}

	log.Println()

	m := MapsSort{
		Key:     "num",
		MapList: jarr,
	}

	jarr2 := m.Sort()

	for idx, _ := range jarr2 {
		log.Println(jarr2[idx])
	}

	strArr := []string{"1", "2", "3", "4", "5", "0"}
	sort.Strings(strArr)

	log.Println(strArr)

	// how to sort string value of map

}

func TestStringSliceSort() {
	slc := []string{
		"apple",
		"tomato",
		"banana",
		"potato",
	}

	sort.StringSlice(slc).Sort()

	fmt.Println(
		slc,
	)
}
