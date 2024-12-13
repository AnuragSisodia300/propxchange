package schema

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID            primitive.ObjectID   `json:"id" bson:"_id"`
	Name          string               `json:"name" bson:"name"`
	Email         string               `json:"email" bson:"email"`
	WalletBalance float64              `json:"wallet_balance" bson:"wallet_balance"`
	Purchased     []UserProperty       `json:"purchased" bson:"purchased"`
	Favorites     []primitive.ObjectID `json:"favorites" bson:"favorites"`
	KYC           KYCInfo              `json:"kyc" bson:"kyc"`
}

type UserProperty struct {
	PropertyID primitive.ObjectID `json:"property_id" bson:"property_id"`
	Units      int                `json:"units" bson:"units"`
	AmountPaid float64            `json:"amount_paid" bson:"amount_paid"`
	PaymentID  string             `json:"payment_id" bson:"payment_id"`
}

type KYCInfo struct {
	PersonalInfo    PersonalInfo    `json:"personal_info" bson:"personal_info"`
	ResidentialInfo ResidentialInfo `json:"residential_info" bson:"residential_info"`
	FinancialInfo   FinancialInfo   `json:"financial_info" bson:"financial_info"`
	Documents       Documents       `json:"documents" bson:"documents"`
	IsCompleted     bool            `json:"is_completed" bson:"is_completed"`
}

type PersonalInfo struct {
	FirstName   string `json:"first_name" bson:"first_name"`
	LastName    string `json:"last_name" bson:"last_name"`
	DateOfBirth string `json:"dob" bson:"dob"`
	Email       string `json:"email" bson:"email"`
	PhoneNumber string `json:"phone_number" bson:"phone_number"`
}

type ResidentialInfo struct {
	Address    string `json:"address" bson:"address"`
	StreetName string `json:"street_name" bson:"street_name"`
	City       string `json:"city" bson:"city"`
	State      string `json:"state" bson:"state"`
	PostalCode string `json:"postal_code" bson:"postal_code"`
	Country    string `json:"country" bson:"country"`
}

type FinancialInfo struct {
	IncomeSource string  `json:"income_source" bson:"income_source"`
	AnnualIncome float64 `json:"annual_income" bson:"annual_income"`
}

type Documents struct {
	IDProof      string `json:"id_proof" bson:"id_proof"`
	AddressProof string `json:"address_proof" bson:"address_proof"`
	IncomeProof  string `json:"income_proof" bson:"income_proof"`
}
