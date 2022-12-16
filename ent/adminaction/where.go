// Code generated by ent, DO NOT EDIT.

package adminaction

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/vmkevv/rigelapi/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.AdminAction {
	return predicate.AdminAction(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.AdminAction {
	return predicate.AdminAction(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.AdminAction {
	return predicate.AdminAction(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.AdminAction {
	return predicate.AdminAction(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.AdminAction {
	return predicate.AdminAction(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.AdminAction {
	return predicate.AdminAction(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.AdminAction {
	return predicate.AdminAction(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.AdminAction {
	return predicate.AdminAction(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.AdminAction {
	return predicate.AdminAction(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Action applies equality check predicate on the "action" field. It's identical to ActionEQ.
func Action(v string) predicate.AdminAction {
	return predicate.AdminAction(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAction), v))
	})
}

// Info applies equality check predicate on the "info" field. It's identical to InfoEQ.
func Info(v string) predicate.AdminAction {
	return predicate.AdminAction(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldInfo), v))
	})
}

// ActionEQ applies the EQ predicate on the "action" field.
func ActionEQ(v string) predicate.AdminAction {
	return predicate.AdminAction(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAction), v))
	})
}

// ActionNEQ applies the NEQ predicate on the "action" field.
func ActionNEQ(v string) predicate.AdminAction {
	return predicate.AdminAction(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAction), v))
	})
}

// ActionIn applies the In predicate on the "action" field.
func ActionIn(vs ...string) predicate.AdminAction {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AdminAction(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldAction), v...))
	})
}

// ActionNotIn applies the NotIn predicate on the "action" field.
func ActionNotIn(vs ...string) predicate.AdminAction {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AdminAction(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldAction), v...))
	})
}

// ActionGT applies the GT predicate on the "action" field.
func ActionGT(v string) predicate.AdminAction {
	return predicate.AdminAction(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAction), v))
	})
}

// ActionGTE applies the GTE predicate on the "action" field.
func ActionGTE(v string) predicate.AdminAction {
	return predicate.AdminAction(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAction), v))
	})
}

// ActionLT applies the LT predicate on the "action" field.
func ActionLT(v string) predicate.AdminAction {
	return predicate.AdminAction(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAction), v))
	})
}

// ActionLTE applies the LTE predicate on the "action" field.
func ActionLTE(v string) predicate.AdminAction {
	return predicate.AdminAction(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAction), v))
	})
}

// ActionContains applies the Contains predicate on the "action" field.
func ActionContains(v string) predicate.AdminAction {
	return predicate.AdminAction(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldAction), v))
	})
}

// ActionHasPrefix applies the HasPrefix predicate on the "action" field.
func ActionHasPrefix(v string) predicate.AdminAction {
	return predicate.AdminAction(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldAction), v))
	})
}

// ActionHasSuffix applies the HasSuffix predicate on the "action" field.
func ActionHasSuffix(v string) predicate.AdminAction {
	return predicate.AdminAction(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldAction), v))
	})
}

// ActionEqualFold applies the EqualFold predicate on the "action" field.
func ActionEqualFold(v string) predicate.AdminAction {
	return predicate.AdminAction(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldAction), v))
	})
}

// ActionContainsFold applies the ContainsFold predicate on the "action" field.
func ActionContainsFold(v string) predicate.AdminAction {
	return predicate.AdminAction(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldAction), v))
	})
}

// InfoEQ applies the EQ predicate on the "info" field.
func InfoEQ(v string) predicate.AdminAction {
	return predicate.AdminAction(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldInfo), v))
	})
}

// InfoNEQ applies the NEQ predicate on the "info" field.
func InfoNEQ(v string) predicate.AdminAction {
	return predicate.AdminAction(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldInfo), v))
	})
}

// InfoIn applies the In predicate on the "info" field.
func InfoIn(vs ...string) predicate.AdminAction {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AdminAction(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldInfo), v...))
	})
}

// InfoNotIn applies the NotIn predicate on the "info" field.
func InfoNotIn(vs ...string) predicate.AdminAction {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AdminAction(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldInfo), v...))
	})
}

// InfoGT applies the GT predicate on the "info" field.
func InfoGT(v string) predicate.AdminAction {
	return predicate.AdminAction(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldInfo), v))
	})
}

// InfoGTE applies the GTE predicate on the "info" field.
func InfoGTE(v string) predicate.AdminAction {
	return predicate.AdminAction(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldInfo), v))
	})
}

// InfoLT applies the LT predicate on the "info" field.
func InfoLT(v string) predicate.AdminAction {
	return predicate.AdminAction(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldInfo), v))
	})
}

// InfoLTE applies the LTE predicate on the "info" field.
func InfoLTE(v string) predicate.AdminAction {
	return predicate.AdminAction(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldInfo), v))
	})
}

// InfoContains applies the Contains predicate on the "info" field.
func InfoContains(v string) predicate.AdminAction {
	return predicate.AdminAction(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldInfo), v))
	})
}

// InfoHasPrefix applies the HasPrefix predicate on the "info" field.
func InfoHasPrefix(v string) predicate.AdminAction {
	return predicate.AdminAction(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldInfo), v))
	})
}

// InfoHasSuffix applies the HasSuffix predicate on the "info" field.
func InfoHasSuffix(v string) predicate.AdminAction {
	return predicate.AdminAction(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldInfo), v))
	})
}

// InfoEqualFold applies the EqualFold predicate on the "info" field.
func InfoEqualFold(v string) predicate.AdminAction {
	return predicate.AdminAction(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldInfo), v))
	})
}

// InfoContainsFold applies the ContainsFold predicate on the "info" field.
func InfoContainsFold(v string) predicate.AdminAction {
	return predicate.AdminAction(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldInfo), v))
	})
}

// HasTeacher applies the HasEdge predicate on the "teacher" edge.
func HasTeacher() predicate.AdminAction {
	return predicate.AdminAction(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(TeacherTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, TeacherTable, TeacherColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTeacherWith applies the HasEdge predicate on the "teacher" edge with a given conditions (other predicates).
func HasTeacherWith(preds ...predicate.Teacher) predicate.AdminAction {
	return predicate.AdminAction(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(TeacherInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, TeacherTable, TeacherColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.AdminAction) predicate.AdminAction {
	return predicate.AdminAction(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.AdminAction) predicate.AdminAction {
	return predicate.AdminAction(func(s *sql.Selector) {
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
func Not(p predicate.AdminAction) predicate.AdminAction {
	return predicate.AdminAction(func(s *sql.Selector) {
		p(s.Not())
	})
}