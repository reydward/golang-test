package main

type Order struct {
	ID                  int             `json:"id"`
	Packages            []*Package      `json:"packages"`
	SalesChannelOrderID int             `json:"salesChannelOrderId"`
}

type Package struct {
	ID                    int           `json:"id"`
	OrderID               int           `json:"orderId"`
	VendorID              int           `json:"vendorId"`
	PackageVendorID       *string       `json:"packageVendorId"`
}

func main()  {
	println("test")

	order := Order{
		ID: 1,
		SalesChannelOrderID: 11,
	}

	o, err := applyRoutingRules(&order)

	if err != nil {
		o.SalesChannelOrderID = 666
		return
	}

	if len(order.Packages) == 0 {
		println("Order has no packages after routing")
	}

	println("Order with packages after routing")
}

func initializePackage() (p Package) {
	p.OrderID = 333
	p.VendorID = 111
	return
}

func applyRoutingRules(input *Order) (order *Order, err error) {
	newP := initializePackage()
	input.Packages = append(input.Packages, &newP)
	return
}
