package metrics

import (
	"fmt"
	"net"
	"time"
)

func WriteMetricWithPlaintext(graphiteConn net.Conn, name string, value float64) {
	if _, err := fmt.Fprintf(graphiteConn, "%s %f %d\n", name, value, time.Now().Unix()); err != nil {
		fmt.Println("error while wrapping metrics to Graphite:", err.Error())
	}
}
