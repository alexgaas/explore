package console

import (
	"explore/config"
	"flag"
	"fmt"
	"math/rand"
	"time"

	"github.com/jedib0t/go-pretty/v6/progress"
	"github.com/jedib0t/go-pretty/v6/text"
)

var (
	flagAutoStop           = flag.Bool("auto-stop", true, "Auto-stop rendering?")
	flagHideETA            = flag.Bool("hide-eta", true, "Hide the ETA?")
	flagHideETAOverall     = flag.Bool("hide-eta-overall", true, "Hide the ETA in the overall tracker?")
	flagHideOverallTracker = flag.Bool("hide-overall", true, "Hide the Overall Tracker?")
	flagHidePercentage     = flag.Bool("hide-percentage", false, "Hide the console percent?")
	flagHideTime           = flag.Bool("hide-time", true, "Hide the time taken?")
	flagHideValue          = flag.Bool("hide-value", false, "Hide the tracker value?")
	flagNumTrackers        = flag.Int("num-trackers", 1, "Number of Trackers")
	flagShowSpeed          = flag.Bool("show-speed", false, "Show the tracker speed?")
	flagShowSpeedOverall   = flag.Bool("show-speed-overall", false, "Show the overall tracker speed?")
	flagShowPinned         = flag.Bool("show-pinned", false, "Show a pinned message?")
	flagRandomFail         = flag.Bool("rnd-fail", false, "Introduce random failures in tracking")
	flagRandomLogs         = flag.Bool("rnd-logs", false, "Output random logs in the middle of tracking")

	messageColors = []text.Color{
		text.FgGreen,
		text.FgWhite,
	}
	timeStart = time.Now()
)

func getMessage(idx int64, units *progress.Units, config *config.Config) string {
	var message string
	switch units {
	default:
		message = fmt.Sprintf("Compressing file \"%s\" using %v codecs %d times #", config.FilePath, config.Codecs, config.Count)
	}
	return message
}

func getUnits() *progress.Units {
	var units *progress.Units
	units = &progress.Units{
		Notation:         " iterations",
		Formatter:        nil,
		NotationPosition: progress.UnitsNotationPositionAfter,
	}
	return units
}

func trackExploration(pw progress.Writer, idx int64, updateMessage bool, config *config.Config) {
	total := idx * idx * idx * int64(config.Count)
	incrementPerCycle := idx * int64(*flagNumTrackers) * 100

	units := getUnits()
	message := getMessage(idx, units, config)
	tracker := progress.Tracker{Message: message, Total: total, Units: *units}
	if idx == int64(*flagNumTrackers) {
		tracker.Total = 0
	}

	pw.AppendTracker(&tracker)

	ticker := time.Tick(time.Millisecond * 500)
	updateTicker := time.Tick(time.Millisecond * 250)
	for !tracker.IsDone() {
		select {
		case <-ticker:
			tracker.Increment(incrementPerCycle)
			if idx == int64(*flagNumTrackers) && tracker.Value() >= total {
				tracker.MarkAsDone()
			} else if *flagRandomFail && rand.Float64() < 0.1 {
				tracker.MarkAsErrored()
			}
			pw.SetPinnedMessages(
				fmt.Sprintf(">> Current Time: %-32s", time.Now().Format(time.RFC3339)),
				fmt.Sprintf(">>   Total Time: %-32s", time.Now().Sub(timeStart).Round(time.Millisecond)),
			)
		case <-updateTicker:
			if updateMessage {
				rndIdx := rand.Intn(len(messageColors))
				if rndIdx == len(messageColors) {
					rndIdx--
				}
				tracker.UpdateMessage(messageColors[rndIdx].Sprint(message))
			}
		}
	}
}

func TrackProgress(pw progress.Writer, config *config.Config) {
	for idx := int64(1); idx <= int64(*flagNumTrackers); idx++ {
		go trackExploration(pw, idx, idx == int64(*flagNumTrackers), config)

		if !*flagAutoStop {
			time.Sleep(time.Millisecond * 100)
		}
	}
}

func StopProgress(pw progress.Writer) {
	time.Sleep(time.Second)
	messagesLogged := make(map[string]bool)
	for pw.IsRenderInProgress() {
		if *flagRandomLogs && pw.LengthDone()%3 == 0 {
			logMsg := text.Faint.Sprintf("[INFO] done with %d trackers", pw.LengthDone())
			if !messagesLogged[logMsg] {
				pw.Log(logMsg)
				messagesLogged[logMsg] = true
			}
		}

		if !*flagAutoStop && pw.LengthActive() == 0 {
			pw.Stop()
		}
		time.Sleep(time.Millisecond * 100)
	}
}

func InitProgress() progress.Writer {
	flag.Parse()

	pw := progress.NewWriter()
	pw.SetAutoStop(*flagAutoStop)
	pw.SetTrackerLength(25)
	pw.SetMessageWidth(125)
	pw.SetNumTrackersExpected(*flagNumTrackers)
	pw.SetSortBy(progress.SortByPercentDsc)
	pw.SetStyle(progress.StyleDefault)
	pw.SetTrackerPosition(progress.PositionRight)
	pw.SetUpdateFrequency(time.Millisecond * 100)
	pw.Style().Colors = progress.StyleColorsExample
	pw.Style().Options.PercentFormat = "%4.1f%%"
	pw.Style().Visibility.ETA = !*flagHideETA
	pw.Style().Visibility.ETAOverall = !*flagHideETAOverall
	pw.Style().Visibility.Percentage = !*flagHidePercentage
	pw.Style().Visibility.Speed = *flagShowSpeed
	pw.Style().Visibility.SpeedOverall = *flagShowSpeedOverall
	pw.Style().Visibility.Time = !*flagHideTime
	pw.Style().Visibility.TrackerOverall = !*flagHideOverallTracker
	pw.Style().Visibility.Value = !*flagHideValue
	pw.Style().Visibility.Pinned = *flagShowPinned

	go pw.Render()
	return pw
}
