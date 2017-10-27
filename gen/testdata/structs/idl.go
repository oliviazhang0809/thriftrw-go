// Code generated by thriftrw v1.9.0. DO NOT EDIT.
// @generated

package structs

import (
	"go.uber.org/thriftrw/gen/testdata/enums"
	"go.uber.org/thriftrw/thriftreflect"
)

// ThriftModule represents the IDL file used to generate this package.
var ThriftModule = &thriftreflect.ThriftModule{
	Name:     "structs",
	Package:  "go.uber.org/thriftrw/gen/testdata/structs",
	FilePath: "structs.thrift",
	SHA1:     "dc082cf16eb7cd106b0f0134604e12be0a0431a7",
	Includes: []*thriftreflect.ThriftModule{
		enums.ThriftModule,
	},
	Raw: rawIDL,
}

const rawIDL = "include \"./enums.thrift\"\n\nstruct EmptyStruct {}\n\n//////////////////////////////////////////////////////////////////////////////\n// Structs with primitives\n\n/**\n * A struct that contains primitive fields exclusively.\n *\n * All fields are required.\n */\nstruct PrimitiveRequiredStruct {\n    1: required bool boolField\n    2: required byte byteField\n    3: required i16 int16Field\n    4: required i32 int32Field\n    5: required i64 int64Field\n    6: required double doubleField\n    7: required string stringField\n    8: required binary binaryField\n}\n\n/**\n * A struct that contains primitive fields exclusively.\n *\n * All fields are optional.\n */\nstruct PrimitiveOptionalStruct {\n    1: optional bool boolField\n    2: optional byte byteField\n    3: optional i16 int16Field\n    4: optional i32 int32Field\n    5: optional i64 int64Field\n    6: optional double doubleField\n    7: optional string stringField\n    8: optional binary binaryField\n}\n\n//////////////////////////////////////////////////////////////////////////////\n// Nested structs (Required)\n\n/**\n * A point in 2D space.\n */\nstruct Point {\n    1: required double x\n    2: required double y\n}\n\n/**\n * Size of something.\n */\nstruct Size {\n    /**\n     * Width in pixels.\n     */\n    1: required double width\n    /** Height in pixels. */\n    2: required double height\n}\n\nstruct Frame {\n    1: required Point topLeft\n    2: required Size size\n}\n\nstruct Edge {\n    1: required Point startPoint\n    2: required Point endPoint\n}\n\n/**\n * A graph is comprised of zero or more edges.\n */\nstruct Graph {\n    /**\n     * List of edges in the graph.\n     *\n     * May be empty.\n     */\n    1: required list<Edge> edges\n}\n\n//////////////////////////////////////////////////////////////////////////////\n// Nested structs (Optional)\n\ntypedef i64 Timestamp (json.type = \"Date\")\ntypedef i64 (json.type = \"Long\") Long\n\nstruct I64TestStruct {\n    1: required Timestamp time\n    2: required Long longNumber\n    3: required i64 i64Field\n}\n\nstruct ContactInfo {\n    1: required string emailAddress\n}\n\nstruct User {\n    1: required string name\n    2: optional ContactInfo contact\n}\n\n//////////////////////////////////////////////////////////////////////////////\n// self-referential struct\n\ntypedef Node List\n\n/**\n * Node is linked list of values.\n * All values are 32-bit integers.\n */\nstruct Node {\n    1: required i32 value\n    2: optional List tail\n}\n\n//////////////////////////////////////////////////////////////////////////////\n// JSON tagged structs\n\nstruct Rename {\n    1: required string Default (go.tag = 'json:\"default\"')\n    2: required string camelCase (go.tag = 'json:\"snake_case\"')\n}\n\nstruct Omit {\n    1: required string serialized\n    2: required string hidden (go.tag = 'json:\"-\"')\n}\n\nstruct GoTags {\n        1: required string Foo (go.tag = 'json:\"-\" foo:\"bar\"')\n        2: optional string Bar (go.tag = 'bar:\"foo\"')\n        3: required string FooBar (go.tag = 'json:\"foobar,option1,option2\" bar:\"foo,option1\" foo:\"foobar\"')\n        4: required string FooBarWithSpace (go.tag = 'json:\"foobarWithSpace\" foo:\"foo bar foobar barfoo\"')\n        5: optional string FooBarWithOmitEmpty (go.tag = 'json:\"foobarWithOmitEmpty,omitempty\"')\n        6: required string FooBarWithRequired (go.tag = 'json:\"foobarWithRequired,required\"')\n}\n\n//////////////////////////////////////////////////////////////////////////////\n// Default values\n\nstruct DefaultsStruct {\n    1: required i32 requiredPrimitive = 100\n    2: optional i32 optionalPrimitive = 200\n\n    3: required enums.EnumDefault requiredEnum = enums.EnumDefault.Bar\n    4: optional enums.EnumDefault optionalEnum = 2\n\n    5: required list<string> requiredList = [\"hello\", \"world\"]\n    6: optional list<double> optionalList = [1, 2.0, 3]\n\n    7: required Frame requiredStruct = {\n        \"topLeft\": {\"x\": 1, \"y\": 2},\n        \"size\": {\"width\": 100, \"height\": 200},\n    }\n    8: optional Edge optionalStruct = {\n        \"startPoint\": {\"x\": 1, \"y\": 2},\n        \"endPoint\":   {\"x\": 3, \"y\": 4},\n    }\n}\n"
