// Code generated by ent, DO NOT EDIT.

package classperiodsync

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/vmkevv/rigelapi/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.ClassPeriodSync {
	return predicate.ClassPeriodSync(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.ClassPeriodSync {
	return predicate.ClassPeriodSync(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.ClassPeriodSync {
	return predicate.ClassPeriodSync(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.ClassPeriodSync {
	return predicate.ClassPeriodSync(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.ClassPeriodSync {
	return predicate.ClassPeriodSync(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.ClassPeriodSync {
	return predicate.ClassPeriodSync(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.ClassPeriodSync {
	return predicate.ClassPeriodSync(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.ClassPeriodSync {
	return predicate.ClassPeriodSync(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.ClassPeriodSync {
	return predicate.ClassPeriodSync(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// LastSyncID applies equality check predicate on the "last_sync_id" field. It's identical to LastSyncIDEQ.
func LastSyncID(v string) predicate.ClassPeriodSync {
	return predicate.ClassPeriodSync(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLastSyncID), v))
	})
}

// LastSyncIDEQ applies the EQ predicate on the "last_sync_id" field.
func LastSyncIDEQ(v string) predicate.ClassPeriodSync {
	return predicate.ClassPeriodSync(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLastSyncID), v))
	})
}

// LastSyncIDNEQ applies the NEQ predicate on the "last_sync_id" field.
func LastSyncIDNEQ(v string) predicate.ClassPeriodSync {
	return predicate.ClassPeriodSync(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldLastSyncID), v))
	})
}

// LastSyncIDIn applies the In predicate on the "last_sync_id" field.
func LastSyncIDIn(vs ...string) predicate.ClassPeriodSync {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ClassPeriodSync(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldLastSyncID), v...))
	})
}

// LastSyncIDNotIn applies the NotIn predicate on the "last_sync_id" field.
func LastSyncIDNotIn(vs ...string) predicate.ClassPeriodSync {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ClassPeriodSync(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldLastSyncID), v...))
	})
}

// LastSyncIDGT applies the GT predicate on the "last_sync_id" field.
func LastSyncIDGT(v string) predicate.ClassPeriodSync {
	return predicate.ClassPeriodSync(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldLastSyncID), v))
	})
}

// LastSyncIDGTE applies the GTE predicate on the "last_sync_id" field.
func LastSyncIDGTE(v string) predicate.ClassPeriodSync {
	return predicate.ClassPeriodSync(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldLastSyncID), v))
	})
}

// LastSyncIDLT applies the LT predicate on the "last_sync_id" field.
func LastSyncIDLT(v string) predicate.ClassPeriodSync {
	return predicate.ClassPeriodSync(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldLastSyncID), v))
	})
}

// LastSyncIDLTE applies the LTE predicate on the "last_sync_id" field.
func LastSyncIDLTE(v string) predicate.ClassPeriodSync {
	return predicate.ClassPeriodSync(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldLastSyncID), v))
	})
}

// LastSyncIDContains applies the Contains predicate on the "last_sync_id" field.
func LastSyncIDContains(v string) predicate.ClassPeriodSync {
	return predicate.ClassPeriodSync(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldLastSyncID), v))
	})
}

// LastSyncIDHasPrefix applies the HasPrefix predicate on the "last_sync_id" field.
func LastSyncIDHasPrefix(v string) predicate.ClassPeriodSync {
	return predicate.ClassPeriodSync(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldLastSyncID), v))
	})
}

// LastSyncIDHasSuffix applies the HasSuffix predicate on the "last_sync_id" field.
func LastSyncIDHasSuffix(v string) predicate.ClassPeriodSync {
	return predicate.ClassPeriodSync(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldLastSyncID), v))
	})
}

// LastSyncIDEqualFold applies the EqualFold predicate on the "last_sync_id" field.
func LastSyncIDEqualFold(v string) predicate.ClassPeriodSync {
	return predicate.ClassPeriodSync(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldLastSyncID), v))
	})
}

// LastSyncIDContainsFold applies the ContainsFold predicate on the "last_sync_id" field.
func LastSyncIDContainsFold(v string) predicate.ClassPeriodSync {
	return predicate.ClassPeriodSync(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldLastSyncID), v))
	})
}

// HasClass applies the HasEdge predicate on the "class" edge.
func HasClass() predicate.ClassPeriodSync {
	return predicate.ClassPeriodSync(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ClassTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ClassTable, ClassColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasClassWith applies the HasEdge predicate on the "class" edge with a given conditions (other predicates).
func HasClassWith(preds ...predicate.Class) predicate.ClassPeriodSync {
	return predicate.ClassPeriodSync(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ClassInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ClassTable, ClassColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.ClassPeriodSync) predicate.ClassPeriodSync {
	return predicate.ClassPeriodSync(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.ClassPeriodSync) predicate.ClassPeriodSync {
	return predicate.ClassPeriodSync(func(s *sql.Selector) {
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
func Not(p predicate.ClassPeriodSync) predicate.ClassPeriodSync {
	return predicate.ClassPeriodSync(func(s *sql.Selector) {
		p(s.Not())
	})
}