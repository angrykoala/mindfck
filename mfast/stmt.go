package mfast

import (
	"mindfck/codegen"
)

type Stmt interface {
	EvalStmt(cmd *codegen.CommandHandler) error
}
