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

	switch len(ss[0]) {
	case 8:
		return splitDay(ss)
	case 6:
		return splitMonth(ss)
	}
}

func splitDay(ss []string) ([]string, error) {
	from, err := time.Parse("20060102", ss[0])
	if err != nil {
		return nil, err
	}

	to, err := time.Parse("20060102", ss[1])
	if err != nil {
		return nil, err
	}

	var ps []string
	ps = append(ps, from.Format("20060102"))
	current := from
	for current.Before(to) {
		current = current.Add(time.Hour * 24)
		ps = append(ps, current.Format("20060102"))
	}

	return ps, nil
}

func splitMonth(ss []string) ([]string, error) {
	from, err := time.Parse("200601", ss[0])
	if err != nil {
		return nil, err
	}

	to, err := time.Parse("200601", ss[1])
	if err != nil {
		return nil, err
	}

	var ps []string
	ps = append(ps, from.Format("200601"))
	current := from
	for current.Before(to) {
		current = time.Date(current.Year(), current.Month()+1, 1, 0, 0, 0, 0, time.UTC)
		ps = append(ps, current.Format("200601"))
	}

	return ps, nil
}
