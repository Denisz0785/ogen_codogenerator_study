// Code generated by ogen, DO NOT EDIT.

package api

import (
	"math/bits"
	"strconv"
	"time"

	"github.com/go-faster/errors"
	"github.com/go-faster/jx"

	"github.com/ogen-go/ogen/json"
	"github.com/ogen-go/ogen/validate"
)

// Encode implements json.Marshaler.
func (s *AllExpenses) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *AllExpenses) encodeFields(e *jx.Encoder) {
	{
		if s.Data != nil {
			e.FieldStart("data")
			e.ArrStart()
			for _, elem := range s.Data {
				elem.Encode(e)
			}
			e.ArrEnd()
		}
	}
}

var jsonFieldsNameOfAllExpenses = [1]string{
	0: "data",
}

// Decode decodes AllExpenses from json.
func (s *AllExpenses) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode AllExpenses to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "data":
			if err := func() error {
				s.Data = make([]Expense, 0)
				if err := d.Arr(func(d *jx.Decoder) error {
					var elem Expense
					if err := elem.Decode(d); err != nil {
						return err
					}
					s.Data = append(s.Data, elem)
					return nil
				}); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"data\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode AllExpenses")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *AllExpenses) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *AllExpenses) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *ErrorResponse) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *ErrorResponse) encodeFields(e *jx.Encoder) {
	{
		if s.Message.Set {
			e.FieldStart("message")
			s.Message.Encode(e)
		}
	}
}

var jsonFieldsNameOfErrorResponse = [1]string{
	0: "message",
}

// Decode decodes ErrorResponse from json.
func (s *ErrorResponse) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode ErrorResponse to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "message":
			if err := func() error {
				s.Message.Reset()
				if err := s.Message.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"message\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode ErrorResponse")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *ErrorResponse) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *ErrorResponse) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *Expense) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *Expense) encodeFields(e *jx.Encoder) {
	{
		e.FieldStart("id")
		e.Int(s.ID)
	}
	{
		e.FieldStart("expense_type_id")
		e.Str(s.ExpenseTypeID)
	}
	{
		if s.ReatedAt.Set {
			e.FieldStart("reated_at")
			s.ReatedAt.Encode(e, json.EncodeDateTime)
		}
	}
	{
		e.FieldStart("spent_money")
		e.Float64(s.SpentMoney)
	}
}

var jsonFieldsNameOfExpense = [4]string{
	0: "id",
	1: "expense_type_id",
	2: "reated_at",
	3: "spent_money",
}

// Decode decodes Expense from json.
func (s *Expense) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode Expense to nil")
	}
	var requiredBitSet [1]uint8

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "id":
			requiredBitSet[0] |= 1 << 0
			if err := func() error {
				v, err := d.Int()
				s.ID = int(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"id\"")
			}
		case "expense_type_id":
			requiredBitSet[0] |= 1 << 1
			if err := func() error {
				v, err := d.Str()
				s.ExpenseTypeID = string(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"expense_type_id\"")
			}
		case "reated_at":
			if err := func() error {
				s.ReatedAt.Reset()
				if err := s.ReatedAt.Decode(d, json.DecodeDateTime); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"reated_at\"")
			}
		case "spent_money":
			requiredBitSet[0] |= 1 << 3
			if err := func() error {
				v, err := d.Float64()
				s.SpentMoney = float64(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"spent_money\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode Expense")
	}
	// Validate required fields.
	var failures []validate.FieldError
	for i, mask := range [1]uint8{
		0b00001011,
	} {
		if result := (requiredBitSet[i] & mask) ^ mask; result != 0 {
			// Mask only required fields and check equality to mask using XOR.
			//
			// If XOR result is not zero, result is not equal to expected, so some fields are missed.
			// Bits of fields which would be set are actually bits of missed fields.
			missed := bits.OnesCount8(result)
			for bitN := 0; bitN < missed; bitN++ {
				bitIdx := bits.TrailingZeros8(result)
				fieldIdx := i*8 + bitIdx
				var name string
				if fieldIdx < len(jsonFieldsNameOfExpense) {
					name = jsonFieldsNameOfExpense[fieldIdx]
				} else {
					name = strconv.Itoa(fieldIdx)
				}
				failures = append(failures, validate.FieldError{
					Name:  name,
					Error: validate.ErrFieldRequired,
				})
				// Reset bit.
				result &^= 1 << bitIdx
			}
		}
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *Expense) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *Expense) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes GetAllExpensesBadRequest as json.
func (s *GetAllExpensesBadRequest) Encode(e *jx.Encoder) {
	unwrapped := (*ErrorResponse)(s)

	unwrapped.Encode(e)
}

// Decode decodes GetAllExpensesBadRequest from json.
func (s *GetAllExpensesBadRequest) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode GetAllExpensesBadRequest to nil")
	}
	var unwrapped ErrorResponse
	if err := func() error {
		if err := unwrapped.Decode(d); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return errors.Wrap(err, "alias")
	}
	*s = GetAllExpensesBadRequest(unwrapped)
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *GetAllExpensesBadRequest) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *GetAllExpensesBadRequest) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes GetAllExpensesInternalServerError as json.
func (s *GetAllExpensesInternalServerError) Encode(e *jx.Encoder) {
	unwrapped := (*ErrorResponse)(s)

	unwrapped.Encode(e)
}

// Decode decodes GetAllExpensesInternalServerError from json.
func (s *GetAllExpensesInternalServerError) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode GetAllExpensesInternalServerError to nil")
	}
	var unwrapped ErrorResponse
	if err := func() error {
		if err := unwrapped.Decode(d); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return errors.Wrap(err, "alias")
	}
	*s = GetAllExpensesInternalServerError(unwrapped)
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *GetAllExpensesInternalServerError) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *GetAllExpensesInternalServerError) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes time.Time as json.
func (o OptDateTime) Encode(e *jx.Encoder, format func(*jx.Encoder, time.Time)) {
	if !o.Set {
		return
	}
	format(e, o.Value)
}

// Decode decodes time.Time from json.
func (o *OptDateTime) Decode(d *jx.Decoder, format func(*jx.Decoder) (time.Time, error)) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptDateTime to nil")
	}
	o.Set = true
	v, err := format(d)
	if err != nil {
		return err
	}
	o.Value = v
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptDateTime) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e, json.EncodeDateTime)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptDateTime) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d, json.DecodeDateTime)
}

// Encode encodes string as json.
func (o OptString) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	e.Str(string(o.Value))
}

// Decode decodes string from json.
func (o *OptString) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptString to nil")
	}
	o.Set = true
	v, err := d.Str()
	if err != nil {
		return err
	}
	o.Value = string(v)
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptString) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptString) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}