package csv_test

import (
	"errors"
	"io"
	"io/ioutil"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/simox-83/CodeClinic/01-parsingdata/formats/csv"
)

func TestRead(t *testing.T) {

	t.Run("Success test", func(t *testing.T) {
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
	})

	t.Run("Error", func(t *testing.T) {
		assert := assert.New(t)
		g := testGetter{
			err: errors.New("This is an error"),
		}

		table, err := csv.Read(&g)

		assert.EqualError(err, g.err.Error())
		assert.Nil(table)

	})

	t.Run("Error2", func(t *testing.T) {
		assert := assert.New(t)
		g := testGetter{
			err: errors.New("Hello"),
		}

		table, err := csv.Read(&g)

		assert.EqualError(err, g.err.Error())
		assert.Nil(table)

	})
}

type testGetter struct {
	data string
	err  error
}

func (g *testGetter) Get() (io.ReadCloser, error) {
	if g.err != nil {
		return nil, g.err
	}
	return ioutil.NopCloser(strings.NewReader(g.data)), nil
}
