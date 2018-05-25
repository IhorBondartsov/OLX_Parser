package olx_client

import "testing"

func TestGetHTMLPage(t *testing.T) {
	GetHTMLPage("https://www.olx.ua/zapchasti-dlya-transporta/kharkov/?search%5Bdistrict_id%5D=79")
}