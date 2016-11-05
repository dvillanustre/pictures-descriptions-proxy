package main

import (
	"bytes"

	"pictures-descriptions-proxy/src/proxy/scaler"

	"net/http"

	"github.com/gin-gonic/gin"
)

func getDescription(c *gin.Context) {
	html := "<table width='906' align='center' border='0'><tbody><tr><td width='1078'> <div align='center'> <p><img width='900' height='156' src='http://www.allmobile.mx/ml/all3/1.jpg' border='0' usemap='#Map' /><img width='900' height='431' src='http://www.allmobile.mx/ml/all3/fleipo4gh.jpg' /><img width='900' height='211' src='http://www.allmobile.mx/ml/all3/3.jpg' /></p> <p> </p> <p></p> <p><img width='900' height='134' src='http://www.allmobile.mx/ml/all3/5.jpg' /></p> <p> </p> <p><img width='900' height='1372' align='absmiddle' src='http://www.allmobile.mx/ml/all3/4.jpg' /></p> </div> </td></tr></tbody></table> <p><map name='Map'> <area href='http://eshops.mercadolibre.com.mx/allmobile&#43;mx/' shape='rect' coords='825,4,881,21' target='_new' /> <area href='http://eshops.mercadolibre.com.mx/allmobile&#43;mx/politicas&#43;de&#43;venta/' shape='rect' coords='755,4,824,22' target='_new' /> <area href='http://listado.mercadolibre.com.mx/_DisplayType_G_CustId_25871251' shape='rect' coords='676,4,754,22' target='_new' /> <area href='http://www.mercadolibre.com.mx/jm/profile?act&#61;ver&id&#61;25871251&oper&#61;B&isMot&#61;N&tipo&#61;1&RP2&#61;Y&nick&#61;' shape='rect' coords='577,4,673,22' target='_new' /> </map></p>"

	buffer := bytes.NewBufferString(html)

	byteArr, err := scaler.GetPictureDefault(buffer.Bytes(), "D")

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "error"})
		return
	}

	c.Data(http.StatusOK, "image/jpeg", byteArr)
}

func main() {
	initRouter()
}

func initRouter() {
	r := gin.Default()

	r.GET("/description", func(c *gin.Context) { getDescription(c) })

	r.Run(":8080")
}
