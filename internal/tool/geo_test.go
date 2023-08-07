package tool

import "testing"
import "github.com/lishimeng/geo/geoconvert"

func TestGeoConvert(t *testing.T) {
	var lat = 31.3109455
	var lng = 121.5040130

	lng, lat = geoconvert.WGS84toBD09(lng, lat)
	t.Logf("%f:%f", lng, lat)
}
