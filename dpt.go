package nmea0183

import (
	"fmt"

	goNMEA "github.com/adrianmo/go-nmea"
)

type DPT struct {
	goNMEA.BaseSentence
	Depth      Float64
	Offset     Float64
	RangeScale Float64
}

func init() {
	goNMEA.MustRegisterParser("DPT", func(s goNMEA.BaseSentence) (goNMEA.Sentence, error) {
		p := goNMEA.NewParser(s)
		result := DPT{
			BaseSentence: s,
		}
		if p.Fields[0] != "" {
			result.Depth = NewFloat64WithValue(p.Float64(0, "depth"))
		} else {
			result.Depth = NewFloat64()
		}
		if p.Fields[2] != "" {
			result.Offset = NewFloat64WithValue(p.Float64(1, "offset"))
		} else {
			result.Offset = NewFloat64()
		}
		if p.Fields[4] != "" {
			result.RangeScale = NewFloat64WithValue(p.Float64(2, "range scale"))
		} else {
			result.RangeScale = NewFloat64()
		}
		return result, p.Err()
	})
}

// GetDepthBelowTransducer retrieves the depth below the keel from the sentence
func (s DPT) GetDepthBelowTransducer() (float64, error) {
	if v, err := s.Depth.GetValue(); err == nil {
		return v, nil
	}
	return 0, fmt.Errorf("value is unavailable")
}

// GetDepthBelowKeel retrieves the depth below the keel from the sentence
func (s DPT) GetDepthBelowKeel() (float64, error) {
	if vDepth, err := s.Depth.GetValue(); err == nil {
		if vOffset, err := s.Offset.GetValue(); err == nil && vOffset < 0 {
			return vDepth + vOffset, nil
		}
	}
	return 0, fmt.Errorf("value is unavailable")
}

// GetDepthBelowSurface retrieves the depth below surface from the sentence
func (s DPT) GetDepthBelowSurface() (float64, error) {
	if vDepth, err := s.Depth.GetValue(); err == nil {
		if vOffset, err := s.Offset.GetValue(); err == nil && vOffset > 0 {
			return vDepth + vOffset, nil
		}
	}
	return 0, fmt.Errorf("value is unavailable")
}
