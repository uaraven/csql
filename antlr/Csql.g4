/*
 *    Copyright 2021 Oleksiy Voronin
 *
 *    Licensed under the Apache License, Version 2.0 (the "License");
 *    you may not use this file except in compliance with the License.
 *    You may obtain a copy of the License at
 *
 *        http://www.apache.org/licenses/LICENSE-2.0
 *
 *    Unless required by applicable law or agreed to in writing, software
 *    distributed under the License is distributed on an "AS IS" BASIS,
 *    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *    See the License for the specific language governing permissions and
 *    limitations under the License.
 *
 *    SPDX-License-Identifier: Apache-2.0
 *    SPDX-FileCopyrightText: (c) 2022 Oleksiy Voronin <ovoronin@gmail.com>
 */
grammar Csql;

query: selectStatement ';'? EOF;

comparisonOperator: '<' | '<=' | '>' | '>=' | '=' | '!=' | '<>';

binaryOperation: '+' | '-' | '*' | '/' | '%' | '||';

list: '(' valueExpr (',' valueExpr)* ')';

term: compoundName | literalValue;

valueExpr
    : valueExpr binaryOperation valueExpr                 # valueBinaryExpr
    | term                                                # valueTerm
    | '(' valueExpr ')'                                   # valueParensExpr
    ;

whereExpr
    : term										     # termItem
    | valueExpr                                      # valueExprItem
    | whereExpr comparisonOperator whereExpr		 # condition
    | whereExpr K_AND whereExpr						 # andExpr
    | whereExpr K_OR whereExpr						 # orExpr
    | K_NOT whereExpr								 # notExpr
    | valueExpr K_BETWEEN valueExpr K_AND valueExpr	 # betweenExpr
    | valueExpr K_NOT? K_LIKE valueExpr	             # likeExpr
    | valueExpr K_NOT? K_MATCH valueExpr             # matchExpr
    | valueExpr K_NOT? K_IN list					 # inExpr
    | valueExpr K_IS K_NOT? K_NULL                   # isNullExpr
    | '(' whereExpr ')'								 # parensExpr
    ;

where: K_WHERE whereExpr;

distinct: K_DISTINCT;

projection: distinct? projectionField (',' projectionField)*;

projectionField: (projectionFieldName | valueExpr ) (K_AS? alias)?;

projectionFieldName: (qualifier '.')? fieldName;

fieldName: IDENTIFIER | '*';

innerJoin: K_INNER? K_JOIN;
leftJoin: K_LEFT K_OUTER? K_JOIN;
rightJoin: K_RIGHT K_OUTER? K_JOIN;
//fullJoin: K_FULL K_OUTER? K_JOIN;
crossJoin: K_CROSS K_JOIN;

conditionalJoinType
        : innerJoin
		| leftJoin
		| rightJoin;
//		| fullJoin


conditionalJoinTarget: dataSource K_ON whereExpr;

unconditionalJoinTarget: dataSource;

dataSource:
	sourceName												# dataSourceNamed
	| dataSource conditionalJoinType conditionalJoinTarget	# dataSourceConditionalJoin
	| dataSource crossJoin unconditionalJoinTarget			# dataSourceCrossJoin
	| '(' dataSource ')'									# dataSourceItem;

sources: dataSource (',' dataSource)*;

selectStatement: K_SELECT projection K_FROM sources where? limit? orderBy?;

signedNumber: ( '+' | '-')? NUMERIC_LITERAL;

stringValue: STRING_LITERAL;

nullValue: K_NULL;

literalValue: signedNumber | stringValue | nullValue;

alias: IDENTIFIER;

sourceName: name alias?;

compoundName: (qualifier '.')? name;

name: IDENTIFIER;

qualifier: IDENTIFIER;

limit: K_LIMIT limitValue;

limitValue: NUMERIC_LITERAL;

orderBy: K_ORDER K_BY orderByField (',' orderByField)*;

orderByField
    : (compoundName | fieldIndex) ( K_ASC | K_DESC )?
    ;

fieldIndex: NUMERIC_LITERAL;

K_AND: A N D;
K_AS: A S;
K_BETWEEN: B E T W E E N;
K_FROM: F R O M;
K_IN: I N;
K_IS: I S;
K_NOT: N O T;
K_NULL: N U L L;
K_OR: O R;
K_REGEX: R E G E X;
K_SELECT: S E L E C T;
K_MATCH: M A T C H;
K_WHERE: W H E R E;
K_TRUE: T R U E;
K_FALSE: F A L S E;
K_JOIN: J O I N;
K_LEFT: L E F T;
K_RIGHT: R I G H T;
K_OUTER: O U T E R;
K_INNER: I N N E R;
K_CROSS: C R O S S;
K_ON: O N;
K_DISTINCT: D I S T I N C T;
K_LIKE: L I K E;
K_FULL: F U L L;
K_LIMIT: L I M I T;
K_ORDER: O R D E R;
K_BY: B Y;
K_ASC: A S C;
K_DESC: D E S C;

IDENTIFIER: SIMPLE_IDENTIFIER | QUOTED_IDENTIFIER;

SIMPLE_IDENTIFIER: [a-zA-Z] [a-zA-Z_0-9]*;

QUOTED_IDENTIFIER: '"' ( ~'"' | '""')* '"';

NUMERIC_LITERAL:
	DIGIT+ ('.' DIGIT*)? (E [-+]? DIGIT+)?
	| '.' DIGIT+ ( E [-+]? DIGIT+)?;

STRING_LITERAL: '\'' (~'\'' | '\'\'')* '\'';

SINGLE_LINE_COMMENT: '--' ~[\r\n]* -> channel(HIDDEN);

MULTILINE_COMMENT: '/*' .*? ( '*/' | EOF) -> channel(HIDDEN);

SPACES: [ \u000B\t\r\n] -> channel(HIDDEN);

UNEXPECTED_CHAR: .;

fragment DIGIT: [0-9];

fragment A: [aA];
fragment B: [bB];
fragment C: [cC];
fragment D: [dD];
fragment E: [eE];
fragment F: [fF];
fragment G: [gG];
fragment H: [hH];
fragment I: [iI];
fragment J: [jJ];
fragment K: [kK];
fragment L: [lL];
fragment M: [mM];
fragment N: [nN];
fragment O: [oO];
fragment P: [pP];
fragment Q: [qQ];
fragment R: [rR];
fragment S: [sS];
fragment T: [tT];
fragment U: [uU];
fragment V: [vV];
fragment W: [wW];
fragment X: [xX];
fragment Y: [yY];
fragment Z: [zZ];