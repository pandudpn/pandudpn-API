package visitorrepo

import "pandudpn/api/src/model"

const queryVisitorByIpAndUserAgent = `
SELECT id, ip, user_agent, total_visit, created_at
FROM visitors
WHERE ip=$1 AND user_agent=$2
AND DATE(created_at) = DATE(now())
ORDER BY created_at DESC
`

func (v *VisitorRepository) FindVisitorByIpAndUserAgent(ip, userAgent string) (*model.Visitor, error) {
	var visitor model.Visitor

	err := v.db.Get(&visitor, queryVisitorByIpAndUserAgent, ip, userAgent)
	if err != nil {
		return nil, err
	}

	return &visitor, nil
}
