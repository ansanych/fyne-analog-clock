package main

import (
	"image/color"
	"math"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

func main() {
	var radius float32 = 300
	rSecond := radius * 0.45
	rMinute := radius * 0.35
	rHour := radius * 0.25
	rTable := radius * 0.475
	rTable5 := rTable - 10
	rTable1 := rTable - 5

	myApp := app.New()
	w := myApp.NewWindow("Analog Clock")

	centerX, centerY := radius/2, radius/2

	secondHand := canvas.NewLine(color.White)
	secondHand.StrokeWidth = 2
	secondHand.Position1.X = centerX
	secondHand.Position1.Y = centerY

	minuteHand := canvas.NewLine(color.White)
	minuteHand.StrokeWidth = 4
	minuteHand.Position1.X = centerX
	minuteHand.Position1.Y = centerY

	hourHand := canvas.NewLine(color.White)
	hourHand.StrokeWidth = 6
	hourHand.Position1.X = centerX
	hourHand.Position1.Y = centerY

	clockContainer := container.NewWithoutLayout()
	clockContainer.Add(secondHand)
	clockContainer.Add(minuteHand)
	clockContainer.Add(hourHand)

	tableContainer := container.NewWithoutLayout()
	tableAngle := (math.Pi / 30)

	for i := range 60 {

		tableMark := canvas.NewLine(color.White)

		if i%5 == 0 {
			tableMark.StrokeWidth = 5
			tableMark.Position1 = getPos(centerX, centerY, rTable5, tableAngle*float64(i))
		} else {
			tableMark.StrokeWidth = 1
			tableMark.Position1 = getPos(centerX, centerY, rTable1, tableAngle*float64(i))
		}

		tableMark.Position2 = getPos(centerX, centerY, rTable, tableAngle*float64(i))

		tableContainer.Add(tableMark)
	}

	w.SetContent(container.NewStack(clockContainer, tableContainer))

	go func() {
		for {
			now := time.Now()
			second := float64(now.Second())
			minute := float64(now.Minute()) + second/60
			hour := float64(now.Hour()%12) + minute/60

			angleSecond := (math.Pi / 30) * second
			angleMinute := (math.Pi / 30) * minute
			angleHour := (math.Pi / 6) * hour

			fyne.Do(func() {
				secondHand.Position2 = getPos(centerX, centerY, rSecond, angleSecond)
				minuteHand.Position2 = getPos(centerX, centerY, rMinute, angleMinute)
				hourHand.Position2 = getPos(centerX, centerY, rHour, angleHour)

				clockContainer.Refresh()
			})

			time.Sleep(100 * time.Millisecond)
		}
	}()

	w.Resize(fyne.NewSize(radius, radius))
	w.ShowAndRun()
}

func getPos(cx float32, cy float32, r float32, ang float64) fyne.Position {
	pos := fyne.Position{
		X: cx + r*float32(math.Sin(ang)),
		Y: cy - r*float32(math.Cos(ang)),
	}

	return pos
}
