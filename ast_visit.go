package sqlparser

import "log"

// Visitor methods are all put in this source file to
// reduce work when updating ast sources from vitesss upstream.

type ASTVisitor interface {
	VisitSelect(n *Select)
	VisitUnion(n *Union)
	VisitInsert(n *Insert)
	VisitUpdate(n *Update)
	VisitDelete(n *Delete)
	VisitSet(n *Set)
	VisitDDL(n *DDL)
	VisitColumnDefinition(n *ColumnDefinition)
	VisitCreateTable(n *CreateTable)
	VisitOther(n *Other)
	VisitStarExpr(n *StarExpr)
	VisitNonStarExpr(n *NonStarExpr)
	VisitAliasedTableExpr(*AliasedTableExpr)
	VisitParenTableExpr(*ParenTableExpr)
	VisitJoinTableExpr(*JoinTableExpr)
	VisitWhere(*Where)
	VisitOrder(*Order)
}

// might not need the accept methods...
func (n *Union) Accept(v ASTVisitor)            { v.VisitUnion(n) }
func (n *Select) Accept(v ASTVisitor)           { v.VisitSelect(n) }
func (n *Insert) Accept(v ASTVisitor)           { v.VisitInsert(n) }
func (n *Update) Accept(v ASTVisitor)           { v.VisitUpdate(n) }
func (n *Delete) Accept(v ASTVisitor)           { v.VisitDelete(n) }
func (n *Set) Accept(v ASTVisitor)              { v.VisitSet(n) }
func (n *DDL) Accept(v ASTVisitor)              { v.VisitDDL(n) }
func (n *ColumnDefinition) Accept(v ASTVisitor) { v.VisitColumnDefinition(n) }
func (n *CreateTable) Accept(v ASTVisitor)      { v.VisitCreateTable(n) }
func (n *Other) Accept(v ASTVisitor)            { v.VisitOther(n) }
func (n *StarExpr) Accept(v ASTVisitor)         { v.VisitStarExpr(n) }
func (n *NonStarExpr) Accept(v ASTVisitor)      { v.VisitNonStarExpr(n) }

func StatementAccept(n Statement, v ASTVisitor) {
	switch n.(type) {
	case *Union:
		v.VisitUnion(n.(*Union))
	case *Select:
		v.VisitSelect(n.(*Select))
	default:
		log.Printf("StatementAccept TODO %T", n)
	}
}

func SelectExprAccept(n SelectExpr, v ASTVisitor) {
	switch n.(type) {
	case *StarExpr:
		v.VisitStarExpr(n.(*StarExpr))
	case *NonStarExpr:
		v.VisitNonStarExpr(n.(*NonStarExpr))
	default:
		log.Printf("SelectExprAccept TODO %T", n)
	}
}

func TableExprAccept(n TableExpr, v ASTVisitor) {
	switch n.(type) {
	case *AliasedTableExpr:
		v.VisitAliasedTableExpr(n.(*AliasedTableExpr))
	case *ParenTableExpr:
		v.VisitParenTableExpr(n.(*ParenTableExpr))
	case *JoinTableExpr:
		v.VisitJoinTableExpr(n.(*JoinTableExpr))
	default:
		log.Printf("TableExprAccept TODO %T", n)
	}
}
