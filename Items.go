package ridder

import (
	"fmt"
	"net/http"
	"net/url"

	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
	r_types "github.com/leapforce-libraries/go_ridder/types"
)

type Item struct {
	Id                        int32                   `json:"Id"`
	Code                      string                  `json:"Code"`
	DefaultSalesPrice         *float32                `json:"DefaultSalesPrice"`
	Description               map[string]string       `json:"Description"`
	DrawingNumber             *string                 `json:"DrawingNumber"`
	EanCode                   *string                 `json:"EanCode"`
	ItemGroupCode             *string                 `json:"ItemGroupCode"`
	Image                     *string                 `json:"Image"`
	SalesDescription          map[string]string       `json:"SalesDescription"`
	SalesDrawingnumber        *string                 `json:"SalesDrawingnumber"`
	SalesItemCode             *string                 `json:"SalesItemCode"`
	AppliancePriority         *int32                  `json:"AppliancePriority"`
	Barcode2                  *string                 `json:"Barcode2"`
	DateChanged               *r_types.DateTimeString `json:"DateChanged"`
	DateCreated               *r_types.DateTimeString `json:"DateCreated"`
	DefaultSawingCode         *string                 `json:"DefaultSawingCode"`
	ExternalKey               *string                 `json:"ExternalKey"`
	FixedScalePriceNumber     *int32                  `json:"FixedScalePriceNumber"`
	BrandCode                 *string                 `json:"BrandCode"`
	ItemUnitCode              *string                 `json:"ItemUnitCode"`
	MaterialCode              *string                 `json:"MaterialCode"`
	Memo                      map[string]string       `json:"Memo"`
	PlainTextMemo             *string                 `json:"PlainTextMemo"`
	SalesMemo                 map[string]string       `json:"SalesMemo"`
	PlainTextSalesMemo        *string                 `json:"PlainTextSalesMemo"`
	StatisticCode             *string                 `json:"StatisticCode"`
	IndustryStandard          *string                 `json:"IndustryStandard"`
	Keywords                  *string                 `json:"Keywords"`
	PaintArea                 *float32                `json:"PaintArea"`
	SparePart                 *bool                   `json:"SparePart"`
	SurfaceArea               *float32                `json:"SurfaceArea"`
	SurfaceTreatment          *string                 `json:"SurfaceTreatment"`
	TextureAngle              *float32                `json:"TextureAngle"`
	Thickness                 *float32                `json:"Thickness"`
	TradeLength               *float32                `json:"TradeLength"`
	TradeWidth                *float32                `json:"TradeWidth"`
	TypeNumber                *string                 `json:"TypeNumber"`
	Unmarketable              *bool                   `json:"Unmarketable"`
	Volume                    *float32                `json:"Volume"`
	Weight                    *float32                `json:"Weight"`
	SalesPriceUnitDescription *string                 `json:"SalesPriceUnitDescription"`
	TotalFutureStock          *float32                `json:"TotalFutureStock"`
	TotalStockIn              *float32                `json:"TotalStockIn"`
	TotalStockOut             *float32                `json:"TotalStockOut"`
	TotalStockReservation     *float32                `json:"TotalStockReservation"`
	ExternalInfo1             *string                 `json:"ExternalInfo1"`
	ExternalInfo2             *string                 `json:"ExternalInfo2"`
	ExternalInfo3             *string                 `json:"ExternalInfo3"`
	ExternalInfo4             *string                 `json:"ExternalInfo4"`
	ExternalInfo5             *string                 `json:"ExternalInfo5"`
	ExternalInfo6             *string                 `json:"ExternalInfo6"`
	ExternalInfo7             *string                 `json:"ExternalInfo7"`
	ExternalInfo8             *string                 `json:"ExternalInfo8"`
	ExternalInfo9             *string                 `json:"ExternalInfo9"`
	ExternalInfo10            *string                 `json:"ExternalInfo10"`
}

func (service *Service) GetItems(itemGroupId int32) (*[]Item, *errortools.Error) {
	var values = url.Values{}
	values.Set("filter", fmt.Sprintf("FK_ITEMGROUP=%v", itemGroupId))

	var items []Item

	requestConfig := go_http.RequestConfig{
		Method:        http.MethodGet,
		Url:           service.url(fmt.Sprintf("items?%s", values.Encode())),
		ResponseModel: &items,
	}
	_, _, e := service.httpRequest(&requestConfig)
	if e != nil {
		return nil, e
	}

	return &items, nil
}
