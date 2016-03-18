package sqlparser

import (
	"fmt"
	"testing"
)

func TestVisitSelect(t *testing.T) {
	sql := "select * from persons"
	stat, _ := Parse(sql)
	r := new(recorder)
	StatementAccept(stat, r)
	if got, want := fmt.Sprintf("%v", r.seen), "[Select StarExpr AliasedTableExpr]"; got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

type recorder struct {
	seen []string
}

func (r *recorder) VisitSelect(n *Select) {
	r.seen = append(r.seen, "Select")
	for _, each := range n.SelectExprs {
		SelectExprAccept(each, r)
	}
	for _, each := range n.From {
		TableExprAccept(each, r)
	}
}
func (r *recorder) VisitUnion(n *Union) {
	r.seen = append(r.seen, "Union")
}
func (r *recorder) VisitInsert(n *Insert) {
	r.seen = append(r.seen, "Insert")
}
func (r *recorder) VisitUpdate(n *Update) {
	r.seen = append(r.seen, "Update")
}
func (r *recorder) VisitDelete(n *Delete) {
	r.seen = append(r.seen, "Delete")
}
func (r *recorder) VisitSet(n *Set) {
	r.seen = append(r.seen, "Set")
}
func (r *recorder) VisitDDL(n *DDL) {
	r.seen = append(r.seen, "DDL")
}
func (r *recorder) VisitColumnDefinition(n *ColumnDefinition) {
	r.seen = append(r.seen, "ColumnDefinition")
}
func (r *recorder) VisitCreateTable(n *CreateTable) {
	r.seen = append(r.seen, "CreateTable")
}
func (r *recorder) VisitOther(n *Other) {
	r.seen = append(r.seen, "Other")
}
func (r *recorder) VisitStarExpr(n *StarExpr) {
	r.seen = append(r.seen, "StarExpr")
}
func (r *recorder) VisitNonStarExpr(n *NonStarExpr) {
	r.seen = append(r.seen, "NonStarExpr")
}
func (r *recorder) VisitTableExpr(n *TableExpr) {
	r.seen = append(r.seen, "TableExpr")
}
func (r *recorder) VisitAliasedTableExpr(*AliasedTableExpr) {
	r.seen = append(r.seen, "AliasedTableExpr")
}
func (r *recorder) VisitParenTableExpr(*ParenTableExpr) {
	r.seen = append(r.seen, "ParenTableExpr")
}
func (r *recorder) VisitJoinTableExpr(*JoinTableExpr) {
	r.seen = append(r.seen, "JoinTableExpr")
}
