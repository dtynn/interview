package handler

import (
	"context"
	"net/http"

	"github.com/dtynn/interview/db"
	"github.com/dtynn/interview/proto"
	"github.com/dtynn/interview/util"
)

func Schedule(rw http.ResponseWriter, req *http.Request) {
	var resp proto.Response

	teams, err := db.Team.List(context.Background(), 16)
	if err != nil {
		resp.Code = 1
		resp.Message = err.Error()
		util.Resp(rw, resp)
		return
	}

	if len(teams) != 16 {
		resp.Code = 1
		resp.Message = "16 teams required"
		util.Resp(rw, resp)
		return
	}

	var schedule proto.Schedule

	times := []string{"23:00", "02:00"}

	for gi := 0; gi < 4; gi++ {
		inner := teams[gi*4 : gi*4+4]
		matches := pick(inner, 2)

		group := proto.ScheduleGroup{
			ID:    gi + 1,
			Teams: inner,
		}
		for mi := range matches {
			group.Matches = append(group.Matches, proto.ScheduleMatch{
				Day:  mi/2 + 1,
				Time: times[mi%2],
				Home: matches[mi][0],
				Away: matches[mi][1],
			})
		}

		schedule.Groups = append(schedule.Groups, group)
	}

	util.Resp(rw, schedule)
	return
}

func pick(teams []proto.Team, n int) [][]proto.Team {
	res := make([][]proto.Team, 0)
	doPick(nil, teams, n, &res)
	return res
}

func doPick(prefix, left []proto.Team, n int, res *[][]proto.Team) {
	if len(prefix) == n {
		*res = append(*res, prefix)
		return
	}

	for i := range left {
		next := make([]proto.Team, len(prefix)+1)
		copy(next, prefix)
		next[len(next)-1] = left[i]
		doPick(next, left[i+1:], n, res)
	}
}
