package eventstore

import (
	"time"
	"database/sql"
	"github.com/pkg/errors"
	"github.com/innoxchain/ixstorage/config"
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
	if(err!=nil) {
		log.Fatal("Error connecting to database: ", err)
	}

	createTables()
}

func createTables() {
	if _, err := db.Exec(
		"CREATE TABLE IF NOT EXISTS events (event_seq INT PRIMARY KEY, aggregateid STRING, eventtype STRING, aggregatetype STRING, eventdata STRING, creationtime TIMESTAMP)");
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
		rs = append(rs, event_seq)
		rs = append(rs, aggregateid)
		rs = append(rs, eventtype)
		rs = append(rs, aggregatetype)
		rs = append(rs, eventdata)
		rs = append(rs, string(creationtime))
	}
	return rs
}