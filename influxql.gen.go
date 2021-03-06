// Generated by tmpl
// https://github.com/benbjohnson/tmpl
//
// DO NOT EDIT!
// Source: influxql.gen.go.tmpl

package kapacitor

import (
	"fmt"
	"time"

	"github.com/influxdata/influxdb/influxql"
	"github.com/influxdata/kapacitor/models"
	"github.com/influxdata/kapacitor/pipeline"
)

type floatPointAggregator struct {
	field         string
	topBottomInfo *pipeline.TopBottomCallInfo
	aggregator    influxql.FloatPointAggregator
}

func floatPopulateAuxFieldsAndTags(ap *influxql.FloatPoint, fieldsAndTags []string, fields models.Fields, tags models.Tags) {
	ap.Aux = make([]interface{}, len(fieldsAndTags))
	for i, name := range fieldsAndTags {
		if f, ok := fields[name]; ok {
			ap.Aux[i] = f
		} else {
			ap.Aux[i] = tags[name]
		}
	}
}

func (a *floatPointAggregator) AggregateBatch(b *models.Batch) {
	for _, p := range b.Points {
		ap := &influxql.FloatPoint{
			Name:  b.Name,
			Tags:  influxql.NewTags(p.Tags),
			Time:  p.Time.UnixNano(),
			Value: p.Fields[a.field].(float64),
		}
		if a.topBottomInfo != nil {
			// We need to populate the Aux fields
			floatPopulateAuxFieldsAndTags(ap, a.topBottomInfo.FieldsAndTags, p.Fields, p.Tags)
		}
		a.aggregator.AggregateFloat(ap)
	}
}

func (a *floatPointAggregator) AggregatePoint(p *models.Point) {
	ap := &influxql.FloatPoint{
		Name:  p.Name,
		Tags:  influxql.NewTags(p.Tags),
		Time:  p.Time.UnixNano(),
		Value: p.Fields[a.field].(float64),
	}
	if a.topBottomInfo != nil {
		// We need to populate the Aux fields
		floatPopulateAuxFieldsAndTags(ap, a.topBottomInfo.FieldsAndTags, p.Fields, p.Tags)
	}
	a.aggregator.AggregateFloat(ap)
}

type floatPointBulkAggregator struct {
	field         string
	topBottomInfo *pipeline.TopBottomCallInfo
	aggregator    pipeline.FloatBulkPointAggregator
}

func (a *floatPointBulkAggregator) AggregateBatch(b *models.Batch) {
	slice := make([]influxql.FloatPoint, len(b.Points))
	for i, p := range b.Points {
		slice[i] = influxql.FloatPoint{
			Name:  b.Name,
			Tags:  influxql.NewTags(p.Tags),
			Time:  p.Time.UnixNano(),
			Value: p.Fields[a.field].(float64),
		}
		if a.topBottomInfo != nil {
			// We need to populate the Aux fields
			floatPopulateAuxFieldsAndTags(&slice[i], a.topBottomInfo.FieldsAndTags, p.Fields, p.Tags)
		}
	}
	a.aggregator.AggregateFloatBulk(slice)
}

func (a *floatPointBulkAggregator) AggregatePoint(p *models.Point) {
	ap := &influxql.FloatPoint{
		Name:  p.Name,
		Tags:  influxql.NewTags(p.Tags),
		Time:  p.Time.UnixNano(),
		Value: p.Fields[a.field].(float64),
	}
	if a.topBottomInfo != nil {
		// We need to populate the Aux fields
		floatPopulateAuxFieldsAndTags(ap, a.topBottomInfo.FieldsAndTags, p.Fields, p.Tags)
	}
	a.aggregator.AggregateFloat(ap)
}

type floatPointEmitter struct {
	baseReduceContext
	emitter influxql.FloatPointEmitter
}

func (e *floatPointEmitter) EmitPoint() (models.Point, error) {
	slice := e.emitter.Emit()
	if len(slice) != 1 {
		return models.Point{}, fmt.Errorf("unexpected result from InfluxQL function, got %d points expected 1", len(slice))
	}
	ap := slice[0]
	var t time.Time
	if e.pointTimes {
		if ap.Time == influxql.ZeroTime {
			t = e.time
		} else {
			t = time.Unix(0, ap.Time).UTC()
		}
	} else {
		t = e.time
	}
	return models.Point{
		Name:       e.name,
		Time:       t,
		Group:      e.group,
		Dimensions: e.dimensions,
		Tags:       e.tags,
		Fields:     map[string]interface{}{e.as: ap.Value},
	}, nil
}

func (e *floatPointEmitter) EmitBatch() models.Batch {
	slice := e.emitter.Emit()
	b := models.Batch{
		Name:   e.name,
		TMax:   e.time,
		Group:  e.group,
		Tags:   e.tags,
		Points: make([]models.BatchPoint, len(slice)),
	}
	var t time.Time
	for i, ap := range slice {
		if e.pointTimes {
			if ap.Time == influxql.ZeroTime {
				t = e.time
			} else {
				t = time.Unix(0, ap.Time).UTC()
			}
		} else {
			t = e.time
		}
		b.Points[i] = models.BatchPoint{
			Time:   t,
			Tags:   ap.Tags.KeyValues(),
			Fields: map[string]interface{}{e.as: ap.Value},
		}
	}
	return b
}

type integerPointAggregator struct {
	field         string
	topBottomInfo *pipeline.TopBottomCallInfo
	aggregator    influxql.IntegerPointAggregator
}

func integerPopulateAuxFieldsAndTags(ap *influxql.IntegerPoint, fieldsAndTags []string, fields models.Fields, tags models.Tags) {
	ap.Aux = make([]interface{}, len(fieldsAndTags))
	for i, name := range fieldsAndTags {
		if f, ok := fields[name]; ok {
			ap.Aux[i] = f
		} else {
			ap.Aux[i] = tags[name]
		}
	}
}

func (a *integerPointAggregator) AggregateBatch(b *models.Batch) {
	for _, p := range b.Points {
		ap := &influxql.IntegerPoint{
			Name:  b.Name,
			Tags:  influxql.NewTags(p.Tags),
			Time:  p.Time.UnixNano(),
			Value: p.Fields[a.field].(int64),
		}
		if a.topBottomInfo != nil {
			// We need to populate the Aux fields
			integerPopulateAuxFieldsAndTags(ap, a.topBottomInfo.FieldsAndTags, p.Fields, p.Tags)
		}
		a.aggregator.AggregateInteger(ap)
	}
}

func (a *integerPointAggregator) AggregatePoint(p *models.Point) {
	ap := &influxql.IntegerPoint{
		Name:  p.Name,
		Tags:  influxql.NewTags(p.Tags),
		Time:  p.Time.UnixNano(),
		Value: p.Fields[a.field].(int64),
	}
	if a.topBottomInfo != nil {
		// We need to populate the Aux fields
		integerPopulateAuxFieldsAndTags(ap, a.topBottomInfo.FieldsAndTags, p.Fields, p.Tags)
	}
	a.aggregator.AggregateInteger(ap)
}

type integerPointBulkAggregator struct {
	field         string
	topBottomInfo *pipeline.TopBottomCallInfo
	aggregator    pipeline.IntegerBulkPointAggregator
}

func (a *integerPointBulkAggregator) AggregateBatch(b *models.Batch) {
	slice := make([]influxql.IntegerPoint, len(b.Points))
	for i, p := range b.Points {
		slice[i] = influxql.IntegerPoint{
			Name:  b.Name,
			Tags:  influxql.NewTags(p.Tags),
			Time:  p.Time.UnixNano(),
			Value: p.Fields[a.field].(int64),
		}
		if a.topBottomInfo != nil {
			// We need to populate the Aux fields
			integerPopulateAuxFieldsAndTags(&slice[i], a.topBottomInfo.FieldsAndTags, p.Fields, p.Tags)
		}
	}
	a.aggregator.AggregateIntegerBulk(slice)
}

func (a *integerPointBulkAggregator) AggregatePoint(p *models.Point) {
	ap := &influxql.IntegerPoint{
		Name:  p.Name,
		Tags:  influxql.NewTags(p.Tags),
		Time:  p.Time.UnixNano(),
		Value: p.Fields[a.field].(int64),
	}
	if a.topBottomInfo != nil {
		// We need to populate the Aux fields
		integerPopulateAuxFieldsAndTags(ap, a.topBottomInfo.FieldsAndTags, p.Fields, p.Tags)
	}
	a.aggregator.AggregateInteger(ap)
}

type integerPointEmitter struct {
	baseReduceContext
	emitter influxql.IntegerPointEmitter
}

func (e *integerPointEmitter) EmitPoint() (models.Point, error) {
	slice := e.emitter.Emit()
	if len(slice) != 1 {
		return models.Point{}, fmt.Errorf("unexpected result from InfluxQL function, got %d points expected 1", len(slice))
	}
	ap := slice[0]
	var t time.Time
	if e.pointTimes {
		if ap.Time == influxql.ZeroTime {
			t = e.time
		} else {
			t = time.Unix(0, ap.Time).UTC()
		}
	} else {
		t = e.time
	}
	return models.Point{
		Name:       e.name,
		Time:       t,
		Group:      e.group,
		Dimensions: e.dimensions,
		Tags:       e.tags,
		Fields:     map[string]interface{}{e.as: ap.Value},
	}, nil
}

func (e *integerPointEmitter) EmitBatch() models.Batch {
	slice := e.emitter.Emit()
	b := models.Batch{
		Name:   e.name,
		TMax:   e.time,
		Group:  e.group,
		Tags:   e.tags,
		Points: make([]models.BatchPoint, len(slice)),
	}
	var t time.Time
	for i, ap := range slice {
		if e.pointTimes {
			if ap.Time == influxql.ZeroTime {
				t = e.time
			} else {
				t = time.Unix(0, ap.Time).UTC()
			}
		} else {
			t = e.time
		}
		b.Points[i] = models.BatchPoint{
			Time:   t,
			Tags:   ap.Tags.KeyValues(),
			Fields: map[string]interface{}{e.as: ap.Value},
		}
	}
	return b
}

// floatReduceContext uses composition to implement the reduceContext interface
type floatReduceContext struct {
	floatPointAggregator
	floatPointEmitter
}

// floatBulkReduceContext uses composition to implement the reduceContext interface
type floatBulkReduceContext struct {
	floatPointBulkAggregator
	floatPointEmitter
}

// floatIntegerReduceContext uses composition to implement the reduceContext interface
type floatIntegerReduceContext struct {
	floatPointAggregator
	integerPointEmitter
}

// floatBulkIntegerReduceContext uses composition to implement the reduceContext interface
type floatBulkIntegerReduceContext struct {
	floatPointBulkAggregator
	integerPointEmitter
}

// integerFloatReduceContext uses composition to implement the reduceContext interface
type integerFloatReduceContext struct {
	integerPointAggregator
	floatPointEmitter
}

// integerBulkFloatReduceContext uses composition to implement the reduceContext interface
type integerBulkFloatReduceContext struct {
	integerPointBulkAggregator
	floatPointEmitter
}

// integerReduceContext uses composition to implement the reduceContext interface
type integerReduceContext struct {
	integerPointAggregator
	integerPointEmitter
}

// integerBulkReduceContext uses composition to implement the reduceContext interface
type integerBulkReduceContext struct {
	integerPointBulkAggregator
	integerPointEmitter
}

func determineReduceContextCreateFn(method string, value interface{}, rc pipeline.ReduceCreater) (fn createReduceContextFunc, err error) {
	switch value.(type) {

	case float64:
		switch {

		case rc.CreateFloatReducer != nil:
			fn = func(c baseReduceContext) reduceContext {
				a, e := rc.CreateFloatReducer()
				return &floatReduceContext{
					floatPointAggregator: floatPointAggregator{
						field:         c.field,
						topBottomInfo: rc.TopBottomCallInfo,
						aggregator:    a,
					},
					floatPointEmitter: floatPointEmitter{
						baseReduceContext: c,
						emitter:           e,
					},
				}
			}
		case rc.CreateFloatBulkReducer != nil:
			fn = func(c baseReduceContext) reduceContext {
				a, e := rc.CreateFloatBulkReducer()
				return &floatBulkReduceContext{
					floatPointBulkAggregator: floatPointBulkAggregator{
						field:         c.field,
						topBottomInfo: rc.TopBottomCallInfo,
						aggregator:    a,
					},
					floatPointEmitter: floatPointEmitter{
						baseReduceContext: c,
						emitter:           e,
					},
				}
			}

		case rc.CreateFloatIntegerReducer != nil:
			fn = func(c baseReduceContext) reduceContext {
				a, e := rc.CreateFloatIntegerReducer()
				return &floatIntegerReduceContext{
					floatPointAggregator: floatPointAggregator{
						field:         c.field,
						topBottomInfo: rc.TopBottomCallInfo,
						aggregator:    a,
					},
					integerPointEmitter: integerPointEmitter{
						baseReduceContext: c,
						emitter:           e,
					},
				}
			}
		case rc.CreateFloatBulkIntegerReducer != nil:
			fn = func(c baseReduceContext) reduceContext {
				a, e := rc.CreateFloatBulkIntegerReducer()
				return &floatBulkIntegerReduceContext{
					floatPointBulkAggregator: floatPointBulkAggregator{
						field:         c.field,
						topBottomInfo: rc.TopBottomCallInfo,
						aggregator:    a,
					},
					integerPointEmitter: integerPointEmitter{
						baseReduceContext: c,
						emitter:           e,
					},
				}
			}

		default:
			err = fmt.Errorf("cannot apply %s to float64 field", method)
		}

	case int64:
		switch {

		case rc.CreateIntegerFloatReducer != nil:
			fn = func(c baseReduceContext) reduceContext {
				a, e := rc.CreateIntegerFloatReducer()
				return &integerFloatReduceContext{
					integerPointAggregator: integerPointAggregator{
						field:         c.field,
						topBottomInfo: rc.TopBottomCallInfo,
						aggregator:    a,
					},
					floatPointEmitter: floatPointEmitter{
						baseReduceContext: c,
						emitter:           e,
					},
				}
			}
		case rc.CreateIntegerBulkFloatReducer != nil:
			fn = func(c baseReduceContext) reduceContext {
				a, e := rc.CreateIntegerBulkFloatReducer()
				return &integerBulkFloatReduceContext{
					integerPointBulkAggregator: integerPointBulkAggregator{
						field:         c.field,
						topBottomInfo: rc.TopBottomCallInfo,
						aggregator:    a,
					},
					floatPointEmitter: floatPointEmitter{
						baseReduceContext: c,
						emitter:           e,
					},
				}
			}

		case rc.CreateIntegerReducer != nil:
			fn = func(c baseReduceContext) reduceContext {
				a, e := rc.CreateIntegerReducer()
				return &integerReduceContext{
					integerPointAggregator: integerPointAggregator{
						field:         c.field,
						topBottomInfo: rc.TopBottomCallInfo,
						aggregator:    a,
					},
					integerPointEmitter: integerPointEmitter{
						baseReduceContext: c,
						emitter:           e,
					},
				}
			}
		case rc.CreateIntegerBulkReducer != nil:
			fn = func(c baseReduceContext) reduceContext {
				a, e := rc.CreateIntegerBulkReducer()
				return &integerBulkReduceContext{
					integerPointBulkAggregator: integerPointBulkAggregator{
						field:         c.field,
						topBottomInfo: rc.TopBottomCallInfo,
						aggregator:    a,
					},
					integerPointEmitter: integerPointEmitter{
						baseReduceContext: c,
						emitter:           e,
					},
				}
			}

		default:
			err = fmt.Errorf("cannot apply %s to int64 field", method)
		}

	default:
		err = fmt.Errorf("invalid field type: %T", value)
	}
	return
}
