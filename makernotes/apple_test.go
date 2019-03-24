package makernotes

import (
	"os"
	"testing"

	"github.com/rwcarlsen/goexif/exif"
	"github.com/stretchr/testify/require"
)

func init() {
	exif.RegisterParsers(Apple)
}

func TestExtraction(t *testing.T) {
	t.Run("Burst UUID", func(t *testing.T) {
		f, err := os.Open("../testdata/apple_burstuuid.jpg")
		require.NoError(t, err)

		x, err := exif.Decode(f)
		require.NoError(t, err)

		data, err := x.Get(Apple_AccelerationVector)
		require.NoError(t, err)
		j, _ := data.MarshalJSON()
		require.Equal(t, string(j), `["-5691/6191","-677/22055","-1538/3835"]`)

		data, err = x.Get(Apple_BurstUUID)
		require.NoError(t, err)
		j, _ = data.MarshalJSON()
		require.Equal(t, string(j), `"81F26A91-7A4F-4968-AA5A-9189D3681D91"`)
	})

	t.Run("Content ID", func(t *testing.T) {
		f, err := os.Open("../testdata/apple_contentid.jpg")
		require.NoError(t, err)

		x, err := exif.Decode(f)
		require.NoError(t, err)

		data, err := x.Get(Apple_AccelerationVector)
		require.NoError(t, err)
		j, _ := data.MarshalJSON()
		require.Equal(t, string(j), `["-1340/16057","-812/905","-1205/2874"]`)

		data, err = x.Get(Apple_ContentIdentifier)
		require.NoError(t, err)
		j, _ = data.MarshalJSON()
		require.Equal(t, string(j), `"D4122CE4-5DA9-4EBA-9475-01579D01ADCE"`)
	})

}
