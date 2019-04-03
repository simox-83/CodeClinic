package csv_test

import (
	"errors"
	"io"
	"strings"
	"testing"

	"github.com/simox-83/CodeClinic/01-parsingdata/formats/csv"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestRead(t *testing.T) {

	t.Run("Success test", func(t *testing.T) {
		assert := assert.New(t)

		r := strings.NewReader(`date	time	Air_Temp	Barometric_Press	Dew_Point	Relative_Humidity	Wind_Dir	Wind_Gust	Wind_Speed
2015_01_01	00:02:43	19.50	30.62	14.78	81.60	159.78	14.00	 9.20
2015_01_01	00:02:52	19.50	30.62	14.78	81.60	159.78	14.00	 9.20
2015_01_01	00:07:43	19.50	30.61	14.66	81.20	155.63	11.00	 8.60`)

		rc := &ClosedMock{Reader: r}
		rc.On("Close").Return(nil)
		m := &MyMockedObject{}
		m.On("Get").Return(rc, nil)
		table, err := csv.Read(m)

		rc.AssertExpectations(t)

		assert.IsType([][]string{}, table)
		assert.NoError(err)

		require.Len(t, table, 4)
		assert.Len(table[1], 9)

		assert.Equal(table[3][4], "14.66")
	})

	t.Run("Error - Get returns nil", func(t *testing.T) {
		assert := assert.New(t)
		sampleErr := errors.New("Sample error")
		var rc *ClosedMock
		m := &MyMockedObject{}
		m.On("Get").Return(rc, sampleErr)
		table, err := csv.Read(m)

		assert.Nil(rc)
		m.AssertNumberOfCalls(t, "Get", 1)

		assert.EqualError(err, sampleErr.Error())
		assert.Nil(table)
	})

	t.Run("Error - Get returns initialized object", func(t *testing.T) {
		assert := assert.New(t)
		sampleErr := errors.New("Sample error")
		r := strings.NewReader("")
		rc := &ClosedMock{Reader: r}
		rc.On("Close").Return(nil)
		m := &MyMockedObject{}
		m.On("Get").Return(rc, sampleErr)
		table, err := csv.Read(m)

		rc.AssertNumberOfCalls(t, "Close", 0)
		m.AssertNumberOfCalls(t, "Get", 1)

		assert.EqualError(err, sampleErr.Error())
		assert.Nil(table)
	})
}

type MyMockedObject struct {
	mock.Mock
}

func (m *MyMockedObject) Get() (io.ReadCloser, error) { // questa e' l'implementazione dell'interfaccia - e' il nostro producer

	args := m.Called()
	if o, ok := args.Get(0).(*ClosedMock); ok {
		///fmt.Printf("%T: %v\n", o, o)
		return o, args.Error(1)
	}
	return args.Get(0).(io.ReadCloser), args.Error(1)

}

type ClosedMock struct {
	mock.Mock
	io.Reader
}

func (m *ClosedMock) Close() error {
	args := m.Called()
	return args.Error(0)
}
