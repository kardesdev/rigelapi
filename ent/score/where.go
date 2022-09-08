// Code generated by ent, DO NOT EDIT.

package score

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/vmkevv/rigelapi/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.Score {
	return predicate.Score(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.Score {
	return predicate.Score(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.Score {
	return predicate.Score(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.Score {
	return predicate.Score(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.Score {
	return predicate.Score(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.Score {
	return predicate.Score(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.Score {
	return predicate.Score(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.Score {
	return predicate.Score(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.Score {
	return predicate.Score(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Points applies equality check predicate on the "points" field. It's identical to PointsEQ.
func Points(v int) predicate.Score {
	return predicate.Score(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPoints), v))
	})
}

// PointsEQ applies the EQ predicate on the "points" field.
func PointsEQ(v int) predicate.Score {
	return predicate.Score(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPoints), v))
	})
}

// PointsNEQ applies the NEQ predicate on the "points" field.
func PointsNEQ(v int) predicate.Score {
	return predicate.Score(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPoints), v))
	})
}

// PointsIn applies the In predicate on the "points" field.
func PointsIn(vs ...int) predicate.Score {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Score(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldPoints), v...))
	})
}

// PointsNotIn applies the NotIn predicate on the "points" field.
func PointsNotIn(vs ...int) predicate.Score {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Score(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldPoints), v...))
	})
}

// PointsGT applies the GT predicate on the "points" field.
func PointsGT(v int) predicate.Score {
	return predicate.Score(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPoints), v))
	})
}

// PointsGTE applies the GTE predicate on the "points" field.
func PointsGTE(v int) predicate.Score {
	return predicate.Score(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPoints), v))
	})
}

// PointsLT applies the LT predicate on the "points" field.
func PointsLT(v int) predicate.Score {
	return predicate.Score(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPoints), v))
	})
}

// PointsLTE applies the LTE predicate on the "points" field.
func PointsLTE(v int) predicate.Score {
	return predicate.Score(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPoints), v))
	})
}

// HasActivity applies the HasEdge predicate on the "activity" edge.
func HasActivity() predicate.Score {
	return predicate.Score(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ActivityTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ActivityTable, ActivityColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasActivityWith applies the HasEdge predicate on the "activity" edge with a given conditions (other predicates).
func HasActivityWith(preds ...predicate.Activity) predicate.Score {
	return predicate.Score(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ActivityInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ActivityTable, ActivityColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasStudent applies the HasEdge predicate on the "student" edge.
func HasStudent() predicate.Score {
	return predicate.Score(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(StudentTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, StudentTable, StudentColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasStudentWith applies the HasEdge predicate on the "student" edge with a given conditions (other predicates).
func HasStudentWith(preds ...predicate.Student) predicate.Score {
	return predicate.Score(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(StudentInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, StudentTable, StudentColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Score) predicate.Score {
	return predicate.Score(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Score) predicate.Score {
	return predicate.Score(func(s *sql.Selector) {
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
func Not(p predicate.Score) predicate.Score {
	return predicate.Score(func(s *sql.Selector) {
		p(s.Not())
	})
}