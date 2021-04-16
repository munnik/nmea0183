package nmea0183

import (
	"fmt"

	goNMEA "github.com/adrianmo/go-nmea"
	"github.com/martinlindhe/unit"
)

type DBT struct {
	goNMEA.BaseSentence
	DepthFeet    Float64
	DepthMeters  Float64
	DepthFathoms Float64
}

func init() {
	goNMEA.MustRegisterParser("DBT", func(s goNMEA.BaseSentence) (goNMEA.Sentence, error) {
		p := goNMEA.NewParser(s)
		result := DBT{
			BaseSentence: s,
		}
		if p.Fields[0] != "" {
			result.DepthFeet = NewFloat64WithValue(p.Float64(0, "depth_feet"))
		} else {
			result.DepthFeet = NewFloat64()
		}
		if p.Fields[2] != "" {
			result.DepthMeters = NewFloat64WithValue(p.Float64(2, "depth_meters"))
		} else {
			result.DepthMeters = NewFloat64()
		}
		if p.Fields[4] != "" {
			result.DepthFathoms = NewFloat64WithValue(p.Float64(4, "depth_fathoms"))
		} else {
			result.DepthFathoms = NewFloat64()
		}
		return result, p.Err()
	})
}

// GetDepthBelowTransducer retrieves the depth below the transducer from the sentence
func (s DBT) GetDepthBelowTransducer() (float64, error) {
	if v, err := s.DepthMeters.GetValue(); err == nil {
		return v, nil
	}
	if v, err := s.DepthFeet.GetValue(); err == nil {
		return (unit.Length(v) * unit.Foot).Meters(), nil
	}
	if v, err := s.DepthFathoms.GetValue(); err == nil {
		return (unit.Length(v) * unit.Fathom).Meters(), nil
	}
	return 0, fmt.Errorf("value is unavailable")
}
