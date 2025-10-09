package converter

import (
	"order/api"
	repoModel "order/internal/repository/model"
	"order/internal/service/model"
	v1 "order/pkg/inventory/inventory"
	v2 "order/pkg/payment/payment"
	"time"

	"github.com/google/uuid"
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

func PartFilterToProto(filter repoModel.PartsFilter) *v1.PartsFilter {
	categories := make([]v1.Category, 0, len(filter.Categories))
	for _, v := range filter.Categories {
		categories = append(categories, convertCategoryFromPB(v))
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

func UUIDToString(uuid uuid.UUID) string {
	tmp := uuid.String()
	return tmp
}
func StringToUUID(s string) uuid.UUID {
	return uuid.MustParse(s)
}

func OptNilUUIDToUUID(str *string) api.OptNilUUID {
	return api.OptNilUUID{
		Value: StringToUUID(*str),
		Set:   true,
		Null:  false,
	}
}
func OptNilStringToString(str *string) api.OptNilString {
	return api.OptNilString{
		Value: *str,
		Set:   true,
		Null:  false,
	}
}
func ConvertPaymentMethodToString(method repoModel.PaymentMethod) string {
	switch method {
	case repoModel.PaymentMethodCard:
		return "CARD"
	case repoModel.PaymentMethodSBP:
		return "PaymentMethodSBP"
	case repoModel.PaymentMethodCreditCard:
		return "CREDIT_CARD"
	case repoModel.PaymentMethodInvestorMoney:
		return "INVESTOR_MONEY"
	default:
		return "UNKNOWN"
	}
}
func ConvertPaymentMethod(method string) repoModel.PaymentMethod {
	switch method {
	case "CARD":
		return repoModel.PaymentMethodCard
	case "PaymentMethodSBP":
		return repoModel.PaymentMethodSBP
	case "CREDIT_CARD":
		return repoModel.PaymentMethodCreditCard
	case "INVESTOR_MONEY":
		return repoModel.PaymentMethodInvestorMoney
	default:
		return repoModel.PaymentMethodUnknown
	}
}
func convertCategoryFromPB(category repoModel.Category) v1.Category {
	switch category {
	case repoModel.CategoryEngine:
		return v1.Category_CATEGORY_ENGINE
	case repoModel.CategoryFuel:
		return v1.Category_CATEGORY_FUEL
	case repoModel.CategoryWing:
		return v1.Category_CATEGORY_WING
	case repoModel.CategoryPorthole:
		return v1.Category_CATEGORY_PORTHOLE
	default:
		return v1.Category_CATEGORY_UNKNOWN
	}
}
