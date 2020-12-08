package apipb

import (
	"fmt"
	"gonum.org/v1/gonum/floats"
	"google.golang.org/protobuf/types/known/structpb"
	"sort"
	"strings"
)

const (
	Any = "*"
)

func NewStruct(data map[string]interface{}) *structpb.Struct {
	x, _ := structpb.NewStruct(data)
	return x
}

type Mapper interface {
	AsMap() map[string]interface{}
}

func (m *Ref) AsMap() map[string]interface{} {
	if m == nil {
		return map[string]interface{}{}
	}
	return map[string]interface{}{
		"gid":   m.GetGid(),
		"gtype": m.GetGtype(),
	}
}

func (n *Doc) AsMap() map[string]interface{} {
	if n == nil {
		return map[string]interface{}{}
	}
	return map[string]interface{}{
		"ref":        n.GetRef().AsMap(),
		"attributes": n.GetAttributes().AsMap(),
	}
}

func (n *Connection) AsMap() map[string]interface{} {
	if n == nil {
		return map[string]interface{}{}
	}
	return map[string]interface{}{
		"ref":        n.GetRef().AsMap(),
		"attributes": n.GetAttributes().AsMap(),
		"directed":   n.GetDirected(),
		"from":       n.GetFrom().AsMap(),
		"to":         n.GetTo().AsMap(),
	}
}

func (n *Message) AsMap() map[string]interface{} {
	if n == nil {
		return map[string]interface{}{}
	}
	return map[string]interface{}{
		"channel":   n.GetChannel(),
		"sender":    n.GetSender().AsMap(),
		"data":      n.GetData().AsMap(),
		"timestamp": n.GetTimestamp(),
	}
}

func (p *Refs) AsMap() map[string]interface{} {
	if p == nil {
		return map[string]interface{}{}
	}
	var refs []interface{}
	for _, ref := range p.GetRefs() {
		refs = append(refs, ref.AsMap())
	}
	return map[string]interface{}{
		"refs": refs,
	}
}

func (n *Filter) AsMap() map[string]interface{} {
	if n == nil {
		return map[string]interface{}{}
	}
	return map[string]interface{}{
		"gtype":      n.GetGtype(),
		"expression": n.GetExpression(),
		"limit":      n.GetLimit(),
		"sort":       n.GetSort(),
		"reverse":    n.GetReverse(),
		"seek":       n.GetSeek(),
	}
}

func (n *CFilter) AsMap() map[string]interface{} {
	if n == nil {
		return map[string]interface{}{}
	}
	return map[string]interface{}{
		"gtype":      n.GetGtype(),
		"expression": n.GetExpression(),
		"doc_ref":    n.GetDocRef(),
		"limit":      n.GetLimit(),
		"sort":       n.GetSort(),
		"reverse":    n.GetReverse(),
		"seek":       n.GetSeek(),
	}
}

func (n *ChanFilter) AsMap() map[string]interface{} {
	if n == nil {
		return map[string]interface{}{}
	}
	return map[string]interface{}{
		"channel":    n.GetChannel(),
		"expression": n.GetExpression(),
	}
}

func (n *ExprFilter) AsMap() map[string]interface{} {
	if n == nil {
		return map[string]interface{}{}
	}
	return map[string]interface{}{
		"expression": n.GetExpression(),
	}
}

func (p *Edit) AsMap() map[string]interface{} {
	if p == nil {
		return map[string]interface{}{}
	}
	return map[string]interface{}{
		"ref":        p.GetRef(),
		"attributes": p.GetAttributes(),
	}
}

func (n *EFilter) AsMap() map[string]interface{} {
	if n == nil {
		return map[string]interface{}{}
	}
	return map[string]interface{}{
		"patch":  n.GetAttributes().AsMap(),
		"filter": n.GetFilter().AsMap(),
	}
}

func (e *RefConstructor) AsMap() map[string]interface{} {
	if e == nil {
		return map[string]interface{}{}
	}
	return map[string]interface{}{
		"gtype": e.GetGtype(),
		"gid":   e.GetGid(),
	}
}

func (e *ConnectionConstructor) AsMap() map[string]interface{} {
	if e == nil {
		return map[string]interface{}{}
	}
	return map[string]interface{}{
		"ref":        e.GetRef().AsMap(),
		"attributes": e.GetAttributes().AsMap(),
		"directed":   e.GetDirected(),
		"from":       e.GetFrom().AsMap(),
		"to":         e.GetTo().AsMap(),
	}
}

func (e *DocConstructor) AsMap() map[string]interface{} {
	if e == nil {
		return map[string]interface{}{}
	}
	return map[string]interface{}{
		"ref":        e.GetRef().AsMap(),
		"attributes": e.GetAttributes().AsMap(),
	}
}

func (o *OutboundMessage) AsMap() map[string]interface{} {
	if o == nil {
		return map[string]interface{}{}
	}
	return map[string]interface{}{
		"channel": o.GetChannel(),
		"data":    o.GetData().AsMap(),
	}
}

func (e *Connections) AsMap() map[string]interface{} {
	var connections []interface{}
	for _, connection := range e.GetConnections() {
		connections = append(connections, connection.AsMap())
	}
	return map[string]interface{}{
		"connections": connections,
	}
}

func (e *Docs) AsMap() map[string]interface{} {
	var docs []interface{}
	for _, connection := range e.GetDocs() {
		docs = append(docs, connection.AsMap())
	}
	return map[string]interface{}{
		"docs": docs,
	}
}

func (g *Graph) AsMap() map[string]interface{} {
	return map[string]interface{}{
		"connections": g.GetConnections().AsMap(),
		"docs":        g.GetDocs().AsMap(),
	}
}

func (p *Refs) Sort(field string) {
	if p == nil {
		return
	}
	switch field {
	case "gid":
		sort.Slice(p.GetRefs(), func(i, j int) bool {
			return p.GetRefs()[i].GetGid() < p.GetRefs()[j].GetGid()
		})
	default:
		sort.Slice(p.GetRefs(), func(i, j int) bool {
			return p.GetRefs()[i].GetGtype() < p.GetRefs()[j].GetGtype()
		})
	}
}

func (n *Docs) Sort(field string) {
	if n == nil {
		return
	}
	switch {
	case field == "ref.gid":
		sort.Slice(n.GetDocs(), func(i, j int) bool {
			return n.GetDocs()[i].GetRef().GetGid() < n.GetDocs()[j].GetRef().GetGid()
		})
	case field == "ref.gtype":
		sort.Slice(n.GetDocs(), func(i, j int) bool {
			return n.GetDocs()[i].GetRef().GetGtype() < n.GetDocs()[j].GetRef().GetGtype()
		})
	case strings.Contains(field, "attributes."):
		split := strings.Split(field, "attributes.")
		if len(split) == 2 {
			key := split[1]
			sort.Slice(n.GetDocs(), func(i, j int) bool {
				fields := n.GetDocs()[i].GetAttributes().GetFields()
				if fields == nil {
					return true
				}
				if fields[key] == nil {
					return true
				}
				switch n.GetDocs()[i].GetAttributes().GetFields()[key].GetKind() {
				case &structpb.Value_NumberValue{}:
					return n.GetDocs()[i].GetAttributes().GetFields()[key].GetNumberValue() < n.GetDocs()[j].GetAttributes().GetFields()[key].GetNumberValue()
				case &structpb.Value_StringValue{}:
					return n.GetDocs()[i].GetAttributes().GetFields()[key].GetStringValue() < n.GetDocs()[j].GetAttributes().GetFields()[key].GetStringValue()
				default:
					return fmt.Sprint(n.GetDocs()[i].GetAttributes().GetFields()[key].AsInterface()) < fmt.Sprint(n.GetDocs()[j].GetAttributes().GetFields()[key].AsInterface())
				}
			})
		}
	}
}

func (e *Connections) Sort(field string) {
	if e == nil {
		return
	}
	switch {
	case field == "ref.gid":
		sort.Slice(e.GetConnections(), func(i, j int) bool {
			return e.GetConnections()[i].GetRef().GetGid() < e.GetConnections()[j].GetRef().GetGid()
		})
	case field == "ref.gtype":
		sort.Slice(e.GetConnections(), func(i, j int) bool {
			return e.GetConnections()[i].GetRef().GetGtype() < e.GetConnections()[j].GetRef().GetGtype()
		})
	case strings.Contains(field, "attributes."):
		split := strings.Split(field, "attributes.")
		if len(split) == 2 {
			key := split[1]
			sort.Slice(e.GetConnections(), func(i, j int) bool {
				fields := e.GetConnections()[i].GetAttributes().GetFields()
				if fields == nil {
					return true
				}
				if fields[key] == nil {
					return true
				}
				switch e.GetConnections()[i].GetAttributes().GetFields()[key].GetKind() {
				case &structpb.Value_NumberValue{}:
					return e.GetConnections()[i].GetAttributes().GetFields()[key].GetNumberValue() < e.GetConnections()[j].GetAttributes().GetFields()[key].GetNumberValue()
				case &structpb.Value_StringValue{}:
					return e.GetConnections()[i].GetAttributes().GetFields()[key].GetStringValue() < e.GetConnections()[j].GetAttributes().GetFields()[key].GetStringValue()
				default:
					return fmt.Sprint(e.GetConnections()[i].GetAttributes().GetFields()[key].AsInterface()) < fmt.Sprint(e.GetConnections()[j].GetAttributes().GetFields()[key].AsInterface())
				}
			})
		}
	}
}

func (d *Docs) Range(fn func(d *Doc) bool) {
	for _, d := range d.GetDocs() {
		if !fn(d) {
			return
		}
	}
}

func (d *Connections) Range(fn func(d *Connection) bool) {
	for _, d := range d.GetConnections() {
		if !fn(d) {
			return
		}
	}
}

func (d *Docs) Aggregate(aggregate string, field string) float64 {
	if aggregate == "count" {
		return float64(len(d.GetDocs()))
	}
	var values []float64
	d.Range(func(d *Doc) bool {
		switch {
		case strings.Contains(field, "attributes."):
			if fields := d.GetAttributes().GetFields(); fields != nil {
				split := strings.Split(field, "attributes.")
				if len(split) == 2 {
					key := split[1]
					if val := fields[key]; val != nil {
						values = append(values, val.GetNumberValue())
					}
				}
			}
		}
		return true
	})
	if len(values) == 0 {
		return 0
	}
	switch aggregate {
	case "sum":
		return floats.Sum(values)
	case "min":
		return floats.Min(values)
	case "max":
		return floats.Max(values)
	case "prod":
		return floats.Prod(values)
	case "avg":
		return floats.Sum(values) / float64(len(values))
	}
	return 0
}

func (c *Connections) Aggregate(aggregate string, field string) float64 {
	if aggregate == "count" {
		return float64(len(c.GetConnections()))
	}
	var values []float64
	c.Range(func(c *Connection) bool {
		switch {
		case strings.Contains(field, "attributes."):
			if fields := c.GetAttributes().GetFields(); fields != nil {
				split := strings.Split(field, "attributes.")
				if len(split) == 2 {
					key := split[1]
					if val := fields[key]; val != nil {
						values = append(values, val.GetNumberValue())
					}
				}
			}
		}
		return true
	})
	if len(values) == 0 {
		return 0
	}
	switch aggregate {
	case "sum":
		return floats.Sum(values)
	case "min":
		return floats.Min(values)
	case "max":
		return floats.Max(values)
	case "count":
		return float64(floats.Count(func(f float64) bool {
			return true
		}, values))
	case "prod":
		return floats.Prod(values)
	case "avg":
		return floats.Sum(values) / float64(len(values))
	}
	return 0
}