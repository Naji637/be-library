package borrowing

import "errors"

const (
	StatusBorrowed = "dipinjam"
	StatusReturned = "dikembalikan"
)

var (
	ErrorNoBookRequested         = errors.New("tidak ada buku yang dipinjam")
	ErrorEmptyBookHubId          = errors.New("book hub id tidak boleh kosong")
	ErrorDuplicateBookHubId      = errors.New("book hub id duplikat dalam satu request")
	ErrorMemberNotFound          = errors.New("member tidak ditemukan")
	ErrorMaxBooksExceeded        = errors.New("jumlah buku melebihi batas pinjaman member")
	ErrorBookHubNotFound         = errors.New("buku tidak ditemukan")
	ErrorBookAlreadyBorrowed     = errors.New("buku sedang dipinjam")
	ErrorBookAlreadyHeldByMember = errors.New("member meminjam buku dengan id yang sama")
)

var borrowingRuleErrors = []error{
	ErrorNoBookRequested,
	ErrorEmptyBookHubId,
	ErrorDuplicateBookHubId,
	ErrorMemberNotFound,
	ErrorMaxBooksExceeded,
	ErrorBookHubNotFound,
	ErrorBookAlreadyBorrowed,
	ErrorBookAlreadyHeldByMember,
}

func IsRuleError(err error) bool {
	for _, ruleError := range borrowingRuleErrors {
		if errors.Is(err, ruleError) {
			return true
		}
	}
	return false
}
