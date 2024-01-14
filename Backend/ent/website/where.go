// Code generated by ent, DO NOT EDIT.

package website

import (
	"Backend/ent/predicate"

	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Website {
	return predicate.Website(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Website {
	return predicate.Website(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Website {
	return predicate.Website(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Website {
	return predicate.Website(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Website {
	return predicate.Website(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Website {
	return predicate.Website(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Website {
	return predicate.Website(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Website {
	return predicate.Website(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Website {
	return predicate.Website(sql.FieldLTE(FieldID, id))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Website {
	return predicate.Website(sql.FieldEQ(FieldName, v))
}

// FilePath applies equality check predicate on the "filePath" field. It's identical to FilePathEQ.
func FilePath(v string) predicate.Website {
	return predicate.Website(sql.FieldEQ(FieldFilePath, v))
}

// URL applies equality check predicate on the "url" field. It's identical to URLEQ.
func URL(v string) predicate.Website {
	return predicate.Website(sql.FieldEQ(FieldURL, v))
}

// Flashy applies equality check predicate on the "flashy" field. It's identical to FlashyEQ.
func Flashy(v float64) predicate.Website {
	return predicate.Website(sql.FieldEQ(FieldFlashy, v))
}

// Adult applies equality check predicate on the "adult" field. It's identical to AdultEQ.
func Adult(v float64) predicate.Website {
	return predicate.Website(sql.FieldEQ(FieldAdult, v))
}

// Smart applies equality check predicate on the "smart" field. It's identical to SmartEQ.
func Smart(v float64) predicate.Website {
	return predicate.Website(sql.FieldEQ(FieldSmart, v))
}

// Beautiful applies equality check predicate on the "beautiful" field. It's identical to BeautifulEQ.
func Beautiful(v float64) predicate.Website {
	return predicate.Website(sql.FieldEQ(FieldBeautiful, v))
}

// Like applies equality check predicate on the "like" field. It's identical to LikeEQ.
func Like(v float64) predicate.Website {
	return predicate.Website(sql.FieldEQ(FieldLike, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Website {
	return predicate.Website(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Website {
	return predicate.Website(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Website {
	return predicate.Website(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Website {
	return predicate.Website(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Website {
	return predicate.Website(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Website {
	return predicate.Website(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Website {
	return predicate.Website(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Website {
	return predicate.Website(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Website {
	return predicate.Website(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Website {
	return predicate.Website(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Website {
	return predicate.Website(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Website {
	return predicate.Website(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Website {
	return predicate.Website(sql.FieldContainsFold(FieldName, v))
}

// FilePathEQ applies the EQ predicate on the "filePath" field.
func FilePathEQ(v string) predicate.Website {
	return predicate.Website(sql.FieldEQ(FieldFilePath, v))
}

// FilePathNEQ applies the NEQ predicate on the "filePath" field.
func FilePathNEQ(v string) predicate.Website {
	return predicate.Website(sql.FieldNEQ(FieldFilePath, v))
}

// FilePathIn applies the In predicate on the "filePath" field.
func FilePathIn(vs ...string) predicate.Website {
	return predicate.Website(sql.FieldIn(FieldFilePath, vs...))
}

// FilePathNotIn applies the NotIn predicate on the "filePath" field.
func FilePathNotIn(vs ...string) predicate.Website {
	return predicate.Website(sql.FieldNotIn(FieldFilePath, vs...))
}

// FilePathGT applies the GT predicate on the "filePath" field.
func FilePathGT(v string) predicate.Website {
	return predicate.Website(sql.FieldGT(FieldFilePath, v))
}

// FilePathGTE applies the GTE predicate on the "filePath" field.
func FilePathGTE(v string) predicate.Website {
	return predicate.Website(sql.FieldGTE(FieldFilePath, v))
}

// FilePathLT applies the LT predicate on the "filePath" field.
func FilePathLT(v string) predicate.Website {
	return predicate.Website(sql.FieldLT(FieldFilePath, v))
}

// FilePathLTE applies the LTE predicate on the "filePath" field.
func FilePathLTE(v string) predicate.Website {
	return predicate.Website(sql.FieldLTE(FieldFilePath, v))
}

// FilePathContains applies the Contains predicate on the "filePath" field.
func FilePathContains(v string) predicate.Website {
	return predicate.Website(sql.FieldContains(FieldFilePath, v))
}

// FilePathHasPrefix applies the HasPrefix predicate on the "filePath" field.
func FilePathHasPrefix(v string) predicate.Website {
	return predicate.Website(sql.FieldHasPrefix(FieldFilePath, v))
}

// FilePathHasSuffix applies the HasSuffix predicate on the "filePath" field.
func FilePathHasSuffix(v string) predicate.Website {
	return predicate.Website(sql.FieldHasSuffix(FieldFilePath, v))
}

// FilePathEqualFold applies the EqualFold predicate on the "filePath" field.
func FilePathEqualFold(v string) predicate.Website {
	return predicate.Website(sql.FieldEqualFold(FieldFilePath, v))
}

// FilePathContainsFold applies the ContainsFold predicate on the "filePath" field.
func FilePathContainsFold(v string) predicate.Website {
	return predicate.Website(sql.FieldContainsFold(FieldFilePath, v))
}

// URLEQ applies the EQ predicate on the "url" field.
func URLEQ(v string) predicate.Website {
	return predicate.Website(sql.FieldEQ(FieldURL, v))
}

// URLNEQ applies the NEQ predicate on the "url" field.
func URLNEQ(v string) predicate.Website {
	return predicate.Website(sql.FieldNEQ(FieldURL, v))
}

// URLIn applies the In predicate on the "url" field.
func URLIn(vs ...string) predicate.Website {
	return predicate.Website(sql.FieldIn(FieldURL, vs...))
}

// URLNotIn applies the NotIn predicate on the "url" field.
func URLNotIn(vs ...string) predicate.Website {
	return predicate.Website(sql.FieldNotIn(FieldURL, vs...))
}

// URLGT applies the GT predicate on the "url" field.
func URLGT(v string) predicate.Website {
	return predicate.Website(sql.FieldGT(FieldURL, v))
}

// URLGTE applies the GTE predicate on the "url" field.
func URLGTE(v string) predicate.Website {
	return predicate.Website(sql.FieldGTE(FieldURL, v))
}

// URLLT applies the LT predicate on the "url" field.
func URLLT(v string) predicate.Website {
	return predicate.Website(sql.FieldLT(FieldURL, v))
}

// URLLTE applies the LTE predicate on the "url" field.
func URLLTE(v string) predicate.Website {
	return predicate.Website(sql.FieldLTE(FieldURL, v))
}

// URLContains applies the Contains predicate on the "url" field.
func URLContains(v string) predicate.Website {
	return predicate.Website(sql.FieldContains(FieldURL, v))
}

// URLHasPrefix applies the HasPrefix predicate on the "url" field.
func URLHasPrefix(v string) predicate.Website {
	return predicate.Website(sql.FieldHasPrefix(FieldURL, v))
}

// URLHasSuffix applies the HasSuffix predicate on the "url" field.
func URLHasSuffix(v string) predicate.Website {
	return predicate.Website(sql.FieldHasSuffix(FieldURL, v))
}

// URLEqualFold applies the EqualFold predicate on the "url" field.
func URLEqualFold(v string) predicate.Website {
	return predicate.Website(sql.FieldEqualFold(FieldURL, v))
}

// URLContainsFold applies the ContainsFold predicate on the "url" field.
func URLContainsFold(v string) predicate.Website {
	return predicate.Website(sql.FieldContainsFold(FieldURL, v))
}

// FlashyEQ applies the EQ predicate on the "flashy" field.
func FlashyEQ(v float64) predicate.Website {
	return predicate.Website(sql.FieldEQ(FieldFlashy, v))
}

// FlashyNEQ applies the NEQ predicate on the "flashy" field.
func FlashyNEQ(v float64) predicate.Website {
	return predicate.Website(sql.FieldNEQ(FieldFlashy, v))
}

// FlashyIn applies the In predicate on the "flashy" field.
func FlashyIn(vs ...float64) predicate.Website {
	return predicate.Website(sql.FieldIn(FieldFlashy, vs...))
}

// FlashyNotIn applies the NotIn predicate on the "flashy" field.
func FlashyNotIn(vs ...float64) predicate.Website {
	return predicate.Website(sql.FieldNotIn(FieldFlashy, vs...))
}

// FlashyGT applies the GT predicate on the "flashy" field.
func FlashyGT(v float64) predicate.Website {
	return predicate.Website(sql.FieldGT(FieldFlashy, v))
}

// FlashyGTE applies the GTE predicate on the "flashy" field.
func FlashyGTE(v float64) predicate.Website {
	return predicate.Website(sql.FieldGTE(FieldFlashy, v))
}

// FlashyLT applies the LT predicate on the "flashy" field.
func FlashyLT(v float64) predicate.Website {
	return predicate.Website(sql.FieldLT(FieldFlashy, v))
}

// FlashyLTE applies the LTE predicate on the "flashy" field.
func FlashyLTE(v float64) predicate.Website {
	return predicate.Website(sql.FieldLTE(FieldFlashy, v))
}

// AdultEQ applies the EQ predicate on the "adult" field.
func AdultEQ(v float64) predicate.Website {
	return predicate.Website(sql.FieldEQ(FieldAdult, v))
}

// AdultNEQ applies the NEQ predicate on the "adult" field.
func AdultNEQ(v float64) predicate.Website {
	return predicate.Website(sql.FieldNEQ(FieldAdult, v))
}

// AdultIn applies the In predicate on the "adult" field.
func AdultIn(vs ...float64) predicate.Website {
	return predicate.Website(sql.FieldIn(FieldAdult, vs...))
}

// AdultNotIn applies the NotIn predicate on the "adult" field.
func AdultNotIn(vs ...float64) predicate.Website {
	return predicate.Website(sql.FieldNotIn(FieldAdult, vs...))
}

// AdultGT applies the GT predicate on the "adult" field.
func AdultGT(v float64) predicate.Website {
	return predicate.Website(sql.FieldGT(FieldAdult, v))
}

// AdultGTE applies the GTE predicate on the "adult" field.
func AdultGTE(v float64) predicate.Website {
	return predicate.Website(sql.FieldGTE(FieldAdult, v))
}

// AdultLT applies the LT predicate on the "adult" field.
func AdultLT(v float64) predicate.Website {
	return predicate.Website(sql.FieldLT(FieldAdult, v))
}

// AdultLTE applies the LTE predicate on the "adult" field.
func AdultLTE(v float64) predicate.Website {
	return predicate.Website(sql.FieldLTE(FieldAdult, v))
}

// SmartEQ applies the EQ predicate on the "smart" field.
func SmartEQ(v float64) predicate.Website {
	return predicate.Website(sql.FieldEQ(FieldSmart, v))
}

// SmartNEQ applies the NEQ predicate on the "smart" field.
func SmartNEQ(v float64) predicate.Website {
	return predicate.Website(sql.FieldNEQ(FieldSmart, v))
}

// SmartIn applies the In predicate on the "smart" field.
func SmartIn(vs ...float64) predicate.Website {
	return predicate.Website(sql.FieldIn(FieldSmart, vs...))
}

// SmartNotIn applies the NotIn predicate on the "smart" field.
func SmartNotIn(vs ...float64) predicate.Website {
	return predicate.Website(sql.FieldNotIn(FieldSmart, vs...))
}

// SmartGT applies the GT predicate on the "smart" field.
func SmartGT(v float64) predicate.Website {
	return predicate.Website(sql.FieldGT(FieldSmart, v))
}

// SmartGTE applies the GTE predicate on the "smart" field.
func SmartGTE(v float64) predicate.Website {
	return predicate.Website(sql.FieldGTE(FieldSmart, v))
}

// SmartLT applies the LT predicate on the "smart" field.
func SmartLT(v float64) predicate.Website {
	return predicate.Website(sql.FieldLT(FieldSmart, v))
}

// SmartLTE applies the LTE predicate on the "smart" field.
func SmartLTE(v float64) predicate.Website {
	return predicate.Website(sql.FieldLTE(FieldSmart, v))
}

// BeautifulEQ applies the EQ predicate on the "beautiful" field.
func BeautifulEQ(v float64) predicate.Website {
	return predicate.Website(sql.FieldEQ(FieldBeautiful, v))
}

// BeautifulNEQ applies the NEQ predicate on the "beautiful" field.
func BeautifulNEQ(v float64) predicate.Website {
	return predicate.Website(sql.FieldNEQ(FieldBeautiful, v))
}

// BeautifulIn applies the In predicate on the "beautiful" field.
func BeautifulIn(vs ...float64) predicate.Website {
	return predicate.Website(sql.FieldIn(FieldBeautiful, vs...))
}

// BeautifulNotIn applies the NotIn predicate on the "beautiful" field.
func BeautifulNotIn(vs ...float64) predicate.Website {
	return predicate.Website(sql.FieldNotIn(FieldBeautiful, vs...))
}

// BeautifulGT applies the GT predicate on the "beautiful" field.
func BeautifulGT(v float64) predicate.Website {
	return predicate.Website(sql.FieldGT(FieldBeautiful, v))
}

// BeautifulGTE applies the GTE predicate on the "beautiful" field.
func BeautifulGTE(v float64) predicate.Website {
	return predicate.Website(sql.FieldGTE(FieldBeautiful, v))
}

// BeautifulLT applies the LT predicate on the "beautiful" field.
func BeautifulLT(v float64) predicate.Website {
	return predicate.Website(sql.FieldLT(FieldBeautiful, v))
}

// BeautifulLTE applies the LTE predicate on the "beautiful" field.
func BeautifulLTE(v float64) predicate.Website {
	return predicate.Website(sql.FieldLTE(FieldBeautiful, v))
}

// LikeEQ applies the EQ predicate on the "like" field.
func LikeEQ(v float64) predicate.Website {
	return predicate.Website(sql.FieldEQ(FieldLike, v))
}

// LikeNEQ applies the NEQ predicate on the "like" field.
func LikeNEQ(v float64) predicate.Website {
	return predicate.Website(sql.FieldNEQ(FieldLike, v))
}

// LikeIn applies the In predicate on the "like" field.
func LikeIn(vs ...float64) predicate.Website {
	return predicate.Website(sql.FieldIn(FieldLike, vs...))
}

// LikeNotIn applies the NotIn predicate on the "like" field.
func LikeNotIn(vs ...float64) predicate.Website {
	return predicate.Website(sql.FieldNotIn(FieldLike, vs...))
}

// LikeGT applies the GT predicate on the "like" field.
func LikeGT(v float64) predicate.Website {
	return predicate.Website(sql.FieldGT(FieldLike, v))
}

// LikeGTE applies the GTE predicate on the "like" field.
func LikeGTE(v float64) predicate.Website {
	return predicate.Website(sql.FieldGTE(FieldLike, v))
}

// LikeLT applies the LT predicate on the "like" field.
func LikeLT(v float64) predicate.Website {
	return predicate.Website(sql.FieldLT(FieldLike, v))
}

// LikeLTE applies the LTE predicate on the "like" field.
func LikeLTE(v float64) predicate.Website {
	return predicate.Website(sql.FieldLTE(FieldLike, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Website) predicate.Website {
	return predicate.Website(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Website) predicate.Website {
	return predicate.Website(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Website) predicate.Website {
	return predicate.Website(sql.NotPredicates(p))
}
