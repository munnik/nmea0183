# nmea0183

This project is based on https://github.com/adrianmo/go-nmea and has some additions to make it a better fit for https://github.com/munnik/gosk.

## Main additions
### Interfaces
Interfaces are defined to get the data from nmea0183 sentences, e.g. the interface `DepthBelowTransducer` with one method `GetDepthBelowTransducer`. All sentences that can produce this data should implement this interface. All data is represented in SI units according to https://signalk.org/specification/1.5.0/doc/data_model.html.

### Custom Float64 and Int64 types
The Float64 and Int64 types have an extra boolean property `isSet` that keeps track if the data is really available. If the sentence had an empty string then isSet will be `false` and the value will be `0`.
