package converter

import (
	"order/internal/repository/model"
	v1 "order/pkg/inventory/inventory"
	v2 "order/pkg/payment/payment"
	"time"
)

func PaymentMethodToEnum(s string) v2.PaymentMethod {
	switch s {
	case "CARD":
		return v2.PaymentMethod_PAYMENT_METHOD_CARD
	case "SBP":
		return v2.PaymentMethod_PAYMENT_METHOD_SBP
	case "CREDITCARD":
		return v2.PaymentMethod_PAYMENT_METHOD_CARD
	case "INVESTORMONEY":
		return v2.PaymentMethod_PAYMENT_METHOD_INVESTOR_MONEY
	default:
		return v2.PaymentMethod_PAYMENT_METHOD_UNKNOWN
	}
}

func PartFilterToProto(filter model.PartsFilter) *v1.PartsFilter {
	categories := make([]v1.Category, 0, len(filter.Categories))
	for _, v := range filter.Categories {
		categories = append(categories, v1.Category(v))
	}
	return &v1.PartsFilter{
		Uuids:                 filter.Uuids,
		Categories:            categories,
		Names:                 filter.Names,
		Tags:                  filter.Tags,
		ManufacturerCountries: filter.ManufacturerCountries,
	}
}

func PartProtoToModel(part *v1.Part) *model.Part {
	var updatedAt *time.Time
	if part.UpdatedAt != nil {
		tmp := part.UpdatedAt.AsTime()
		updatedAt = &tmp
	}
	var createdAt *time.Time
	if part.CreatedAt != nil {
		tmp := part.CreatedAt.AsTime()
		createdAt = &tmp
	}
	return &model.Part{
		UUID:          part.UUID,
		Name:          part.Name,
		Description:   part.Description,
		Price:         part.Price,
		StockQuantity: part.StockQuantity,
		Category:      int8(part.Category),
		Dimensions:    DimensionsToModel(part.Dimensions),
		Manufacturer:  ManufacturerToModel(part.Manufacturer),
		Tags:          part.Tags,
		MetaData:      MetaDataToModel(part.Metadata),
		CreatedAt:     createdAt,
		UpdatedAt:     updatedAt,
	}
}

func DimensionsToModel(dimen *v1.Dimensions) model.Dimensions {
	return model.Dimensions{
		Height: dimen.Height,
		Width:  dimen.Width,
		Length: dimen.Length,
		Weight: dimen.Weight,
	}
}

func ManufacturerToModel(manu *v1.Manufacturer) model.Manufacturer {
	return model.Manufacturer{
		Name:    manu.Name,
		Country: manu.Country,
		Website: manu.Website,
	}
}
func MetaDataToModel(metaD map[string]*v1.Value) map[string]any {
	res := make(map[string]any, len(metaD))
	for k, value := range metaD {
		if value == nil {
			continue
		}
		switch {
		case value.StringValue != "":
			res[k] = value.StringValue
		case value.BoolValue:
			res[k] = value.BoolValue
		case value.Int64Value != 0:
			res[k] = value.Int64Value
		case value.DoubleValue != 0:
			res[k] = value.DoubleValue
		default:
			res[k] = nil
		}
	}
	return res
}
func PartsListToModel(parts []*v1.Part) []*model.Part {
	res := make([]*model.Part, len(parts))
	for _, part := range parts {
		res = append(res, PartProtoToModel(part))
	}
	return res
}

func