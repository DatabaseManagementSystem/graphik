// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
	"time"
)

type AggFilter struct {
	Filter    *Filter `json:"filter"`
	Aggregate string  `json:"aggregate"`
	Field     *string `json:"field"`
}

type Authorizer struct {
	Name       string `json:"name"`
	Expression string `json:"expression"`
}

type AuthorizerInput struct {
	Name       string `json:"name"`
	Expression string `json:"expression"`
}

type Authorizers struct {
	Authorizers []*Authorizer `json:"authorizers"`
}

type AuthorizersInput struct {
	Authorizers []*AuthorizerInput `json:"authorizers"`
}

type CFilter struct {
	DocRef     *RefInput `json:"doc_ref"`
	Gtype      string    `json:"gtype"`
	Expression *string   `json:"expression"`
	Limit      int       `json:"limit"`
	Sort       *string   `json:"sort"`
	Seek       *string   `json:"seek"`
	Reverse    *bool     `json:"reverse"`
}

type ChanFilter struct {
	Channel    string  `json:"channel"`
	Expression *string `json:"expression"`
}

type Connection struct {
	Ref        *Ref                   `json:"ref"`
	Attributes map[string]interface{} `json:"attributes"`
	Directed   bool                   `json:"directed"`
	From       *Ref                   `json:"from"`
	To         *Ref                   `json:"to"`
}

type ConnectionConstructor struct {
	Ref        *RefConstructor        `json:"ref"`
	Directed   bool                   `json:"directed"`
	Attributes map[string]interface{} `json:"attributes"`
	From       *RefInput              `json:"from"`
	To         *RefInput              `json:"to"`
}

type Connections struct {
	Connections []*Connection `json:"connections"`
	SeekNext    string        `json:"seek_next"`
}

type Doc struct {
	Ref        *Ref                   `json:"ref"`
	Attributes map[string]interface{} `json:"attributes"`
}

type DocConstructor struct {
	Ref        *RefConstructor        `json:"ref"`
	Attributes map[string]interface{} `json:"attributes"`
}

type Docs struct {
	Docs     []*Doc `json:"docs"`
	SeekNext string `json:"seek_next"`
}

type EFilter struct {
	Filter     *Filter                `json:"filter"`
	Attributes map[string]interface{} `json:"attributes"`
}

type Edit struct {
	Ref        *RefInput              `json:"ref"`
	Attributes map[string]interface{} `json:"attributes"`
}

type ExprFilter struct {
	Expression *string `json:"expression"`
}

type Filter struct {
	Gtype      string  `json:"gtype"`
	Expression *string `json:"expression"`
	Limit      int     `json:"limit"`
	Sort       *string `json:"sort"`
	Seek       *string `json:"seek"`
	Reverse    *bool   `json:"reverse"`
	Index      *string `json:"index"`
}

type Index struct {
	Name        string `json:"name"`
	Gtype       string `json:"gtype"`
	Expression  string `json:"expression"`
	Docs        *bool  `json:"docs"`
	Connections *bool  `json:"connections"`
}

type IndexInput struct {
	Name        string `json:"name"`
	Gtype       string `json:"gtype"`
	Expression  string `json:"expression"`
	Docs        *bool  `json:"docs"`
	Connections *bool  `json:"connections"`
}

type Indexes struct {
	Indexes []*Index `json:"indexes"`
}

type IndexesInput struct {
	Indexes []*IndexInput `json:"indexes"`
}

type Message struct {
	Channel   string                 `json:"channel"`
	Data      map[string]interface{} `json:"data"`
	Sender    *Ref                   `json:"sender"`
	Timestamp time.Time              `json:"timestamp"`
}

type OutboundMessage struct {
	Channel string                 `json:"channel"`
	Data    map[string]interface{} `json:"data"`
}

type Pong struct {
	Message string `json:"message"`
}

type Ref struct {
	Gtype string `json:"gtype"`
	Gid   string `json:"gid"`
}

type RefConstructor struct {
	Gtype string  `json:"gtype"`
	Gid   *string `json:"gid"`
}

type RefInput struct {
	Gtype string `json:"gtype"`
	Gid   string `json:"gid"`
}

type Refs struct {
	Refs []*Ref `json:"refs"`
}

type SConnectFilter struct {
	Filter     *Filter                `json:"filter"`
	Gtype      string                 `json:"gtype"`
	Attributes map[string]interface{} `json:"attributes"`
	Directed   bool                   `json:"directed"`
	From       *RefInput              `json:"from"`
}

type Schema struct {
	ConnectionTypes []string        `json:"connection_types"`
	DocTypes        []string        `json:"doc_types"`
	Authorizers     *Authorizers    `json:"authorizers"`
	Validators      *TypeValidators `json:"validators"`
	Indexes         *Indexes        `json:"indexes"`
}

type TFilter struct {
	Root                 *RefInput  `json:"root"`
	DocExpression        *string    `json:"doc_expression"`
	ConnectionExpression *string    `json:"connection_expression"`
	Limit                int        `json:"limit"`
	Sort                 *string    `json:"sort"`
	Reverse              *bool      `json:"reverse"`
	Algorithm            *Algorithm `json:"algorithm"`
}

type Traversal struct {
	Doc           *Doc   `json:"doc"`
	TraversalPath []*Ref `json:"traversal_path"`
}

type Traversals struct {
	Traversals []*Traversal `json:"traversals"`
}

type TypeValidator struct {
	Name        string `json:"name"`
	Gtype       string `json:"gtype"`
	Expression  string `json:"expression"`
	Docs        *bool  `json:"docs"`
	Connections *bool  `json:"connections"`
}

type TypeValidatorInput struct {
	Name        string `json:"name"`
	Gtype       string `json:"gtype"`
	Expression  string `json:"expression"`
	Docs        *bool  `json:"docs"`
	Connections *bool  `json:"connections"`
}

type TypeValidators struct {
	Validators []*TypeValidator `json:"validators"`
}

type TypeValidatorsInput struct {
	Validators []*TypeValidatorInput `json:"validators"`
}

type Algorithm string

const (
	AlgorithmBfs Algorithm = "BFS"
	AlgorithmDfs Algorithm = "DFS"
)

var AllAlgorithm = []Algorithm{
	AlgorithmBfs,
	AlgorithmDfs,
}

func (e Algorithm) IsValid() bool {
	switch e {
	case AlgorithmBfs, AlgorithmDfs:
		return true
	}
	return false
}

func (e Algorithm) String() string {
	return string(e)
}

func (e *Algorithm) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Algorithm(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Algorithm", str)
	}
	return nil
}

func (e Algorithm) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
