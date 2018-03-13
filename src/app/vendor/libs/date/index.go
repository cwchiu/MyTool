package date

import (
    "time"
    "fmt"
)

type Date struct {
    Data time.Time
}

const (
    FORMAT  = "2006-01-02T15:04:05-0700"
    ISO8601 = "2006-01-02T15:04:05Z0700"
    RFC850  = "Monday, 02-Jan-06 15:04:05 MST"
    RFC1036 = "Mon, 02 Jan 06 15:04:05 Z0700"
    RFC2822 = "Mon, 02 Jan 2006 15:04:05 Z0700"
    RFC3339 = "2006-01-02T15:04:05Z07:00"
    RFC822  = RFC1036
    RSS     = RFC2822
    COOKIE  = RFC850
    W3C     = RFC3339
    ATOM    = RFC3339
)

func FromTime(t time.Time) (*Date, error){
    return &Date{Data: t}, nil
}

func FromUnix(v int64) (*Date, error){
    dt := time.Unix(v, 0)
    return &Date{Data: dt}, nil
}

func FromString(s string) (*Date, error){
    
    dt, err := time.Parse(FORMAT, s)
    if err != nil {
        return nil, err
    }
    return &Date{Data: dt}, nil
}

func (this Date) Iso8601() string {
    return this.Data.UTC().Format(ISO8601)
}

func (this Date) Rfc850() string {
    return this.Data.UTC().Format(RFC850)
}

func (this Date) Rfc822() string {
    return this.Data.UTC().Format(RFC822)
}

func (this Date) Rfc2822() string {
    return this.Data.UTC().Format(RFC2822)
}

func (this Date) Rfc1123() string {
    return this.Data.UTC().Format(time.RFC1123)
}

func (this Date) Rfc1036() string {
    return this.Data.UTC().Format(RFC1036)
}

func (this Date) Rfc3339() string {
    return this.Data.UTC().Format(RFC3339)
}


func (this Date) Unix() string {
    return this.Data.UTC().Format(time.UnixDate)
}

func (this Date) Ruby() string {
    return this.Data.UTC().Format(time.RubyDate)
}

func (this Date) AnsiC() string {
    return this.Data.UTC().Format(time.ANSIC)
}

func (this Date) Cookie() string {
    return this.Data.UTC().Format(COOKIE)
}

func (this Date) Epoch() string {
    return fmt.Sprintf("%v", this.Data.UTC().Unix())
}

func (this Date) Local() string {
    return this.Data.Local().Format(ISO8601)
}