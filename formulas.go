/*
Citili, a program for generating CTL formulas for the model checking contest
Copyright (C) 2020  Loïg Jezequel

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see https://www.gnu.org/licenses/.
*/
package main

import "math/rand"

type formula struct {
	operator operator
	operand  []formula
}

func genBooleanFormula(maxDepth int) (f formula) {
	if maxDepth <= 1 {
		f = formula{operator: booleanOperators[0]}
		return f
	}

	// choose operator
	opNum := rand.Intn(len(booleanOperators))
	f = formula{operator: booleanOperators[opNum]}

	// generate subformulas
	arity := rand.Intn(f.operator.maxArity+1-f.operator.minArity) + f.operator.minArity
	f.operand = make([]formula, arity)
	if f.operator.isOverBooleans {
		for i := 0; i < arity; i++ {
			f.operand[i] = genBooleanFormula(maxDepth - 1)
		}
	} else {
		for i := 0; i < arity; i++ {
			f.operand[i] = genPathFormula(maxDepth - 1)
		}
	}

	return f
}

func genPathFormula(maxDepth int) (f formula) {
	// choose operator
	opNum := rand.Intn(len(pathOperators))
	f = formula{operator: pathOperators[opNum]}

	// generate subformulas
	arity := rand.Intn(f.operator.maxArity+1-f.operator.minArity) + f.operator.minArity
	f.operand = make([]formula, arity)
	for i := 0; i < arity; i++ {
		f.operand[i] = genBooleanFormula(maxDepth - 1)
	}

	return f
}

func genCTLFormula(maxDepth int) formula {
	return genBooleanFormula(maxDepth)
}