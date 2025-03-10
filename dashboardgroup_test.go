package signalfx

import (
	"context"
	"net/http"
	"net/url"
	"strconv"
	"testing"

	"github.com/signalfx/signalfx-go/dashboard_group"
	"github.com/stretchr/testify/assert"
)

// TODO Create simple dashboard group

func TestCreateDashboardGroup(t *testing.T) {
	teardown := setup()
	defer teardown()

	mux.HandleFunc("/v2/dashboardgroup", verifyRequest(t, "POST", true, http.StatusOK, nil, "dashboardgroup/create_success.json"))

	result, err := client.CreateDashboardGroup(context.Background(), &dashboard_group.CreateUpdateDashboardGroupRequest{
		Name: "string",
	}, false)
	assert.NoError(t, err, "Unexpected error creating dashboard group")
	assert.Equal(t, "string", result.Name, "Name does not match")
}

func TestCreateEmptyDashboardGroup(t *testing.T) {
	teardown := setup()
	defer teardown()

	params := url.Values{}
	params.Add("empty", "true")

	mux.HandleFunc("/v2/dashboardgroup", verifyRequest(t, "POST", true, http.StatusOK, params, "dashboardgroup/create_success.json"))

	result, err := client.CreateDashboardGroup(context.Background(), &dashboard_group.CreateUpdateDashboardGroupRequest{
		Name: "string",
	}, true)
	assert.NoError(t, err, "Unexpected error creating dashboard group")
	assert.Equal(t, "string", result.Name, "Name does not match")
}

func TestCreateBadDashboardGroup(t *testing.T) {
	teardown := setup()
	defer teardown()

	mux.HandleFunc("/v2/dashboardgroup", verifyRequest(t, "POST", true, http.StatusBadRequest, nil, ""))

	result, err := client.CreateDashboardGroup(context.Background(), &dashboard_group.CreateUpdateDashboardGroupRequest{
		Name: "string",
	}, false)
	assert.Error(t, err, "Should get an error from bad dashboard group")
	assert.Nil(t, result, "Should get nil result from bad dashboard group")
}

func TestDeleteDashboardGroup(t *testing.T) {
	teardown := setup()
	defer teardown()

	mux.HandleFunc("/v2/dashboardgroup/string", verifyRequest(t, "DELETE", true, http.StatusNoContent, nil, ""))

	err := client.DeleteDashboardGroup(context.Background(), "string")
	assert.NoError(t, err, "Unexpected error getting dashboard group")
}

func TestDeleteMissingDashboardGroup(t *testing.T) {
	teardown := setup()
	defer teardown()

	mux.HandleFunc("/v2/dashboardgroup/string", verifyRequest(t, "DELETE", true, http.StatusNotFound, nil, ""))

	err := client.DeleteDashboardGroup(context.Background(), "string")
	assert.Error(t, err, "Should get an error getting missing dashboard group")
}

func TestGetDashboardGroup(t *testing.T) {
	teardown := setup()
	defer teardown()

	mux.HandleFunc("/v2/dashboardgroup/string", verifyRequest(t, "GET", true, http.StatusOK, nil, "dashboardgroup/get_success.json"))

	result, err := client.GetDashboardGroup(context.Background(), "string")
	assert.NoError(t, err, "Unexpected error getting dashboard group")
	assert.Equal(t, result.Name, "string", "Name does not match")
}

func TestGetMissingDashboardGroup(t *testing.T) {
	teardown := setup()
	defer teardown()

	mux.HandleFunc("/v2/dashboardgroup/string", verifyRequest(t, "GET", true, http.StatusNotFound, nil, ""))

	result, err := client.GetDashboardGroup(context.Background(), "string")
	assert.Error(t, err, "Should get error getting missing dashboard group")
	assert.Nil(t, result, "Result should be nil")
}

func TestSearchDashboardGroup(t *testing.T) {
	teardown := setup()
	defer teardown()

	limit := 10
	name := "foo"
	offset := 2
	params := url.Values{}
	params.Add("limit", strconv.Itoa(limit))
	params.Add("name", name)
	params.Add("offset", strconv.Itoa(offset))

	mux.HandleFunc("/v2/dashboardgroup", verifyRequest(t, "GET", true, http.StatusOK, params, "dashboardgroup/search_success.json"))

	results, err := client.SearchDashboardGroups(context.Background(), limit, name, offset)
	assert.NoError(t, err, "Unexpected error search dashboard group")
	assert.Equal(t, int32(1), results.Count, "Incorrect number of results")
}

func TestSearchDashboardGroupBad(t *testing.T) {
	teardown := setup()
	defer teardown()

	limit := 10
	name := "foo"
	offset := 2
	params := url.Values{}
	params.Add("limit", strconv.Itoa(limit))
	params.Add("name", name)
	params.Add("offset", strconv.Itoa(offset))

	mux.HandleFunc("/v2/dashboardgroup", verifyRequest(t, "GET", true, http.StatusBadRequest, params, ""))

	_, err := client.SearchDashboardGroups(context.Background(), limit, name, offset)
	assert.Error(t, err, "Unexpected error search dashboard group")
}

func TestUpdateDashboardGroup(t *testing.T) {
	teardown := setup()
	defer teardown()

	mux.HandleFunc("/v2/dashboardgroup/string", verifyRequest(t, "PUT", true, http.StatusBadRequest, nil, ""))

	result, err := client.UpdateDashboardGroup(context.Background(), "string", &dashboard_group.CreateUpdateDashboardGroupRequest{
		Name: "string",
	})
	assert.Error(t, err, "Should have error updating missing dashboard group")
	assert.Nil(t, result, "Should have nil result")
}

func TestValidateDashboardGroup(t *testing.T) {
	teardown := setup()
	defer teardown()

	mux.HandleFunc("/v2/dashboardgroup/validate", verifyRequest(t, "POST", true, http.StatusNoContent, nil, ""))

	err := client.ValidateDashboardGroup(context.Background(), &dashboard_group.CreateUpdateDashboardGroupRequest{
		Name: "string",
	})
	assert.NoError(t, err, "Unexpected error validating dashboard group")
}

func TestValidateDashboardGroupBad(t *testing.T) {
	teardown := setup()
	defer teardown()

	mux.HandleFunc("/v2/dashboardgroup/validate", verifyRequest(t, "POST", true, http.StatusBadRequest, nil, ""))

	err := client.ValidateDashboardGroup(context.Background(), &dashboard_group.CreateUpdateDashboardGroupRequest{
		Name: "string",
	})
	assert.Error(t, err, "Should have gotten an error from invalid dashboard group")
}
