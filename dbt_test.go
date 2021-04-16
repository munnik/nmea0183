package nmea0183_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/munnik/nmea0183"
)

var _ = Describe("DBT", func() {
	var (
		parsed DBT
	)
	Describe("Getting data from a $__DBT sentence", func() {
		BeforeEach(func() {
			parsed = DBT{
				DepthFeet:    NewFloat64WithValue(DepthBelowSurfaceFeet - DepthTransducerFeet),
				DepthMeters:  NewFloat64WithValue(DepthBelowSurfaceMeters - DepthTransducerMeters),
				DepthFathoms: NewFloat64WithValue(DepthBelowSurfaceFathoms - DepthTransducerFanthoms),
			}
		})
		Context("When having a parsed sentence", func() {
			It("should give a valid depth below surface", func() {
				Expect(parsed.GetDepthBelowTransducer()).To(Float64Equal(DepthBelowSurfaceMeters-DepthTransducerMeters, 0.00001))
			})
		})
		Context("When having a parsed sentence with only depth in feet set", func() {
			JustBeforeEach(func() {
				parsed.DepthMeters = NewFloat64()
				parsed.DepthFathoms = NewFloat64()
			})
			It("should give a valid depth below surface", func() {
				Expect(parsed.GetDepthBelowTransducer()).To(Float64Equal(DepthBelowSurfaceMeters-DepthTransducerMeters, 0.00001))
			})
		})
		Context("When having a parsed sentence with only depth in fathoms set", func() {
			JustBeforeEach(func() {
				parsed.DepthFeet = NewFloat64()
				parsed.DepthMeters = NewFloat64()
			})
			It("should give a valid depth below surface", func() {
				Expect(parsed.GetDepthBelowTransducer()).To(Float64Equal(DepthBelowSurfaceMeters-DepthTransducerMeters, 0.00001))
			})
		})
		Context("When having a parsed sentence with only depth in meters set", func() {
			JustBeforeEach(func() {
				parsed.DepthFeet = NewFloat64()
				parsed.DepthFathoms = NewFloat64()
			})
			It("should give a valid depth below surface", func() {
				Expect(parsed.GetDepthBelowTransducer()).To(Float64Equal(DepthBelowSurfaceMeters-DepthTransducerMeters, 0.00001))
			})
		})
		Context("When having a parsed sentence with missing depth values", func() {
			JustBeforeEach(func() {
				parsed.DepthFeet = NewFloat64()
				parsed.DepthMeters = NewFloat64()
				parsed.DepthFathoms = NewFloat64()
			})
			Specify("an error is returned", func() {
				_, err := parsed.GetDepthBelowTransducer()
				Expect(err).To(HaveOccurred())
			})
		})
	})
})
