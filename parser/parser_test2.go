package parser

// func TestParsingHashLiteralsStringKeys(t *testing.T) {
// 	input := `{"one": 1, "two": 2, "three": 3}`

// 	l := lexer.New(input)
// 	p := New(l)
// 	program := p.ParseProgram()
// 	checkParserErrors(t, p)

// 	stmt := program.Statements[0].(*ast.ExpressionStatement)
// 	hash, ok := stmt.Expression.(*ast.HashLiteral)
// 	if !ok {
// 		t.Fatalf("exp is not %T. got=%T", ast.HashLiteral{}, stmt.Expression)
// 	}

// 	if len(hash.Pairs) != 3 {
// 		t.Errorf("hash.Pairs has wrong length. got=%d", len(hash.Pairs))
// 	}

// 	expected := map[string]int64{
// 		"one":   1,
// 		"two":   2,
// 		"three": 3,
// 	}

// 	for key, value := range hash.Pairs {
// 		literal, ok := key.(*ast.StringLiteral)
// 		if !ok {
// 			t.Errorf("key is not %T. got=%T", ast.StringLiteral{}, key)
// 		}

// 		expectedValue := expected[literal.String()]

// 		testIntegerLiteral(t, value, expectedValue)
// 	}
// }

// func TestParsingEmptyHashLiteral(t *testing.T) {
// 	input := `{}`

// 	l := lexer.New(input)
// 	p := New(l)
// 	program := p.ParseProgram()
// 	checkParserErrors(t, p)

// 	stmt := program.Statements[0].(*ast.ExpressionStatement)
// 	hash, ok := stmt.Expression.(*ast.HashLiteral)
// 	if !ok {
// 		t.Fatalf("exp is not %T. got=%T", ast.HashLiteral{}, stmt.Expression)
// 	}

// 	if len(hash.Pairs) != 0 {
// 		t.Errorf("hash.Pairs has wrong length. got=%d", len(hash.Pairs))
// 	}
// }

// func TestParsingHashLiteralsWithExpressions(t *testing.T) {
// 	input := `{"one": 0 + 1, "two": 10 - 8, "three": 15 / 5}`

// 	l := lexer.New(input)
// 	p := New(l)
// 	program := p.ParseProgram()
// 	checkParserErrors(t, p)

// 	stmt := program.Statements[0].(*ast.ExpressionStatement)
// 	hash, ok := stmt.Expression.(*ast.HashLiteral)
// 	if !ok {
// 		t.Fatalf("exp is not %T. got=%T", ast.HashLiteral{}, stmt.Expression)
// 	}

// 	if len(hash.Pairs) != 3 {
// 		t.Errorf("hash.Pairs has wrong length. got=%d", len(hash.Pairs))
// 	}

// 	tests := map[string]func(ast.Expression){
// 		"one": func(e ast.Expression) {
// 			testInfixExpression(t, e, 0, "+", 1)
// 		},
// 		"two": func(e ast.Expression) {
// 			testInfixExpression(t, e, 10, "-", 8)
// 		},
// 		"three": func(e ast.Expression) {
// 			testInfixExpression(t, e, 15, "/", 5)
// 		},
// 	}

// 	for key, value := range hash.Pairs {
// 		literal, ok := key.(*ast.StringLiteral)
// 		if !ok {
// 			t.Errorf("key is not %T. got=%T", ast.StringLiteral{}, key)
// 			continue
// 		}

// 		testFunc, ok := tests[literal.String()]
// 		if !ok {
// 			t.Errorf("No test function for key %q found", literal.String())
// 			continue
// 		}

// 		testFunc(value)
// 	}
// }

// func TestParsingArrayLiteral(t *testing.T) {
// 	input := "[1, 2 * 2, 3 + 3]"

// 	l := lexer.New(input)
// 	p := New(l)
// 	program := p.ParseProgram()
// 	checkParserErrors(t, p)

// 	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
// 	array, ok := stmt.Expression.(*ast.ArrayLiteral)
// 	if !ok {
// 		t.Fatalf("exp not %T. got=%T", &ast.ArrayLiteral{}, stmt.Expression)
// 	}

// 	if len(array.Elements) != 3 {
// 		t.Fatalf("len(array.Elements) not 3. got=%d", len(array.Elements))
// 	}

// 	testIntegerLiteral(t, array.Elements[0], 1)
// 	testInfixExpression(t, array.Elements[1], 2, "*", 2)
// 	testInfixExpression(t, array.Elements[2], 3, "+", 3)
// }

// func TestParsingIndexExpressions(t *testing.T) {
// 	input := "myArray[1 + 1]"

// 	l := lexer.New(input)
// 	p := New(l)
// 	program := p.ParseProgram()
// 	checkParserErrors(t, p)

// 	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
// 	indexExp, ok := stmt.Expression.(*ast.IndexExpression)
// 	if !ok {
// 		t.Fatalf("exp not %T. got=%T", &ast.IndexExpression{}, stmt.Expression)
// 	}

// 	if !testIdentifier(t, indexExp.Left, "myArray") {
// 		return
// 	}

// 	if !testInfixExpression(t, indexExp.Index, 1, "+", 1) {
// 		return
// 	}
// }
