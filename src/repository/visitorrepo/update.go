package visitorrepo

import (
	"fmt"

	"pandudpn/api/src/model"
)

const queryUpdateVisitor = `
UPDATE visitors SET total_visit=$2, updated_at=$3
WHERE id=$1
`

func (v *VisitorRepository) UpdateVisitor(visitor *model.Visitor) error {
	stmt, err := v.db.Preparex(queryUpdateVisitor)
	if err != nil {
		return err
	}
	defer stmt.Close()

	res, err := stmt.Exec(visitor.Id, visitor.TotalVisit, visitor.UpdatedAt)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected < 1 {
		err = fmt.Errorf("error update visitor")
		return err
	}

	return nil
}
