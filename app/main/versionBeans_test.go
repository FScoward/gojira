package main

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestParseVersionBeans(t *testing.T) {
	jsonString := `
[
    {
        "self": "http://your-domain.atlassian.net/rest/api/2/version/10000",
        "id": "10000",
        "description": "An excellent version",
        "name": "New Version 1",
        "archived": false,
        "released": true,
        "releaseDate": "2010-07-06",
        "overdue": true,
        "userReleaseDate": "6/Jul/2010",
        "projectId": 10000
    },
    {
        "self": "http://your-domain.atlassian.net/rest/api/2/version/10010",
        "id": "10010",
        "description": "Minor Bugfix version",
        "name": "Next Version",
        "archived": false,
        "released": false,
        "overdue": false,
        "projectId": 10000,
        "issuesStatusForFixVersion": {
            "unmapped": 0,
            "toDo": 10,
            "inProgress": 20,
            "done": 100
        }
    }
]
`
	var versionBeans []VersionBean
	err := json.Unmarshal([]byte(jsonString), &versionBeans)
	if err != nil {
		t.Error(err)
	}

	expect := []VersionBean{
		VersionBean{
			"http://your-domain.atlassian.net/rest/api/2/version/10000",
			"10000",
			"An excellent version",
			"New Version 1",
			false,
			true,
			"2010-07-06",
			true,
			"6/Jul/2010",
			10000,
		},
		VersionBean{
			"http://your-domain.atlassian.net/rest/api/2/version/10010",
			"10010",
			"Minor Bugfix version",
			"Next Version",
			false,
			false,
			"",
			false,
			"",
			10000,
		},
	}

	for i, _ := range versionBeans {
		if !reflect.DeepEqual(expect, versionBeans) {
			t.Errorf("index: %v, actual: %v, expect: %v", i, versionBeans, expect)
		}
	}
}
