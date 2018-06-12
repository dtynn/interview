package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/dtynn/interview/db"
	"github.com/dtynn/interview/proto"
	"github.com/dtynn/interview/util"
)

func AddTeam(rw http.ResponseWriter, req *http.Request) {
	dec := json.NewDecoder(req.Body)
	defer req.Body.Close()

	teams := proto.AddTeamInput{}
	if err := dec.Decode(&teams); err != nil {
		util.Resp(rw, proto.Response{
			Code:    1,
			Message: err.Error(),
		})
		return
	}

	if len(teams.Teams) > 0 {
		if err := db.Team.Add(context.Background(), teams.Teams...); err != nil {
			util.Resp(rw, proto.Response{
				Code:    1,
				Message: err.Error(),
			})
			return
		}
	}

	util.Resp(rw, proto.Response{})
	return
}
