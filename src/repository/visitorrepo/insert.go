package visitorrepo

import "pandudpn/api/src/model"

const createNewVisitor = `
INSERT INTO visitors
(ip, user_agent, created_at) VALUES ($1, $2, $3)
returning id
`

func (v *VisitorRepository) NewVisitor(visitor *model.Visitor) error {
	stmt, err := v.db.Preparex(createNewVisitor)
	if err != nil {
		return err
	}
	defer stmt.Close()

	err = stmt.QueryRowx(visitor.Ip, visitor.UserAgent, visitor.CreatedAt).StructScan(visitor)
	return err
}
