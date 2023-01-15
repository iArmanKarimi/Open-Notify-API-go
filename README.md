# Open Notify API - Go

Open Nofity API client for `Go`

> Open Notify is an open source project to provide a simple programming interface for some of NASA’s awesome data.

For other languages, see [Open Notify API clients](https://github.com/iArmanKarimi/Open-Notify-API-clients)

## Installation

```bash

```

## Examples

Number of People in Space:

```go
ppl, err := GetPeopleInSpace()
fmt.Println("People in space right now:")
for _, v := range ppl.People {
    fmt.Printf("Craft: %s, Name: %s\n", v.Craft, v.Name)
}
fmt.Printf("Number of people in space: %d\n", ppl.Number)
```

Current Location of the International Space Station:

```go
iss, err := GetISSLocation()
fmt.Println("current location of ISS:")
fmt.Printf("latitude: %f\n", iss.Location.Latitude)
fmt.Printf("longitude: %f\n", iss.Location.Longitude)
```

## References

[Open Notify Website](http://open-notify.org/)

[Official API documentation](http://open-notify.org/Open-Notify-API/)

## License

[MIT](https://github.com/iArmanKarimi/Open-Notify-API-go/blob/main/LICENSE)
