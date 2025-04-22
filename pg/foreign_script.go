package pg

import "fmt"

// Foreign is used to avoid "import cycle not allowed" errors in Go.
// It contains foreign key details, but the main purpose is to break circular dependencies.
type Foreign struct {
	Table       string // Table containing the foreign key
	ForeignID   string // Column in the table acting as the foreign key
	Reference   string // Referenced table
	ReferenceID string // Referenced column
}

// GetScript generates a SQL script for adding a foreign key constraint, but its primary purpose
// is to break import cycles in Go by creating an indirect dependency between packages.
func (f *Foreign) GetScript() string {

	fkname := fmt.Sprintf("fk_%s_%s", f.Table, f.ForeignID)

	return `
	DO $$ BEGIN
		IF NOT EXISTS (SELECT 1 FROM pg_constraint WHERE conname= '` + fkname + `') THEN
			ALTER TABLE public.` + f.Table + ` ADD CONSTRAINT ` + fkname + `
			FOREIGN KEY (` + f.ForeignID + `) REFERENCES ` + f.Reference + `(` + f.ReferenceID + `) 
			ON UPDATE CASCADE ON DELETE CASCADE;
		END IF;
	END $$;
	`
}
