package camera

import (
	"testing"
)

func TestInfraredCameraRetrieveData(t *testing.T) {
	t.Run("Found", func(t *testing.T) {
		c := NewInfraredCamera(1, "Canon", []Data{{1, "Коза", "Побігла вліво"}})
		data, _ := c.retrieveData()

		if len(*data) == 0 {
			t.Errorf("retrieveData() should not return empty data.")
		}

		for _, d := range *data {
			if d.Animal != "Коза" || d.Movement != "Побігла вліво" {
				t.Errorf("retrieveData() = %+v", d)
			}
		}
	})

	t.Run("Not found", func(t *testing.T) {
		c := NewInfraredCamera(1, "Canon", []Data{})
		_, err := c.retrieveData()
		if err == nil {
			t.Errorf("retrieveData() should return error = %v", err)
		}
	})
}

func TestInfraredCameraProcessData(t *testing.T) {
}
