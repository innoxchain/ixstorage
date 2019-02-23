package event

import (
	"strconv"
	"database/sql"
	"github.com/innoxchain/ixstorage/config"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"time"

	_ "github.com/lib/pq"
)

const JSON_DB_CONFIG_PATH = "../../../../../config/local_db.json"

type DbConfig struct {
	ConnectString string
	DatabaseName  string
	SslMode       string
}

type EventStore struct{}

var db *sql.DB

func init() {
	cfg := DbConfig{}
	err := config.LoadConfig(JSON_DB_CONFIG_PATH, &cfg)

	if err != nil {
		log.Fatal("Couldn't load JSON database configuration")
	}

	log.WithFields(log.Fields{
		"ConnectString": cfg.ConnectString,
		"DatabaseName":  cfg.DatabaseName,
		"SslMode":       cfg.SslMode,
	}).Info("DB Configuration")

	db, err = sql.Open("postgres", cfg.ConnectString+"/"+cfg.DatabaseName+"?sslmode="+cfg.SslMode)

	log.Info("Connected to Database")

	if err != nil {
		log.Fatal("Error connecting to database: ", err)
	}

	createTables()
}

func createTables() {
	if _, err := db.Exec(
		"CREATE TABLE IF NOT EXISTS events (event_seq INT, aggregateid STRING, eventtype STRING, aggregatetype STRING, eventdata STRING, creationtime TIMESTAMP, PRIMARY KEY (event_seq, aggregateid))"); err != nil {
		log.Fatal(err)
	}

	if _, err := db.Exec(
		"CREATE TABLE IF NOT EXISTS snapshots (aggregateid STRING, snapshot_event_seq INT, aggregatestate STRING, PRIMARY KEY (aggregateid, snapshot_event_seq))"); err != nil {
		log.Fatal(err)
	}
}

//CreateEvent is deprecated
func (es EventStore) CreateEvent(event_seq int, aggregateid, eventtype, aggregatetype, data string, creationtime time.Time) error {
	sql := `
		INSERT INTO events (event_seq, aggregateid, eventtype, aggregatetype, eventdata, creationtime) 
		VALUES ($1, $2, $3, $4, $5, $6)`

	_, err := db.Exec(sql, event_seq, aggregateid, eventtype, aggregatetype, data, creationtime)

	if err != nil {
		return errors.Wrap(err, "error occured when inserting event")
	}
	return nil
}

func (es EventStore) Persist(aggregate Aggregate) error {

	sql := `
	INSERT INTO events (event_seq, aggregateid, eventtype, aggregatetype, eventdata, creationtime) 
	VALUES ($1, $2, $3, $4, $5, $6)`

	for _, e := range aggregate.GetChanges() {
		persistentEvent, error := e.Serialize()

		if error != nil {
			log.Fatal(error)
		}

		log.Info("Persisting Event: ")
		log.Info("persistentEvent.Sequence ", persistentEvent.Sequence)
		log.Info("persistentEvent.AggregateID ", persistentEvent.AggregateID)
		log.Info("persistentEvent.EventType ", persistentEvent.EventType)
		log.Info("persistentEvent.AggregateType ", persistentEvent.AggregateType)
		log.Info("persistentEvent.RawData ", persistentEvent.RawData)
		log.Info("persistentEvent.CreatedAt ", persistentEvent.CreatedAt)
		
		_, err := db.Exec(sql, persistentEvent.Sequence, persistentEvent.AggregateID, persistentEvent.EventType, persistentEvent.AggregateType, persistentEvent.RawData, persistentEvent.CreatedAt)

		if err != nil {
			return errors.Wrap(err, "error occured when inserting event")
		}
	}

	aggregate.MarkAsCommited()

	return nil
}

//PersistEvent persists an Event in the event store
func (es EventStore) PersistEvent(e Event) error {
	sql := `
		INSERT INTO events (event_seq, aggregateid, eventtype, aggregatetype, eventdata, creationtime) 
		VALUES ($1, $2, $3, $4, $5, $6)`

	persistentEvent, error := e.Serialize()

	if error != nil {
		log.Fatal(error)
	}

	_, err := db.Exec(sql, persistentEvent.Sequence, persistentEvent.AggregateID, persistentEvent.EventType, persistentEvent.AggregateType, persistentEvent.RawData, persistentEvent.CreatedAt)

	if err != nil {
		return errors.Wrap(err, "error occured when inserting event")
	}
	return nil
}

func (es EventStore) CreateSnapshot(aggregateid, aggregatestate string, snapshot_event_seq int) error {
	sql := `
		INSERT INTO snapshots (aggregateid, snapshot_event_seq, aggregatestate) 
		VALUES ($1, $2, $3)`

	_, err := db.Exec(sql, aggregateid, snapshot_event_seq, aggregatestate)

	if err != nil {
		return errors.Wrap(err, "error occured when inserting snapshot")
	}
	return nil
}

func (es EventStore) GetSnapshot(aggregateid string) string {
	aggregate := ""

	rows, err := db.Query("select aggregatestate from snapshots where aggregateid=$1 order by snapshot_event_seq desc limit 1", aggregateid)

	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var aggregatestate string
		if err := rows.Scan(&aggregatestate); err != nil {
			log.Fatal(err)
		}
		aggregate = aggregatestate
	}
	return aggregate
}

func (es EventStore) EventsForAggregate(aggregateid string, eventSeq int) []Event {

	events := []Event{}

	rows, err := db.Query("SELECT event_seq, aggregateid, aggregatetype, eventtype, eventdata, creationtime FROM events where aggregateid=$1 and event_seq>$2", aggregateid, eventSeq)

	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var event_seq, aggregateid, aggregatetype, eventtype, eventdata string
		var creationtime time.Time
		if err := rows.Scan(&event_seq, &aggregateid, &aggregatetype, &eventtype, &eventdata, &creationtime); err != nil {
			log.Fatal(err)
		}

		seq, _ := strconv.Atoi(event_seq)

		pe := PersistentEvent{
			AggregateID:   aggregateid,
			AggregateType: aggregatetype,
			EventType:     eventtype,
			CreatedAt:     creationtime,
			Sequence:      seq,
			RawData:       eventdata,
		}

		e, err := pe.Deserialize()
		if err != nil {
			log.Fatal("couldn't deserialize PersistentEvent: ", pe)
		}

		events = append(events, e)
	}
	return events
}