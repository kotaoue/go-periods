package periods

import (
	"errors"
	"strings"
	"time"
)

func Split(period string) ([]string, error) {
	ss := strings.Split(period, "-")
	if len(ss) == 1 {
		return ss, nil
	}

	if len(ss) != 2 {
		return nil, errors.New("3 or more periods specified")
	}

	from, err := time.Parse("20060102", ss[0])
	if err != nil {
		return nil, err
	}

	to, err := time.Parse("20060102", ss[1])
	if err != nil {
		return nil, err
	}

	var periods []string
	periods = append(periods, from.Format("20060102"))
	current := from
	for current.Before(to) {
		current = current.Add(time.Hour * 24)
		periods = append(periods, current.Format("20060102"))
	}

	return periods, nil
}
