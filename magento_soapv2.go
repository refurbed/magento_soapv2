package binding

import (
	"github.com/fiorix/wsdl2go/soap"
)

// Namespace was auto-generated from WSDL.
var Namespace = "urn:Magento"

// NewPortType creates an initializes a PortType.
func NewPortType(cli *soap.Client) PortType {
	return &portType{cli}
}

// PortType was auto-generated from WSDL
// and defines interface for the remote service. Useful for testing.
type PortType interface {
	// Create an order from Adcurve shopping cart
	AdcurveShoppingCartCreateAdcurveOrder(sessionId string, quoteId int, storeId string, licenses *ArrayOfString, orderAdditionalAttributes *ShoppingCartOrderAdditionalAttributesEntity) (string, error)

	// Add product(s) to adcurve shopping cart
	AdcurveShoppingCartProductAdcurveProductAdd(sessionId string, quoteId int, products *ShoppingCartAdcurveProductEntityArray, storeId string) (bool, error)

	// Update product(s) quantities in adcurve shopping cart
	AdcurveShoppingCartProductAdcurveProductUpdate(sessionId string, quoteId int, products *ShoppingCartAdcurveProductEntityArray, storeId string) (bool, error)

	// Assign product to category
	CatalogCategoryAssignProduct(sessionId string, categoryId int, product string, position string, identifierType string) (bool, error)

	// Retrieve list of assigned products
	CatalogCategoryAssignedProducts(sessionId string, categoryId int) (*CatalogAssignedProductArray, error)

	// Set/Get current store view
	CatalogCategoryAttributeCurrentStore(sessionId string, storeView string) (int, error)

	// Retrieve category attributes
	CatalogCategoryAttributeList(sessionId string) (*CatalogAttributeEntityArray, error)

	// Retrieve attribute options
	CatalogCategoryAttributeOptions(sessionId string, attributeId string, storeView string) (*CatalogAttributeOptionEntityArray, error)

	// Create new category and return its id.
	CatalogCategoryCreate(sessionId string, parentId int, categoryData *CatalogCategoryEntityCreate, storeView string) (int, error)

	// Set_Get current store view
	CatalogCategoryCurrentStore(sessionId string, storeView string) (int, error)

	// Delete category
	CatalogCategoryDelete(sessionId string, categoryId int) (bool, error)

	// Retrieve hierarchical tree of categories.
	CatalogCategoryInfo(sessionId string, categoryId int, storeView string, attributes *ArrayOfString) (*CatalogCategoryInfo, error)

	// Retrieve hierarchical tree of categories.
	CatalogCategoryLevel(sessionId string, website string, storeView string, parentCategory string) (*ArrayOfCatalogCategoryEntitiesNoChildren, error)

	// Move category in tree
	CatalogCategoryMove(sessionId string, categoryId int, parentId int, afterId string) (bool, error)

	// Remove product assignment from category
	CatalogCategoryRemoveProduct(sessionId string, categoryId int, product string, identifierType string) (bool, error)

	// Retrieve hierarchical tree of categories.
	CatalogCategoryTree(sessionId string, parentId string, storeView string) (*CatalogCategoryTree, error)

	// Update category
	CatalogCategoryUpdate(sessionId string, categoryId int, categoryData *CatalogCategoryEntityCreate, storeView string) (bool, error)

	// Update assigned product
	CatalogCategoryUpdateProduct(sessionId string, categoryId int, product string, position string, identifierType string) (bool, error)

	// Retrieve stock data by product ids
	CatalogInventoryStockItemList(sessionId string, products *ArrayOfString) (*CatalogInventoryStockItemEntityArray, error)

	// Multi stock item update
	CatalogInventoryStockItemMultiUpdate(sessionId string, productIds *ArrayOfString, productData *CatalogInventoryStockItemUpdateEntityArray) (bool, error)

	// Update product stock data
	CatalogInventoryStockItemUpdate(sessionId string, product string, data *CatalogInventoryStockItemUpdateEntity) (int, error)

	// Add option to attribute
	CatalogProductAttributeAddOption(sessionId string, attribute string, data *CatalogProductAttributeOptionEntityToAdd) (bool, error)

	// Create new attribute
	CatalogProductAttributeCreate(sessionId string, data *CatalogProductAttributeEntityToCreate) (int, error)

	// Set/Get current store view
	CatalogProductAttributeCurrentStore(sessionId string, storeView string) (int, error)

	// Get full information about attribute with list of options
	CatalogProductAttributeInfo(sessionId string, attribute string) (*CatalogProductAttributeEntity, error)

	// Retrieve attribute list
	CatalogProductAttributeList(sessionId string, setId int) (*CatalogAttributeEntityArray, error)

	// Upload new product image
	CatalogProductAttributeMediaCreate(sessionId string, product string, data *CatalogProductAttributeMediaCreateEntity, storeView string, identifierType string) (string, error)

	// Set/Get current store view
	CatalogProductAttributeMediaCurrentStore(sessionId string, storeView string) (int, error)

	// Retrieve product image data
	CatalogProductAttributeMediaInfo(sessionId string, product string, file string, storeView string, identifierType string) (*CatalogProductImageEntity, error)

	// Retrieve product image list
	CatalogProductAttributeMediaList(sessionId string, product string, storeView string, identifierType string) (*CatalogProductImageEntityArray, error)

	// Remove product image
	CatalogProductAttributeMediaRemove(sessionId string, product string, file string, identifierType string) (bool, error)

	// Retrieve product image types
	CatalogProductAttributeMediaTypes(sessionId string, setId string) (*CatalogProductAttributeMediaTypeEntityArray, error)

	// Update product image
	CatalogProductAttributeMediaUpdate(sessionId string, product string, file string, data *CatalogProductAttributeMediaCreateEntity, storeView string, identifierType string) (bool, error)

	// Retrieve attribute options
	CatalogProductAttributeOptions(sessionId string, attributeId string, storeView string) (*CatalogAttributeOptionEntityArray, error)

	// Delete attribute
	CatalogProductAttributeRemove(sessionId string, attribute string) (bool, error)

	// Remove option from attribute
	CatalogProductAttributeRemoveOption(sessionId string, attribute string, optionId string) (bool, error)

	// Add attribute into attribute set
	CatalogProductAttributeSetAttributeAdd(sessionId string, attributeId string, attributeSetId string, attributeGroupId string, sortOrder string) (bool, error)

	// Remove attribute from attribute set
	CatalogProductAttributeSetAttributeRemove(sessionId string, attributeId string, attributeSetId string) (bool, error)

	// Create product attribute set based on another set
	CatalogProductAttributeSetCreate(sessionId string, attributeSetName string, skeletonSetId string) (int, error)

	// Create group within existing attribute set
	CatalogProductAttributeSetGroupAdd(sessionId string, attributeSetId string, groupName string) (int, error)

	// Remove group from attribute set
	CatalogProductAttributeSetGroupRemove(sessionId string, attributeGroupId string) (bool, error)

	// Rename existing group
	CatalogProductAttributeSetGroupRename(sessionId string, groupId string, groupName string) (bool, error)

	// Retrieve product attribute sets
	CatalogProductAttributeSetList(sessionId string) (*CatalogProductAttributeSetEntityArray, error)

	// Remove product attribute set
	CatalogProductAttributeSetRemove(sessionId string, attributeSetId string, forceProductsRemove string) (bool, error)

	// Retrieve product tier prices
	CatalogProductAttributeTierPriceInfo(sessionId string, product string, identifierType string) (*CatalogProductTierPriceEntityArray, error)

	// Update product tier prices
	CatalogProductAttributeTierPriceUpdate(sessionId string, product string, tier_price *CatalogProductTierPriceEntityArray, identifierType string) (int, error)

	// Get list of possible attribute types
	CatalogProductAttributeTypes(sessionId string) (*CatalogAttributeOptionEntityArray, error)

	// Update attribute
	CatalogProductAttributeUpdate(sessionId string, attribute string, data *CatalogProductAttributeEntityToUpdate) (bool, error)

	// Create new product and return product id
	CatalogProductCreate(sessionId string, _type string, set string, sku string, productData *CatalogProductCreateEntity, storeView string) (int, error)

	// Set/Get current store view
	CatalogProductCurrentStore(sessionId string, storeView string) (int, error)

	// Add new custom option into product
	CatalogProductCustomOptionAdd(sessionId string, productId string, data *CatalogProductCustomOptionToAdd, store string) (bool, error)

	// Get full information about custom option in product
	CatalogProductCustomOptionInfo(sessionId string, optionId string, store string) (*CatalogProductCustomOptionInfoEntity, error)

	// Retrieve list of product custom options
	CatalogProductCustomOptionList(sessionId string, productId string, store string) (*CatalogProductCustomOptionListArray, error)

	// Remove custom option
	CatalogProductCustomOptionRemove(sessionId string, optionId string) (bool, error)

	// Get list of available custom option types
	CatalogProductCustomOptionTypes(sessionId string) (*CatalogProductCustomOptionTypesArray, error)

	// Update product custom option
	CatalogProductCustomOptionUpdate(sessionId string, optionId string, data *CatalogProductCustomOptionToUpdate, store string) (bool, error)

	// Add new custom option values
	CatalogProductCustomOptionValueAdd(sessionId string, optionId string, data *CatalogProductCustomOptionValueAddArray, store string) (bool, error)

	// Retrieve custom option value info
	CatalogProductCustomOptionValueInfo(sessionId string, valueId string, store string) (*CatalogProductCustomOptionValueInfoEntity, error)

	// Retrieve custom option values list
	CatalogProductCustomOptionValueList(sessionId string, optionId string, store string) (*CatalogProductCustomOptionValueListArray, error)

	// Remove value from custom option
	CatalogProductCustomOptionValueRemove(sessionId string, valueId string) (bool, error)

	// Update custom option value
	CatalogProductCustomOptionValueUpdate(sessionId string, valueId string, data *CatalogProductCustomOptionValueUpdateEntity, storeId string) (bool, error)

	// Delete product
	CatalogProductDelete(sessionId string, product string, identifierType string) (int, error)

	// Add links to downloadable product
	CatalogProductDownloadableLinkAdd(sessionId string, productId string, resource *CatalogProductDownloadableLinkAddEntity, resourceType string, store string, identifierType string) (int, error)

	// Retrieve list of links and samples for downloadable product
	CatalogProductDownloadableLinkList(sessionId string, productId string, store string, identifierType string) (*CatalogProductDownloadableLinkInfoEntity, error)

	// Remove links and samples from downloadable product
	CatalogProductDownloadableLinkRemove(sessionId string, linkId string, resourceType string) (bool, error)

	// Get product special price data
	CatalogProductGetSpecialPrice(sessionId string, product string, storeView string, identifierType string) (*CatalogProductSpecialPriceReturnEntity, error)

	// Retrieve product
	CatalogProductInfo(sessionId string, productId string, storeView string, attributes *CatalogProductRequestAttributes, identifierType string) (*CatalogProductReturnEntity, error)

	// Assign product link
	CatalogProductLinkAssign(sessionId string, _type string, product string, linkedProduct string, data *CatalogProductLinkEntity, identifierType string) (bool, error)

	// Retrieve product link type attributes
	CatalogProductLinkAttributes(sessionId string, _type string) (*CatalogProductLinkAttributeEntityArray, error)

	// Retrieve linked products
	CatalogProductLinkList(sessionId string, _type string, product string, identifierType string) (*CatalogProductLinkEntityArray, error)

	// Remove product link
	CatalogProductLinkRemove(sessionId string, _type string, product string, linkedProduct string, identifierType string) (bool, error)

	// Retrieve product link types
	CatalogProductLinkTypes(sessionId string) (*ArrayOfString, error)

	// Update product link
	CatalogProductLinkUpdate(sessionId string, _type string, product string, linkedProduct string, data *CatalogProductLinkEntity, identifierType string) (bool, error)

	// Retrieve products list by filters
	CatalogProductList(sessionId string, filters *Filters, storeView string) (*CatalogProductEntityArray, error)

	// Get list of additional attributes which are not in default create/update
	// list
	CatalogProductListOfAdditionalAttributes(sessionId string, productType string, attributeSetId string) (*CatalogAttributeEntityArray, error)

	// Product multi update
	CatalogProductMultiUpdate(sessionId string, productIds *ArrayOfString, productData *CatalogProductCreateEntityArray, store string, identifierType string) (bool, error)

	// Update product special price
	CatalogProductSetSpecialPrice(sessionId string, product string, specialPrice string, fromDate string, toDate string, storeView string, identifierType string) (int, error)

	// Add tag(s) to product
	CatalogProductTagAdd(sessionId string, data *CatalogProductTagAddEntity) (*AssociativeArray, error)

	// Retrieve product tag info
	CatalogProductTagInfo(sessionId string, tagId string, store string) (*CatalogProductTagInfoEntity, error)

	// Retrieve list of tags by product
	CatalogProductTagList(sessionId string, productId string, store string) (*CatalogProductTagListEntityArray, error)

	// Remove product tag
	CatalogProductTagRemove(sessionId string, tagId string) (bool, error)

	// Update product tag
	CatalogProductTagUpdate(sessionId string, tagId string, data *CatalogProductTagUpdateEntity, store string) (bool, error)

	// Retrieve product types
	CatalogProductTypeList(sessionId string) (*CatalogProductTypeEntityArray, error)

	// Update product
	CatalogProductUpdate(sessionId string, product string, productData *CatalogProductCreateEntity, storeView string, identifierType string) (bool, error)

	// Create customer address
	CustomerAddressCreate(sessionId string, customerId int, addressData *CustomerAddressEntityCreate) (int, error)

	// Delete customer address
	CustomerAddressDelete(sessionId string, addressId int) (bool, error)

	// Retrieve customer address data
	CustomerAddressInfo(sessionId string, addressId int) (*CustomerAddressEntityItem, error)

	// Retrieve customer addresses
	CustomerAddressList(sessionId string, customerId int) (*CustomerAddressEntityArray, error)

	// Update customer address data
	CustomerAddressUpdate(sessionId string, addressId int, addressData *CustomerAddressEntityCreate) (bool, error)

	// Create customer
	CustomerCustomerCreate(sessionId string, customerData *CustomerCustomerEntityToCreate) (int, error)

	// Delete customer
	CustomerCustomerDelete(sessionId string, customerId int) (bool, error)

	// Retrieve customer data
	CustomerCustomerInfo(sessionId string, customerId int, attributes *ArrayOfString) (*CustomerCustomerEntity, error)

	// Retrieve customers
	CustomerCustomerList(sessionId string, filters *Filters) (*CustomerCustomerEntityArray, error)

	// Update customer data
	CustomerCustomerUpdate(sessionId string, customerId int, customerData *CustomerCustomerEntityToCreate) (bool, error)

	// Retrieve customer groups
	CustomerGroupList(sessionId string) (*CustomerGroupEntityArray, error)

	// List of countries
	DirectoryCountryList(sessionId string) (*DirectoryCountryEntityArray, error)

	// List of regions in specified country
	DirectoryRegionList(sessionId string, country string) (*DirectoryRegionEntityArray, error)

	// End web service session
	EndSession(sessionId string) (bool, error)

	// Set a gift message to the cart
	GiftMessageSetForQuote(sessionId string, quoteId string, giftMessage *GiftMessageEntity, storeId string) (*GiftMessageResponse, error)

	// Setting a gift messages to the quote item
	GiftMessageSetForQuoteItem(sessionId string, quoteItemId string, giftMessage *GiftMessageEntity, storeId string) (*GiftMessageResponse, error)

	// Setting a gift messages to the quote items by products
	GiftMessageSetForQuoteProduct(sessionId string, quoteId string, productsAndMessages *GiftMessageAssociativeProductsEntityArray, storeId string) (*GiftMessageResponseArray, error)

	// List of global faults
	GlobalFaults(sessionId string) (*ArrayOfExistsFaltures, error)

	// Login user and retrive session id
	Login(username string, apiKey string) (string, error)

	// Info about current Magento installation
	MagentoInfo(sessionId string) (*MagentoInfoEntity, error)

	// List of resource faults
	ResourceFaults(resourceName string, sessionId string) (*ArrayOfExistsFaltures, error)

	// List of available resources
	Resources(sessionId string) (*ArrayOfApis, error)

	// Add comment to order
	SalesOrderAddComment(sessionId string, orderIncrementId string, status string, comment string, notify string) (bool, error)

	// Cancel order
	SalesOrderCancel(sessionId string, orderIncrementId string) (bool, error)

	// Add new comment to shipment
	SalesOrderCreditmemoAddComment(sessionId string, creditmemoIncrementId string, comment string, notifyCustomer int, includeComment int) (bool, error)

	// Cancel creditmemo
	SalesOrderCreditmemoCancel(sessionId string, creditmemoIncrementId string) (bool, error)

	// Create new creditmemo for order
	SalesOrderCreditmemoCreate(sessionId string, orderIncrementId string, creditmemoData *SalesOrderCreditmemoData, comment string, notifyCustomer int, includeComment int, refundToStoreCreditAmount string) (string, error)

	// Retrieve creditmemo information
	SalesOrderCreditmemoInfo(sessionId string, creditmemoIncrementId string) (*SalesOrderCreditmemoEntity, error)

	// Retrieve list of creditmemos by filters
	SalesOrderCreditmemoList(sessionId string, filters *Filters) (*SalesOrderCreditmemoEntityArray, error)

	// Hold order
	SalesOrderHold(sessionId string, orderIncrementId string) (bool, error)

	// Retrieve order information
	SalesOrderInfo(sessionId string, orderIncrementId string) (*SalesOrderEntity, error)

	// Add new comment to shipment
	SalesOrderInvoiceAddComment(sessionId string, invoiceIncrementId string, comment string, email string, includeComment string) (bool, error)

	// Cancel invoice
	SalesOrderInvoiceCancel(sessionId string, invoiceIncrementId string) (bool, error)

	// Capture invoice
	SalesOrderInvoiceCapture(sessionId string, invoiceIncrementId string) (bool, error)

	// Create new invoice for order
	SalesOrderInvoiceCreate(sessionId string, invoiceIncrementId string, itemsQty *OrderItemIdQtyArray, comment string, email string, includeComment string) (string, error)

	// Retrieve invoice information
	SalesOrderInvoiceInfo(sessionId string, invoiceIncrementId string) (*SalesOrderInvoiceEntity, error)

	// Retrieve list of invoices by filters
	SalesOrderInvoiceList(sessionId string, filters *Filters) (*SalesOrderInvoiceEntityArray, error)

	// Void invoice
	SalesOrderInvoiceVoid(sessionId string, invoiceIncrementId string) (bool, error)

	// Retrieve list of orders by filters
	SalesOrderList(sessionId string, filters *Filters) (*SalesOrderListEntityArray, error)

	// Add new comment to shipment
	SalesOrderShipmentAddComment(sessionId string, shipmentIncrementId string, comment string, email string, includeInEmail string) (bool, error)

	// Add new tracking number
	SalesOrderShipmentAddTrack(sessionId string, shipmentIncrementId string, carrier string, title string, trackNumber string) (int, error)

	// Create new shipment for order
	SalesOrderShipmentCreate(sessionId string, orderIncrementId string, itemsQty *OrderItemIdQtyArray, comment string, email int, includeComment int) (string, error)

	// Retrieve list of allowed carriers for order
	SalesOrderShipmentGetCarriers(sessionId string, orderIncrementId string) (*AssociativeArray, error)

	// Retrieve shipment information
	SalesOrderShipmentInfo(sessionId string, shipmentIncrementId string) (*SalesOrderShipmentEntity, error)

	// Retrieve list of shipments by filters
	SalesOrderShipmentList(sessionId string, filters *Filters) (*SalesOrderShipmentEntityArray, error)

	// Remove tracking number
	SalesOrderShipmentRemoveTrack(sessionId string, shipmentIncrementId string, trackId string) (bool, error)

	// Send shipment info
	SalesOrderShipmentSendInfo(sessionId string, shipmentIncrementId string, comment string) (bool, error)

	// Unhold order
	SalesOrderUnhold(sessionId string, orderIncrementId string) (bool, error)

	// Add coupon code for shopping cart
	ShoppingCartCouponAdd(sessionId string, quoteId int, couponCode string, storeId string) (bool, error)

	// Remove coupon code from shopping cart
	ShoppingCartCouponRemove(sessionId string, quoteId int, storeId string) (bool, error)

	// Create shopping cart
	ShoppingCartCreate(sessionId string, storeId string) (int, error)

	// Set customer's addresses in shopping cart
	ShoppingCartCustomerAddresses(sessionId string, quoteId int, customer *ShoppingCartCustomerAddressEntityArray, storeId string) (bool, error)

	// Set customer for shopping cart
	ShoppingCartCustomerSet(sessionId string, quoteId int, customer *ShoppingCartCustomerEntity, storeId string) (bool, error)

	// Retrieve information about shopping cart
	ShoppingCartInfo(sessionId string, quoteId int, storeId string) (*ShoppingCartInfoEntity, error)

	// Get terms and conditions
	ShoppingCartLicense(sessionId string, quoteId int, storeId string) (*ShoppingCartLicenseEntityArray, error)

	// Create an order from shopping cart
	ShoppingCartOrder(sessionId string, quoteId int, storeId string, licenses *ArrayOfString) (string, error)

	// Get list of available payment methods
	ShoppingCartPaymentList(sessionId string, quoteId int, store string) (*ShoppingCartPaymentMethodResponseEntityArray, error)

	// Set payment method
	ShoppingCartPaymentMethod(sessionId string, quoteId int, method *ShoppingCartPaymentMethodEntity, storeId string) (bool, error)

	// Add product(s) to shopping cart
	ShoppingCartProductAdd(sessionId string, quoteId int, products *ShoppingCartProductEntityArray, storeId string) (bool, error)

	// Get list of products in shopping cart
	ShoppingCartProductList(sessionId string, quoteId int, storeId string) (*ShoppingCartProductResponseEntityArray, error)

	// Move product(s) to customer quote
	ShoppingCartProductMoveToCustomerQuote(sessionId string, quoteId int, products *ShoppingCartProductEntityArray, storeId string) (bool, error)

	// Remove product(s) from shopping cart
	ShoppingCartProductRemove(sessionId string, quoteId int, products *ShoppingCartProductEntityArray, storeId string) (bool, error)

	// Update product(s) quantities in shopping cart
	ShoppingCartProductUpdate(sessionId string, quoteId int, products *ShoppingCartProductEntityArray, storeId string) (bool, error)

	// Get list of available shipping methods
	ShoppingCartShippingList(sessionId string, quoteId int, storeId string) (*ShoppingCartShippingMethodEntityArray, error)

	// Set shipping method
	ShoppingCartShippingMethod(sessionId string, quoteId int, method string, storeId string) (bool, error)

	// Get total prices for shopping cart
	ShoppingCartTotals(sessionId string, quoteId int, storeId string) (*ShoppingCartTotalsEntityArray, error)

	// Start web service session
	StartSession() (string, error)

	// Store view info
	StoreInfo(sessionId string, storeId string) (*StoreEntity, error)

	// List of stores
	StoreList(sessionId string) (*StoreEntityArray, error)
}

// ArrayCoordinate was auto-generated from WSDL.
type ArrayCoordinate string

// Base64 was auto-generated from WSDL.
type Base64 []byte

// 	   'Array' is a complex type for accessors identified by position
//
type Array struct {
}

// ArrayOfApiMethods was auto-generated from WSDL.
type ArrayOfApiMethods struct {
	Items []*ApiMethodEntity `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
}

// ArrayOfApis was auto-generated from WSDL.
type ArrayOfApis struct {
	Items []*ApiEntity `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
}

// ArrayOfCatalogCategoryEntities was auto-generated from WSDL.
type ArrayOfCatalogCategoryEntities struct {
	Items []*CatalogCategoryEntity `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
}

// ArrayOfCatalogCategoryEntitiesNoChildren was auto-generated
// from WSDL.
type ArrayOfCatalogCategoryEntitiesNoChildren struct {
	Items []*CatalogCategoryEntityNoChildren `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
}

// ArrayOfExistsFaltures was auto-generated from WSDL.
type ArrayOfExistsFaltures struct {
	Items []*ExistsFaltureEntity `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
}

// ArrayOfInt was auto-generated from WSDL.
type ArrayOfInt struct {
	Items []*int `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
}

// ArrayOfString was auto-generated from WSDL.
type ArrayOfString struct {
	Items []*string `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
}

// ENTITIES was auto-generated from WSDL.
type ENTITIES struct {
}

// ENTITY was auto-generated from WSDL.
type ENTITY struct {
}

// ID was auto-generated from WSDL.
type ID struct {
}

// IDREF was auto-generated from WSDL.
type IDREF struct {
}

// IDREFS was auto-generated from WSDL.
type IDREFS struct {
}

// NCName was auto-generated from WSDL.
type NCName struct {
}

// NMTOKEN was auto-generated from WSDL.
type NMTOKEN struct {
}

// NMTOKENS was auto-generated from WSDL.
type NMTOKENS struct {
}

// NOTATION was auto-generated from WSDL.
type NOTATION struct {
}

// Name was auto-generated from WSDL.
type Name struct {
}

// QName was auto-generated from WSDL.
type QName struct {
}

// Struct was auto-generated from WSDL.
type Struct struct {
}

// AnyURI was auto-generated from WSDL.
type AnyURI struct {
}

// ApiEntity was auto-generated from WSDL.
type ApiEntity struct {
	Title   *string            `xml:"title,omitempty" json:"title,omitempty" yaml:"title,omitempty"`
	Name    *string            `xml:"name,omitempty" json:"name,omitempty" yaml:"name,omitempty"`
	Aliases *ArrayOfString     `xml:"aliases,omitempty" json:"aliases,omitempty" yaml:"aliases,omitempty"`
	Methods *ArrayOfApiMethods `xml:"methods,omitempty" json:"methods,omitempty" yaml:"methods,omitempty"`
}

// ApiMethodEntity was auto-generated from WSDL.
type ApiMethodEntity struct {
	Title   *string        `xml:"title,omitempty" json:"title,omitempty" yaml:"title,omitempty"`
	Path    *string        `xml:"path,omitempty" json:"path,omitempty" yaml:"path,omitempty"`
	Name    *string        `xml:"name,omitempty" json:"name,omitempty" yaml:"name,omitempty"`
	Aliases *ArrayOfString `xml:"aliases,omitempty" json:"aliases,omitempty" yaml:"aliases,omitempty"`
}

// AssociativeArray was auto-generated from WSDL.
type AssociativeArray struct {
	Items []*AssociativeEntity `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
}

// AssociativeEntity was auto-generated from WSDL.
type AssociativeEntity struct {
	Key   *string `xml:"key,omitempty" json:"key,omitempty" yaml:"key,omitempty"`
	Value *string `xml:"value,omitempty" json:"value,omitempty" yaml:"value,omitempty"`
}

// AssociativeMultiArray was auto-generated from WSDL.
type AssociativeMultiArray struct {
	Items []*AssociativeMultiEntity `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
}

// AssociativeMultiEntity was auto-generated from WSDL.
type AssociativeMultiEntity struct {
	Key   *string        `xml:"key,omitempty" json:"key,omitempty" yaml:"key,omitempty"`
	Value *ArrayOfString `xml:"value,omitempty" json:"value,omitempty" yaml:"value,omitempty"`
}

// Base64Binary was auto-generated from WSDL.
type Base64Binary struct {
}

// Boolean was auto-generated from WSDL.
type Boolean struct {
}

// Byte was auto-generated from WSDL.
type Byte struct {
}

// CatalogAssignedProduct was auto-generated from WSDL.
type CatalogAssignedProduct struct {
	Product_id *int    `xml:"product_id,omitempty" json:"product_id,omitempty" yaml:"product_id,omitempty"`
	Type       *string `xml:"type,omitempty" json:"type,omitempty" yaml:"type,omitempty"`
	Set        *int    `xml:"set,omitempty" json:"set,omitempty" yaml:"set,omitempty"`
	Sku        *string `xml:"sku,omitempty" json:"sku,omitempty" yaml:"sku,omitempty"`
	Position   *int    `xml:"position,omitempty" json:"position,omitempty" yaml:"position,omitempty"`
}

// CatalogAssignedProductArray was auto-generated from WSDL.
type CatalogAssignedProductArray struct {
	Items []*CatalogAssignedProduct `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
}

// CatalogAttributeEntity was auto-generated from WSDL.
type CatalogAttributeEntity struct {
	Attribute_id *int    `xml:"attribute_id,omitempty" json:"attribute_id,omitempty" yaml:"attribute_id,omitempty"`
	Code         *string `xml:"code,omitempty" json:"code,omitempty" yaml:"code,omitempty"`
	Type         *string `xml:"type,omitempty" json:"type,omitempty" yaml:"type,omitempty"`
	Required     *string `xml:"required,omitempty" json:"required,omitempty" yaml:"required,omitempty"`
	Scope        *string `xml:"scope,omitempty" json:"scope,omitempty" yaml:"scope,omitempty"`
}

// CatalogAttributeEntityArray was auto-generated from WSDL.
type CatalogAttributeEntityArray struct {
	Items []*CatalogAttributeEntity `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
}

// CatalogAttributeOptionEntity was auto-generated from WSDL.
type CatalogAttributeOptionEntity struct {
	Label *string `xml:"label,omitempty" json:"label,omitempty" yaml:"label,omitempty"`
	Value *string `xml:"value,omitempty" json:"value,omitempty" yaml:"value,omitempty"`
}

// CatalogAttributeOptionEntityArray was auto-generated from WSDL.
type CatalogAttributeOptionEntityArray struct {
	Items []*CatalogAttributeOptionEntity `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
}

// CatalogCategoryEntity was auto-generated from WSDL.
type CatalogCategoryEntity struct {
	Category_id *int                            `xml:"category_id,omitempty" json:"category_id,omitempty" yaml:"category_id,omitempty"`
	Parent_id   *int                            `xml:"parent_id,omitempty" json:"parent_id,omitempty" yaml:"parent_id,omitempty"`
	Name        *string                         `xml:"name,omitempty" json:"name,omitempty" yaml:"name,omitempty"`
	Is_active   *int                            `xml:"is_active,omitempty" json:"is_active,omitempty" yaml:"is_active,omitempty"`
	Position    *int                            `xml:"position,omitempty" json:"position,omitempty" yaml:"position,omitempty"`
	Level       *int                            `xml:"level,omitempty" json:"level,omitempty" yaml:"level,omitempty"`
	Children    *ArrayOfCatalogCategoryEntities `xml:"children,omitempty" json:"children,omitempty" yaml:"children,omitempty"`
}

// CatalogCategoryEntityCreate was auto-generated from WSDL.
type CatalogCategoryEntityCreate struct {
	Name                 *string        `xml:"name,omitempty" json:"name,omitempty" yaml:"name,omitempty"`
	Is_active            *int           `xml:"is_active,omitempty" json:"is_active,omitempty" yaml:"is_active,omitempty"`
	Position             *int           `xml:"position,omitempty" json:"position,omitempty" yaml:"position,omitempty"`
	Available_sort_by    *ArrayOfString `xml:"available_sort_by,omitempty" json:"available_sort_by,omitempty" yaml:"available_sort_by,omitempty"`
	Custom_design        *string        `xml:"custom_design,omitempty" json:"custom_design,omitempty" yaml:"custom_design,omitempty"`
	Custom_design_apply  *int           `xml:"custom_design_apply,omitempty" json:"custom_design_apply,omitempty" yaml:"custom_design_apply,omitempty"`
	Custom_design_from   *string        `xml:"custom_design_from,omitempty" json:"custom_design_from,omitempty" yaml:"custom_design_from,omitempty"`
	Custom_design_to     *string        `xml:"custom_design_to,omitempty" json:"custom_design_to,omitempty" yaml:"custom_design_to,omitempty"`
	Custom_layout_update *string        `xml:"custom_layout_update,omitempty" json:"custom_layout_update,omitempty" yaml:"custom_layout_update,omitempty"`
	Default_sort_by      *string        `xml:"default_sort_by,omitempty" json:"default_sort_by,omitempty" yaml:"default_sort_by,omitempty"`
	Description          *string        `xml:"description,omitempty" json:"description,omitempty" yaml:"description,omitempty"`
	Display_mode         *string        `xml:"display_mode,omitempty" json:"display_mode,omitempty" yaml:"display_mode,omitempty"`
	Is_anchor            *int           `xml:"is_anchor,omitempty" json:"is_anchor,omitempty" yaml:"is_anchor,omitempty"`
	Landing_page         *int           `xml:"landing_page,omitempty" json:"landing_page,omitempty" yaml:"landing_page,omitempty"`
	Meta_description     *string        `xml:"meta_description,omitempty" json:"meta_description,omitempty" yaml:"meta_description,omitempty"`
	Meta_keywords        *string        `xml:"meta_keywords,omitempty" json:"meta_keywords,omitempty" yaml:"meta_keywords,omitempty"`
	Meta_title           *string        `xml:"meta_title,omitempty" json:"meta_title,omitempty" yaml:"meta_title,omitempty"`
	Page_layout          *string        `xml:"page_layout,omitempty" json:"page_layout,omitempty" yaml:"page_layout,omitempty"`
	Url_key              *string        `xml:"url_key,omitempty" json:"url_key,omitempty" yaml:"url_key,omitempty"`
	Include_in_menu      *int           `xml:"include_in_menu,omitempty" json:"include_in_menu,omitempty" yaml:"include_in_menu,omitempty"`
}

// CatalogCategoryEntityNoChildren was auto-generated from WSDL.
type CatalogCategoryEntityNoChildren struct {
	Category_id *int    `xml:"category_id,omitempty" json:"category_id,omitempty" yaml:"category_id,omitempty"`
	Parent_id   *int    `xml:"parent_id,omitempty" json:"parent_id,omitempty" yaml:"parent_id,omitempty"`
	Name        *string `xml:"name,omitempty" json:"name,omitempty" yaml:"name,omitempty"`
	Is_active   *int    `xml:"is_active,omitempty" json:"is_active,omitempty" yaml:"is_active,omitempty"`
	Position    *int    `xml:"position,omitempty" json:"position,omitempty" yaml:"position,omitempty"`
	Level       *int    `xml:"level,omitempty" json:"level,omitempty" yaml:"level,omitempty"`
}

// CatalogCategoryInfo was auto-generated from WSDL.
type CatalogCategoryInfo struct {
	Category_id                *string        `xml:"category_id,omitempty" json:"category_id,omitempty" yaml:"category_id,omitempty"`
	Is_active                  *int           `xml:"is_active,omitempty" json:"is_active,omitempty" yaml:"is_active,omitempty"`
	Position                   *string        `xml:"position,omitempty" json:"position,omitempty" yaml:"position,omitempty"`
	Level                      *string        `xml:"level,omitempty" json:"level,omitempty" yaml:"level,omitempty"`
	Parent_id                  *string        `xml:"parent_id,omitempty" json:"parent_id,omitempty" yaml:"parent_id,omitempty"`
	All_children               *string        `xml:"all_children,omitempty" json:"all_children,omitempty" yaml:"all_children,omitempty"`
	Children                   *string        `xml:"children,omitempty" json:"children,omitempty" yaml:"children,omitempty"`
	Created_at                 *string        `xml:"created_at,omitempty" json:"created_at,omitempty" yaml:"created_at,omitempty"`
	Updated_at                 *string        `xml:"updated_at,omitempty" json:"updated_at,omitempty" yaml:"updated_at,omitempty"`
	Name                       *string        `xml:"name,omitempty" json:"name,omitempty" yaml:"name,omitempty"`
	Url_key                    *string        `xml:"url_key,omitempty" json:"url_key,omitempty" yaml:"url_key,omitempty"`
	Description                *string        `xml:"description,omitempty" json:"description,omitempty" yaml:"description,omitempty"`
	Meta_title                 *string        `xml:"meta_title,omitempty" json:"meta_title,omitempty" yaml:"meta_title,omitempty"`
	Meta_keywords              *string        `xml:"meta_keywords,omitempty" json:"meta_keywords,omitempty" yaml:"meta_keywords,omitempty"`
	Meta_description           *string        `xml:"meta_description,omitempty" json:"meta_description,omitempty" yaml:"meta_description,omitempty"`
	Path                       *string        `xml:"path,omitempty" json:"path,omitempty" yaml:"path,omitempty"`
	Url_path                   *string        `xml:"url_path,omitempty" json:"url_path,omitempty" yaml:"url_path,omitempty"`
	Children_count             *int           `xml:"children_count,omitempty" json:"children_count,omitempty" yaml:"children_count,omitempty"`
	Display_mode               *string        `xml:"display_mode,omitempty" json:"display_mode,omitempty" yaml:"display_mode,omitempty"`
	Is_anchor                  *int           `xml:"is_anchor,omitempty" json:"is_anchor,omitempty" yaml:"is_anchor,omitempty"`
	Available_sort_by          *ArrayOfString `xml:"available_sort_by,omitempty" json:"available_sort_by,omitempty" yaml:"available_sort_by,omitempty"`
	Custom_design              *string        `xml:"custom_design,omitempty" json:"custom_design,omitempty" yaml:"custom_design,omitempty"`
	Custom_design_apply        *string        `xml:"custom_design_apply,omitempty" json:"custom_design_apply,omitempty" yaml:"custom_design_apply,omitempty"`
	Custom_design_from         *string        `xml:"custom_design_from,omitempty" json:"custom_design_from,omitempty" yaml:"custom_design_from,omitempty"`
	Custom_design_to           *string        `xml:"custom_design_to,omitempty" json:"custom_design_to,omitempty" yaml:"custom_design_to,omitempty"`
	Page_layout                *string        `xml:"page_layout,omitempty" json:"page_layout,omitempty" yaml:"page_layout,omitempty"`
	Custom_layout_update       *string        `xml:"custom_layout_update,omitempty" json:"custom_layout_update,omitempty" yaml:"custom_layout_update,omitempty"`
	Default_sort_by            *string        `xml:"default_sort_by,omitempty" json:"default_sort_by,omitempty" yaml:"default_sort_by,omitempty"`
	Landing_page               *int           `xml:"landing_page,omitempty" json:"landing_page,omitempty" yaml:"landing_page,omitempty"`
	Include_in_menu            *int           `xml:"include_in_menu,omitempty" json:"include_in_menu,omitempty" yaml:"include_in_menu,omitempty"`
	Custom_use_parent_settings *int           `xml:"custom_use_parent_settings,omitempty" json:"custom_use_parent_settings,omitempty" yaml:"custom_use_parent_settings,omitempty"`
	Custom_apply_to_products   *int           `xml:"custom_apply_to_products,omitempty" json:"custom_apply_to_products,omitempty" yaml:"custom_apply_to_products,omitempty"`
	Filter_price_range         *string        `xml:"filter_price_range,omitempty" json:"filter_price_range,omitempty" yaml:"filter_price_range,omitempty"`
}

// CatalogCategoryTree was auto-generated from WSDL.
type CatalogCategoryTree struct {
	Category_id *int                            `xml:"category_id,omitempty" json:"category_id,omitempty" yaml:"category_id,omitempty"`
	Parent_id   *int                            `xml:"parent_id,omitempty" json:"parent_id,omitempty" yaml:"parent_id,omitempty"`
	Name        *string                         `xml:"name,omitempty" json:"name,omitempty" yaml:"name,omitempty"`
	Position    *int                            `xml:"position,omitempty" json:"position,omitempty" yaml:"position,omitempty"`
	Level       *int                            `xml:"level,omitempty" json:"level,omitempty" yaml:"level,omitempty"`
	Children    *ArrayOfCatalogCategoryEntities `xml:"children,omitempty" json:"children,omitempty" yaml:"children,omitempty"`
}

// CatalogInventoryStockItemEntity was auto-generated from WSDL.
type CatalogInventoryStockItemEntity struct {
	Product_id  *string `xml:"product_id,omitempty" json:"product_id,omitempty" yaml:"product_id,omitempty"`
	Sku         *string `xml:"sku,omitempty" json:"sku,omitempty" yaml:"sku,omitempty"`
	Qty         *string `xml:"qty,omitempty" json:"qty,omitempty" yaml:"qty,omitempty"`
	Is_in_stock *string `xml:"is_in_stock,omitempty" json:"is_in_stock,omitempty" yaml:"is_in_stock,omitempty"`
}

// CatalogInventoryStockItemEntityArray was auto-generated from
// WSDL.
type CatalogInventoryStockItemEntityArray struct {
	Items []*CatalogInventoryStockItemEntity `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
}

// CatalogInventoryStockItemUpdateEntity was auto-generated from
// WSDL.
type CatalogInventoryStockItemUpdateEntity struct {
	Qty                         *string `xml:"qty,omitempty" json:"qty,omitempty" yaml:"qty,omitempty"`
	Is_in_stock                 *int    `xml:"is_in_stock,omitempty" json:"is_in_stock,omitempty" yaml:"is_in_stock,omitempty"`
	Manage_stock                *int    `xml:"manage_stock,omitempty" json:"manage_stock,omitempty" yaml:"manage_stock,omitempty"`
	Use_config_manage_stock     *int    `xml:"use_config_manage_stock,omitempty" json:"use_config_manage_stock,omitempty" yaml:"use_config_manage_stock,omitempty"`
	Min_qty                     *int    `xml:"min_qty,omitempty" json:"min_qty,omitempty" yaml:"min_qty,omitempty"`
	Use_config_min_qty          *int    `xml:"use_config_min_qty,omitempty" json:"use_config_min_qty,omitempty" yaml:"use_config_min_qty,omitempty"`
	Min_sale_qty                *int    `xml:"min_sale_qty,omitempty" json:"min_sale_qty,omitempty" yaml:"min_sale_qty,omitempty"`
	Use_config_min_sale_qty     *int    `xml:"use_config_min_sale_qty,omitempty" json:"use_config_min_sale_qty,omitempty" yaml:"use_config_min_sale_qty,omitempty"`
	Max_sale_qty                *int    `xml:"max_sale_qty,omitempty" json:"max_sale_qty,omitempty" yaml:"max_sale_qty,omitempty"`
	Use_config_max_sale_qty     *int    `xml:"use_config_max_sale_qty,omitempty" json:"use_config_max_sale_qty,omitempty" yaml:"use_config_max_sale_qty,omitempty"`
	Is_qty_decimal              *int    `xml:"is_qty_decimal,omitempty" json:"is_qty_decimal,omitempty" yaml:"is_qty_decimal,omitempty"`
	Backorders                  *int    `xml:"backorders,omitempty" json:"backorders,omitempty" yaml:"backorders,omitempty"`
	Use_config_backorders       *int    `xml:"use_config_backorders,omitempty" json:"use_config_backorders,omitempty" yaml:"use_config_backorders,omitempty"`
	Notify_stock_qty            *int    `xml:"notify_stock_qty,omitempty" json:"notify_stock_qty,omitempty" yaml:"notify_stock_qty,omitempty"`
	Use_config_notify_stock_qty *int    `xml:"use_config_notify_stock_qty,omitempty" json:"use_config_notify_stock_qty,omitempty" yaml:"use_config_notify_stock_qty,omitempty"`
}

// CatalogInventoryStockItemUpdateEntityArray was auto-generated
// from WSDL.
type CatalogInventoryStockItemUpdateEntityArray struct {
	Items []*CatalogInventoryStockItemUpdateEntity `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
}

// CatalogProductAdditionalAttributesEntity was auto-generated
// from WSDL.
type CatalogProductAdditionalAttributesEntity struct {
	Multi_data  *AssociativeMultiArray `xml:"multi_data,omitempty" json:"multi_data,omitempty" yaml:"multi_data,omitempty"`
	Single_data *AssociativeArray      `xml:"single_data,omitempty" json:"single_data,omitempty" yaml:"single_data,omitempty"`
}

// CatalogProductAttributeEntity was auto-generated from WSDL.
type CatalogProductAttributeEntity struct {
	Attribute_id                  string                                     `xml:"attribute_id" json:"attribute_id" yaml:"attribute_id"`
	Attribute_code                string                                     `xml:"attribute_code" json:"attribute_code" yaml:"attribute_code"`
	Frontend_input                string                                     `xml:"frontend_input" json:"frontend_input" yaml:"frontend_input"`
	Scope                         *string                                    `xml:"scope,omitempty" json:"scope,omitempty" yaml:"scope,omitempty"`
	Default_value                 *string                                    `xml:"default_value,omitempty" json:"default_value,omitempty" yaml:"default_value,omitempty"`
	Is_unique                     *int                                       `xml:"is_unique,omitempty" json:"is_unique,omitempty" yaml:"is_unique,omitempty"`
	Is_required                   *int                                       `xml:"is_required,omitempty" json:"is_required,omitempty" yaml:"is_required,omitempty"`
	Apply_to                      *ArrayOfString                             `xml:"apply_to,omitempty" json:"apply_to,omitempty" yaml:"apply_to,omitempty"`
	Is_configurable               *int                                       `xml:"is_configurable,omitempty" json:"is_configurable,omitempty" yaml:"is_configurable,omitempty"`
	Is_searchable                 *int                                       `xml:"is_searchable,omitempty" json:"is_searchable,omitempty" yaml:"is_searchable,omitempty"`
	Is_visible_in_advanced_search *int                                       `xml:"is_visible_in_advanced_search,omitempty" json:"is_visible_in_advanced_search,omitempty" yaml:"is_visible_in_advanced_search,omitempty"`
	Is_comparable                 *int                                       `xml:"is_comparable,omitempty" json:"is_comparable,omitempty" yaml:"is_comparable,omitempty"`
	Is_used_for_promo_rules       *int                                       `xml:"is_used_for_promo_rules,omitempty" json:"is_used_for_promo_rules,omitempty" yaml:"is_used_for_promo_rules,omitempty"`
	Is_visible_on_front           *int                                       `xml:"is_visible_on_front,omitempty" json:"is_visible_on_front,omitempty" yaml:"is_visible_on_front,omitempty"`
	Used_in_product_listing       *int                                       `xml:"used_in_product_listing,omitempty" json:"used_in_product_listing,omitempty" yaml:"used_in_product_listing,omitempty"`
	Additional_fields             *AssociativeArray                          `xml:"additional_fields,omitempty" json:"additional_fields,omitempty" yaml:"additional_fields,omitempty"`
	Options                       *CatalogAttributeOptionEntityArray         `xml:"options,omitempty" json:"options,omitempty" yaml:"options,omitempty"`
	Frontend_label                *CatalogProductAttributeFrontendLabelArray `xml:"frontend_label" json:"frontend_label" yaml:"frontend_label"`
}

// CatalogProductAttributeEntityToCreate was auto-generated from
// WSDL.
type CatalogProductAttributeEntityToCreate struct {
	Attribute_code                string                                     `xml:"attribute_code" json:"attribute_code" yaml:"attribute_code"`
	Frontend_input                string                                     `xml:"frontend_input" json:"frontend_input" yaml:"frontend_input"`
	Scope                         *string                                    `xml:"scope,omitempty" json:"scope,omitempty" yaml:"scope,omitempty"`
	Default_value                 *string                                    `xml:"default_value,omitempty" json:"default_value,omitempty" yaml:"default_value,omitempty"`
	Is_unique                     *int                                       `xml:"is_unique,omitempty" json:"is_unique,omitempty" yaml:"is_unique,omitempty"`
	Is_required                   *int                                       `xml:"is_required,omitempty" json:"is_required,omitempty" yaml:"is_required,omitempty"`
	Apply_to                      *ArrayOfString                             `xml:"apply_to,omitempty" json:"apply_to,omitempty" yaml:"apply_to,omitempty"`
	Is_configurable               *int                                       `xml:"is_configurable,omitempty" json:"is_configurable,omitempty" yaml:"is_configurable,omitempty"`
	Is_searchable                 *int                                       `xml:"is_searchable,omitempty" json:"is_searchable,omitempty" yaml:"is_searchable,omitempty"`
	Is_visible_in_advanced_search *int                                       `xml:"is_visible_in_advanced_search,omitempty" json:"is_visible_in_advanced_search,omitempty" yaml:"is_visible_in_advanced_search,omitempty"`
	Is_comparable                 *int                                       `xml:"is_comparable,omitempty" json:"is_comparable,omitempty" yaml:"is_comparable,omitempty"`
	Is_used_for_promo_rules       *int                                       `xml:"is_used_for_promo_rules,omitempty" json:"is_used_for_promo_rules,omitempty" yaml:"is_used_for_promo_rules,omitempty"`
	Is_visible_on_front           *int                                       `xml:"is_visible_on_front,omitempty" json:"is_visible_on_front,omitempty" yaml:"is_visible_on_front,omitempty"`
	Used_in_product_listing       *int                                       `xml:"used_in_product_listing,omitempty" json:"used_in_product_listing,omitempty" yaml:"used_in_product_listing,omitempty"`
	Additional_fields             *AssociativeArray                          `xml:"additional_fields,omitempty" json:"additional_fields,omitempty" yaml:"additional_fields,omitempty"`
	Frontend_label                *CatalogProductAttributeFrontendLabelArray `xml:"frontend_label" json:"frontend_label" yaml:"frontend_label"`
}

// CatalogProductAttributeEntityToUpdate was auto-generated from
// WSDL.
type CatalogProductAttributeEntityToUpdate struct {
	Scope                         *string                                    `xml:"scope,omitempty" json:"scope,omitempty" yaml:"scope,omitempty"`
	Default_value                 *string                                    `xml:"default_value,omitempty" json:"default_value,omitempty" yaml:"default_value,omitempty"`
	Is_unique                     *int                                       `xml:"is_unique,omitempty" json:"is_unique,omitempty" yaml:"is_unique,omitempty"`
	Is_required                   *int                                       `xml:"is_required,omitempty" json:"is_required,omitempty" yaml:"is_required,omitempty"`
	Apply_to                      *ArrayOfString                             `xml:"apply_to,omitempty" json:"apply_to,omitempty" yaml:"apply_to,omitempty"`
	Is_configurable               *int                                       `xml:"is_configurable,omitempty" json:"is_configurable,omitempty" yaml:"is_configurable,omitempty"`
	Is_searchable                 *int                                       `xml:"is_searchable,omitempty" json:"is_searchable,omitempty" yaml:"is_searchable,omitempty"`
	Is_visible_in_advanced_search *int                                       `xml:"is_visible_in_advanced_search,omitempty" json:"is_visible_in_advanced_search,omitempty" yaml:"is_visible_in_advanced_search,omitempty"`
	Is_comparable                 *int                                       `xml:"is_comparable,omitempty" json:"is_comparable,omitempty" yaml:"is_comparable,omitempty"`
	Is_used_for_promo_rules       *int                                       `xml:"is_used_for_promo_rules,omitempty" json:"is_used_for_promo_rules,omitempty" yaml:"is_used_for_promo_rules,omitempty"`
	Is_visible_on_front           *int                                       `xml:"is_visible_on_front,omitempty" json:"is_visible_on_front,omitempty" yaml:"is_visible_on_front,omitempty"`
	Used_in_product_listing       *int                                       `xml:"used_in_product_listing,omitempty" json:"used_in_product_listing,omitempty" yaml:"used_in_product_listing,omitempty"`
	Additional_fields             *AssociativeArray                          `xml:"additional_fields,omitempty" json:"additional_fields,omitempty" yaml:"additional_fields,omitempty"`
	Frontend_label                *CatalogProductAttributeFrontendLabelArray `xml:"frontend_label" json:"frontend_label" yaml:"frontend_label"`
}

// CatalogProductAttributeFrontendLabelArray was auto-generated
// from WSDL.
type CatalogProductAttributeFrontendLabelArray struct {
	Items []*CatalogProductAttributeFrontendLabelEntity `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
}

// CatalogProductAttributeFrontendLabelEntity was auto-generated
// from WSDL.
type CatalogProductAttributeFrontendLabelEntity struct {
	Store_id *string `xml:"store_id,omitempty" json:"store_id,omitempty" yaml:"store_id,omitempty"`
	Label    *string `xml:"label,omitempty" json:"label,omitempty" yaml:"label,omitempty"`
}

// CatalogProductAttributeMediaCreateEntity was auto-generated
// from WSDL.
type CatalogProductAttributeMediaCreateEntity struct {
	File     *CatalogProductImageFileEntity `xml:"file,omitempty" json:"file,omitempty" yaml:"file,omitempty"`
	Label    *string                        `xml:"label,omitempty" json:"label,omitempty" yaml:"label,omitempty"`
	Position *string                        `xml:"position,omitempty" json:"position,omitempty" yaml:"position,omitempty"`
	Types    *ArrayOfString                 `xml:"types,omitempty" json:"types,omitempty" yaml:"types,omitempty"`
	Exclude  *string                        `xml:"exclude,omitempty" json:"exclude,omitempty" yaml:"exclude,omitempty"`
	Remove   *string                        `xml:"remove,omitempty" json:"remove,omitempty" yaml:"remove,omitempty"`
}

// CatalogProductAttributeMediaTypeEntity was auto-generated from
// WSDL.
type CatalogProductAttributeMediaTypeEntity struct {
	Code  *string `xml:"code,omitempty" json:"code,omitempty" yaml:"code,omitempty"`
	Scope *string `xml:"scope,omitempty" json:"scope,omitempty" yaml:"scope,omitempty"`
}

// CatalogProductAttributeMediaTypeEntityArray was auto-generated
// from WSDL.
type CatalogProductAttributeMediaTypeEntityArray struct {
	Items []*CatalogProductAttributeMediaTypeEntity `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
}

// CatalogProductAttributeOptionEntityToAdd was auto-generated
// from WSDL.
type CatalogProductAttributeOptionEntityToAdd struct {
	Label      *CatalogProductAttributeOptionLabelArray `xml:"label,omitempty" json:"label,omitempty" yaml:"label,omitempty"`
	Order      *int                                     `xml:"order,omitempty" json:"order,omitempty" yaml:"order,omitempty"`
	Is_default *int                                     `xml:"is_default,omitempty" json:"is_default,omitempty" yaml:"is_default,omitempty"`
}

// CatalogProductAttributeOptionLabelArray was auto-generated from
// WSDL.
type CatalogProductAttributeOptionLabelArray struct {
	Items []*CatalogProductAttributeOptionLabelEntity `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
}

// CatalogProductAttributeOptionLabelEntity was auto-generated
// from WSDL.
type CatalogProductAttributeOptionLabelEntity struct {
	Store_id *ArrayOfString `xml:"store_id,omitempty" json:"store_id,omitempty" yaml:"store_id,omitempty"`
	Value    *string        `xml:"value,omitempty" json:"value,omitempty" yaml:"value,omitempty"`
}

// CatalogProductAttributeSetEntity was auto-generated from WSDL.
type CatalogProductAttributeSetEntity struct {
	Set_id *int    `xml:"set_id,omitempty" json:"set_id,omitempty" yaml:"set_id,omitempty"`
	Name   *string `xml:"name,omitempty" json:"name,omitempty" yaml:"name,omitempty"`
}

// CatalogProductAttributeSetEntityArray was auto-generated from
// WSDL.
type CatalogProductAttributeSetEntityArray struct {
	Items []*CatalogProductAttributeSetEntity `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
}

// CatalogProductCreateEntity was auto-generated from WSDL.
type CatalogProductCreateEntity struct {
	Categories             *ArrayOfString                            `xml:"categories,omitempty" json:"categories,omitempty" yaml:"categories,omitempty"`
	Websites               *ArrayOfString                            `xml:"websites,omitempty" json:"websites,omitempty" yaml:"websites,omitempty"`
	Name                   *string                                   `xml:"name,omitempty" json:"name,omitempty" yaml:"name,omitempty"`
	Description            *string                                   `xml:"description,omitempty" json:"description,omitempty" yaml:"description,omitempty"`
	Short_description      *string                                   `xml:"short_description,omitempty" json:"short_description,omitempty" yaml:"short_description,omitempty"`
	Weight                 *string                                   `xml:"weight,omitempty" json:"weight,omitempty" yaml:"weight,omitempty"`
	Status                 *string                                   `xml:"status,omitempty" json:"status,omitempty" yaml:"status,omitempty"`
	Url_key                *string                                   `xml:"url_key,omitempty" json:"url_key,omitempty" yaml:"url_key,omitempty"`
	Url_path               *string                                   `xml:"url_path,omitempty" json:"url_path,omitempty" yaml:"url_path,omitempty"`
	Visibility             *string                                   `xml:"visibility,omitempty" json:"visibility,omitempty" yaml:"visibility,omitempty"`
	Category_ids           *ArrayOfString                            `xml:"category_ids,omitempty" json:"category_ids,omitempty" yaml:"category_ids,omitempty"`
	Website_ids            *ArrayOfString                            `xml:"website_ids,omitempty" json:"website_ids,omitempty" yaml:"website_ids,omitempty"`
	Has_options            *string                                   `xml:"has_options,omitempty" json:"has_options,omitempty" yaml:"has_options,omitempty"`
	Gift_message_available *string                                   `xml:"gift_message_available,omitempty" json:"gift_message_available,omitempty" yaml:"gift_message_available,omitempty"`
	Price                  *string                                   `xml:"price,omitempty" json:"price,omitempty" yaml:"price,omitempty"`
	Special_price          *string                                   `xml:"special_price,omitempty" json:"special_price,omitempty" yaml:"special_price,omitempty"`
	Special_from_date      *string                                   `xml:"special_from_date,omitempty" json:"special_from_date,omitempty" yaml:"special_from_date,omitempty"`
	Special_to_date        *string                                   `xml:"special_to_date,omitempty" json:"special_to_date,omitempty" yaml:"special_to_date,omitempty"`
	Tax_class_id           *string                                   `xml:"tax_class_id,omitempty" json:"tax_class_id,omitempty" yaml:"tax_class_id,omitempty"`
	Tier_price             *CatalogProductTierPriceEntityArray       `xml:"tier_price,omitempty" json:"tier_price,omitempty" yaml:"tier_price,omitempty"`
	Meta_title             *string                                   `xml:"meta_title,omitempty" json:"meta_title,omitempty" yaml:"meta_title,omitempty"`
	Meta_keyword           *string                                   `xml:"meta_keyword,omitempty" json:"meta_keyword,omitempty" yaml:"meta_keyword,omitempty"`
	Meta_description       *string                                   `xml:"meta_description,omitempty" json:"meta_description,omitempty" yaml:"meta_description,omitempty"`
	Custom_design          *string                                   `xml:"custom_design,omitempty" json:"custom_design,omitempty" yaml:"custom_design,omitempty"`
	Custom_layout_update   *string                                   `xml:"custom_layout_update,omitempty" json:"custom_layout_update,omitempty" yaml:"custom_layout_update,omitempty"`
	Options_container      *string                                   `xml:"options_container,omitempty" json:"options_container,omitempty" yaml:"options_container,omitempty"`
	Additional_attributes  *CatalogProductAdditionalAttributesEntity `xml:"additional_attributes,omitempty" json:"additional_attributes,omitempty" yaml:"additional_attributes,omitempty"`
	Stock_data             *CatalogInventoryStockItemUpdateEntity    `xml:"stock_data,omitempty" json:"stock_data,omitempty" yaml:"stock_data,omitempty"`
}

// CatalogProductCreateEntityArray was auto-generated from WSDL.
type CatalogProductCreateEntityArray struct {
	Items []*CatalogProductCreateEntity `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
}

// CatalogProductCustomOptionAdditionalFieldsArray was auto-generated
// from WSDL.
type CatalogProductCustomOptionAdditionalFieldsArray struct {
	Items []*CatalogProductCustomOptionAdditionalFieldsEntity `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
}

// CatalogProductCustomOptionAdditionalFieldsEntity was auto-generated
// from WSDL.
type CatalogProductCustomOptionAdditionalFieldsEntity struct {
	Title          *string `xml:"title,omitempty" json:"title,omitempty" yaml:"title,omitempty"`
	Price          *string `xml:"price,omitempty" json:"price,omitempty" yaml:"price,omitempty"`
	Price_type     *string `xml:"price_type,omitempty" json:"price_type,omitempty" yaml:"price_type,omitempty"`
	Sku            *string `xml:"sku,omitempty" json:"sku,omitempty" yaml:"sku,omitempty"`
	Max_characters *string `xml:"max_characters,omitempty" json:"max_characters,omitempty" yaml:"max_characters,omitempty"`
	Sort_order     *string `xml:"sort_order,omitempty" json:"sort_order,omitempty" yaml:"sort_order,omitempty"`
	File_extension *string `xml:"file_extension,omitempty" json:"file_extension,omitempty" yaml:"file_extension,omitempty"`
	Image_size_x   *string `xml:"image_size_x,omitempty" json:"image_size_x,omitempty" yaml:"image_size_x,omitempty"`
	Image_size_y   *string `xml:"image_size_y,omitempty" json:"image_size_y,omitempty" yaml:"image_size_y,omitempty"`
	Value_id       *string `xml:"value_id,omitempty" json:"value_id,omitempty" yaml:"value_id,omitempty"`
}

// CatalogProductCustomOptionInfoEntity was auto-generated from
// WSDL.
type CatalogProductCustomOptionInfoEntity struct {
	Title             string                                           `xml:"title" json:"title" yaml:"title"`
	Type              string                                           `xml:"type" json:"type" yaml:"type"`
	Sort_order        string                                           `xml:"sort_order" json:"sort_order" yaml:"sort_order"`
	Is_require        int                                              `xml:"is_require" json:"is_require" yaml:"is_require"`
	Additional_fields *CatalogProductCustomOptionAdditionalFieldsArray `xml:"additional_fields" json:"additional_fields" yaml:"additional_fields"`
}

// CatalogProductCustomOptionListArray was auto-generated from
// WSDL.
type CatalogProductCustomOptionListArray struct {
	Items []*CatalogProductCustomOptionListEntity `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
}

// CatalogProductCustomOptionListEntity was auto-generated from
// WSDL.
type CatalogProductCustomOptionListEntity struct {
	Option_id  string `xml:"option_id" json:"option_id" yaml:"option_id"`
	Title      string `xml:"title" json:"title" yaml:"title"`
	Type       string `xml:"type" json:"type" yaml:"type"`
	Sort_order string `xml:"sort_order" json:"sort_order" yaml:"sort_order"`
	Is_require int    `xml:"is_require" json:"is_require" yaml:"is_require"`
}

// CatalogProductCustomOptionToAdd was auto-generated from WSDL.
type CatalogProductCustomOptionToAdd struct {
	Title             string                                           `xml:"title" json:"title" yaml:"title"`
	Type              string                                           `xml:"type" json:"type" yaml:"type"`
	Sort_order        *string                                          `xml:"sort_order,omitempty" json:"sort_order,omitempty" yaml:"sort_order,omitempty"`
	Is_require        *int                                             `xml:"is_require,omitempty" json:"is_require,omitempty" yaml:"is_require,omitempty"`
	Additional_fields *CatalogProductCustomOptionAdditionalFieldsArray `xml:"additional_fields" json:"additional_fields" yaml:"additional_fields"`
}

// CatalogProductCustomOptionToUpdate was auto-generated from WSDL.
type CatalogProductCustomOptionToUpdate struct {
	Title             *string                                          `xml:"title,omitempty" json:"title,omitempty" yaml:"title,omitempty"`
	Type              *string                                          `xml:"type,omitempty" json:"type,omitempty" yaml:"type,omitempty"`
	Sort_order        *string                                          `xml:"sort_order,omitempty" json:"sort_order,omitempty" yaml:"sort_order,omitempty"`
	Is_require        *int                                             `xml:"is_require,omitempty" json:"is_require,omitempty" yaml:"is_require,omitempty"`
	Additional_fields *CatalogProductCustomOptionAdditionalFieldsArray `xml:"additional_fields,omitempty" json:"additional_fields,omitempty" yaml:"additional_fields,omitempty"`
}

// CatalogProductCustomOptionTypesArray was auto-generated from
// WSDL.
type CatalogProductCustomOptionTypesArray struct {
	Items []*CatalogProductCustomOptionTypesEntity `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
}

// CatalogProductCustomOptionTypesEntity was auto-generated from
// WSDL.
type CatalogProductCustomOptionTypesEntity struct {
	Label *string `xml:"label,omitempty" json:"label,omitempty" yaml:"label,omitempty"`
	Value *string `xml:"value,omitempty" json:"value,omitempty" yaml:"value,omitempty"`
}

// CatalogProductCustomOptionValueAddArray was auto-generated from
// WSDL.
type CatalogProductCustomOptionValueAddArray struct {
	Items []*CatalogProductCustomOptionValueAddEntity `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
}

// CatalogProductCustomOptionValueAddEntity was auto-generated
// from WSDL.
type CatalogProductCustomOptionValueAddEntity struct {
	Title      string  `xml:"title" json:"title" yaml:"title"`
	Price      string  `xml:"price" json:"price" yaml:"price"`
	Price_type string  `xml:"price_type" json:"price_type" yaml:"price_type"`
	Sku        string  `xml:"sku" json:"sku" yaml:"sku"`
	Sort_order *string `xml:"sort_order,omitempty" json:"sort_order,omitempty" yaml:"sort_order,omitempty"`
}

// CatalogProductCustomOptionValueInfoEntity was auto-generated
// from WSDL.
type CatalogProductCustomOptionValueInfoEntity struct {
	Value_id           string `xml:"value_id" json:"value_id" yaml:"value_id"`
	Option_id          string `xml:"option_id" json:"option_id" yaml:"option_id"`
	Sku                string `xml:"sku" json:"sku" yaml:"sku"`
	Sort_order         string `xml:"sort_order" json:"sort_order" yaml:"sort_order"`
	Default_price      string `xml:"default_price" json:"default_price" yaml:"default_price"`
	Default_price_type string `xml:"default_price_type" json:"default_price_type" yaml:"default_price_type"`
	Store_price        string `xml:"store_price" json:"store_price" yaml:"store_price"`
	Store_price_type   string `xml:"store_price_type" json:"store_price_type" yaml:"store_price_type"`
	Price              string `xml:"price" json:"price" yaml:"price"`
	Price_type         string `xml:"price_type" json:"price_type" yaml:"price_type"`
	Default_title      string `xml:"default_title" json:"default_title" yaml:"default_title"`
	Store_title        string `xml:"store_title" json:"store_title" yaml:"store_title"`
	Title              string `xml:"title" json:"title" yaml:"title"`
}

// CatalogProductCustomOptionValueListArray was auto-generated
// from WSDL.
type CatalogProductCustomOptionValueListArray struct {
	Items []*CatalogProductCustomOptionValueListEntity `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
}

// CatalogProductCustomOptionValueListEntity was auto-generated
// from WSDL.
type CatalogProductCustomOptionValueListEntity struct {
	Value_id   string `xml:"value_id" json:"value_id" yaml:"value_id"`
	Title      string `xml:"title" json:"title" yaml:"title"`
	Price      string `xml:"price" json:"price" yaml:"price"`
	Price_type string `xml:"price_type" json:"price_type" yaml:"price_type"`
	Sku        string `xml:"sku" json:"sku" yaml:"sku"`
	Sort_order string `xml:"sort_order" json:"sort_order" yaml:"sort_order"`
}

// CatalogProductCustomOptionValueUpdateEntity was auto-generated
// from WSDL.
type CatalogProductCustomOptionValueUpdateEntity struct {
	Title      string  `xml:"title" json:"title" yaml:"title"`
	Price      string  `xml:"price" json:"price" yaml:"price"`
	Price_type string  `xml:"price_type" json:"price_type" yaml:"price_type"`
	Sku        string  `xml:"sku" json:"sku" yaml:"sku"`
	Sort_order *string `xml:"sort_order,omitempty" json:"sort_order,omitempty" yaml:"sort_order,omitempty"`
}

// CatalogProductDownloadableLinkAddEntity was auto-generated from
// WSDL.
type CatalogProductDownloadableLinkAddEntity struct {
	Title               string                                         `xml:"title" json:"title" yaml:"title"`
	Price               *string                                        `xml:"price,omitempty" json:"price,omitempty" yaml:"price,omitempty"`
	Is_unlimited        *int                                           `xml:"is_unlimited,omitempty" json:"is_unlimited,omitempty" yaml:"is_unlimited,omitempty"`
	Number_of_downloads *int                                           `xml:"number_of_downloads,omitempty" json:"number_of_downloads,omitempty" yaml:"number_of_downloads,omitempty"`
	Is_shareable        *int                                           `xml:"is_shareable,omitempty" json:"is_shareable,omitempty" yaml:"is_shareable,omitempty"`
	Sample              *CatalogProductDownloadableLinkAddSampleEntity `xml:"sample,omitempty" json:"sample,omitempty" yaml:"sample,omitempty"`
	Type                *string                                        `xml:"type,omitempty" json:"type,omitempty" yaml:"type,omitempty"`
	File                *CatalogProductDownloadableLinkFileEntity      `xml:"file,omitempty" json:"file,omitempty" yaml:"file,omitempty"`
	Link_url            *string                                        `xml:"link_url,omitempty" json:"link_url,omitempty" yaml:"link_url,omitempty"`
	Sample_url          *string                                        `xml:"sample_url,omitempty" json:"sample_url,omitempty" yaml:"sample_url,omitempty"`
	Sort_order          *int                                           `xml:"sort_order,omitempty" json:"sort_order,omitempty" yaml:"sort_order,omitempty"`
}

// CatalogProductDownloadableLinkAddSampleEntity was auto-generated
// from WSDL.
type CatalogProductDownloadableLinkAddSampleEntity struct {
	Type *string                                   `xml:"type,omitempty" json:"type,omitempty" yaml:"type,omitempty"`
	File *CatalogProductDownloadableLinkFileEntity `xml:"file,omitempty" json:"file,omitempty" yaml:"file,omitempty"`
	Url  *string                                   `xml:"url,omitempty" json:"url,omitempty" yaml:"url,omitempty"`
}

// CatalogProductDownloadableLinkEntity was auto-generated from
// WSDL.
type CatalogProductDownloadableLinkEntity struct {
	Link_id             *string                                            `xml:"link_id,omitempty" json:"link_id,omitempty" yaml:"link_id,omitempty"`
	Title               *string                                            `xml:"title,omitempty" json:"title,omitempty" yaml:"title,omitempty"`
	Price               *string                                            `xml:"price,omitempty" json:"price,omitempty" yaml:"price,omitempty"`
	Number_of_downloads *int                                               `xml:"number_of_downloads,omitempty" json:"number_of_downloads,omitempty" yaml:"number_of_downloads,omitempty"`
	Is_unlimited        *int                                               `xml:"is_unlimited,omitempty" json:"is_unlimited,omitempty" yaml:"is_unlimited,omitempty"`
	Is_shareable        *int                                               `xml:"is_shareable,omitempty" json:"is_shareable,omitempty" yaml:"is_shareable,omitempty"`
	Link_url            *string                                            `xml:"link_url,omitempty" json:"link_url,omitempty" yaml:"link_url,omitempty"`
	Link_type           *string                                            `xml:"link_type,omitempty" json:"link_type,omitempty" yaml:"link_type,omitempty"`
	Sample_file         *string                                            `xml:"sample_file,omitempty" json:"sample_file,omitempty" yaml:"sample_file,omitempty"`
	Sample_url          *string                                            `xml:"sample_url,omitempty" json:"sample_url,omitempty" yaml:"sample_url,omitempty"`
	Sample_type         *string                                            `xml:"sample_type,omitempty" json:"sample_type,omitempty" yaml:"sample_type,omitempty"`
	Sort_order          *int                                               `xml:"sort_order,omitempty" json:"sort_order,omitempty" yaml:"sort_order,omitempty"`
	File_save           *CatalogProductDownloadableLinkFileInfoEntityArray `xml:"file_save,omitempty" json:"file_save,omitempty" yaml:"file_save,omitempty"`
	Sample_file_save    *CatalogProductDownloadableLinkFileInfoEntityArray `xml:"sample_file_save,omitempty" json:"sample_file_save,omitempty" yaml:"sample_file_save,omitempty"`
}

// CatalogProductDownloadableLinkEntityArray was auto-generated
// from WSDL.
type CatalogProductDownloadableLinkEntityArray struct {
	Items []*CatalogProductDownloadableLinkEntity `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
}

// CatalogProductDownloadableLinkFileEntity was auto-generated
// from WSDL.
type CatalogProductDownloadableLinkFileEntity struct {
	Name           *string `xml:"name,omitempty" json:"name,omitempty" yaml:"name,omitempty"`
	Base64_content *string `xml:"base64_content,omitempty" json:"base64_content,omitempty" yaml:"base64_content,omitempty"`
}

// CatalogProductDownloadableLinkFileInfoEntity was auto-generated
// from WSDL.
type CatalogProductDownloadableLinkFileInfoEntity struct {
	File   *string `xml:"file,omitempty" json:"file,omitempty" yaml:"file,omitempty"`
	Name   *string `xml:"name,omitempty" json:"name,omitempty" yaml:"name,omitempty"`
	Size   *int    `xml:"size,omitempty" json:"size,omitempty" yaml:"size,omitempty"`
	Status *string `xml:"status,omitempty" json:"status,omitempty" yaml:"status,omitempty"`
}

// CatalogProductDownloadableLinkFileInfoEntityArray was auto-generated
// from WSDL.
type CatalogProductDownloadableLinkFileInfoEntityArray struct {
	Items []*CatalogProductDownloadableLinkFileInfoEntity `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
}

// CatalogProductDownloadableLinkInfoEntity was auto-generated
// from WSDL.
type CatalogProductDownloadableLinkInfoEntity struct {
	Links   *CatalogProductDownloadableLinkEntityArray       `xml:"links,omitempty" json:"links,omitempty" yaml:"links,omitempty"`
	Samples *CatalogProductDownloadableLinkSampleEntityArray `xml:"samples,omitempty" json:"samples,omitempty" yaml:"samples,omitempty"`
}

// CatalogProductDownloadableLinkSampleEntity was auto-generated
// from WSDL.
type CatalogProductDownloadableLinkSampleEntity struct {
	Sample_id     *string `xml:"sample_id,omitempty" json:"sample_id,omitempty" yaml:"sample_id,omitempty"`
	Product_id    *string `xml:"product_id,omitempty" json:"product_id,omitempty" yaml:"product_id,omitempty"`
	Sample_file   *string `xml:"sample_file,omitempty" json:"sample_file,omitempty" yaml:"sample_file,omitempty"`
	Sample_url    *string `xml:"sample_url,omitempty" json:"sample_url,omitempty" yaml:"sample_url,omitempty"`
	Sample_type   *string `xml:"sample_type,omitempty" json:"sample_type,omitempty" yaml:"sample_type,omitempty"`
	Sort_order    *string `xml:"sort_order,omitempty" json:"sort_order,omitempty" yaml:"sort_order,omitempty"`
	Default_title *string `xml:"default_title,omitempty" json:"default_title,omitempty" yaml:"default_title,omitempty"`
	Store_title   *string `xml:"store_title,omitempty" json:"store_title,omitempty" yaml:"store_title,omitempty"`
	Title         *string `xml:"title,omitempty" json:"title,omitempty" yaml:"title,omitempty"`
}

// CatalogProductDownloadableLinkSampleEntityArray was auto-generated
// from WSDL.
type CatalogProductDownloadableLinkSampleEntityArray struct {
	Items []*CatalogProductDownloadableLinkSampleEntity `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
}

// CatalogProductEntity was auto-generated from WSDL.
type CatalogProductEntity struct {
	Product_id   *string        `xml:"product_id,omitempty" json:"product_id,omitempty" yaml:"product_id,omitempty"`
	Sku          *string        `xml:"sku,omitempty" json:"sku,omitempty" yaml:"sku,omitempty"`
	Name         *string        `xml:"name,omitempty" json:"name,omitempty" yaml:"name,omitempty"`
	Set          *string        `xml:"set,omitempty" json:"set,omitempty" yaml:"set,omitempty"`
	Type         *string        `xml:"type,omitempty" json:"type,omitempty" yaml:"type,omitempty"`
	Category_ids *ArrayOfString `xml:"category_ids,omitempty" json:"category_ids,omitempty" yaml:"category_ids,omitempty"`
	Website_ids  *ArrayOfString `xml:"website_ids,omitempty" json:"website_ids,omitempty" yaml:"website_ids,omitempty"`
}

// CatalogProductEntityArray was auto-generated from WSDL.
type CatalogProductEntityArray struct {
	Items []*CatalogProductEntity `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
}

// CatalogProductImageEntity was auto-generated from WSDL.
type CatalogProductImageEntity struct {
	File     *string        `xml:"file,omitempty" json:"file,omitempty" yaml:"file,omitempty"`
	Label    *string        `xml:"label,omitempty" json:"label,omitempty" yaml:"label,omitempty"`
	Position *string        `xml:"position,omitempty" json:"position,omitempty" yaml:"position,omitempty"`
	Exclude  *string        `xml:"exclude,omitempty" json:"exclude,omitempty" yaml:"exclude,omitempty"`
	Url      *string        `xml:"url,omitempty" json:"url,omitempty" yaml:"url,omitempty"`
	Types    *ArrayOfString `xml:"types,omitempty" json:"types,omitempty" yaml:"types,omitempty"`
}

// CatalogProductImageEntityArray was auto-generated from WSDL.
type CatalogProductImageEntityArray struct {
	Items []*CatalogProductImageEntity `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
}

// CatalogProductImageFileEntity was auto-generated from WSDL.
type CatalogProductImageFileEntity struct {
	Content *string `xml:"content,omitempty" json:"content,omitempty" yaml:"content,omitempty"`
	Mime    *string `xml:"mime,omitempty" json:"mime,omitempty" yaml:"mime,omitempty"`
	Name    *string `xml:"name,omitempty" json:"name,omitempty" yaml:"name,omitempty"`
}

// CatalogProductLinkAttributeEntity was auto-generated from WSDL.
type CatalogProductLinkAttributeEntity struct {
	Code *string `xml:"code,omitempty" json:"code,omitempty" yaml:"code,omitempty"`
	Type *string `xml:"type,omitempty" json:"type,omitempty" yaml:"type,omitempty"`
}

// CatalogProductLinkAttributeEntityArray was auto-generated from
// WSDL.
type CatalogProductLinkAttributeEntityArray struct {
	Items []*CatalogProductLinkAttributeEntity `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
}

// CatalogProductLinkEntity was auto-generated from WSDL.
type CatalogProductLinkEntity struct {
	Product_id *string `xml:"product_id,omitempty" json:"product_id,omitempty" yaml:"product_id,omitempty"`
	Type       *string `xml:"type,omitempty" json:"type,omitempty" yaml:"type,omitempty"`
	Set        *string `xml:"set,omitempty" json:"set,omitempty" yaml:"set,omitempty"`
	Sku        *string `xml:"sku,omitempty" json:"sku,omitempty" yaml:"sku,omitempty"`
	Position   *string `xml:"position,omitempty" json:"position,omitempty" yaml:"position,omitempty"`
	Qty        *string `xml:"qty,omitempty" json:"qty,omitempty" yaml:"qty,omitempty"`
}

// CatalogProductLinkEntityArray was auto-generated from WSDL.
type CatalogProductLinkEntityArray struct {
	Items []*CatalogProductLinkEntity `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
}

// CatalogProductRequestAttributes was auto-generated from WSDL.
type CatalogProductRequestAttributes struct {
	Attributes            *ArrayOfString `xml:"attributes,omitempty" json:"attributes,omitempty" yaml:"attributes,omitempty"`
	Additional_attributes *ArrayOfString `xml:"additional_attributes,omitempty" json:"additional_attributes,omitempty" yaml:"additional_attributes,omitempty"`
}

// CatalogProductReturnEntity was auto-generated from WSDL.
type CatalogProductReturnEntity struct {
	Product_id             *string                             `xml:"product_id,omitempty" json:"product_id,omitempty" yaml:"product_id,omitempty"`
	Sku                    *string                             `xml:"sku,omitempty" json:"sku,omitempty" yaml:"sku,omitempty"`
	Set                    *string                             `xml:"set,omitempty" json:"set,omitempty" yaml:"set,omitempty"`
	Type                   *string                             `xml:"type,omitempty" json:"type,omitempty" yaml:"type,omitempty"`
	Categories             *ArrayOfString                      `xml:"categories,omitempty" json:"categories,omitempty" yaml:"categories,omitempty"`
	Websites               *ArrayOfString                      `xml:"websites,omitempty" json:"websites,omitempty" yaml:"websites,omitempty"`
	Created_at             *string                             `xml:"created_at,omitempty" json:"created_at,omitempty" yaml:"created_at,omitempty"`
	Updated_at             *string                             `xml:"updated_at,omitempty" json:"updated_at,omitempty" yaml:"updated_at,omitempty"`
	Type_id                *string                             `xml:"type_id,omitempty" json:"type_id,omitempty" yaml:"type_id,omitempty"`
	Name                   *string                             `xml:"name,omitempty" json:"name,omitempty" yaml:"name,omitempty"`
	Description            *string                             `xml:"description,omitempty" json:"description,omitempty" yaml:"description,omitempty"`
	Short_description      *string                             `xml:"short_description,omitempty" json:"short_description,omitempty" yaml:"short_description,omitempty"`
	Weight                 *string                             `xml:"weight,omitempty" json:"weight,omitempty" yaml:"weight,omitempty"`
	Status                 *string                             `xml:"status,omitempty" json:"status,omitempty" yaml:"status,omitempty"`
	Url_key                *string                             `xml:"url_key,omitempty" json:"url_key,omitempty" yaml:"url_key,omitempty"`
	Url_path               *string                             `xml:"url_path,omitempty" json:"url_path,omitempty" yaml:"url_path,omitempty"`
	Visibility             *string                             `xml:"visibility,omitempty" json:"visibility,omitempty" yaml:"visibility,omitempty"`
	Category_ids           *ArrayOfString                      `xml:"category_ids,omitempty" json:"category_ids,omitempty" yaml:"category_ids,omitempty"`
	Website_ids            *ArrayOfString                      `xml:"website_ids,omitempty" json:"website_ids,omitempty" yaml:"website_ids,omitempty"`
	Has_options            *string                             `xml:"has_options,omitempty" json:"has_options,omitempty" yaml:"has_options,omitempty"`
	Gift_message_available *string                             `xml:"gift_message_available,omitempty" json:"gift_message_available,omitempty" yaml:"gift_message_available,omitempty"`
	Price                  *string                             `xml:"price,omitempty" json:"price,omitempty" yaml:"price,omitempty"`
	Special_price          *string                             `xml:"special_price,omitempty" json:"special_price,omitempty" yaml:"special_price,omitempty"`
	Special_from_date      *string                             `xml:"special_from_date,omitempty" json:"special_from_date,omitempty" yaml:"special_from_date,omitempty"`
	Special_to_date        *string                             `xml:"special_to_date,omitempty" json:"special_to_date,omitempty" yaml:"special_to_date,omitempty"`
	Tax_class_id           *string                             `xml:"tax_class_id,omitempty" json:"tax_class_id,omitempty" yaml:"tax_class_id,omitempty"`
	Tier_price             *CatalogProductTierPriceEntityArray `xml:"tier_price,omitempty" json:"tier_price,omitempty" yaml:"tier_price,omitempty"`
	Meta_title             *string                             `xml:"meta_title,omitempty" json:"meta_title,omitempty" yaml:"meta_title,omitempty"`
	Meta_keyword           *string                             `xml:"meta_keyword,omitempty" json:"meta_keyword,omitempty" yaml:"meta_keyword,omitempty"`
	Meta_description       *string                             `xml:"meta_description,omitempty" json:"meta_description,omitempty" yaml:"meta_description,omitempty"`
	Custom_design          *string                             `xml:"custom_design,omitempty" json:"custom_design,omitempty" yaml:"custom_design,omitempty"`
	Custom_layout_update   *string                             `xml:"custom_layout_update,omitempty" json:"custom_layout_update,omitempty" yaml:"custom_layout_update,omitempty"`
	Options_container      *string                             `xml:"options_container,omitempty" json:"options_container,omitempty" yaml:"options_container,omitempty"`
	Additional_attributes  *AssociativeArray                   `xml:"additional_attributes,omitempty" json:"additional_attributes,omitempty" yaml:"additional_attributes,omitempty"`
}

// CatalogProductSpecialPriceReturnEntity was auto-generated from
// WSDL.
type CatalogProductSpecialPriceReturnEntity struct {
	Special_price     string `xml:"special_price" json:"special_price" yaml:"special_price"`
	Special_from_date string `xml:"special_from_date" json:"special_from_date" yaml:"special_from_date"`
	Special_to_date   string `xml:"special_to_date" json:"special_to_date" yaml:"special_to_date"`
}

// CatalogProductTagAddEntity was auto-generated from WSDL.
type CatalogProductTagAddEntity struct {
	Tag         string `xml:"tag" json:"tag" yaml:"tag"`
	Product_id  string `xml:"product_id" json:"product_id" yaml:"product_id"`
	Customer_id string `xml:"customer_id" json:"customer_id" yaml:"customer_id"`
	Store       string `xml:"store" json:"store" yaml:"store"`
}

// CatalogProductTagInfoEntity was auto-generated from WSDL.
type CatalogProductTagInfoEntity struct {
	Name            string            `xml:"name" json:"name" yaml:"name"`
	Status          string            `xml:"status" json:"status" yaml:"status"`
	Base_popularity string            `xml:"base_popularity" json:"base_popularity" yaml:"base_popularity"`
	Products        *AssociativeArray `xml:"products" json:"products" yaml:"products"`
}

// CatalogProductTagListEntity was auto-generated from WSDL.
type CatalogProductTagListEntity struct {
	Tag_id string `xml:"tag_id" json:"tag_id" yaml:"tag_id"`
	Name   string `xml:"name" json:"name" yaml:"name"`
}

// CatalogProductTagListEntityArray was auto-generated from WSDL.
type CatalogProductTagListEntityArray struct {
	Items []*CatalogProductTagListEntity `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
}

// CatalogProductTagUpdateEntity was auto-generated from WSDL.
type CatalogProductTagUpdateEntity struct {
	Name            *string `xml:"name,omitempty" json:"name,omitempty" yaml:"name,omitempty"`
	Status          *string `xml:"status,omitempty" json:"status,omitempty" yaml:"status,omitempty"`
	Base_popularity *string `xml:"base_popularity,omitempty" json:"base_popularity,omitempty" yaml:"base_popularity,omitempty"`
}

// CatalogProductTierPriceEntity was auto-generated from WSDL.
type CatalogProductTierPriceEntity struct {
	Customer_group_id *string  `xml:"customer_group_id,omitempty" json:"customer_group_id,omitempty" yaml:"customer_group_id,omitempty"`
	Website           *string  `xml:"website,omitempty" json:"website,omitempty" yaml:"website,omitempty"`
	Qty               *int     `xml:"qty,omitempty" json:"qty,omitempty" yaml:"qty,omitempty"`
	Price             *float64 `xml:"price,omitempty" json:"price,omitempty" yaml:"price,omitempty"`
}

// CatalogProductTierPriceEntityArray was auto-generated from WSDL.
type CatalogProductTierPriceEntityArray struct {
	Items []*CatalogProductTierPriceEntity `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
}

// CatalogProductTypeEntity was auto-generated from WSDL.
type CatalogProductTypeEntity struct {
	Type  *string `xml:"type,omitempty" json:"type,omitempty" yaml:"type,omitempty"`
	Label *string `xml:"label,omitempty" json:"label,omitempty" yaml:"label,omitempty"`
}

// CatalogProductTypeEntityArray was auto-generated from WSDL.
type CatalogProductTypeEntityArray struct {
	Items []*CatalogProductTypeEntity `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
}

// ComplexFilter was auto-generated from WSDL.
type ComplexFilter struct {
	Key   *string            `xml:"key,omitempty" json:"key,omitempty" yaml:"key,omitempty"`
	Value *AssociativeEntity `xml:"value,omitempty" json:"value,omitempty" yaml:"value,omitempty"`
}

// ComplexFilterArray was auto-generated from WSDL.
type ComplexFilterArray struct {
	Items []*ComplexFilter `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
}

// CustomerAddressEntityArray was auto-generated from WSDL.
type CustomerAddressEntityArray struct {
	Items []*CustomerAddressEntityItem `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
}

// CustomerAddressEntityCreate was auto-generated from WSDL.
type CustomerAddressEntityCreate struct {
	City                *string        `xml:"city,omitempty" json:"city,omitempty" yaml:"city,omitempty"`
	Company             *string        `xml:"company,omitempty" json:"company,omitempty" yaml:"company,omitempty"`
	Country_id          *string        `xml:"country_id,omitempty" json:"country_id,omitempty" yaml:"country_id,omitempty"`
	Fax                 *string        `xml:"fax,omitempty" json:"fax,omitempty" yaml:"fax,omitempty"`
	Firstname           *string        `xml:"firstname,omitempty" json:"firstname,omitempty" yaml:"firstname,omitempty"`
	Lastname            *string        `xml:"lastname,omitempty" json:"lastname,omitempty" yaml:"lastname,omitempty"`
	Middlename          *string        `xml:"middlename,omitempty" json:"middlename,omitempty" yaml:"middlename,omitempty"`
	Postcode            *string        `xml:"postcode,omitempty" json:"postcode,omitempty" yaml:"postcode,omitempty"`
	Prefix              *string        `xml:"prefix,omitempty" json:"prefix,omitempty" yaml:"prefix,omitempty"`
	Region_id           *int           `xml:"region_id,omitempty" json:"region_id,omitempty" yaml:"region_id,omitempty"`
	Region              *string        `xml:"region,omitempty" json:"region,omitempty" yaml:"region,omitempty"`
	Street              *ArrayOfString `xml:"street,omitempty" json:"street,omitempty" yaml:"street,omitempty"`
	Suffix              *string        `xml:"suffix,omitempty" json:"suffix,omitempty" yaml:"suffix,omitempty"`
	Telephone           *string        `xml:"telephone,omitempty" json:"telephone,omitempty" yaml:"telephone,omitempty"`
	Is_default_billing  *bool          `xml:"is_default_billing,omitempty" json:"is_default_billing,omitempty" yaml:"is_default_billing,omitempty"`
	Is_default_shipping *bool          `xml:"is_default_shipping,omitempty" json:"is_default_shipping,omitempty" yaml:"is_default_shipping,omitempty"`
}

// CustomerAddressEntityItem was auto-generated from WSDL.
type CustomerAddressEntityItem struct {
	Customer_address_id *int    `xml:"customer_address_id,omitempty" json:"customer_address_id,omitempty" yaml:"customer_address_id,omitempty"`
	Created_at          *string `xml:"created_at,omitempty" json:"created_at,omitempty" yaml:"created_at,omitempty"`
	Updated_at          *string `xml:"updated_at,omitempty" json:"updated_at,omitempty" yaml:"updated_at,omitempty"`
	Increment_id        *string `xml:"increment_id,omitempty" json:"increment_id,omitempty" yaml:"increment_id,omitempty"`
	City                *string `xml:"city,omitempty" json:"city,omitempty" yaml:"city,omitempty"`
	Company             *string `xml:"company,omitempty" json:"company,omitempty" yaml:"company,omitempty"`
	Country_id          *string `xml:"country_id,omitempty" json:"country_id,omitempty" yaml:"country_id,omitempty"`
	Fax                 *string `xml:"fax,omitempty" json:"fax,omitempty" yaml:"fax,omitempty"`
	Firstname           *string `xml:"firstname,omitempty" json:"firstname,omitempty" yaml:"firstname,omitempty"`
	Lastname            *string `xml:"lastname,omitempty" json:"lastname,omitempty" yaml:"lastname,omitempty"`
	Middlename          *string `xml:"middlename,omitempty" json:"middlename,omitempty" yaml:"middlename,omitempty"`
	Postcode            *string `xml:"postcode,omitempty" json:"postcode,omitempty" yaml:"postcode,omitempty"`
	Prefix              *string `xml:"prefix,omitempty" json:"prefix,omitempty" yaml:"prefix,omitempty"`
	Region              *string `xml:"region,omitempty" json:"region,omitempty" yaml:"region,omitempty"`
	Region_id           *int    `xml:"region_id,omitempty" json:"region_id,omitempty" yaml:"region_id,omitempty"`
	Street              *string `xml:"street,omitempty" json:"street,omitempty" yaml:"street,omitempty"`
	Suffix              *string `xml:"suffix,omitempty" json:"suffix,omitempty" yaml:"suffix,omitempty"`
	Telephone           *string `xml:"telephone,omitempty" json:"telephone,omitempty" yaml:"telephone,omitempty"`
	Is_default_billing  *bool   `xml:"is_default_billing,omitempty" json:"is_default_billing,omitempty" yaml:"is_default_billing,omitempty"`
	Is_default_shipping *bool   `xml:"is_default_shipping,omitempty" json:"is_default_shipping,omitempty" yaml:"is_default_shipping,omitempty"`
}

// CustomerCustomerEntity was auto-generated from WSDL.
type CustomerCustomerEntity struct {
	Customer_id   *int    `xml:"customer_id,omitempty" json:"customer_id,omitempty" yaml:"customer_id,omitempty"`
	Created_at    *string `xml:"created_at,omitempty" json:"created_at,omitempty" yaml:"created_at,omitempty"`
	Updated_at    *string `xml:"updated_at,omitempty" json:"updated_at,omitempty" yaml:"updated_at,omitempty"`
	Increment_id  *string `xml:"increment_id,omitempty" json:"increment_id,omitempty" yaml:"increment_id,omitempty"`
	Store_id      *int    `xml:"store_id,omitempty" json:"store_id,omitempty" yaml:"store_id,omitempty"`
	Website_id    *int    `xml:"website_id,omitempty" json:"website_id,omitempty" yaml:"website_id,omitempty"`
	Created_in    *string `xml:"created_in,omitempty" json:"created_in,omitempty" yaml:"created_in,omitempty"`
	Email         *string `xml:"email,omitempty" json:"email,omitempty" yaml:"email,omitempty"`
	Firstname     *string `xml:"firstname,omitempty" json:"firstname,omitempty" yaml:"firstname,omitempty"`
	Middlename    *string `xml:"middlename,omitempty" json:"middlename,omitempty" yaml:"middlename,omitempty"`
	Lastname      *string `xml:"lastname,omitempty" json:"lastname,omitempty" yaml:"lastname,omitempty"`
	Group_id      *int    `xml:"group_id,omitempty" json:"group_id,omitempty" yaml:"group_id,omitempty"`
	Prefix        *string `xml:"prefix,omitempty" json:"prefix,omitempty" yaml:"prefix,omitempty"`
	Suffix        *string `xml:"suffix,omitempty" json:"suffix,omitempty" yaml:"suffix,omitempty"`
	Dob           *string `xml:"dob,omitempty" json:"dob,omitempty" yaml:"dob,omitempty"`
	Taxvat        *string `xml:"taxvat,omitempty" json:"taxvat,omitempty" yaml:"taxvat,omitempty"`
	Confirmation  *bool   `xml:"confirmation,omitempty" json:"confirmation,omitempty" yaml:"confirmation,omitempty"`
	Password_hash *string `xml:"password_hash,omitempty" json:"password_hash,omitempty" yaml:"password_hash,omitempty"`
}

// CustomerCustomerEntityArray was auto-generated from WSDL.
type CustomerCustomerEntityArray struct {
	Items []*CustomerCustomerEntity `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
}

// CustomerCustomerEntityToCreate was auto-generated from WSDL.
type CustomerCustomerEntityToCreate struct {
	Customer_id *int    `xml:"customer_id,omitempty" json:"customer_id,omitempty" yaml:"customer_id,omitempty"`
	Email       *string `xml:"email,omitempty" json:"email,omitempty" yaml:"email,omitempty"`
	Firstname   *string `xml:"firstname,omitempty" json:"firstname,omitempty" yaml:"firstname,omitempty"`
	Lastname    *string `xml:"lastname,omitempty" json:"lastname,omitempty" yaml:"lastname,omitempty"`
	Middlename  *string `xml:"middlename,omitempty" json:"middlename,omitempty" yaml:"middlename,omitempty"`
	Password    *string `xml:"password,omitempty" json:"password,omitempty" yaml:"password,omitempty"`
	Website_id  *int    `xml:"website_id,omitempty" json:"website_id,omitempty" yaml:"website_id,omitempty"`
	Store_id    *int    `xml:"store_id,omitempty" json:"store_id,omitempty" yaml:"store_id,omitempty"`
	Group_id    *int    `xml:"group_id,omitempty" json:"group_id,omitempty" yaml:"group_id,omitempty"`
	Prefix      *string `xml:"prefix,omitempty" json:"prefix,omitempty" yaml:"prefix,omitempty"`
	Suffix      *string `xml:"suffix,omitempty" json:"suffix,omitempty" yaml:"suffix,omitempty"`
	Dob         *string `xml:"dob,omitempty" json:"dob,omitempty" yaml:"dob,omitempty"`
	Taxvat      *string `xml:"taxvat,omitempty" json:"taxvat,omitempty" yaml:"taxvat,omitempty"`
	Gender      *int    `xml:"gender,omitempty" json:"gender,omitempty" yaml:"gender,omitempty"`
}

// CustomerGroupEntity was auto-generated from WSDL.
type CustomerGroupEntity struct {
	Customer_group_id   *int    `xml:"customer_group_id,omitempty" json:"customer_group_id,omitempty" yaml:"customer_group_id,omitempty"`
	Customer_group_code *string `xml:"customer_group_code,omitempty" json:"customer_group_code,omitempty" yaml:"customer_group_code,omitempty"`
}

// CustomerGroupEntityArray was auto-generated from WSDL.
type CustomerGroupEntityArray struct {
	Items []*CustomerGroupEntity `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
}

// Date was auto-generated from WSDL.
type Date struct {
}

// DateTime was auto-generated from WSDL.
type DateTime struct {
}

// Decimal was auto-generated from WSDL.
type Decimal struct {
}

// DirectoryCountryEntity was auto-generated from WSDL.
type DirectoryCountryEntity struct {
	Country_id *string `xml:"country_id,omitempty" json:"country_id,omitempty" yaml:"country_id,omitempty"`
	Iso2_code  *string `xml:"iso2_code,omitempty" json:"iso2_code,omitempty" yaml:"iso2_code,omitempty"`
	Iso3_code  *string `xml:"iso3_code,omitempty" json:"iso3_code,omitempty" yaml:"iso3_code,omitempty"`
	Name       *string `xml:"name,omitempty" json:"name,omitempty" yaml:"name,omitempty"`
}

// DirectoryCountryEntityArray was auto-generated from WSDL.
type DirectoryCountryEntityArray struct {
	Items []*DirectoryCountryEntity `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
}

// DirectoryRegionEntity was auto-generated from WSDL.
type DirectoryRegionEntity struct {
	Region_id *string `xml:"region_id,omitempty" json:"region_id,omitempty" yaml:"region_id,omitempty"`
	Code      *string `xml:"code,omitempty" json:"code,omitempty" yaml:"code,omitempty"`
	Name      *string `xml:"name,omitempty" json:"name,omitempty" yaml:"name,omitempty"`
}

// DirectoryRegionEntityArray was auto-generated from WSDL.
type DirectoryRegionEntityArray struct {
	Items []*DirectoryRegionEntity `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
}

// Double was auto-generated from WSDL.
type Double struct {
}

// Duration was auto-generated from WSDL.
type Duration struct {
}

// ExistsFaltureEntity was auto-generated from WSDL.
type ExistsFaltureEntity struct {
	Code    *string `xml:"code,omitempty" json:"code,omitempty" yaml:"code,omitempty"`
	Message *string `xml:"message,omitempty" json:"message,omitempty" yaml:"message,omitempty"`
}

// Filters was auto-generated from WSDL.
type Filters struct {
	Filter         *AssociativeArray   `xml:"filter,omitempty" json:"filter,omitempty" yaml:"filter,omitempty"`
	Complex_filter *ComplexFilterArray `xml:"complex_filter,omitempty" json:"complex_filter,omitempty" yaml:"complex_filter,omitempty"`
}

// Float was auto-generated from WSDL.
type Float struct {
}

// GDay was auto-generated from WSDL.
type GDay struct {
}

// GMonth was auto-generated from WSDL.
type GMonth struct {
}

// GMonthDay was auto-generated from WSDL.
type GMonthDay struct {
}

// GYear was auto-generated from WSDL.
type GYear struct {
}

// GYearMonth was auto-generated from WSDL.
type GYearMonth struct {
}

// GiftMessageAssociativeProductsEntity was auto-generated from
// WSDL.
type GiftMessageAssociativeProductsEntity struct {
	Product *ShoppingCartProductEntity `xml:"product,omitempty" json:"product,omitempty" yaml:"product,omitempty"`
	Message *GiftMessageEntity         `xml:"message,omitempty" json:"message,omitempty" yaml:"message,omitempty"`
}

// GiftMessageAssociativeProductsEntityArray was auto-generated
// from WSDL.
type GiftMessageAssociativeProductsEntityArray struct {
	Items []*GiftMessageAssociativeProductsEntity `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
}

// GiftMessageEntity was auto-generated from WSDL.
type GiftMessageEntity struct {
	From    *string `xml:"from,omitempty" json:"from,omitempty" yaml:"from,omitempty"`
	To      *string `xml:"to,omitempty" json:"to,omitempty" yaml:"to,omitempty"`
	Message *string `xml:"message,omitempty" json:"message,omitempty" yaml:"message,omitempty"`
}

// GiftMessageResponse was auto-generated from WSDL.
type GiftMessageResponse struct {
	EntityId *string `xml:"entityId,omitempty" json:"entityId,omitempty" yaml:"entityId,omitempty"`
	Result   *bool   `xml:"result,omitempty" json:"result,omitempty" yaml:"result,omitempty"`
	Error    *string `xml:"error,omitempty" json:"error,omitempty" yaml:"error,omitempty"`
}

// GiftMessageResponseArray was auto-generated from WSDL.
type GiftMessageResponseArray struct {
	Items []*GiftMessageResponse `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
}

// HexBinary was auto-generated from WSDL.
type HexBinary struct {
}

// Int was auto-generated from WSDL.
type Int struct {
}

// Integer was auto-generated from WSDL.
type Integer struct {
}

// Language was auto-generated from WSDL.
type Language struct {
}

// Long was auto-generated from WSDL.
type Long struct {
}

// MagentoInfoEntity was auto-generated from WSDL.
type MagentoInfoEntity struct {
	Magento_version *string `xml:"magento_version,omitempty" json:"magento_version,omitempty" yaml:"magento_version,omitempty"`
	Magento_edition *string `xml:"magento_edition,omitempty" json:"magento_edition,omitempty" yaml:"magento_edition,omitempty"`
}

// NegativeInteger was auto-generated from WSDL.
type NegativeInteger struct {
}

// NonNegativeInteger was auto-generated from WSDL.
type NonNegativeInteger struct {
}

// NonPositiveInteger was auto-generated from WSDL.
type NonPositiveInteger struct {
}

// NormalizedString was auto-generated from WSDL.
type NormalizedString struct {
}

// OrderItemIdQty was auto-generated from WSDL.
type OrderItemIdQty struct {
	Order_item_id *int     `xml:"order_item_id,omitempty" json:"order_item_id,omitempty" yaml:"order_item_id,omitempty"`
	Qty           *float64 `xml:"qty,omitempty" json:"qty,omitempty" yaml:"qty,omitempty"`
}

// OrderItemIdQtyArray was auto-generated from WSDL.
type OrderItemIdQtyArray struct {
	Items []*OrderItemIdQty `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
}

// PositiveInteger was auto-generated from WSDL.
type PositiveInteger struct {
}

// SalesOrderAddressEntity was auto-generated from WSDL.
type SalesOrderAddressEntity struct {
	Increment_id *string `xml:"increment_id,omitempty" json:"increment_id,omitempty" yaml:"increment_id,omitempty"`
	Parent_id    *string `xml:"parent_id,omitempty" json:"parent_id,omitempty" yaml:"parent_id,omitempty"`
	Created_at   *string `xml:"created_at,omitempty" json:"created_at,omitempty" yaml:"created_at,omitempty"`
	Updated_at   *string `xml:"updated_at,omitempty" json:"updated_at,omitempty" yaml:"updated_at,omitempty"`
	Is_active    *string `xml:"is_active,omitempty" json:"is_active,omitempty" yaml:"is_active,omitempty"`
	Address_type *string `xml:"address_type,omitempty" json:"address_type,omitempty" yaml:"address_type,omitempty"`
	Firstname    *string `xml:"firstname,omitempty" json:"firstname,omitempty" yaml:"firstname,omitempty"`
	Lastname     *string `xml:"lastname,omitempty" json:"lastname,omitempty" yaml:"lastname,omitempty"`
	Company      *string `xml:"company,omitempty" json:"company,omitempty" yaml:"company,omitempty"`
	Street       *string `xml:"street,omitempty" json:"street,omitempty" yaml:"street,omitempty"`
	City         *string `xml:"city,omitempty" json:"city,omitempty" yaml:"city,omitempty"`
	Region       *string `xml:"region,omitempty" json:"region,omitempty" yaml:"region,omitempty"`
	Postcode     *string `xml:"postcode,omitempty" json:"postcode,omitempty" yaml:"postcode,omitempty"`
	Country_id   *string `xml:"country_id,omitempty" json:"country_id,omitempty" yaml:"country_id,omitempty"`
	Telephone    *string `xml:"telephone,omitempty" json:"telephone,omitempty" yaml:"telephone,omitempty"`
	Fax          *string `xml:"fax,omitempty" json:"fax,omitempty" yaml:"fax,omitempty"`
	Region_id    *string `xml:"region_id,omitempty" json:"region_id,omitempty" yaml:"region_id,omitempty"`
	Address_id   *string `xml:"address_id,omitempty" json:"address_id,omitempty" yaml:"address_id,omitempty"`
}

// SalesOrderCreditmemoCommentEntity was auto-generated from WSDL.
type SalesOrderCreditmemoCommentEntity struct {
	Parent_id            *string `xml:"parent_id,omitempty" json:"parent_id,omitempty" yaml:"parent_id,omitempty"`
	Created_at           *string `xml:"created_at,omitempty" json:"created_at,omitempty" yaml:"created_at,omitempty"`
	Comment              *string `xml:"comment,omitempty" json:"comment,omitempty" yaml:"comment,omitempty"`
	Is_customer_notified *string `xml:"is_customer_notified,omitempty" json:"is_customer_notified,omitempty" yaml:"is_customer_notified,omitempty"`
	Comment_id           *string `xml:"comment_id,omitempty" json:"comment_id,omitempty" yaml:"comment_id,omitempty"`
	Is_visible_on_front  *string `xml:"is_visible_on_front,omitempty" json:"is_visible_on_front,omitempty" yaml:"is_visible_on_front,omitempty"`
}

// SalesOrderCreditmemoCommentEntityArray was auto-generated from
// WSDL.
type SalesOrderCreditmemoCommentEntityArray struct {
	Items []*SalesOrderCreditmemoCommentEntity `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
}

// SalesOrderCreditmemoData was auto-generated from WSDL.
type SalesOrderCreditmemoData struct {
	Qtys                *OrderItemIdQtyArray `xml:"qtys,omitempty" json:"qtys,omitempty" yaml:"qtys,omitempty"`
	Shipping_amount     *float64             `xml:"shipping_amount,omitempty" json:"shipping_amount,omitempty" yaml:"shipping_amount,omitempty"`
	Adjustment_positive *float64             `xml:"adjustment_positive,omitempty" json:"adjustment_positive,omitempty" yaml:"adjustment_positive,omitempty"`
	Adjustment_negative *float64             `xml:"adjustment_negative,omitempty" json:"adjustment_negative,omitempty" yaml:"adjustment_negative,omitempty"`
}

// SalesOrderCreditmemoEntity was auto-generated from WSDL.
type SalesOrderCreditmemoEntity struct {
	Updated_at                     *string                                 `xml:"updated_at,omitempty" json:"updated_at,omitempty" yaml:"updated_at,omitempty"`
	Created_at                     *string                                 `xml:"created_at,omitempty" json:"created_at,omitempty" yaml:"created_at,omitempty"`
	Increment_id                   *string                                 `xml:"increment_id,omitempty" json:"increment_id,omitempty" yaml:"increment_id,omitempty"`
	Transaction_id                 *string                                 `xml:"transaction_id,omitempty" json:"transaction_id,omitempty" yaml:"transaction_id,omitempty"`
	Global_currency_code           *string                                 `xml:"global_currency_code,omitempty" json:"global_currency_code,omitempty" yaml:"global_currency_code,omitempty"`
	Base_currency_code             *string                                 `xml:"base_currency_code,omitempty" json:"base_currency_code,omitempty" yaml:"base_currency_code,omitempty"`
	Order_currency_code            *string                                 `xml:"order_currency_code,omitempty" json:"order_currency_code,omitempty" yaml:"order_currency_code,omitempty"`
	Store_currency_code            *string                                 `xml:"store_currency_code,omitempty" json:"store_currency_code,omitempty" yaml:"store_currency_code,omitempty"`
	Cybersource_token              *string                                 `xml:"cybersource_token,omitempty" json:"cybersource_token,omitempty" yaml:"cybersource_token,omitempty"`
	Invoice_id                     *string                                 `xml:"invoice_id,omitempty" json:"invoice_id,omitempty" yaml:"invoice_id,omitempty"`
	Billing_address_id             *string                                 `xml:"billing_address_id,omitempty" json:"billing_address_id,omitempty" yaml:"billing_address_id,omitempty"`
	Shipping_address_id            *string                                 `xml:"shipping_address_id,omitempty" json:"shipping_address_id,omitempty" yaml:"shipping_address_id,omitempty"`
	State                          *string                                 `xml:"state,omitempty" json:"state,omitempty" yaml:"state,omitempty"`
	Creditmemo_status              *string                                 `xml:"creditmemo_status,omitempty" json:"creditmemo_status,omitempty" yaml:"creditmemo_status,omitempty"`
	Email_sent                     *string                                 `xml:"email_sent,omitempty" json:"email_sent,omitempty" yaml:"email_sent,omitempty"`
	Order_id                       *string                                 `xml:"order_id,omitempty" json:"order_id,omitempty" yaml:"order_id,omitempty"`
	Tax_amount                     *string                                 `xml:"tax_amount,omitempty" json:"tax_amount,omitempty" yaml:"tax_amount,omitempty"`
	Shipping_tax_amount            *string                                 `xml:"shipping_tax_amount,omitempty" json:"shipping_tax_amount,omitempty" yaml:"shipping_tax_amount,omitempty"`
	Base_tax_amount                *string                                 `xml:"base_tax_amount,omitempty" json:"base_tax_amount,omitempty" yaml:"base_tax_amount,omitempty"`
	Base_adjustment_positive       *string                                 `xml:"base_adjustment_positive,omitempty" json:"base_adjustment_positive,omitempty" yaml:"base_adjustment_positive,omitempty"`
	Base_grand_total               *string                                 `xml:"base_grand_total,omitempty" json:"base_grand_total,omitempty" yaml:"base_grand_total,omitempty"`
	Adjustment                     *string                                 `xml:"adjustment,omitempty" json:"adjustment,omitempty" yaml:"adjustment,omitempty"`
	Subtotal                       *string                                 `xml:"subtotal,omitempty" json:"subtotal,omitempty" yaml:"subtotal,omitempty"`
	Discount_amount                *string                                 `xml:"discount_amount,omitempty" json:"discount_amount,omitempty" yaml:"discount_amount,omitempty"`
	Base_subtotal                  *string                                 `xml:"base_subtotal,omitempty" json:"base_subtotal,omitempty" yaml:"base_subtotal,omitempty"`
	Base_adjustment                *string                                 `xml:"base_adjustment,omitempty" json:"base_adjustment,omitempty" yaml:"base_adjustment,omitempty"`
	Base_to_global_rate            *string                                 `xml:"base_to_global_rate,omitempty" json:"base_to_global_rate,omitempty" yaml:"base_to_global_rate,omitempty"`
	Store_to_base_rate             *string                                 `xml:"store_to_base_rate,omitempty" json:"store_to_base_rate,omitempty" yaml:"store_to_base_rate,omitempty"`
	Base_shipping_amount           *string                                 `xml:"base_shipping_amount,omitempty" json:"base_shipping_amount,omitempty" yaml:"base_shipping_amount,omitempty"`
	Adjustment_negative            *string                                 `xml:"adjustment_negative,omitempty" json:"adjustment_negative,omitempty" yaml:"adjustment_negative,omitempty"`
	Subtotal_incl_tax              *string                                 `xml:"subtotal_incl_tax,omitempty" json:"subtotal_incl_tax,omitempty" yaml:"subtotal_incl_tax,omitempty"`
	Shipping_amount                *string                                 `xml:"shipping_amount,omitempty" json:"shipping_amount,omitempty" yaml:"shipping_amount,omitempty"`
	Base_subtotal_incl_tax         *string                                 `xml:"base_subtotal_incl_tax,omitempty" json:"base_subtotal_incl_tax,omitempty" yaml:"base_subtotal_incl_tax,omitempty"`
	Base_adjustment_negative       *string                                 `xml:"base_adjustment_negative,omitempty" json:"base_adjustment_negative,omitempty" yaml:"base_adjustment_negative,omitempty"`
	Grand_total                    *string                                 `xml:"grand_total,omitempty" json:"grand_total,omitempty" yaml:"grand_total,omitempty"`
	Base_discount_amount           *string                                 `xml:"base_discount_amount,omitempty" json:"base_discount_amount,omitempty" yaml:"base_discount_amount,omitempty"`
	Base_to_order_rate             *string                                 `xml:"base_to_order_rate,omitempty" json:"base_to_order_rate,omitempty" yaml:"base_to_order_rate,omitempty"`
	Store_to_order_rate            *string                                 `xml:"store_to_order_rate,omitempty" json:"store_to_order_rate,omitempty" yaml:"store_to_order_rate,omitempty"`
	Base_shipping_tax_amount       *string                                 `xml:"base_shipping_tax_amount,omitempty" json:"base_shipping_tax_amount,omitempty" yaml:"base_shipping_tax_amount,omitempty"`
	Adjustment_positive            *string                                 `xml:"adjustment_positive,omitempty" json:"adjustment_positive,omitempty" yaml:"adjustment_positive,omitempty"`
	Store_id                       *string                                 `xml:"store_id,omitempty" json:"store_id,omitempty" yaml:"store_id,omitempty"`
	Hidden_tax_amount              *string                                 `xml:"hidden_tax_amount,omitempty" json:"hidden_tax_amount,omitempty" yaml:"hidden_tax_amount,omitempty"`
	Base_hidden_tax_amount         *string                                 `xml:"base_hidden_tax_amount,omitempty" json:"base_hidden_tax_amount,omitempty" yaml:"base_hidden_tax_amount,omitempty"`
	Shipping_hidden_tax_amount     *string                                 `xml:"shipping_hidden_tax_amount,omitempty" json:"shipping_hidden_tax_amount,omitempty" yaml:"shipping_hidden_tax_amount,omitempty"`
	Base_shipping_hidden_tax_amnt  *string                                 `xml:"base_shipping_hidden_tax_amnt,omitempty" json:"base_shipping_hidden_tax_amnt,omitempty" yaml:"base_shipping_hidden_tax_amnt,omitempty"`
	Shipping_incl_tax              *string                                 `xml:"shipping_incl_tax,omitempty" json:"shipping_incl_tax,omitempty" yaml:"shipping_incl_tax,omitempty"`
	Base_shipping_incl_tax         *string                                 `xml:"base_shipping_incl_tax,omitempty" json:"base_shipping_incl_tax,omitempty" yaml:"base_shipping_incl_tax,omitempty"`
	Base_customer_balance_amount   *string                                 `xml:"base_customer_balance_amount,omitempty" json:"base_customer_balance_amount,omitempty" yaml:"base_customer_balance_amount,omitempty"`
	Customer_balance_amount        *string                                 `xml:"customer_balance_amount,omitempty" json:"customer_balance_amount,omitempty" yaml:"customer_balance_amount,omitempty"`
	Bs_customer_bal_total_refunded *string                                 `xml:"bs_customer_bal_total_refunded,omitempty" json:"bs_customer_bal_total_refunded,omitempty" yaml:"bs_customer_bal_total_refunded,omitempty"`
	Customer_bal_total_refunded    *string                                 `xml:"customer_bal_total_refunded,omitempty" json:"customer_bal_total_refunded,omitempty" yaml:"customer_bal_total_refunded,omitempty"`
	Base_gift_cards_amount         *string                                 `xml:"base_gift_cards_amount,omitempty" json:"base_gift_cards_amount,omitempty" yaml:"base_gift_cards_amount,omitempty"`
	Gift_cards_amount              *string                                 `xml:"gift_cards_amount,omitempty" json:"gift_cards_amount,omitempty" yaml:"gift_cards_amount,omitempty"`
	Gw_base_price                  *string                                 `xml:"gw_base_price,omitempty" json:"gw_base_price,omitempty" yaml:"gw_base_price,omitempty"`
	Gw_price                       *string                                 `xml:"gw_price,omitempty" json:"gw_price,omitempty" yaml:"gw_price,omitempty"`
	Gw_items_base_price            *string                                 `xml:"gw_items_base_price,omitempty" json:"gw_items_base_price,omitempty" yaml:"gw_items_base_price,omitempty"`
	Gw_items_price                 *string                                 `xml:"gw_items_price,omitempty" json:"gw_items_price,omitempty" yaml:"gw_items_price,omitempty"`
	Gw_card_base_price             *string                                 `xml:"gw_card_base_price,omitempty" json:"gw_card_base_price,omitempty" yaml:"gw_card_base_price,omitempty"`
	Gw_card_price                  *string                                 `xml:"gw_card_price,omitempty" json:"gw_card_price,omitempty" yaml:"gw_card_price,omitempty"`
	Gw_base_tax_amount             *string                                 `xml:"gw_base_tax_amount,omitempty" json:"gw_base_tax_amount,omitempty" yaml:"gw_base_tax_amount,omitempty"`
	Gw_tax_amount                  *string                                 `xml:"gw_tax_amount,omitempty" json:"gw_tax_amount,omitempty" yaml:"gw_tax_amount,omitempty"`
	Gw_items_base_tax_amount       *string                                 `xml:"gw_items_base_tax_amount,omitempty" json:"gw_items_base_tax_amount,omitempty" yaml:"gw_items_base_tax_amount,omitempty"`
	Gw_items_tax_amount            *string                                 `xml:"gw_items_tax_amount,omitempty" json:"gw_items_tax_amount,omitempty" yaml:"gw_items_tax_amount,omitempty"`
	Gw_card_base_tax_amount        *string                                 `xml:"gw_card_base_tax_amount,omitempty" json:"gw_card_base_tax_amount,omitempty" yaml:"gw_card_base_tax_amount,omitempty"`
	Gw_card_tax_amount             *string                                 `xml:"gw_card_tax_amount,omitempty" json:"gw_card_tax_amount,omitempty" yaml:"gw_card_tax_amount,omitempty"`
	Base_reward_currency_amount    *string                                 `xml:"base_reward_currency_amount,omitempty" json:"base_reward_currency_amount,omitempty" yaml:"base_reward_currency_amount,omitempty"`
	Reward_currency_amount         *string                                 `xml:"reward_currency_amount,omitempty" json:"reward_currency_amount,omitempty" yaml:"reward_currency_amount,omitempty"`
	Reward_points_balance          *string                                 `xml:"reward_points_balance,omitempty" json:"reward_points_balance,omitempty" yaml:"reward_points_balance,omitempty"`
	Reward_points_balance_refund   *string                                 `xml:"reward_points_balance_refund,omitempty" json:"reward_points_balance_refund,omitempty" yaml:"reward_points_balance_refund,omitempty"`
	Creditmemo_id                  *string                                 `xml:"creditmemo_id,omitempty" json:"creditmemo_id,omitempty" yaml:"creditmemo_id,omitempty"`
	Items                          *SalesOrderCreditmemoItemEntityArray    `xml:"items,omitempty" json:"items,omitempty" yaml:"items,omitempty"`
	Comments                       *SalesOrderCreditmemoCommentEntityArray `xml:"comments,omitempty" json:"comments,omitempty" yaml:"comments,omitempty"`
}

// SalesOrderCreditmemoEntityArray was auto-generated from WSDL.
type SalesOrderCreditmemoEntityArray struct {
	Items []*SalesOrderCreditmemoEntity `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
}

// SalesOrderCreditmemoItemEntity was auto-generated from WSDL.
type SalesOrderCreditmemoItemEntity struct {
	Item_id                          *string `xml:"item_id,omitempty" json:"item_id,omitempty" yaml:"item_id,omitempty"`
	Parent_id                        *string `xml:"parent_id,omitempty" json:"parent_id,omitempty" yaml:"parent_id,omitempty"`
	Weee_tax_applied_row_amount      *string `xml:"weee_tax_applied_row_amount,omitempty" json:"weee_tax_applied_row_amount,omitempty" yaml:"weee_tax_applied_row_amount,omitempty"`
	Base_price                       *string `xml:"base_price,omitempty" json:"base_price,omitempty" yaml:"base_price,omitempty"`
	Base_weee_tax_row_disposition    *string `xml:"base_weee_tax_row_disposition,omitempty" json:"base_weee_tax_row_disposition,omitempty" yaml:"base_weee_tax_row_disposition,omitempty"`
	Tax_amount                       *string `xml:"tax_amount,omitempty" json:"tax_amount,omitempty" yaml:"tax_amount,omitempty"`
	Base_weee_tax_applied_amount     *string `xml:"base_weee_tax_applied_amount,omitempty" json:"base_weee_tax_applied_amount,omitempty" yaml:"base_weee_tax_applied_amount,omitempty"`
	Weee_tax_row_disposition         *string `xml:"weee_tax_row_disposition,omitempty" json:"weee_tax_row_disposition,omitempty" yaml:"weee_tax_row_disposition,omitempty"`
	Base_row_total                   *string `xml:"base_row_total,omitempty" json:"base_row_total,omitempty" yaml:"base_row_total,omitempty"`
	Discount_amount                  *string `xml:"discount_amount,omitempty" json:"discount_amount,omitempty" yaml:"discount_amount,omitempty"`
	Row_total                        *string `xml:"row_total,omitempty" json:"row_total,omitempty" yaml:"row_total,omitempty"`
	Weee_tax_applied_amount          *string `xml:"weee_tax_applied_amount,omitempty" json:"weee_tax_applied_amount,omitempty" yaml:"weee_tax_applied_amount,omitempty"`
	Base_discount_amount             *string `xml:"base_discount_amount,omitempty" json:"base_discount_amount,omitempty" yaml:"base_discount_amount,omitempty"`
	Base_weee_tax_disposition        *string `xml:"base_weee_tax_disposition,omitempty" json:"base_weee_tax_disposition,omitempty" yaml:"base_weee_tax_disposition,omitempty"`
	Price_incl_tax                   *string `xml:"price_incl_tax,omitempty" json:"price_incl_tax,omitempty" yaml:"price_incl_tax,omitempty"`
	Base_tax_amount                  *string `xml:"base_tax_amount,omitempty" json:"base_tax_amount,omitempty" yaml:"base_tax_amount,omitempty"`
	Weee_tax_disposition             *string `xml:"weee_tax_disposition,omitempty" json:"weee_tax_disposition,omitempty" yaml:"weee_tax_disposition,omitempty"`
	Base_price_incl_tax              *string `xml:"base_price_incl_tax,omitempty" json:"base_price_incl_tax,omitempty" yaml:"base_price_incl_tax,omitempty"`
	Qty                              *string `xml:"qty,omitempty" json:"qty,omitempty" yaml:"qty,omitempty"`
	Base_cost                        *string `xml:"base_cost,omitempty" json:"base_cost,omitempty" yaml:"base_cost,omitempty"`
	Base_weee_tax_applied_row_amount *string `xml:"base_weee_tax_applied_row_amount,omitempty" json:"base_weee_tax_applied_row_amount,omitempty" yaml:"base_weee_tax_applied_row_amount,omitempty"`
	Price                            *string `xml:"price,omitempty" json:"price,omitempty" yaml:"price,omitempty"`
	Base_row_total_incl_tax          *string `xml:"base_row_total_incl_tax,omitempty" json:"base_row_total_incl_tax,omitempty" yaml:"base_row_total_incl_tax,omitempty"`
	Row_total_incl_tax               *string `xml:"row_total_incl_tax,omitempty" json:"row_total_incl_tax,omitempty" yaml:"row_total_incl_tax,omitempty"`
	Product_id                       *string `xml:"product_id,omitempty" json:"product_id,omitempty" yaml:"product_id,omitempty"`
	Order_item_id                    *string `xml:"order_item_id,omitempty" json:"order_item_id,omitempty" yaml:"order_item_id,omitempty"`
	Additional_data                  *string `xml:"additional_data,omitempty" json:"additional_data,omitempty" yaml:"additional_data,omitempty"`
	Description                      *string `xml:"description,omitempty" json:"description,omitempty" yaml:"description,omitempty"`
	Weee_tax_applied                 *string `xml:"weee_tax_applied,omitempty" json:"weee_tax_applied,omitempty" yaml:"weee_tax_applied,omitempty"`
	Sku                              *string `xml:"sku,omitempty" json:"sku,omitempty" yaml:"sku,omitempty"`
	Name                             *string `xml:"name,omitempty" json:"name,omitempty" yaml:"name,omitempty"`
	Hidden_tax_amount                *string `xml:"hidden_tax_amount,omitempty" json:"hidden_tax_amount,omitempty" yaml:"hidden_tax_amount,omitempty"`
	Base_hidden_tax_amount           *string `xml:"base_hidden_tax_amount,omitempty" json:"base_hidden_tax_amount,omitempty" yaml:"base_hidden_tax_amount,omitempty"`
}

// SalesOrderCreditmemoItemEntityArray was auto-generated from
// WSDL.
type SalesOrderCreditmemoItemEntityArray struct {
	Items []*SalesOrderCreditmemoItemEntity `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
}

// SalesOrderEntity was auto-generated from WSDL.
type SalesOrderEntity struct {
	Increment_id                *string                             `xml:"increment_id,omitempty" json:"increment_id,omitempty" yaml:"increment_id,omitempty"`
	Parent_id                   *string                             `xml:"parent_id,omitempty" json:"parent_id,omitempty" yaml:"parent_id,omitempty"`
	Store_id                    *string                             `xml:"store_id,omitempty" json:"store_id,omitempty" yaml:"store_id,omitempty"`
	Created_at                  *string                             `xml:"created_at,omitempty" json:"created_at,omitempty" yaml:"created_at,omitempty"`
	Updated_at                  *string                             `xml:"updated_at,omitempty" json:"updated_at,omitempty" yaml:"updated_at,omitempty"`
	Is_active                   *string                             `xml:"is_active,omitempty" json:"is_active,omitempty" yaml:"is_active,omitempty"`
	Customer_id                 *string                             `xml:"customer_id,omitempty" json:"customer_id,omitempty" yaml:"customer_id,omitempty"`
	Tax_amount                  *string                             `xml:"tax_amount,omitempty" json:"tax_amount,omitempty" yaml:"tax_amount,omitempty"`
	Shipping_amount             *string                             `xml:"shipping_amount,omitempty" json:"shipping_amount,omitempty" yaml:"shipping_amount,omitempty"`
	Discount_amount             *string                             `xml:"discount_amount,omitempty" json:"discount_amount,omitempty" yaml:"discount_amount,omitempty"`
	Subtotal                    *string                             `xml:"subtotal,omitempty" json:"subtotal,omitempty" yaml:"subtotal,omitempty"`
	Grand_total                 *string                             `xml:"grand_total,omitempty" json:"grand_total,omitempty" yaml:"grand_total,omitempty"`
	Total_paid                  *string                             `xml:"total_paid,omitempty" json:"total_paid,omitempty" yaml:"total_paid,omitempty"`
	Total_refunded              *string                             `xml:"total_refunded,omitempty" json:"total_refunded,omitempty" yaml:"total_refunded,omitempty"`
	Total_qty_ordered           *string                             `xml:"total_qty_ordered,omitempty" json:"total_qty_ordered,omitempty" yaml:"total_qty_ordered,omitempty"`
	Total_canceled              *string                             `xml:"total_canceled,omitempty" json:"total_canceled,omitempty" yaml:"total_canceled,omitempty"`
	Total_invoiced              *string                             `xml:"total_invoiced,omitempty" json:"total_invoiced,omitempty" yaml:"total_invoiced,omitempty"`
	Total_online_refunded       *string                             `xml:"total_online_refunded,omitempty" json:"total_online_refunded,omitempty" yaml:"total_online_refunded,omitempty"`
	Total_offline_refunded      *string                             `xml:"total_offline_refunded,omitempty" json:"total_offline_refunded,omitempty" yaml:"total_offline_refunded,omitempty"`
	Base_tax_amount             *string                             `xml:"base_tax_amount,omitempty" json:"base_tax_amount,omitempty" yaml:"base_tax_amount,omitempty"`
	Base_shipping_amount        *string                             `xml:"base_shipping_amount,omitempty" json:"base_shipping_amount,omitempty" yaml:"base_shipping_amount,omitempty"`
	Base_discount_amount        *string                             `xml:"base_discount_amount,omitempty" json:"base_discount_amount,omitempty" yaml:"base_discount_amount,omitempty"`
	Base_subtotal               *string                             `xml:"base_subtotal,omitempty" json:"base_subtotal,omitempty" yaml:"base_subtotal,omitempty"`
	Base_grand_total            *string                             `xml:"base_grand_total,omitempty" json:"base_grand_total,omitempty" yaml:"base_grand_total,omitempty"`
	Base_total_paid             *string                             `xml:"base_total_paid,omitempty" json:"base_total_paid,omitempty" yaml:"base_total_paid,omitempty"`
	Base_total_refunded         *string                             `xml:"base_total_refunded,omitempty" json:"base_total_refunded,omitempty" yaml:"base_total_refunded,omitempty"`
	Base_total_qty_ordered      *string                             `xml:"base_total_qty_ordered,omitempty" json:"base_total_qty_ordered,omitempty" yaml:"base_total_qty_ordered,omitempty"`
	Base_total_canceled         *string                             `xml:"base_total_canceled,omitempty" json:"base_total_canceled,omitempty" yaml:"base_total_canceled,omitempty"`
	Base_total_invoiced         *string                             `xml:"base_total_invoiced,omitempty" json:"base_total_invoiced,omitempty" yaml:"base_total_invoiced,omitempty"`
	Base_total_online_refunded  *string                             `xml:"base_total_online_refunded,omitempty" json:"base_total_online_refunded,omitempty" yaml:"base_total_online_refunded,omitempty"`
	Base_total_offline_refunded *string                             `xml:"base_total_offline_refunded,omitempty" json:"base_total_offline_refunded,omitempty" yaml:"base_total_offline_refunded,omitempty"`
	Billing_address_id          *string                             `xml:"billing_address_id,omitempty" json:"billing_address_id,omitempty" yaml:"billing_address_id,omitempty"`
	Billing_firstname           *string                             `xml:"billing_firstname,omitempty" json:"billing_firstname,omitempty" yaml:"billing_firstname,omitempty"`
	Billing_lastname            *string                             `xml:"billing_lastname,omitempty" json:"billing_lastname,omitempty" yaml:"billing_lastname,omitempty"`
	Shipping_address_id         *string                             `xml:"shipping_address_id,omitempty" json:"shipping_address_id,omitempty" yaml:"shipping_address_id,omitempty"`
	Shipping_firstname          *string                             `xml:"shipping_firstname,omitempty" json:"shipping_firstname,omitempty" yaml:"shipping_firstname,omitempty"`
	Shipping_lastname           *string                             `xml:"shipping_lastname,omitempty" json:"shipping_lastname,omitempty" yaml:"shipping_lastname,omitempty"`
	Billing_name                *string                             `xml:"billing_name,omitempty" json:"billing_name,omitempty" yaml:"billing_name,omitempty"`
	Shipping_name               *string                             `xml:"shipping_name,omitempty" json:"shipping_name,omitempty" yaml:"shipping_name,omitempty"`
	Store_to_base_rate          *string                             `xml:"store_to_base_rate,omitempty" json:"store_to_base_rate,omitempty" yaml:"store_to_base_rate,omitempty"`
	Store_to_order_rate         *string                             `xml:"store_to_order_rate,omitempty" json:"store_to_order_rate,omitempty" yaml:"store_to_order_rate,omitempty"`
	Base_to_global_rate         *string                             `xml:"base_to_global_rate,omitempty" json:"base_to_global_rate,omitempty" yaml:"base_to_global_rate,omitempty"`
	Base_to_order_rate          *string                             `xml:"base_to_order_rate,omitempty" json:"base_to_order_rate,omitempty" yaml:"base_to_order_rate,omitempty"`
	Weight                      *string                             `xml:"weight,omitempty" json:"weight,omitempty" yaml:"weight,omitempty"`
	Store_name                  *string                             `xml:"store_name,omitempty" json:"store_name,omitempty" yaml:"store_name,omitempty"`
	Remote_ip                   *string                             `xml:"remote_ip,omitempty" json:"remote_ip,omitempty" yaml:"remote_ip,omitempty"`
	Status                      *string                             `xml:"status,omitempty" json:"status,omitempty" yaml:"status,omitempty"`
	State                       *string                             `xml:"state,omitempty" json:"state,omitempty" yaml:"state,omitempty"`
	Applied_rule_ids            *string                             `xml:"applied_rule_ids,omitempty" json:"applied_rule_ids,omitempty" yaml:"applied_rule_ids,omitempty"`
	Global_currency_code        *string                             `xml:"global_currency_code,omitempty" json:"global_currency_code,omitempty" yaml:"global_currency_code,omitempty"`
	Base_currency_code          *string                             `xml:"base_currency_code,omitempty" json:"base_currency_code,omitempty" yaml:"base_currency_code,omitempty"`
	Store_currency_code         *string                             `xml:"store_currency_code,omitempty" json:"store_currency_code,omitempty" yaml:"store_currency_code,omitempty"`
	Order_currency_code         *string                             `xml:"order_currency_code,omitempty" json:"order_currency_code,omitempty" yaml:"order_currency_code,omitempty"`
	Shipping_method             *string                             `xml:"shipping_method,omitempty" json:"shipping_method,omitempty" yaml:"shipping_method,omitempty"`
	Shipping_description        *string                             `xml:"shipping_description,omitempty" json:"shipping_description,omitempty" yaml:"shipping_description,omitempty"`
	Customer_email              *string                             `xml:"customer_email,omitempty" json:"customer_email,omitempty" yaml:"customer_email,omitempty"`
	Customer_firstname          *string                             `xml:"customer_firstname,omitempty" json:"customer_firstname,omitempty" yaml:"customer_firstname,omitempty"`
	Customer_lastname           *string                             `xml:"customer_lastname,omitempty" json:"customer_lastname,omitempty" yaml:"customer_lastname,omitempty"`
	Quote_id                    *string                             `xml:"quote_id,omitempty" json:"quote_id,omitempty" yaml:"quote_id,omitempty"`
	Is_virtual                  *string                             `xml:"is_virtual,omitempty" json:"is_virtual,omitempty" yaml:"is_virtual,omitempty"`
	Customer_group_id           *string                             `xml:"customer_group_id,omitempty" json:"customer_group_id,omitempty" yaml:"customer_group_id,omitempty"`
	Customer_note_notify        *string                             `xml:"customer_note_notify,omitempty" json:"customer_note_notify,omitempty" yaml:"customer_note_notify,omitempty"`
	Customer_is_guest           *string                             `xml:"customer_is_guest,omitempty" json:"customer_is_guest,omitempty" yaml:"customer_is_guest,omitempty"`
	Email_sent                  *string                             `xml:"email_sent,omitempty" json:"email_sent,omitempty" yaml:"email_sent,omitempty"`
	Order_id                    *string                             `xml:"order_id,omitempty" json:"order_id,omitempty" yaml:"order_id,omitempty"`
	Gift_message_id             *string                             `xml:"gift_message_id,omitempty" json:"gift_message_id,omitempty" yaml:"gift_message_id,omitempty"`
	Gift_message                *string                             `xml:"gift_message,omitempty" json:"gift_message,omitempty" yaml:"gift_message,omitempty"`
	Shipping_address            *SalesOrderAddressEntity            `xml:"shipping_address,omitempty" json:"shipping_address,omitempty" yaml:"shipping_address,omitempty"`
	Billing_address             *SalesOrderAddressEntity            `xml:"billing_address,omitempty" json:"billing_address,omitempty" yaml:"billing_address,omitempty"`
	Items                       *SalesOrderItemEntityArray          `xml:"items,omitempty" json:"items,omitempty" yaml:"items,omitempty"`
	Payment                     *SalesOrderPaymentEntity            `xml:"payment,omitempty" json:"payment,omitempty" yaml:"payment,omitempty"`
	Status_history              *SalesOrderStatusHistoryEntityArray `xml:"status_history,omitempty" json:"status_history,omitempty" yaml:"status_history,omitempty"`
}

// SalesOrderInvoiceCommentEntity was auto-generated from WSDL.
type SalesOrderInvoiceCommentEntity struct {
	Increment_id         *string `xml:"increment_id,omitempty" json:"increment_id,omitempty" yaml:"increment_id,omitempty"`
	Parent_id            *string `xml:"parent_id,omitempty" json:"parent_id,omitempty" yaml:"parent_id,omitempty"`
	Created_at           *string `xml:"created_at,omitempty" json:"created_at,omitempty" yaml:"created_at,omitempty"`
	Updated_at           *string `xml:"updated_at,omitempty" json:"updated_at,omitempty" yaml:"updated_at,omitempty"`
	Is_active            *string `xml:"is_active,omitempty" json:"is_active,omitempty" yaml:"is_active,omitempty"`
	Comment              *string `xml:"comment,omitempty" json:"comment,omitempty" yaml:"comment,omitempty"`
	Is_customer_notified *string `xml:"is_customer_notified,omitempty" json:"is_customer_notified,omitempty" yaml:"is_customer_notified,omitempty"`
	Comment_id           *string `xml:"comment_id,omitempty" json:"comment_id,omitempty" yaml:"comment_id,omitempty"`
}

// SalesOrderInvoiceCommentEntityArray was auto-generated from
// WSDL.
type SalesOrderInvoiceCommentEntityArray struct {
	Items []*SalesOrderInvoiceCommentEntity `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
}

// SalesOrderInvoiceEntity was auto-generated from WSDL.
type SalesOrderInvoiceEntity struct {
	Increment_id         *string                              `xml:"increment_id,omitempty" json:"increment_id,omitempty" yaml:"increment_id,omitempty"`
	Parent_id            *string                              `xml:"parent_id,omitempty" json:"parent_id,omitempty" yaml:"parent_id,omitempty"`
	Store_id             *string                              `xml:"store_id,omitempty" json:"store_id,omitempty" yaml:"store_id,omitempty"`
	Created_at           *string                              `xml:"created_at,omitempty" json:"created_at,omitempty" yaml:"created_at,omitempty"`
	Updated_at           *string                              `xml:"updated_at,omitempty" json:"updated_at,omitempty" yaml:"updated_at,omitempty"`
	Is_active            *string                              `xml:"is_active,omitempty" json:"is_active,omitempty" yaml:"is_active,omitempty"`
	Global_currency_code *string                              `xml:"global_currency_code,omitempty" json:"global_currency_code,omitempty" yaml:"global_currency_code,omitempty"`
	Base_currency_code   *string                              `xml:"base_currency_code,omitempty" json:"base_currency_code,omitempty" yaml:"base_currency_code,omitempty"`
	Store_currency_code  *string                              `xml:"store_currency_code,omitempty" json:"store_currency_code,omitempty" yaml:"store_currency_code,omitempty"`
	Order_currency_code  *string                              `xml:"order_currency_code,omitempty" json:"order_currency_code,omitempty" yaml:"order_currency_code,omitempty"`
	Store_to_base_rate   *string                              `xml:"store_to_base_rate,omitempty" json:"store_to_base_rate,omitempty" yaml:"store_to_base_rate,omitempty"`
	Store_to_order_rate  *string                              `xml:"store_to_order_rate,omitempty" json:"store_to_order_rate,omitempty" yaml:"store_to_order_rate,omitempty"`
	Base_to_global_rate  *string                              `xml:"base_to_global_rate,omitempty" json:"base_to_global_rate,omitempty" yaml:"base_to_global_rate,omitempty"`
	Base_to_order_rate   *string                              `xml:"base_to_order_rate,omitempty" json:"base_to_order_rate,omitempty" yaml:"base_to_order_rate,omitempty"`
	Subtotal             *string                              `xml:"subtotal,omitempty" json:"subtotal,omitempty" yaml:"subtotal,omitempty"`
	Base_subtotal        *string                              `xml:"base_subtotal,omitempty" json:"base_subtotal,omitempty" yaml:"base_subtotal,omitempty"`
	Base_grand_total     *string                              `xml:"base_grand_total,omitempty" json:"base_grand_total,omitempty" yaml:"base_grand_total,omitempty"`
	Discount_amount      *string                              `xml:"discount_amount,omitempty" json:"discount_amount,omitempty" yaml:"discount_amount,omitempty"`
	Base_discount_amount *string                              `xml:"base_discount_amount,omitempty" json:"base_discount_amount,omitempty" yaml:"base_discount_amount,omitempty"`
	Shipping_amount      *string                              `xml:"shipping_amount,omitempty" json:"shipping_amount,omitempty" yaml:"shipping_amount,omitempty"`
	Base_shipping_amount *string                              `xml:"base_shipping_amount,omitempty" json:"base_shipping_amount,omitempty" yaml:"base_shipping_amount,omitempty"`
	Tax_amount           *string                              `xml:"tax_amount,omitempty" json:"tax_amount,omitempty" yaml:"tax_amount,omitempty"`
	Base_tax_amount      *string                              `xml:"base_tax_amount,omitempty" json:"base_tax_amount,omitempty" yaml:"base_tax_amount,omitempty"`
	Billing_address_id   *string                              `xml:"billing_address_id,omitempty" json:"billing_address_id,omitempty" yaml:"billing_address_id,omitempty"`
	Billing_firstname    *string                              `xml:"billing_firstname,omitempty" json:"billing_firstname,omitempty" yaml:"billing_firstname,omitempty"`
	Billing_lastname     *string                              `xml:"billing_lastname,omitempty" json:"billing_lastname,omitempty" yaml:"billing_lastname,omitempty"`
	Order_id             *string                              `xml:"order_id,omitempty" json:"order_id,omitempty" yaml:"order_id,omitempty"`
	Order_increment_id   *string                              `xml:"order_increment_id,omitempty" json:"order_increment_id,omitempty" yaml:"order_increment_id,omitempty"`
	Order_created_at     *string                              `xml:"order_created_at,omitempty" json:"order_created_at,omitempty" yaml:"order_created_at,omitempty"`
	State                *string                              `xml:"state,omitempty" json:"state,omitempty" yaml:"state,omitempty"`
	Grand_total          *string                              `xml:"grand_total,omitempty" json:"grand_total,omitempty" yaml:"grand_total,omitempty"`
	Invoice_id           *string                              `xml:"invoice_id,omitempty" json:"invoice_id,omitempty" yaml:"invoice_id,omitempty"`
	Items                *SalesOrderInvoiceItemEntityArray    `xml:"items,omitempty" json:"items,omitempty" yaml:"items,omitempty"`
	Comments             *SalesOrderInvoiceCommentEntityArray `xml:"comments,omitempty" json:"comments,omitempty" yaml:"comments,omitempty"`
}

// SalesOrderInvoiceEntityArray was auto-generated from WSDL.
type SalesOrderInvoiceEntityArray struct {
	Items []*SalesOrderInvoiceEntity `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
}

// SalesOrderInvoiceItemEntity was auto-generated from WSDL.
type SalesOrderInvoiceItemEntity struct {
	Increment_id                     *string `xml:"increment_id,omitempty" json:"increment_id,omitempty" yaml:"increment_id,omitempty"`
	Parent_id                        *string `xml:"parent_id,omitempty" json:"parent_id,omitempty" yaml:"parent_id,omitempty"`
	Created_at                       *string `xml:"created_at,omitempty" json:"created_at,omitempty" yaml:"created_at,omitempty"`
	Updated_at                       *string `xml:"updated_at,omitempty" json:"updated_at,omitempty" yaml:"updated_at,omitempty"`
	Is_active                        *string `xml:"is_active,omitempty" json:"is_active,omitempty" yaml:"is_active,omitempty"`
	Weee_tax_applied                 *string `xml:"weee_tax_applied,omitempty" json:"weee_tax_applied,omitempty" yaml:"weee_tax_applied,omitempty"`
	Qty                              *string `xml:"qty,omitempty" json:"qty,omitempty" yaml:"qty,omitempty"`
	Cost                             *string `xml:"cost,omitempty" json:"cost,omitempty" yaml:"cost,omitempty"`
	Price                            *string `xml:"price,omitempty" json:"price,omitempty" yaml:"price,omitempty"`
	Tax_amount                       *string `xml:"tax_amount,omitempty" json:"tax_amount,omitempty" yaml:"tax_amount,omitempty"`
	Row_total                        *string `xml:"row_total,omitempty" json:"row_total,omitempty" yaml:"row_total,omitempty"`
	Base_price                       *string `xml:"base_price,omitempty" json:"base_price,omitempty" yaml:"base_price,omitempty"`
	Base_tax_amount                  *string `xml:"base_tax_amount,omitempty" json:"base_tax_amount,omitempty" yaml:"base_tax_amount,omitempty"`
	Base_row_total                   *string `xml:"base_row_total,omitempty" json:"base_row_total,omitempty" yaml:"base_row_total,omitempty"`
	Base_weee_tax_applied_amount     *string `xml:"base_weee_tax_applied_amount,omitempty" json:"base_weee_tax_applied_amount,omitempty" yaml:"base_weee_tax_applied_amount,omitempty"`
	Base_weee_tax_applied_row_amount *string `xml:"base_weee_tax_applied_row_amount,omitempty" json:"base_weee_tax_applied_row_amount,omitempty" yaml:"base_weee_tax_applied_row_amount,omitempty"`
	Weee_tax_applied_amount          *string `xml:"weee_tax_applied_amount,omitempty" json:"weee_tax_applied_amount,omitempty" yaml:"weee_tax_applied_amount,omitempty"`
	Weee_tax_applied_row_amount      *string `xml:"weee_tax_applied_row_amount,omitempty" json:"weee_tax_applied_row_amount,omitempty" yaml:"weee_tax_applied_row_amount,omitempty"`
	Weee_tax_disposition             *string `xml:"weee_tax_disposition,omitempty" json:"weee_tax_disposition,omitempty" yaml:"weee_tax_disposition,omitempty"`
	Weee_tax_row_disposition         *string `xml:"weee_tax_row_disposition,omitempty" json:"weee_tax_row_disposition,omitempty" yaml:"weee_tax_row_disposition,omitempty"`
	Base_weee_tax_disposition        *string `xml:"base_weee_tax_disposition,omitempty" json:"base_weee_tax_disposition,omitempty" yaml:"base_weee_tax_disposition,omitempty"`
	Base_weee_tax_row_disposition    *string `xml:"base_weee_tax_row_disposition,omitempty" json:"base_weee_tax_row_disposition,omitempty" yaml:"base_weee_tax_row_disposition,omitempty"`
	Sku                              *string `xml:"sku,omitempty" json:"sku,omitempty" yaml:"sku,omitempty"`
	Name                             *string `xml:"name,omitempty" json:"name,omitempty" yaml:"name,omitempty"`
	Order_item_id                    *string `xml:"order_item_id,omitempty" json:"order_item_id,omitempty" yaml:"order_item_id,omitempty"`
	Product_id                       *string `xml:"product_id,omitempty" json:"product_id,omitempty" yaml:"product_id,omitempty"`
	Item_id                          *string `xml:"item_id,omitempty" json:"item_id,omitempty" yaml:"item_id,omitempty"`
}

// SalesOrderInvoiceItemEntityArray was auto-generated from WSDL.
type SalesOrderInvoiceItemEntityArray struct {
	Items []*SalesOrderInvoiceItemEntity `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
}

// SalesOrderItemEntity was auto-generated from WSDL.
type SalesOrderItemEntity struct {
	Item_id                          *string `xml:"item_id,omitempty" json:"item_id,omitempty" yaml:"item_id,omitempty"`
	Order_id                         *string `xml:"order_id,omitempty" json:"order_id,omitempty" yaml:"order_id,omitempty"`
	Quote_item_id                    *string `xml:"quote_item_id,omitempty" json:"quote_item_id,omitempty" yaml:"quote_item_id,omitempty"`
	Created_at                       *string `xml:"created_at,omitempty" json:"created_at,omitempty" yaml:"created_at,omitempty"`
	Updated_at                       *string `xml:"updated_at,omitempty" json:"updated_at,omitempty" yaml:"updated_at,omitempty"`
	Product_id                       *string `xml:"product_id,omitempty" json:"product_id,omitempty" yaml:"product_id,omitempty"`
	Product_type                     *string `xml:"product_type,omitempty" json:"product_type,omitempty" yaml:"product_type,omitempty"`
	Product_options                  *string `xml:"product_options,omitempty" json:"product_options,omitempty" yaml:"product_options,omitempty"`
	Weight                           *string `xml:"weight,omitempty" json:"weight,omitempty" yaml:"weight,omitempty"`
	Is_virtual                       *string `xml:"is_virtual,omitempty" json:"is_virtual,omitempty" yaml:"is_virtual,omitempty"`
	Sku                              *string `xml:"sku,omitempty" json:"sku,omitempty" yaml:"sku,omitempty"`
	Name                             *string `xml:"name,omitempty" json:"name,omitempty" yaml:"name,omitempty"`
	Applied_rule_ids                 *string `xml:"applied_rule_ids,omitempty" json:"applied_rule_ids,omitempty" yaml:"applied_rule_ids,omitempty"`
	Free_shipping                    *string `xml:"free_shipping,omitempty" json:"free_shipping,omitempty" yaml:"free_shipping,omitempty"`
	Is_qty_decimal                   *string `xml:"is_qty_decimal,omitempty" json:"is_qty_decimal,omitempty" yaml:"is_qty_decimal,omitempty"`
	No_discount                      *string `xml:"no_discount,omitempty" json:"no_discount,omitempty" yaml:"no_discount,omitempty"`
	Qty_canceled                     *string `xml:"qty_canceled,omitempty" json:"qty_canceled,omitempty" yaml:"qty_canceled,omitempty"`
	Qty_invoiced                     *string `xml:"qty_invoiced,omitempty" json:"qty_invoiced,omitempty" yaml:"qty_invoiced,omitempty"`
	Qty_ordered                      *string `xml:"qty_ordered,omitempty" json:"qty_ordered,omitempty" yaml:"qty_ordered,omitempty"`
	Qty_refunded                     *string `xml:"qty_refunded,omitempty" json:"qty_refunded,omitempty" yaml:"qty_refunded,omitempty"`
	Qty_shipped                      *string `xml:"qty_shipped,omitempty" json:"qty_shipped,omitempty" yaml:"qty_shipped,omitempty"`
	Cost                             *string `xml:"cost,omitempty" json:"cost,omitempty" yaml:"cost,omitempty"`
	Price                            *string `xml:"price,omitempty" json:"price,omitempty" yaml:"price,omitempty"`
	Base_price                       *string `xml:"base_price,omitempty" json:"base_price,omitempty" yaml:"base_price,omitempty"`
	Original_price                   *string `xml:"original_price,omitempty" json:"original_price,omitempty" yaml:"original_price,omitempty"`
	Base_original_price              *string `xml:"base_original_price,omitempty" json:"base_original_price,omitempty" yaml:"base_original_price,omitempty"`
	Tax_percent                      *string `xml:"tax_percent,omitempty" json:"tax_percent,omitempty" yaml:"tax_percent,omitempty"`
	Tax_amount                       *string `xml:"tax_amount,omitempty" json:"tax_amount,omitempty" yaml:"tax_amount,omitempty"`
	Base_tax_amount                  *string `xml:"base_tax_amount,omitempty" json:"base_tax_amount,omitempty" yaml:"base_tax_amount,omitempty"`
	Tax_invoiced                     *string `xml:"tax_invoiced,omitempty" json:"tax_invoiced,omitempty" yaml:"tax_invoiced,omitempty"`
	Base_tax_invoiced                *string `xml:"base_tax_invoiced,omitempty" json:"base_tax_invoiced,omitempty" yaml:"base_tax_invoiced,omitempty"`
	Discount_percent                 *string `xml:"discount_percent,omitempty" json:"discount_percent,omitempty" yaml:"discount_percent,omitempty"`
	Discount_amount                  *string `xml:"discount_amount,omitempty" json:"discount_amount,omitempty" yaml:"discount_amount,omitempty"`
	Base_discount_amount             *string `xml:"base_discount_amount,omitempty" json:"base_discount_amount,omitempty" yaml:"base_discount_amount,omitempty"`
	Discount_invoiced                *string `xml:"discount_invoiced,omitempty" json:"discount_invoiced,omitempty" yaml:"discount_invoiced,omitempty"`
	Base_discount_invoiced           *string `xml:"base_discount_invoiced,omitempty" json:"base_discount_invoiced,omitempty" yaml:"base_discount_invoiced,omitempty"`
	Amount_refunded                  *string `xml:"amount_refunded,omitempty" json:"amount_refunded,omitempty" yaml:"amount_refunded,omitempty"`
	Base_amount_refunded             *string `xml:"base_amount_refunded,omitempty" json:"base_amount_refunded,omitempty" yaml:"base_amount_refunded,omitempty"`
	Row_total                        *string `xml:"row_total,omitempty" json:"row_total,omitempty" yaml:"row_total,omitempty"`
	Base_row_total                   *string `xml:"base_row_total,omitempty" json:"base_row_total,omitempty" yaml:"base_row_total,omitempty"`
	Row_invoiced                     *string `xml:"row_invoiced,omitempty" json:"row_invoiced,omitempty" yaml:"row_invoiced,omitempty"`
	Base_row_invoiced                *string `xml:"base_row_invoiced,omitempty" json:"base_row_invoiced,omitempty" yaml:"base_row_invoiced,omitempty"`
	Row_weight                       *string `xml:"row_weight,omitempty" json:"row_weight,omitempty" yaml:"row_weight,omitempty"`
	Gift_message_id                  *string `xml:"gift_message_id,omitempty" json:"gift_message_id,omitempty" yaml:"gift_message_id,omitempty"`
	Gift_message                     *string `xml:"gift_message,omitempty" json:"gift_message,omitempty" yaml:"gift_message,omitempty"`
	Gift_message_available           *string `xml:"gift_message_available,omitempty" json:"gift_message_available,omitempty" yaml:"gift_message_available,omitempty"`
	Base_tax_before_discount         *string `xml:"base_tax_before_discount,omitempty" json:"base_tax_before_discount,omitempty" yaml:"base_tax_before_discount,omitempty"`
	Tax_before_discount              *string `xml:"tax_before_discount,omitempty" json:"tax_before_discount,omitempty" yaml:"tax_before_discount,omitempty"`
	Weee_tax_applied                 *string `xml:"weee_tax_applied,omitempty" json:"weee_tax_applied,omitempty" yaml:"weee_tax_applied,omitempty"`
	Weee_tax_applied_amount          *string `xml:"weee_tax_applied_amount,omitempty" json:"weee_tax_applied_amount,omitempty" yaml:"weee_tax_applied_amount,omitempty"`
	Weee_tax_applied_row_amount      *string `xml:"weee_tax_applied_row_amount,omitempty" json:"weee_tax_applied_row_amount,omitempty" yaml:"weee_tax_applied_row_amount,omitempty"`
	Base_weee_tax_applied_amount     *string `xml:"base_weee_tax_applied_amount,omitempty" json:"base_weee_tax_applied_amount,omitempty" yaml:"base_weee_tax_applied_amount,omitempty"`
	Base_weee_tax_applied_row_amount *string `xml:"base_weee_tax_applied_row_amount,omitempty" json:"base_weee_tax_applied_row_amount,omitempty" yaml:"base_weee_tax_applied_row_amount,omitempty"`
	Weee_tax_disposition             *string `xml:"weee_tax_disposition,omitempty" json:"weee_tax_disposition,omitempty" yaml:"weee_tax_disposition,omitempty"`
	Weee_tax_row_disposition         *string `xml:"weee_tax_row_disposition,omitempty" json:"weee_tax_row_disposition,omitempty" yaml:"weee_tax_row_disposition,omitempty"`
	Base_weee_tax_disposition        *string `xml:"base_weee_tax_disposition,omitempty" json:"base_weee_tax_disposition,omitempty" yaml:"base_weee_tax_disposition,omitempty"`
	Base_weee_tax_row_disposition    *string `xml:"base_weee_tax_row_disposition,omitempty" json:"base_weee_tax_row_disposition,omitempty" yaml:"base_weee_tax_row_disposition,omitempty"`
}

// SalesOrderItemEntityArray was auto-generated from WSDL.
type SalesOrderItemEntityArray struct {
	Items []*SalesOrderItemEntity `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
}

// SalesOrderListEntity was auto-generated from WSDL.
type SalesOrderListEntity struct {
	Increment_id                         *string `xml:"increment_id,omitempty" json:"increment_id,omitempty" yaml:"increment_id,omitempty"`
	Store_id                             *string `xml:"store_id,omitempty" json:"store_id,omitempty" yaml:"store_id,omitempty"`
	Created_at                           *string `xml:"created_at,omitempty" json:"created_at,omitempty" yaml:"created_at,omitempty"`
	Updated_at                           *string `xml:"updated_at,omitempty" json:"updated_at,omitempty" yaml:"updated_at,omitempty"`
	Customer_id                          *string `xml:"customer_id,omitempty" json:"customer_id,omitempty" yaml:"customer_id,omitempty"`
	Tax_amount                           *string `xml:"tax_amount,omitempty" json:"tax_amount,omitempty" yaml:"tax_amount,omitempty"`
	Shipping_amount                      *string `xml:"shipping_amount,omitempty" json:"shipping_amount,omitempty" yaml:"shipping_amount,omitempty"`
	Discount_amount                      *string `xml:"discount_amount,omitempty" json:"discount_amount,omitempty" yaml:"discount_amount,omitempty"`
	Subtotal                             *string `xml:"subtotal,omitempty" json:"subtotal,omitempty" yaml:"subtotal,omitempty"`
	Grand_total                          *string `xml:"grand_total,omitempty" json:"grand_total,omitempty" yaml:"grand_total,omitempty"`
	Total_paid                           *string `xml:"total_paid,omitempty" json:"total_paid,omitempty" yaml:"total_paid,omitempty"`
	Total_refunded                       *string `xml:"total_refunded,omitempty" json:"total_refunded,omitempty" yaml:"total_refunded,omitempty"`
	Total_qty_ordered                    *string `xml:"total_qty_ordered,omitempty" json:"total_qty_ordered,omitempty" yaml:"total_qty_ordered,omitempty"`
	Total_canceled                       *string `xml:"total_canceled,omitempty" json:"total_canceled,omitempty" yaml:"total_canceled,omitempty"`
	Total_invoiced                       *string `xml:"total_invoiced,omitempty" json:"total_invoiced,omitempty" yaml:"total_invoiced,omitempty"`
	Total_online_refunded                *string `xml:"total_online_refunded,omitempty" json:"total_online_refunded,omitempty" yaml:"total_online_refunded,omitempty"`
	Total_offline_refunded               *string `xml:"total_offline_refunded,omitempty" json:"total_offline_refunded,omitempty" yaml:"total_offline_refunded,omitempty"`
	Base_tax_amount                      *string `xml:"base_tax_amount,omitempty" json:"base_tax_amount,omitempty" yaml:"base_tax_amount,omitempty"`
	Base_shipping_amount                 *string `xml:"base_shipping_amount,omitempty" json:"base_shipping_amount,omitempty" yaml:"base_shipping_amount,omitempty"`
	Base_discount_amount                 *string `xml:"base_discount_amount,omitempty" json:"base_discount_amount,omitempty" yaml:"base_discount_amount,omitempty"`
	Base_subtotal                        *string `xml:"base_subtotal,omitempty" json:"base_subtotal,omitempty" yaml:"base_subtotal,omitempty"`
	Base_grand_total                     *string `xml:"base_grand_total,omitempty" json:"base_grand_total,omitempty" yaml:"base_grand_total,omitempty"`
	Base_total_paid                      *string `xml:"base_total_paid,omitempty" json:"base_total_paid,omitempty" yaml:"base_total_paid,omitempty"`
	Base_total_refunded                  *string `xml:"base_total_refunded,omitempty" json:"base_total_refunded,omitempty" yaml:"base_total_refunded,omitempty"`
	Base_total_qty_ordered               *string `xml:"base_total_qty_ordered,omitempty" json:"base_total_qty_ordered,omitempty" yaml:"base_total_qty_ordered,omitempty"`
	Base_total_canceled                  *string `xml:"base_total_canceled,omitempty" json:"base_total_canceled,omitempty" yaml:"base_total_canceled,omitempty"`
	Base_total_invoiced                  *string `xml:"base_total_invoiced,omitempty" json:"base_total_invoiced,omitempty" yaml:"base_total_invoiced,omitempty"`
	Base_total_online_refunded           *string `xml:"base_total_online_refunded,omitempty" json:"base_total_online_refunded,omitempty" yaml:"base_total_online_refunded,omitempty"`
	Base_total_offline_refunded          *string `xml:"base_total_offline_refunded,omitempty" json:"base_total_offline_refunded,omitempty" yaml:"base_total_offline_refunded,omitempty"`
	Billing_address_id                   *string `xml:"billing_address_id,omitempty" json:"billing_address_id,omitempty" yaml:"billing_address_id,omitempty"`
	Billing_firstname                    *string `xml:"billing_firstname,omitempty" json:"billing_firstname,omitempty" yaml:"billing_firstname,omitempty"`
	Billing_lastname                     *string `xml:"billing_lastname,omitempty" json:"billing_lastname,omitempty" yaml:"billing_lastname,omitempty"`
	Shipping_address_id                  *string `xml:"shipping_address_id,omitempty" json:"shipping_address_id,omitempty" yaml:"shipping_address_id,omitempty"`
	Shipping_firstname                   *string `xml:"shipping_firstname,omitempty" json:"shipping_firstname,omitempty" yaml:"shipping_firstname,omitempty"`
	Shipping_lastname                    *string `xml:"shipping_lastname,omitempty" json:"shipping_lastname,omitempty" yaml:"shipping_lastname,omitempty"`
	Billing_name                         *string `xml:"billing_name,omitempty" json:"billing_name,omitempty" yaml:"billing_name,omitempty"`
	Shipping_name                        *string `xml:"shipping_name,omitempty" json:"shipping_name,omitempty" yaml:"shipping_name,omitempty"`
	Store_to_base_rate                   *string `xml:"store_to_base_rate,omitempty" json:"store_to_base_rate,omitempty" yaml:"store_to_base_rate,omitempty"`
	Store_to_order_rate                  *string `xml:"store_to_order_rate,omitempty" json:"store_to_order_rate,omitempty" yaml:"store_to_order_rate,omitempty"`
	Base_to_global_rate                  *string `xml:"base_to_global_rate,omitempty" json:"base_to_global_rate,omitempty" yaml:"base_to_global_rate,omitempty"`
	Base_to_order_rate                   *string `xml:"base_to_order_rate,omitempty" json:"base_to_order_rate,omitempty" yaml:"base_to_order_rate,omitempty"`
	Weight                               *string `xml:"weight,omitempty" json:"weight,omitempty" yaml:"weight,omitempty"`
	Store_name                           *string `xml:"store_name,omitempty" json:"store_name,omitempty" yaml:"store_name,omitempty"`
	Remote_ip                            *string `xml:"remote_ip,omitempty" json:"remote_ip,omitempty" yaml:"remote_ip,omitempty"`
	Status                               *string `xml:"status,omitempty" json:"status,omitempty" yaml:"status,omitempty"`
	State                                *string `xml:"state,omitempty" json:"state,omitempty" yaml:"state,omitempty"`
	Applied_rule_ids                     *string `xml:"applied_rule_ids,omitempty" json:"applied_rule_ids,omitempty" yaml:"applied_rule_ids,omitempty"`
	Global_currency_code                 *string `xml:"global_currency_code,omitempty" json:"global_currency_code,omitempty" yaml:"global_currency_code,omitempty"`
	Base_currency_code                   *string `xml:"base_currency_code,omitempty" json:"base_currency_code,omitempty" yaml:"base_currency_code,omitempty"`
	Store_currency_code                  *string `xml:"store_currency_code,omitempty" json:"store_currency_code,omitempty" yaml:"store_currency_code,omitempty"`
	Order_currency_code                  *string `xml:"order_currency_code,omitempty" json:"order_currency_code,omitempty" yaml:"order_currency_code,omitempty"`
	Shipping_method                      *string `xml:"shipping_method,omitempty" json:"shipping_method,omitempty" yaml:"shipping_method,omitempty"`
	Shipping_description                 *string `xml:"shipping_description,omitempty" json:"shipping_description,omitempty" yaml:"shipping_description,omitempty"`
	Customer_email                       *string `xml:"customer_email,omitempty" json:"customer_email,omitempty" yaml:"customer_email,omitempty"`
	Customer_firstname                   *string `xml:"customer_firstname,omitempty" json:"customer_firstname,omitempty" yaml:"customer_firstname,omitempty"`
	Customer_lastname                    *string `xml:"customer_lastname,omitempty" json:"customer_lastname,omitempty" yaml:"customer_lastname,omitempty"`
	Quote_id                             *string `xml:"quote_id,omitempty" json:"quote_id,omitempty" yaml:"quote_id,omitempty"`
	Is_virtual                           *string `xml:"is_virtual,omitempty" json:"is_virtual,omitempty" yaml:"is_virtual,omitempty"`
	Customer_group_id                    *string `xml:"customer_group_id,omitempty" json:"customer_group_id,omitempty" yaml:"customer_group_id,omitempty"`
	Customer_note_notify                 *string `xml:"customer_note_notify,omitempty" json:"customer_note_notify,omitempty" yaml:"customer_note_notify,omitempty"`
	Customer_is_guest                    *string `xml:"customer_is_guest,omitempty" json:"customer_is_guest,omitempty" yaml:"customer_is_guest,omitempty"`
	Email_sent                           *string `xml:"email_sent,omitempty" json:"email_sent,omitempty" yaml:"email_sent,omitempty"`
	Order_id                             *string `xml:"order_id,omitempty" json:"order_id,omitempty" yaml:"order_id,omitempty"`
	Gift_message_id                      *string `xml:"gift_message_id,omitempty" json:"gift_message_id,omitempty" yaml:"gift_message_id,omitempty"`
	Coupon_code                          *string `xml:"coupon_code,omitempty" json:"coupon_code,omitempty" yaml:"coupon_code,omitempty"`
	Protect_code                         *string `xml:"protect_code,omitempty" json:"protect_code,omitempty" yaml:"protect_code,omitempty"`
	Base_discount_canceled               *string `xml:"base_discount_canceled,omitempty" json:"base_discount_canceled,omitempty" yaml:"base_discount_canceled,omitempty"`
	Base_discount_invoiced               *string `xml:"base_discount_invoiced,omitempty" json:"base_discount_invoiced,omitempty" yaml:"base_discount_invoiced,omitempty"`
	Base_discount_refunded               *string `xml:"base_discount_refunded,omitempty" json:"base_discount_refunded,omitempty" yaml:"base_discount_refunded,omitempty"`
	Base_shipping_canceled               *string `xml:"base_shipping_canceled,omitempty" json:"base_shipping_canceled,omitempty" yaml:"base_shipping_canceled,omitempty"`
	Base_shipping_invoiced               *string `xml:"base_shipping_invoiced,omitempty" json:"base_shipping_invoiced,omitempty" yaml:"base_shipping_invoiced,omitempty"`
	Base_shipping_refunded               *string `xml:"base_shipping_refunded,omitempty" json:"base_shipping_refunded,omitempty" yaml:"base_shipping_refunded,omitempty"`
	Base_shipping_tax_amount             *string `xml:"base_shipping_tax_amount,omitempty" json:"base_shipping_tax_amount,omitempty" yaml:"base_shipping_tax_amount,omitempty"`
	Base_shipping_tax_refunded           *string `xml:"base_shipping_tax_refunded,omitempty" json:"base_shipping_tax_refunded,omitempty" yaml:"base_shipping_tax_refunded,omitempty"`
	Base_subtotal_canceled               *string `xml:"base_subtotal_canceled,omitempty" json:"base_subtotal_canceled,omitempty" yaml:"base_subtotal_canceled,omitempty"`
	Base_subtotal_invoiced               *string `xml:"base_subtotal_invoiced,omitempty" json:"base_subtotal_invoiced,omitempty" yaml:"base_subtotal_invoiced,omitempty"`
	Base_subtotal_refunded               *string `xml:"base_subtotal_refunded,omitempty" json:"base_subtotal_refunded,omitempty" yaml:"base_subtotal_refunded,omitempty"`
	Base_tax_canceled                    *string `xml:"base_tax_canceled,omitempty" json:"base_tax_canceled,omitempty" yaml:"base_tax_canceled,omitempty"`
	Base_tax_invoiced                    *string `xml:"base_tax_invoiced,omitempty" json:"base_tax_invoiced,omitempty" yaml:"base_tax_invoiced,omitempty"`
	Base_tax_refunded                    *string `xml:"base_tax_refunded,omitempty" json:"base_tax_refunded,omitempty" yaml:"base_tax_refunded,omitempty"`
	Base_total_invoiced_cost             *string `xml:"base_total_invoiced_cost,omitempty" json:"base_total_invoiced_cost,omitempty" yaml:"base_total_invoiced_cost,omitempty"`
	Discount_canceled                    *string `xml:"discount_canceled,omitempty" json:"discount_canceled,omitempty" yaml:"discount_canceled,omitempty"`
	Discount_invoiced                    *string `xml:"discount_invoiced,omitempty" json:"discount_invoiced,omitempty" yaml:"discount_invoiced,omitempty"`
	Discount_refunded                    *string `xml:"discount_refunded,omitempty" json:"discount_refunded,omitempty" yaml:"discount_refunded,omitempty"`
	Shipping_canceled                    *string `xml:"shipping_canceled,omitempty" json:"shipping_canceled,omitempty" yaml:"shipping_canceled,omitempty"`
	Shipping_invoiced                    *string `xml:"shipping_invoiced,omitempty" json:"shipping_invoiced,omitempty" yaml:"shipping_invoiced,omitempty"`
	Shipping_refunded                    *string `xml:"shipping_refunded,omitempty" json:"shipping_refunded,omitempty" yaml:"shipping_refunded,omitempty"`
	Shipping_tax_amount                  *string `xml:"shipping_tax_amount,omitempty" json:"shipping_tax_amount,omitempty" yaml:"shipping_tax_amount,omitempty"`
	Shipping_tax_refunded                *string `xml:"shipping_tax_refunded,omitempty" json:"shipping_tax_refunded,omitempty" yaml:"shipping_tax_refunded,omitempty"`
	Subtotal_canceled                    *string `xml:"subtotal_canceled,omitempty" json:"subtotal_canceled,omitempty" yaml:"subtotal_canceled,omitempty"`
	Subtotal_invoiced                    *string `xml:"subtotal_invoiced,omitempty" json:"subtotal_invoiced,omitempty" yaml:"subtotal_invoiced,omitempty"`
	Subtotal_refunded                    *string `xml:"subtotal_refunded,omitempty" json:"subtotal_refunded,omitempty" yaml:"subtotal_refunded,omitempty"`
	Tax_canceled                         *string `xml:"tax_canceled,omitempty" json:"tax_canceled,omitempty" yaml:"tax_canceled,omitempty"`
	Tax_invoiced                         *string `xml:"tax_invoiced,omitempty" json:"tax_invoiced,omitempty" yaml:"tax_invoiced,omitempty"`
	Tax_refunded                         *string `xml:"tax_refunded,omitempty" json:"tax_refunded,omitempty" yaml:"tax_refunded,omitempty"`
	Can_ship_partially                   *string `xml:"can_ship_partially,omitempty" json:"can_ship_partially,omitempty" yaml:"can_ship_partially,omitempty"`
	Can_ship_partially_item              *string `xml:"can_ship_partially_item,omitempty" json:"can_ship_partially_item,omitempty" yaml:"can_ship_partially_item,omitempty"`
	Edit_increment                       *string `xml:"edit_increment,omitempty" json:"edit_increment,omitempty" yaml:"edit_increment,omitempty"`
	Forced_do_shipment_with_invoice      *string `xml:"forced_do_shipment_with_invoice,omitempty" json:"forced_do_shipment_with_invoice,omitempty" yaml:"forced_do_shipment_with_invoice,omitempty"`
	Payment_authorization_expiration     *string `xml:"payment_authorization_expiration,omitempty" json:"payment_authorization_expiration,omitempty" yaml:"payment_authorization_expiration,omitempty"`
	Paypal_ipn_customer_notified         *string `xml:"paypal_ipn_customer_notified,omitempty" json:"paypal_ipn_customer_notified,omitempty" yaml:"paypal_ipn_customer_notified,omitempty"`
	Quote_address_id                     *string `xml:"quote_address_id,omitempty" json:"quote_address_id,omitempty" yaml:"quote_address_id,omitempty"`
	Adjustment_negative                  *string `xml:"adjustment_negative,omitempty" json:"adjustment_negative,omitempty" yaml:"adjustment_negative,omitempty"`
	Adjustment_positive                  *string `xml:"adjustment_positive,omitempty" json:"adjustment_positive,omitempty" yaml:"adjustment_positive,omitempty"`
	Base_adjustment_negative             *string `xml:"base_adjustment_negative,omitempty" json:"base_adjustment_negative,omitempty" yaml:"base_adjustment_negative,omitempty"`
	Base_adjustment_positive             *string `xml:"base_adjustment_positive,omitempty" json:"base_adjustment_positive,omitempty" yaml:"base_adjustment_positive,omitempty"`
	Base_shipping_discount_amount        *string `xml:"base_shipping_discount_amount,omitempty" json:"base_shipping_discount_amount,omitempty" yaml:"base_shipping_discount_amount,omitempty"`
	Base_subtotal_incl_tax               *string `xml:"base_subtotal_incl_tax,omitempty" json:"base_subtotal_incl_tax,omitempty" yaml:"base_subtotal_incl_tax,omitempty"`
	Base_total_due                       *string `xml:"base_total_due,omitempty" json:"base_total_due,omitempty" yaml:"base_total_due,omitempty"`
	Payment_authorization_amount         *string `xml:"payment_authorization_amount,omitempty" json:"payment_authorization_amount,omitempty" yaml:"payment_authorization_amount,omitempty"`
	Shipping_discount_amount             *string `xml:"shipping_discount_amount,omitempty" json:"shipping_discount_amount,omitempty" yaml:"shipping_discount_amount,omitempty"`
	Subtotal_incl_tax                    *string `xml:"subtotal_incl_tax,omitempty" json:"subtotal_incl_tax,omitempty" yaml:"subtotal_incl_tax,omitempty"`
	Total_due                            *string `xml:"total_due,omitempty" json:"total_due,omitempty" yaml:"total_due,omitempty"`
	Customer_dob                         *string `xml:"customer_dob,omitempty" json:"customer_dob,omitempty" yaml:"customer_dob,omitempty"`
	Customer_middlename                  *string `xml:"customer_middlename,omitempty" json:"customer_middlename,omitempty" yaml:"customer_middlename,omitempty"`
	Customer_prefix                      *string `xml:"customer_prefix,omitempty" json:"customer_prefix,omitempty" yaml:"customer_prefix,omitempty"`
	Customer_suffix                      *string `xml:"customer_suffix,omitempty" json:"customer_suffix,omitempty" yaml:"customer_suffix,omitempty"`
	Customer_taxvat                      *string `xml:"customer_taxvat,omitempty" json:"customer_taxvat,omitempty" yaml:"customer_taxvat,omitempty"`
	Discount_description                 *string `xml:"discount_description,omitempty" json:"discount_description,omitempty" yaml:"discount_description,omitempty"`
	Ext_customer_id                      *string `xml:"ext_customer_id,omitempty" json:"ext_customer_id,omitempty" yaml:"ext_customer_id,omitempty"`
	Ext_order_id                         *string `xml:"ext_order_id,omitempty" json:"ext_order_id,omitempty" yaml:"ext_order_id,omitempty"`
	Hold_before_state                    *string `xml:"hold_before_state,omitempty" json:"hold_before_state,omitempty" yaml:"hold_before_state,omitempty"`
	Hold_before_status                   *string `xml:"hold_before_status,omitempty" json:"hold_before_status,omitempty" yaml:"hold_before_status,omitempty"`
	Original_increment_id                *string `xml:"original_increment_id,omitempty" json:"original_increment_id,omitempty" yaml:"original_increment_id,omitempty"`
	Relation_child_id                    *string `xml:"relation_child_id,omitempty" json:"relation_child_id,omitempty" yaml:"relation_child_id,omitempty"`
	Relation_child_real_id               *string `xml:"relation_child_real_id,omitempty" json:"relation_child_real_id,omitempty" yaml:"relation_child_real_id,omitempty"`
	Relation_parent_id                   *string `xml:"relation_parent_id,omitempty" json:"relation_parent_id,omitempty" yaml:"relation_parent_id,omitempty"`
	Relation_parent_real_id              *string `xml:"relation_parent_real_id,omitempty" json:"relation_parent_real_id,omitempty" yaml:"relation_parent_real_id,omitempty"`
	X_forwarded_for                      *string `xml:"x_forwarded_for,omitempty" json:"x_forwarded_for,omitempty" yaml:"x_forwarded_for,omitempty"`
	Customer_note                        *string `xml:"customer_note,omitempty" json:"customer_note,omitempty" yaml:"customer_note,omitempty"`
	Total_item_count                     *string `xml:"total_item_count,omitempty" json:"total_item_count,omitempty" yaml:"total_item_count,omitempty"`
	Customer_gender                      *string `xml:"customer_gender,omitempty" json:"customer_gender,omitempty" yaml:"customer_gender,omitempty"`
	Hidden_tax_amount                    *string `xml:"hidden_tax_amount,omitempty" json:"hidden_tax_amount,omitempty" yaml:"hidden_tax_amount,omitempty"`
	Base_hidden_tax_amount               *string `xml:"base_hidden_tax_amount,omitempty" json:"base_hidden_tax_amount,omitempty" yaml:"base_hidden_tax_amount,omitempty"`
	Shipping_hidden_tax_amount           *string `xml:"shipping_hidden_tax_amount,omitempty" json:"shipping_hidden_tax_amount,omitempty" yaml:"shipping_hidden_tax_amount,omitempty"`
	Base_shipping_hidden_tax_amount      *string `xml:"base_shipping_hidden_tax_amount,omitempty" json:"base_shipping_hidden_tax_amount,omitempty" yaml:"base_shipping_hidden_tax_amount,omitempty"`
	Hidden_tax_invoiced                  *string `xml:"hidden_tax_invoiced,omitempty" json:"hidden_tax_invoiced,omitempty" yaml:"hidden_tax_invoiced,omitempty"`
	Base_hidden_tax_invoiced             *string `xml:"base_hidden_tax_invoiced,omitempty" json:"base_hidden_tax_invoiced,omitempty" yaml:"base_hidden_tax_invoiced,omitempty"`
	Hidden_tax_refunded                  *string `xml:"hidden_tax_refunded,omitempty" json:"hidden_tax_refunded,omitempty" yaml:"hidden_tax_refunded,omitempty"`
	Base_hidden_tax_refunded             *string `xml:"base_hidden_tax_refunded,omitempty" json:"base_hidden_tax_refunded,omitempty" yaml:"base_hidden_tax_refunded,omitempty"`
	Shipping_incl_tax                    *string `xml:"shipping_incl_tax,omitempty" json:"shipping_incl_tax,omitempty" yaml:"shipping_incl_tax,omitempty"`
	Base_shipping_incl_tax               *string `xml:"base_shipping_incl_tax,omitempty" json:"base_shipping_incl_tax,omitempty" yaml:"base_shipping_incl_tax,omitempty"`
	Base_customer_balance_amount         *string `xml:"base_customer_balance_amount,omitempty" json:"base_customer_balance_amount,omitempty" yaml:"base_customer_balance_amount,omitempty"`
	Customer_balance_amount              *string `xml:"customer_balance_amount,omitempty" json:"customer_balance_amount,omitempty" yaml:"customer_balance_amount,omitempty"`
	Base_customer_balance_invoiced       *string `xml:"base_customer_balance_invoiced,omitempty" json:"base_customer_balance_invoiced,omitempty" yaml:"base_customer_balance_invoiced,omitempty"`
	Customer_balance_invoiced            *string `xml:"customer_balance_invoiced,omitempty" json:"customer_balance_invoiced,omitempty" yaml:"customer_balance_invoiced,omitempty"`
	Base_customer_balance_refunded       *string `xml:"base_customer_balance_refunded,omitempty" json:"base_customer_balance_refunded,omitempty" yaml:"base_customer_balance_refunded,omitempty"`
	Customer_balance_refunded            *string `xml:"customer_balance_refunded,omitempty" json:"customer_balance_refunded,omitempty" yaml:"customer_balance_refunded,omitempty"`
	Base_customer_balance_total_refunded *string `xml:"base_customer_balance_total_refunded,omitempty" json:"base_customer_balance_total_refunded,omitempty" yaml:"base_customer_balance_total_refunded,omitempty"`
	Customer_balance_total_refunded      *string `xml:"customer_balance_total_refunded,omitempty" json:"customer_balance_total_refunded,omitempty" yaml:"customer_balance_total_refunded,omitempty"`
	Gift_cards                           *string `xml:"gift_cards,omitempty" json:"gift_cards,omitempty" yaml:"gift_cards,omitempty"`
	Base_gift_cards_amount               *string `xml:"base_gift_cards_amount,omitempty" json:"base_gift_cards_amount,omitempty" yaml:"base_gift_cards_amount,omitempty"`
	Gift_cards_amount                    *string `xml:"gift_cards_amount,omitempty" json:"gift_cards_amount,omitempty" yaml:"gift_cards_amount,omitempty"`
	Base_gift_cards_invoiced             *string `xml:"base_gift_cards_invoiced,omitempty" json:"base_gift_cards_invoiced,omitempty" yaml:"base_gift_cards_invoiced,omitempty"`
	Gift_cards_invoiced                  *string `xml:"gift_cards_invoiced,omitempty" json:"gift_cards_invoiced,omitempty" yaml:"gift_cards_invoiced,omitempty"`
	Base_gift_cards_refunded             *string `xml:"base_gift_cards_refunded,omitempty" json:"base_gift_cards_refunded,omitempty" yaml:"base_gift_cards_refunded,omitempty"`
	Gift_cards_refunded                  *string `xml:"gift_cards_refunded,omitempty" json:"gift_cards_refunded,omitempty" yaml:"gift_cards_refunded,omitempty"`
	Reward_points_balance                *string `xml:"reward_points_balance,omitempty" json:"reward_points_balance,omitempty" yaml:"reward_points_balance,omitempty"`
	Base_reward_currency_amount          *string `xml:"base_reward_currency_amount,omitempty" json:"base_reward_currency_amount,omitempty" yaml:"base_reward_currency_amount,omitempty"`
	Reward_currency_amount               *string `xml:"reward_currency_amount,omitempty" json:"reward_currency_amount,omitempty" yaml:"reward_currency_amount,omitempty"`
	Base_reward_currency_amount_invoiced *string `xml:"base_reward_currency_amount_invoiced,omitempty" json:"base_reward_currency_amount_invoiced,omitempty" yaml:"base_reward_currency_amount_invoiced,omitempty"`
	Reward_currency_amount_invoiced      *string `xml:"reward_currency_amount_invoiced,omitempty" json:"reward_currency_amount_invoiced,omitempty" yaml:"reward_currency_amount_invoiced,omitempty"`
	Base_reward_currency_amount_refunded *string `xml:"base_reward_currency_amount_refunded,omitempty" json:"base_reward_currency_amount_refunded,omitempty" yaml:"base_reward_currency_amount_refunded,omitempty"`
	Reward_currency_amount_refunded      *string `xml:"reward_currency_amount_refunded,omitempty" json:"reward_currency_amount_refunded,omitempty" yaml:"reward_currency_amount_refunded,omitempty"`
	Reward_points_balance_refunded       *string `xml:"reward_points_balance_refunded,omitempty" json:"reward_points_balance_refunded,omitempty" yaml:"reward_points_balance_refunded,omitempty"`
	Reward_points_balance_to_refund      *string `xml:"reward_points_balance_to_refund,omitempty" json:"reward_points_balance_to_refund,omitempty" yaml:"reward_points_balance_to_refund,omitempty"`
	Reward_salesrule_points              *string `xml:"reward_salesrule_points,omitempty" json:"reward_salesrule_points,omitempty" yaml:"reward_salesrule_points,omitempty"`
	Firstname                            *string `xml:"firstname,omitempty" json:"firstname,omitempty" yaml:"firstname,omitempty"`
	Lastname                             *string `xml:"lastname,omitempty" json:"lastname,omitempty" yaml:"lastname,omitempty"`
	Telephone                            *string `xml:"telephone,omitempty" json:"telephone,omitempty" yaml:"telephone,omitempty"`
	Postcode                             *string `xml:"postcode,omitempty" json:"postcode,omitempty" yaml:"postcode,omitempty"`
}

// SalesOrderListEntityArray was auto-generated from WSDL.
type SalesOrderListEntityArray struct {
	Items []*SalesOrderListEntity `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
}

// SalesOrderPaymentEntity was auto-generated from WSDL.
type SalesOrderPaymentEntity struct {
	Increment_id         *string `xml:"increment_id,omitempty" json:"increment_id,omitempty" yaml:"increment_id,omitempty"`
	Parent_id            *string `xml:"parent_id,omitempty" json:"parent_id,omitempty" yaml:"parent_id,omitempty"`
	Created_at           *string `xml:"created_at,omitempty" json:"created_at,omitempty" yaml:"created_at,omitempty"`
	Updated_at           *string `xml:"updated_at,omitempty" json:"updated_at,omitempty" yaml:"updated_at,omitempty"`
	Is_active            *string `xml:"is_active,omitempty" json:"is_active,omitempty" yaml:"is_active,omitempty"`
	Amount_ordered       *string `xml:"amount_ordered,omitempty" json:"amount_ordered,omitempty" yaml:"amount_ordered,omitempty"`
	Shipping_amount      *string `xml:"shipping_amount,omitempty" json:"shipping_amount,omitempty" yaml:"shipping_amount,omitempty"`
	Base_amount_ordered  *string `xml:"base_amount_ordered,omitempty" json:"base_amount_ordered,omitempty" yaml:"base_amount_ordered,omitempty"`
	Base_shipping_amount *string `xml:"base_shipping_amount,omitempty" json:"base_shipping_amount,omitempty" yaml:"base_shipping_amount,omitempty"`
	Method               *string `xml:"method,omitempty" json:"method,omitempty" yaml:"method,omitempty"`
	Po_number            *string `xml:"po_number,omitempty" json:"po_number,omitempty" yaml:"po_number,omitempty"`
	Cc_type              *string `xml:"cc_type,omitempty" json:"cc_type,omitempty" yaml:"cc_type,omitempty"`
	Cc_number_enc        *string `xml:"cc_number_enc,omitempty" json:"cc_number_enc,omitempty" yaml:"cc_number_enc,omitempty"`
	Cc_last4             *string `xml:"cc_last4,omitempty" json:"cc_last4,omitempty" yaml:"cc_last4,omitempty"`
	Cc_owner             *string `xml:"cc_owner,omitempty" json:"cc_owner,omitempty" yaml:"cc_owner,omitempty"`
	Cc_exp_month         *string `xml:"cc_exp_month,omitempty" json:"cc_exp_month,omitempty" yaml:"cc_exp_month,omitempty"`
	Cc_exp_year          *string `xml:"cc_exp_year,omitempty" json:"cc_exp_year,omitempty" yaml:"cc_exp_year,omitempty"`
	Cc_ss_start_month    *string `xml:"cc_ss_start_month,omitempty" json:"cc_ss_start_month,omitempty" yaml:"cc_ss_start_month,omitempty"`
	Cc_ss_start_year     *string `xml:"cc_ss_start_year,omitempty" json:"cc_ss_start_year,omitempty" yaml:"cc_ss_start_year,omitempty"`
	Payment_id           *string `xml:"payment_id,omitempty" json:"payment_id,omitempty" yaml:"payment_id,omitempty"`
}

// SalesOrderShipmentCommentEntity was auto-generated from WSDL.
type SalesOrderShipmentCommentEntity struct {
	Increment_id         *string `xml:"increment_id,omitempty" json:"increment_id,omitempty" yaml:"increment_id,omitempty"`
	Parent_id            *string `xml:"parent_id,omitempty" json:"parent_id,omitempty" yaml:"parent_id,omitempty"`
	Created_at           *string `xml:"created_at,omitempty" json:"created_at,omitempty" yaml:"created_at,omitempty"`
	Updated_at           *string `xml:"updated_at,omitempty" json:"updated_at,omitempty" yaml:"updated_at,omitempty"`
	Is_active            *string `xml:"is_active,omitempty" json:"is_active,omitempty" yaml:"is_active,omitempty"`
	Comment              *string `xml:"comment,omitempty" json:"comment,omitempty" yaml:"comment,omitempty"`
	Is_customer_notified *string `xml:"is_customer_notified,omitempty" json:"is_customer_notified,omitempty" yaml:"is_customer_notified,omitempty"`
	Comment_id           *string `xml:"comment_id,omitempty" json:"comment_id,omitempty" yaml:"comment_id,omitempty"`
}

// SalesOrderShipmentCommentEntityArray was auto-generated from
// WSDL.
type SalesOrderShipmentCommentEntityArray struct {
	Items []*SalesOrderShipmentCommentEntity `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
}

// SalesOrderShipmentEntity was auto-generated from WSDL.
type SalesOrderShipmentEntity struct {
	Increment_id        *string                               `xml:"increment_id,omitempty" json:"increment_id,omitempty" yaml:"increment_id,omitempty"`
	Parent_id           *string                               `xml:"parent_id,omitempty" json:"parent_id,omitempty" yaml:"parent_id,omitempty"`
	Store_id            *string                               `xml:"store_id,omitempty" json:"store_id,omitempty" yaml:"store_id,omitempty"`
	Created_at          *string                               `xml:"created_at,omitempty" json:"created_at,omitempty" yaml:"created_at,omitempty"`
	Updated_at          *string                               `xml:"updated_at,omitempty" json:"updated_at,omitempty" yaml:"updated_at,omitempty"`
	Is_active           *string                               `xml:"is_active,omitempty" json:"is_active,omitempty" yaml:"is_active,omitempty"`
	Shipping_address_id *string                               `xml:"shipping_address_id,omitempty" json:"shipping_address_id,omitempty" yaml:"shipping_address_id,omitempty"`
	Shipping_firstname  *string                               `xml:"shipping_firstname,omitempty" json:"shipping_firstname,omitempty" yaml:"shipping_firstname,omitempty"`
	Shipping_lastname   *string                               `xml:"shipping_lastname,omitempty" json:"shipping_lastname,omitempty" yaml:"shipping_lastname,omitempty"`
	Order_id            *string                               `xml:"order_id,omitempty" json:"order_id,omitempty" yaml:"order_id,omitempty"`
	Order_increment_id  *string                               `xml:"order_increment_id,omitempty" json:"order_increment_id,omitempty" yaml:"order_increment_id,omitempty"`
	Order_created_at    *string                               `xml:"order_created_at,omitempty" json:"order_created_at,omitempty" yaml:"order_created_at,omitempty"`
	Total_qty           *string                               `xml:"total_qty,omitempty" json:"total_qty,omitempty" yaml:"total_qty,omitempty"`
	Shipment_id         *string                               `xml:"shipment_id,omitempty" json:"shipment_id,omitempty" yaml:"shipment_id,omitempty"`
	Items               *SalesOrderShipmentItemEntityArray    `xml:"items,omitempty" json:"items,omitempty" yaml:"items,omitempty"`
	Tracks              *SalesOrderShipmentTrackEntityArray   `xml:"tracks,omitempty" json:"tracks,omitempty" yaml:"tracks,omitempty"`
	Comments            *SalesOrderShipmentCommentEntityArray `xml:"comments,omitempty" json:"comments,omitempty" yaml:"comments,omitempty"`
}

// SalesOrderShipmentEntityArray was auto-generated from WSDL.
type SalesOrderShipmentEntityArray struct {
	Items []*SalesOrderShipmentEntity `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
}

// SalesOrderShipmentItemEntity was auto-generated from WSDL.
type SalesOrderShipmentItemEntity struct {
	Increment_id  *string `xml:"increment_id,omitempty" json:"increment_id,omitempty" yaml:"increment_id,omitempty"`
	Parent_id     *string `xml:"parent_id,omitempty" json:"parent_id,omitempty" yaml:"parent_id,omitempty"`
	Created_at    *string `xml:"created_at,omitempty" json:"created_at,omitempty" yaml:"created_at,omitempty"`
	Updated_at    *string `xml:"updated_at,omitempty" json:"updated_at,omitempty" yaml:"updated_at,omitempty"`
	Is_active     *string `xml:"is_active,omitempty" json:"is_active,omitempty" yaml:"is_active,omitempty"`
	Sku           *string `xml:"sku,omitempty" json:"sku,omitempty" yaml:"sku,omitempty"`
	Name          *string `xml:"name,omitempty" json:"name,omitempty" yaml:"name,omitempty"`
	Order_item_id *string `xml:"order_item_id,omitempty" json:"order_item_id,omitempty" yaml:"order_item_id,omitempty"`
	Product_id    *string `xml:"product_id,omitempty" json:"product_id,omitempty" yaml:"product_id,omitempty"`
	Weight        *string `xml:"weight,omitempty" json:"weight,omitempty" yaml:"weight,omitempty"`
	Price         *string `xml:"price,omitempty" json:"price,omitempty" yaml:"price,omitempty"`
	Qty           *string `xml:"qty,omitempty" json:"qty,omitempty" yaml:"qty,omitempty"`
	Item_id       *string `xml:"item_id,omitempty" json:"item_id,omitempty" yaml:"item_id,omitempty"`
}

// SalesOrderShipmentItemEntityArray was auto-generated from WSDL.
type SalesOrderShipmentItemEntityArray struct {
	Items []*SalesOrderShipmentItemEntity `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
}

// SalesOrderShipmentTrackEntity was auto-generated from WSDL.
type SalesOrderShipmentTrackEntity struct {
	Increment_id *string `xml:"increment_id,omitempty" json:"increment_id,omitempty" yaml:"increment_id,omitempty"`
	Parent_id    *string `xml:"parent_id,omitempty" json:"parent_id,omitempty" yaml:"parent_id,omitempty"`
	Created_at   *string `xml:"created_at,omitempty" json:"created_at,omitempty" yaml:"created_at,omitempty"`
	Updated_at   *string `xml:"updated_at,omitempty" json:"updated_at,omitempty" yaml:"updated_at,omitempty"`
	Is_active    *string `xml:"is_active,omitempty" json:"is_active,omitempty" yaml:"is_active,omitempty"`
	Carrier_code *string `xml:"carrier_code,omitempty" json:"carrier_code,omitempty" yaml:"carrier_code,omitempty"`
	Title        *string `xml:"title,omitempty" json:"title,omitempty" yaml:"title,omitempty"`
	Number       *string `xml:"number,omitempty" json:"number,omitempty" yaml:"number,omitempty"`
	Order_id     *string `xml:"order_id,omitempty" json:"order_id,omitempty" yaml:"order_id,omitempty"`
	Track_id     *string `xml:"track_id,omitempty" json:"track_id,omitempty" yaml:"track_id,omitempty"`
}

// SalesOrderShipmentTrackEntityArray was auto-generated from WSDL.
type SalesOrderShipmentTrackEntityArray struct {
	Items []*SalesOrderShipmentTrackEntity `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
}

// SalesOrderStatusHistoryEntity was auto-generated from WSDL.
type SalesOrderStatusHistoryEntity struct {
	Increment_id         *string `xml:"increment_id,omitempty" json:"increment_id,omitempty" yaml:"increment_id,omitempty"`
	Parent_id            *string `xml:"parent_id,omitempty" json:"parent_id,omitempty" yaml:"parent_id,omitempty"`
	Created_at           *string `xml:"created_at,omitempty" json:"created_at,omitempty" yaml:"created_at,omitempty"`
	Updated_at           *string `xml:"updated_at,omitempty" json:"updated_at,omitempty" yaml:"updated_at,omitempty"`
	Is_active            *string `xml:"is_active,omitempty" json:"is_active,omitempty" yaml:"is_active,omitempty"`
	Is_customer_notified *string `xml:"is_customer_notified,omitempty" json:"is_customer_notified,omitempty" yaml:"is_customer_notified,omitempty"`
	Status               *string `xml:"status,omitempty" json:"status,omitempty" yaml:"status,omitempty"`
	Comment              *string `xml:"comment,omitempty" json:"comment,omitempty" yaml:"comment,omitempty"`
}

// SalesOrderStatusHistoryEntityArray was auto-generated from WSDL.
type SalesOrderStatusHistoryEntityArray struct {
	Items []*SalesOrderStatusHistoryEntity `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
}

// ShoppingCartAdcurveProductEntity was auto-generated from WSDL.
type ShoppingCartAdcurveProductEntity struct {
	Product_id        *string           `xml:"product_id,omitempty" json:"product_id,omitempty" yaml:"product_id,omitempty"`
	Sku               *string           `xml:"sku,omitempty" json:"sku,omitempty" yaml:"sku,omitempty"`
	Qty               *float64          `xml:"qty,omitempty" json:"qty,omitempty" yaml:"qty,omitempty"`
	Price             *float64          `xml:"price,omitempty" json:"price,omitempty" yaml:"price,omitempty"`
	Options           *AssociativeArray `xml:"options,omitempty" json:"options,omitempty" yaml:"options,omitempty"`
	Bundle_option     *AssociativeArray `xml:"bundle_option,omitempty" json:"bundle_option,omitempty" yaml:"bundle_option,omitempty"`
	Bundle_option_qty *AssociativeArray `xml:"bundle_option_qty,omitempty" json:"bundle_option_qty,omitempty" yaml:"bundle_option_qty,omitempty"`
	Links             *ArrayOfString    `xml:"links,omitempty" json:"links,omitempty" yaml:"links,omitempty"`
}

// ShoppingCartAdcurveProductEntityArray was auto-generated from
// WSDL.
type ShoppingCartAdcurveProductEntityArray struct {
	Items []*ShoppingCartAdcurveProductEntity `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
}

// ShoppingCartAddressEntity was auto-generated from WSDL.
type ShoppingCartAddressEntity struct {
	Address_id           *string  `xml:"address_id,omitempty" json:"address_id,omitempty" yaml:"address_id,omitempty"`
	Created_at           *string  `xml:"created_at,omitempty" json:"created_at,omitempty" yaml:"created_at,omitempty"`
	Updated_at           *string  `xml:"updated_at,omitempty" json:"updated_at,omitempty" yaml:"updated_at,omitempty"`
	Customer_id          *string  `xml:"customer_id,omitempty" json:"customer_id,omitempty" yaml:"customer_id,omitempty"`
	Save_in_address_book *int     `xml:"save_in_address_book,omitempty" json:"save_in_address_book,omitempty" yaml:"save_in_address_book,omitempty"`
	Customer_address_id  *string  `xml:"customer_address_id,omitempty" json:"customer_address_id,omitempty" yaml:"customer_address_id,omitempty"`
	Address_type         *string  `xml:"address_type,omitempty" json:"address_type,omitempty" yaml:"address_type,omitempty"`
	Email                *string  `xml:"email,omitempty" json:"email,omitempty" yaml:"email,omitempty"`
	Prefix               *string  `xml:"prefix,omitempty" json:"prefix,omitempty" yaml:"prefix,omitempty"`
	Firstname            *string  `xml:"firstname,omitempty" json:"firstname,omitempty" yaml:"firstname,omitempty"`
	Middlename           *string  `xml:"middlename,omitempty" json:"middlename,omitempty" yaml:"middlename,omitempty"`
	Lastname             *string  `xml:"lastname,omitempty" json:"lastname,omitempty" yaml:"lastname,omitempty"`
	Suffix               *string  `xml:"suffix,omitempty" json:"suffix,omitempty" yaml:"suffix,omitempty"`
	Company              *string  `xml:"company,omitempty" json:"company,omitempty" yaml:"company,omitempty"`
	Street               *string  `xml:"street,omitempty" json:"street,omitempty" yaml:"street,omitempty"`
	City                 *string  `xml:"city,omitempty" json:"city,omitempty" yaml:"city,omitempty"`
	Region               *string  `xml:"region,omitempty" json:"region,omitempty" yaml:"region,omitempty"`
	Region_id            *string  `xml:"region_id,omitempty" json:"region_id,omitempty" yaml:"region_id,omitempty"`
	Postcode             *string  `xml:"postcode,omitempty" json:"postcode,omitempty" yaml:"postcode,omitempty"`
	Country_id           *string  `xml:"country_id,omitempty" json:"country_id,omitempty" yaml:"country_id,omitempty"`
	Telephone            *string  `xml:"telephone,omitempty" json:"telephone,omitempty" yaml:"telephone,omitempty"`
	Fax                  *string  `xml:"fax,omitempty" json:"fax,omitempty" yaml:"fax,omitempty"`
	Same_as_billing      *int     `xml:"same_as_billing,omitempty" json:"same_as_billing,omitempty" yaml:"same_as_billing,omitempty"`
	Free_shipping        *int     `xml:"free_shipping,omitempty" json:"free_shipping,omitempty" yaml:"free_shipping,omitempty"`
	Shipping_method      *string  `xml:"shipping_method,omitempty" json:"shipping_method,omitempty" yaml:"shipping_method,omitempty"`
	Shipping_description *string  `xml:"shipping_description,omitempty" json:"shipping_description,omitempty" yaml:"shipping_description,omitempty"`
	Weight               *float64 `xml:"weight,omitempty" json:"weight,omitempty" yaml:"weight,omitempty"`
}

// ShoppingCartCustomerAddressEntity was auto-generated from WSDL.
type ShoppingCartCustomerAddressEntity struct {
	Mode                *string `xml:"mode,omitempty" json:"mode,omitempty" yaml:"mode,omitempty"`
	Address_id          *string `xml:"address_id,omitempty" json:"address_id,omitempty" yaml:"address_id,omitempty"`
	Firstname           *string `xml:"firstname,omitempty" json:"firstname,omitempty" yaml:"firstname,omitempty"`
	Lastname            *string `xml:"lastname,omitempty" json:"lastname,omitempty" yaml:"lastname,omitempty"`
	Company             *string `xml:"company,omitempty" json:"company,omitempty" yaml:"company,omitempty"`
	Street              *string `xml:"street,omitempty" json:"street,omitempty" yaml:"street,omitempty"`
	City                *string `xml:"city,omitempty" json:"city,omitempty" yaml:"city,omitempty"`
	Region              *string `xml:"region,omitempty" json:"region,omitempty" yaml:"region,omitempty"`
	Region_id           *string `xml:"region_id,omitempty" json:"region_id,omitempty" yaml:"region_id,omitempty"`
	Postcode            *string `xml:"postcode,omitempty" json:"postcode,omitempty" yaml:"postcode,omitempty"`
	Country_id          *string `xml:"country_id,omitempty" json:"country_id,omitempty" yaml:"country_id,omitempty"`
	Telephone           *string `xml:"telephone,omitempty" json:"telephone,omitempty" yaml:"telephone,omitempty"`
	Fax                 *string `xml:"fax,omitempty" json:"fax,omitempty" yaml:"fax,omitempty"`
	Is_default_billing  *int    `xml:"is_default_billing,omitempty" json:"is_default_billing,omitempty" yaml:"is_default_billing,omitempty"`
	Is_default_shipping *int    `xml:"is_default_shipping,omitempty" json:"is_default_shipping,omitempty" yaml:"is_default_shipping,omitempty"`
}

// ShoppingCartCustomerAddressEntityArray was auto-generated from
// WSDL.
type ShoppingCartCustomerAddressEntityArray struct {
	Items []*ShoppingCartCustomerAddressEntity `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
}

// ShoppingCartCustomerEntity was auto-generated from WSDL.
type ShoppingCartCustomerEntity struct {
	Mode         *string `xml:"mode,omitempty" json:"mode,omitempty" yaml:"mode,omitempty"`
	Customer_id  *int    `xml:"customer_id,omitempty" json:"customer_id,omitempty" yaml:"customer_id,omitempty"`
	Email        *string `xml:"email,omitempty" json:"email,omitempty" yaml:"email,omitempty"`
	Firstname    *string `xml:"firstname,omitempty" json:"firstname,omitempty" yaml:"firstname,omitempty"`
	Lastname     *string `xml:"lastname,omitempty" json:"lastname,omitempty" yaml:"lastname,omitempty"`
	Password     *string `xml:"password,omitempty" json:"password,omitempty" yaml:"password,omitempty"`
	Confirmation *string `xml:"confirmation,omitempty" json:"confirmation,omitempty" yaml:"confirmation,omitempty"`
	Website_id   *int    `xml:"website_id,omitempty" json:"website_id,omitempty" yaml:"website_id,omitempty"`
	Store_id     *int    `xml:"store_id,omitempty" json:"store_id,omitempty" yaml:"store_id,omitempty"`
	Group_id     *int    `xml:"group_id,omitempty" json:"group_id,omitempty" yaml:"group_id,omitempty"`
}

// ShoppingCartInfoEntity was auto-generated from WSDL.
type ShoppingCartInfoEntity struct {
	Store_id                          *string                      `xml:"store_id,omitempty" json:"store_id,omitempty" yaml:"store_id,omitempty"`
	Created_at                        *string                      `xml:"created_at,omitempty" json:"created_at,omitempty" yaml:"created_at,omitempty"`
	Updated_at                        *string                      `xml:"updated_at,omitempty" json:"updated_at,omitempty" yaml:"updated_at,omitempty"`
	Converted_at                      *string                      `xml:"converted_at,omitempty" json:"converted_at,omitempty" yaml:"converted_at,omitempty"`
	Quote_id                          *int                         `xml:"quote_id,omitempty" json:"quote_id,omitempty" yaml:"quote_id,omitempty"`
	Is_active                         *int                         `xml:"is_active,omitempty" json:"is_active,omitempty" yaml:"is_active,omitempty"`
	Is_virtual                        *int                         `xml:"is_virtual,omitempty" json:"is_virtual,omitempty" yaml:"is_virtual,omitempty"`
	Is_multi_shipping                 *int                         `xml:"is_multi_shipping,omitempty" json:"is_multi_shipping,omitempty" yaml:"is_multi_shipping,omitempty"`
	Items_count                       *float64                     `xml:"items_count,omitempty" json:"items_count,omitempty" yaml:"items_count,omitempty"`
	Items_qty                         *float64                     `xml:"items_qty,omitempty" json:"items_qty,omitempty" yaml:"items_qty,omitempty"`
	Orig_order_id                     *string                      `xml:"orig_order_id,omitempty" json:"orig_order_id,omitempty" yaml:"orig_order_id,omitempty"`
	Store_to_base_rate                *string                      `xml:"store_to_base_rate,omitempty" json:"store_to_base_rate,omitempty" yaml:"store_to_base_rate,omitempty"`
	Store_to_quote_rate               *string                      `xml:"store_to_quote_rate,omitempty" json:"store_to_quote_rate,omitempty" yaml:"store_to_quote_rate,omitempty"`
	Base_currency_code                *string                      `xml:"base_currency_code,omitempty" json:"base_currency_code,omitempty" yaml:"base_currency_code,omitempty"`
	Store_currency_code               *string                      `xml:"store_currency_code,omitempty" json:"store_currency_code,omitempty" yaml:"store_currency_code,omitempty"`
	Quote_currency_code               *string                      `xml:"quote_currency_code,omitempty" json:"quote_currency_code,omitempty" yaml:"quote_currency_code,omitempty"`
	Grand_total                       *string                      `xml:"grand_total,omitempty" json:"grand_total,omitempty" yaml:"grand_total,omitempty"`
	Base_grand_total                  *string                      `xml:"base_grand_total,omitempty" json:"base_grand_total,omitempty" yaml:"base_grand_total,omitempty"`
	Checkout_method                   *string                      `xml:"checkout_method,omitempty" json:"checkout_method,omitempty" yaml:"checkout_method,omitempty"`
	Customer_id                       *string                      `xml:"customer_id,omitempty" json:"customer_id,omitempty" yaml:"customer_id,omitempty"`
	Customer_tax_class_id             *string                      `xml:"customer_tax_class_id,omitempty" json:"customer_tax_class_id,omitempty" yaml:"customer_tax_class_id,omitempty"`
	Customer_group_id                 *int                         `xml:"customer_group_id,omitempty" json:"customer_group_id,omitempty" yaml:"customer_group_id,omitempty"`
	Customer_email                    *string                      `xml:"customer_email,omitempty" json:"customer_email,omitempty" yaml:"customer_email,omitempty"`
	Customer_prefix                   *string                      `xml:"customer_prefix,omitempty" json:"customer_prefix,omitempty" yaml:"customer_prefix,omitempty"`
	Customer_firstname                *string                      `xml:"customer_firstname,omitempty" json:"customer_firstname,omitempty" yaml:"customer_firstname,omitempty"`
	Customer_middlename               *string                      `xml:"customer_middlename,omitempty" json:"customer_middlename,omitempty" yaml:"customer_middlename,omitempty"`
	Customer_lastname                 *string                      `xml:"customer_lastname,omitempty" json:"customer_lastname,omitempty" yaml:"customer_lastname,omitempty"`
	Customer_suffix                   *string                      `xml:"customer_suffix,omitempty" json:"customer_suffix,omitempty" yaml:"customer_suffix,omitempty"`
	Customer_note                     *string                      `xml:"customer_note,omitempty" json:"customer_note,omitempty" yaml:"customer_note,omitempty"`
	Customer_note_notify              *string                      `xml:"customer_note_notify,omitempty" json:"customer_note_notify,omitempty" yaml:"customer_note_notify,omitempty"`
	Customer_is_guest                 *string                      `xml:"customer_is_guest,omitempty" json:"customer_is_guest,omitempty" yaml:"customer_is_guest,omitempty"`
	Applied_rule_ids                  *string                      `xml:"applied_rule_ids,omitempty" json:"applied_rule_ids,omitempty" yaml:"applied_rule_ids,omitempty"`
	Reserved_order_id                 *string                      `xml:"reserved_order_id,omitempty" json:"reserved_order_id,omitempty" yaml:"reserved_order_id,omitempty"`
	Password_hash                     *string                      `xml:"password_hash,omitempty" json:"password_hash,omitempty" yaml:"password_hash,omitempty"`
	Coupon_code                       *string                      `xml:"coupon_code,omitempty" json:"coupon_code,omitempty" yaml:"coupon_code,omitempty"`
	Global_currency_code              *string                      `xml:"global_currency_code,omitempty" json:"global_currency_code,omitempty" yaml:"global_currency_code,omitempty"`
	Base_to_global_rate               *float64                     `xml:"base_to_global_rate,omitempty" json:"base_to_global_rate,omitempty" yaml:"base_to_global_rate,omitempty"`
	Base_to_quote_rate                *float64                     `xml:"base_to_quote_rate,omitempty" json:"base_to_quote_rate,omitempty" yaml:"base_to_quote_rate,omitempty"`
	Customer_taxvat                   *string                      `xml:"customer_taxvat,omitempty" json:"customer_taxvat,omitempty" yaml:"customer_taxvat,omitempty"`
	Customer_gender                   *string                      `xml:"customer_gender,omitempty" json:"customer_gender,omitempty" yaml:"customer_gender,omitempty"`
	Subtotal                          *float64                     `xml:"subtotal,omitempty" json:"subtotal,omitempty" yaml:"subtotal,omitempty"`
	Base_subtotal                     *float64                     `xml:"base_subtotal,omitempty" json:"base_subtotal,omitempty" yaml:"base_subtotal,omitempty"`
	Subtotal_with_discount            *float64                     `xml:"subtotal_with_discount,omitempty" json:"subtotal_with_discount,omitempty" yaml:"subtotal_with_discount,omitempty"`
	Base_subtotal_with_discount       *float64                     `xml:"base_subtotal_with_discount,omitempty" json:"base_subtotal_with_discount,omitempty" yaml:"base_subtotal_with_discount,omitempty"`
	Ext_shipping_info                 *string                      `xml:"ext_shipping_info,omitempty" json:"ext_shipping_info,omitempty" yaml:"ext_shipping_info,omitempty"`
	Gift_message_id                   *string                      `xml:"gift_message_id,omitempty" json:"gift_message_id,omitempty" yaml:"gift_message_id,omitempty"`
	Gift_message                      *string                      `xml:"gift_message,omitempty" json:"gift_message,omitempty" yaml:"gift_message,omitempty"`
	Customer_balance_amount_used      *float64                     `xml:"customer_balance_amount_used,omitempty" json:"customer_balance_amount_used,omitempty" yaml:"customer_balance_amount_used,omitempty"`
	Base_customer_balance_amount_used *float64                     `xml:"base_customer_balance_amount_used,omitempty" json:"base_customer_balance_amount_used,omitempty" yaml:"base_customer_balance_amount_used,omitempty"`
	Use_customer_balance              *string                      `xml:"use_customer_balance,omitempty" json:"use_customer_balance,omitempty" yaml:"use_customer_balance,omitempty"`
	Gift_cards_amount                 *string                      `xml:"gift_cards_amount,omitempty" json:"gift_cards_amount,omitempty" yaml:"gift_cards_amount,omitempty"`
	Base_gift_cards_amount            *string                      `xml:"base_gift_cards_amount,omitempty" json:"base_gift_cards_amount,omitempty" yaml:"base_gift_cards_amount,omitempty"`
	Gift_cards_amount_used            *string                      `xml:"gift_cards_amount_used,omitempty" json:"gift_cards_amount_used,omitempty" yaml:"gift_cards_amount_used,omitempty"`
	Use_reward_points                 *string                      `xml:"use_reward_points,omitempty" json:"use_reward_points,omitempty" yaml:"use_reward_points,omitempty"`
	Reward_points_balance             *string                      `xml:"reward_points_balance,omitempty" json:"reward_points_balance,omitempty" yaml:"reward_points_balance,omitempty"`
	Base_reward_currency_amount       *string                      `xml:"base_reward_currency_amount,omitempty" json:"base_reward_currency_amount,omitempty" yaml:"base_reward_currency_amount,omitempty"`
	Reward_currency_amount            *string                      `xml:"reward_currency_amount,omitempty" json:"reward_currency_amount,omitempty" yaml:"reward_currency_amount,omitempty"`
	Shipping_address                  *ShoppingCartAddressEntity   `xml:"shipping_address,omitempty" json:"shipping_address,omitempty" yaml:"shipping_address,omitempty"`
	Billing_address                   *ShoppingCartAddressEntity   `xml:"billing_address,omitempty" json:"billing_address,omitempty" yaml:"billing_address,omitempty"`
	Items                             *ShoppingCartItemEntityArray `xml:"items,omitempty" json:"items,omitempty" yaml:"items,omitempty"`
	Payment                           *ShoppingCartPaymentEntity   `xml:"payment,omitempty" json:"payment,omitempty" yaml:"payment,omitempty"`
}

// ShoppingCartItemEntity was auto-generated from WSDL.
type ShoppingCartItemEntity struct {
	Item_id                          *string  `xml:"item_id,omitempty" json:"item_id,omitempty" yaml:"item_id,omitempty"`
	Created_at                       *string  `xml:"created_at,omitempty" json:"created_at,omitempty" yaml:"created_at,omitempty"`
	Updated_at                       *string  `xml:"updated_at,omitempty" json:"updated_at,omitempty" yaml:"updated_at,omitempty"`
	Product_id                       *string  `xml:"product_id,omitempty" json:"product_id,omitempty" yaml:"product_id,omitempty"`
	Store_id                         *string  `xml:"store_id,omitempty" json:"store_id,omitempty" yaml:"store_id,omitempty"`
	Parent_item_id                   *string  `xml:"parent_item_id,omitempty" json:"parent_item_id,omitempty" yaml:"parent_item_id,omitempty"`
	Is_virtual                       *int     `xml:"is_virtual,omitempty" json:"is_virtual,omitempty" yaml:"is_virtual,omitempty"`
	Sku                              *string  `xml:"sku,omitempty" json:"sku,omitempty" yaml:"sku,omitempty"`
	Name                             *string  `xml:"name,omitempty" json:"name,omitempty" yaml:"name,omitempty"`
	Description                      *string  `xml:"description,omitempty" json:"description,omitempty" yaml:"description,omitempty"`
	Applied_rule_ids                 *string  `xml:"applied_rule_ids,omitempty" json:"applied_rule_ids,omitempty" yaml:"applied_rule_ids,omitempty"`
	Additional_data                  *string  `xml:"additional_data,omitempty" json:"additional_data,omitempty" yaml:"additional_data,omitempty"`
	Free_shipping                    *string  `xml:"free_shipping,omitempty" json:"free_shipping,omitempty" yaml:"free_shipping,omitempty"`
	Is_qty_decimal                   *string  `xml:"is_qty_decimal,omitempty" json:"is_qty_decimal,omitempty" yaml:"is_qty_decimal,omitempty"`
	No_discount                      *string  `xml:"no_discount,omitempty" json:"no_discount,omitempty" yaml:"no_discount,omitempty"`
	Weight                           *float64 `xml:"weight,omitempty" json:"weight,omitempty" yaml:"weight,omitempty"`
	Qty                              *float64 `xml:"qty,omitempty" json:"qty,omitempty" yaml:"qty,omitempty"`
	Price                            *float64 `xml:"price,omitempty" json:"price,omitempty" yaml:"price,omitempty"`
	Base_price                       *float64 `xml:"base_price,omitempty" json:"base_price,omitempty" yaml:"base_price,omitempty"`
	Custom_price                     *float64 `xml:"custom_price,omitempty" json:"custom_price,omitempty" yaml:"custom_price,omitempty"`
	Discount_percent                 *float64 `xml:"discount_percent,omitempty" json:"discount_percent,omitempty" yaml:"discount_percent,omitempty"`
	Discount_amount                  *float64 `xml:"discount_amount,omitempty" json:"discount_amount,omitempty" yaml:"discount_amount,omitempty"`
	Base_discount_amount             *float64 `xml:"base_discount_amount,omitempty" json:"base_discount_amount,omitempty" yaml:"base_discount_amount,omitempty"`
	Tax_percent                      *float64 `xml:"tax_percent,omitempty" json:"tax_percent,omitempty" yaml:"tax_percent,omitempty"`
	Tax_amount                       *float64 `xml:"tax_amount,omitempty" json:"tax_amount,omitempty" yaml:"tax_amount,omitempty"`
	Base_tax_amount                  *float64 `xml:"base_tax_amount,omitempty" json:"base_tax_amount,omitempty" yaml:"base_tax_amount,omitempty"`
	Row_total                        *float64 `xml:"row_total,omitempty" json:"row_total,omitempty" yaml:"row_total,omitempty"`
	Base_row_total                   *float64 `xml:"base_row_total,omitempty" json:"base_row_total,omitempty" yaml:"base_row_total,omitempty"`
	Row_total_with_discount          *float64 `xml:"row_total_with_discount,omitempty" json:"row_total_with_discount,omitempty" yaml:"row_total_with_discount,omitempty"`
	Row_weight                       *float64 `xml:"row_weight,omitempty" json:"row_weight,omitempty" yaml:"row_weight,omitempty"`
	Product_type                     *string  `xml:"product_type,omitempty" json:"product_type,omitempty" yaml:"product_type,omitempty"`
	Base_tax_before_discount         *float64 `xml:"base_tax_before_discount,omitempty" json:"base_tax_before_discount,omitempty" yaml:"base_tax_before_discount,omitempty"`
	Tax_before_discount              *float64 `xml:"tax_before_discount,omitempty" json:"tax_before_discount,omitempty" yaml:"tax_before_discount,omitempty"`
	Original_custom_price            *float64 `xml:"original_custom_price,omitempty" json:"original_custom_price,omitempty" yaml:"original_custom_price,omitempty"`
	Base_cost                        *float64 `xml:"base_cost,omitempty" json:"base_cost,omitempty" yaml:"base_cost,omitempty"`
	Price_incl_tax                   *float64 `xml:"price_incl_tax,omitempty" json:"price_incl_tax,omitempty" yaml:"price_incl_tax,omitempty"`
	Base_price_incl_tax              *float64 `xml:"base_price_incl_tax,omitempty" json:"base_price_incl_tax,omitempty" yaml:"base_price_incl_tax,omitempty"`
	Row_total_incl_tax               *float64 `xml:"row_total_incl_tax,omitempty" json:"row_total_incl_tax,omitempty" yaml:"row_total_incl_tax,omitempty"`
	Base_row_total_incl_tax          *float64 `xml:"base_row_total_incl_tax,omitempty" json:"base_row_total_incl_tax,omitempty" yaml:"base_row_total_incl_tax,omitempty"`
	Gift_message_id                  *string  `xml:"gift_message_id,omitempty" json:"gift_message_id,omitempty" yaml:"gift_message_id,omitempty"`
	Gift_message                     *string  `xml:"gift_message,omitempty" json:"gift_message,omitempty" yaml:"gift_message,omitempty"`
	Gift_message_available           *string  `xml:"gift_message_available,omitempty" json:"gift_message_available,omitempty" yaml:"gift_message_available,omitempty"`
	Weee_tax_applied                 *float64 `xml:"weee_tax_applied,omitempty" json:"weee_tax_applied,omitempty" yaml:"weee_tax_applied,omitempty"`
	Weee_tax_applied_amount          *float64 `xml:"weee_tax_applied_amount,omitempty" json:"weee_tax_applied_amount,omitempty" yaml:"weee_tax_applied_amount,omitempty"`
	Weee_tax_applied_row_amount      *float64 `xml:"weee_tax_applied_row_amount,omitempty" json:"weee_tax_applied_row_amount,omitempty" yaml:"weee_tax_applied_row_amount,omitempty"`
	Base_weee_tax_applied_amount     *float64 `xml:"base_weee_tax_applied_amount,omitempty" json:"base_weee_tax_applied_amount,omitempty" yaml:"base_weee_tax_applied_amount,omitempty"`
	Base_weee_tax_applied_row_amount *float64 `xml:"base_weee_tax_applied_row_amount,omitempty" json:"base_weee_tax_applied_row_amount,omitempty" yaml:"base_weee_tax_applied_row_amount,omitempty"`
	Weee_tax_disposition             *float64 `xml:"weee_tax_disposition,omitempty" json:"weee_tax_disposition,omitempty" yaml:"weee_tax_disposition,omitempty"`
	Weee_tax_row_disposition         *float64 `xml:"weee_tax_row_disposition,omitempty" json:"weee_tax_row_disposition,omitempty" yaml:"weee_tax_row_disposition,omitempty"`
	Base_weee_tax_disposition        *float64 `xml:"base_weee_tax_disposition,omitempty" json:"base_weee_tax_disposition,omitempty" yaml:"base_weee_tax_disposition,omitempty"`
	Base_weee_tax_row_disposition    *float64 `xml:"base_weee_tax_row_disposition,omitempty" json:"base_weee_tax_row_disposition,omitempty" yaml:"base_weee_tax_row_disposition,omitempty"`
	Tax_class_id                     *string  `xml:"tax_class_id,omitempty" json:"tax_class_id,omitempty" yaml:"tax_class_id,omitempty"`
}

// ShoppingCartItemEntityArray was auto-generated from WSDL.
type ShoppingCartItemEntityArray struct {
	Items []*ShoppingCartItemEntity `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
}

// ShoppingCartLicenseEntity was auto-generated from WSDL.
type ShoppingCartLicenseEntity struct {
	Agreement_id *string `xml:"agreement_id,omitempty" json:"agreement_id,omitempty" yaml:"agreement_id,omitempty"`
	Name         *string `xml:"name,omitempty" json:"name,omitempty" yaml:"name,omitempty"`
	Content      *string `xml:"content,omitempty" json:"content,omitempty" yaml:"content,omitempty"`
	Is_active    *int    `xml:"is_active,omitempty" json:"is_active,omitempty" yaml:"is_active,omitempty"`
	Is_html      *int    `xml:"is_html,omitempty" json:"is_html,omitempty" yaml:"is_html,omitempty"`
}

// ShoppingCartLicenseEntityArray was auto-generated from WSDL.
type ShoppingCartLicenseEntityArray struct {
	Items []*ShoppingCartLicenseEntity `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
}

// ShoppingCartOrderAdditionalAttributesEntity was auto-generated
// from WSDL.
type ShoppingCartOrderAdditionalAttributesEntity struct {
	AdcurveOrderSource   *string  `xml:"adcurveOrderSource,omitempty" json:"adcurveOrderSource,omitempty" yaml:"adcurveOrderSource,omitempty"`
	AdcurveOrderId       *string  `xml:"adcurveOrderId,omitempty" json:"adcurveOrderId,omitempty" yaml:"adcurveOrderId,omitempty"`
	AdcurveShippingPrice *float64 `xml:"adcurveShippingPrice,omitempty" json:"adcurveShippingPrice,omitempty" yaml:"adcurveShippingPrice,omitempty"`
}

// ShoppingCartPaymentEntity was auto-generated from WSDL.
type ShoppingCartPaymentEntity struct {
	Payment_id             *string `xml:"payment_id,omitempty" json:"payment_id,omitempty" yaml:"payment_id,omitempty"`
	Created_at             *string `xml:"created_at,omitempty" json:"created_at,omitempty" yaml:"created_at,omitempty"`
	Updated_at             *string `xml:"updated_at,omitempty" json:"updated_at,omitempty" yaml:"updated_at,omitempty"`
	Method                 *string `xml:"method,omitempty" json:"method,omitempty" yaml:"method,omitempty"`
	Cc_type                *string `xml:"cc_type,omitempty" json:"cc_type,omitempty" yaml:"cc_type,omitempty"`
	Cc_number_enc          *string `xml:"cc_number_enc,omitempty" json:"cc_number_enc,omitempty" yaml:"cc_number_enc,omitempty"`
	Cc_last4               *string `xml:"cc_last4,omitempty" json:"cc_last4,omitempty" yaml:"cc_last4,omitempty"`
	Cc_cid_enc             *string `xml:"cc_cid_enc,omitempty" json:"cc_cid_enc,omitempty" yaml:"cc_cid_enc,omitempty"`
	Cc_owner               *string `xml:"cc_owner,omitempty" json:"cc_owner,omitempty" yaml:"cc_owner,omitempty"`
	Cc_exp_month           *string `xml:"cc_exp_month,omitempty" json:"cc_exp_month,omitempty" yaml:"cc_exp_month,omitempty"`
	Cc_exp_year            *string `xml:"cc_exp_year,omitempty" json:"cc_exp_year,omitempty" yaml:"cc_exp_year,omitempty"`
	Cc_ss_owner            *string `xml:"cc_ss_owner,omitempty" json:"cc_ss_owner,omitempty" yaml:"cc_ss_owner,omitempty"`
	Cc_ss_start_month      *string `xml:"cc_ss_start_month,omitempty" json:"cc_ss_start_month,omitempty" yaml:"cc_ss_start_month,omitempty"`
	Cc_ss_start_year       *string `xml:"cc_ss_start_year,omitempty" json:"cc_ss_start_year,omitempty" yaml:"cc_ss_start_year,omitempty"`
	Cc_ss_issue            *string `xml:"cc_ss_issue,omitempty" json:"cc_ss_issue,omitempty" yaml:"cc_ss_issue,omitempty"`
	Po_number              *string `xml:"po_number,omitempty" json:"po_number,omitempty" yaml:"po_number,omitempty"`
	Additional_data        *string `xml:"additional_data,omitempty" json:"additional_data,omitempty" yaml:"additional_data,omitempty"`
	Additional_information *string `xml:"additional_information,omitempty" json:"additional_information,omitempty" yaml:"additional_information,omitempty"`
}

// ShoppingCartPaymentMethodEntity was auto-generated from WSDL.
type ShoppingCartPaymentMethodEntity struct {
	Po_number    *string `xml:"po_number,omitempty" json:"po_number,omitempty" yaml:"po_number,omitempty"`
	Method       *string `xml:"method,omitempty" json:"method,omitempty" yaml:"method,omitempty"`
	Cc_cid       *string `xml:"cc_cid,omitempty" json:"cc_cid,omitempty" yaml:"cc_cid,omitempty"`
	Cc_owner     *string `xml:"cc_owner,omitempty" json:"cc_owner,omitempty" yaml:"cc_owner,omitempty"`
	Cc_number    *string `xml:"cc_number,omitempty" json:"cc_number,omitempty" yaml:"cc_number,omitempty"`
	Cc_type      *string `xml:"cc_type,omitempty" json:"cc_type,omitempty" yaml:"cc_type,omitempty"`
	Cc_exp_year  *string `xml:"cc_exp_year,omitempty" json:"cc_exp_year,omitempty" yaml:"cc_exp_year,omitempty"`
	Cc_exp_month *string `xml:"cc_exp_month,omitempty" json:"cc_exp_month,omitempty" yaml:"cc_exp_month,omitempty"`
}

// ShoppingCartPaymentMethodResponseEntity was auto-generated from
// WSDL.
type ShoppingCartPaymentMethodResponseEntity struct {
	Code     *string           `xml:"code,omitempty" json:"code,omitempty" yaml:"code,omitempty"`
	Title    *string           `xml:"title,omitempty" json:"title,omitempty" yaml:"title,omitempty"`
	Cc_types *AssociativeArray `xml:"cc_types,omitempty" json:"cc_types,omitempty" yaml:"cc_types,omitempty"`
}

// ShoppingCartPaymentMethodResponseEntityArray was auto-generated
// from WSDL.
type ShoppingCartPaymentMethodResponseEntityArray struct {
	Items []*ShoppingCartPaymentMethodResponseEntity `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
}

// ShoppingCartProductEntity was auto-generated from WSDL.
type ShoppingCartProductEntity struct {
	Product_id        *string           `xml:"product_id,omitempty" json:"product_id,omitempty" yaml:"product_id,omitempty"`
	Sku               *string           `xml:"sku,omitempty" json:"sku,omitempty" yaml:"sku,omitempty"`
	Qty               *float64          `xml:"qty,omitempty" json:"qty,omitempty" yaml:"qty,omitempty"`
	Options           *AssociativeArray `xml:"options,omitempty" json:"options,omitempty" yaml:"options,omitempty"`
	Bundle_option     *AssociativeArray `xml:"bundle_option,omitempty" json:"bundle_option,omitempty" yaml:"bundle_option,omitempty"`
	Bundle_option_qty *AssociativeArray `xml:"bundle_option_qty,omitempty" json:"bundle_option_qty,omitempty" yaml:"bundle_option_qty,omitempty"`
	Links             *ArrayOfString    `xml:"links,omitempty" json:"links,omitempty" yaml:"links,omitempty"`
}

// ShoppingCartProductEntityArray was auto-generated from WSDL.
type ShoppingCartProductEntityArray struct {
	Items []*ShoppingCartProductEntity `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
}

// ShoppingCartProductResponseEntityArray was auto-generated from
// WSDL.
type ShoppingCartProductResponseEntityArray struct {
	Items []*CatalogProductEntity `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
}

// ShoppingCartShippingMethodEntity was auto-generated from WSDL.
type ShoppingCartShippingMethodEntity struct {
	Code               *string  `xml:"code,omitempty" json:"code,omitempty" yaml:"code,omitempty"`
	Carrier            *string  `xml:"carrier,omitempty" json:"carrier,omitempty" yaml:"carrier,omitempty"`
	Carrier_title      *string  `xml:"carrier_title,omitempty" json:"carrier_title,omitempty" yaml:"carrier_title,omitempty"`
	Method             *string  `xml:"method,omitempty" json:"method,omitempty" yaml:"method,omitempty"`
	Method_title       *string  `xml:"method_title,omitempty" json:"method_title,omitempty" yaml:"method_title,omitempty"`
	Method_description *string  `xml:"method_description,omitempty" json:"method_description,omitempty" yaml:"method_description,omitempty"`
	Price              *float64 `xml:"price,omitempty" json:"price,omitempty" yaml:"price,omitempty"`
}

// ShoppingCartShippingMethodEntityArray was auto-generated from
// WSDL.
type ShoppingCartShippingMethodEntityArray struct {
	Items []*ShoppingCartShippingMethodEntity `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
}

// ShoppingCartTotalsEntity was auto-generated from WSDL.
type ShoppingCartTotalsEntity struct {
	Title  *string  `xml:"title,omitempty" json:"title,omitempty" yaml:"title,omitempty"`
	Amount *float64 `xml:"amount,omitempty" json:"amount,omitempty" yaml:"amount,omitempty"`
}

// ShoppingCartTotalsEntityArray was auto-generated from WSDL.
type ShoppingCartTotalsEntityArray struct {
	Items []*ShoppingCartTotalsEntity `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
}

// Short was auto-generated from WSDL.
type Short struct {
}

// StoreEntity was auto-generated from WSDL.
type StoreEntity struct {
	Store_id   *int    `xml:"store_id,omitempty" json:"store_id,omitempty" yaml:"store_id,omitempty"`
	Code       *string `xml:"code,omitempty" json:"code,omitempty" yaml:"code,omitempty"`
	Website_id *int    `xml:"website_id,omitempty" json:"website_id,omitempty" yaml:"website_id,omitempty"`
	Group_id   *int    `xml:"group_id,omitempty" json:"group_id,omitempty" yaml:"group_id,omitempty"`
	Name       *string `xml:"name,omitempty" json:"name,omitempty" yaml:"name,omitempty"`
	Sort_order *int    `xml:"sort_order,omitempty" json:"sort_order,omitempty" yaml:"sort_order,omitempty"`
	Is_active  *int    `xml:"is_active,omitempty" json:"is_active,omitempty" yaml:"is_active,omitempty"`
}

// StoreEntityArray was auto-generated from WSDL.
type StoreEntityArray struct {
	Items []*StoreEntity `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
}

// String was auto-generated from WSDL.
type String struct {
}

// Time was auto-generated from WSDL.
type Time struct {
}

// Token was auto-generated from WSDL.
type Token struct {
}

// UnsignedByte was auto-generated from WSDL.
type UnsignedByte struct {
}

// UnsignedInt was auto-generated from WSDL.
type UnsignedInt struct {
}

// UnsignedLong was auto-generated from WSDL.
type UnsignedLong struct {
}

// UnsignedShort was auto-generated from WSDL.
type UnsignedShort struct {
}

// Operation wrapper for AdcurveShoppingCartCreateAdcurveOrder.
// OperationAdcurveShoppingCartCreateAdcurveOrderRequest was auto-generated
// from WSDL.
type OperationAdcurveShoppingCartCreateAdcurveOrderRequest struct {
	SessionId                 *string                                      `xml:"sessionId,omitempty" json:"sessionId,omitempty" yaml:"sessionId,omitempty"`
	QuoteId                   *int                                         `xml:"quoteId,omitempty" json:"quoteId,omitempty" yaml:"quoteId,omitempty"`
	StoreId                   *string                                      `xml:"storeId,omitempty" json:"storeId,omitempty" yaml:"storeId,omitempty"`
	Licenses                  *ArrayOfString                               `xml:"licenses,omitempty" json:"licenses,omitempty" yaml:"licenses,omitempty"`
	OrderAdditionalAttributes *ShoppingCartOrderAdditionalAttributesEntity `xml:"orderAdditionalAttributes,omitempty" json:"orderAdditionalAttributes,omitempty" yaml:"orderAdditionalAttributes,omitempty"`
}

// Operation wrapper for AdcurveShoppingCartCreateAdcurveOrder.
// OperationAdcurveShoppingCartCreateAdcurveOrderResponse was auto-generated
// from WSDL.
type OperationAdcurveShoppingCartCreateAdcurveOrderResponse struct {
	Result *string `xml:"result,omitempty" json:"result,omitempty" yaml:"result,omitempty"`
}

// Operation wrapper for AdcurveShoppingCartProductAdcurveProductAdd.
// OperationAdcurveShoppingCartProductAdcurveProductAddRequest
// was auto-generated from WSDL.
type OperationAdcurveShoppingCartProductAdcurveProductAddRequest struct {
	SessionId *string                                `xml:"sessionId,omitempty" json:"sessionId,omitempty" yaml:"sessionId,omitempty"`
	QuoteId   *int                                   `xml:"quoteId,omitempty" json:"quoteId,omitempty" yaml:"quoteId,omitempty"`
	Products  *ShoppingCartAdcurveProductEntityArray `xml:"products,omitempty" json:"products,omitempty" yaml:"products,omitempty"`
	StoreId   *string                                `xml:"storeId,omitempty" json:"storeId,omitempty" yaml:"storeId,omitempty"`
}

// Operation wrapper for AdcurveShoppingCartProductAdcurveProductAdd.
// OperationAdcurveShoppingCartProductAdcurveProductAddResponse
// was auto-generated from WSDL.
type OperationAdcurveShoppingCartProductAdcurveProductAddResponse struct {
	Result *bool `xml:"result,omitempty" json:"result,omitempty" yaml:"result,omitempty"`
}

// Operation wrapper for AdcurveShoppingCartProductAdcurveProductUpdate.
// OperationAdcurveShoppingCartProductAdcurveProductUpdateRequest
// was auto-generated from WSDL.
type OperationAdcurveShoppingCartProductAdcurveProductUpdateRequest struct {
	SessionId *string                                `xml:"sessionId,omitempty" json:"sessionId,omitempty" yaml:"sessionId,omitempty"`
	QuoteId   *int                                   `xml:"quoteId,omitempty" json:"quoteId,omitempty" yaml:"quoteId,omitempty"`
	Products  *ShoppingCartAdcurveProductEntityArray `xml:"products,omitempty" json:"products,omitempty" yaml:"products,omitempty"`
	StoreId   *string                                `xml:"storeId,omitempty" json:"storeId,omitempty" yaml:"storeId,omitempty"`
}

// Operation wrapper for AdcurveShoppingCartProductAdcurveProductUpdate.
// OperationAdcurveShoppingCartProductAdcurveProductUpdateResponse
// was auto-generated from WSDL.
type OperationAdcurveShoppingCartProductAdcurveProductUpdateResponse struct {
	Result *bool `xml:"result,omitempty" json:"result,omitempty" yaml:"result,omitempty"`
}

// Operation wrapper for CatalogCategoryAssignProduct.
// OperationCatalogCategoryAssignProductRequest was auto-generated
// from WSDL.
type OperationCatalogCategoryAssignProductRequest struct {
	SessionId      *string `xml:"sessionId,omitempty" json:"sessionId,omitempty" yaml:"sessionId,omitempty"`
	CategoryId     *int    `xml:"categoryId,omitempty" json:"categoryId,omitempty" yaml:"categoryId,omitempty"`
	Product        *string `xml:"product,omitempty" json:"product,omitempty" yaml:"product,omitempty"`
	Position       *string `xml:"position,omitempty" json:"position,omitempty" yaml:"position,omitempty"`
	IdentifierType *string `xml:"identifierType,omitempty" json:"identifierType,omitempty" yaml:"identifierType,omitempty"`
}

// Operation wrapper for CatalogCategoryAssignProduct.
// OperationCatalogCategoryAssignProductResponse was auto-generated
// from WSDL.
type OperationCatalogCategoryAssignProductResponse struct {
	Result *bool `xml:"result,omitempty" json:"result,omitempty" yaml:"result,omitempty"`
}

// Operation wrapper for CatalogCategoryAssignedProducts.
// OperationCatalogCategoryAssignedProductsRequest was auto-generated
// from WSDL.
type OperationCatalogCategoryAssignedProductsRequest struct {
	SessionId  *string `xml:"sessionId,omitempty" json:"sessionId,omitempty" yaml:"sessionId,omitempty"`
	CategoryId *int    `xml:"categoryId,omitempty" json:"categoryId,omitempty" yaml:"categoryId,omitempty"`
}

// Operation wrapper for CatalogCategoryAssignedProducts.
// OperationCatalogCategoryAssignedProductsResponse was auto-generated
// from WSDL.
type OperationCatalogCategoryAssignedProductsResponse struct {
	Result *CatalogAssignedProductArray `xml:"result,omitempty" json:"result,omitempty" yaml:"result,omitempty"`
}

// Operation wrapper for CatalogCategoryAttributeCurrentStore.
// OperationCatalogCategoryAttributeCurrentStoreRequest was auto-generated
// from WSDL.
type OperationCatalogCategoryAttributeCurrentStoreRequest struct {
	SessionId *string `xml:"sessionId,omitempty" json:"sessionId,omitempty" yaml:"sessionId,omitempty"`
	StoreView *string `xml:"storeView,omitempty" json:"storeView,omitempty" yaml:"storeView,omitempty"`
}

// Operation wrapper for CatalogCategoryAttributeCurrentStore.
// OperationCatalogCategoryAttributeCurrentStoreResponse was auto-generated
// from WSDL.
type OperationCatalogCategoryAttributeCurrentStoreResponse struct {
	StoreView *int `xml:"storeView,omitempty" json:"storeView,omitempty" yaml:"storeView,omitempty"`
}

// Operation wrapper for CatalogCategoryAttributeList.
// OperationCatalogCategoryAttributeListRequest was auto-generated
// from WSDL.
type OperationCatalogCategoryAttributeListRequest struct {
	SessionId *string `xml:"sessionId,omitempty" json:"sessionId,omitempty" yaml:"sessionId,omitempty"`
}

// Operation wrapper for CatalogCategoryAttributeList.
// OperationCatalogCategoryAttributeListResponse was auto-generated
// from WSDL.
type OperationCatalogCategoryAttributeListResponse struct {
	Result *CatalogAttributeEntityArray `xml:"result,omitempty" json:"result,omitempty" yaml:"result,omitempty"`
}

// Operation wrapper for CatalogCategoryAttributeOptions.
// OperationCatalogCategoryAttributeOptionsRequest was auto-generated
// from WSDL.
type OperationCatalogCategoryAttributeOptionsRequest struct {
	SessionId   *string `xml:"sessionId,omitempty" json:"sessionId,omitempty" yaml:"sessionId,omitempty"`
	AttributeId *string `xml:"attributeId,omitempty" json:"attributeId,omitempty" yaml:"attributeId,omitempty"`
	StoreView   *string `xml:"storeView,omitempty" json:"storeView,omitempty" yaml:"storeView,omitempty"`
}

// Operation wrapper for CatalogCategoryAttributeOptions.
// OperationCatalogCategoryAttributeOptionsResponse was auto-generated
// from WSDL.
type OperationCatalogCategoryAttributeOptionsResponse struct {
	Result *CatalogAttributeOptionEntityArray `xml:"result,omitempty" json:"result,omitempty" yaml:"result,omitempty"`
}

// Operation wrapper for CatalogCategoryCreate.
// OperationCatalogCategoryCreateRequest was auto-generated from
// WSDL.
type OperationCatalogCategoryCreateRequest struct {
	SessionId    *string                      `xml:"sessionId,omitempty" json:"sessionId,omitempty" yaml:"sessionId,omitempty"`
	ParentId     *int                         `xml:"parentId,omitempty" json:"parentId,omitempty" yaml:"parentId,omitempty"`
	CategoryData *CatalogCategoryEntityCreate `xml:"categoryData,omitempty" json:"categoryData,omitempty" yaml:"categoryData,omitempty"`
	StoreView    *string                      `xml:"storeView,omitempty" json:"storeView,omitempty" yaml:"storeView,omitempty"`
}

// Operation wrapper for CatalogCategoryCreate.
// OperationCatalogCategoryCreateResponse was auto-generated from
// WSDL.
type OperationCatalogCategoryCreateResponse struct {
	Attribute_id *int `xml:"attribute_id,omitempty" json:"attribute_id,omitempty" yaml:"attribute_id,omitempty"`
}

// Operation wrapper for CatalogCategoryCurrentStore.
// OperationCatalogCategoryCurrentStoreRequest was auto-generated
// from WSDL.
type OperationCatalogCategoryCurrentStoreRequest struct {
	SessionId *string `xml:"sessionId,omitempty" json:"sessionId,omitempty" yaml:"sessionId,omitempty"`
	StoreView *string `xml:"storeView,omitempty" json:"storeView,omitempty" yaml:"storeView,omitempty"`
}

// Operation wrapper for CatalogCategoryCurrentStore.
// OperationCatalogCategoryCurrentStoreResponse was auto-generated
// from WSDL.
type OperationCatalogCategoryCurrentStoreResponse struct {
	StoreView *int `xml:"storeView,omitempty" json:"storeView,omitempty" yaml:"storeView,omitempty"`
}

// Operation wrapper for CatalogCategoryDelete.
// OperationCatalogCategoryDeleteRequest was auto-generated from
// WSDL.
type OperationCatalogCategoryDeleteRequest struct {
	SessionId  *string `xml:"sessionId,omitempty" json:"sessionId,omitempty" yaml:"sessionId,omitempty"`
	CategoryId *int    `xml:"categoryId,omitempty" json:"categoryId,omitempty" yaml:"categoryId,omitempty"`
}

// Operation wrapper for CatalogCategoryDelete.
// OperationCatalogCategoryDeleteResponse was auto-generated from
// WSDL.
type OperationCatalogCategoryDeleteResponse struct {
	Result *bool `xml:"result,omitempty" json:"result,omitempty" yaml:"result,omitempty"`
}

// Operation wrapper for CatalogCategoryInfo.
// OperationCatalogCategoryInfoRequest was auto-generated from
// WSDL.
type OperationCatalogCategoryInfoRequest struct {
	SessionId  *string        `xml:"sessionId,omitempty" json:"sessionId,omitempty" yaml:"sessionId,omitempty"`
	CategoryId *int           `xml:"categoryId,omitempty" json:"categoryId,omitempty" yaml:"categoryId,omitempty"`
	StoreView  *string        `xml:"storeView,omitempty" json:"storeView,omitempty" yaml:"storeView,omitempty"`
	Attributes *ArrayOfString `xml:"attributes,omitempty" json:"attributes,omitempty" yaml:"attributes,omitempty"`
}

// Operation wrapper for CatalogCategoryInfo.
// OperationCatalogCategoryInfoResponse was auto-generated from
// WSDL.
type OperationCatalogCategoryInfoResponse struct {
	Info *CatalogCategoryInfo `xml:"info,omitempty" json:"info,omitempty" yaml:"info,omitempty"`
}

// Operation wrapper for CatalogCategoryLevel.
// OperationCatalogCategoryLevelRequest was auto-generated from
// WSDL.
type OperationCatalogCategoryLevelRequest struct {
	SessionId      *string `xml:"sessionId,omitempty" json:"sessionId,omitempty" yaml:"sessionId,omitempty"`
	Website        *string `xml:"website,omitempty" json:"website,omitempty" yaml:"website,omitempty"`
	StoreView      *string `xml:"storeView,omitempty" json:"storeView,omitempty" yaml:"storeView,omitempty"`
	ParentCategory *string `xml:"parentCategory,omitempty" json:"parentCategory,omitempty" yaml:"parentCategory,omitempty"`
}

// Operation wrapper for CatalogCategoryLevel.
// OperationCatalogCategoryLevelResponse was auto-generated from
// WSDL.
type OperationCatalogCategoryLevelResponse struct {
	Tree *ArrayOfCatalogCategoryEntitiesNoChildren `xml:"tree,omitempty" json:"tree,omitempty" yaml:"tree,omitempty"`
}

// Operation wrapper for CatalogCategoryMove.
// OperationCatalogCategoryMoveRequest was auto-generated from
// WSDL.
type OperationCatalogCategoryMoveRequest struct {
	SessionId  *string `xml:"sessionId,omitempty" json:"sessionId,omitempty" yaml:"sessionId,omitempty"`
	CategoryId *int    `xml:"categoryId,omitempty" json:"categoryId,omitempty" yaml:"categoryId,omitempty"`
	ParentId   *int    `xml:"parentId,omitempty" json:"parentId,omitempty" yaml:"parentId,omitempty"`
	AfterId    *string `xml:"afterId,omitempty" json:"afterId,omitempty" yaml:"afterId,omitempty"`
}

// Operation wrapper for CatalogCategoryMove.
// OperationCatalogCategoryMoveResponse was auto-generated from
// WSDL.
type OperationCatalogCategoryMoveResponse struct {
	Id *bool `xml:"id,omitempty" json:"id,omitempty" yaml:"id,omitempty"`
}

// Operation wrapper for CatalogCategoryRemoveProduct.
// OperationCatalogCategoryRemoveProductRequest was auto-generated
// from WSDL.
type OperationCatalogCategoryRemoveProductRequest struct {
	SessionId      *string `xml:"sessionId,omitempty" json:"sessionId,omitempty" yaml:"sessionId,omitempty"`
	CategoryId     *int    `xml:"categoryId,omitempty" json:"categoryId,omitempty" yaml:"categoryId,omitempty"`
	Product        *string `xml:"product,omitempty" json:"product,omitempty" yaml:"product,omitempty"`
	IdentifierType *string `xml:"identifierType,omitempty" json:"identifierType,omitempty" yaml:"identifierType,omitempty"`
}

// Operation wrapper for CatalogCategoryRemoveProduct.
// OperationCatalogCategoryRemoveProductResponse was auto-generated
// from WSDL.
type OperationCatalogCategoryRemoveProductResponse struct {
	Result *bool `xml:"result,omitempty" json:"result,omitempty" yaml:"result,omitempty"`
}

// Operation wrapper for CatalogCategoryTree.
// OperationCatalogCategoryTreeRequest was auto-generated from
// WSDL.
type OperationCatalogCategoryTreeRequest struct {
	SessionId *string `xml:"sessionId,omitempty" json:"sessionId,omitempty" yaml:"sessionId,omitempty"`
	ParentId  *string `xml:"parentId,omitempty" json:"parentId,omitempty" yaml:"parentId,omitempty"`
	StoreView *string `xml:"storeView,omitempty" json:"storeView,omitempty" yaml:"storeView,omitempty"`
}

// Operation wrapper for CatalogCategoryTree.
// OperationCatalogCategoryTreeResponse was auto-generated from
// WSDL.
type OperationCatalogCategoryTreeResponse struct {
	Tree *CatalogCategoryTree `xml:"tree,omitempty" json:"tree,omitempty" yaml:"tree,omitempty"`
}

// Operation wrapper for CatalogCategoryUpdate.
// OperationCatalogCategoryUpdateRequest was auto-generated from
// WSDL.
type OperationCatalogCategoryUpdateRequest struct {
	SessionId    *string                      `xml:"sessionId,omitempty" json:"sessionId,omitempty" yaml:"sessionId,omitempty"`
	CategoryId   *int                         `xml:"categoryId,omitempty" json:"categoryId,omitempty" yaml:"categoryId,omitempty"`
	CategoryData *CatalogCategoryEntityCreate `xml:"categoryData,omitempty" json:"categoryData,omitempty" yaml:"categoryData,omitempty"`
	StoreView    *string                      `xml:"storeView,omitempty" json:"storeView,omitempty" yaml:"storeView,omitempty"`
}

// Operation wrapper for CatalogCategoryUpdate.
// OperationCatalogCategoryUpdateResponse was auto-generated from
// WSDL.
type OperationCatalogCategoryUpdateResponse struct {
	Id *bool `xml:"id,omitempty" json:"id,omitempty" yaml:"id,omitempty"`
}

// Operation wrapper for CatalogCategoryUpdateProduct.
// OperationCatalogCategoryUpdateProductRequest was auto-generated
// from WSDL.
type OperationCatalogCategoryUpdateProductRequest struct {
	SessionId      *string `xml:"sessionId,omitempty" json:"sessionId,omitempty" yaml:"sessionId,omitempty"`
	CategoryId     *int    `xml:"categoryId,omitempty" json:"categoryId,omitempty" yaml:"categoryId,omitempty"`
	Product        *string `xml:"product,omitempty" json:"product,omitempty" yaml:"product,omitempty"`
	Position       *string `xml:"position,omitempty" json:"position,omitempty" yaml:"position,omitempty"`
	IdentifierType *string `xml:"identifierType,omitempty" json:"identifierType,omitempty" yaml:"identifierType,omitempty"`
}

// Operation wrapper for CatalogCategoryUpdateProduct.
// OperationCatalogCategoryUpdateProductResponse was auto-generated
// from WSDL.
type OperationCatalogCategoryUpdateProductResponse struct {
	Result *bool `xml:"result,omitempty" json:"result,omitempty" yaml:"result,omitempty"`
}

// Operation wrapper for CatalogInventoryStockItemList.
// OperationCatalogInventoryStockItemListRequest was auto-generated
// from WSDL.
type OperationCatalogInventoryStockItemListRequest struct {
	SessionId *string        `xml:"sessionId,omitempty" json:"sessionId,omitempty" yaml:"sessionId,omitempty"`
	Products  *ArrayOfString `xml:"products,omitempty" json:"products,omitempty" yaml:"products,omitempty"`
}

// Operation wrapper for CatalogInventoryStockItemList.
// OperationCatalogInventoryStockItemListResponse was auto-generated
// from WSDL.
type OperationCatalogInventoryStockItemListResponse struct {
	Result *CatalogInventoryStockItemEntityArray `xml:"result,omitempty" json:"result,omitempty" yaml:"result,omitempty"`
}

// Operation wrapper for CatalogInventoryStockItemMultiUpdate.
// OperationCatalogInventoryStockItemMultiUpdateRequest was auto-generated
// from WSDL.
type OperationCatalogInventoryStockItemMultiUpdateRequest struct {
	SessionId   *string                                     `xml:"sessionId,omitempty" json:"sessionId,omitempty" yaml:"sessionId,omitempty"`
	ProductIds  *ArrayOfString                              `xml:"productIds,omitempty" json:"productIds,omitempty" yaml:"productIds,omitempty"`
	ProductData *CatalogInventoryStockItemUpdateEntityArray `xml:"productData,omitempty" json:"productData,omitempty" yaml:"productData,omitempty"`
}

// Operation wrapper for CatalogInventoryStockItemMultiUpdate.
// OperationCatalogInventoryStockItemMultiUpdateResponse was auto-generated
// from WSDL.
type OperationCatalogInventoryStockItemMultiUpdateResponse struct {
	Result *bool `xml:"result,omitempty" json:"result,omitempty" yaml:"result,omitempty"`
}

// Operation wrapper for CatalogInventoryStockItemUpdate.
// OperationCatalogInventoryStockItemUpdateRequest was auto-generated
// from WSDL.
type OperationCatalogInventoryStockItemUpdateRequest struct {
	SessionId *string                                `xml:"sessionId,omitempty" json:"sessionId,omitempty" yaml:"sessionId,omitempty"`
	Product   *string                                `xml:"product,omitempty" json:"product,omitempty" yaml:"product,omitempty"`
	Data      *CatalogInventoryStockItemUpdateEntity `xml:"data,omitempty" json:"data,omitempty" yaml:"data,omitempty"`
}

// Operation wrapper for CatalogInventoryStockItemUpdate.
// OperationCatalogInventoryStockItemUpdateResponse was auto-generated
// from WSDL.
type OperationCatalogInventoryStockItemUpdateResponse struct {
	Result *int `xml:"result,omitempty" json:"result,omitempty" yaml:"result,omitempty"`
}

// Operation wrapper for CatalogProductAttributeAddOption.
// OperationCatalogProductAttributeAddOptionRequest was auto-generated
// from WSDL.
type OperationCatalogProductAttributeAddOptionRequest struct {
	SessionId *string                                   `xml:"sessionId,omitempty" json:"sessionId,omitempty" yaml:"sessionId,omitempty"`
	Attribute *string                                   `xml:"attribute,omitempty" json:"attribute,omitempty" yaml:"attribute,omitempty"`
	Data      *CatalogProductAttributeOptionEntityToAdd `xml:"data,omitempty" json:"data,omitempty" yaml:"data,omitempty"`
}

// Operation wrapper for CatalogProductAttributeAddOption.
// OperationCatalogProductAttributeAddOptionResponse was auto-generated
// from WSDL.
type OperationCatalogProductAttributeAddOptionResponse struct {
	Result *bool `xml:"result,omitempty" json:"result,omitempty" yaml:"result,omitempty"`
}

// Operation wrapper for CatalogProductAttributeCreate.
// OperationCatalogProductAttributeCreateRequest was auto-generated
// from WSDL.
type OperationCatalogProductAttributeCreateRequest struct {
	SessionId *string                                `xml:"sessionId,omitempty" json:"sessionId,omitempty" yaml:"sessionId,omitempty"`
	Data      *CatalogProductAttributeEntityToCreate `xml:"data,omitempty" json:"data,omitempty" yaml:"data,omitempty"`
}

// Operation wrapper for CatalogProductAttributeCreate.
// OperationCatalogProductAttributeCreateResponse was auto-generated
// from WSDL.
type OperationCatalogProductAttributeCreateResponse struct {
	Result *int `xml:"result,omitempty" json:"result,omitempty" yaml:"result,omitempty"`
}

// Operation wrapper for CatalogProductAttributeCurrentStore.
// OperationCatalogProductAttributeCurrentStoreRequest was auto-generated
// from WSDL.
type OperationCatalogProductAttributeCurrentStoreRequest struct {
	SessionId *string `xml:"sessionId,omitempty" json:"sessionId,omitempty" yaml:"sessionId,omitempty"`
	StoreView *string `xml:"storeView,omitempty" json:"storeView,omitempty" yaml:"storeView,omitempty"`
}

// Operation wrapper for CatalogProductAttributeCurrentStore.
// OperationCatalogProductAttributeCurrentStoreResponse was auto-generated
// from WSDL.
type OperationCatalogProductAttributeCurrentStoreResponse struct {
	StoreView *int `xml:"storeView,omitempty" json:"storeView,omitempty" yaml:"storeView,omitempty"`
}

// Operation wrapper for CatalogProductAttributeInfo.
// OperationCatalogProductAttributeInfoRequest was auto-generated
// from WSDL.
type OperationCatalogProductAttributeInfoRequest struct {
	SessionId *string `xml:"sessionId,omitempty" json:"sessionId,omitempty" yaml:"sessionId,omitempty"`
	Attribute *string `xml:"attribute,omitempty" json:"attribute,omitempty" yaml:"attribute,omitempty"`
}

// Operation wrapper for CatalogProductAttributeInfo.
// OperationCatalogProductAttributeInfoResponse was auto-generated
// from WSDL.
type OperationCatalogProductAttributeInfoResponse struct {
	Result *CatalogProductAttributeEntity `xml:"result,omitempty" json:"result,omitempty" yaml:"result,omitempty"`
}

// Operation wrapper for CatalogProductAttributeList.
// OperationCatalogProductAttributeListRequest was auto-generated
// from WSDL.
type OperationCatalogProductAttributeListRequest struct {
	SessionId *string `xml:"sessionId,omitempty" json:"sessionId,omitempty" yaml:"sessionId,omitempty"`
	SetId     *int    `xml:"setId,omitempty" json:"setId,omitempty" yaml:"setId,omitempty"`
}

// Operation wrapper for CatalogProductAttributeList.
// OperationCatalogProductAttributeListResponse was auto-generated
// from WSDL.
type OperationCatalogProductAttributeListResponse struct {
	Result *CatalogAttributeEntityArray `xml:"result,omitempty" json:"result,omitempty" yaml:"result,omitempty"`
}

// Operation wrapper for CatalogProductAttributeMediaCreate.
// OperationCatalogProductAttributeMediaCreateRequest was auto-generated
// from WSDL.
type OperationCatalogProductAttributeMediaCreateRequest struct {
	SessionId      *string                                   `xml:"sessionId,omitempty" json:"sessionId,omitempty" yaml:"sessionId,omitempty"`
	Product        *string                                   `xml:"product,omitempty" json:"product,omitempty" yaml:"product,omitempty"`
	Data           *CatalogProductAttributeMediaCreateEntity `xml:"data,omitempty" json:"data,omitempty" yaml:"data,omitempty"`
	StoreView      *string                                   `xml:"storeView,omitempty" json:"storeView,omitempty" yaml:"storeView,omitempty"`
	IdentifierType *string                                   `xml:"identifierType,omitempty" json:"identifierType,omitempty" yaml:"identifierType,omitempty"`
}

// Operation wrapper for CatalogProductAttributeMediaCreate.
// OperationCatalogProductAttributeMediaCreateResponse was auto-generated
// from WSDL.
type OperationCatalogProductAttributeMediaCreateResponse struct {
	Result *string `xml:"result,omitempty" json:"result,omitempty" yaml:"result,omitempty"`
}

// Operation wrapper for CatalogProductAttributeMediaCurrentStore.
// OperationCatalogProductAttributeMediaCurrentStoreRequest was
// auto-generated from WSDL.
type OperationCatalogProductAttributeMediaCurrentStoreRequest struct {
	SessionId *string `xml:"sessionId,omitempty" json:"sessionId,omitempty" yaml:"sessionId,omitempty"`
	StoreView *string `xml:"storeView,omitempty" json:"storeView,omitempty" yaml:"storeView,omitempty"`
}

// Operation wrapper for CatalogProductAttributeMediaCurrentStore.
// OperationCatalogProductAttributeMediaCurrentStoreResponse was
// auto-generated from WSDL.
type OperationCatalogProductAttributeMediaCurrentStoreResponse struct {
	StoreView *int `xml:"storeView,omitempty" json:"storeView,omitempty" yaml:"storeView,omitempty"`
}

// Operation wrapper for CatalogProductAttributeMediaInfo.
// OperationCatalogProductAttributeMediaInfoRequest was auto-generated
// from WSDL.
type OperationCatalogProductAttributeMediaInfoRequest struct {
	SessionId      *string `xml:"sessionId,omitempty" json:"sessionId,omitempty" yaml:"sessionId,omitempty"`
	Product        *string `xml:"product,omitempty" json:"product,omitempty" yaml:"product,omitempty"`
	File           *string `xml:"file,omitempty" json:"file,omitempty" yaml:"file,omitempty"`
	StoreView      *string `xml:"storeView,omitempty" json:"storeView,omitempty" yaml:"storeView,omitempty"`
	IdentifierType *string `xml:"identifierType,omitempty" json:"identifierType,omitempty" yaml:"identifierType,omitempty"`
}

// Operation wrapper for CatalogProductAttributeMediaInfo.
// OperationCatalogProductAttributeMediaInfoResponse was auto-generated
// from WSDL.
type OperationCatalogProductAttributeMediaInfoResponse struct {
	Result *CatalogProductImageEntity `xml:"result,omitempty" json:"result,omitempty" yaml:"result,omitempty"`
}

// Operation wrapper for CatalogProductAttributeMediaList.
// OperationCatalogProductAttributeMediaListRequest was auto-generated
// from WSDL.
type OperationCatalogProductAttributeMediaListRequest struct {
	SessionId      *string `xml:"sessionId,omitempty" json:"sessionId,omitempty" yaml:"sessionId,omitempty"`
	Product        *string `xml:"product,omitempty" json:"product,omitempty" yaml:"product,omitempty"`
	StoreView      *string `xml:"storeView,omitempty" json:"storeView,omitempty" yaml:"storeView,omitempty"`
	IdentifierType *string `xml:"identifierType,omitempty" json:"identifierType,omitempty" yaml:"identifierType,omitempty"`
}

// Operation wrapper for CatalogProductAttributeMediaList.
// OperationCatalogProductAttributeMediaListResponse was auto-generated
// from WSDL.
type OperationCatalogProductAttributeMediaListResponse struct {
	Result *CatalogProductImageEntityArray `xml:"result,omitempty" json:"result,omitempty" yaml:"result,omitempty"`
}

// Operation wrapper for CatalogProductAttributeMediaRemove.
// OperationCatalogProductAttributeMediaRemoveRequest was auto-generated
// from WSDL.
type OperationCatalogProductAttributeMediaRemoveRequest struct {
	SessionId      *string `xml:"sessionId,omitempty" json:"sessionId,omitempty" yaml:"sessionId,omitempty"`
	Product        *string `xml:"product,omitempty" json:"product,omitempty" yaml:"product,omitempty"`
	File           *string `xml:"file,omitempty" json:"file,omitempty" yaml:"file,omitempty"`
	IdentifierType *string `xml:"identifierType,omitempty" json:"identifierType,omitempty" yaml:"identifierType,omitempty"`
}

// Operation wrapper for CatalogProductAttributeMediaRemove.
// OperationCatalogProductAttributeMediaRemoveResponse was auto-generated
// from WSDL.
type OperationCatalogProductAttributeMediaRemoveResponse struct {
	Result *bool `xml:"result,omitempty" json:"result,omitempty" yaml:"result,omitempty"`
}

// Operation wrapper for CatalogProductAttributeMediaTypes.
// OperationCatalogProductAttributeMediaTypesRequest was auto-generated
// from WSDL.
type OperationCatalogProductAttributeMediaTypesRequest struct {
	SessionId *string `xml:"sessionId,omitempty" json:"sessionId,omitempty" yaml:"sessionId,omitempty"`
	SetId     *string `xml:"setId,omitempty" json:"setId,omitempty" yaml:"setId,omitempty"`
}

// Operation wrapper for CatalogProductAttributeMediaTypes.
// OperationCatalogProductAttributeMediaTypesResponse was auto-generated
// from WSDL.
type OperationCatalogProductAttributeMediaTypesResponse struct {
	Result *CatalogProductAttributeMediaTypeEntityArray `xml:"result,omitempty" json:"result,omitempty" yaml:"result,omitempty"`
}

// Operation wrapper for CatalogProductAttributeMediaUpdate.
// OperationCatalogProductAttributeMediaUpdateRequest was auto-generated
// from WSDL.
type OperationCatalogProductAttributeMediaUpdateRequest struct {
	SessionId      *string                                   `xml:"sessionId,omitempty" json:"sessionId,omitempty" yaml:"sessionId,omitempty"`
	Product        *string                                   `xml:"product,omitempty" json:"product,omitempty" yaml:"product,omitempty"`
	File           *string                                   `xml:"file,omitempty" json:"file,omitempty" yaml:"file,omitempty"`
	Data           *CatalogProductAttributeMediaCreateEntity `xml:"data,omitempty" json:"data,omitempty" yaml:"data,omitempty"`
	StoreView      *string                                   `xml:"storeView,omitempty" json:"storeView,omitempty" yaml:"storeView,omitempty"`
	IdentifierType *string                                   `xml:"identifierType,omitempty" json:"identifierType,omitempty" yaml:"identifierType,omitempty"`
}

// Operation wrapper for CatalogProductAttributeMediaUpdate.
// OperationCatalogProductAttributeMediaUpdateResponse was auto-generated
// from WSDL.
type OperationCatalogProductAttributeMediaUpdateResponse struct {
	Result *bool `xml:"result,omitempty" json:"result,omitempty" yaml:"result,omitempty"`
}

// Operation wrapper for CatalogProductAttributeOptions.
// OperationCatalogProductAttributeOptionsRequest was auto-generated
// from WSDL.
type OperationCatalogProductAttributeOptionsRequest struct {
	SessionId   *string `xml:"sessionId,omitempty" json:"sessionId,omitempty" yaml:"sessionId,omitempty"`
	AttributeId *string `xml:"attributeId,omitempty" json:"attributeId,omitempty" yaml:"attributeId,omitempty"`
	StoreView   *string `xml:"storeView,omitempty" json:"storeView,omitempty" yaml:"storeView,omitempty"`
}

// Operation wrapper for CatalogProductAttributeOptions.
// OperationCatalogProductAttributeOptionsResponse was auto-generated
// from WSDL.
type OperationCatalogProductAttributeOptionsResponse struct {
	Result *CatalogAttributeOptionEntityArray `xml:"result,omitempty" json:"result,omitempty" yaml:"result,omitempty"`
}

// Operation wrapper for CatalogProductAttributeRemove.
// OperationCatalogProductAttributeRemoveRequest was auto-generated
// from WSDL.
type OperationCatalogProductAttributeRemoveRequest struct {
	SessionId *string `xml:"sessionId,omitempty" json:"sessionId,omitempty" yaml:"sessionId,omitempty"`
	Attribute *string `xml:"attribute,omitempty" json:"attribute,omitempty" yaml:"attribute,omitempty"`
}

// Operation wrapper for CatalogProductAttributeRemove.
// OperationCatalogProductAttributeRemoveResponse was auto-generated
// from WSDL.
type OperationCatalogProductAttributeRemoveResponse struct {
	Result *bool `xml:"result,omitempty" json:"result,omitempty" yaml:"result,omitempty"`
}

// Operation wrapper for CatalogProductAttributeRemoveOption.
// OperationCatalogProductAttributeRemoveOptionRequest was auto-generated
// from WSDL.
type OperationCatalogProductAttributeRemoveOptionRequest struct {
	SessionId *string `xml:"sessionId,omitempty" json:"sessionId,omitempty" yaml:"sessionId,omitempty"`
	Attribute *string `xml:"attribute,omitempty" json:"attribute,omitempty" yaml:"attribute,omitempty"`
	OptionId  *string `xml:"optionId,omitempty" json:"optionId,omitempty" yaml:"optionId,omitempty"`
}

// Operation wrapper for CatalogProductAttributeRemoveOption.
// OperationCatalogProductAttributeRemoveOptionResponse was auto-generated
// from WSDL.
type OperationCatalogProductAttributeRemoveOptionResponse struct {
	Result *bool `xml:"result,omitempty" json:"result,omitempty" yaml:"result,omitempty"`
}

// Operation wrapper for CatalogProductAttributeSetAttributeAdd.
// OperationCatalogProductAttributeSetAttributeAddRequest was auto-generated
// from WSDL.
type OperationCatalogProductAttributeSetAttributeAddRequest struct {
	SessionId        *string `xml:"sessionId,omitempty" json:"sessionId,omitempty" yaml:"sessionId,omitempty"`
	AttributeId      *string `xml:"attributeId,omitempty" json:"attributeId,omitempty" yaml:"attributeId,omitempty"`
	AttributeSetId   *string `xml:"attributeSetId,omitempty" json:"attributeSetId,omitempty" yaml:"attributeSetId,omitempty"`
	AttributeGroupId *string `xml:"attributeGroupId,omitempty" json:"attributeGroupId,omitempty" yaml:"attributeGroupId,omitempty"`
	SortOrder        *string `xml:"sortOrder,omitempty" json:"sortOrder,omitempty" yaml:"sortOrder,omitempty"`
}

// Operation wrapper for CatalogProductAttributeSetAttributeAdd.
// OperationCatalogProductAttributeSetAttributeAddResponse was
// auto-generated from WSDL.
type OperationCatalogProductAttributeSetAttributeAddResponse struct {
	IsAdded *bool `xml:"isAdded,omitempty" json:"isAdded,omitempty" yaml:"isAdded,omitempty"`
}

// Operation wrapper for CatalogProductAttributeSetAttributeRemove.
// OperationCatalogProductAttributeSetAttributeRemoveRequest was
// auto-generated from WSDL.
type OperationCatalogProductAttributeSetAttributeRemoveRequest struct {
	SessionId      *string `xml:"sessionId,omitempty" json:"sessionId,omitempty" yaml:"sessionId,omitempty"`
	AttributeId    *string `xml:"attributeId,omitempty" json:"attributeId,omitempty" yaml:"attributeId,omitempty"`
	AttributeSetId *string `xml:"attributeSetId,omitempty" json:"attributeSetId,omitempty" yaml:"attributeSetId,omitempty"`
}

// Operation wrapper for CatalogProductAttributeSetAttributeRemove.
// OperationCatalogProductAttributeSetAttributeRemoveResponse was
// auto-generated from WSDL.
type OperationCatalogProductAttributeSetAttributeRemoveResponse struct {
	IsRemoved *bool `xml:"isRemoved,omitempty" json:"isRemoved,omitempty" yaml:"isRemoved,omitempty"`
}

// Operation wrapper for CatalogProductAttributeSetCreate.
// OperationCatalogProductAttributeSetCreateRequest was auto-generated
// from WSDL.
type OperationCatalogProductAttributeSetCreateRequest struct {
	SessionId        *string `xml:"sessionId,omitempty" json:"sessionId,omitempty" yaml:"sessionId,omitempty"`
	AttributeSetName *string `xml:"attributeSetName,omitempty" json:"attributeSetName,omitempty" yaml:"attributeSetName,omitempty"`
	SkeletonSetId    *string `xml:"skeletonSetId,omitempty" json:"skeletonSetId,omitempty" yaml:"skeletonSetId,omitempty"`
}

// Operation wrapper for CatalogProductAttributeSetCreate.
// OperationCatalogProductAttributeSetCreateResponse was auto-generated
// from WSDL.
type OperationCatalogProductAttributeSetCreateResponse struct {
	SetId *int `xml:"setId,omitempty" json:"setId,omitempty" yaml:"setId,omitempty"`
}

// Operation wrapper for CatalogProductAttributeSetGroupAdd.
// OperationCatalogProductAttributeSetGroupAddRequest was auto-generated
// from WSDL.
type OperationCatalogProductAttributeSetGroupAddRequest struct {
	SessionId      *string `xml:"sessionId,omitempty" json:"sessionId,omitempty" yaml:"sessionId,omitempty"`
	AttributeSetId *string `xml:"attributeSetId,omitempty" json:"attributeSetId,omitempty" yaml:"attributeSetId,omitempty"`
	GroupName      *string `xml:"groupName,omitempty" json:"groupName,omitempty" yaml:"groupName,omitempty"`
}

// Operation wrapper for CatalogProductAttributeSetGroupAdd.
// OperationCatalogProductAttributeSetGroupAddResponse was auto-generated
// from WSDL.
type OperationCatalogProductAttributeSetGroupAddResponse struct {
	Result *int `xml:"result,omitempty" json:"result,omitempty" yaml:"result,omitempty"`
}

// Operation wrapper for CatalogProductAttributeSetGroupRemove.
// OperationCatalogProductAttributeSetGroupRemoveRequest was auto-generated
// from WSDL.
type OperationCatalogProductAttributeSetGroupRemoveRequest struct {
	SessionId        *string `xml:"sessionId,omitempty" json:"sessionId,omitempty" yaml:"sessionId,omitempty"`
	AttributeGroupId *string `xml:"attributeGroupId,omitempty" json:"attributeGroupId,omitempty" yaml:"attributeGroupId,omitempty"`
}

// Operation wrapper for CatalogProductAttributeSetGroupRemove.
// OperationCatalogProductAttributeSetGroupRemoveResponse was auto-generated
// from WSDL.
type OperationCatalogProductAttributeSetGroupRemoveResponse struct {
	Result *bool `xml:"result,omitempty" json:"result,omitempty" yaml:"result,omitempty"`
}

// Operation wrapper for CatalogProductAttributeSetGroupRename.
// OperationCatalogProductAttributeSetGroupRenameRequest was auto-generated
// from WSDL.
type OperationCatalogProductAttributeSetGroupRenameRequest struct {
	SessionId *string `xml:"sessionId,omitempty" json:"sessionId,omitempty" yaml:"sessionId,omitempty"`
	GroupId   *string `xml:"groupId,omitempty" json:"groupId,omitempty" yaml:"groupId,omitempty"`
	GroupName *string `xml:"groupName,omitempty" json:"groupName,omitempty" yaml:"groupName,omitempty"`
}

// Operation wrapper for CatalogProductAttributeSetGroupRename.
// OperationCatalogProductAttributeSetGroupRenameResponse was auto-generated
// from WSDL.
type OperationCatalogProductAttributeSetGroupRenameResponse struct {
	Result *bool `xml:"result,omitempty" json:"result,omitempty" yaml:"result,omitempty"`
}

// Operation wrapper for CatalogProductAttributeSetList.
// OperationCatalogProductAttributeSetListRequest was auto-generated
// from WSDL.
type OperationCatalogProductAttributeSetListRequest struct {
	SessionId *string `xml:"sessionId,omitempty" json:"sessionId,omitempty" yaml:"sessionId,omitempty"`
}

// Operation wrapper for CatalogProductAttributeSetList.
// OperationCatalogProductAttributeSetListResponse was auto-generated
// from WSDL.
type OperationCatalogProductAttributeSetListResponse struct {
	Result *CatalogProductAttributeSetEntityArray `xml:"result,omitempty" json:"result,omitempty" yaml:"result,omitempty"`
}

// Operation wrapper for CatalogProductAttributeSetRemove.
// OperationCatalogProductAttributeSetRemoveRequest was auto-generated
// from WSDL.
type OperationCatalogProductAttributeSetRemoveRequest struct {
	SessionId           *string `xml:"sessionId,omitempty" json:"sessionId,omitempty" yaml:"sessionId,omitempty"`
	AttributeSetId      *string `xml:"attributeSetId,omitempty" json:"attributeSetId,omitempty" yaml:"attributeSetId,omitempty"`
	ForceProductsRemove *string `xml:"forceProductsRemove,omitempty" json:"forceProductsRemove,omitempty" yaml:"forceProductsRemove,omitempty"`
}

// Operation wrapper for CatalogProductAttributeSetRemove.
// OperationCatalogProductAttributeSetRemoveResponse was auto-generated
// from WSDL.
type OperationCatalogProductAttributeSetRemoveResponse struct {
	IsRemoved *bool `xml:"isRemoved,omitempty" json:"isRemoved,omitempty" yaml:"isRemoved,omitempty"`
}

// Operation wrapper for CatalogProductAttributeTierPriceInfo.
// OperationCatalogProductAttributeTierPriceInfoRequest was auto-generated
// from WSDL.
type OperationCatalogProductAttributeTierPriceInfoRequest struct {
	SessionId      *string `xml:"sessionId,omitempty" json:"sessionId,omitempty" yaml:"sessionId,omitempty"`
	Product        *string `xml:"product,omitempty" json:"product,omitempty" yaml:"product,omitempty"`
	IdentifierType *string `xml:"identifierType,omitempty" json:"identifierType,omitempty" yaml:"identifierType,omitempty"`
}

// Operation wrapper for CatalogProductAttributeTierPriceInfo.
// OperationCatalogProductAttributeTierPriceInfoResponse was auto-generated
// from WSDL.
type OperationCatalogProductAttributeTierPriceInfoResponse struct {
	Result *CatalogProductTierPriceEntityArray `xml:"result,omitempty" json:"result,omitempty" yaml:"result,omitempty"`
}

// Operation wrapper for CatalogProductAttributeTierPriceUpdate.
// OperationCatalogProductAttributeTierPriceUpdateRequest was auto-generated
// from WSDL.
type OperationCatalogProductAttributeTierPriceUpdateRequest struct {
	SessionId      *string                             `xml:"sessionId,omitempty" json:"sessionId,omitempty" yaml:"sessionId,omitempty"`
	Product        *string                             `xml:"product,omitempty" json:"product,omitempty" yaml:"product,omitempty"`
	Tier_price     *CatalogProductTierPriceEntityArray `xml:"tier_price,omitempty" json:"tier_price,omitempty" yaml:"tier_price,omitempty"`
	IdentifierType *string                             `xml:"identifierType,omitempty" json:"identifierType,omitempty" yaml:"identifierType,omitempty"`
}

// Operation wrapper for CatalogProductAttributeTierPriceUpdate.
// OperationCatalogProductAttributeTierPriceUpdateResponse was
// auto-generated from WSDL.
type OperationCatalogProductAttributeTierPriceUpdateResponse struct {
	Result *int `xml:"result,omitempty" json:"result,omitempty" yaml:"result,omitempty"`
}

// Operation wrapper for CatalogProductAttributeTypes.
// OperationCatalogProductAttributeTypesRequest was auto-generated
// from WSDL.
type OperationCatalogProductAttributeTypesRequest struct {
	SessionId *string `xml:"sessionId,omitempty" json:"sessionId,omitempty" yaml:"sessionId,omitempty"`
}

// Operation wrapper for CatalogProductAttributeTypes.
// OperationCatalogProductAttributeTypesResponse was auto-generated
// from WSDL.
type OperationCatalogProductAttributeTypesResponse struct {
	Result *CatalogAttributeOptionEntityArray `xml:"result,omitempty" json:"result,omitempty" yaml:"result,omitempty"`
}

// Operation wrapper for CatalogProductAttributeUpdate.
// OperationCatalogProductAttributeUpdateRequest was auto-generated
// from WSDL.
type OperationCatalogProductAttributeUpdateRequest struct {
	SessionId *string                                `xml:"sessionId,omitempty" json:"sessionId,omitempty" yaml:"sessionId,omitempty"`
	Attribute *string                                `xml:"attribute,omitempty" json:"attribute,omitempty" yaml:"attribute,omitempty"`
	Data      *CatalogProductAttributeEntityToUpdate `xml:"data,omitempty" json:"data,omitempty" yaml:"data,omitempty"`
}

// Operation wrapper for CatalogProductAttributeUpdate.
// OperationCatalogProductAttributeUpdateResponse was auto-generated
// from WSDL.
type OperationCatalogProductAttributeUpdateResponse struct {
	Result *bool `xml:"result,omitempty" json:"result,omitempty" yaml:"result,omitempty"`
}

// Operation wrapper for CatalogProductCreate.
// OperationCatalogProductCreateRequest was auto-generated from
// WSDL.
type OperationCatalogProductCreateRequest struct {
	SessionId   *string                     `xml:"sessionId,omitempty" json:"sessionId,omitempty" yaml:"sessionId,omitempty"`
	Type        *string                     `xml:"type,omitempty" json:"type,omitempty" yaml:"type,omitempty"`
	Set         *string                     `xml:"set,omitempty" json:"set,omitempty" yaml:"set,omitempty"`
	Sku         *string                     `xml:"sku,omitempty" json:"sku,omitempty" yaml:"sku,omitempty"`
	ProductData *CatalogProductCreateEntity `xml:"productData,omitempty" json:"productData,omitempty" yaml:"productData,omitempty"`
	StoreView   *string                     `xml:"storeView,omitempty" json:"storeView,omitempty" yaml:"storeView,omitempty"`
}

// Operation wrapper for CatalogProductCreate.
// OperationCatalogProductCreateResponse was auto-generated from
// WSDL.
type OperationCatalogProductCreateResponse struct {
	Result *int `xml:"result,omitempty" json:"result,omitempty" yaml:"result,omitempty"`
}

// Operation wrapper for CatalogProductCurrentStore.
// OperationCatalogProductCurrentStoreRequest was auto-generated
// from WSDL.
type OperationCatalogProductCurrentStoreRequest struct {
	SessionId *string `xml:"sessionId,omitempty" json:"sessionId,omitempty" yaml:"sessionId,omitempty"`
	StoreView *string `xml:"storeView,omitempty" json:"storeView,omitempty" yaml:"storeView,omitempty"`
}

// Operation wrapper for CatalogProductCurrentStore.
// OperationCatalogProductCurrentStoreResponse was auto-generated
// from WSDL.
type OperationCatalogProductCurrentStoreResponse struct {
	StoreView *int `xml:"storeView,omitempty" json:"storeView,omitempty" yaml:"storeView,omitempty"`
}

// Operation wrapper for CatalogProductCustomOptionAdd.
// OperationCatalogProductCustomOptionAddRequest was auto-generated
// from WSDL.
type OperationCatalogProductCustomOptionAddRequest struct {
	SessionId *string                          `xml:"sessionId,omitempty" json:"sessionId,omitempty" yaml:"sessionId,omitempty"`
	ProductId *string                          `xml:"productId,omitempty" json:"productId,omitempty" yaml:"productId,omitempty"`
	Data      *CatalogProductCustomOptionToAdd `xml:"data,omitempty" json:"data,omitempty" yaml:"data,omitempty"`
	Store     *string                          `xml:"store,omitempty" json:"store,omitempty" yaml:"store,omitempty"`
}

// Operation wrapper for CatalogProductCustomOptionAdd.
// OperationCatalogProductCustomOptionAddResponse was auto-generated
// from WSDL.
type OperationCatalogProductCustomOptionAddResponse struct {
	Result *bool `xml:"result,omitempty" json:"result,omitempty" yaml:"result,omitempty"`
}

// Operation wrapper for CatalogProductCustomOptionInfo.
// OperationCatalogProductCustomOptionInfoRequest was auto-generated
// from WSDL.
type OperationCatalogProductCustomOptionInfoRequest struct {
	SessionId *string `xml:"sessionId,omitempty" json:"sessionId,omitempty" yaml:"sessionId,omitempty"`
	OptionId  *string `xml:"optionId,omitempty" json:"optionId,omitempty" yaml:"optionId,omitempty"`
	Store     *string `xml:"store,omitempty" json:"store,omitempty" yaml:"store,omitempty"`
}

// Operation wrapper for CatalogProductCustomOptionInfo.
// OperationCatalogProductCustomOptionInfoResponse was auto-generated
// from WSDL.
type OperationCatalogProductCustomOptionInfoResponse struct {
	Result *CatalogProductCustomOptionInfoEntity `xml:"result,omitempty" json:"result,omitempty" yaml:"result,omitempty"`
}

// Operation wrapper for CatalogProductCustomOptionList.
// OperationCatalogProductCustomOptionListRequest was auto-generated
// from WSDL.
type OperationCatalogProductCustomOptionListRequest struct {
	SessionId *string `xml:"sessionId,omitempty" json:"sessionId,omitempty" yaml:"sessionId,omitempty"`
	ProductId *string `xml:"productId,omitempty" json:"productId,omitempty" yaml:"productId,omitempty"`
	Store     *string `xml:"store,omitempty" json:"store,omitempty" yaml:"store,omitempty"`
}

// Operation wrapper for CatalogProductCustomOptionList.
// OperationCatalogProductCustomOptionListResponse was auto-generated
// from WSDL.
type OperationCatalogProductCustomOptionListResponse struct {
	Result *CatalogProductCustomOptionListArray `xml:"result,omitempty" json:"result,omitempty" yaml:"result,omitempty"`
}

// Operation wrapper for CatalogProductCustomOptionRemove.
// OperationCatalogProductCustomOptionRemoveRequest was auto-generated
// from WSDL.
type OperationCatalogProductCustomOptionRemoveRequest struct {
	SessionId *string `xml:"sessionId,omitempty" json:"sessionId,omitempty" yaml:"sessionId,omitempty"`
	OptionId  *string `xml:"optionId,omitempty" json:"optionId,omitempty" yaml:"optionId,omitempty"`
}

// Operation wrapper for CatalogProductCustomOptionRemove.
// OperationCatalogProductCustomOptionRemoveResponse was auto-generated
// from WSDL.
type OperationCatalogProductCustomOptionRemoveResponse struct {
	Result *bool `xml:"result,omitempty" json:"result,omitempty" yaml:"result,omitempty"`
}

// Operation wrapper for CatalogProductCustomOptionTypes.
// OperationCatalogProductCustomOptionTypesRequest was auto-generated
// from WSDL.
type OperationCatalogProductCustomOptionTypesRequest struct {
	SessionId *string `xml:"sessionId,omitempty" json:"sessionId,omitempty" yaml:"sessionId,omitempty"`
}

// Operation wrapper for CatalogProductCustomOptionTypes.
// OperationCatalogProductCustomOptionTypesResponse was auto-generated
// from WSDL.
type OperationCatalogProductCustomOptionTypesResponse struct {
	Result *CatalogProductCustomOptionTypesArray `xml:"result,omitempty" json:"result,omitempty" yaml:"result,omitempty"`
}

// Operation wrapper for CatalogProductCustomOptionUpdate.
// OperationCatalogProductCustomOptionUpdateRequest was auto-generated
// from WSDL.
type OperationCatalogProductCustomOptionUpdateRequest struct {
	SessionId *string                             `xml:"sessionId,omitempty" json:"sessionId,omitempty" yaml:"sessionId,omitempty"`
	OptionId  *string                             `xml:"optionId,omitempty" json:"optionId,omitempty" yaml:"optionId,omitempty"`
	Data      *CatalogProductCustomOptionToUpdate `xml:"data,omitempty" json:"data,omitempty" yaml:"data,omitempty"`
	Store     *string                             `xml:"store,omitempty" json:"store,omitempty" yaml:"store,omitempty"`
}

// Operation wrapper for CatalogProductCustomOptionUpdate.
// OperationCatalogProductCustomOptionUpdateResponse was auto-generated
// from WSDL.
type OperationCatalogProductCustomOptionUpdateResponse struct {
	Result *bool `xml:"result,omitempty" json:"result,omitempty" yaml:"result,omitempty"`
}

// Operation wrapper for CatalogProductCustomOptionValueAdd.
// OperationCatalogProductCustomOptionValueAddRequest was auto-generated
// from WSDL.
type OperationCatalogProductCustomOptionValueAddRequest struct {
	SessionId *string                                  `xml:"sessionId,omitempty" json:"sessionId,omitempty" yaml:"sessionId,omitempty"`
	OptionId  *string                                  `xml:"optionId,omitempty" json:"optionId,omitempty" yaml:"optionId,omitempty"`
	Data      *CatalogProductCustomOptionValueAddArray `xml:"data,omitempty" json:"data,omitempty" yaml:"data,omitempty"`
	Store     *string                                  `xml:"store,omitempty" json:"store,omitempty" yaml:"store,omitempty"`
}

// Operation wrapper for CatalogProductCustomOptionValueAdd.
// OperationCatalogProductCustomOptionValueAddResponse was auto-generated
// from WSDL.
type OperationCatalogProductCustomOptionValueAddResponse struct {
	Result *bool `xml:"result,omitempty" json:"result,omitempty" yaml:"result,omitempty"`
}

// Operation wrapper for CatalogProductCustomOptionValueInfo.
// OperationCatalogProductCustomOptionValueInfoRequest was auto-generated
// from WSDL.
type OperationCatalogProductCustomOptionValueInfoRequest struct {
	SessionId *string `xml:"sessionId,omitempty" json:"sessionId,omitempty" yaml:"sessionId,omitempty"`
	ValueId   *string `xml:"valueId,omitempty" json:"valueId,omitempty" yaml:"valueId,omitempty"`
	Store     *string `xml:"store,omitempty" json:"store,omitempty" yaml:"store,omitempty"`
}

// Operation wrapper for CatalogProductCustomOptionValueInfo.
// OperationCatalogProductCustomOptionValueInfoResponse was auto-generated
// from WSDL.
type OperationCatalogProductCustomOptionValueInfoResponse struct {
	Result *CatalogProductCustomOptionValueInfoEntity `xml:"result,omitempty" json:"result,omitempty" yaml:"result,omitempty"`
}

// Operation wrapper for CatalogProductCustomOptionValueList.
// OperationCatalogProductCustomOptionValueListRequest was auto-generated
// from WSDL.
type OperationCatalogProductCustomOptionValueListRequest struct {
	SessionId *string `xml:"sessionId,omitempty" json:"sessionId,omitempty" yaml:"sessionId,omitempty"`
	OptionId  *string `xml:"optionId,omitempty" json:"optionId,omitempty" yaml:"optionId,omitempty"`
	Store     *string `xml:"store,omitempty" json:"store,omitempty" yaml:"store,omitempty"`
}

// Operation wrapper for CatalogProductCustomOptionValueList.
// OperationCatalogProductCustomOptionValueListResponse was auto-generated
// from WSDL.
type OperationCatalogProductCustomOptionValueListResponse struct {
	Result *CatalogProductCustomOptionValueListArray `xml:"result,omitempty" json:"result,omitempty" yaml:"result,omitempty"`
}

// Operation wrapper for CatalogProductCustomOptionValueRemove.
// OperationCatalogProductCustomOptionValueRemoveRequest was auto-generated
// from WSDL.
type OperationCatalogProductCustomOptionValueRemoveRequest struct {
	SessionId *string `xml:"sessionId,omitempty" json:"sessionId,omitempty" yaml:"sessionId,omitempty"`
	ValueId   *string `xml:"valueId,omitempty" json:"valueId,omitempty" yaml:"valueId,omitempty"`
}

// Operation wrapper for CatalogProductCustomOptionValueRemove.
// OperationCatalogProductCustomOptionValueRemoveResponse was auto-generated
// from WSDL.
type OperationCatalogProductCustomOptionValueRemoveResponse struct {
	Result *bool `xml:"result,omitempty" json:"result,omitempty" yaml:"result,omitempty"`
}

// Operation wrapper for CatalogProductCustomOptionValueUpdate.
// OperationCatalogProductCustomOptionValueUpdateRequest was auto-generated
// from WSDL.
type OperationCatalogProductCustomOptionValueUpdateRequest struct {
	SessionId *string                                      `xml:"sessionId,omitempty" json:"sessionId,omitempty" yaml:"sessionId,omitempty"`
	ValueId   *string                                      `xml:"valueId,omitempty" json:"valueId,omitempty" yaml:"valueId,omitempty"`
	Data      *CatalogProductCustomOptionValueUpdateEntity `xml:"data,omitempty" json:"data,omitempty" yaml:"data,omitempty"`
	StoreId   *string                                      `xml:"storeId,omitempty" json:"storeId,omitempty" yaml:"storeId,omitempty"`
}

// Operation wrapper for CatalogProductCustomOptionValueUpdate.
// OperationCatalogProductCustomOptionValueUpdateResponse was auto-generated
// from WSDL.
type OperationCatalogProductCustomOptionValueUpdateResponse struct {
	Result *bool `xml:"result,omitempty" json:"result,omitempty" yaml:"result,omitempty"`
}

// Operation wrapper for CatalogProductDelete.
// OperationCatalogProductDeleteRequest was auto-generated from
// WSDL.
type OperationCatalogProductDeleteRequest struct {
	SessionId      *string `xml:"sessionId,omitempty" json:"sessionId,omitempty" yaml:"sessionId,omitempty"`
	Product        *string `xml:"product,omitempty" json:"product,omitempty" yaml:"product,omitempty"`
	IdentifierType *string `xml:"identifierType,omitempty" json:"identifierType,omitempty" yaml:"identifierType,omitempty"`
}

// Operation wrapper for CatalogProductDelete.
// OperationCatalogProductDeleteResponse was auto-generated from
// WSDL.
type OperationCatalogProductDeleteResponse struct {
	Result *int `xml:"result,omitempty" json:"result,omitempty" yaml:"result,omitempty"`
}

// Operation wrapper for CatalogProductDownloadableLinkAdd.
// OperationCatalogProductDownloadableLinkAddRequest was auto-generated
// from WSDL.
type OperationCatalogProductDownloadableLinkAddRequest struct {
	SessionId      *string                                  `xml:"sessionId,omitempty" json:"sessionId,omitempty" yaml:"sessionId,omitempty"`
	ProductId      *string                                  `xml:"productId,omitempty" json:"productId,omitempty" yaml:"productId,omitempty"`
	Resource       *CatalogProductDownloadableLinkAddEntity `xml:"resource,omitempty" json:"resource,omitempty" yaml:"resource,omitempty"`
	ResourceType   *string                                  `xml:"resourceType,omitempty" json:"resourceType,omitempty" yaml:"resourceType,omitempty"`
	Store          *string                                  `xml:"store,omitempty" json:"store,omitempty" yaml:"store,omitempty"`
	IdentifierType *string                                  `xml:"identifierType,omitempty" json:"identifierType,omitempty" yaml:"identifierType,omitempty"`
}

// Operation wrapper for CatalogProductDownloadableLinkAdd.
// OperationCatalogProductDownloadableLinkAddResponse was auto-generated
// from WSDL.
type OperationCatalogProductDownloadableLinkAddResponse struct {
	Respons *int `xml:"respons,omitempty" json:"respons,omitempty" yaml:"respons,omitempty"`
}

// Operation wrapper for CatalogProductDownloadableLinkList.
// OperationCatalogProductDownloadableLinkListRequest was auto-generated
// from WSDL.
type OperationCatalogProductDownloadableLinkListRequest struct {
	SessionId      *string `xml:"sessionId,omitempty" json:"sessionId,omitempty" yaml:"sessionId,omitempty"`
	ProductId      *string `xml:"productId,omitempty" json:"productId,omitempty" yaml:"productId,omitempty"`
	Store          *string `xml:"store,omitempty" json:"store,omitempty" yaml:"store,omitempty"`
	IdentifierType *string `xml:"identifierType,omitempty" json:"identifierType,omitempty" yaml:"identifierType,omitempty"`
}

// Operation wrapper for CatalogProductDownloadableLinkList.
// OperationCatalogProductDownloadableLinkListResponse was auto-generated
// from WSDL.
type OperationCatalogProductDownloadableLinkListResponse struct {
	Respons *CatalogProductDownloadableLinkInfoEntity `xml:"respons,omitempty" json:"respons,omitempty" yaml:"respons,omitempty"`
}

// Operation wrapper for CatalogProductDownloadableLinkRemove.
// OperationCatalogProductDownloadableLinkRemoveRequest was auto-generated
// from WSDL.
type OperationCatalogProductDownloadableLinkRemoveRequest struct {
	SessionId    *string `xml:"sessionId,omitempty" json:"sessionId,omitempty" yaml:"sessionId,omitempty"`
	LinkId       *string `xml:"linkId,omitempty" json:"linkId,omitempty" yaml:"linkId,omitempty"`
	ResourceType *string `xml:"resourceType,omitempty" json:"resourceType,omitempty" yaml:"resourceType,omitempty"`
}

// Operation wrapper for CatalogProductDownloadableLinkRemove.
// OperationCatalogProductDownloadableLinkRemoveResponse was auto-generated
// from WSDL.
type OperationCatalogProductDownloadableLinkRemoveResponse struct {
	Respons *bool `xml:"respons,omitempty" json:"respons,omitempty" yaml:"respons,omitempty"`
}

// Operation wrapper for CatalogProductGetSpecialPrice.
// OperationCatalogProductGetSpecialPriceRequest was auto-generated
// from WSDL.
type OperationCatalogProductGetSpecialPriceRequest struct {
	SessionId      *string `xml:"sessionId,omitempty" json:"sessionId,omitempty" yaml:"sessionId,omitempty"`
	Product        *string `xml:"product,omitempty" json:"product,omitempty" yaml:"product,omitempty"`
	StoreView      *string `xml:"storeView,omitempty" json:"storeView,omitempty" yaml:"storeView,omitempty"`
	IdentifierType *string `xml:"identifierType,omitempty" json:"identifierType,omitempty" yaml:"identifierType,omitempty"`
}

// Operation wrapper for CatalogProductGetSpecialPrice.
// OperationCatalogProductGetSpecialPriceResponse was auto-generated
// from WSDL.
type OperationCatalogProductGetSpecialPriceResponse struct {
	Result *CatalogProductSpecialPriceReturnEntity `xml:"result,omitempty" json:"result,omitempty" yaml:"result,omitempty"`
}

// Operation wrapper for CatalogProductInfo.
// OperationCatalogProductInfoRequest was auto-generated from WSDL.
type OperationCatalogProductInfoRequest struct {
	SessionId      *string                          `xml:"sessionId,omitempty" json:"sessionId,omitempty" yaml:"sessionId,omitempty"`
	ProductId      *string                          `xml:"productId,omitempty" json:"productId,omitempty" yaml:"productId,omitempty"`
	StoreView      *string                          `xml:"storeView,omitempty" json:"storeView,omitempty" yaml:"storeView,omitempty"`
	Attributes     *CatalogProductRequestAttributes `xml:"attributes,omitempty" json:"attributes,omitempty" yaml:"attributes,omitempty"`
	IdentifierType *string                          `xml:"identifierType,omitempty" json:"identifierType,omitempty" yaml:"identifierType,omitempty"`
}

// Operation wrapper for CatalogProductInfo.
// OperationCatalogProductInfoResponse was auto-generated from
// WSDL.
type OperationCatalogProductInfoResponse struct {
	Info *CatalogProductReturnEntity `xml:"info,omitempty" json:"info,omitempty" yaml:"info,omitempty"`
}

// Operation wrapper for CatalogProductLinkAssign.
// OperationCatalogProductLinkAssignRequest was auto-generated
// from WSDL.
type OperationCatalogProductLinkAssignRequest struct {
	SessionId      *string                   `xml:"sessionId,omitempty" json:"sessionId,omitempty" yaml:"sessionId,omitempty"`
	Type           *string                   `xml:"type,omitempty" json:"type,omitempty" yaml:"type,omitempty"`
	Product        *string                   `xml:"product,omitempty" json:"product,omitempty" yaml:"product,omitempty"`
	LinkedProduct  *string                   `xml:"linkedProduct,omitempty" json:"linkedProduct,omitempty" yaml:"linkedProduct,omitempty"`
	Data           *CatalogProductLinkEntity `xml:"data,omitempty" json:"data,omitempty" yaml:"data,omitempty"`
	IdentifierType *string                   `xml:"identifierType,omitempty" json:"identifierType,omitempty" yaml:"identifierType,omitempty"`
}

// Operation wrapper for CatalogProductLinkAssign.
// OperationCatalogProductLinkAssignResponse was auto-generated
// from WSDL.
type OperationCatalogProductLinkAssignResponse struct {
	Result *bool `xml:"result,omitempty" json:"result,omitempty" yaml:"result,omitempty"`
}

// Operation wrapper for CatalogProductLinkAttributes.
// OperationCatalogProductLinkAttributesRequest was auto-generated
// from WSDL.
type OperationCatalogProductLinkAttributesRequest struct {
	SessionId *string `xml:"sessionId,omitempty" json:"sessionId,omitempty" yaml:"sessionId,omitempty"`
	Type      *string `xml:"type,omitempty" json:"type,omitempty" yaml:"type,omitempty"`
}

// Operation wrapper for CatalogProductLinkAttributes.
// OperationCatalogProductLinkAttributesResponse was auto-generated
// from WSDL.
type OperationCatalogProductLinkAttributesResponse struct {
	Result *CatalogProductLinkAttributeEntityArray `xml:"result,omitempty" json:"result,omitempty" yaml:"result,omitempty"`
}

// Operation wrapper for CatalogProductLinkList.
// OperationCatalogProductLinkListRequest was auto-generated from
// WSDL.
type OperationCatalogProductLinkListRequest struct {
	SessionId      *string `xml:"sessionId,omitempty" json:"sessionId,omitempty" yaml:"sessionId,omitempty"`
	Type           *string `xml:"type,omitempty" json:"type,omitempty" yaml:"type,omitempty"`
	Product        *string `xml:"product,omitempty" json:"product,omitempty" yaml:"product,omitempty"`
	IdentifierType *string `xml:"identifierType,omitempty" json:"identifierType,omitempty" yaml:"identifierType,omitempty"`
}

// Operation wrapper for CatalogProductLinkList.
// OperationCatalogProductLinkListResponse was auto-generated from
// WSDL.
type OperationCatalogProductLinkListResponse struct {
	Result *CatalogProductLinkEntityArray `xml:"result,omitempty" json:"result,omitempty" yaml:"result,omitempty"`
}

// Operation wrapper for CatalogProductLinkRemove.
// OperationCatalogProductLinkRemoveRequest was auto-generated
// from WSDL.
type OperationCatalogProductLinkRemoveRequest struct {
	SessionId      *string `xml:"sessionId,omitempty" json:"sessionId,omitempty" yaml:"sessionId,omitempty"`
	Type           *string `xml:"type,omitempty" json:"type,omitempty" yaml:"type,omitempty"`
	Product        *string `xml:"product,omitempty" json:"product,omitempty" yaml:"product,omitempty"`
	LinkedProduct  *string `xml:"linkedProduct,omitempty" json:"linkedProduct,omitempty" yaml:"linkedProduct,omitempty"`
	IdentifierType *string `xml:"identifierType,omitempty" json:"identifierType,omitempty" yaml:"identifierType,omitempty"`
}

// Operation wrapper for CatalogProductLinkRemove.
// OperationCatalogProductLinkRemoveResponse was auto-generated
// from WSDL.
type OperationCatalogProductLinkRemoveResponse struct {
	Result *bool `xml:"result,omitempty" json:"result,omitempty" yaml:"result,omitempty"`
}

// Operation wrapper for CatalogProductLinkTypes.
// OperationCatalogProductLinkTypesRequest was auto-generated from
// WSDL.
type OperationCatalogProductLinkTypesRequest struct {
	SessionId *string `xml:"sessionId,omitempty" json:"sessionId,omitempty" yaml:"sessionId,omitempty"`
}

// Operation wrapper for CatalogProductLinkTypes.
// OperationCatalogProductLinkTypesResponse was auto-generated
// from WSDL.
type OperationCatalogProductLinkTypesResponse struct {
	Result *ArrayOfString `xml:"result,omitempty" json:"result,omitempty" yaml:"result,omitempty"`
}

// Operation wrapper for CatalogProductLinkUpdate.
// OperationCatalogProductLinkUpdateRequest was auto-generated
// from WSDL.
type OperationCatalogProductLinkUpdateRequest struct {
	SessionId      *string                   `xml:"sessionId,omitempty" json:"sessionId,omitempty" yaml:"sessionId,omitempty"`
	Type           *string                   `xml:"type,omitempty" json:"type,omitempty" yaml:"type,omitempty"`
	Product        *string                   `xml:"product,omitempty" json:"product,omitempty" yaml:"product,omitempty"`
	LinkedProduct  *string                   `xml:"linkedProduct,omitempty" json:"linkedProduct,omitempty" yaml:"linkedProduct,omitempty"`
	Data           *CatalogProductLinkEntity `xml:"data,omitempty" json:"data,omitempty" yaml:"data,omitempty"`
	IdentifierType *string                   `xml:"identifierType,omitempty" json:"identifierType,omitempty" yaml:"identifierType,omitempty"`
}

// Operation wrapper for CatalogProductLinkUpdate.
// OperationCatalogProductLinkUpdateResponse was auto-generated
// from WSDL.
type OperationCatalogProductLinkUpdateResponse struct {
	Result *bool `xml:"result,omitempty" json:"result,omitempty" yaml:"result,omitempty"`
}

// Operation wrapper for CatalogProductList.
// OperationCatalogProductListRequest was auto-generated from WSDL.
type OperationCatalogProductListRequest struct {
	SessionId *string  `xml:"sessionId,omitempty" json:"sessionId,omitempty" yaml:"sessionId,omitempty"`
	Filters   *Filters `xml:"filters,omitempty" json:"filters,omitempty" yaml:"filters,omitempty"`
	StoreView *string  `xml:"storeView,omitempty" json:"storeView,omitempty" yaml:"storeView,omitempty"`
}

// Operation wrapper for CatalogProductList.
// OperationCatalogProductListResponse was auto-generated from
// WSDL.
type OperationCatalogProductListResponse struct {
	StoreView *CatalogProductEntityArray `xml:"storeView,omitempty" json:"storeView,omitempty" yaml:"storeView,omitempty"`
}

// Operation wrapper for CatalogProductListOfAdditionalAttributes.
// OperationCatalogProductListOfAdditionalAttributesRequest was
// auto-generated from WSDL.
type OperationCatalogProductListOfAdditionalAttributesRequest struct {
	SessionId      *string `xml:"sessionId,omitempty" json:"sessionId,omitempty" yaml:"sessionId,omitempty"`
	ProductType    *string `xml:"productType,omitempty" json:"productType,omitempty" yaml:"productType,omitempty"`
	AttributeSetId *string `xml:"attributeSetId,omitempty" json:"attributeSetId,omitempty" yaml:"attributeSetId,omitempty"`
}

// Operation wrapper for CatalogProductListOfAdditionalAttributes.
// OperationCatalogProductListOfAdditionalAttributesResponse was
// auto-generated from WSDL.
type OperationCatalogProductListOfAdditionalAttributesResponse struct {
	Result *CatalogAttributeEntityArray `xml:"result,omitempty" json:"result,omitempty" yaml:"result,omitempty"`
}

// Operation wrapper for CatalogProductMultiUpdate.
// OperationCatalogProductMultiUpdateRequest was auto-generated
// from WSDL.
type OperationCatalogProductMultiUpdateRequest struct {
	SessionId      *string                          `xml:"sessionId,omitempty" json:"sessionId,omitempty" yaml:"sessionId,omitempty"`
	ProductIds     *ArrayOfString                   `xml:"productIds,omitempty" json:"productIds,omitempty" yaml:"productIds,omitempty"`
	ProductData    *CatalogProductCreateEntityArray `xml:"productData,omitempty" json:"productData,omitempty" yaml:"productData,omitempty"`
	Store          *string                          `xml:"store,omitempty" json:"store,omitempty" yaml:"store,omitempty"`
	IdentifierType *string                          `xml:"identifierType,omitempty" json:"identifierType,omitempty" yaml:"identifierType,omitempty"`
}

// Operation wrapper for CatalogProductMultiUpdate.
// OperationCatalogProductMultiUpdateResponse was auto-generated
// from WSDL.
type OperationCatalogProductMultiUpdateResponse struct {
	Result *bool `xml:"result,omitempty" json:"result,omitempty" yaml:"result,omitempty"`
}

// Operation wrapper for CatalogProductSetSpecialPrice.
// OperationCatalogProductSetSpecialPriceRequest was auto-generated
// from WSDL.
type OperationCatalogProductSetSpecialPriceRequest struct {
	SessionId      *string `xml:"sessionId,omitempty" json:"sessionId,omitempty" yaml:"sessionId,omitempty"`
	Product        *string `xml:"product,omitempty" json:"product,omitempty" yaml:"product,omitempty"`
	SpecialPrice   *string `xml:"specialPrice,omitempty" json:"specialPrice,omitempty" yaml:"specialPrice,omitempty"`
	FromDate       *string `xml:"fromDate,omitempty" json:"fromDate,omitempty" yaml:"fromDate,omitempty"`
	ToDate         *string `xml:"toDate,omitempty" json:"toDate,omitempty" yaml:"toDate,omitempty"`
	StoreView      *string `xml:"storeView,omitempty" json:"storeView,omitempty" yaml:"storeView,omitempty"`
	IdentifierType *string `xml:"identifierType,omitempty" json:"identifierType,omitempty" yaml:"identifierType,omitempty"`
}

// Operation wrapper for CatalogProductSetSpecialPrice.
// OperationCatalogProductSetSpecialPriceResponse was auto-generated
// from WSDL.
type OperationCatalogProductSetSpecialPriceResponse struct {
	Result *int `xml:"result,omitempty" json:"result,omitempty" yaml:"result,omitempty"`
}

// Operation wrapper for CatalogProductTagAdd.
// OperationCatalogProductTagAddRequest was auto-generated from
// WSDL.
type OperationCatalogProductTagAddRequest struct {
	SessionId *string                     `xml:"sessionId,omitempty" json:"sessionId,omitempty" yaml:"sessionId,omitempty"`
	Data      *CatalogProductTagAddEntity `xml:"data,omitempty" json:"data,omitempty" yaml:"data,omitempty"`
}

// Operation wrapper for CatalogProductTagAdd.
// OperationCatalogProductTagAddResponse was auto-generated from
// WSDL.
type OperationCatalogProductTagAddResponse struct {
	Result *AssociativeArray `xml:"result,omitempty" json:"result,omitempty" yaml:"result,omitempty"`
}

// Operation wrapper for CatalogProductTagInfo.
// OperationCatalogProductTagInfoRequest was auto-generated from
// WSDL.
type OperationCatalogProductTagInfoRequest struct {
	SessionId *string `xml:"sessionId,omitempty" json:"sessionId,omitempty" yaml:"sessionId,omitempty"`
	TagId     *string `xml:"tagId,omitempty" json:"tagId,omitempty" yaml:"tagId,omitempty"`
	Store     *string `xml:"store,omitempty" json:"store,omitempty" yaml:"store,omitempty"`
}

// Operation wrapper for CatalogProductTagInfo.
// OperationCatalogProductTagInfoResponse was auto-generated from
// WSDL.
type OperationCatalogProductTagInfoResponse struct {
	Result *CatalogProductTagInfoEntity `xml:"result,omitempty" json:"result,omitempty" yaml:"result,omitempty"`
}

// Operation wrapper for CatalogProductTagList.
// OperationCatalogProductTagListRequest was auto-generated from
// WSDL.
type OperationCatalogProductTagListRequest struct {
	SessionId *string `xml:"sessionId,omitempty" json:"sessionId,omitempty" yaml:"sessionId,omitempty"`
	ProductId *string `xml:"productId,omitempty" json:"productId,omitempty" yaml:"productId,omitempty"`
	Store     *string `xml:"store,omitempty" json:"store,omitempty" yaml:"store,omitempty"`
}

// Operation wrapper for CatalogProductTagList.
// OperationCatalogProductTagListResponse was auto-generated from
// WSDL.
type OperationCatalogProductTagListResponse struct {
	Result *CatalogProductTagListEntityArray `xml:"result,omitempty" json:"result,omitempty" yaml:"result,omitempty"`
}

// Operation wrapper for CatalogProductTagRemove.
// OperationCatalogProductTagRemoveRequest was auto-generated from
// WSDL.
type OperationCatalogProductTagRemoveRequest struct {
	SessionId *string `xml:"sessionId,omitempty" json:"sessionId,omitempty" yaml:"sessionId,omitempty"`
	TagId     *string `xml:"tagId,omitempty" json:"tagId,omitempty" yaml:"tagId,omitempty"`
}

// Operation wrapper for CatalogProductTagRemove.
// OperationCatalogProductTagRemoveResponse was auto-generated
// from WSDL.
type OperationCatalogProductTagRemoveResponse struct {
	Result *bool `xml:"result,omitempty" json:"result,omitempty" yaml:"result,omitempty"`
}

// Operation wrapper for CatalogProductTagUpdate.
// OperationCatalogProductTagUpdateRequest was auto-generated from
// WSDL.
type OperationCatalogProductTagUpdateRequest struct {
	SessionId *string                        `xml:"sessionId,omitempty" json:"sessionId,omitempty" yaml:"sessionId,omitempty"`
	TagId     *string                        `xml:"tagId,omitempty" json:"tagId,omitempty" yaml:"tagId,omitempty"`
	Data      *CatalogProductTagUpdateEntity `xml:"data,omitempty" json:"data,omitempty" yaml:"data,omitempty"`
	Store     *string                        `xml:"store,omitempty" json:"store,omitempty" yaml:"store,omitempty"`
}

// Operation wrapper for CatalogProductTagUpdate.
// OperationCatalogProductTagUpdateResponse was auto-generated
// from WSDL.
type OperationCatalogProductTagUpdateResponse struct {
	Result *bool `xml:"result,omitempty" json:"result,omitempty" yaml:"result,omitempty"`
}

// Operation wrapper for CatalogProductTypeList.
// OperationCatalogProductTypeListRequest was auto-generated from
// WSDL.
type OperationCatalogProductTypeListRequest struct {
	SessionId *string `xml:"sessionId,omitempty" json:"sessionId,omitempty" yaml:"sessionId,omitempty"`
}

// Operation wrapper for CatalogProductTypeList.
// OperationCatalogProductTypeListResponse was auto-generated from
// WSDL.
type OperationCatalogProductTypeListResponse struct {
	Result *CatalogProductTypeEntityArray `xml:"result,omitempty" json:"result,omitempty" yaml:"result,omitempty"`
}

// Operation wrapper for CatalogProductUpdate.
// OperationCatalogProductUpdateRequest was auto-generated from
// WSDL.
type OperationCatalogProductUpdateRequest struct {
	SessionId      *string                     `xml:"sessionId,omitempty" json:"sessionId,omitempty" yaml:"sessionId,omitempty"`
	Product        *string                     `xml:"product,omitempty" json:"product,omitempty" yaml:"product,omitempty"`
	ProductData    *CatalogProductCreateEntity `xml:"productData,omitempty" json:"productData,omitempty" yaml:"productData,omitempty"`
	StoreView      *string                     `xml:"storeView,omitempty" json:"storeView,omitempty" yaml:"storeView,omitempty"`
	IdentifierType *string                     `xml:"identifierType,omitempty" json:"identifierType,omitempty" yaml:"identifierType,omitempty"`
}

// Operation wrapper for CatalogProductUpdate.
// OperationCatalogProductUpdateResponse was auto-generated from
// WSDL.
type OperationCatalogProductUpdateResponse struct {
	Result *bool `xml:"result,omitempty" json:"result,omitempty" yaml:"result,omitempty"`
}

// Operation wrapper for CustomerAddressCreate.
// OperationCustomerAddressCreateRequest was auto-generated from
// WSDL.
type OperationCustomerAddressCreateRequest struct {
	SessionId   *string                      `xml:"sessionId,omitempty" json:"sessionId,omitempty" yaml:"sessionId,omitempty"`
	CustomerId  *int                         `xml:"customerId,omitempty" json:"customerId,omitempty" yaml:"customerId,omitempty"`
	AddressData *CustomerAddressEntityCreate `xml:"addressData,omitempty" json:"addressData,omitempty" yaml:"addressData,omitempty"`
}

// Operation wrapper for CustomerAddressCreate.
// OperationCustomerAddressCreateResponse was auto-generated from
// WSDL.
type OperationCustomerAddressCreateResponse struct {
	Result *int `xml:"result,omitempty" json:"result,omitempty" yaml:"result,omitempty"`
}

// Operation wrapper for CustomerAddressDelete.
// OperationCustomerAddressDeleteRequest was auto-generated from
// WSDL.
type OperationCustomerAddressDeleteRequest struct {
	SessionId *string `xml:"sessionId,omitempty" json:"sessionId,omitempty" yaml:"sessionId,omitempty"`
	AddressId *int    `xml:"addressId,omitempty" json:"addressId,omitempty" yaml:"addressId,omitempty"`
}

// Operation wrapper for CustomerAddressDelete.
// OperationCustomerAddressDeleteResponse was auto-generated from
// WSDL.
type OperationCustomerAddressDeleteResponse struct {
	Info *bool `xml:"info,omitempty" json:"info,omitempty" yaml:"info,omitempty"`
}

// Operation wrapper for CustomerAddressInfo.
// OperationCustomerAddressInfoRequest was auto-generated from
// WSDL.
type OperationCustomerAddressInfoRequest struct {
	SessionId *string `xml:"sessionId,omitempty" json:"sessionId,omitempty" yaml:"sessionId,omitempty"`
	AddressId *int    `xml:"addressId,omitempty" json:"addressId,omitempty" yaml:"addressId,omitempty"`
}

// Operation wrapper for CustomerAddressInfo.
// OperationCustomerAddressInfoResponse was auto-generated from
// WSDL.
type OperationCustomerAddressInfoResponse struct {
	Info *CustomerAddressEntityItem `xml:"info,omitempty" json:"info,omitempty" yaml:"info,omitempty"`
}

// Operation wrapper for CustomerAddressList.
// OperationCustomerAddressListRequest was auto-generated from
// WSDL.
type OperationCustomerAddressListRequest struct {
	SessionId  *string `xml:"sessionId,omitempty" json:"sessionId,omitempty" yaml:"sessionId,omitempty"`
	CustomerId *int    `xml:"customerId,omitempty" json:"customerId,omitempty" yaml:"customerId,omitempty"`
}

// Operation wrapper for CustomerAddressList.
// OperationCustomerAddressListResponse was auto-generated from
// WSDL.
type OperationCustomerAddressListResponse struct {
	Result *CustomerAddressEntityArray `xml:"result,omitempty" json:"result,omitempty" yaml:"result,omitempty"`
}

// Operation wrapper for CustomerAddressUpdate.
// OperationCustomerAddressUpdateRequest was auto-generated from
// WSDL.
type OperationCustomerAddressUpdateRequest struct {
	SessionId   *string                      `xml:"sessionId,omitempty" json:"sessionId,omitempty" yaml:"sessionId,omitempty"`
	AddressId   *int                         `xml:"addressId,omitempty" json:"addressId,omitempty" yaml:"addressId,omitempty"`
	AddressData *CustomerAddressEntityCreate `xml:"addressData,omitempty" json:"addressData,omitempty" yaml:"addressData,omitempty"`
}

// Operation wrapper for CustomerAddressUpdate.
// OperationCustomerAddressUpdateResponse was auto-generated from
// WSDL.
type OperationCustomerAddressUpdateResponse struct {
	Info *bool `xml:"info,omitempty" json:"info,omitempty" yaml:"info,omitempty"`
}

// Operation wrapper for CustomerCustomerCreate.
// OperationCustomerCustomerCreateRequest was auto-generated from
// WSDL.
type OperationCustomerCustomerCreateRequest struct {
	SessionId    *string                         `xml:"sessionId,omitempty" json:"sessionId,omitempty" yaml:"sessionId,omitempty"`
	CustomerData *CustomerCustomerEntityToCreate `xml:"customerData,omitempty" json:"customerData,omitempty" yaml:"customerData,omitempty"`
}

// Operation wrapper for CustomerCustomerCreate.
// OperationCustomerCustomerCreateResponse was auto-generated from
// WSDL.
type OperationCustomerCustomerCreateResponse struct {
	Result *int `xml:"result,omitempty" json:"result,omitempty" yaml:"result,omitempty"`
}

// Operation wrapper for CustomerCustomerDelete.
// OperationCustomerCustomerDeleteRequest was auto-generated from
// WSDL.
type OperationCustomerCustomerDeleteRequest struct {
	SessionId  *string `xml:"sessionId,omitempty" json:"sessionId,omitempty" yaml:"sessionId,omitempty"`
	CustomerId *int    `xml:"customerId,omitempty" json:"customerId,omitempty" yaml:"customerId,omitempty"`
}

// Operation wrapper for CustomerCustomerDelete.
// OperationCustomerCustomerDeleteResponse was auto-generated from
// WSDL.
type OperationCustomerCustomerDeleteResponse struct {
	Result *bool `xml:"result,omitempty" json:"result,omitempty" yaml:"result,omitempty"`
}

// Operation wrapper for CustomerCustomerInfo.
// OperationCustomerCustomerInfoRequest was auto-generated from
// WSDL.
type OperationCustomerCustomerInfoRequest struct {
	SessionId  *string        `xml:"sessionId,omitempty" json:"sessionId,omitempty" yaml:"sessionId,omitempty"`
	CustomerId *int           `xml:"customerId,omitempty" json:"customerId,omitempty" yaml:"customerId,omitempty"`
	Attributes *ArrayOfString `xml:"attributes,omitempty" json:"attributes,omitempty" yaml:"attributes,omitempty"`
}

// Operation wrapper for CustomerCustomerInfo.
// OperationCustomerCustomerInfoResponse was auto-generated from
// WSDL.
type OperationCustomerCustomerInfoResponse struct {
	CustomerInfo *CustomerCustomerEntity `xml:"customerInfo,omitempty" json:"customerInfo,omitempty" yaml:"customerInfo,omitempty"`
}

// Operation wrapper for CustomerCustomerList.
// OperationCustomerCustomerListRequest was auto-generated from
// WSDL.
type OperationCustomerCustomerListRequest struct {
	SessionId *string  `xml:"sessionId,omitempty" json:"sessionId,omitempty" yaml:"sessionId,omitempty"`
	Filters   *Filters `xml:"filters,omitempty" json:"filters,omitempty" yaml:"filters,omitempty"`
}

// Operation wrapper for CustomerCustomerList.
// OperationCustomerCustomerListResponse was auto-generated from
// WSDL.
type OperationCustomerCustomerListResponse struct {
	StoreView *CustomerCustomerEntityArray `xml:"storeView,omitempty" json:"storeView,omitempty" yaml:"storeView,omitempty"`
}

// Operation wrapper for CustomerCustomerUpdate.
// OperationCustomerCustomerUpdateRequest was auto-generated from
// WSDL.
type OperationCustomerCustomerUpdateRequest struct {
	SessionId    *string                         `xml:"sessionId,omitempty" json:"sessionId,omitempty" yaml:"sessionId,omitempty"`
	CustomerId   *int                            `xml:"customerId,omitempty" json:"customerId,omitempty" yaml:"customerId,omitempty"`
	CustomerData *CustomerCustomerEntityToCreate `xml:"customerData,omitempty" json:"customerData,omitempty" yaml:"customerData,omitempty"`
}

// Operation wrapper for CustomerCustomerUpdate.
// OperationCustomerCustomerUpdateResponse was auto-generated from
// WSDL.
type OperationCustomerCustomerUpdateResponse struct {
	Result *bool `xml:"result,omitempty" json:"result,omitempty" yaml:"result,omitempty"`
}

// Operation wrapper for CustomerGroupList.
// OperationCustomerGroupListRequest was auto-generated from WSDL.
type OperationCustomerGroupListRequest struct {
	SessionId *string `xml:"sessionId,omitempty" json:"sessionId,omitempty" yaml:"sessionId,omitempty"`
}

// Operation wrapper for CustomerGroupList.
// OperationCustomerGroupListResponse was auto-generated from WSDL.
type OperationCustomerGroupListResponse struct {
	Result *CustomerGroupEntityArray `xml:"result,omitempty" json:"result,omitempty" yaml:"result,omitempty"`
}

// Operation wrapper for DirectoryCountryList.
// OperationDirectoryCountryListRequest was auto-generated from
// WSDL.
type OperationDirectoryCountryListRequest struct {
	SessionId *string `xml:"sessionId,omitempty" json:"sessionId,omitempty" yaml:"sessionId,omitempty"`
}

// Operation wrapper for DirectoryCountryList.
// OperationDirectoryCountryListResponse was auto-generated from
// WSDL.
type OperationDirectoryCountryListResponse struct {
	Countries *DirectoryCountryEntityArray `xml:"countries,omitempty" json:"countries,omitempty" yaml:"countries,omitempty"`
}

// Operation wrapper for DirectoryRegionList.
// OperationDirectoryRegionListRequest was auto-generated from
// WSDL.
type OperationDirectoryRegionListRequest struct {
	SessionId *string `xml:"sessionId,omitempty" json:"sessionId,omitempty" yaml:"sessionId,omitempty"`
	Country   *string `xml:"country,omitempty" json:"country,omitempty" yaml:"country,omitempty"`
}

// Operation wrapper for DirectoryRegionList.
// OperationDirectoryRegionListResponse was auto-generated from
// WSDL.
type OperationDirectoryRegionListResponse struct {
	Countries *DirectoryRegionEntityArray `xml:"countries,omitempty" json:"countries,omitempty" yaml:"countries,omitempty"`
}

// Operation wrapper for EndSession.
// OperationEndSession was auto-generated from WSDL.
type OperationEndSession struct {
	SessionId *string `xml:"sessionId,omitempty" json:"sessionId,omitempty" yaml:"sessionId,omitempty"`
}

// Operation wrapper for EndSession.
// OperationEndSessionResponse was auto-generated from WSDL.
type OperationEndSessionResponse struct {
	EndSessionReturn *bool `xml:"endSessionReturn,omitempty" json:"endSessionReturn,omitempty" yaml:"endSessionReturn,omitempty"`
}

// Operation wrapper for GiftMessageSetForQuote.
// OperationGiftMessageForQuoteRequest was auto-generated from
// WSDL.
type OperationGiftMessageForQuoteRequest struct {
	SessionId   *string            `xml:"sessionId,omitempty" json:"sessionId,omitempty" yaml:"sessionId,omitempty"`
	QuoteId     *string            `xml:"quoteId,omitempty" json:"quoteId,omitempty" yaml:"quoteId,omitempty"`
	GiftMessage *GiftMessageEntity `xml:"giftMessage,omitempty" json:"giftMessage,omitempty" yaml:"giftMessage,omitempty"`
	StoreId     *string            `xml:"storeId,omitempty" json:"storeId,omitempty" yaml:"storeId,omitempty"`
}

// Operation wrapper for GiftMessageSetForQuote.
// OperationGiftMessageForQuoteResponse was auto-generated from
// WSDL.
type OperationGiftMessageForQuoteResponse struct {
	Result *GiftMessageResponse `xml:"result,omitempty" json:"result,omitempty" yaml:"result,omitempty"`
}

// Operation wrapper for GiftMessageSetForQuoteItem.
// OperationGiftMessageForQuoteItemRequest was auto-generated from
// WSDL.
type OperationGiftMessageForQuoteItemRequest struct {
	SessionId   *string            `xml:"sessionId,omitempty" json:"sessionId,omitempty" yaml:"sessionId,omitempty"`
	QuoteItemId *string            `xml:"quoteItemId,omitempty" json:"quoteItemId,omitempty" yaml:"quoteItemId,omitempty"`
	GiftMessage *GiftMessageEntity `xml:"giftMessage,omitempty" json:"giftMessage,omitempty" yaml:"giftMessage,omitempty"`
	StoreId     *string            `xml:"storeId,omitempty" json:"storeId,omitempty" yaml:"storeId,omitempty"`
}

// Operation wrapper for GiftMessageSetForQuoteItem.
// OperationGiftMessageForQuoteItemResponse was auto-generated
// from WSDL.
type OperationGiftMessageForQuoteItemResponse struct {
	Result *GiftMessageResponse `xml:"result,omitempty" json:"result,omitempty" yaml:"result,omitempty"`
}

// Operation wrapper for GiftMessageSetForQuoteProduct.
// OperationGiftMessageForQuoteProductRequest was auto-generated
// from WSDL.
type OperationGiftMessageForQuoteProductRequest struct {
	SessionId           *string                                    `xml:"sessionId,omitempty" json:"sessionId,omitempty" yaml:"sessionId,omitempty"`
	QuoteId             *string                                    `xml:"quoteId,omitempty" json:"quoteId,omitempty" yaml:"quoteId,omitempty"`
	ProductsAndMessages *GiftMessageAssociativeProductsEntityArray `xml:"productsAndMessages,omitempty" json:"productsAndMessages,omitempty" yaml:"productsAndMessages,omitempty"`
	StoreId             *string                                    `xml:"storeId,omitempty" json:"storeId,omitempty" yaml:"storeId,omitempty"`
}

// Operation wrapper for GiftMessageSetForQuoteProduct.
// OperationGiftMessageForQuoteProductResponse was auto-generated
// from WSDL.
type OperationGiftMessageForQuoteProductResponse struct {
	Result *GiftMessageResponseArray `xml:"result,omitempty" json:"result,omitempty" yaml:"result,omitempty"`
}

// Operation wrapper for GlobalFaults.
// OperationGlobalFaults was auto-generated from WSDL.
type OperationGlobalFaults struct {
	SessionId *string `xml:"sessionId,omitempty" json:"sessionId,omitempty" yaml:"sessionId,omitempty"`
}

// Operation wrapper for GlobalFaults.
// OperationGlobalFaultsResponse was auto-generated from WSDL.
type OperationGlobalFaultsResponse struct {
	GlobalFaultsReturn *ArrayOfExistsFaltures `xml:"globalFaultsReturn,omitempty" json:"globalFaultsReturn,omitempty" yaml:"globalFaultsReturn,omitempty"`
}

// Operation wrapper for Login.
// OperationLogin was auto-generated from WSDL.
type OperationLogin struct {
	Username *string `xml:"username,omitempty" json:"username,omitempty" yaml:"username,omitempty"`
	ApiKey   *string `xml:"apiKey,omitempty" json:"apiKey,omitempty" yaml:"apiKey,omitempty"`
}

// Operation wrapper for Login.
// OperationLoginResponse was auto-generated from WSDL.
type OperationLoginResponse struct {
	LoginReturn *string `xml:"loginReturn,omitempty" json:"loginReturn,omitempty" yaml:"loginReturn,omitempty"`
}

// Operation wrapper for MagentoInfo.
// OperationMagentoInfoRequest was auto-generated from WSDL.
type OperationMagentoInfoRequest struct {
	SessionId *string `xml:"sessionId,omitempty" json:"sessionId,omitempty" yaml:"sessionId,omitempty"`
}

// Operation wrapper for MagentoInfo.
// OperationMagentoInfoResponse was auto-generated from WSDL.
type OperationMagentoInfoResponse struct {
	Info *MagentoInfoEntity `xml:"info,omitempty" json:"info,omitempty" yaml:"info,omitempty"`
}

// Operation wrapper for ResourceFaults.
// OperationResourceFaults was auto-generated from WSDL.
type OperationResourceFaults struct {
	ResourceName *string `xml:"resourceName,omitempty" json:"resourceName,omitempty" yaml:"resourceName,omitempty"`
	SessionId    *string `xml:"sessionId,omitempty" json:"sessionId,omitempty" yaml:"sessionId,omitempty"`
}

// Operation wrapper for ResourceFaults.
// OperationResourceFaultsResponse was auto-generated from WSDL.
type OperationResourceFaultsResponse struct {
	ResourceFaultsReturn *ArrayOfExistsFaltures `xml:"resourceFaultsReturn,omitempty" json:"resourceFaultsReturn,omitempty" yaml:"resourceFaultsReturn,omitempty"`
}

// Operation wrapper for Resources.
// OperationResourcesRequest was auto-generated from WSDL.
type OperationResourcesRequest struct {
	SessionId *string `xml:"sessionId,omitempty" json:"sessionId,omitempty" yaml:"sessionId,omitempty"`
}

// Operation wrapper for Resources.
// OperationResourcesResponse was auto-generated from WSDL.
type OperationResourcesResponse struct {
	ResourcesReturn *ArrayOfApis `xml:"resourcesReturn,omitempty" json:"resourcesReturn,omitempty" yaml:"resourcesReturn,omitempty"`
}

// Operation wrapper for SalesOrderAddComment.
// OperationSalesOrderAddCommentRequest was auto-generated from
// WSDL.
type OperationSalesOrderAddCommentRequest struct {
	SessionId        *string `xml:"sessionId,omitempty" json:"sessionId,omitempty" yaml:"sessionId,omitempty"`
	OrderIncrementId *string `xml:"orderIncrementId,omitempty" json:"orderIncrementId,omitempty" yaml:"orderIncrementId,omitempty"`
	Status           *string `xml:"status,omitempty" json:"status,omitempty" yaml:"status,omitempty"`
	Comment          *string `xml:"comment,omitempty" json:"comment,omitempty" yaml:"comment,omitempty"`
	Notify           *string `xml:"notify,omitempty" json:"notify,omitempty" yaml:"notify,omitempty"`
}

// Operation wrapper for SalesOrderAddComment.
// OperationSalesOrderAddCommentResponse was auto-generated from
// WSDL.
type OperationSalesOrderAddCommentResponse struct {
	Result *bool `xml:"result,omitempty" json:"result,omitempty" yaml:"result,omitempty"`
}

// Operation wrapper for SalesOrderCancel.
// OperationSalesOrderCancelRequest was auto-generated from WSDL.
type OperationSalesOrderCancelRequest struct {
	SessionId        *string `xml:"sessionId,omitempty" json:"sessionId,omitempty" yaml:"sessionId,omitempty"`
	OrderIncrementId *string `xml:"orderIncrementId,omitempty" json:"orderIncrementId,omitempty" yaml:"orderIncrementId,omitempty"`
}

// Operation wrapper for SalesOrderCancel.
// OperationSalesOrderCancelResponse was auto-generated from WSDL.
type OperationSalesOrderCancelResponse struct {
	Result *bool `xml:"result,omitempty" json:"result,omitempty" yaml:"result,omitempty"`
}

// Operation wrapper for SalesOrderCreditmemoAddComment.
// OperationSalesOrderCreditmemoAddCommentRequest was auto-generated
// from WSDL.
type OperationSalesOrderCreditmemoAddCommentRequest struct {
	SessionId             *string `xml:"sessionId,omitempty" json:"sessionId,omitempty" yaml:"sessionId,omitempty"`
	CreditmemoIncrementId *string `xml:"creditmemoIncrementId,omitempty" json:"creditmemoIncrementId,omitempty" yaml:"creditmemoIncrementId,omitempty"`
	Comment               *string `xml:"comment,omitempty" json:"comment,omitempty" yaml:"comment,omitempty"`
	NotifyCustomer        *int    `xml:"notifyCustomer,omitempty" json:"notifyCustomer,omitempty" yaml:"notifyCustomer,omitempty"`
	IncludeComment        *int    `xml:"includeComment,omitempty" json:"includeComment,omitempty" yaml:"includeComment,omitempty"`
}

// Operation wrapper for SalesOrderCreditmemoAddComment.
// OperationSalesOrderCreditmemoAddCommentResponse was auto-generated
// from WSDL.
type OperationSalesOrderCreditmemoAddCommentResponse struct {
	Result *bool `xml:"result,omitempty" json:"result,omitempty" yaml:"result,omitempty"`
}

// Operation wrapper for SalesOrderCreditmemoCancel.
// OperationSalesOrderCreditmemoCancelRequest was auto-generated
// from WSDL.
type OperationSalesOrderCreditmemoCancelRequest struct {
	SessionId             *string `xml:"sessionId,omitempty" json:"sessionId,omitempty" yaml:"sessionId,omitempty"`
	CreditmemoIncrementId *string `xml:"creditmemoIncrementId,omitempty" json:"creditmemoIncrementId,omitempty" yaml:"creditmemoIncrementId,omitempty"`
}

// Operation wrapper for SalesOrderCreditmemoCancel.
// OperationSalesOrderCreditmemoCancelResponse was auto-generated
// from WSDL.
type OperationSalesOrderCreditmemoCancelResponse struct {
	Result *bool `xml:"result,omitempty" json:"result,omitempty" yaml:"result,omitempty"`
}

// Operation wrapper for SalesOrderCreditmemoCreate.
// OperationSalesOrderCreditmemoCreateRequest was auto-generated
// from WSDL.
type OperationSalesOrderCreditmemoCreateRequest struct {
	SessionId                 *string                   `xml:"sessionId,omitempty" json:"sessionId,omitempty" yaml:"sessionId,omitempty"`
	OrderIncrementId          *string                   `xml:"orderIncrementId,omitempty" json:"orderIncrementId,omitempty" yaml:"orderIncrementId,omitempty"`
	CreditmemoData            *SalesOrderCreditmemoData `xml:"creditmemoData,omitempty" json:"creditmemoData,omitempty" yaml:"creditmemoData,omitempty"`
	Comment                   *string                   `xml:"comment,omitempty" json:"comment,omitempty" yaml:"comment,omitempty"`
	NotifyCustomer            *int                      `xml:"notifyCustomer,omitempty" json:"notifyCustomer,omitempty" yaml:"notifyCustomer,omitempty"`
	IncludeComment            *int                      `xml:"includeComment,omitempty" json:"includeComment,omitempty" yaml:"includeComment,omitempty"`
	RefundToStoreCreditAmount *string                   `xml:"refundToStoreCreditAmount,omitempty" json:"refundToStoreCreditAmount,omitempty" yaml:"refundToStoreCreditAmount,omitempty"`
}

// Operation wrapper for SalesOrderCreditmemoCreate.
// OperationSalesOrderCreditmemoCreateResponse was auto-generated
// from WSDL.
type OperationSalesOrderCreditmemoCreateResponse struct {
	Result *string `xml:"result,omitempty" json:"result,omitempty" yaml:"result,omitempty"`
}

// Operation wrapper for SalesOrderCreditmemoInfo.
// OperationSalesOrderCreditmemoInfoRequest was auto-generated
// from WSDL.
type OperationSalesOrderCreditmemoInfoRequest struct {
	SessionId             *string `xml:"sessionId,omitempty" json:"sessionId,omitempty" yaml:"sessionId,omitempty"`
	CreditmemoIncrementId *string `xml:"creditmemoIncrementId,omitempty" json:"creditmemoIncrementId,omitempty" yaml:"creditmemoIncrementId,omitempty"`
}

// Operation wrapper for SalesOrderCreditmemoInfo.
// OperationSalesOrderCreditmemoInfoResponse was auto-generated
// from WSDL.
type OperationSalesOrderCreditmemoInfoResponse struct {
	Result *SalesOrderCreditmemoEntity `xml:"result,omitempty" json:"result,omitempty" yaml:"result,omitempty"`
}

// Operation wrapper for SalesOrderCreditmemoList.
// OperationSalesOrderCreditmemoListRequest was auto-generated
// from WSDL.
type OperationSalesOrderCreditmemoListRequest struct {
	SessionId *string  `xml:"sessionId,omitempty" json:"sessionId,omitempty" yaml:"sessionId,omitempty"`
	Filters   *Filters `xml:"filters,omitempty" json:"filters,omitempty" yaml:"filters,omitempty"`
}

// Operation wrapper for SalesOrderCreditmemoList.
// OperationSalesOrderCreditmemoListResponse was auto-generated
// from WSDL.
type OperationSalesOrderCreditmemoListResponse struct {
	Result *SalesOrderCreditmemoEntityArray `xml:"result,omitempty" json:"result,omitempty" yaml:"result,omitempty"`
}

// Operation wrapper for SalesOrderHold.
// OperationSalesOrderHoldRequest was auto-generated from WSDL.
type OperationSalesOrderHoldRequest struct {
	SessionId        *string `xml:"sessionId,omitempty" json:"sessionId,omitempty" yaml:"sessionId,omitempty"`
	OrderIncrementId *string `xml:"orderIncrementId,omitempty" json:"orderIncrementId,omitempty" yaml:"orderIncrementId,omitempty"`
}

// Operation wrapper for SalesOrderHold.
// OperationSalesOrderHoldResponse was auto-generated from WSDL.
type OperationSalesOrderHoldResponse struct {
	Result *bool `xml:"result,omitempty" json:"result,omitempty" yaml:"result,omitempty"`
}

// Operation wrapper for SalesOrderInfo.
// OperationSalesOrderInfoRequest was auto-generated from WSDL.
type OperationSalesOrderInfoRequest struct {
	SessionId        *string `xml:"sessionId,omitempty" json:"sessionId,omitempty" yaml:"sessionId,omitempty"`
	OrderIncrementId *string `xml:"orderIncrementId,omitempty" json:"orderIncrementId,omitempty" yaml:"orderIncrementId,omitempty"`
}

// Operation wrapper for SalesOrderInfo.
// OperationSalesOrderInfoResponse was auto-generated from WSDL.
type OperationSalesOrderInfoResponse struct {
	Result *SalesOrderEntity `xml:"result,omitempty" json:"result,omitempty" yaml:"result,omitempty"`
}

// Operation wrapper for SalesOrderInvoiceAddComment.
// OperationSalesOrderInvoiceAddCommentRequest was auto-generated
// from WSDL.
type OperationSalesOrderInvoiceAddCommentRequest struct {
	SessionId          *string `xml:"sessionId,omitempty" json:"sessionId,omitempty" yaml:"sessionId,omitempty"`
	InvoiceIncrementId *string `xml:"invoiceIncrementId,omitempty" json:"invoiceIncrementId,omitempty" yaml:"invoiceIncrementId,omitempty"`
	Comment            *string `xml:"comment,omitempty" json:"comment,omitempty" yaml:"comment,omitempty"`
	Email              *string `xml:"email,omitempty" json:"email,omitempty" yaml:"email,omitempty"`
	IncludeComment     *string `xml:"includeComment,omitempty" json:"includeComment,omitempty" yaml:"includeComment,omitempty"`
}

// Operation wrapper for SalesOrderInvoiceAddComment.
// OperationSalesOrderInvoiceAddCommentResponse was auto-generated
// from WSDL.
type OperationSalesOrderInvoiceAddCommentResponse struct {
	Result *bool `xml:"result,omitempty" json:"result,omitempty" yaml:"result,omitempty"`
}

// Operation wrapper for SalesOrderInvoiceCancel.
// OperationSalesOrderInvoiceCancelRequest was auto-generated from
// WSDL.
type OperationSalesOrderInvoiceCancelRequest struct {
	SessionId          *string `xml:"sessionId,omitempty" json:"sessionId,omitempty" yaml:"sessionId,omitempty"`
	InvoiceIncrementId *string `xml:"invoiceIncrementId,omitempty" json:"invoiceIncrementId,omitempty" yaml:"invoiceIncrementId,omitempty"`
}

// Operation wrapper for SalesOrderInvoiceCancel.
// OperationSalesOrderInvoiceCancelResponse was auto-generated
// from WSDL.
type OperationSalesOrderInvoiceCancelResponse struct {
	Result *bool `xml:"result,omitempty" json:"result,omitempty" yaml:"result,omitempty"`
}

// Operation wrapper for SalesOrderInvoiceCapture.
// OperationSalesOrderInvoiceCaptureRequest was auto-generated
// from WSDL.
type OperationSalesOrderInvoiceCaptureRequest struct {
	SessionId          *string `xml:"sessionId,omitempty" json:"sessionId,omitempty" yaml:"sessionId,omitempty"`
	InvoiceIncrementId *string `xml:"invoiceIncrementId,omitempty" json:"invoiceIncrementId,omitempty" yaml:"invoiceIncrementId,omitempty"`
}

// Operation wrapper for SalesOrderInvoiceCapture.
// OperationSalesOrderInvoiceCaptureResponse was auto-generated
// from WSDL.
type OperationSalesOrderInvoiceCaptureResponse struct {
	Result *bool `xml:"result,omitempty" json:"result,omitempty" yaml:"result,omitempty"`
}

// Operation wrapper for SalesOrderInvoiceCreate.
// OperationSalesOrderInvoiceCreateRequest was auto-generated from
// WSDL.
type OperationSalesOrderInvoiceCreateRequest struct {
	SessionId          *string              `xml:"sessionId,omitempty" json:"sessionId,omitempty" yaml:"sessionId,omitempty"`
	InvoiceIncrementId *string              `xml:"invoiceIncrementId,omitempty" json:"invoiceIncrementId,omitempty" yaml:"invoiceIncrementId,omitempty"`
	ItemsQty           *OrderItemIdQtyArray `xml:"itemsQty,omitempty" json:"itemsQty,omitempty" yaml:"itemsQty,omitempty"`
	Comment            *string              `xml:"comment,omitempty" json:"comment,omitempty" yaml:"comment,omitempty"`
	Email              *string              `xml:"email,omitempty" json:"email,omitempty" yaml:"email,omitempty"`
	IncludeComment     *string              `xml:"includeComment,omitempty" json:"includeComment,omitempty" yaml:"includeComment,omitempty"`
}

// Operation wrapper for SalesOrderInvoiceCreate.
// OperationSalesOrderInvoiceCreateResponse was auto-generated
// from WSDL.
type OperationSalesOrderInvoiceCreateResponse struct {
	Result *string `xml:"result,omitempty" json:"result,omitempty" yaml:"result,omitempty"`
}

// Operation wrapper for SalesOrderInvoiceInfo.
// OperationSalesOrderInvoiceInfoRequest was auto-generated from
// WSDL.
type OperationSalesOrderInvoiceInfoRequest struct {
	SessionId          *string `xml:"sessionId,omitempty" json:"sessionId,omitempty" yaml:"sessionId,omitempty"`
	InvoiceIncrementId *string `xml:"invoiceIncrementId,omitempty" json:"invoiceIncrementId,omitempty" yaml:"invoiceIncrementId,omitempty"`
}

// Operation wrapper for SalesOrderInvoiceInfo.
// OperationSalesOrderInvoiceInfoResponse was auto-generated from
// WSDL.
type OperationSalesOrderInvoiceInfoResponse struct {
	Result *SalesOrderInvoiceEntity `xml:"result,omitempty" json:"result,omitempty" yaml:"result,omitempty"`
}

// Operation wrapper for SalesOrderInvoiceList.
// OperationSalesOrderInvoiceListRequest was auto-generated from
// WSDL.
type OperationSalesOrderInvoiceListRequest struct {
	SessionId *string  `xml:"sessionId,omitempty" json:"sessionId,omitempty" yaml:"sessionId,omitempty"`
	Filters   *Filters `xml:"filters,omitempty" json:"filters,omitempty" yaml:"filters,omitempty"`
}

// Operation wrapper for SalesOrderInvoiceList.
// OperationSalesOrderInvoiceListResponse was auto-generated from
// WSDL.
type OperationSalesOrderInvoiceListResponse struct {
	Result *SalesOrderInvoiceEntityArray `xml:"result,omitempty" json:"result,omitempty" yaml:"result,omitempty"`
}

// Operation wrapper for SalesOrderInvoiceVoid.
// OperationSalesOrderInvoiceVoidRequest was auto-generated from
// WSDL.
type OperationSalesOrderInvoiceVoidRequest struct {
	SessionId          *string `xml:"sessionId,omitempty" json:"sessionId,omitempty" yaml:"sessionId,omitempty"`
	InvoiceIncrementId *string `xml:"invoiceIncrementId,omitempty" json:"invoiceIncrementId,omitempty" yaml:"invoiceIncrementId,omitempty"`
}

// Operation wrapper for SalesOrderInvoiceVoid.
// OperationSalesOrderInvoiceVoidResponse was auto-generated from
// WSDL.
type OperationSalesOrderInvoiceVoidResponse struct {
	Result *bool `xml:"result,omitempty" json:"result,omitempty" yaml:"result,omitempty"`
}

// Operation wrapper for SalesOrderList.
// OperationSalesOrderListRequest was auto-generated from WSDL.
type OperationSalesOrderListRequest struct {
	SessionId *string  `xml:"sessionId,omitempty" json:"sessionId,omitempty" yaml:"sessionId,omitempty"`
	Filters   *Filters `xml:"filters,omitempty" json:"filters,omitempty" yaml:"filters,omitempty"`
}

// Operation wrapper for SalesOrderList.
// OperationSalesOrderListResponse was auto-generated from WSDL.
type OperationSalesOrderListResponse struct {
	Result *SalesOrderListEntityArray `xml:"result,omitempty" json:"result,omitempty" yaml:"result,omitempty"`
}

// Operation wrapper for SalesOrderShipmentAddComment.
// OperationSalesOrderShipmentAddCommentRequest was auto-generated
// from WSDL.
type OperationSalesOrderShipmentAddCommentRequest struct {
	SessionId           *string `xml:"sessionId,omitempty" json:"sessionId,omitempty" yaml:"sessionId,omitempty"`
	ShipmentIncrementId *string `xml:"shipmentIncrementId,omitempty" json:"shipmentIncrementId,omitempty" yaml:"shipmentIncrementId,omitempty"`
	Comment             *string `xml:"comment,omitempty" json:"comment,omitempty" yaml:"comment,omitempty"`
	Email               *string `xml:"email,omitempty" json:"email,omitempty" yaml:"email,omitempty"`
	IncludeInEmail      *string `xml:"includeInEmail,omitempty" json:"includeInEmail,omitempty" yaml:"includeInEmail,omitempty"`
}

// Operation wrapper for SalesOrderShipmentAddComment.
// OperationSalesOrderShipmentAddCommentResponse was auto-generated
// from WSDL.
type OperationSalesOrderShipmentAddCommentResponse struct {
	ShipmentIncrementId *bool `xml:"shipmentIncrementId,omitempty" json:"shipmentIncrementId,omitempty" yaml:"shipmentIncrementId,omitempty"`
}

// Operation wrapper for SalesOrderShipmentAddTrack.
// OperationSalesOrderShipmentAddTrackRequest was auto-generated
// from WSDL.
type OperationSalesOrderShipmentAddTrackRequest struct {
	SessionId           *string `xml:"sessionId,omitempty" json:"sessionId,omitempty" yaml:"sessionId,omitempty"`
	ShipmentIncrementId *string `xml:"shipmentIncrementId,omitempty" json:"shipmentIncrementId,omitempty" yaml:"shipmentIncrementId,omitempty"`
	Carrier             *string `xml:"carrier,omitempty" json:"carrier,omitempty" yaml:"carrier,omitempty"`
	Title               *string `xml:"title,omitempty" json:"title,omitempty" yaml:"title,omitempty"`
	TrackNumber         *string `xml:"trackNumber,omitempty" json:"trackNumber,omitempty" yaml:"trackNumber,omitempty"`
}

// Operation wrapper for SalesOrderShipmentAddTrack.
// OperationSalesOrderShipmentAddTrackResponse was auto-generated
// from WSDL.
type OperationSalesOrderShipmentAddTrackResponse struct {
	Result *int `xml:"result,omitempty" json:"result,omitempty" yaml:"result,omitempty"`
}

// Operation wrapper for SalesOrderShipmentCreate.
// OperationSalesOrderShipmentCreateRequest was auto-generated
// from WSDL.
type OperationSalesOrderShipmentCreateRequest struct {
	SessionId        *string              `xml:"sessionId,omitempty" json:"sessionId,omitempty" yaml:"sessionId,omitempty"`
	OrderIncrementId *string              `xml:"orderIncrementId,omitempty" json:"orderIncrementId,omitempty" yaml:"orderIncrementId,omitempty"`
	ItemsQty         *OrderItemIdQtyArray `xml:"itemsQty,omitempty" json:"itemsQty,omitempty" yaml:"itemsQty,omitempty"`
	Comment          *string              `xml:"comment,omitempty" json:"comment,omitempty" yaml:"comment,omitempty"`
	Email            *int                 `xml:"email,omitempty" json:"email,omitempty" yaml:"email,omitempty"`
	IncludeComment   *int                 `xml:"includeComment,omitempty" json:"includeComment,omitempty" yaml:"includeComment,omitempty"`
}

// Operation wrapper for SalesOrderShipmentCreate.
// OperationSalesOrderShipmentCreateResponse was auto-generated
// from WSDL.
type OperationSalesOrderShipmentCreateResponse struct {
	ShipmentIncrementId *string `xml:"shipmentIncrementId,omitempty" json:"shipmentIncrementId,omitempty" yaml:"shipmentIncrementId,omitempty"`
}

// Operation wrapper for SalesOrderShipmentGetCarriers.
// OperationSalesOrderShipmentGetCarriersRequest was auto-generated
// from WSDL.
type OperationSalesOrderShipmentGetCarriersRequest struct {
	SessionId        *string `xml:"sessionId,omitempty" json:"sessionId,omitempty" yaml:"sessionId,omitempty"`
	OrderIncrementId *string `xml:"orderIncrementId,omitempty" json:"orderIncrementId,omitempty" yaml:"orderIncrementId,omitempty"`
}

// Operation wrapper for SalesOrderShipmentGetCarriers.
// OperationSalesOrderShipmentGetCarriersResponse was auto-generated
// from WSDL.
type OperationSalesOrderShipmentGetCarriersResponse struct {
	Result *AssociativeArray `xml:"result,omitempty" json:"result,omitempty" yaml:"result,omitempty"`
}

// Operation wrapper for SalesOrderShipmentInfo.
// OperationSalesOrderShipmentInfoRequest was auto-generated from
// WSDL.
type OperationSalesOrderShipmentInfoRequest struct {
	SessionId           *string `xml:"sessionId,omitempty" json:"sessionId,omitempty" yaml:"sessionId,omitempty"`
	ShipmentIncrementId *string `xml:"shipmentIncrementId,omitempty" json:"shipmentIncrementId,omitempty" yaml:"shipmentIncrementId,omitempty"`
}

// Operation wrapper for SalesOrderShipmentInfo.
// OperationSalesOrderShipmentInfoResponse was auto-generated from
// WSDL.
type OperationSalesOrderShipmentInfoResponse struct {
	Result *SalesOrderShipmentEntity `xml:"result,omitempty" json:"result,omitempty" yaml:"result,omitempty"`
}

// Operation wrapper for SalesOrderShipmentList.
// OperationSalesOrderShipmentListRequest was auto-generated from
// WSDL.
type OperationSalesOrderShipmentListRequest struct {
	SessionId *string  `xml:"sessionId,omitempty" json:"sessionId,omitempty" yaml:"sessionId,omitempty"`
	Filters   *Filters `xml:"filters,omitempty" json:"filters,omitempty" yaml:"filters,omitempty"`
}

// Operation wrapper for SalesOrderShipmentList.
// OperationSalesOrderShipmentListResponse was auto-generated from
// WSDL.
type OperationSalesOrderShipmentListResponse struct {
	Result *SalesOrderShipmentEntityArray `xml:"result,omitempty" json:"result,omitempty" yaml:"result,omitempty"`
}

// Operation wrapper for SalesOrderShipmentRemoveTrack.
// OperationSalesOrderShipmentRemoveTrackRequest was auto-generated
// from WSDL.
type OperationSalesOrderShipmentRemoveTrackRequest struct {
	SessionId           *string `xml:"sessionId,omitempty" json:"sessionId,omitempty" yaml:"sessionId,omitempty"`
	ShipmentIncrementId *string `xml:"shipmentIncrementId,omitempty" json:"shipmentIncrementId,omitempty" yaml:"shipmentIncrementId,omitempty"`
	TrackId             *string `xml:"trackId,omitempty" json:"trackId,omitempty" yaml:"trackId,omitempty"`
}

// Operation wrapper for SalesOrderShipmentRemoveTrack.
// OperationSalesOrderShipmentRemoveTrackResponse was auto-generated
// from WSDL.
type OperationSalesOrderShipmentRemoveTrackResponse struct {
	Result *bool `xml:"result,omitempty" json:"result,omitempty" yaml:"result,omitempty"`
}

// Operation wrapper for SalesOrderShipmentSendInfo.
// OperationSalesOrderShipmentSendInfoRequest was auto-generated
// from WSDL.
type OperationSalesOrderShipmentSendInfoRequest struct {
	SessionId           *string `xml:"sessionId,omitempty" json:"sessionId,omitempty" yaml:"sessionId,omitempty"`
	ShipmentIncrementId *string `xml:"shipmentIncrementId,omitempty" json:"shipmentIncrementId,omitempty" yaml:"shipmentIncrementId,omitempty"`
	Comment             *string `xml:"comment,omitempty" json:"comment,omitempty" yaml:"comment,omitempty"`
}

// Operation wrapper for SalesOrderShipmentSendInfo.
// OperationSalesOrderShipmentSendInfoResponse was auto-generated
// from WSDL.
type OperationSalesOrderShipmentSendInfoResponse struct {
	Result *bool `xml:"result,omitempty" json:"result,omitempty" yaml:"result,omitempty"`
}

// Operation wrapper for SalesOrderUnhold.
// OperationSalesOrderUnholdRequest was auto-generated from WSDL.
type OperationSalesOrderUnholdRequest struct {
	SessionId        *string `xml:"sessionId,omitempty" json:"sessionId,omitempty" yaml:"sessionId,omitempty"`
	OrderIncrementId *string `xml:"orderIncrementId,omitempty" json:"orderIncrementId,omitempty" yaml:"orderIncrementId,omitempty"`
}

// Operation wrapper for SalesOrderUnhold.
// OperationSalesOrderUnholdResponse was auto-generated from WSDL.
type OperationSalesOrderUnholdResponse struct {
	Result *bool `xml:"result,omitempty" json:"result,omitempty" yaml:"result,omitempty"`
}

// Operation wrapper for ShoppingCartCouponAdd.
// OperationShoppingCartCouponAddRequest was auto-generated from
// WSDL.
type OperationShoppingCartCouponAddRequest struct {
	SessionId  *string `xml:"sessionId,omitempty" json:"sessionId,omitempty" yaml:"sessionId,omitempty"`
	QuoteId    *int    `xml:"quoteId,omitempty" json:"quoteId,omitempty" yaml:"quoteId,omitempty"`
	CouponCode *string `xml:"couponCode,omitempty" json:"couponCode,omitempty" yaml:"couponCode,omitempty"`
	StoreId    *string `xml:"storeId,omitempty" json:"storeId,omitempty" yaml:"storeId,omitempty"`
}

// Operation wrapper for ShoppingCartCouponAdd.
// OperationShoppingCartCouponAddResponse was auto-generated from
// WSDL.
type OperationShoppingCartCouponAddResponse struct {
	Result *bool `xml:"result,omitempty" json:"result,omitempty" yaml:"result,omitempty"`
}

// Operation wrapper for ShoppingCartCouponRemove.
// OperationShoppingCartCouponRemoveRequest was auto-generated
// from WSDL.
type OperationShoppingCartCouponRemoveRequest struct {
	SessionId *string `xml:"sessionId,omitempty" json:"sessionId,omitempty" yaml:"sessionId,omitempty"`
	QuoteId   *int    `xml:"quoteId,omitempty" json:"quoteId,omitempty" yaml:"quoteId,omitempty"`
	StoreId   *string `xml:"storeId,omitempty" json:"storeId,omitempty" yaml:"storeId,omitempty"`
}

// Operation wrapper for ShoppingCartCouponRemove.
// OperationShoppingCartCouponRemoveResponse was auto-generated
// from WSDL.
type OperationShoppingCartCouponRemoveResponse struct {
	Result *bool `xml:"result,omitempty" json:"result,omitempty" yaml:"result,omitempty"`
}

// Operation wrapper for ShoppingCartCreate.
// OperationShoppingCartCreateRequest was auto-generated from WSDL.
type OperationShoppingCartCreateRequest struct {
	SessionId *string `xml:"sessionId,omitempty" json:"sessionId,omitempty" yaml:"sessionId,omitempty"`
	StoreId   *string `xml:"storeId,omitempty" json:"storeId,omitempty" yaml:"storeId,omitempty"`
}

// Operation wrapper for ShoppingCartCreate.
// OperationShoppingCartCreateResponse was auto-generated from
// WSDL.
type OperationShoppingCartCreateResponse struct {
	QuoteId *int `xml:"quoteId,omitempty" json:"quoteId,omitempty" yaml:"quoteId,omitempty"`
}

// Operation wrapper for ShoppingCartCustomerAddresses.
// OperationShoppingCartCustomerAddressesRequest was auto-generated
// from WSDL.
type OperationShoppingCartCustomerAddressesRequest struct {
	SessionId *string                                 `xml:"sessionId,omitempty" json:"sessionId,omitempty" yaml:"sessionId,omitempty"`
	QuoteId   *int                                    `xml:"quoteId,omitempty" json:"quoteId,omitempty" yaml:"quoteId,omitempty"`
	Customer  *ShoppingCartCustomerAddressEntityArray `xml:"customer,omitempty" json:"customer,omitempty" yaml:"customer,omitempty"`
	StoreId   *string                                 `xml:"storeId,omitempty" json:"storeId,omitempty" yaml:"storeId,omitempty"`
}

// Operation wrapper for ShoppingCartCustomerAddresses.
// OperationShoppingCartCustomerAddressesResponse was auto-generated
// from WSDL.
type OperationShoppingCartCustomerAddressesResponse struct {
	Result *bool `xml:"result,omitempty" json:"result,omitempty" yaml:"result,omitempty"`
}

// Operation wrapper for ShoppingCartCustomerSet.
// OperationShoppingCartCustomerSetRequest was auto-generated from
// WSDL.
type OperationShoppingCartCustomerSetRequest struct {
	SessionId *string                     `xml:"sessionId,omitempty" json:"sessionId,omitempty" yaml:"sessionId,omitempty"`
	QuoteId   *int                        `xml:"quoteId,omitempty" json:"quoteId,omitempty" yaml:"quoteId,omitempty"`
	Customer  *ShoppingCartCustomerEntity `xml:"customer,omitempty" json:"customer,omitempty" yaml:"customer,omitempty"`
	StoreId   *string                     `xml:"storeId,omitempty" json:"storeId,omitempty" yaml:"storeId,omitempty"`
}

// Operation wrapper for ShoppingCartCustomerSet.
// OperationShoppingCartCustomerSetResponse was auto-generated
// from WSDL.
type OperationShoppingCartCustomerSetResponse struct {
	Result *bool `xml:"result,omitempty" json:"result,omitempty" yaml:"result,omitempty"`
}

// Operation wrapper for ShoppingCartInfo.
// OperationShoppingCartInfoRequest was auto-generated from WSDL.
type OperationShoppingCartInfoRequest struct {
	SessionId *string `xml:"sessionId,omitempty" json:"sessionId,omitempty" yaml:"sessionId,omitempty"`
	QuoteId   *int    `xml:"quoteId,omitempty" json:"quoteId,omitempty" yaml:"quoteId,omitempty"`
	StoreId   *string `xml:"storeId,omitempty" json:"storeId,omitempty" yaml:"storeId,omitempty"`
}

// Operation wrapper for ShoppingCartInfo.
// OperationShoppingCartInfoResponse was auto-generated from WSDL.
type OperationShoppingCartInfoResponse struct {
	Result *ShoppingCartInfoEntity `xml:"result,omitempty" json:"result,omitempty" yaml:"result,omitempty"`
}

// Operation wrapper for ShoppingCartLicense.
// OperationShoppingCartLicenseRequest was auto-generated from
// WSDL.
type OperationShoppingCartLicenseRequest struct {
	SessionId *string `xml:"sessionId,omitempty" json:"sessionId,omitempty" yaml:"sessionId,omitempty"`
	QuoteId   *int    `xml:"quoteId,omitempty" json:"quoteId,omitempty" yaml:"quoteId,omitempty"`
	StoreId   *string `xml:"storeId,omitempty" json:"storeId,omitempty" yaml:"storeId,omitempty"`
}

// Operation wrapper for ShoppingCartLicense.
// OperationShoppingCartLicenseResponse was auto-generated from
// WSDL.
type OperationShoppingCartLicenseResponse struct {
	Result *ShoppingCartLicenseEntityArray `xml:"result,omitempty" json:"result,omitempty" yaml:"result,omitempty"`
}

// Operation wrapper for ShoppingCartOrder.
// OperationShoppingCartOrderRequest was auto-generated from WSDL.
type OperationShoppingCartOrderRequest struct {
	SessionId *string        `xml:"sessionId,omitempty" json:"sessionId,omitempty" yaml:"sessionId,omitempty"`
	QuoteId   *int           `xml:"quoteId,omitempty" json:"quoteId,omitempty" yaml:"quoteId,omitempty"`
	StoreId   *string        `xml:"storeId,omitempty" json:"storeId,omitempty" yaml:"storeId,omitempty"`
	Licenses  *ArrayOfString `xml:"licenses,omitempty" json:"licenses,omitempty" yaml:"licenses,omitempty"`
}

// Operation wrapper for ShoppingCartOrder.
// OperationShoppingCartOrderResponse was auto-generated from WSDL.
type OperationShoppingCartOrderResponse struct {
	Result *string `xml:"result,omitempty" json:"result,omitempty" yaml:"result,omitempty"`
}

// Operation wrapper for ShoppingCartPaymentList.
// OperationShoppingCartPaymentListRequest was auto-generated from
// WSDL.
type OperationShoppingCartPaymentListRequest struct {
	SessionId *string `xml:"sessionId,omitempty" json:"sessionId,omitempty" yaml:"sessionId,omitempty"`
	QuoteId   *int    `xml:"quoteId,omitempty" json:"quoteId,omitempty" yaml:"quoteId,omitempty"`
	Store     *string `xml:"store,omitempty" json:"store,omitempty" yaml:"store,omitempty"`
}

// Operation wrapper for ShoppingCartPaymentList.
// OperationShoppingCartPaymentListResponse was auto-generated
// from WSDL.
type OperationShoppingCartPaymentListResponse struct {
	Result *ShoppingCartPaymentMethodResponseEntityArray `xml:"result,omitempty" json:"result,omitempty" yaml:"result,omitempty"`
}

// Operation wrapper for ShoppingCartPaymentMethod.
// OperationShoppingCartPaymentMethodRequest was auto-generated
// from WSDL.
type OperationShoppingCartPaymentMethodRequest struct {
	SessionId *string                          `xml:"sessionId,omitempty" json:"sessionId,omitempty" yaml:"sessionId,omitempty"`
	QuoteId   *int                             `xml:"quoteId,omitempty" json:"quoteId,omitempty" yaml:"quoteId,omitempty"`
	Method    *ShoppingCartPaymentMethodEntity `xml:"method,omitempty" json:"method,omitempty" yaml:"method,omitempty"`
	StoreId   *string                          `xml:"storeId,omitempty" json:"storeId,omitempty" yaml:"storeId,omitempty"`
}

// Operation wrapper for ShoppingCartPaymentMethod.
// OperationShoppingCartPaymentMethodResponse was auto-generated
// from WSDL.
type OperationShoppingCartPaymentMethodResponse struct {
	Result *bool `xml:"result,omitempty" json:"result,omitempty" yaml:"result,omitempty"`
}

// Operation wrapper for ShoppingCartProductAdd.
// OperationShoppingCartProductAddRequest was auto-generated from
// WSDL.
type OperationShoppingCartProductAddRequest struct {
	SessionId *string                         `xml:"sessionId,omitempty" json:"sessionId,omitempty" yaml:"sessionId,omitempty"`
	QuoteId   *int                            `xml:"quoteId,omitempty" json:"quoteId,omitempty" yaml:"quoteId,omitempty"`
	Products  *ShoppingCartProductEntityArray `xml:"products,omitempty" json:"products,omitempty" yaml:"products,omitempty"`
	StoreId   *string                         `xml:"storeId,omitempty" json:"storeId,omitempty" yaml:"storeId,omitempty"`
}

// Operation wrapper for ShoppingCartProductAdd.
// OperationShoppingCartProductAddResponse was auto-generated from
// WSDL.
type OperationShoppingCartProductAddResponse struct {
	Result *bool `xml:"result,omitempty" json:"result,omitempty" yaml:"result,omitempty"`
}

// Operation wrapper for ShoppingCartProductList.
// OperationShoppingCartProductListRequest was auto-generated from
// WSDL.
type OperationShoppingCartProductListRequest struct {
	SessionId *string `xml:"sessionId,omitempty" json:"sessionId,omitempty" yaml:"sessionId,omitempty"`
	QuoteId   *int    `xml:"quoteId,omitempty" json:"quoteId,omitempty" yaml:"quoteId,omitempty"`
	StoreId   *string `xml:"storeId,omitempty" json:"storeId,omitempty" yaml:"storeId,omitempty"`
}

// Operation wrapper for ShoppingCartProductList.
// OperationShoppingCartProductListResponse was auto-generated
// from WSDL.
type OperationShoppingCartProductListResponse struct {
	Result *ShoppingCartProductResponseEntityArray `xml:"result,omitempty" json:"result,omitempty" yaml:"result,omitempty"`
}

// Operation wrapper for ShoppingCartProductMoveToCustomerQuote.
// OperationShoppingCartProductMoveToCustomerQuoteRequest was auto-generated
// from WSDL.
type OperationShoppingCartProductMoveToCustomerQuoteRequest struct {
	SessionId *string                         `xml:"sessionId,omitempty" json:"sessionId,omitempty" yaml:"sessionId,omitempty"`
	QuoteId   *int                            `xml:"quoteId,omitempty" json:"quoteId,omitempty" yaml:"quoteId,omitempty"`
	Products  *ShoppingCartProductEntityArray `xml:"products,omitempty" json:"products,omitempty" yaml:"products,omitempty"`
	StoreId   *string                         `xml:"storeId,omitempty" json:"storeId,omitempty" yaml:"storeId,omitempty"`
}

// Operation wrapper for ShoppingCartProductMoveToCustomerQuote.
// OperationShoppingCartProductMoveToCustomerQuoteResponse was
// auto-generated from WSDL.
type OperationShoppingCartProductMoveToCustomerQuoteResponse struct {
	Result *bool `xml:"result,omitempty" json:"result,omitempty" yaml:"result,omitempty"`
}

// Operation wrapper for ShoppingCartProductRemove.
// OperationShoppingCartProductRemoveRequest was auto-generated
// from WSDL.
type OperationShoppingCartProductRemoveRequest struct {
	SessionId *string                         `xml:"sessionId,omitempty" json:"sessionId,omitempty" yaml:"sessionId,omitempty"`
	QuoteId   *int                            `xml:"quoteId,omitempty" json:"quoteId,omitempty" yaml:"quoteId,omitempty"`
	Products  *ShoppingCartProductEntityArray `xml:"products,omitempty" json:"products,omitempty" yaml:"products,omitempty"`
	StoreId   *string                         `xml:"storeId,omitempty" json:"storeId,omitempty" yaml:"storeId,omitempty"`
}

// Operation wrapper for ShoppingCartProductRemove.
// OperationShoppingCartProductRemoveResponse was auto-generated
// from WSDL.
type OperationShoppingCartProductRemoveResponse struct {
	Result *bool `xml:"result,omitempty" json:"result,omitempty" yaml:"result,omitempty"`
}

// Operation wrapper for ShoppingCartProductUpdate.
// OperationShoppingCartProductUpdateRequest was auto-generated
// from WSDL.
type OperationShoppingCartProductUpdateRequest struct {
	SessionId *string                         `xml:"sessionId,omitempty" json:"sessionId,omitempty" yaml:"sessionId,omitempty"`
	QuoteId   *int                            `xml:"quoteId,omitempty" json:"quoteId,omitempty" yaml:"quoteId,omitempty"`
	Products  *ShoppingCartProductEntityArray `xml:"products,omitempty" json:"products,omitempty" yaml:"products,omitempty"`
	StoreId   *string                         `xml:"storeId,omitempty" json:"storeId,omitempty" yaml:"storeId,omitempty"`
}

// Operation wrapper for ShoppingCartProductUpdate.
// OperationShoppingCartProductUpdateResponse was auto-generated
// from WSDL.
type OperationShoppingCartProductUpdateResponse struct {
	Result *bool `xml:"result,omitempty" json:"result,omitempty" yaml:"result,omitempty"`
}

// Operation wrapper for ShoppingCartShippingList.
// OperationShoppingCartShippingListRequest was auto-generated
// from WSDL.
type OperationShoppingCartShippingListRequest struct {
	SessionId *string `xml:"sessionId,omitempty" json:"sessionId,omitempty" yaml:"sessionId,omitempty"`
	QuoteId   *int    `xml:"quoteId,omitempty" json:"quoteId,omitempty" yaml:"quoteId,omitempty"`
	StoreId   *string `xml:"storeId,omitempty" json:"storeId,omitempty" yaml:"storeId,omitempty"`
}

// Operation wrapper for ShoppingCartShippingList.
// OperationShoppingCartShippingListResponse was auto-generated
// from WSDL.
type OperationShoppingCartShippingListResponse struct {
	Result *ShoppingCartShippingMethodEntityArray `xml:"result,omitempty" json:"result,omitempty" yaml:"result,omitempty"`
}

// Operation wrapper for ShoppingCartShippingMethod.
// OperationShoppingCartShippingMethodRequest was auto-generated
// from WSDL.
type OperationShoppingCartShippingMethodRequest struct {
	SessionId *string `xml:"sessionId,omitempty" json:"sessionId,omitempty" yaml:"sessionId,omitempty"`
	QuoteId   *int    `xml:"quoteId,omitempty" json:"quoteId,omitempty" yaml:"quoteId,omitempty"`
	Method    *string `xml:"method,omitempty" json:"method,omitempty" yaml:"method,omitempty"`
	StoreId   *string `xml:"storeId,omitempty" json:"storeId,omitempty" yaml:"storeId,omitempty"`
}

// Operation wrapper for ShoppingCartShippingMethod.
// OperationShoppingCartShippingMethodResponse was auto-generated
// from WSDL.
type OperationShoppingCartShippingMethodResponse struct {
	Result *bool `xml:"result,omitempty" json:"result,omitempty" yaml:"result,omitempty"`
}

// Operation wrapper for ShoppingCartTotals.
// OperationShoppingCartTotalsRequest was auto-generated from WSDL.
type OperationShoppingCartTotalsRequest struct {
	SessionId *string `xml:"sessionId,omitempty" json:"sessionId,omitempty" yaml:"sessionId,omitempty"`
	QuoteId   *int    `xml:"quoteId,omitempty" json:"quoteId,omitempty" yaml:"quoteId,omitempty"`
	StoreId   *string `xml:"storeId,omitempty" json:"storeId,omitempty" yaml:"storeId,omitempty"`
}

// Operation wrapper for ShoppingCartTotals.
// OperationShoppingCartTotalsResponse was auto-generated from
// WSDL.
type OperationShoppingCartTotalsResponse struct {
	Result *ShoppingCartTotalsEntityArray `xml:"result,omitempty" json:"result,omitempty" yaml:"result,omitempty"`
}

// Operation wrapper for StartSession.
// OperationStartSessionResponse was auto-generated from WSDL.
type OperationStartSessionResponse struct {
	StartSessionReturn *string `xml:"startSessionReturn,omitempty" json:"startSessionReturn,omitempty" yaml:"startSessionReturn,omitempty"`
}

// Operation wrapper for StoreInfo.
// OperationStoreInfoRequest was auto-generated from WSDL.
type OperationStoreInfoRequest struct {
	SessionId *string `xml:"sessionId,omitempty" json:"sessionId,omitempty" yaml:"sessionId,omitempty"`
	StoreId   *string `xml:"storeId,omitempty" json:"storeId,omitempty" yaml:"storeId,omitempty"`
}

// Operation wrapper for StoreInfo.
// OperationStoreInfoResponse was auto-generated from WSDL.
type OperationStoreInfoResponse struct {
	Info *StoreEntity `xml:"info,omitempty" json:"info,omitempty" yaml:"info,omitempty"`
}

// Operation wrapper for StoreList.
// OperationStoreListRequest was auto-generated from WSDL.
type OperationStoreListRequest struct {
	SessionId *string `xml:"sessionId,omitempty" json:"sessionId,omitempty" yaml:"sessionId,omitempty"`
}

// Operation wrapper for StoreList.
// OperationStoreListResponse was auto-generated from WSDL.
type OperationStoreListResponse struct {
	Stores *StoreEntityArray `xml:"stores,omitempty" json:"stores,omitempty" yaml:"stores,omitempty"`
}

// portType implements the PortType interface.
type portType struct {
	cli *soap.Client
}

// Create an order from Adcurve shopping cart
func (p *portType) AdcurveShoppingCartCreateAdcurveOrder(sessionId string, quoteId int, storeId string, licenses *ArrayOfString, orderAdditionalAttributes *ShoppingCartOrderAdditionalAttributesEntity) (string, error) {
	α := struct {
		M OperationAdcurveShoppingCartCreateAdcurveOrderRequest `xml:"typens:adcurveShoppingCartCreateAdcurveOrder"`
	}{
		OperationAdcurveShoppingCartCreateAdcurveOrderRequest{
			&sessionId,
			&quoteId,
			&storeId,
			licenses,
			orderAdditionalAttributes,
		},
	}

	γ := struct {
		M OperationAdcurveShoppingCartCreateAdcurveOrderResponse `xml:"adcurveShoppingCartCreateAdcurveOrderResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("urn:Action", α, &γ); err != nil {
		return "", err
	}
	return *γ.M.Result, nil
}

// Add product(s) to adcurve shopping cart
func (p *portType) AdcurveShoppingCartProductAdcurveProductAdd(sessionId string, quoteId int, products *ShoppingCartAdcurveProductEntityArray, storeId string) (bool, error) {
	α := struct {
		M OperationAdcurveShoppingCartProductAdcurveProductAddRequest `xml:"typens:adcurveShoppingCartProductAdcurveProductAdd"`
	}{
		OperationAdcurveShoppingCartProductAdcurveProductAddRequest{
			&sessionId,
			&quoteId,
			products,
			&storeId,
		},
	}

	γ := struct {
		M OperationAdcurveShoppingCartProductAdcurveProductAddResponse `xml:"adcurveShoppingCartProductAdcurveProductAddResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("urn:Action", α, &γ); err != nil {
		return false, err
	}
	return *γ.M.Result, nil
}

// Update product(s) quantities in adcurve shopping cart
func (p *portType) AdcurveShoppingCartProductAdcurveProductUpdate(sessionId string, quoteId int, products *ShoppingCartAdcurveProductEntityArray, storeId string) (bool, error) {
	α := struct {
		M OperationAdcurveShoppingCartProductAdcurveProductUpdateRequest `xml:"typens:adcurveShoppingCartProductAdcurveProductUpdate"`
	}{
		OperationAdcurveShoppingCartProductAdcurveProductUpdateRequest{
			&sessionId,
			&quoteId,
			products,
			&storeId,
		},
	}

	γ := struct {
		M OperationAdcurveShoppingCartProductAdcurveProductUpdateResponse `xml:"adcurveShoppingCartProductAdcurveProductUpdateResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("urn:Action", α, &γ); err != nil {
		return false, err
	}
	return *γ.M.Result, nil
}

// Assign product to category
func (p *portType) CatalogCategoryAssignProduct(sessionId string, categoryId int, product string, position string, identifierType string) (bool, error) {
	α := struct {
		M OperationCatalogCategoryAssignProductRequest `xml:"typens:catalogCategoryAssignProduct"`
	}{
		OperationCatalogCategoryAssignProductRequest{
			&sessionId,
			&categoryId,
			&product,
			&position,
			&identifierType,
		},
	}

	γ := struct {
		M OperationCatalogCategoryAssignProductResponse `xml:"catalogCategoryAssignProductResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("urn:Action", α, &γ); err != nil {
		return false, err
	}
	return *γ.M.Result, nil
}

// Retrieve list of assigned products
func (p *portType) CatalogCategoryAssignedProducts(sessionId string, categoryId int) (*CatalogAssignedProductArray, error) {
	α := struct {
		M OperationCatalogCategoryAssignedProductsRequest `xml:"typens:catalogCategoryAssignedProducts"`
	}{
		OperationCatalogCategoryAssignedProductsRequest{
			&sessionId,
			&categoryId,
		},
	}

	γ := struct {
		M OperationCatalogCategoryAssignedProductsResponse `xml:"catalogCategoryAssignedProductsResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("urn:Action", α, &γ); err != nil {
		return nil, err
	}
	return γ.M.Result, nil
}

// Set/Get current store view
func (p *portType) CatalogCategoryAttributeCurrentStore(sessionId string, storeView string) (int, error) {
	α := struct {
		M OperationCatalogCategoryAttributeCurrentStoreRequest `xml:"typens:catalogCategoryAttributeCurrentStore"`
	}{
		OperationCatalogCategoryAttributeCurrentStoreRequest{
			&sessionId,
			&storeView,
		},
	}

	γ := struct {
		M OperationCatalogCategoryAttributeCurrentStoreResponse `xml:"catalogCategoryAttributeCurrentStoreResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("urn:Action", α, &γ); err != nil {
		return 0, err
	}
	return *γ.M.StoreView, nil
}

// Retrieve category attributes
func (p *portType) CatalogCategoryAttributeList(sessionId string) (*CatalogAttributeEntityArray, error) {
	α := struct {
		M OperationCatalogCategoryAttributeListRequest `xml:"typens:catalogCategoryAttributeList"`
	}{
		OperationCatalogCategoryAttributeListRequest{
			&sessionId,
		},
	}

	γ := struct {
		M OperationCatalogCategoryAttributeListResponse `xml:"catalogCategoryAttributeListResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("urn:Action", α, &γ); err != nil {
		return nil, err
	}
	return γ.M.Result, nil
}

// Retrieve attribute options
func (p *portType) CatalogCategoryAttributeOptions(sessionId string, attributeId string, storeView string) (*CatalogAttributeOptionEntityArray, error) {
	α := struct {
		M OperationCatalogCategoryAttributeOptionsRequest `xml:"typens:catalogCategoryAttributeOptions"`
	}{
		OperationCatalogCategoryAttributeOptionsRequest{
			&sessionId,
			&attributeId,
			&storeView,
		},
	}

	γ := struct {
		M OperationCatalogCategoryAttributeOptionsResponse `xml:"catalogCategoryAttributeOptionsResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("urn:Action", α, &γ); err != nil {
		return nil, err
	}
	return γ.M.Result, nil
}

// Create new category and return its id.
func (p *portType) CatalogCategoryCreate(sessionId string, parentId int, categoryData *CatalogCategoryEntityCreate, storeView string) (int, error) {
	α := struct {
		M OperationCatalogCategoryCreateRequest `xml:"typens:catalogCategoryCreate"`
	}{
		OperationCatalogCategoryCreateRequest{
			&sessionId,
			&parentId,
			categoryData,
			&storeView,
		},
	}

	γ := struct {
		M OperationCatalogCategoryCreateResponse `xml:"catalogCategoryCreateResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("urn:Action", α, &γ); err != nil {
		return 0, err
	}
	return *γ.M.Attribute_id, nil
}

// Set_Get current store view
func (p *portType) CatalogCategoryCurrentStore(sessionId string, storeView string) (int, error) {
	α := struct {
		M OperationCatalogCategoryCurrentStoreRequest `xml:"typens:catalogCategoryCurrentStore"`
	}{
		OperationCatalogCategoryCurrentStoreRequest{
			&sessionId,
			&storeView,
		},
	}

	γ := struct {
		M OperationCatalogCategoryCurrentStoreResponse `xml:"catalogCategoryCurrentStoreResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("urn:Action", α, &γ); err != nil {
		return 0, err
	}
	return *γ.M.StoreView, nil
}

// Delete category
func (p *portType) CatalogCategoryDelete(sessionId string, categoryId int) (bool, error) {
	α := struct {
		M OperationCatalogCategoryDeleteRequest `xml:"typens:catalogCategoryDelete"`
	}{
		OperationCatalogCategoryDeleteRequest{
			&sessionId,
			&categoryId,
		},
	}

	γ := struct {
		M OperationCatalogCategoryDeleteResponse `xml:"catalogCategoryDeleteResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("urn:Action", α, &γ); err != nil {
		return false, err
	}
	return *γ.M.Result, nil
}

// Retrieve hierarchical tree of categories.
func (p *portType) CatalogCategoryInfo(sessionId string, categoryId int, storeView string, attributes *ArrayOfString) (*CatalogCategoryInfo, error) {
	α := struct {
		M OperationCatalogCategoryInfoRequest `xml:"typens:catalogCategoryInfo"`
	}{
		OperationCatalogCategoryInfoRequest{
			&sessionId,
			&categoryId,
			&storeView,
			attributes,
		},
	}

	γ := struct {
		M OperationCatalogCategoryInfoResponse `xml:"catalogCategoryInfoResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("urn:Action", α, &γ); err != nil {
		return nil, err
	}
	return γ.M.Info, nil
}

// Retrieve hierarchical tree of categories.
func (p *portType) CatalogCategoryLevel(sessionId string, website string, storeView string, parentCategory string) (*ArrayOfCatalogCategoryEntitiesNoChildren, error) {
	α := struct {
		M OperationCatalogCategoryLevelRequest `xml:"typens:catalogCategoryLevel"`
	}{
		OperationCatalogCategoryLevelRequest{
			&sessionId,
			&website,
			&storeView,
			&parentCategory,
		},
	}

	γ := struct {
		M OperationCatalogCategoryLevelResponse `xml:"catalogCategoryLevelResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("urn:Action", α, &γ); err != nil {
		return nil, err
	}
	return γ.M.Tree, nil
}

// Move category in tree
func (p *portType) CatalogCategoryMove(sessionId string, categoryId int, parentId int, afterId string) (bool, error) {
	α := struct {
		M OperationCatalogCategoryMoveRequest `xml:"typens:catalogCategoryMove"`
	}{
		OperationCatalogCategoryMoveRequest{
			&sessionId,
			&categoryId,
			&parentId,
			&afterId,
		},
	}

	γ := struct {
		M OperationCatalogCategoryMoveResponse `xml:"catalogCategoryMoveResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("urn:Action", α, &γ); err != nil {
		return false, err
	}
	return *γ.M.Id, nil
}

// Remove product assignment from category
func (p *portType) CatalogCategoryRemoveProduct(sessionId string, categoryId int, product string, identifierType string) (bool, error) {
	α := struct {
		M OperationCatalogCategoryRemoveProductRequest `xml:"typens:catalogCategoryRemoveProduct"`
	}{
		OperationCatalogCategoryRemoveProductRequest{
			&sessionId,
			&categoryId,
			&product,
			&identifierType,
		},
	}

	γ := struct {
		M OperationCatalogCategoryRemoveProductResponse `xml:"catalogCategoryRemoveProductResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("urn:Action", α, &γ); err != nil {
		return false, err
	}
	return *γ.M.Result, nil
}

// Retrieve hierarchical tree of categories.
func (p *portType) CatalogCategoryTree(sessionId string, parentId string, storeView string) (*CatalogCategoryTree, error) {
	α := struct {
		M OperationCatalogCategoryTreeRequest `xml:"typens:catalogCategoryTree"`
	}{
		OperationCatalogCategoryTreeRequest{
			&sessionId,
			&parentId,
			&storeView,
		},
	}

	γ := struct {
		M OperationCatalogCategoryTreeResponse `xml:"catalogCategoryTreeResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("urn:Action", α, &γ); err != nil {
		return nil, err
	}
	return γ.M.Tree, nil
}

// Update category
func (p *portType) CatalogCategoryUpdate(sessionId string, categoryId int, categoryData *CatalogCategoryEntityCreate, storeView string) (bool, error) {
	α := struct {
		M OperationCatalogCategoryUpdateRequest `xml:"typens:catalogCategoryUpdate"`
	}{
		OperationCatalogCategoryUpdateRequest{
			&sessionId,
			&categoryId,
			categoryData,
			&storeView,
		},
	}

	γ := struct {
		M OperationCatalogCategoryUpdateResponse `xml:"catalogCategoryUpdateResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("urn:Action", α, &γ); err != nil {
		return false, err
	}
	return *γ.M.Id, nil
}

// Update assigned product
func (p *portType) CatalogCategoryUpdateProduct(sessionId string, categoryId int, product string, position string, identifierType string) (bool, error) {
	α := struct {
		M OperationCatalogCategoryUpdateProductRequest `xml:"typens:catalogCategoryUpdateProduct"`
	}{
		OperationCatalogCategoryUpdateProductRequest{
			&sessionId,
			&categoryId,
			&product,
			&position,
			&identifierType,
		},
	}

	γ := struct {
		M OperationCatalogCategoryUpdateProductResponse `xml:"catalogCategoryUpdateProductResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("urn:Action", α, &γ); err != nil {
		return false, err
	}
	return *γ.M.Result, nil
}

// Retrieve stock data by product ids
func (p *portType) CatalogInventoryStockItemList(sessionId string, products *ArrayOfString) (*CatalogInventoryStockItemEntityArray, error) {
	α := struct {
		M OperationCatalogInventoryStockItemListRequest `xml:"typens:catalogInventoryStockItemList"`
	}{
		OperationCatalogInventoryStockItemListRequest{
			&sessionId,
			products,
		},
	}

	γ := struct {
		M OperationCatalogInventoryStockItemListResponse `xml:"catalogInventoryStockItemListResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("urn:Action", α, &γ); err != nil {
		return nil, err
	}
	return γ.M.Result, nil
}

// Multi stock item update
func (p *portType) CatalogInventoryStockItemMultiUpdate(sessionId string, productIds *ArrayOfString, productData *CatalogInventoryStockItemUpdateEntityArray) (bool, error) {
	α := struct {
		M OperationCatalogInventoryStockItemMultiUpdateRequest `xml:"typens:catalogInventoryStockItemMultiUpdate"`
	}{
		OperationCatalogInventoryStockItemMultiUpdateRequest{
			&sessionId,
			productIds,
			productData,
		},
	}

	γ := struct {
		M OperationCatalogInventoryStockItemMultiUpdateResponse `xml:"catalogInventoryStockItemMultiUpdateResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("urn:Action", α, &γ); err != nil {
		return false, err
	}
	return *γ.M.Result, nil
}

// Update product stock data
func (p *portType) CatalogInventoryStockItemUpdate(sessionId string, product string, data *CatalogInventoryStockItemUpdateEntity) (int, error) {
	α := struct {
		M OperationCatalogInventoryStockItemUpdateRequest `xml:"typens:catalogInventoryStockItemUpdate"`
	}{
		OperationCatalogInventoryStockItemUpdateRequest{
			&sessionId,
			&product,
			data,
		},
	}

	γ := struct {
		M OperationCatalogInventoryStockItemUpdateResponse `xml:"catalogInventoryStockItemUpdateResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("urn:Action", α, &γ); err != nil {
		return 0, err
	}
	return *γ.M.Result, nil
}

// Add option to attribute
func (p *portType) CatalogProductAttributeAddOption(sessionId string, attribute string, data *CatalogProductAttributeOptionEntityToAdd) (bool, error) {
	α := struct {
		M OperationCatalogProductAttributeAddOptionRequest `xml:"typens:catalogProductAttributeAddOption"`
	}{
		OperationCatalogProductAttributeAddOptionRequest{
			&sessionId,
			&attribute,
			data,
		},
	}

	γ := struct {
		M OperationCatalogProductAttributeAddOptionResponse `xml:"catalogProductAttributeAddOptionResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("urn:Action", α, &γ); err != nil {
		return false, err
	}
	return *γ.M.Result, nil
}

// Create new attribute
func (p *portType) CatalogProductAttributeCreate(sessionId string, data *CatalogProductAttributeEntityToCreate) (int, error) {
	α := struct {
		M OperationCatalogProductAttributeCreateRequest `xml:"typens:catalogProductAttributeCreate"`
	}{
		OperationCatalogProductAttributeCreateRequest{
			&sessionId,
			data,
		},
	}

	γ := struct {
		M OperationCatalogProductAttributeCreateResponse `xml:"catalogProductAttributeCreateResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("urn:Action", α, &γ); err != nil {
		return 0, err
	}
	return *γ.M.Result, nil
}

// Set/Get current store view
func (p *portType) CatalogProductAttributeCurrentStore(sessionId string, storeView string) (int, error) {
	α := struct {
		M OperationCatalogProductAttributeCurrentStoreRequest `xml:"typens:catalogProductAttributeCurrentStore"`
	}{
		OperationCatalogProductAttributeCurrentStoreRequest{
			&sessionId,
			&storeView,
		},
	}

	γ := struct {
		M OperationCatalogProductAttributeCurrentStoreResponse `xml:"catalogProductAttributeCurrentStoreResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("urn:Action", α, &γ); err != nil {
		return 0, err
	}
	return *γ.M.StoreView, nil
}

// Get full information about attribute with list of options
func (p *portType) CatalogProductAttributeInfo(sessionId string, attribute string) (*CatalogProductAttributeEntity, error) {
	α := struct {
		M OperationCatalogProductAttributeInfoRequest `xml:"typens:catalogProductAttributeInfo"`
	}{
		OperationCatalogProductAttributeInfoRequest{
			&sessionId,
			&attribute,
		},
	}

	γ := struct {
		M OperationCatalogProductAttributeInfoResponse `xml:"catalogProductAttributeInfoResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("urn:Action", α, &γ); err != nil {
		return nil, err
	}
	return γ.M.Result, nil
}

// Retrieve attribute list
func (p *portType) CatalogProductAttributeList(sessionId string, setId int) (*CatalogAttributeEntityArray, error) {
	α := struct {
		M OperationCatalogProductAttributeListRequest `xml:"typens:catalogProductAttributeList"`
	}{
		OperationCatalogProductAttributeListRequest{
			&sessionId,
			&setId,
		},
	}

	γ := struct {
		M OperationCatalogProductAttributeListResponse `xml:"catalogProductAttributeListResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("urn:Action", α, &γ); err != nil {
		return nil, err
	}
	return γ.M.Result, nil
}

// Upload new product image
func (p *portType) CatalogProductAttributeMediaCreate(sessionId string, product string, data *CatalogProductAttributeMediaCreateEntity, storeView string, identifierType string) (string, error) {
	α := struct {
		M OperationCatalogProductAttributeMediaCreateRequest `xml:"typens:catalogProductAttributeMediaCreate"`
	}{
		OperationCatalogProductAttributeMediaCreateRequest{
			&sessionId,
			&product,
			data,
			&storeView,
			&identifierType,
		},
	}

	γ := struct {
		M OperationCatalogProductAttributeMediaCreateResponse `xml:"catalogProductAttributeMediaCreateResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("urn:Action", α, &γ); err != nil {
		return "", err
	}
	return *γ.M.Result, nil
}

// Set/Get current store view
func (p *portType) CatalogProductAttributeMediaCurrentStore(sessionId string, storeView string) (int, error) {
	α := struct {
		M OperationCatalogProductAttributeMediaCurrentStoreRequest `xml:"typens:catalogProductAttributeMediaCurrentStore"`
	}{
		OperationCatalogProductAttributeMediaCurrentStoreRequest{
			&sessionId,
			&storeView,
		},
	}

	γ := struct {
		M OperationCatalogProductAttributeMediaCurrentStoreResponse `xml:"catalogProductAttributeMediaCurrentStoreResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("urn:Action", α, &γ); err != nil {
		return 0, err
	}
	return *γ.M.StoreView, nil
}

// Retrieve product image data
func (p *portType) CatalogProductAttributeMediaInfo(sessionId string, product string, file string, storeView string, identifierType string) (*CatalogProductImageEntity, error) {
	α := struct {
		M OperationCatalogProductAttributeMediaInfoRequest `xml:"typens:catalogProductAttributeMediaInfo"`
	}{
		OperationCatalogProductAttributeMediaInfoRequest{
			&sessionId,
			&product,
			&file,
			&storeView,
			&identifierType,
		},
	}

	γ := struct {
		M OperationCatalogProductAttributeMediaInfoResponse `xml:"catalogProductAttributeMediaInfoResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("urn:Action", α, &γ); err != nil {
		return nil, err
	}
	return γ.M.Result, nil
}

// Retrieve product image list
func (p *portType) CatalogProductAttributeMediaList(sessionId string, product string, storeView string, identifierType string) (*CatalogProductImageEntityArray, error) {
	α := struct {
		M OperationCatalogProductAttributeMediaListRequest `xml:"typens:catalogProductAttributeMediaList"`
	}{
		OperationCatalogProductAttributeMediaListRequest{
			&sessionId,
			&product,
			&storeView,
			&identifierType,
		},
	}

	γ := struct {
		M OperationCatalogProductAttributeMediaListResponse `xml:"catalogProductAttributeMediaListResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("urn:Action", α, &γ); err != nil {
		return nil, err
	}
	return γ.M.Result, nil
}

// Remove product image
func (p *portType) CatalogProductAttributeMediaRemove(sessionId string, product string, file string, identifierType string) (bool, error) {
	α := struct {
		M OperationCatalogProductAttributeMediaRemoveRequest `xml:"typens:catalogProductAttributeMediaRemove"`
	}{
		OperationCatalogProductAttributeMediaRemoveRequest{
			&sessionId,
			&product,
			&file,
			&identifierType,
		},
	}

	γ := struct {
		M OperationCatalogProductAttributeMediaRemoveResponse `xml:"catalogProductAttributeMediaRemoveResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("urn:Action", α, &γ); err != nil {
		return false, err
	}
	return *γ.M.Result, nil
}

// Retrieve product image types
func (p *portType) CatalogProductAttributeMediaTypes(sessionId string, setId string) (*CatalogProductAttributeMediaTypeEntityArray, error) {
	α := struct {
		M OperationCatalogProductAttributeMediaTypesRequest `xml:"typens:catalogProductAttributeMediaTypes"`
	}{
		OperationCatalogProductAttributeMediaTypesRequest{
			&sessionId,
			&setId,
		},
	}

	γ := struct {
		M OperationCatalogProductAttributeMediaTypesResponse `xml:"catalogProductAttributeMediaTypesResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("urn:Action", α, &γ); err != nil {
		return nil, err
	}
	return γ.M.Result, nil
}

// Update product image
func (p *portType) CatalogProductAttributeMediaUpdate(sessionId string, product string, file string, data *CatalogProductAttributeMediaCreateEntity, storeView string, identifierType string) (bool, error) {
	α := struct {
		M OperationCatalogProductAttributeMediaUpdateRequest `xml:"typens:catalogProductAttributeMediaUpdate"`
	}{
		OperationCatalogProductAttributeMediaUpdateRequest{
			&sessionId,
			&product,
			&file,
			data,
			&storeView,
			&identifierType,
		},
	}

	γ := struct {
		M OperationCatalogProductAttributeMediaUpdateResponse `xml:"catalogProductAttributeMediaUpdateResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("urn:Action", α, &γ); err != nil {
		return false, err
	}
	return *γ.M.Result, nil
}

// Retrieve attribute options
func (p *portType) CatalogProductAttributeOptions(sessionId string, attributeId string, storeView string) (*CatalogAttributeOptionEntityArray, error) {
	α := struct {
		M OperationCatalogProductAttributeOptionsRequest `xml:"typens:catalogProductAttributeOptions"`
	}{
		OperationCatalogProductAttributeOptionsRequest{
			&sessionId,
			&attributeId,
			&storeView,
		},
	}

	γ := struct {
		M OperationCatalogProductAttributeOptionsResponse `xml:"catalogProductAttributeOptionsResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("urn:Action", α, &γ); err != nil {
		return nil, err
	}
	return γ.M.Result, nil
}

// Delete attribute
func (p *portType) CatalogProductAttributeRemove(sessionId string, attribute string) (bool, error) {
	α := struct {
		M OperationCatalogProductAttributeRemoveRequest `xml:"typens:catalogProductAttributeRemove"`
	}{
		OperationCatalogProductAttributeRemoveRequest{
			&sessionId,
			&attribute,
		},
	}

	γ := struct {
		M OperationCatalogProductAttributeRemoveResponse `xml:"catalogProductAttributeRemoveResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("urn:Action", α, &γ); err != nil {
		return false, err
	}
	return *γ.M.Result, nil
}

// Remove option from attribute
func (p *portType) CatalogProductAttributeRemoveOption(sessionId string, attribute string, optionId string) (bool, error) {
	α := struct {
		M OperationCatalogProductAttributeRemoveOptionRequest `xml:"typens:catalogProductAttributeRemoveOption"`
	}{
		OperationCatalogProductAttributeRemoveOptionRequest{
			&sessionId,
			&attribute,
			&optionId,
		},
	}

	γ := struct {
		M OperationCatalogProductAttributeRemoveOptionResponse `xml:"catalogProductAttributeRemoveOptionResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("urn:Action", α, &γ); err != nil {
		return false, err
	}
	return *γ.M.Result, nil
}

// Add attribute into attribute set
func (p *portType) CatalogProductAttributeSetAttributeAdd(sessionId string, attributeId string, attributeSetId string, attributeGroupId string, sortOrder string) (bool, error) {
	α := struct {
		M OperationCatalogProductAttributeSetAttributeAddRequest `xml:"typens:catalogProductAttributeSetAttributeAdd"`
	}{
		OperationCatalogProductAttributeSetAttributeAddRequest{
			&sessionId,
			&attributeId,
			&attributeSetId,
			&attributeGroupId,
			&sortOrder,
		},
	}

	γ := struct {
		M OperationCatalogProductAttributeSetAttributeAddResponse `xml:"catalogProductAttributeSetAttributeAddResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("urn:Action", α, &γ); err != nil {
		return false, err
	}
	return *γ.M.IsAdded, nil
}

// Remove attribute from attribute set
func (p *portType) CatalogProductAttributeSetAttributeRemove(sessionId string, attributeId string, attributeSetId string) (bool, error) {
	α := struct {
		M OperationCatalogProductAttributeSetAttributeRemoveRequest `xml:"typens:catalogProductAttributeSetAttributeRemove"`
	}{
		OperationCatalogProductAttributeSetAttributeRemoveRequest{
			&sessionId,
			&attributeId,
			&attributeSetId,
		},
	}

	γ := struct {
		M OperationCatalogProductAttributeSetAttributeRemoveResponse `xml:"catalogProductAttributeSetAttributeRemoveResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("urn:Action", α, &γ); err != nil {
		return false, err
	}
	return *γ.M.IsRemoved, nil
}

// Create product attribute set based on another set
func (p *portType) CatalogProductAttributeSetCreate(sessionId string, attributeSetName string, skeletonSetId string) (int, error) {
	α := struct {
		M OperationCatalogProductAttributeSetCreateRequest `xml:"typens:catalogProductAttributeSetCreate"`
	}{
		OperationCatalogProductAttributeSetCreateRequest{
			&sessionId,
			&attributeSetName,
			&skeletonSetId,
		},
	}

	γ := struct {
		M OperationCatalogProductAttributeSetCreateResponse `xml:"catalogProductAttributeSetCreateResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("urn:Action", α, &γ); err != nil {
		return 0, err
	}
	return *γ.M.SetId, nil
}

// Create group within existing attribute set
func (p *portType) CatalogProductAttributeSetGroupAdd(sessionId string, attributeSetId string, groupName string) (int, error) {
	α := struct {
		M OperationCatalogProductAttributeSetGroupAddRequest `xml:"typens:catalogProductAttributeSetGroupAdd"`
	}{
		OperationCatalogProductAttributeSetGroupAddRequest{
			&sessionId,
			&attributeSetId,
			&groupName,
		},
	}

	γ := struct {
		M OperationCatalogProductAttributeSetGroupAddResponse `xml:"catalogProductAttributeSetGroupAddResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("urn:Action", α, &γ); err != nil {
		return 0, err
	}
	return *γ.M.Result, nil
}

// Remove group from attribute set
func (p *portType) CatalogProductAttributeSetGroupRemove(sessionId string, attributeGroupId string) (bool, error) {
	α := struct {
		M OperationCatalogProductAttributeSetGroupRemoveRequest `xml:"typens:catalogProductAttributeSetGroupRemove"`
	}{
		OperationCatalogProductAttributeSetGroupRemoveRequest{
			&sessionId,
			&attributeGroupId,
		},
	}

	γ := struct {
		M OperationCatalogProductAttributeSetGroupRemoveResponse `xml:"catalogProductAttributeSetGroupRemoveResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("urn:Action", α, &γ); err != nil {
		return false, err
	}
	return *γ.M.Result, nil
}

// Rename existing group
func (p *portType) CatalogProductAttributeSetGroupRename(sessionId string, groupId string, groupName string) (bool, error) {
	α := struct {
		M OperationCatalogProductAttributeSetGroupRenameRequest `xml:"typens:catalogProductAttributeSetGroupRename"`
	}{
		OperationCatalogProductAttributeSetGroupRenameRequest{
			&sessionId,
			&groupId,
			&groupName,
		},
	}

	γ := struct {
		M OperationCatalogProductAttributeSetGroupRenameResponse `xml:"catalogProductAttributeSetGroupRenameResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("urn:Action", α, &γ); err != nil {
		return false, err
	}
	return *γ.M.Result, nil
}

// Retrieve product attribute sets
func (p *portType) CatalogProductAttributeSetList(sessionId string) (*CatalogProductAttributeSetEntityArray, error) {
	α := struct {
		M OperationCatalogProductAttributeSetListRequest `xml:"typens:catalogProductAttributeSetList"`
	}{
		OperationCatalogProductAttributeSetListRequest{
			&sessionId,
		},
	}

	γ := struct {
		M OperationCatalogProductAttributeSetListResponse `xml:"catalogProductAttributeSetListResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("urn:Action", α, &γ); err != nil {
		return nil, err
	}
	return γ.M.Result, nil
}

// Remove product attribute set
func (p *portType) CatalogProductAttributeSetRemove(sessionId string, attributeSetId string, forceProductsRemove string) (bool, error) {
	α := struct {
		M OperationCatalogProductAttributeSetRemoveRequest `xml:"typens:catalogProductAttributeSetRemove"`
	}{
		OperationCatalogProductAttributeSetRemoveRequest{
			&sessionId,
			&attributeSetId,
			&forceProductsRemove,
		},
	}

	γ := struct {
		M OperationCatalogProductAttributeSetRemoveResponse `xml:"catalogProductAttributeSetRemoveResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("urn:Action", α, &γ); err != nil {
		return false, err
	}
	return *γ.M.IsRemoved, nil
}

// Retrieve product tier prices
func (p *portType) CatalogProductAttributeTierPriceInfo(sessionId string, product string, identifierType string) (*CatalogProductTierPriceEntityArray, error) {
	α := struct {
		M OperationCatalogProductAttributeTierPriceInfoRequest `xml:"typens:catalogProductAttributeTierPriceInfo"`
	}{
		OperationCatalogProductAttributeTierPriceInfoRequest{
			&sessionId,
			&product,
			&identifierType,
		},
	}

	γ := struct {
		M OperationCatalogProductAttributeTierPriceInfoResponse `xml:"catalogProductAttributeTierPriceInfoResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("urn:Action", α, &γ); err != nil {
		return nil, err
	}
	return γ.M.Result, nil
}

// Update product tier prices
func (p *portType) CatalogProductAttributeTierPriceUpdate(sessionId string, product string, tier_price *CatalogProductTierPriceEntityArray, identifierType string) (int, error) {
	α := struct {
		M OperationCatalogProductAttributeTierPriceUpdateRequest `xml:"typens:catalogProductAttributeTierPriceUpdate"`
	}{
		OperationCatalogProductAttributeTierPriceUpdateRequest{
			&sessionId,
			&product,
			tier_price,
			&identifierType,
		},
	}

	γ := struct {
		M OperationCatalogProductAttributeTierPriceUpdateResponse `xml:"catalogProductAttributeTierPriceUpdateResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("urn:Action", α, &γ); err != nil {
		return 0, err
	}
	return *γ.M.Result, nil
}

// Get list of possible attribute types
func (p *portType) CatalogProductAttributeTypes(sessionId string) (*CatalogAttributeOptionEntityArray, error) {
	α := struct {
		M OperationCatalogProductAttributeTypesRequest `xml:"typens:catalogProductAttributeTypes"`
	}{
		OperationCatalogProductAttributeTypesRequest{
			&sessionId,
		},
	}

	γ := struct {
		M OperationCatalogProductAttributeTypesResponse `xml:"catalogProductAttributeTypesResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("urn:Action", α, &γ); err != nil {
		return nil, err
	}
	return γ.M.Result, nil
}

// Update attribute
func (p *portType) CatalogProductAttributeUpdate(sessionId string, attribute string, data *CatalogProductAttributeEntityToUpdate) (bool, error) {
	α := struct {
		M OperationCatalogProductAttributeUpdateRequest `xml:"typens:catalogProductAttributeUpdate"`
	}{
		OperationCatalogProductAttributeUpdateRequest{
			&sessionId,
			&attribute,
			data,
		},
	}

	γ := struct {
		M OperationCatalogProductAttributeUpdateResponse `xml:"catalogProductAttributeUpdateResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("urn:Action", α, &γ); err != nil {
		return false, err
	}
	return *γ.M.Result, nil
}

// Create new product and return product id
func (p *portType) CatalogProductCreate(sessionId string, _type string, set string, sku string, productData *CatalogProductCreateEntity, storeView string) (int, error) {
	α := struct {
		M OperationCatalogProductCreateRequest `xml:"typens:catalogProductCreate"`
	}{
		OperationCatalogProductCreateRequest{
			&sessionId,
			&_type,
			&set,
			&sku,
			productData,
			&storeView,
		},
	}

	γ := struct {
		M OperationCatalogProductCreateResponse `xml:"catalogProductCreateResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("urn:Action", α, &γ); err != nil {
		return 0, err
	}
	return *γ.M.Result, nil
}

// Set/Get current store view
func (p *portType) CatalogProductCurrentStore(sessionId string, storeView string) (int, error) {
	α := struct {
		M OperationCatalogProductCurrentStoreRequest `xml:"typens:catalogProductCurrentStore"`
	}{
		OperationCatalogProductCurrentStoreRequest{
			&sessionId,
			&storeView,
		},
	}

	γ := struct {
		M OperationCatalogProductCurrentStoreResponse `xml:"catalogProductCurrentStoreResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("urn:Action", α, &γ); err != nil {
		return 0, err
	}
	return *γ.M.StoreView, nil
}

// Add new custom option into product
func (p *portType) CatalogProductCustomOptionAdd(sessionId string, productId string, data *CatalogProductCustomOptionToAdd, store string) (bool, error) {
	α := struct {
		M OperationCatalogProductCustomOptionAddRequest `xml:"typens:catalogProductCustomOptionAdd"`
	}{
		OperationCatalogProductCustomOptionAddRequest{
			&sessionId,
			&productId,
			data,
			&store,
		},
	}

	γ := struct {
		M OperationCatalogProductCustomOptionAddResponse `xml:"catalogProductCustomOptionAddResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("urn:Action", α, &γ); err != nil {
		return false, err
	}
	return *γ.M.Result, nil
}

// Get full information about custom option in product
func (p *portType) CatalogProductCustomOptionInfo(sessionId string, optionId string, store string) (*CatalogProductCustomOptionInfoEntity, error) {
	α := struct {
		M OperationCatalogProductCustomOptionInfoRequest `xml:"typens:catalogProductCustomOptionInfo"`
	}{
		OperationCatalogProductCustomOptionInfoRequest{
			&sessionId,
			&optionId,
			&store,
		},
	}

	γ := struct {
		M OperationCatalogProductCustomOptionInfoResponse `xml:"catalogProductCustomOptionInfoResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("urn:Action", α, &γ); err != nil {
		return nil, err
	}
	return γ.M.Result, nil
}

// Retrieve list of product custom options
func (p *portType) CatalogProductCustomOptionList(sessionId string, productId string, store string) (*CatalogProductCustomOptionListArray, error) {
	α := struct {
		M OperationCatalogProductCustomOptionListRequest `xml:"typens:catalogProductCustomOptionList"`
	}{
		OperationCatalogProductCustomOptionListRequest{
			&sessionId,
			&productId,
			&store,
		},
	}

	γ := struct {
		M OperationCatalogProductCustomOptionListResponse `xml:"catalogProductCustomOptionListResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("urn:Action", α, &γ); err != nil {
		return nil, err
	}
	return γ.M.Result, nil
}

// Remove custom option
func (p *portType) CatalogProductCustomOptionRemove(sessionId string, optionId string) (bool, error) {
	α := struct {
		M OperationCatalogProductCustomOptionRemoveRequest `xml:"typens:catalogProductCustomOptionRemove"`
	}{
		OperationCatalogProductCustomOptionRemoveRequest{
			&sessionId,
			&optionId,
		},
	}

	γ := struct {
		M OperationCatalogProductCustomOptionRemoveResponse `xml:"catalogProductCustomOptionRemoveResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("urn:Action", α, &γ); err != nil {
		return false, err
	}
	return *γ.M.Result, nil
}

// Get list of available custom option types
func (p *portType) CatalogProductCustomOptionTypes(sessionId string) (*CatalogProductCustomOptionTypesArray, error) {
	α := struct {
		M OperationCatalogProductCustomOptionTypesRequest `xml:"typens:catalogProductCustomOptionTypes"`
	}{
		OperationCatalogProductCustomOptionTypesRequest{
			&sessionId,
		},
	}

	γ := struct {
		M OperationCatalogProductCustomOptionTypesResponse `xml:"catalogProductCustomOptionTypesResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("urn:Action", α, &γ); err != nil {
		return nil, err
	}
	return γ.M.Result, nil
}

// Update product custom option
func (p *portType) CatalogProductCustomOptionUpdate(sessionId string, optionId string, data *CatalogProductCustomOptionToUpdate, store string) (bool, error) {
	α := struct {
		M OperationCatalogProductCustomOptionUpdateRequest `xml:"typens:catalogProductCustomOptionUpdate"`
	}{
		OperationCatalogProductCustomOptionUpdateRequest{
			&sessionId,
			&optionId,
			data,
			&store,
		},
	}

	γ := struct {
		M OperationCatalogProductCustomOptionUpdateResponse `xml:"catalogProductCustomOptionUpdateResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("urn:Action", α, &γ); err != nil {
		return false, err
	}
	return *γ.M.Result, nil
}

// Add new custom option values
func (p *portType) CatalogProductCustomOptionValueAdd(sessionId string, optionId string, data *CatalogProductCustomOptionValueAddArray, store string) (bool, error) {
	α := struct {
		M OperationCatalogProductCustomOptionValueAddRequest `xml:"typens:catalogProductCustomOptionValueAdd"`
	}{
		OperationCatalogProductCustomOptionValueAddRequest{
			&sessionId,
			&optionId,
			data,
			&store,
		},
	}

	γ := struct {
		M OperationCatalogProductCustomOptionValueAddResponse `xml:"catalogProductCustomOptionValueAddResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("urn:Action", α, &γ); err != nil {
		return false, err
	}
	return *γ.M.Result, nil
}

// Retrieve custom option value info
func (p *portType) CatalogProductCustomOptionValueInfo(sessionId string, valueId string, store string) (*CatalogProductCustomOptionValueInfoEntity, error) {
	α := struct {
		M OperationCatalogProductCustomOptionValueInfoRequest `xml:"typens:catalogProductCustomOptionValueInfo"`
	}{
		OperationCatalogProductCustomOptionValueInfoRequest{
			&sessionId,
			&valueId,
			&store,
		},
	}

	γ := struct {
		M OperationCatalogProductCustomOptionValueInfoResponse `xml:"catalogProductCustomOptionValueInfoResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("urn:Action", α, &γ); err != nil {
		return nil, err
	}
	return γ.M.Result, nil
}

// Retrieve custom option values list
func (p *portType) CatalogProductCustomOptionValueList(sessionId string, optionId string, store string) (*CatalogProductCustomOptionValueListArray, error) {
	α := struct {
		M OperationCatalogProductCustomOptionValueListRequest `xml:"typens:catalogProductCustomOptionValueList"`
	}{
		OperationCatalogProductCustomOptionValueListRequest{
			&sessionId,
			&optionId,
			&store,
		},
	}

	γ := struct {
		M OperationCatalogProductCustomOptionValueListResponse `xml:"catalogProductCustomOptionValueListResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("urn:Action", α, &γ); err != nil {
		return nil, err
	}
	return γ.M.Result, nil
}

// Remove value from custom option
func (p *portType) CatalogProductCustomOptionValueRemove(sessionId string, valueId string) (bool, error) {
	α := struct {
		M OperationCatalogProductCustomOptionValueRemoveRequest `xml:"typens:catalogProductCustomOptionValueRemove"`
	}{
		OperationCatalogProductCustomOptionValueRemoveRequest{
			&sessionId,
			&valueId,
		},
	}

	γ := struct {
		M OperationCatalogProductCustomOptionValueRemoveResponse `xml:"catalogProductCustomOptionValueRemoveResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("urn:Action", α, &γ); err != nil {
		return false, err
	}
	return *γ.M.Result, nil
}

// Update custom option value
func (p *portType) CatalogProductCustomOptionValueUpdate(sessionId string, valueId string, data *CatalogProductCustomOptionValueUpdateEntity, storeId string) (bool, error) {
	α := struct {
		M OperationCatalogProductCustomOptionValueUpdateRequest `xml:"typens:catalogProductCustomOptionValueUpdate"`
	}{
		OperationCatalogProductCustomOptionValueUpdateRequest{
			&sessionId,
			&valueId,
			data,
			&storeId,
		},
	}

	γ := struct {
		M OperationCatalogProductCustomOptionValueUpdateResponse `xml:"catalogProductCustomOptionValueUpdateResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("urn:Action", α, &γ); err != nil {
		return false, err
	}
	return *γ.M.Result, nil
}

// Delete product
func (p *portType) CatalogProductDelete(sessionId string, product string, identifierType string) (int, error) {
	α := struct {
		M OperationCatalogProductDeleteRequest `xml:"typens:catalogProductDelete"`
	}{
		OperationCatalogProductDeleteRequest{
			&sessionId,
			&product,
			&identifierType,
		},
	}

	γ := struct {
		M OperationCatalogProductDeleteResponse `xml:"catalogProductDeleteResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("urn:Action", α, &γ); err != nil {
		return 0, err
	}
	return *γ.M.Result, nil
}

// Add links to downloadable product
func (p *portType) CatalogProductDownloadableLinkAdd(sessionId string, productId string, resource *CatalogProductDownloadableLinkAddEntity, resourceType string, store string, identifierType string) (int, error) {
	α := struct {
		M OperationCatalogProductDownloadableLinkAddRequest `xml:"typens:catalogProductDownloadableLinkAdd"`
	}{
		OperationCatalogProductDownloadableLinkAddRequest{
			&sessionId,
			&productId,
			resource,
			&resourceType,
			&store,
			&identifierType,
		},
	}

	γ := struct {
		M OperationCatalogProductDownloadableLinkAddResponse `xml:"catalogProductDownloadableLinkAddResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("urn:Action", α, &γ); err != nil {
		return 0, err
	}
	return *γ.M.Respons, nil
}

// Retrieve list of links and samples for downloadable product
func (p *portType) CatalogProductDownloadableLinkList(sessionId string, productId string, store string, identifierType string) (*CatalogProductDownloadableLinkInfoEntity, error) {
	α := struct {
		M OperationCatalogProductDownloadableLinkListRequest `xml:"typens:catalogProductDownloadableLinkList"`
	}{
		OperationCatalogProductDownloadableLinkListRequest{
			&sessionId,
			&productId,
			&store,
			&identifierType,
		},
	}

	γ := struct {
		M OperationCatalogProductDownloadableLinkListResponse `xml:"catalogProductDownloadableLinkListResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("urn:Action", α, &γ); err != nil {
		return nil, err
	}
	return γ.M.Respons, nil
}

// Remove links and samples from downloadable product
func (p *portType) CatalogProductDownloadableLinkRemove(sessionId string, linkId string, resourceType string) (bool, error) {
	α := struct {
		M OperationCatalogProductDownloadableLinkRemoveRequest `xml:"typens:catalogProductDownloadableLinkRemove"`
	}{
		OperationCatalogProductDownloadableLinkRemoveRequest{
			&sessionId,
			&linkId,
			&resourceType,
		},
	}

	γ := struct {
		M OperationCatalogProductDownloadableLinkRemoveResponse `xml:"catalogProductDownloadableLinkRemoveResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("urn:Action", α, &γ); err != nil {
		return false, err
	}
	return *γ.M.Respons, nil
}

// Get product special price data
func (p *portType) CatalogProductGetSpecialPrice(sessionId string, product string, storeView string, identifierType string) (*CatalogProductSpecialPriceReturnEntity, error) {
	α := struct {
		M OperationCatalogProductGetSpecialPriceRequest `xml:"typens:catalogProductGetSpecialPrice"`
	}{
		OperationCatalogProductGetSpecialPriceRequest{
			&sessionId,
			&product,
			&storeView,
			&identifierType,
		},
	}

	γ := struct {
		M OperationCatalogProductGetSpecialPriceResponse `xml:"catalogProductGetSpecialPriceResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("urn:Action", α, &γ); err != nil {
		return nil, err
	}
	return γ.M.Result, nil
}

// Retrieve product
func (p *portType) CatalogProductInfo(sessionId string, productId string, storeView string, attributes *CatalogProductRequestAttributes, identifierType string) (*CatalogProductReturnEntity, error) {
	α := struct {
		M OperationCatalogProductInfoRequest `xml:"typens:catalogProductInfo"`
	}{
		OperationCatalogProductInfoRequest{
			&sessionId,
			&productId,
			&storeView,
			attributes,
			&identifierType,
		},
	}

	γ := struct {
		M OperationCatalogProductInfoResponse `xml:"catalogProductInfoResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("urn:Action", α, &γ); err != nil {
		return nil, err
	}
	return γ.M.Info, nil
}

// Assign product link
func (p *portType) CatalogProductLinkAssign(sessionId string, _type string, product string, linkedProduct string, data *CatalogProductLinkEntity, identifierType string) (bool, error) {
	α := struct {
		M OperationCatalogProductLinkAssignRequest `xml:"typens:catalogProductLinkAssign"`
	}{
		OperationCatalogProductLinkAssignRequest{
			&sessionId,
			&_type,
			&product,
			&linkedProduct,
			data,
			&identifierType,
		},
	}

	γ := struct {
		M OperationCatalogProductLinkAssignResponse `xml:"catalogProductLinkAssignResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("urn:Action", α, &γ); err != nil {
		return false, err
	}
	return *γ.M.Result, nil
}

// Retrieve product link type attributes
func (p *portType) CatalogProductLinkAttributes(sessionId string, _type string) (*CatalogProductLinkAttributeEntityArray, error) {
	α := struct {
		M OperationCatalogProductLinkAttributesRequest `xml:"typens:catalogProductLinkAttributes"`
	}{
		OperationCatalogProductLinkAttributesRequest{
			&sessionId,
			&_type,
		},
	}

	γ := struct {
		M OperationCatalogProductLinkAttributesResponse `xml:"catalogProductLinkAttributesResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("urn:Action", α, &γ); err != nil {
		return nil, err
	}
	return γ.M.Result, nil
}

// Retrieve linked products
func (p *portType) CatalogProductLinkList(sessionId string, _type string, product string, identifierType string) (*CatalogProductLinkEntityArray, error) {
	α := struct {
		M OperationCatalogProductLinkListRequest `xml:"typens:catalogProductLinkList"`
	}{
		OperationCatalogProductLinkListRequest{
			&sessionId,
			&_type,
			&product,
			&identifierType,
		},
	}

	γ := struct {
		M OperationCatalogProductLinkListResponse `xml:"catalogProductLinkListResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("urn:Action", α, &γ); err != nil {
		return nil, err
	}
	return γ.M.Result, nil
}

// Remove product link
func (p *portType) CatalogProductLinkRemove(sessionId string, _type string, product string, linkedProduct string, identifierType string) (bool, error) {
	α := struct {
		M OperationCatalogProductLinkRemoveRequest `xml:"typens:catalogProductLinkRemove"`
	}{
		OperationCatalogProductLinkRemoveRequest{
			&sessionId,
			&_type,
			&product,
			&linkedProduct,
			&identifierType,
		},
	}

	γ := struct {
		M OperationCatalogProductLinkRemoveResponse `xml:"catalogProductLinkRemoveResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("urn:Action", α, &γ); err != nil {
		return false, err
	}
	return *γ.M.Result, nil
}

// Retrieve product link types
func (p *portType) CatalogProductLinkTypes(sessionId string) (*ArrayOfString, error) {
	α := struct {
		M OperationCatalogProductLinkTypesRequest `xml:"typens:catalogProductLinkTypes"`
	}{
		OperationCatalogProductLinkTypesRequest{
			&sessionId,
		},
	}

	γ := struct {
		M OperationCatalogProductLinkTypesResponse `xml:"catalogProductLinkTypesResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("urn:Action", α, &γ); err != nil {
		return nil, err
	}
	return γ.M.Result, nil
}

// Update product link
func (p *portType) CatalogProductLinkUpdate(sessionId string, _type string, product string, linkedProduct string, data *CatalogProductLinkEntity, identifierType string) (bool, error) {
	α := struct {
		M OperationCatalogProductLinkUpdateRequest `xml:"typens:catalogProductLinkUpdate"`
	}{
		OperationCatalogProductLinkUpdateRequest{
			&sessionId,
			&_type,
			&product,
			&linkedProduct,
			data,
			&identifierType,
		},
	}

	γ := struct {
		M OperationCatalogProductLinkUpdateResponse `xml:"catalogProductLinkUpdateResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("urn:Action", α, &γ); err != nil {
		return false, err
	}
	return *γ.M.Result, nil
}

// Retrieve products list by filters
func (p *portType) CatalogProductList(sessionId string, filters *Filters, storeView string) (*CatalogProductEntityArray, error) {
	α := struct {
		M OperationCatalogProductListRequest `xml:"typens:catalogProductList"`
	}{
		OperationCatalogProductListRequest{
			&sessionId,
			filters,
			&storeView,
		},
	}

	γ := struct {
		M OperationCatalogProductListResponse `xml:"catalogProductListResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("urn:Action", α, &γ); err != nil {
		return nil, err
	}
	return γ.M.StoreView, nil
}

// Get list of additional attributes which are not in default create/update
// list
func (p *portType) CatalogProductListOfAdditionalAttributes(sessionId string, productType string, attributeSetId string) (*CatalogAttributeEntityArray, error) {
	α := struct {
		M OperationCatalogProductListOfAdditionalAttributesRequest `xml:"typens:catalogProductListOfAdditionalAttributes"`
	}{
		OperationCatalogProductListOfAdditionalAttributesRequest{
			&sessionId,
			&productType,
			&attributeSetId,
		},
	}

	γ := struct {
		M OperationCatalogProductListOfAdditionalAttributesResponse `xml:"catalogProductListOfAdditionalAttributesResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("urn:Action", α, &γ); err != nil {
		return nil, err
	}
	return γ.M.Result, nil
}

// Product multi update
func (p *portType) CatalogProductMultiUpdate(sessionId string, productIds *ArrayOfString, productData *CatalogProductCreateEntityArray, store string, identifierType string) (bool, error) {
	α := struct {
		M OperationCatalogProductMultiUpdateRequest `xml:"typens:catalogProductMultiUpdate"`
	}{
		OperationCatalogProductMultiUpdateRequest{
			&sessionId,
			productIds,
			productData,
			&store,
			&identifierType,
		},
	}

	γ := struct {
		M OperationCatalogProductMultiUpdateResponse `xml:"catalogProductMultiUpdateResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("urn:Action", α, &γ); err != nil {
		return false, err
	}
	return *γ.M.Result, nil
}

// Update product special price
func (p *portType) CatalogProductSetSpecialPrice(sessionId string, product string, specialPrice string, fromDate string, toDate string, storeView string, identifierType string) (int, error) {
	α := struct {
		M OperationCatalogProductSetSpecialPriceRequest `xml:"typens:catalogProductSetSpecialPrice"`
	}{
		OperationCatalogProductSetSpecialPriceRequest{
			&sessionId,
			&product,
			&specialPrice,
			&fromDate,
			&toDate,
			&storeView,
			&identifierType,
		},
	}

	γ := struct {
		M OperationCatalogProductSetSpecialPriceResponse `xml:"catalogProductSetSpecialPriceResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("urn:Action", α, &γ); err != nil {
		return 0, err
	}
	return *γ.M.Result, nil
}

// Add tag(s) to product
func (p *portType) CatalogProductTagAdd(sessionId string, data *CatalogProductTagAddEntity) (*AssociativeArray, error) {
	α := struct {
		M OperationCatalogProductTagAddRequest `xml:"typens:catalogProductTagAdd"`
	}{
		OperationCatalogProductTagAddRequest{
			&sessionId,
			data,
		},
	}

	γ := struct {
		M OperationCatalogProductTagAddResponse `xml:"catalogProductTagAddResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("urn:Action", α, &γ); err != nil {
		return nil, err
	}
	return γ.M.Result, nil
}

// Retrieve product tag info
func (p *portType) CatalogProductTagInfo(sessionId string, tagId string, store string) (*CatalogProductTagInfoEntity, error) {
	α := struct {
		M OperationCatalogProductTagInfoRequest `xml:"typens:catalogProductTagInfo"`
	}{
		OperationCatalogProductTagInfoRequest{
			&sessionId,
			&tagId,
			&store,
		},
	}

	γ := struct {
		M OperationCatalogProductTagInfoResponse `xml:"catalogProductTagInfoResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("urn:Action", α, &γ); err != nil {
		return nil, err
	}
	return γ.M.Result, nil
}

// Retrieve list of tags by product
func (p *portType) CatalogProductTagList(sessionId string, productId string, store string) (*CatalogProductTagListEntityArray, error) {
	α := struct {
		M OperationCatalogProductTagListRequest `xml:"typens:catalogProductTagList"`
	}{
		OperationCatalogProductTagListRequest{
			&sessionId,
			&productId,
			&store,
		},
	}

	γ := struct {
		M OperationCatalogProductTagListResponse `xml:"catalogProductTagListResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("urn:Action", α, &γ); err != nil {
		return nil, err
	}
	return γ.M.Result, nil
}

// Remove product tag
func (p *portType) CatalogProductTagRemove(sessionId string, tagId string) (bool, error) {
	α := struct {
		M OperationCatalogProductTagRemoveRequest `xml:"typens:catalogProductTagRemove"`
	}{
		OperationCatalogProductTagRemoveRequest{
			&sessionId,
			&tagId,
		},
	}

	γ := struct {
		M OperationCatalogProductTagRemoveResponse `xml:"catalogProductTagRemoveResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("urn:Action", α, &γ); err != nil {
		return false, err
	}
	return *γ.M.Result, nil
}

// Update product tag
func (p *portType) CatalogProductTagUpdate(sessionId string, tagId string, data *CatalogProductTagUpdateEntity, store string) (bool, error) {
	α := struct {
		M OperationCatalogProductTagUpdateRequest `xml:"typens:catalogProductTagUpdate"`
	}{
		OperationCatalogProductTagUpdateRequest{
			&sessionId,
			&tagId,
			data,
			&store,
		},
	}

	γ := struct {
		M OperationCatalogProductTagUpdateResponse `xml:"catalogProductTagUpdateResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("urn:Action", α, &γ); err != nil {
		return false, err
	}
	return *γ.M.Result, nil
}

// Retrieve product types
func (p *portType) CatalogProductTypeList(sessionId string) (*CatalogProductTypeEntityArray, error) {
	α := struct {
		M OperationCatalogProductTypeListRequest `xml:"typens:catalogProductTypeList"`
	}{
		OperationCatalogProductTypeListRequest{
			&sessionId,
		},
	}

	γ := struct {
		M OperationCatalogProductTypeListResponse `xml:"catalogProductTypeListResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("urn:Action", α, &γ); err != nil {
		return nil, err
	}
	return γ.M.Result, nil
}

// Update product
func (p *portType) CatalogProductUpdate(sessionId string, product string, productData *CatalogProductCreateEntity, storeView string, identifierType string) (bool, error) {
	α := struct {
		M OperationCatalogProductUpdateRequest `xml:"typens:catalogProductUpdate"`
	}{
		OperationCatalogProductUpdateRequest{
			&sessionId,
			&product,
			productData,
			&storeView,
			&identifierType,
		},
	}

	γ := struct {
		M OperationCatalogProductUpdateResponse `xml:"catalogProductUpdateResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("urn:Action", α, &γ); err != nil {
		return false, err
	}
	return *γ.M.Result, nil
}

// Create customer address
func (p *portType) CustomerAddressCreate(sessionId string, customerId int, addressData *CustomerAddressEntityCreate) (int, error) {
	α := struct {
		M OperationCustomerAddressCreateRequest `xml:"typens:customerAddressCreate"`
	}{
		OperationCustomerAddressCreateRequest{
			&sessionId,
			&customerId,
			addressData,
		},
	}

	γ := struct {
		M OperationCustomerAddressCreateResponse `xml:"customerAddressCreateResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("urn:Action", α, &γ); err != nil {
		return 0, err
	}
	return *γ.M.Result, nil
}

// Delete customer address
func (p *portType) CustomerAddressDelete(sessionId string, addressId int) (bool, error) {
	α := struct {
		M OperationCustomerAddressDeleteRequest `xml:"typens:customerAddressDelete"`
	}{
		OperationCustomerAddressDeleteRequest{
			&sessionId,
			&addressId,
		},
	}

	γ := struct {
		M OperationCustomerAddressDeleteResponse `xml:"customerAddressDeleteResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("urn:Action", α, &γ); err != nil {
		return false, err
	}
	return *γ.M.Info, nil
}

// Retrieve customer address data
func (p *portType) CustomerAddressInfo(sessionId string, addressId int) (*CustomerAddressEntityItem, error) {
	α := struct {
		M OperationCustomerAddressInfoRequest `xml:"typens:customerAddressInfo"`
	}{
		OperationCustomerAddressInfoRequest{
			&sessionId,
			&addressId,
		},
	}

	γ := struct {
		M OperationCustomerAddressInfoResponse `xml:"customerAddressInfoResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("urn:Action", α, &γ); err != nil {
		return nil, err
	}
	return γ.M.Info, nil
}

// Retrieve customer addresses
func (p *portType) CustomerAddressList(sessionId string, customerId int) (*CustomerAddressEntityArray, error) {
	α := struct {
		M OperationCustomerAddressListRequest `xml:"typens:customerAddressList"`
	}{
		OperationCustomerAddressListRequest{
			&sessionId,
			&customerId,
		},
	}

	γ := struct {
		M OperationCustomerAddressListResponse `xml:"customerAddressListResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("urn:Action", α, &γ); err != nil {
		return nil, err
	}
	return γ.M.Result, nil
}

// Update customer address data
func (p *portType) CustomerAddressUpdate(sessionId string, addressId int, addressData *CustomerAddressEntityCreate) (bool, error) {
	α := struct {
		M OperationCustomerAddressUpdateRequest `xml:"typens:customerAddressUpdate"`
	}{
		OperationCustomerAddressUpdateRequest{
			&sessionId,
			&addressId,
			addressData,
		},
	}

	γ := struct {
		M OperationCustomerAddressUpdateResponse `xml:"customerAddressUpdateResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("urn:Action", α, &γ); err != nil {
		return false, err
	}
	return *γ.M.Info, nil
}

// Create customer
func (p *portType) CustomerCustomerCreate(sessionId string, customerData *CustomerCustomerEntityToCreate) (int, error) {
	α := struct {
		M OperationCustomerCustomerCreateRequest `xml:"typens:customerCustomerCreate"`
	}{
		OperationCustomerCustomerCreateRequest{
			&sessionId,
			customerData,
		},
	}

	γ := struct {
		M OperationCustomerCustomerCreateResponse `xml:"customerCustomerCreateResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("urn:Action", α, &γ); err != nil {
		return 0, err
	}
	return *γ.M.Result, nil
}

// Delete customer
func (p *portType) CustomerCustomerDelete(sessionId string, customerId int) (bool, error) {
	α := struct {
		M OperationCustomerCustomerDeleteRequest `xml:"typens:customerCustomerDelete"`
	}{
		OperationCustomerCustomerDeleteRequest{
			&sessionId,
			&customerId,
		},
	}

	γ := struct {
		M OperationCustomerCustomerDeleteResponse `xml:"customerCustomerDeleteResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("urn:Action", α, &γ); err != nil {
		return false, err
	}
	return *γ.M.Result, nil
}

// Retrieve customer data
func (p *portType) CustomerCustomerInfo(sessionId string, customerId int, attributes *ArrayOfString) (*CustomerCustomerEntity, error) {
	α := struct {
		M OperationCustomerCustomerInfoRequest `xml:"typens:customerCustomerInfo"`
	}{
		OperationCustomerCustomerInfoRequest{
			&sessionId,
			&customerId,
			attributes,
		},
	}

	γ := struct {
		M OperationCustomerCustomerInfoResponse `xml:"customerCustomerInfoResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("urn:Action", α, &γ); err != nil {
		return nil, err
	}
	return γ.M.CustomerInfo, nil
}

// Retrieve customers
func (p *portType) CustomerCustomerList(sessionId string, filters *Filters) (*CustomerCustomerEntityArray, error) {
	α := struct {
		M OperationCustomerCustomerListRequest `xml:"typens:customerCustomerList"`
	}{
		OperationCustomerCustomerListRequest{
			&sessionId,
			filters,
		},
	}

	γ := struct {
		M OperationCustomerCustomerListResponse `xml:"customerCustomerListResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("urn:Action", α, &γ); err != nil {
		return nil, err
	}
	return γ.M.StoreView, nil
}

// Update customer data
func (p *portType) CustomerCustomerUpdate(sessionId string, customerId int, customerData *CustomerCustomerEntityToCreate) (bool, error) {
	α := struct {
		M OperationCustomerCustomerUpdateRequest `xml:"typens:customerCustomerUpdate"`
	}{
		OperationCustomerCustomerUpdateRequest{
			&sessionId,
			&customerId,
			customerData,
		},
	}

	γ := struct {
		M OperationCustomerCustomerUpdateResponse `xml:"customerCustomerUpdateResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("urn:Action", α, &γ); err != nil {
		return false, err
	}
	return *γ.M.Result, nil
}

// Retrieve customer groups
func (p *portType) CustomerGroupList(sessionId string) (*CustomerGroupEntityArray, error) {
	α := struct {
		M OperationCustomerGroupListRequest `xml:"typens:customerGroupList"`
	}{
		OperationCustomerGroupListRequest{
			&sessionId,
		},
	}

	γ := struct {
		M OperationCustomerGroupListResponse `xml:"customerGroupListResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("urn:Action", α, &γ); err != nil {
		return nil, err
	}
	return γ.M.Result, nil
}

// List of countries
func (p *portType) DirectoryCountryList(sessionId string) (*DirectoryCountryEntityArray, error) {
	α := struct {
		M OperationDirectoryCountryListRequest `xml:"typens:directoryCountryList"`
	}{
		OperationDirectoryCountryListRequest{
			&sessionId,
		},
	}

	γ := struct {
		M OperationDirectoryCountryListResponse `xml:"directoryCountryListResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("urn:Action", α, &γ); err != nil {
		return nil, err
	}
	return γ.M.Countries, nil
}

// List of regions in specified country
func (p *portType) DirectoryRegionList(sessionId string, country string) (*DirectoryRegionEntityArray, error) {
	α := struct {
		M OperationDirectoryRegionListRequest `xml:"typens:directoryRegionList"`
	}{
		OperationDirectoryRegionListRequest{
			&sessionId,
			&country,
		},
	}

	γ := struct {
		M OperationDirectoryRegionListResponse `xml:"directoryRegionListResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("urn:Action", α, &γ); err != nil {
		return nil, err
	}
	return γ.M.Countries, nil
}

// End web service session
func (p *portType) EndSession(sessionId string) (bool, error) {
	α := struct {
		M OperationEndSession `xml:"typens:endSession"`
	}{
		OperationEndSession{
			&sessionId,
		},
	}

	γ := struct {
		M OperationEndSessionResponse `xml:"endSessionResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("urn:Action", α, &γ); err != nil {
		return false, err
	}
	return *γ.M.EndSessionReturn, nil
}

// Set a gift message to the cart
func (p *portType) GiftMessageSetForQuote(sessionId string, quoteId string, giftMessage *GiftMessageEntity, storeId string) (*GiftMessageResponse, error) {
	α := struct {
		M OperationGiftMessageForQuoteRequest `xml:"typens:giftMessageSetForQuote"`
	}{
		OperationGiftMessageForQuoteRequest{
			&sessionId,
			&quoteId,
			giftMessage,
			&storeId,
		},
	}

	γ := struct {
		M OperationGiftMessageForQuoteResponse `xml:"giftMessageSetForQuoteResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("urn:Action", α, &γ); err != nil {
		return nil, err
	}
	return γ.M.Result, nil
}

// Setting a gift messages to the quote item
func (p *portType) GiftMessageSetForQuoteItem(sessionId string, quoteItemId string, giftMessage *GiftMessageEntity, storeId string) (*GiftMessageResponse, error) {
	α := struct {
		M OperationGiftMessageForQuoteItemRequest `xml:"typens:giftMessageSetForQuoteItem"`
	}{
		OperationGiftMessageForQuoteItemRequest{
			&sessionId,
			&quoteItemId,
			giftMessage,
			&storeId,
		},
	}

	γ := struct {
		M OperationGiftMessageForQuoteItemResponse `xml:"giftMessageSetForQuoteItemResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("urn:Action", α, &γ); err != nil {
		return nil, err
	}
	return γ.M.Result, nil
}

// Setting a gift messages to the quote items by products
func (p *portType) GiftMessageSetForQuoteProduct(sessionId string, quoteId string, productsAndMessages *GiftMessageAssociativeProductsEntityArray, storeId string) (*GiftMessageResponseArray, error) {
	α := struct {
		M OperationGiftMessageForQuoteProductRequest `xml:"typens:giftMessageSetForQuoteProduct"`
	}{
		OperationGiftMessageForQuoteProductRequest{
			&sessionId,
			&quoteId,
			productsAndMessages,
			&storeId,
		},
	}

	γ := struct {
		M OperationGiftMessageForQuoteProductResponse `xml:"giftMessageSetForQuoteProductResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("urn:Action", α, &γ); err != nil {
		return nil, err
	}
	return γ.M.Result, nil
}

// List of global faults
func (p *portType) GlobalFaults(sessionId string) (*ArrayOfExistsFaltures, error) {
	α := struct {
		M OperationGlobalFaults `xml:"typens:globalFaults"`
	}{
		OperationGlobalFaults{
			&sessionId,
		},
	}

	γ := struct {
		M OperationGlobalFaultsResponse `xml:"globalFaultsResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("urn:Action", α, &γ); err != nil {
		return nil, err
	}
	return γ.M.GlobalFaultsReturn, nil
}

// Login user and retrive session id
func (p *portType) Login(username string, apiKey string) (string, error) {
	α := struct {
		M OperationLogin `xml:"typens:login"`
	}{
		OperationLogin{
			&username,
			&apiKey,
		},
	}

	γ := struct {
		M OperationLoginResponse `xml:"loginResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("urn:Action", α, &γ); err != nil {
		return "", err
	}
	return *γ.M.LoginReturn, nil
}

// Info about current Magento installation
func (p *portType) MagentoInfo(sessionId string) (*MagentoInfoEntity, error) {
	α := struct {
		M OperationMagentoInfoRequest `xml:"typens:magentoInfo"`
	}{
		OperationMagentoInfoRequest{
			&sessionId,
		},
	}

	γ := struct {
		M OperationMagentoInfoResponse `xml:"magentoInfoResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("urn:Action", α, &γ); err != nil {
		return nil, err
	}
	return γ.M.Info, nil
}

// List of resource faults
func (p *portType) ResourceFaults(resourceName string, sessionId string) (*ArrayOfExistsFaltures, error) {
	α := struct {
		M OperationResourceFaults `xml:"typens:resourceFaults"`
	}{
		OperationResourceFaults{
			&resourceName,
			&sessionId,
		},
	}

	γ := struct {
		M OperationResourceFaultsResponse `xml:"resourceFaultsResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("urn:Action", α, &γ); err != nil {
		return nil, err
	}
	return γ.M.ResourceFaultsReturn, nil
}

// List of available resources
func (p *portType) Resources(sessionId string) (*ArrayOfApis, error) {
	α := struct {
		M OperationResourcesRequest `xml:"typens:resources"`
	}{
		OperationResourcesRequest{
			&sessionId,
		},
	}

	γ := struct {
		M OperationResourcesResponse `xml:"resourcesResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("urn:Action", α, &γ); err != nil {
		return nil, err
	}
	return γ.M.ResourcesReturn, nil
}

// Add comment to order
func (p *portType) SalesOrderAddComment(sessionId string, orderIncrementId string, status string, comment string, notify string) (bool, error) {
	α := struct {
		M OperationSalesOrderAddCommentRequest `xml:"typens:salesOrderAddComment"`
	}{
		OperationSalesOrderAddCommentRequest{
			&sessionId,
			&orderIncrementId,
			&status,
			&comment,
			&notify,
		},
	}

	γ := struct {
		M OperationSalesOrderAddCommentResponse `xml:"salesOrderAddCommentResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("urn:Action", α, &γ); err != nil {
		return false, err
	}
	return *γ.M.Result, nil
}

// Cancel order
func (p *portType) SalesOrderCancel(sessionId string, orderIncrementId string) (bool, error) {
	α := struct {
		M OperationSalesOrderCancelRequest `xml:"typens:salesOrderCancel"`
	}{
		OperationSalesOrderCancelRequest{
			&sessionId,
			&orderIncrementId,
		},
	}

	γ := struct {
		M OperationSalesOrderCancelResponse `xml:"salesOrderCancelResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("urn:Action", α, &γ); err != nil {
		return false, err
	}
	return *γ.M.Result, nil
}

// Add new comment to shipment
func (p *portType) SalesOrderCreditmemoAddComment(sessionId string, creditmemoIncrementId string, comment string, notifyCustomer int, includeComment int) (bool, error) {
	α := struct {
		M OperationSalesOrderCreditmemoAddCommentRequest `xml:"typens:salesOrderCreditmemoAddComment"`
	}{
		OperationSalesOrderCreditmemoAddCommentRequest{
			&sessionId,
			&creditmemoIncrementId,
			&comment,
			&notifyCustomer,
			&includeComment,
		},
	}

	γ := struct {
		M OperationSalesOrderCreditmemoAddCommentResponse `xml:"salesOrderCreditmemoAddCommentResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("urn:Action", α, &γ); err != nil {
		return false, err
	}
	return *γ.M.Result, nil
}

// Cancel creditmemo
func (p *portType) SalesOrderCreditmemoCancel(sessionId string, creditmemoIncrementId string) (bool, error) {
	α := struct {
		M OperationSalesOrderCreditmemoCancelRequest `xml:"typens:salesOrderCreditmemoCancel"`
	}{
		OperationSalesOrderCreditmemoCancelRequest{
			&sessionId,
			&creditmemoIncrementId,
		},
	}

	γ := struct {
		M OperationSalesOrderCreditmemoCancelResponse `xml:"salesOrderCreditmemoCancelResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("urn:Action", α, &γ); err != nil {
		return false, err
	}
	return *γ.M.Result, nil
}

// Create new creditmemo for order
func (p *portType) SalesOrderCreditmemoCreate(sessionId string, orderIncrementId string, creditmemoData *SalesOrderCreditmemoData, comment string, notifyCustomer int, includeComment int, refundToStoreCreditAmount string) (string, error) {
	α := struct {
		M OperationSalesOrderCreditmemoCreateRequest `xml:"typens:salesOrderCreditmemoCreate"`
	}{
		OperationSalesOrderCreditmemoCreateRequest{
			&sessionId,
			&orderIncrementId,
			creditmemoData,
			&comment,
			&notifyCustomer,
			&includeComment,
			&refundToStoreCreditAmount,
		},
	}

	γ := struct {
		M OperationSalesOrderCreditmemoCreateResponse `xml:"salesOrderCreditmemoCreateResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("urn:Action", α, &γ); err != nil {
		return "", err
	}
	return *γ.M.Result, nil
}

// Retrieve creditmemo information
func (p *portType) SalesOrderCreditmemoInfo(sessionId string, creditmemoIncrementId string) (*SalesOrderCreditmemoEntity, error) {
	α := struct {
		M OperationSalesOrderCreditmemoInfoRequest `xml:"typens:salesOrderCreditmemoInfo"`
	}{
		OperationSalesOrderCreditmemoInfoRequest{
			&sessionId,
			&creditmemoIncrementId,
		},
	}

	γ := struct {
		M OperationSalesOrderCreditmemoInfoResponse `xml:"salesOrderCreditmemoInfoResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("urn:Action", α, &γ); err != nil {
		return nil, err
	}
	return γ.M.Result, nil
}

// Retrieve list of creditmemos by filters
func (p *portType) SalesOrderCreditmemoList(sessionId string, filters *Filters) (*SalesOrderCreditmemoEntityArray, error) {
	α := struct {
		M OperationSalesOrderCreditmemoListRequest `xml:"typens:salesOrderCreditmemoList"`
	}{
		OperationSalesOrderCreditmemoListRequest{
			&sessionId,
			filters,
		},
	}

	γ := struct {
		M OperationSalesOrderCreditmemoListResponse `xml:"salesOrderCreditmemoListResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("urn:Action", α, &γ); err != nil {
		return nil, err
	}
	return γ.M.Result, nil
}

// Hold order
func (p *portType) SalesOrderHold(sessionId string, orderIncrementId string) (bool, error) {
	α := struct {
		M OperationSalesOrderHoldRequest `xml:"typens:salesOrderHold"`
	}{
		OperationSalesOrderHoldRequest{
			&sessionId,
			&orderIncrementId,
		},
	}

	γ := struct {
		M OperationSalesOrderHoldResponse `xml:"salesOrderHoldResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("urn:Action", α, &γ); err != nil {
		return false, err
	}
	return *γ.M.Result, nil
}

// Retrieve order information
func (p *portType) SalesOrderInfo(sessionId string, orderIncrementId string) (*SalesOrderEntity, error) {
	α := struct {
		M OperationSalesOrderInfoRequest `xml:"typens:salesOrderInfo"`
	}{
		OperationSalesOrderInfoRequest{
			&sessionId,
			&orderIncrementId,
		},
	}

	γ := struct {
		M OperationSalesOrderInfoResponse `xml:"salesOrderInfoResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("urn:Action", α, &γ); err != nil {
		return nil, err
	}
	return γ.M.Result, nil
}

// Add new comment to shipment
func (p *portType) SalesOrderInvoiceAddComment(sessionId string, invoiceIncrementId string, comment string, email string, includeComment string) (bool, error) {
	α := struct {
		M OperationSalesOrderInvoiceAddCommentRequest `xml:"typens:salesOrderInvoiceAddComment"`
	}{
		OperationSalesOrderInvoiceAddCommentRequest{
			&sessionId,
			&invoiceIncrementId,
			&comment,
			&email,
			&includeComment,
		},
	}

	γ := struct {
		M OperationSalesOrderInvoiceAddCommentResponse `xml:"salesOrderInvoiceAddCommentResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("urn:Action", α, &γ); err != nil {
		return false, err
	}
	return *γ.M.Result, nil
}

// Cancel invoice
func (p *portType) SalesOrderInvoiceCancel(sessionId string, invoiceIncrementId string) (bool, error) {
	α := struct {
		M OperationSalesOrderInvoiceCancelRequest `xml:"typens:salesOrderInvoiceCancel"`
	}{
		OperationSalesOrderInvoiceCancelRequest{
			&sessionId,
			&invoiceIncrementId,
		},
	}

	γ := struct {
		M OperationSalesOrderInvoiceCancelResponse `xml:"salesOrderInvoiceCancelResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("urn:Action", α, &γ); err != nil {
		return false, err
	}
	return *γ.M.Result, nil
}

// Capture invoice
func (p *portType) SalesOrderInvoiceCapture(sessionId string, invoiceIncrementId string) (bool, error) {
	α := struct {
		M OperationSalesOrderInvoiceCaptureRequest `xml:"typens:salesOrderInvoiceCapture"`
	}{
		OperationSalesOrderInvoiceCaptureRequest{
			&sessionId,
			&invoiceIncrementId,
		},
	}

	γ := struct {
		M OperationSalesOrderInvoiceCaptureResponse `xml:"salesOrderInvoiceCaptureResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("urn:Action", α, &γ); err != nil {
		return false, err
	}
	return *γ.M.Result, nil
}

// Create new invoice for order
func (p *portType) SalesOrderInvoiceCreate(sessionId string, invoiceIncrementId string, itemsQty *OrderItemIdQtyArray, comment string, email string, includeComment string) (string, error) {
	α := struct {
		M OperationSalesOrderInvoiceCreateRequest `xml:"typens:salesOrderInvoiceCreate"`
	}{
		OperationSalesOrderInvoiceCreateRequest{
			&sessionId,
			&invoiceIncrementId,
			itemsQty,
			&comment,
			&email,
			&includeComment,
		},
	}

	γ := struct {
		M OperationSalesOrderInvoiceCreateResponse `xml:"salesOrderInvoiceCreateResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("urn:Action", α, &γ); err != nil {
		return "", err
	}
	return *γ.M.Result, nil
}

// Retrieve invoice information
func (p *portType) SalesOrderInvoiceInfo(sessionId string, invoiceIncrementId string) (*SalesOrderInvoiceEntity, error) {
	α := struct {
		M OperationSalesOrderInvoiceInfoRequest `xml:"typens:salesOrderInvoiceInfo"`
	}{
		OperationSalesOrderInvoiceInfoRequest{
			&sessionId,
			&invoiceIncrementId,
		},
	}

	γ := struct {
		M OperationSalesOrderInvoiceInfoResponse `xml:"salesOrderInvoiceInfoResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("urn:Action", α, &γ); err != nil {
		return nil, err
	}
	return γ.M.Result, nil
}

// Retrieve list of invoices by filters
func (p *portType) SalesOrderInvoiceList(sessionId string, filters *Filters) (*SalesOrderInvoiceEntityArray, error) {
	α := struct {
		M OperationSalesOrderInvoiceListRequest `xml:"typens:salesOrderInvoiceList"`
	}{
		OperationSalesOrderInvoiceListRequest{
			&sessionId,
			filters,
		},
	}

	γ := struct {
		M OperationSalesOrderInvoiceListResponse `xml:"salesOrderInvoiceListResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("urn:Action", α, &γ); err != nil {
		return nil, err
	}
	return γ.M.Result, nil
}

// Void invoice
func (p *portType) SalesOrderInvoiceVoid(sessionId string, invoiceIncrementId string) (bool, error) {
	α := struct {
		M OperationSalesOrderInvoiceVoidRequest `xml:"typens:salesOrderInvoiceVoid"`
	}{
		OperationSalesOrderInvoiceVoidRequest{
			&sessionId,
			&invoiceIncrementId,
		},
	}

	γ := struct {
		M OperationSalesOrderInvoiceVoidResponse `xml:"salesOrderInvoiceVoidResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("urn:Action", α, &γ); err != nil {
		return false, err
	}
	return *γ.M.Result, nil
}

// Retrieve list of orders by filters
func (p *portType) SalesOrderList(sessionId string, filters *Filters) (*SalesOrderListEntityArray, error) {
	α := struct {
		M OperationSalesOrderListRequest `xml:"typens:salesOrderList"`
	}{
		OperationSalesOrderListRequest{
			&sessionId,
			filters,
		},
	}

	γ := struct {
		M OperationSalesOrderListResponse `xml:"salesOrderListResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("urn:Action", α, &γ); err != nil {
		return nil, err
	}
	return γ.M.Result, nil
}

// Add new comment to shipment
func (p *portType) SalesOrderShipmentAddComment(sessionId string, shipmentIncrementId string, comment string, email string, includeInEmail string) (bool, error) {
	α := struct {
		M OperationSalesOrderShipmentAddCommentRequest `xml:"typens:salesOrderShipmentAddComment"`
	}{
		OperationSalesOrderShipmentAddCommentRequest{
			&sessionId,
			&shipmentIncrementId,
			&comment,
			&email,
			&includeInEmail,
		},
	}

	γ := struct {
		M OperationSalesOrderShipmentAddCommentResponse `xml:"salesOrderShipmentAddCommentResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("urn:Action", α, &γ); err != nil {
		return false, err
	}
	return *γ.M.ShipmentIncrementId, nil
}

// Add new tracking number
func (p *portType) SalesOrderShipmentAddTrack(sessionId string, shipmentIncrementId string, carrier string, title string, trackNumber string) (int, error) {
	α := struct {
		M OperationSalesOrderShipmentAddTrackRequest `xml:"typens:salesOrderShipmentAddTrack"`
	}{
		OperationSalesOrderShipmentAddTrackRequest{
			&sessionId,
			&shipmentIncrementId,
			&carrier,
			&title,
			&trackNumber,
		},
	}

	γ := struct {
		M OperationSalesOrderShipmentAddTrackResponse `xml:"salesOrderShipmentAddTrackResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("urn:Action", α, &γ); err != nil {
		return 0, err
	}
	return *γ.M.Result, nil
}

// Create new shipment for order
func (p *portType) SalesOrderShipmentCreate(sessionId string, orderIncrementId string, itemsQty *OrderItemIdQtyArray, comment string, email int, includeComment int) (string, error) {
	α := struct {
		M OperationSalesOrderShipmentCreateRequest `xml:"typens:salesOrderShipmentCreate"`
	}{
		OperationSalesOrderShipmentCreateRequest{
			&sessionId,
			&orderIncrementId,
			itemsQty,
			&comment,
			&email,
			&includeComment,
		},
	}

	γ := struct {
		M OperationSalesOrderShipmentCreateResponse `xml:"salesOrderShipmentCreateResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("urn:Action", α, &γ); err != nil {
		return "", err
	}
	return *γ.M.ShipmentIncrementId, nil
}

// Retrieve list of allowed carriers for order
func (p *portType) SalesOrderShipmentGetCarriers(sessionId string, orderIncrementId string) (*AssociativeArray, error) {
	α := struct {
		M OperationSalesOrderShipmentGetCarriersRequest `xml:"typens:salesOrderShipmentGetCarriers"`
	}{
		OperationSalesOrderShipmentGetCarriersRequest{
			&sessionId,
			&orderIncrementId,
		},
	}

	γ := struct {
		M OperationSalesOrderShipmentGetCarriersResponse `xml:"salesOrderShipmentGetCarriersResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("urn:Action", α, &γ); err != nil {
		return nil, err
	}
	return γ.M.Result, nil
}

// Retrieve shipment information
func (p *portType) SalesOrderShipmentInfo(sessionId string, shipmentIncrementId string) (*SalesOrderShipmentEntity, error) {
	α := struct {
		M OperationSalesOrderShipmentInfoRequest `xml:"typens:salesOrderShipmentInfo"`
	}{
		OperationSalesOrderShipmentInfoRequest{
			&sessionId,
			&shipmentIncrementId,
		},
	}

	γ := struct {
		M OperationSalesOrderShipmentInfoResponse `xml:"salesOrderShipmentInfoResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("urn:Action", α, &γ); err != nil {
		return nil, err
	}
	return γ.M.Result, nil
}

// Retrieve list of shipments by filters
func (p *portType) SalesOrderShipmentList(sessionId string, filters *Filters) (*SalesOrderShipmentEntityArray, error) {
	α := struct {
		M OperationSalesOrderShipmentListRequest `xml:"typens:salesOrderShipmentList"`
	}{
		OperationSalesOrderShipmentListRequest{
			&sessionId,
			filters,
		},
	}

	γ := struct {
		M OperationSalesOrderShipmentListResponse `xml:"salesOrderShipmentListResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("urn:Action", α, &γ); err != nil {
		return nil, err
	}
	return γ.M.Result, nil
}

// Remove tracking number
func (p *portType) SalesOrderShipmentRemoveTrack(sessionId string, shipmentIncrementId string, trackId string) (bool, error) {
	α := struct {
		M OperationSalesOrderShipmentRemoveTrackRequest `xml:"typens:salesOrderShipmentRemoveTrack"`
	}{
		OperationSalesOrderShipmentRemoveTrackRequest{
			&sessionId,
			&shipmentIncrementId,
			&trackId,
		},
	}

	γ := struct {
		M OperationSalesOrderShipmentRemoveTrackResponse `xml:"salesOrderShipmentRemoveTrackResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("urn:Action", α, &γ); err != nil {
		return false, err
	}
	return *γ.M.Result, nil
}

// Send shipment info
func (p *portType) SalesOrderShipmentSendInfo(sessionId string, shipmentIncrementId string, comment string) (bool, error) {
	α := struct {
		M OperationSalesOrderShipmentSendInfoRequest `xml:"typens:salesOrderShipmentSendInfo"`
	}{
		OperationSalesOrderShipmentSendInfoRequest{
			&sessionId,
			&shipmentIncrementId,
			&comment,
		},
	}

	γ := struct {
		M OperationSalesOrderShipmentSendInfoResponse `xml:"salesOrderShipmentSendInfoResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("urn:Action", α, &γ); err != nil {
		return false, err
	}
	return *γ.M.Result, nil
}

// Unhold order
func (p *portType) SalesOrderUnhold(sessionId string, orderIncrementId string) (bool, error) {
	α := struct {
		M OperationSalesOrderUnholdRequest `xml:"typens:salesOrderUnhold"`
	}{
		OperationSalesOrderUnholdRequest{
			&sessionId,
			&orderIncrementId,
		},
	}

	γ := struct {
		M OperationSalesOrderUnholdResponse `xml:"salesOrderUnholdResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("urn:Action", α, &γ); err != nil {
		return false, err
	}
	return *γ.M.Result, nil
}

// Add coupon code for shopping cart
func (p *portType) ShoppingCartCouponAdd(sessionId string, quoteId int, couponCode string, storeId string) (bool, error) {
	α := struct {
		M OperationShoppingCartCouponAddRequest `xml:"typens:shoppingCartCouponAdd"`
	}{
		OperationShoppingCartCouponAddRequest{
			&sessionId,
			&quoteId,
			&couponCode,
			&storeId,
		},
	}

	γ := struct {
		M OperationShoppingCartCouponAddResponse `xml:"shoppingCartCouponAddResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("urn:Action", α, &γ); err != nil {
		return false, err
	}
	return *γ.M.Result, nil
}

// Remove coupon code from shopping cart
func (p *portType) ShoppingCartCouponRemove(sessionId string, quoteId int, storeId string) (bool, error) {
	α := struct {
		M OperationShoppingCartCouponRemoveRequest `xml:"typens:shoppingCartCouponRemove"`
	}{
		OperationShoppingCartCouponRemoveRequest{
			&sessionId,
			&quoteId,
			&storeId,
		},
	}

	γ := struct {
		M OperationShoppingCartCouponRemoveResponse `xml:"shoppingCartCouponRemoveResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("urn:Action", α, &γ); err != nil {
		return false, err
	}
	return *γ.M.Result, nil
}

// Create shopping cart
func (p *portType) ShoppingCartCreate(sessionId string, storeId string) (int, error) {
	α := struct {
		M OperationShoppingCartCreateRequest `xml:"typens:shoppingCartCreate"`
	}{
		OperationShoppingCartCreateRequest{
			&sessionId,
			&storeId,
		},
	}

	γ := struct {
		M OperationShoppingCartCreateResponse `xml:"shoppingCartCreateResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("urn:Action", α, &γ); err != nil {
		return 0, err
	}
	return *γ.M.QuoteId, nil
}

// Set customer's addresses in shopping cart
func (p *portType) ShoppingCartCustomerAddresses(sessionId string, quoteId int, customer *ShoppingCartCustomerAddressEntityArray, storeId string) (bool, error) {
	α := struct {
		M OperationShoppingCartCustomerAddressesRequest `xml:"typens:shoppingCartCustomerAddresses"`
	}{
		OperationShoppingCartCustomerAddressesRequest{
			&sessionId,
			&quoteId,
			customer,
			&storeId,
		},
	}

	γ := struct {
		M OperationShoppingCartCustomerAddressesResponse `xml:"shoppingCartCustomerAddressesResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("urn:Action", α, &γ); err != nil {
		return false, err
	}
	return *γ.M.Result, nil
}

// Set customer for shopping cart
func (p *portType) ShoppingCartCustomerSet(sessionId string, quoteId int, customer *ShoppingCartCustomerEntity, storeId string) (bool, error) {
	α := struct {
		M OperationShoppingCartCustomerSetRequest `xml:"typens:shoppingCartCustomerSet"`
	}{
		OperationShoppingCartCustomerSetRequest{
			&sessionId,
			&quoteId,
			customer,
			&storeId,
		},
	}

	γ := struct {
		M OperationShoppingCartCustomerSetResponse `xml:"shoppingCartCustomerSetResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("urn:Action", α, &γ); err != nil {
		return false, err
	}
	return *γ.M.Result, nil
}

// Retrieve information about shopping cart
func (p *portType) ShoppingCartInfo(sessionId string, quoteId int, storeId string) (*ShoppingCartInfoEntity, error) {
	α := struct {
		M OperationShoppingCartInfoRequest `xml:"typens:shoppingCartInfo"`
	}{
		OperationShoppingCartInfoRequest{
			&sessionId,
			&quoteId,
			&storeId,
		},
	}

	γ := struct {
		M OperationShoppingCartInfoResponse `xml:"shoppingCartInfoResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("urn:Action", α, &γ); err != nil {
		return nil, err
	}
	return γ.M.Result, nil
}

// Get terms and conditions
func (p *portType) ShoppingCartLicense(sessionId string, quoteId int, storeId string) (*ShoppingCartLicenseEntityArray, error) {
	α := struct {
		M OperationShoppingCartLicenseRequest `xml:"typens:shoppingCartLicense"`
	}{
		OperationShoppingCartLicenseRequest{
			&sessionId,
			&quoteId,
			&storeId,
		},
	}

	γ := struct {
		M OperationShoppingCartLicenseResponse `xml:"shoppingCartLicenseResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("urn:Action", α, &γ); err != nil {
		return nil, err
	}
	return γ.M.Result, nil
}

// Create an order from shopping cart
func (p *portType) ShoppingCartOrder(sessionId string, quoteId int, storeId string, licenses *ArrayOfString) (string, error) {
	α := struct {
		M OperationShoppingCartOrderRequest `xml:"typens:shoppingCartOrder"`
	}{
		OperationShoppingCartOrderRequest{
			&sessionId,
			&quoteId,
			&storeId,
			licenses,
		},
	}

	γ := struct {
		M OperationShoppingCartOrderResponse `xml:"shoppingCartOrderResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("urn:Action", α, &γ); err != nil {
		return "", err
	}
	return *γ.M.Result, nil
}

// Get list of available payment methods
func (p *portType) ShoppingCartPaymentList(sessionId string, quoteId int, store string) (*ShoppingCartPaymentMethodResponseEntityArray, error) {
	α := struct {
		M OperationShoppingCartPaymentListRequest `xml:"typens:shoppingCartPaymentList"`
	}{
		OperationShoppingCartPaymentListRequest{
			&sessionId,
			&quoteId,
			&store,
		},
	}

	γ := struct {
		M OperationShoppingCartPaymentListResponse `xml:"shoppingCartPaymentListResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("urn:Action", α, &γ); err != nil {
		return nil, err
	}
	return γ.M.Result, nil
}

// Set payment method
func (p *portType) ShoppingCartPaymentMethod(sessionId string, quoteId int, method *ShoppingCartPaymentMethodEntity, storeId string) (bool, error) {
	α := struct {
		M OperationShoppingCartPaymentMethodRequest `xml:"typens:shoppingCartPaymentMethod"`
	}{
		OperationShoppingCartPaymentMethodRequest{
			&sessionId,
			&quoteId,
			method,
			&storeId,
		},
	}

	γ := struct {
		M OperationShoppingCartPaymentMethodResponse `xml:"shoppingCartPaymentMethodResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("urn:Action", α, &γ); err != nil {
		return false, err
	}
	return *γ.M.Result, nil
}

// Add product(s) to shopping cart
func (p *portType) ShoppingCartProductAdd(sessionId string, quoteId int, products *ShoppingCartProductEntityArray, storeId string) (bool, error) {
	α := struct {
		M OperationShoppingCartProductAddRequest `xml:"typens:shoppingCartProductAdd"`
	}{
		OperationShoppingCartProductAddRequest{
			&sessionId,
			&quoteId,
			products,
			&storeId,
		},
	}

	γ := struct {
		M OperationShoppingCartProductAddResponse `xml:"shoppingCartProductAddResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("urn:Action", α, &γ); err != nil {
		return false, err
	}
	return *γ.M.Result, nil
}

// Get list of products in shopping cart
func (p *portType) ShoppingCartProductList(sessionId string, quoteId int, storeId string) (*ShoppingCartProductResponseEntityArray, error) {
	α := struct {
		M OperationShoppingCartProductListRequest `xml:"typens:shoppingCartProductList"`
	}{
		OperationShoppingCartProductListRequest{
			&sessionId,
			&quoteId,
			&storeId,
		},
	}

	γ := struct {
		M OperationShoppingCartProductListResponse `xml:"shoppingCartProductListResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("urn:Action", α, &γ); err != nil {
		return nil, err
	}
	return γ.M.Result, nil
}

// Move product(s) to customer quote
func (p *portType) ShoppingCartProductMoveToCustomerQuote(sessionId string, quoteId int, products *ShoppingCartProductEntityArray, storeId string) (bool, error) {
	α := struct {
		M OperationShoppingCartProductMoveToCustomerQuoteRequest `xml:"typens:shoppingCartProductMoveToCustomerQuote"`
	}{
		OperationShoppingCartProductMoveToCustomerQuoteRequest{
			&sessionId,
			&quoteId,
			products,
			&storeId,
		},
	}

	γ := struct {
		M OperationShoppingCartProductMoveToCustomerQuoteResponse `xml:"shoppingCartProductMoveToCustomerQuoteResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("urn:Action", α, &γ); err != nil {
		return false, err
	}
	return *γ.M.Result, nil
}

// Remove product(s) from shopping cart
func (p *portType) ShoppingCartProductRemove(sessionId string, quoteId int, products *ShoppingCartProductEntityArray, storeId string) (bool, error) {
	α := struct {
		M OperationShoppingCartProductRemoveRequest `xml:"typens:shoppingCartProductRemove"`
	}{
		OperationShoppingCartProductRemoveRequest{
			&sessionId,
			&quoteId,
			products,
			&storeId,
		},
	}

	γ := struct {
		M OperationShoppingCartProductRemoveResponse `xml:"shoppingCartProductRemoveResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("urn:Action", α, &γ); err != nil {
		return false, err
	}
	return *γ.M.Result, nil
}

// Update product(s) quantities in shopping cart
func (p *portType) ShoppingCartProductUpdate(sessionId string, quoteId int, products *ShoppingCartProductEntityArray, storeId string) (bool, error) {
	α := struct {
		M OperationShoppingCartProductUpdateRequest `xml:"typens:shoppingCartProductUpdate"`
	}{
		OperationShoppingCartProductUpdateRequest{
			&sessionId,
			&quoteId,
			products,
			&storeId,
		},
	}

	γ := struct {
		M OperationShoppingCartProductUpdateResponse `xml:"shoppingCartProductUpdateResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("urn:Action", α, &γ); err != nil {
		return false, err
	}
	return *γ.M.Result, nil
}

// Get list of available shipping methods
func (p *portType) ShoppingCartShippingList(sessionId string, quoteId int, storeId string) (*ShoppingCartShippingMethodEntityArray, error) {
	α := struct {
		M OperationShoppingCartShippingListRequest `xml:"typens:shoppingCartShippingList"`
	}{
		OperationShoppingCartShippingListRequest{
			&sessionId,
			&quoteId,
			&storeId,
		},
	}

	γ := struct {
		M OperationShoppingCartShippingListResponse `xml:"shoppingCartShippingListResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("urn:Action", α, &γ); err != nil {
		return nil, err
	}
	return γ.M.Result, nil
}

// Set shipping method
func (p *portType) ShoppingCartShippingMethod(sessionId string, quoteId int, method string, storeId string) (bool, error) {
	α := struct {
		M OperationShoppingCartShippingMethodRequest `xml:"typens:shoppingCartShippingMethod"`
	}{
		OperationShoppingCartShippingMethodRequest{
			&sessionId,
			&quoteId,
			&method,
			&storeId,
		},
	}

	γ := struct {
		M OperationShoppingCartShippingMethodResponse `xml:"shoppingCartShippingMethodResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("urn:Action", α, &γ); err != nil {
		return false, err
	}
	return *γ.M.Result, nil
}

// Get total prices for shopping cart
func (p *portType) ShoppingCartTotals(sessionId string, quoteId int, storeId string) (*ShoppingCartTotalsEntityArray, error) {
	α := struct {
		M OperationShoppingCartTotalsRequest `xml:"typens:shoppingCartTotals"`
	}{
		OperationShoppingCartTotalsRequest{
			&sessionId,
			&quoteId,
			&storeId,
		},
	}

	γ := struct {
		M OperationShoppingCartTotalsResponse `xml:"shoppingCartTotalsResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("urn:Action", α, &γ); err != nil {
		return nil, err
	}
	return γ.M.Result, nil
}

// Start web service session
func (p *portType) StartSession() (string, error) {
	α := struct {
		M struct{} `xml:"typens:startSession"`
	}{
		struct{}{},
	}

	γ := struct {
		M OperationStartSessionResponse `xml:"startSessionResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("urn:Action", α, &γ); err != nil {
		return "", err
	}
	return *γ.M.StartSessionReturn, nil
}

// Store view info
func (p *portType) StoreInfo(sessionId string, storeId string) (*StoreEntity, error) {
	α := struct {
		M OperationStoreInfoRequest `xml:"typens:storeInfo"`
	}{
		OperationStoreInfoRequest{
			&sessionId,
			&storeId,
		},
	}

	γ := struct {
		M OperationStoreInfoResponse `xml:"storeInfoResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("urn:Action", α, &γ); err != nil {
		return nil, err
	}
	return γ.M.Info, nil
}

// List of stores
func (p *portType) StoreList(sessionId string) (*StoreEntityArray, error) {
	α := struct {
		M OperationStoreListRequest `xml:"typens:storeList"`
	}{
		OperationStoreListRequest{
			&sessionId,
		},
	}

	γ := struct {
		M OperationStoreListResponse `xml:"storeListResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("urn:Action", α, &γ); err != nil {
		return nil, err
	}
	return γ.M.Stores, nil
}
