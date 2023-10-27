// Code generated by ent, DO NOT EDIT.

package adminuser

import (
	"frame/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.AdminUser {
	return predicate.AdminUser(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.AdminUser {
	return predicate.AdminUser(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.AdminUser {
	return predicate.AdminUser(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.AdminUser {
	return predicate.AdminUser(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.AdminUser {
	return predicate.AdminUser(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.AdminUser {
	return predicate.AdminUser(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.AdminUser {
	return predicate.AdminUser(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.AdminUser {
	return predicate.AdminUser(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.AdminUser {
	return predicate.AdminUser(sql.FieldLTE(FieldID, id))
}

// CreateTime applies equality check predicate on the "create_time" field. It's identical to CreateTimeEQ.
func CreateTime(v time.Time) predicate.AdminUser {
	return predicate.AdminUser(sql.FieldEQ(FieldCreateTime, v))
}

// UpdateTime applies equality check predicate on the "update_time" field. It's identical to UpdateTimeEQ.
func UpdateTime(v time.Time) predicate.AdminUser {
	return predicate.AdminUser(sql.FieldEQ(FieldUpdateTime, v))
}

// WalletAddress applies equality check predicate on the "wallet_address" field. It's identical to WalletAddressEQ.
func WalletAddress(v string) predicate.AdminUser {
	return predicate.AdminUser(sql.FieldEQ(FieldWalletAddress, v))
}

// CreateTimeEQ applies the EQ predicate on the "create_time" field.
func CreateTimeEQ(v time.Time) predicate.AdminUser {
	return predicate.AdminUser(sql.FieldEQ(FieldCreateTime, v))
}

// CreateTimeNEQ applies the NEQ predicate on the "create_time" field.
func CreateTimeNEQ(v time.Time) predicate.AdminUser {
	return predicate.AdminUser(sql.FieldNEQ(FieldCreateTime, v))
}

// CreateTimeIn applies the In predicate on the "create_time" field.
func CreateTimeIn(vs ...time.Time) predicate.AdminUser {
	return predicate.AdminUser(sql.FieldIn(FieldCreateTime, vs...))
}

// CreateTimeNotIn applies the NotIn predicate on the "create_time" field.
func CreateTimeNotIn(vs ...time.Time) predicate.AdminUser {
	return predicate.AdminUser(sql.FieldNotIn(FieldCreateTime, vs...))
}

// CreateTimeGT applies the GT predicate on the "create_time" field.
func CreateTimeGT(v time.Time) predicate.AdminUser {
	return predicate.AdminUser(sql.FieldGT(FieldCreateTime, v))
}

// CreateTimeGTE applies the GTE predicate on the "create_time" field.
func CreateTimeGTE(v time.Time) predicate.AdminUser {
	return predicate.AdminUser(sql.FieldGTE(FieldCreateTime, v))
}

// CreateTimeLT applies the LT predicate on the "create_time" field.
func CreateTimeLT(v time.Time) predicate.AdminUser {
	return predicate.AdminUser(sql.FieldLT(FieldCreateTime, v))
}

// CreateTimeLTE applies the LTE predicate on the "create_time" field.
func CreateTimeLTE(v time.Time) predicate.AdminUser {
	return predicate.AdminUser(sql.FieldLTE(FieldCreateTime, v))
}

// UpdateTimeEQ applies the EQ predicate on the "update_time" field.
func UpdateTimeEQ(v time.Time) predicate.AdminUser {
	return predicate.AdminUser(sql.FieldEQ(FieldUpdateTime, v))
}

// UpdateTimeNEQ applies the NEQ predicate on the "update_time" field.
func UpdateTimeNEQ(v time.Time) predicate.AdminUser {
	return predicate.AdminUser(sql.FieldNEQ(FieldUpdateTime, v))
}

// UpdateTimeIn applies the In predicate on the "update_time" field.
func UpdateTimeIn(vs ...time.Time) predicate.AdminUser {
	return predicate.AdminUser(sql.FieldIn(FieldUpdateTime, vs...))
}

// UpdateTimeNotIn applies the NotIn predicate on the "update_time" field.
func UpdateTimeNotIn(vs ...time.Time) predicate.AdminUser {
	return predicate.AdminUser(sql.FieldNotIn(FieldUpdateTime, vs...))
}

// UpdateTimeGT applies the GT predicate on the "update_time" field.
func UpdateTimeGT(v time.Time) predicate.AdminUser {
	return predicate.AdminUser(sql.FieldGT(FieldUpdateTime, v))
}

// UpdateTimeGTE applies the GTE predicate on the "update_time" field.
func UpdateTimeGTE(v time.Time) predicate.AdminUser {
	return predicate.AdminUser(sql.FieldGTE(FieldUpdateTime, v))
}

// UpdateTimeLT applies the LT predicate on the "update_time" field.
func UpdateTimeLT(v time.Time) predicate.AdminUser {
	return predicate.AdminUser(sql.FieldLT(FieldUpdateTime, v))
}

// UpdateTimeLTE applies the LTE predicate on the "update_time" field.
func UpdateTimeLTE(v time.Time) predicate.AdminUser {
	return predicate.AdminUser(sql.FieldLTE(FieldUpdateTime, v))
}

// WalletAddressEQ applies the EQ predicate on the "wallet_address" field.
func WalletAddressEQ(v string) predicate.AdminUser {
	return predicate.AdminUser(sql.FieldEQ(FieldWalletAddress, v))
}

// WalletAddressNEQ applies the NEQ predicate on the "wallet_address" field.
func WalletAddressNEQ(v string) predicate.AdminUser {
	return predicate.AdminUser(sql.FieldNEQ(FieldWalletAddress, v))
}

// WalletAddressIn applies the In predicate on the "wallet_address" field.
func WalletAddressIn(vs ...string) predicate.AdminUser {
	return predicate.AdminUser(sql.FieldIn(FieldWalletAddress, vs...))
}

// WalletAddressNotIn applies the NotIn predicate on the "wallet_address" field.
func WalletAddressNotIn(vs ...string) predicate.AdminUser {
	return predicate.AdminUser(sql.FieldNotIn(FieldWalletAddress, vs...))
}

// WalletAddressGT applies the GT predicate on the "wallet_address" field.
func WalletAddressGT(v string) predicate.AdminUser {
	return predicate.AdminUser(sql.FieldGT(FieldWalletAddress, v))
}

// WalletAddressGTE applies the GTE predicate on the "wallet_address" field.
func WalletAddressGTE(v string) predicate.AdminUser {
	return predicate.AdminUser(sql.FieldGTE(FieldWalletAddress, v))
}

// WalletAddressLT applies the LT predicate on the "wallet_address" field.
func WalletAddressLT(v string) predicate.AdminUser {
	return predicate.AdminUser(sql.FieldLT(FieldWalletAddress, v))
}

// WalletAddressLTE applies the LTE predicate on the "wallet_address" field.
func WalletAddressLTE(v string) predicate.AdminUser {
	return predicate.AdminUser(sql.FieldLTE(FieldWalletAddress, v))
}

// WalletAddressContains applies the Contains predicate on the "wallet_address" field.
func WalletAddressContains(v string) predicate.AdminUser {
	return predicate.AdminUser(sql.FieldContains(FieldWalletAddress, v))
}

// WalletAddressHasPrefix applies the HasPrefix predicate on the "wallet_address" field.
func WalletAddressHasPrefix(v string) predicate.AdminUser {
	return predicate.AdminUser(sql.FieldHasPrefix(FieldWalletAddress, v))
}

// WalletAddressHasSuffix applies the HasSuffix predicate on the "wallet_address" field.
func WalletAddressHasSuffix(v string) predicate.AdminUser {
	return predicate.AdminUser(sql.FieldHasSuffix(FieldWalletAddress, v))
}

// WalletAddressEqualFold applies the EqualFold predicate on the "wallet_address" field.
func WalletAddressEqualFold(v string) predicate.AdminUser {
	return predicate.AdminUser(sql.FieldEqualFold(FieldWalletAddress, v))
}

// WalletAddressContainsFold applies the ContainsFold predicate on the "wallet_address" field.
func WalletAddressContainsFold(v string) predicate.AdminUser {
	return predicate.AdminUser(sql.FieldContainsFold(FieldWalletAddress, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.AdminUser) predicate.AdminUser {
	return predicate.AdminUser(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.AdminUser) predicate.AdminUser {
	return predicate.AdminUser(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.AdminUser) predicate.AdminUser {
	return predicate.AdminUser(sql.NotPredicates(p))
}