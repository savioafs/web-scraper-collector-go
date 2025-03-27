package main

import (
	"fmt"
	"github.com/savioafs/web-scraper-collector-go/internal/services/collectors"
	"log"
)

var url = "https://www.mercadolivre.com.br/alto-falante-jbl-selenium-ambientare-c621-branco-50w-rms-de-parede/p/MLB15789972?pdp_filters=item_id:MLB2076197876#polycard_client=recommendations_pdp-pads-reviews&reco_backend=recos-pads-retrieval-4stars&reco_model=rk_ent_v2_retsys_4stars&reco_client=pdp-pads-reviews&reco_item_pos=0&reco_backend_type=low_level&reco_id=f29c022a-f1c6-422f-976f-74f386a819f1&wid=MLB2076197876&sid=recos&is_advertising=true&ad_domain=PDPDESKTOP_RECOMMENDED&ad_position=1&ad_click_id=NTM4MTk4ZTctMTdkMy00MzM4LTgzZjMtYzNkYjdjMjRhYTlj"

func main() {
	name, price, err := collectors.Run(url)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("name:::::::", name)
	fmt.Println("price:::::::", price)

}
