package compiler

import (
	"mindfck/codegen"
	"mindfck/parser/mfast"
)

func Compile(ast []mfast.Stmt) (string, error) {
	cmd := codegen.New()
	for _, stmt := range ast {
		err := stmt.EvalStmt(cmd)
		if err != nil {
			return "", err
		}
	}
	return cmd.Compile(), nil
}
