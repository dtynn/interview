package proto

type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type Schedule struct {
	Groups []ScheduleGroup `json:"groups"`
}

type ScheduleGroup struct {
	ID      int             `json:"id"`
	Teams   []Team          `json:"team"`
	Matches []ScheduleMatch `json:"matches"`
}

type ScheduleMatch struct {
	Day  int    `json:"day"`
	Time string `json:"time"`
	Home Team   `json:"home"`
	Away Team   `json:"away"`
}
