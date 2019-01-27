package eventstore

import (
	"encoding/json"
	"time"
	"database/sql"
	"github.com/pkg/errors"
	"github.com/innoxchain/ixstorage/config"
	"github.com/innoxchain/ixstorage/pkg/apps/ixservice/domain/event"
	log "github.com/sirupsen/logrus"

	_ "github.com/lib/pq"
)

const JSON_DB_CONFIG_PATH = "../../../../config/local_db.json"

type DbConfig struct {
	ConnectString string
	DatabaseName string
	SslMode string
}

type EventStore struct{}

var db *sql.DB

func init() {
	cfg := DbConfig{}
	err := config.LoadConfig(JSON_DB_CONFIG_PATH, &cfg)

	if(err!=nil) {
		log.Fatal("Couldn't load JSON database configuration")
	}

	log.WithFields(log.Fields{
		"ConnectString": cfg.ConnectString,
		"DatabaseName": cfg.DatabaseName,
		"SslMode": cfg.SslMode,
	  }).Info("DB Configuration")

	db, err = sql.Open("postgres", cfg.ConnectString + "/" + cfg.DatabaseName + "?sslmode=" + cfg.SslMode)

	log.Info("Connected to Database")

	if(err!=nil) {
		log.Fatal("Error connecting to database: ", err)
	}

	createTables()
}

func createTables() {
	if _, err := db.Exec(
		"CREATE TABLE IF NOT EXISTS events (event_seq INT, aggregateid STRING, eventtype STRING, aggregatetype STRING, eventdata STRING, creationtime TIMESTAMP, PRIMARY KEY (event_seq, aggregateid))");
		err != nil {
			log.Fatal(err)
		}

	if _, err := db.Exec(
		"CREATE TABLE IF NOT EXISTS snapshots (aggregateid STRING, snapshot_event_seq INT, aggregatestate STRING, PRIMARY KEY (aggregateid, snapshot_event_seq))");
			err != nil {
				log.Fatal(err)
			}
}

func (es EventStore) CreateEvent(event_seq int, aggregateid, eventtype, aggregatetype, data string, creationtime time.Time) error {
	sql := `
		INSERT INTO events (event_seq, aggregateid, eventtype, aggregatetype, eventdata, creationtime) 
		VALUES ($1, $2, $3, $4, $5, $6)`
	
		_, err := db.Exec(sql, event_seq, aggregateid, eventtype, aggregatetype, data, creationtime)

		if(err!=nil) {
			return errors.Wrap(err, "error occured when inserting event")
		}
	return nil
}

func (es EventStore) CreateSnapshot(aggregateid, aggregatestate string, snapshot_event_seq int) error {
	sql := `
		INSERT INTO snapshots (aggregateid, snapshot_event_seq, aggregatestate) 
		VALUES ($1, $2, $3)`
	
		_, err := db.Exec(sql, aggregateid, snapshot_event_seq, aggregatestate)

		if(err!=nil) {
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
		aggregate=aggregatestate
	}
	return aggregate
}

func (es EventStore) GetEventsForAggregate(aggregateid string, eventSeq int) []event.DomainEvent {

	events := []event.DomainEvent{}

	rows, err := db.Query("SELECT event_seq, aggregateid, aggregatetype, eventtype, eventdata, creationtime FROM events where aggregateid=$1 and event_seq>$2", aggregateid, eventSeq)

	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var event_seq, aggregateid, aggregatetype, eventtype, eventdata, creationtime string
		if err := rows.Scan(&event_seq, &aggregateid, &aggregatetype, &eventtype, &eventdata, &creationtime); err != nil {
			log.Fatal(err)
		}

		/*
		pe := event.PersistentEvent{
			AggregateID: aggregateid,
			AggregateType: aggregatetype,
			EventType: eventtype,
			CreatedAt: creationtime,
			RawData: eventdata,
		}

		e, err := pe.Deserialize()
		if err!=nil {
			log.Fatal("couldn't deserialize PersistentEvent: ", pe)
		}
		evs = append(evs, e)
		*/

		switch eventtype {
			case "OrderCreated":
				deserializedEvent := &event.OrderCreatedEvent{}
				err := json.Unmarshal([]byte(eventdata), deserializedEvent)
				if(err!=nil) {
					log.Fatal("Error deserializing event! ", err)
				}
				
				events = append(events, 
					&event.OrderCreatedEvent{
						Event: event.Event{AggregateID: aggregateid, EventType: eventtype, CreatedAt: creationtime}, 
						Capacity: deserializedEvent.Capacity})

			case "OrderConfirmed":
				deserializedEvent := &event.OrderConfirmedEvent{}
				err := json.Unmarshal([]byte(eventdata), deserializedEvent)
				if(err!=nil) {
					log.Fatal("Error deserializing event! ", err)
				}

				events = append(events, 
					&event.OrderConfirmedEvent{
						Event: event.Event{AggregateID: aggregateid, EventType: eventtype, CreatedAt: creationtime}, 
						ConfirmedBy: deserializedEvent.ConfirmedBy})
		}
	}
	return events
}

func (es EventStore) GetEvents() []string {
	rs := []string{}

	rows, err := db.Query("SELECT event_seq, aggregateid, eventtype, aggregatetype, eventdata, creationtime FROM events")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var event_seq, aggregateid, eventtype, aggregatetype, eventdata, creationtime string
		if err := rows.Scan(&event_seq, &aggregateid, &eventtype, &aggregatetype, &eventdata, &creationtime); err != nil {
			log.Fatal(err)
		}
		//switch eventtype {
		//case "order.created":
		//	log.Fatal("Found Eventtype oder.created")
		//}
		rs = append(rs, event_seq)
		rs = append(rs, aggregateid)
		rs = append(rs, eventtype)
		rs = append(rs, aggregatetype)
		rs = append(rs, eventdata)
		rs = append(rs, string(creationtime))
	}
	return rs
}