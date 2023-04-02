package handler

import "github.com/michael1011/clnurl/clnurl"

type getClnurl = func() (*clnurl.ClnUrl, error)
