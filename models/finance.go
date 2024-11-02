package models

import "time"

type Income struct {
	ID          int       `bun:",pk,autoincrement" json:"id"`
	UserID      int       `bun:"user_id,notnull" json:"user_id"`
	User        *User     `bun:"rel:belongs-to" json:"user"`
	Amount      float64   `bun:"amount,notnull" json:"amount"`            // Gelir miktarı
	Description string    `bun:"description,nullzero" json:"description"` // Gelir açıklaması (örn. maaş, ek gelir vb.)
	Date        time.Time `bun:"date,notnull" json:"date"`                // Gelir tarihi
}

type Expense struct {
	ID          int       `bun:",pk,autoincrement" json:"id"`
	UserID      int       `bun:"user_id,notnull" json:"user_id"`
	User        *User     `bun:"rel:belongs-to" json:"user"`
	Amount      float64   `bun:"amount,notnull" json:"amount"`            // Gider miktarı
	Description string    `bun:"description,nullzero" json:"description"` // Gider açıklaması (örn. kira, fatura vb.)
	Category    string    `bun:"category,nullzero" json:"category"`       // Gider kategorisi (örn. ulaşım, yiyecek)
	Date        time.Time `bun:"date,notnull" json:"date"`                // Gider tarihi
	BudgetID    int       `bun:"budget_id,nullzero" json:"budget_id"`     // Giderin ait olduğu bütçe
}

type Budget struct {
	ID        int       `bun:",pk,autoincrement" json:"id"`
	UserID    int       `bun:"user_id,notnull" json:"user_id"`
	User      *User     `bun:"rel:belongs-to" json:"user"`
	Amount    float64   `bun:"amount,notnull" json:"amount"`         // Bütçe miktarı
	Category  string    `bun:"category,nullzero" json:"category"`    // Bütçenin kategorisi (örn. ulaşım, yiyecek)
	StartDate time.Time `bun:"start_date,notnull" json:"start_date"` // Bütçe başlangıç tarihi
	EndDate   time.Time `bun:"end_date,notnull" json:"end_date"`     // Bütçe bitiş tarihi
	//Expenses  []Expense `bun:"rel:has-many,join:budget_id=id" json:"expenses"` // Bütçeye dahil giderler
}

type Report struct {
	ID           int       `bun:",pk,autoincrement" json:"id"`
	UserID       int       `bun:"user_id,notnull" json:"user_id"`
	User         *User     `bun:"rel:belongs-to" json:"user"`
	Month        int       `bun:"month,notnull" json:"month"`                 // Raporun ait olduğu ay (1-12)
	Year         int       `bun:"year,notnull" json:"year"`                   // Raporun ait olduğu yıl
	TotalIncome  float64   `bun:"total_income,notnull" json:"total_income"`   // Toplam gelir
	TotalExpense float64   `bun:"total_expense,notnull" json:"total_expense"` // Toplam gider
	Balance      float64   `bun:"balance,notnull" json:"balance"`             // Gelir-gider dengesi (gelir - gider)
	GeneratedAt  time.Time `bun:"generated_at,notnull" json:"generated_at"`   // Rapor oluşturulma tarihi
}

type GenerateReport struct {
	UserID int `json:"user_id"`
	Month  int `json:"month"`
	Year   int `json:"year"`
}
