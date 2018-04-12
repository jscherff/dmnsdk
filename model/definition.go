// Copyright 2018 John Scherff
//
// Licensed under the Apache License, version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package model

import (
	`bytes`
	`encoding/xml`
	`encoding/json`
	`io`
	`net/http`
	`os`
)

// ==============================================================================
// See https://docs.camunda.org/manual/7.4/reference/dmn11/decision-table/
// ==============================================================================

// DefinitionList is a collection of Decision Model and Notation (DMN) objects.
type DefinitionList []*DefinitionInfo

// Read unmarshals JSON from an io.Reader into an slice of objects.
func (this *DefinitionList) Read(r io.Reader) (error) {
	return readJson(this, r)
}

// ReadUrl unmarshals JSON from a url into an slice of objects.
func (this *DefinitionList) ReadUrl(u string) (error) {
	return readJsonFromUrl(this, u)
}

// ReadFile unmarshals JSON from a file into a slice of objects.
func (this *DefinitionList) ReadFile(f string) (error) {
	return readJsonFromFile(this, f)
}

// DMN is a Decision Model and Notation Object. It contains DMN metadata, the 
// raw DMN XML, and the DMN Definition as a hierarchy of objects corresponding
// to DMN XML elements.
type DefinitionInfo struct {
	Id			string		`json:"id"`
	Key			string		`json:"key"`
	Category		string		`json:"category"`
	Name			string		`json:"name"`
	Version			int		`json:"version"`
	Resource		string		`json:"resource"`
	DeploymentId		string		`json:"deploymentId"`
	TenantId		string		`json:"tenantId"`
	DecisionReqDefId	string		`json:"decisionRequirementsDefinitionId"`
	DecisionReqDefKey	string		`json:"decisionRequirementsDefinitionKey"`
	HistoryTtl		string		`json:"historyTimeToLive"`
}

// Read unmarshals JSON from an io.Reader into an object.
func (this *DefinitionInfo) Read(r io.Reader) (error) {
	return readJson(this, r)
}

// ReadUrl unmarshals JSON from a url into an object.
func (this *DefinitionInfo) ReadUrl(u string) (error) {
	return readJsonFromUrl(this, u)
}

// ReadFile unmarshals JSON from a file into an object.
func (this *DefinitionInfo) ReadFile(f string) (error) {
	return readJsonFromFile(this, f)
}

// DmnXml is the DMN XML document describing the Decision Definition.
type DmnXml struct {
	Id			string		`json:"id"`
	dmnXml			string		`json:"dmnXml`
}

// Read unmarshals JSON from an io.Reader into an object.
func (this *DmnXml) Read(r io.Reader) (error) {
	return readJson(this, r)
}

// ReadUrl unmarshals JSON from a url into an object.
func (this *DmnXml) ReadUrl(u string) (error) {
	return readJsonFromUrl(this, u)
}

// ReadFile unmarshals JSON from a file into an object.
func (this *DmnXml) ReadFile(f string) (error) {
	return readJsonFromFile(this, f)
}

// Definition is a Decision Model and Notation DefinitionInfo. 
type Definition struct {
	XMLName			xml.Name
	Xmlns			string		`xml:"xmlns,attr"`
	Id			string		`xml:"id,attr"`
	Name			string		`xml:"name,attr"`
	ExpressionLang		string		`xml:"expressionLanguage"`
	Namespace		string		`xml:"namespace,attr"`
	Decision		*Decision	`xml:"decision"`
}

// Read unmarshals XML from an io.Reader into an object.
func (this *Definition) Read(r io.Reader) (error) {
	return readXml(this, r)
}

// ReadString unmarshals XML from a string into an object.
func (this *Definition) ReadString(s string) (error) {
	return readXmlFromString(this, s)
}

// A DecisionTable is decision logic which can be depicted as a table in
// DMN 1.1. It consists of inputs, outputs and rules and is represented
// by a <decisionTable> element inside a <decision> element. 

// The name describes the decision for which the decision table provides
// the decision logic. It is set as the name attribute on the decision
// element. The id is the technical identifier of the decision. It is
// set in the id attribute on the decision element.

type Decision struct {
	XMLName			xml.Name
	Id			string		`xml:"id,attr"`
	Name			string		`xml:"name,attr"`
	DecisionTable		*DecisionTable	`xml:"decisionTable"`
}

type DecisionTable struct {
	XMLName			xml.Name
	Id			string		`xml:"id,attr"`
	HitPolicy		string		`xml:"hitPolicy,attr"`
	Inputs			[]*Input	`xml:"input"`
	Outputs			[]*Output	`xml:"output"`
	Rules			[]*Rule		`xml:"rule"`
}

// A decision table can have one or more inputs, also called input
// clauses. An input clause defines the id, label, expression and type
// of a decision table input. An input clause is represented by an input
// element inside a decisionTable XML element.

// The input id is an unique identifier of the decision table input.
// It is used by the Camunda BPMN platform to reference the input in the
// history of evaluated decisions. Therefore it is required by the Camunda
// DMN engine. It is set as the id attribute of the input XML element.

// An input label is a short description of the input. It is set on the
// input XML element in the label attribute. Note that the label is not
// required but recommended since it helps to understand the decision.

type Input struct {
	XMLName			xml.Name
	Id			string		`xml:"id,attr"`
	Label			string		`xml:"label,attr"`
	InputExpressions	[]*InputExpression `xml:"inputExpression"`
}

// An input expression specifies how the value of the input clause is
// generated. It is an expression which will be evaluated by the DMN
// engine. It is usually simple and references a variable which is
// available during the evaluation. The expression is set inside a text
// element that is a child of the inputExpression XML element.

// The type of the input clause can be specified by the typeRef attribute
// on the inputExpression XML element. After the input expression is
// evaluated by the DMN engine it converts the result to the specified
// type.

// The expression language of the input expression can be specified by
// the expressionLanguage attribute on the inputExpression XML element.
// If no expression language is set then the global expression language
// is used which is set on the definitions XML element. In case no global
// expression language is set, the default expression language is used
// instead. The default expression language for input expressions is JUEL.

type InputExpression struct {
	XMLName			xml.Name
	Id			string		`xml:"id,attr"`
	TypeRef			string		`xml:"typeRef,attr"`
	ExpressionLang		string		`xml:"expressionLanguage,attr"`
	Text			string		`xml:"text"`
}

// A decision table can have one or more output, also called output clauses.
// An output clause defines the id, label, name and type of a decision table
// output. An output clause is represented by an output element inside a
// decisionTable XML element.

// The output id is an unique identifier of the decision table output. It
// is used by the Camunda BPMN platform to reference the output in the
// history of evaluated decisions. Therefore it is required by the Camunda
// DMN engine. It is set as the id attribute of the output XML element.

// An output label is a short description of the output. It is set on the
// output XML element in the label attribute. Note that the label is not
// required but recommended since it helps to understand the decision.

// The name of the output is used to reference the value of the output in
// the decision table result. It is specified by the name attribute on the
// output XML element. If the decision table has more than one output then
// all outputs must have an unique name.

// The type of the output clause can be specified by the typeRef attribute
// on the output XML element. After an output entry is evaluated by the DMN
// engine it converts the result to the specified type. Note that the type
// is not required but recommended since it provides a type safety of the
// output values. Additionally, the type can be used to transform the output
// value into another type. For example, transform the output value 80% of
// type String into a Double using a custom data type.

type Output struct {
	XMLName			xml.Name
	Id			string		`xml:"id,attr"`
	Label			string		`xml:"label,attr"`
	Name			string		`xml:"name,attr"`
	TypeRef			string		`xml:"typeRef,attr"`
}

// A decision table can have one or more rules. Each rule contains input
// and output entries. The input entries are the condition and the output
// entries the conclusion of the rule. If each input entry (condition) is
// satisfied then the rule is satisfied and the decision result contains
// the output entries (conclusion) of this rule. A rule is represented by
// a rule element inside a decisionTable XML element.

type Rule struct {
	XMLName			xml.Name
	Id			string		`xml:"id,attr"`
	InputEntries		[]*InputEntry	`xml:"inputEntry"`
	OutputEntries		[]*OutputEntry	`xml:"outputEntry"`
}

// A rule can have one or more input entries which are the conditions of
// the rule. Each input entry contains an expression in a text element as
// child of an inputEntry XML element. The input entry is satisfied when
// the evaluated expression returns true. In case an input entry is
// irrelevant for a rule, the expression is empty which is always satisfied.

// The expression language of the input entry can be specified by the
// expressionLanguage attribute on the inputEntry XML element. If no
// expression language is set then the global expression language is
// used which is set on the definitions XML element. In case no global
// expression language is set, the default expression language is used
// instead. 

type InputEntry struct {
	XMLName			xml.Name
	Id			string		`xml:"id,attr"`
	ExpressionLang		string		`xml:"expressionLanguage,attr"`
	Text			string		`xml:"text"`
}

// A rule can have one or more output entries which are the conclusions
// of the rule. Each output entry contains an expression in a text element
// as child of an outputEntry XML element. If the output entry is empty then
// the output is ignored and not part of the decision table result.

// The expression language of the expression can be specified by the
// expressionLanguage attribute on the outputEntry XML element. If no
// expression language is set then the global expression language is
// used which is set on the definitions XML element. In case no global
// expression language is set, the default expression language is used
// instead. 

// A rule can be annotated with a description that provides additional
// information. The description text is set inside the description XML
// element.

type OutputEntry struct {
	XMLName			xml.Name
	Id			string		`xml:"id,attr"`
	ExpressionLang		string		`xml:"expressionLanguage,attr"`
	Description		string		`xml:"description"`
	Text			string		`xml:"text"`
}

// readJson is a helper function for package/object methods.
func readJson(t interface{}, r io.Reader) (error) {
	return json.NewDecoder(r).Decode(&t)
}

// readXml is a helper function for package/object methods.
func readXml(t interface{}, r io.Reader) (error) {
	return xml.NewDecoder(r).Decode(&t)
}

// readXmlFromString is a helper function for package/object methods.
func readXmlFromString(t interface{}, s string) (error) {
	return readXml(t, bytes.NewBufferString(s))
}

// readJsonFromUrl is a helper function for package/object methods.
func readJsonFromUrl(t interface{}, u string) (error) {

	if resp, err := http.Get(u); err != nil {
		return err
	} else {
		defer resp.Body.Close()
		return readJson(t, resp.Body)
	}
}

// readJsonFromFile is a helper function for package/object methods.
func readJsonFromFile(t interface{}, f string) (error) {

        if fh, err := os.Open(f); err != nil {
                return err
	} else {
                defer fh.Close()
		return readJson(t, fh)
        }
}