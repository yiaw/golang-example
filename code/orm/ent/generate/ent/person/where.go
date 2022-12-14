// Code generated by ent, DO NOT EDIT.

package person

import (
	"ent/generate/ent/predicate"

	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.Person {
	return predicate.Person(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.Person {
	return predicate.Person(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.Person {
	return predicate.Person(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.Person {
	return predicate.Person(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.Person {
	return predicate.Person(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.Person {
	return predicate.Person(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.Person {
	return predicate.Person(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.Person {
	return predicate.Person(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.Person {
	return predicate.Person(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// FirstName applies equality check predicate on the "first_name" field. It's identical to FirstNameEQ.
func FirstName(v string) predicate.Person {
	return predicate.Person(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFirstName), v))
	})
}

// SecondName applies equality check predicate on the "second_name" field. It's identical to SecondNameEQ.
func SecondName(v string) predicate.Person {
	return predicate.Person(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSecondName), v))
	})
}

// FirstNameEQ applies the EQ predicate on the "first_name" field.
func FirstNameEQ(v string) predicate.Person {
	return predicate.Person(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFirstName), v))
	})
}

// FirstNameNEQ applies the NEQ predicate on the "first_name" field.
func FirstNameNEQ(v string) predicate.Person {
	return predicate.Person(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldFirstName), v))
	})
}

// FirstNameIn applies the In predicate on the "first_name" field.
func FirstNameIn(vs ...string) predicate.Person {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Person(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldFirstName), v...))
	})
}

// FirstNameNotIn applies the NotIn predicate on the "first_name" field.
func FirstNameNotIn(vs ...string) predicate.Person {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Person(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldFirstName), v...))
	})
}

// FirstNameGT applies the GT predicate on the "first_name" field.
func FirstNameGT(v string) predicate.Person {
	return predicate.Person(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldFirstName), v))
	})
}

// FirstNameGTE applies the GTE predicate on the "first_name" field.
func FirstNameGTE(v string) predicate.Person {
	return predicate.Person(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldFirstName), v))
	})
}

// FirstNameLT applies the LT predicate on the "first_name" field.
func FirstNameLT(v string) predicate.Person {
	return predicate.Person(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldFirstName), v))
	})
}

// FirstNameLTE applies the LTE predicate on the "first_name" field.
func FirstNameLTE(v string) predicate.Person {
	return predicate.Person(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldFirstName), v))
	})
}

// FirstNameContains applies the Contains predicate on the "first_name" field.
func FirstNameContains(v string) predicate.Person {
	return predicate.Person(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldFirstName), v))
	})
}

// FirstNameHasPrefix applies the HasPrefix predicate on the "first_name" field.
func FirstNameHasPrefix(v string) predicate.Person {
	return predicate.Person(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldFirstName), v))
	})
}

// FirstNameHasSuffix applies the HasSuffix predicate on the "first_name" field.
func FirstNameHasSuffix(v string) predicate.Person {
	return predicate.Person(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldFirstName), v))
	})
}

// FirstNameEqualFold applies the EqualFold predicate on the "first_name" field.
func FirstNameEqualFold(v string) predicate.Person {
	return predicate.Person(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldFirstName), v))
	})
}

// FirstNameContainsFold applies the ContainsFold predicate on the "first_name" field.
func FirstNameContainsFold(v string) predicate.Person {
	return predicate.Person(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldFirstName), v))
	})
}

// SecondNameEQ applies the EQ predicate on the "second_name" field.
func SecondNameEQ(v string) predicate.Person {
	return predicate.Person(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSecondName), v))
	})
}

// SecondNameNEQ applies the NEQ predicate on the "second_name" field.
func SecondNameNEQ(v string) predicate.Person {
	return predicate.Person(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSecondName), v))
	})
}

// SecondNameIn applies the In predicate on the "second_name" field.
func SecondNameIn(vs ...string) predicate.Person {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Person(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldSecondName), v...))
	})
}

// SecondNameNotIn applies the NotIn predicate on the "second_name" field.
func SecondNameNotIn(vs ...string) predicate.Person {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Person(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldSecondName), v...))
	})
}

// SecondNameGT applies the GT predicate on the "second_name" field.
func SecondNameGT(v string) predicate.Person {
	return predicate.Person(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSecondName), v))
	})
}

// SecondNameGTE applies the GTE predicate on the "second_name" field.
func SecondNameGTE(v string) predicate.Person {
	return predicate.Person(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSecondName), v))
	})
}

// SecondNameLT applies the LT predicate on the "second_name" field.
func SecondNameLT(v string) predicate.Person {
	return predicate.Person(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSecondName), v))
	})
}

// SecondNameLTE applies the LTE predicate on the "second_name" field.
func SecondNameLTE(v string) predicate.Person {
	return predicate.Person(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSecondName), v))
	})
}

// SecondNameContains applies the Contains predicate on the "second_name" field.
func SecondNameContains(v string) predicate.Person {
	return predicate.Person(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldSecondName), v))
	})
}

// SecondNameHasPrefix applies the HasPrefix predicate on the "second_name" field.
func SecondNameHasPrefix(v string) predicate.Person {
	return predicate.Person(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldSecondName), v))
	})
}

// SecondNameHasSuffix applies the HasSuffix predicate on the "second_name" field.
func SecondNameHasSuffix(v string) predicate.Person {
	return predicate.Person(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldSecondName), v))
	})
}

// SecondNameEqualFold applies the EqualFold predicate on the "second_name" field.
func SecondNameEqualFold(v string) predicate.Person {
	return predicate.Person(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldSecondName), v))
	})
}

// SecondNameContainsFold applies the ContainsFold predicate on the "second_name" field.
func SecondNameContainsFold(v string) predicate.Person {
	return predicate.Person(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldSecondName), v))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Person) predicate.Person {
	return predicate.Person(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Person) predicate.Person {
	return predicate.Person(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Person) predicate.Person {
	return predicate.Person(func(s *sql.Selector) {
		p(s.Not())
	})
}
