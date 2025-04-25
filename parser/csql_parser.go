// Code generated from Csql.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // Csql

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type CsqlParser struct {
	*antlr.BaseParser
}

var CsqlParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func csqlParserInit() {
	staticData := &CsqlParserStaticData
	staticData.LiteralNames = []string{
		"", "';'", "'<'", "'<='", "'>'", "'>='", "'='", "'!='", "'<>'", "'+'",
		"'-'", "'*'", "'/'", "'%'", "'||'", "'('", "','", "')'", "'.'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "K_AND", "K_AS", "K_BETWEEN", "K_FROM", "K_IN", "K_IS", "K_NOT",
		"K_NULL", "K_OR", "K_REGEX", "K_SELECT", "K_MATCH", "K_WHERE", "K_TRUE",
		"K_FALSE", "K_JOIN", "K_LEFT", "K_RIGHT", "K_OUTER", "K_INNER", "K_CROSS",
		"K_ON", "K_DISTINCT", "K_LIKE", "K_FULL", "K_LIMIT", "K_ORDER", "K_BY",
		"K_ASC", "K_DESC", "K_UNION", "K_ALL", "K_TO_STRING", "K_TO_FLOAT",
		"K_TO_INT", "K_ROUND", "K_LEN", "K_TRUNC", "K_FRAC", "K_SUBSTRING",
		"K_TO_UPPER", "K_TO_LOWER", "K_POW", "K_SQRT", "K_COALESCE", "K_COUNT",
		"K_SUM", "K_AVG", "K_MIN", "K_MAX", "K_GROUP", "K_HAVING", "IDENTIFIER",
		"SIMPLE_IDENTIFIER", "QUOTED_IDENTIFIER", "NUMERIC_LITERAL", "STRING_LITERAL",
		"SINGLE_LINE_COMMENT", "MULTILINE_COMMENT", "SPACES", "UNEXPECTED_CHAR",
	}
	staticData.RuleNames = []string{
		"query", "comparisonOperator", "binaryOperation", "list", "term", "valueExpr",
		"whereExpr", "where", "having", "distinct", "funCall", "aggregateFunCall",
		"countFunc", "aggregateFunc", "projection", "projectionField", "projectionFieldName",
		"fieldName", "innerJoin", "leftJoin", "rightJoin", "fullJoin", "crossJoin",
		"conditionalJoinType", "conditionalJoinTarget", "unconditionalJoinTarget",
		"dataSource", "sources", "unionSelects", "selectStatement", "signedNumber",
		"stringValue", "nullValue", "literalValue", "alias", "sourceName", "compoundName",
		"name", "qualifier", "limit", "limitValue", "orderBy", "orderByField",
		"fieldIndex", "groupBy", "groupByField", "function",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 79, 436, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36, 7, 36,
		2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7, 41, 2,
		42, 7, 42, 2, 43, 7, 43, 2, 44, 7, 44, 2, 45, 7, 45, 2, 46, 7, 46, 1, 0,
		1, 0, 3, 0, 97, 8, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1,
		3, 1, 3, 5, 3, 109, 8, 3, 10, 3, 12, 3, 112, 9, 3, 1, 3, 1, 3, 1, 4, 1,
		4, 3, 4, 118, 8, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 3,
		5, 128, 8, 5, 1, 5, 1, 5, 1, 5, 1, 5, 5, 5, 134, 8, 5, 10, 5, 12, 5, 137,
		9, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6,
		1, 6, 1, 6, 3, 6, 152, 8, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 3, 6, 159, 8,
		6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 3, 6, 166, 8, 6, 1, 6, 1, 6, 1, 6, 1,
		6, 1, 6, 1, 6, 3, 6, 174, 8, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 3,
		6, 182, 8, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1,
		6, 5, 6, 194, 8, 6, 10, 6, 12, 6, 197, 9, 6, 1, 7, 1, 7, 1, 7, 1, 8, 1,
		8, 1, 8, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 5, 10, 212, 8,
		10, 10, 10, 12, 10, 215, 9, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 11,
		1, 11, 1, 11, 3, 11, 225, 8, 11, 1, 12, 1, 12, 1, 12, 3, 12, 230, 8, 12,
		1, 12, 1, 12, 1, 12, 1, 13, 1, 13, 1, 14, 3, 14, 238, 8, 14, 1, 14, 1,
		14, 1, 14, 5, 14, 243, 8, 14, 10, 14, 12, 14, 246, 9, 14, 1, 15, 1, 15,
		1, 15, 3, 15, 251, 8, 15, 1, 15, 3, 15, 254, 8, 15, 1, 15, 3, 15, 257,
		8, 15, 1, 16, 1, 16, 1, 16, 3, 16, 262, 8, 16, 1, 16, 1, 16, 1, 17, 1,
		17, 1, 18, 3, 18, 269, 8, 18, 1, 18, 1, 18, 1, 19, 1, 19, 3, 19, 275, 8,
		19, 1, 19, 1, 19, 1, 20, 1, 20, 3, 20, 281, 8, 20, 1, 20, 1, 20, 1, 21,
		1, 21, 3, 21, 287, 8, 21, 1, 21, 1, 21, 1, 22, 1, 22, 1, 22, 1, 23, 1,
		23, 1, 23, 1, 23, 3, 23, 298, 8, 23, 1, 24, 1, 24, 1, 24, 1, 24, 1, 25,
		1, 25, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 3, 26, 312, 8, 26, 1,
		26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 5, 26, 322, 8, 26,
		10, 26, 12, 26, 325, 9, 26, 1, 27, 1, 27, 1, 27, 5, 27, 330, 8, 27, 10,
		27, 12, 27, 333, 9, 27, 1, 28, 1, 28, 1, 28, 3, 28, 338, 8, 28, 1, 28,
		5, 28, 341, 8, 28, 10, 28, 12, 28, 344, 9, 28, 1, 29, 1, 29, 1, 29, 1,
		29, 1, 29, 3, 29, 351, 8, 29, 1, 29, 3, 29, 354, 8, 29, 1, 29, 3, 29, 357,
		8, 29, 1, 29, 3, 29, 360, 8, 29, 1, 29, 3, 29, 363, 8, 29, 1, 30, 3, 30,
		366, 8, 30, 1, 30, 1, 30, 1, 31, 1, 31, 1, 32, 1, 32, 1, 33, 1, 33, 1,
		33, 3, 33, 377, 8, 33, 1, 34, 1, 34, 1, 35, 1, 35, 3, 35, 383, 8, 35, 1,
		36, 1, 36, 1, 36, 3, 36, 388, 8, 36, 1, 36, 1, 36, 1, 37, 1, 37, 1, 38,
		1, 38, 1, 39, 1, 39, 1, 39, 1, 40, 1, 40, 1, 41, 1, 41, 1, 41, 1, 41, 1,
		41, 5, 41, 406, 8, 41, 10, 41, 12, 41, 409, 9, 41, 1, 42, 1, 42, 3, 42,
		413, 8, 42, 1, 42, 3, 42, 416, 8, 42, 1, 43, 1, 43, 1, 44, 1, 44, 1, 44,
		1, 44, 1, 44, 5, 44, 425, 8, 44, 10, 44, 12, 44, 428, 9, 44, 1, 45, 1,
		45, 3, 45, 432, 8, 45, 1, 46, 1, 46, 1, 46, 0, 3, 10, 12, 52, 47, 0, 2,
		4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40,
		42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 66, 68, 70, 72, 74, 76,
		78, 80, 82, 84, 86, 88, 90, 92, 0, 7, 1, 0, 2, 8, 1, 0, 9, 14, 1, 0, 65,
		68, 2, 0, 11, 11, 71, 71, 1, 0, 9, 10, 1, 0, 47, 48, 1, 0, 51, 63, 448,
		0, 94, 1, 0, 0, 0, 2, 100, 1, 0, 0, 0, 4, 102, 1, 0, 0, 0, 6, 104, 1, 0,
		0, 0, 8, 117, 1, 0, 0, 0, 10, 127, 1, 0, 0, 0, 12, 181, 1, 0, 0, 0, 14,
		198, 1, 0, 0, 0, 16, 201, 1, 0, 0, 0, 18, 204, 1, 0, 0, 0, 20, 206, 1,
		0, 0, 0, 22, 224, 1, 0, 0, 0, 24, 226, 1, 0, 0, 0, 26, 234, 1, 0, 0, 0,
		28, 237, 1, 0, 0, 0, 30, 250, 1, 0, 0, 0, 32, 261, 1, 0, 0, 0, 34, 265,
		1, 0, 0, 0, 36, 268, 1, 0, 0, 0, 38, 272, 1, 0, 0, 0, 40, 278, 1, 0, 0,
		0, 42, 284, 1, 0, 0, 0, 44, 290, 1, 0, 0, 0, 46, 297, 1, 0, 0, 0, 48, 299,
		1, 0, 0, 0, 50, 303, 1, 0, 0, 0, 52, 311, 1, 0, 0, 0, 54, 326, 1, 0, 0,
		0, 56, 334, 1, 0, 0, 0, 58, 345, 1, 0, 0, 0, 60, 365, 1, 0, 0, 0, 62, 369,
		1, 0, 0, 0, 64, 371, 1, 0, 0, 0, 66, 376, 1, 0, 0, 0, 68, 378, 1, 0, 0,
		0, 70, 380, 1, 0, 0, 0, 72, 387, 1, 0, 0, 0, 74, 391, 1, 0, 0, 0, 76, 393,
		1, 0, 0, 0, 78, 395, 1, 0, 0, 0, 80, 398, 1, 0, 0, 0, 82, 400, 1, 0, 0,
		0, 84, 412, 1, 0, 0, 0, 86, 417, 1, 0, 0, 0, 88, 419, 1, 0, 0, 0, 90, 431,
		1, 0, 0, 0, 92, 433, 1, 0, 0, 0, 94, 96, 3, 56, 28, 0, 95, 97, 5, 1, 0,
		0, 96, 95, 1, 0, 0, 0, 96, 97, 1, 0, 0, 0, 97, 98, 1, 0, 0, 0, 98, 99,
		5, 0, 0, 1, 99, 1, 1, 0, 0, 0, 100, 101, 7, 0, 0, 0, 101, 3, 1, 0, 0, 0,
		102, 103, 7, 1, 0, 0, 103, 5, 1, 0, 0, 0, 104, 105, 5, 15, 0, 0, 105, 110,
		3, 10, 5, 0, 106, 107, 5, 16, 0, 0, 107, 109, 3, 10, 5, 0, 108, 106, 1,
		0, 0, 0, 109, 112, 1, 0, 0, 0, 110, 108, 1, 0, 0, 0, 110, 111, 1, 0, 0,
		0, 111, 113, 1, 0, 0, 0, 112, 110, 1, 0, 0, 0, 113, 114, 5, 17, 0, 0, 114,
		7, 1, 0, 0, 0, 115, 118, 3, 72, 36, 0, 116, 118, 3, 66, 33, 0, 117, 115,
		1, 0, 0, 0, 117, 116, 1, 0, 0, 0, 118, 9, 1, 0, 0, 0, 119, 120, 6, 5, -1,
		0, 120, 128, 3, 8, 4, 0, 121, 128, 3, 20, 10, 0, 122, 128, 3, 22, 11, 0,
		123, 124, 5, 15, 0, 0, 124, 125, 3, 10, 5, 0, 125, 126, 5, 17, 0, 0, 126,
		128, 1, 0, 0, 0, 127, 119, 1, 0, 0, 0, 127, 121, 1, 0, 0, 0, 127, 122,
		1, 0, 0, 0, 127, 123, 1, 0, 0, 0, 128, 135, 1, 0, 0, 0, 129, 130, 10, 5,
		0, 0, 130, 131, 3, 4, 2, 0, 131, 132, 3, 10, 5, 6, 132, 134, 1, 0, 0, 0,
		133, 129, 1, 0, 0, 0, 134, 137, 1, 0, 0, 0, 135, 133, 1, 0, 0, 0, 135,
		136, 1, 0, 0, 0, 136, 11, 1, 0, 0, 0, 137, 135, 1, 0, 0, 0, 138, 139, 6,
		6, -1, 0, 139, 182, 3, 8, 4, 0, 140, 182, 3, 10, 5, 0, 141, 142, 5, 25,
		0, 0, 142, 182, 3, 12, 6, 7, 143, 144, 3, 10, 5, 0, 144, 145, 5, 21, 0,
		0, 145, 146, 3, 10, 5, 0, 146, 147, 5, 19, 0, 0, 147, 148, 3, 10, 5, 0,
		148, 182, 1, 0, 0, 0, 149, 151, 3, 10, 5, 0, 150, 152, 5, 25, 0, 0, 151,
		150, 1, 0, 0, 0, 151, 152, 1, 0, 0, 0, 152, 153, 1, 0, 0, 0, 153, 154,
		5, 42, 0, 0, 154, 155, 3, 10, 5, 0, 155, 182, 1, 0, 0, 0, 156, 158, 3,
		10, 5, 0, 157, 159, 5, 25, 0, 0, 158, 157, 1, 0, 0, 0, 158, 159, 1, 0,
		0, 0, 159, 160, 1, 0, 0, 0, 160, 161, 5, 30, 0, 0, 161, 162, 3, 10, 5,
		0, 162, 182, 1, 0, 0, 0, 163, 165, 3, 10, 5, 0, 164, 166, 5, 25, 0, 0,
		165, 164, 1, 0, 0, 0, 165, 166, 1, 0, 0, 0, 166, 167, 1, 0, 0, 0, 167,
		168, 5, 23, 0, 0, 168, 169, 3, 6, 3, 0, 169, 182, 1, 0, 0, 0, 170, 171,
		3, 10, 5, 0, 171, 173, 5, 24, 0, 0, 172, 174, 5, 25, 0, 0, 173, 172, 1,
		0, 0, 0, 173, 174, 1, 0, 0, 0, 174, 175, 1, 0, 0, 0, 175, 176, 5, 26, 0,
		0, 176, 182, 1, 0, 0, 0, 177, 178, 5, 15, 0, 0, 178, 179, 3, 12, 6, 0,
		179, 180, 5, 17, 0, 0, 180, 182, 1, 0, 0, 0, 181, 138, 1, 0, 0, 0, 181,
		140, 1, 0, 0, 0, 181, 141, 1, 0, 0, 0, 181, 143, 1, 0, 0, 0, 181, 149,
		1, 0, 0, 0, 181, 156, 1, 0, 0, 0, 181, 163, 1, 0, 0, 0, 181, 170, 1, 0,
		0, 0, 181, 177, 1, 0, 0, 0, 182, 195, 1, 0, 0, 0, 183, 184, 10, 10, 0,
		0, 184, 185, 3, 2, 1, 0, 185, 186, 3, 12, 6, 11, 186, 194, 1, 0, 0, 0,
		187, 188, 10, 9, 0, 0, 188, 189, 5, 19, 0, 0, 189, 194, 3, 12, 6, 10, 190,
		191, 10, 8, 0, 0, 191, 192, 5, 27, 0, 0, 192, 194, 3, 12, 6, 9, 193, 183,
		1, 0, 0, 0, 193, 187, 1, 0, 0, 0, 193, 190, 1, 0, 0, 0, 194, 197, 1, 0,
		0, 0, 195, 193, 1, 0, 0, 0, 195, 196, 1, 0, 0, 0, 196, 13, 1, 0, 0, 0,
		197, 195, 1, 0, 0, 0, 198, 199, 5, 31, 0, 0, 199, 200, 3, 12, 6, 0, 200,
		15, 1, 0, 0, 0, 201, 202, 5, 70, 0, 0, 202, 203, 3, 12, 6, 0, 203, 17,
		1, 0, 0, 0, 204, 205, 5, 41, 0, 0, 205, 19, 1, 0, 0, 0, 206, 207, 3, 92,
		46, 0, 207, 208, 5, 15, 0, 0, 208, 213, 3, 10, 5, 0, 209, 210, 5, 16, 0,
		0, 210, 212, 3, 10, 5, 0, 211, 209, 1, 0, 0, 0, 212, 215, 1, 0, 0, 0, 213,
		211, 1, 0, 0, 0, 213, 214, 1, 0, 0, 0, 214, 216, 1, 0, 0, 0, 215, 213,
		1, 0, 0, 0, 216, 217, 5, 17, 0, 0, 217, 21, 1, 0, 0, 0, 218, 219, 3, 26,
		13, 0, 219, 220, 5, 15, 0, 0, 220, 221, 3, 10, 5, 0, 221, 222, 5, 17, 0,
		0, 222, 225, 1, 0, 0, 0, 223, 225, 3, 24, 12, 0, 224, 218, 1, 0, 0, 0,
		224, 223, 1, 0, 0, 0, 225, 23, 1, 0, 0, 0, 226, 227, 5, 64, 0, 0, 227,
		229, 5, 15, 0, 0, 228, 230, 5, 41, 0, 0, 229, 228, 1, 0, 0, 0, 229, 230,
		1, 0, 0, 0, 230, 231, 1, 0, 0, 0, 231, 232, 3, 32, 16, 0, 232, 233, 5,
		17, 0, 0, 233, 25, 1, 0, 0, 0, 234, 235, 7, 2, 0, 0, 235, 27, 1, 0, 0,
		0, 236, 238, 3, 18, 9, 0, 237, 236, 1, 0, 0, 0, 237, 238, 1, 0, 0, 0, 238,
		239, 1, 0, 0, 0, 239, 244, 3, 30, 15, 0, 240, 241, 5, 16, 0, 0, 241, 243,
		3, 30, 15, 0, 242, 240, 1, 0, 0, 0, 243, 246, 1, 0, 0, 0, 244, 242, 1,
		0, 0, 0, 244, 245, 1, 0, 0, 0, 245, 29, 1, 0, 0, 0, 246, 244, 1, 0, 0,
		0, 247, 251, 3, 32, 16, 0, 248, 251, 3, 10, 5, 0, 249, 251, 3, 22, 11,
		0, 250, 247, 1, 0, 0, 0, 250, 248, 1, 0, 0, 0, 250, 249, 1, 0, 0, 0, 251,
		256, 1, 0, 0, 0, 252, 254, 5, 20, 0, 0, 253, 252, 1, 0, 0, 0, 253, 254,
		1, 0, 0, 0, 254, 255, 1, 0, 0, 0, 255, 257, 3, 68, 34, 0, 256, 253, 1,
		0, 0, 0, 256, 257, 1, 0, 0, 0, 257, 31, 1, 0, 0, 0, 258, 259, 3, 76, 38,
		0, 259, 260, 5, 18, 0, 0, 260, 262, 1, 0, 0, 0, 261, 258, 1, 0, 0, 0, 261,
		262, 1, 0, 0, 0, 262, 263, 1, 0, 0, 0, 263, 264, 3, 34, 17, 0, 264, 33,
		1, 0, 0, 0, 265, 266, 7, 3, 0, 0, 266, 35, 1, 0, 0, 0, 267, 269, 5, 38,
		0, 0, 268, 267, 1, 0, 0, 0, 268, 269, 1, 0, 0, 0, 269, 270, 1, 0, 0, 0,
		270, 271, 5, 34, 0, 0, 271, 37, 1, 0, 0, 0, 272, 274, 5, 35, 0, 0, 273,
		275, 5, 37, 0, 0, 274, 273, 1, 0, 0, 0, 274, 275, 1, 0, 0, 0, 275, 276,
		1, 0, 0, 0, 276, 277, 5, 34, 0, 0, 277, 39, 1, 0, 0, 0, 278, 280, 5, 36,
		0, 0, 279, 281, 5, 37, 0, 0, 280, 279, 1, 0, 0, 0, 280, 281, 1, 0, 0, 0,
		281, 282, 1, 0, 0, 0, 282, 283, 5, 34, 0, 0, 283, 41, 1, 0, 0, 0, 284,
		286, 5, 43, 0, 0, 285, 287, 5, 37, 0, 0, 286, 285, 1, 0, 0, 0, 286, 287,
		1, 0, 0, 0, 287, 288, 1, 0, 0, 0, 288, 289, 5, 34, 0, 0, 289, 43, 1, 0,
		0, 0, 290, 291, 5, 39, 0, 0, 291, 292, 5, 34, 0, 0, 292, 45, 1, 0, 0, 0,
		293, 298, 3, 36, 18, 0, 294, 298, 3, 38, 19, 0, 295, 298, 3, 40, 20, 0,
		296, 298, 3, 42, 21, 0, 297, 293, 1, 0, 0, 0, 297, 294, 1, 0, 0, 0, 297,
		295, 1, 0, 0, 0, 297, 296, 1, 0, 0, 0, 298, 47, 1, 0, 0, 0, 299, 300, 3,
		52, 26, 0, 300, 301, 5, 40, 0, 0, 301, 302, 3, 12, 6, 0, 302, 49, 1, 0,
		0, 0, 303, 304, 3, 52, 26, 0, 304, 51, 1, 0, 0, 0, 305, 306, 6, 26, -1,
		0, 306, 312, 3, 70, 35, 0, 307, 308, 5, 15, 0, 0, 308, 309, 3, 52, 26,
		0, 309, 310, 5, 17, 0, 0, 310, 312, 1, 0, 0, 0, 311, 305, 1, 0, 0, 0, 311,
		307, 1, 0, 0, 0, 312, 323, 1, 0, 0, 0, 313, 314, 10, 3, 0, 0, 314, 315,
		3, 46, 23, 0, 315, 316, 3, 48, 24, 0, 316, 322, 1, 0, 0, 0, 317, 318, 10,
		2, 0, 0, 318, 319, 3, 44, 22, 0, 319, 320, 3, 50, 25, 0, 320, 322, 1, 0,
		0, 0, 321, 313, 1, 0, 0, 0, 321, 317, 1, 0, 0, 0, 322, 325, 1, 0, 0, 0,
		323, 321, 1, 0, 0, 0, 323, 324, 1, 0, 0, 0, 324, 53, 1, 0, 0, 0, 325, 323,
		1, 0, 0, 0, 326, 331, 3, 52, 26, 0, 327, 328, 5, 16, 0, 0, 328, 330, 3,
		52, 26, 0, 329, 327, 1, 0, 0, 0, 330, 333, 1, 0, 0, 0, 331, 329, 1, 0,
		0, 0, 331, 332, 1, 0, 0, 0, 332, 55, 1, 0, 0, 0, 333, 331, 1, 0, 0, 0,
		334, 342, 3, 58, 29, 0, 335, 337, 5, 49, 0, 0, 336, 338, 5, 50, 0, 0, 337,
		336, 1, 0, 0, 0, 337, 338, 1, 0, 0, 0, 338, 339, 1, 0, 0, 0, 339, 341,
		3, 56, 28, 0, 340, 335, 1, 0, 0, 0, 341, 344, 1, 0, 0, 0, 342, 340, 1,
		0, 0, 0, 342, 343, 1, 0, 0, 0, 343, 57, 1, 0, 0, 0, 344, 342, 1, 0, 0,
		0, 345, 346, 5, 29, 0, 0, 346, 347, 3, 28, 14, 0, 347, 348, 5, 22, 0, 0,
		348, 350, 3, 54, 27, 0, 349, 351, 3, 14, 7, 0, 350, 349, 1, 0, 0, 0, 350,
		351, 1, 0, 0, 0, 351, 353, 1, 0, 0, 0, 352, 354, 3, 82, 41, 0, 353, 352,
		1, 0, 0, 0, 353, 354, 1, 0, 0, 0, 354, 356, 1, 0, 0, 0, 355, 357, 3, 78,
		39, 0, 356, 355, 1, 0, 0, 0, 356, 357, 1, 0, 0, 0, 357, 359, 1, 0, 0, 0,
		358, 360, 3, 88, 44, 0, 359, 358, 1, 0, 0, 0, 359, 360, 1, 0, 0, 0, 360,
		362, 1, 0, 0, 0, 361, 363, 3, 16, 8, 0, 362, 361, 1, 0, 0, 0, 362, 363,
		1, 0, 0, 0, 363, 59, 1, 0, 0, 0, 364, 366, 7, 4, 0, 0, 365, 364, 1, 0,
		0, 0, 365, 366, 1, 0, 0, 0, 366, 367, 1, 0, 0, 0, 367, 368, 5, 74, 0, 0,
		368, 61, 1, 0, 0, 0, 369, 370, 5, 75, 0, 0, 370, 63, 1, 0, 0, 0, 371, 372,
		5, 26, 0, 0, 372, 65, 1, 0, 0, 0, 373, 377, 3, 60, 30, 0, 374, 377, 3,
		62, 31, 0, 375, 377, 3, 64, 32, 0, 376, 373, 1, 0, 0, 0, 376, 374, 1, 0,
		0, 0, 376, 375, 1, 0, 0, 0, 377, 67, 1, 0, 0, 0, 378, 379, 5, 71, 0, 0,
		379, 69, 1, 0, 0, 0, 380, 382, 3, 74, 37, 0, 381, 383, 3, 68, 34, 0, 382,
		381, 1, 0, 0, 0, 382, 383, 1, 0, 0, 0, 383, 71, 1, 0, 0, 0, 384, 385, 3,
		76, 38, 0, 385, 386, 5, 18, 0, 0, 386, 388, 1, 0, 0, 0, 387, 384, 1, 0,
		0, 0, 387, 388, 1, 0, 0, 0, 388, 389, 1, 0, 0, 0, 389, 390, 3, 74, 37,
		0, 390, 73, 1, 0, 0, 0, 391, 392, 5, 71, 0, 0, 392, 75, 1, 0, 0, 0, 393,
		394, 5, 71, 0, 0, 394, 77, 1, 0, 0, 0, 395, 396, 5, 44, 0, 0, 396, 397,
		3, 80, 40, 0, 397, 79, 1, 0, 0, 0, 398, 399, 5, 74, 0, 0, 399, 81, 1, 0,
		0, 0, 400, 401, 5, 45, 0, 0, 401, 402, 5, 46, 0, 0, 402, 407, 3, 84, 42,
		0, 403, 404, 5, 16, 0, 0, 404, 406, 3, 84, 42, 0, 405, 403, 1, 0, 0, 0,
		406, 409, 1, 0, 0, 0, 407, 405, 1, 0, 0, 0, 407, 408, 1, 0, 0, 0, 408,
		83, 1, 0, 0, 0, 409, 407, 1, 0, 0, 0, 410, 413, 3, 72, 36, 0, 411, 413,
		3, 86, 43, 0, 412, 410, 1, 0, 0, 0, 412, 411, 1, 0, 0, 0, 413, 415, 1,
		0, 0, 0, 414, 416, 7, 5, 0, 0, 415, 414, 1, 0, 0, 0, 415, 416, 1, 0, 0,
		0, 416, 85, 1, 0, 0, 0, 417, 418, 5, 74, 0, 0, 418, 87, 1, 0, 0, 0, 419,
		420, 5, 69, 0, 0, 420, 421, 5, 46, 0, 0, 421, 426, 3, 90, 45, 0, 422, 423,
		5, 16, 0, 0, 423, 425, 3, 90, 45, 0, 424, 422, 1, 0, 0, 0, 425, 428, 1,
		0, 0, 0, 426, 424, 1, 0, 0, 0, 426, 427, 1, 0, 0, 0, 427, 89, 1, 0, 0,
		0, 428, 426, 1, 0, 0, 0, 429, 432, 3, 72, 36, 0, 430, 432, 3, 86, 43, 0,
		431, 429, 1, 0, 0, 0, 431, 430, 1, 0, 0, 0, 432, 91, 1, 0, 0, 0, 433, 434,
		7, 6, 0, 0, 434, 93, 1, 0, 0, 0, 46, 96, 110, 117, 127, 135, 151, 158,
		165, 173, 181, 193, 195, 213, 224, 229, 237, 244, 250, 253, 256, 261, 268,
		274, 280, 286, 297, 311, 321, 323, 331, 337, 342, 350, 353, 356, 359, 362,
		365, 376, 382, 387, 407, 412, 415, 426, 431,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// CsqlParserInit initializes any static state used to implement CsqlParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewCsqlParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func CsqlParserInit() {
	staticData := &CsqlParserStaticData
	staticData.once.Do(csqlParserInit)
}

// NewCsqlParser produces a new parser instance for the optional input antlr.TokenStream.
func NewCsqlParser(input antlr.TokenStream) *CsqlParser {
	CsqlParserInit()
	this := new(CsqlParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &CsqlParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "Csql.g4"

	return this
}

// CsqlParser tokens.
const (
	CsqlParserEOF                 = antlr.TokenEOF
	CsqlParserT__0                = 1
	CsqlParserT__1                = 2
	CsqlParserT__2                = 3
	CsqlParserT__3                = 4
	CsqlParserT__4                = 5
	CsqlParserT__5                = 6
	CsqlParserT__6                = 7
	CsqlParserT__7                = 8
	CsqlParserT__8                = 9
	CsqlParserT__9                = 10
	CsqlParserT__10               = 11
	CsqlParserT__11               = 12
	CsqlParserT__12               = 13
	CsqlParserT__13               = 14
	CsqlParserT__14               = 15
	CsqlParserT__15               = 16
	CsqlParserT__16               = 17
	CsqlParserT__17               = 18
	CsqlParserK_AND               = 19
	CsqlParserK_AS                = 20
	CsqlParserK_BETWEEN           = 21
	CsqlParserK_FROM              = 22
	CsqlParserK_IN                = 23
	CsqlParserK_IS                = 24
	CsqlParserK_NOT               = 25
	CsqlParserK_NULL              = 26
	CsqlParserK_OR                = 27
	CsqlParserK_REGEX             = 28
	CsqlParserK_SELECT            = 29
	CsqlParserK_MATCH             = 30
	CsqlParserK_WHERE             = 31
	CsqlParserK_TRUE              = 32
	CsqlParserK_FALSE             = 33
	CsqlParserK_JOIN              = 34
	CsqlParserK_LEFT              = 35
	CsqlParserK_RIGHT             = 36
	CsqlParserK_OUTER             = 37
	CsqlParserK_INNER             = 38
	CsqlParserK_CROSS             = 39
	CsqlParserK_ON                = 40
	CsqlParserK_DISTINCT          = 41
	CsqlParserK_LIKE              = 42
	CsqlParserK_FULL              = 43
	CsqlParserK_LIMIT             = 44
	CsqlParserK_ORDER             = 45
	CsqlParserK_BY                = 46
	CsqlParserK_ASC               = 47
	CsqlParserK_DESC              = 48
	CsqlParserK_UNION             = 49
	CsqlParserK_ALL               = 50
	CsqlParserK_TO_STRING         = 51
	CsqlParserK_TO_FLOAT          = 52
	CsqlParserK_TO_INT            = 53
	CsqlParserK_ROUND             = 54
	CsqlParserK_LEN               = 55
	CsqlParserK_TRUNC             = 56
	CsqlParserK_FRAC              = 57
	CsqlParserK_SUBSTRING         = 58
	CsqlParserK_TO_UPPER          = 59
	CsqlParserK_TO_LOWER          = 60
	CsqlParserK_POW               = 61
	CsqlParserK_SQRT              = 62
	CsqlParserK_COALESCE          = 63
	CsqlParserK_COUNT             = 64
	CsqlParserK_SUM               = 65
	CsqlParserK_AVG               = 66
	CsqlParserK_MIN               = 67
	CsqlParserK_MAX               = 68
	CsqlParserK_GROUP             = 69
	CsqlParserK_HAVING            = 70
	CsqlParserIDENTIFIER          = 71
	CsqlParserSIMPLE_IDENTIFIER   = 72
	CsqlParserQUOTED_IDENTIFIER   = 73
	CsqlParserNUMERIC_LITERAL     = 74
	CsqlParserSTRING_LITERAL      = 75
	CsqlParserSINGLE_LINE_COMMENT = 76
	CsqlParserMULTILINE_COMMENT   = 77
	CsqlParserSPACES              = 78
	CsqlParserUNEXPECTED_CHAR     = 79
)

// CsqlParser rules.
const (
	CsqlParserRULE_query                   = 0
	CsqlParserRULE_comparisonOperator      = 1
	CsqlParserRULE_binaryOperation         = 2
	CsqlParserRULE_list                    = 3
	CsqlParserRULE_term                    = 4
	CsqlParserRULE_valueExpr               = 5
	CsqlParserRULE_whereExpr               = 6
	CsqlParserRULE_where                   = 7
	CsqlParserRULE_having                  = 8
	CsqlParserRULE_distinct                = 9
	CsqlParserRULE_funCall                 = 10
	CsqlParserRULE_aggregateFunCall        = 11
	CsqlParserRULE_countFunc               = 12
	CsqlParserRULE_aggregateFunc           = 13
	CsqlParserRULE_projection              = 14
	CsqlParserRULE_projectionField         = 15
	CsqlParserRULE_projectionFieldName     = 16
	CsqlParserRULE_fieldName               = 17
	CsqlParserRULE_innerJoin               = 18
	CsqlParserRULE_leftJoin                = 19
	CsqlParserRULE_rightJoin               = 20
	CsqlParserRULE_fullJoin                = 21
	CsqlParserRULE_crossJoin               = 22
	CsqlParserRULE_conditionalJoinType     = 23
	CsqlParserRULE_conditionalJoinTarget   = 24
	CsqlParserRULE_unconditionalJoinTarget = 25
	CsqlParserRULE_dataSource              = 26
	CsqlParserRULE_sources                 = 27
	CsqlParserRULE_unionSelects            = 28
	CsqlParserRULE_selectStatement         = 29
	CsqlParserRULE_signedNumber            = 30
	CsqlParserRULE_stringValue             = 31
	CsqlParserRULE_nullValue               = 32
	CsqlParserRULE_literalValue            = 33
	CsqlParserRULE_alias                   = 34
	CsqlParserRULE_sourceName              = 35
	CsqlParserRULE_compoundName            = 36
	CsqlParserRULE_name                    = 37
	CsqlParserRULE_qualifier               = 38
	CsqlParserRULE_limit                   = 39
	CsqlParserRULE_limitValue              = 40
	CsqlParserRULE_orderBy                 = 41
	CsqlParserRULE_orderByField            = 42
	CsqlParserRULE_fieldIndex              = 43
	CsqlParserRULE_groupBy                 = 44
	CsqlParserRULE_groupByField            = 45
	CsqlParserRULE_function                = 46
)

// IQueryContext is an interface to support dynamic dispatch.
type IQueryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	UnionSelects() IUnionSelectsContext
	EOF() antlr.TerminalNode

	// IsQueryContext differentiates from other interfaces.
	IsQueryContext()
}

type QueryContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyQueryContext() *QueryContext {
	var p = new(QueryContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CsqlParserRULE_query
	return p
}

func InitEmptyQueryContext(p *QueryContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CsqlParserRULE_query
}

func (*QueryContext) IsQueryContext() {}

func NewQueryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *QueryContext {
	var p = new(QueryContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CsqlParserRULE_query

	return p
}

func (s *QueryContext) GetParser() antlr.Parser { return s.parser }

func (s *QueryContext) UnionSelects() IUnionSelectsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUnionSelectsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUnionSelectsContext)
}

func (s *QueryContext) EOF() antlr.TerminalNode {
	return s.GetToken(CsqlParserEOF, 0)
}

func (s *QueryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *QueryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *QueryContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CsqlVisitor:
		return t.VisitQuery(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CsqlParser) Query() (localctx IQueryContext) {
	localctx = NewQueryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, CsqlParserRULE_query)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(94)
		p.UnionSelects()
	}
	p.SetState(96)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == CsqlParserT__0 {
		{
			p.SetState(95)
			p.Match(CsqlParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(98)
		p.Match(CsqlParserEOF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IComparisonOperatorContext is an interface to support dynamic dispatch.
type IComparisonOperatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsComparisonOperatorContext differentiates from other interfaces.
	IsComparisonOperatorContext()
}

type ComparisonOperatorContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyComparisonOperatorContext() *ComparisonOperatorContext {
	var p = new(ComparisonOperatorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CsqlParserRULE_comparisonOperator
	return p
}

func InitEmptyComparisonOperatorContext(p *ComparisonOperatorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CsqlParserRULE_comparisonOperator
}

func (*ComparisonOperatorContext) IsComparisonOperatorContext() {}

func NewComparisonOperatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ComparisonOperatorContext {
	var p = new(ComparisonOperatorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CsqlParserRULE_comparisonOperator

	return p
}

func (s *ComparisonOperatorContext) GetParser() antlr.Parser { return s.parser }
func (s *ComparisonOperatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ComparisonOperatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ComparisonOperatorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CsqlVisitor:
		return t.VisitComparisonOperator(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CsqlParser) ComparisonOperator() (localctx IComparisonOperatorContext) {
	localctx = NewComparisonOperatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, CsqlParserRULE_comparisonOperator)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(100)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&508) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBinaryOperationContext is an interface to support dynamic dispatch.
type IBinaryOperationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsBinaryOperationContext differentiates from other interfaces.
	IsBinaryOperationContext()
}

type BinaryOperationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBinaryOperationContext() *BinaryOperationContext {
	var p = new(BinaryOperationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CsqlParserRULE_binaryOperation
	return p
}

func InitEmptyBinaryOperationContext(p *BinaryOperationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CsqlParserRULE_binaryOperation
}

func (*BinaryOperationContext) IsBinaryOperationContext() {}

func NewBinaryOperationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BinaryOperationContext {
	var p = new(BinaryOperationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CsqlParserRULE_binaryOperation

	return p
}

func (s *BinaryOperationContext) GetParser() antlr.Parser { return s.parser }
func (s *BinaryOperationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BinaryOperationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BinaryOperationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CsqlVisitor:
		return t.VisitBinaryOperation(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CsqlParser) BinaryOperation() (localctx IBinaryOperationContext) {
	localctx = NewBinaryOperationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, CsqlParserRULE_binaryOperation)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(102)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&32256) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IListContext is an interface to support dynamic dispatch.
type IListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllValueExpr() []IValueExprContext
	ValueExpr(i int) IValueExprContext

	// IsListContext differentiates from other interfaces.
	IsListContext()
}

type ListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyListContext() *ListContext {
	var p = new(ListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CsqlParserRULE_list
	return p
}

func InitEmptyListContext(p *ListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CsqlParserRULE_list
}

func (*ListContext) IsListContext() {}

func NewListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListContext {
	var p = new(ListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CsqlParserRULE_list

	return p
}

func (s *ListContext) GetParser() antlr.Parser { return s.parser }

func (s *ListContext) AllValueExpr() []IValueExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IValueExprContext); ok {
			len++
		}
	}

	tst := make([]IValueExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IValueExprContext); ok {
			tst[i] = t.(IValueExprContext)
			i++
		}
	}

	return tst
}

func (s *ListContext) ValueExpr(i int) IValueExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueExprContext)
}

func (s *ListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CsqlVisitor:
		return t.VisitList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CsqlParser) List() (localctx IListContext) {
	localctx = NewListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, CsqlParserRULE_list)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(104)
		p.Match(CsqlParserT__14)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(105)
		p.valueExpr(0)
	}
	p.SetState(110)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == CsqlParserT__15 {
		{
			p.SetState(106)
			p.Match(CsqlParserT__15)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(107)
			p.valueExpr(0)
		}

		p.SetState(112)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(113)
		p.Match(CsqlParserT__16)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITermContext is an interface to support dynamic dispatch.
type ITermContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	CompoundName() ICompoundNameContext
	LiteralValue() ILiteralValueContext

	// IsTermContext differentiates from other interfaces.
	IsTermContext()
}

type TermContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTermContext() *TermContext {
	var p = new(TermContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CsqlParserRULE_term
	return p
}

func InitEmptyTermContext(p *TermContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CsqlParserRULE_term
}

func (*TermContext) IsTermContext() {}

func NewTermContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TermContext {
	var p = new(TermContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CsqlParserRULE_term

	return p
}

func (s *TermContext) GetParser() antlr.Parser { return s.parser }

func (s *TermContext) CompoundName() ICompoundNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICompoundNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICompoundNameContext)
}

func (s *TermContext) LiteralValue() ILiteralValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILiteralValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILiteralValueContext)
}

func (s *TermContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TermContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TermContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CsqlVisitor:
		return t.VisitTerm(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CsqlParser) Term() (localctx ITermContext) {
	localctx = NewTermContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, CsqlParserRULE_term)
	p.SetState(117)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case CsqlParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(115)
			p.CompoundName()
		}

	case CsqlParserT__8, CsqlParserT__9, CsqlParserK_NULL, CsqlParserNUMERIC_LITERAL, CsqlParserSTRING_LITERAL:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(116)
			p.LiteralValue()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IValueExprContext is an interface to support dynamic dispatch.
type IValueExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsValueExprContext differentiates from other interfaces.
	IsValueExprContext()
}

type ValueExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValueExprContext() *ValueExprContext {
	var p = new(ValueExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CsqlParserRULE_valueExpr
	return p
}

func InitEmptyValueExprContext(p *ValueExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CsqlParserRULE_valueExpr
}

func (*ValueExprContext) IsValueExprContext() {}

func NewValueExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueExprContext {
	var p = new(ValueExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CsqlParserRULE_valueExpr

	return p
}

func (s *ValueExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueExprContext) CopyAll(ctx *ValueExprContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ValueExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ValueTermContext struct {
	ValueExprContext
}

func NewValueTermContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ValueTermContext {
	var p = new(ValueTermContext)

	InitEmptyValueExprContext(&p.ValueExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ValueExprContext))

	return p
}

func (s *ValueTermContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueTermContext) Term() ITermContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITermContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITermContext)
}

func (s *ValueTermContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CsqlVisitor:
		return t.VisitValueTerm(s)

	default:
		return t.VisitChildren(s)
	}
}

type ValueBinaryExprContext struct {
	ValueExprContext
}

func NewValueBinaryExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ValueBinaryExprContext {
	var p = new(ValueBinaryExprContext)

	InitEmptyValueExprContext(&p.ValueExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ValueExprContext))

	return p
}

func (s *ValueBinaryExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueBinaryExprContext) AllValueExpr() []IValueExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IValueExprContext); ok {
			len++
		}
	}

	tst := make([]IValueExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IValueExprContext); ok {
			tst[i] = t.(IValueExprContext)
			i++
		}
	}

	return tst
}

func (s *ValueBinaryExprContext) ValueExpr(i int) IValueExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueExprContext)
}

func (s *ValueBinaryExprContext) BinaryOperation() IBinaryOperationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBinaryOperationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBinaryOperationContext)
}

func (s *ValueBinaryExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CsqlVisitor:
		return t.VisitValueBinaryExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type ValueParensExprContext struct {
	ValueExprContext
}

func NewValueParensExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ValueParensExprContext {
	var p = new(ValueParensExprContext)

	InitEmptyValueExprContext(&p.ValueExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ValueExprContext))

	return p
}

func (s *ValueParensExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueParensExprContext) ValueExpr() IValueExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueExprContext)
}

func (s *ValueParensExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CsqlVisitor:
		return t.VisitValueParensExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type FunctionExprContext struct {
	ValueExprContext
}

func NewFunctionExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FunctionExprContext {
	var p = new(FunctionExprContext)

	InitEmptyValueExprContext(&p.ValueExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ValueExprContext))

	return p
}

func (s *FunctionExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionExprContext) FunCall() IFunCallContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunCallContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunCallContext)
}

func (s *FunctionExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CsqlVisitor:
		return t.VisitFunctionExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type AggregateFuncExprContext struct {
	ValueExprContext
}

func NewAggregateFuncExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AggregateFuncExprContext {
	var p = new(AggregateFuncExprContext)

	InitEmptyValueExprContext(&p.ValueExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ValueExprContext))

	return p
}

func (s *AggregateFuncExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AggregateFuncExprContext) AggregateFunCall() IAggregateFunCallContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAggregateFunCallContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAggregateFunCallContext)
}

func (s *AggregateFuncExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CsqlVisitor:
		return t.VisitAggregateFuncExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CsqlParser) ValueExpr() (localctx IValueExprContext) {
	return p.valueExpr(0)
}

func (p *CsqlParser) valueExpr(_p int) (localctx IValueExprContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewValueExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IValueExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 10
	p.EnterRecursionRule(localctx, 10, CsqlParserRULE_valueExpr, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(127)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case CsqlParserT__8, CsqlParserT__9, CsqlParserK_NULL, CsqlParserIDENTIFIER, CsqlParserNUMERIC_LITERAL, CsqlParserSTRING_LITERAL:
		localctx = NewValueTermContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(120)
			p.Term()
		}

	case CsqlParserK_TO_STRING, CsqlParserK_TO_FLOAT, CsqlParserK_TO_INT, CsqlParserK_ROUND, CsqlParserK_LEN, CsqlParserK_TRUNC, CsqlParserK_FRAC, CsqlParserK_SUBSTRING, CsqlParserK_TO_UPPER, CsqlParserK_TO_LOWER, CsqlParserK_POW, CsqlParserK_SQRT, CsqlParserK_COALESCE:
		localctx = NewFunctionExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(121)
			p.FunCall()
		}

	case CsqlParserK_COUNT, CsqlParserK_SUM, CsqlParserK_AVG, CsqlParserK_MIN, CsqlParserK_MAX:
		localctx = NewAggregateFuncExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(122)
			p.AggregateFunCall()
		}

	case CsqlParserT__14:
		localctx = NewValueParensExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(123)
			p.Match(CsqlParserT__14)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(124)
			p.valueExpr(0)
		}
		{
			p.SetState(125)
			p.Match(CsqlParserT__16)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(135)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewValueBinaryExprContext(p, NewValueExprContext(p, _parentctx, _parentState))
			p.PushNewRecursionContext(localctx, _startState, CsqlParserRULE_valueExpr)
			p.SetState(129)

			if !(p.Precpred(p.GetParserRuleContext(), 5)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
				goto errorExit
			}
			{
				p.SetState(130)
				p.BinaryOperation()
			}
			{
				p.SetState(131)
				p.valueExpr(6)
			}

		}
		p.SetState(137)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IWhereExprContext is an interface to support dynamic dispatch.
type IWhereExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsWhereExprContext differentiates from other interfaces.
	IsWhereExprContext()
}

type WhereExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWhereExprContext() *WhereExprContext {
	var p = new(WhereExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CsqlParserRULE_whereExpr
	return p
}

func InitEmptyWhereExprContext(p *WhereExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CsqlParserRULE_whereExpr
}

func (*WhereExprContext) IsWhereExprContext() {}

func NewWhereExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WhereExprContext {
	var p = new(WhereExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CsqlParserRULE_whereExpr

	return p
}

func (s *WhereExprContext) GetParser() antlr.Parser { return s.parser }

func (s *WhereExprContext) CopyAll(ctx *WhereExprContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *WhereExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WhereExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ConditionContext struct {
	WhereExprContext
}

func NewConditionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ConditionContext {
	var p = new(ConditionContext)

	InitEmptyWhereExprContext(&p.WhereExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*WhereExprContext))

	return p
}

func (s *ConditionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConditionContext) AllWhereExpr() []IWhereExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IWhereExprContext); ok {
			len++
		}
	}

	tst := make([]IWhereExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IWhereExprContext); ok {
			tst[i] = t.(IWhereExprContext)
			i++
		}
	}

	return tst
}

func (s *ConditionContext) WhereExpr(i int) IWhereExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWhereExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWhereExprContext)
}

func (s *ConditionContext) ComparisonOperator() IComparisonOperatorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IComparisonOperatorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IComparisonOperatorContext)
}

func (s *ConditionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CsqlVisitor:
		return t.VisitCondition(s)

	default:
		return t.VisitChildren(s)
	}
}

type NotExprContext struct {
	WhereExprContext
}

func NewNotExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NotExprContext {
	var p = new(NotExprContext)

	InitEmptyWhereExprContext(&p.WhereExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*WhereExprContext))

	return p
}

func (s *NotExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NotExprContext) K_NOT() antlr.TerminalNode {
	return s.GetToken(CsqlParserK_NOT, 0)
}

func (s *NotExprContext) WhereExpr() IWhereExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWhereExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWhereExprContext)
}

func (s *NotExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CsqlVisitor:
		return t.VisitNotExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type TermItemContext struct {
	WhereExprContext
}

func NewTermItemContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TermItemContext {
	var p = new(TermItemContext)

	InitEmptyWhereExprContext(&p.WhereExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*WhereExprContext))

	return p
}

func (s *TermItemContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TermItemContext) Term() ITermContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITermContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITermContext)
}

func (s *TermItemContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CsqlVisitor:
		return t.VisitTermItem(s)

	default:
		return t.VisitChildren(s)
	}
}

type InExprContext struct {
	WhereExprContext
}

func NewInExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *InExprContext {
	var p = new(InExprContext)

	InitEmptyWhereExprContext(&p.WhereExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*WhereExprContext))

	return p
}

func (s *InExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InExprContext) ValueExpr() IValueExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueExprContext)
}

func (s *InExprContext) K_IN() antlr.TerminalNode {
	return s.GetToken(CsqlParserK_IN, 0)
}

func (s *InExprContext) List() IListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListContext)
}

func (s *InExprContext) K_NOT() antlr.TerminalNode {
	return s.GetToken(CsqlParserK_NOT, 0)
}

func (s *InExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CsqlVisitor:
		return t.VisitInExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type IsNullExprContext struct {
	WhereExprContext
}

func NewIsNullExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IsNullExprContext {
	var p = new(IsNullExprContext)

	InitEmptyWhereExprContext(&p.WhereExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*WhereExprContext))

	return p
}

func (s *IsNullExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IsNullExprContext) ValueExpr() IValueExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueExprContext)
}

func (s *IsNullExprContext) K_IS() antlr.TerminalNode {
	return s.GetToken(CsqlParserK_IS, 0)
}

func (s *IsNullExprContext) K_NULL() antlr.TerminalNode {
	return s.GetToken(CsqlParserK_NULL, 0)
}

func (s *IsNullExprContext) K_NOT() antlr.TerminalNode {
	return s.GetToken(CsqlParserK_NOT, 0)
}

func (s *IsNullExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CsqlVisitor:
		return t.VisitIsNullExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type LikeExprContext struct {
	WhereExprContext
}

func NewLikeExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LikeExprContext {
	var p = new(LikeExprContext)

	InitEmptyWhereExprContext(&p.WhereExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*WhereExprContext))

	return p
}

func (s *LikeExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LikeExprContext) AllValueExpr() []IValueExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IValueExprContext); ok {
			len++
		}
	}

	tst := make([]IValueExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IValueExprContext); ok {
			tst[i] = t.(IValueExprContext)
			i++
		}
	}

	return tst
}

func (s *LikeExprContext) ValueExpr(i int) IValueExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueExprContext)
}

func (s *LikeExprContext) K_LIKE() antlr.TerminalNode {
	return s.GetToken(CsqlParserK_LIKE, 0)
}

func (s *LikeExprContext) K_NOT() antlr.TerminalNode {
	return s.GetToken(CsqlParserK_NOT, 0)
}

func (s *LikeExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CsqlVisitor:
		return t.VisitLikeExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type MatchExprContext struct {
	WhereExprContext
}

func NewMatchExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MatchExprContext {
	var p = new(MatchExprContext)

	InitEmptyWhereExprContext(&p.WhereExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*WhereExprContext))

	return p
}

func (s *MatchExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MatchExprContext) AllValueExpr() []IValueExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IValueExprContext); ok {
			len++
		}
	}

	tst := make([]IValueExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IValueExprContext); ok {
			tst[i] = t.(IValueExprContext)
			i++
		}
	}

	return tst
}

func (s *MatchExprContext) ValueExpr(i int) IValueExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueExprContext)
}

func (s *MatchExprContext) K_MATCH() antlr.TerminalNode {
	return s.GetToken(CsqlParserK_MATCH, 0)
}

func (s *MatchExprContext) K_NOT() antlr.TerminalNode {
	return s.GetToken(CsqlParserK_NOT, 0)
}

func (s *MatchExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CsqlVisitor:
		return t.VisitMatchExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type OrExprContext struct {
	WhereExprContext
}

func NewOrExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *OrExprContext {
	var p = new(OrExprContext)

	InitEmptyWhereExprContext(&p.WhereExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*WhereExprContext))

	return p
}

func (s *OrExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OrExprContext) AllWhereExpr() []IWhereExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IWhereExprContext); ok {
			len++
		}
	}

	tst := make([]IWhereExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IWhereExprContext); ok {
			tst[i] = t.(IWhereExprContext)
			i++
		}
	}

	return tst
}

func (s *OrExprContext) WhereExpr(i int) IWhereExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWhereExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWhereExprContext)
}

func (s *OrExprContext) K_OR() antlr.TerminalNode {
	return s.GetToken(CsqlParserK_OR, 0)
}

func (s *OrExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CsqlVisitor:
		return t.VisitOrExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type ParensExprContext struct {
	WhereExprContext
}

func NewParensExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParensExprContext {
	var p = new(ParensExprContext)

	InitEmptyWhereExprContext(&p.WhereExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*WhereExprContext))

	return p
}

func (s *ParensExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParensExprContext) WhereExpr() IWhereExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWhereExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWhereExprContext)
}

func (s *ParensExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CsqlVisitor:
		return t.VisitParensExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type ValueExprItemContext struct {
	WhereExprContext
}

func NewValueExprItemContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ValueExprItemContext {
	var p = new(ValueExprItemContext)

	InitEmptyWhereExprContext(&p.WhereExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*WhereExprContext))

	return p
}

func (s *ValueExprItemContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueExprItemContext) ValueExpr() IValueExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueExprContext)
}

func (s *ValueExprItemContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CsqlVisitor:
		return t.VisitValueExprItem(s)

	default:
		return t.VisitChildren(s)
	}
}

type BetweenExprContext struct {
	WhereExprContext
}

func NewBetweenExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BetweenExprContext {
	var p = new(BetweenExprContext)

	InitEmptyWhereExprContext(&p.WhereExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*WhereExprContext))

	return p
}

func (s *BetweenExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BetweenExprContext) AllValueExpr() []IValueExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IValueExprContext); ok {
			len++
		}
	}

	tst := make([]IValueExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IValueExprContext); ok {
			tst[i] = t.(IValueExprContext)
			i++
		}
	}

	return tst
}

func (s *BetweenExprContext) ValueExpr(i int) IValueExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueExprContext)
}

func (s *BetweenExprContext) K_BETWEEN() antlr.TerminalNode {
	return s.GetToken(CsqlParserK_BETWEEN, 0)
}

func (s *BetweenExprContext) K_AND() antlr.TerminalNode {
	return s.GetToken(CsqlParserK_AND, 0)
}

func (s *BetweenExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CsqlVisitor:
		return t.VisitBetweenExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type AndExprContext struct {
	WhereExprContext
}

func NewAndExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AndExprContext {
	var p = new(AndExprContext)

	InitEmptyWhereExprContext(&p.WhereExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*WhereExprContext))

	return p
}

func (s *AndExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AndExprContext) AllWhereExpr() []IWhereExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IWhereExprContext); ok {
			len++
		}
	}

	tst := make([]IWhereExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IWhereExprContext); ok {
			tst[i] = t.(IWhereExprContext)
			i++
		}
	}

	return tst
}

func (s *AndExprContext) WhereExpr(i int) IWhereExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWhereExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWhereExprContext)
}

func (s *AndExprContext) K_AND() antlr.TerminalNode {
	return s.GetToken(CsqlParserK_AND, 0)
}

func (s *AndExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CsqlVisitor:
		return t.VisitAndExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CsqlParser) WhereExpr() (localctx IWhereExprContext) {
	return p.whereExpr(0)
}

func (p *CsqlParser) whereExpr(_p int) (localctx IWhereExprContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewWhereExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IWhereExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 12
	p.EnterRecursionRule(localctx, 12, CsqlParserRULE_whereExpr, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(181)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 9, p.GetParserRuleContext()) {
	case 1:
		localctx = NewTermItemContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(139)
			p.Term()
		}

	case 2:
		localctx = NewValueExprItemContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(140)
			p.valueExpr(0)
		}

	case 3:
		localctx = NewNotExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(141)
			p.Match(CsqlParserK_NOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(142)
			p.whereExpr(7)
		}

	case 4:
		localctx = NewBetweenExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(143)
			p.valueExpr(0)
		}
		{
			p.SetState(144)
			p.Match(CsqlParserK_BETWEEN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(145)
			p.valueExpr(0)
		}
		{
			p.SetState(146)
			p.Match(CsqlParserK_AND)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(147)
			p.valueExpr(0)
		}

	case 5:
		localctx = NewLikeExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(149)
			p.valueExpr(0)
		}
		p.SetState(151)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == CsqlParserK_NOT {
			{
				p.SetState(150)
				p.Match(CsqlParserK_NOT)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(153)
			p.Match(CsqlParserK_LIKE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(154)
			p.valueExpr(0)
		}

	case 6:
		localctx = NewMatchExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(156)
			p.valueExpr(0)
		}
		p.SetState(158)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == CsqlParserK_NOT {
			{
				p.SetState(157)
				p.Match(CsqlParserK_NOT)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(160)
			p.Match(CsqlParserK_MATCH)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(161)
			p.valueExpr(0)
		}

	case 7:
		localctx = NewInExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(163)
			p.valueExpr(0)
		}
		p.SetState(165)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == CsqlParserK_NOT {
			{
				p.SetState(164)
				p.Match(CsqlParserK_NOT)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(167)
			p.Match(CsqlParserK_IN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(168)
			p.List()
		}

	case 8:
		localctx = NewIsNullExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(170)
			p.valueExpr(0)
		}
		{
			p.SetState(171)
			p.Match(CsqlParserK_IS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(173)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == CsqlParserK_NOT {
			{
				p.SetState(172)
				p.Match(CsqlParserK_NOT)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(175)
			p.Match(CsqlParserK_NULL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 9:
		localctx = NewParensExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(177)
			p.Match(CsqlParserT__14)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(178)
			p.whereExpr(0)
		}
		{
			p.SetState(179)
			p.Match(CsqlParserT__16)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(195)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 11, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(193)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 10, p.GetParserRuleContext()) {
			case 1:
				localctx = NewConditionContext(p, NewWhereExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, CsqlParserRULE_whereExpr)
				p.SetState(183)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
					goto errorExit
				}
				{
					p.SetState(184)
					p.ComparisonOperator()
				}
				{
					p.SetState(185)
					p.whereExpr(11)
				}

			case 2:
				localctx = NewAndExprContext(p, NewWhereExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, CsqlParserRULE_whereExpr)
				p.SetState(187)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
					goto errorExit
				}
				{
					p.SetState(188)
					p.Match(CsqlParserK_AND)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(189)
					p.whereExpr(10)
				}

			case 3:
				localctx = NewOrExprContext(p, NewWhereExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, CsqlParserRULE_whereExpr)
				p.SetState(190)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
					goto errorExit
				}
				{
					p.SetState(191)
					p.Match(CsqlParserK_OR)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(192)
					p.whereExpr(9)
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(197)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 11, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IWhereContext is an interface to support dynamic dispatch.
type IWhereContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	K_WHERE() antlr.TerminalNode
	WhereExpr() IWhereExprContext

	// IsWhereContext differentiates from other interfaces.
	IsWhereContext()
}

type WhereContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWhereContext() *WhereContext {
	var p = new(WhereContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CsqlParserRULE_where
	return p
}

func InitEmptyWhereContext(p *WhereContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CsqlParserRULE_where
}

func (*WhereContext) IsWhereContext() {}

func NewWhereContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WhereContext {
	var p = new(WhereContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CsqlParserRULE_where

	return p
}

func (s *WhereContext) GetParser() antlr.Parser { return s.parser }

func (s *WhereContext) K_WHERE() antlr.TerminalNode {
	return s.GetToken(CsqlParserK_WHERE, 0)
}

func (s *WhereContext) WhereExpr() IWhereExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWhereExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWhereExprContext)
}

func (s *WhereContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WhereContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WhereContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CsqlVisitor:
		return t.VisitWhere(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CsqlParser) Where() (localctx IWhereContext) {
	localctx = NewWhereContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, CsqlParserRULE_where)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(198)
		p.Match(CsqlParserK_WHERE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(199)
		p.whereExpr(0)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IHavingContext is an interface to support dynamic dispatch.
type IHavingContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	K_HAVING() antlr.TerminalNode
	WhereExpr() IWhereExprContext

	// IsHavingContext differentiates from other interfaces.
	IsHavingContext()
}

type HavingContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyHavingContext() *HavingContext {
	var p = new(HavingContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CsqlParserRULE_having
	return p
}

func InitEmptyHavingContext(p *HavingContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CsqlParserRULE_having
}

func (*HavingContext) IsHavingContext() {}

func NewHavingContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *HavingContext {
	var p = new(HavingContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CsqlParserRULE_having

	return p
}

func (s *HavingContext) GetParser() antlr.Parser { return s.parser }

func (s *HavingContext) K_HAVING() antlr.TerminalNode {
	return s.GetToken(CsqlParserK_HAVING, 0)
}

func (s *HavingContext) WhereExpr() IWhereExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWhereExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWhereExprContext)
}

func (s *HavingContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HavingContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *HavingContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CsqlVisitor:
		return t.VisitHaving(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CsqlParser) Having() (localctx IHavingContext) {
	localctx = NewHavingContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, CsqlParserRULE_having)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(201)
		p.Match(CsqlParserK_HAVING)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(202)
		p.whereExpr(0)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDistinctContext is an interface to support dynamic dispatch.
type IDistinctContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	K_DISTINCT() antlr.TerminalNode

	// IsDistinctContext differentiates from other interfaces.
	IsDistinctContext()
}

type DistinctContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDistinctContext() *DistinctContext {
	var p = new(DistinctContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CsqlParserRULE_distinct
	return p
}

func InitEmptyDistinctContext(p *DistinctContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CsqlParserRULE_distinct
}

func (*DistinctContext) IsDistinctContext() {}

func NewDistinctContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DistinctContext {
	var p = new(DistinctContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CsqlParserRULE_distinct

	return p
}

func (s *DistinctContext) GetParser() antlr.Parser { return s.parser }

func (s *DistinctContext) K_DISTINCT() antlr.TerminalNode {
	return s.GetToken(CsqlParserK_DISTINCT, 0)
}

func (s *DistinctContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DistinctContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DistinctContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CsqlVisitor:
		return t.VisitDistinct(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CsqlParser) Distinct() (localctx IDistinctContext) {
	localctx = NewDistinctContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, CsqlParserRULE_distinct)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(204)
		p.Match(CsqlParserK_DISTINCT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFunCallContext is an interface to support dynamic dispatch.
type IFunCallContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Function() IFunctionContext
	AllValueExpr() []IValueExprContext
	ValueExpr(i int) IValueExprContext

	// IsFunCallContext differentiates from other interfaces.
	IsFunCallContext()
}

type FunCallContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunCallContext() *FunCallContext {
	var p = new(FunCallContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CsqlParserRULE_funCall
	return p
}

func InitEmptyFunCallContext(p *FunCallContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CsqlParserRULE_funCall
}

func (*FunCallContext) IsFunCallContext() {}

func NewFunCallContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunCallContext {
	var p = new(FunCallContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CsqlParserRULE_funCall

	return p
}

func (s *FunCallContext) GetParser() antlr.Parser { return s.parser }

func (s *FunCallContext) Function() IFunctionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunctionContext)
}

func (s *FunCallContext) AllValueExpr() []IValueExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IValueExprContext); ok {
			len++
		}
	}

	tst := make([]IValueExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IValueExprContext); ok {
			tst[i] = t.(IValueExprContext)
			i++
		}
	}

	return tst
}

func (s *FunCallContext) ValueExpr(i int) IValueExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueExprContext)
}

func (s *FunCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunCallContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunCallContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CsqlVisitor:
		return t.VisitFunCall(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CsqlParser) FunCall() (localctx IFunCallContext) {
	localctx = NewFunCallContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, CsqlParserRULE_funCall)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(206)
		p.Function()
	}
	{
		p.SetState(207)
		p.Match(CsqlParserT__14)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(208)
		p.valueExpr(0)
	}
	p.SetState(213)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == CsqlParserT__15 {
		{
			p.SetState(209)
			p.Match(CsqlParserT__15)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(210)
			p.valueExpr(0)
		}

		p.SetState(215)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(216)
		p.Match(CsqlParserT__16)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAggregateFunCallContext is an interface to support dynamic dispatch.
type IAggregateFunCallContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AggregateFunc() IAggregateFuncContext
	ValueExpr() IValueExprContext
	CountFunc() ICountFuncContext

	// IsAggregateFunCallContext differentiates from other interfaces.
	IsAggregateFunCallContext()
}

type AggregateFunCallContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAggregateFunCallContext() *AggregateFunCallContext {
	var p = new(AggregateFunCallContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CsqlParserRULE_aggregateFunCall
	return p
}

func InitEmptyAggregateFunCallContext(p *AggregateFunCallContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CsqlParserRULE_aggregateFunCall
}

func (*AggregateFunCallContext) IsAggregateFunCallContext() {}

func NewAggregateFunCallContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AggregateFunCallContext {
	var p = new(AggregateFunCallContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CsqlParserRULE_aggregateFunCall

	return p
}

func (s *AggregateFunCallContext) GetParser() antlr.Parser { return s.parser }

func (s *AggregateFunCallContext) AggregateFunc() IAggregateFuncContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAggregateFuncContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAggregateFuncContext)
}

func (s *AggregateFunCallContext) ValueExpr() IValueExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueExprContext)
}

func (s *AggregateFunCallContext) CountFunc() ICountFuncContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICountFuncContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICountFuncContext)
}

func (s *AggregateFunCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AggregateFunCallContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AggregateFunCallContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CsqlVisitor:
		return t.VisitAggregateFunCall(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CsqlParser) AggregateFunCall() (localctx IAggregateFunCallContext) {
	localctx = NewAggregateFunCallContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, CsqlParserRULE_aggregateFunCall)
	p.SetState(224)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case CsqlParserK_SUM, CsqlParserK_AVG, CsqlParserK_MIN, CsqlParserK_MAX:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(218)
			p.AggregateFunc()
		}
		{
			p.SetState(219)
			p.Match(CsqlParserT__14)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(220)
			p.valueExpr(0)
		}
		{
			p.SetState(221)
			p.Match(CsqlParserT__16)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case CsqlParserK_COUNT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(223)
			p.CountFunc()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICountFuncContext is an interface to support dynamic dispatch.
type ICountFuncContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	K_COUNT() antlr.TerminalNode
	ProjectionFieldName() IProjectionFieldNameContext
	K_DISTINCT() antlr.TerminalNode

	// IsCountFuncContext differentiates from other interfaces.
	IsCountFuncContext()
}

type CountFuncContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCountFuncContext() *CountFuncContext {
	var p = new(CountFuncContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CsqlParserRULE_countFunc
	return p
}

func InitEmptyCountFuncContext(p *CountFuncContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CsqlParserRULE_countFunc
}

func (*CountFuncContext) IsCountFuncContext() {}

func NewCountFuncContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CountFuncContext {
	var p = new(CountFuncContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CsqlParserRULE_countFunc

	return p
}

func (s *CountFuncContext) GetParser() antlr.Parser { return s.parser }

func (s *CountFuncContext) K_COUNT() antlr.TerminalNode {
	return s.GetToken(CsqlParserK_COUNT, 0)
}

func (s *CountFuncContext) ProjectionFieldName() IProjectionFieldNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IProjectionFieldNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IProjectionFieldNameContext)
}

func (s *CountFuncContext) K_DISTINCT() antlr.TerminalNode {
	return s.GetToken(CsqlParserK_DISTINCT, 0)
}

func (s *CountFuncContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CountFuncContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CountFuncContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CsqlVisitor:
		return t.VisitCountFunc(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CsqlParser) CountFunc() (localctx ICountFuncContext) {
	localctx = NewCountFuncContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, CsqlParserRULE_countFunc)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(226)
		p.Match(CsqlParserK_COUNT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(227)
		p.Match(CsqlParserT__14)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(229)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == CsqlParserK_DISTINCT {
		{
			p.SetState(228)
			p.Match(CsqlParserK_DISTINCT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(231)
		p.ProjectionFieldName()
	}
	{
		p.SetState(232)
		p.Match(CsqlParserT__16)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAggregateFuncContext is an interface to support dynamic dispatch.
type IAggregateFuncContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	K_SUM() antlr.TerminalNode
	K_AVG() antlr.TerminalNode
	K_MIN() antlr.TerminalNode
	K_MAX() antlr.TerminalNode

	// IsAggregateFuncContext differentiates from other interfaces.
	IsAggregateFuncContext()
}

type AggregateFuncContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAggregateFuncContext() *AggregateFuncContext {
	var p = new(AggregateFuncContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CsqlParserRULE_aggregateFunc
	return p
}

func InitEmptyAggregateFuncContext(p *AggregateFuncContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CsqlParserRULE_aggregateFunc
}

func (*AggregateFuncContext) IsAggregateFuncContext() {}

func NewAggregateFuncContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AggregateFuncContext {
	var p = new(AggregateFuncContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CsqlParserRULE_aggregateFunc

	return p
}

func (s *AggregateFuncContext) GetParser() antlr.Parser { return s.parser }

func (s *AggregateFuncContext) K_SUM() antlr.TerminalNode {
	return s.GetToken(CsqlParserK_SUM, 0)
}

func (s *AggregateFuncContext) K_AVG() antlr.TerminalNode {
	return s.GetToken(CsqlParserK_AVG, 0)
}

func (s *AggregateFuncContext) K_MIN() antlr.TerminalNode {
	return s.GetToken(CsqlParserK_MIN, 0)
}

func (s *AggregateFuncContext) K_MAX() antlr.TerminalNode {
	return s.GetToken(CsqlParserK_MAX, 0)
}

func (s *AggregateFuncContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AggregateFuncContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AggregateFuncContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CsqlVisitor:
		return t.VisitAggregateFunc(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CsqlParser) AggregateFunc() (localctx IAggregateFuncContext) {
	localctx = NewAggregateFuncContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, CsqlParserRULE_aggregateFunc)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(234)
		_la = p.GetTokenStream().LA(1)

		if !((int64((_la-65)) & ^0x3f) == 0 && ((int64(1)<<(_la-65))&15) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IProjectionContext is an interface to support dynamic dispatch.
type IProjectionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllProjectionField() []IProjectionFieldContext
	ProjectionField(i int) IProjectionFieldContext
	Distinct() IDistinctContext

	// IsProjectionContext differentiates from other interfaces.
	IsProjectionContext()
}

type ProjectionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProjectionContext() *ProjectionContext {
	var p = new(ProjectionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CsqlParserRULE_projection
	return p
}

func InitEmptyProjectionContext(p *ProjectionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CsqlParserRULE_projection
}

func (*ProjectionContext) IsProjectionContext() {}

func NewProjectionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProjectionContext {
	var p = new(ProjectionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CsqlParserRULE_projection

	return p
}

func (s *ProjectionContext) GetParser() antlr.Parser { return s.parser }

func (s *ProjectionContext) AllProjectionField() []IProjectionFieldContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IProjectionFieldContext); ok {
			len++
		}
	}

	tst := make([]IProjectionFieldContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IProjectionFieldContext); ok {
			tst[i] = t.(IProjectionFieldContext)
			i++
		}
	}

	return tst
}

func (s *ProjectionContext) ProjectionField(i int) IProjectionFieldContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IProjectionFieldContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IProjectionFieldContext)
}

func (s *ProjectionContext) Distinct() IDistinctContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDistinctContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDistinctContext)
}

func (s *ProjectionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProjectionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProjectionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CsqlVisitor:
		return t.VisitProjection(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CsqlParser) Projection() (localctx IProjectionContext) {
	localctx = NewProjectionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, CsqlParserRULE_projection)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(237)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == CsqlParserK_DISTINCT {
		{
			p.SetState(236)
			p.Distinct()
		}

	}
	{
		p.SetState(239)
		p.ProjectionField()
	}
	p.SetState(244)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == CsqlParserT__15 {
		{
			p.SetState(240)
			p.Match(CsqlParserT__15)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(241)
			p.ProjectionField()
		}

		p.SetState(246)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IProjectionFieldContext is an interface to support dynamic dispatch.
type IProjectionFieldContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ProjectionFieldName() IProjectionFieldNameContext
	ValueExpr() IValueExprContext
	AggregateFunCall() IAggregateFunCallContext
	Alias() IAliasContext
	K_AS() antlr.TerminalNode

	// IsProjectionFieldContext differentiates from other interfaces.
	IsProjectionFieldContext()
}

type ProjectionFieldContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProjectionFieldContext() *ProjectionFieldContext {
	var p = new(ProjectionFieldContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CsqlParserRULE_projectionField
	return p
}

func InitEmptyProjectionFieldContext(p *ProjectionFieldContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CsqlParserRULE_projectionField
}

func (*ProjectionFieldContext) IsProjectionFieldContext() {}

func NewProjectionFieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProjectionFieldContext {
	var p = new(ProjectionFieldContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CsqlParserRULE_projectionField

	return p
}

func (s *ProjectionFieldContext) GetParser() antlr.Parser { return s.parser }

func (s *ProjectionFieldContext) ProjectionFieldName() IProjectionFieldNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IProjectionFieldNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IProjectionFieldNameContext)
}

func (s *ProjectionFieldContext) ValueExpr() IValueExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueExprContext)
}

func (s *ProjectionFieldContext) AggregateFunCall() IAggregateFunCallContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAggregateFunCallContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAggregateFunCallContext)
}

func (s *ProjectionFieldContext) Alias() IAliasContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAliasContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAliasContext)
}

func (s *ProjectionFieldContext) K_AS() antlr.TerminalNode {
	return s.GetToken(CsqlParserK_AS, 0)
}

func (s *ProjectionFieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProjectionFieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProjectionFieldContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CsqlVisitor:
		return t.VisitProjectionField(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CsqlParser) ProjectionField() (localctx IProjectionFieldContext) {
	localctx = NewProjectionFieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, CsqlParserRULE_projectionField)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(250)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 17, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(247)
			p.ProjectionFieldName()
		}

	case 2:
		{
			p.SetState(248)
			p.valueExpr(0)
		}

	case 3:
		{
			p.SetState(249)
			p.AggregateFunCall()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.SetState(256)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == CsqlParserK_AS || _la == CsqlParserIDENTIFIER {
		p.SetState(253)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == CsqlParserK_AS {
			{
				p.SetState(252)
				p.Match(CsqlParserK_AS)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(255)
			p.Alias()
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IProjectionFieldNameContext is an interface to support dynamic dispatch.
type IProjectionFieldNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FieldName() IFieldNameContext
	Qualifier() IQualifierContext

	// IsProjectionFieldNameContext differentiates from other interfaces.
	IsProjectionFieldNameContext()
}

type ProjectionFieldNameContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProjectionFieldNameContext() *ProjectionFieldNameContext {
	var p = new(ProjectionFieldNameContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CsqlParserRULE_projectionFieldName
	return p
}

func InitEmptyProjectionFieldNameContext(p *ProjectionFieldNameContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CsqlParserRULE_projectionFieldName
}

func (*ProjectionFieldNameContext) IsProjectionFieldNameContext() {}

func NewProjectionFieldNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProjectionFieldNameContext {
	var p = new(ProjectionFieldNameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CsqlParserRULE_projectionFieldName

	return p
}

func (s *ProjectionFieldNameContext) GetParser() antlr.Parser { return s.parser }

func (s *ProjectionFieldNameContext) FieldName() IFieldNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldNameContext)
}

func (s *ProjectionFieldNameContext) Qualifier() IQualifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IQualifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IQualifierContext)
}

func (s *ProjectionFieldNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProjectionFieldNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProjectionFieldNameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CsqlVisitor:
		return t.VisitProjectionFieldName(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CsqlParser) ProjectionFieldName() (localctx IProjectionFieldNameContext) {
	localctx = NewProjectionFieldNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, CsqlParserRULE_projectionFieldName)
	p.EnterOuterAlt(localctx, 1)
	p.SetState(261)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 20, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(258)
			p.Qualifier()
		}
		{
			p.SetState(259)
			p.Match(CsqlParserT__17)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}
	{
		p.SetState(263)
		p.FieldName()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFieldNameContext is an interface to support dynamic dispatch.
type IFieldNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode

	// IsFieldNameContext differentiates from other interfaces.
	IsFieldNameContext()
}

type FieldNameContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldNameContext() *FieldNameContext {
	var p = new(FieldNameContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CsqlParserRULE_fieldName
	return p
}

func InitEmptyFieldNameContext(p *FieldNameContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CsqlParserRULE_fieldName
}

func (*FieldNameContext) IsFieldNameContext() {}

func NewFieldNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldNameContext {
	var p = new(FieldNameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CsqlParserRULE_fieldName

	return p
}

func (s *FieldNameContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldNameContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(CsqlParserIDENTIFIER, 0)
}

func (s *FieldNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldNameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CsqlVisitor:
		return t.VisitFieldName(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CsqlParser) FieldName() (localctx IFieldNameContext) {
	localctx = NewFieldNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, CsqlParserRULE_fieldName)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(265)
		_la = p.GetTokenStream().LA(1)

		if !(_la == CsqlParserT__10 || _la == CsqlParserIDENTIFIER) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IInnerJoinContext is an interface to support dynamic dispatch.
type IInnerJoinContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	K_JOIN() antlr.TerminalNode
	K_INNER() antlr.TerminalNode

	// IsInnerJoinContext differentiates from other interfaces.
	IsInnerJoinContext()
}

type InnerJoinContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInnerJoinContext() *InnerJoinContext {
	var p = new(InnerJoinContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CsqlParserRULE_innerJoin
	return p
}

func InitEmptyInnerJoinContext(p *InnerJoinContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CsqlParserRULE_innerJoin
}

func (*InnerJoinContext) IsInnerJoinContext() {}

func NewInnerJoinContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InnerJoinContext {
	var p = new(InnerJoinContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CsqlParserRULE_innerJoin

	return p
}

func (s *InnerJoinContext) GetParser() antlr.Parser { return s.parser }

func (s *InnerJoinContext) K_JOIN() antlr.TerminalNode {
	return s.GetToken(CsqlParserK_JOIN, 0)
}

func (s *InnerJoinContext) K_INNER() antlr.TerminalNode {
	return s.GetToken(CsqlParserK_INNER, 0)
}

func (s *InnerJoinContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InnerJoinContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InnerJoinContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CsqlVisitor:
		return t.VisitInnerJoin(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CsqlParser) InnerJoin() (localctx IInnerJoinContext) {
	localctx = NewInnerJoinContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, CsqlParserRULE_innerJoin)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(268)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == CsqlParserK_INNER {
		{
			p.SetState(267)
			p.Match(CsqlParserK_INNER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(270)
		p.Match(CsqlParserK_JOIN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILeftJoinContext is an interface to support dynamic dispatch.
type ILeftJoinContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	K_LEFT() antlr.TerminalNode
	K_JOIN() antlr.TerminalNode
	K_OUTER() antlr.TerminalNode

	// IsLeftJoinContext differentiates from other interfaces.
	IsLeftJoinContext()
}

type LeftJoinContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLeftJoinContext() *LeftJoinContext {
	var p = new(LeftJoinContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CsqlParserRULE_leftJoin
	return p
}

func InitEmptyLeftJoinContext(p *LeftJoinContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CsqlParserRULE_leftJoin
}

func (*LeftJoinContext) IsLeftJoinContext() {}

func NewLeftJoinContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LeftJoinContext {
	var p = new(LeftJoinContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CsqlParserRULE_leftJoin

	return p
}

func (s *LeftJoinContext) GetParser() antlr.Parser { return s.parser }

func (s *LeftJoinContext) K_LEFT() antlr.TerminalNode {
	return s.GetToken(CsqlParserK_LEFT, 0)
}

func (s *LeftJoinContext) K_JOIN() antlr.TerminalNode {
	return s.GetToken(CsqlParserK_JOIN, 0)
}

func (s *LeftJoinContext) K_OUTER() antlr.TerminalNode {
	return s.GetToken(CsqlParserK_OUTER, 0)
}

func (s *LeftJoinContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LeftJoinContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LeftJoinContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CsqlVisitor:
		return t.VisitLeftJoin(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CsqlParser) LeftJoin() (localctx ILeftJoinContext) {
	localctx = NewLeftJoinContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, CsqlParserRULE_leftJoin)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(272)
		p.Match(CsqlParserK_LEFT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(274)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == CsqlParserK_OUTER {
		{
			p.SetState(273)
			p.Match(CsqlParserK_OUTER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(276)
		p.Match(CsqlParserK_JOIN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IRightJoinContext is an interface to support dynamic dispatch.
type IRightJoinContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	K_RIGHT() antlr.TerminalNode
	K_JOIN() antlr.TerminalNode
	K_OUTER() antlr.TerminalNode

	// IsRightJoinContext differentiates from other interfaces.
	IsRightJoinContext()
}

type RightJoinContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRightJoinContext() *RightJoinContext {
	var p = new(RightJoinContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CsqlParserRULE_rightJoin
	return p
}

func InitEmptyRightJoinContext(p *RightJoinContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CsqlParserRULE_rightJoin
}

func (*RightJoinContext) IsRightJoinContext() {}

func NewRightJoinContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RightJoinContext {
	var p = new(RightJoinContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CsqlParserRULE_rightJoin

	return p
}

func (s *RightJoinContext) GetParser() antlr.Parser { return s.parser }

func (s *RightJoinContext) K_RIGHT() antlr.TerminalNode {
	return s.GetToken(CsqlParserK_RIGHT, 0)
}

func (s *RightJoinContext) K_JOIN() antlr.TerminalNode {
	return s.GetToken(CsqlParserK_JOIN, 0)
}

func (s *RightJoinContext) K_OUTER() antlr.TerminalNode {
	return s.GetToken(CsqlParserK_OUTER, 0)
}

func (s *RightJoinContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RightJoinContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RightJoinContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CsqlVisitor:
		return t.VisitRightJoin(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CsqlParser) RightJoin() (localctx IRightJoinContext) {
	localctx = NewRightJoinContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, CsqlParserRULE_rightJoin)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(278)
		p.Match(CsqlParserK_RIGHT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(280)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == CsqlParserK_OUTER {
		{
			p.SetState(279)
			p.Match(CsqlParserK_OUTER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(282)
		p.Match(CsqlParserK_JOIN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFullJoinContext is an interface to support dynamic dispatch.
type IFullJoinContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	K_FULL() antlr.TerminalNode
	K_JOIN() antlr.TerminalNode
	K_OUTER() antlr.TerminalNode

	// IsFullJoinContext differentiates from other interfaces.
	IsFullJoinContext()
}

type FullJoinContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFullJoinContext() *FullJoinContext {
	var p = new(FullJoinContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CsqlParserRULE_fullJoin
	return p
}

func InitEmptyFullJoinContext(p *FullJoinContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CsqlParserRULE_fullJoin
}

func (*FullJoinContext) IsFullJoinContext() {}

func NewFullJoinContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FullJoinContext {
	var p = new(FullJoinContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CsqlParserRULE_fullJoin

	return p
}

func (s *FullJoinContext) GetParser() antlr.Parser { return s.parser }

func (s *FullJoinContext) K_FULL() antlr.TerminalNode {
	return s.GetToken(CsqlParserK_FULL, 0)
}

func (s *FullJoinContext) K_JOIN() antlr.TerminalNode {
	return s.GetToken(CsqlParserK_JOIN, 0)
}

func (s *FullJoinContext) K_OUTER() antlr.TerminalNode {
	return s.GetToken(CsqlParserK_OUTER, 0)
}

func (s *FullJoinContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FullJoinContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FullJoinContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CsqlVisitor:
		return t.VisitFullJoin(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CsqlParser) FullJoin() (localctx IFullJoinContext) {
	localctx = NewFullJoinContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, CsqlParserRULE_fullJoin)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(284)
		p.Match(CsqlParserK_FULL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(286)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == CsqlParserK_OUTER {
		{
			p.SetState(285)
			p.Match(CsqlParserK_OUTER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(288)
		p.Match(CsqlParserK_JOIN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICrossJoinContext is an interface to support dynamic dispatch.
type ICrossJoinContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	K_CROSS() antlr.TerminalNode
	K_JOIN() antlr.TerminalNode

	// IsCrossJoinContext differentiates from other interfaces.
	IsCrossJoinContext()
}

type CrossJoinContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCrossJoinContext() *CrossJoinContext {
	var p = new(CrossJoinContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CsqlParserRULE_crossJoin
	return p
}

func InitEmptyCrossJoinContext(p *CrossJoinContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CsqlParserRULE_crossJoin
}

func (*CrossJoinContext) IsCrossJoinContext() {}

func NewCrossJoinContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CrossJoinContext {
	var p = new(CrossJoinContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CsqlParserRULE_crossJoin

	return p
}

func (s *CrossJoinContext) GetParser() antlr.Parser { return s.parser }

func (s *CrossJoinContext) K_CROSS() antlr.TerminalNode {
	return s.GetToken(CsqlParserK_CROSS, 0)
}

func (s *CrossJoinContext) K_JOIN() antlr.TerminalNode {
	return s.GetToken(CsqlParserK_JOIN, 0)
}

func (s *CrossJoinContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CrossJoinContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CrossJoinContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CsqlVisitor:
		return t.VisitCrossJoin(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CsqlParser) CrossJoin() (localctx ICrossJoinContext) {
	localctx = NewCrossJoinContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, CsqlParserRULE_crossJoin)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(290)
		p.Match(CsqlParserK_CROSS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(291)
		p.Match(CsqlParserK_JOIN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IConditionalJoinTypeContext is an interface to support dynamic dispatch.
type IConditionalJoinTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	InnerJoin() IInnerJoinContext
	LeftJoin() ILeftJoinContext
	RightJoin() IRightJoinContext
	FullJoin() IFullJoinContext

	// IsConditionalJoinTypeContext differentiates from other interfaces.
	IsConditionalJoinTypeContext()
}

type ConditionalJoinTypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConditionalJoinTypeContext() *ConditionalJoinTypeContext {
	var p = new(ConditionalJoinTypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CsqlParserRULE_conditionalJoinType
	return p
}

func InitEmptyConditionalJoinTypeContext(p *ConditionalJoinTypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CsqlParserRULE_conditionalJoinType
}

func (*ConditionalJoinTypeContext) IsConditionalJoinTypeContext() {}

func NewConditionalJoinTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConditionalJoinTypeContext {
	var p = new(ConditionalJoinTypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CsqlParserRULE_conditionalJoinType

	return p
}

func (s *ConditionalJoinTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *ConditionalJoinTypeContext) InnerJoin() IInnerJoinContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInnerJoinContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInnerJoinContext)
}

func (s *ConditionalJoinTypeContext) LeftJoin() ILeftJoinContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILeftJoinContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILeftJoinContext)
}

func (s *ConditionalJoinTypeContext) RightJoin() IRightJoinContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRightJoinContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRightJoinContext)
}

func (s *ConditionalJoinTypeContext) FullJoin() IFullJoinContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFullJoinContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFullJoinContext)
}

func (s *ConditionalJoinTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConditionalJoinTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConditionalJoinTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CsqlVisitor:
		return t.VisitConditionalJoinType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CsqlParser) ConditionalJoinType() (localctx IConditionalJoinTypeContext) {
	localctx = NewConditionalJoinTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, CsqlParserRULE_conditionalJoinType)
	p.SetState(297)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case CsqlParserK_JOIN, CsqlParserK_INNER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(293)
			p.InnerJoin()
		}

	case CsqlParserK_LEFT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(294)
			p.LeftJoin()
		}

	case CsqlParserK_RIGHT:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(295)
			p.RightJoin()
		}

	case CsqlParserK_FULL:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(296)
			p.FullJoin()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IConditionalJoinTargetContext is an interface to support dynamic dispatch.
type IConditionalJoinTargetContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DataSource() IDataSourceContext
	K_ON() antlr.TerminalNode
	WhereExpr() IWhereExprContext

	// IsConditionalJoinTargetContext differentiates from other interfaces.
	IsConditionalJoinTargetContext()
}

type ConditionalJoinTargetContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConditionalJoinTargetContext() *ConditionalJoinTargetContext {
	var p = new(ConditionalJoinTargetContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CsqlParserRULE_conditionalJoinTarget
	return p
}

func InitEmptyConditionalJoinTargetContext(p *ConditionalJoinTargetContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CsqlParserRULE_conditionalJoinTarget
}

func (*ConditionalJoinTargetContext) IsConditionalJoinTargetContext() {}

func NewConditionalJoinTargetContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConditionalJoinTargetContext {
	var p = new(ConditionalJoinTargetContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CsqlParserRULE_conditionalJoinTarget

	return p
}

func (s *ConditionalJoinTargetContext) GetParser() antlr.Parser { return s.parser }

func (s *ConditionalJoinTargetContext) DataSource() IDataSourceContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDataSourceContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDataSourceContext)
}

func (s *ConditionalJoinTargetContext) K_ON() antlr.TerminalNode {
	return s.GetToken(CsqlParserK_ON, 0)
}

func (s *ConditionalJoinTargetContext) WhereExpr() IWhereExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWhereExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWhereExprContext)
}

func (s *ConditionalJoinTargetContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConditionalJoinTargetContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConditionalJoinTargetContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CsqlVisitor:
		return t.VisitConditionalJoinTarget(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CsqlParser) ConditionalJoinTarget() (localctx IConditionalJoinTargetContext) {
	localctx = NewConditionalJoinTargetContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, CsqlParserRULE_conditionalJoinTarget)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(299)
		p.dataSource(0)
	}
	{
		p.SetState(300)
		p.Match(CsqlParserK_ON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(301)
		p.whereExpr(0)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IUnconditionalJoinTargetContext is an interface to support dynamic dispatch.
type IUnconditionalJoinTargetContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DataSource() IDataSourceContext

	// IsUnconditionalJoinTargetContext differentiates from other interfaces.
	IsUnconditionalJoinTargetContext()
}

type UnconditionalJoinTargetContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnconditionalJoinTargetContext() *UnconditionalJoinTargetContext {
	var p = new(UnconditionalJoinTargetContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CsqlParserRULE_unconditionalJoinTarget
	return p
}

func InitEmptyUnconditionalJoinTargetContext(p *UnconditionalJoinTargetContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CsqlParserRULE_unconditionalJoinTarget
}

func (*UnconditionalJoinTargetContext) IsUnconditionalJoinTargetContext() {}

func NewUnconditionalJoinTargetContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UnconditionalJoinTargetContext {
	var p = new(UnconditionalJoinTargetContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CsqlParserRULE_unconditionalJoinTarget

	return p
}

func (s *UnconditionalJoinTargetContext) GetParser() antlr.Parser { return s.parser }

func (s *UnconditionalJoinTargetContext) DataSource() IDataSourceContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDataSourceContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDataSourceContext)
}

func (s *UnconditionalJoinTargetContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnconditionalJoinTargetContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UnconditionalJoinTargetContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CsqlVisitor:
		return t.VisitUnconditionalJoinTarget(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CsqlParser) UnconditionalJoinTarget() (localctx IUnconditionalJoinTargetContext) {
	localctx = NewUnconditionalJoinTargetContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, CsqlParserRULE_unconditionalJoinTarget)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(303)
		p.dataSource(0)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDataSourceContext is an interface to support dynamic dispatch.
type IDataSourceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsDataSourceContext differentiates from other interfaces.
	IsDataSourceContext()
}

type DataSourceContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDataSourceContext() *DataSourceContext {
	var p = new(DataSourceContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CsqlParserRULE_dataSource
	return p
}

func InitEmptyDataSourceContext(p *DataSourceContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CsqlParserRULE_dataSource
}

func (*DataSourceContext) IsDataSourceContext() {}

func NewDataSourceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DataSourceContext {
	var p = new(DataSourceContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CsqlParserRULE_dataSource

	return p
}

func (s *DataSourceContext) GetParser() antlr.Parser { return s.parser }

func (s *DataSourceContext) CopyAll(ctx *DataSourceContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *DataSourceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DataSourceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type DataSourceConditionalJoinContext struct {
	DataSourceContext
}

func NewDataSourceConditionalJoinContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DataSourceConditionalJoinContext {
	var p = new(DataSourceConditionalJoinContext)

	InitEmptyDataSourceContext(&p.DataSourceContext)
	p.parser = parser
	p.CopyAll(ctx.(*DataSourceContext))

	return p
}

func (s *DataSourceConditionalJoinContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DataSourceConditionalJoinContext) DataSource() IDataSourceContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDataSourceContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDataSourceContext)
}

func (s *DataSourceConditionalJoinContext) ConditionalJoinType() IConditionalJoinTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConditionalJoinTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConditionalJoinTypeContext)
}

func (s *DataSourceConditionalJoinContext) ConditionalJoinTarget() IConditionalJoinTargetContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConditionalJoinTargetContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConditionalJoinTargetContext)
}

func (s *DataSourceConditionalJoinContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CsqlVisitor:
		return t.VisitDataSourceConditionalJoin(s)

	default:
		return t.VisitChildren(s)
	}
}

type DataSourceItemContext struct {
	DataSourceContext
}

func NewDataSourceItemContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DataSourceItemContext {
	var p = new(DataSourceItemContext)

	InitEmptyDataSourceContext(&p.DataSourceContext)
	p.parser = parser
	p.CopyAll(ctx.(*DataSourceContext))

	return p
}

func (s *DataSourceItemContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DataSourceItemContext) DataSource() IDataSourceContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDataSourceContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDataSourceContext)
}

func (s *DataSourceItemContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CsqlVisitor:
		return t.VisitDataSourceItem(s)

	default:
		return t.VisitChildren(s)
	}
}

type DataSourceCrossJoinContext struct {
	DataSourceContext
}

func NewDataSourceCrossJoinContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DataSourceCrossJoinContext {
	var p = new(DataSourceCrossJoinContext)

	InitEmptyDataSourceContext(&p.DataSourceContext)
	p.parser = parser
	p.CopyAll(ctx.(*DataSourceContext))

	return p
}

func (s *DataSourceCrossJoinContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DataSourceCrossJoinContext) DataSource() IDataSourceContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDataSourceContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDataSourceContext)
}

func (s *DataSourceCrossJoinContext) CrossJoin() ICrossJoinContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICrossJoinContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICrossJoinContext)
}

func (s *DataSourceCrossJoinContext) UnconditionalJoinTarget() IUnconditionalJoinTargetContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUnconditionalJoinTargetContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUnconditionalJoinTargetContext)
}

func (s *DataSourceCrossJoinContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CsqlVisitor:
		return t.VisitDataSourceCrossJoin(s)

	default:
		return t.VisitChildren(s)
	}
}

type DataSourceNamedContext struct {
	DataSourceContext
}

func NewDataSourceNamedContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DataSourceNamedContext {
	var p = new(DataSourceNamedContext)

	InitEmptyDataSourceContext(&p.DataSourceContext)
	p.parser = parser
	p.CopyAll(ctx.(*DataSourceContext))

	return p
}

func (s *DataSourceNamedContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DataSourceNamedContext) SourceName() ISourceNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISourceNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISourceNameContext)
}

func (s *DataSourceNamedContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CsqlVisitor:
		return t.VisitDataSourceNamed(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CsqlParser) DataSource() (localctx IDataSourceContext) {
	return p.dataSource(0)
}

func (p *CsqlParser) dataSource(_p int) (localctx IDataSourceContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewDataSourceContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IDataSourceContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 52
	p.EnterRecursionRule(localctx, 52, CsqlParserRULE_dataSource, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(311)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case CsqlParserIDENTIFIER:
		localctx = NewDataSourceNamedContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(306)
			p.SourceName()
		}

	case CsqlParserT__14:
		localctx = NewDataSourceItemContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(307)
			p.Match(CsqlParserT__14)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(308)
			p.dataSource(0)
		}
		{
			p.SetState(309)
			p.Match(CsqlParserT__16)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(323)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 28, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(321)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 27, p.GetParserRuleContext()) {
			case 1:
				localctx = NewDataSourceConditionalJoinContext(p, NewDataSourceContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, CsqlParserRULE_dataSource)
				p.SetState(313)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
					goto errorExit
				}
				{
					p.SetState(314)
					p.ConditionalJoinType()
				}
				{
					p.SetState(315)
					p.ConditionalJoinTarget()
				}

			case 2:
				localctx = NewDataSourceCrossJoinContext(p, NewDataSourceContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, CsqlParserRULE_dataSource)
				p.SetState(317)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
					goto errorExit
				}
				{
					p.SetState(318)
					p.CrossJoin()
				}
				{
					p.SetState(319)
					p.UnconditionalJoinTarget()
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(325)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 28, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISourcesContext is an interface to support dynamic dispatch.
type ISourcesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllDataSource() []IDataSourceContext
	DataSource(i int) IDataSourceContext

	// IsSourcesContext differentiates from other interfaces.
	IsSourcesContext()
}

type SourcesContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySourcesContext() *SourcesContext {
	var p = new(SourcesContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CsqlParserRULE_sources
	return p
}

func InitEmptySourcesContext(p *SourcesContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CsqlParserRULE_sources
}

func (*SourcesContext) IsSourcesContext() {}

func NewSourcesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SourcesContext {
	var p = new(SourcesContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CsqlParserRULE_sources

	return p
}

func (s *SourcesContext) GetParser() antlr.Parser { return s.parser }

func (s *SourcesContext) AllDataSource() []IDataSourceContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IDataSourceContext); ok {
			len++
		}
	}

	tst := make([]IDataSourceContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IDataSourceContext); ok {
			tst[i] = t.(IDataSourceContext)
			i++
		}
	}

	return tst
}

func (s *SourcesContext) DataSource(i int) IDataSourceContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDataSourceContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDataSourceContext)
}

func (s *SourcesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SourcesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SourcesContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CsqlVisitor:
		return t.VisitSources(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CsqlParser) Sources() (localctx ISourcesContext) {
	localctx = NewSourcesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, CsqlParserRULE_sources)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(326)
		p.dataSource(0)
	}
	p.SetState(331)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == CsqlParserT__15 {
		{
			p.SetState(327)
			p.Match(CsqlParserT__15)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(328)
			p.dataSource(0)
		}

		p.SetState(333)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IUnionSelectsContext is an interface to support dynamic dispatch.
type IUnionSelectsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SelectStatement() ISelectStatementContext
	AllK_UNION() []antlr.TerminalNode
	K_UNION(i int) antlr.TerminalNode
	AllUnionSelects() []IUnionSelectsContext
	UnionSelects(i int) IUnionSelectsContext
	AllK_ALL() []antlr.TerminalNode
	K_ALL(i int) antlr.TerminalNode

	// IsUnionSelectsContext differentiates from other interfaces.
	IsUnionSelectsContext()
}

type UnionSelectsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnionSelectsContext() *UnionSelectsContext {
	var p = new(UnionSelectsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CsqlParserRULE_unionSelects
	return p
}

func InitEmptyUnionSelectsContext(p *UnionSelectsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CsqlParserRULE_unionSelects
}

func (*UnionSelectsContext) IsUnionSelectsContext() {}

func NewUnionSelectsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UnionSelectsContext {
	var p = new(UnionSelectsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CsqlParserRULE_unionSelects

	return p
}

func (s *UnionSelectsContext) GetParser() antlr.Parser { return s.parser }

func (s *UnionSelectsContext) SelectStatement() ISelectStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISelectStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISelectStatementContext)
}

func (s *UnionSelectsContext) AllK_UNION() []antlr.TerminalNode {
	return s.GetTokens(CsqlParserK_UNION)
}

func (s *UnionSelectsContext) K_UNION(i int) antlr.TerminalNode {
	return s.GetToken(CsqlParserK_UNION, i)
}

func (s *UnionSelectsContext) AllUnionSelects() []IUnionSelectsContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IUnionSelectsContext); ok {
			len++
		}
	}

	tst := make([]IUnionSelectsContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IUnionSelectsContext); ok {
			tst[i] = t.(IUnionSelectsContext)
			i++
		}
	}

	return tst
}

func (s *UnionSelectsContext) UnionSelects(i int) IUnionSelectsContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUnionSelectsContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUnionSelectsContext)
}

func (s *UnionSelectsContext) AllK_ALL() []antlr.TerminalNode {
	return s.GetTokens(CsqlParserK_ALL)
}

func (s *UnionSelectsContext) K_ALL(i int) antlr.TerminalNode {
	return s.GetToken(CsqlParserK_ALL, i)
}

func (s *UnionSelectsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnionSelectsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UnionSelectsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CsqlVisitor:
		return t.VisitUnionSelects(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CsqlParser) UnionSelects() (localctx IUnionSelectsContext) {
	localctx = NewUnionSelectsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, CsqlParserRULE_unionSelects)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(334)
		p.SelectStatement()
	}
	p.SetState(342)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 31, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(335)
				p.Match(CsqlParserK_UNION)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			p.SetState(337)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			if _la == CsqlParserK_ALL {
				{
					p.SetState(336)
					p.Match(CsqlParserK_ALL)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			}
			{
				p.SetState(339)
				p.UnionSelects()
			}

		}
		p.SetState(344)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 31, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISelectStatementContext is an interface to support dynamic dispatch.
type ISelectStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	K_SELECT() antlr.TerminalNode
	Projection() IProjectionContext
	K_FROM() antlr.TerminalNode
	Sources() ISourcesContext
	Where() IWhereContext
	OrderBy() IOrderByContext
	Limit() ILimitContext
	GroupBy() IGroupByContext
	Having() IHavingContext

	// IsSelectStatementContext differentiates from other interfaces.
	IsSelectStatementContext()
}

type SelectStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySelectStatementContext() *SelectStatementContext {
	var p = new(SelectStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CsqlParserRULE_selectStatement
	return p
}

func InitEmptySelectStatementContext(p *SelectStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CsqlParserRULE_selectStatement
}

func (*SelectStatementContext) IsSelectStatementContext() {}

func NewSelectStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SelectStatementContext {
	var p = new(SelectStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CsqlParserRULE_selectStatement

	return p
}

func (s *SelectStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *SelectStatementContext) K_SELECT() antlr.TerminalNode {
	return s.GetToken(CsqlParserK_SELECT, 0)
}

func (s *SelectStatementContext) Projection() IProjectionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IProjectionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IProjectionContext)
}

func (s *SelectStatementContext) K_FROM() antlr.TerminalNode {
	return s.GetToken(CsqlParserK_FROM, 0)
}

func (s *SelectStatementContext) Sources() ISourcesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISourcesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISourcesContext)
}

func (s *SelectStatementContext) Where() IWhereContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWhereContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWhereContext)
}

func (s *SelectStatementContext) OrderBy() IOrderByContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOrderByContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOrderByContext)
}

func (s *SelectStatementContext) Limit() ILimitContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILimitContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILimitContext)
}

func (s *SelectStatementContext) GroupBy() IGroupByContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGroupByContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGroupByContext)
}

func (s *SelectStatementContext) Having() IHavingContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IHavingContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IHavingContext)
}

func (s *SelectStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SelectStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SelectStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CsqlVisitor:
		return t.VisitSelectStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CsqlParser) SelectStatement() (localctx ISelectStatementContext) {
	localctx = NewSelectStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, CsqlParserRULE_selectStatement)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(345)
		p.Match(CsqlParserK_SELECT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(346)
		p.Projection()
	}
	{
		p.SetState(347)
		p.Match(CsqlParserK_FROM)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(348)
		p.Sources()
	}
	p.SetState(350)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == CsqlParserK_WHERE {
		{
			p.SetState(349)
			p.Where()
		}

	}
	p.SetState(353)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == CsqlParserK_ORDER {
		{
			p.SetState(352)
			p.OrderBy()
		}

	}
	p.SetState(356)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == CsqlParserK_LIMIT {
		{
			p.SetState(355)
			p.Limit()
		}

	}
	p.SetState(359)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == CsqlParserK_GROUP {
		{
			p.SetState(358)
			p.GroupBy()
		}

	}
	p.SetState(362)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == CsqlParserK_HAVING {
		{
			p.SetState(361)
			p.Having()
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISignedNumberContext is an interface to support dynamic dispatch.
type ISignedNumberContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NUMERIC_LITERAL() antlr.TerminalNode

	// IsSignedNumberContext differentiates from other interfaces.
	IsSignedNumberContext()
}

type SignedNumberContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySignedNumberContext() *SignedNumberContext {
	var p = new(SignedNumberContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CsqlParserRULE_signedNumber
	return p
}

func InitEmptySignedNumberContext(p *SignedNumberContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CsqlParserRULE_signedNumber
}

func (*SignedNumberContext) IsSignedNumberContext() {}

func NewSignedNumberContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SignedNumberContext {
	var p = new(SignedNumberContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CsqlParserRULE_signedNumber

	return p
}

func (s *SignedNumberContext) GetParser() antlr.Parser { return s.parser }

func (s *SignedNumberContext) NUMERIC_LITERAL() antlr.TerminalNode {
	return s.GetToken(CsqlParserNUMERIC_LITERAL, 0)
}

func (s *SignedNumberContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SignedNumberContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SignedNumberContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CsqlVisitor:
		return t.VisitSignedNumber(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CsqlParser) SignedNumber() (localctx ISignedNumberContext) {
	localctx = NewSignedNumberContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, CsqlParserRULE_signedNumber)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(365)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == CsqlParserT__8 || _la == CsqlParserT__9 {
		{
			p.SetState(364)
			_la = p.GetTokenStream().LA(1)

			if !(_la == CsqlParserT__8 || _la == CsqlParserT__9) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	}
	{
		p.SetState(367)
		p.Match(CsqlParserNUMERIC_LITERAL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStringValueContext is an interface to support dynamic dispatch.
type IStringValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	STRING_LITERAL() antlr.TerminalNode

	// IsStringValueContext differentiates from other interfaces.
	IsStringValueContext()
}

type StringValueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStringValueContext() *StringValueContext {
	var p = new(StringValueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CsqlParserRULE_stringValue
	return p
}

func InitEmptyStringValueContext(p *StringValueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CsqlParserRULE_stringValue
}

func (*StringValueContext) IsStringValueContext() {}

func NewStringValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StringValueContext {
	var p = new(StringValueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CsqlParserRULE_stringValue

	return p
}

func (s *StringValueContext) GetParser() antlr.Parser { return s.parser }

func (s *StringValueContext) STRING_LITERAL() antlr.TerminalNode {
	return s.GetToken(CsqlParserSTRING_LITERAL, 0)
}

func (s *StringValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StringValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CsqlVisitor:
		return t.VisitStringValue(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CsqlParser) StringValue() (localctx IStringValueContext) {
	localctx = NewStringValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, CsqlParserRULE_stringValue)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(369)
		p.Match(CsqlParserSTRING_LITERAL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// INullValueContext is an interface to support dynamic dispatch.
type INullValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	K_NULL() antlr.TerminalNode

	// IsNullValueContext differentiates from other interfaces.
	IsNullValueContext()
}

type NullValueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNullValueContext() *NullValueContext {
	var p = new(NullValueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CsqlParserRULE_nullValue
	return p
}

func InitEmptyNullValueContext(p *NullValueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CsqlParserRULE_nullValue
}

func (*NullValueContext) IsNullValueContext() {}

func NewNullValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NullValueContext {
	var p = new(NullValueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CsqlParserRULE_nullValue

	return p
}

func (s *NullValueContext) GetParser() antlr.Parser { return s.parser }

func (s *NullValueContext) K_NULL() antlr.TerminalNode {
	return s.GetToken(CsqlParserK_NULL, 0)
}

func (s *NullValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NullValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NullValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CsqlVisitor:
		return t.VisitNullValue(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CsqlParser) NullValue() (localctx INullValueContext) {
	localctx = NewNullValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, CsqlParserRULE_nullValue)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(371)
		p.Match(CsqlParserK_NULL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILiteralValueContext is an interface to support dynamic dispatch.
type ILiteralValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SignedNumber() ISignedNumberContext
	StringValue() IStringValueContext
	NullValue() INullValueContext

	// IsLiteralValueContext differentiates from other interfaces.
	IsLiteralValueContext()
}

type LiteralValueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLiteralValueContext() *LiteralValueContext {
	var p = new(LiteralValueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CsqlParserRULE_literalValue
	return p
}

func InitEmptyLiteralValueContext(p *LiteralValueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CsqlParserRULE_literalValue
}

func (*LiteralValueContext) IsLiteralValueContext() {}

func NewLiteralValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LiteralValueContext {
	var p = new(LiteralValueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CsqlParserRULE_literalValue

	return p
}

func (s *LiteralValueContext) GetParser() antlr.Parser { return s.parser }

func (s *LiteralValueContext) SignedNumber() ISignedNumberContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISignedNumberContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISignedNumberContext)
}

func (s *LiteralValueContext) StringValue() IStringValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStringValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStringValueContext)
}

func (s *LiteralValueContext) NullValue() INullValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INullValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INullValueContext)
}

func (s *LiteralValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LiteralValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CsqlVisitor:
		return t.VisitLiteralValue(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CsqlParser) LiteralValue() (localctx ILiteralValueContext) {
	localctx = NewLiteralValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, CsqlParserRULE_literalValue)
	p.SetState(376)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case CsqlParserT__8, CsqlParserT__9, CsqlParserNUMERIC_LITERAL:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(373)
			p.SignedNumber()
		}

	case CsqlParserSTRING_LITERAL:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(374)
			p.StringValue()
		}

	case CsqlParserK_NULL:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(375)
			p.NullValue()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAliasContext is an interface to support dynamic dispatch.
type IAliasContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode

	// IsAliasContext differentiates from other interfaces.
	IsAliasContext()
}

type AliasContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAliasContext() *AliasContext {
	var p = new(AliasContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CsqlParserRULE_alias
	return p
}

func InitEmptyAliasContext(p *AliasContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CsqlParserRULE_alias
}

func (*AliasContext) IsAliasContext() {}

func NewAliasContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AliasContext {
	var p = new(AliasContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CsqlParserRULE_alias

	return p
}

func (s *AliasContext) GetParser() antlr.Parser { return s.parser }

func (s *AliasContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(CsqlParserIDENTIFIER, 0)
}

func (s *AliasContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AliasContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AliasContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CsqlVisitor:
		return t.VisitAlias(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CsqlParser) Alias() (localctx IAliasContext) {
	localctx = NewAliasContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, CsqlParserRULE_alias)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(378)
		p.Match(CsqlParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISourceNameContext is an interface to support dynamic dispatch.
type ISourceNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Name() INameContext
	Alias() IAliasContext

	// IsSourceNameContext differentiates from other interfaces.
	IsSourceNameContext()
}

type SourceNameContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySourceNameContext() *SourceNameContext {
	var p = new(SourceNameContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CsqlParserRULE_sourceName
	return p
}

func InitEmptySourceNameContext(p *SourceNameContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CsqlParserRULE_sourceName
}

func (*SourceNameContext) IsSourceNameContext() {}

func NewSourceNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SourceNameContext {
	var p = new(SourceNameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CsqlParserRULE_sourceName

	return p
}

func (s *SourceNameContext) GetParser() antlr.Parser { return s.parser }

func (s *SourceNameContext) Name() INameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *SourceNameContext) Alias() IAliasContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAliasContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAliasContext)
}

func (s *SourceNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SourceNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SourceNameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CsqlVisitor:
		return t.VisitSourceName(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CsqlParser) SourceName() (localctx ISourceNameContext) {
	localctx = NewSourceNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, CsqlParserRULE_sourceName)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(380)
		p.Name()
	}
	p.SetState(382)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 39, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(381)
			p.Alias()
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICompoundNameContext is an interface to support dynamic dispatch.
type ICompoundNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Name() INameContext
	Qualifier() IQualifierContext

	// IsCompoundNameContext differentiates from other interfaces.
	IsCompoundNameContext()
}

type CompoundNameContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCompoundNameContext() *CompoundNameContext {
	var p = new(CompoundNameContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CsqlParserRULE_compoundName
	return p
}

func InitEmptyCompoundNameContext(p *CompoundNameContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CsqlParserRULE_compoundName
}

func (*CompoundNameContext) IsCompoundNameContext() {}

func NewCompoundNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CompoundNameContext {
	var p = new(CompoundNameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CsqlParserRULE_compoundName

	return p
}

func (s *CompoundNameContext) GetParser() antlr.Parser { return s.parser }

func (s *CompoundNameContext) Name() INameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *CompoundNameContext) Qualifier() IQualifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IQualifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IQualifierContext)
}

func (s *CompoundNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CompoundNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CompoundNameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CsqlVisitor:
		return t.VisitCompoundName(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CsqlParser) CompoundName() (localctx ICompoundNameContext) {
	localctx = NewCompoundNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, CsqlParserRULE_compoundName)
	p.EnterOuterAlt(localctx, 1)
	p.SetState(387)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 40, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(384)
			p.Qualifier()
		}
		{
			p.SetState(385)
			p.Match(CsqlParserT__17)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}
	{
		p.SetState(389)
		p.Name()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// INameContext is an interface to support dynamic dispatch.
type INameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode

	// IsNameContext differentiates from other interfaces.
	IsNameContext()
}

type NameContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNameContext() *NameContext {
	var p = new(NameContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CsqlParserRULE_name
	return p
}

func InitEmptyNameContext(p *NameContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CsqlParserRULE_name
}

func (*NameContext) IsNameContext() {}

func NewNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NameContext {
	var p = new(NameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CsqlParserRULE_name

	return p
}

func (s *NameContext) GetParser() antlr.Parser { return s.parser }

func (s *NameContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(CsqlParserIDENTIFIER, 0)
}

func (s *NameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CsqlVisitor:
		return t.VisitName(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CsqlParser) Name() (localctx INameContext) {
	localctx = NewNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, CsqlParserRULE_name)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(391)
		p.Match(CsqlParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IQualifierContext is an interface to support dynamic dispatch.
type IQualifierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode

	// IsQualifierContext differentiates from other interfaces.
	IsQualifierContext()
}

type QualifierContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyQualifierContext() *QualifierContext {
	var p = new(QualifierContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CsqlParserRULE_qualifier
	return p
}

func InitEmptyQualifierContext(p *QualifierContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CsqlParserRULE_qualifier
}

func (*QualifierContext) IsQualifierContext() {}

func NewQualifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *QualifierContext {
	var p = new(QualifierContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CsqlParserRULE_qualifier

	return p
}

func (s *QualifierContext) GetParser() antlr.Parser { return s.parser }

func (s *QualifierContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(CsqlParserIDENTIFIER, 0)
}

func (s *QualifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *QualifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *QualifierContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CsqlVisitor:
		return t.VisitQualifier(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CsqlParser) Qualifier() (localctx IQualifierContext) {
	localctx = NewQualifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 76, CsqlParserRULE_qualifier)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(393)
		p.Match(CsqlParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILimitContext is an interface to support dynamic dispatch.
type ILimitContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	K_LIMIT() antlr.TerminalNode
	LimitValue() ILimitValueContext

	// IsLimitContext differentiates from other interfaces.
	IsLimitContext()
}

type LimitContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLimitContext() *LimitContext {
	var p = new(LimitContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CsqlParserRULE_limit
	return p
}

func InitEmptyLimitContext(p *LimitContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CsqlParserRULE_limit
}

func (*LimitContext) IsLimitContext() {}

func NewLimitContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LimitContext {
	var p = new(LimitContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CsqlParserRULE_limit

	return p
}

func (s *LimitContext) GetParser() antlr.Parser { return s.parser }

func (s *LimitContext) K_LIMIT() antlr.TerminalNode {
	return s.GetToken(CsqlParserK_LIMIT, 0)
}

func (s *LimitContext) LimitValue() ILimitValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILimitValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILimitValueContext)
}

func (s *LimitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LimitContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LimitContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CsqlVisitor:
		return t.VisitLimit(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CsqlParser) Limit() (localctx ILimitContext) {
	localctx = NewLimitContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 78, CsqlParserRULE_limit)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(395)
		p.Match(CsqlParserK_LIMIT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(396)
		p.LimitValue()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILimitValueContext is an interface to support dynamic dispatch.
type ILimitValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NUMERIC_LITERAL() antlr.TerminalNode

	// IsLimitValueContext differentiates from other interfaces.
	IsLimitValueContext()
}

type LimitValueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLimitValueContext() *LimitValueContext {
	var p = new(LimitValueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CsqlParserRULE_limitValue
	return p
}

func InitEmptyLimitValueContext(p *LimitValueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CsqlParserRULE_limitValue
}

func (*LimitValueContext) IsLimitValueContext() {}

func NewLimitValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LimitValueContext {
	var p = new(LimitValueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CsqlParserRULE_limitValue

	return p
}

func (s *LimitValueContext) GetParser() antlr.Parser { return s.parser }

func (s *LimitValueContext) NUMERIC_LITERAL() antlr.TerminalNode {
	return s.GetToken(CsqlParserNUMERIC_LITERAL, 0)
}

func (s *LimitValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LimitValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LimitValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CsqlVisitor:
		return t.VisitLimitValue(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CsqlParser) LimitValue() (localctx ILimitValueContext) {
	localctx = NewLimitValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 80, CsqlParserRULE_limitValue)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(398)
		p.Match(CsqlParserNUMERIC_LITERAL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IOrderByContext is an interface to support dynamic dispatch.
type IOrderByContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	K_ORDER() antlr.TerminalNode
	K_BY() antlr.TerminalNode
	AllOrderByField() []IOrderByFieldContext
	OrderByField(i int) IOrderByFieldContext

	// IsOrderByContext differentiates from other interfaces.
	IsOrderByContext()
}

type OrderByContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOrderByContext() *OrderByContext {
	var p = new(OrderByContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CsqlParserRULE_orderBy
	return p
}

func InitEmptyOrderByContext(p *OrderByContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CsqlParserRULE_orderBy
}

func (*OrderByContext) IsOrderByContext() {}

func NewOrderByContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OrderByContext {
	var p = new(OrderByContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CsqlParserRULE_orderBy

	return p
}

func (s *OrderByContext) GetParser() antlr.Parser { return s.parser }

func (s *OrderByContext) K_ORDER() antlr.TerminalNode {
	return s.GetToken(CsqlParserK_ORDER, 0)
}

func (s *OrderByContext) K_BY() antlr.TerminalNode {
	return s.GetToken(CsqlParserK_BY, 0)
}

func (s *OrderByContext) AllOrderByField() []IOrderByFieldContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IOrderByFieldContext); ok {
			len++
		}
	}

	tst := make([]IOrderByFieldContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IOrderByFieldContext); ok {
			tst[i] = t.(IOrderByFieldContext)
			i++
		}
	}

	return tst
}

func (s *OrderByContext) OrderByField(i int) IOrderByFieldContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOrderByFieldContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOrderByFieldContext)
}

func (s *OrderByContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OrderByContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OrderByContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CsqlVisitor:
		return t.VisitOrderBy(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CsqlParser) OrderBy() (localctx IOrderByContext) {
	localctx = NewOrderByContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 82, CsqlParserRULE_orderBy)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(400)
		p.Match(CsqlParserK_ORDER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(401)
		p.Match(CsqlParserK_BY)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(402)
		p.OrderByField()
	}
	p.SetState(407)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == CsqlParserT__15 {
		{
			p.SetState(403)
			p.Match(CsqlParserT__15)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(404)
			p.OrderByField()
		}

		p.SetState(409)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IOrderByFieldContext is an interface to support dynamic dispatch.
type IOrderByFieldContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	CompoundName() ICompoundNameContext
	FieldIndex() IFieldIndexContext
	K_ASC() antlr.TerminalNode
	K_DESC() antlr.TerminalNode

	// IsOrderByFieldContext differentiates from other interfaces.
	IsOrderByFieldContext()
}

type OrderByFieldContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOrderByFieldContext() *OrderByFieldContext {
	var p = new(OrderByFieldContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CsqlParserRULE_orderByField
	return p
}

func InitEmptyOrderByFieldContext(p *OrderByFieldContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CsqlParserRULE_orderByField
}

func (*OrderByFieldContext) IsOrderByFieldContext() {}

func NewOrderByFieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OrderByFieldContext {
	var p = new(OrderByFieldContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CsqlParserRULE_orderByField

	return p
}

func (s *OrderByFieldContext) GetParser() antlr.Parser { return s.parser }

func (s *OrderByFieldContext) CompoundName() ICompoundNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICompoundNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICompoundNameContext)
}

func (s *OrderByFieldContext) FieldIndex() IFieldIndexContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldIndexContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldIndexContext)
}

func (s *OrderByFieldContext) K_ASC() antlr.TerminalNode {
	return s.GetToken(CsqlParserK_ASC, 0)
}

func (s *OrderByFieldContext) K_DESC() antlr.TerminalNode {
	return s.GetToken(CsqlParserK_DESC, 0)
}

func (s *OrderByFieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OrderByFieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OrderByFieldContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CsqlVisitor:
		return t.VisitOrderByField(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CsqlParser) OrderByField() (localctx IOrderByFieldContext) {
	localctx = NewOrderByFieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 84, CsqlParserRULE_orderByField)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(412)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case CsqlParserIDENTIFIER:
		{
			p.SetState(410)
			p.CompoundName()
		}

	case CsqlParserNUMERIC_LITERAL:
		{
			p.SetState(411)
			p.FieldIndex()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	p.SetState(415)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == CsqlParserK_ASC || _la == CsqlParserK_DESC {
		{
			p.SetState(414)
			_la = p.GetTokenStream().LA(1)

			if !(_la == CsqlParserK_ASC || _la == CsqlParserK_DESC) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFieldIndexContext is an interface to support dynamic dispatch.
type IFieldIndexContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NUMERIC_LITERAL() antlr.TerminalNode

	// IsFieldIndexContext differentiates from other interfaces.
	IsFieldIndexContext()
}

type FieldIndexContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldIndexContext() *FieldIndexContext {
	var p = new(FieldIndexContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CsqlParserRULE_fieldIndex
	return p
}

func InitEmptyFieldIndexContext(p *FieldIndexContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CsqlParserRULE_fieldIndex
}

func (*FieldIndexContext) IsFieldIndexContext() {}

func NewFieldIndexContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldIndexContext {
	var p = new(FieldIndexContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CsqlParserRULE_fieldIndex

	return p
}

func (s *FieldIndexContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldIndexContext) NUMERIC_LITERAL() antlr.TerminalNode {
	return s.GetToken(CsqlParserNUMERIC_LITERAL, 0)
}

func (s *FieldIndexContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldIndexContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldIndexContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CsqlVisitor:
		return t.VisitFieldIndex(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CsqlParser) FieldIndex() (localctx IFieldIndexContext) {
	localctx = NewFieldIndexContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 86, CsqlParserRULE_fieldIndex)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(417)
		p.Match(CsqlParserNUMERIC_LITERAL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IGroupByContext is an interface to support dynamic dispatch.
type IGroupByContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	K_GROUP() antlr.TerminalNode
	K_BY() antlr.TerminalNode
	AllGroupByField() []IGroupByFieldContext
	GroupByField(i int) IGroupByFieldContext

	// IsGroupByContext differentiates from other interfaces.
	IsGroupByContext()
}

type GroupByContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGroupByContext() *GroupByContext {
	var p = new(GroupByContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CsqlParserRULE_groupBy
	return p
}

func InitEmptyGroupByContext(p *GroupByContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CsqlParserRULE_groupBy
}

func (*GroupByContext) IsGroupByContext() {}

func NewGroupByContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GroupByContext {
	var p = new(GroupByContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CsqlParserRULE_groupBy

	return p
}

func (s *GroupByContext) GetParser() antlr.Parser { return s.parser }

func (s *GroupByContext) K_GROUP() antlr.TerminalNode {
	return s.GetToken(CsqlParserK_GROUP, 0)
}

func (s *GroupByContext) K_BY() antlr.TerminalNode {
	return s.GetToken(CsqlParserK_BY, 0)
}

func (s *GroupByContext) AllGroupByField() []IGroupByFieldContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IGroupByFieldContext); ok {
			len++
		}
	}

	tst := make([]IGroupByFieldContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IGroupByFieldContext); ok {
			tst[i] = t.(IGroupByFieldContext)
			i++
		}
	}

	return tst
}

func (s *GroupByContext) GroupByField(i int) IGroupByFieldContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGroupByFieldContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGroupByFieldContext)
}

func (s *GroupByContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GroupByContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GroupByContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CsqlVisitor:
		return t.VisitGroupBy(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CsqlParser) GroupBy() (localctx IGroupByContext) {
	localctx = NewGroupByContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 88, CsqlParserRULE_groupBy)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(419)
		p.Match(CsqlParserK_GROUP)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(420)
		p.Match(CsqlParserK_BY)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(421)
		p.GroupByField()
	}
	p.SetState(426)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == CsqlParserT__15 {
		{
			p.SetState(422)
			p.Match(CsqlParserT__15)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(423)
			p.GroupByField()
		}

		p.SetState(428)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IGroupByFieldContext is an interface to support dynamic dispatch.
type IGroupByFieldContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	CompoundName() ICompoundNameContext
	FieldIndex() IFieldIndexContext

	// IsGroupByFieldContext differentiates from other interfaces.
	IsGroupByFieldContext()
}

type GroupByFieldContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGroupByFieldContext() *GroupByFieldContext {
	var p = new(GroupByFieldContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CsqlParserRULE_groupByField
	return p
}

func InitEmptyGroupByFieldContext(p *GroupByFieldContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CsqlParserRULE_groupByField
}

func (*GroupByFieldContext) IsGroupByFieldContext() {}

func NewGroupByFieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GroupByFieldContext {
	var p = new(GroupByFieldContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CsqlParserRULE_groupByField

	return p
}

func (s *GroupByFieldContext) GetParser() antlr.Parser { return s.parser }

func (s *GroupByFieldContext) CompoundName() ICompoundNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICompoundNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICompoundNameContext)
}

func (s *GroupByFieldContext) FieldIndex() IFieldIndexContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldIndexContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldIndexContext)
}

func (s *GroupByFieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GroupByFieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GroupByFieldContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CsqlVisitor:
		return t.VisitGroupByField(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CsqlParser) GroupByField() (localctx IGroupByFieldContext) {
	localctx = NewGroupByFieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 90, CsqlParserRULE_groupByField)
	p.EnterOuterAlt(localctx, 1)
	p.SetState(431)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case CsqlParserIDENTIFIER:
		{
			p.SetState(429)
			p.CompoundName()
		}

	case CsqlParserNUMERIC_LITERAL:
		{
			p.SetState(430)
			p.FieldIndex()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFunctionContext is an interface to support dynamic dispatch.
type IFunctionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	K_ROUND() antlr.TerminalNode
	K_LEN() antlr.TerminalNode
	K_TRUNC() antlr.TerminalNode
	K_FRAC() antlr.TerminalNode
	K_TO_STRING() antlr.TerminalNode
	K_TO_FLOAT() antlr.TerminalNode
	K_TO_INT() antlr.TerminalNode
	K_SUBSTRING() antlr.TerminalNode
	K_TO_UPPER() antlr.TerminalNode
	K_TO_LOWER() antlr.TerminalNode
	K_POW() antlr.TerminalNode
	K_SQRT() antlr.TerminalNode
	K_COALESCE() antlr.TerminalNode

	// IsFunctionContext differentiates from other interfaces.
	IsFunctionContext()
}

type FunctionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionContext() *FunctionContext {
	var p = new(FunctionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CsqlParserRULE_function
	return p
}

func InitEmptyFunctionContext(p *FunctionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CsqlParserRULE_function
}

func (*FunctionContext) IsFunctionContext() {}

func NewFunctionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionContext {
	var p = new(FunctionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CsqlParserRULE_function

	return p
}

func (s *FunctionContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionContext) K_ROUND() antlr.TerminalNode {
	return s.GetToken(CsqlParserK_ROUND, 0)
}

func (s *FunctionContext) K_LEN() antlr.TerminalNode {
	return s.GetToken(CsqlParserK_LEN, 0)
}

func (s *FunctionContext) K_TRUNC() antlr.TerminalNode {
	return s.GetToken(CsqlParserK_TRUNC, 0)
}

func (s *FunctionContext) K_FRAC() antlr.TerminalNode {
	return s.GetToken(CsqlParserK_FRAC, 0)
}

func (s *FunctionContext) K_TO_STRING() antlr.TerminalNode {
	return s.GetToken(CsqlParserK_TO_STRING, 0)
}

func (s *FunctionContext) K_TO_FLOAT() antlr.TerminalNode {
	return s.GetToken(CsqlParserK_TO_FLOAT, 0)
}

func (s *FunctionContext) K_TO_INT() antlr.TerminalNode {
	return s.GetToken(CsqlParserK_TO_INT, 0)
}

func (s *FunctionContext) K_SUBSTRING() antlr.TerminalNode {
	return s.GetToken(CsqlParserK_SUBSTRING, 0)
}

func (s *FunctionContext) K_TO_UPPER() antlr.TerminalNode {
	return s.GetToken(CsqlParserK_TO_UPPER, 0)
}

func (s *FunctionContext) K_TO_LOWER() antlr.TerminalNode {
	return s.GetToken(CsqlParserK_TO_LOWER, 0)
}

func (s *FunctionContext) K_POW() antlr.TerminalNode {
	return s.GetToken(CsqlParserK_POW, 0)
}

func (s *FunctionContext) K_SQRT() antlr.TerminalNode {
	return s.GetToken(CsqlParserK_SQRT, 0)
}

func (s *FunctionContext) K_COALESCE() antlr.TerminalNode {
	return s.GetToken(CsqlParserK_COALESCE, 0)
}

func (s *FunctionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CsqlVisitor:
		return t.VisitFunction(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CsqlParser) Function() (localctx IFunctionContext) {
	localctx = NewFunctionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 92, CsqlParserRULE_function)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(433)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&-2251799813685248) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

func (p *CsqlParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 5:
		var t *ValueExprContext = nil
		if localctx != nil {
			t = localctx.(*ValueExprContext)
		}
		return p.ValueExpr_Sempred(t, predIndex)

	case 6:
		var t *WhereExprContext = nil
		if localctx != nil {
			t = localctx.(*WhereExprContext)
		}
		return p.WhereExpr_Sempred(t, predIndex)

	case 26:
		var t *DataSourceContext = nil
		if localctx != nil {
			t = localctx.(*DataSourceContext)
		}
		return p.DataSource_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *CsqlParser) ValueExpr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 5)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *CsqlParser) WhereExpr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 1:
		return p.Precpred(p.GetParserRuleContext(), 10)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 9)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 8)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *CsqlParser) DataSource_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 4:
		return p.Precpred(p.GetParserRuleContext(), 3)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
