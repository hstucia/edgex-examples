package transforms

import (
	"testing"

	"github.com/edgexfoundry/app-functions-sdk-go/v2/pkg"
	"github.com/edgexfoundry/app-functions-sdk-go/v2/pkg/interfaces"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/dtos"
	"github.com/google/uuid"

	"github.com/edgexfoundry/go-mod-core-contracts/v2/clients/logger"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/common"
	"github.com/stretchr/testify/assert"
)

var context interfaces.AppFunctionContext
var lc logger.LoggingClient

func init() {
	lc := logger.NewMockClient()
	context = pkg.NewAppFuncContextForTest(uuid.NewString(), lc)
}
func TestTransformToInflux(t *testing.T) {

	reading1 := dtos.BaseReading{
		Id:            "id-1",
		ResourceName:  "field1",
		ValueType:     common.ValueTypeInt32,
		SimpleReading: dtos.SimpleReading{Value: "1"},
	}
	reading2 := dtos.BaseReading{
		Id:            "id-2",
		ResourceName:  "field2",
		ValueType:     common.ValueTypeInt32,
		SimpleReading: dtos.SimpleReading{Value: "2"},
	}

	tags := map[string]interface{}{"tag1": "toto"}
	var devID1 = "id1"
	var event = dtos.Event{
		Id:         "event-" + devID1,
		DeviceName: devID1,
		Origin:     123,
		Tags:       tags,
		Readings: []dtos.BaseReading{
			reading1,
			reading2,
		},
	}
	//sample influx line protocol
	want := `id1,tag1=toto field1=1,field2=2 123`
	conv := NewConversion()
	continuePipeline, result := conv.TransformToInflux(context, event)

	assert.NotNil(t, result)
	assert.True(t, continuePipeline)
	assert.Equal(t, want, result)
}
