package csv_test

import (
	"io"
	"io/ioutil"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/simox-83/CodeClinic/01-parsingdata/formats/csv"
)

func TestRead(t *testing.T) {
	assert := assert.New(t)

	table, err := csv.Read(&testGetter{
		data: `date	time	Air_Temp	Barometric_Press	Dew_Point	Relative_Humidity	Wind_Dir	Wind_Gust	Wind_Speed
2015_01_01	00:02:43	19.50	30.62	14.78	81.60	159.78	14.00	 9.20
2015_01_01	00:02:52	19.50	30.62	14.78	81.60	159.78	14.00	 9.20
2015_01_01	00:07:43	19.50	30.61	14.66	81.20	155.63	11.00	 8.60`})

	assert.IsType([][]string{}, table)
	assert.NoError(err)

	assert.Len(table, 4)
	assert.Len(table[1], 9)

	assert.Equal(table[3][4], "14.66")

}

type testGetter struct {
	data string
}

func (g *testGetter) Get() (io.ReadCloser, error) {
	return ioutil.NopCloser(strings.NewReader(g.data)), nil
}
