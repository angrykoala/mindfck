package parser

import (
	"mindfck/mfast"
	mindfck "mindfck/parser/antlr"
	"mindfck/utils"

	"github.com/antlr4-go/antlr/v4"
)

func Parse(input string) ([]mfast.Stmt, error) {
	inputStream := antlr.NewInputStream(input)
	lexer := mindfck.NewmindfckLexer(inputStream)
	tokenStream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	parser := mindfck.NewmindfckParser(tokenStream)
	tree := parser.Statements()

	listener := &AstGenerator{
		scope:      newScope(),
		scopeStack: []*Scope{},
	}

	visitor := &AstGeneratorVisitor{}

	// TODO How to do the error handling here?
	antlr.ParseTreeWalkerDefault.Walk(listener, tree)
	ast := visitor.Visit(tree)

	return ast.([]mfast.Stmt), nil
}

type AstGeneratorVisitor struct {
	mindfck.BasemindfckVisitor
}

func (v *AstGeneratorVisitor) Visit(tree antlr.ParseTree) interface{} {
	switch tree.(type) {
	case *mindfck.StatementsContext:
		return v.VisitStatements(tree.(*mindfck.StatementsContext))
	}

	return nil
}

func (v *AstGeneratorVisitor) VisitStatements(ctx *mindfck.StatementsContext) []mfast.Stmt {
	result := []mfast.Stmt{}
	if ctx.AllStatement() != nil {
		for i, _ := range ctx.AllStatement() {
			result = append(result, v.VisitStatement(ctx.Statement(i).(*mindfck.StatementContext))...)
		}
	}

	return result
}

func (v *AstGeneratorVisitor) VisitStatement(ctx *mindfck.StatementContext) []mfast.Stmt {
	if ctx.Declaration() != nil {
		return v.VisitDeclaration(ctx.Declaration().(*mindfck.DeclarationContext))
	}

	return []mfast.Stmt{}
}

func (v *AstGeneratorVisitor) VisitDeclaration(ctx *mindfck.DeclarationContext) []mfast.Stmt {
	return []mfast.Stmt{&mfast.Declare{
		Label: ctx.Identifier().IDENTIFIER().GetText(),
	}}
}

type AstGenerator struct {
	*mindfck.BasemindfckListener

	scope          *Scope
	lastInnerScope *Scope // Scope to be set on exitBlock

	scopeStack []*Scope
}

func (l *AstGenerator) Ast() []mfast.Stmt {
	return l.scope.stmts
}

func (l *AstGenerator) EnterStatement(c *mindfck.StatementContext) {
	l.scope.clearExpr()
}

func (l *AstGenerator) ExitDeclaration(decl *mindfck.DeclarationContext) {
	current := &mfast.Declare{
		Label: decl.Identifier().IDENTIFIER().GetText(),
	}
	l.scope.appendStmt(current)
}

func (l *AstGenerator) ExitAssignment(assignment *mindfck.AssignmentContext) {
	current := &mfast.Assign{
		To:   assignment.Identifier().GetText(),
		From: l.scope.popLeftExpr(),
	}

	l.scope.appendStmt(current)
}

func (l *AstGenerator) ExitPrint(print *mindfck.PrintContext) {
	current := &mfast.Print{
		Value: l.scope.popLeftExpr(),
	}

	l.scope.appendStmt(current)
}

// EnterBlock is called when entering the block production.
func (l *AstGenerator) EnterBlock(c *mindfck.BlockContext) {
	l.beginScope()
}

// ExitBlock is called when exiting the block production.
func (l *AstGenerator) ExitBlock(c *mindfck.BlockContext) {
	l.lastInnerScope = l.popScope()
}

func (l *AstGenerator) ExitIfConditional(cond *mindfck.IfConditionalContext) {
	innerScope := l.lastInnerScope
	if innerScope == nil {
		panic("Cannot find inner scope in if block")
	}

	current := &mfast.If{
		Condition: l.scope.popLeftExpr(),
		Block:     innerScope.stmts,
	}
	l.scope.appendStmt(current)
}

// ExitWhileLoop is called when exiting the whileLoop production.
func (l *AstGenerator) ExitWhileLoop(c *mindfck.WhileLoopContext) {
	innerScope := l.lastInnerScope
	if innerScope == nil {
		panic("Cannot find inner scope in if block")
	}

	current := &mfast.While{
		Condition: l.scope.popLeftExpr(),
		Block:     innerScope.stmts,
	}
	l.scope.appendStmt(current)
}

func (l *AstGenerator) ExitExpression(c *mindfck.ExpressionContext) {
	if c.Literal() != nil {
		l.scope.pushExpr(&mfast.Literal{
			Value: utils.ToInt(c.Literal().GetText()),
		})

	} else if c.Identifier() != nil {
		l.scope.pushExpr(&mfast.VariableExpr{
			Label: c.Identifier().GetText(),
		})
	} else if c.Expression(0) != nil && c.Expression(1) != nil {
		l.scope.pushExpr(&mfast.BinaryExpr{
			Operator: mfast.Operand(c.Operand().GetText()),
			Left:     l.scope.popLeftExpr(),
			Right:    l.scope.popLeftExpr(),
		})
	}
}

func (l *AstGenerator) beginScope() {
	l.scopeStack = append(l.scopeStack, l.scope)
	l.scope = newScope()
}

// Pops last scope into current, returns current scope
func (l *AstGenerator) popScope() *Scope {
	current := l.scope
	l.scope = l.scopeStack[len(l.scopeStack)-1]
	l.scopeStack = l.scopeStack[:len(l.scopeStack)-1]
	return current
}
