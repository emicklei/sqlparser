package sqlparser

// Visitor methods are all put in this source file to
// reduce work when updating ast sources from vitesss upstream.

type ASTVisitor interface {
	VisitSelect(n Select)
	VisitUnion(n Union)
	VisitInsert(n Insert)
	VisitUpdate(n Update)
	VisitDelete(n Delete)
	VisitSet(n Set)
	VisitDDL(n DDL)
	VisitColumnDefinition(n ColumnDefinition)
	VisitCreateTable(n CreateTable)
	VisitOther(n Other)
	VisitSelectExprs(n SelectExprs)
	VisitStarExpr(n StarExpr)
}

func (n Union) Accept(v ASTVisitor)            { v.VisitUnion(n) }
func (n Select) Accept(v ASTVisitor)           { v.VisitSelect(n) }
func (n Insert) Accept(v ASTVisitor)           { v.VisitInsert(n) }
func (n Update) Accept(v ASTVisitor)           { v.VisitUpdate(n) }
func (n Delete) Accept(v ASTVisitor)           { v.VisitDelete(n) }
func (n Set) Accept(v ASTVisitor)              { v.VisitSet(n) }
func (n DDL) Accept(v ASTVisitor)              { v.VisitDDL(n) }
func (n ColumnDefinition) Accept(v ASTVisitor) { v.VisitColumnDefinition(n) }
func (n CreateTable) Accept(v ASTVisitor)      { v.VisitCreateTable(n) }
func (n Other) Accept(v ASTVisitor)            { v.VisitOther(n) }
func (n SelectExprs) Accept(v ASTVisitor)      { v.VisitSelectExprs(n) }
func (n StarExpr) Accept(v ASTVisitor)         { v.VisitStarExpr(n) }
