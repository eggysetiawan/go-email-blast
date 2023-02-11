package domain

type Csv struct {
	TrxTime string `csv:"trx_time"`
	Tid     string `csv:"terminal_id"`
}
