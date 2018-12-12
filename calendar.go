package ical

type Calendar struct {
	Items  []CalendarEvent
	ProdID string
	Name   string
	Method string
}

func (this *Calendar) Serialize() string {
	serializer := calSerializer{
		calendar: this,
		buffer:   new(strBuffer),
	}
	return serializer.serialize()
}

func (this *Calendar) ToICS() string {
	return this.Serialize()
}
