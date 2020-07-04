package payconfig

import (
	"github.com/iGoogle-ink/gopay"
	"github.com/iGoogle-ink/gopay/alipay"
)

var (
	Client    *alipay.Client
	BodyMap   = make(gopay.BodyMap)
	AppId     = "2016101600703675"
	gateway   = "https://openapi.alipaydev.com/gateway.do"
	ReturnUrl = "http://hizjlhi.com:9000/CompleteOrder"
	NotifiyUrl = "http://hizjlhi.com:9000/Verity" // https://openapi.alipay.com/gateway.do
	key = "i6zI0XxjkSzzy"

	privateKey = "MIIEogIBAAKCAQEAwtP+z4NnWykjfA2U5jVJ8ySyyJhw4wgmTcUFbHTLpWT3FurmAr" +
		"WSfhHgmXSyQl3S1IyvQxGgDH6BwiAHU3PqtzyXMIZl35G6J2sWAuhubVSj4ILN8g8vIk6HdPEcnffJ" +
		"lXLwjD+frm2xvCEAPDlMqh1y6vYh+YrBGG1yZK+tpK+7BdYahgqT6TTT26VadIHdCoR0wpNjwfnFTgQkD" +
		"Qy0cLPzuHXQpZyMLlXJonRb74LNOEcyf9ULkMYutAp9QOUyyWXU7JGPIB6K4a3Qkyk9mdrKjPmsAQzXLDhF/" +
		"zl4QC34izlxJaaF4SyL4B/7B4VeITYV2wigUhG/ceTbBYEd5wIDAQABAoIBABIdwF3vbBbv9lq8sJHPdrbjPLSma" +
		"CjUQyuMHHr/SUhx4QACi6zI0XxjkSzzyvec3dPh2RI1e1puEQbKD6RU7Qho2+4pMSPe2x57OKrAjjQgYLSptRjDDLD1" +
		"+GaJXQ3bUqVQ8Mk5yVFg5dwGfY0cDuLqphrvFC2uF1J8Ktzt8QmVHBg+6qDyb35yhXzeSKQVLbcslOrS6XyeokZRIYbbnYl" +
		"fsPVMqVVE9OEmzTJaDY92q2aN8XUO1IenszaNbtSwpqvt9EIfCagNNhkIDpUwntETVo7Vpx90mDbRnrjk7cEewmvc77IHXhhEIj" +
		"Jms0NVOiCpYlniBPcya3jlPuhK52ECgYEA4ERa2I+U+clPbGs2Lq8cu+tYN/AHKXSUhH96E1rzxl/gL7ygMSzg2cxNTGv9AWQsntv+L" +
		"u7rEVbzkZcVd64LqwIlhrdmcMQjYsRYbJ8Ve3xtUGJwavW2NldgcqCCeV+OWzF+cwJOWz3xwp8php3Db/TZnBhHT8duqiMNVOALPZkCgYEA" +
		"3mVFs7pCKkLDNjsX0FhZNRpfEGcDWdQR3uMQf8RH9oKI5HTrRdQMLcyqCg+ZDFg+nrlR1dGYVIYZvV10gGbDNWrOXtcDztVZauXgM+1jqIClMigB" +
		"vIC426sCcQxkjDEsnVMJOy2MwYwzCYEdM2HsYbOQsRMk1tVwJAordyBtZ38CgYA8y4rtBg9Wn3H+bBnmEeZyMtxZXaIzZL0WJhCLyq4m6rq1dibe7dGO" +
		"eUnDG8scY3GNJwoC4xWqP64Lm69gIDhdhVly6ajFjQYisiNeXsnODe78SZM8C5v3ozwsFsMH8BWBNyuWNdvT9Djjj44MRhxC71dGb3Z1dBTV3mawyYOCkQKB" +
		"gCN5mBLBRUikeasu5e+QCDrWSA+/vuMXVvps0fbvgmMqE1gN5nGvD3pGyWDteAZmFScQfKNU2a2x7m0seSb6PW1J6j1qourdUXQh0w+1cE4ypydHBm/hQJuZvbv2" +
		"tBtAxNMbqb9M1sUQ7hs4A0Zs+l3jXNNCMIAsb8Tv8lsASzNVAoGAXXLcBwC/hfYhFOIFE6uULqR/O0QtrusMTjzzOtdyBsuJXP5YCZca7iqJNrXLCkiTFjWfr+Dz4bjh" +
		"l6ApY+WfQArxtc9fz3muLu8cxReSi3nK9UMbnZddfjvrutUxqYB3yM7SjIFVKV/OhIzGgODUWTSWX9gYS4Mv4ncbN5PBYVY="

	Aipaykey = "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAsHdemCWwBjEeb2oeEcYUt4+OYc0PQs" +
		"IG9Gjf+JJ0ZFtBs2NZjmTZbABPk56WbYpDyuVf1C1H56+ySJTHi0XEnuEvgkRSBzu1+NDmv8tAeBag3ki3J2" +
		"f1bw0Y79lFMLhdzWRCx8Tn5Jd2F3NJRH4+nDXRLXIZsQ9e3cEL9IJ4YZCqgVNrjqzl5HhfLOHLuDcnJjMvQTtidh" +
		"eA9oC5zU36x26VpbmHHwXTdcubgBtjbe2Mf59BpdzQTjWkkVmGhF/lo/4mv+9tn/AiHci9iOraUqcMQXfLQGq6V01Vg8" +
		"/ha5n+SLA4MEVnrhD1FB93rSvEBG5nzl4D/E1I6U1HyJNEMQIDAQAB"
)