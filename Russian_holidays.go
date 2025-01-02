package goical

import (
	"fmt"
	"io"
	"net/url"
	"time"
)

func RussianHolidays(tz *time.Location, w io.Writer) error {
	holidays := New(tz)
	now := time.Now()
	holidays.AddEvent(Event{
		UID:         fmt.Sprintf("new_year_%v", now.Year()),
		Timestamp:   time.Now(),
		Summary:     "Happy New year",
		Description: fmt.Sprintf("Celebration of beginning of %v year", now.Year()),
		Location:    "your home",
		Start:       time.Date(now.Year(), time.January, 1, 0, 0, 0, 0, time.Local),
		End:         time.Date(now.Year(), time.January, 1, 23, 59, 59, 0, time.Local),
	})

	holidays.AddEvent(Event{
		UID:         fmt.Sprintf("christmas_%v", now.Year()),
		Timestamp:   time.Now(),
		Summary:     "Рождество Христово",
		Description: "Рождество Христово",
		Start:       time.Date(now.Year(), time.January, 7, 0, 0, 0, 0, time.Local),
		End:         time.Date(now.Year(), time.January, 7, 23, 59, 59, 0, time.Local),
	})

	holidays.AddEvent(Event{
		UID:         fmt.Sprintf("motherland_protector_%v", now.Year()),
		Timestamp:   time.Now(),
		Summary:     "День Защитника Отечества",
		Description: "Поздравляем всех причастных",
		Location:    "везде",
		Start:       time.Date(now.Year(), time.February, 23, 0, 0, 0, 0, time.Local),
		End:         time.Date(now.Year(), time.February, 23, 23, 59, 59, 0, time.Local),
	})

	holidays.AddEvent(Event{
		UID:         fmt.Sprintf("womans_day_%v", now.Year()),
		Timestamp:   time.Now(),
		Summary:     "Международный Женский День",
		Description: "Поздравляем всех причастных",
		Location:    "везде",
		Start:       time.Date(now.Year(), time.March, 8, 0, 0, 0, 0, time.Local),
		End:         time.Date(now.Year(), time.March, 8, 23, 59, 59, 0, time.Local),
	})

	holidays.AddEvent(Event{
		UID:         fmt.Sprintf("cosmonautics_day%v", now.Year()),
		Timestamp:   time.Now(),
		Summary:     "День Космонавтики",
		Description: "12 апреля 1961 года советский космонавт Юрий Гагарин совершил первый полёт в Космос, сделав 3 витка по орбите вокруг Земли на космическом корабле Восток 1",
		Location:    "Байканур",
		Start:       time.Date(now.Year(), time.April, 12, 0, 0, 0, 0, time.Local),
		End:         time.Date(now.Year(), time.April, 12, 23, 59, 59, 0, time.Local),
	})

	holidays.AddEvent(Event{
		UID:         fmt.Sprintf("labour_day_%v", now.Year()),
		Timestamp:   time.Now(),
		Summary:     "День Весны и Труда",
		Description: "День Весны и Труда",
		Start:       time.Date(now.Year(), time.May, 1, 0, 0, 0, 0, time.Local),
		End:         time.Date(now.Year(), time.May, 1, 23, 59, 00, 0, time.Local),
	})

	vd_url, _ := url.Parse("https://ruxpert.ru/%D0%9C%D0%B8%D1%84%D1%8B_%D0%BE_%D0%92%D0%B5%D0%BB%D0%B8%D0%BA%D0%BE%D0%B9_%D0%9E%D1%82%D0%B5%D1%87%D0%B5%D1%81%D1%82%D0%B2%D0%B5%D0%BD%D0%BD%D0%BE%D0%B9_%D0%B2%D0%BE%D0%B9%D0%BD%D0%B5")
	holidays.AddEvent(Event{
		UID:         fmt.Sprintf("Victory_day_%v", now.Year()),
		Timestamp:   time.Now(),
		Summary:     "День Победы",
		Description: "В ознаменование победоносного завершения Великой Отечественной войны советского народа против немецко-фашистских захватчиков и одержанных исторических побед Красной Армии, увенчавшихся полным разгромом гитлеровской Германии, заявившей о безоговорочной капитуляции, установить, что 9 мая является днём всенародного торжества — ПРАЗДНИКОМ ПОБЕДЫ.",
		URL:         vd_url,
		Location:    "Москва, Красная Площадь",
		Organizer: Person{
			CommonName: "Иосиф Сталин",
			Email:      "stalin@kremlin.ru",
		},
		Start: time.Date(now.Year(), time.May, 9, 9, 0, 0, 0, time.Local),
		End:   time.Date(now.Year(), time.May, 9, 11, 00, 00, 0, time.Local),
	})

	rd_url, _ := url.Parse("https://ruxpert.ru/%D0%A0%D0%BE%D1%81%D1%81%D0%B8%D1%8F_%D0%B7%D0%B0%D0%BD%D0%B8%D0%BC%D0%B0%D0%B5%D1%82_%D0%BF%D0%B5%D1%80%D0%B2%D0%BE%D0%B5_%D0%BC%D0%B5%D1%81%D1%82%D0%BE_%D0%B2_%D0%BC%D0%B8%D1%80%D0%B5")
	holidays.AddEvent(Event{
		UID:         fmt.Sprintf("Russia_day_%v", now.Year()),
		Timestamp:   time.Now(),
		Summary:     "День России",
		Description: "Горжусь Россией!",
		URL:         rd_url,
		Start:       time.Date(now.Year(), time.June, 12, 0, 0, 0, 0, time.Local),
		End:         time.Date(now.Year(), time.June, 12, 23, 59, 59, 0, time.Local),
	})

	holidays.AddEvent(Event{
		UID:         fmt.Sprintf("unity_day_%v", now.Year()),
		Timestamp:   time.Now(),
		Summary:     "День Народного Единства",
		Description: "В 1612 году польско-литовские захватчики были изгнаны из Кремля вторым народным ополчением под руководством Минина И Пожарского",
		Start:       time.Date(now.Year(), time.November, 4, 0, 0, 0, 0, time.Local),
		End:         time.Date(now.Year(), time.November, 4, 23, 59, 59, 0, time.Local),
	})

	holidays.AddEvent(Event{
		UID:         fmt.Sprintf("revolution_day_%v", now.Year()),
		Timestamp:   time.Now(),
		Summary:     "День Октябрьской Революции",
		Description: "В ночь с 7 по 8 ноября в Петрограде было свергнуто временное правительство и провозглашена Власть Советов",
		Start:       time.Date(now.Year(), time.November, 7, 0, 0, 0, 0, time.Local),
		End:         time.Date(now.Year(), time.November, 7, 23, 59, 59, 0, time.Local),
	})

	err := holidays.Render(w)
	if err != nil {
		fmt.Errorf("error rendering calendar: %w", err)
	}
	return nil
}
