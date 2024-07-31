package internal

import (
	"database/sql"
	"errors"
	"log"
)

func (app *MyApp) GetPersonById(id string) (data Person, err error) {

	query := `select p.name, ph.number as phone_number, 
	ad.city, ad.state, ad.street1, ad.street2, ad.zip_code
	from person p 
	left join phone ph on p.id = ph.person_id
	left join address_join aj on aj.person_id = p.id
	left join address ad on ad.id = aj.address_id where p.id =?
	`
	row := app.DB.QueryRow(query, id)
	err = row.Scan(&data.Name, &data.PhoneNumber, &data.City, &data.State, &data.Street1, &data.Street2, &data.ZipCode)
	if err == sql.ErrNoRows {
		return data, errors.New(`Person not found`)
	} else if err != nil {
		return data, err
	}
	return data, nil
}

func (app *MyApp) CreatePersonData(data Person) error {

	var (
		personId  int64
		addressId int64
	)

	txn, _ := app.DB.Begin()
	defer txn.Rollback()

	personQuery := `insert into person(name) values (?)`
	phoneQuery := `insert into phone( person_id, number) values (?,?)`
	addressQuery := `insert into address(city , state , street1 , street2 , zip_code) values (?,?,?,?,?)`
	addressJoinQuery := `insert into address_join (person_id, address_id) values (?,?)`

	query, err := txn.Exec(personQuery, data.Name)
	if err != nil {
		log.Println(err)
		return err
	}
	personId, _ = query.LastInsertId()

	_, err = txn.Exec(phoneQuery, personId, data.PhoneNumber)

	if err != nil {
		log.Println(err)
		return err
	}

	query, err = txn.Exec(addressQuery, data.City, data.State, data.Street1, data.Street2, data.ZipCode)
	if err != nil {
		log.Println(err)
		return err
	}
	addressId, _ = query.LastInsertId()

	_, err = txn.Exec(addressJoinQuery, personId, addressId)
	if err != nil {
		log.Println(err)
		return err
	}

	txn.Commit()
	return nil
}
