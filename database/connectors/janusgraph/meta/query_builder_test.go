package meta

import (
	"fmt"
	"strings"
	"testing"

	"github.com/creativesoftwarefdn/weaviate/database/connectors/janusgraph/state"
	"github.com/creativesoftwarefdn/weaviate/database/schema"
	"github.com/creativesoftwarefdn/weaviate/database/schema/kind"
	gm "github.com/creativesoftwarefdn/weaviate/graphqlapi/local/getmeta"
	"github.com/creativesoftwarefdn/weaviate/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// This file contains only a single test to verify that combining multiple
// props together works as intended and helpers for other tests. Each
// individual property type however, is also extensively tested, please see the
// test files for individual props with the format query_builder_<type>_test.go

func Test_QueryBuilder_MultipleProps(t *testing.T) {
	tests := testCases{
		testCase{
			name: "with multiple props",
			inputProps: []gm.MetaProperty{
				gm.MetaProperty{
					Name: "isCapital",
					StatisticalAnalyses: []gm.StatisticalAnalysis{
						gm.Count, gm.TotalTrue, gm.TotalFalse, gm.PercentageTrue, gm.PercentageFalse,
					},
				},
				gm.MetaProperty{
					Name: "population",
					StatisticalAnalyses: []gm.StatisticalAnalysis{
						gm.Average, gm.Sum, gm.Highest, gm.Lowest, gm.Count,
					},
				},
			},
			expectedQuery: `
				.union(
					union(
						has("isCapital").count()
							.as("count").project("count").by(select("count")),
						groupCount().by("isCapital")
							.as("boolGroupCount").project("boolGroupCount").by(select("boolGroupCount"))
					)
						.as("isCapital").project("isCapital").by(select("isCapital")),
					aggregate("aggregation").by("population").cap("aggregation").limit(1)
						.as("average", "sum", "highest", "lowest", "count")
						.select("average", "sum", "highest", "lowest", "count")
						.by(mean(local)).by(sum(local)).by(max(local)).by(min(local)).by(count(local))
						.as("population").project("population").by(select("population"))
				)
			`,
		},
	}

	tests.AssertQuery(t, nil)

}

type fakeNameSource struct{}

func (f *fakeNameSource) GetMappedPropertyName(className schema.ClassName,
	propName schema.PropertyName) state.MappedPropertyName {
	switch propName {
	case schema.PropertyName("inCountry"):
		return "prop_15"
	}
	return state.MappedPropertyName("prop_20")
}

func (f *fakeNameSource) GetMappedClassName(className schema.ClassName) state.MappedClassName {
	return state.MappedClassName("class_18")
}

type fakeTypeSource struct{}

func (f *fakeTypeSource) GetProperty(kind kind.Kind, className schema.ClassName,
	propName schema.PropertyName) (error, *models.SemanticSchemaClassProperty) {

	switch propName {
	case "isCapital":
		return nil, &models.SemanticSchemaClassProperty{AtDataType: []string{"bool"}}
	case "population":
		return nil, &models.SemanticSchemaClassProperty{AtDataType: []string{"int"}}
	case "area":
		return nil, &models.SemanticSchemaClassProperty{AtDataType: []string{"number"}}
	}

	return fmt.Errorf("fake type source does not have an implementation for prop '%s'", propName), nil
}

func (f *fakeTypeSource) FindPropertyDataType(dataType []string) (schema.PropertyDataType, error) {
	switch dataType[0] {
	case "bool":
		return &fakeDataType{dataType: schema.DataTypeBoolean}, nil
	case "int":
		return &fakeDataType{dataType: schema.DataTypeInt}, nil
	case "number":
		return &fakeDataType{dataType: schema.DataTypeNumber}, nil
	}

	return nil, fmt.Errorf("fake type source does not have an implementation for dataType '%v'", dataType)
}

type fakeDataType struct {
	dataType schema.DataType
}

func (p *fakeDataType) Kind() schema.PropertyKind {
	panic("not implemented")
}

func (p *fakeDataType) IsPrimitive() bool {
	return true
}

func (p *fakeDataType) AsPrimitive() schema.DataType {
	return p.dataType
}

func (p *fakeDataType) IsReference() bool {
	return false
}

func (p *fakeDataType) Classes() []schema.ClassName {
	panic("not implemented")
}

func (p *fakeDataType) ContainsClass(needle schema.ClassName) bool {
	panic("not implemented")
}

type testCase struct {
	name          string
	inputProps    []gm.MetaProperty
	expectedQuery string
}

type testCases []testCase

func (tests testCases) AssertQuery(t *testing.T, nameSource nameSource) {
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			params := &gm.Params{
				Properties: test.inputProps,
			}
			query, err := NewQuery(params, nameSource, &fakeTypeSource{}).String()
			require.Nil(t, err, "should not error")
			assert.Equal(t, stripAll(test.expectedQuery), stripAll(query), "should match the query")
		})
	}
}

func stripAll(input string) string {
	input = strings.Replace(input, " ", "", -1)
	input = strings.Replace(input, "\t", "", -1)
	input = strings.Replace(input, "\n", "", -1)
	return input
}