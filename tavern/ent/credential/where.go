// Code generated by entc, DO NOT EDIT.

package credential

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/kcarretto/realm/tavern/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Credential {
	return predicate.Credential(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Credential {
	return predicate.Credential(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Credential {
	return predicate.Credential(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Credential {
	return predicate.Credential(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Credential {
	return predicate.Credential(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Credential {
	return predicate.Credential(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Credential {
	return predicate.Credential(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Credential {
	return predicate.Credential(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Credential {
	return predicate.Credential(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Principal applies equality check predicate on the "principal" field. It's identical to PrincipalEQ.
func Principal(v string) predicate.Credential {
	return predicate.Credential(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPrincipal), v))
	})
}

// Secret applies equality check predicate on the "secret" field. It's identical to SecretEQ.
func Secret(v string) predicate.Credential {
	return predicate.Credential(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSecret), v))
	})
}

// PrincipalEQ applies the EQ predicate on the "principal" field.
func PrincipalEQ(v string) predicate.Credential {
	return predicate.Credential(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPrincipal), v))
	})
}

// PrincipalNEQ applies the NEQ predicate on the "principal" field.
func PrincipalNEQ(v string) predicate.Credential {
	return predicate.Credential(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPrincipal), v))
	})
}

// PrincipalIn applies the In predicate on the "principal" field.
func PrincipalIn(vs ...string) predicate.Credential {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Credential(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldPrincipal), v...))
	})
}

// PrincipalNotIn applies the NotIn predicate on the "principal" field.
func PrincipalNotIn(vs ...string) predicate.Credential {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Credential(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldPrincipal), v...))
	})
}

// PrincipalGT applies the GT predicate on the "principal" field.
func PrincipalGT(v string) predicate.Credential {
	return predicate.Credential(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPrincipal), v))
	})
}

// PrincipalGTE applies the GTE predicate on the "principal" field.
func PrincipalGTE(v string) predicate.Credential {
	return predicate.Credential(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPrincipal), v))
	})
}

// PrincipalLT applies the LT predicate on the "principal" field.
func PrincipalLT(v string) predicate.Credential {
	return predicate.Credential(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPrincipal), v))
	})
}

// PrincipalLTE applies the LTE predicate on the "principal" field.
func PrincipalLTE(v string) predicate.Credential {
	return predicate.Credential(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPrincipal), v))
	})
}

// PrincipalContains applies the Contains predicate on the "principal" field.
func PrincipalContains(v string) predicate.Credential {
	return predicate.Credential(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldPrincipal), v))
	})
}

// PrincipalHasPrefix applies the HasPrefix predicate on the "principal" field.
func PrincipalHasPrefix(v string) predicate.Credential {
	return predicate.Credential(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldPrincipal), v))
	})
}

// PrincipalHasSuffix applies the HasSuffix predicate on the "principal" field.
func PrincipalHasSuffix(v string) predicate.Credential {
	return predicate.Credential(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldPrincipal), v))
	})
}

// PrincipalEqualFold applies the EqualFold predicate on the "principal" field.
func PrincipalEqualFold(v string) predicate.Credential {
	return predicate.Credential(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldPrincipal), v))
	})
}

// PrincipalContainsFold applies the ContainsFold predicate on the "principal" field.
func PrincipalContainsFold(v string) predicate.Credential {
	return predicate.Credential(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldPrincipal), v))
	})
}

// SecretEQ applies the EQ predicate on the "secret" field.
func SecretEQ(v string) predicate.Credential {
	return predicate.Credential(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSecret), v))
	})
}

// SecretNEQ applies the NEQ predicate on the "secret" field.
func SecretNEQ(v string) predicate.Credential {
	return predicate.Credential(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSecret), v))
	})
}

// SecretIn applies the In predicate on the "secret" field.
func SecretIn(vs ...string) predicate.Credential {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Credential(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldSecret), v...))
	})
}

// SecretNotIn applies the NotIn predicate on the "secret" field.
func SecretNotIn(vs ...string) predicate.Credential {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Credential(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldSecret), v...))
	})
}

// SecretGT applies the GT predicate on the "secret" field.
func SecretGT(v string) predicate.Credential {
	return predicate.Credential(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSecret), v))
	})
}

// SecretGTE applies the GTE predicate on the "secret" field.
func SecretGTE(v string) predicate.Credential {
	return predicate.Credential(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSecret), v))
	})
}

// SecretLT applies the LT predicate on the "secret" field.
func SecretLT(v string) predicate.Credential {
	return predicate.Credential(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSecret), v))
	})
}

// SecretLTE applies the LTE predicate on the "secret" field.
func SecretLTE(v string) predicate.Credential {
	return predicate.Credential(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSecret), v))
	})
}

// SecretContains applies the Contains predicate on the "secret" field.
func SecretContains(v string) predicate.Credential {
	return predicate.Credential(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldSecret), v))
	})
}

// SecretHasPrefix applies the HasPrefix predicate on the "secret" field.
func SecretHasPrefix(v string) predicate.Credential {
	return predicate.Credential(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldSecret), v))
	})
}

// SecretHasSuffix applies the HasSuffix predicate on the "secret" field.
func SecretHasSuffix(v string) predicate.Credential {
	return predicate.Credential(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldSecret), v))
	})
}

// SecretEqualFold applies the EqualFold predicate on the "secret" field.
func SecretEqualFold(v string) predicate.Credential {
	return predicate.Credential(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldSecret), v))
	})
}

// SecretContainsFold applies the ContainsFold predicate on the "secret" field.
func SecretContainsFold(v string) predicate.Credential {
	return predicate.Credential(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldSecret), v))
	})
}

// KindEQ applies the EQ predicate on the "kind" field.
func KindEQ(v Kind) predicate.Credential {
	return predicate.Credential(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldKind), v))
	})
}

// KindNEQ applies the NEQ predicate on the "kind" field.
func KindNEQ(v Kind) predicate.Credential {
	return predicate.Credential(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldKind), v))
	})
}

// KindIn applies the In predicate on the "kind" field.
func KindIn(vs ...Kind) predicate.Credential {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Credential(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldKind), v...))
	})
}

// KindNotIn applies the NotIn predicate on the "kind" field.
func KindNotIn(vs ...Kind) predicate.Credential {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Credential(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldKind), v...))
	})
}

// HasTarget applies the HasEdge predicate on the "target" edge.
func HasTarget() predicate.Credential {
	return predicate.Credential(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(TargetTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, TargetTable, TargetColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTargetWith applies the HasEdge predicate on the "target" edge with a given conditions (other predicates).
func HasTargetWith(preds ...predicate.Target) predicate.Credential {
	return predicate.Credential(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(TargetInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, TargetTable, TargetColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Credential) predicate.Credential {
	return predicate.Credential(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Credential) predicate.Credential {
	return predicate.Credential(func(s *sql.Selector) {
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
func Not(p predicate.Credential) predicate.Credential {
	return predicate.Credential(func(s *sql.Selector) {
		p(s.Not())
	})
}