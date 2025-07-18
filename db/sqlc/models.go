// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.29.0

package db

import (
	"database/sql"
)

type Account struct {
	ID        int64        `json:"id"`
	Username  string       `json:"username"`
	Password  string       `json:"password"`
	Role      string       `json:"role"`
	CreatedAt sql.NullTime `json:"created_at"`
}

type Address struct {
	ID        int64        `json:"id"`
	Name      string       `json:"name"`
	Phone     string       `json:"phone"`
	Address   string       `json:"address"`
	Status    string       `json:"status"`
	CreatedAt sql.NullTime `json:"created_at"`
}

type ApiKey struct {
	ID        int64         `json:"id"`
	ClientID  sql.NullInt64 `json:"client_id"`
	ApiKey    string        `json:"api_key"`
	CreatedAt sql.NullTime  `json:"created_at"`
	UpdatedAt sql.NullTime  `json:"updated_at"`
}

type Client struct {
	ID           int64         `json:"id"`
	Name         string        `json:"name"`
	AccountID    sql.NullInt32 `json:"account_id"`
	ContactEmail string        `json:"contact_email"`
	CreatedAt    sql.NullTime  `json:"created_at"`
}

type Shipment struct {
	ID            int64          `json:"id"`
	ClientID      sql.NullInt64  `json:"client_id"`
	FromAddressID sql.NullInt64  `json:"from_address_id"`
	ToAddressID   sql.NullInt64  `json:"to_address_id"`
	ShipperID     sql.NullInt64  `json:"shipper_id"`
	ShipmentCode  sql.NullString `json:"shipment_code"`
	Fee           int32          `json:"fee"`
	// pending, accepted, in_transit, delivered, canceled
	Status    sql.NullString `json:"status"`
	CreatedAt sql.NullTime   `json:"created_at"`
	UpdatedAt sql.NullTime   `json:"updated_at"`
}

type ShipmentStatusLog struct {
	ID         int64         `json:"id"`
	ShipmentID sql.NullInt64 `json:"shipment_id"`
	Status     string        `json:"status"`
	Note       string        `json:"note"`
	CreatedAt  sql.NullTime  `json:"created_at"`
}

type Shipper struct {
	ID        int64        `json:"id"`
	Name      string       `json:"name"`
	Phone     string       `json:"phone"`
	Active    sql.NullBool `json:"active"`
	CreatedAt sql.NullTime `json:"created_at"`
}
