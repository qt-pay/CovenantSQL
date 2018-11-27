/*
 * Copyright 2018 The CovenantSQL Authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package resolver

import "github.com/CovenantSQL/sqlparser"

// Query defines a resolver result query.
type Query struct {
	Stmt          sqlparser.Statement
	Query         string
	ParamCount    int
	ResultColumns []string
}

// IsDDL returns whether a resolved query is DDL or not.
func (q *Query) IsDDL() bool {
	if q.Stmt != nil {
		_, ok := q.Stmt.(*sqlparser.DDL)
		return ok
	}

	return false
}

// IsRead returns whether a resolved query is READ query.
func (q *Query) IsRead() bool {
	if q.Stmt != nil {
		switch q.Stmt.(type) {
		case *sqlparser.Show, *sqlparser.Explain, sqlparser.SelectStatement:
			return true
		}
	}

	return false
}