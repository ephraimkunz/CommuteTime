package main

import (
	"reflect"
	"testing"
	"time"

	"github.com/wcharczuk/go-chart/util"
)

func Test_formatTime(t *testing.T) {
	now := time.Date(2018, time.August, 11, 12, 0, 0, 0, time.Now().Location())

	type args struct {
		v interface{}
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Time", args{now}, "Sat - 12:00 PM"},
		{"Other types", args{now.Unix()}, ""},
		{"Float", args{util.Time.ToFloat64(now)}, "Sat - 12:00 PM"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := formatTime(tt.args.v); got != tt.want {
				t.Errorf("formatTime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_roundToNearestHalfHourInFuture(t *testing.T) {
	now := time.Now()
	type args struct {
		now time.Time
	}
	tests := []struct {
		name string
		args args
		want time.Time
	}{
		{"New hour", args{time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), 0, 0, 0, now.Location())}, time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), 30, 0, 0, now.Location())},
		{"New hour + 5", args{time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), 5, 0, 0, now.Location())}, time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), 30, 0, 0, now.Location())},
		{"Just before 1/2 hour", args{time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), 29, 59, 0, now.Location())}, time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), 30, 0, 0, now.Location())},
		{"Just after 1/2 hour", args{time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), 30, 1, 0, now.Location())}, time.Date(now.Year(), now.Month(), now.Day(), now.Hour()+1, 0, 0, 0, now.Location())},
		{"Just before new hour", args{time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), 59, 59, 0, now.Location())}, time.Date(now.Year(), now.Month(), now.Day(), now.Hour()+1, 0, 0, 0, now.Location())},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := roundToNearestHalfHourInFuture(tt.args.now); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("roundToNearestHalfHourInFuture() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_timeOfDay(t *testing.T) {
	now := time.Now()
	noon := time.Date(now.Year(), now.Month(), now.Day(), 12, 0, 0, 0, now.Location())
	type args struct {
		now  time.Time
		noon time.Time
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Morning", args{time.Date(now.Year(), now.Month(), now.Day(), 9, 0, 0, 0, now.Location()), noon}, "this morning"},
		{"Afternoon", args{time.Date(now.Year(), now.Month(), now.Day(), 13, 0, 0, 0, now.Location()), noon}, "this afternoon"},
		{"Noon", args{time.Date(now.Year(), now.Month(), now.Day(), 12, 0, 0, 0, now.Location()), noon}, "this afternoon"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := timeOfDay(tt.args.now, tt.args.noon); got != tt.want {
				t.Errorf("timeOfDay() = %v, want %v", got, tt.want)
			}
		})
	}
}
