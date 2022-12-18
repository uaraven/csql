// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package parser // Csql

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type CsqlParser struct {
	*antlr.BaseParser
}

var csqlParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func csqlParserInit() {
	staticData := &csqlParserStaticData
	staticData.literalNames = []string{
		"", "';'", "'<'", "'<='", "'>'", "'>='", "'='", "'!='", "'<>'", "'+'",
		"'-'", "'*'", "'/'", "'%'", "'||'", "'('", "','", "')'", "'.'",
	}
	staticData.symbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "K_AND", "K_AS", "K_BETWEEN", "K_FROM", "K_IN", "K_IS", "K_NOT",
		"K_NULL", "K_OR", "K_REGEX", "K_SELECT", "K_MATCH", "K_WHERE", "K_TRUE",
		"K_FALSE", "K_JOIN", "K_LEFT", "K_RIGHT", "K_OUTER", "K_INNER", "K_CROSS",
		"K_ON", "K_DISTINCT", "K_LIKE", "K_FULL", "K_LIMIT", "K_ORDER", "K_BY",
		"K_ASC", "K_DESC", "K_UNION", "K_ALL", "K_TO_STRING", "K_TO_FLOAT",
		"K_TO_INT", "K_ROUND", "K_LEN", "K_TRUNC", "K_FRAC", "K_SUBSTRING",
		"K_TO_UPPER", "K_TO_LOWER", "K_POW", "K_SQRT", "IDENTIFIER", "SIMPLE_IDENTIFIER",
		"QUOTED_IDENTIFIER", "NUMERIC_LITERAL", "STRING_LITERAL", "SINGLE_LINE_COMMENT",
		"MULTILINE_COMMENT", "SPACES", "UNEXPECTED_CHAR",
	}
	staticData.ruleNames = []string{
		"query", "comparisonOperator", "binaryOperation", "list", "term", "valueExpr",
		"whereExpr", "where", "distinct", "funCall", "projection", "projectionField",
		"projectionFieldName", "fieldName", "innerJoin", "leftJoin", "rightJoin",
		"fullJoin", "crossJoin", "conditionalJoinType", "conditionalJoinTarget",
		"unconditionalJoinTarget", "dataSource", "sources", "unionSelects",
		"selectStatement", "signedNumber", "stringValue", "nullValue", "literalValue",
		"alias", "sourceName", "compoundName", "name", "qualifier", "limit",
		"limitValue", "orderBy", "orderByField", "fieldIndex", "function",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 71, 381, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36, 7, 36,
		2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 1, 0, 1, 0, 3,
		0, 85, 8, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3,
		5, 3, 97, 8, 3, 10, 3, 12, 3, 100, 9, 3, 1, 3, 1, 3, 1, 4, 1, 4, 3, 4,
		106, 8, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 3, 5, 115, 8, 5, 1,
		5, 1, 5, 1, 5, 1, 5, 5, 5, 121, 8, 5, 10, 5, 12, 5, 124, 9, 5, 1, 6, 1,
		6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 3,
		6, 139, 8, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 3, 6, 146, 8, 6, 1, 6, 1, 6,
		1, 6, 1, 6, 1, 6, 3, 6, 153, 8, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6,
		3, 6, 161, 8, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 3, 6, 169, 8, 6, 1,
		6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 5, 6, 181, 8,
		6, 10, 6, 12, 6, 184, 9, 6, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 9, 1, 9, 1,
		9, 1, 9, 1, 9, 5, 9, 196, 8, 9, 10, 9, 12, 9, 199, 9, 9, 1, 9, 1, 9, 1,
		10, 3, 10, 204, 8, 10, 1, 10, 1, 10, 1, 10, 5, 10, 209, 8, 10, 10, 10,
		12, 10, 212, 9, 10, 1, 11, 1, 11, 3, 11, 216, 8, 11, 1, 11, 3, 11, 219,
		8, 11, 1, 11, 3, 11, 222, 8, 11, 1, 12, 1, 12, 1, 12, 3, 12, 227, 8, 12,
		1, 12, 1, 12, 1, 13, 1, 13, 1, 14, 3, 14, 234, 8, 14, 1, 14, 1, 14, 1,
		15, 1, 15, 3, 15, 240, 8, 15, 1, 15, 1, 15, 1, 16, 1, 16, 3, 16, 246, 8,
		16, 1, 16, 1, 16, 1, 17, 1, 17, 3, 17, 252, 8, 17, 1, 17, 1, 17, 1, 18,
		1, 18, 1, 18, 1, 19, 1, 19, 1, 19, 1, 19, 3, 19, 263, 8, 19, 1, 20, 1,
		20, 1, 20, 1, 20, 1, 21, 1, 21, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22,
		3, 22, 277, 8, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1,
		22, 5, 22, 287, 8, 22, 10, 22, 12, 22, 290, 9, 22, 1, 23, 1, 23, 1, 23,
		5, 23, 295, 8, 23, 10, 23, 12, 23, 298, 9, 23, 1, 24, 1, 24, 1, 24, 3,
		24, 303, 8, 24, 1, 24, 5, 24, 306, 8, 24, 10, 24, 12, 24, 309, 9, 24, 1,
		25, 1, 25, 1, 25, 1, 25, 1, 25, 3, 25, 316, 8, 25, 1, 25, 3, 25, 319, 8,
		25, 1, 25, 3, 25, 322, 8, 25, 1, 26, 3, 26, 325, 8, 26, 1, 26, 1, 26, 1,
		27, 1, 27, 1, 28, 1, 28, 1, 29, 1, 29, 1, 29, 3, 29, 336, 8, 29, 1, 30,
		1, 30, 1, 31, 1, 31, 3, 31, 342, 8, 31, 1, 32, 1, 32, 1, 32, 3, 32, 347,
		8, 32, 1, 32, 1, 32, 1, 33, 1, 33, 1, 34, 1, 34, 1, 35, 1, 35, 1, 35, 1,
		36, 1, 36, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 5, 37, 365, 8, 37, 10, 37,
		12, 37, 368, 9, 37, 1, 38, 1, 38, 3, 38, 372, 8, 38, 1, 38, 3, 38, 375,
		8, 38, 1, 39, 1, 39, 1, 40, 1, 40, 1, 40, 0, 3, 10, 12, 44, 41, 0, 2, 4,
		6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42,
		44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 66, 68, 70, 72, 74, 76, 78,
		80, 0, 6, 1, 0, 2, 8, 1, 0, 9, 14, 2, 0, 11, 11, 63, 63, 1, 0, 9, 10, 1,
		0, 47, 48, 1, 0, 51, 62, 391, 0, 82, 1, 0, 0, 0, 2, 88, 1, 0, 0, 0, 4,
		90, 1, 0, 0, 0, 6, 92, 1, 0, 0, 0, 8, 105, 1, 0, 0, 0, 10, 114, 1, 0, 0,
		0, 12, 168, 1, 0, 0, 0, 14, 185, 1, 0, 0, 0, 16, 188, 1, 0, 0, 0, 18, 190,
		1, 0, 0, 0, 20, 203, 1, 0, 0, 0, 22, 215, 1, 0, 0, 0, 24, 226, 1, 0, 0,
		0, 26, 230, 1, 0, 0, 0, 28, 233, 1, 0, 0, 0, 30, 237, 1, 0, 0, 0, 32, 243,
		1, 0, 0, 0, 34, 249, 1, 0, 0, 0, 36, 255, 1, 0, 0, 0, 38, 262, 1, 0, 0,
		0, 40, 264, 1, 0, 0, 0, 42, 268, 1, 0, 0, 0, 44, 276, 1, 0, 0, 0, 46, 291,
		1, 0, 0, 0, 48, 299, 1, 0, 0, 0, 50, 310, 1, 0, 0, 0, 52, 324, 1, 0, 0,
		0, 54, 328, 1, 0, 0, 0, 56, 330, 1, 0, 0, 0, 58, 335, 1, 0, 0, 0, 60, 337,
		1, 0, 0, 0, 62, 339, 1, 0, 0, 0, 64, 346, 1, 0, 0, 0, 66, 350, 1, 0, 0,
		0, 68, 352, 1, 0, 0, 0, 70, 354, 1, 0, 0, 0, 72, 357, 1, 0, 0, 0, 74, 359,
		1, 0, 0, 0, 76, 371, 1, 0, 0, 0, 78, 376, 1, 0, 0, 0, 80, 378, 1, 0, 0,
		0, 82, 84, 3, 48, 24, 0, 83, 85, 5, 1, 0, 0, 84, 83, 1, 0, 0, 0, 84, 85,
		1, 0, 0, 0, 85, 86, 1, 0, 0, 0, 86, 87, 5, 0, 0, 1, 87, 1, 1, 0, 0, 0,
		88, 89, 7, 0, 0, 0, 89, 3, 1, 0, 0, 0, 90, 91, 7, 1, 0, 0, 91, 5, 1, 0,
		0, 0, 92, 93, 5, 15, 0, 0, 93, 98, 3, 10, 5, 0, 94, 95, 5, 16, 0, 0, 95,
		97, 3, 10, 5, 0, 96, 94, 1, 0, 0, 0, 97, 100, 1, 0, 0, 0, 98, 96, 1, 0,
		0, 0, 98, 99, 1, 0, 0, 0, 99, 101, 1, 0, 0, 0, 100, 98, 1, 0, 0, 0, 101,
		102, 5, 17, 0, 0, 102, 7, 1, 0, 0, 0, 103, 106, 3, 64, 32, 0, 104, 106,
		3, 58, 29, 0, 105, 103, 1, 0, 0, 0, 105, 104, 1, 0, 0, 0, 106, 9, 1, 0,
		0, 0, 107, 108, 6, 5, -1, 0, 108, 115, 3, 8, 4, 0, 109, 115, 3, 18, 9,
		0, 110, 111, 5, 15, 0, 0, 111, 112, 3, 10, 5, 0, 112, 113, 5, 17, 0, 0,
		113, 115, 1, 0, 0, 0, 114, 107, 1, 0, 0, 0, 114, 109, 1, 0, 0, 0, 114,
		110, 1, 0, 0, 0, 115, 122, 1, 0, 0, 0, 116, 117, 10, 4, 0, 0, 117, 118,
		3, 4, 2, 0, 118, 119, 3, 10, 5, 5, 119, 121, 1, 0, 0, 0, 120, 116, 1, 0,
		0, 0, 121, 124, 1, 0, 0, 0, 122, 120, 1, 0, 0, 0, 122, 123, 1, 0, 0, 0,
		123, 11, 1, 0, 0, 0, 124, 122, 1, 0, 0, 0, 125, 126, 6, 6, -1, 0, 126,
		169, 3, 8, 4, 0, 127, 169, 3, 10, 5, 0, 128, 129, 5, 25, 0, 0, 129, 169,
		3, 12, 6, 7, 130, 131, 3, 10, 5, 0, 131, 132, 5, 21, 0, 0, 132, 133, 3,
		10, 5, 0, 133, 134, 5, 19, 0, 0, 134, 135, 3, 10, 5, 0, 135, 169, 1, 0,
		0, 0, 136, 138, 3, 10, 5, 0, 137, 139, 5, 25, 0, 0, 138, 137, 1, 0, 0,
		0, 138, 139, 1, 0, 0, 0, 139, 140, 1, 0, 0, 0, 140, 141, 5, 42, 0, 0, 141,
		142, 3, 10, 5, 0, 142, 169, 1, 0, 0, 0, 143, 145, 3, 10, 5, 0, 144, 146,
		5, 25, 0, 0, 145, 144, 1, 0, 0, 0, 145, 146, 1, 0, 0, 0, 146, 147, 1, 0,
		0, 0, 147, 148, 5, 30, 0, 0, 148, 149, 3, 10, 5, 0, 149, 169, 1, 0, 0,
		0, 150, 152, 3, 10, 5, 0, 151, 153, 5, 25, 0, 0, 152, 151, 1, 0, 0, 0,
		152, 153, 1, 0, 0, 0, 153, 154, 1, 0, 0, 0, 154, 155, 5, 23, 0, 0, 155,
		156, 3, 6, 3, 0, 156, 169, 1, 0, 0, 0, 157, 158, 3, 10, 5, 0, 158, 160,
		5, 24, 0, 0, 159, 161, 5, 25, 0, 0, 160, 159, 1, 0, 0, 0, 160, 161, 1,
		0, 0, 0, 161, 162, 1, 0, 0, 0, 162, 163, 5, 26, 0, 0, 163, 169, 1, 0, 0,
		0, 164, 165, 5, 15, 0, 0, 165, 166, 3, 12, 6, 0, 166, 167, 5, 17, 0, 0,
		167, 169, 1, 0, 0, 0, 168, 125, 1, 0, 0, 0, 168, 127, 1, 0, 0, 0, 168,
		128, 1, 0, 0, 0, 168, 130, 1, 0, 0, 0, 168, 136, 1, 0, 0, 0, 168, 143,
		1, 0, 0, 0, 168, 150, 1, 0, 0, 0, 168, 157, 1, 0, 0, 0, 168, 164, 1, 0,
		0, 0, 169, 182, 1, 0, 0, 0, 170, 171, 10, 10, 0, 0, 171, 172, 3, 2, 1,
		0, 172, 173, 3, 12, 6, 11, 173, 181, 1, 0, 0, 0, 174, 175, 10, 9, 0, 0,
		175, 176, 5, 19, 0, 0, 176, 181, 3, 12, 6, 10, 177, 178, 10, 8, 0, 0, 178,
		179, 5, 27, 0, 0, 179, 181, 3, 12, 6, 9, 180, 170, 1, 0, 0, 0, 180, 174,
		1, 0, 0, 0, 180, 177, 1, 0, 0, 0, 181, 184, 1, 0, 0, 0, 182, 180, 1, 0,
		0, 0, 182, 183, 1, 0, 0, 0, 183, 13, 1, 0, 0, 0, 184, 182, 1, 0, 0, 0,
		185, 186, 5, 31, 0, 0, 186, 187, 3, 12, 6, 0, 187, 15, 1, 0, 0, 0, 188,
		189, 5, 41, 0, 0, 189, 17, 1, 0, 0, 0, 190, 191, 3, 80, 40, 0, 191, 192,
		5, 15, 0, 0, 192, 197, 3, 10, 5, 0, 193, 194, 5, 16, 0, 0, 194, 196, 3,
		10, 5, 0, 195, 193, 1, 0, 0, 0, 196, 199, 1, 0, 0, 0, 197, 195, 1, 0, 0,
		0, 197, 198, 1, 0, 0, 0, 198, 200, 1, 0, 0, 0, 199, 197, 1, 0, 0, 0, 200,
		201, 5, 17, 0, 0, 201, 19, 1, 0, 0, 0, 202, 204, 3, 16, 8, 0, 203, 202,
		1, 0, 0, 0, 203, 204, 1, 0, 0, 0, 204, 205, 1, 0, 0, 0, 205, 210, 3, 22,
		11, 0, 206, 207, 5, 16, 0, 0, 207, 209, 3, 22, 11, 0, 208, 206, 1, 0, 0,
		0, 209, 212, 1, 0, 0, 0, 210, 208, 1, 0, 0, 0, 210, 211, 1, 0, 0, 0, 211,
		21, 1, 0, 0, 0, 212, 210, 1, 0, 0, 0, 213, 216, 3, 24, 12, 0, 214, 216,
		3, 10, 5, 0, 215, 213, 1, 0, 0, 0, 215, 214, 1, 0, 0, 0, 216, 221, 1, 0,
		0, 0, 217, 219, 5, 20, 0, 0, 218, 217, 1, 0, 0, 0, 218, 219, 1, 0, 0, 0,
		219, 220, 1, 0, 0, 0, 220, 222, 3, 60, 30, 0, 221, 218, 1, 0, 0, 0, 221,
		222, 1, 0, 0, 0, 222, 23, 1, 0, 0, 0, 223, 224, 3, 68, 34, 0, 224, 225,
		5, 18, 0, 0, 225, 227, 1, 0, 0, 0, 226, 223, 1, 0, 0, 0, 226, 227, 1, 0,
		0, 0, 227, 228, 1, 0, 0, 0, 228, 229, 3, 26, 13, 0, 229, 25, 1, 0, 0, 0,
		230, 231, 7, 2, 0, 0, 231, 27, 1, 0, 0, 0, 232, 234, 5, 38, 0, 0, 233,
		232, 1, 0, 0, 0, 233, 234, 1, 0, 0, 0, 234, 235, 1, 0, 0, 0, 235, 236,
		5, 34, 0, 0, 236, 29, 1, 0, 0, 0, 237, 239, 5, 35, 0, 0, 238, 240, 5, 37,
		0, 0, 239, 238, 1, 0, 0, 0, 239, 240, 1, 0, 0, 0, 240, 241, 1, 0, 0, 0,
		241, 242, 5, 34, 0, 0, 242, 31, 1, 0, 0, 0, 243, 245, 5, 36, 0, 0, 244,
		246, 5, 37, 0, 0, 245, 244, 1, 0, 0, 0, 245, 246, 1, 0, 0, 0, 246, 247,
		1, 0, 0, 0, 247, 248, 5, 34, 0, 0, 248, 33, 1, 0, 0, 0, 249, 251, 5, 43,
		0, 0, 250, 252, 5, 37, 0, 0, 251, 250, 1, 0, 0, 0, 251, 252, 1, 0, 0, 0,
		252, 253, 1, 0, 0, 0, 253, 254, 5, 34, 0, 0, 254, 35, 1, 0, 0, 0, 255,
		256, 5, 39, 0, 0, 256, 257, 5, 34, 0, 0, 257, 37, 1, 0, 0, 0, 258, 263,
		3, 28, 14, 0, 259, 263, 3, 30, 15, 0, 260, 263, 3, 32, 16, 0, 261, 263,
		3, 34, 17, 0, 262, 258, 1, 0, 0, 0, 262, 259, 1, 0, 0, 0, 262, 260, 1,
		0, 0, 0, 262, 261, 1, 0, 0, 0, 263, 39, 1, 0, 0, 0, 264, 265, 3, 44, 22,
		0, 265, 266, 5, 40, 0, 0, 266, 267, 3, 12, 6, 0, 267, 41, 1, 0, 0, 0, 268,
		269, 3, 44, 22, 0, 269, 43, 1, 0, 0, 0, 270, 271, 6, 22, -1, 0, 271, 277,
		3, 62, 31, 0, 272, 273, 5, 15, 0, 0, 273, 274, 3, 44, 22, 0, 274, 275,
		5, 17, 0, 0, 275, 277, 1, 0, 0, 0, 276, 270, 1, 0, 0, 0, 276, 272, 1, 0,
		0, 0, 277, 288, 1, 0, 0, 0, 278, 279, 10, 3, 0, 0, 279, 280, 3, 38, 19,
		0, 280, 281, 3, 40, 20, 0, 281, 287, 1, 0, 0, 0, 282, 283, 10, 2, 0, 0,
		283, 284, 3, 36, 18, 0, 284, 285, 3, 42, 21, 0, 285, 287, 1, 0, 0, 0, 286,
		278, 1, 0, 0, 0, 286, 282, 1, 0, 0, 0, 287, 290, 1, 0, 0, 0, 288, 286,
		1, 0, 0, 0, 288, 289, 1, 0, 0, 0, 289, 45, 1, 0, 0, 0, 290, 288, 1, 0,
		0, 0, 291, 296, 3, 44, 22, 0, 292, 293, 5, 16, 0, 0, 293, 295, 3, 44, 22,
		0, 294, 292, 1, 0, 0, 0, 295, 298, 1, 0, 0, 0, 296, 294, 1, 0, 0, 0, 296,
		297, 1, 0, 0, 0, 297, 47, 1, 0, 0, 0, 298, 296, 1, 0, 0, 0, 299, 307, 3,
		50, 25, 0, 300, 302, 5, 49, 0, 0, 301, 303, 5, 50, 0, 0, 302, 301, 1, 0,
		0, 0, 302, 303, 1, 0, 0, 0, 303, 304, 1, 0, 0, 0, 304, 306, 3, 48, 24,
		0, 305, 300, 1, 0, 0, 0, 306, 309, 1, 0, 0, 0, 307, 305, 1, 0, 0, 0, 307,
		308, 1, 0, 0, 0, 308, 49, 1, 0, 0, 0, 309, 307, 1, 0, 0, 0, 310, 311, 5,
		29, 0, 0, 311, 312, 3, 20, 10, 0, 312, 313, 5, 22, 0, 0, 313, 315, 3, 46,
		23, 0, 314, 316, 3, 14, 7, 0, 315, 314, 1, 0, 0, 0, 315, 316, 1, 0, 0,
		0, 316, 318, 1, 0, 0, 0, 317, 319, 3, 74, 37, 0, 318, 317, 1, 0, 0, 0,
		318, 319, 1, 0, 0, 0, 319, 321, 1, 0, 0, 0, 320, 322, 3, 70, 35, 0, 321,
		320, 1, 0, 0, 0, 321, 322, 1, 0, 0, 0, 322, 51, 1, 0, 0, 0, 323, 325, 7,
		3, 0, 0, 324, 323, 1, 0, 0, 0, 324, 325, 1, 0, 0, 0, 325, 326, 1, 0, 0,
		0, 326, 327, 5, 66, 0, 0, 327, 53, 1, 0, 0, 0, 328, 329, 5, 67, 0, 0, 329,
		55, 1, 0, 0, 0, 330, 331, 5, 26, 0, 0, 331, 57, 1, 0, 0, 0, 332, 336, 3,
		52, 26, 0, 333, 336, 3, 54, 27, 0, 334, 336, 3, 56, 28, 0, 335, 332, 1,
		0, 0, 0, 335, 333, 1, 0, 0, 0, 335, 334, 1, 0, 0, 0, 336, 59, 1, 0, 0,
		0, 337, 338, 5, 63, 0, 0, 338, 61, 1, 0, 0, 0, 339, 341, 3, 66, 33, 0,
		340, 342, 3, 60, 30, 0, 341, 340, 1, 0, 0, 0, 341, 342, 1, 0, 0, 0, 342,
		63, 1, 0, 0, 0, 343, 344, 3, 68, 34, 0, 344, 345, 5, 18, 0, 0, 345, 347,
		1, 0, 0, 0, 346, 343, 1, 0, 0, 0, 346, 347, 1, 0, 0, 0, 347, 348, 1, 0,
		0, 0, 348, 349, 3, 66, 33, 0, 349, 65, 1, 0, 0, 0, 350, 351, 5, 63, 0,
		0, 351, 67, 1, 0, 0, 0, 352, 353, 5, 63, 0, 0, 353, 69, 1, 0, 0, 0, 354,
		355, 5, 44, 0, 0, 355, 356, 3, 72, 36, 0, 356, 71, 1, 0, 0, 0, 357, 358,
		5, 66, 0, 0, 358, 73, 1, 0, 0, 0, 359, 360, 5, 45, 0, 0, 360, 361, 5, 46,
		0, 0, 361, 366, 3, 76, 38, 0, 362, 363, 5, 16, 0, 0, 363, 365, 3, 76, 38,
		0, 364, 362, 1, 0, 0, 0, 365, 368, 1, 0, 0, 0, 366, 364, 1, 0, 0, 0, 366,
		367, 1, 0, 0, 0, 367, 75, 1, 0, 0, 0, 368, 366, 1, 0, 0, 0, 369, 372, 3,
		64, 32, 0, 370, 372, 3, 78, 39, 0, 371, 369, 1, 0, 0, 0, 371, 370, 1, 0,
		0, 0, 372, 374, 1, 0, 0, 0, 373, 375, 7, 4, 0, 0, 374, 373, 1, 0, 0, 0,
		374, 375, 1, 0, 0, 0, 375, 77, 1, 0, 0, 0, 376, 377, 5, 66, 0, 0, 377,
		79, 1, 0, 0, 0, 378, 379, 7, 5, 0, 0, 379, 81, 1, 0, 0, 0, 40, 84, 98,
		105, 114, 122, 138, 145, 152, 160, 168, 180, 182, 197, 203, 210, 215, 218,
		221, 226, 233, 239, 245, 251, 262, 276, 286, 288, 296, 302, 307, 315, 318,
		321, 324, 335, 341, 346, 366, 371, 374,
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
	staticData := &csqlParserStaticData
	staticData.once.Do(csqlParserInit)
}

// NewCsqlParser produces a new parser instance for the optional input antlr.TokenStream.
func NewCsqlParser(input antlr.TokenStream) *CsqlParser {
	CsqlParserInit()
	this := new(CsqlParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &csqlParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "java-escape"

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
	CsqlParserIDENTIFIER          = 63
	CsqlParserSIMPLE_IDENTIFIER   = 64
	CsqlParserQUOTED_IDENTIFIER   = 65
	CsqlParserNUMERIC_LITERAL     = 66
	CsqlParserSTRING_LITERAL      = 67
	CsqlParserSINGLE_LINE_COMMENT = 68
	CsqlParserMULTILINE_COMMENT   = 69
	CsqlParserSPACES              = 70
	CsqlParserUNEXPECTED_CHAR     = 71
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
	CsqlParserRULE_distinct                = 8
	CsqlParserRULE_funCall                 = 9
	CsqlParserRULE_projection              = 10
	CsqlParserRULE_projectionField         = 11
	CsqlParserRULE_projectionFieldName     = 12
	CsqlParserRULE_fieldName               = 13
	CsqlParserRULE_innerJoin               = 14
	CsqlParserRULE_leftJoin                = 15
	CsqlParserRULE_rightJoin               = 16
	CsqlParserRULE_fullJoin                = 17
	CsqlParserRULE_crossJoin               = 18
	CsqlParserRULE_conditionalJoinType     = 19
	CsqlParserRULE_conditionalJoinTarget   = 20
	CsqlParserRULE_unconditionalJoinTarget = 21
	CsqlParserRULE_dataSource              = 22
	CsqlParserRULE_sources                 = 23
	CsqlParserRULE_unionSelects            = 24
	CsqlParserRULE_selectStatement         = 25
	CsqlParserRULE_signedNumber            = 26
	CsqlParserRULE_stringValue             = 27
	CsqlParserRULE_nullValue               = 28
	CsqlParserRULE_literalValue            = 29
	CsqlParserRULE_alias                   = 30
	CsqlParserRULE_sourceName              = 31
	CsqlParserRULE_compoundName            = 32
	CsqlParserRULE_name                    = 33
	CsqlParserRULE_qualifier               = 34
	CsqlParserRULE_limit                   = 35
	CsqlParserRULE_limitValue              = 36
	CsqlParserRULE_orderBy                 = 37
	CsqlParserRULE_orderByField            = 38
	CsqlParserRULE_fieldIndex              = 39
	CsqlParserRULE_function                = 40
)

// IQueryContext is an interface to support dynamic dispatch.
type IQueryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsQueryContext differentiates from other interfaces.
	IsQueryContext()
}

type QueryContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyQueryContext() *QueryContext {
	var p = new(QueryContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CsqlParserRULE_query
	return p
}

func (*QueryContext) IsQueryContext() {}

func NewQueryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *QueryContext {
	var p = new(QueryContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

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
	this := p
	_ = this

	localctx = NewQueryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, CsqlParserRULE_query)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(82)
		p.UnionSelects()
	}
	p.SetState(84)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == CsqlParserT__0 {
		{
			p.SetState(83)
			p.Match(CsqlParserT__0)
		}

	}
	{
		p.SetState(86)
		p.Match(CsqlParserEOF)
	}

	return localctx
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
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyComparisonOperatorContext() *ComparisonOperatorContext {
	var p = new(ComparisonOperatorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CsqlParserRULE_comparisonOperator
	return p
}

func (*ComparisonOperatorContext) IsComparisonOperatorContext() {}

func NewComparisonOperatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ComparisonOperatorContext {
	var p = new(ComparisonOperatorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

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
	this := p
	_ = this

	localctx = NewComparisonOperatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, CsqlParserRULE_comparisonOperator)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(88)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&508) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
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
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBinaryOperationContext() *BinaryOperationContext {
	var p = new(BinaryOperationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CsqlParserRULE_binaryOperation
	return p
}

func (*BinaryOperationContext) IsBinaryOperationContext() {}

func NewBinaryOperationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BinaryOperationContext {
	var p = new(BinaryOperationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

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
	this := p
	_ = this

	localctx = NewBinaryOperationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, CsqlParserRULE_binaryOperation)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(90)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&32256) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IListContext is an interface to support dynamic dispatch.
type IListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsListContext differentiates from other interfaces.
	IsListContext()
}

type ListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyListContext() *ListContext {
	var p = new(ListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CsqlParserRULE_list
	return p
}

func (*ListContext) IsListContext() {}

func NewListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListContext {
	var p = new(ListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

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
	this := p
	_ = this

	localctx = NewListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, CsqlParserRULE_list)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(92)
		p.Match(CsqlParserT__14)
	}
	{
		p.SetState(93)
		p.valueExpr(0)
	}
	p.SetState(98)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == CsqlParserT__15 {
		{
			p.SetState(94)
			p.Match(CsqlParserT__15)
		}
		{
			p.SetState(95)
			p.valueExpr(0)
		}

		p.SetState(100)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(101)
		p.Match(CsqlParserT__16)
	}

	return localctx
}

// ITermContext is an interface to support dynamic dispatch.
type ITermContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTermContext differentiates from other interfaces.
	IsTermContext()
}

type TermContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTermContext() *TermContext {
	var p = new(TermContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CsqlParserRULE_term
	return p
}

func (*TermContext) IsTermContext() {}

func NewTermContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TermContext {
	var p = new(TermContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

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
	this := p
	_ = this

	localctx = NewTermContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, CsqlParserRULE_term)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(105)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case CsqlParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(103)
			p.CompoundName()
		}

	case CsqlParserT__8, CsqlParserT__9, CsqlParserK_NULL, CsqlParserNUMERIC_LITERAL, CsqlParserSTRING_LITERAL:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(104)
			p.LiteralValue()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
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
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValueExprContext() *ValueExprContext {
	var p = new(ValueExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CsqlParserRULE_valueExpr
	return p
}

func (*ValueExprContext) IsValueExprContext() {}

func NewValueExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueExprContext {
	var p = new(ValueExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CsqlParserRULE_valueExpr

	return p
}

func (s *ValueExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueExprContext) CopyFrom(ctx *ValueExprContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ValueExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ValueTermContext struct {
	*ValueExprContext
}

func NewValueTermContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ValueTermContext {
	var p = new(ValueTermContext)

	p.ValueExprContext = NewEmptyValueExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ValueExprContext))

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
	*ValueExprContext
}

func NewValueBinaryExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ValueBinaryExprContext {
	var p = new(ValueBinaryExprContext)

	p.ValueExprContext = NewEmptyValueExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ValueExprContext))

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
	*ValueExprContext
}

func NewValueParensExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ValueParensExprContext {
	var p = new(ValueParensExprContext)

	p.ValueExprContext = NewEmptyValueExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ValueExprContext))

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
	*ValueExprContext
}

func NewFunctionExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FunctionExprContext {
	var p = new(FunctionExprContext)

	p.ValueExprContext = NewEmptyValueExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ValueExprContext))

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

func (p *CsqlParser) ValueExpr() (localctx IValueExprContext) {
	return p.valueExpr(0)
}

func (p *CsqlParser) valueExpr(_p int) (localctx IValueExprContext) {
	this := p
	_ = this

	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewValueExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IValueExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 10
	p.EnterRecursionRule(localctx, 10, CsqlParserRULE_valueExpr, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(114)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case CsqlParserT__8, CsqlParserT__9, CsqlParserK_NULL, CsqlParserIDENTIFIER, CsqlParserNUMERIC_LITERAL, CsqlParserSTRING_LITERAL:
		localctx = NewValueTermContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(108)
			p.Term()
		}

	case CsqlParserK_TO_STRING, CsqlParserK_TO_FLOAT, CsqlParserK_TO_INT, CsqlParserK_ROUND, CsqlParserK_LEN, CsqlParserK_TRUNC, CsqlParserK_FRAC, CsqlParserK_SUBSTRING, CsqlParserK_TO_UPPER, CsqlParserK_TO_LOWER, CsqlParserK_POW, CsqlParserK_SQRT:
		localctx = NewFunctionExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(109)
			p.FunCall()
		}

	case CsqlParserT__14:
		localctx = NewValueParensExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(110)
			p.Match(CsqlParserT__14)
		}
		{
			p.SetState(111)
			p.valueExpr(0)
		}
		{
			p.SetState(112)
			p.Match(CsqlParserT__16)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(122)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewValueBinaryExprContext(p, NewValueExprContext(p, _parentctx, _parentState))
			p.PushNewRecursionContext(localctx, _startState, CsqlParserRULE_valueExpr)
			p.SetState(116)

			if !(p.Precpred(p.GetParserRuleContext(), 4)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
			}
			{
				p.SetState(117)
				p.BinaryOperation()
			}
			{
				p.SetState(118)
				p.valueExpr(5)
			}

		}
		p.SetState(124)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext())
	}

	return localctx
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
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWhereExprContext() *WhereExprContext {
	var p = new(WhereExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CsqlParserRULE_whereExpr
	return p
}

func (*WhereExprContext) IsWhereExprContext() {}

func NewWhereExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WhereExprContext {
	var p = new(WhereExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CsqlParserRULE_whereExpr

	return p
}

func (s *WhereExprContext) GetParser() antlr.Parser { return s.parser }

func (s *WhereExprContext) CopyFrom(ctx *WhereExprContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *WhereExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WhereExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ConditionContext struct {
	*WhereExprContext
}

func NewConditionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ConditionContext {
	var p = new(ConditionContext)

	p.WhereExprContext = NewEmptyWhereExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*WhereExprContext))

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
	*WhereExprContext
}

func NewNotExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NotExprContext {
	var p = new(NotExprContext)

	p.WhereExprContext = NewEmptyWhereExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*WhereExprContext))

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
	*WhereExprContext
}

func NewTermItemContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TermItemContext {
	var p = new(TermItemContext)

	p.WhereExprContext = NewEmptyWhereExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*WhereExprContext))

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
	*WhereExprContext
}

func NewInExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *InExprContext {
	var p = new(InExprContext)

	p.WhereExprContext = NewEmptyWhereExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*WhereExprContext))

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
	*WhereExprContext
}

func NewIsNullExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IsNullExprContext {
	var p = new(IsNullExprContext)

	p.WhereExprContext = NewEmptyWhereExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*WhereExprContext))

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
	*WhereExprContext
}

func NewLikeExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LikeExprContext {
	var p = new(LikeExprContext)

	p.WhereExprContext = NewEmptyWhereExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*WhereExprContext))

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
	*WhereExprContext
}

func NewMatchExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MatchExprContext {
	var p = new(MatchExprContext)

	p.WhereExprContext = NewEmptyWhereExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*WhereExprContext))

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
	*WhereExprContext
}

func NewOrExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *OrExprContext {
	var p = new(OrExprContext)

	p.WhereExprContext = NewEmptyWhereExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*WhereExprContext))

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
	*WhereExprContext
}

func NewParensExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParensExprContext {
	var p = new(ParensExprContext)

	p.WhereExprContext = NewEmptyWhereExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*WhereExprContext))

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
	*WhereExprContext
}

func NewValueExprItemContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ValueExprItemContext {
	var p = new(ValueExprItemContext)

	p.WhereExprContext = NewEmptyWhereExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*WhereExprContext))

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
	*WhereExprContext
}

func NewBetweenExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BetweenExprContext {
	var p = new(BetweenExprContext)

	p.WhereExprContext = NewEmptyWhereExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*WhereExprContext))

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
	*WhereExprContext
}

func NewAndExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AndExprContext {
	var p = new(AndExprContext)

	p.WhereExprContext = NewEmptyWhereExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*WhereExprContext))

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
	this := p
	_ = this

	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewWhereExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IWhereExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 12
	p.EnterRecursionRule(localctx, 12, CsqlParserRULE_whereExpr, _p)
	var _la int

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(168)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext()) {
	case 1:
		localctx = NewTermItemContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(126)
			p.Term()
		}

	case 2:
		localctx = NewValueExprItemContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(127)
			p.valueExpr(0)
		}

	case 3:
		localctx = NewNotExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(128)
			p.Match(CsqlParserK_NOT)
		}
		{
			p.SetState(129)
			p.whereExpr(7)
		}

	case 4:
		localctx = NewBetweenExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(130)
			p.valueExpr(0)
		}
		{
			p.SetState(131)
			p.Match(CsqlParserK_BETWEEN)
		}
		{
			p.SetState(132)
			p.valueExpr(0)
		}
		{
			p.SetState(133)
			p.Match(CsqlParserK_AND)
		}
		{
			p.SetState(134)
			p.valueExpr(0)
		}

	case 5:
		localctx = NewLikeExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(136)
			p.valueExpr(0)
		}
		p.SetState(138)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == CsqlParserK_NOT {
			{
				p.SetState(137)
				p.Match(CsqlParserK_NOT)
			}

		}
		{
			p.SetState(140)
			p.Match(CsqlParserK_LIKE)
		}
		{
			p.SetState(141)
			p.valueExpr(0)
		}

	case 6:
		localctx = NewMatchExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(143)
			p.valueExpr(0)
		}
		p.SetState(145)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == CsqlParserK_NOT {
			{
				p.SetState(144)
				p.Match(CsqlParserK_NOT)
			}

		}
		{
			p.SetState(147)
			p.Match(CsqlParserK_MATCH)
		}
		{
			p.SetState(148)
			p.valueExpr(0)
		}

	case 7:
		localctx = NewInExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(150)
			p.valueExpr(0)
		}
		p.SetState(152)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == CsqlParserK_NOT {
			{
				p.SetState(151)
				p.Match(CsqlParserK_NOT)
			}

		}
		{
			p.SetState(154)
			p.Match(CsqlParserK_IN)
		}
		{
			p.SetState(155)
			p.List()
		}

	case 8:
		localctx = NewIsNullExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(157)
			p.valueExpr(0)
		}
		{
			p.SetState(158)
			p.Match(CsqlParserK_IS)
		}
		p.SetState(160)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == CsqlParserK_NOT {
			{
				p.SetState(159)
				p.Match(CsqlParserK_NOT)
			}

		}
		{
			p.SetState(162)
			p.Match(CsqlParserK_NULL)
		}

	case 9:
		localctx = NewParensExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(164)
			p.Match(CsqlParserT__14)
		}
		{
			p.SetState(165)
			p.whereExpr(0)
		}
		{
			p.SetState(166)
			p.Match(CsqlParserT__16)
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(182)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(180)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext()) {
			case 1:
				localctx = NewConditionContext(p, NewWhereExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, CsqlParserRULE_whereExpr)
				p.SetState(170)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
				}
				{
					p.SetState(171)
					p.ComparisonOperator()
				}
				{
					p.SetState(172)
					p.whereExpr(11)
				}

			case 2:
				localctx = NewAndExprContext(p, NewWhereExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, CsqlParserRULE_whereExpr)
				p.SetState(174)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
				}
				{
					p.SetState(175)
					p.Match(CsqlParserK_AND)
				}
				{
					p.SetState(176)
					p.whereExpr(10)
				}

			case 3:
				localctx = NewOrExprContext(p, NewWhereExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, CsqlParserRULE_whereExpr)
				p.SetState(177)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
				}
				{
					p.SetState(178)
					p.Match(CsqlParserK_OR)
				}
				{
					p.SetState(179)
					p.whereExpr(9)
				}

			}

		}
		p.SetState(184)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext())
	}

	return localctx
}

// IWhereContext is an interface to support dynamic dispatch.
type IWhereContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsWhereContext differentiates from other interfaces.
	IsWhereContext()
}

type WhereContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWhereContext() *WhereContext {
	var p = new(WhereContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CsqlParserRULE_where
	return p
}

func (*WhereContext) IsWhereContext() {}

func NewWhereContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WhereContext {
	var p = new(WhereContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

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
	this := p
	_ = this

	localctx = NewWhereContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, CsqlParserRULE_where)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(185)
		p.Match(CsqlParserK_WHERE)
	}
	{
		p.SetState(186)
		p.whereExpr(0)
	}

	return localctx
}

// IDistinctContext is an interface to support dynamic dispatch.
type IDistinctContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDistinctContext differentiates from other interfaces.
	IsDistinctContext()
}

type DistinctContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDistinctContext() *DistinctContext {
	var p = new(DistinctContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CsqlParserRULE_distinct
	return p
}

func (*DistinctContext) IsDistinctContext() {}

func NewDistinctContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DistinctContext {
	var p = new(DistinctContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

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
	this := p
	_ = this

	localctx = NewDistinctContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, CsqlParserRULE_distinct)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(188)
		p.Match(CsqlParserK_DISTINCT)
	}

	return localctx
}

// IFunCallContext is an interface to support dynamic dispatch.
type IFunCallContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunCallContext differentiates from other interfaces.
	IsFunCallContext()
}

type FunCallContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunCallContext() *FunCallContext {
	var p = new(FunCallContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CsqlParserRULE_funCall
	return p
}

func (*FunCallContext) IsFunCallContext() {}

func NewFunCallContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunCallContext {
	var p = new(FunCallContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

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
	this := p
	_ = this

	localctx = NewFunCallContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, CsqlParserRULE_funCall)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(190)
		p.Function()
	}
	{
		p.SetState(191)
		p.Match(CsqlParserT__14)
	}
	{
		p.SetState(192)
		p.valueExpr(0)
	}
	p.SetState(197)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == CsqlParserT__15 {
		{
			p.SetState(193)
			p.Match(CsqlParserT__15)
		}
		{
			p.SetState(194)
			p.valueExpr(0)
		}

		p.SetState(199)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(200)
		p.Match(CsqlParserT__16)
	}

	return localctx
}

// IProjectionContext is an interface to support dynamic dispatch.
type IProjectionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsProjectionContext differentiates from other interfaces.
	IsProjectionContext()
}

type ProjectionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProjectionContext() *ProjectionContext {
	var p = new(ProjectionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CsqlParserRULE_projection
	return p
}

func (*ProjectionContext) IsProjectionContext() {}

func NewProjectionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProjectionContext {
	var p = new(ProjectionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

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
	this := p
	_ = this

	localctx = NewProjectionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, CsqlParserRULE_projection)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(203)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == CsqlParserK_DISTINCT {
		{
			p.SetState(202)
			p.Distinct()
		}

	}
	{
		p.SetState(205)
		p.ProjectionField()
	}
	p.SetState(210)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == CsqlParserT__15 {
		{
			p.SetState(206)
			p.Match(CsqlParserT__15)
		}
		{
			p.SetState(207)
			p.ProjectionField()
		}

		p.SetState(212)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IProjectionFieldContext is an interface to support dynamic dispatch.
type IProjectionFieldContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsProjectionFieldContext differentiates from other interfaces.
	IsProjectionFieldContext()
}

type ProjectionFieldContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProjectionFieldContext() *ProjectionFieldContext {
	var p = new(ProjectionFieldContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CsqlParserRULE_projectionField
	return p
}

func (*ProjectionFieldContext) IsProjectionFieldContext() {}

func NewProjectionFieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProjectionFieldContext {
	var p = new(ProjectionFieldContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

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
	this := p
	_ = this

	localctx = NewProjectionFieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, CsqlParserRULE_projectionField)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(215)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(213)
			p.ProjectionFieldName()
		}

	case 2:
		{
			p.SetState(214)
			p.valueExpr(0)
		}

	}
	p.SetState(221)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == CsqlParserK_AS || _la == CsqlParserIDENTIFIER {
		p.SetState(218)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == CsqlParserK_AS {
			{
				p.SetState(217)
				p.Match(CsqlParserK_AS)
			}

		}
		{
			p.SetState(220)
			p.Alias()
		}

	}

	return localctx
}

// IProjectionFieldNameContext is an interface to support dynamic dispatch.
type IProjectionFieldNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsProjectionFieldNameContext differentiates from other interfaces.
	IsProjectionFieldNameContext()
}

type ProjectionFieldNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProjectionFieldNameContext() *ProjectionFieldNameContext {
	var p = new(ProjectionFieldNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CsqlParserRULE_projectionFieldName
	return p
}

func (*ProjectionFieldNameContext) IsProjectionFieldNameContext() {}

func NewProjectionFieldNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProjectionFieldNameContext {
	var p = new(ProjectionFieldNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

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
	this := p
	_ = this

	localctx = NewProjectionFieldNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, CsqlParserRULE_projectionFieldName)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(226)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 18, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(223)
			p.Qualifier()
		}
		{
			p.SetState(224)
			p.Match(CsqlParserT__17)
		}

	}
	{
		p.SetState(228)
		p.FieldName()
	}

	return localctx
}

// IFieldNameContext is an interface to support dynamic dispatch.
type IFieldNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFieldNameContext differentiates from other interfaces.
	IsFieldNameContext()
}

type FieldNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldNameContext() *FieldNameContext {
	var p = new(FieldNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CsqlParserRULE_fieldName
	return p
}

func (*FieldNameContext) IsFieldNameContext() {}

func NewFieldNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldNameContext {
	var p = new(FieldNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

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
	this := p
	_ = this

	localctx = NewFieldNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, CsqlParserRULE_fieldName)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(230)
		_la = p.GetTokenStream().LA(1)

		if !(_la == CsqlParserT__10 || _la == CsqlParserIDENTIFIER) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IInnerJoinContext is an interface to support dynamic dispatch.
type IInnerJoinContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInnerJoinContext differentiates from other interfaces.
	IsInnerJoinContext()
}

type InnerJoinContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInnerJoinContext() *InnerJoinContext {
	var p = new(InnerJoinContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CsqlParserRULE_innerJoin
	return p
}

func (*InnerJoinContext) IsInnerJoinContext() {}

func NewInnerJoinContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InnerJoinContext {
	var p = new(InnerJoinContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

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
	this := p
	_ = this

	localctx = NewInnerJoinContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, CsqlParserRULE_innerJoin)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(233)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == CsqlParserK_INNER {
		{
			p.SetState(232)
			p.Match(CsqlParserK_INNER)
		}

	}
	{
		p.SetState(235)
		p.Match(CsqlParserK_JOIN)
	}

	return localctx
}

// ILeftJoinContext is an interface to support dynamic dispatch.
type ILeftJoinContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLeftJoinContext differentiates from other interfaces.
	IsLeftJoinContext()
}

type LeftJoinContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLeftJoinContext() *LeftJoinContext {
	var p = new(LeftJoinContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CsqlParserRULE_leftJoin
	return p
}

func (*LeftJoinContext) IsLeftJoinContext() {}

func NewLeftJoinContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LeftJoinContext {
	var p = new(LeftJoinContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

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
	this := p
	_ = this

	localctx = NewLeftJoinContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, CsqlParserRULE_leftJoin)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(237)
		p.Match(CsqlParserK_LEFT)
	}
	p.SetState(239)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == CsqlParserK_OUTER {
		{
			p.SetState(238)
			p.Match(CsqlParserK_OUTER)
		}

	}
	{
		p.SetState(241)
		p.Match(CsqlParserK_JOIN)
	}

	return localctx
}

// IRightJoinContext is an interface to support dynamic dispatch.
type IRightJoinContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRightJoinContext differentiates from other interfaces.
	IsRightJoinContext()
}

type RightJoinContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRightJoinContext() *RightJoinContext {
	var p = new(RightJoinContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CsqlParserRULE_rightJoin
	return p
}

func (*RightJoinContext) IsRightJoinContext() {}

func NewRightJoinContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RightJoinContext {
	var p = new(RightJoinContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

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
	this := p
	_ = this

	localctx = NewRightJoinContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, CsqlParserRULE_rightJoin)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(243)
		p.Match(CsqlParserK_RIGHT)
	}
	p.SetState(245)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == CsqlParserK_OUTER {
		{
			p.SetState(244)
			p.Match(CsqlParserK_OUTER)
		}

	}
	{
		p.SetState(247)
		p.Match(CsqlParserK_JOIN)
	}

	return localctx
}

// IFullJoinContext is an interface to support dynamic dispatch.
type IFullJoinContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFullJoinContext differentiates from other interfaces.
	IsFullJoinContext()
}

type FullJoinContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFullJoinContext() *FullJoinContext {
	var p = new(FullJoinContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CsqlParserRULE_fullJoin
	return p
}

func (*FullJoinContext) IsFullJoinContext() {}

func NewFullJoinContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FullJoinContext {
	var p = new(FullJoinContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

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
	this := p
	_ = this

	localctx = NewFullJoinContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, CsqlParserRULE_fullJoin)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(249)
		p.Match(CsqlParserK_FULL)
	}
	p.SetState(251)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == CsqlParserK_OUTER {
		{
			p.SetState(250)
			p.Match(CsqlParserK_OUTER)
		}

	}
	{
		p.SetState(253)
		p.Match(CsqlParserK_JOIN)
	}

	return localctx
}

// ICrossJoinContext is an interface to support dynamic dispatch.
type ICrossJoinContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCrossJoinContext differentiates from other interfaces.
	IsCrossJoinContext()
}

type CrossJoinContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCrossJoinContext() *CrossJoinContext {
	var p = new(CrossJoinContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CsqlParserRULE_crossJoin
	return p
}

func (*CrossJoinContext) IsCrossJoinContext() {}

func NewCrossJoinContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CrossJoinContext {
	var p = new(CrossJoinContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

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
	this := p
	_ = this

	localctx = NewCrossJoinContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, CsqlParserRULE_crossJoin)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(255)
		p.Match(CsqlParserK_CROSS)
	}
	{
		p.SetState(256)
		p.Match(CsqlParserK_JOIN)
	}

	return localctx
}

// IConditionalJoinTypeContext is an interface to support dynamic dispatch.
type IConditionalJoinTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsConditionalJoinTypeContext differentiates from other interfaces.
	IsConditionalJoinTypeContext()
}

type ConditionalJoinTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConditionalJoinTypeContext() *ConditionalJoinTypeContext {
	var p = new(ConditionalJoinTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CsqlParserRULE_conditionalJoinType
	return p
}

func (*ConditionalJoinTypeContext) IsConditionalJoinTypeContext() {}

func NewConditionalJoinTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConditionalJoinTypeContext {
	var p = new(ConditionalJoinTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

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
	this := p
	_ = this

	localctx = NewConditionalJoinTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, CsqlParserRULE_conditionalJoinType)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(262)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case CsqlParserK_JOIN, CsqlParserK_INNER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(258)
			p.InnerJoin()
		}

	case CsqlParserK_LEFT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(259)
			p.LeftJoin()
		}

	case CsqlParserK_RIGHT:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(260)
			p.RightJoin()
		}

	case CsqlParserK_FULL:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(261)
			p.FullJoin()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IConditionalJoinTargetContext is an interface to support dynamic dispatch.
type IConditionalJoinTargetContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsConditionalJoinTargetContext differentiates from other interfaces.
	IsConditionalJoinTargetContext()
}

type ConditionalJoinTargetContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConditionalJoinTargetContext() *ConditionalJoinTargetContext {
	var p = new(ConditionalJoinTargetContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CsqlParserRULE_conditionalJoinTarget
	return p
}

func (*ConditionalJoinTargetContext) IsConditionalJoinTargetContext() {}

func NewConditionalJoinTargetContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConditionalJoinTargetContext {
	var p = new(ConditionalJoinTargetContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

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
	this := p
	_ = this

	localctx = NewConditionalJoinTargetContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, CsqlParserRULE_conditionalJoinTarget)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(264)
		p.dataSource(0)
	}
	{
		p.SetState(265)
		p.Match(CsqlParserK_ON)
	}
	{
		p.SetState(266)
		p.whereExpr(0)
	}

	return localctx
}

// IUnconditionalJoinTargetContext is an interface to support dynamic dispatch.
type IUnconditionalJoinTargetContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsUnconditionalJoinTargetContext differentiates from other interfaces.
	IsUnconditionalJoinTargetContext()
}

type UnconditionalJoinTargetContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnconditionalJoinTargetContext() *UnconditionalJoinTargetContext {
	var p = new(UnconditionalJoinTargetContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CsqlParserRULE_unconditionalJoinTarget
	return p
}

func (*UnconditionalJoinTargetContext) IsUnconditionalJoinTargetContext() {}

func NewUnconditionalJoinTargetContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UnconditionalJoinTargetContext {
	var p = new(UnconditionalJoinTargetContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

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
	this := p
	_ = this

	localctx = NewUnconditionalJoinTargetContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, CsqlParserRULE_unconditionalJoinTarget)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(268)
		p.dataSource(0)
	}

	return localctx
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
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDataSourceContext() *DataSourceContext {
	var p = new(DataSourceContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CsqlParserRULE_dataSource
	return p
}

func (*DataSourceContext) IsDataSourceContext() {}

func NewDataSourceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DataSourceContext {
	var p = new(DataSourceContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CsqlParserRULE_dataSource

	return p
}

func (s *DataSourceContext) GetParser() antlr.Parser { return s.parser }

func (s *DataSourceContext) CopyFrom(ctx *DataSourceContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *DataSourceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DataSourceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type DataSourceConditionalJoinContext struct {
	*DataSourceContext
}

func NewDataSourceConditionalJoinContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DataSourceConditionalJoinContext {
	var p = new(DataSourceConditionalJoinContext)

	p.DataSourceContext = NewEmptyDataSourceContext()
	p.parser = parser
	p.CopyFrom(ctx.(*DataSourceContext))

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
	*DataSourceContext
}

func NewDataSourceItemContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DataSourceItemContext {
	var p = new(DataSourceItemContext)

	p.DataSourceContext = NewEmptyDataSourceContext()
	p.parser = parser
	p.CopyFrom(ctx.(*DataSourceContext))

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
	*DataSourceContext
}

func NewDataSourceCrossJoinContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DataSourceCrossJoinContext {
	var p = new(DataSourceCrossJoinContext)

	p.DataSourceContext = NewEmptyDataSourceContext()
	p.parser = parser
	p.CopyFrom(ctx.(*DataSourceContext))

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
	*DataSourceContext
}

func NewDataSourceNamedContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DataSourceNamedContext {
	var p = new(DataSourceNamedContext)

	p.DataSourceContext = NewEmptyDataSourceContext()
	p.parser = parser
	p.CopyFrom(ctx.(*DataSourceContext))

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
	this := p
	_ = this

	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewDataSourceContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IDataSourceContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 44
	p.EnterRecursionRule(localctx, 44, CsqlParserRULE_dataSource, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(276)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case CsqlParserIDENTIFIER:
		localctx = NewDataSourceNamedContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(271)
			p.SourceName()
		}

	case CsqlParserT__14:
		localctx = NewDataSourceItemContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(272)
			p.Match(CsqlParserT__14)
		}
		{
			p.SetState(273)
			p.dataSource(0)
		}
		{
			p.SetState(274)
			p.Match(CsqlParserT__16)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(288)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 26, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(286)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 25, p.GetParserRuleContext()) {
			case 1:
				localctx = NewDataSourceConditionalJoinContext(p, NewDataSourceContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, CsqlParserRULE_dataSource)
				p.SetState(278)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(279)
					p.ConditionalJoinType()
				}
				{
					p.SetState(280)
					p.ConditionalJoinTarget()
				}

			case 2:
				localctx = NewDataSourceCrossJoinContext(p, NewDataSourceContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, CsqlParserRULE_dataSource)
				p.SetState(282)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				}
				{
					p.SetState(283)
					p.CrossJoin()
				}
				{
					p.SetState(284)
					p.UnconditionalJoinTarget()
				}

			}

		}
		p.SetState(290)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 26, p.GetParserRuleContext())
	}

	return localctx
}

// ISourcesContext is an interface to support dynamic dispatch.
type ISourcesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSourcesContext differentiates from other interfaces.
	IsSourcesContext()
}

type SourcesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySourcesContext() *SourcesContext {
	var p = new(SourcesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CsqlParserRULE_sources
	return p
}

func (*SourcesContext) IsSourcesContext() {}

func NewSourcesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SourcesContext {
	var p = new(SourcesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

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
	this := p
	_ = this

	localctx = NewSourcesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, CsqlParserRULE_sources)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(291)
		p.dataSource(0)
	}
	p.SetState(296)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == CsqlParserT__15 {
		{
			p.SetState(292)
			p.Match(CsqlParserT__15)
		}
		{
			p.SetState(293)
			p.dataSource(0)
		}

		p.SetState(298)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IUnionSelectsContext is an interface to support dynamic dispatch.
type IUnionSelectsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsUnionSelectsContext differentiates from other interfaces.
	IsUnionSelectsContext()
}

type UnionSelectsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnionSelectsContext() *UnionSelectsContext {
	var p = new(UnionSelectsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CsqlParserRULE_unionSelects
	return p
}

func (*UnionSelectsContext) IsUnionSelectsContext() {}

func NewUnionSelectsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UnionSelectsContext {
	var p = new(UnionSelectsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

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
	this := p
	_ = this

	localctx = NewUnionSelectsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, CsqlParserRULE_unionSelects)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(299)
		p.SelectStatement()
	}
	p.SetState(307)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 29, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(300)
				p.Match(CsqlParserK_UNION)
			}
			p.SetState(302)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			if _la == CsqlParserK_ALL {
				{
					p.SetState(301)
					p.Match(CsqlParserK_ALL)
				}

			}
			{
				p.SetState(304)
				p.UnionSelects()
			}

		}
		p.SetState(309)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 29, p.GetParserRuleContext())
	}

	return localctx
}

// ISelectStatementContext is an interface to support dynamic dispatch.
type ISelectStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSelectStatementContext differentiates from other interfaces.
	IsSelectStatementContext()
}

type SelectStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySelectStatementContext() *SelectStatementContext {
	var p = new(SelectStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CsqlParserRULE_selectStatement
	return p
}

func (*SelectStatementContext) IsSelectStatementContext() {}

func NewSelectStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SelectStatementContext {
	var p = new(SelectStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

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
	this := p
	_ = this

	localctx = NewSelectStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, CsqlParserRULE_selectStatement)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(310)
		p.Match(CsqlParserK_SELECT)
	}
	{
		p.SetState(311)
		p.Projection()
	}
	{
		p.SetState(312)
		p.Match(CsqlParserK_FROM)
	}
	{
		p.SetState(313)
		p.Sources()
	}
	p.SetState(315)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == CsqlParserK_WHERE {
		{
			p.SetState(314)
			p.Where()
		}

	}
	p.SetState(318)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == CsqlParserK_ORDER {
		{
			p.SetState(317)
			p.OrderBy()
		}

	}
	p.SetState(321)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == CsqlParserK_LIMIT {
		{
			p.SetState(320)
			p.Limit()
		}

	}

	return localctx
}

// ISignedNumberContext is an interface to support dynamic dispatch.
type ISignedNumberContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSignedNumberContext differentiates from other interfaces.
	IsSignedNumberContext()
}

type SignedNumberContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySignedNumberContext() *SignedNumberContext {
	var p = new(SignedNumberContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CsqlParserRULE_signedNumber
	return p
}

func (*SignedNumberContext) IsSignedNumberContext() {}

func NewSignedNumberContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SignedNumberContext {
	var p = new(SignedNumberContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

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
	this := p
	_ = this

	localctx = NewSignedNumberContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, CsqlParserRULE_signedNumber)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(324)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == CsqlParserT__8 || _la == CsqlParserT__9 {
		{
			p.SetState(323)
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
		p.SetState(326)
		p.Match(CsqlParserNUMERIC_LITERAL)
	}

	return localctx
}

// IStringValueContext is an interface to support dynamic dispatch.
type IStringValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStringValueContext differentiates from other interfaces.
	IsStringValueContext()
}

type StringValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStringValueContext() *StringValueContext {
	var p = new(StringValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CsqlParserRULE_stringValue
	return p
}

func (*StringValueContext) IsStringValueContext() {}

func NewStringValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StringValueContext {
	var p = new(StringValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

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
	this := p
	_ = this

	localctx = NewStringValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, CsqlParserRULE_stringValue)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(328)
		p.Match(CsqlParserSTRING_LITERAL)
	}

	return localctx
}

// INullValueContext is an interface to support dynamic dispatch.
type INullValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNullValueContext differentiates from other interfaces.
	IsNullValueContext()
}

type NullValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNullValueContext() *NullValueContext {
	var p = new(NullValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CsqlParserRULE_nullValue
	return p
}

func (*NullValueContext) IsNullValueContext() {}

func NewNullValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NullValueContext {
	var p = new(NullValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

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
	this := p
	_ = this

	localctx = NewNullValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, CsqlParserRULE_nullValue)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(330)
		p.Match(CsqlParserK_NULL)
	}

	return localctx
}

// ILiteralValueContext is an interface to support dynamic dispatch.
type ILiteralValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLiteralValueContext differentiates from other interfaces.
	IsLiteralValueContext()
}

type LiteralValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLiteralValueContext() *LiteralValueContext {
	var p = new(LiteralValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CsqlParserRULE_literalValue
	return p
}

func (*LiteralValueContext) IsLiteralValueContext() {}

func NewLiteralValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LiteralValueContext {
	var p = new(LiteralValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

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
	this := p
	_ = this

	localctx = NewLiteralValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, CsqlParserRULE_literalValue)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(335)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case CsqlParserT__8, CsqlParserT__9, CsqlParserNUMERIC_LITERAL:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(332)
			p.SignedNumber()
		}

	case CsqlParserSTRING_LITERAL:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(333)
			p.StringValue()
		}

	case CsqlParserK_NULL:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(334)
			p.NullValue()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IAliasContext is an interface to support dynamic dispatch.
type IAliasContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAliasContext differentiates from other interfaces.
	IsAliasContext()
}

type AliasContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAliasContext() *AliasContext {
	var p = new(AliasContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CsqlParserRULE_alias
	return p
}

func (*AliasContext) IsAliasContext() {}

func NewAliasContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AliasContext {
	var p = new(AliasContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

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
	this := p
	_ = this

	localctx = NewAliasContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, CsqlParserRULE_alias)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(337)
		p.Match(CsqlParserIDENTIFIER)
	}

	return localctx
}

// ISourceNameContext is an interface to support dynamic dispatch.
type ISourceNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSourceNameContext differentiates from other interfaces.
	IsSourceNameContext()
}

type SourceNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySourceNameContext() *SourceNameContext {
	var p = new(SourceNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CsqlParserRULE_sourceName
	return p
}

func (*SourceNameContext) IsSourceNameContext() {}

func NewSourceNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SourceNameContext {
	var p = new(SourceNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

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
	this := p
	_ = this

	localctx = NewSourceNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, CsqlParserRULE_sourceName)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(339)
		p.Name()
	}
	p.SetState(341)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 35, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(340)
			p.Alias()
		}

	}

	return localctx
}

// ICompoundNameContext is an interface to support dynamic dispatch.
type ICompoundNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCompoundNameContext differentiates from other interfaces.
	IsCompoundNameContext()
}

type CompoundNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCompoundNameContext() *CompoundNameContext {
	var p = new(CompoundNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CsqlParserRULE_compoundName
	return p
}

func (*CompoundNameContext) IsCompoundNameContext() {}

func NewCompoundNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CompoundNameContext {
	var p = new(CompoundNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

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
	this := p
	_ = this

	localctx = NewCompoundNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, CsqlParserRULE_compoundName)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(346)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 36, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(343)
			p.Qualifier()
		}
		{
			p.SetState(344)
			p.Match(CsqlParserT__17)
		}

	}
	{
		p.SetState(348)
		p.Name()
	}

	return localctx
}

// INameContext is an interface to support dynamic dispatch.
type INameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNameContext differentiates from other interfaces.
	IsNameContext()
}

type NameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNameContext() *NameContext {
	var p = new(NameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CsqlParserRULE_name
	return p
}

func (*NameContext) IsNameContext() {}

func NewNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NameContext {
	var p = new(NameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

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
	this := p
	_ = this

	localctx = NewNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, CsqlParserRULE_name)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(350)
		p.Match(CsqlParserIDENTIFIER)
	}

	return localctx
}

// IQualifierContext is an interface to support dynamic dispatch.
type IQualifierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsQualifierContext differentiates from other interfaces.
	IsQualifierContext()
}

type QualifierContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyQualifierContext() *QualifierContext {
	var p = new(QualifierContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CsqlParserRULE_qualifier
	return p
}

func (*QualifierContext) IsQualifierContext() {}

func NewQualifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *QualifierContext {
	var p = new(QualifierContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

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
	this := p
	_ = this

	localctx = NewQualifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, CsqlParserRULE_qualifier)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(352)
		p.Match(CsqlParserIDENTIFIER)
	}

	return localctx
}

// ILimitContext is an interface to support dynamic dispatch.
type ILimitContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLimitContext differentiates from other interfaces.
	IsLimitContext()
}

type LimitContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLimitContext() *LimitContext {
	var p = new(LimitContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CsqlParserRULE_limit
	return p
}

func (*LimitContext) IsLimitContext() {}

func NewLimitContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LimitContext {
	var p = new(LimitContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

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
	this := p
	_ = this

	localctx = NewLimitContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, CsqlParserRULE_limit)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(354)
		p.Match(CsqlParserK_LIMIT)
	}
	{
		p.SetState(355)
		p.LimitValue()
	}

	return localctx
}

// ILimitValueContext is an interface to support dynamic dispatch.
type ILimitValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLimitValueContext differentiates from other interfaces.
	IsLimitValueContext()
}

type LimitValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLimitValueContext() *LimitValueContext {
	var p = new(LimitValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CsqlParserRULE_limitValue
	return p
}

func (*LimitValueContext) IsLimitValueContext() {}

func NewLimitValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LimitValueContext {
	var p = new(LimitValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

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
	this := p
	_ = this

	localctx = NewLimitValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, CsqlParserRULE_limitValue)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(357)
		p.Match(CsqlParserNUMERIC_LITERAL)
	}

	return localctx
}

// IOrderByContext is an interface to support dynamic dispatch.
type IOrderByContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOrderByContext differentiates from other interfaces.
	IsOrderByContext()
}

type OrderByContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOrderByContext() *OrderByContext {
	var p = new(OrderByContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CsqlParserRULE_orderBy
	return p
}

func (*OrderByContext) IsOrderByContext() {}

func NewOrderByContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OrderByContext {
	var p = new(OrderByContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

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
	this := p
	_ = this

	localctx = NewOrderByContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, CsqlParserRULE_orderBy)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(359)
		p.Match(CsqlParserK_ORDER)
	}
	{
		p.SetState(360)
		p.Match(CsqlParserK_BY)
	}
	{
		p.SetState(361)
		p.OrderByField()
	}
	p.SetState(366)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == CsqlParserT__15 {
		{
			p.SetState(362)
			p.Match(CsqlParserT__15)
		}
		{
			p.SetState(363)
			p.OrderByField()
		}

		p.SetState(368)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IOrderByFieldContext is an interface to support dynamic dispatch.
type IOrderByFieldContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOrderByFieldContext differentiates from other interfaces.
	IsOrderByFieldContext()
}

type OrderByFieldContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOrderByFieldContext() *OrderByFieldContext {
	var p = new(OrderByFieldContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CsqlParserRULE_orderByField
	return p
}

func (*OrderByFieldContext) IsOrderByFieldContext() {}

func NewOrderByFieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OrderByFieldContext {
	var p = new(OrderByFieldContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

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
	this := p
	_ = this

	localctx = NewOrderByFieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 76, CsqlParserRULE_orderByField)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(371)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case CsqlParserIDENTIFIER:
		{
			p.SetState(369)
			p.CompoundName()
		}

	case CsqlParserNUMERIC_LITERAL:
		{
			p.SetState(370)
			p.FieldIndex()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.SetState(374)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == CsqlParserK_ASC || _la == CsqlParserK_DESC {
		{
			p.SetState(373)
			_la = p.GetTokenStream().LA(1)

			if !(_la == CsqlParserK_ASC || _la == CsqlParserK_DESC) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	}

	return localctx
}

// IFieldIndexContext is an interface to support dynamic dispatch.
type IFieldIndexContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFieldIndexContext differentiates from other interfaces.
	IsFieldIndexContext()
}

type FieldIndexContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldIndexContext() *FieldIndexContext {
	var p = new(FieldIndexContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CsqlParserRULE_fieldIndex
	return p
}

func (*FieldIndexContext) IsFieldIndexContext() {}

func NewFieldIndexContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldIndexContext {
	var p = new(FieldIndexContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

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
	this := p
	_ = this

	localctx = NewFieldIndexContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 78, CsqlParserRULE_fieldIndex)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(376)
		p.Match(CsqlParserNUMERIC_LITERAL)
	}

	return localctx
}

// IFunctionContext is an interface to support dynamic dispatch.
type IFunctionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunctionContext differentiates from other interfaces.
	IsFunctionContext()
}

type FunctionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionContext() *FunctionContext {
	var p = new(FunctionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CsqlParserRULE_function
	return p
}

func (*FunctionContext) IsFunctionContext() {}

func NewFunctionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionContext {
	var p = new(FunctionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

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
	this := p
	_ = this

	localctx = NewFunctionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 80, CsqlParserRULE_function)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(378)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&9221120237041090560) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
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

	case 22:
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
	this := p
	_ = this

	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 4)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *CsqlParser) WhereExpr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	this := p
	_ = this

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
	this := p
	_ = this

	switch predIndex {
	case 4:
		return p.Precpred(p.GetParserRuleContext(), 3)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
