package main

//sk_test_51JMja8H47rNLCGON29JCa7YfnwwWFnDradk4edkDkVIrbDHlvcgxT5RpaexLFyLMTJA86iMsSNYnH5xZlN0s6fkr00R5DnpqgY

import(
	"fmt"
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/customer"
)

func main(){
	stripe.Key = "sk_test_51JMja8H47rNLCGON29JCa7YfnwwWFnDradk4edkDkVIrbDHlvcgxT5RpaexLFyLMTJA86iMsSNYnH5xZlN0s6fkr00R5DnpqgY"
	
	fmt.Println("making requests!")
	
	//c,_  := customer.New(nil) //instance of the object created, error
	//fmt.Println(c)
	
	//cus_K13WlSHISjkImI 
	//c, _ := customer.Get("cus_K13WlSHISjkImI", nil)
	//fmt.Println(c)
	
	//params := &stripe.CustomerParams{
	//	Email: stripe.String("amp3453@example.com"),
	//	Name: stripe.String("AMP"),
	//}
	
	//stripe API request is form encoded, response is JSON
	
	//c, _ := customer.New(params)
	//fmt.Println(c)
	
	//params := &stripe.CustomerParams{
		//TaxExempt: stripe.String(string(stripe.CustomerTaxExemptNone)),
	//}
	
	//c,_ := customer.New(params)
	//fmt.Println(c.TaxExempt)
	
	//creating an object with a nested object
	params := &stripe.CustomerParams{
		PaymentMethod: stripe.String("pm_card_visa"),
		InvoiceSettings: &stripe.CustomerInvoiceSettingsParams{
			DefaultPaymentMethod: stripe.String("pm_card_visa"),
		},
	}
	
	c, _ := customer.New(params)
	fmt.Println(c.InvoiceSettings.DefaultPaymentMethod)
	
	
}