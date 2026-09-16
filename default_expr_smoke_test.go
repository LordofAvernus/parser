package parser_test

import (
	"strings"
	"testing"

	"github.com/pingcap/parser"
	"github.com/pingcap/parser/ast"
	"github.com/pingcap/parser/format"
	_ "github.com/pingcap/parser/test_driver"
)

func TestDefaultExprColumnOption(t *testing.T) {
	cases := []string{
		"CREATE TABLE t1 (id int NOT NULL, j json NOT NULL DEFAULT (JSON_ARRAY()));",
		"CREATE TABLE t2 (id int NOT NULL, j json NOT NULL DEFAULT (JSON_OBJECT()));",
		"CREATE TABLE t3 (id int NOT NULL, d date NOT NULL DEFAULT (CURRENT_DATE));",
		"CREATE TABLE t4 (id int NOT NULL, n int NOT NULL DEFAULT ((1 + 2)));",
		"CREATE TABLE t5 (id int NOT NULL, u varchar(36) NOT NULL DEFAULT (UUID()));",
		"ALTER TABLE t1 ADD COLUMN j2 json NOT NULL DEFAULT (JSON_ARRAY());",
		"ALTER TABLE t1 MODIFY COLUMN j json NOT NULL DEFAULT (JSON_OBJECT());",
		"CREATE TABLE t6 (id int DEFAULT 1, ts timestamp DEFAULT CURRENT_TIMESTAMP);",
		"CREATE TABLE t7 (c json NOT NULL DEFAULT(JSON_ARRAY()));",
	}
	p := parser.New()
	for _, sql := range cases {
		stmts, _, err := p.Parse(sql, "", "")
		if err != nil {
			t.Errorf("FAIL parse: %s err=%v", sql, err)
			continue
		}
		var b strings.Builder
		ctx := format.NewRestoreCtx(format.DefaultRestoreFlags, &b)
		if err := stmts[0].(ast.Node).Restore(ctx); err != nil {
			t.Errorf("FAIL restore: %s err=%v", sql, err)
			continue
		}
		restored := b.String()
		if _, _, err := p.Parse(restored, "", ""); err != nil {
			t.Errorf("FAIL reparse: %s restored=%s err=%v", sql, restored, err)
			continue
		}
		t.Logf("OK: %s -> %s", sql, restored)
	}
}
